# Contributing

Thanks for helping keep this list short and honest. The whole point of the list is that every entry earned its place, so the bar is deliberately high and most of it is enforced by CI rather than by arguing in comments.

## How the list works

- `entries/<owner>--<name>.json` is one file per listed project. That is the file you add, edit or delete.
- `list.json` holds the categories, the policy thresholds and the list metadata. It changes rarely.
- `README.md`, `metadata.json` and everything under `badges/` are generated. CI rejects pull requests that touch them.
- A nightly job refreshes star counts, removes repositories that were archived or deleted, and regenerates the README and badges. Nobody has to remember to clean up.
- Within each category the most-starred projects take the podium (🥇🥈🥉) and the rest are listed as contenders. Stars decide the order, humans decide who is on the list at all.

## Adding an entry

1. Make sure the project is not already listed and passes the [entry rules](#entry-rules).
2. Create `entries/<owner>--<name>.json`, lowercase, with the two slashes of the repository path replaced by `--`:

   ```json
   {
     "$schema": "../schema/entry.schema.json",
     "repo": "owner/name",
     "description": "What it does, in one plain sentence.",
     "category": "existing-category-id"
   }
   ```

   Category ids are in `list.json`. `name` is optional and defaults to the repository name. Do not set `added`; the bot fills it in on acceptance.

3. Run the tooling and fix what it reports:

   ```sh
   go run ./cmd/awesome fmt        # canonical formatting and file naming
   go run ./cmd/awesome validate   # offline rules
   GITHUB_TOKEN=$(gh auth token) go run ./cmd/awesome validate -remote   # optional: the GitHub checks CI will run
   ```

4. Open a pull request with **exactly one new entry** and a sentence or two on why the project is awesome from your own experience. CI rejects pull requests that add more than one file under `entries/`; splitting them makes each one reviewable.

Fully AI-generated pull requests are not accepted.

## Entry rules

Enforced by CI on every pull request. The numbers live in the `policy` section of `list.json`, so if they change the README and this document follow.

| Rule | Why |
| --- | --- |
| Repository is on GitHub, referenced as `owner/name` | The tooling needs one canonical identity to check stars, status and renames. |
| At least **90 days** old | Age filters out the weekend project that will be abandoned next weekend. |
| At least **200 stars** | A low bar that still means real people have found it useful. |
| Pushed to within the last **365 days** | Not required to be busy, just alive. |
| Has a license detected by GitHub | Unlicensed code cannot be used. |
| Primary language is Go, or there is a `go.mod` at the root | This is a Go list. |
| Not archived, not disabled, not a fork | Forks belong to the upstream entry; archived projects are removed automatically anyway. |
| Not already listed, including under an old name | Renamed repositories are followed through GitHub redirects. |

Description rules, also enforced:

- One sentence, at most 100 characters, starts with an uppercase letter, ends with a period.
- Says what the project does. No marketing ("blazing fast", "best", "powerful"), no leading article ("A", "The"), and no repeating the project name.
- Does not mention Go or Golang. Everything here is Go.
- Plain text: no Markdown, links or HTML.

Two rules can be waived by a maintainer with an explicit `"exempt"` list on the entry, because automation gets them wrong for good projects:

- `"fork"`: the repository is a fork that became the maintained successor of its upstream (for example `go-viper/mapstructure`).
- `"inactive"`: a finished library that has not needed a commit in a while but is still the right answer (for example `google/uuid`).

Age and star thresholds cannot be waived. Every exemption must be justified in the pull request, and the nightly job keeps listing inactive projects in its summary so they get looked at again.

Rules that need a human:

- The project must be something you would recommend to a colleague without caveats.
- The category must fit. Suggest a new category in the PR if none does, with at least three candidate entries for it.
- Deprecated, "maintenance mode" and thin wrappers around another listed project do not qualify even if they pass every automated check.

## Removing an entry

Archived and deleted repositories are removed automatically every night. Anything else (unmaintained, superseded, no longer recommendable) is removed through a pull request that deletes the entry and explains why. The nightly job also lists projects with no pushes in the policy window in its job summary so they can be reviewed.

## Badge

Every listed project gets its own badge: a small SVG with the project name in a colour of its own.

![badge example](badges/gin-gonic--gin.svg)

Embed it in your README:

```markdown
[![Awesome Go](https://raw.githubusercontent.com/floatdrop/awesome-go/main/badges/OWNER--NAME.svg)](https://github.com/floatdrop/awesome-go)
```

Replace `OWNER--NAME` with your repository slug in lowercase. The badge only exists while the project is listed; if the entry is removed the image disappears with the next sync.

## Tooling

Everything is stdlib Go, run through `go run ./cmd/awesome <command>`:

| Command | What it does |
| --- | --- |
| `validate` | Offline rules; `-remote` adds the GitHub checks, `-base DIR` limits them to entries not present under that directory. |
| `fmt` | Rewrites `list.json` and `entries/` in canonical form and fixes file names; `-check` only verifies. |
| `refresh` | Fetches stars, activity and status into `metadata.json`, follows renames, stamps `added` dates. |
| `prune` | Removes archived, disabled and deleted repositories; reports stale ones. |
| `generate` | Writes `README.md` and `badges/*.svg`. |
| `sync` | `refresh`, `prune` and `generate` in sequence; what the nightly job runs. |

Set `GITHUB_TOKEN` for anything that talks to GitHub. Unauthenticated requests are limited to 60 per hour.
