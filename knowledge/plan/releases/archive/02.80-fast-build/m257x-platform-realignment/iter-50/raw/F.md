# Seat F — M257x clause-5 KB-fidelity reading

## 1. Header

**Corpus:** `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfded8791fcb12a4651d747247ce9d04d7f0` (verified via `git rev-parse HEAD`).

**Ground-truth clones consulted** (all re-verified with `git -C <dir> rev-parse --short HEAD`; every one
matched the briefing table):

| clone | sha |
|---|---|
| `stack-demo/app` | `5ba17044` |
| `stack-demo/app/studio` | `aeec036a` |
| `stack-demo/platform` | `2adcf714` |
| `stack-demo/next-web-app` | `bb3313bc` |
| `stack-demo/sentinel` | `88bc5592` |
| `stack-demo/storage` | `4ce8ece5` |
| `stack-demo/messenger` | `fa47850d` |
| `stack-demo/cms` | `ca50c817` |
| `stack-demo/jobsimulation` | `462343b0` |
| `stack-demo/roadrunner` | `87d8d443` |
| `stack-demo/graphql-wundergraph` | `60c229f3` |
| `stack-demo/studio-desk` | `14a5442a` |
| `stack-demo/ant-academy` | `9c3843cd` |

Also consulted: `.agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py`.

### Positive control — `wc -l` on every assigned file

Produced by a single invocation:
`cd /Users/marco/workspace/anthropos/rosetta && wc -l corpus/services/ant-academy.md corpus/services/studio-desk.md corpus/services/jobsimulation.md corpus/architecture/platform-migration-status.md corpus/services/skillpath.md corpus/architecture/frontend_architecture.md`

| file | lines | briefing said | read in full? |
|---|---|---|---|
| `corpus/services/ant-academy.md` | 436 | 436 | yes, 1→436 |
| `corpus/services/studio-desk.md` | 435 | 435 | yes, 1→435 |
| `corpus/services/jobsimulation.md` | 226 | 226 | yes, 1→226 |
| `corpus/architecture/platform-migration-status.md` | 189 | 189 | yes, 1→189 |
| `corpus/services/skillpath.md` | 107 | 107 | yes, 1→107 |
| `corpus/architecture/frontend_architecture.md` | 105 | 105 | yes, 1→105 |
| **total** | **1498** | 1498 | — |

No file read short or empty. No pipeline failure at the read layer.

### Pipeline-integrity notes (method rule 3 / 4)

Three searches in this session returned a **false zero** and were caught and redone. Recording them
because each is exactly the failure mode the briefing names:

1. `grep -rn ... --include=*.ts` under zsh → `(eval):1: no matches found: --include=*.ts`, printed `0`.
   An unquoted glob, not an absence. Redone quoted.
2. `grep -rn "SkillPathSessionService" --include=*.go` → same zsh glob rejection, printed `0` — which
   would have "confirmed" `skillpath.md`'s boldest claim for the wrong reason. Redone with a
   `JobSimulationService` positive control (returned 7) alongside; the real answer for
   `SkillPathSessionService` in `.go` is genuinely 0.
3. `grep -n 'SilenceUsage\|SilenceErrors' cmd/root.go | head` with `echo "(grep exit $?)"` — `$?` was
   `head`'s status, so the probe **could not fail**. Redone with `grep -rn ... cmd/` unpiped
   (rc=1, no matches) plus a `grep -c 'RunE' cmd/root.go` = 1 positive control.

A fourth was a **false absence from a wrong symbol name** (method rule 4/5) and is the single most
important thing I did this pass: `grep -n 'Coming soon'` on
`InsightsBySkillPathStudentSimulationsContainer.tsx` returned nothing, which appeared to refute
`skillpath.md:92-99` outright. Reading the file in full showed the component renders
`t('enterprise.insights.comingSoon')` — a translation key resolving via
`configs/i18n/messages/en/enterprise.json:1862` to `"Coming soon"`. The corpus claim is **correct**;
my grep was wrong. Had I stopped at the grep I would have filed a false blocker.

---

## 2. BLOCKERS

**None.** Zero claims found that a reader would act on and that are false, and zero load-bearing
anchors that fail to resolve to a degree that would harm a reader.

I actively hunted for them in the highest-risk shape the briefing names — the merge banners,
retractions and bolded "**not** X but Y" corrections in `jobsimulation.md`, `skillpath.md` and
`platform-migration-status.md`. Every one of those over-correction candidates checked out
(see §4). The one anchor that genuinely points at the wrong adjacent code block
(`jobsimulation.md:34`) carries a substantive claim that is independently true and correctly
anchored 60 lines later, so it is graded MINOR, not BLOCKER — see MINOR 5 and the reasoning there.

---

