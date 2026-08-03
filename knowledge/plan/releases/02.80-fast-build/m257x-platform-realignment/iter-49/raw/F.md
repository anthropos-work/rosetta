# Seat F — M257x iter-49, clause-5 reading #9

## Shas consulted

| repo | sha |
|---|---|
| `rosetta` (this repo, branch `m257x/platform-realignment`) | `2fc633a` |
| `stack-demo/app` | `5ba17044` (tag `v1.363.2`) |
| `stack-demo/platform` | `2adcf71` |
| `stack-demo/next-web-app` | `bb3313bc0` |
| `stack-demo/studio-desk` | `14a5442` |
| `stack-demo/ant-academy` | `9c3843cd` |
| `stack-demo/jobsimulation` (frozen/legacy clone) | `462343b0` |
| `stack-demo/cms`, `sentinel`, `storage`, `messenger`, `roadrunner`, `graphql-wundergraph` | present, read for terraform/supergraph anchors |
| `.agentspace/rosetta-extensions` | `4d03b53` |

Fence run: `PLATFORM_REPOS_YML=stack-demo/platform/repos.yml python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py`
→ `platform_alignment_guard: OK — platform-migration-status.md and repos.yml agree in both directions.` (exit 0)

---

## Coverage

| # | file | `wc -l` | lines read |
|---|---|---|---|
| 1 | `corpus/services/ant-academy.md` | 436 | all 436 |
| 2 | `corpus/services/studio-desk.md` | 435 | all 435 |
| 3 | `corpus/services/jobsimulation.md` | 226 | all 226 |
| 4 | `corpus/architecture/platform-migration-status.md` | 189 | all 189 |
| 5 | `corpus/services/skillpath.md` | 107 | all 107 |
| 6 | `corpus/architecture/frontend_architecture.md` | 105 | all 105 |
| | **total** | **1498** | **1498** |

Each file was read top-to-bottom in a single `Read` with no `offset`/`limit`, so the whole body was in
context before any claim was checked.

---

## BLOCKERS

**None found.** See §"Audited zero" below for what was checked hardest and how.

| # | site | the false claim | what is true |
|---|---|---|---|
| — | — | — | — |

---

## MINORS

1. **`ant-academy.md:137`** — `emptyCatalogView() = { chapters: [], skillPaths: {}, series: [] }` is
   written as a literal equality. Actual (`stack-demo/ant-academy/code/src/lib/serverTenant.js:115-117`):
   `{ chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }`
   — two keys short. The conclusion it supports (**0 cards**) is unaffected; `getServerCatalogView()` at
   `:143-145` is verbatim as quoted.
2. **`ant-academy.md:246`** — "~26 Playwright e2e spec files (tests/e2e/)". Actual
   `code/tests/e2e/*.spec.js` = **31** (34 dir entries). The paired "1000+ Vitest tests" is true and
   heavily understated (2,830 `it(`/`test(` sites).
3. **`ant-academy.md:146`** — "`code/public/catalog.json` (~2,667 entries)". Actual `courses[]` length =
   **2,715** (`json.load(...)['courses']`). Drifts with content authoring; the surrounding claim (an
   FS-derived index the grid never reads) is correct.
4. **`ant-academy.md:324` + `:328`** — "The **app's** env file is `code/.env`" / `cp .env.example .env`.
   The repo ships `code/.env.example` and `code/.env.local`; there is no `code/.env`. Next.js does load
   `.env`, so the recipe works — but the same document says `code/.env.local` at `:197` and `:319`, and
   `.env.local` **wins over `.env`**, so on any box the demo tooling has touched (`demo-stack/ant-academy.sh`
   writes `code/.env.local` truncatingly) the §2 recipe is silently shadowed. Worth one sentence.
5. **`ant-academy.md:90-94`** — the mermaid `subgraph Core["Core Backend (Tier 1, Docker)"]` still draws
   `CMS` and `Jobsim` as peer services of `App`. Per the same corpus
   (`platform-migration-status.md:61-62`) both are `running_but_unfederated` **husks**, merged into `app`.
   Stale diagram against the document's own prose.
