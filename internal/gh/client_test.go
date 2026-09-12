package gh

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("missing auth header")
		}
		switch r.URL.Path {
		case "/repos/old/name":
			http.Redirect(w, r, "/repos/new/name", http.StatusMovedPermanently)
		case "/repos/new/name":
			w.Write([]byte(`{"full_name":"new/name","stargazers_count":42,"language":"Go","license":{"spdx_id":"MIT","name":"MIT License"},"created_at":"2020-01-01T00:00:00Z","pushed_at":"2026-01-01T00:00:00Z"}`))
		case "/repos/new/name/contents/go.mod":
			w.WriteHeader(http.StatusOK)
		case "/repos/new/name/contents/missing":
			w.WriteHeader(http.StatusNotFound)
		case "/repos/limited/repo":
			w.Header().Set("X-RateLimit-Remaining", "0")
			w.Header().Set("X-RateLimit-Reset", "1900000000")
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, Token: "tok", HTTPClient: srv.Client()}
	ctx := context.Background()

	r, err := c.Repo(ctx, "old/name")
	if err != nil {
		t.Fatal(err)
	}
	if r.FullName != "new/name" || r.Stars != 42 || r.License.SPDXID != "MIT" || r.CreatedAt.Year() != 2020 {
		t.Fatalf("unexpected repo: %+v", r)
	}
	if _, err := c.Repo(ctx, "gone/gone"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
	if ok, err := c.FileExists(ctx, "new/name", "go.mod"); err != nil || !ok {
		t.Fatalf("go.mod should exist: %v %v", ok, err)
	}
	if ok, err := c.FileExists(ctx, "new/name", "missing"); err != nil || ok {
		t.Fatalf("missing should not exist: %v %v", ok, err)
	}
	var rl *RateLimitError
	if _, err := c.Repo(ctx, "limited/repo"); !errors.As(err, &rl) || rl.Reset.Unix() != 1900000000 {
		t.Fatalf("expected rate limit error, got %v", err)
	}
}
