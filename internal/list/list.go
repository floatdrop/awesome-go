// Package list holds the data model behind list.json, entries/ and
// metadata.json.
package list

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	// ListFileName holds the list metadata, the policy and the categories.
	ListFileName = "list.json"
	// EntriesDir holds one JSON file per entry, named after the repository slug.
	EntriesDir = "entries"

	listSchema  = "schema/list.schema.json"
	entrySchema = "../schema/entry.schema.json"
)

// List is the in-memory view of the whole data set.
type List struct {
	Schema     string     `json:"$schema,omitempty"`
	Meta       Meta       `json:"list"`
	Policy     Policy     `json:"policy"`
	Categories []Category `json:"categories"`
	Entries    []Entry    `json:"-"`
}

// Meta describes the list itself.
type Meta struct {
	Title       string `json:"title"`
	Repo        string `json:"repo"`
	Branch      string `json:"branch"`
	Description string `json:"description"`
	// Site overrides the GitHub Pages URL, for a custom domain. Defaults to
	// https://<owner>.github.io/<name>/.
	Site string `json:"site,omitempty"`
	// Podium is how many entries per category are shown before the rest is
	// collapsed into a "More" block. The first three get medals. Zero shows a
	// flat list.
	Podium int `json:"podium"`
}

// SiteURL is where the generated docs/ site is published.
func (m Meta) SiteURL() string {
	if m.Site != "" {
		return m.Site
	}
	owner, name, _ := strings.Cut(m.Repo, "/")
	return "https://" + strings.ToLower(owner) + ".github.io/" + name + "/"
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

// Entry is a single listed repository, stored as entries/<owner>--<name>.json.
type Entry struct {
	Schema      string `json:"$schema,omitempty"`
	Repo        string `json:"repo"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Added       string `json:"added,omitempty"`
	// Exempt lists policy rules a maintainer has consciously waived for this
	// entry (see ExemptFork, ExemptInactive). Age and stars cannot be waived.
	Exempt []string `json:"exempt,omitempty"`

	file string // basename the entry was loaded from; empty when created in memory
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

// FileName is the canonical basename for the entry under EntriesDir.
func (e Entry) FileName() string { return Slug(e.Repo) + ".json" }

// File is the basename the entry was loaded from, or "" for in-memory entries.
func (e Entry) File() string { return e.file }

// Path is the repository-relative path the entry is (or should be) stored at.
func (e Entry) Path() string {
	if e.file != "" {
		return EntriesDir + "/" + e.file
	}
	return EntriesDir + "/" + e.FileName()
}

// Slug turns "Owner/Repo" into a filesystem and URL safe "owner--repo".
func Slug(repo string) string {
	return strings.ToLower(strings.ReplaceAll(repo, "/", "--"))
}

// Load reads list.json and every entries/*.json under root.
func Load(root string) (*List, error) {
	data, err := os.ReadFile(filepath.Join(root, ListFileName))
	if err != nil {
		return nil, err
	}
	l, err := ParseList(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ListFileName, err)
	}
	paths, err := filepath.Glob(filepath.Join(root, EntriesDir, "*.json"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		e, err := ParseEntry(data)
		if err != nil {
			return nil, fmt.Errorf("%s/%s: %w", EntriesDir, filepath.Base(path), err)
		}
		e.file = filepath.Base(path)
		l.Entries = append(l.Entries, e)
	}
	l.Sort()
	return l, nil
}

func strictDecode(data []byte, out any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(out); err != nil {
		return err
	}
	if dec.More() {
		return errors.New("trailing data after JSON document")
	}
	return nil
}

// ParseList decodes list.json strictly: unknown fields are errors so typos are
// caught instead of silently ignored.
func ParseList(data []byte) (*List, error) {
	var l List
	if err := strictDecode(data, &l); err != nil {
		return nil, err
	}
	if l.Categories == nil {
		l.Categories = []Category{}
	}
	l.Entries = []Entry{}
	return &l, nil
}

// ParseEntry decodes one entry file strictly.
func ParseEntry(data []byte) (Entry, error) {
	var e Entry
	err := strictDecode(data, &e)
	return e, err
}

// Sort orders entries by repository name, case-insensitively.
func (l *List) Sort() {
	sort.SliceStable(l.Entries, func(i, j int) bool {
		return strings.ToLower(l.Entries[i].Repo) < strings.ToLower(l.Entries[j].Repo)
	})
}

func encode(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// FormatList returns the canonical encoding of list.json.
func (l *List) FormatList() ([]byte, error) {
	c := *l
	c.Schema = listSchema
	if c.Categories == nil {
		c.Categories = []Category{}
	}
	return encode(&c)
}

// Format returns the canonical encoding of an entry file.
func (e Entry) Format() ([]byte, error) {
	e.Schema = entrySchema
	return encode(&e)
}

// Save writes list.json and one file per entry in canonical form, and deletes
// entry files that no longer correspond to an entry (removed or renamed).
func (l *List) Save(root string) error {
	data, err := l.FormatList()
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, ListFileName), data, 0o644); err != nil {
		return err
	}
	dir := filepath.Join(root, EntriesDir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	wanted := map[string]bool{}
	for i := range l.Entries {
		e := &l.Entries[i]
		data, err := e.Format()
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, e.FileName()), data, 0o644); err != nil {
			return err
		}
		wanted[e.FileName()] = true
		e.file = e.FileName()
	}
	existing, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return err
	}
	for _, path := range existing {
		if !wanted[filepath.Base(path)] {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}
	return nil
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
