# Seat D — M257x iter-49 KB-fidelity audit (ninth clause-5 reading)

## Header — shas consulted

| what | sha | note |
|---|---|---|
| `rosetta` (audited repo) | `2fc633a2c5c09a6034e5ab4e29d509dfcadcbd8a` | branch `m257x/platform-realignment` |
| `stack-demo/app` | `5ba1704482cf812b130c2d3673afd09f4f7f22e5` | backend monolith |
| `stack-demo/app/studio` | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` | full non-shallow `anthropos-studio-room` clone |
| `stack-demo/platform` | `2adcf714bd877a205e8948f59a23db49b884c054` | compose / repos.yml / git history (`045857c`, `fdfa189`, `2adcf71`) |
| `stack-demo/next-web-app` | `bb3313bc0133ee5728ce83fda485e95bfea1a6c6` | |
| `stack-demo/roadrunner` | `87d8d44382ef07a9f165869530cbac9e5e0a4332` | |
| `stack-demo/jobsimulation` | (working clone) | for `09631fb2`, `internal/runner/runner.go` |
| `stack-demo/studio-desk` | `14a5442` | for `STUDIO_ACCESS_ROLES` anchors |
| `stack-demo/graphql-wundergraph` | (working clone) | for `915da06` |
| `.agentspace/rosetta-extensions` | `4d03b53a5e524e9abb020c1a4534ec968c25072b` | `clerkenstein/`, `alignment/`, `stack-seeding/` |

**No `chronos` and no `intelligence` clone exists on this box** (`find` over the repo root returns only
`corpus/services/chronos.md`). Their repo-internal claims are therefore checkable only against the
platform's own git history (the removal commits), which I did. `gh` is **not installed** on this box, so
GitHub archive-status claims (chronos NOT archived, skiller/skillpath archived) could **not** be
re-verified — I state that rather than let it read as confirmed.

## Coverage

| # | file | wc -l | lines read |
|---|---|---|---|
| 1 | `corpus/services/studio-room.md` | 473 | all 473 |
| 2 | `corpus/services/clerkenstein.md` | 366 | all 366 |
| 3 | `corpus/services/chronos.md` | 245 | all 245 |
| 4 | `corpus/services/roadrunner.md` | 171 | all 171 |
| 5 | `corpus/services/next-web-app.md` | 126 | all 126 |
| 6 | `corpus/services/gotenberg.md` | 82 | all 82 |
| 7 | `corpus/services/intelligence.md` | 18 | all 18 |
| | **total** | **1481** | **1481** |

Each file was read top-to-bottom in a single `Read` (no `limit`/`offset`), then re-opened at specific
line ranges while verifying anchors.

## BLOCKERS

| # | site (file:line) | the false claim | what is true (with platform file:line) |
|---|---|---|---|
| D-1 | `corpus/services/roadrunner.md:87` | RPC table row: **"`SubmissionPackage(...)` \| Submit a batch of runs in one call"** | `SubmissionPackage` submits **exactly ONE Judge0 submission** and returns **one** token. `stack-demo/roadrunner/internal/rpcsrv/rpc.go:43-57` calls `s.runnerManager.CreateSubmission(ctx, "zip", nil, codePackagePtr, nil)` — the *same* single-submission function `Submission` calls, only with the source moved into `additional_files` and the runtime set to `"zip"`. `internal/runner/runner.go:62` POSTs `{base}/submissions?base64_encoded=true&wait=false` — Judge0's **single**-submission endpoint, **not** `/submissions/batch`. `internal/runner/languages.go:32-33` maps `"zip" → 89`, i.e. Judge0's *multi-file program* language, not a batch mode. So the method is "submit one multi-**file** program", never "a batch of runs". A reader wiring a batch call against this contract gets one token back for one run. |
| D-2 | `corpus/services/roadrunner.md:49` | "It also runs an **Asynq** worker pool for asynchronous **batch submissions**." | The Asynq pool handles **one poll task per single submission**. `internal/runner/runner.go:126-127` enqueues exactly one `roadrunner:submissionresult` task per `CreateSubmission` (`asynq.Queue(queues.DefaultQueue)`, `asynq.MaxRetry(3)`); `internal/worker/worker.go:47-48` registers that **one** task type on the mux (`internal/worker/tasks/tasks.go:4` is the only task constant in the repo). There is no batch path anywhere. This also **contradicts this same file at `:97`**, which correctly states *"Every submission enqueues exactly one poll task on the `roadrunner:default` queue (MaxRetry 3) from `runner.CreateSubmission`"*. |

Both blockers share one root: the word "batch" was attached to `SubmissionPackage`/the worker pool
somewhere upstream and never checked against `rpc.go`. The correct reading is *multi-file package*, not
*multiple runs*.

## MINORS

1. **`chronos.md:27`** — "**Ports**: `8080` (HTTP/health), `8081` (RPC)". The platform ran chronos at
   `PORT=8500` / `RPC_PORT=8501`, and jobsimulation dialled `CHRONOS_RPC_ADDR=http://chronos:8501`
   (both visible in the removed block of `platform` `045857c`, `git show 045857c -- docker-compose.yml`).
   8080/8081 is almost certainly the *repo's code default* — the identical colony service roadrunner
   defaults to exactly those (`roadrunner/cmd/root.go:84` `RPC_PORT`→`"8081"`, `:110` `PORT`→`"8080"`) —
   but the doc never says the platform overrode them, so a reader reconstructing the historical topology
   gets the wrong port. **Not escalated to blocker**: no chronos clone exists here, so the code default
   cannot be confirmed, and the doc is explicitly historical-only.
