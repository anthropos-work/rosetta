# Jobsimulation Service

> [!IMPORTANT]
> **This service holds NO simulation content.** "Jobsimulation" the *service* ≠ simulation *content*. It is a **runtime/session engine** that *runs* a simulation; the simulation **definition/blueprint** it runs — roles, sequences, tasks, validation criteria, knowledge assets, library categories — is **owned by the CMS service** (the `simulations` Directus collection + the Studio `StudioDocument`/`StudioTask` authoring model) and fetched **by ID** over Connect-RPC (`cms.GetSimulation`). Jobsimulation does **not** hold a `DIRECTUS_BASE_ADDR` of its own — all its content reads flow *through* CMS. See **[CMS](./cms.md)** for the content side. (This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).)

## Role & Responsibility

Jobsimulation runs **AI-powered workplace simulations** end-to-end: it loads simulation **definitions** from CMS (the content layer), hosts the interactive **session** (voice via LiveKit, chat, code, documents), records the interaction, generates post-session insights, and reports outcomes via Redis Streams to the App (which now hosts the in-process skill-path engine, formerly the standalone skillpath service). Its own `jobsimulation` DB schema holds the **run/session state** (sessions, interactions, recordings, validation/anti-cheat results) — never the definition.

This is the user-facing "experience" service. Everything else (skills, content, auth, scoring) feeds it or consumes its outputs.

## Architecture & Code Map

* **Codebase**: `jobsimulation` (Local directory; repo `git@github.com:anthropos-work/jobsimulation`)
* **Language**: Go
* **Database**: PostgreSQL `jobsimulation` schema (via Ent + Atlas migrations)
* **Ports**: 8400 (GraphQL/HTTP), 8401 (Connect-RPC) — **as deployed by the platform**, which sets `PORT=8400` / `RPC_PORT=8401` and publishes `8400:8400` / `8401:8401` (`platform/docker-compose.yml`). Note the **repo's own defaults differ**: with those env vars unset `cmd/root.go` falls back to `8080`/`8081` (and the Dockerfiles `EXPOSE 8080`), which is what the in-repo `CLAUDE.md` documents. Both are correct in their own context — use 8400/8401 for anything driven through `platform`, and add the stack offset for a `dev-N`/`demo-N`.
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
* **RPC**: `internal/rpcsrv` — consumed by Backend (incl. the in-process skill-path engine) and Messenger via `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`.

> **Session/result READ-MODEL — this doc is not the home for it.** Two things a reader looking for "how does a
> played session render?" will not find here. (1) The **player** result page `/sim/<slug>/result/<sessionId>` is a
> **persisted read**, not a live recompute — `internal/graph/queries.resolvers.go:70` does plain Ent SELECTs over
> `validation_attempt_results`, so a seeded result fan-out renders a full result. (2) The **manager** view does
> **not** read this service's `sessions` table at all — it reads an `app`-side MIRROR, `public.local_jobsimulation_sessions`
> (the analog of skill-path's `local_skill_path_session`). Seed the runtime rows only and the manager scoreboard
> is blank. Full route-by-route treatment lives in
> [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md); the write side is
> [`../ops/demo/session-clone-spec.md`](../ops/demo/session-clone-spec.md).

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

## Startup contract — read this before diagnosing a crash (M217)

**The cobra ROOT command's `RunE` *is* the server.** There is **no `serve` and no `run` subcommand.**

- The image is `ENTRYPOINT ["./application"]` with **no CMD**; docker-compose passes **no `command:`**.
- Running the binary with **zero arguments is correct** — that starts the server.
- The optional subcommands are `aggregate`, `clone-session`, `test-command`, `validate`. **None of them starts
  the service.**

> ⚠️ **`command: serve` would BREAK it** — cobra would reject `unknown command "serve"` and exit 1. The repo's own
> `CLAUDE.md` documents `go run . serve`; **that command does not exist.** (It is a platform repo — don't trust
> it here, and don't edit it.)

### "It printed the CLI help" means an INIT ERROR — not a missing subcommand

The root command sets neither `SilenceUsage` nor `SilenceErrors`. So **any** error returned from `RunE` makes
cobra print `Error: …` **followed by the full usage/help block**, then exit 1.

**That usage block is a symptom of a failed init, not of a wrong command.** It was misread as "the container
needs a subcommand" for an entire release cycle, and the proposed fix would have broken the service.

**Always read the FIRST line of `docker logs`, never the help block:**

```bash
docker logs demo-<N>-jobsimulation-1 2>&1 | head -3
# Error: can't init AI: can't load AWS config: failed to load shared config file, ...
```

### The `$HOME/.aws/credentials` landmine (why it died in every demo)

`docker-compose.yml` binds `$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the
file. **When the host path does not exist, Docker auto-creates it as an empty DIRECTORY.** The container then
sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s `config.LoadDefaultConfig()` **opens it
successfully** (opening a directory succeeds!) before failing `EISDIR` on the read — so it is *not* skipped as
an unreadable file. That error propagates out of `ai.NewAIManager` → the root `RunE` → cobra's usage block →
`exit 1`.

**With the path simply absent, `LoadDefaultConfig` returns `nil`.** The mount is the bug.

- **On a workstation** with a real `~/.aws/credentials` file, it works — which is why this never showed up in
  local dev and only bit a fresh Linux box.
- **In a demo/dev stack**, rext's **generated compose override drops the bind** (`volumes: !reset null` on the
  demo path; an `!override`-tagged empty list on the dev path). Zero platform-repo edits. A stack carries **no
  AWS credentials at all**, so that mount could only ever *be* the broken empty directory.

> ⚠️ **A bare `volumes: []` does NOT remove it** — compose *merges* volume sequences and the inherited bind
> survives. Only the `!reset` / `!override` tags remove it. Verified against the compose binary.

**Downstream while it is dead:** the AI-Simulations surface is gone; its GraphQL subgraph errors; the
`pt-aisim-chat-launch` playthrough cannot pass; no session-completed events reach the Redis stream, so the
skill-path engine (now in `app`) never sees completions. And it is the service behind the nameless *"1 check(s) FAILED"* the
bring-up's autoverify used to report.

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
