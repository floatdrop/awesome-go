package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/floatdrop/awesome-go/internal/list"
)

var linkClient = &http.Client{Timeout: 20 * time.Second}

// checkLinks requests every link not exempt from the live check and returns a
// problem for each one that is unreachable or answers with an HTTP error.
func checkLinks(ctx context.Context, links []list.Link) []list.Problem {
	var checked []list.Link
	for _, k := range links {
		if !k.IsExempt(list.LinkExemptCheck) {
			checked = append(checked, k)
		}
	}
	links = checked
	results := make([]string, len(links))
	sem := make(chan struct{}, fetchWorkers)
	var wg sync.WaitGroup
	for i, k := range links {
		wg.Add(1)
		go func(i int, k list.Link) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			var failed []string
			for _, u := range k.URLs() {
				if msg := checkLink(ctx, linkClient, u); msg != "" {
					failed = append(failed, u+": "+msg)
				}
			}
			results[i] = strings.Join(failed, "; ")
		}(i, k)
	}
	wg.Wait()
	var problems []list.Problem
	for i, msg := range results {
		if msg != "" {
			problems = append(problems, list.Problem{File: links[i].Path(), Msg: msg})
		}
	}
	return problems
}

// checkLink returns "" when the URL answers, after redirects, with a status
// below 400, and a human-readable reason otherwise.
func checkLink(ctx context.Context, c *http.Client, url string) string {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "invalid link: " + err.Error()
	}
	req.Header.Set("User-Agent", "awesome-go-list link checker")
	resp, err := c.Do(req)
	if err != nil {
		return "link is unreachable: " + err.Error()
	}
	resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Sprintf("link answers HTTP %d", resp.StatusCode)
	}
	return ""
}

// changedLinks returns links with at least one URL, including version URLs,
// that is not present under baseRoot.
func changedLinks(l *list.List, baseRoot string) ([]list.Link, error) {
	base, err := list.Load(baseRoot)
	if err != nil {
		return nil, fmt.Errorf("base list: %w", err)
	}
	seen := map[string]bool{}
	for _, k := range base.Links {
		for _, u := range k.URLs() {
			seen[u] = true
		}
	}
	var changed []list.Link
	for _, k := range l.Links {
		for _, u := range k.URLs() {
			if !seen[u] {
				changed = append(changed, k)
				break
			}
		}
	}
	return changed, nil
}