2. **`chronos.md:202`** — env table gives `REDIS_STREAMS_INDEX` example `2`; the platform set `4`
   (`045857c` removed block; `4` is also what backend/jobsimulation/roadrunner use today).
3. **`roadrunner.md:21`** — "and **no other platform repo references roadrunner at all**" is an absolute
   refuted by grep: `stack-demo/next-web-app/knowledge/service-dependencies.md:77` and
   `stack-demo/studio-desk/knowledge/04-integration/platform-integration.md:137` (plus
   `studio-desk/knowledge/00-product/simulation-experience.md` and
   `.../01-platform-context/anthropos-platform.md`) all still name it. They are stale *docs*, not code —
   the Go-code half of the sentence is correct and I re-verified it (zero hits for
   `ROADRUNNER_RPC_ADDR` / `RoadRunnerService` / `roadrunner:10401` across `app` + `jobsimulation`
   `*.go`; positive control: a case-insensitive `roadrunner` grep over the same trees *does* hit
   `app/main.go`, `app/internal/jobsimulation/runner/runner.go`, etc., so the grep works).
4. **`roadrunner.md:148-149`** — the env table's "Default" column gives `PORT 10400` / `RPC_PORT 10401`.
   Those are the **compose** values (`platform/docker-compose.yml:295,301`); the **code** defaults are
   `8080`/`8081` (`roadrunner/cmd/root.go:110`, `:84`) — which matters precisely for the "Run natively"
   recipe two sections above it.
5. **`roadrunner.md:9-10`** — "`backend` reads `JUDGE0_BASE_URL` and calls Judge0 directly" is true in
   code (`app/internal/jobsimwiring/wiring.go:118` `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"),
   getenv("JUDGE0_BASE_URL"))`), but the **local `backend` compose block sets neither var**
   (`platform/docker-compose.yml:43-67`); the file's only `JUDGE0_BASE_URL` is line 297, inside the
   **roadrunner** block. Unless `platform/.env` supplies it, in-process code execution on a local stack
   resolves an empty base URL. Worth a fence in the merged-service banner.
6. **`clerkenstein.md:3`** — "**Last updated:** 2026-07-14" and a status line that stops at
   *v2.3 "cue to cue" M218*, while the body carries v2.4 M224 (`isHiring` threading), v2.8 M256
   (the `signedOut`/D81 block, single-tenant seat disclosure) and v2.8 M257x iter-23 (the colony-pin
   drift fence). The metadata under-reports the page by two releases.
7. **`clerkenstein.md:17-18`** — "ONE private monorepo with sections (`clerkenstein`, `demo-stack`,
   `stack-injection`, `stack-core`, `stack-seeding`, `alignment`)" enumerates **6**; the monorepo has
   **11** (`ls .agentspace/rosetta-extensions`: alignment, clerkenstein, demo-stack, **dev-stack**,
   **playthroughs**, stack-core, stack-injection, **stack-secrets**, stack-seeding, **stack-snapshot**,
   **stack-verify**). Presented as a complete parenthetical.
8. **`clerkenstein.md:101`** — `cmd/` is listed as `mintpk` · `fake-fapi` / `fake-bapi`; there is a
   fourth binary, `clerkenstein/cmd/jwtkey`.
