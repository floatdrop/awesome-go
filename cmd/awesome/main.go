// Command awesome maintains the list: validates awesome.json, refreshes GitHub
// metadata, prunes dead repositories and generates README.md and badges.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/floatdrop/awesome-go/internal/gh"
	"github.com/floatdrop/awesome-go/internal/list"
)

const usage = `usage: awesome <command> [flags]

commands:
  validate   check awesome.json against the entry rules (-remote also asks GitHub)
  fmt        rewrite awesome.json in canonical form (-check only verifies)
  refresh    fetch stars and status of every repository into metadata.json
  prune      remove archived, disabled and deleted repositories
  generate   write README.md and badges/ from awesome.json and metadata.json
  sync       refresh, prune and generate in one go (what the nightly job runs)

flags common to all commands:
  -root DIR  repository root (default ".")
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	cmd, args := os.Args[1], os.Args[2:]
	var err error
	switch cmd {
	case "validate":
		err = runValidate(args)
	case "fmt":
		err = runFmt(args)
	case "refresh":
		err = runRefresh(args)
	case "prune":
		err = runPrune(args)
	case "generate":
		err = runGenerate(args)
	case "sync":
		err = runSync(args)
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n%s", cmd, usage)
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// paths resolves the well-known files under the repository root.
type paths struct{ root string }

func (p paths) list() string     { return filepath.Join(p.root, list.FileName) }
func (p paths) metadata() string { return filepath.Join(p.root, list.MetadataFileName) }
func (p paths) readme() string   { return filepath.Join(p.root, "README.md") }
func (p paths) badges() string   { return filepath.Join(p.root, "badges") }

func newFlags(name string) (*flag.FlagSet, *paths) {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	p := &paths{}
	fs.StringVar(&p.root, "root", ".", "repository root")
	return fs, p
}

func runValidate(args []string) error {
	fs, p := newFlags("validate")
	remote := fs.Bool("remote", false, "also check repositories against GitHub")
	base := fs.String("base", "", "with -remote: awesome.json of the base branch; only new or changed entries are checked")
	if err := fs.Parse(args); err != nil {
		return err
	}
	now := time.Now().UTC()
	l, err := list.Load(p.list())
	if err != nil {
		return err
	}
	problems := l.Validate(now)

	if *remote && len(problems) == 0 {
		targets := l.Entries
		if *base != "" {
			targets, err = changedEntries(l, *base)
			if err != nil {
				return err
			}
		}
		fmt.Printf("checking %d repositories against GitHub\n", len(targets))
		client := gh.New(os.Getenv("GITHUB_TOKEN"))
		for _, r := range fetchAll(context.Background(), client, targets, now) {
			if r.err != nil {
				return r.err
			}
			if r.canonical != "" && r.canonical != r.entry.Repo {
				problems = append(problems, fmt.Sprintf("%s: repository has moved to %s, use the new name", r.entry.Repo, r.canonical))
			}
			for _, msg := range l.Policy.Check(r.meta, now, r.entry) {
				problems = append(problems, r.entry.Repo+": "+msg)
			}
		}
	}

	if len(problems) > 0 {
		for _, msg := range problems {
			fmt.Printf("✗ %s\n", msg)
			annotate("error", msg)
		}
		return fmt.Errorf("%d problem(s) found", len(problems))
	}
	fmt.Printf("✓ %s: %d entries in %d categories\n", list.FileName, len(l.Entries), len(l.Categories))
	return nil
}

// changedEntries returns entries that do not exist verbatim in the base file.
func changedEntries(l *list.List, basePath string) ([]list.Entry, error) {
	base, err := list.Load(basePath)
	if err != nil {
		return nil, fmt.Errorf("base list: %w", err)
	}
	seen := map[string]list.Entry{}
	for _, e := range base.Entries {
		seen[strings.ToLower(e.Repo)] = e
	}
	var changed []list.Entry
	for _, e := range l.Entries {
		if old, ok := seen[strings.ToLower(e.Repo)]; !ok || old.Repo != e.Repo {
			changed = append(changed, e)
		}
	}
	return changed, nil
}

func runFmt(args []string) error {
	fs, p := newFlags("fmt")
	check := fs.Bool("check", false, "exit non-zero if the file is not formatted instead of rewriting it")
	if err := fs.Parse(args); err != nil {
		return err
	}
	original, err := os.ReadFile(p.list())
	if err != nil {
		return err
	}
	l, err := list.Parse(original)
	if err != nil {
		return fmt.Errorf("%s: %w", p.list(), err)
	}
	formatted, err := l.Format()
	if err != nil {
		return err
	}
	if string(original) == string(formatted) {
		fmt.Printf("✓ %s is formatted\n", list.FileName)
		return nil
	}
	if *check {
		msg := list.FileName + " is not in canonical form; run `go run ./cmd/awesome fmt` and commit the result"
		annotate("error", msg)
		return errors.New(msg)
	}
	if err := os.WriteFile(p.list(), formatted, 0o644); err != nil {
		return err
	}
	fmt.Printf("formatted %s\n", list.FileName)
	return nil
}

func runSync(args []string) error {
	for _, step := range []func([]string) error{runRefresh, runPrune, runGenerate} {
		if err := step(args); err != nil {
			return err
		}
	}
	return nil
}

// annotate emits a GitHub Actions workflow command so problems show up inline
// in the pull request. Outside Actions it is silent.
func annotate(level, msg string) {
	if os.Getenv("GITHUB_ACTIONS") != "true" {
		return
	}
	fmt.Printf("::%s file=%s::%s\n", level, list.FileName, msg)
}

// summary appends Markdown to the GitHub Actions job summary when available
// and always echoes it to stdout.
func summary(lines ...string) {
	text := strings.Join(lines, "\n") + "\n"
	fmt.Print(text)
	if path := os.Getenv("GITHUB_STEP_SUMMARY"); path != "" {
		if f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644); err == nil {
			_, _ = f.WriteString(text)
			_ = f.Close()
		}
	}
}
