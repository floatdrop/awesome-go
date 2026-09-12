package list

import (
	"strings"
	"testing"
	"time"
)

var now = time.Date(2026, 9, 12, 0, 0, 0, 0, time.UTC)

func valid() *List {
	return &List{
		Meta:       Meta{Title: "Awesome Go", Repo: "floatdrop/awesome-go"},
		Policy:     Policy{MaxDescriptionLength: 100},
		Categories: []Category{{ID: "web", Name: "Web"}, {ID: "cli", Name: "CLI"}},
		Entries: []Entry{
			{Repo: "gin-gonic/gin", Description: "HTTP web framework with a martini-like API.", Category: "web", Added: "2026-01-02"},
			{Repo: "spf13/cobra", Description: "Framework for building CLI applications.", Category: "cli"},
		},
	}
}

func TestValidateAcceptsGoodList(t *testing.T) {
	if p := valid().Validate(now); len(p) != 0 {
		t.Fatalf("unexpected problems: %v", p)
	}
}

func TestValidateRejects(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*List)
		want   string
	}{
		{"url as repo", func(l *List) { l.Entries[0].Repo = "https://github.com/gin-gonic/gin" }, "owner/name"},
		{"duplicate repo", func(l *List) { l.Entries[1].Repo = "GIN-GONIC/gin" }, "duplicate"},
		{"unknown category", func(l *List) { l.Entries[0].Category = "nope" }, "unknown category"},
		{"unsorted", func(l *List) { l.Entries[0], l.Entries[1] = l.Entries[1], l.Entries[0] }, "not sorted"},
		{"lowercase start", func(l *List) { l.Entries[0].Description = "http framework." }, "uppercase"},
		{"no period", func(l *List) { l.Entries[0].Description = "HTTP framework" }, "period"},
		{"too long", func(l *List) { l.Entries[0].Description = "X" + strings.Repeat("y", 100) + "." }, "maximum"},
		{"starts with name", func(l *List) { l.Entries[0].Description = "Gin is a web framework." }, "project name"},
		{"article", func(l *List) { l.Entries[0].Description = "A web framework." }, "article"},
		{"buzzword", func(l *List) { l.Entries[0].Description = "Blazingly fast web framework." }, "marketing"},
		{"redundant go", func(l *List) { l.Entries[0].Description = "Web framework for Go." }, "everything here is Go"},
		{"golang", func(l *List) { l.Entries[0].Description = "Web framework in golang style." }, "everything here is Go"},
		{"markdown", func(l *List) { l.Entries[0].Description = "Web framework, see [docs](https://x)." }, "plain text"},
		{"bad date", func(l *List) { l.Entries[0].Added = "12.09.2026" }, "YYYY-MM-DD"},
		{"future date", func(l *List) { l.Entries[0].Added = "2099-01-01" }, "future"},
		{"bad category id", func(l *List) { l.Categories[0].ID = "Web Stuff"; l.Entries[0].Category = "Web Stuff" }, "kebab-case"},
		{"bad list repo", func(l *List) { l.Meta.Repo = "awesome-go" }, "list.repo"},
		{"unknown exemption", func(l *List) { l.Entries[0].Exempt = []string{"stars"} }, "unknown exemption"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := valid()
			tc.mutate(l)
			problems := l.Validate(now)
			for _, p := range problems {
				if strings.Contains(p, tc.want) {
					return
				}
			}
			t.Fatalf("want a problem containing %q, got %v", tc.want, problems)
		})
	}
}

func TestParseRejectsUnknownFields(t *testing.T) {
	_, err := Parse([]byte(`{"list":{"title":"x"},"entries":[{"repo":"a/b","descripton":"typo."}]}`))
	if err == nil || !strings.Contains(err.Error(), "descripton") {
		t.Fatalf("expected unknown field error, got %v", err)
	}
}

func TestFormatIsCanonicalAndIdempotent(t *testing.T) {
	l := valid()
	l.Entries[0], l.Entries[1] = l.Entries[1], l.Entries[0]
	once, err := l.Format()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(string(once), "}\n") {
		t.Fatalf("expected trailing newline")
	}
	again, err := l.Format()
	if err != nil {
		t.Fatal(err)
	}
	if string(once) != string(again) {
		t.Fatalf("format is not idempotent")
	}
	parsed, err := Parse(once)
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Entries[0].Repo != "gin-gonic/gin" {
		t.Fatalf("entries not sorted in output: %v", parsed.Entries)
	}
}

func TestPolicyCheck(t *testing.T) {
	p := Policy{MinAgeDays: 90, MinStars: 100, MaxInactiveDays: 365, RequireLicense: true, RequireGo: true}
	good := RepoMeta{Stars: 500, IsGo: true, License: "MIT", CreatedAt: now.AddDate(-1, 0, 0), PushedAt: now.AddDate(0, -1, 0)}
	if got := p.Check(good, now, Entry{}); len(got) != 0 {
		t.Fatalf("good repo rejected: %v", got)
	}
	cases := []struct {
		name string
		meta RepoMeta
		want string
	}{
		{"young", func() RepoMeta { m := good; m.CreatedAt = now.AddDate(0, 0, -30); return m }(), "days old"},
		{"few stars", func() RepoMeta { m := good; m.Stars = 3; return m }(), "stars"},
		{"inactive", func() RepoMeta { m := good; m.PushedAt = now.AddDate(-2, 0, 0); return m }(), "last push"},
		{"no license", func() RepoMeta { m := good; m.License = ""; return m }(), "license"},
		{"not go", func() RepoMeta { m := good; m.IsGo = false; return m }(), "not Go"},
		{"fork", func() RepoMeta { m := good; m.Fork = true; return m }(), "fork"},
		{"archived", func() RepoMeta { m := good; m.Archived = true; return m }(), "archived"},
		{"missing", RepoMeta{NotFound: true}, "no longer exists"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, msg := range p.Check(tc.meta, now, Entry{}) {
				if strings.Contains(msg, tc.want) {
					return
				}
			}
			t.Fatalf("want %q in problems, got %v", tc.want, p.Check(tc.meta, now, Entry{}))
		})
	}
	waived := good
	waived.Fork = true
	waived.PushedAt = now.AddDate(-2, 0, 0)
	if got := p.Check(waived, now, Entry{Exempt: []string{ExemptFork, ExemptInactive}}); len(got) != 0 {
		t.Fatalf("exemptions not honoured: %v", got)
	}
	young := good
	young.CreatedAt = now
	if got := p.Check(young, now, Entry{Exempt: []string{ExemptFork, ExemptInactive}}); len(got) == 0 {
		t.Fatal("age must not be waivable")
	}
	if _, dead := p.Dead(good); dead {
		t.Fatal("good repo reported dead")
	}
	if days, stale := p.Stale(RepoMeta{PushedAt: now.AddDate(-3, 0, 0)}, now); !stale || days < 1000 {
		t.Fatalf("expected stale, got %d %v", days, stale)
	}
}