9. **`clerkenstein.md:307`** — anchor drift: the M220 clerk-js block is `clerk-frontend/server.go:35-68`
   (`var clerkJSClient` is line 68, the doc says `35-67`). Everything the sentence asserts is exact —
   `clerkJSFetchTimeout = 15 * time.Second`, the `"Explicitly NOT http.DefaultClient"` comment,
   `FAKE_FAPI_CLERKJS_CACHE`, and the source-level fence test
   (`clerk-frontend/clerkjs_cache_test.go:196` bans `http.Get(` / `http.DefaultClient`).
10. **`studio-room.md:388`** — "its only outbound API call is to the skills taxonomy service
    (`api.anthropos.work`)". True of *platform* APIs (`services/taxonomy.py:11`
    `BASE_URL = "https://api.anthropos.work/api"`; no GraphQL/Directus calls anywhere), but the pipeline
    also calls the AI providers over the network (`services/ai.py`; `requirements.txt` ships `openai`,
    `anthropic`, `mistralai`). "Only outbound API call" over-reads.
11. **`studio-room.md:337-341`** — "against the managed venv at `studio/studio-venv`
    (`studioManager.go:94-96`)". The constants are exact (`app/internal/cms/studio/studioManager.go:94-96`),
    but the venv is used **only** when `s.env == colony.Development` (`:112-118`); otherwise `pyBin` stays
    plain `python3`.
12. **`studio-room.md:406`** — anchor `gen.py:450-456`; the `work_paths` loop is `gen.py:450-455`.
13. **`studio-room.md:426`** — anchor `exporter.py:513-550`; `_create_export_package` runs
    `agents/simulation/postgen/exporter.py:514-552` (the `rmtree` of the scratch dir is line 552).
    The load-bearing `:518-519` two-zip anchor **is** exact.
14. **`studio-room.md:60-93`** — the project-structure tree shows `workspace/` (with four subdirs) as a
    repo directory; it does not exist in the repo and is created at runtime
    (`gen.py:450-454` `os.makedirs`). The tree also omits `tests/`, `tools/`, `pytest.ini`, `CLAUDE.md`.
