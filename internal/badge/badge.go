// Package badge generates the per-project badge: a flat SVG in the style of
// shields.io with an 8-bit gopher in sunglasses and "awesome" on the left and
// the project name on the right, coloured per project so no two badges look alike.
package badge

import (
	"crypto/sha256"
	"fmt"
	"html"
	"strings"
	"unicode"
)

// Label is the left-hand text of every badge. "Go" is carried by the gopher
// icon, so the visible text stays short.
const Label = "awesome"

// accessibleLabel names the badge for screen readers, which cannot see the
// gopher, so it keeps the full name of the list.
const accessibleLabel = "awesome go"

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

// Color picks a palette colour deterministically from a seed such as the
// repository name.
func Color(seed string) string {
	sum := sha256.Sum256([]byte(strings.ToLower(seed)))
	return palette[int(sum[0])%len(palette)]
}

// Options describe a single badge.
type Options struct {
	Name  string // project name shown in the right segment
	Seed  string // what the colour is derived from, normally the repository
	Title string // accessible description embedded in the SVG
}

type segment struct {
	text  string
	color string
	width int
	lead  int // space reserved before the text, for the gopher icon
}

// gopher is the 8-bit gopher in pixel sunglasses drawn at the left of every
// badge, so listed projects stand out among generic shields. Each byte is one
// pixel: '.' is transparent, anything else is a key of gopherPalette.
var gopher = [...]string{
	"..OO......OO..",
	".OTBOOOOOOBTO.",
	".OBBBBBBBBBBO.",
	"OKKKKKKKKKKKKO",
	"OBKWKKBBKWKKBO",
	"OBBKKKBBKKKBBO",
	"OBBBBBNNBBBBBO",
	"OBBBBTTTTBBBBO",
	"OBBBBBWWBBBBBO",
	"TBBBBBBBBBBBBT",
	"OBBBBBBBBBBBBO",
	"OBBBBBBBBBBBBO",
	".OBBBBBBBBBBO.",
	"..TO......OT..",
}

var gopherPalette = map[byte]string{
	'O': "#1d1d1d", // outline
	'B': "#93e6fd", // body, the blue of the list logo
	'K': "#000",    // sunglasses
	'W': "#fff",    // glint and teeth
	'T': "#f4d6b2", // ears, snout, hands and feet
	'N': "#111",    // nose
}

// Icon placement inside the 20px-tall badge.
const (
	iconX    = 5
	iconY    = 3
	iconSize = 14
	iconGap  = 0 // the label segment's own 5px text padding separates icon and text
)

// writeGopher draws the icon, merging horizontal runs of one colour into a
// single rect so every badge stays a few kilobytes.
func writeGopher(b *strings.Builder) {
	b.WriteString(`<g shape-rendering="crispEdges">`)
	for y, row := range gopher {
		for x := 0; x < len(row); {
			c, ok := gopherPalette[row[x]]
			run := 1
			for x+run < len(row) && row[x+run] == row[x] {
				run++
			}
			if ok {
				fmt.Fprintf(b, `<rect x="%d" y="%d" width="%d" height="1" fill="%s"/>`, iconX+x, iconY+y, run, c)
			}
			x += run
		}
	}
	b.WriteString(`</g>`)
}

// SVG renders the badge.
func SVG(o Options) []byte {
	segs := []segment{
		{text: Label, color: "#555", lead: iconX + iconSize + iconGap},
		{text: o.Name, color: Color(o.Seed)},
	}
	total := 0
	for i := range segs {
		segs[i].width = segs[i].lead + textWidth(segs[i].text) + 10
		total += segs[i].width
	}

	var b strings.Builder
	w := func(format string, args ...any) { fmt.Fprintf(&b, format, args...) }
	aria := html.EscapeString(fmt.Sprintf("%s: %s", accessibleLabel, o.Name))
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
	writeGopher(&b)
	w(`<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="11">`)
	x = 0
	for _, s := range segs {
		cx := x + s.lead + (s.width-s.lead)/2
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