6. **`jobsimulation.md:34`** — anchor drift. "The local re-point onto `app` is **M809**, not yet done — see
   `app/main.go:1196-1202`." Lines 1196-1202 are the **cms**-in-app M807 comment (messenger's
   `CMS_RPC_ADDR` until the M809 re-point). The **jobsim** edge comment is `app/main.go:1190-1195`, with the
   handler registration at `:1195` — which `jobsimulation.md:95` already cites correctly. The M809 claim
   itself is true (`docker-compose.yml:52` + `:258` still point at `http://jobsimulation:8401`).
7. **`jobsimulation.md:69`** — "**Profile**: `graphql` (default) and `jobsimulation`". Compose declares
   `profiles: [graphql, jobsimulation, all]` (`platform/docker-compose.yml:141`); `all` is omitted.
8. **`jobsimulation.md:112`** — stray markdown: "…reaches the husk cms over RPC.**  **The M23 content
   cutover…" — unbalanced `**` renders a literal asterisk pair.
9. **`frontend_architecture.md:39`** — "~15 sites hitting `NEXT_PUBLIC_BACKEND_API_URL`". Actual: **35
   occurrences across 24 source files** (excluding the two `.env.example`s and one `.md`) in
   `apps/` + `packages/`. Every example the sentence names is real (invite page, `useAssignmentBuilder.ts`,
   `useStripe.tsx`, `EnterpriseAddMultipleEmployeeModal`, `internal/tools/*`), and the point it makes
   ("*GraphQL only* is the wrong mental model") is correct — the number is ~60% low.
10. **`frontend_architecture.md:29-35`** — the "Core Packages (`packages/`)" table omits
    **`packages/design`**, which exists in the workspace alongside `ui` / `graphql` / `core-js`.
11. **`platform-migration-status.md:52`** — the self-audit recipe says "a name they return that has no row
    is a gap." Re-running it verbatim returns 26 names, one of which is **`app-network`** — a compose
    *network*, not a service — which correctly has no row. A reader following the instruction literally
    reports a false gap. (The 14-name `repos.yml` half reproduces **exactly**, and all five named
    pre-history compose services — `nats`, `web-app`, `chromedp`, `simulator`, `realtime` — are present.)
12. **`studio-desk.md:112`** — "Anthropic `claude-opus-4-5` / `claude-sonnet-4-5` / `claude-haiku-4-5`".
    The in-code constants carry date suffixes: `claude-opus-4-5-20251101` / `claude-sonnet-4-5-20250929` /
    `claude-haiku-4-5-20251001` (`studio-desk/src/services/ai/config.ts:35-38`). Also, four tiers map onto
    three listed model names because `thinking_slow` and `thinking_fast` are **both** `gpt-5.2`
    (`config.ts:21-22`) — true but compressed.

---

## Audited zero — what was checked hardest

### `platform-migration-status.md` — every row verified individually

Guard exit 0, **plus** each `file:line` opened by hand:

- **terraform desired counts, all eight, exact line hits:** `app/terraform/main.tf:44 = 1` ·
  `cms:39 = 0` · `jobsimulation:40 = 0` · `roadrunner:19 = 1` · `sentinel:19 = 1` · `storage:19 = 1` ·
  `messenger:19 = 1` · `graphql-wundergraph:20 = 1`. The `roadrunner` "contradiction, recorded not
  resolved" row is genuinely contradictory in source and correctly labelled.
- **`repos.yml`:** exactly **9** entries; `app:10-13` (`migrations: true`, `schema: public`),
  `cms:14-16`, `jobsimulation:17-19`, `sentinel:20-22`, `storage:23-25`, `messenger:26-28`,
  `roadrunner:29-31`, `next-web-app:34-36`, `studio-desk:37-39` — every cited range is byte-correct.
  The `repos.yml:14-31` "legacy — folded into app" comments are real.
- **compose:** `sentinel:5`, `backend:28`, `jobsimulation:83`, `cms:144`, `storage:189`,
  `customerio-sync:220-222` (`context: git@github.com:anthropos-work/customerio-sync.git#main`) +
  `profiles:238`, `messenger:240`, `roadrunner:281`, `studio-desk:311`, `next-web-app:344` + `:352`,
  `gotenberg:371-372` (`gotenberg/gotenberg:8`) — **all exact**. `include: - common.yml` at `:1-2`;
  `postgresql` at `common.yml:2`, `redis` at `common.yml:20`. Sentinel's `search_path=sentinel` is on
  `docker-compose.yml:18` under `migrations: false` — the Trap-A row is real.
