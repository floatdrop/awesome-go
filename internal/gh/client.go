// Package gh is a minimal GitHub REST client: only what the list needs.
package gh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// ErrNotFound is returned for 404 responses.
var ErrNotFound = errors.New("not found")

// RateLimitError is returned when the API rate limit is exhausted.
type RateLimitError struct{ Reset time.Time }

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("GitHub API rate limit exceeded, resets at %s", e.Reset.Format(time.RFC3339))
}

// Client talks to the GitHub REST API.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// New returns a client for api.github.com. An empty token works but is
// limited to 60 requests per hour.
func New(token string) *Client {
	return &Client{
		BaseURL:    "https://api.github.com",
		Token:      token,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// Repo is the subset of the repository object the list cares about.
type Repo struct {
	FullName  string    `json:"full_name"`
	Stars     int       `json:"stargazers_count"`
	Archived  bool      `json:"archived"`
	Disabled  bool      `json:"disabled"`
	Fork      bool      `json:"fork"`
	Language  string    `json:"language"`
	License   *License  `json:"license"`
	CreatedAt time.Time `json:"created_at"`
	PushedAt  time.Time `json:"pushed_at"`
}

// License as detected by GitHub. SPDXID is "NOASSERTION" for unrecognised
// license files, which still counts as having a license.
type License struct {
	SPDXID string `json:"spdx_id"`
	Name   string `json:"name"`
}

// Repo fetches a repository. Renamed or transferred repositories are followed
// through GitHub's redirect, so FullName may differ from the requested name.
func (c *Client) Repo(ctx context.Context, fullName string) (*Repo, error) {
	var r Repo
	if err := c.get(ctx, "/repos/"+fullName, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", fullName, err)
	}
	return &r, nil
}

// FileExists reports whether a file exists at the given path in the default
// branch of a repository.
func (c *Client) FileExists(ctx context.Context, fullName, path string) (bool, error) {
	err := c.get(ctx, "/repos/"+fullName+"/contents/"+path, nil)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, ErrNotFound):
		return false, nil
	default:
		return false, fmt.Errorf("%s/%s: %w", fullName, path, err)
	}
}

func (c *Client) get(ctx context.Context, path string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "awesome-go-list")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	hc := c.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == http.StatusOK:
		if out == nil {
			return nil
		}
		return json.NewDecoder(resp.Body).Decode(out)
	case resp.StatusCode == http.StatusNotFound:
		return ErrNotFound
	case resp.StatusCode == http.StatusTooManyRequests,
		resp.StatusCode == http.StatusForbidden && resp.Header.Get("X-RateLimit-Remaining") == "0":
		reset := time.Now().Add(time.Minute)
		if s, err := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64); err == nil {
			reset = time.Unix(s, 0)
		}
		return &RateLimitError{Reset: reset}
	default:
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("GitHub API returned %s: %s", resp.Status, string(body))
	}
}