## 3. MINORS — 10

1. **`corpus/services/ant-academy.md:137`** — renders `emptyCatalogView()` as
   `{ chapters: [], skillPaths: {}, series: [] }`. Actual (`ant-academy/code/src/lib/serverTenant.js:115-117`)
   is `{ chapters: [], skillPaths: {}, series: [], bundles: PUBLIC_BUNDLES, catalogVersion: CATALOG_VERSION }`
   (`PUBLIC_BUNDLES` / `CATALOG_VERSION` imported from `../../ucourses/catalog.js` at `serverTenant.js:32`).
   Omitted list members. The operative "→ **0 cards**" conclusion is unaffected — cards come from
   chapters/series.
2. **`corpus/services/ant-academy.md:246`** — "~26 Playwright e2e spec files (tests/e2e/)". Actual:
   **31**, from `ls code/tests/e2e/*.spec.js | wc -l`. Undercount.
3. **`corpus/services/ant-academy.md:146`** — "`code/public/catalog.json` (~2,667 entries)". Actual
   `courses` array length is **2,715**, from
   `node -e "const c=require('./public/catalog.json'); ..."` (which also reports `languages: 7`).
   Stale approximation; the "~" softens it.
4. **`corpus/services/ant-academy.md:63`** — describes the progress mutations as
   "`upsertChapterProgress[Batch]` / `setLastActivity`, posted from `code/app/api/academy/beacon/route.js`".
   Both mutations and the `[Batch]` form exist (`src/graphql/schema.graphql:2207`,
   `src/graphql/query/academyProgress.js:19`), and the beacon route does re-issue them. But that route's
   own header (`code/app/api/academy/beacon/route.js:1-18`) says it is the **on-unload last-ditch flush**,
   and that "the real academy mutations go to the CROSS-ORIGIN WunderGraph supergraph" directly from the
   client harness. Naming only the beacon route points a reader debugging academy writes at the fallback
   path rather than the primary one. Incomplete, not false.
5. **`corpus/services/jobsimulation.md:34`** — "The local re-point onto `app` is **M809**, not yet done —
   see `app/main.go:1196-1202`." Lines 1196-1202 are the **cms-in-app M807** comment block, about
   messenger keeping `CMS_RPC_ADDR` on the standalone cms until the M809 re-point. The **jobsimulation**
   handler's comment is `app/main.go:1190-1194` and names no milestone; the handler itself is `:1195`.
   Graded minor, not blocker, because the substantive claim (the local `JOBSIMULATION_RPC_ADDR` still
   resolves to the husk, re-point not done) is **true** and is separately and correctly anchored at
   `jobsimulation.md:95` → `docker-compose.yml:52` (backend) and `:258` (messenger), both
   `http://jobsimulation:8401`, plus `app/main.go:1195`. A reader is not misled about the state of the
   world, only about which comment to read.
6. **`corpus/architecture/frontend_architecture.md:39`** — "~15 sites hitting `NEXT_PUBLIC_BACKEND_API_URL`".
   Actual, from `grep -rn 'NEXT_PUBLIC_BACKEND_API_URL' apps packages` in `next-web-app`: **35 occurrences
   across 26 files** (23 files excluding the two `.env.example`s and one `.md`). Every example the doc
   names — `invite/[token]/page.tsx`, `useAssignmentBuilder.ts`, `useStripe.tsx`, the bulk-import and admin
   backfill tools — is real. Undercount only; the "GraphQL only is the wrong mental model" point stands
   and is if anything understated.
7. **`corpus/architecture/frontend_architecture.md:29-35`** — the "Core Packages" table omits
   `packages/design`, which exists (`ls next-web-app/packages` → `core-js design graphql ui`). Partial
   list; the table is titled "Core Packages", so arguably intentional.
8. **`corpus/architecture/platform-migration-status.md:51-54`** — the section states its own audit rule as
   "a name they return that has no row is a gap", then cites the compose-history union as 26 names. I
   reproduced that union exactly (`git log -p --follow -- docker-compose.yml | grep -E '^\+  [a-z0-9-]+:$' | sort -u | wc -l`
   → **26**, including `nats`, `web-app`, `chromedp`, `simulator`, `realtime` as claimed). But that union
   also contains `graphql`, `wundergraph` and `app-network`, none of which is a row label — the router is
   covered by the `graphql-wundergraph` row and `app-network` is a `networks:` key, not a service. Strictly
   read, the doc's own rule flags three non-gaps. Self-consistency nit; no reader is misled.
