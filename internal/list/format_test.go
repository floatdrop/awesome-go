package list

import (
	"strings"
	"testing"
)

func TestFormatListSortsCategoriesAndSectionsByName(t *testing.T) {
	l := &List{
		Categories:   []Category{{ID: "web", Name: "Web Frameworks"}, {ID: "cli", Name: "Command Line"}, {ID: "audio", Name: "audio"}},
		LinkSections: []LinkSection{{ID: "videos", Name: "Videos"}, {ID: "books", Name: "Books"}},
	}
	out, err := l.FormatList()
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	for _, pair := range [][2]string{{`"audio"`, `"Command Line"`}, {`"Command Line"`, `"Web Frameworks"`}, {`"Books"`, `"Videos"`}} {
		if strings.Index(s, pair[0]) > strings.Index(s, pair[1]) {
			t.Errorf("%s must come before %s:\n%s", pair[0], pair[1], s)
		}
	}
	if strings.Index(s, `"Web Frameworks"`) > strings.Index(s, `"Books"`) {
		t.Errorf("categories must stay ahead of link sections:\n%s", s)
	}
	if l.Categories[0].ID != "web" {
		t.Errorf("FormatList must not reorder the caller's slice")
	}
}
