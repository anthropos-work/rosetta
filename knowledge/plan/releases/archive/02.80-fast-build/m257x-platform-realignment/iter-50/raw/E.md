# Seat E — M257x clause-5 KB-fidelity reading

## 1. Header

**Corpus under audit:** `/Users/marco/workspace/anthropos/rosetta`, branch `m257x/platform-realignment`,
HEAD `57dfbfd` (confirmed via `git rev-parse --short HEAD` + `git branch --show-current`).

**Ground-truth clones consulted** (sha confirmed with `git log -1 --format=%h` in each):

| clone | sha seen | used for |
|---|---|---|
| `stack-demo/platform` | `2adcf71` | compose, `repos.yml`, `common.yml`, git history (`a2a3ee6`, `b56d731`, `360efd4`, `045857c`, `fdfa189`) |
| `stack-demo/app` | `5ba17044` (`.version` = `v1.363.2`) | migrations, Ent schema/enums, resolvers, `IntelligenceManager`, askengine, main.go wiring, go.mod, CI |
| `stack-demo/app/studio` | `aeec036` | `gen.py` argparse surface, `studio/CLAUDE.md` |
| `stack-demo/next-web-app` | `bb3313bc0` | `apps/web` + `apps/hiring` hiring anchors |
| `stack-demo/storage` | `4ce8ece` | storage.md end-to-end |
| `stack-demo/studio-desk` | `14a5442` | framework/model/language claims |
| `stack-demo/ant-academy` | `9c3843cd` | academy GraphQL/serwist/manifest claims |
| `stack-demo/graphql-wundergraph` | `60c229f` | terraform ports, supergraph-config history |
| `stack-demo/{sentinel,messenger,cms,jobsimulation,roadrunner}` | — | `go.mod` shared-library pins |
| `.agentspace/rosetta-extensions` | `a91f8f7` | seeder/patch/e2e-gate claims |
| `.agentspace/snapshots/{taxonomy,directus}` | — | taxonomy row counts, hiring-sim pool |

**Positive control — `wc -l` on every assigned file** (one invocation:
`wc -l corpus/architecture/service_taxonomy.md corpus/services/hiring.md corpus/architecture/shared_libraries.md corpus/services/storage.md corpus/services/askengine.md corpus/architecture/dependency_map.md`):

| file | lines | read in full |
|---|---:|---|
| `corpus/architecture/service_taxonomy.md` | 440 | yes, 1→440 |
| `corpus/services/hiring.md` | 398 | yes, 1→398 |
| `corpus/architecture/shared_libraries.md` | 242 | yes, 1→242 |
| `corpus/services/storage.md` | 175 | yes, 1→175 |
| `corpus/services/askengine.md` | 121 | yes, 1→121 |
| `corpus/architecture/dependency_map.md` | 103 | yes, 1→103 |
| **total** | **1479** | |

**Search-pipeline hygiene.** Two searches in this pass silently failed and were re-run:
(a) `grep -rn ROADRUNNER_RPC_ADDR app/ --include=*.go` died with zsh
`no matches found: --include=*.go` — an engine rejection that reads identically to "0 hits"; re-run
with the pattern quoted (`--include='*.go'`) against a positive control (`STORAGE_RPC_ADDR` → 3 hits)
before concluding absence. (b) a `grep` pattern using `\|` under `ugrep` errored
(`invalid escape`) — re-run as separate patterns. Every "0 hits" reported below was paired with a
control that returned non-zero in the same invocation.

---

## 2. BLOCKERS

**None. 0 blockers.**

Every load-bearing `file:line` anchor in these six files that I could reach resolved to what the text
says is there, and every actionable claim I could measure was true. The anchor set I re-derived
independently and confirmed exact (not exhaustive, but the load-bearing spine):