9. **`corpus/services/studio-desk.md:112`** — "Tier defaults: OpenAI/Azure `gpt-5.2` / `gpt-5-mini` /
   `gpt-5-nano`; Anthropic `claude-opus-4-5` / `claude-sonnet-4-5` / `claude-haiku-4-5`" gives three model
   names for four tiers and drops the Anthropic date suffixes. Actual
   (`studio-desk/src/services/ai/config.ts:20-38`): OpenAI/Azure `thinking_slow` **and** `thinking_fast`
   are both `gpt-5.2`, `fast`=`gpt-5-mini`, `instant`=`gpt-5-nano`; Anthropic
   `claude-opus-4-5-20251101` / `claude-sonnet-4-5-20250929` (× 2, `thinking_fast` and `fast`) /
   `claude-haiku-4-5-20251001`. Compression, not error.
10. **`corpus/services/skillpath.md:98`** — "the body renders the literal string **"Coming soon"**".
    It renders `t('enterprise.insights.comingSoon')`
    (`InsightsBySkillPathStudentSimulationsContainer.tsx:138`), resolved through
    `configs/i18n/messages/en/enterprise.json:1862` → `"Coming soon"`. True in effect, imprecise in
    letter — and load-bearing for the next reader, because a literal-string grep returns **0** and looks
    like a refutation. Recommend "renders the `enterprise.insights.comingSoon` string ("Coming soon")".

---

## 4. Audited zeros — read in full, found clean

Naming where I looked, per the briefing.

### `corpus/services/skillpath.md` (107 lines) — clean, and unusually well-evidenced

Every anchor in the merge banner and in "Still-true domain knowledge" resolves:

- `:14`, `:57` — `app/internal/skillpath/session.go:205-207` is the comment
  `// cms-in-app deseam: cms is in-process` (`:205-206`) immediately followed by
  `u.cms.GetSkillPathDomain(...)` (`:207`). `app/internal/skillpaths/skillpaths.go:88-95` is the same
  comment + the same call. ✔
- `:26-28` — `public.skill_path_sessions`; Ent schema files
  `app/internal/data/ent/schema/skill_path_session.go` + `skillpath_mixins.go` present. ✔
- `:30-33` — **the boldest claim in the file, and it holds.** `SkillPathSessionService` appears **0**
  times in `.go` across `stack-demo/app` (positive control: `JobSimulationService` = 7 in the same
  invocation); `grep -rln` across the whole app clone returns only `CLAUDE.md` and
  `knowledge/architecture.md`. No `skillpath…v1connect` import exists. ✔
- `:35` — the "Trap C" citation is exact: `app/CLAUDE.md:72` and `app/knowledge/architecture.md:28`
  both list `SkillPathSessionService` on the RPC mux. The doc's instruction "grade against `main.go`"
  is correct. ✔
- `:42` — `app/internal/web/backend/graphql/graph/schemas/skillpath_sessions.graphqls` exists. ✔
- `:45` — `repos.yml` at `2adcf71` has **9** entries, none of them skillpath. ✔
- `:76-77` — `InsightsSkillPathByMemberships` is at `app/internal/organization/intelligence.go:1144`;
  `:1159-1170` is exactly the `m.ent.SkillPathSession.Query()` block with `SkillPathID`,
  `StatusIn(Active, Completed)` and the `TenantIDIsNil / TenantID(organizationID)` predicate. ✔
- **`:79-87` — the RETRACTION is correct, which is the point worth recording.** It reverses earlier
  guidance and the reversal is right: `app/terraform/migrations/20260729133514.sql` contains
  `-- 5. Drop the mirrors.` at `:58`, `DROP TABLE "local_jobsimulation_sessions";` at **`:62`** and
  `DROP TABLE "local_skill_path_sessions";` at **`:63`** — matching the doc line-for-line. It is the
  last migration in the repo (`ls terraform/migrations | tail` → `20260729133514.sql`, then
  `atlas.sum`). No `local_skill_path_session.go` Ent schema exists. ✔
- `:92-99` — verified by reading `InsightsBySkillPathStudentSimulationsContainer.tsx` end-to-end:
  `userData` is `return null as unknown as MembershipEnriched` at `:33` with the real query commented
  out at `:32`; the results `<Table>` is commented out at `:140-149`; the totals block at `:150-160`;
  the body renders the comingSoon key at `:138`; the only live query is `useGetSkillPathDetails`
  (path metadata, not the session) — so "no query touches the seeded session there" is exact. ✔
  (See MINOR 10 for the one wording nit.)

### `corpus/architecture/platform-migration-status.md` (189 lines) — clean; the fence is real and green

I checked every `service_desired_count` and every `docker-compose.yml` / `repos.yml` anchor in the
fenced table, and both completeness counts:

- **All eight terraform counts exact**: `app:44 = 1`, `cms:39 = 0`, `jobsimulation:40 = 0`,
  `roadrunner:19 = 1`, `sentinel:19 = 1`, `storage:19 = 1`, `messenger:19 = 1`,
  `graphql-wundergraph:20 = 1`. The `roadrunner` row's "contradiction, recorded not resolved" is a
  faithful description of a real contradiction. ✔
- **All `repos.yml` line ranges exact**: app `10-13`, cms `14-16`, jobsimulation `17-19`, sentinel
  `20-22`, storage `23-25`, messenger `26-28`, roadrunner `29-31`, next-web-app `34-36`, studio-desk
  `37-39`; and `:14-31` is indeed the cms→roadrunner legacy span cited at `:39`. ✔
- **All compose anchors exact**: `backend:28`, `jobsimulation:83`, `cms:144`, `storage:189`,
  `customerio-sync:220-222` (`context: git@github.com:anthropos-work/customerio-sync.git#main` at
  `:222`, profile at `:238`), `messenger:240`, `roadrunner:281`, `studio-desk:311`,
  `next-web-app:344`, `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` at `:352` (the build arg — it also recurs at
  `:361` as an env var; the cited line resolves), `gotenberg:371-372` (`image: gotenberg/gotenberg:8`
  at `:372`), sentinel's `search_path=sentinel` at `:18`, `include: - common.yml` at `:1-2`,
  `common.yml:2` postgresql, `common.yml:20` redis. ✔
