# Roadrunner Service

## Role & Responsibility

> **⚠️ MERGED INTO `app` / ORPHANED — nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7 "july jitter"
> M247 against the CONSOLIDATED platform — the ~386-commit `app` bump).** Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**. `backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly, and **nothing calls the
> roadrunner service any more.**
>
> **⚠️ Precision, because the declarations disagree (v2.8 M257x).** *"There is no roadrunner service in
> production"* overstates it: `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1` and has
> not been touched since `87d8d44` (2026-06-19, before the fold), while the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` **and** `- name: roadrunner`. The message is wrong; the diff is the fact. At
> `0dab54d` there is neither a `repos.yml` entry nor a `roadrunner` compose service to start locally. This is the **one row where prod
> and the platform's own declaration contradict each other** — recorded, not resolved. Say *orphaned* (nothing
> calls it), not *absent*. See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> It is not part of the *logical* platform stack. On current `origin/main` there is **no `ROADRUNNER_RPC_ADDR`
> / `RoadRunnerService` / `roadrunner:10401` read in any service's Go code** (M247 re-grepped `app` + `jobsimulation`
> on the consolidated clones — zero hits outside CHANGELOG), and no other platform repo references roadrunner at all.
>
> **ORPHANED *and* de-orchestrated — M247's "still in `repos.yml`" no longer holds (corrected v2.8 M257x).**
> That line said roadrunner was *"still in `repos.yml` (1 of the 9 repos)"* — a claim about `2adcf71`, and
> **false at `0dab54d`** — unlike intelligence / skiller /
> skillpath, which were *removed* from `repos.yml` + `docker-compose.yml` **and whose GitHub repos really are
> archived** (skiller 2026-07-01, skillpath 2026-07-31). It was true when written — the entry sat at
> `repos.yml:29-31` **at `2adcf71`**, one of **9** entries at that ref. Platform `d11a403` (2026-08-03) then
> deleted it. **At `0dab54d` `repos.yml` has 6 entries — app, sentinel, storage, messenger, next-web-app,
> studio-desk — and roadrunner is not among them.** What still separates roadrunner from that list is only
> its **GitHub repo, which is NOT archived**, and prod's `service_desired_count = 1`.
>
> **`jobsimulation` was previously listed above as also *removed* from both files; it is not (corrected
> M257x iter-46, and **superseded by the platform at iter-77**).** Its GitHub repo IS archived (2026-07-31), and
> as late as `2adcf71` it did remain in both files — `jobsimulation`'s entry at `repos.yml:17-19` and
> its service at `docker-compose.yml:83`, with
> `profiles: [graphql, jobsimulation, all]`, **at that ref** — so its container then still started on a bare
> `make up`. Platform `d11a403` (2026-08-03) removed **both**: at `0dab54d` there is no `jobsimulation`
> `repos.yml` entry and no `jobsimulation` compose service, so nothing of it starts. **The two ended up in
> exactly the same place** — orphaned, de-orchestrated, uncloned — which is why the old
> orphaned-but-still-orchestrated distinction this blockquote drew between them has been withdrawn. Related
> statements live at `../architecture/architecture_overview.md` and `README.md`.
>
> **Correction (v2.8 M257x, measured):** `chronos` does **not** belong in that list — its GitHub repo is **NOT
> archived** (last push 2026-04-23), only removed from orchestration at platform `045857c`. The corpus called it
> archived; the org disagrees. See [`../architecture/platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> roadrunner is **no longer a `docker-compose.yml` service at all**, and `ROADRUNNER_RPC_ADDR` is set by **no**
> compose file and is absent from `.env_example` — measured at platform `0dab54d`. (Both were true as recently as
> `2adcf71`; the container and its env went together.) There is no roadrunner profile either, and asking for one
> does not fail — it exits 0 and starts the 3-service floor. **Everything below describes the
> service as built, not as used.** Treat retirement as pending, not done.

Roadrunner is the **code-execution proxy** for the platform. When a simulation includes a coding task, jobsimulation hands the user's source code to Roadrunner, which forwards it to **Judge0** (a sandboxed code-execution API) and returns the results (stdout, stderr, status, time).

Roadrunner exists for one reason: it gives the platform a clean, language-agnostic boundary for running untrusted code without ever executing it in our own services or on our own infrastructure.

It also runs an **Asynq** worker pool that polls Judge0 for results — **one task per single submission**, as § Async tasks below states.

## Architecture & Code Map

* **Codebase**: `roadrunner` — repo `git@github.com:anthropos-work/roadrunner` (**not** archived). **Not cloned by `make init`**: no `repos.yml` entry since `d11a403`. Clone it by hand; the surviving code path is `app`'s in-process Judge0 runner
* **Language**: Go 1.25
* **Frameworks**: Connect-RPC, [Asynq](https://github.com/hibiken/asynq) (`v0.25.1` background tasks), `gorilla/websocket`
* **Ports**: **8080 (HTTP — `/_meta` health only), 8081 (Connect-RPC) — the binary's own defaults, and now the only ones there are**: `cmd/root.go:110` `cmp.Or(os.Getenv("PORT"), "8080")`, `:84` `cmp.Or(os.Getenv("RPC_PORT"), "8081")`. The **10400 / 10401** pair quoted throughout this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=10400` (`:298`) / `RPC_PORT=10401` (`:302`) and published `10400:10400` / `10401:10401` (`:291-292`) — **at `2adcf71`**. At `0dab54d` there is no `roadrunner` service, so nothing sets them and nothing is published. **Treat 10400/10401 as historical, not as an address**
* **Profile**: **none — there is no `roadrunner` compose service.** Deleted by platform `d11a403`; the line that stood here named the `graphql` profile, which `0dab54d` renamed `core`, for a service that had already been removed. Historical only (corrected M257x iter-68)
* **Execution backend**: [Judge0](https://judge0.com/) — external sandboxed API at `JUDGE0_BASE_URL`

### Key directories

```
main.go                       Entry point
cmd/
  root.go                     Server startup (HTTP + RPC + worker)
  runcode/                    Debug CLI subcommand (runcode.go + launch.go) — lists Judge0 languages
internal/
  lsp/lsp.go                  Experimental WebSocket LSP proxy — NOT wired into any running server
  rpcsrv/                     Connect-RPC handlers
  runner/
    runner.go                 Judge0 client + execution loop
    languages.go              Supported language IDs (matches Judge0)
  worker/
    worker.go                 Asynq server bootstrap
    client/                   Asynq client (called by handlers)
    queues/                   Queue/priority definitions
    tasks/                    Task-type constant only ('roadrunner:submissionresult'); handler lives in internal/runner/runner.go
```

## Interface Discovery

### Connect-RPC (`RoadRunnerService`)

| Method | Purpose |
|--------|---------|
| `Submission(runtime, source_code, stdin)` | Submit a single execution; returns a `token` |
| `SubmissionPackage(...)` | Submit one **multi-file** program in one call; returns a single `token` |
| `SubmissionResult(token)` | Poll for execution result (output, errors, status, time) |

### HTTP / WebSocket

* The HTTP server (`PORT`, default `8080`) exposes only the `/_meta` health endpoint. All code submission goes through Connect-RPC on `RPC_PORT` (default `8081`). The 10400/10401 pair this line used to quote was compose-supplied — see § Ports.
* The repo contains an experimental WebSocket LSP proxy (`internal/lsp/lsp.go`) that is NOT wired into any running server — there is no reachable LSP endpoint today.

### Async tasks

Every submission enqueues exactly one poll task on the `roadrunner:default` queue (MaxRetry 3) from `runner.CreateSubmission`; the worker (10 concurrent, `internal/worker/worker.go`) runs `HandleSubmissionResultTask`, which polls Judge0 up to 15 times at 1s intervals, then publishes a `RoadrunnerSubmissionCompleted` event. The RPC handlers call the runner directly and never invoke the Asynq client; there are no HTTP handlers.

On completion the worker publishes a `RoadrunnerSubmissionCompleted` event (carrying the Judge0 token) to Redis Streams (`REDIS_STREAMS_INDEX`) via colony pubsub; jobsimulation consumes it as the async signal that execution finished.

## Dependencies

* **Upstream consumers**: **none (orphaned — see the banner at the top).** Historically jobsimulation was the only caller via `ROADRUNNER_RPC_ADDR`; at platform `0dab54d` that variable is set by **no** compose file, is absent from `.env_example`, and is read by no Go code — code execution having moved in-process to `jobsimulation/internal/runner/`
* **Downstream**: Judge0 at `JUDGE0_BASE_URL=http://52.48.139.23:2358` (default in compose), Redis (Asynq backend)
* **No database** — roadrunner owns no Postgres schema and stores no persistent state of its own. Judge0 holds submission state by token.

## Local Development

### Run in Docker

```bash
cd platform
make up                  # the `core` profile — `backend` calls Judge0 directly
# There is NO roadrunner profile and no roadrunner container. Asking for one does NOT fail:
# it exits 0 and starts only postgresql, redis and sentinel.
```

### Run natively

Native is the **only** way to run it. There is no `make dev S=roadrunner` step any more — it would stop
nothing and exit 0 — and `make init` no longer clones the repo (no `repos.yml` entry since `d11a403`), so
clone it by hand.

```bash
cd roadrunner
go run main.go           # binds :8080 / :8081 unless you export PORT / RPC_PORT
```

Native runs require the platform `.env` to be sourced (or `REDIS_ADDR`, `REDIS_STREAMS_INDEX`, `REDIS_WORKER_INDEX`, `JUDGE0_BASE_URL`, `JUDGE0_API_KEY` exported). `REDIS_WORKER_INDEX` must be a valid integer — if unset/non-numeric the process exits immediately (`strconv.Atoi` error in `cmd/root.go`). `main.go` auto-loads a local `.env` (`godotenv/autoload`) if one is present in the working directory.

### Smoke-test execution

> **⚠️ Nothing is listening.** There is no `roadrunner` container on any stack, so **these calls have no
> endpoint to reach** — they are kept as a record of the wire contract, not as a recipe. They work only
> against a binary you started yourself by hand (§ Run natively), at whatever `RPC_PORT` you gave it —
> `8081` if you gave it none. **`localhost:10401` resolves to nothing at `0dab54d`.**

There is no REST submission endpoint — submission is Connect-RPC only. The language map accepts `py`, not `python`. Note: proto contracts are NOT vendored in the roadrunner repo; they come from the shared `github.com/anthropos-work/proto` module (`proto/roadrunner/v1/roadrunner.proto`). Rely on server reflection rather than a local `--schema` flag.

```bash
# Against a hand-started binary on its default RPC_PORT. Submit a Python script (returns a token)
buf curl http://localhost:8081/roadrunner.v1.RoadRunnerService/Submission \
  -d '{"runtime":"py","source_code":"print(2+2)","stdin":""}'
# → {"token":"..."}

# Fetch result (poll until status != "in_queue")
buf curl http://localhost:8081/roadrunner.v1.RoadRunnerService/SubmissionResult \
  -d '{"token":"..."}'
```

## Environment Variables

> **Two different columns used to be conflated here.** Only `PORT` and `RPC_PORT` have a **built-in** default
> (`cmp.Or` in `cmd/root.go`); every other row is a bare `os.Getenv` with **no fallback at all**. The values
> once printed as their "defaults" were **supplied by the compose service**, from
> `docker-compose.yml:296-302` **at `2adcf71`** — and that service was deleted at `d11a403`, so **nothing
> supplies them now.** Export them yourself or the binary starts degraded (or, for `REDIS_WORKER_INDEX`, not
> at all).

| Variable | Built-in default | Value compose supplied (`:296-302` @ `2adcf71`) | Description |
|----------|------------------|-----------------------------------|-------------|
| `PORT` | `8080` | `10400` | HTTP health port (`/_meta` only) |
| `RPC_PORT` | `8081` | `10401` | Connect-RPC port |
| `JUDGE0_BASE_URL` | — none | `http://52.48.139.23:2358` | Judge0 API endpoint |
| `JUDGE0_API_KEY` | — none (required) | — not set here | Judge0 `X-Auth-Token`; the one Judge0 var the compose block never set — supplied via platform/.env |
| `SENTRY_DSN` | — none (optional) | — not set here | Sentry error-tracking DSN |
| `REDIS_ADDR` | — none | `redis:6379` | Redis address for Asynq |
| `REDIS_STREAMS_INDEX` | — none | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | — none | `0` | Redis DB index for Asynq — **must parse as an int or the process exits immediately** |
| `ENVIRONMENT` | — none | `development` | Environment name |

## Testing

Roadrunner currently has NO test suite — there are zero `*_test.go` files, so `go test ./...` (also run at Docker build time, `Dockerfile` line 18) is a no-op that passes vacuously.

```bash
cd roadrunner
go test ./...   # currently no tests defined
```

## Related Documentation

* [Jobsimulation Service](./jobsimulation.md) — the historical (now severed) consumer; code execution moved in-process into jobsimulation (M247), so nothing calls roadrunner today
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)
