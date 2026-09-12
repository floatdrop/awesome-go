// Package badge generates the per-project "badge of proof": a flat SVG in the
// style of shields.io, with the project name, a colour derived from its id and
// the id itself, so anyone can verify the badge against badges/index.json.
package badge

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html"
	"strconv"
	"strings"
	"unicode"
)

// Label is the left-hand text of every badge.
const Label = "awesome go"

// ID derives a stable 7-hex-character identifier for a listing. It changes if
// the entry is removed and re-added later, which is the point: the id proves
// a specific act of acceptance.
func ID(listRepo, repo, since string) string {
	sum := sha256.Sum256([]byte(strings.ToLower(listRepo) + "\n" + strings.ToLower(repo) + "\n" + since))
	return hex.EncodeToString(sum[:])[:7]
}

var palette = []string{
	"#007ec6", // blue
	"#4c1",    // bright green
	"#97ca00", // green
	"#dfb317", // yellow
	"#fe7d37", // orange
	"#e05d44", // red
	"#8e44ad", // purple
	"#e91e63", // pink
	"#00b5ad", // teal
	"#2f80ed", // azure
}

// Color picks a palette colour from the id so every badge looks distinct.
func Color(id string) string {
	n, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		return palette[0]
	}
	return palette[n%uint64(len(palette))]
}

// Options describe a single badge.
type Options struct {
	Name  string // project name shown in the middle segment
	ID    string // proof id shown in the right segment
	Title string // accessible description embedded in the SVG
}

type segment struct {
	text  string
	color string
	width int
}

// SVG renders the badge.
func SVG(o Options) []byte {
	segs := []segment{
		{text: Label, color: "#555"},
		{text: o.Name, color: Color(o.ID)},
		{text: o.ID, color: "#2b2b2b"},
	}
	total := 0
	for i := range segs {
		segs[i].width = textWidth(segs[i].text) + 10
		total += segs[i].width
	}

	var b strings.Builder
	w := func(format string, args ...any) { fmt.Fprintf(&b, format, args...) }
	aria := html.EscapeString(fmt.Sprintf("%s: %s, id %s", Label, o.Name, o.ID))
	w(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="20" role="img" aria-label="%s">`, total, aria)
	w(`<title>%s</title>`, html.EscapeString(o.Title))
	w(`<linearGradient id="s" x2="0" y2="100%%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient>`)
	w(`<clipPath id="r"><rect width="%d" height="20" rx="3" fill="#fff"/></clipPath>`, total)
	w(`<g clip-path="url(#r)">`)
	x := 0
	for _, s := range segs {
		w(`<rect x="%d" width="%d" height="20" fill="%s"/>`, x, s.width, s.color)
		x += s.width
	}
	w(`<rect width="%d" height="20" fill="url(#s)"/></g>`, total)
	w(`<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="11">`)
	x = 0
	for _, s := range segs {
		cx := x + s.width/2
		t := html.EscapeString(s.text)
		w(`<text x="%d" y="15" fill="#010101" fill-opacity=".3">%s</text>`, cx, t)
		w(`<text x="%d" y="14">%s</text>`, cx, t)
		x += s.width
	}
	w(`</g></svg>`)
	b.WriteString("\n")
	return []byte(b.String())
}

// textWidth approximates Verdana 11px metrics; exact enough for a badge.
func textWidth(s string) int {
	total := 0.0
	for _, r := range s {
		switch {
		case strings.ContainsRune("ilj!|.,:;'", r):
			total += 3.5
		case strings.ContainsRune("tfrI-", r):
			total += 4.5
		case strings.ContainsRune("mwMW", r):
			total += 10
		case r == ' ':
			total += 4
		case unicode.IsUpper(r):
			total += 8
		case unicode.IsDigit(r):
			total += 7
		case unicode.IsLower(r):
			total += 6.6
		default:
			total += 7.5
		}
	}
	return int(total + 0.5)
}
