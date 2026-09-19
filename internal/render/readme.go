// Package render turns the list and its metadata into README.md.
package render

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/floatdrop/awesome-go/internal/list"
)

// StarsDir is the directory under badges/ that holds one star pill per entry.
const StarsDir = "stars"

// StarsPath is the README-relative path of an entry's star pill.
func StarsPath(repo string) string {
	return "badges/" + StarsDir + "/" + list.Slug(repo) + ".svg"
}

// StarsText is what the star pill shows: the rounded count, or "new" for an
// entry whose stars have not been fetched yet.
func StarsText(stars int, known bool) string {
	if !known {
		return "new"
	}
	return FormatStars(stars)
}

// Header is the first line of the generated README; CI uses it to recognise
// the file as generated.
const Header = "<!-- Generated from list.json and entries/ by `go run ./cmd/awesome generate`. Do not edit by hand. -->"

type ranked struct {
	entry list.Entry
	stars int
	known bool
}

// README renders the full README.md.
func README(l *list.List, m *list.Metadata, now time.Time) []byte {
	var b strings.Builder
	w := func(format string, args ...any) { fmt.Fprintf(&b, format, args...) }

	groups := map[string][]ranked{}
	for _, e := range l.Entries {
		r := ranked{entry: e}
		if rm, ok := m.Repos[e.Repo]; ok {
			r.stars, r.known = rm.Stars, true
		}
		groups[e.Category] = append(groups[e.Category], r)
	}
	var categories []list.Category
	for _, c := range l.Categories {
		if len(groups[c.ID]) > 0 {
			categories = append(categories, c)
		}
	}

	w("%s\n", Header)
	if l.Meta.Logo != "" {
		w("<p align=\"center\">\n  <img src=\"%s\" alt=\"The Go gopher wearing pixel sunglasses\" width=\"180\">\n</p>\n\n", l.Meta.Logo)
		w("<h1 align=\"center\">%s</h1>\n\n", l.Meta.Title)
	} else {
		w("# %s\n\n", l.Meta.Title)
	}
	if d := l.Meta.Discord; d != nil {
		w("<p align=\"center\">\n  <a href=\"%s\"><img src=\"https://img.shields.io/discord/%s?logo=discord&logoColor=white&label=discord&color=5865F2\" alt=\"Discord\"></a>\n</p>\n\n", d.Link(), d.Server)
	}
	w("> %s\n\n", l.Meta.Description)
	w("%d projects in %d categories. ", len(l.Entries), len(categories))
	w("Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. ")
	w("An entry earns its place by covering something no listed project does, or by doing it better. ")
	if l.Meta.Podium > 0 {
		w("Within each category the %d most-starred projects are shown first and the rest are folded under More. ", l.Meta.Podium)
	} else {
		w("Entries are ordered by GitHub stars. ")
	}
	w("Archived repositories are removed automatically.")
	if m.UpdatedAt != "" {
		w(" Stars were last refreshed on %s.", m.UpdatedAt)
	}
	w("\n\n")
	w("Search, filter and sort the full list at [%s](%s).\n\n", strings.TrimSuffix(strings.TrimPrefix(l.Meta.SiteURL(), "https://"), "/"), l.Meta.SiteURL())

	var linkSections []list.LinkSection
	for _, s := range l.LinkSections {
		if len(l.LinksIn(s.ID)) > 0 {
			linkSections = append(linkSections, s)
		}
	}

	w("## Contents\n\n")
	for _, c := range categories {
		w("- [%s](#%s)\n", c.Name, Slugify(c.Name))
	}
	for _, s := range linkSections {
		w("- [%s](#%s)\n", s.Name, Slugify(s.Name))
	}

	for _, c := range categories {
		w("\n## %s\n\n", c.Name)
		if c.Description != "" {
			w("%s\n\n", c.Description)
		}
		rs := groups[c.ID]
		sort.SliceStable(rs, func(i, j int) bool {
			if rs[i].stars != rs[j].stars {
				return rs[i].stars > rs[j].stars
			}
			return strings.ToLower(rs[i].entry.DisplayName()) < strings.ToLower(rs[j].entry.DisplayName())
		})
		podium := l.Meta.Podium
		if podium <= 0 || podium >= len(rs) {
			podium = len(rs)
		}
		for _, r := range rs[:podium] {
			w("- %s\n", item(r))
		}
		if rest := rs[podium:]; len(rest) > 0 {
			w("\n<details>\n<summary>More (%d)</summary>\n\n", len(rest))
			for _, r := range rest {
				w("- %s\n", item(r))
			}
			w("\n</details>\n")
		}
	}

	for _, s := range linkSections {
		w("\n## %s\n\n", s.Name)
		if s.Description != "" {
			w("%s\n\n", s.Description)
		}
		for _, k := range l.LinksIn(s.ID) {
			w("- [%s](%s) - %s", k.Title, k.URL, k.Description)
			if len(k.Versions) > 0 {
				parts := make([]string, len(k.Versions))
				for i, v := range k.Versions {
					parts[i] = fmt.Sprintf("[%s](%s)", v.Label, v.URL)
				}
				w(" Versions: %s", strings.Join(parts, " · "))
			}
			w("\n")
		}
	}

	w("\n## Contributing\n\n")
	w("Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.\n")
	w("\n## License\n\n")
	w("[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)\n\n")
	w("To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work.")
	if l.Meta.Logo != "" {
		w(" The logo is excluded: it is based on the Go gopher by Renée French, licensed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).")
	}
	w("\n")
	return []byte(b.String())
}

// item renders one entry line. Every line starts with the entry's star pill,
// a fixed-width image, so the names line up in one column. It is an HTML img
// rather than Markdown because GitHub puts Markdown images on the text
// baseline, which leaves a 20px pill floating above the line. GitHub's
// sanitizer keeps the align attribute; "absmiddle" is the value every engine
// maps to vertical-align: middle, whereas WebKit treats "middle" as
// baseline-middle and drops the pill below the text.
func item(r ranked) string {
	alt := "new"
	if r.known {
		alt = FormatStars(r.stars) + " stars"
	}
	return fmt.Sprintf(`<img src="%s" alt="%s" align="absmiddle"> [%s](%s) - %s`, StarsPath(r.entry.Repo), alt, r.entry.DisplayName(), r.entry.URL(), r.entry.Description)
}

// FormatStars renders 1234 as "1.2k" and 80123 as "80k".
func FormatStars(n int) string {
	switch {
	case n < 1000:
		return fmt.Sprintf("%d", n)
	case n < 10000:
		tenths := (n + 50) / 100 // round half up to one decimal
		if tenths%10 == 0 {
			return fmt.Sprintf("%dk", tenths/10)
		}
		return fmt.Sprintf("%d.%dk", tenths/10, tenths%10)
	default:
		return fmt.Sprintf("%dk", (n+500)/1000)
	}
}

var slugStrip = regexp.MustCompile(`[^a-z0-9 -]`)

// Slugify mimics GitHub's heading anchors closely enough for a table of
// contents: lowercase, drop punctuation, spaces to hyphens.
func Slugify(heading string) string {
	s := strings.ToLower(heading)
	s = slugStrip.ReplaceAllString(s, "")
	return strings.ReplaceAll(s, " ", "-")
}
