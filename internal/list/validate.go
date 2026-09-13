package list

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	repoRe      = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?/[A-Za-z0-9_.-]+$`)
	idRe        = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	discordIDRe = regexp.MustCompile(`^[0-9]{17,20}$`)
	dateRe      = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	buzzRe      = regexp.MustCompile(`(?i)\b(blazing(?:ly)?|fastest|best|awesome|amazing|revolutionary|ultimate|world-class|state-of-the-art|next-gen(?:eration)?|cutting-edge|powerful)\b`)
	redundRe    = regexp.MustCompile(`(?i)\b(golang|(?:written|implemented) in go|for go)\b`)
	articleRe   = regexp.MustCompile(`^(?i:a|an|the)\s`)
	markdownRe  = regexp.MustCompile("[`*_\\[\\]<>]|https?://")
)

// Problem is a validation failure attached to the file it was found in.
type Problem struct {
	File string
	Msg  string
}

func (p Problem) String() string {
	if p.File == "" {
		return p.Msg
	}
	return p.File + ": " + p.Msg
}

// Validate applies every rule that can be checked without network access and
// returns human-readable problems. An empty result means the data is valid.
func (l *List) Validate(now time.Time) []Problem {
	var problems []Problem
	file := ListFileName
	add := func(format string, args ...any) {
		problems = append(problems, Problem{File: file, Msg: fmt.Sprintf(format, args...)})
	}

	if l.Meta.Title == "" {
		add("list.title is required")
	}
	if !repoRe.MatchString(l.Meta.Repo) {
		add("list.repo must be \"owner/name\", got %q", l.Meta.Repo)
	}
	if l.Meta.Podium < 0 {
		add("list.podium must not be negative")
	}
	if d := l.Meta.Discord; d != nil {
		if !discordIDRe.MatchString(d.Server) {
			add("list.discord.server must be the numeric server id, got %q", d.Server)
		}
		if d.Invite != "" && !strings.HasPrefix(d.Invite, "https://discord.gg/") && !strings.HasPrefix(d.Invite, "https://discord.com/invite/") {
			add("list.discord.invite must be a discord.gg or discord.com/invite URL")
		}
	}
	if strings.HasPrefix(l.Meta.Logo, "/") || strings.Contains(l.Meta.Logo, "://") || strings.HasPrefix(l.Meta.Logo, "docs/") {
		add("list.logo must be a repository-relative path outside docs/, got %q", l.Meta.Logo)
	}

	seenCategory := map[string]bool{}
	for i, c := range l.Categories {
		where := fmt.Sprintf("categories[%d]", i)
		if !idRe.MatchString(c.ID) {
			add("%s: id %q must be kebab-case", where, c.ID)
		}
		if seenCategory[c.ID] {
			add("%s: duplicate category id %q", where, c.ID)
		}
		seenCategory[c.ID] = true
		if strings.TrimSpace(c.Name) == "" {
			add("%s: name is required", where)
		}
	}

	seenRepo := map[string]string{}
	for _, e := range l.Entries {
		file = e.Path()
		where := e.Repo
		switch {
		case e.Repo == "":
			add("repo is required")
			continue
		case !repoRe.MatchString(e.Repo):
			add("repo must be \"owner/name\" without a URL")
			continue
		case strings.HasSuffix(strings.ToLower(e.Repo), ".git"):
			add("%s: repo must not end with .git", where)
		}
		if e.File() != "" && e.File() != e.FileName() {
			add("file must be named %s to match repo %s (run `go run ./cmd/awesome fmt`)", e.FileName(), e.Repo)
		}
		key := strings.ToLower(e.Repo)
		if other, dup := seenRepo[key]; dup {
			add("%s: duplicate of %s", where, other)
		}
		seenRepo[key] = e.Path()

		if !seenCategory[e.Category] {
			add("%s: unknown category %q", where, e.Category)
		}
		if e.Name != "" && (strings.TrimSpace(e.Name) != e.Name || markdownRe.MatchString(e.Name)) {
			add("%s: name must be plain text without surrounding whitespace", where)
		}
		for _, p := range describeProblems(e, l.Policy.maxDescription()) {
			add("%s: %s", where, p)
		}
		seenExempt := map[string]bool{}
		for _, x := range e.Exempt {
			if !KnownExemption(x) {
				add("%s: unknown exemption %q (allowed: %s)", where, x, strings.Join(Exemptions, ", "))
			}
			if seenExempt[x] {
				add("%s: duplicate exemption %q", where, x)
			}
			seenExempt[x] = true
		}
		if e.Added != "" {
			t, err := time.Parse("2006-01-02", e.Added)
			if !dateRe.MatchString(e.Added) || err != nil {
				add("%s: added must be a YYYY-MM-DD date", where)
			} else if t.After(now) {
				add("%s: added date is in the future", where)
			}
		}
	}

	file = ListFileName
	seenSection := map[string]bool{}
	for i, s := range l.LinkSections {
		where := fmt.Sprintf("link_sections[%d]", i)
		if !idRe.MatchString(s.ID) {
			add("%s: id %q must be kebab-case", where, s.ID)
		}
		if seenSection[s.ID] {
			add("%s: duplicate link section id %q", where, s.ID)
		}
		if seenCategory[s.ID] {
			add("%s: id %q is already used by a category", where, s.ID)
		}
		seenSection[s.ID] = true
		if strings.TrimSpace(s.Name) == "" {
			add("%s: name is required", where)
		}
	}
	seenURL := map[string]string{}
	for _, k := range l.Links {
		file = k.Path()
		if strings.TrimSpace(k.Title) == "" || strings.TrimSpace(k.Title) != k.Title || markdownRe.MatchString(k.Title) {
			add("title must be plain text without surrounding whitespace")
		} else if k.File() != "" && k.File() != k.FileName() {
			add("file must be named %s to match its title (run `go run ./cmd/awesome fmt`)", k.FileName())
		}
		u, err := url.Parse(k.URL)
		switch {
		case err != nil || u.Scheme != "https" || u.Host == "":
			add("url must be an absolute https URL, got %q", k.URL)
		case strings.EqualFold(u.Host, "github.com") && len(strings.Split(strings.Trim(u.Path, "/"), "/")) == 2:
			add("%s is a repository; it belongs in entries/, not links/", k.URL)
		}
		key := strings.TrimSuffix(strings.ToLower(k.URL), "/")
		if other, dup := seenURL[key]; dup {
			add("duplicate of %s", other)
		}
		seenURL[key] = k.Path()
		if !seenSection[k.Section] {
			add("unknown link section %q", k.Section)
		}
		for _, x := range k.Exempt {
			if x != LinkExemptCheck {
				add("unknown link exemption %q (allowed: %s)", x, LinkExemptCheck)
			}
		}
		for _, p := range describeText(k.Title, k.Description, l.Policy.maxDescription()) {
			add("%s", p)
		}
	}

	return problems
}

