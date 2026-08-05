# Jobsimulation Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"jobsim-in-app"** program (platform milestone **M810** tracks the final teardown), the standalone
> `jobsimulation` Go microservice has been **merged into the `app` monolith** (the service the platform calls
> "backend"). Jobsimulation no longer runs as a separate service **in production**
> (`jobsimulation/terraform/main.tf:40` `service_desired_count = 0`), and its subgraph is gone from the
> supergraph.
>
> **✅ The husk is GONE locally too (measured at platform `0dab54d`).** There is no `jobsimulation` compose
> service, no `jobsimulation` entry in `repos.yml` (6 entries: app, sentinel, storage, messenger,
> next-web-app, studio-desk) and no `jobsimulation` profile. Platform **`d11a403`** (2026-08-03) deleted
> both in one commit — its `repos.yml` diff removes `- name: cms`, `- name: jobsimulation` **and**
> `- name: roadrunner`.
> *This banner used to read "**but locally the husk still starts**", and it was right at `2adcf71`:
> `docker-compose.yml:83` @ that ref defined a `jobsimulation` service with
> `profiles: [graphql, jobsimulation, all]` (`:140`), `graphql` was the default (`Makefile:10`
> `PROFILE ?= graphql` **at that ref**), and `repos.yml:17-19` @ `2adcf71` still listed the repo (marked
> `migrations: false # legacy`).* The **GitHub repo was archived 2026-07-31**. State: **frozen legacy repo,
> no local container, no clone entry**; what **M810** tears down is the *production* rollback path
> (`module.jobsimulation_euwest1`), not a local husk. See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> This is the same pattern as the earlier [skiller-in-app](./skiller.md) and
> [skillpath-in-app](./skillpath.md) merges.
>
> Where everything went:
>
> * **Domain / engine** — the whole simulation engine lives in `app/internal/jobsimulation/` (activity, agent,
>   ai, analytics, anticheat, bunny, calls, graph, inbound, recording, …), constructed by the single wiring
>   entry point **`app/internal/jobsimwiring/wiring.go`**.
> * **Data** — the 23 session/run tables (`sessions`, `actors`, `interactions`, `validation_*`, `anticheat_*`,
>   `recordings`, `chime_recordings`, `code_submissions`, …) were re-created in the **`public` schema** by
>   `app/terraform/migrations/20260722081626_jobsim_data_model.sql`. **Most kept their names — but the
>   headline one did NOT:** the very next migration, `20260722104506.sql`, creates
>   `job_simulation_sessions` (`:2`) and `DROP TABLE "sessions"` (`:79`). **`public.sessions` does not
>   exist**; the session table is `public.job_simulation_sessions`. The old
>   `jobsimulation` DB schema is **legacy — no longer authoritative**.
> * **RPC** — `JobSimulationService` is served on `app`'s single RPC mux. `messenger` reaches it at
>   `JOBSIMULATION_RPC_ADDR=`**`http://backend:8083`** locally (`docker-compose.yml:176` @ platform
>   `0dab54d`); `http://backend.internal.anthropos:8081` in production.
>   **The local re-point onto `app` — M809 — HAS landed**, and there is no husk container left to reach:
>   `0dab54d`'s compose declares **eight** services — ten effective, once `include: common.yml` adds
>   the `postgresql`/`redis` floor — and `jobsimulation` is not one of them.
>   (`http://jobsimulation:8401` was true at `2adcf71`.) The in-app edge is registered at
>   `app/main.go:1204` (@ `b948604` v1.366.0). `app` itself makes
>   **no** outbound jobsim RPC — those are in-process calls now.
> * **GraphQL** — the jobsimulation subgraph was removed from the federation; its types/queries are served by
>   `app`'s sole `backend` subgraph.
> * **Events** — `app` owns the `JOBSIMULATION_STREAM` subscriber. The ported engine's handlers are merged onto
>   app's **existing** subscriber via `.AddHandler(...)` (a second `AddSubscriber` for the same stream would
>   silently overwrite the first — colony keys by stream name).
> * **Dependencies that changed** — chronos is gone (session timers are Asynq jobs); roadrunner is gone (the
>   in-process Judge0 runner executes code directly via `JUDGE0_BASE_URL`); the `BACKEND_USERS_RPC_ADDR`
>   loopback is replaced by an in-process users reader.
> * **Infrastructure** — `module.jobsimulation_euwest1` is **still declared** in
>   `infrastructure/terraform/production/services.tf` as the **rollback path** and takes no traffic. It still
>   **owns the LiveKit and Chime recording S3 buckets**, which `backend` reuses by literal name — move
>   ownership before destroying it. Teardown is **M810**.
> * **Repo** — the `jobsimulation` git repo still exists but is **frozen/legacy**; make changes in `app`.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

