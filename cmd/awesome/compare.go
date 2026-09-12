package main

import (
	"fmt"

	"github.com/floatdrop/awesome-go/internal/list"
	"github.com/floatdrop/awesome-go/internal/render"
)

// runCompare prints a Markdown comparison of new or changed entries against
// what their categories already hold. CI posts it on the pull request.
func runCompare(args []string) error {
	fs, p := newFlags("compare")
	base := fs.String("base", "", "directory holding the base branch's list.json and entries/; only entries not present there are compared")
	if err := fs.Parse(args); err != nil {
		return err
	}
	l, err := list.Load(p.root)
	if err != nil {
		return err
	}
	meta, err := list.LoadMetadata(p.metadata())
	if err != nil {
		return err
	}
	targets := l.Entries
	if *base != "" {
		if targets, err = changedEntries(l, *base); err != nil {
			return err
		}
	}
	fmt.Print(render.Compare(l, meta, targets))
	return nil
}
