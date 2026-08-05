# CustomerIO Sync Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4`).
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`).
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

It runs on `app`'s **shared** analytics pool (`copilotDB`), not a pool of its own — one query every
ten minutes does not earn a standing allocation against the platform's connection budget
(`app/main.go:393-396`).

### Compose definition — **HISTORICAL**

This block was deleted at `838d907`. It is kept because the build pattern was unique in the platform
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
(`ssh: ["default"]`) with `$GH_PAT` available. **It was the only service built that way, and the
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
**exits 0** and starts only the always-on floor (`postgresql`, `redis`, `sentinel`), which looks like
a live stack.

The domain is inside `backend` and **off unless you say otherwise**:

```bash
# In platform/.env — NOT in docker-compose.yml, which deliberately pins no value:
#   CUSTOMERIO_SYNC_ENABLED=true
#   BREVO_KEY=<a real key>
# Then `make up`. Both are required: app log.Fatalf's if the switch is on with an
# empty key (app/main.go:295), because the sender has no console fallback.
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
is **fatal** rather than off (`app/main.go:284`), so a deployed environment must state its intent.
**Scope note:** whether its own ECS task / image / terraform module have been torn down was **not
measured** in this pass — the fold and the container deletion are local-compose and `app`-source
facts. Do not read them as the production teardown.

## Related Documentation

* [Backend (app)](./backend.md) — the host process, and the source of user/org data
* [Service Taxonomy](../architecture/service_taxonomy.md) — orchestration profiles
* [External Services](../architecture/external_services.md) — Customer.io as an integrated SaaS

## Notes

The "build from GitHub URL" pattern was intentional while it lasted: the service was operationally
simple and rarely changed, so day-to-day developers did not need it cloned. It was the **only**
compose service built that way, and `838d907` took the pattern with it — do not reach for it as a
precedent when adding a service.
