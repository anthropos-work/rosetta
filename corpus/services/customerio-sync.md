# CustomerIO Sync Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).

## Role & Responsibility

CustomerIO Sync is the **background data-sync domain** that pushes platform users into the marketing
platform as contacts, for marketing automation, lifecycle email campaigns, and product analytics.

It's a one-directional pipeline: PostgreSQL `public` schema → the marketing API. No inbound traffic,
no consumers inside the platform.

## Architecture

* **Live code**: `app/internal/customeriosync/` (`sync.go` the push, `store.go` the read, `sync_query.sql` the query, `doc.go` the port record) + the manual tool `app/cmd/customerio-resync`
* **Frozen repo**: `git@github.com:anthropos-work/customerio-sync` (private) — not deleted, but not built, cloned or deployed any more
* **Language**: Go
* **Profile**: **none — there is no `customerio-sync` compose service.** `838d907` deleted the service block outright, and the `customerio-sync` profile is gone with it; the same commit dropped it from `all`, which is where it also used to appear
* **Local port**: none. The standalone published 8080 (HTTP — health/metrics); the folded domain has no port of its own

### What changed in the port

Two things deliberately did not survive as-is, both recorded in `doc.go` rather than buried:

* **The sync window is now stateless.** The standalone held `lastSyncTime` in memory and advanced it per tick. `app`'s worker is multi-replica and restarts on every deploy, so each run reads a fixed overlap window instead (wider than the 10-minute schedule). Pushing a contact is idempotent (`CreateContact` with `UpdateEnabled`), so the overlap is free.
* **The scan is `database/sql`, not sqlc/pgx.** `app` carries no sqlc toolchain. The consequence is NULL handling, not shape: a NULL now sends `""` / `0` / `false` where sqlc's pgtype wrappers left the attribute unset.

Two more structural moves: the standalone's 10-minute ticker became an **asynq scheduled task**, and
the DB **view** it read (`public.customer_io_sync_table`, never under migration control) was replaced
by `sync_query.sql` — verified row-for-row against the view, identical for 7,392 of 7,397 rows and
correct where it differs.

It runs on `app`'s **shared** analytics pool (`copilotDB`, `app/main.go:393-396` @ `2035f9a` — re-pinned M257x iter-126; the block named no ref, so the anchor could drift silently), not a pool of its own — one query every
ten minutes does not earn a standing allocation against the platform's connection budget
(`app/main.go:393-396`).

### Compose definition — **HISTORICAL**

This block was deleted at `838d907`. It is kept because the build pattern was **the LAST of its kind** in the platform
and a reader will meet it in older runbooks and in `git log`:

```yaml
# DELETED at platform 838d907 (2026-08-05) — no longer in docker-compose.yml
customerio-sync:
  build:
    context: git@github.com:anthropos-work/customerio-sync.git#main
    ssh: ["default"]
    args:
      VERSION: dev
      GH_ACCESS_TOKEN: $GH_PAT
  ports: ["8080:8080"]
  environment:
    - DB_CONNECTION_BACKEND=postgresql://postgres@postgresql:5432/postgres?sslmode=disable&search_path=public
  depends_on:
    postgresql: { condition: service_healthy }
```

Note `context: git@github.com:...#main` — Docker BuildKit cloned the repo at build time, no local
checkout needed, which worked because the build ran inside an SSH-agent-forwarded context
(`ssh: ["default"]`) with `$GH_PAT` available. ⚠️ **CORRECTED M257x iter-115 — it was the LAST service built that way, never the only one.** Re-derived over the whole history of `platform`, every `context:` line matching a git URL in every commit touching `docker-compose.yml`, de-duplicated: **18 distinct repo URLs** — app, chronos, cms, customerio-sync, graphql, graphql-wundergraph, graphqltmp, intelligence, jobsimulation, messenger, realtime, roadrunner, sentinel, simulator, skiller, skillpath, storage, studio-desk. Building from `git@github.com:anthropos-work/<repo>.git#main` was the platform **default** until `a2a3ee6` (2026-02-27, *"add Makefile, repos.yml, and switch to local Dockerfile.dev builds"*), and even after it **two** services kept a git URL — `customerio-sync` and `realtime`, the latter until `c17cc9a` (2026-04-15). Only from that date was customerio-sync the sole one, and at `838d907^` it is indeed the only git-URL context among seven. **The corpus refutes the superlative one file away**: [`external_services.md`](../architecture/external_services.md)'s *Build Context* note records that the router *"was changed from the old 'git+url' build"* too. **The
pattern died with it**: every remaining compose build takes a local sibling directory as its context.