- **The §1 "trap 2" claim is TRUE and is the most valuable thing in the file.** `repos.yml` calls
  cms/jobsimulation/roadrunner "legacy — folded into app" while `docker-compose.yml` still defines all
  three **in the default `graphql` profile** — confirmed: `profiles: [graphql, jobsimulation, all]`
  (`:140`), `[graphql, cms, all]` (`:187`), `[graphql, roadrunner, all]` (`:309`). This directly
  contradicts the repo-root `CLAUDE.md` banner ("There is no cms / jobsimulation / skiller / skillpath
  / roadrunner container, profile, port or subgraph"). **This file is the one that is right**; I note
  the conflict for whoever owns `CLAUDE.md`, which is outside my seat. ✔
- **`app/main.go` per-domain wiring — the M257x iter-46 correction is correct.** `:573` is
  `skiller.NewSkillerManager(...)`, `:604` is `jobsimwiring.Wire(...)`, `:634` is
  `skillpath.NewSessionManager(...)`, `:1034` is `appcms.Wire(appcms.Deps{`. Four distinct call sites,
  exactly as claimed. ✔
- **`app/internal/roadrunner/` genuinely does not exist** (`ls -d` → No such file or directory), and
  `app/internal/jobsimulation/runner/` does; `app/internal/jobsimwiring/wiring.go:118` is
  `runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`. ✔
- **Supergraph = one subgraph, verified three ways**: `graphql-wundergraph/supergraph-config-prod.yaml`
  lists `backend` alone (with an in-file comment "cms-in-app (v8.0): the cms subgraph is folded into
  backend (supergraph 2→1)"), `schemas/` holds `backend.graphqls` alone,
  `subgraphs.conf` is `BACKEND=v1.360.0`. ✔
- **Both completeness counts reproduce exactly.** `git log -p --follow -- repos.yml` → **14** unique
  names, and they are precisely the 14 the doc lists in order. `git log -p --follow --
  docker-compose.yml` → **26** unique service names, including the five "pre-history" names cited. ✔
- **§4's fence executes and passes.** `PLATFORM_REPOS_YML=stack-demo/platform/repos.yml python3
  .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py` →
  `platform_alignment_guard: OK — platform-migration-status.md and repos.yml agree in both directions.`,
  exit 0. The guard is not vapor. ✔

### `corpus/services/jobsimulation.md` (226 lines) — clean

- `:8`, `:11-14` — `jobsimulation/terraform/main.tf:40 = 0`; `docker-compose.yml:83`; `repos.yml:17-19`
  with the `migrations: false # legacy` comment. The `running_but_unfederated` framing matches the
  compose reality I confirmed above. ✔
- `:24-30` — **the "headline table renamed" correction is exact and important.** The jobsim data-model
  migration creates exactly **23** tables (`grep -c '^CREATE TABLE'` → 23), and the names the doc cites
  (`sessions`, `actors`, `interactions`, `validation_*`, `anticheat_*`, `recordings`,
  `chime_recordings`, `code_submissions`) are all in the list. `20260722104506.sql:2` is
  `CREATE TABLE "job_simulation_sessions" (` and `:79` is `DROP TABLE "sessions";` — so
  "**`public.sessions` does not exist**" is a correctly-derived net statement, not an assumption. ✔
- `:31-33`, `:95` — `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` at `docker-compose.yml:52`
  (backend) and `:258` (messenger); `app/main.go:1195` registers
  `jobsimulationv1connect.NewJobSimulationServiceHandler`. The doc's emphatic "**That address is
  CURRENT, not stale text**" is right. ✔ (anchor nit at `:34` → MINOR 5)
- `:68` — the dual-port explanation holds: compose sets `PORT=8400`/`RPC_PORT=8401` and publishes
  `8400:8400`/`8401:8401`, while the repo's `cmd/root.go:77-78` is
  `cmp.Or(os.Getenv("PORT"), "8080")` / `cmp.Or(os.Getenv("RPC_PORT"), "8081")`. Both-are-correct
  framing verified. ✔
- `:81` — worker pools: `internal/worker/worker.go:27` `Concurrency: 10`, `:104` `Concurrency: 25`. ✔
- `:88-90` — `hibiken/asynq v0.26.0` in `go.mod:24`; the interview migrations `20260402145459` and
  `20260409131539_add_summary_to_extraction_results` exist under `terraform/migrations/` (**note:** my
  first probe looked in a nonexistent top-level `migrations/` and found nothing — a false absence I
  caught and corrected); `interview_extraction_results` appears in the Ent layer
  (`internal/ent/interviewextractionresult/`, `internal/ent/migrate/schema.go`). ✔
- `:99` — `internal/graph/queries.resolvers.go:70` is exactly
  `func (r *queryResolver) JobSimulationResult(ctx context.Context, sessionID uuid.UUID)`. ✔
- `:101-103` — the mirrors-are-gone correction matches `20260729133514.sql:58-62` (`-- 5. Drop the
  mirrors.` at `:58`, `DROP TABLE "local_jobsimulation_sessions"` at `:62`) and
  `app/internal/organization/intelligence.go:1700` = `query := m.ent.JobSimulationSession.Query().`. ✔
- `:112` — `CMS_RPC_ADDR=http://cms:8091` is at `docker-compose.yml:104` (inside the jobsim block);
  `app/cms_reader_switch.go` exists; `app/main.go:971-973` is
  `directusBaseAddr := os.Getenv("DIRECTUS_BASE_ADDR")` / `if directusBaseAddr == ""` /
  `log.Fatalf("DIRECTUS_BASE_ADDR is required …")` — the cited `log.Fatalf` is exactly there. ✔
- `:116` — `jobsimulation/internal/runner/runner.go` header reads *"…(formerly the standalone
  \"roadrunner\" service)"*, word-for-word as quoted. ✔
- `:136-148` — **the startup contract is exactly right.** `cmd/` holds `aggregate.go`,
  `clone_session.go`, `test.go`, `validate.go` with `Use:` values `aggregate`, `clone-session`,
  `test-command`, `validate` — four subcommands, **no `serve`, no `run`**. Root `Use: "jobimulation"`
  (sic, upstream typo). ✔
- `:151` — "sets neither `SilenceUsage` nor `SilenceErrors`": `grep -rn` over `cmd/` returns nothing,
  rc=1, against a `grep -c 'RunE' cmd/root.go` = 1 positive control. ✔
- `:166` — "the **only** AWS bind in the file": `grep -n '.aws' platform/docker-compose.yml` returns
  exactly one line, `:142`, `$HOME/.aws/credentials:/root/.aws/credentials:ro`. ✔

### `corpus/services/ant-academy.md` (436 lines) — clean apart from MINORS 1-4

The section the doc itself flags as highest-risk is the one I checked hardest, and it is complete:

- **`:250-267` the `isPublic` matcher — the doc says "Do not paraphrase this from memory", and it
  didn't.** I enumerated `code/proxy.js:112-188` in full. Every one of the 25 static patterns plus both
  conditional spreads is in the doc's table, and the table adds no pattern that isn't in the code:
  `/api/_meta(.*)`, `/api/meta(.*)`, `/robots.txt`, `/sitemap.xml`, `/sitemap(.*)`, `/llms.txt`,
  `/llms-full.txt`, `/.well-known/(.*)`, `/courses`, `/courses/(.*)`, `/sign-in(.*)`,
  `/no-organization`, `/verify/(.*)`, `/api/verify/(.*)`, `/api/ai/chat`, `/library`, `/library/(.*)`,
  `/free`, `/free/(.*)`, `/local-content/(.*)`, `/catalog.json`, `/academy-manifest.json`, `/`,
  `/latest(.*)`, `/chapters/(.*)`, `…(DEV_LOGIN_ENABLED ? ["/api/dev/login-as", "/dev/accept"] : [])`,
  `…(VISUAL_BYPASS ? ["/my-certificates", "/my-activity", "/bookmarks"] : [])`. The doc's grouping and
  each "why public" rationale match the in-code comments. ✔
- `:278-280` — "**There is no `/library/[slug]` route**; the per-course page is `/courses/[slug]`":
  `app/(public)/library/` contains only `page.jsx`, `app/(public)/free/` only `page.jsx`, and
  `app/(authed)/courses/[slug]/page.jsx` exists. The CTA guidance is correct. ✔
- `:258` — "The pages still live in the `(authed)` route group": confirmed, `/courses` is under
  `app/(authed)/courses/`. ✔
- `:66-68` — the "⚠️ It does *not* fall back to the committed FS catalog" correction is right:
  `serverTenant.js:115-145` contains both `emptyCatalogView()` (`:115`) and the quoted in-code line
  *"the cutover is intentional, not reversible-on-error"*, and `getServerCatalogView()` at `:143-146`
  is `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView()`. ✔ (see MINOR 1)
- `:124-134` — every file in the read chain exists: `src/lib/serverTenant.js`,
  `src/lib/backendContent.js`, `src/graphql/server.js`. ✔
- `:156-157` — `src/lib/draftMode.js:46-50` is
  `NODE_ENV === 'development' && isDraftOptIn(process.env.ACADEMY_SHOW_DRAFTS)`; `src/lib/draftCatalog.js`
  exists. The "production hard-block, whitelisted on 'development'" framing is accurate. ✔
- `:171-180` — `src/lib/serverChapterBody.js:51-52` is
  `export async function resolveServerChapterBody(slug, locale) { const body = await getBackendChapterBody(slug, locale)`;
  `return { notFound: true }` at `:67`; `app/not-found.jsx:43` reads `You wandered off the trail.` ✔
- `:219-233` — `src/i18n/locale.js`, `src/i18n/LocaleSwitch.jsx` and `src/i18n/translate.js` all exist;
  `LocaleSwitch.jsx:10` is `const target = locale === 'it' ? 'en' : 'it'` — a 2-way EN↔IT toggle, not a
  dropdown, exactly as claimed. ✔
- `:34` — `src/components/RegisterServiceWorker.jsx` is verbatim a kill-switch: its header reads
  "Service-worker KILL-SWITCH (v0.5 "direct line" M1 — offline removal)" and "This component no longer
  REGISTERS anything." ✔
- `:240` — `REQUIRE_ORGANIZATION_MEMBERSHIP` default-on confirmed (`proxy.js:5-6` "default on",
  `:90-98` the `ORG_GATE_ENABLED` warn-when-off block). ✔
- `:243-248` — FA Pro vendored at `code/public/assets/fontawesome/{css,webfonts}` ✔;
  `code/app/layout.jsx:132` is `manifest: "/academy-manifest.json"` ✔; `code/vercel.json` is exactly
  `{"framework": "nextjs"}` ✔; `code/package.json:8` `"node": ">=22"` ✔.
- `:352` — `"dev": "next dev --port 3077"` at `code/package.json:21`. ✔
- `:408` — the Cosmo claims are all exact:
  `ucourses/ucourse-engine/assistant/agent.js:12` `API_URL = 'https://api.openai.com/v1/responses'`,
  `:13` `MODEL = 'gpt-5.2'`, `:32` `localStorage.getItem('openai_api_key')`, and
  `src/lib/featureFlags.js:8` gates on `NEXT_PUBLIC_FEATURE_TRAINING_COACH`. ✔
- `:26`, `:293`, `:424-426` — the thrice-repeated "NOT in `repos.yml`" claim holds: the 9 entries at
  `2adcf71` are app, cms, jobsimulation, sentinel, storage, messenger, roadrunner, next-web-app,
  studio-desk. ✔
- `:392` — `tests/e2e/i18n-language-toggle.spec.js` exists. ✔
- `:383-390` — the skills table matches `ant-academy/.claude/skills/`; the directory also holds
  `export-path` and `_shared`, which the doc's open-ended "Other content-pipeline helpers" row covers.

### `corpus/services/studio-desk.md` (435 lines) — clean

- `:21`, `:140` — `docker-compose.yml:337-341` is precisely the `depends_on: backend / cms` block
  (`:337` `depends_on:`, `:338-339` backend, `:340-341` cms), and `:334` is
  `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`. The "**not** `graphql`, which is no
  longer a compose service" claim is right — there is no `graphql` service at `2adcf71`. ✔
- `:34` — `vite.config.ts:28-42` `rollupOptions.input` lists exactly the nine prod entries the doc
  names (home/simulation-builder/sim-advanced-builder/sim-guided-builder/builder-skill-path/generation/
  catalog/academy/skills) with `dev-accept` behind `...(isProduction ? {} : {...})`. ✔
- `:107` — `src/routes/youtube.ts:43` is `const apiKey = process.env.YOUTUBE_API_KEY;` with the
  `_mock: true` fallback at `:46`; **`GCLOUD_SERVICE_ACCOUNT` is read by nothing in `src/`**
  (`grep -rn` empty, in a pass where sibling greps returned hits), while it is declared at
  `.env.example:120` and injected at `terraform/main.tf:129`. The "vestigial, not a second YouTube
  credential" verdict is correct. ✔
- `:112` — `AI_PROVIDER_CHAIN=azure-openai,openai` at `.env.example:57`; `AI_DEFAULT_TIER=fast` at
  `.env.example:61`; and the in-code fallback is exactly
  `defaultTier: (env.AI_DEFAULT_TIER as ModelTier) || 'thinking_fast'`
  (`src/services/ai/config.ts:182`) — the doc's parenthetical is precisely right. ✔ (models: MINOR 9)
- `:161` — `package.json:7` `"node": ">=24"`, `Dockerfile:1` and `Dockerfile.dev:1` both
  `FROM node:24-alpine`. ✔
- `:172` — "in-code fallback for PORT is 9100": `src/index.ts:60` is
  `const backendPort = process.env.PORT || 9100;`. A genuinely counter-intuitive fact, correctly
  recorded. ✔
- `:279-287` — **the no-datastore claim is right, negatively verified.** `package.json` matches none of
  `pg|postgres|prisma|sqlite|mysql|mongodb|mongoose|knex|drizzle-orm|typeorm|sequelize` (rc=1) against a
  positive control that found `"express"` and `"@clerk/*"` in the same file.
  `app/services/studioDB.js` exists and is a facade; `GET_USER_STUDIO_PREFERENCES` /
  `SET_USER_STUDIO_PREFERENCES` are real (`app/services/graphql/queries.ts:205`, `mutations.ts:58`).
  The Tailscale-funnel aside is exact: `app/core/main.ts:105` is the GlitchTip comment inside a
  `NODE_ENV === 'production'` Sentry block. ✔
- `:305` — `src/index.ts:96` is
  `const STUDIO_ACCESS_ROLES = ['admin', 'org:admin', 'content_creator', 'org:content_creator'];`,
  consumed by `checkEnterpriseAndAdmin` at `:99` / `:115`. The "content creators, not only org admins"
  correction is right. ✔
- `:313-315` — **all three prod-host hardcodes exist, exactly where claimed**:
  `app/core/scaffold/pageWrapper.js:149` (logo link),
  `app/core/scaffold/userProfile.js:148` (menu "Back") and `:302` (logout redirect), all
  `https://app.anthropos.work`. ✔
- `:388`, `:393` — `app/core/main.ts:97` is `preloadCriticalCSS();` and `:206` is `new PageWrapper();`,
  matching the boot-order diagram's two cited line numbers. ✔
- `:420` — `canAccess()` reaching `clerk.user.getOrganizationMemberships()` is confirmed
  (`app/services/userService.ts:213`, `:259`). ✔
- `:432` — the `./studio-room.md` relative link resolves (`corpus/services/studio-room.md` exists).

### `corpus/architecture/frontend_architecture.md` (105 lines) — clean apart from MINORS 6-7

- `:11` — `docker-compose.yml:311` is the `studio-desk:` service block; `repos.yml` @ `2adcf71` holds
  exactly **9** entries; ant-academy is not among them. All three sub-claims verified. ✔
- `:21-25` — ports confirmed from each app's `dev` script: web `--port 3000`, hiring `--port 3001`,
  integration `--port 3002`; `maintenance` has no `dev` script (doc shows "—"); mobile is
  `"dev:mobile": "expo start --port 3031 -c"`. ✔
- `:24` — `!apps/mobile` is present in `pnpm-workspace.yaml`, under the comment "exclude mobile app to
  speed up development". ✔
- `:31-35` — package identities exact: `@anthropos/ui`, `@anthropos/graphql`, `@anthropos/core-js`,
  `@anthropos/tsconfig` (at `configs/tsconfig`), `@anthropos/i18n` (at `configs/i18n`). The 8 locales
  are exactly `de en es fr it ja nl pt` (`ls configs/i18n/messages`). ✔
  *(Aside: `next-web-app/CLAUDE.md` says "7 locales" in three places — the corpus's 8 is the one that
  matches the filesystem. The corpus is right and the repo's own guide is stale.)*
- `:39` — `NEXT_PUBLIC_BACKEND_API_URL` is at `docker-compose.yml:362` and its value is
  `http://…:8082`; `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` is `…:8082/graphql/query`. The "`:5050` was only
  ever the deleted LOCAL compose host mapping, never a production port" retraction is consistent with
  `2adcf71`, where no `:5050` mapping and no `graphql` service exist. ✔ (count: MINOR 6)
- `:44` — `packages/graphql/src/__generated__/` exists with the client-preset output
  (`fragment-masking.ts`, `gql.ts`, `graphql.ts`); `graphql-request ^7.4.0` and
  `@tanstack/react-query ^5.99.2` are both declared. ✔
- `:48-52` — `apps/web/package.json:46` is `"next": "^16.2.7"` (the cited line number is exact);
  root `package.json:20` `"packageManager": "pnpm@10.30.3"`; `:14` `"node": ">=24.0.0"`. ✔

---

## 5. Unverified — could not check, and why

These are stated as unverified, **not** as passes and **not** as blockers, per the briefing.

1. **All GitHub repo-archival claims.** `gh` is unavailable and these repos aren't reachable from the
   clones. Affects: `jobsimulation` "ARCHIVED 2026-07-31" (`jobsimulation.md:13`,
   `platform-migration-status.md:62`), `cms` "**not** archived" (`:61`), `roadrunner` "not archived"
   (`:63`), `skiller` "ARCHIVED 2026-07-01" (`:70`), `skillpath` "ARCHIVED 2026-07-31" (`:71`),
   `chronos` "NOT archived" (`:72`), `intelligence` "ARCHIVED 2026-04-02" (`:73`),
   `graphql-wundergraph` "repo is ARCHIVED 2026-07-30" (`:69`).
2. **The whole §3 net-new census** (`platform-migration-status.md:97-136`) — "93 repositories", "46
   named by no corpus document", and all 19 per-repo last-push dates, plus the `auth` and `AI-Labs`
   dispositions. Requires the org API.
3. **`app` PR #1103** (the v9.0 storage+messenger fold, cited at `:65`, `:66`, `:170`) and
   **jobsimulation PR #395** (`jobsimulation.md:88`). Not checkable without `gh`.
4. **Individual platform/app commit shas → dates.** I reproduced the two `git log` **aggregate** counts
   the doc offers as its completeness proof (14 and 26, both exact), which exercises the same history;
   but I did not map each cited sha (`236771f`, `21429b7`, `a4db680`, `045857c`, `fdfa189`, `b56d731`,
   `360efd4`, `915da06`, `a2a3ee6`, `cb6ebf5`, `1474b1f`, `84862d1`, `b43b99a`, `c17cc9a`, `467965a`,
   `8770fe6`, `ef4b449`, `05b4035`, `09631fb2`) to its claimed date.
5. **All demo-runtime and measured-performance claims.** No stack was brought up. Affects: the "65 real
   cards" grid (`ant-academy.md:74`), the M230/M238/M245 demopatch behaviours and their "proven live on
   `billion`" HTTP 404→200 results, the M249/M252/M253 studio-desk demo sections
   (`studio-desk.md:309-428`) including "skeleton-visible p95 4669 ms → 817 ms" and the
   "`canAccess` 4049 → 38 ms / FCP 6936 → 2152 ms" table, and the M257x iter-24 "96 Directus log lines,
   all 403" measurement (`jobsimulation.md:112`).
6. **`rosetta-extensions` internals** beyond the one guard I executed. I ran
   `platform_alignment_guard.py` (exit 0) but did not exercise `demopatch`, `ensure-clones.sh`,
   `demo-stack/ant-academy.sh`, `gen_injected_override.py` or the named pytest cases
   (`test_studio_desk_env_clerkenstein_no_mock_and_offset_sign_in`,
   `test_studio_desk_block_shape_single_port_clerkenstein_wired`).
7. **`colony` / `proto` / `taxonomy`** — not cloned, per the briefing. Affects the library rows at
   `platform-migration-status.md:79-83` and the "colony keys by stream name" assertion at
   `jobsimulation.md:40`.
8. **"1000+ Vitest tests"** (`ant-academy.md:246`). I counted **244** test *files* under `tests/` +
   `src/`; the claim is about individual test cases, which needs a run. Plausible at ~4 cases/file, but
   not measured. (The Playwright half of the same sentence *is* measured — see MINOR 2.)
9. **Production-only runtime claims** — the Cosmo Router serving `:8080/graphql` in prod
   (`frontend_architecture.md:39`, `studio-desk.md:132`), and
   `JOBSIMULATION_RPC_ADDR=http://backend.internal.anthropos:8081` in production
   (`jobsimulation.md:33`). The prod routing_url in `supergraph-config-prod.yaml` is
   `http://backend.internal.anthropos:8080/graphql/query`, which is consistent, but no prod environment
   was reachable.
10. **Historical-state claims** that describe a past sha rather than `2adcf71`, e.g.
    `skillpath.md:38` "the supergraph is **3 subgraphs** at the time". The *present*-tense half of that
    sentence ("it is **1** now") is verified; the historical half is not.
