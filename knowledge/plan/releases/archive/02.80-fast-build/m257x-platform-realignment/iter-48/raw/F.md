# Seat F — iter-48

Repo: `/Users/marco/workspace/anthropos/rosetta` @ `m257x/platform-realignment` (`cabc3b1`).
Ground truth opened directly: `stack-demo/app` @ **`5ba17044`** (v1.363.2), `stack-demo/platform` @ **`2adcf71`**,
`stack-demo/jobsimulation` @ `462343b0`, `stack-demo/studio-desk` @ `14a5442`, `stack-demo/next-web-app` @ `bb3313bc0`,
`stack-demo/ant-academy` @ `9c3843cd`, `stack-demo/graphql-wundergraph` @ `60c229f`, plus
`stack-demo/{cms,sentinel,storage,messenger,roadrunner}/terraform/main.tf` and the rext authoring copy
`.agentspace/rosetta-extensions` @ `932554e`.

## Coverage (file, `wc -l`, lines read)

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/services/ant-academy.md` | 436 | **all 436** (reader displayed a trailing 437th empty line) |
| 2 | `corpus/services/studio-desk.md` | 435 | **all 435** (reader displayed a trailing 436th empty line) |
| 3 | `corpus/services/jobsimulation.md` | 226 | **all 226** |
| 4 | `corpus/architecture/platform-migration-status.md` | 189 | **all 189** |
| 5 | `corpus/services/skillpath.md` | 107 | **all 107** |
| 6 | `corpus/architecture/frontend_architecture.md` | 105 | **all 105** |

Every file was read top-to-bottom in full with `Read` (no offset/limit), before any grep. Total 1498 lines.

**Rule compliance.** (1) Every search's stderr was read — one zsh glob rejection (`--include=*.go` unquoted)
was caught and re-run quoted; the "no matches" it produced was NOT trusted. (2) Positive controls run in the
same pass: `JobSimulationService` (matched, 3 hits) alongside `SkillPathSessionService` (0 hits);
`check:types` (matched, 5 package.json) alongside `check:deprecations`. (3) Every quoted `file:line` was read
with 5–20 lines of surrounding context. (4) Counts were re-derived, not copied (23 `CREATE TABLE`s, 9
`repos.yml` entries, 8 i18n locales, 7 studio l12n locales, 31 academy e2e specs, 32 `NEXT_PUBLIC_BACKEND_API_URL`
references). (5) I also **executed** the corpus's own fence:
`PLATFORM_REPOS_YML=stack-demo/platform/repos.yml python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py`
→ **exit 0**, "platform-migration-status.md and repos.yml agree in both directions".

## Blockers

**None. 0 blockers.**

Every substantive claim in all six files that is checkable against the cloned platform source verified TRUE.
The highest-risk items were re-derived from source rather than accepted:

- `app/main.go` per-domain wiring call sites — skiller `:573` (`skiller.NewSkillerManager`), jobsimulation
  `:604` (`jobsimwiring.Wire`), skillpath `:634` (`skillpath.NewSessionManager`), cms `:1034` (`appcms.Wire`) —
  all four exact (migration-status:60).
- `app/internal/roadrunner/` does not exist; the Judge0 runner is `app/internal/jobsimulation/runner/`,
  constructed at `internal/jobsimwiring/wiring.go:118` `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))` — exact.
- 23 `CREATE TABLE`s in `20260722081626_jobsim_data_model.sql`; `job_simulation_sessions` created at
  `20260722104506.sql:2` and `DROP TABLE "sessions"` at `:79` — `public.sessions` really does not exist.
- Mirror drops: `20260729133514.sql:62` `local_jobsimulation_sessions`, `:63` `local_skill_path_sessions`
  (skillpath.md's citation is exact); it IS the last migration in the repo.
- `InsightsSkillPathByMemberships` at `internal/organization/intelligence.go:1144`, its `SkillPathSession`
  query at `:1159-1170` (skill_path_id + status ∈ {active,completed} + tenant predicate) — exact to the line;
  `intelligence.go:1700` is a `m.ent.JobSimulationSession.Query()` — exact.
- `SkillPathSessionService`: **0** hits across app/cms/jobsimulation/sentinel/storage/messenger/roadrunner,
  no `skillpath…v1connect` import (positive control passed) — skillpath.md:30-33 confirmed.
- All compose anchors exact: sentinel `:5`/`:18`, backend `:28`/`:52`, jobsimulation `:83`/`:104`, cms `:144`,
  storage `:189`, customerio-sync `:220-222`/`:238`, messenger `:240`/`:258`, roadrunner `:281`,
  studio-desk `:311`/`:334`/`:337-341`, next-web-app `:344`/`:352`/`:362`, gotenberg `:371-372`,
  `common.yml:2` postgresql / `:20` redis, AWS bind at `:142` (the file's ONLY aws bind).
- All terraform desired counts: app `:44`=1, cms `:39`=0, jobsimulation `:40`=0, roadrunner `:19`=1,
  sentinel `:19`=1, storage `:19`=1, messenger `:19`=1, router `:20`=1.
- Supergraph: `supergraph-config-prod.yaml` = `backend` alone, `schemas/` = `backend.graphqls` alone,
  `subgraphs.conf` = `BACKEND=v1.360.0`, `config.prod.yaml` `listen_addr: 0.0.0.0:8080` + `graphql_path: /graphql`
  (so frontend_architecture's ":8080/graphql in prod" is right).
- All 17 cited platform shas resolve with the exact dates claimed (incl. the pre-history set: `cb6ebf5`,
  `8770fe6`, `1474b1f`, `84862d1`, `467965a`, `ef4b449`, `b43b99a`, `c17cc9a`, `a2a3ee6` — and I diffed
  `a2a3ee6` to confirm it really removes the `directus:` service, `b43b99a`/`c17cc9a` really add/remove `realtime:`).
- ant-academy: the public-route table matches `code/proxy.js`'s `isPublic` matcher **entry for entry**
  (including both dev-only conditional groups); `getServerCatalogView()` is verbatim
  `(await getBackendCatalogView(eids)) ?? emptyCatalogView()`; `serverTenant.js:115-145` really contains
  *"the cutover is intentional, not reversible-on-error"*; `serverChapterBody.js:52/67` is the backend-null →
  `{notFound:true}` path; `app/not-found.jsx:43` is "You wandered off the trail."; `?lang=` is query-param-only
  (`src/i18n/locale.js`) and `LocaleSwitch.jsx` is a 2-way EN↔IT `<Link>`; `academyCatalogSeries` /
  `academyCatalogSkillPaths` / `upsertChapterProgress[Batch]` / `setLastActivity` all exist in
  `app/internal/web/backend/graphql/graph/schemas/academy.graphqls`; the table is
  `academy_chapter_progresses` (plural, `20260603182114_add_academy_progress.sql:21`); `cmd/academy-seed` exists;
  the four named rext academy patches all exist in `demo-stack/patches/`.
- studio-desk: `STUDIO_ACCESS_ROLES = ['admin','org:admin','content_creator','org:content_creator']` at
  `src/index.ts:96`; `checkEnterpriseAndAdmin` at `:99`; PORT in-code fallback `9100` at `:60`;
  `YOUTUBE_API_KEY` read at `src/routes/youtube.ts:43` with the `_mock:true` fallback; `skillpath.ts` = 61,080 B;
  `GCLOUD_SERVICE_ACCOUNT` at `.env.example:120` + `terraform/main.tf:129` and read by **no** code in `src/`;
  `main.ts:97` `preloadCriticalCSS()`, `:105` the tailscale/GlitchTip comment, `:206` `new PageWrapper()` — all
  three exact; the three hardcoded `https://app.anthropos.work` sites (pageWrapper.js:149, userProfile.js:148, :302)
  are present in the canonical repo as described; both pinned rext regression test names exist
  (`stack-injection/tests/test_injection.py:1594` and `:1651`); all five studio demopatches exist.
