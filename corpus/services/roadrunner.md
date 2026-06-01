# Roadrunner Service

## Role & Responsibility

Roadrunner is the **code-execution proxy** for the platform. When a simulation includes a coding task, jobsimulation hands the user's source code to Roadrunner, which forwards it to **Judge0** (a sandboxed code-execution API) and returns the results (stdout, stderr, status, time).

Roadrunner exists for one reason: it gives the platform a clean, language-agnostic boundary for running untrusted code without ever executing it in our own services or on our own infrastructure.

It also runs an **Asynq** worker pool for asynchronous batch submissions.

## Architecture & Code Map

* **Codebase**: `roadrunner` (local) — repo `git@github.com:anthropos-work/roadrunner`
* **Language**: Go 1.25
* **Frameworks**: Connect-RPC, [Asynq](https://github.com/hibiken/asynq) (`v0.25.1` background tasks), `gorilla/websocket`
* **Ports**: 10400 (HTTP — `/_meta` health only), 10401 (Connect-RPC) — same on host and inside container per `platform/docker-compose.yml` (`PORT=10400`, `RPC_PORT=10401`)
* **Profile**: `graphql` (default) and `roadrunner`
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
| `SubmissionPackage(...)` | Submit a batch of runs in one call |
| `SubmissionResult(token)` | Poll for execution result (output, errors, status, time) |

### HTTP / WebSocket

* The HTTP server (`PORT` 10400) exposes only the `/_meta` health endpoint. All code submission goes through Connect-RPC on `RPC_PORT` 10401.
* The repo contains an experimental WebSocket LSP proxy (`internal/lsp/lsp.go`) that is NOT wired into any running server — there is no reachable LSP endpoint today.

### Async tasks

Every submission enqueues exactly one poll task on the `roadrunner:default` queue (MaxRetry 3) from `runner.CreateSubmission`; the worker (10 concurrent, `internal/worker/worker.go`) runs `HandleSubmissionResultTask`, which polls Judge0 up to 15 times at 1s intervals, then publishes a `RoadrunnerSubmissionCompleted` event. The RPC handlers call the runner directly and never invoke the Asynq client; there are no HTTP handlers.

On completion the worker publishes a `RoadrunnerSubmissionCompleted` event (carrying the Judge0 token) to Redis Streams (`REDIS_STREAMS_INDEX`) via colony pubsub; jobsimulation consumes it as the async signal that execution finished.

## Dependencies

* **Upstream consumers**: jobsimulation (the only caller — `ROADRUNNER_RPC_ADDR=http://roadrunner:10401`)
* **Downstream**: Judge0 at `JUDGE0_BASE_URL=http://52.48.139.23:2358` (default in compose), Redis (Asynq backend)
* **No database** — roadrunner owns no Postgres schema and stores no persistent state of its own. Judge0 holds submission state by token.

## Local Development

### Run in Docker

```bash
cd platform
make up                  # default graphql profile — includes roadrunner
# or just roadrunner:
make up PROFILE=roadrunner
```

### Run natively

```bash
cd platform
make dev S=roadrunner
cd ../roadrunner
go run main.go
```

Native runs require the platform `.env` to be sourced (or `REDIS_ADDR`, `REDIS_STREAMS_INDEX`, `REDIS_WORKER_INDEX`, `JUDGE0_BASE_URL`, `JUDGE0_API_KEY` exported). `REDIS_WORKER_INDEX` must be a valid integer — if unset/non-numeric the process exits immediately (`strconv.Atoi` error in `cmd/root.go`). `main.go` auto-loads a local `.env` (`godotenv/autoload`) if one is present in the working directory.

### Smoke-test execution

There is no REST submission endpoint — submit via Connect-RPC on port 10401. The language map accepts `py`, not `python`. Note: proto contracts are NOT vendored in the roadrunner repo; they come from the shared `github.com/anthropos-work/proto` module (`proto/roadrunner/v1/roadrunner.proto`). Rely on server reflection rather than a local `--schema` flag.

```bash
# Submit a Python script (returns a token)
buf curl http://localhost:10401/roadrunner.v1.RoadRunnerService/Submission \
  -d '{"runtime":"py","source_code":"print(2+2)","stdin":""}'
# → {"token":"..."}

# Fetch result (poll until status != "in_queue")
buf curl http://localhost:10401/roadrunner.v1.RoadRunnerService/SubmissionResult \
  -d '{"token":"..."}'
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `10400` | HTTP health port (`/_meta` only) |
| `RPC_PORT` | `10401` | Connect-RPC port |
| `JUDGE0_BASE_URL` | `http://52.48.139.23:2358` | Judge0 API endpoint |
| `JUDGE0_API_KEY` | — (required) | Judge0 `X-Auth-Token`; the one Judge0 var NOT set in the compose environment block — supplied via platform/.env |
| `SENTRY_DSN` | — (optional) | Sentry error-tracking DSN |
| `REDIS_ADDR` | `redis:6379` | Redis address for Asynq |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | `0` | Redis DB index for Asynq |
| `ENVIRONMENT` | `development` | Environment name |

## Testing

Roadrunner currently has NO test suite — there are zero `*_test.go` files, so `go test ./...` (also run at Docker build time, `Dockerfile` line 18) is a no-op that passes vacuously.

```bash
cd roadrunner
go test ./...   # currently no tests defined
```

## Related Documentation

* [Jobsimulation Service](./jobsimulation.md) — the only consumer
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)
