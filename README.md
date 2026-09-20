<!-- Generated from list.json and entries/ by `go run ./cmd/awesome generate`. Do not edit by hand. -->
<p align="center">
  <img src="assets/logotype.svg" alt="The Go gopher wearing pixel sunglasses" width="180">
</p>

<h1 align="center">Awesome Go</h1>

<p align="center">
  <a href="https://discord.com/channels/1548529967288029294"><img src="https://img.shields.io/discord/1548529967288029294?logo=discord&logoColor=white&label=discord&color=5865F2" alt="Discord"></a>
</p>

> A short, hand-curated list of Go libraries and tools that are actually worth using. Not a directory of everything that exists.

47 projects in 19 categories. Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. An entry earns its place by covering something no listed project does, or by doing it better. Within each category the 10 most-starred projects are shown first and the rest are folded under More. Archived repositories are removed automatically. Stars were last refreshed on 2026-09-20.

Search, filter and sort the full list at [floatdrop.github.io/awesome-go](https://floatdrop.github.io/awesome-go/).

## Contents

- [Audio & Video](#audio--video)
- [Caching](#caching)
- [Command Line](#command-line)
- [Concurrency](#concurrency)
- [Configuration](#configuration)
- [Data Structures](#data-structures)
- [Databases & SQL](#databases--sql)
- [Dependency Injection](#dependency-injection)
- [Developer Tools](#developer-tools)
- [Embedded & Hardware](#embedded--hardware)
- [GUI & Desktop](#gui--desktop)
- [Logging](#logging)
- [Messaging & Queues](#messaging--queues)
- [Networking & Protocols](#networking--protocols)
- [Observability](#observability)
- [Resilience](#resilience)
- [Testing](#testing)
- [Validation](#validation)
- [Web Frameworks & Routers](#web-frameworks--routers)
- [Blog Posts](#blog-posts)
- [Books](#books)
- [Documentation](#documentation)
- [Podcasts](#podcasts)
- [Videos](#videos)

## Audio & Video

Codecs, containers and media processing.

- <img src="badges/stars/pion--opus.svg" alt="560 stars" align="absmiddle"> [opus](https://github.com/pion/opus) - Implementation of the Opus audio codec per RFC 6716, with bitstream internals exported for analysis.
- <img src="badges/stars/eyevinn--hi264.svg" alt="22 stars" align="absmiddle"> [hi264](https://github.com/Eyevinn/hi264) - Decoder for H.264 IDR frames and generator of test bitstreams, with fMP4 fragment extension.
- <img src="badges/stars/tphakala--go-flac.svg" alt="1 stars" align="absmiddle"> [go-flac](https://github.com/tphakala/go-flac) - FLAC encoder and decoder with SIMD acceleration, no cgo, and output bit-exact with libFLAC.

## Caching

In-memory caches and eviction policies.

- <img src="badges/stars/allegro--bigcache.svg" alt="8.2k stars" align="absmiddle"> [bigcache](https://github.com/allegro/bigcache) - Efficient in-memory cache for gigabytes of data without GC overhead.
- <img src="badges/stars/maypok86--otter.svg" alt="2.7k stars" align="absmiddle"> [otter](https://github.com/maypok86/otter) - Concurrent in-memory cache tuned for high hit ratio and low contention.
- <img src="badges/stars/opencoff--go-sieve.svg" alt="41 stars" align="absmiddle"> [go-sieve](https://github.com/opencoff/go-sieve) - Generic in-memory cache with SIEVE eviction and a lock-free read path.

## Command Line

Argument parsing, terminal UIs and everything else for building CLIs.

- <img src="badges/stars/charmbracelet--bubbletea.svg" alt="45k stars" align="absmiddle"> [bubbletea](https://github.com/charmbracelet/bubbletea) - Framework for terminal user interfaces based on The Elm Architecture.
- <img src="badges/stars/spf13--cobra.svg" alt="45k stars" align="absmiddle"> [cobra](https://github.com/spf13/cobra) - Framework for CLI applications with subcommands, flags, shell completions and generated docs.
- <img src="badges/stars/alecthomas--kong.svg" alt="3.2k stars" align="absmiddle"> [kong](https://github.com/alecthomas/kong) - Command-line parser where the whole grammar is a struct and commands are typed Run methods.
- <img src="badges/stars/peterbourgon--ff.svg" alt="1.4k stars" align="absmiddle"> [ff](https://github.com/peterbourgon/ff) - Extends the standard flag package with env vars, config files and subcommands, without replacing it.

## Concurrency

Goroutine pools, structured concurrency, synchronization helpers and actor frameworks.

- <img src="badges/stars/ergo-services--ergo.svg" alt="4.7k stars" align="absmiddle"> [ergo](https://github.com/ergo-services/ergo) - Actor framework with supervision trees and network transparency, modeled on Erlang/OTP.
- <img src="badges/stars/alitto--pond.svg" alt="2.2k stars" align="absmiddle"> [pond](https://github.com/alitto/pond) - Bounded worker pool with generics, result groups, context cancellation and subpools.

## Configuration

Loading settings from files, environment variables and flags.

- <img src="badges/stars/knadh--koanf.svg" alt="4.2k stars" align="absmiddle"> [koanf](https://github.com/knadh/koanf) - Configuration management with pluggable providers and parsers merged in explicit layers.
- <img src="badges/stars/yandex--confetti.svg" alt="7 stars" align="absmiddle"> [confetti](https://github.com/yandex/confetti) - Configuration loader that chains sources per value or struct, with backends as plain functions.

## Data Structures

Bitsets, compressed bitmaps and other containers beyond slices and maps.

- <img src="badges/stars/roaringbitmap--roaring.svg" alt="2.9k stars" align="absmiddle"> [roaring](https://github.com/RoaringBitmap/roaring) - Compressed bitmaps for sparse or huge integer sets, with fast set operations and serialization.
- <img src="badges/stars/bits-and-blooms--bloom.svg" alt="2.8k stars" align="absmiddle"> [bloom](https://github.com/bits-and-blooms/bloom) - Probabilistic set membership with Bloom filters sized from a target false positive rate.
- <img src="badges/stars/bits-and-blooms--bitset.svg" alt="1.5k stars" align="absmiddle"> [bitset](https://github.com/bits-and-blooms/bitset) - Bitsets with set algebra, population counts, fast iteration over set bits and serialization.

## Databases & SQL

Drivers, query builders, ORMs, migrations, embedded and distributed stores.

- <img src="badges/stars/etcd-io--etcd.svg" alt="52k stars" align="absmiddle"> [etcd](https://github.com/etcd-io/etcd) - Distributed key-value store with Raft consensus, watches and leases, plus its client.
- <img src="badges/stars/redis--go-redis.svg" alt="22k stars" align="absmiddle"> [go-redis](https://github.com/redis/go-redis) - Redis client with cluster, sentinel and ring support, pipelines, pub/sub and OpenTelemetry hooks.
- <img src="badges/stars/sqlc-dev--sqlc.svg" alt="18k stars" align="absmiddle"> [sqlc](https://github.com/sqlc-dev/sqlc) - Compiler that generates type-safe code from SQL queries.
- <img src="badges/stars/jackc--pgx.svg" alt="14k stars" align="absmiddle"> [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit with a native API, connection pool, binary protocol and COPY support.
- <img src="badges/stars/blockloop--scan.svg" alt="615 stars" align="absmiddle"> [scan](https://github.com/blockloop/scan) - Map database/sql rows onto structs, slices and primitives using struct tags.
- <img src="badges/stars/duckdb--duckdb-go.svg" alt="297 stars" align="absmiddle"> [duckdb-go](https://github.com/duckdb/duckdb-go) - Official database/sql driver for DuckDB, the embedded analytical engine, with Appender and Arrow.
- <img src="badges/stars/modernc-org--sqlite.svg" alt="99 stars" align="absmiddle"> [modernc.org/sqlite](https://github.com/modernc-org/sqlite) - SQLite translated from C, so database/sql gets an embedded database without cgo or a C toolchain.

## Dependency Injection

Wiring applications together.

- <img src="badges/stars/floatdrop--di.svg" alt="3 stars" align="absmiddle"> [di](https://github.com/floatdrop/di) - Dependency injection container with typed constructors, lifecycle hooks and scopes.

## Developer Tools

Linters, build tools, debuggers, code generators and terminal tools for daily development work.

- <img src="badges/stars/junegunn--fzf.svg" alt="83k stars" align="absmiddle"> [fzf](https://github.com/junegunn/fzf) - Interactive fuzzy finder for the command line, with shell integration for files and history.
- <img src="badges/stars/jesseduffield--lazygit.svg" alt="83k stars" align="absmiddle"> [lazygit](https://github.com/jesseduffield/lazygit) - Terminal UI for git with staging, rebasing, stashing and worktrees a keystroke away.
- <img src="badges/stars/go-delve--delve.svg" alt="25k stars" align="absmiddle"> [delve](https://github.com/go-delve/delve) - Debugger with goroutine-aware breakpoints, core dump analysis and remote debugging over DAP.
- <img src="badges/stars/golangci--golangci-lint.svg" alt="19k stars" align="absmiddle"> [golangci-lint](https://github.com/golangci/golangci-lint) - Linters runner that aggregates dozens of linters with caching, parallelism and a single config.

## Embedded & Hardware

Compilers, frameworks and drivers for microcontrollers, boards and the devices attached to them.

- <img src="badges/stars/tinygo-org--tinygo.svg" alt="18k stars" align="absmiddle"> [tinygo](https://github.com/tinygo-org/tinygo) - Compiler for microcontrollers, WebAssembly and small binaries, built on LLVM.

## GUI & Desktop

Desktop application frameworks and GUI toolkits.

- <img src="badges/stars/wailsapp--wails.svg" alt="36k stars" align="absmiddle"> [wails](https://github.com/wailsapp/wails) - Desktop applications with a Go backend and a web frontend, using the native webview.

## Logging

Structured and leveled logging.

- <img src="badges/stars/charmbracelet--log.svg" alt="3.4k stars" align="absmiddle"> [log](https://github.com/charmbracelet/log) - Minimal, colorful logger with a slog handler and structured key-value output.

## Messaging & Queues

Message brokers, streaming platforms and task queues.

- <img src="badges/stars/nats-io--nats.go.svg" alt="6.8k stars" align="absmiddle"> [nats.go](https://github.com/nats-io/nats.go) - Client for the NATS messaging system, including JetStream streams, key-value and object stores.

## Networking & Protocols

Protocol implementations, transports and low-level network plumbing.

- <img src="badges/stars/pion--webrtc.svg" alt="17k stars" align="absmiddle"> [webrtc](https://github.com/pion/webrtc) - Pure implementation of the WebRTC stack: ICE, DTLS, SRTP, data channels and media tracks.
- <img src="badges/stars/floatdrop--moq-go.svg" alt="3 stars" align="absmiddle"> [moq-go](https://github.com/floatdrop/moq-go) - Media over QUIC transport library and relay tracking the IETF MoQ drafts.

## Observability

Metrics, tracing and instrumentation.

- <img src="badges/stars/open-telemetry--opentelemetry-go.svg" alt="6.6k stars" align="absmiddle"> [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go) - Vendor-neutral SDK for traces, metrics and logs with OTLP and Prometheus exporters.

## Resilience

Retries, backoff, circuit breakers and other patterns for surviving failing dependencies.

- <img src="badges/stars/cenkalti--backoff.svg" alt="4.1k stars" align="absmiddle"> [backoff](https://github.com/cenkalti/backoff) - Exponential backoff with jitter for retries, honoring context cancellation and permanent errors.
- <img src="badges/stars/sony--gobreaker.svg" alt="3.7k stars" align="absmiddle"> [gobreaker](https://github.com/sony/gobreaker) - Circuit breaker with configurable trip conditions, state change hooks and a generic two-step API.

## Testing

Assertions, mocks, fixtures and integration test helpers.

- <img src="badges/stars/stretchr--testify.svg" alt="26k stars" align="absmiddle"> [testify](https://github.com/stretchr/testify) - Assertions, mocks and suites that work with the standard testing package.
- <img src="badges/stars/testcontainers--testcontainers-go.svg" alt="5k stars" align="absmiddle"> [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - Throwaway Docker containers for integration tests, with ready modules for databases and brokers.
- <img src="badges/stars/google--go-cmp.svg" alt="4.7k stars" align="absmiddle"> [go-cmp](https://github.com/google/go-cmp) - Equality comparison of arbitrary values for tests, with readable diffs and custom options.

## Validation

Validating structs and inputs.

- <img src="badges/stars/go-playground--validator.svg" alt="20k stars" align="absmiddle"> [validator](https://github.com/go-playground/validator) - Struct and field validation using tags, with cross-field rules, custom validators and translations.
- <img src="badges/stars/go-ozzo--ozzo-validation.svg" alt="4.1k stars" align="absmiddle"> [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Validation rules written as code rather than tags, with conditional rules and field-keyed errors.
- <img src="badges/stars/santhosh-tekuri--jsonschema.svg" alt="1.3k stars" align="absmiddle"> [jsonschema](https://github.com/santhosh-tekuri/jsonschema) - Validation against JSON Schema drafts 4 through 2020-12, with detailed errors and custom formats.
- <img src="badges/stars/oudwins--zog.svg" alt="1.2k stars" align="absmiddle"> [zog](https://github.com/Oudwins/zog) - Schema builder that parses untyped input into structs and validates it in one step, like Zod.

## Web Frameworks & Routers

HTTP servers, routers and full-stack web frameworks.

- <img src="badges/stars/gin-gonic--gin.svg" alt="89k stars" align="absmiddle"> [gin](https://github.com/gin-gonic/gin) - HTTP web framework with a martini-like API and strong routing performance.
- <img src="badges/stars/go-chi--chi.svg" alt="23k stars" align="absmiddle"> [chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router built on net/http.

## Blog Posts

Articles that changed how people write Go at work.

- [Go for Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/) - Peter Bourgon on structuring team codebases: configuration, dependencies, logging and testing.

## Books

Books worth reading cover to cover.

- [100 Go Mistakes and How to Avoid Them](https://100go.co/book/) - Teiva Harsanyi's catalogue of common pitfalls, from slices and maps to concurrency and testing.
- [Learning Go, 2nd Edition](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/) - Jon Bodner's guide to idiomatic modern practice, from generics to concurrency and testing.

## Documentation

Official guides and references for learning and using the language.

- [A Tour of Go](https://go.dev/tour/) - Interactive introduction that runs in the browser, from basic syntax to generics and concurrency.
- [Effective Go](https://go.dev/doc/effective_go) - Official guide to clear, idiomatic code: naming, formatting, interfaces, errors and concurrency.
- [Go Recipes](https://github.com/nikolaydubina/go-recipes/blob/main/README.md) - Cookbook of tool recipes for testing, dependencies, code generation, profiling and static analysis.
- [Go Release Interactive Tours](https://victoriametrics.com/blog/go-1-27/) - Runnable examples of what changed in each release, from language features to the standard library. Versions: [1.27](https://victoriametrics.com/blog/go-1-27/) · [1.26](https://antonz.org/go-1-26/) · [1.25](https://antonz.org/go-1-25/) · [1.24](https://antonz.org/go-1-24/) · [1.23](https://antonz.org/go-1-23/) · [1.22](https://antonz.org/go-1-22/)

## Podcasts

Shows worth a spot in your queue.

- [Cup o' Go](https://cupogo.dev/) - Weekly fifteen-minute rundown of news, releases and proposals with Jonathan Hall and Shay Nehmad.

## Videos

Talks and courses worth watching from start to finish.

- [Go Class by Matt KØDVB](https://www.youtube.com/playlist?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6) - Matt Holiday's university-style lecture course, from basic types to concurrency and testing.

## Contributing

Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work. The logo is excluded: it is based on the Go gopher by Renée French, licensed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