- jobsimulation husk: cobra root has no `serve`; subcommands are exactly aggregate/clone-session/test/validate;
  `cmd/root.go:77-78` falls back to `8080`/`8081`; `internal/runner/runner.go`'s header says *"formerly the
  standalone 'roadrunner' service"*; the `!reset null` demo emitter is at
  `stack-injection/gen_injected_override.py:590` exactly as described.
- next-web-app: `apps/web/package.json:46` `"next": "^16.2.7"`; `pnpm@10.30.3`; `engines.node >=24.0.0`;
  `!apps/mobile` workspace exclusion; ports 3000/3001/3002/3031; 8 i18n locales; codegen schema
  `http://localhost:8082/graphql/query` with `documents: ['src/query/**']` and the client preset into
  `src/__generated__/`; 441 commits in July 2026 (">300/month" holds).

## Minors

1. **`corpus/architecture/platform-migration-status.md:63`** — wrong evidentiary sha. The roadrunner row says
   *"that file has not been touched since `87d8d44` (2026-06-19)"*. `87d8d44` **is the roadrunner repo's HEAD**
   (2026-06-19, *"ci: pass GitHub App secrets to bump workflow"*) and it touches only
   `.github/workflows/bump-version.yml`. `terraform/main.tf` was last touched at **`e45eb61` (2026-05-27)**
   (`git log -1 -- terraform/main.tf`). The claim it supports is still TRUE (and is understated: the file is
   even older than stated, and `:19` still reads `= 1`), but in a document whose thesis is *"every claim is
   cited to a sha"* the sha should be `e45eb61`, or the sentence should say "the repo has been dormant since
   `87d8d44`".
