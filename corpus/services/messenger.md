# Messenger Service

## Role & Responsibility

Messenger is the **centralized notification service**. It sends and schedules transactional emails on behalf of every other service, using **Brevo** (formerly Sendinblue) as the delivery backend and **Liquid** templating for the bodies.

Other services don't talk to Brevo directly — they fire a Messenger RPC. Messenger then decides whether to send immediately, schedule for later, apply org-level whitelabel branding, or skip the message entirely based on per-domain notification rules (e.g., suppress repeat emails for the same simulation within 12 hours).

> **Default-off in local development.** Messenger is in the `messenger` Docker profile, not the default `graphql` profile. `make up` does **not** start it. Bring it up explicitly when iterating on notification flows.

## Architecture & Code Map

* **Codebase**: `messenger` (local) — repo `git@github.com:anthropos-work/messenger`
* **Language**: Go 1.25
* **Framework**: Connect-RPC
* **Email backend**: Brevo via `getbrevo/brevo-go v1.1.3`
* **Templating**: `osteele/liquid v1.8.1`
* **Ports**: `8200` (host) → `8200` (container, HTTP); `8201` (host) → `8201` (container, Connect-RPC)
* **Profile**: `messenger` only (NOT in default `graphql`). Opt-in.

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
    jobsimulations.go        Job-simulation completion / reminder rules
    organizations.go         Org invitation / membership rules
    organizations_db.go      Org DB lookups (read-only)
    whitelabel.go            Per-org whitelabel rendering (subject + body)
```

### Whitelabel rendering (2026-Q2)

Recent work in v0.34.0 added **whitelabel support**: when an org has custom branding (logo URL, custom invitation templates), Messenger renders subject and body separately so the Brevo send can include the org's logo and styling. The whitelabel lookup uses a **read-only Postgres connection** (`READONLY_DB_CONN`, formerly `COPILOT_DB_CONN`) so the rendering path doesn't contend with the write-heavy backend load.

## Interface Discovery

### Connect-RPC (`MessengerService`)

| Method | Purpose |
|--------|---------|
| `Send(message)` | Send an email immediately |
| `Schedule(message, schedule_for)` | Schedule a future email |
| `CancelScheduledMessage(id)` | Cancel a previously scheduled message |

Messages carry user info, template ID, and template params; the body is rendered through Liquid against those params before the Brevo send.

### What triggers Messenger?

Most messenger sends are reactive — driven by **Redis Streams** events from other services (`jobsimulation`, `cms`, `backend`). The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply suppression rules (e.g., 12-hour cooldown between repeated job-sim notifications). See `internal/flow/jobsimulations.go` for examples.

## Dependencies

* **Upstream consumers**: every service that needs to send mail — primarily `app`, with secondary callers via stream events from `cms`, `jobsimulation`, `skillpath`, `skiller`
* **Downstream**:
  * **Brevo API** — outbound email delivery (`BREVO_KEY`)
  * **PostgreSQL** — read-only `public` schema access for org / whitelabel lookups
  * **Redis** — Watermill stream subscriber + scheduled-message storage
  * **Sentinel** — for any RPC that requires authorization

> **Staging safety**: if you ever restore a production DB dump into local staging, `BREVO_KEY` **must be blanked** in `platform/.env` before `make up` to prevent real customer emails from going out. See [staging_from_dump.md](../ops/staging_from_dump.md).

## Local Development

### Run in Docker (opt-in)

```bash
cd platform
make up PROFILE=messenger
# or include alongside the default stack:
docker compose --profile graphql --profile messenger up --build -d
```

Messenger depends on `backend`, `cms`, `jobsimulation`, `skiller`, `skillpath` at startup (compose `depends_on`), so bringing it up implicitly brings the rest of the stack.

### Run natively

```bash
cd platform
make dev S=messenger
cd ../messenger
go run main.go
```

For local development, set `BREVO_KEY=""` to route through the **console sender** (`internal/messenger/console/`) instead of hitting Brevo — emails print to stdout.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8200` | HTTP port |
| `RPC_PORT` | `8201` | Connect-RPC port |
| `BREVO_KEY` | (empty) | Brevo API key. Empty → console sender. **MUST be empty for prod-dump staging.** |
| `REDIS_ADDR` | `redis:6379` | Redis address |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | `0` | Redis DB index for workers |
| `BACKEND_USERS_RPC_ADDR` | `http://backend:8083` | Backend RPC for user lookups |
| `CMS_RPC_ADDR` | `http://cms:8091` | CMS RPC |
| `JOBSIMULATION_RPC_ADDR` | `http://jobsimulation:8401` | Jobsimulation RPC |
| `SKILLER_RPC_ADDR` | `http://skiller:8086` | Skiller RPC |
| `SKILLPATH_RPC_ADDR` | `http://skillpath:8101` | Skillpath RPC |

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
