package render

import (
	"fmt"
	"sort"
	"strings"

	"github.com/floatdrop/awesome-go/internal/list"
)

// CompareMarker starts every comparison comment so CI can find and update
// its own comment instead of posting a new one per push.
const CompareMarker = "<!-- awesome-compare -->"

// Compare renders, for each target entry, the entries already listed in the
// same category, so a reviewer can judge whether the newcomer complements or
// beats them. It returns "" when there is nothing to compare.
func Compare(l *list.List, m *list.Metadata, targets []list.Entry) string {
	if len(targets) == 0 {
		return ""
	}
	isTarget := map[string]bool{}
	for _, t := range targets {
		isTarget[strings.ToLower(t.Repo)] = true
	}
	var b strings.Builder
	b.WriteString(CompareMarker + "\n### What is already listed in the category\n\n")
	b.WriteString("An entry earns its place by covering something no listed project does, or by doing it better. For the reviewer:\n")
	for _, t := range targets {
		name := t.Category
		if c, ok := l.Category(t.Category); ok {
			name = c.Name
		}
		fmt.Fprintf(&b, "\n**%s** joins _%s_", t.Repo, name)
		var others []list.Entry
		for _, e := range l.Entries {
			if e.Category == t.Category && !isTarget[strings.ToLower(e.Repo)] {
				others = append(others, e)
			}
		}
		if len(others) == 0 {
			b.WriteString(", which has no other entries yet. This one defines the category.\n")
			continue
		}
		stars := func(e list.Entry) int { return m.Repos[e.Repo].Stars }
		sort.SliceStable(others, func(i, j int) bool { return stars(others[i]) > stars(others[j]) })
		b.WriteString(":\n\n")
		for _, e := range others {
			fmt.Fprintf(&b, "- [%s](%s)", e.Repo, e.URL())
			if rm, ok := m.Repos[e.Repo]; ok {
				fmt.Fprintf(&b, " ★ %s", FormatStars(rm.Stars))
			}
			fmt.Fprintf(&b, " — %s\n", e.Description)
		}
	}
	return b.String()
}
