package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/floatdrop/awesome-go/internal/list"
)

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