> [!IMPORTANT]
> **This service holds NO simulation content.** "Jobsimulation" the *service* ≠ simulation *content*. It is a **runtime/session engine** that *runs* a simulation; the simulation **definition/blueprint** it runs — roles, sequences, tasks, validation criteria, knowledge assets, library categories — is **owned by the CMS service** (the `simulations` Directus collection + the Studio `StudioDocument`/`StudioTask` authoring model) and fetched **by ID** — since the merge this is an **in-process call** into the folded cms domain (it was
> `cms.GetSimulation` over Connect-RPC). The jobsim domain still holds no `DIRECTUS_BASE_ADDR` of its own — all
> its content reads flow *through* the cms domain. See **[CMS](./cms.md)** for the content side. (This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).)

## Role & Responsibility

Jobsimulation runs **AI-powered workplace simulations** end-to-end: it loads simulation **definitions** from CMS (the content layer), hosts the interactive **session** (voice via LiveKit, chat, code, documents), records the interaction, generates post-session insights, and reports outcomes via Redis Streams to the App (which now hosts the in-process skill-path engine, formerly the standalone skillpath service). Its run/session state (sessions, interactions, recordings, validation/anti-cheat results) now lives in the shared **`public`** schema — never the definition.

This is the user-facing "experience" service. Everything else (skills, content, auth, scoring) feeds it or consumes its outputs.

## Architecture & Code Map

* **Codebase**: `jobsimulation` — repo `git@github.com:anthropos-work/jobsimulation` (archived 2026-07-31). **Not cloned by `make init`**: no `repos.yml` entry since `d11a403`. Clone it by hand to read the pre-merge source; the live code is `app/internal/jobsimulation/`
* **Language**: Go
* **Database**: ~~PostgreSQL `jobsimulation` schema~~ → the 23 run-state tables live in **`public`**, created by **`app`**'s migrations (`app/terraform/migrations/20260722081626_jobsim_data_model.sql`). The legacy `jobsimulation` schema is **not authoritative** — consistent with the **Data** bullet, :31-38 above
* **Ports**: **8080 (GraphQL/HTTP), 8081 (Connect-RPC) — the binary's own defaults**, and now the only ones there are: `cmd/root.go:77` `cmp.Or(os.Getenv("PORT"), "8080")` / `:78` `cmp.Or(os.Getenv("RPC_PORT"), "8081")` (the Dockerfiles `EXPOSE 8080`), which is what the in-repo `CLAUDE.md` documents. The **8400 / 8401** pair quoted all over this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=8400` (`:113`) / `RPC_PORT=8401` (`:119`) and published `8400:8400` / `8401:8401` (`:93-94`) — **at `2adcf71`**. At `0dab54d` there is no `jobsimulation` service, so nothing sets those values and nothing is published; **8400/8401 are historical, not an address you can reach**, with or without a `dev-N`/`demo-N` offset. The engine's live HTTP/GraphQL surface is `backend`'s.
* **Profile**: **none — there is no `jobsimulation` compose service.** Deleted by platform `d11a403` (2026-08-03), the compose clean-up that followed the fold; the line that stood here named the `graphql` profile, which `0dab54d` renamed `core`, for a service that had already been removed. Historical only (corrected M257x iter-68)

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

