package render

import (
	"strings"
	"testing"

	"github.com/floatdrop/awesome-go/internal/list"
)

func TestCompare(t *testing.T) {
	l, m := fixture(3)
	l.Entries = append(l.Entries, list.Entry{Repo: "z/new", Description: "New.", Category: "web"})
	l.Entries = append(l.Entries, list.Entry{Repo: "q/lonely", Description: "Lonely.", Category: "empty"})

	out := Compare(l, m, []list.Entry{l.Entries[len(l.Entries)-2], l.Entries[len(l.Entries)-1]})
	if !strings.HasPrefix(out, CompareMarker) {
		t.Fatalf("comment must start with the marker")
	}
	for _, want := range []string{
		"**z/new** joins _Web & HTTP_:",
		"- [b/big](https://github.com/b/big) ★ 46k — Big.\n- [c/mid](https://github.com/c/mid) ★ 2.1k — Mid.\n- [a/small](https://github.com/a/small) ★ 120 — Small.\n- [d/unknown](https://github.com/d/unknown) — Unknown.\n",
		"**q/lonely** joins _Empty_, which has no other entries yet.",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
	if strings.Contains(out, "[z/new]") {
		t.Errorf("the target must not be listed among the existing entries")
	}
	if Compare(l, m, nil) != "" {
		t.Errorf("no targets must render nothing")
	}
}
