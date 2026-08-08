# Roadrunner Service

## Role & Responsibility

> **⚠️ DELETED AND REPLACED — NOT "merged into `app`" (corrected M257x iter-137; this banner said *"MERGED
> INTO `app` / ORPHANED"*).** Nothing calls this service any more (verified v2.5 M231 KB-6; re-verified v2.7
> "july jitter" M247 against the CONSOLIDATED platform — the ~386-commit `app` bump). **The distinction is
> not pedantry — it is what predicts where the code is.** Seven services were folded into `app` and each
> has a package to show for it (`app/internal/{cms,customeriosync,jobsimulation,messenger,skiller,skillpath,storage}/`).
> **`app/internal/roadrunner/` exists at no ref and was never added** — `git log --all --diff-filter=A --
> internal/roadrunner` returns **0 commits, ever**, in a full 6,728-ref clone at `app` `ad9f3c498` (positive
> control: `jobsimwiring` → 3 paths). Code execution moved **in-process into
> jobsimulation** (`jobsimulation/internal/runner/runner.go`, an in-process Judge0 client whose own header comment
> reads *"formerly the standalone 'roadrunner' service"*) — and with the **jobsim-in-app** merge that runner now
> lives inside **`app`**, still **inside the jobsimulation domain**: `app/internal/jobsimwiring/wiring.go:123`
> wires `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))` against
> `app/internal/jobsimulation/runner`, and the comment above that line says it *"replaces the removed
> roadrunner RPC edge"* — the platform's own word for it is **removed**. `backend` reads `JUDGE0_BASE_URL`
> and calls Judge0 directly, and **nothing calls the roadrunner service any more.**
>
> **⚠️ SETTLED — there is no roadrunner service in production, and this banner spent four days saying the
> question was open (corrected M257x iter-137).** It read *"'There is no roadrunner service in production'
> overstates it"*, on the strength of `roadrunner/terraform/main.tf:19` `service_desired_count = 1`. **That
> line is an input to a module nothing instantiates.** Measured at `infrastructure` `13c248e6` (re-derived
> at source by iter-137; first measured iter-123): `terraform/production/services.tf` declares **exactly
> ten** service modules — `sentinel`, `directus`, `acm_media_certificate`, `storage-service`,
> `next-webapp`, `backend`, `jobsimulation`, `studio_desk`, `db-backup`, `metabase` — and **`module
> "roadrunner"` is not among them.** `roadrunner/terraform/main.tf` is 95 lines at `87d8d443` whose `:10-11`
> is `module "roadrunner" { source = ".../modules/services/base_internal_service" }` fed from **unbound
> `var.*`** (`var.environment`, `var.platform_cluster_id`, `var.platform_vpc_id`, …) — a module awaiting a
> caller that does not exist. **This is the same orphaned-dead-code class as `cms`, `messenger` and
> `graphql-wundergraph`, and `org-repos.md` § 3 was written to close exactly it: a service repo's own
> `service_desired_count` is not evidence of production state.**
>
> **What DOES survive in `infrastructure` is the NAME, in seven places, none of them terraform.** Two CI
> workflows inject the Judge0 credentials from `production_roadrunner_judge0_*` secrets under a
> `# Roadrunner` comment (`infrastructure/.github/workflows/wf-terraform-deploy.yml:209-211`,
> `infrastructure/.github/workflows/wf-terraform-plan-preview.yml:241-243`) into `TF_VAR_judge0_{api_key,base_url}` — and those are consumed
> by **`module "backend_euwest1"`** (`infrastructure/terraform/production/services.tf:384-385`). **Production wiring Judge0
> straight into `backend` under roadrunner-named secrets IS the fold, visible at the config layer**; it is
> better evidence than the count the corpus was reading. The seventh hit is the platform's own words at
> `infrastructure/knowledge/service-dependencies.md:119`: *"**Judge0** (code execution — called directly
> now; `roadrunner` is off this path)"*.
>
> **On the sha, and it still stands:** `roadrunner/terraform/main.tf:19` has
> **not been touched since `84a4b4f` (2025-12-15)** — the commit that first added `terraform/main.tf`,
> seven months before the fold. ⚠️ **This said "`87d8d44` (2026-06-19)" until M257x iter-115, and that is not
> the line's provenance**: `87d8d44` is the repo's HEAD and touches exactly one file,
> `.github/workflows/bump-version.yml` (3 insertions) — it never goes near terraform, so *"not touched since
> it"* is vacuous by construction while the parenthetical presented it as the date of the last touch. The
> subject of *"has not been touched"* is the **line**, not the repo. `git blame -L 19,19 87d8d443 --
> terraform/main.tf` names `84a4b4f`; a file-level `git log` is not line provenance (the file's own most recent
> touch is `e45eb61`, 2026-05-27, a line-11 module-source swap). The corpus's own fenced authority,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md), had this right all along and
> this document never named `84a4b4f` anywhere. The conclusion — *before the fold* — survives; the sha did not.
> Meanwhile the platform has removed it from its own
> clone set: `roadrunner` had a `repos.yml` entry reading *"legacy — folded into app"* as late as `2adcf71`
> (`repos.yml:29-31` **at that ref**), and platform `d11a403` (2026-08-03) **deleted the entry outright** —
> **grade that commit by its diff, not its message.** Its message asserts roadrunner's *"repos.yml clone entry
> was already gone"*; `git show d11a403 -- repos.yml` shows **that very commit** removing `- name: cms`,
> `- name: jobsimulation` **and** `- name: roadrunner`. The message is wrong; the diff is the fact. At
> `0dab54d` there is neither a `repos.yml` entry nor a `roadrunner` compose service to start locally.
> **⚠️ This sentence called it *"the one row where prod and the platform's own declaration contradict each
> other — recorded, not resolved"* until M257x iter-137. There was never a contradiction: both
> declarations say the service is gone, and only a misread of an uninstantiated module input made them
> look opposed.** Say *absent from production and deleted locally*; *orphaned* is right only of the
> **repo**, which is unarchived and still carries dead terraform.
> See [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> It is not part of the *logical* platform stack. On current `origin/main` there is **no `ROADRUNNER_RPC_ADDR`
> / `RoadRunnerService` / `roadrunner:10401` read in any service's Go code** — re-measured M257x iter-98 at
> `app b948604f` / `jobsimulation 462343b0`: **0 Go hits for all three strings at both repos**. **The old
> parenthetical "zero hits outside CHANGELOG" and "no other platform repo references roadrunner at all" is
> RETRACTED** — each is false: outside Go the name is widespread — `jobsimulation/knowledge/operational.md:68`
> tabulates `ROADRUNNER_RPC_ADDR`, `app/knowledge/plan/releases/07.00-jobsim-in-app/RE-PORT-CHECKLIST.md:10`
> names `RoadRunnerService`, and a case-insensitive sweep of every clone at its own ref returns hits in
> **app (25 files), jobsimulation (8), platform (3), studio-desk (3), next-web-app (1)**. **The claim that
> holds is scoped to Go source; the repo-wide form never did.**
>
> **ORPHANED *and* de-orchestrated — M247's "still in `repos.yml`" no longer holds (corrected v2.8 M257x).**
> That line said roadrunner was *"still in `repos.yml` (1 of the 9 repos)"* — a claim about `2adcf71`, and
> **false at `0dab54d`** — unlike intelligence / skiller /
> skillpath, which were *removed* from `repos.yml` + `docker-compose.yml` **and whose GitHub repos really are
> archived** (skiller 2026-07-01, skillpath 2026-07-31). It was true when written — `repos.yml` still
> carried a roadrunner entry **at `2adcf71`**, one of **9** there. Platform `d11a403` (2026-08-03) then
> deleted it. **At `0dab54d` `repos.yml` had 6 entries — app, sentinel, storage, messenger, next-web-app,
> studio-desk — and roadrunner was not among them. It has 4 since `838d907` took `storage` and
> `messenger` out too: app, sentinel, next-web-app, studio-desk.** What still separates roadrunner from that list is only
> its **GitHub repo, which is NOT archived** — and *not* a production deployment. **This clause named
> *"prod's `service_desired_count = 1`"* as the second separator until M257x iter-137; that value is a
> module input nobody instantiates, so it separates roadrunner from nothing** (see the SETTLED banner
> above).
>
> **`jobsimulation` was previously listed above as also *removed* from both files; it is not (corrected
> M257x iter-46, and **superseded by the platform at iter-77**).** Its GitHub repo's archive state is **not visible to this corpus** (this line asserted an archive on 2026-07-31; `origin/main` took four commits on 2026-08-04 — see [`jobsimulation.md`](./jobsimulation.md)), and
> as late as `2adcf71` it did remain in both files — `jobsimulation`'s entry at `repos.yml:17-19` and
> its service at `docker-compose.yml:83`, with
> `profiles: [graphql, jobsimulation, all]`, **at that ref** — so its container then still started on a bare
> `make up`. Platform `d11a403` (2026-08-03) removed **both**: at `0c91421` there is no `jobsimulation`
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
> service as built, not as used.** **Retirement is DONE, on both sides** — corrected M257x iter-137, where
> this line read *"Treat retirement as pending, not done"*; it was the conclusion drawn from the
> misattributed module input above, and it did not survive reading `infrastructure`.

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

On completion the worker publishes a `RoadrunnerSubmissionCompleted` event (carrying the Judge0 token) to Redis Streams (`REDIS_STREAMS_INDEX`) via colony pubsub — **and nothing consumes it.** The jobsimulation consumer was **deleted, not moved**: at `jobsimulation 462343b0` the repo's **Go source** contains exactly one `roadrunner` mention (`internal/runner/runner.go:3`, a comment), with no handler and no event reference — **scope that to Go, not to the repo**: `git -C stack-demo/jobsimulation grep -in roadrunner 462343b0` returns **14 lines across 8 files** (5 exact-case across 3), the rest being CHANGELOG and `knowledge/*.md`; in `app`, `internal/jobsimulation/simulator/stream_handlers.go:30-34` states that the roadrunner-submission pubsub event was removed upstream and the code-submission result now arrives as `HandleCodeSubmissionResultTask` on `CodeRunQueue` — *"NOT stream handlers."* Consistent with the **Upstream consumers** bullet under § *Dependencies* below (*"none (orphaned — see the banner at the top)"*) — **named, not pinned.** ⚠️ **This clause used to carry a bare line pin as its own worked example of a bad one, and the example rotted twice**: it named a line "below" that M257x iter-120 found was *above*, and M257x iter-137's repair then shifted the file until that same pin landed on a blank line and turned two fences RED. **A retraction that quotes the retracted pin re-publishes it** — the anchor guard cannot tell the two apart, and neither can a reader who copies it. The pin is therefore gone entirely; the construct name is the citation.

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
> once printed as their "defaults" were **supplied by the compose service**, from its
> `environment:` block **at `2adcf71`** — and that service was deleted at `d11a403`, so **nothing
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