- **hiring.md's whole read-path table** — `resolver_queries.go:1034` (decl), `:1035`
  (`OrgFeatureInsights`), `:1053` (status ∈ {active,invited}), `:1080` (manager call), `:1085`/`:1089`
  (the *neighbouring* resolver's gate + `GetMembership` — the distinction the doc draws is real);
  `intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`), `:1820` (`RoundFloat(float64(ls.Score))`),
  `:1844`, `:1846`, `:885-886` (`InsightsSortFieldCompletitionStatus → FieldCompletionStatus`);
  `ent/schema/job_simulation_session.go:39,41,45`; `enum/jobsimulation.go:29-35` + `Values()` `:37-43`
  (exactly 5 lowercase members); `ent/jobsimulationsession.go:181-186` (unconditional cast, cannot error);
  `graph.go:129546-129554` **and** the proto-bound twin `:129392-129400` (bare `MarshalString`, no
  membership check); `jobsimulations.graphqls:14` + `:128`.
- **The M257 migration facts.** `20260729133514.sql:58` (`-- 5. Drop the mirrors.`) → `:62`
  `DROP TABLE "local_jobsimulation_sessions"`, preceded by the FK re-point at `:51-56`; **`SET "score"`
  returns 0 hits across `terraform/migrations/`** (positive control in the same pass: `DROP TABLE` → 11
  hits) — so "no back-fill" is measured, not asserted. `20260722104506.sql:2` (CREATE
  `job_simulation_sessions`), `:12` (`completion_status`, correctly spelled, plain `varchar`, **no
  CHECK**), `:13` (`token` NOT NULL), `:29` (UNIQUE index on token), `:53` (anticheat FK re-point),
  `:79` (`DROP TABLE "sessions"`). `anticheat_summary` exists only on the dropped mirror
  (`20250416091037.sql:5`) and on no Ent field — the "do NOT write it" warning is correct.
  `app/atlas.hcl:8` pins `search_path=public`; the only `CREATE SCHEMA` in the whole set is `auth`
  (`20230817154747_supabase_baseline.sql:2`); `askengine/registry.go:192` does say the `jobsimulation`
  schema stays frozen until M710. The iter-49 corrections to this passage are all sound.
- **hiring.md's grep-based negative claims, re-run independently:** `grep -in hiring` over
  `resolver_queries.go` → 4 hits, all sim-TYPE filters/content-permission strings (`:811,843,991,992`);
  over `intelligence.go` → 0. The doc's stated positive controls reproduce **exactly**:
  `OrgFeatureInsights` ×8, `JobSimulationSession` ×44.
- **The `is_hiring` blast radius.** `resolver_cms_queries.go:95,210,258,295` (`GetOrganizationIsHiring`)
  and `:99-103` (`hiringLibraryTypes` vs `workforceLibraryTypes`); `organization/manager.go:448`
  (`switch org.IsHiring` → `RoleCandidate`) and `:485` + `siminvitationlink.go:62` (both hard-error
  `"organization is not hiring"`). Client side: `apps/web/src/hooks/useGetClerkOrganization.tsx:16-18,20-21`
  (verbatim as quoted), `useNavbarSections.tsx:459-466` with the label ternary at `:460`,
  `template.tsx:90` `isEnterprise = Boolean(organization)`, `FreeTrialContainer.tsx:29`
  `Boolean(!isHiringOrg && organizationId)`, `apps/web/.../UserStatusContext.tsx:144-149` + `:168-172`
  (`targetProduct:'hiring'`), `apps/hiring/.../UserStatusContext.tsx:125` + `:144-145`
  (`targetProduct:'workforce'`). The "zero intercepting routes in `apps/`" claim holds
  (`find apps -type d -name "(.*"` → empty; control `-name "(*"` → `apps/mobile/app/(protected)` etc.).
