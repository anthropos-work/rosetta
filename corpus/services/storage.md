# Storage Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **v9.0 "support-in-app"** (2026-08-04), the standalone `storage` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"), alongside
> [messenger](./messenger.md) and [customerio-sync](./customerio-sync.md). `backend` reads and writes
> both S3 buckets **directly**; there is no object-storage RPC hop left anywhere on the platform.
>
> Where everything went:
>
> * **Domain** — `app/internal/storage/` holds both managers (`NewManager` / `NewPublicManager`),
>   constructed once at boot and handed to every consumer as an in-process dependency. The former
>   callers are all inside the same binary now: the jobsimulation domain (recordings, anti-cheat
>   captures), the cms domain (content assets, media) and app itself (user files, profile images).
> * **Config** — `backend` reads **`STORAGE_S3_BUCKET`** (private) and
>   **`STORAGE_S3_PUBLIC_BUCKET`** (public), plus `AWS_REGION`/`AWS_DEFAULT_REGION`.
>   **`STORAGE_RPC_ADDR` is gone** — no code reads it; the only surviving occurrences are comments
>   saying so.
> * **Boot guard** — `backend` proves at boot that its task role can actually reach both buckets under
>   the names it was given, because both are created with terraform `bucket_prefix` and carry a
>   generated suffix. The guard is **disarmed by `ENVIRONMENT=development`**, which is the trap worth
>   knowing: with empty bucket names a local `backend` boots "fine" and silently writes every upload to
>   the container's ephemeral disk.
> * **Infrastructure — the ECS service is gone; the terraform module is NOT, and must not be deleted.**
>   `module "storage-service_euwest1"` now declares only the platform's object-storage **assets**: both
>   S3 buckets (~92 GiB), their versioning and SSE, the CloudFront distribution + OAI + bucket policy,
>   and the `media.anthropos.work` CNAME — all of which `backend` reads and writes directly.
>   `prevent_destroy` will **not** save you here: it is read from *configuration*, so removing the block
>   removes the guards along with the resources they guard. `backend`'s `storage_s3_bucket` /
>   `storage_s3_public_bucket` / `media_url` inputs are wired from this module's outputs.
>   **Name collision:** `module.storage-service_euwest1` (the buckets) is **not**
>   `module.storage_euwest1` (`modules/core/storage` — the RDS instance and the ElastiCache replication
>   group). Never abbreviate either name; never target them with a wildcard.
> * **Repo** — the `storage` git repo still exists but is **frozen/legacy**; make changes in `app`.
>   It is still in `repos.yml` and still startable from the `storage-legacy` compose profile, as the
>   rollback path.
>
> **Everything below this banner describes the standalone service.** The object layout, namespaces and
> manager semantics carried over unchanged and are still accurate; the RPC surface, the SDK and the
> ports are the rollback target's, not the live path's.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

## Role & Responsibility

Storage is the **centralized file/blob service** for the platform. Other services (`jobsimulation`, `cms`, `app`) push and pull binary objects through it instead of dealing with S3 themselves. It has two parallel storage managers — **private** (internal files, recordings, documents) and **public** (CDN-served assets) — each backed by its own S3 bucket and accessed by namespace + UUID.

Storage is stateless and owns no database: all state lives in S3 (the private manager falls back to local filesystem in dev when `STORAGE_S3_BUCKET` is unset; the public manager is wired to production S3 in compose).

## Architecture & Code Map

* **Codebase**: `storage` (local) — repo `git@github.com:anthropos-work/storage`
* **Language**: Go 1.25
* **Framework**: Connect-RPC (via the shared `colony` library), Cobra CLI
* **Database**: none — all state lives in S3 (or local filesystem in dev)
* **Ports**: 8300 (HTTP health), 8301 (Connect-RPC) — `PORT=8300` and `RPC_PORT=8301` in compose, mapped 1:1 to host (CLAUDE.md mentions different defaults at the binary level, but the platform compose pins them to 8300/8301 in both directions)
* **Profile**: `storage-legacy` **only**, since v9.0 — **not** in the default selection. (The old line named two profiles the platform no longer has: `graphql` was renamed `core` at platform `0dab54d`, and there was never a bare `storage` profile.)

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

* **Upstream consumers**: **`app` only** — the jobsimulation domain (recordings, simulation documents), the cms domain (content assets, media) and app itself (user files, profile images) all call from inside the `backend` binary. Since v9.0 the manager itself is in-process, so those calls no longer cross a network at all
* **Downstream**: AWS S3 (production), CloudFront (public bucket), `colony` shared library, `proto` for RPC contracts
* **No outbound RPC** to other platform services — storage is a leaf

## Local Development

### Run in Docker

```bash
cd platform
make up                       # the `core` profile — which does NOT include storage any more
# To start the standalone (rollback comparison only — app serves storage in-process,
# and running both means two writers on one bucket):
docker compose --profile storage-legacy up storage
```

> **The local storage story changed with the fold.** `backend`'s compose env now sets **both** bucket
> names to the **real production buckets**, so a default local stack reads and writes production S3
> directly — the private manager no longer falls back to `/tmp` because `STORAGE_S3_BUCKET` is no
> longer empty. The paragraph below describes the standalone container's behaviour, which still holds
> when you start `storage-legacy`.

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
| `STORAGE_S3_PUBLIC_BUCKET` | `production-storage-public20240919130721114900000001` | Public bucket — hardcoded to a real PRODUCTION S3 bucket in compose (`docker-compose.yml:324`). NOT empty in local dev. |
| `AWS_REGION` / `AWS_DEFAULT_REGION` | `eu-west-1` | AWS region (EU-first) |
| `ENVIRONMENT` | (empty) | Environment name |
| `SERVICE_NAME` | `storage` | Logging label |
| `SENTRY_DSN` | (empty) | Sentry error tracking |

### The in-app variables (what you actually set now)

Read by **`backend`**, not by this container:

| Variable | Compose value | Description |
|----------|---------------|-------------|
| `STORAGE_S3_BUCKET` | `production-storage20240826131618541000000005` | Private bucket. **Not empty any more** — the local default is the real production private bucket |
| `STORAGE_S3_PUBLIC_BUCKET` | `production-storage-public20240919130721114900000001` | Public bucket, fronted by CloudFront at `media.anthropos.work` |
| `AWS_REGION` / `AWS_DEFAULT_REGION` | `eu-west-1` | Required — the buckets live in eu-west-1 |
| `MEDIA_URL` | `https://media.anthropos.work` | The public read URL `backend` hands out for public objects |
| ~~`STORAGE_RPC_ADDR`~~ | *(gone)* | Set by no compose file, absent from `.env_example`, read by no code. Its only remaining occurrences are comments recording that it is gone |

> **FOOTGUN, and a bigger one than the standalone's.** Both bucket names default to **production**
> buckets on a local stack, so `backend` writes to real S3 out of the box (given credentials — the
> compose file mounts `~/.aws/credentials` read-only). Override both to empty for a fully local run,
> and expect the boot-time bucket-access guard to be silent about it: `ENVIRONMENT=development`
> disarms the guard, so empty buckets look like a healthy boot while every upload goes to the
> container's ephemeral disk.

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
