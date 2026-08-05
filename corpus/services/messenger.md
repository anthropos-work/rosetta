# Messenger Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **v9.0 "support-in-app"** (2026-08-04), the standalone `messenger` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"). It was folded in the
> same program — and on the same day — as [storage](./storage.md) and
> [customerio-sync](./customerio-sync.md). With those three gone, **[sentinel](./sentinel.md) is the
> only support service left running out-of-process.**
>
> Where everything went:
>
> * **Domain** — the whole mailer lives in `app/internal/messenger/` (`flow/`, `adapters/`, `sender/`,
>   `message/`, `brevo/`, `console/`, `aireadinessemail/`) — the Brevo client, the Liquid templates, the
>   whitelabel renderer and all **24 event handlers**, ported as-is.
> * **Events — `app` TAKES OVER messenger's own consumer group; it does not merge the handlers.**
>   This is the one place the messenger fold breaks the pattern the earlier folds used. `app` starts a
>   **second, dedicated `SubscriberServer`** on messenger's **own** Redis consumer group (the literal
>   `messenger`, because the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and
>   terraform never set `SERVICE_NAME` for it). Attaching to the existing group means Redis keeps the
>   cursor, so there is **no gap** at cutover. A dedicated server is also what makes it safe: messenger
>   registers three `AddSubscriber` calls for streams `app` already subscribes to, and on a shared
>   server those would have **silently replaced** app's own handlers (colony keys by stream name).
>   Boot verifies the group exists rather than silently creating a fresh one.
> * **Switch** — the whole block is gated by **`MESSENGER_ENABLED`** (`app/env_guards.go`). It is
>   **off when unset on a developer machine** — folding the mailer in deleted the "it's a different
>   binary with its own credentials" barrier, so consent is now explicit. In a **deployed** environment
>   unset is a **boot failure**, not a default-off: silently-unsent mail passes every health check.
>   An unparseable value is an error everywhere. `BREVO_KEY` is **required** whenever the switch is on.
> * **RPC** — `MessengerService` (`Send`, plus the two unimplemented scheduling stubs) is served on
>   `app`'s single RPC mux. Nothing calls it: messenger was the **last external caller** of that mux, so
>   `backend`'s Connect-RPC surface now has no out-of-process consumers at all. In-process senders make
>   plain Go calls.
> * **Infrastructure** — the ECS module was **deleted**. The ECR repo was preserved through the removal
>   (`removed { destroy = false }`) and is now **unmanaged in AWS** — it exists, terraform does not
>   know about it.
> * **Repo** — the `messenger` git repo still exists but is **frozen/legacy**; make changes in `app`.
>   Its Go module is still published, and `proto` still carries the `MessengerService` contract, but on
>   `main` **nothing imports the module**. It is still in `repos.yml` and still startable from the
>   `messenger` compose profile, as the rollback path.
>
> **Everything below this banner describes the standalone service** — which is still the code that runs
> if you start that profile, and still the best description of the templates, the Brevo client and the
> notification rules that were ported. Read it as a description of a rollback target, not of the
> default path.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

## Role & Responsibility

Messenger is the **centralized notification service**. It sends and schedules transactional emails on behalf of every other service, using **Brevo** (formerly Sendinblue) as the delivery backend and **Liquid** templating for the bodies.

Other services don't talk to Brevo directly — they fire a Messenger RPC. Messenger then decides whether to send immediately, apply org-level whitelabel branding, or skip the message entirely based on per-domain notification rules (e.g., it skips job-sim emails for stale/re-triggered sessions). (Scheduling RPCs exist in the proto but are not yet implemented — they return Unimplemented.)

> **Default-off in local development.** Messenger is in the `messenger` Docker profile — not the
> default profile, which is **`core`** (platform `0dab54d` renamed `graphql` → `core`; there is no
> `graphql` profile any more). `make up` does **not** start it, and since the v9.0 fold `backend` is
> already doing its work in-process.
>
> The same commit also **dropped messenger from the `all` profile**: `backend` now consumes
> messenger's own Redis consumer group, so running both puts **two consumers on one group**.

## Architecture & Code Map

* **Codebase**: `messenger` (local) — repo `git@github.com:anthropos-work/messenger`
* **Language**: Go 1.25
* **Framework**: Connect-RPC
* **Email backend**: Brevo via `getbrevo/brevo-go v1.1.3`
* **Templating**: `osteele/liquid v1.8.1`
* **Ports**: `8200` (host) → `8200` (container, HTTP); `8201` (host) → `8201` (container, Connect-RPC)
* **Profile**: `messenger` only (NOT in the default `core`, and dropped from `all` at platform `0dab54d`). Opt-in — and since the v9.0 fold, a rollback path rather than the default route.

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

