# Jobsimulation Service

## Role & Responsibility

Jobsimulation runs **AI-powered workplace simulations** end-to-end: it loads simulation definitions from CMS, hosts the interactive session (voice via LiveKit, chat, code, documents), records the interaction, generates post-session insights, and reports outcomes via Redis Streams to Skillpath and the App.

This is the user-facing "experience" service. Everything else (skills, content, auth, scoring) feeds it or consumes its outputs.

## Architecture & Code Map

* **Codebase**: `jobsimulation` (Local directory; repo `git@github.com:anthropos-work/jobsimulation`)
* **Language**: Go
* **Database**: PostgreSQL `jobsimulation` schema (via Ent + Atlas migrations)
* **Ports**: 8400 (GraphQL/HTTP), 8401 (Connect-RPC)
* **Profile**: `graphql` (default) and `jobsimulation`

### Key directories

```
cmd/                    Entrypoints
internal/
  graph/                GraphQL layer
    schemas/*.graphqls  simulation.graphqls is the main contract
  rpcsrv/               Connect-RPC server
  simulator/            Core simulation runtime
    manager/            Session lifecycle, interview extraction reports
  worker/               Background workers (Redis Streams + Asynq)
  ent/                  Generated Ent code
ent/schema/             Ent entity definitions
```

## Recent structural changes (2026-Q2)

* **Chronos removed**: session timeouts no longer scheduled via the chronos service. Replaced by **in-process [Asynq](https://github.com/hibiken/asynq)** (Redis-backed task queue, `hibiken/asynq v0.26.0`). See commit `09631fb2` ("remove Chronos references and update documentation to reflect Asynq integration for session timeout management") and PR `#395` (`feat/remove-chronos-and-realtime`).
* **Interview extraction pipeline added**: new entity `interview_extraction_results` (migrations `20260402145459`, `20260409131539`) stores per-session `user_report`, `manager_report`, and `summary` JSON blobs linked to a `session_id`. Exposed via CSV export with language arg (see `internal/simulator/manager/interview_report_csv*.go`).
* **READONLY_DB_CONNECTION env var added** (platform commit `05b4035`): a separate read-only connection string for reporting/extraction queries that should not contend with write traffic.

## Interface Discovery

* **GraphQL**: schemas at `internal/graph/schemas/simulation.graphqls`. Federated into the platform schema by Cosmo Router.
* **RPC**: `internal/rpcsrv` — consumed by Backend, Skillpath, Messenger via `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`.

### Direct dependencies (from compose `depends_on` + env)

* **Backend (app)** — user context, organization scoping
* **CMS** — simulation definitions, content, studio entities
* **Sentinel** — authz
* **Storage** — file uploads, recordings
* **Skiller** — skill metadata
* **Roadrunner** — code execution sandboxing (for code-task simulations)
* **PostgreSQL**, **Redis** — base infra

### External integrations

* **LiveKit** — primary voice engine (`LIVEKIT_HOST_URL`, `LIVEKIT_RECORDINGS_BUCKET_NAME`)
* **AWS Chime SDK** — video/camera/screensharing recording (`CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo`)
* **ElevenLabs** — legacy voice agents (`ELEVENLABS_TEMPLATE_AGENT_ID`, `ELEVENLABS_EU_TEMPLATE_AGENT_ID`); replaced by LiveKit + GPT Realtime for new sessions
* **AI providers** — via the shared `ai` library

### Redis Streams

* Producer: `jobsimulation` stream (session completed, insights generated)
* Consumer: `cms` (content events), `skiller` (skill updates)

## Local Development

### Run in Docker

```bash
cd platform
make up                           # default graphql profile
# or just jobsimulation:
make up PROFILE=jobsimulation
```

### Run natively

```bash
cd platform
make dev S=jobsimulation          # stops the docker container
cd ../jobsimulation
make setup                        # installs ent, atlas, gqlgen
make gen                          # regenerates Ent + GraphQL code
go run .
```

Make sure `.env` has the LiveKit + AWS credentials and that Postgres/Redis are reachable on `localhost`.

### Migrations

```bash
cd platform
make migrate S=jobsimulation
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — voice engines, recording, model routing
* [CMS](./cms.md) — content source
* [Dependency Map](../architecture/dependency_map.md) — RPC + event-stream relationships
