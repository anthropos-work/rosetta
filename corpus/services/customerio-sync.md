# CustomerIO Sync Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **customerio-sync-in-app** fold (2026-08-04, the v9.0 "support-in-app" program that also
> took in [messenger](./messenger.md) and [storage](./storage.md)), the marketing-contact sync runs
> **inside the `app` monolith**.
>
> **The name is now doubly historical.** It was already misleading — the destination has been **Brevo**,
> not Customer.io, since the platform moved off Customer.io; the repo name outlived the vendor. Now the
> service is gone too.
>
> Where everything went:
>
> * **Domain** — `app/internal/customeriosync/` (`sync.go`, `store.go`, `sync_query.sql`): the same
>   one-directional push, PostgreSQL `public` → **Brevo** marketing contacts.
> * **Schedule** — it runs on `app`'s **asynq scheduler** on a **10-minute** period, replacing the
>   standalone's own in-process ticker. The sync window became **stateless** in the move: the
>   standalone held `lastSyncTime` in memory, which is neither shared nor durable across a
>   multi-replica worker that restarts on every deploy. Each run now reads a fixed overlap window
>   wider than the schedule, so a missed or crashed run is covered by the next one. Pushing a contact
>   is idempotent (`CreateContact` with `UpdateEnabled`), so the overlap costs nothing.
> * **The `public.customer_io_sync_table` view is gone, and that matters beyond this service.** The
>   standalone read a DB view that was **never under migration control** — its DDL existed only in the
>   production database and had been hand-edited. It was also the **only cross-schema dependency
>   `public` had**, which made it a silent casualty of any legacy-schema drop: `DROP SCHEMA skiller
>   CASCADE` would have taken the view, and this sync, with it. The logic now lives in
>   `internal/customeriosync/sync_query.sql`, computed against final `public` tables (plus the two
>   Directus content joins, which stay — there is no `public` mirror of simulation and skill-path
>   titles).
> * **Manual runs** — `go run ./cmd/customerio-resync` in `app`. That path needs **no** switch (you
>   typed the tool's name, which is consent), and its `--dry-run` constructs no Brevo client at all.
> * **Switch** — gated by **`CUSTOMERIO_SYNC_ENABLED`** (`app/env_guards.go`), the same strict switch
>   as `MESSENGER_ENABLED`: **unset ⇒ off** on a developer machine (folding it in removed the
>   "different binary, different credentials" barrier, and holding a credential is not consent to use
>   it); **unset in a deployed environment ⇒ `backend` refuses to boot**; an unparseable value is an
>   error everywhere. `BREVO_KEY` is required when it is on.
> * **Infrastructure** — the terraform module was **fully deleted** and the ECR repo destroyed. Unlike
>   messenger and storage, **there is no rollback path left** for this one.
> * **Repo** — the `customerio-sync` git repo still exists but is **frozen/legacy**; make changes in
>   `app`. It was never cloned by `make init`, so nothing changed there.
>
> **This is the one that used to self-trigger.** The other folded services stopped mattering the moment
> their callers were re-pointed; this one ran its own ticker, so stopping the deployment is what ended
> it. That has happened. If anyone ever redeploys the standalone alongside a
> `CUSTOMERIO_SYNC_ENABLED=true` backend, the two pushes are **idempotent** — the contacts do not
> corrupt — but the Brevo rate spend doubles.
>
> **Local compose has not caught up.** At platform `0dab54d` the `customerio-sync` service is still
> declared and is still in the **`all`** profile — so `make up PROFILE=all` starts it, and the
> `git@github.com:…#main` build context means it builds the **frozen** repo at whatever `main` is. It
> is not in the default `core` profile.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

## Role & Responsibility

CustomerIO Sync is a **background data-sync service** that pushes user and organization data from the Anthropos backend database into [Customer.io](https://customer.io/) for marketing automation, lifecycle email campaigns, and product analytics.

It's a one-directional pipeline: PostgreSQL `public` schema → Customer.io API. No inbound traffic, no consumers inside the platform.

## Architecture

* **Repo**: `git@github.com:anthropos-work/customerio-sync` (private) — **frozen**
* **Language**: Go
* **Local port**: 8080 (HTTP — health/metrics)
* **Profile**: `customerio-sync` and `all` (NOT in the default profile, which is `core`, not `graphql` — platform `0dab54d` renamed it)
* **Build pattern**: **unique among Anthropos services** — Docker builds it directly from the GitHub URL, the repo is not cloned locally by `make init`.

### Compose definition

```yaml
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
  profiles: [customerio-sync, all]
  depends_on:
    postgresql: { condition: service_healthy }
```

Note `context: git@github.com:...#main` — Docker BuildKit clones the repo at build time, no local checkout needed. This works because the build runs inside an SSH-agent-forwarded context (`ssh: ["default"]`) with `$GH_PAT` available.

### Dependencies

* **PostgreSQL** (`public` schema, read access via `DB_CONNECTION_BACKEND`)
* **Customer.io API** (external — credentials live in `platform/.env`)

## Interface Discovery

The service does not expose business APIs to other platform services — it's a pure sync worker. Port 8080 likely serves a health/metrics endpoint.

For protocol and field-mapping details, see the repo:

```bash
gh repo clone anthropos-work/customerio-sync
```

## Local Development

This service is **off by default**. To run it locally:

```bash
cd platform
docker compose --profile customerio-sync up --build -d customerio-sync
```

You'll need Customer.io API credentials in `platform/.env`. For most local-development tasks you do not need this service.

## Production

**No longer deployed.** The sync runs inside `backend` on the asynq scheduler, gated by
`CUSTOMERIO_SYNC_ENABLED`. The standalone's terraform module was deleted and its ECR repo destroyed —
there is no rollback path. (Historically: it ran as its own ECS task, configured via `platform/.env`
and Terraform-managed secrets.)

## Related Documentation

* [Backend (app)](./backend.md) — source of user/org data
* [Service Taxonomy](../architecture/service_taxonomy.md) — orchestration profiles
* [External Services](../architecture/external_services.md) — Customer.io as an integrated SaaS

## Notes

The "build from GitHub URL" pattern is intentional: this service is operationally simple and rarely changes, so day-to-day developers do not need it cloned. If you need to iterate on it, clone it as a sibling of `platform/` and add it to `repos.yml` temporarily.
