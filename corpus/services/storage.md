# Storage Service

## Role & Responsibility

Storage is the **centralized file/blob service** for the platform.

> **⚠️ MERGED INTO `app` — storage is served in-process since the v9.0 fold (2026-08-04), and there is no live RPC caller at all.** The jobsimulation and cms domains run
> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,
> `app/internal/jobsimulation/anticheat/anticheat.go:30` — spelled in full because there is **no** `anticheat.go`
> beside `recording.go`; the storage-SDK import lives in the sibling `anticheat/` package, so reading the bare
> filename against `recording/` resolves to nothing — `app/main.go:1048`
> `internalstorage.NewClient(storageManager, storagens.CMS)` @ `app`
> `9d00a313` v1.367.0); at platform `0dab54d` they have **no compose containers at all** — the local
> husks are gone, along with the `cms` and `jobsimulation` `repos.yml` entries. (Local compose only;
> the production **M810** rollback-path teardown was not measured here.) Since v9.0 the object-storage
> *manager itself* is in-process too, so that call site no longer crosses a network at all.
>
> **⚠️⚠️ And the v9.0 fold COMPLETED on 2026-08-04 — the container went the next morning.** This block read MID-FOLD for four M257x
> iterations; the half it was waiting on landed in a single working morning, and platform `838d907`
> (merged `0c91421`, 2026-08-05, *"drop the storage, messenger and customerio-sync containers"*) then removed the local service and
> its clone entry. Re-derived on BOTH sides, because one side alone is not a claim (`D-M257x-59-4`).
>
> | side | measured at platform `0c91421` / app `2035f9a` / storage `9f8cb53` (a **prod** claim is settled by that repo's `origin/main` — `9f8cb53` for `storage`; `2035f9a` was `app`'s `origin/main` when this table was written and `ad9f3c49` is on 2026-08-06, with every `app` anchor below byte-identical across the two) |
> |---|---|
> | **prod** | the ECS service is **DELETED, not scaled to zero** — `storage/terraform/main.tf` is 18 lines and declares no service block at all. Its own comment: *"The ECS service that used to live here is GONE (v9.0 'support-in-app'): app serves object storage in-process. What remains is the ASSETS — the two buckets, the CloudFront distribution and the media DNS record, all in storage.tf."* The module deliberately survives: deleting the block would destroy those assets along with their `prevent_destroy` guards, which are read from configuration (`storage/terraform/main.tf:13`). **`:18` is not a custody clause** — it reads *"See outputs.tf — consumers should reference these by output, never by literal name."*, and *custody* occurs **0** times in the storage repo at `9f8cb53`. **M903 was never executed and is superseded**: its only mention in the repo is `storage/terraform/storage.tf:22-25` — *"An earlier plan, M903, instead proposed `moved`-ing these resources into a local module in the infrastructure repo. That was never executed, and the shipped design supersedes it: the assets stay here and the module stays declared."* No `moved` block exists in the repo. The plan itself is `app/knowledge/plan/releases/09.00-support-in-app/m903-s3-custody/` @ app `2035f9a` — its `overview.md` front-matter says `status: planned`, its `progress.md` says *"Planned, not started — no terraform applied, no branch"*. **State the ref or this flips:** at the checked-out storage ref `4ce8ece5` (20 behind `9f8cb53`) the same header still reads M903 as a live instruction — *"relocate the assets out of this module BEFORE M907"* — and `main.tf` is 100 lines still declaring the ECS module. `d3e6d32` (2026-08-05) retired it: *"M903 never ran."* |
> | **config** | `STORAGE_RPC_ADDR` is set by **no** compose file and is **absent from `.env_example`** — 0 occurrences across `docker-compose.yml`, `common.yml`, `.env_example` |
> | **compose** | there is **no `storage` service to start.** `0dab54d` parked it behind a rollback-only profile for one release; `838d907` then deleted the service block outright, and the `storage-legacy` profile is gone with it |
> | **`repos.yml`** | **no `storage` entry** — `838d907` removed it, so `make init` does not clone the repo any more. Four entries remain: `app`, `sentinel`, `next-web-app`, `studio-desk`. (It was `repos.yml:18-20` at `0dab54d`, kept then as the rollback path.) The repo itself is not deleted — clone it by hand to read the pre-merge source |
> | **consumer** | `app` serves object storage **in-process**: `internalstorage.NewManager` / `NewPublicManager` at `app/main.go:524`, `:525`, consumed at `:547` (`resource.NewManager`) and `:1102` (`cmsStorage`); the constants at `app/internal/storage/service.go:22`, `:24` are the bucket **env-var NAMES** (`EnvBucket = "STORAGE_S3_BUCKET"`, `EnvPublicBucket = "STORAGE_S3_PUBLIC_BUCKET"`), not the bucket names — those are still read from the environment, at `app/main.go:516`, `:517`. `STORAGE_RPC_ADDR` is read by **nothing** *at `app` `2035f9a`* — `git -C stack-demo/app grep -n STORAGE_RPC_ADDR 2035f9a -- '*.go'` returns **3 hits, all of them comments**. **That ref used to be written here as `origin/main`, and the LABEL expired on 2026-08-06:** `app`'s `origin/main` is now **`ad9f3c49`**, five commits on. The sha still means what it meant — a pin is a pin, a branch name is not (M257x iter-102). Naming the newer ref costs nothing *here* because the anchors do not move across it: `git -C stack-demo/app diff --stat 2035f9a4 ad9f3c49 -- main.go internal/storage/` is **empty**, and the same grep at `ad9f3c49` returns the identical 3 comment hits. **The ref is still load-bearing and this cell states it deliberately:** until 2026-08-06 the demo's build pin was the older `b948604f`, where that grep returns **15 hits, 7 of them live env reads** — so a ref-less command contradicted this sentence on the clone a reader actually had. The pin has since advanced to `ad9f3c49` (= `origin/main`) and a ref-less grep now agrees; the ref stays written down anyway, because a build pin can move again. The older side's per-hit breakdown lives in [`platform-migration-status.md`](../architecture/platform-migration-status.md)'s `storage` row and is **deliberately not repeated here** — a block that names two refs its anchors resolve *differently* at makes every anchor in it ungradeable (M257x run-53). The three comment hits: `app/main.go:504` (*"the standalone service takes no traffic and STORAGE_RPC_ADDR is gone"*), `app/internal/jobsimwiring/wiring.go:101` and `app/internal/storagens/callsites_test.go:189`. (Those three paths are in the **`app`** repo, not this one; a bare `main.go` here resolves to `storage`'s own 18-line `main.go`.) **Zero** `os.Getenv` sites, in `main.go` or in any of the three `cmd/` tools |
>
> The mid-fold hazard this block used to describe — a client built against an empty address, failing
> at call time rather than boot time, and two `cmd/` tools hard-failing outright — **is gone**, because
> there is no longer a client to build. `platform_predicate_guard.py` G6 derives that consumer side
> **at a named ref** (M257x iter-68), and the ref is the whole story: at the demo's *former* build pin
> `b948604f` the guard reads a mid-fold, because `STORAGE_RPC_ADDR` is live there (15 Go hits, 7 of them
> env reads); at the demo's build pin **today** — `ad9f3c49`, level with `origin/main` since 2026-08-06 —
> it is 3 hits, all comments. **Two opposite verdicts, and what selects between them is only which
> checkout you happen to be looking at** — which is why G6 names its ref and why this block does too.
> (Corrected M257x iter-102: this passage named `b948604` as the *current* demo pin, which it stopped
> being when the clone advanced to `ad9f3c49`.)

Callers push and pull binary objects through it instead of dealing with S3 themselves. It has two parallel storage managers — **private** (internal files, recordings, documents) and **public** (CDN-served assets) — each backed by its own S3 bucket and accessed by namespace + UUID.

Storage is stateless and owns no database: all state lives in S3 — and since platform `0dab54d` **both** managers are wired to **production** buckets in compose (`docker-compose.yml:82`, `:83` @ `0c91421`, on `backend`), not just the public one. Each manager falls back to local filesystem only when ITS bucket variable is set **empty**; on a stock stack neither is. See the hazard note under "Two storage managers".

## Architecture & Code Map

* **Codebase**: `storage` — repo `git@github.com:anthropos-work/storage`. **Not cloned by `make init`**: `838d907` removed the `repos.yml` entry along with the container. Clone it by hand to read the pre-merge source; the live code is `app/internal/storage/` (+ `app/internal/storagens/`, `app/internal/publicstorage/`)
* **Language**: Go 1.25
* **Framework**: Connect-RPC (via the shared `colony` library), Cobra CLI
* **Database**: none — all state lives in S3 (or local filesystem in dev)
* **Ports**: 8300 (HTTP health), 8301 (Connect-RPC) — `PORT=8300` and `RPC_PORT=8301` were injected by the `storage` compose block, mapped 1:1 to host, until `838d907` deleted that block. The repo's own CLAUDE.md documents different binary-level defaults; **nothing publishes 8300/8301 on a stack now**
* **Profile**: **none — there is no `storage` compose service.** Platform `838d907` (merged `0c91421`, 2026-08-05) deleted the service block outright; `0dab54d` had parked it behind a rollback-only profile for one release, and the `storage-legacy` profile is gone with it. Corrected M257x iter-87 — iter-68 had corrected the same line for naming two profiles the platform did not have

### Two storage managers

| Manager | Bucket env | Access pattern |
|---------|------------|----------------|
| Private | `STORAGE_S3_BUCKET` | Internal data: session recordings, documents. Reads via RPC or presigned URLs. |
| Public | `STORAGE_S3_PUBLIC_BUCKET` | Public assets served via CloudFront at `media.<root_domain>`. |

Each manager falls back to local filesystem only when ITS bucket env var is empty (private → `/tmp/anthropos-storage/`, public → `/tmp/anthropos-public-storage/`) — `getKeyPath` branches on `s3Bucket != ""` (`app/internal/storage/storage.go:193-200` @ app `2035f9a`), so any **non-empty** value routes to `s3://…` unconditionally, for both managers.

> **⚠️ HAZARD — on a stock stack NEITHER manager uses local FS, and the private one writes to production.**
> In the platform compose **both** buckets are hardcoded to **production** buckets, on the `backend` service
> block: `STORAGE_S3_BUCKET=production-storage20240826131618541000000005` (`docker-compose.yml:82`) and
> `STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001` (`:83`) @ platform
> `0c91421`, with the reason in-comment at `:73-79` (*"These MUST be set"*). The **private** line arrived
> with `0dab54d` (2026-08-03, *"run without the standalone storage; rename graphql -> core"*) — that commit
> is when the private manager stopped being local. **This document previously said the private manager uses
> local FS while only the public one talks to real S3; that is RETRACTED.**
>
> The credentials are there by design, not by accident: `backend` mounts `$HOME/.aws/credentials` read-only
> (`docker-compose.yml:100`) and platform's own `README.md:81-87` instructs you to put a live
> `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`. **So a local stack with
> working AWS credentials writes its private uploads into the production private bucket.** Nothing warns:
> app's two boot guards (`main.go:518-523` empty-bucket fatal, `:529-535` `verifyBucketAccess`) both run
> only `if deployedEnvironment()`, and `deployedEnvironment()` returns **false** for
> `ENVIRONMENT=development` (`app/env_guards.go:37-44`). **All three of those anchors are at `app`
> `ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06, and byte-identical at `2035f9a4`.
> **The ref is not decoration here** (M257x iter-102): at the demo's *former* pin `b948604f`,
> `git -C stack-demo/app ls-tree b948604f -- env_guards.go` returns **nothing** — the file does not exist —
> and both `main.go` anchors resolve there onto the wrong constructs (`:518-523` is the **public-storage
> clients** block, `:529-535` the **academy asset uploader**), so the whole "nothing warns" mechanism is
> unverifiable on that checkout. And compose supplies the disarming value itself:
> `- ENVIRONMENT=development` on the `backend` block, `docker-compose.yml:56` @ platform `0c91421`.
>
> The empty-env escape still works — set the variable to empty **explicitly** and that manager falls back to
> its `/tmp` root. **Disposition of this hazard is an open escalated item
> (`DEF-M257x-iter80-storage-prod-bucket`, severity high), FILED at M257x iter-102 as
> `PLATFORM-M257x-compose-points-local-backend-at-the-PRODUCTION-S3-buckets` in
> [`knowledge/plan/platform-defect-register.md`](../../knowledge/plan/platform-defect-register.md).** It is a
> platform/compose fact this repo cannot fix under its zero-platform-edit rule, not a corpus fact; this
> document records the exposure and deliberately prescribes no change.

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
  inside the `backend` binary. There are no `jobsimulation`/`cms` containers left to call anything — `d11a403` deleted both compose services and both `repos.yml` entries. In prod their fates now differ: `jobsimulation`'s ECS service is already destroyed (`6092c6d2` — M810 landed for that row), while `cms`'s ECS module survives **in cms's own repo** at `service_desired_count = 0` — but its prod state is **UNMEASURABLE** (`6efa1d5` deleted cms's build workflow saying the ECR *"is decommissioned (M810)"*, and the deletion itself lands in `infrastructure`, never in a clone set).
* **Downstream**: AWS S3 (production), CloudFront (public bucket), `colony` shared library, `proto` for RPC contracts
* **No outbound RPC** to other platform services — storage is a leaf

## Local Development

### Run in Docker

```bash
cd platform
make up                       # the `core` profile — and there is no storage container to include
# You cannot start one. Platform 838d907 (merged 0c91421, 2026-08-05) deleted the service block
# outright, and the `storage-legacy` profile is gone with it — 0dab54d had parked storage behind
# that profile for a single release, as a rollback path, while app took object storage in-process.
# Asking for the retired `storage` token does NOT fail: it exits 0 and starts only the floor.
```

**What the container used to do with its buckets, and what the binary still does with them.** With `STORAGE_S3_BUCKET` empty the PRIVATE manager falls back to `/tmp/anthropos-storage/` automatically, and its presigned URLs return empty strings in that mode (`storage.go:122`). FOOTGUN: the PUBLIC manager is not sandboxed by that fallback — the deleted `storage` compose block hardcoded `STORAGE_S3_PUBLIC_BUCKET` to the production public bucket, so `PutPublicObject`/`GetPublicObject` hit real S3 and failed without AWS credentials. (That parenthetical used to say none were set in `platform/.env` — **no longer true**: platform `README.md:81-87` @ `0c91421` instructs you to put live `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_SESSION_TOKEN` in `.env`, and `docker-compose.yml:100` mounts `$HOME/.aws/credentials` into `backend`, so on a current stack the credentials are generally present and the write **succeeds**.) Running the binary by hand, override `STORAGE_S3_PUBLIC_BUCKET` to empty; it then falls back to `/tmp/anthropos-public-storage/` (a separate path from the private fallback).

### Run natively

```bash
cd storage
go run main.go   # or: go run .
```

The old first step — `cd platform && make dev S=storage`, which stopped the container so the native
process could take its port — is moot: `838d907` removed the container, and the repo is not a sibling
clone any more (`make init` does not fetch it), so `cd storage` assumes you cloned it by hand.

`make setup`/`make gen` exist in the Makefile but are legacy no-ops — the repo has no codegen (no `//go:generate` directives, no gqlgen/graphql usage; gqlgen is vestigial).

### Sync between backends

The `storage sync` CLI moves objects between two configured backends (e.g., local FS → S3 for an initial seed):

```bash
storage sync /tmp/anthropos-storage s3://anthropos-private-bucket --dry-run
```

## Environment Variables

> **HISTORICAL — there is no `storage` container to inject any of these into.** The middle column
> records what `docker-compose.yml` set on the `storage` service block, **read at platform `0dab54d`**
> (the last ref that had one — the block's seven `environment:` entries are identical at `2adcf71`);
> `838d907` (merged `0c91421`, 2026-08-05) deleted that block, so nothing sets them for this binary any
> more. They still describe the binary's own inputs if you run it by hand. A cell reading *(empty)* is a
> claim that **the block did not set it** — not that it set it to an empty string.

| Variable | Compose value | Description |
|----------|---------------|-------------|
| `PORT` | `8300` | HTTP health port (binary default 8080, overridden in compose) |
| `RPC_PORT` | `8301` | Connect-RPC port (binary default 8081, overridden in compose) |
| `STORAGE_S3_BUCKET` | (empty) | Private bucket. The deleted `storage` service block never set it, so **this binary** fell back to `/tmp/anthropos-storage/`. **Do not read that across to the live path:** `backend` sets it to the production private bucket (`docker-compose.yml:82` @ `0c91421`) — see the hazard note under "Two storage managers". |
| `STORAGE_S3_PUBLIC_BUCKET` | `production-storage-public20240919130721114900000001` | Public bucket — hardcoded to a real PRODUCTION S3 bucket in compose (**`docker-compose.yml:210`** @ platform `2adcf71`; `:324` is inside the *studio-desk* block). NOT empty in local dev. |
| `AWS_REGION` / `AWS_DEFAULT_REGION` | `eu-west-1` | AWS region (EU-first) |
| `ENVIRONMENT` | `development` | Environment name. **The block DID set it** — `- ENVIRONMENT=development` at `docker-compose.yml:119` @ platform `0dab54d` and `:206` @ `2adcf71`. (Corrected M257x iter-102; this cell read *(empty)*, i.e. "never set by compose". The error mattered: `development` is precisely the value that makes `deployedEnvironment()` return false and **disarms** app's boot guards — see the hazard note under "Two storage managers" — so recording it as unset hid the mechanism.) |
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