- **the four in-process wiring call sites, the row corrected at iter-46:** `app/main.go:573` =
  `skiller.NewSkillerManager(...)`, `:604` = `jobsimwiring.Wire(...)`, `:634` =
  `skillpath.NewSessionManager(...)`, `:1034` = `appcms.Wire(...)`. **All four land exactly.**
- **`app/internal/roadrunner/` does not exist** — confirmed against a full `ls internal/`; the runner is
  `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:118` =
  `runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))` —
  **exact line**.
- **supergraph = one subgraph:** `graphql-wundergraph/supergraph-config-prod.yaml` lists `backend` alone,
  `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0`. All three ✓.
- **`app` @ `5ba17044` is `v1.363.2`** — `git describe --tags` + CHANGELOG head agree.
- **directus removed at `a2a3ee6`** — commit exists, dated 2026-02-27 ✓.
- **§2 completeness recipe re-run:** `git log -p --follow -- repos.yml` → the **same 14 names**, all with
  rows.

### `jobsimulation.md` — the merge AND the remnant state

- **The table rename claim, the load-bearing one:** `20260722081626_jobsim_data_model.sql` has exactly
  **23** `CREATE TABLE` statements ✓; `20260722104506.sql:2` creates `job_simulation_sessions` ✓ and
  `:79` is `DROP TABLE "sessions";` — **both anchors exact**. `public.sessions` really is gone.
- **The mirrors:** `20260729133514.sql:58` = `-- 5. Drop the mirrors.`, `:62` =
  `DROP TABLE "local_jobsimulation_sessions";`, `:63` = `DROP TABLE "local_skill_path_sessions";` —
  exact. `intelligence.go:1700` = `query := m.ent.JobSimulationSession.Query().` — **exact line**. The
  supersession of the old "seed the mirror" guidance is correct.
- **Husk remnant state:** container still defined at `docker-compose.yml:83` in the default `graphql`
  profile ✓; `CMS_RPC_ADDR=http://cms:8091` at `:104` ✓; `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`
  at `:52` (backend) and `:258` (messenger) ✓ — **not** re-pointed, exactly as claimed.
  `$HOME/.aws/credentials` is the only AWS bind in the file ✓.
- **Startup contract:** `cmd/root.go:59` root `cobra.Command` with `Use: "jobimulation"`; the only
  subcommands are `aggregate` / `clone-session` / `test-command` / `validate` — **no `serve`, no `run`** ✓.
  `cmd/root.go:77-78` fall back to `8080`/`8081`; both Dockerfiles `EXPOSE 8080`; compose sets `PORT=8400`
  / `RPC_PORT=8401` and publishes `8400:8400` / `8401:8401` — the "both are correct in their own context"
  framing checks out.
- **Worker pools:** `internal/worker/worker.go:27` `Concurrency: 10`, `:104` `Concurrency: 25` ✓.
- **Redis streams:** `cmd/root.go:284` `AddSubscriber(versionConfig.Name, jobSimSubscriber)`, `:285`
  `AddSubscriber(cmp.Or(os.Getenv("CMS_STREAM"), "cms"), cmsSubscriber)`; `SubscriberServer` at `:121` in
  `cmd/root.go`, not `internal/worker/` ✓.
