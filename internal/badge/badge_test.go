package badge

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
	"testing"
)

func TestIDIsStableAndSpecific(t *testing.T) {
	a := ID("floatdrop/awesome-go", "gin-gonic/gin", "2026-09-12")
	b := ID("Floatdrop/Awesome-Go", "GIN-GONIC/GIN", "2026-09-12")
	c := ID("floatdrop/awesome-go", "gin-gonic/gin", "2026-09-13")
	if len(a) != 7 || a != b {
		t.Fatalf("id must be 7 chars and case-insensitive: %q %q", a, b)
	}
	if a == c {
		t.Fatalf("id must depend on the acceptance date")
	}
}

func TestSVGIsWellFormed(t *testing.T) {
	svg := SVG(Options{Name: "gin <&> co", ID: "abc1234", Title: "gin-gonic/gin is listed"})
	dec := xml.NewDecoder(bytes.NewReader(svg))
	for {
		if _, err := dec.Token(); err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("invalid XML: %v\n%s", err, svg)
		}
	}
	s := string(svg)
	for _, want := range []string{"awesome go", "gin &lt;&amp;&gt; co", "abc1234", Color("abc1234"), "<title>gin-gonic/gin is listed</title>"} {
		if !strings.Contains(s, want) {
			t.Errorf("badge missing %q", want)
		}
	}
	if !bytes.Equal(svg, SVG(Options{Name: "gin <&> co", ID: "abc1234", Title: "gin-gonic/gin is listed"})) {
		t.Errorf("badge output must be deterministic")
	}
}

func TestColorFromID(t *testing.T) {
	if Color("0000000") != palette[0] || Color("zzz") != palette[0] {
		t.Fatalf("unexpected colour mapping")
	}
	seen := map[string]bool{}
	for _, id := range []string{"0000001", "0000002", "0000003"} {
		seen[Color(id)] = true
	}
	if len(seen) != 3 {
		t.Fatalf("consecutive ids should map to different colours")
	}
}