Most messenger sends are reactive — driven by **Redis Streams** events from other services (`jobsimulation`, `cms`, `backend`). The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply staleness guards (e.g., for job-sim completions it drops the email if the session ended >2h ago, or has no end time and started >12h ago — `internal/flow/jobsimulations.go:140-151`). See `internal/flow/jobsimulations.go` for examples.

## Dependencies

* **RPC clients**: messenger still constructs four Connect-RPC clients — `cms`, `backend` (users + organizations), `skiller` and `jobsimulation` — but **all four addresses now resolve to the one `backend` mux** (`http://backend:8083` in compose). The `cms`, `jobsimulation` and `skiller` services those clients are named for no longer exist; their surfaces are registered on `app`'s RPC server. Skill-path notifications arrive as Redis Streams events on the `backend` subscriber (`OrgSkillPath*` handlers in `internal/flow/flow.go:72-87`), not via a direct Skillpath RPC.
* **Downstream**:
  * **Brevo API** — outbound email delivery (`BREVO_KEY`)
  * **PostgreSQL** — read-only `public` schema access for org / whitelabel lookups
  * **Redis** — Watermill stream subscriber + scheduled-message storage

> **Staging safety**: if you ever restore a production DB dump into local staging, `BREVO_KEY` **must be blanked** in `platform/.env` before `make up` to prevent real customer emails from going out. See [staging_from_dump.md](../ops/staging_from_dump.md).
>
> **Since v9.0 the kill switch moved.** Blanking `BREVO_KEY` is still necessary but is no longer the
> primary control, and restarting the `messenger` container no longer does anything on a default
> stack — there is no messenger container running. The mailer is `backend`, and the switch is
> **`MESSENGER_ENABLED`** (unset ⇒ off on a developer machine). `backend` also **refuses to boot** if
> the switch is on with an empty `BREVO_KEY`, so on a prod-dump stack leave the switch off rather than
> trying to neuter the key.

## Local Development

### Run in Docker (opt-in)

```bash
cd platform
# NB: `make up PROFILE=messenger` alone exits 1 — messenger declares `depends_on: backend`,
# which the `messenger` profile does not select, so compose rejects the project as invalid.
# Bring it up alongside the default stack instead:
docker compose --profile core --profile messenger up --build -d
```

Messenger's compose `depends_on` is now **redis, postgresql and `backend`** — the `cms` and
`jobsimulation` entries went with the services themselves; `skillpath` had already gone when it merged
into `app`. Bringing messenger up therefore implicitly brings the rest of the stack.

> **Don't run it alongside a `MESSENGER_ENABLED=true` backend.** Both attach to the same Redis
> consumer group, and entries claimed by one are not seen by the other — you get a coin-flip over which
> process sends each email. Pick one.

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
| `CMS_RPC_ADDR` | `http://backend:8083` | CMS RPC — **re-pointed at `backend`**; there is no `cms` container left to address. The old `http://cms:8091` is historical |
| `JOBSIMULATION_RPC_ADDR` | `http://backend:8083` | Jobsimulation RPC — **re-pointed at `backend`**, like all four values compose sets. The old `http://jobsimulation:8401` is historical |
| `SKILLER_RPC_ADDR` | `http://backend:8083` | Skiller RPC surface — served by `backend` since the skiller→app merge |
| ~~`SKILLPATH_RPC_ADDR`~~ | *(removed)* | **Gone from docker-compose** since skillpath was decommissioned into `app` ("skillpath-in-app", M502→M507) — only the residual `SKILLPATH_STREAM=skillpath` remains. Messenger never had a Skillpath RPC client anyway; skill-path data is read via the CMS client (`internal/flow/assignments.go:815`). |

> Values shown are what docker-compose injects. The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).

### The in-app variables (what you actually set now)

These are read by **`backend`**, not by this container:

| Variable | Read by | Description |
|----------|---------|-------------|
| `MESSENGER_ENABLED` | `backend` | Master switch for the folded mailer. **Unset ⇒ off** on a developer machine; **unset in a deployed environment ⇒ `backend` refuses to boot** (an unset switch would silently drop every email while every health check stayed green). An unparseable value is an error everywhere |
| `BREVO_KEY` | `backend` | **Required** when `MESSENGER_ENABLED` (or `CUSTOMERIO_SYNC_ENABLED`) is on — `backend` fails fast rather than starting a mailer that cannot send. One key now covers transactional mail, product tracking and the marketing-contact sync |
| `REDIS_ADDR` / `REDIS_STREAMS_INDEX` | `backend` | Where the takeover subscriber attaches. It joins the **existing** `messenger` consumer group; boot verifies the group exists rather than creating a fresh one |

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