15. **`studio-room.md:219`** — "the repo's own `CLAUDE.md:12-14`"; the command is on line **12** alone
    (11 and 13 are the ``` fences). The quoted command text is verbatim-correct.
16. **`next-web-app.md:36-43`** — the "Shared packages (`packages/` + `configs/`)" table omits
    `packages/design`.
17. **`intelligence.md:18`** — "Exposed an HTTP server on :8080 (PORT env override)". The parenthetical
    saves it, but the platform ran it at `PORT=9002` (`fdfa189` removed block) and the doc never says so.

## What I checked hardest, and what came back clean

Because the blockers landed in one file, here is the audited-clean surface, so a later seat does not
re-spend the same effort:

**`gotenberg.md` — clean, fully verified.** Image `gotenberg/gotenberg:8`, port `3200`, profiles
`[graphql, backend, all]`, and the four-flag `command` all match `platform/docker-compose.yml:371-384`
verbatim. `ConvertToPDF(ctx, gotenbergURL, document, filename) ([]byte, error)` matches
`app/internal/converter/gotenberg.go:16` exactly; the **90 s client-side timeout** is
`gotenberg.go:13` (`&http.Client{Timeout: 90 * time.Second}`); the endpoint is
`gotenbergURL+"/forms/libreoffice/convert"` (`:31`). "The backend service (`app`) is the only consumer"
holds — `GOTENBERG_URL` appears in `app` only (`main.go:243`,
`internal/web/backend/coursebuilder/handler.go:242`, `terraform/main.tf:281`) and a `gotenberg` grep
over cms / jobsimulation / sentinel / storage / messenger / roadrunner / studio-desk / next-web-app
returns nothing.

**`next-web-app.md` — clean.** Next `^16.2.7` and React `^19.2.7` in all four of
`apps/{web,hiring,integration,maintenance}/package.json`; `engines.node ">=24.0.0"`,
`packageManager pnpm@10.30.3`, `turbo ^2.9.6`; **no `middleware.ts` anywhere** and `proxy.ts` in
web/hiring/integration; the public allowlist items (`/login`, `/sign-up`, `/checkout`, `/free-trial`,
`/monitoring`, `/print`, `/api/bunny/thumbnail`) are all in `apps/web/src/proxy.ts:7-57`; `/print` is
HMAC-gated via `PRINT_ROUTE_SECRET` (`packages/core-js/src/security/printToken.ts:16,49`); **storybook
is genuinely gone** (no script, no `.storybook/`, only `configs/tailwind/storybooks.css` survives);
**8 locale dirs** on disk vs the repo's own `CLAUDE.md:23` still saying 7 — the doc's "some docs say 7"
is correct; only `Dockerfile.dev` at root while `CLAUDE.md:82-84` still claims two — the doc's "stale"
call is correct; `CLAUDE.md:15` and `:55` are cited at the exact right lines; `graphql-request` +
`@tanstack/react-query` with **no Apollo**; mobile excluded at `pnpm-workspace.yaml:11` and Expo on
3031; `test` script only in web+hiring; `docker-compose.yml:352` bakes the endpoint; `:5050/graphql`
**was** the pre-`2adcf71` value (`git show 2adcf71~1:docker-compose.yml:377`); `915da06` resolves in
`graphql-wundergraph` ("fold cms subgraph into backend").

**`clerkenstein.md` — clean on every load-bearing number and anchor.** DNA gene/capability counts read
straight out of the JSON: `clerk-2.6.0` **27 genes / 14 capabilities**, `clerk-deploy-1` **7**,
`clerk-express-1` **13 / 5**, `clerk-js-5` **9**, `clerk-multi-1` **9** — all as claimed.
`clerk-backend/store.go:138` `SeedOrgIdentity` / `:151` `LookupOrgEid` are **exact**;
`alignment/dna/clerk-2.6.0.json:131` is **exact** and contains the quoted
*"taking the Go surface 97.2% -> 100%"*. `alignment/cmd/alignctl/run.go:134-135` is **exact**
(`ExitRegressed = 2`, `ExitUnmeasurable = 3`). The M256/D81 block is exact on every point: `signedOut`
at `clerk-frontend/server.go:107`, the `_method=DELETE` dispatch at `:583`, `establishLocked` minting
`"sess_clerkenstein"` at `:665`, `clientID: "client_clerkenstein"` at `:125`,
`handleSelectIdentity` setting `signedIn=false; sessID=""` at `:646-647` while clearing `signedOut` at
`:654`, and `handleClient`/`handleToken`/`handleMe`/`handleSignOut` all declared with `_ *http.Request`
(`:241`, `:467`, `:488`, `:554`). **`r.Cookie(`/`.Cookie(` appears nowhere in `clerkenstein/*.go`** —
verified with a positive control (a plain `Cookie` grep hits `server.go:348,394` and
`handshake_test.go:39`). All five named `server_test.go` tests exist at `:256/:286/:390/:427/:461`.
The colony pin fence is correct in **both** directions: `clerkenstein/go.mod:8` = `colony v0.34.3` and
`go.mod:9` = `clerk-sdk-go/v2 v2.6.0`, while `app/go.mod:16` = `colony v0.35.2` and `:31` =
`clerk-sdk-go/v2 v2.7.0`. `studio-desk/src/index.ts:96` and
`studio-desk/app/services/userService.ts:16` are **exact line hits** for `STUDIO_ACCESS_ROLES` with the
quoted comment. The BAPI/FAPI twin split is real (`clerk-backend/server.go:47` vs
`clerk-frontend/server.go:186`), and the "not yet a measured gene" fence is correct — `clerk-js-5.json`
has a `Me` capability (`universal-user`, `unauthenticated`) and **no** gene for
`/v1/me/organization_memberships`. `@clerk/express ^1.3.47` and the `dotless-pk-rejected` gene are in
`clerk-express-1.json:3,24`. `clerkenstein/.github/workflows/alignment.yml` exists while the
**monorepo root has no `.github/`** — so the "git-tracked but inert" claim holds.
`clerkenstein/alignment/scripts/` exists with `gate.sh`/`drift-check.sh`/`drift-test.sh` and
`ALIGN_DIR` defaults to `$base/../../alignment` (`gate.sh:30`), while the rext `alignment/` section has
**no** `scripts/` — exactly as documented. Roster threading verified end-to-end:
`stack-seeding/seeders/roster.go:53-54,77,140,155` (producer) →
`clerk-frontend/registry.go:125-126,144` (consumer) → `resources.go:50-51,265-276`
(`orgNameDefault`/`orgSlugDefault` fallback + the conditional `isHiring` emit).

**`studio-room.md` — clean on every checkable claim.** `gen.py:484-492` is **exactly nine**
`add_argument` calls, in the documented order; `parse_argument` at `gen.py:18-28` really does
`parse_known_args` + fold unknown `--k v` pairs, so the `--template` swallowing story is literally
true; `validate_blueprint_exclusivity` at `:241` with the documented whitelist at `:247-250` and the
verbatim *"Cannot combine --blueprint with content parameters"* message; `translate_legacy_blueprint`
at `:212` with `_LEGACY_TEMPLATE_DEFAULTS` at `:205` naming **exactly** the three legacy shapes, and the
verbatim *"is ignored; asset type is now inferred from task interactions"* warning; **zero** consumers of
a `template` key outside that bridge. Two state files per sim (`save_state`, `gen.py:114-117`) with the
`worklog/` in-code fallback at `:106` and `worklog_path = workspace/trace` in the shipped configs. All
three tracked configs carry `max_tokens = 4000`; the `{MODE}_AI_{BRANCH}_MODEL` rows and the env-var
override in `load_services_settings` (`gen.py:41-57`) match. `GenMode` (`services/ai.py:35-42`) has the
five documented members with `DEFAULT = EXECUTION`. `postgen.py:396-398` really does mark
`--media`/`--simid`/`--target` `required=True`, and `agents/simulation/postgen/` holds
`guidance/metadata/translation/toolkit` **plus** `testing.py`, with `exporter.py` the only writer —
`:518-519` is the exact two-zip line pair and the `simulation.json`-inside-the-zip narrative
(`unpack → write → make_archive → rmtree`) is literally what `_create_export_package` does. The
`model.py:59,467-469` taxonomy memo anchor is **exact**, and it is the only cache in the pipeline.
`requirements.txt` matches all nine listed deps with no `aiohttp`. The blueprint example matches
`tests/e2e/blueprints/technical.json` key-for-key, and both blueprint dirs hold exactly the listed
files. The embedding claims check out: `.github/workflows/build-production.yml:29`
`additional_repo: "anthropos-studio-room:studio"`, `Dockerfile:28` / `Dockerfile.dev:26`
`FROM python:3.11-slim`, `CHANGELOG.md:80` = v1.360.1 *"pull studio via additional_repo like cms"*.
The Asynq numbers are exact: `app/internal/cms/worker/worker.go:29` `Concurrency: 5`, `:31-33`
`ai_video: 7` / `studio: 3`; `studioManager.go:119` invokes `studio/gen.py`. The "parallel API calls"
claim is substantiated (12 `asyncio.gather` sites across `prep.py`/`story.py`/`evaluation.py`/postgen).

**`chronos.md` / `intelligence.md` — the removal claims are confirmed against platform git.**
`045857c` (2026-04-17, *"remove chronos service from orchestration…"*) removes chronos from **both**
`docker-compose.yml` (-37) and `repos.yml` (-4); `fdfa189` (2026-04-17, subject **verbatim** as quoted)
removes intelligence from **both** (-25 / -3). A grep over the whole `platform` tree (excluding `.git`)
returns **zero** files mentioning either — positive control: the same grep for `roadrunner` returns
`docker-compose.yml` + `repos.yml`. Jobsimulation `09631fb2` exists with the **verbatim** subject the
doc quotes, and the Asynq replacement is real
(`jobsimulation/internal/worker/tasks/tasks.go:17` `JobsimulationSessionTimeout`). Intelligence's
`DB_CONNECTION_BACKEND` / `DB_CONNECTION_SKILLER` are exactly as documented in the removed block.

**Also re-confirmed for `roadrunner.md` (the parts that are right):** `terraform/main.tf:19`
`service_desired_count = 1` is **exact**; `repos.yml:29-31` carries the *"legacy — folded into app"*
comment and the file lists **9** repos (the router entry was deleted by `2adcf71`, `-5` lines in
`repos.yml` — verified in the diff); `docker-compose.yml:281` starts the container with
`profiles: [graphql, roadrunner, all]`; jobsimulation is at `repos.yml:17` **and**
`docker-compose.yml:83` with `profiles: [graphql, jobsimulation, all]`, and still carries the dead
`ROADRUNNER_RPC_ADDR=http://roadrunner:10401`; the jobsim runner header comment reads *"(formerly the
standalone \"roadrunner\" service)"* verbatim; **zero** `*_test.go` in the roadrunner repo with
`RUN go test -v ./...` at `Dockerfile:18` (so the vacuous-pass claim is exact); worker
`Concurrency: 10`, queue `roadrunner:default`, `MaxRetry(3)`, `maxRetries = 15` with `time.Sleep(1s)`,
and the `RoadrunnerSubmissionCompleted` publish at `runner.go:276-278`; `godotenv/autoload` at
`main.go:9`; the `strconv.Atoi(os.Getenv("REDIS_WORKER_INDEX"))` hard-exit at `cmd/root.go:68-71`;
`"py"` (not `"python"`) at `languages.go:26`; `internal/lsp/lsp.go` present and unreferenced by
`cmd/root.go`; the file tree matches one-for-one.
