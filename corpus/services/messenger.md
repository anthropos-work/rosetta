# Messenger Service

## Role & Responsibility

Messenger is the **centralized notification service**. It sends and schedules transactional emails on behalf of every other service, using **Brevo** (formerly Sendinblue) as the delivery backend and **Liquid** templating for the bodies.

Other services don't talk to Brevo directly — they fire a Messenger RPC. Messenger then decides whether to send immediately, apply org-level whitelabel branding, or skip the message entirely based on per-domain notification rules (e.g., it skips job-sim emails for stale/re-triggered sessions). (Scheduling RPCs exist in the proto but are not yet implemented — they return Unimplemented.)

> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning. Re-derived at platform `0dab54d` / `app` `9d00a313` v1.367.0 /
> `messenger` `a0ec933`.
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:61`, `:62`) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1387`, wired at `:1423` with `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap (`:1330-1340`). The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1362-1365`) |
> | **local** | still in `repos.yml:21-23` and still defined in compose (`docker-compose.yml:156`, `messenger` profile at `:195`) — startable, as the rollback path. `0dab54d` also dropped it from the `all` profile, because running both puts **two consumers on one group** |
>
> **Everything below this banner describes the standalone service**, which is still the code that
> runs when you start the profile, and is still where the templates, the Brevo client and the
> notification rules live. Read it as the description of a rollback target, not of the default path.

> **Default-off in local development.** Messenger is in the `messenger` Docker profile, not the
> default `core` profile — `core`, not `graphql`: platform `0dab54d` **renamed** it, and there is no
> `graphql` profile any more. `make up` does **not** start it, and since the fold `backend` is already
> doing its work in-process.

## Architecture & Code Map

* **Codebase**: `messenger` (local) — repo `git@github.com:anthropos-work/messenger`
* **Language**: Go 1.25
* **Framework**: Connect-RPC
* **Email backend**: Brevo via `getbrevo/brevo-go v1.1.3`
* **Templating**: `osteele/liquid v1.8.1`
* **Ports**: `8200` (host) → `8200` (container, HTTP); `8201` (host) → `8201` (container, Connect-RPC)
* **Profile**: `messenger` **only** — `profiles: [messenger]` (`docker-compose.yml:195`, derived from `docker-compose.yml` @ platform `0dab54d`). Not in the default `core`, and `0dab54d` also dropped it from `all` (two consumers on one Redis group). Opt-in, and since the v9.0 fold it is a rollback path rather than the default route.

### Key directories

```
cmd/                         Entrypoints
internal/
  rpcsrv/rpcsrv.go           Connect-RPC handler (Send, Schedule, Cancel)
  messenger/
    messenger.go             Top-level Messenger dispatcher
    brevo/                   Brevo client
    console/                 Console sender for local dev
    message/                 Message types + Liquid rendering
  flow/
    flow.go                  Notification-flow dispatcher
    assignments.go           Assignment notification rules
    cms.go                   CMS studio-task simulation completion rules (success/failure email)
    jobsimulations.go        Job-simulation completion / reminder rules
    organizations.go         Org invitation / membership rules
    organizations_db.go      Org DB lookups (read-only)
    whitelabel.go            Per-org whitelabel rendering (subject + body)
```

### Whitelabel rendering (2026-Q2)

Recent work in v0.34.0 added **whitelabel support**: when an org has custom branding (logo URL, custom invitation templates), Messenger renders subject and body separately so the Brevo send can include the org's logo and styling. The org lookup uses a **read-only Postgres connection** (`READONLY_DB_CONNECTION`, formerly `COPILOT_DB_CONNECTION` — see `cmd/root.go:147`) so the rendering path doesn't contend with the write-heavy backend load.

## Interface Discovery

### Connect-RPC (`MessengerService`)

| Method | Purpose | Status |
|--------|---------|--------|
| `Send(message)` | Send an email immediately | Implemented |
| `Schedule(message, schedule_for)` | Schedule a future email | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |
| `CancelScheduledMessage(id)` | Cancel a previously scheduled message | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |

Messages carry user info, template ID, and template params; the body is rendered through Liquid against those params before the Brevo send.

### What triggers Messenger?

Most messenger sends are reactive — driven by **Redis Streams** events on the `jobsimulation`, `cms` and `backend` streams. The stream *names* outlived the services: since the merges they are published from inside `app` (e.g. the `CMS_STREAM` publisher at `app/main.go:1095`, and the whole subscriber stream binding at `:1478-1484` @ `app` `9d00a313` v1.367.0), so there is no separate producer service in compose behind any of them. The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply staleness guards (e.g., for job-sim completions it drops the email if the session ended >2h ago, or has no end time and started >12h ago — `internal/flow/jobsimulations.go:140-151`). See `internal/flow/jobsimulations.go` for examples.

