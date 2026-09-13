package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/floatdrop/awesome-go/internal/list"
)

func TestCheckLinksChecksEveryVersion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/old" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	links := []list.Link{{
		Title: "Series",
		URL:   srv.URL + "/new",
		Versions: []list.LinkVersion{
			{Label: "2", URL: srv.URL + "/new"},
			{Label: "1", URL: srv.URL + "/old"},
		},
	}}
	problems := checkLinks(context.Background(), links)
	if len(problems) != 1 || !strings.Contains(problems[0].Msg, "/old: link answers HTTP 404") || strings.Contains(problems[0].Msg, "/new") {
		t.Fatalf("want only the broken version reported, got %v", problems)
	}
}

func TestCheckLinksSkipsExempt(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()
	links := []list.Link{
		{Title: "Blocked", URL: srv.URL + "/blocked"},
		{Title: "Waived", URL: srv.URL + "/waived", Exempt: []string{list.LinkExemptCheck}},
	}
	problems := checkLinks(context.Background(), links)
	if len(problems) != 1 || problems[0].File != "links/blocked.json" {
		t.Fatalf("want only the non-exempt link reported, got %v", problems)
	}
}