- **The supergraph ladder** (`service_taxonomy.md:339`, a repaired passage that corrects a commit's own
  title — and the correction is right): `749dc86^` prod config lists 5 subgraph names, `7c17e63^` lists
  4, `915da06^` lists 3, and `915da06 --stat` deletes **both** `schemas/cms.graphqls` (762 lines) **and**
  `schemas/jobsimulation.graphqls` (860 lines) in one commit. So "cms-in-app is the 3 → 1 step, not
  2 → 1" is measured. Router ports: `terraform/locals.tf:8` `port = 8080`, `main.tf:48-49` map
  `${local.port}`→`${local.port}`, `config.prod.yaml:5` `listen_addr: 0.0.0.0:8080`, `main.tf:20`
  `service_desired_count = 1`. `grep -rn 5050 platform/` returns **only** `platform/CLAUDE.md:29`
  ("Port 5050 is free") — no compose mapping anywhere.
- **The Directus retraction-of-a-retraction** (`service_taxonomy.md:300-305`): `git show
  a2a3ee6^:docker-compose.yml` → `:384 image: directus/directus:10.10.1`, `:386 - 8055:8055`,
  `:409 ADMIN_PASSWORD=password` — all three exact; and `git log --all -S "admin@example.com"` → 0
  commits, matching the doc's own "only the email is unfound". The over-correction is correctly fenced.
- **compose/`repos.yml` shape:** services at `sentinel:5, backend:28, jobsimulation:83, cms:144,
  storage:189, customerio-sync:220, messenger:240, roadrunner:281, studio-desk:311, next-web-app:344,
  gotenberg:371` — **no `graphql` service, no `directus` service**; `postgresql`/`redis` come from
  `common.yml` (`redis` image `bitnamilegacy/redis:latest`, `:5432`/`:6379`). `repos.yml:10-13` = the
  `app`/`migrations: true`/`schema: public` block; `:14-19` = cms + jobsimulation; no
  `graphql-wundergraph` entry. `docker-compose.yml:18` is sentinel's `search_path=sentinel`.
  Frontend endpoints baked at `:318`/`:334` (studio-desk `VITE_GRAPHQL_ENDPOINT`) and `:352`/`:361`
  (next-web `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`), both `…:8082/graphql/query`. Backend `depends_on`
  `:70-80` really does include the cms husk. Messenger `:255,256,258,265` exactly as documented.
  Studio-desk `depends_on` `:337-341` = backend + cms; `profiles` at `:342`.
- **shared-library pins — every single one exact** from the seven `go.mod`s: colony
  `app`/`messenger` `v0.35.2`, `cms`/`jobsimulation` `v0.35.1`, `sentinel`/`storage`/`roadrunner`
  `v0.34.3`; proto `1.210.0 / 1.207.0 / 1.205.0 / 1.200.0 / 1.196.0`; ai `v1.40.2` at `cms/go.mod:9`
  and `jobsimulation/go.mod:11`; taxonomy `v1.2.0` at `cms/go.mod:13` and `jobsimulation/go.mod:15`
  (neither `// indirect`), `// indirect` at `sentinel:21` / `storage:25`; roadrunner requires only
  colony + proto. `app/main.go` registers exactly six Connect handlers at `:1178, :1179, :1187, :1195,
  :1204, :1218` — no `SkillPathSessionService`, no RoadRunner. `skillpaths.go:27-31` does call itself
  "the drop-in for the **removed** skillpath RPC client"; `jobsimwiring/wiring.go:118` is the Judge0
  call; `ROADRUNNER_RPC_ADDR` (present at `docker-compose.yml:118`, in the *jobsimulation* block) is
  read by **no** Go file in `app` (verified with a working control). `roadrunner/main.go:7` imports
  colony. No `go.work` committed anywhere. `storage/internal/migration/migration.go:13` imports
  `go/simulator/storage/v1` as `legacyStorage`.
- **The taxonomy figures.** `.agentspace/snapshots/taxonomy/…/manifest.json` (`source: primary-read`,
  `public_only: true`, `predicate: org-null`, `captured_at: 2026-06-29`) gives `skills` **42,790**,
  `job_roles` **22,470**, `job_role_embeddings` **18,919** — i.e. 18K is below the public floor
  (REFUTED) and 60K is neither supported nor excluded (UNVERIFIED). The doc's framing is exactly right.
