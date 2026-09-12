package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// MetadataFileName caches what we last learned about each repository from the
// GitHub API. It is generated, never edited by hand.
const MetadataFileName = "metadata.json"

// Metadata maps "owner/repo" to the facts fetched about it. It deliberately
// holds no timestamps finer than a day: the file must only change when a real
// value changes, so a sync run that learns nothing new commits nothing.
type Metadata struct {
	UpdatedAt string              `json:"updated_at"` // YYYY-MM-DD of the last refresh
	Repos     map[string]RepoMeta `json:"repos"`
}

// RepoMeta is everything the policy and the README need to know about a
// repository.
type RepoMeta struct {
	Stars     int       `json:"stars"`
	Archived  bool      `json:"archived"`
	Disabled  bool      `json:"disabled,omitempty"`
	NotFound  bool      `json:"not_found,omitempty"`
	Fork      bool      `json:"fork,omitempty"`
	IsGo      bool      `json:"is_go"`
	Language  string    `json:"language,omitempty"`
	License   string    `json:"license,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	PushedAt  time.Time `json:"pushed_at"`
}

// LoadMetadata reads the cache; a missing file is an empty cache.
func LoadMetadata(path string) (*Metadata, error) {
	m := &Metadata{Repos: map[string]RepoMeta{}}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return m, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, m); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if m.Repos == nil {
		m.Repos = map[string]RepoMeta{}
	}
	return m, nil
}

// Save writes the cache with sorted keys so diffs stay readable.
func (m *Metadata) Save(path string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}
