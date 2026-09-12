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

var medals = []string{"🥇", "🥈", "🥉"}

// Header is the first line of the generated README; CI uses it to recognise
// the file as generated.
const Header = "<!-- Generated from awesome.json by `go run ./cmd/awesome generate`. Do not edit by hand. -->"

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
	w("# %s\n\n", l.Meta.Title)
	w("> %s\n\n", l.Meta.Description)
	w("%d projects in %d categories. ", len(l.Entries), len(categories))
	w("Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. ")
	if l.Meta.Podium > 0 {
		w("Within each category the %d most-starred projects take the podium and the rest are listed as contenders. ", l.Meta.Podium)
	} else {
		w("Entries are ordered by GitHub stars. ")
	}
	w("Archived repositories are removed automatically.")
	if !m.UpdatedAt.IsZero() {
		w(" Stars were last refreshed on %s.", m.UpdatedAt.UTC().Format("2006-01-02"))
	}
	w("\n\n")

	w("## Contents\n\n")
	for _, c := range categories {
		w("- [%s](#%s)\n", c.Name, Slugify(c.Name))
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
		for i, r := range rs[:podium] {
			w("- %s\n", item(r, medal(i, l.Meta.Podium)))
		}
		if rest := rs[podium:]; len(rest) > 0 {
			w("\n<details>\n<summary>Contenders (%d)</summary>\n\n", len(rest))
			for _, r := range rest {
				w("- %s\n", item(r, ""))
			}
			w("\n</details>\n")
		}
	}

	w("\n## Contributing\n\n")
	w("Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.\n")
	w("\n## License\n\n")
	w("[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)\n\n")
	w("To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work.\n")
	return []byte(b.String())
}

func medal(i, podium int) string {
	if podium <= 0 || i >= len(medals) {
		return ""
	}
	return medals[i]
}

func item(r ranked, prefix string) string {
	var b strings.Builder
	if prefix != "" {
		b.WriteString(prefix + " ")
	}
	fmt.Fprintf(&b, "[%s](%s) - %s", r.entry.DisplayName(), r.entry.URL(), r.entry.Description)
	if r.known {
		fmt.Fprintf(&b, " ★ %s", FormatStars(r.stars))
	}
	return b.String()
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