- **storage.md end-to-end** (see Audited zeros).
- **askengine.md end-to-end** (see Audited zeros) — including two places where the corpus is *more*
  correct than the platform's own comments (noted below).
- **dependency_map.md's Redis-Stream table** against `app/main.go:1265-1311`: skillpath `:1265`,
  skiller `:1267`, jobsimulation on **ONE** subscriber merging JSManager + skillPathSessionManager +
  jobsimEngine `:1275-1276`, cms on **ONE** subscriber `:1278-1294`, AI `:1296`, backend self-stream
  `:1304-1311`; and `:1297-1298` confirms the roadrunner stream was removed upstream (no producer, no
  consumer). Storage really has no `DB_CONNECTION`/`REDIS_ADDR` in compose and no redis in `go.mod`.
- **hiring seeder/tooling claims:** `stack-seeding/seeders/persona_write.go:91`
  (`{"public","job_simulation_sessions",sessionCols(),…}`), `sessionCols()` at `:152` with `"token"` in
  the list at `:155`; `hiringComparableFloor = 6` (`contentref.go:158`) and
  `RENDER_GATE_FLOOR ?? '6'` (`render-hiring-comparison.spec.ts:56`) + `RENDER_ONLY_SIM` (`:245`);
  `TestGenericActivitySeeders_SkipHiringOrg` enumerates exactly **8** seeders (`hiring_scope_test.go:81-95`,
  including the two M228 additions feedback + succession); `skipGenericActivityForHiringOrg` →
  `st.IsHiringOrg()` (`hiring_scope.go:45-46`); the four demo-patches named at hiring.md:350 all exist
  under `demo-stack/patches/`; the no-net-new-grant claim holds — `policy_grants.go:82`
  `{"default","admin","org:feature:insights"}` is a **platform-default** p3 row.
- **The hiring content pool.** Parsing the captured `directus.simulations.copy` against its manifest
  column list: 307 published+public rows, of which **87** are `SIMULATION_TYPE_HIRING` — the doc's
  number, exact. No `job_position` **table** is in the capture at all, so "0 `job_position` rows" holds.
