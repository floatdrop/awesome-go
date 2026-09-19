<!-- Generated from list.json and entries/ by `go run ./cmd/awesome generate`. Do not edit by hand. -->
<p align="center">
  <img src="assets/logotype.svg" alt="The Go gopher wearing pixel sunglasses" width="180">
</p>

<h1 align="center">Awesome Go</h1>

<p align="center">
  <a href="https://discord.com/channels/1548529967288029294"><img src="https://img.shields.io/discord/1548529967288029294?logo=discord&logoColor=white&label=discord&color=5865F2" alt="Discord"></a>
</p>

> A short, hand-curated list of Go libraries and tools that are actually worth using. Not a directory of everything that exists.

46 projects in 17 categories. Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. An entry earns its place by covering something no listed project does, or by doing it better. Within each category the 5 most-starred projects are shown first and the rest are folded under More. Archived repositories are removed automatically. Stars were last refreshed on 2026-09-19.

Search, filter and sort the full list at [floatdrop.github.io/awesome-go](https://floatdrop.github.io/awesome-go/).

## Contents

- [Web Frameworks & Routers](#web-frameworks--routers)
- [Command Line](#command-line)
- [GUI & Desktop](#gui--desktop)
- [Configuration](#configuration)
- [Logging](#logging)
- [Testing](#testing)
- [Databases & SQL](#databases--sql)
- [Validation](#validation)
- [Dependency Injection](#dependency-injection)
- [Concurrency](#concurrency)
- [Resilience](#resilience)
- [Caching](#caching)
- [Networking & Protocols](#networking--protocols)
- [Audio & Video](#audio--video)
- [Messaging & Queues](#messaging--queues)
- [Data Structures](#data-structures)
- [Developer Tools](#developer-tools)
- [Documentation](#documentation)
- [Books](#books)
- [Videos](#videos)

## Web Frameworks & Routers

HTTP servers, routers and full-stack web frameworks.

- 🥇 [gin](https://github.com/gin-gonic/gin) - HTTP web framework with a martini-like API and strong routing performance. ★&nbsp;89k
- 🥈 [chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router built on net/http. ★&nbsp;23k

## Command Line

Argument parsing, terminal UIs and everything else for building CLIs.

- 🥇 [bubbletea](https://github.com/charmbracelet/bubbletea) - Framework for terminal user interfaces based on The Elm Architecture. ★&nbsp;45k
- 🥈 [cobra](https://github.com/spf13/cobra) - Framework for CLI applications with subcommands, flags, shell completions and generated docs. ★&nbsp;45k
- 🥉 [kong](https://github.com/alecthomas/kong) - Command-line parser where the whole grammar is a struct and commands are typed Run methods. ★&nbsp;3.2k
- &emsp;&#8196; [ff](https://github.com/peterbourgon/ff) - Extends the standard flag package with env vars, config files and subcommands, without replacing it. ★&nbsp;1.4k

## GUI & Desktop

Desktop application frameworks and GUI toolkits.

- 🥇 [wails](https://github.com/wailsapp/wails) - Desktop applications with a Go backend and a web frontend, using the native webview. ★&nbsp;36k

## Configuration

Loading settings from files, environment variables and flags.

- 🥇 [koanf](https://github.com/knadh/koanf) - Configuration management with pluggable providers and parsers merged in explicit layers. ★&nbsp;4.2k
- 🥈 [confetti](https://github.com/yandex/confetti) - Configuration loader that chains sources per value or struct, with backends as plain functions. ★&nbsp;7

## Logging

Structured and leveled logging.

- 🥇 [log](https://github.com/charmbracelet/log) - Minimal, colorful logger with a slog handler and structured key-value output. ★&nbsp;3.4k

## Testing

Assertions, mocks, fixtures and integration test helpers.

- 🥇 [testify](https://github.com/stretchr/testify) - Assertions, mocks and suites that work with the standard testing package. ★&nbsp;26k
- 🥈 [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - Throwaway Docker containers for integration tests, with ready modules for databases and brokers. ★&nbsp;5k
- 🥉 [go-cmp](https://github.com/google/go-cmp) - Equality comparison of arbitrary values for tests, with readable diffs and custom options. ★&nbsp;4.7k

## Databases & SQL

Drivers, query builders, ORMs, migrations, embedded and distributed stores.

- 🥇 [etcd](https://github.com/etcd-io/etcd) - Distributed key-value store with Raft consensus, watches and leases, plus its client. ★&nbsp;52k
- 🥈 [go-redis](https://github.com/redis/go-redis) - Redis client with cluster, sentinel and ring support, pipelines, pub/sub and OpenTelemetry hooks. ★&nbsp;22k
- 🥉 [sqlc](https://github.com/sqlc-dev/sqlc) - Compiler that generates type-safe code from SQL queries. ★&nbsp;18k
- &emsp;&#8196; [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit with a native API, connection pool, binary protocol and COPY support. ★&nbsp;14k
- &emsp;&#8196; [scan](https://github.com/blockloop/scan) - Map database/sql rows onto structs, slices and primitives using struct tags. ★&nbsp;615

<details>
<summary>More (1)</summary>

- [modernc.org/sqlite](https://github.com/modernc-org/sqlite) - SQLite translated from C, so database/sql gets an embedded database without cgo or a C toolchain. ★&nbsp;99

</details>

## Validation

Validating structs and inputs.

- 🥇 [validator](https://github.com/go-playground/validator) - Struct and field validation using tags, with cross-field rules, custom validators and translations. ★&nbsp;20k
- 🥈 [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Validation rules written as code rather than tags, with conditional rules and field-keyed errors. ★&nbsp;4.1k
- 🥉 [jsonschema](https://github.com/santhosh-tekuri/jsonschema) - Validation against JSON Schema drafts 4 through 2020-12, with detailed errors and custom formats. ★&nbsp;1.3k
- &emsp;&#8196; [zog](https://github.com/Oudwins/zog) - Schema builder that parses untyped input into structs and validates it in one step, like Zod. ★&nbsp;1.2k

## Dependency Injection

Wiring applications together.

- 🥇 [di](https://github.com/floatdrop/di) - Dependency injection container with typed constructors, lifecycle hooks and scopes. ★&nbsp;3

## Concurrency

Goroutine pools, structured concurrency, synchronization helpers and actor frameworks.

- 🥇 [conc](https://github.com/sourcegraph/conc) - Structured concurrency: scoped groups with panic propagation, bounded pools and ordered streams. ★&nbsp;10k
- 🥈 [ergo](https://github.com/ergo-services/ergo) - Actor framework with supervision trees and network transparency, modeled on Erlang/OTP. ★&nbsp;4.7k
- 🥉 [pond](https://github.com/alitto/pond) - Bounded worker pool with generics, result groups, context cancellation and subpools. ★&nbsp;2.2k
- &emsp;&#8196; [suture](https://github.com/thejerf/suture) - Supervisor trees that restart failed long-running services with backoff, modeled on Erlang/OTP. ★&nbsp;1.4k

## Resilience

Retries, backoff, circuit breakers and other patterns for surviving failing dependencies.

- 🥇 [backoff](https://github.com/cenkalti/backoff) - Exponential backoff with jitter for retries, honoring context cancellation and permanent errors. ★&nbsp;4.1k
- 🥈 [gobreaker](https://github.com/sony/gobreaker) - Circuit breaker with configurable trip conditions, state change hooks and a generic two-step API. ★&nbsp;3.7k

## Caching

In-memory caches and eviction policies.

- 🥇 [bigcache](https://github.com/allegro/bigcache) - Efficient in-memory cache for gigabytes of data without GC overhead. ★&nbsp;8.2k
- 🥈 [otter](https://github.com/maypok86/otter) - Concurrent in-memory cache tuned for high hit ratio and low contention. ★&nbsp;2.7k
- 🥉 [go-sieve](https://github.com/opencoff/go-sieve) - Generic in-memory cache with SIEVE eviction and a lock-free read path. ★&nbsp;41

## Networking & Protocols

Protocol implementations, transports and low-level network plumbing.

- 🥇 [webrtc](https://github.com/pion/webrtc) - Pure implementation of the WebRTC stack: ICE, DTLS, SRTP, data channels and media tracks. ★&nbsp;17k
- 🥈 [moq-go](https://github.com/floatdrop/moq-go) - Media over QUIC transport library and relay tracking the IETF MoQ drafts. ★&nbsp;3

## Audio & Video

Codecs, containers and media processing.

- 🥇 [opus](https://github.com/pion/opus) - Implementation of the Opus audio codec per RFC 6716, with bitstream internals exported for analysis. ★&nbsp;560
- 🥈 [hi264](https://github.com/Eyevinn/hi264) - Decoder for H.264 IDR frames and generator of test bitstreams, with fMP4 fragment extension. ★&nbsp;22
- 🥉 [go-flac](https://github.com/tphakala/go-flac) - FLAC encoder and decoder with SIMD acceleration, no cgo, and output bit-exact with libFLAC. ★&nbsp;1

## Messaging & Queues

Message brokers, streaming platforms and task queues.

- 🥇 [nats.go](https://github.com/nats-io/nats.go) - Client for the NATS messaging system, including JetStream streams, key-value and object stores. ★&nbsp;6.8k

## Data Structures

Bitsets, compressed bitmaps and other containers beyond slices and maps.

- 🥇 [roaring](https://github.com/RoaringBitmap/roaring) - Compressed bitmaps for sparse or huge integer sets, with fast set operations and serialization. ★&nbsp;2.9k
- 🥈 [bloom](https://github.com/bits-and-blooms/bloom) - Probabilistic set membership with Bloom filters sized from a target false positive rate. ★&nbsp;2.8k
- 🥉 [bitset](https://github.com/bits-and-blooms/bitset) - Bitsets with set algebra, population counts, fast iteration over set bits and serialization. ★&nbsp;1.5k

## Developer Tools

Linters, build tools, debuggers, code generators and terminal tools for daily development work.

- 🥇 [fzf](https://github.com/junegunn/fzf) - Interactive fuzzy finder for the command line, with shell integration for files and history. ★&nbsp;83k
- 🥈 [lazygit](https://github.com/jesseduffield/lazygit) - Terminal UI for git with staging, rebasing, stashing and worktrees a keystroke away. ★&nbsp;82k
- 🥉 [delve](https://github.com/go-delve/delve) - Debugger with goroutine-aware breakpoints, core dump analysis and remote debugging over DAP. ★&nbsp;25k
- &emsp;&#8196; [golangci-lint](https://github.com/golangci/golangci-lint) - Linters runner that aggregates dozens of linters with caching, parallelism and a single config. ★&nbsp;19k

## Documentation

Official guides and references for learning and using the language.

- [A Tour of Go](https://go.dev/tour/) - Interactive introduction that runs in the browser, from basic syntax to generics and concurrency.
- [Effective Go](https://go.dev/doc/effective_go) - Official guide to clear, idiomatic code: naming, formatting, interfaces, errors and concurrency.
- [Go Recipes](https://github.com/nikolaydubina/go-recipes/blob/main/README.md) - Cookbook of tool recipes for testing, dependencies, code generation, profiling and static analysis.
- [Go Release Interactive Tours](https://victoriametrics.com/blog/go-1-27/) - Runnable examples of what changed in each release, from language features to the standard library. Versions: [1.27](https://victoriametrics.com/blog/go-1-27/) · [1.26](https://antonz.org/go-1-26/) · [1.25](https://antonz.org/go-1-25/) · [1.24](https://antonz.org/go-1-24/) · [1.23](https://antonz.org/go-1-23/) · [1.22](https://antonz.org/go-1-22/)

## Books

Books worth reading cover to cover.

- [Learning Go, 2nd Edition](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/) - Jon Bodner's guide to idiomatic modern practice, from generics to concurrency and testing.

## Videos

Talks and courses worth watching from start to finish.

- [Go Class by Matt KØDVB](https://www.youtube.com/playlist?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6) - Matt Holiday's university-style lecture course, from basic types to concurrency and testing.

## Contributing

Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work. The logo is excluded: it is based on the Go gopher by Renée French, licensed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
