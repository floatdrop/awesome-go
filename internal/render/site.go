package render

import (
	"bytes"
	_ "embed"
	"html/template"
	"sort"
	"strings"
	"time"

	"github.com/floatdrop/awesome-go/internal/list"
)

//go:embed site.tmpl
var siteTemplate string

var siteTpl = template.Must(template.New("site").Parse(siteTemplate))

type sitePage struct {
	Title       string
	Description string
	Repo        string
	RepoURL     string
	Branch      string
	SiteURL     string
	Updated     string
	Total       int
	Podium      int
	Categories  []siteCategory
}

type siteCategory struct {
	ID          string
	Name        string
	Description string
	Entries     []siteEntry
}

type siteEntry struct {
	Rank        int
	Medal       string
	Slug        string
	Name        string
	Repo        string
	URL         string
	Description string
	Stars       int
	StarsText   string
	Added       string
	Search      string
}

// Site renders the GitHub Pages site: one self-contained HTML file with
// client-side search, filtering and sorting over the same data as the README.
func Site(l *list.List, m *list.Metadata, now time.Time) ([]byte, error) {
	page := sitePage{
		Title:       l.Meta.Title,
		Description: l.Meta.Description,
		Repo:        l.Meta.Repo,
		RepoURL:     "https://github.com/" + l.Meta.Repo,
		Branch:      l.Meta.BranchOrMain(),
		SiteURL:     l.Meta.SiteURL(),
		Total:       len(l.Entries),
		Podium:      l.Meta.Podium,
	}
	page.Updated = m.UpdatedAt

	byCategory := map[string][]list.Entry{}
	for _, e := range l.Entries {
		byCategory[e.Category] = append(byCategory[e.Category], e)
	}
	for _, c := range l.Categories {
		entries := byCategory[c.ID]
		if len(entries) == 0 {
			continue
		}
		sc := siteCategory{ID: c.ID, Name: c.Name, Description: c.Description}
		for _, e := range entries {
			se := siteEntry{
				Slug:        list.Slug(e.Repo),
				Name:        e.DisplayName(),
				Repo:        e.Repo,
				URL:         e.URL(),
				Description: e.Description,
				Added:       e.Added,
				Search:      strings.ToLower(e.DisplayName() + " " + e.Repo + " " + e.Description),
			}
			if rm, ok := m.Repos[e.Repo]; ok {
				se.Stars = rm.Stars
				se.StarsText = FormatStars(rm.Stars)
			}
			sc.Entries = append(sc.Entries, se)
		}
		sort.SliceStable(sc.Entries, func(i, j int) bool {
			if sc.Entries[i].Stars != sc.Entries[j].Stars {
				return sc.Entries[i].Stars > sc.Entries[j].Stars
			}
			return strings.ToLower(sc.Entries[i].Name) < strings.ToLower(sc.Entries[j].Name)
		})
		for i := range sc.Entries {
			sc.Entries[i].Rank = i + 1
			if l.Meta.Podium > 0 && i < len(medals) && i < l.Meta.Podium {
				sc.Entries[i].Medal = medals[i]
			}
		}
		page.Categories = append(page.Categories, sc)
	}

	var buf bytes.Buffer
	if err := siteTpl.Execute(&buf, page); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
