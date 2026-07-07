# Jobsimulation Service

> [!IMPORTANT]
> **This service holds NO simulation content.** "Jobsimulation" the *service* ≠ simulation *content*. It is a **runtime/session engine** that *runs* a simulation; the simulation **definition/blueprint** it runs — roles, sequences, tasks, validation criteria, knowledge assets, library categories — is **owned by the CMS service** (the `simulations` Directus collection + the Studio `StudioDocument`/`StudioTask` authoring model) and fetched **by ID** over Connect-RPC (`cms.GetSimulation`). Jobsimulation does **not** hold a `DIRECTUS_BASE_ADDR` of its own — all its content reads flow *through* CMS. See **[CMS](./cms.md)** for the content side. (This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).)

## Role & Responsibility

Jobsimulation runs **AI-powered workplace simulations** end-to-end: it loads simulation **definitions** from CMS (the content layer), hosts the interactive **session** (voice via LiveKit, chat, code, documents), records the interaction, generates post-session insights, and reports outcomes via Redis Streams to Skillpath and the App. Its own `jobsimulation` DB schema holds the **run/session state** (sessions, interactions, recordings, validation/anti-cheat results) — never the definition.

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
    schemas/*.graphqls  schema.graphqls is the main contract (also mutations/queries/subscriptions/activites)
  rpcsrv/               Connect-RPC server
  simulator/            Core simulation runtime
    manager/            Session lifecycle, interview extraction reports
  worker/               Asynq background workers (two pools: standard concurrency=10, real-time concurrency=25)
  ent/                  Generated Ent code (internal/ent/)
  ent/schema/           Ent entity definitions — source of truth (internal/ent/schema/)
```

## Recent structural changes (2026-Q2)

* **Chronos removed**: session timeouts no longer scheduled via the chronos service. Replaced by **in-process [Asynq](https://github.com/hibiken/asynq)** (Redis-backed task queue, `hibiken/asynq v0.26.0`). See commit `09631fb2` ("remove Chronos references and update documentation to reflect Asynq integration for session timeout management") and PR `#395` (`feat/remove-chronos-and-realtime`).
* **Interview extraction pipeline added**: new entity `interview_extraction_results` (migrations `20260402145459`, `20260409131539`) stores per-session `user_report`, `manager_report`, and `summary` JSON blobs linked to a `session_id`. Exposed via CSV export with language arg (see `internal/simulator/manager/interview_report_csv*.go`).
* **READONLY_DB_CONNECTION env var added** (platform commit `05b4035`): a separate read-only connection string for reporting/extraction queries that should not contend with write traffic.

## Interface Discovery

* **GraphQL**: schemas at `internal/graph/schemas/` (main contract: `schema.graphqls`). Federated into the platform schema by Cosmo Router.
* **RPC**: `internal/rpcsrv` — consumed by Backend, Skillpath, Messenger via `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`.

### Direct dependencies (from compose `depends_on` + env)

* **Backend (app)** — user context, organization scoping
* **CMS** — simulation definitions, content, studio entities. **Jobsimulation reads Directus content *through* CMS over RPC — it does NOT hold a `DIRECTUS_BASE_ADDR`/`DIRECTUS_TOKEN` of its own.** So the M23 content cutover (re-pointing CMS's `DIRECTUS_BASE_ADDR` at the per-stack Directus) carries jobsimulation's content reads to local automatically; no jobsimulation env change is needed.
* **Sentinel** — authz
* **Storage** — file uploads, recordings
* **Skiller RPC surface** — skill metadata; served by **Backend (app)** since the skiller→app merge (July 2026): `SKILLER_RPC_ADDR=http://backend:8083`
* **Roadrunner** — code execution sandboxing (for code-task simulations)
* **PostgreSQL**, **Redis** — base infra

### External integrations

* **LiveKit** — primary voice engine (`LIVEKIT_HOST_URL`, `LIVEKIT_RECORDINGS_BUCKET_NAME`)
* **AWS Chime SDK** — video/camera/screensharing recording (`CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo`)
* **ElevenLabs** — voice agents still used in the call/reply pipeline (`ELEVENLABS_TEMPLATE_AGENT_ID`, `ELEVENLABS_EU_TEMPLATE_AGENT_ID`); new sessions increasingly use LiveKit + OpenAI Realtime (gated by the `flag_use_realtime_openai` PostHog flag)
* **AssemblyAI** — EU voice transcription for call recordings (`ASSEMBLYAI_API_KEY`)
* **Bunny.net** — video stream hosting / tokenized playback (`BUNNY_REC_STREAM_API_KEY`, `BUNNY_TOKEN_HASH_KEY`)
* **PostHog** — feature flags + telemetry (`POSTHOG_API_KEY`); the OpenAI Realtime voice path is gated by the `flag_use_realtime_openai` PostHog flag
* **AI providers** — via the shared `ai` library

### Redis Streams

* Producer: `jobsimulation` stream (session completed, insights generated)
* Consumer (subscribes to): `cms` (content events) and `roadrunner` (code-execution events)

Redis Streams consumption is handled by the colony pubsub `SubscriberServer` wired up in `cmd/root.go`, not by `internal/worker/` (which is Asynq-only).

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
make setup                        # installs ent, atlas, gqlgen, goverter
make gen                          # regenerates Ent + Goverter + gqlgen
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