* **GraphQL**: schemas at `internal/graph/schemas/` (main contract: `schema.graphqls`). ~~Federated into the platform schema by Cosmo Router~~ — **the jobsimulation subgraph is folded into `backend`**; the supergraph is one subgraph (`backend.graphqls`).
* **RPC**: `internal/rpcsrv` — reached **in-process** by Backend (incl. the in-process skill-path engine), and over the wire by **`messenger` alone**, the only service left that reads `JOBSIMULATION_RPC_ADDR`. At platform `0dab54d` that value is **`http://backend:8083`**, like all four addresses compose sets — and compose sets them only on `messenger`; `d11a403` dropped `JOBSIMULATION_RPC_ADDR` from `backend` outright, having verified zero reads in `app`. **M809 has landed** and there is no husk container left to resolve to. `app` registers its own in-app `JobSimulationService` handler (`app/main.go:1204` @ `app` `b948604` v1.366.0).
  > **This line used to say the opposite, emphatically — keep the note (M257x iter-60).** Until `2adcf71` it read *"That address is **CURRENT, not stale text**"*, and it was **right at that ref**: only `SKILLER_RPC_ADDR` had been re-pointed then. A refutation is a measurement and expires exactly like the claim it refuted — and anti-repair wording is the kind that survives readings, because it looks already-adjudicated. See [`platform-alignment.md`](../ops/platform-alignment.md) §5 rule 31.

> **Session/result READ-MODEL — this doc is not the home for it.** Two things a reader looking for "how does a
> played session render?" will not find here. (1) The **player** result page `/sim/<slug>/result/<sessionId>` is a
> **persisted read**, not a live recompute — `internal/graph/queries.resolvers.go:70` does plain Ent SELECTs over
> `validation_attempt_results`, so a seeded result fan-out renders a full result. (2) The **manager** view
> reads the **same** table — **the mirrors are GONE.** `app/terraform/migrations/20260729133514.sql:58-62`
> (*"5. Drop the mirrors."*) **re-points** `organization_assignment_sessions`' two foreign keys off the mirror
> ids (`:15-23`), NULLs the orphans (`:36-44`), then `DROP TABLE`s both `local_jobsimulation_sessions` and
> `local_skill_path_sessions` (`:62-63`), and `intelligence.go:1700` now reads `m.ent.JobSimulationSession.Query()`.
> **No session row is back-filled — the file contains 0 `INSERT`s** (this said *"back-fills then DROPs"* until
> M257x iter-52).
> **There is one row to seed, not a pair** — the older "seed the mirror or the scoreboard is blank"
> guidance is superseded. Full route-by-route treatment lives in
> [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md); the write side is
> [`../ops/demo/session-clone-spec.md`](../ops/demo/session-clone-spec.md).

### Direct dependencies

