package render

import (
	"strings"
	"testing"
	"time"
)

func TestSite(t *testing.T) {
	l, m := fixture(3)
	l.Entries[0].Description = "Small <b>&</b> escaped."
	out, err := Site(l, m, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	for _, want := range []string{
		"<title>Awesome Go</title>",
		`<link rel="canonical" href="https://x.github.io/y/">`,
		`<option value="web">Web &amp; HTTP</option>`,
		`<section class="category" id="web" data-id="web">`,
		`id="b--big" data-name="big" data-stars="45678" data-added=""`,
		`<span class="rank">🥇</span>`,
		`<span class="rank">4</span>`,
		"Small &lt;b&gt;&amp;&lt;/b&gt; escaped.",
		"stars refreshed 2026-09-12",
		"4 projects in 1 categories",
	} {
		if !strings.Contains(s, want) {
			t.Errorf("site missing %q", want)
		}
	}
	if strings.Contains(s, `id="empty" data-id`) {
		t.Errorf("empty category must be skipped")
	}
	if strings.Index(s, `id="b--big"`) > strings.Index(s, `id="a--small"`) {
		t.Errorf("entries must be ordered by stars")
	}
}