- **studio / academy claims:** `gen.py` registers exactly the nine listed arguments (`:484-492`), uses
  `parse_known_args` (`:19`), and has no `--template`; `studio/CLAUDE.md:12-14` is quoted verbatim;
  `additional_repo: "anthropos-studio-room:studio"` at `build-production.yml:29`; `app/internal/cms/studio`
  exists. studio-desk: **0** react/vue/angular entries in `package.json`, **0** `.tsx`/`.jsx`
  (control: 71 `.ts` files), gpt-5.x model map, and exactly **7** language flags in
  `app/public/languages/`. ant-academy: `code/src/graphql/server.js:14,18` (throws when unset),
  `backendContent.js:36,102-103`, `serverTenant.js:115-145` (the "NO FS-as-published fallback …
  not reversible-on-error" text is at `:129,132`, the code at `:145`), `beacon/route.js:36,41-55`,
  `layout.jsx:132`, `academy-manifest.json` `display: standalone`, **no** serwist/workbox dependency,
  `next-scaffold.test.js:106,111` + `react-compiler-config.test.js:41` are the regression fence,
  `RegisterServiceWorker.jsx` is a kill-switch (`:43` `r.unregister()`), Next `^16.2.6` / React
  `^19.2.5` / Expo `~54.0.33`, ports 3077 / 8555, `gpt-5.2` at `ucourses/…/agent.js:13`, and the rext
  patch `demo-stack/patches/academy-fs-published-fallback` exists.

---

## 3. MINORS

Nine. None gate the milestone; all are cheap to fix.

1. **`corpus/services/hiring.md:206-208`** — *"`token` … is the **only** required-and-undefaulted column
   in the table."* False as a superlative: `20260722104506.sql:6` (`owner_id`), `:7` (`sim_id`) and
   `:10` (`sim_type`) are equally `NOT NULL` with no `DEFAULT`. The *actionable* guidance is unharmed —
   all three are already listed as required in the same sentence, and "write a unique token per row" is
   correct (`ent/schema/job_simulation_session.go:41` also constrains it to `^[a-z0-9]+$`, 5–10 chars,
   which the doc does not mention). Suggested fix: drop "the only", or say "the only one this contract
   previously omitted".
2. **`corpus/services/hiring.md:304`** — *"Meridian Talent … 5 admins + 45 candidates"*. The preset is
   `size: 50, role_mix: { admin: 0.14, member: 0.0, candidate: 0.86 }`
   (`rosetta-extensions/stack-seeding/presets/stories.seed.yaml:249`), and `roleForIndex`
   (`stack-seeding/seeders/users.go:571-573`) computes `adminCount = int(50 × 0.14) = 7` → **7 admins +
   43 candidates**. hiring.md's own render tally at `:381` (8+8+9+9+8) sums to **43**, so the document
   already contradicts itself. Note the preset's own header comment (`stories.seed.yaml:3-4`) carries the
   same stale figure — fix both or neither.
3. **`corpus/services/hiring.md:127`** — *"`role_mix ≈ 0.1 admin / 0.9 candidate`"*; actual `0.14 / 0.86`.
   Hedged with `≈`, so borderline; folded here with #2 because they share one root.
4. **`corpus/services/hiring.md:169`** (read-path table row 5) — cites `intelligence.go:1728-1735` for
   *"best-attempt: `row_number() ORDER BY score DESC` per candidate"*. That range is the
   `JsSessionID` collection plus the **call** at `:1733`; the `ROW_NUMBER() OVER (PARTITION BY … ORDER BY
   … DESC …)` literal lives in `usersBestOrFirstJobSimulationSession` at `intelligence.go:2158`. Anchor
   points at the call site, not the mechanism it names.
5. **`corpus/services/hiring.md:173`** — cites `intelligence.go:1738-1751` for the sort trio
   *"score DESC, completition_status ASC, session_started_at DESC"*. The `SortFields` literals are at
   `:1751-1764`; `:1738-1746` is the `RowNumber == 1` reduction. Range is one line short at the head of
   the sort block.
6. **`corpus/services/hiring.md:188`** — cites `persona_write.go:69-71,143-167` for the PersonaSeeder's
   `validation_*` writes. The writes are at `:92-94`; `:66-68` is the comment enumerating the three
   tables and `:69-71` is the tail of an unrelated note about the *removed* `jobsimulation.sessions` step.
7. **`corpus/services/hiring.md:165`** (read-path table row 1) — `simulationScoreColumn.tsx:54` is
   `accessorKey,` inside the column factory, not a render of `row.score`. The co-cited `:95-97` **is**
   correct (`const { score, … } = info.row.original; const finalScore = formatNumberToString(score, 0)`).
8. **`corpus/architecture/shared_libraries.md:126-127`** — *"Vendor selection lives in each consumer's
   own `internal/ai/ai.go` wrapper"*. There is **no `app/internal/ai/ai.go`**
   (`find app -path "*/internal/ai/ai.go"` → empty; `app/internal` has `aireadiness`, `aiusage`,
   `cms/aivideo`, `jobsimulation/ai`). Since the merges the live wrappers are
   `app/internal/jobsimulation/ai/ai.go` (EU-Azure default + `flag_use_azure_us` at `:259-277`, the
   429→OpenAI fallback at `:127-137,167,297-300`, `AudioTranscriptions` at `:340-353`) and
   `app/internal/skillerai/ai.go:347`. The path is still literally true of the frozen
   `jobsimulation` husk repo (`jobsimulation/internal/ai/ai.go` exists), which is why I graded it MINOR
   rather than escalating — **and the substance of the correction is fully verified correct**. But a
   reader who `ls`es `app/internal/ai/` finds nothing.
9. **`corpus/architecture/shared_libraries.md:122`** — *"written to the `ai_usage` Postgres table"*. The
   table is **`ai_usages`** (plural) — `app/terraform/migrations/20250611092754.sql:2`:
   `CREATE TABLE "ai_usages" (… "cost" double precision NOT NULL …)`. The hardcoded model→price switch
   the same sentence claims is real (`internal/aiusage/ai_usage.go:34-54`), as is the `Event_AiUsage`
   Redis path (`main.go:1296`). Most directly actionable of the minors: a `/db-query` reader gets
   `relation "ai_usage" does not exist`.

### Verified-not-a-defect (checked because they *looked* wrong)

- **`askengine.md:63-64`** documents the admin routes as `POST /ask/admin/auto-rules/distill` etc.,
  while the platform's own handler comments say `/admin/ask/auto-rules/…` (`ask/admin.go:77,102,129,173`).
  **The corpus is right and the code comments are stale**: the mount is
  `askRouter := e.Group("/ask", …)` (`internal/web/backend/backend.go:171`) with
  `askRouter.POST("/admin/auto-rules/distill", …)` at `:185-188`.
- **`askengine.md:44`** says `WrapQuery` *"defends four documented bypass vectors"*, while
  `sandbox.go:350` says "three". The code comment then enumerates **four** bullets (`:352-362`) — the
  corpus counted the bullets, the platform prose was never updated. Corpus right.
- **`service_taxonomy.md:339`'s "3 → 1, not 2 → 1"** contradicts the upstream commit's own title
  (`915da06`: "fold cms subgraph into backend (supergraph 2→1)"). Measured: `915da06^` prod config
  carried three subgraphs and the commit deleted two schema files. Corpus right, commit title wrong.

---

## 4. Audited zeros

Read in full, top-to-bottom, and found clean at the blocker bar:

- **`corpus/services/storage.md` (175 lines) — zero findings, every claim measured.** Ports/env
  (`docker-compose.yml:198-210`, `PORT=8300`/`RPC_PORT=8301` mapped 1:1); the "binary default 8080/8081"
  aside (`storage/cmd/root.go:45-46`); `STORAGE_S3_PUBLIC_BUCKET` hardcoded to the production bucket at
  **`docker-compose.yml:210`** and the doc's own disambiguation that `:324` is inside the studio-desk
  block (it is — `env_file:` at `:324`, block spans `:311-342`); per-bucket FS fallbacks
  `/tmp/anthropos-storage/` (`internal/storage/storage.go:36`) and `/tmp/anthropos-public-storage/`
  (`:174`); presigned URLs return `""` with no bucket (`storage.go:121-124` — the doc's `:122` is the
  `if s.S3Bucket == ""` line); `viant/afs v1.30.0` (`go.mod:17`); Go 1.25.0; `NewClient`/`NewPublicClient`
  (`sdk/storage/client.go:10,27`); **no `*_test.go` anywhere** and `Dockerfile:18` is
  `RUN go test -v ./...` — so the "build gate that tests nothing" warning is exactly right; no
  `//go:generate` (control ran clean); gqlgen present only as `// indirect` (`go.mod:23`) = vestigial;
  no redis, no `DB_CONNECTION`. The merge banner's anchors resolve:
  `app/internal/jobsimulation/recording/recording.go:12` and
  `app/internal/jobsimulation/anticheat/anticheat.go:34` are both the storage-SDK import (note there are
  two `anticheat.go` files under `app/internal/jobsimulation/`; the doc means the
  `anticheat/` one, not `repository/`), and `app/main.go:983` is
  `storage.NewClient(os.Getenv("STORAGE_RPC_ADDR"), storagens.CMS)` verbatim.
- **`corpus/services/askengine.md` (121 lines) — zero findings.** `registry.go` holds **60** literal
  `TableDef`s ("~60" ✓) spanning `public` + `jobsimulation.*` (16) + `directus.*` (3), plus the
  programmatic `skiller.*` (`:551-561`) and `skillpath.skill_path_sessions` (`:571-574`) transition
  aliases that resolve to `public` — exactly as described. `bedrock.go:25-26`
  (`DefaultModelID = "eu.anthropic.claude-sonnet-4-6"`, `DefaultRegion = "eu-west-1"`), `Temperature: 0`
  (`:315,405`), `ASK_MODEL_ID` override (`:163`). `sandbox.go`: `ValidateSQL:163`, `BuildCTEs:227` with
  `StrategyGlobal/Direct/Indirect` (`:257-268`), `WrapQuery:363`, `$1`=orgID/`$2`=callerUserID (`:392,423`).
  `executor.go`: `MaxInlineRows = 200` (`:18`), `MaxCellLength = 400` (`:21`), `QueryTimeout = 10 *
  time.Second` (`:24`) applied as `SET LOCAL statement_timeout` (`:97`), `crossValidate` warnings for
  percent-range (`:215`), negatives (`:225`), headcount (`:37`) and duplicate-`user_id`/missing-GROUP-BY
  (`:253`). `main.go:149` `askEngineMaxConns = 6`, `:312` `COPILOT_DB_CONN`, `:335`
  `SetMaxOpenConns(askEngineMaxConns)`. `handler.go:34` `maxAgenticIterations = 15`, `:41`
  `loopTimeout = 10 * time.Minute`, detached via `context.WithoutCancel` (`:243-248`).
  `rules.md` is **146,219 bytes** ("~146 KB" ✓). All five `ask_*` tables are `CREATE TABLE` in the
  `public` migration set (`20260505133528.sql:2,19`, `20260506145258.sql:2`, `20260518125439.sql:6,23`,
  the last with `extensions.vector(1536)`). Test counts reproduce **exactly**: sandbox 49, executor 13,
  prompt 10, followups 14, and `TestBuildSystemPrompt_RulesBlockIsIdentical` exists
  (`prompt_test.go:176`). Route group + per-route mounts confirmed at `backend.go:171-188`.
- **`corpus/architecture/dependency_map.md` (103 lines) — zero findings.** The matrix, the husk
  annotations, the messenger four-address split, the studio-desk `depends_on`, the storage row (the
  iter-49 correction is right), the stream table (all six streams + the roadrunner strike-through), the
  six key flows, and the shared-library table all check out against compose + `app/main.go:1256-1311`.
  Spot-verified extras: pgvector really is built into the custom Postgres image
  (`platform/postgresql/Dockerfile`, `PGVECTOR_VERSION=v0.4.4`) and the vector columns live in the
  `extensions` schema (`20260518125439.sql:30`); `/api/webhook/directus` is mounted at
  `backend.go:324` with the `DIRECTUS_WEBHOOK_SECRET` handler built at `main.go:1080`;
  `app/internal/converter/gotenberg.go:31` posts to `/forms/libreoffice/convert`;
  `academy.graphqls` exists at the stated path.
- **`corpus/architecture/service_taxonomy.md` (440 lines) — zero blockers.** Tier-1 table, husk rows,
  profile table (`:412-425`), summary table (`:429-441`), the two banner blockquotes, the
  content-vs-runtime callout, all three Tier-2 sub-sections, the Tier-3 Clerk/Directus/Router sections.
  Profile membership independently reconstructed from the compose `profiles:` keys and matches every
  row of `:412-425`, including "sentinel is always on" (it has no `profiles:` key) and "six Go services
  + Gotenberg" under `graphql`.
- **`corpus/architecture/shared_libraries.md` (242 lines) — zero blockers**, minors #8/#9 above. The
  consumption model, the per-repo version-pin tables, the `proto` "removed, not re-hosted" argument, the
  authn-is-a-colony-subpackage correction, and the entire taxonomy-figures section are all measured-true.
- **`corpus/services/hiring.md` (398 lines) — zero blockers**, minors #1–#7 above. Notably the three
  facts the iter-23/iter-49 re-grounding banner claims to have changed are each independently confirmed,
  including the subtle one: `20260722104506.sql:79` drops **`public.sessions`** (bare `DROP TABLE
  "sessions"` under `search_path=public`), *not* `jobsimulation.sessions`, and no `app` migration
  touches the `jobsimulation` schema at all.

---

## 5. Unverified

Claims I could not check, with the reason. None are reported as passed and none as blockers.

1. **All GitHub repo-archival dates** — `service_taxonomy.md:3` (`graphql-wundergraph` archived
   2026-07-30), `:95` (skiller 2026-07-01), `:96` (skillpath 2026-07-31), `:97` (jobsimulation
   2026-07-31), `:98` (cms "frozen, **not** archived"), `:64`/`:66`. `gh` is unavailable and archival
   state is not in a clone. The *local* consequences of each (absent from `repos.yml`/compose, or
   present as a husk) were verified independently and are correct.
2. **`colony` internals** — the whole package table at `shared_libraries.md:48-54`, the Sentry-only-in-
   production note, and the "GraphQL rate limiter only logs, reject path commented out" claim
   (`:56-59`). `colony` is not cloned. Only its *version pins* and *import graph* were verifiable
   (and are correct).
3. **`proto` internals** — "12 Connect-RPC services are defined" and the enumerated list
   (`shared_libraries.md:77-83`), the buf-v1 single-module layout and `go/simulator/*`-without-`.proto`
   note (`:91-94`). `proto` is not cloned. Two halves *were* checkable and hold: `app/main.go`
   serves exactly six of them, with neither `SkillPathSessionService` nor a RoadRunner service among
   them; and `storage/internal/migration/migration.go:13` does import `go/simulator/storage/v1` as
   `legacyStorage`.
4. **`ai` library internals** — the constructor table (`:111-116`), the Anthropic-panics /
   Mistral-OCR-only / no-native-JSON-mode / 10-attempt-retry gotchas (`:130-135`). Not cloned. The
   *consumer-side* half of the same section (cost in `aiusage`, vendor selection in the consumer, the
   Bedrock askengine path) is verified true.
5. **`taxonomy` library internals** — "131-line `node.go`", the `NodeID` format regex, the five
   constructors and the org-scoping `v1.2.0` change, and the three worked examples
   (`shared_libraries.md:215-229`). Not cloned. The pin `v1.2.0` and the full import graph are verified.
6. **`authn` standalone repo** — tag `v1.7.0` and the "missing `return` after the websocket-skip"
   defect (`:143`, `:164-166`). Not cloned. The verifiable half holds: **no** checked-out service
   requires `github.com/anthropos-work/authn`, and sentinel imports neither authn nor `colony/authn`.
7. **Archived-`chronos` colony pin `v0.30.1`** (`shared_libraries.md:41`) — `stack-demo/chronos` does
   not exist (only `app, ant-academy, cms, graphql-wundergraph, jobsimulation, messenger, next-web-app,
   platform, roadrunner, rosetta-extensions, sentinel, storage, studio-desk` are cloned).
8. **Historical measurements and live-run outcomes** — hiring.md's `billion` render results
   (`:343-348`, `:373-382`: the 5/5 per-position tally, "recruiter p95 click→ACCESS 1.27 s", "≥3 cold
   runs, 4/4 flake", "0 prod-eject"), the drawer's first-sim-only detection behaviour, and
   askengine.md's ship dates (`v1.267.0` 2026-05-07 / `v1.340.0` 2026-07-17). These require a running
   stack or a release ledger, not a clone. The *artifacts* they rest on (the render spec, the gate
   floor, the seeders, the enum/flag surfaces) all exist as documented — and the per-position tally
   independently corroborates minor #2 rather than #2's claim.
9. **The rolled-back experiment** at hiring.md:217 (`UPDATE … SET
   completion_status='SIMULATION_COMPLETION_STATUS_PASSED'` is accepted) — not re-run against a live DB.
   Its premise is confirmed statically: the column is a plain `varchar` with no CHECK
   (`20260722104506.sql:12`) and neither Ent's `assignValues` nor the gqlgen marshal can reject a value.
10. **Prod-side counts named as historical** — "the prod '443' [job_position rows] was never captured"
    (hiring.md:140). The capture side is verified (no `job_position` table in the Directus snapshot
    manifest); the prod figure itself is not reachable read-only from here.
