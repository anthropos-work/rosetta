# Storage Service

## Role & Responsibility

Storage is the **centralized file/blob service** for the platform.

> **⚠️ Since the merges, the sole live caller is `app`.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `anticheat.go:34`, `app/main.go:983` `storage.NewClient(…, storagens.CMS)`); their compose containers
> are unfederated husks sitting off every storage path, and stay up only until platform **M810**.

Callers push and pull binary objects through it instead of dealing with S3 themselves. It has two parallel storage managers — **private** (internal files, recordings, documents) and **public** (CDN-served assets) — each backed by its own S3 bucket and accessed by namespace + UUID.

Storage is stateless and owns no database: all state lives in S3 (the private manager falls back to local filesystem in dev when `STORAGE_S3_BUCKET` is unset; the public manager is wired to production S3 in compose).

## Architecture & Code Map

* **Codebase**: `storage` (local) — repo `git@github.com:anthropos-work/storage`
* **Language**: Go 1.25
* **Framework**: Connect-RPC (via the shared `colony` library), Cobra CLI
* **Database**: none — all state lives in S3 (or local filesystem in dev)
* **Ports**: 8300 (HTTP health), 8301 (Connect-RPC) — `PORT=8300` and `RPC_PORT=8301` in compose, mapped 1:1 to host (CLAUDE.md mentions different defaults at the binary level, but the platform compose pins them to 8300/8301 in both directions)
* **Profile**: `graphql` (default) and `storage`

### Two storage managers

| Manager | Bucket env | Access pattern |
|---------|------------|----------------|
| Private | `STORAGE_S3_BUCKET` | Internal data: session recordings, documents. Reads via RPC or presigned URLs. |
| Public | `STORAGE_S3_PUBLIC_BUCKET` | Public assets served via CloudFront at `media.<root_domain>`. |

Each manager falls back to local filesystem only when ITS bucket env var is empty (private → `/tmp/anthropos-storage/`, public → `/tmp/anthropos-public-storage/`). In the platform compose, `STORAGE_S3_PUBLIC_BUCKET` is hardcoded to the production public bucket, so locally the PUBLIC manager talks to real S3 (`PutPublicObject`/`GetPublicObject` require AWS credentials), while the PRIVATE manager uses local FS.

### Object layout

Each stored object is two files:

```
<namespace>/<uuid>                  # raw binary data
<namespace>/<uuid>_metadata.json    # size, content_type, name, created_at
```

Namespaces are arbitrary strings (e.g., `jobsimulation`, `assets`). The `viant/afs` abstraction routes reads/writes to S3 or local FS based on which bucket is configured.

### Key directories

```
main.go                       Entry point
cmd/
  root.go                     Server startup (HTTP + RPC), graceful shutdown
  put.go, get.go, sync.go     CLI: upload, download, bulk-migrate
internal/
  rpcsrv/rpcsrv.go            Connect-RPC handler implementations
  storage/storage.go          StorageManager interface + S3/filesystem backends
  migration/                  Sync engine + transformers (S3 ↔ local migration)
    migration.go
    s3.go
    filesystem.go
sdk/storage/                  Go SDK for in-platform consumers
  client.go                   NewClient / NewPublicClient
  v1/                         Versioned RPC client
terraform/                    ECS, S3, CloudFront, Route53
```

## Interface Discovery

### Connect-RPC (`StorageService`)

Private:

| Method | Request | Response |
|--------|---------|----------|
| `PutObject` | `data`, `metadata`, `namespace` | `key (UUID)`, `namespace` |
| `GetObject` | `key`, `namespace` | `object (data + metadata)` |
| `GetPresignedUrl` | `key`, `namespace`, `expiry_seconds` | `url` (default 15 min) |

Public:

| Method | Request | Response |
|--------|---------|----------|
| `PutPublicObject` | `data`, `metadata`, `namespace` | `key (UUID)`, `namespace` |
| `GetPublicObject` | `key`, `namespace` | `object (data + metadata)` |

### SDK (Go)

Other services use the in-repo Go SDK rather than raw Connect-RPC clients:

```go
import "github.com/anthropos-work/storage/sdk/storage"

// Private
client := storage.NewClient("http://storage:8301", "jobsimulation")
key, _ := client.V1.PutObject(ctx, data, metadata)
obj, _ := client.V1.GetObject(ctx, key)

// Public
pubClient := storage.NewPublicClient("http://storage:8301", "assets")
```

### CLI

```bash
storage                                       # start server
storage put -f /path/to/file -n <namespace>   # upload
storage get -k <uuid> -n <namespace> -o <dir> # download
storage sync <source> <dest> [--dry-run]      # bulk migrate
```

## Dependencies

* **Upstream consumers**: **`app` only** — the jobsimulation domain (recordings, simulation documents),
  the cms domain (content assets, media) and app itself (user files, profile images) all call from
  inside the `backend` binary. The `jobsimulation`/`cms` husk containers call nothing (teardown M810).
* **Downstream**: AWS S3 (production), CloudFront (public bucket), `colony` shared library, `proto` for RPC contracts
* **No outbound RPC** to other platform services — storage is a leaf

## Local Development

### Run in Docker

```bash
cd platform
make up                       # default graphql profile — includes storage
# or just storage:
make up PROFILE=storage
```

In local dev the PRIVATE manager falls back to `/tmp/anthropos-storage/` automatically (`STORAGE_S3_BUCKET` is unset in compose), and its presigned URLs return empty strings in that mode (`storage.go:122`). FOOTGUN: the PUBLIC manager is NOT sandboxed locally — compose hardcodes `STORAGE_S3_PUBLIC_BUCKET` to the production public bucket, so `PutPublicObject`/`GetPublicObject` hit real S3 and fail without AWS credentials (none are set in `platform/.env`). To run public storage fully local, override `STORAGE_S3_PUBLIC_BUCKET` to empty; it then falls back to `/tmp/anthropos-public-storage/` (a separate path from the private fallback).

### Run natively

```bash
cd platform
make dev S=storage
cd ../storage
go run main.go   # or: go run .
```

`make setup`/`make gen` exist in the Makefile but are legacy no-ops — the repo has no codegen (no `//go:generate` directives, no gqlgen/graphql usage; gqlgen is vestigial).

### Sync between backends

The `storage sync` CLI moves objects between two configured backends (e.g., local FS → S3 for an initial seed):

```bash
storage sync /tmp/anthropos-storage s3://anthropos-private-bucket --dry-run
```

## Environment Variables

| Variable | Compose value | Description |
|----------|---------------|-------------|
| `PORT` | `8300` | HTTP health port (binary default 8080, overridden in compose) |
| `RPC_PORT` | `8301` | Connect-RPC port (binary default 8081, overridden in compose) |
| `STORAGE_S3_BUCKET` | (empty) | Private bucket. Absent from compose env and `.env` → local FS fallback at `/tmp/anthropos-storage/`. |
| `STORAGE_S3_PUBLIC_BUCKET` | `production-storage-public20240919130721114900000001` | Public bucket — hardcoded to a real PRODUCTION S3 bucket in compose (**`docker-compose.yml:210`** @ platform `2adcf71`; `:324` is inside the *studio-desk* block). NOT empty in local dev. |
| `AWS_REGION` / `AWS_DEFAULT_REGION` | `eu-west-1` | AWS region (EU-first) |
| `ENVIRONMENT` | (empty) | Environment name |
| `SERVICE_NAME` | `storage` | Logging label |
| `SENTRY_DSN` | (empty) | Sentry error tracking |

## Testing

```bash
cd storage
go test -v ./...
```

Note: the service currently ships NO automated tests (no `*_test.go` files in the repo), so `go test ./...` is a trivial no-op. The same command is baked into the production Dockerfile (`Dockerfile:18`) as a build gate, but it likewise tests nothing — do not read it as evidence of a real suite.

## Related Documentation

* [Backend (app)](./backend.md), [CMS](./cms.md), [Jobsimulation](./jobsimulation.md) — consumers
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)
