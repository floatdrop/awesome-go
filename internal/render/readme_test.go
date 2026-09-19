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
		"- <img src=\"badges/stars/b--big.svg\" alt=\"46k stars\" align=\"middle\"> [big](https://github.com/b/big) - Big.\n",
		"- <img src=\"badges/stars/c--mid.svg\" alt=\"2.1k stars\" align=\"middle\"> [middle](https://github.com/c/mid) - Mid.\n",
		"<summary>More (2)</summary>\n\n- <img src=\"badges/stars/a--small.svg\" alt=\"120 stars\" align=\"middle\"> [small](https://github.com/a/small) - Small.\n- <img src=\"badges/stars/d--unknown.svg\" alt=\"new\" align=\"middle\"> [unknown](https://github.com/d/unknown) - Unknown.\n",
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
	if strings.Contains(out, "🥇") || strings.Contains(out, "★") {
		t.Errorf("medals and inline star counts were replaced by star pills")
	}
}

func TestREADMEPodiumWithoutMore(t *testing.T) {
	l, m := fixture(4)
	out := string(README(l, m, time.Now()))
	if !strings.Contains(out, "- <img src=\"badges/stars/a--small.svg\" alt=\"120 stars\" align=\"middle\"> [small](") || !strings.Contains(out, "- <img src=\"badges/stars/d--unknown.svg\" alt=\"new\" align=\"middle\"> [unknown](") {
		t.Fatalf("every podium entry must carry a star pill:\n%s", out)
	}
	if strings.Contains(out, "<details>") {
		t.Fatalf("no More block expected when every entry is on the podium")
	}
}

func TestREADMEHeaderWithLogoAndDiscord(t *testing.T) {
	l, m := fixture(3)
	l.Meta.Logo = "assets/logotype.svg"
	l.Meta.Discord = &list.Discord{Server: "1548529967288029294"}
	out := string(README(l, m, time.Now()))
	for _, want := range []string{
		"<p align=\"center\">\n  <img src=\"assets/logotype.svg\" alt=\"The Go gopher wearing pixel sunglasses\" width=\"180\">\n</p>\n\n<h1 align=\"center\">Awesome Go</h1>\n\n",
		"<a href=\"https://discord.com/channels/1548529967288029294\"><img src=\"https://img.shields.io/discord/1548529967288029294?logo=discord&logoColor=white&label=discord&color=5865F2\" alt=\"Discord\"></a>",
		"The logo is excluded: it is based on the Go gopher by Renée French",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("README missing %q\n---\n%s", want, out)
		}
	}
	if strings.Contains(out, "# Awesome Go\n") {
		t.Errorf("markdown title must be replaced by the centered heading when a logo is set")
	}
	l.Meta.Discord.Invite = "https://discord.gg/abc"
	if out := string(README(l, m, time.Now())); !strings.Contains(out, `<a href="https://discord.gg/abc">`) {
		t.Errorf("invite must take precedence over the server URL")
	}
}

func TestREADMELinkSections(t *testing.T) {
	l, m := fixture(3)
	l.LinkSections = []list.LinkSection{{ID: "documentation", Name: "Documentation", Description: "Guides."}, {ID: "books", Name: "Books"}}
	l.Links = []list.Link{{Title: "A Tour of Go", URL: "https://go.dev/tour/", Description: "Interactive introduction.", Section: "documentation"}}
	out := string(README(l, m, time.Now()))
	for _, want := range []string{
		"- [Web & HTTP](#web--http)\n- [Documentation](#documentation)\n",
		"\n## Documentation\n\nGuides.\n\n- [A Tour of Go](https://go.dev/tour/) - Interactive introduction.\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("README missing %q\n---\n%s", want, out)
		}
	}
	if strings.Contains(out, "Books") {
		t.Errorf("empty link sections must be skipped")
	}

	l.Links = append(l.Links, list.Link{
		Title: "Release Tours", URL: "https://example.com/2", Description: "Tours.", Section: "documentation",
		Versions: []list.LinkVersion{{Label: "2", URL: "https://example.com/2"}, {Label: "1", URL: "https://example.com/1"}},
	})
	out = string(README(l, m, time.Now()))
	if want := "- [Release Tours](https://example.com/2) - Tours. Versions: [2](https://example.com/2) · [1](https://example.com/1)\n"; !strings.Contains(out, want) {
		t.Errorf("README missing %q\n---\n%s", want, out)
	}
	if strings.Index(out, "## Documentation") > strings.Index(out, "## Contributing") || strings.Index(out, "## Documentation") < strings.Index(out, "## Web & HTTP") {
		t.Errorf("link sections must come after the categories and before Contributing")
	}
}

func TestREADMEFlatWithoutPodium(t *testing.T) {
	l, m := fixture(0)
	out := string(README(l, m, time.Now()))
	if strings.Contains(out, "<details>") {
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