2. **`corpus/services/jobsimulation.md:101`** — anchor range off by one. Cites
   `app/terraform/migrations/20260729133514.sql:58-62` for dropping *both* mirrors; `local_skill_path_sessions`
   is dropped at **`:63`** (`:62` is `local_jobsimulation_sessions`). `skillpath.md:83` cites both lines
   correctly. Also "back-fills then `DROP TABLE`s" is loose: step 2 of the migration **remaps** the link ids
   (`UPDATE … SET session_id = l.skill_path_session_id`), it does not back-fill.
3. **`corpus/services/jobsimulation.md:35`** — mis-targeted anchor. *"The local re-point onto `app` is M809,
   not yet done — see `app/main.go:1196-1202`."* Lines 1196-1202 are the **cms-in-app M807** comment (about
   `CMS_RPC_ADDR` and the cms M809 re-point). The jobsim edge comment is `app/main.go:1190-1195`. Claim true,
   anchor points at the cms twin.
4. **`corpus/architecture/frontend_architecture.md:39`** — stale count. *"~15 sites hitting
   `NEXT_PUBLIC_BACKEND_API_URL`"*. Measured at `next-web-app` `bb3313bc0`: **32 references across 23 files**
   (~21 of them real env-read call sites; the rest comments/tests/error strings). The predicate — the data
   layer is not GraphQL-only — is TRUE, and every named example (invitations, assignment-builder, Stripe,
   CSV bulk import, admin backfill) exists.
5. **`corpus/architecture/frontend_architecture.md:52`** — `check:deprecations` is listed as Turbo-orchestrated
   ("orchestrated by Turbo (`turbo check`, `check:lint`, `check:types`, `check:deprecations`)"). It exists, but
   as a **root `package.json:35` script that invokes eslint directly** (`--no-config-lookup`), and it is **not**
   a task in `turbo.json` (tasks: build/check/check:lint/check:types/codegen/dev/lint/start/test + per-app
   variants). The other three are genuine turbo tasks.
6. **`corpus/architecture/frontend_architecture.md:27-35`** — the "Core Packages (`packages/`)" table omits
   **`packages/design`**, a real workspace package (the other four `packages/*` — ui, graphql, core-js — plus
   two `configs/*` are listed). Incomplete, not false.
7. **`corpus/architecture/frontend_architecture.md:97`** — "Recent UX Work (**May 2026**)" is three months
   stale on a repo doing ~441 commits/month (July 2026 alone). Harmless framing, but it is the only dated
   section in the file.
8. **`corpus/services/ant-academy.md:137`** — incomplete verbatim quote. Says
   `emptyCatalogView() = { chapters: [], skillPaths: {}, series: [] }`; the actual return
   (`code/src/lib/serverTenant.js:116`) is
   `{ chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }`.
   The "→ 0 cards" conclusion still holds (`BundleStripe` returns `null` at `pathCount === 0` and the cards are
   the path triggers threaded from the empty view), so this is a quoting nit, not a wrong mechanism.
9. **`corpus/services/ant-academy.md:50`** — stale repo-layout annotation: `code/tools/  # offline-parity CLI`.
   `code/tools/` now contains only `apply-v3-metadata.mjs`; offline was removed at v0.5 M1, which the same doc
   states at `:34`. Self-inconsistent annotation.
10. **`corpus/services/ant-academy.md:246`** — drifted count: *"~26 Playwright e2e spec files"*; `code/tests/e2e`
    holds **31** `*.spec.js` files at `9c3843cd`. (The companion "1000+ Vitest tests" is true and conservative —
    ~2,701 `it(`/`test(` call sites.)
11. **`corpus/services/ant-academy.md:90-94`** (the mermaid) — draws `CMS` and `Jobsim` as peer boxes inside
    "Core Backend (Tier 1, Docker)". Per the fenced map they are `merged-into-app` in prod and
    `running_but_unfederated` husks locally. The doc's prose is correct everywhere else; only the diagram is
    pre-merge shaped.

### Not defects — recorded so the next pass does not re-open them

- **`ant-academy.md:324` (`code/.env`) vs rosetta `CLAUDE.md` ("the React app reads only from
  `code/.env.local`").** The service doc is **right**: `ant-academy/.env.example:1` itself says *"NOT for the
  React app — that lives in `code/.env`"*, and Next.js loads `.env` and `.env.local` both. If anything is
  drifted it is the root `CLAUDE.md` line — outside this file set.
- **§3 census numbers** in `platform-migration-status.md` (93 org repos / 46 unnamed / the archived-on-GitHub
  dates) could not be re-measured: `gh` is not installed on this box and I did not read a `GH_PAT` value to
  substitute a `curl` (values-blind). They are **unverified**, not refuted; everything in §2 that *is* offline-
  checkable verified exact.
