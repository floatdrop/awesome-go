package main

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/floatdrop/awesome-go/internal/gh"
	"github.com/floatdrop/awesome-go/internal/list"
)

const fetchWorkers = 5

type fetchResult struct {
	entry     list.Entry
	meta      list.RepoMeta
	canonical string // full name as GitHub reports it; differs after a rename
	err       error
}

// fetchAll queries GitHub for every entry with bounded concurrency and returns
// results in input order.
func fetchAll(ctx context.Context, c *gh.Client, entries []list.Entry) []fetchResult {
	results := make([]fetchResult, len(entries))
	sem := make(chan struct{}, fetchWorkers)
	var wg sync.WaitGroup
	for i, e := range entries {
		wg.Add(1)
		go func(i int, e list.Entry) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			m, canonical, err := fetchOne(ctx, c, e)
			results[i] = fetchResult{entry: e, meta: m, canonical: canonical, err: err}
		}(i, e)
	}
	wg.Wait()
	return results
}

func fetchOne(ctx context.Context, c *gh.Client, e list.Entry) (list.RepoMeta, string, error) {
	r, err := c.Repo(ctx, e.Repo)
	if errors.Is(err, gh.ErrNotFound) {
		return list.RepoMeta{NotFound: true}, "", nil
	}
	if err != nil {
		return list.RepoMeta{}, "", err
	}
	m := list.RepoMeta{
		Stars:     r.Stars,
		Archived:  r.Archived,
		Disabled:  r.Disabled,
		Fork:      r.Fork,
		Language:  r.Language,
		CreatedAt: r.CreatedAt,
		PushedAt:  r.PushedAt,
	}
	if r.License != nil {
		m.License = r.License.SPDXID
		if m.License == "" || m.License == "NOASSERTION" {
			m.License = r.License.Name
		}
	}
	m.IsGo = strings.EqualFold(r.Language, "Go")
	if !m.IsGo {
		if ok, err := c.FileExists(ctx, r.FullName, "go.mod"); err == nil {
			m.IsGo = ok
		}
	}
	return m, r.FullName, nil
}