## Dependencies

* **RPC clients**: messenger still constructs four Connect-RPC clients — CMS, backend users + organizations, skiller, and jobsimulation — but at platform `0dab54d` **all four addresses resolve to the one `backend` mux** (`http://backend:8083`, `docker-compose.yml:173`, `:174`, `:176`, `:183`, under compose's own comment at `:171-172`). The `cms` and `jobsimulation` services those clients were named for no longer exist in compose; their surfaces are registered on `app`'s RPC server. Skill-path notifications arrive as Redis Streams events on the `backend` subscriber (`OrgSkillPath*` handlers in `internal/flow/flow.go:72-87`), not via a direct Skillpath RPC.
* **Downstream**:
  * **Brevo API** — outbound email delivery (`BREVO_KEY`)
  * **PostgreSQL** — read-only `public` schema access for org / whitelabel lookups
  * **Redis** — Watermill stream subscriber + scheduled-message storage

> **Staging safety**: if you ever restore a production DB dump into local staging, `BREVO_KEY` **must be blanked** in `platform/.env` before `make up` to prevent real customer emails from going out. See [staging_from_dump.md](../ops/staging_from_dump.md).

## Local Development

### Run in Docker (opt-in)

```bash
cd platform
# NB: `make up PROFILE=messenger` alone EXITS 1 — messenger declares `depends_on: backend`,
# which the `messenger` profile does not select, so compose rejects the project as invalid.
# Bring it up alongside the default stack instead:
docker compose --profile core --profile messenger up --build -d
```

At platform `0dab54d` messenger's `depends_on` is **redis, postgresql and `backend`** (`docker-compose.yml:186-192`), so bringing it up implicitly brings the rest of the stack. The `cms` and `jobsimulation` entries went with the services themselves at `d11a403`; `skillpath` had already gone when it merged into `app`.

### Run natively

```bash
cd platform
make dev S=messenger
cd ../messenger
go run main.go
```

For local development, set `BREVO_KEY=""` to route through the **console sender** (`internal/messenger/console/`) instead of hitting Brevo — emails print to stdout.

## Environment Variables

| Variable | Value (compose) | Description |
|----------|---------|-------------|
| `PORT` | `8200` | HTTP port |
| `RPC_PORT` | `8201` | Connect-RPC port |
| `BREVO_KEY` | (empty) | Brevo API key. Empty → console sender. **MUST be empty for prod-dump staging.** |
| `REDIS_ADDR` | `redis:6379` | Redis address |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | `0` | Set in docker-compose (=0) but NOT read by the code — there is no worker pool / separate worker Redis index; only `REDIS_STREAMS_INDEX` is consumed (`cmd/root.go:107`). |
| `BACKEND_USERS_RPC_ADDR` | `http://backend:8083` | Backend RPC for user lookups |
| `CMS_RPC_ADDR` | `http://backend:8083` | CMS RPC — **re-pointed at `backend`** (`docker-compose.yml:174` @ platform `0dab54d`). **M809 has landed**; there is no `cms` container left to address. The earlier `http://cms:8091` was true at `2adcf71` only. `app`'s own comment at `app/main.go:1205-1211` (@ `b948604` v1.366.0) still says *"additive + DORMANT … until the M809 re-point"* and is **stale in `app`** — grade the address against compose |
| `JOBSIMULATION_RPC_ADDR` | `http://backend:8083` | Jobsimulation RPC — **re-pointed at `backend`** (`docker-compose.yml:176` @ platform `0dab54d`), like all four values compose sets. The earlier `http://jobsimulation:8401` was true at `2adcf71` only; the husk container is gone |
| `SKILLER_RPC_ADDR` | `http://backend:8083` | Skiller RPC surface — served by `backend` since the skiller→app merge |
| ~~`SKILLPATH_RPC_ADDR`~~ | *(removed)* | **Gone from docker-compose** since skillpath was decommissioned into `app` ("skillpath-in-app", M502→M507) — only the residual `SKILLPATH_STREAM=skillpath` remains. Messenger never had a Skillpath RPC client anyway; skill-path data is read via the CMS client (`internal/flow/assignments.go:828`, in `getSkillPath`). |

> Values shown are what docker-compose injects. The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).

## Testing

```bash
cd messenger
go test ./...
```

The flow handlers have unit tests (`assignments_test.go`, `jobsimulations_test.go`, `organizations_test.go`) covering the suppression / whitelabel branches.

## Related Documentation

* [Backend (app)](./backend.md) — main caller
* [staging_from_dump.md](../ops/staging_from_dump.md) — outbound-email kill switch
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)
