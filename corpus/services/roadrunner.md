# Roadrunner Service

## Role & Responsibility

Roadrunner is the **code-execution proxy** for the platform. When a simulation includes a coding task, jobsimulation hands the user's source code to Roadrunner, which forwards it to **Judge0** (a sandboxed code-execution API) and returns the results (stdout, stderr, status, time).

Roadrunner exists for one reason: it gives the platform a clean, language-agnostic boundary for running untrusted code without ever executing it in our own services or on our own infrastructure.

It also hosts a small **LSP** (language-server) helper used by the in-simulation code editor for hints and diagnostics, plus an **Asynq** worker pool for asynchronous batch submissions.

## Architecture & Code Map

* **Codebase**: `roadrunner` (local) — repo `git@github.com:anthropos-work/roadrunner`
* **Language**: Go 1.25
* **Frameworks**: Connect-RPC, [Asynq](https://github.com/hibiken/asynq) (`v0.25.1` background tasks), `gorilla/websocket`
* **Ports**: `10400` (host) → `9000` (container, HTTP/WebSocket); `10401` (host) → `10401` (container, Connect-RPC)
* **Profile**: `graphql` (default) and `roadrunner`
* **Execution backend**: [Judge0](https://judge0.com/) — external sandboxed API at `JUDGE0_BASE_URL`

### Key directories

```
main.go                       Entry point
cmd/
  root.go                     Server startup (HTTP + RPC + worker)
internal/
  lsp/lsp.go                  Language-server helper (over WebSocket)
  rpcsrv/                     Connect-RPC handlers
  runner/
    runner.go                 Judge0 client + execution loop
    languages.go              Supported language IDs (matches Judge0)
  worker/
    worker.go                 Asynq server bootstrap
    client/                   Asynq client (called by handlers)
    queues/                   Queue/priority definitions
    tasks/                    Task type definitions + handlers
```

## Interface Discovery

### Connect-RPC (`RoadRunnerService`)

| Method | Purpose |
|--------|---------|
| `Submission(runtime, source_code, stdin)` | Submit a single execution; returns a `token` |
| `SubmissionPackage(...)` | Submit a batch of runs in one call |
| `SubmissionResult(token)` | Poll for execution result (output, errors, status, time) |

### HTTP / WebSocket

* `POST /run` — REST entrypoint for code submission (alternate to RPC for browser-side callers).
* WebSocket — LSP wire protocol for the in-simulation editor. Implemented in `internal/lsp/lsp.go`.

### Async tasks

Long-running or batch submissions are dispatched onto Asynq queues (`internal/worker/queues`) and processed by the worker pool (`internal/worker/worker.go`). The Asynq client is in `internal/worker/client`, called from the RPC and HTTP handlers when the work shouldn't run synchronously on the request path.

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

### Smoke-test execution

```bash
# Submit a Python script
curl -X POST http://localhost:10400/run \
  -H 'content-type: application/json' \
  -d '{"runtime":"python","source_code":"print(2+2)","stdin":""}'
# → {"token":"..."}

# Fetch result (poll until status != "in_queue")
buf curl --schema ./proto/roadrunner.proto \
  http://localhost:10401/roadrunner.v1.RoadRunnerService/SubmissionResult \
  -d '{"token":"..."}'
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9000` | HTTP + WebSocket port (mapped to 10400 on host) |
| `RPC_PORT` | `10401` | Connect-RPC port (mapped to 10401 on host) |
| `JUDGE0_BASE_URL` | `http://52.48.139.23:2358` | Judge0 API endpoint |
| `REDIS_ADDR` | `redis:6379` | Redis address for Asynq |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | `0` | Redis DB index for Asynq |
| `ENVIRONMENT` | `development` | Environment name |

## Testing

```bash
cd roadrunner
go test ./...
```

## Related Documentation

* [Jobsimulation Service](./jobsimulation.md) — the only consumer
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)
