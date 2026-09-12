<!-- Generated from list.json and entries/ by `go run ./cmd/awesome generate`. Do not edit by hand. -->
# Awesome Go

> A short, hand-curated list of Go libraries and tools that are actually worth using. Not a directory of everything that exists.

24 projects in 13 categories. Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. Within each category the 5 most-starred projects are shown first and the rest are folded under More. Archived repositories are removed automatically. Stars were last refreshed on 2026-09-12.

Search, filter and sort the full list at [floatdrop.github.io/awesome-go](https://floatdrop.github.io/awesome-go/).

## Contents

- [Web Frameworks & Routers](#web-frameworks--routers)
- [Command Line](#command-line)
- [GUI & Desktop](#gui--desktop)
- [Configuration](#configuration)
- [Logging](#logging)
- [Testing](#testing)
- [Databases & SQL](#databases--sql)
- [Dependency Injection](#dependency-injection)
- [Concurrency](#concurrency)
- [Caching](#caching)
- [Networking & Protocols](#networking--protocols)
- [Messaging & Queues](#messaging--queues)
- [Developer Tools](#developer-tools)

## Web Frameworks & Routers

HTTP servers, routers and full-stack web frameworks.

- 🥇 [gin](https://github.com/gin-gonic/gin) - HTTP web framework with a martini-like API and strong routing performance. ★ 89k
- 🥈 [chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router built on net/http. ★ 23k

## Command Line

Argument parsing, terminal UIs and everything else for building CLIs.

- 🥇 [bubbletea](https://github.com/charmbracelet/bubbletea) - Framework for terminal user interfaces based on The Elm Architecture. ★ 45k
- 🥈 [cobra](https://github.com/spf13/cobra) - Framework for CLI applications with subcommands, flags, shell completions and generated docs. ★ 45k
- 🥉 [kong](https://github.com/alecthomas/kong) - Command-line parser where the whole grammar is a struct and commands are typed Run methods. ★ 3.2k

## GUI & Desktop

Desktop application frameworks and GUI toolkits.

- 🥇 [wails](https://github.com/wailsapp/wails) - Desktop applications with a Go backend and a web frontend, using the native webview. ★ 36k

## Configuration

Loading settings from files, environment variables and flags.

- 🥇 [koanf](https://github.com/knadh/koanf) - Configuration management with pluggable providers and parsers merged in explicit layers. ★ 4.2k
- 🥈 [confetti](https://github.com/yandex/confetti) - Configuration loader that chains sources per value or struct, with backends as plain functions. ★ 7

## Logging

Structured and leveled logging.

- 🥇 [log](https://github.com/charmbracelet/log) - Minimal, colorful logger with a slog handler and structured key-value output. ★ 3.4k

## Testing

Assertions, mocks, fixtures and integration test helpers.

- 🥇 [testify](https://github.com/stretchr/testify) - Assertions, mocks and suites that work with the standard testing package. ★ 26k
- 🥈 [go-cmp](https://github.com/google/go-cmp) - Equality comparison of arbitrary values for tests, with readable diffs and custom options. ★ 4.7k

## Databases & SQL

Drivers, query builders, ORMs, migrations and embedded stores.

- 🥇 [sqlc](https://github.com/sqlc-dev/sqlc) - Compiler that generates type-safe code from SQL queries. ★ 18k
- 🥈 [scan](https://github.com/blockloop/scan) - Map database/sql rows onto structs, slices and primitives using struct tags. ★ 615

## Dependency Injection

Wiring applications together.

- 🥇 [di](https://github.com/floatdrop/di) - Dependency injection container with typed constructors, lifecycle hooks and scopes. ★ 2

## Concurrency

Goroutine pools, structured concurrency, synchronization helpers and actor frameworks.

- 🥇 [conc](https://github.com/sourcegraph/conc) - Structured concurrency: scoped groups with panic propagation, bounded pools and ordered streams. ★ 10k
- 🥈 [ergo](https://github.com/ergo-services/ergo) - Actor framework with supervision trees and network transparency, modeled on Erlang/OTP. ★ 4.7k
- 🥉 [pond](https://github.com/alitto/pond) - Bounded worker pool with generics, result groups, context cancellation and subpools. ★ 2.2k

## Caching

In-memory caches and eviction policies.

- 🥇 [bigcache](https://github.com/allegro/bigcache) - Efficient in-memory cache for gigabytes of data without GC overhead. ★ 8.2k
- 🥈 [otter](https://github.com/maypok86/otter) - Concurrent in-memory cache tuned for high hit ratio and low contention. ★ 2.7k
- 🥉 [go-sieve](https://github.com/opencoff/go-sieve) - Generic in-memory cache with SIEVE eviction and a lock-free read path. ★ 41

## Networking & Protocols

Protocol implementations, transports and low-level network plumbing.

- 🥇 [moq-go](https://github.com/floatdrop/moq-go) - Media over QUIC transport library and relay tracking the IETF MoQ drafts. ★ 3

## Messaging & Queues

Message brokers, streaming platforms and task queues.

- 🥇 [nats.go](https://github.com/nats-io/nats.go) - Client for the NATS messaging system, including JetStream streams, key-value and object stores. ★ 6.7k

## Developer Tools

Linters, build tools, debuggers, code generators and terminal tools for daily development work.

- 🥇 [fzf](https://github.com/junegunn/fzf) - Interactive fuzzy finder for the command line, with shell integration for files and history. ★ 83k
- 🥈 [lazygit](https://github.com/jesseduffield/lazygit) - Terminal UI for git with staging, rebasing, stashing and worktrees a keystroke away. ★ 82k

## Contributing

Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work.
