package render

import (
	"strings"
	"testing"
	"time"

	"github.com/floatdrop/awesome-go/internal/list"
)

func fixture(podium int) (*list.List, *list.Metadata) {
	l := &list.List{
		Meta:       list.Meta{Title: "Awesome Go", Repo: "x/y", Description: "Curated.", Podium: podium},
		Categories: []list.Category{{ID: "web", Name: "Web & HTTP", Description: "Servers."}, {ID: "empty", Name: "Empty"}},
		Entries: []list.Entry{
			{Repo: "a/small", Description: "Small.", Category: "web"},
			{Repo: "b/big", Description: "Big.", Category: "web"},
			{Repo: "c/mid", Name: "middle", Description: "Mid.", Category: "web"},
			{Repo: "d/unknown", Description: "Unknown.", Category: "web"},
		},
	}
	m := &list.Metadata{UpdatedAt: "2026-09-12", Repos: map[string]list.RepoMeta{
		"a/small": {Stars: 120},
		"b/big":   {Stars: 45678},
		"c/mid":   {Stars: 2050},
	}}
	return l, m
}

func TestREADMEOrderAndPodium(t *testing.T) {
	l, m := fixture(2)
	out := string(README(l, m, time.Now()))

	if !strings.HasPrefix(out, Header) {
		t.Fatalf("missing generated header")
	}
	for _, want := range []string{
		"# Awesome Go\n\n> Curated.\n",
		"## Contents\n\n- [Web & HTTP](#web--http)\n",
		"## Web & HTTP\n\nServers.\n\n",
		"- 🥇 [big](https://github.com/b/big) - Big. ★ 46k\n",
		"- 🥈 [middle](https://github.com/c/mid) - Mid. ★ 2.1k\n",
		"<summary>More (2)</summary>\n\n- [small](https://github.com/a/small) - Small. ★ 120\n- [unknown](https://github.com/d/unknown) - Unknown.\n",
		"last refreshed on 2026-09-12",
		"An entry earns its place by covering something no listed project does, or by doing it better.",
		"Search, filter and sort the full list at [x.github.io/y](https://x.github.io/y/).",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("README missing %q\n---\n%s", want, out)
		}
	}
	if strings.Contains(out, "## Empty") || strings.Contains(out, "#empty") {
		t.Errorf("empty category should be skipped")
	}
	if strings.Contains(out, "🥉") {
		t.Errorf("podium of 2 must not award bronze")
	}
}

func TestREADMEPadsPodiumEntriesWithoutMedal(t *testing.T) {
	l, m := fixture(4)
	out := string(README(l, m, time.Now()))
	if !strings.Contains(out, "- 🥉 [small](") || !strings.Contains(out, "- &emsp;&#8196; [unknown](") {
		t.Fatalf("fourth podium entry must be padded:\n%s", out)
	}
	if strings.Contains(out, "<details>") {
		t.Fatalf("no More block expected when every entry is on the podium")
	}
}

func TestREADMEFlatWithoutPodium(t *testing.T) {
	l, m := fixture(0)
	out := string(README(l, m, time.Now()))
	if strings.Contains(out, "<details>") || strings.Contains(out, "🥇") {
		t.Fatalf("podium disabled but rendered:\n%s", out)
	}
	big := strings.Index(out, "[big]")
	small := strings.Index(out, "[small]")
	if big < 0 || small < 0 || big > small {
		t.Fatalf("entries not ordered by stars")
	}
}

func TestFormatStars(t *testing.T) {
	cases := map[int]string{0: "0", 999: "999", 1000: "1k", 1234: "1.2k", 9950: "10k", 45678: "46k", 123456: "123k"}
	for n, want := range cases {
		if got := FormatStars(n); got != want {
			t.Errorf("FormatStars(%d) = %q, want %q", n, got, want)
		}
	}
}

func TestSlugify(t *testing.T) {
	cases := map[string]string{"Web Frameworks": "web-frameworks", "Web & HTTP": "web--http", "gRPC/RPC": "grpcrpc"}
	for in, want := range cases {
		if got := Slugify(in); got != want {
			t.Errorf("Slugify(%q) = %q, want %q", in, got, want)
		}
	}
}
