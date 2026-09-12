package list

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	repoRe     = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?/[A-Za-z0-9_.-]+$`)
	idRe       = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	dateRe     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	buzzRe     = regexp.MustCompile(`(?i)\b(blazing(?:ly)?|fastest|best|awesome|amazing|revolutionary|ultimate|world-class|state-of-the-art|next-gen(?:eration)?|cutting-edge|powerful)\b`)
	redundRe   = regexp.MustCompile(`(?i)\b(golang|(?:written|implemented) in go|for go)\b`)
	articleRe  = regexp.MustCompile(`^(?i:a|an|the)\s`)
	markdownRe = regexp.MustCompile("[`*_\\[\\]<>]|https?://")
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
			if x != ExemptFork && x != ExemptInactive {
				add("%s: unknown exemption %q (allowed: %s, %s)", where, x, ExemptFork, ExemptInactive)
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

	return problems
}

func describeProblems(e Entry, maxLen int) []string {
	var problems []string
	d := e.Description
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
	name := strings.ToLower(e.DisplayName())
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