- **Roadrunner orphaned:** `jobsimulation/internal/runner/runner.go:1-3` header reads *"in-process client
  for the Judge0 sandboxed code execution API … (formerly the standalone \"roadrunner\" service)"* — exact
  quote ✓. Zero `ROADRUNNER_RPC_ADDR` reads in Go across both `jobsimulation/` and `app/` (positive
  control: the same grep pattern hits the runner.go header, and `JUDGE0` hits `wiring.go`). `internal/runner/`
  is repo-relative to the frozen `jobsimulation` clone; `platform-migration-status.md:60`'s
  `app/internal/jobsimulation/runner/` is the app-relative path — **both correct, no contradiction**.
- **`internal/graph/queries.resolvers.go:70`** = `func (r *queryResolver) JobSimulationResult(...)` in
  `stack-demo/jobsimulation` @ `462343b0` — **exact line**.

### `skillpath.md` — the strongest claim in my set, and it holds

- **"there is NO `SkillPathSessionService` anywhere … **0** occurrences in Go source across the clone set"**
  — I re-measured it because rosetta's own root `CLAUDE.md` asserts the opposite ("the
  `SkillPathSessionService` RPC … served by `app`'s `backend` subgraph"). Result:
  `grep -rn "SkillPathSessionService" --include="*.go" stack-demo/` → **0 hits**. Positive control on the
  same invocation: `JobSimulationService` → 5 files (`app/main.go`, `app/internal/jobsimulation/rpcsrv/`,
  `app/internal/jobsimwiring/`, `app/internal/skillpath/session.go`, `app/cms_reader_switch.go`). The grep
  works; the zero is real. **`skillpath.md` is right and the root `CLAUDE.md` is the stale one** — that
  defect is outside my file set; flagging it for whoever owns `CLAUDE.md`.
- The Trap-C evidence is exact: `app/CLAUDE.md:72` and `app/knowledge/architecture.md:28` **both** still
  list `SkillPathSessionService` in the RPC-mux sentence, verbatim.
- `app/internal/skillpath/session.go:205-207` = the `// cms-in-app deseam` comment + the
  `u.cms.GetSkillPathDomain(ctx, skillPathId, version)` call — **exact**. `app/internal/skillpaths/skillpaths.go:88-95`
  brackets the same deseam comment through its `GetSkillPathDomain` call — **exact**.
- `InsightsSkillPathByMemberships` is at `app/internal/organization/intelligence.go:1144`, and its
  `m.ent.SkillPathSession.Query()` with the `SkillPathID` + `StatusIn(active, completed)` + tenant
  predicate occupies **`:1159-1170`** — the cited range is exact to the line.
- `app/internal/web/backend/graphql/graph/schemas/skillpath_sessions.graphqls` exists ✓.
- **The "Coming soon" / unimplemented drill-down claim, verified in `next-web-app` @ `bb3313bc0`:**
  `apps/web/src/components/containers/InsightsBySkillPathStudentSimulationsContainer.tsx:31-34` is
  `const userData = useMemo(() => { /* return insightData?.rows[0]?.membership; */ return null as unknown as MembershipEnriched }, [])`
  — hardcoded null ✓; `:138` renders `{t('enterprise.insights.comingSoon')}` ✓; the `<Table>` and the
  totals `<Flex>` are both commented out at `:140` and `:150` ✓. `apps/hiring` has **no** skill-paths
  route (positive control: `apps/hiring/.../@tabs` dirs do exist) ✓.

### `ant-academy.md` — the `isPublic` matcher table, line by line

The doc warns "do not paraphrase this from memory", so I diffed the whole table against
`code/proxy.js:112-188`. **Every group and every pattern matches**: `/api/_meta(.*)` + `/api/meta(.*)`,
`/robots.txt`, `/sitemap.xml`, `/sitemap(.*)`, `/llms.txt`, `/llms-full.txt`, `/.well-known/(.*)`,
`/courses`, `/courses/(.*)`, `/sign-in(.*)`, `/no-organization`, `/verify/(.*)`, `/api/verify/(.*)`,
`/api/ai/chat`, `/library`, `/library/(.*)`, `/free`, `/free/(.*)`, `/local-content/(.*)`,
`/catalog.json`, `/academy-manifest.json`, `/`, `/latest(.*)`, `/chapters/(.*)`, the `DEV_LOGIN_ENABLED`
spread (`/api/dev/login-as`, `/dev/accept`), and the `VISUAL_BYPASS` spread (`/my-certificates`,
`/my-activity`, `/bookmarks`). `VISUAL_BYPASS` is `BENCHMARK_VISUAL_BYPASS === "1" && NODE_ENV === "development"`
— whitelisted, as claimed (`:62-64`). `/api/chapters/*` is **not** in the matcher, as claimed. The
"no `/library/[slug]` route" claim holds: `code/app/(public)/library/` holds only `page.jsx`.

Also verified: `serverTenant.js` carries the exact in-code sentence *"the cutover is intentional, not
reversible-on-error"* ✓ · `serverChapterBody.js:52` `await getBackendChapterBody(slug, locale)` →
`:65` `maybeResolveDraftBody` → `:67` `return { notFound: true }` ✓ · `app/not-found.jsx:43` = "You
wandered off the trail." ✓ · `app/layout.jsx:132` = `manifest: "/academy-manifest.json"` — **exact line** ✓ ·
`src/i18n/LocaleSwitch.jsx` is a 2-way EN↔IT `<Link>` toggle (`target = locale === 'it' ? 'en' : 'it'`),
not a menu, and there is no `[locale]` route dir ✓ · `coerceLocale` falls back to `en` ✓ ·
`RegisterServiceWorker.jsx` is a documented kill-switch that unregisters SWs and deletes
`['academyChapters','academy-courses']` ✓ · FA Pro **is** vendored (29 `webfonts/*.woff2` +
`css/all.min.css`) with **zero** `@fortawesome` npm dependency, so the token-less-install claim is true ✓ ·
`engines.node >= 22` ✓ · `vercel.json` is exactly `{"framework":"nextjs"}` ✓ · `"dev": "next dev --port 3077"` ✓ ·
Cosmo: `ucourses/.../assistant/agent.js:12-13` = `https://api.openai.com/v1/responses` + `gpt-5.2`,
`localStorage.getItem('openai_api_key')` at `:32`, gated on `NEXT_PUBLIC_FEATURE_TRAINING_COACH`
(`src/lib/featureFlags.js:8`) ✓ · backend write side: `academy_chapter_progresses` /
`academy_last_activities` are the real **plural** Ent table names (`app/internal/data/ent/migrate/schema.go:806, :921`),
`upsertChapterProgress` / `upsertChapterProgressBatch` / `setLastActivity` are in
`app/internal/web/backend/graphql/graph/schemas/academy.graphqls:694-700`, and `app/cmd/academy-seed`
exists ✓ · rext side: all four bring-up patches exist (`demo-stack/patches/academy-fs-published-fallback`,
`-public`, `-chapter-body`, `ant-academy-dev-origins`), `stack-injection/apply-academy-fs-published-body.sh`
exists, `demo-stack/ant-academy.sh:58` states in-code that the launcher sets no
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, the M245 "already running" in-place reconcile is at `:330-350`, and
`stack-verify/live/autoverify.sh:459` is assert **(f)** for the academy catalog ✓ ·
`ensure-clones.sh:20` + `:150` are the `(d2)` explicit non-fatal ant-academy clone ✓.

### `studio-desk.md`

`docker-compose.yml:334` `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query` ✓ ·
`:337-341` `depends_on: backend + cms` — **exact range** ✓ · ports `9000`/`9100` ✓ · `profiles: [studio-desk, all]` ✓ ·
`engines.node ">=24"` + `node:24-alpine` in both Dockerfiles ✓ · `vite.config.ts:28-42` prod inputs are
exactly the nine named entries with `dev-accept` spread in only when `!isProduction` ✓ ·
`src/routes/skillpath.ts` = **61,080 bytes, the largest route** ✓ · `src/routes/youtube.ts:43` =
`const apiKey = process.env.YOUTUBE_API_KEY;` with the `_mock: true` fallback, and it is the **only**
YouTube credential read in `src/` ✓ · `GCLOUD_SERVICE_ACCOUNT` at `.env.example:120` and
`terraform/main.tf:129`, read by **no** code in `src/` ✓ · `STUDIO_ACCESS_ROLES = ['admin','org:admin','content_creator','org:content_creator']`
at `src/index.ts:96`, applied as `adminMiddleware` to `/api/ai` (`:158`), `/api/skillpath` (`:161`),
`/api/youtube` (`:164`) and every builder/catalog/skills page (`:179-228`) ✓ · `isMockClerk` bypass ✓ ·
in-code `PORT` fallback is `9100` (`src/index.ts:60`) ✓ · `AI_PROVIDER_CHAIN=azure-openai,openai` in
`.env.example:57`, `AI_DEFAULT_TIER=fast` at `:61`, in-code fallback `'thinking_fast'` at
`src/services/ai/config.ts:182` ✓ · **exactly three** hardcoded `https://app.anthropos.work` sites —
`pageWrapper.js:149` (logo), `userProfile.js:148` (Back), `userProfile.js:302` (logout) ✓ ·
`app/core/main.ts:105` is the **only** `tailscale` mention in the repo (GlitchTip ingest) ✓ ·
M253 anchors: `preloadCriticalCSS()` at **L97**, `new PageWrapper()` at **L206**, with `l12nService.init()`
at `:191` and `userService.canAccess()` at `:199` in between — **all exact** ✓ · rext side:
`demo-stack/patches/studio-desk-{back-to-cockpit,logo-url,logout-url,no-thirdparty,shell-first-paint}` all
exist, `build_frontend_studio_desk()` is at `demo-stack/up-injected.sh:868`, the env_file work is in
`stack-injection/gen_injected_override.py:342-406` (`frontend_lines`), and both named regression tests
exist at `stack-injection/tests/test_injection.py:1594` and `:1651` ✓.

### `frontend_architecture.md`

`repos.yml` holds exactly **9** entries and ant-academy is not one ✓ · `docker-compose.yml:311` studio-desk ✓ ·
ports 3000 / 3001 / 3002 confirmed from each `package.json` `dev` script, mobile `3031` from
`apps/mobile/package.json:70` (`expo start --port 3031`) ✓ · `!apps/mobile` really is in
`pnpm-workspace.yaml` ✓ · `apps/web/package.json:46` = `"next": "^16.2.7"` — **exact line** — and hiring /
integration / maintenance all match ✓ · `packageManager: "pnpm@10.30.3"` ✓ · `engines.node ">=24.0.0"` ✓ ·
i18n = the 8 named locales `de en es fr it ja nl pt` ✓ · **no** `@connectrpc` / `@bufbuild` / `grpc-web`
dependency anywhere (positive control: `graphql-request ^7.4.0` in `apps/web/package.json:39`) ✓ ·
"Cosmo Router at `:8080/graphql` in prod" ✓ — `graphql-wundergraph/config.prod.yaml:5`
`listen_addr: 0.0.0.0:8080` + `:13` `graphql_path: /graphql`, `terraform/locals.tf:8 port = 8080` ·
the `:5050` retraction is right: `git log -S"5050" -- docker-compose.yml` shows it only ever as a compose
host mapping, removed by `b56d731` + `360efd4`.

---

## Explicitly UNVERIFIABLE from this environment — not counted as findings, not counted as passes

- **Every GitHub archive-state claim** in `platform-migration-status.md` (jobsimulation + skillpath
  archived 2026-07-31, skiller 2026-07-01, graphql-wundergraph 2026-07-30, **chronos NOT archived**,
  intelligence 2026-04-02, "repo not archived" for cms/roadrunner) and the whole §3 census (**93**
  repos / **9** in `repos.yml` / **46** uncited, plus the `auth` and `AI-Labs` notes). `gh api
  repos/anthropos-work/<repo>` **errored for all eight repos I tried** — no GitHub credential in this
  sandbox. That is a failed command, **not** evidence of absence. A seat with `gh` auth should re-run it.
- **`app` PR #1103** (the v9.0 storage+messenger fold), cited on `platform-migration-status.md:65` and
  `:170` — GitHub-gated, same reason.
- **`jobsimulation.md:33`'s production address** `http://backend.internal.anthropos:8081`:
  `messenger/terraform/variables.tf:92` declares `jobsimulation_rpc_address` with **no default**; the
  value lives in the separate `infrastructure` repo, which is not in this clone set. Neither confirmable
  nor refutable here. (The *local* half of the same sentence **is** confirmed.)
- **`infrastructure/terraform/production/services.tf`** (`jobsimulation.md:45` — the
  `module.jobsimulation_euwest1` rollback module and its S3 bucket ownership): that repo is not cloned.
- **Runtime/measured numbers** that are inherently environment-scoped and not statically checkable:
  ant-academy's "grid of **65** real cards with 0 Draft chips", the M253 FCP figures (4669→817 ms),
  the `canAccess` 4049→38 ms table, and the "proven live on `billion`" assertions. The FS content tree
  holds 78 second-level path dirs, which neither confirms nor refutes 65 rendered cards.
