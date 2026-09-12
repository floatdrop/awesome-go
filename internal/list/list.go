// Package list holds the data model behind awesome.json and metadata.json.
package list

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

// FileName is the curated source of truth, edited by humans.
const FileName = "awesome.json"

// List is the top-level document stored in awesome.json.
type List struct {
	Schema     string     `json:"$schema,omitempty"`
	Meta       Meta       `json:"list"`
	Policy     Policy     `json:"policy"`
	Categories []Category `json:"categories"`
	Entries    []Entry    `json:"entries"`
}

// Meta describes the list itself.
type Meta struct {
	Title       string `json:"title"`
	Repo        string `json:"repo"`
	Branch      string `json:"branch"`
	Description string `json:"description"`
	// Podium is how many entries per category are shown with medals before the
	// rest is collapsed into a "contenders" block. Zero shows a flat list.
	Podium int `json:"podium"`
}

// BranchOrMain returns the branch that serves generated badges.
func (m Meta) BranchOrMain() string {
	if m.Branch == "" {
		return "main"
	}
	return m.Branch
}

// Category groups entries in the generated README. Order in the file is the
// order in the README.
type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Entry is a single listed repository.
type Entry struct {
	Repo        string `json:"repo"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Added       string `json:"added,omitempty"`
	// Exempt lists policy rules a maintainer has consciously waived for this
	// entry (see ExemptFork, ExemptInactive). Age and stars cannot be waived.
	Exempt []string `json:"exempt,omitempty"`
}

// Exemptions a maintainer may grant. Anything else is a validation error.
const (
	ExemptFork     = "fork"     // a fork that became the maintained successor
	ExemptInactive = "inactive" // a finished library that does not need commits
)

// IsExempt reports whether the entry waives the given rule.
func (e Entry) IsExempt(rule string) bool {
	for _, x := range e.Exempt {
		if x == rule {
			return true
		}
	}
	return false
}

// DisplayName is the name shown in the README: the explicit name or the
// repository name without the owner.
func (e Entry) DisplayName() string {
	if e.Name != "" {
		return e.Name
	}
	if _, name, ok := strings.Cut(e.Repo, "/"); ok {
		return name
	}
	return e.Repo
}

// URL is the repository home page.
func (e Entry) URL() string { return "https://github.com/" + e.Repo }

// Slug turns "Owner/Repo" into a filesystem and URL safe "owner--repo".
func Slug(repo string) string {
	return strings.ToLower(strings.ReplaceAll(repo, "/", "--"))
}

// Load reads and parses a list file.
func Load(path string) (*List, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	l, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return l, nil
}

// Parse decodes a list strictly: unknown fields are errors so typos in
// contributions are caught instead of silently ignored.
func Parse(data []byte) (*List, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var l List
	if err := dec.Decode(&l); err != nil {
		return nil, err
	}
	if dec.More() {
		return nil, errors.New("trailing data after JSON document")
	}
	if l.Categories == nil {
		l.Categories = []Category{}
	}
	if l.Entries == nil {
		l.Entries = []Entry{}
	}
	return &l, nil
}

// Sort orders entries by repository name, case-insensitively. A stable, boring
// order in the source file keeps diffs small and merge conflicts rare.
func (l *List) Sort() {
	sort.SliceStable(l.Entries, func(i, j int) bool {
		return strings.ToLower(l.Entries[i].Repo) < strings.ToLower(l.Entries[j].Repo)
	})
}

// Format returns the canonical encoding of the list: sorted entries, two-space
// indentation, trailing newline.
func (l *List) Format() ([]byte, error) {
	c := *l
	c.Entries = append([]Entry{}, l.Entries...)
	c.Sort()
	if c.Categories == nil {
		c.Categories = []Category{}
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(&c); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Save writes the list in canonical form.
func (l *List) Save(path string) error {
	data, err := l.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Category looks a category up by id.
func (l *List) Category(id string) (Category, bool) {
	for _, c := range l.Categories {
		if c.ID == id {
			return c, true
		}
	}
	return Category{}, false
}

// Remove drops every entry with the given repository name (case-insensitive)
// and reports whether anything was removed.
func (l *List) Remove(repo string) bool {
	kept := l.Entries[:0]
	removed := false
	for _, e := range l.Entries {
		if strings.EqualFold(e.Repo, repo) {
			removed = true
			continue
		}
		kept = append(kept, e)
	}
	l.Entries = kept
	return removed
}