> These are the edges the engine has, **not** a reading of a compose block: at `0dab54d` there is no
> `jobsimulation` service and therefore no `depends_on` list to quote. They are satisfied in-process inside
> `backend` (or, for the two remaining cross-process hops, by `backend`'s own compose entry).

* **Backend (app)** — user context, organization scoping
* **CMS** — simulation definitions, content, studio entities. **The engine holds no `DIRECTUS_BASE_ADDR`/`DIRECTUS_TOKEN` of its own**; it calls the cms domain **in-process** (same binary, no RPC hop). There is no husk container on either end of that edge any more — compose's `CMS_RPC_ADDR`, which only `messenger` reads, is `http://backend:8083` (measured at platform `0dab54d`). **The M23 content cutover does NOT ride on a `cms` container.** `backend` is the in-process Directus reader (`app/cms_reader_switch.go`; `app/main.go:980-982` @ `app` `b948604` v1.366.0 `log.Fatalf`s without `DIRECTUS_BASE_ADDR`), so re-pointing `cms` alone leaves `backend` reading prod — measured live on `demo-1` at M257x iter-24 as 96 Directus log lines, all 403. rext therefore sets `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` in both twins. No jobsimulation env change is needed, but the cutover must include `backend`.
* **Sentinel** — authz
* **Storage** — file uploads, recordings
* **Skiller RPC surface** — skill metadata; served by **Backend (app)** since the skiller→app merge (July 2026): `SKILLER_RPC_ADDR=http://backend:8083`
* **Roadrunner** — **ORPHANED, no longer called** (v2.7 M247). Code execution moved **in-process into jobsimulation** (`internal/runner/runner.go`, an in-process Judge0 client — its header reads *"formerly the standalone 'roadrunner' service"*); `ROADRUNNER_RPC_ADDR` is dead config. See [`roadrunner.md`](roadrunner.md).
* **PostgreSQL**, **Redis** — base infra

### External integrations

* **LiveKit** — primary voice engine (`LIVEKIT_HOST_URL`, `LIVEKIT_RECORDINGS_BUCKET_NAME`)
* **AWS Chime SDK** — video/camera/screensharing recording (`CHIME_RECORDINGS_BUCKET_NAME=ant-prod-chime-demo`)
* **ElevenLabs** — voice agents still used in the call/reply pipeline (`ELEVENLABS_TEMPLATE_AGENT_ID`, `ELEVENLABS_EU_TEMPLATE_AGENT_ID`). Engine choice is per sequence, from the CMS `voice_engine` field, not from a flag; **when that field is nil the default is `gptrealtime`** (`cms/directus/collections/jobsimulation.go:1594-1597`), not ElevenLabs
* **AssemblyAI** — EU voice transcription for call recordings (`ASSEMBLYAI_API_KEY`)
* **Bunny.net** — video stream hosting / tokenized playback (`BUNNY_REC_STREAM_API_KEY`, `BUNNY_TOKEN_HASH_KEY`)
* **PostHog** — feature flags + telemetry (`POSTHOG_API_KEY`); `flag_use_realtime_openai` selects **no engine** — read *inside* `CreateAgentDispatch` (`calls/livekit.go:131-135`), it sets the endpoint to `openai-hosted` **and resets the agent name to the bare `anthropos-agent`** (`:140-144`), silently overriding a US session's `anthropos-agent-us`. See [`../architecture/ai_architecture.md`](../architecture/ai_architecture.md)
* **AI providers** — via the shared `ai` library

### Redis Streams

* Producer: `jobsimulation` stream (session completed, insights generated)
* Consumer (subscribes to): `cms` (content events). (The former `roadrunner` code-execution stream is orphaned — code execution is now in-process; see [`roadrunner.md`](roadrunner.md).)

Redis Streams consumption is handled by the colony pubsub `SubscriberServer` wired up in `cmd/root.go`, not by `internal/worker/` (which is Asynq-only).

## Startup contract — read this before diagnosing a crash (M217)

> **Scope (v2.8 M257x).** The cobra contract below describes the **frozen `jobsimulation` binary**, which no
> compose file starts any more — you will only hit it running the legacy repo by hand. **The
> `$HOME/.aws/credentials` landmine two sections down did NOT retire with it**: the fold carried the bind
> over to **`backend`**, where it is live at `0dab54d`. Read that part as a `backend` bug.

**The cobra ROOT command's `RunE` *is* the server.** There is **no `serve` and no `run` subcommand.**

- The image is `ENTRYPOINT ["./application"]` with **no CMD**; when compose still declared the service it
  passed **no `command:`** either.
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

**Always read the FIRST line of `docker logs`, never the help block.** That rule is for the frozen binary.

> ⚠️ **The signature did NOT survive the fold — do not go looking for a help block in `backend`.** `app` has
> **no cobra root command**: `app/main.go:216` (@ `origin/main` `7177374`, identical at `9d00a313`
> v1.367.0; `:212` at the older `b948604` v1.366.0) is a plain
> `func main()`, and the only `spf13/cobra` import in the whole repo is `cmd/createTaxonomy/main.go`. There is
> no `RunE`, so there is nothing to print `Error: …` and nothing to print a usage block. A failed init in
> `backend` is a single stdlib `log.Fatalf` line — timestamped, no `Error:` prefix, no help — and the container
> exits 1. The jobsim wiring is fatal by design at `app/main.go:670` (`:614` @ `b948604`).

So on the container that actually runs the engine since the fold — `backend`; there is no `…-jobsimulation-1`
container to inspect — the same underlying failure reads like this instead:

```bash
docker logs demo-<N>-backend-1 2>&1 | head -3
# <date time> jobsim-in-app: engine wiring failed (is jobsim env provisioned?): can't load AWS config: failed to load shared config file, ...
#   ^ stdlib log's LstdFlags prefix. No `Error:`, no usage block, nothing after it.
```

That line is assembled from the source strings, not pasted from a capture — but its **shape** is the point:
one line, naming the wiring stage rather than a command. `head -3` is still right (the first line is still the
whole story), but "it printed the CLI help" is, on `backend`, a report of something that cannot happen.

### The `$HOME/.aws/credentials` landmine (why it died in every demo) — now a `backend` bug

**The bind survived the fold and moved onto `backend`.** `docker-compose.yml:91` binds
`$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** AWS bind in the file — under `backend`'s
`volumes:` (`:90`), and compose's own comment says why (`:88-89`: *"jobsim-in-app's Chime/LiveKit recording
managers use the AWS SDK default credential chain — the mount the standalone jobsimulation container had."*).
Measured at platform `0dab54d`. **When the host path does not exist, Docker auto-creates it as an empty
DIRECTORY.** The container then sees a *directory* where a file belongs, and `aws-sdk-go-v2`'s
`config.LoadDefaultConfig()` **opens it successfully** (opening a directory succeeds!) before failing `EISDIR`
on the read — so it is *not* skipped as an unreadable file. In the standalone binary that error propagated out
of `ai.NewAIManager` → the root `RunE` → cobra's usage block → `exit 1`. **The CAUSE is inherited; the
SIGNATURE is not, and the container name is not the only thing that changed.** In `backend` the identical
`config.LoadDefaultConfig` failure comes out of `jsai.NewAIManager` (`app/internal/jobsimulation/ai/ai.go:90`,
`can't load AWS config: %w`), is returned unwrapped by `jobsimwiring.Wire`
(`app/internal/jobsimwiring/wiring.go:147-148`) and dies at `log.Fatalf` in `app/main.go:670` — one timestamped
line, no `Error:` prefix, no usage block (`app` `9d00a313` v1.367.0; the fatal is `:614` @ `b948604`).

**With the path simply absent, `LoadDefaultConfig` returns `nil`.** The mount is the bug.

- **On a workstation** with a real `~/.aws/credentials` file, it works — which is why this never showed up in
  local dev and only bit a fresh Linux box.
- **In a demo/dev stack**, rext's **generated compose override drops the bind** (`volumes: !reset null` on the
  demo path; an `!override`-tagged empty list on the dev path). Zero platform-repo edits. A stack carries **no
  AWS credentials at all**, so that mount could only ever *be* the broken empty directory.

