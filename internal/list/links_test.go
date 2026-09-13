package list

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withLinks() *List {
	l := valid()
	l.LinkSections = []LinkSection{{ID: "documentation", Name: "Documentation"}}
	l.Links = []Link{{
		Title:       "A Tour of Go",
		URL:         "https://go.dev/tour/",
		Description: "Interactive introduction that runs in the browser.",
		Section:     "documentation",
	}}
	return l
}

func TestValidateAcceptsLinks(t *testing.T) {
	if p := withLinks().Validate(now); len(p) != 0 {
		t.Fatalf("unexpected problems: %v", p)
	}
}

func TestValidateRejectsLinks(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*List)
		want   string
	}{
		{"http url", func(l *List) { l.Links[0].URL = "http://go.dev/tour/" }, "https URL"},
		{"relative url", func(l *List) { l.Links[0].URL = "/tour" }, "https URL"},
		{"repository url", func(l *List) { l.Links[0].URL = "https://github.com/golang/tour" }, "belongs in entries/"},
		{"unknown section", func(l *List) { l.Links[0].Section = "books" }, "unknown link section"},
		{"empty title", func(l *List) { l.Links[0].Title = "" }, "title must be plain text"},
		{"markdown title", func(l *List) { l.Links[0].Title = "[Tour](x)" }, "title must be plain text"},
		{"misnamed file", func(l *List) { l.Links[0].file = "tour.json" }, "must be named a-tour-of-go.json"},
		{"description rules", func(l *List) { l.Links[0].Description = "the tour." }, "uppercase"},
		{"duplicate url", func(l *List) {
			d := l.Links[0]
			d.Title = "Tour Again"
			d.URL = "https://go.dev/tour"
			l.Links = append(l.Links, d)
		}, "duplicate of"},
		{"section clashes with category", func(l *List) { l.LinkSections[0].ID = "web"; l.Links[0].Section = "web" }, "already used by a category"},
		{"duplicate section", func(l *List) { l.LinkSections = append(l.LinkSections, l.LinkSections[0]) }, "duplicate link section"},
		{"unknown link exemption", func(l *List) { l.Links[0].Exempt = []string{"stars"} }, "unknown link exemption"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := withLinks()
			tc.mutate(l)
			for _, p := range l.Validate(now) {
				if strings.Contains(p.String(), tc.want) {
					if !strings.HasPrefix(p.File, LinksDir+"/") && !strings.HasPrefix(p.File, ListFileName) {
						t.Fatalf("problem attached to unexpected file %q", p.File)
					}
					return
				}
			}
			t.Fatalf("want a problem containing %q, got %v", tc.want, l.Validate(now))
		})
	}
}

func TestLinksSaveLoadRoundTrip(t *testing.T) {
	root := t.TempDir()
	l := withLinks()
	if err := l.Save(root); err != nil {
		t.Fatal(err)
	}
	stray := filepath.Join(root, LinksDir, "old.json")
	if err := os.WriteFile(stray, []byte(`{"title":"Old","url":"https://example.com/","description":"Old.","section":"documentation"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := l.Save(root); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(stray); !os.IsNotExist(err) {
		t.Fatalf("Save must delete link files that do not belong to a link")
	}
	loaded, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Links) != 1 || loaded.Links[0].File() != "a-tour-of-go.json" || loaded.Links[0].Schema != linkSchema {
		t.Fatalf("unexpected links after reload: %+v", loaded.Links)
	}
	if len(loaded.LinkSections) != 1 || loaded.LinkSections[0].ID != "documentation" {
		t.Fatalf("link sections not persisted in list.json: %+v", loaded.LinkSections)
	}
	if p := loaded.Validate(now); len(p) != 0 {
		t.Fatalf("canonical files must validate: %v", p)
	}
	if _, err := ParseLink([]byte(`{"title":"x","ur":"typo"}`)); err == nil {
		t.Fatalf("unknown fields must be rejected")
	}
}

func TestListWithoutLinksWritesNoLinksDir(t *testing.T) {
	root := t.TempDir()
	if err := valid().Save(root); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, LinksDir)); !os.IsNotExist(err) {
		t.Fatalf("a list without links must not create links/")
	}
}
