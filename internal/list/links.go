package list

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	// LinksDir holds one JSON file per link, named after the link's title.
	LinksDir = "links"

	linkSchema = "../schema/link.schema.json"
)

// LinkSection groups links in the README and on the site, after the library
// categories. Order in list.json is the order of appearance.
type LinkSection struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Link is a resource that is not a GitHub repository, such as documentation.
// It gets the description rules of entries but no star, age or activity
// policy: those only make sense for repositories.
type Link struct {
	Schema      string `json:"$schema,omitempty"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Section     string `json:"section"`
	// Versions turns the link into a series, such as one tour per release.
	// The first version is the newest and its URL must equal URL, so the
	// title always points at the latest installment.
	Versions []LinkVersion `json:"versions,omitempty"`
	// Exempt waives checks a maintainer has decided do not apply to this link,
	// such as the live link check for sites behind bot protection.
	Exempt []string `json:"exempt,omitempty"`

	file string
}

// LinkExemptCheck skips the live HTTP check, for pages that refuse automated
// requests (for example publisher sites behind Akamai) but work in a browser.
const LinkExemptCheck = "link-check"

// LinkVersion is one installment of a series link.
type LinkVersion struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// URLs returns the link's URL followed by every version URL, without
// duplicates, in order.
func (k Link) URLs() []string {
	seen := map[string]bool{}
	var out []string
	for _, u := range append([]string{k.URL}, versionURLs(k.Versions)...) {
		if u != "" && !seen[u] {
			seen[u] = true
			out = append(out, u)
		}
	}
	return out
}

func versionURLs(vs []LinkVersion) []string {
	out := make([]string, len(vs))
	for i, v := range vs {
		out[i] = v.URL
	}
	return out
}

// IsExempt reports whether the link waives the given check.
func (k Link) IsExempt(rule string) bool {
	for _, x := range k.Exempt {
		if x == rule {
			return true
		}
	}
	return false
}

var slugNonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

func titleSlug(title string) string {
	return strings.Trim(slugNonAlnum.ReplaceAllString(strings.ToLower(title), "-"), "-")
}

// Slug is the link's anchor and file stem, derived from its title.
func (k Link) Slug() string { return titleSlug(k.Title) }

// FileName is the canonical basename for the link under LinksDir.
func (k Link) FileName() string { return k.Slug() + ".json" }

// File is the basename the link was loaded from, or "" for in-memory links.
func (k Link) File() string { return k.file }

// Path is the repository-relative path the link is (or should be) stored at.
func (k Link) Path() string {
	if k.file != "" {
		return LinksDir + "/" + k.file
	}
	return LinksDir + "/" + k.FileName()
}

// Host is the link's domain without a leading "www.", for display.
func (k Link) Host() string {
	u, err := url.Parse(k.URL)
	if err != nil {
		return ""
	}
	return strings.TrimPrefix(u.Host, "www.")
}

// ParseLink decodes one link file strictly.
func ParseLink(data []byte) (Link, error) {
	var k Link
	err := strictDecode(data, &k)
	return k, err
}

// Format returns the canonical encoding of a link file.
func (k Link) Format() ([]byte, error) {
	k.Schema = linkSchema
	return encode(&k)
}

// LinksIn returns the links of one section in title order.
func (l *List) LinksIn(section string) []Link {
	var out []Link
	for _, k := range l.Links {
		if k.Section == section {
			out = append(out, k)
		}
	}
	return out
}

func (l *List) sortLinks() {
	sort.SliceStable(l.Links, func(i, j int) bool {
		return strings.ToLower(l.Links[i].Title) < strings.ToLower(l.Links[j].Title)
	})
}

func (l *List) loadLinks(root string) error {
	paths, err := filepath.Glob(filepath.Join(root, LinksDir, "*.json"))
	if err != nil {
		return err
	}
	sort.Strings(paths)
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		k, err := ParseLink(data)
		if err != nil {
			return fmt.Errorf("%s/%s: %w", LinksDir, filepath.Base(path), err)
		}
		k.file = filepath.Base(path)
		l.Links = append(l.Links, k)
	}
	l.sortLinks()
	return nil
}

// saveLinks writes every link in canonical form and deletes link files that
// no longer correspond to a link.
func (l *List) saveLinks(root string) error {
	dir := filepath.Join(root, LinksDir)
	if len(l.Links) == 0 {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return nil
		}
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	wanted := map[string]bool{}
	for i := range l.Links {
		k := &l.Links[i]
		data, err := k.Format()
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, k.FileName()), data, 0o644); err != nil {
			return err
		}
		wanted[k.FileName()] = true
		k.file = k.FileName()
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