> ⚠️ **A bare `volumes: []` does NOT remove it** — compose *merges* volume sequences and the inherited bind
> survives. Only the `!reset` / `!override` tags remove it. Verified against the compose binary.

**Downstream when the mount kills it — and the blast radius GREW with the fold.** It no longer costs you one
surface: `backend` is the only *application* container the `core` profile starts (`sentinel` is authz-only),
so the whole platform goes with it. The AI-Simulations surface is gone and so is every other one;
`:8082/graphql/query` is unreachable rather than "the jobsimulation subgraph erroring" — there is no
jobsimulation subgraph, the supergraph is `backend` alone; the `pt-aisim-chat-launch` playthrough cannot pass;
no session-completed events reach the Redis stream, so the in-process skill-path engine never sees
completions. It is also what used to sit behind the nameless *"1 check(s) FAILED"* the bring-up's autoverify
reported — a symptom that would now present as a wholesale failure instead.

## Local Development

### Run in Docker

```bash
cd platform
make up                           # the `core` profile — `backend` (app) runs the jobsim engine
# There is NO jobsimulation profile and no jobsimulation container. Asking for one does NOT
# fail: it exits 0 and starts only postgresql, redis and sentinel.
```

### Run natively

**To work on the live engine, you run `app`, not this repo:**

```bash
cd platform
make dev S=backend                # stops the backend container
cd ../app
go run .                          # the jobsim engine runs inside this process
```

**To run the frozen pre-merge binary** (reading the old source, reproducing old behaviour) — note there is no
`make dev S=jobsimulation` step any more: it would stop nothing and exit 0, and `make init` no longer clones
the repo (no `repos.yml` entry since `d11a403`), so clone it by hand.

```bash
cd jobsimulation
make setup                        # installs ent, atlas, gqlgen, goverter
make gen                          # regenerates Ent + Goverter + gqlgen
go run .                          # binds :8080 / :8081 — the binary's own defaults, nothing supplies 8400/8401
```

Make sure `.env` has the LiveKit + AWS credentials and that Postgres/Redis are reachable on `localhost`.

### Migrations

```bash
cd platform
make migrate   # NOT `S=jobsimulation`. `jobsimulation` has NO `repos.yml` entry at all since
               # platform `d11a403` (2026-08-03) removed it, so the migrating set is `app`
               # ALONE. Forcing S=jobsimulation runs the frozen repo's atlas env and
               # re-materialises the dead `jobsimulation` schema.
```

## Related Documentation

* [AI Architecture](../architecture/ai_architecture.md) — voice engines, recording, model routing
* [CMS](./cms.md) — content source
* [Dependency Map](../architecture/dependency_map.md) — RPC + event-stream relationships