func describeProblems(e Entry, maxLen int) []string {
	return describeText(e.DisplayName(), e.Description, maxLen)
}

// describeText applies the description rules to anything listed under name.
func describeText(displayName, d string, maxLen int) []string {
	var problems []string
	if d == "" {
		return []string{"description is required"}
	}
	if strings.TrimSpace(d) != d {
		problems = append(problems, "description has leading or trailing whitespace")
	}
	first, _ := utf8.DecodeRuneInString(d)
	if !unicode.IsUpper(first) && !unicode.IsDigit(first) {
		problems = append(problems, "description must start with an uppercase letter")
	}
	if !strings.HasSuffix(d, ".") {
		problems = append(problems, "description must end with a period")
	}
	if n := utf8.RuneCountInString(d); n > maxLen {
		problems = append(problems, fmt.Sprintf("description is %d characters, maximum is %d", n, maxLen))
	}
	if strings.Contains(d, "  ") || strings.ContainsAny(d, "\n\t") {
		problems = append(problems, "description must be a single line with single spaces")
	}
	if markdownRe.MatchString(d) {
		problems = append(problems, "description must be plain text: no markdown, links or HTML")
	}
	lower := strings.ToLower(d)
	name := strings.ToLower(displayName)
	if strings.HasPrefix(lower, name+" ") || strings.HasPrefix(lower, name+",") {
		problems = append(problems, "description must not start with the project name")
	}
	if articleRe.MatchString(d) {
		problems = append(problems, "description must not start with an article (a, an, the)")
	}
	if m := buzzRe.FindString(d); m != "" {
		problems = append(problems, fmt.Sprintf("description contains marketing language: %q", m))
	}
	if m := redundRe.FindString(d); m != "" {
		problems = append(problems, fmt.Sprintf("description says %q; everything here is Go, say what it does instead", m))
	}
	return problems
}
