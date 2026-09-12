package badge

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
	"testing"
)

func TestSVGIsWellFormed(t *testing.T) {
	opts := Options{Name: "gin <&> co", Seed: "gin-gonic/gin", Title: "gin-gonic/gin is listed"}
	svg := SVG(opts)
	dec := xml.NewDecoder(bytes.NewReader(svg))
	for {
		if _, err := dec.Token(); err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("invalid XML: %v\n%s", err, svg)
		}
	}
	s := string(svg)
	for _, want := range []string{"awesome go", "gin &lt;&amp;&gt; co", Color("gin-gonic/gin"), "<title>gin-gonic/gin is listed</title>"} {
		if !strings.Contains(s, want) {
			t.Errorf("badge missing %q", want)
		}
	}
	if !bytes.Equal(svg, SVG(opts)) {
		t.Errorf("badge output must be deterministic")
	}
}

func TestColorIsStableAndVaried(t *testing.T) {
	if Color("gin-gonic/gin") != Color("GIN-GONIC/GIN") {
		t.Fatal("colour must not depend on case")
	}
	seen := map[string]bool{}
	for _, seed := range []string{"a/a", "b/b", "c/c", "d/d", "e/e", "f/f", "g/g", "h/h"} {
		seen[Color(seed)] = true
	}
	if len(seen) < 3 {
		t.Fatalf("expected several colours across seeds, got %d", len(seen))
	}
}
