<!-- Generated from list.json and entries/ by `go run ./cmd/awesome generate`. Do not edit by hand. -->
# Awesome Go

> A short, hand-curated list of Go libraries and tools that are actually worth using. Not a directory of everything that exists.

99 projects in 19 categories. Every entry is added by a human through a pull request and checked against the [entry rules](CONTRIBUTING.md#entry-rules) by CI. Within each category the 5 most-starred projects are shown first and the rest are folded under More. Archived repositories are removed automatically. Stars were last refreshed on 2026-09-12.

## Contents

- [Web Frameworks & Routers](#web-frameworks--routers)
- [Command Line](#command-line)
- [Configuration](#configuration)
- [Logging](#logging)
- [Testing](#testing)
- [Databases & SQL](#databases--sql)
- [Serialization](#serialization)
- [Validation](#validation)
- [Dependency Injection](#dependency-injection)
- [Concurrency](#concurrency)
- [Caching](#caching)
- [HTTP Clients](#http-clients)
- [Messaging & Queues](#messaging--queues)
- [RPC & API](#rpc--api)
- [Authentication & Authorization](#authentication--authorization)
- [Observability](#observability)
- [Text, HTML & Templates](#text-html--templates)
- [Utilities](#utilities)
- [Developer Tools](#developer-tools)

## Web Frameworks & Routers

HTTP servers, routers and full-stack web frameworks.

- 🥇 [gin](https://github.com/gin-gonic/gin) - HTTP web framework with a martini-like API and strong routing performance. ★ 89k
- 🥈 [fiber](https://github.com/gofiber/fiber) - Express-inspired web framework built on top of Fasthttp. ★ 40k
- 🥉 [echo](https://github.com/labstack/echo) - Minimalist web framework with a focus on performance and extensibility. ★ 33k
- &emsp;&thinsp; [fasthttp](https://github.com/valyala/fasthttp) - HTTP implementation tuned for high-throughput workloads. ★ 23k
- &emsp;&thinsp; [chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router built on net/http. ★ 23k

<details>
<summary>More (1)</summary>

- [mux](https://github.com/gorilla/mux) - Request router and dispatcher for matching incoming requests to their handlers. ★ 22k

</details>

## Command Line

Argument parsing, terminal UIs and everything else for building CLIs.

- 🥇 [bubbletea](https://github.com/charmbracelet/bubbletea) - Framework for terminal user interfaces based on The Elm Architecture. ★ 45k
- 🥈 [cobra](https://github.com/spf13/cobra) - Framework for CLI applications with subcommands, flags and shell completions. ★ 45k
- 🥉 [cli](https://github.com/urfave/cli) - Declarative library for building command line applications. ★ 24k
- &emsp;&thinsp; [lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for nice terminal layouts. ★ 12k
- &emsp;&thinsp; [progressbar](https://github.com/schollz/progressbar) - Thread-safe progress bar for terminal applications. ★ 4.7k

<details>
<summary>More (2)</summary>

- [kong](https://github.com/alecthomas/kong) - Command-line parser driven by struct tags. ★ 3.2k
- [pflag](https://github.com/spf13/pflag) - Drop-in replacement for the flag package implementing POSIX and GNU style flags. ★ 2.8k

</details>

## Configuration

Loading settings from files, environment variables and flags.

- 🥇 [viper](https://github.com/spf13/viper) - Complete configuration solution with files, environment variables, flags and remote stores. ★ 30k
- 🥈 [godotenv](https://github.com/joho/godotenv) - Load environment variables from .env files. ★ 11k
- 🥉 [env](https://github.com/caarlos0/env) - Parse environment variables into structs using tags. ★ 6.3k
- &emsp;&thinsp; [koanf](https://github.com/knadh/koanf) - Lightweight, extensible configuration management with pluggable providers and parsers. ★ 4.2k

## Logging

Structured and leveled logging.

- 🥇 [zap](https://github.com/uber-go/zap) - Structured, leveled logging with a focus on performance. ★ 25k
- 🥈 [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. ★ 13k
- 🥉 [tint](https://github.com/lmittmann/tint) - Colorized slog handler for human-readable terminal output. ★ 1.4k

## Testing

Assertions, mocks, fixtures and integration test helpers.

- 🥇 [testify](https://github.com/stretchr/testify) - Assertions, mocks and suites that work with the standard testing package. ★ 26k
- 🥈 [ginkgo](https://github.com/onsi/ginkgo) - BDD-style testing framework paired with the Gomega matcher library. ★ 9.1k
- 🥉 [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - SQL driver mock for testing database interactions without a real database. ★ 6.6k
- &emsp;&thinsp; [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - Throwaway Docker containers for integration tests. ★ 5k
- &emsp;&thinsp; [mock](https://github.com/uber-go/mock) - Mocking framework with code generation via mockgen. ★ 3.4k

## Databases & SQL

Drivers, query builders, ORMs, migrations and embedded stores.

- 🥇 [gorm](https://github.com/go-gorm/gorm) - Developer-friendly ORM with associations, hooks, transactions and migrations. ★ 40k
- 🥈 [go-redis](https://github.com/redis/go-redis) - Redis client supporting clusters, sentinels, pipelines and pub/sub. ★ 22k
- 🥉 [migrate](https://github.com/golang-migrate/migrate) - Database migrations as a CLI and library with many drivers. ★ 19k
- &emsp;&thinsp; [sqlc](https://github.com/sqlc-dev/sqlc) - Compiler that generates type-safe code from SQL queries. ★ 18k
- &emsp;&thinsp; [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql for scanning rows into structs and named queries. ★ 18k

<details>
<summary>More (9)</summary>

- [ent](https://github.com/ent/ent) - Entity framework for modeling data as a graph schema with generated code. ★ 17k
- [badger](https://github.com/dgraph-io/badger) - Embeddable, persistent key-value store built on LSM trees. ★ 16k
- [mysql](https://github.com/go-sql-driver/mysql) - Driver for MySQL over database/sql. ★ 15k
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit with native protocol support. ★ 14k
- [goose](https://github.com/pressly/goose) - Database migration tool supporting SQL and Go migrations. ★ 11k
- [bbolt](https://github.com/etcd-io/bbolt) - Embedded key-value database based on B+ trees. ★ 9.7k
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for database/sql based on cgo. ★ 9.2k
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver. ★ 8.5k
- [bun](https://github.com/uptrace/bun) - SQL-first ORM for PostgreSQL, MySQL, MSSQL and SQLite. ★ 5k

</details>

## Serialization

JSON, YAML, TOML, MessagePack and friends.

- 🥇 [gjson](https://github.com/tidwall/gjson) - JSON value lookup with a single-line path syntax. ★ 16k
- 🥈 [sonic](https://github.com/bytedance/sonic) - JSON serialization and deserialization accelerated with JIT and SIMD. ★ 9.6k
- 🥉 [toml](https://github.com/BurntSushi/toml) - Parser and encoder for TOML with reflection-based mapping to structs. ★ 5k
- &emsp;&thinsp; [go-json](https://github.com/goccy/go-json) - Drop-in replacement for encoding/json with much higher throughput. ★ 3.7k
- &emsp;&thinsp; [go-yaml](https://github.com/goccy/go-yaml) - YAML parser and encoder with precise error reporting and struct tags. ★ 2.2k

## Validation

Validating structs and inputs.

- 🥇 [validator](https://github.com/go-playground/validator) - Struct and field validation using tags, with cross-field and cross-struct rules. ★ 20k

## Dependency Injection

Wiring applications together.

- 🥇 [fx](https://github.com/uber-go/fx) - Application framework built around dependency injection and lifecycle hooks. ★ 7.7k
- 🥈 [dig](https://github.com/uber-go/dig) - Reflection-based dependency injection toolkit. ★ 4.5k
- 🥉 [do](https://github.com/samber/do) - Typesafe dependency injection based on generics. ★ 2.8k

## Concurrency

Goroutine pools, structured concurrency and synchronization helpers.

- 🥇 [ants](https://github.com/panjf2000/ants) - High-performance goroutine pool for managing and recycling goroutines. ★ 15k
- 🥈 [conc](https://github.com/sourcegraph/conc) - Structured concurrency primitives: pools, wait groups and panic propagation. ★ 10k

## Caching

In-memory caches and eviction policies.

- 🥇 [bigcache](https://github.com/allegro/bigcache) - Efficient in-memory cache for gigabytes of data without GC overhead. ★ 8.2k
- 🥈 [ristretto](https://github.com/dgraph-io/ristretto) - Concurrent cache with admission and eviction policies tuned for hit ratio. ★ 7k
- 🥉 [golang-lru](https://github.com/hashicorp/golang-lru) - Thread-safe fixed-size LRU and 2Q caches. ★ 5.1k
- &emsp;&thinsp; [otter](https://github.com/maypok86/otter) - Lock-free cache with S3-FIFO eviction and high hit ratio. ★ 2.7k

## HTTP Clients

Making requests with retries, tracing and convenience.

- 🥇 [resty](https://github.com/go-resty/resty) - Feature-rich REST client with retries, middleware and request tracing. ★ 12k
- 🥈 [req](https://github.com/imroc/req) - HTTP client with a fluent API, debugging tools and HTTP/2 and HTTP/3 support. ★ 4.9k
- 🥉 [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) - Retrying HTTP client with exponential backoff. ★ 2.3k

## Messaging & Queues

Message brokers, streaming platforms and task queues.

- 🥇 [asynq](https://github.com/hibiken/asynq) - Distributed task queue backed by Redis. ★ 14k
- 🥈 [watermill](https://github.com/ThreeDotsLabs/watermill) - Library for building event-driven applications over pub/sub backends. ★ 9.9k
- 🥉 [kafka-go](https://github.com/segmentio/kafka-go) - Kafka client with low-level and high-level APIs. ★ 8.6k
- &emsp;&thinsp; [nats.go](https://github.com/nats-io/nats.go) - Client for the NATS messaging system. ★ 6.7k
- &emsp;&thinsp; [franz-go](https://github.com/twmb/franz-go) - Feature-complete Kafka client with transactions and consumer groups. ★ 3.1k

<details>
<summary>More (1)</summary>

- [amqp091-go](https://github.com/rabbitmq/amqp091-go) - Official RabbitMQ client for AMQP 0.9.1. ★ 2k

</details>

## RPC & API

gRPC, RPC frameworks, GraphQL and OpenAPI tooling.

- 🥇 [grpc-go](https://github.com/grpc/grpc-go) - Official gRPC implementation. ★ 23k
- 🥈 [swag](https://github.com/swaggo/swag) - Generate OpenAPI documentation from source annotations. ★ 13k
- 🥉 [gqlgen](https://github.com/99designs/gqlgen) - Schema-first GraphQL server generator. ★ 11k
- &emsp;&thinsp; [huma](https://github.com/danielgtaylor/huma) - Framework for HTTP APIs with OpenAPI 3 generation and validation. ★ 4.4k
- &emsp;&thinsp; [connect-go](https://github.com/connectrpc/connect-go) - Simple RPC over HTTP/1.1 and HTTP/2 with gRPC compatibility. ★ 4.1k

## Authentication & Authorization

Tokens, sessions, OAuth and access control.

- 🥇 [casbin](https://github.com/apache/casbin) - Authorization library supporting ACL, RBAC and ABAC models. ★ 20k
- 🥈 [jwt](https://github.com/golang-jwt/jwt) - JSON Web Token implementation. ★ 9.2k
- 🥉 [goth](https://github.com/markbates/goth) - Multi-provider OAuth authentication for web applications. ★ 6.6k
- &emsp;&thinsp; [scs](https://github.com/alexedwards/scs) - HTTP session management with pluggable stores. ★ 2.6k
- &emsp;&thinsp; [go-oidc](https://github.com/coreos/go-oidc) - OpenID Connect client with verification of ID tokens. ★ 2.5k

<details>
<summary>More (1)</summary>

- [jwx](https://github.com/lestrrat-go/jwx) - Complete implementation of JWA, JWE, JWK, JWS and JWT. ★ 2.4k

</details>

## Observability

Metrics, tracing and instrumentation.

- 🥇 [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go) - OpenTelemetry API and SDK for traces, metrics and logs. ★ 6.5k
- 🥈 [client_golang](https://github.com/prometheus/client_golang) - Official Prometheus instrumentation library. ★ 6k

## Text, HTML & Templates

Parsing and generating markup and text.

- 🥇 [goquery](https://github.com/PuerkitoBio/goquery) - HTML document traversal and manipulation with a jQuery-like API. ★ 15k
- 🥈 [templ](https://github.com/a-h/templ) - Language for type-safe HTML components that compile to Go code. ★ 11k
- 🥉 [goldmark](https://github.com/yuin/goldmark) - CommonMark-compliant Markdown parser that is easy to extend. ★ 5k

## Utilities

General-purpose helpers that end up in every project.

- 🥇 [lo](https://github.com/samber/lo) - Lodash-style helper functions built on generics. ★ 21k
- 🥈 [fsnotify](https://github.com/fsnotify/fsnotify) - Cross-platform filesystem notifications. ★ 11k
- 🥉 [errors](https://github.com/pkg/errors) - Error handling primitives with stack traces. ★ 8.3k
- &emsp;&thinsp; [afero](https://github.com/spf13/afero) - Filesystem abstraction with in-memory, OS and composable backends. ★ 6.7k
- &emsp;&thinsp; [uuid](https://github.com/google/uuid) - Generation and parsing of UUIDs per RFC 4122 and DCE 1.1. ★ 6.1k

<details>
<summary>More (6)</summary>

- [ulid](https://github.com/oklog/ulid) - Universally Unique Lexicographically Sortable Identifiers. ★ 5k
- [go-humanize](https://github.com/dustin/go-humanize) - Formatters for byte sizes, times, numbers and ordinals in human-readable form. ★ 4.8k
- [backoff](https://github.com/cenkalti/backoff) - Exponential backoff algorithm with retries and contexts. ★ 4.1k
- [cast](https://github.com/spf13/cast) - Safe and easy casting between types. ★ 4k
- [retry-go](https://github.com/avast/retry-go) - Simple retry mechanism with configurable backoff. ★ 3k
- [mapstructure](https://github.com/go-viper/mapstructure) - Decoding of generic map values into native structures. ★ 484

</details>

## Developer Tools

Linters, build tools, debuggers and code generators.

- 🥇 [delve](https://github.com/go-delve/delve) - Debugger with support for goroutines, breakpoints and remote debugging. ★ 25k
- 🥈 [air](https://github.com/air-verse/air) - Live reload for development. ★ 24k
- 🥉 [golangci-lint](https://github.com/golangci/golangci-lint) - Linters runner that aggregates dozens of linters with caching and parallelism. ★ 19k
- &emsp;&thinsp; [task](https://github.com/go-task/task) - Build tool and task runner with a simple YAML syntax. ★ 16k
- &emsp;&thinsp; [goreleaser](https://github.com/goreleaser/goreleaser) - Release automation for building, packaging and publishing binaries. ★ 16k

<details>
<summary>More (4)</summary>

- [buf](https://github.com/bufbuild/buf) - Protobuf build tool with linting, breaking change detection and code generation. ★ 11k
- [mockery](https://github.com/vektra/mockery) - Mock code generator for interfaces. ★ 7.2k
- [go-tools](https://github.com/dominikh/go-tools) - Staticcheck and other advanced static analysis tools. ★ 6.9k
- [gofumpt](https://github.com/mvdan/gofumpt) - Stricter gofmt. ★ 4.1k

</details>

## Contributing

Read the [contribution guidelines](CONTRIBUTING.md) first. Every listed project gets its own [badge](CONTRIBUTING.md#badge) to show off.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the maintainers have waived all copyright and related or neighboring rights to this work.