### Dependencies

* **PostgreSQL** (`public` schema) — read via `app`'s shared analytics pool, not a connection string of its own
* **Brevo API** (external — `BREVO_KEY` in `platform/.env`)

## Interface Discovery

It exposes no business API to other platform services — it is a pure sync worker, and always was.
The standalone published 8080, which most likely served health/metrics (never verified); the folded
domain has no port of its own, and is reached only through the scheduled asynq task and the
`customerio-resync` tool.

For protocol and field-mapping details, read the live code — `app/internal/customeriosync/`, whose
`doc.go` is the port record — or clone the frozen repo for the pre-merge source:

```bash
gh repo clone anthropos-work/customerio-sync
```

## Local Development

**There is nothing to start.** `838d907` deleted the compose service, so no selection of profiles
brings a `customerio-sync` container up — and asking for the retired token does not error: compose
**exits 0** and starts only the always-on floor (`postgresql`, `redis`), which looks like
a live stack.

The domain is inside `backend` and **off unless you say otherwise**:

```bash
# In platform/.env — NOT in docker-compose.yml, which deliberately pins no value:
#   CUSTOMERIO_SYNC_ENABLED=true
#   BREVO_KEY=<a real key>
# Then `make up`. Both are required: app log.Fatalf's if the switch is on with an
# empty key (app/main.go:295-300 — the condition at :295, the log.Fatalf at :296),
# because the sender has no console fallback.
#
# WARNING: this writes REAL marketing contacts. That is why the switch exists.
```

For a one-off run you do not need the switch at all — invoking the tool by name is the consent:

```bash
cd app
go run ./cmd/customerio-resync --dry-run   # --dry-run constructs no Brevo client
```

For most local-development tasks you do not need any of this.

## Production

Runs inside the `backend` ECS task, gated by the same `CUSTOMERIO_SYNC_ENABLED` switch — where unset
is **fatal** rather than off, so a deployed environment must state its intent. The mechanism is
`app/env_guards.go:98-104` (`resolveSubsystemSwitch`'s `case "":` returns an error when `deployed`) via
`mustSubsystemSwitch`'s `log.Fatalf` at `:87`. (`app/main.go:284` is only the **comment** pointing at it,
not the mechanism — anchor corrected M257x iter-102.)
**Scope note:** whether its own ECS task / image / terraform module have been torn down was **not
measured** in this pass — the fold and the container deletion are local-compose and `app`-source
facts. Do not read them as the production teardown.

## Related Documentation

* [Backend (app)](./backend.md) — the host process, and the source of user/org data
* [Service Taxonomy](../architecture/service_taxonomy.md) — orchestration profiles
* [External Services](../architecture/external_services.md) — the third-party-integration index. **It has
  no Customer.io section and no Brevo section**: its per-service `##` sections are Clerk, Directus, the
  Cosmo router, AI Providers, LiveKit and AWS Chime, and `brevo` occurs **0** times in the whole file
  (corpus HEAD, M257x iter-102). This bullet used to gloss it as *"Customer.io as an integrated SaaS"*,
  which was false in **both** directions — there is no such section, and the destination is **Brevo**, not
  Customer.io, as the fossil-name banner at the top of this file says
* [Messenger](./messenger.md) — the **other** Brevo consumer, folded into `backend` in the same v9.0
  program, and the file that actually documents the Brevo integration

## Notes

The "build from GitHub URL" pattern was intentional while it lasted: the service was operationally
simple and rarely changed, so day-to-day developers did not need it cloned. It was the **only**
compose service built that way **at the end** — the pattern was once the platform default across 18 repos (corrected M257x iter-115; see the derivation above) — and `838d907` took it with it. Do not reach for it as a
precedent when adding a service.
