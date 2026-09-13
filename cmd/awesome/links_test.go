package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckLink(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.WriteHeader(http.StatusOK)
		case "/moved":
			http.Redirect(w, r, "/ok", http.StatusMovedPermanently)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	ctx := context.Background()

	if msg := checkLink(ctx, srv.Client(), srv.URL+"/ok"); msg != "" {
		t.Fatalf("ok link reported: %s", msg)
	}
	if msg := checkLink(ctx, srv.Client(), srv.URL+"/moved"); msg != "" {
		t.Fatalf("redirects must be followed: %s", msg)
	}
	if msg := checkLink(ctx, srv.Client(), srv.URL+"/gone"); !strings.Contains(msg, "HTTP 404") {
		t.Fatalf("want HTTP 404, got %q", msg)
	}
	if msg := checkLink(ctx, srv.Client(), "http://127.0.0.1:1/unreachable"); !strings.Contains(msg, "unreachable") {
		t.Fatalf("want unreachable, got %q", msg)
	}
}
