package badge

import (
	"bytes"
	"encoding/xml"
	"fmt"
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
	for _, want := range []string{">awesome</text>", `aria-label="awesome go: gin &lt;&amp;&gt; co"`, "gin &lt;&amp;&gt; co", Color("gin-gonic/gin"), "<title>gin-gonic/gin is listed</title>"} {
		if !strings.Contains(s, want) {
			t.Errorf("badge missing %q", want)
		}
	}
	if !bytes.Equal(svg, SVG(opts)) {
		t.Errorf("badge output must be deterministic")
	}
}

func TestGopherIcon(t *testing.T) {
	if len(gopher) != iconSize {
		t.Fatalf("gopher must be %d rows, got %d", iconSize, len(gopher))
	}
	pixels := 0
	for i, row := range gopher {
		if len(row) != iconSize {
			t.Fatalf("gopher row %d must be %d pixels, got %d", i, iconSize, len(row))
		}
		for j := 0; j < len(row); j++ {
			if row[j] == '.' {
				continue
			}
			if _, ok := gopherPalette[row[j]]; !ok {
				t.Fatalf("gopher row %d uses %q, which is not in the palette", i, row[j])
			}
			pixels++
		}
	}

	s := string(SVG(Options{Name: "gin", Seed: "gin-gonic/gin", Title: "t"}))
	start := strings.Index(s, `<g shape-rendering="crispEdges">`)
	if start < 0 {
		t.Fatal("badge must draw the gopher in a crisp-edged group")
	}
	group := s[start:]
	group = group[:strings.Index(group, "</g>")]
	drawn := 0
	for _, m := range strings.Split(group, `<rect `)[1:] {
		var x, y, w int
		var fill string
		if _, err := fmt.Sscanf(m, `x="%d" y="%d" width="%d" height="1" fill=%s`, &x, &y, &w, &fill); err != nil {
			t.Fatalf("unexpected icon rect %q: %v", m, err)
		}
		if x < iconX || x+w > iconX+iconSize || y < iconY || y >= iconY+iconSize {
			t.Fatalf("icon rect outside the icon box: %q", m)
		}
		drawn += w
	}
	if drawn != pixels {
		t.Fatalf("icon draws %d pixels, the gopher has %d", drawn, pixels)
	}
	for _, c := range []string{"#93e6fd", "#f4d6b2", "#fff"} {
		if !strings.Contains(group, c) {
			t.Errorf("icon missing colour %s", c)
		}
	}

	label := strings.Index(s, `>awesome</text>`)
	xStart := strings.LastIndex(s[:label], `<text x="`) + len(`<text x="`)
	var cx int
	if _, err := fmt.Sscanf(s[xStart:], "%d", &cx); err != nil {
		t.Fatal(err)
	}
	if left := cx - textWidth(Label)/2; left < iconX+iconSize+4 {
		t.Errorf("label text starts at x=%d and overlaps the gopher, which ends at x=%d", left, iconX+iconSize)
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
