package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/floatdrop/awesome-go/internal/badge"
	"github.com/floatdrop/awesome-go/internal/gh"
	"github.com/floatdrop/awesome-go/internal/list"
	"github.com/floatdrop/awesome-go/internal/render"
)

func runRefresh(args []string) error {
	fs, p := newFlags("refresh")
	if err := fs.Parse(args); err != nil {
		return err
	}
	now := time.Now().UTC()
	l, err := list.Load(p.root)
	if err != nil {
		return err
	}
	meta, err := list.LoadMetadata(p.metadata())
	if err != nil {
		return err
	}

	client := gh.New(os.Getenv("GITHUB_TOKEN"))
	fmt.Printf("refreshing %d repositories\n", len(l.Entries))
	results := fetchAll(context.Background(), client, l.Entries)

	var failures []string
	seen := map[string]bool{}
	kept := l.Entries[:0]
	for i, r := range results {
		e := l.Entries[i]
		if r.err != nil {
			failures = append(failures, r.err.Error())
			// Keep whatever we knew before; a transient error must not look like a dead repo.
			if !seen[strings.ToLower(e.Repo)] {
				seen[strings.ToLower(e.Repo)] = true
				kept = append(kept, e)
			}
			continue
		}
		if r.canonical != "" && r.canonical != e.Repo {
			fmt.Printf("renamed: %s -> %s\n", e.Repo, r.canonical)
			delete(meta.Repos, e.Repo)
			e.Repo = r.canonical
		}
		key := strings.ToLower(e.Repo)
		if seen[key] {
			fmt.Printf("dropped duplicate after rename: %s\n", e.Repo)
			continue
		}
		seen[key] = true
		if e.Added == "" {
			e.Added = now.Format("2006-01-02")
		}
		meta.Repos[e.Repo] = r.meta
		kept = append(kept, e)
	}
	l.Entries = kept

	for repo := range meta.Repos {
		if !seen[strings.ToLower(repo)] {
			delete(meta.Repos, repo)
		}
	}
	meta.UpdatedAt = now.Format("2006-01-02")

	if err := l.Save(p.root); err != nil {
		return err
	}
	if err := meta.Save(p.metadata()); err != nil {
		return err
	}
	if len(failures) > 0 {
		return fmt.Errorf("%d repositories could not be refreshed:\n  %s", len(failures), strings.Join(failures, "\n  "))
	}
	fmt.Printf("✓ %s updated\n", list.MetadataFileName)
	return nil
}

func runPrune(args []string) error {
	fs, p := newFlags("prune")
	if err := fs.Parse(args); err != nil {
		return err
	}
	now := time.Now().UTC()
	l, err := list.Load(p.root)
	if err != nil {
		return err
	}
	meta, err := list.LoadMetadata(p.metadata())
	if err != nil {
		return err
	}

	var removed, stale []string
	kept := l.Entries[:0]
	for _, e := range l.Entries {
		m, ok := meta.Repos[e.Repo]
		if !ok {
			kept = append(kept, e)
			continue
		}
		if reason, dead := l.Policy.Dead(m); dead {
			removed = append(removed, fmt.Sprintf("- [%s](%s): %s", e.Repo, e.URL(), reason))
			delete(meta.Repos, e.Repo)
			continue
		}
		if days, isStale := l.Policy.Stale(m, now); isStale {
			note := ""
			if e.IsExempt(list.ExemptInactive) {
				note = " (exempt)"
			}
			stale = append(stale, fmt.Sprintf("- [%s](%s): last push %d days ago%s", e.Repo, e.URL(), days, note))
		}
		kept = append(kept, e)
	}
	l.Entries = kept

	if len(removed) > 0 {
		summary(append([]string{"### Removed", ""}, removed...)...)
		if err := l.Save(p.root); err != nil {
			return err
		}
		if err := meta.Save(p.metadata()); err != nil {
			return err
		}
	} else {
		fmt.Println("nothing to prune")
	}
	if len(stale) > 0 {
		summary(append([]string{"", "### Stale (no pushes within policy window, review manually)", ""}, stale...)...)
	}
	return nil
}

func runGenerate(args []string) error {
	fs, p := newFlags("generate")
	if err := fs.Parse(args); err != nil {
		return err
	}
	now := time.Now().UTC()
	l, err := list.Load(p.root)
	if err != nil {
		return err
	}
	if problems := l.Validate(now); len(problems) > 0 {
		return fmt.Errorf("data is invalid, refusing to generate: %s", problems[0])
	}
	meta, err := list.LoadMetadata(p.metadata())
	if err != nil {
		return err
	}

	if err := os.WriteFile(p.readme(), render.README(l, meta, now), 0o644); err != nil {
		return err
	}
	site, err := render.Site(l, meta, now)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(p.docs(), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(p.docs(), "index.html"), site, 0o644); err != nil {
		return err
	}

	dir := p.badges()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	old, _ := filepath.Glob(filepath.Join(dir, "*.svg"))
	for _, f := range old {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	generated := 0
	skipped := 0
	for _, e := range l.Entries {
		if e.Added == "" {
			skipped++ // refresh assigns the acceptance date; the badge title mentions it.
			continue
		}
		svg := badge.SVG(badge.Options{
			Name:  e.DisplayName(),
			Seed:  e.Repo,
			Title: fmt.Sprintf("%s is listed in %s since %s", e.Repo, l.Meta.Title, e.Added),
		})
		if err := os.WriteFile(filepath.Join(dir, list.Slug(e.Repo)+".svg"), svg, 0o644); err != nil {
			return err
		}
		generated++
	}

	fmt.Printf("✓ README.md, docs/index.html and %d badges generated", generated)
	if skipped > 0 {
		fmt.Printf(" (%d entries have no added date yet; run refresh)", skipped)
	}
	fmt.Println()
	return nil
}
