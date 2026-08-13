# AUDITOR F — 6 files / 1498 lines

**Positive control:** all 6 files read to their final line; counts match `wc -l`
(ant-academy 436 · studio-desk 435 · jobsimulation 226 · platform-migration-status 189 ·
skillpath 107 · frontend_architecture 105).

## BLOCKERS — 0

Every load-bearing anchor checked resolved AND named the construct the sentence claims.

**iter-46's repair passes independent verification.** `platform-migration-status.md:60` now cites four
separate wiring call sites, and all four are exact at `app` @ `5ba17044`:

| cited | `app/main.go` actually says |
|---|---|
| skiller `:573` | `skillerManager := skiller.NewSkillerManager(logger, jobRoleManager, skillTaxonomyManager, localizationManager)` |
| jobsimulation `:604` | `jobsimDj, err := jobsimwiring.Wire(serverContext, logger, serviceName, ent, …)` |
| skillpath `:634` | `skillPathSessionManager := skillpath.NewSessionManager(logger, ent, cmsReaderSw, jobsimDj.SimManager, …)` |
| cms `:1034` | `cmsManagers = appcms.Wire(appcms.Deps{` |

**The over-correction risk (class 2) did not materialize.** The row's companion claim
*"`app/internal/roadrunner/` does not exist"* is true (`ls` → No such file or directory); the runner is at
`app/internal/jobsimulation/runner/`, and `jobsimwiring/wiring.go:118` is exactly
`runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`.

Also re-verified clean in that file: all **eight** `service_desired_count` values (app 1, cms 0, jobsim 0,
roadrunner 1, sentinel 1, storage 1, messenger 1, router 1); every compose anchor (5, 28, 83, 144, 189,
220-222, 240, 281, 311, 344, 352, 371-372) and every profile claim; `common.yml:2`/`:20`; the router row's
supergraph evidence (`supergraph-config-prod.yaml` = `backend` alone, `schemas/` = `backend.graphqls`
alone, `subgraphs.conf` = `BACKEND=v1.360.0`); and both **measured** counts — `git log -p --follow --
repos.yml` returns exactly **14** names, `docker-compose.yml` exactly **26**.

## MINORS — 9

| # | site | what is off |
|---|---|---|
| 1 | jobsimulation.md:35 | cites `app/main.go:1196-1202` for the M809 local re-point; those lines are the **cms**-in-app M807 comment block. The jobsim handler registration is `:1195`, cited correctly at `:95`. Surrounding claim true (compose `:52` + `:258` still point at `http://jobsimulation:8401`) |
| 2 | jobsimulation.md:101-102 | cites `20260729133514.sql:58-62` for dropping **both** mirrors; `DROP TABLE "local_skill_path_sessions"` is at **:63** (skillpath.md:83 cites it correctly) |
| 3 | ant-academy.md:137 | `emptyCatalogView()` literal omits `bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION` (`serverTenant.js:115-117`). The load-bearing conclusion (0 cards) holds |
| 4 | ant-academy.md:147 | `catalog.json` "~2,667 entries" — the `courses` array now holds **2715** |
| 5 | ant-academy.md:246 | "~26 Playwright e2e spec files" — `tests/e2e/*.spec.js` = **31** |
| 6 | ant-academy.md:390 | catch-all skills row omits `export-path` |
| 7 | frontend_architecture.md:39 | "~15 sites hitting `NEXT_PUBLIC_BACKEND_API_URL`" — 26 files reference it (files, not call sites — soft) |
| 8 | studio-desk.md:112 | tier defaults compressed to 3 values for 4 tiers; Anthropic ids drop date suffixes |
| 9 | platform-migration-status.md:62,70,71,73 · jobsimulation.md:13 | GitHub **archive dates** are unverifiable here — no `gh`, no token, private repos. **Not graded either way**; flagged as a claim class with no local evidence path |

## Files read clean

- **`skillpath.md`** — 0 blockers, 0 minors. Headline negative claim measured: **0** `SkillPathSessionService`
  occurrences and **0** `skillpath…v1connect` imports in Go across the clone set; the only two repo-wide hits
  are `app/CLAUDE.md:72` and `app/knowledge/architecture.md:28` — exactly the Trap C the doc names.
- **`platform-migration-status.md`** — 0 blockers, 0 minors. The repaired passage is correct.
- **`frontend_architecture.md`** — one soft minor. Ports, locales, `next ^16.2.7`, `pnpm@10.30.3`, turbo,
  TS, `docker-compose.yml:362` all exact.
- **`studio-desk.md`** — one minor. M253 de-dup claim verified at source (`pageWrapper.js#init`).
- **`ant-academy.md`** — four cosmetic count minors. The most heavily-asserted passage — the
  *"do not paraphrase this from memory"* `isPublic` matcher table — is **complete and exact**, all 21
  patterns plus both conditional blocks, matching `proxy.js:112-187` group for group.
- **`jobsimulation.md`** — two citation minors.
