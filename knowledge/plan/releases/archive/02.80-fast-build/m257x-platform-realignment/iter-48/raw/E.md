# Seat E — iter-48

Repo `rosetta` @ `m257x/platform-realignment` HEAD `cabc3b1`.
Ground truth: `stack-demo/app` @ `5ba17044` (v1.363.2), `stack-demo/platform` @ `2adcf71`,
`stack-demo/next-web-app` @ `bb3313bc0` (v2.133.0), `stack-demo/ant-academy` @ `9c3843cd` (v2.34.2),
`stack-demo/graphql-wundergraph` @ `60c229f`, plus `stack-demo/{cms,jobsimulation,sentinel,storage,messenger,roadrunner,studio-desk}`
and `.agentspace/rosetta-extensions`.

## Coverage (file, wc -l, lines read)

| File | `wc -l` | Physical lines | Lines read |
|:--|--:|--:|:--|
| `corpus/architecture/service_taxonomy.md` | 440 | 441 (last line has no trailing newline) | **all 441** |
| `corpus/services/hiring.md` | 378 | 378 | **all 378** |
| `corpus/architecture/shared_libraries.md` | 242 | 243 (last line has no trailing newline) | **all 243** |
| `corpus/services/storage.md` | 175 | 175 | **all 175** |
| `corpus/services/askengine.md` | 121 | 121 | **all 121** |
| `corpus/architecture/dependency_map.md` | 103 | 104 (last line has no trailing newline) | **all 104** |

Every file was read top-to-bottom in one `Read` call, no `offset`/`limit`, no sampling.

**Not verifiable in this environment** (declared, not silently passed):
- `colony` internals (`shared_libraries.md:44-59`) — the repo is not cloned and is absent from the local
  module cache (`ls $GOMODCACHE/github.com/anthropos-work/` → only `ai@v1.40.1`). The colony *pins* and
  *importers* were verified from each service's `go.mod`; the per-package claims were not.
- `proto`'s 12-service list (`shared_libraries.md:77-83`) — same reason. The six handlers `app` **registers**
  were verified against `app/main.go`.
- GitHub archive states/dates (`service_taxonomy.md:64,95-98`) — no `gh` on this box, no network check run.

**Search hygiene.** Every grep in this pass had its stderr read (one `zsh` glob rejection on
`--include=*.go` was caught and re-run quoted — it would otherwise have read as "no matches"), and each
negative result was paired with a positive control in the same invocation. The controls used:
`grep -c "PORT="` over compose = 16; `JUDGE0_BASE_URL` over `app` = 1 hit (against `ROADRUNNER_RPC_ADDR` = 0);
`anthropos-work/colony` over `storage` = 4 hits (against `DB_CONNECTION|redis` = 0); `CREATE SCHEMA` over
`app/terraform/migrations/` = 1 hit (against `jobsimulation\.` = 0); `find apps -type d` = 539 dirs and
route groups `(protected)`/`(public)` present (against `(.)`/`(..)` = 0); `vite` in `studio-desk/package.json`
(against `react|vue|angular` = 0).

---

## Blockers

| # | Site | The false claim (verbatim) | What is TRUE | Citation |
|:--|:--|:--|:--|:--|
| 1 | `corpus/services/hiring.md:28` (repeated `:148`, `:276`) | "`jobsimulation.sessions` was dropped (`20260722104506.sql:79`) and replaced by `public.job_simulation_sessions` (`:2`)" | Line 79 is `DROP TABLE "sessions";` — and `app`'s migrations run with `search_path=public`, so it dropped **`public.sessions`**, the app-side jobsim table created ~20 minutes earlier in the *same* migration set and renamed to `job_simulation_sessions`. **No `app` migration touches the `jobsimulation` schema at all**, so nothing in `app` could have dropped `jobsimulation.sessions`; the legacy schema survives, frozen, until platform M710. The doc's own supporting clause (`atlas.hcl:8` pins `search_path=public`) is what refutes it. | `app/terraform/migrations/20260722104506.sql:79` (`DROP TABLE "sessions";`) + `:2` (`CREATE TABLE "job_simulation_sessions"`); `app/terraform/migrations/20260722081626_jobsim_data_model.sql:2` (`CREATE TABLE "sessions"`, same set, unqualified); `app/atlas.hcl:8` `…&search_path=public`; `grep -rn 'jobsimulation\.' app/terraform/migrations/` → **0 hits** (positive control `CREATE SCHEMA` → 1 hit, `20230817154747_supabase_baseline.sql:2` `CREATE SCHEMA IF NOT EXISTS auth;`). Platform source that the schema still exists: `app/internal/askengine/registry.go:192` — *"`jobsimulation` stays in rewriteSchemas until M710 drops the (now-frozen, no-longer-synced) jobsimulation schema"*. Corpus twins that contradict the doc: `service_taxonomy.md:52` ("the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks"), `dependency_map.md:78` ("the legacy `jobsimulation` schema is non-authoritative"). The rext seeder records the correct reading in-code: `.agentspace/rosetta-extensions/stack-seeding/seeders/persona_write.go:63-65` — *"`sessions` was created then DROPped by the very next migration as the rename completed; `public.sessions` does not exist"*. |
| 2 | `corpus/services/hiring.md:189-196` | "**Minimal write-set per (candidate × sim):** 1. **`public.job_simulation_sessions`** … Non-null `status`, `started_at`, `ended_at`, `owner_id`, `sim_id`, `sim_type`, plus `score` (0–100), `completion_status` …, `organization_id`, `tenant_id` (NULL or `=org`), `validation_version`." | The list omits **`token`** — `character varying **NOT NULL**` with **no default** and a UNIQUE index. An INSERT of exactly the enumerated column set fails on a NOT-NULL violation, so the stated "minimal write-set" cannot produce a row. It is the **only** required-and-undefaulted column missing (every other omission — `id`, `created_at`, `updated_at`, `interactions_progress`, `language`, `result_status`, `chime_status` — carries a DEFAULT or is NULLable), which is what makes the omission an error rather than deliberate scoping. The word "token" appears **nowhere** in this doc (`grep -n token corpus/services/hiring.md` → exit 1; positive control `grep -c score` → 45). The real seeder writes it. | DDL: `app/terraform/migrations/20260722104506.sql:13` `"token" character varying NOT NULL,` + `:29` `CREATE UNIQUE INDEX "job_simulation_sessions_token_key"`. Ent: `app/internal/data/ent/schema/job_simulation_session.go:41` — `field.String("token").Match(regexp.MustCompile("^[a-z0-9]+$")).MinLen(5).MaxLen(10).Unique().Immutable()` (no `.Default`, no `.Optional`). Seeder ground truth: `.agentspace/rosetta-extensions/stack-seeding/seeders/persona_write.go:152-158` — `sessionCols()` returns 18 columns **including `"token"`** (`:155`), consumed at `:91` `{"public", "job_simulation_sessions", sessionCols(), a.sessions}` — the exact two anchors `hiring.md:232-234` cites. |
| 3 | `corpus/services/hiring.md:20-22` (repeated `:147-148`) | "`app` migration `20260729133514.sql:58-62` — *'5. Drop the mirrors.'* — **back-fills it into the canonical entity** and then `DROP TABLE \"local_jobsimulation_sessions\"`" / "That migration **back-filled the mirror into the canonical entity** and **dropped it** (`:58-62`)" | There is **no back-fill**. `20260729133514.sql` (64 lines, read in full) does exactly four things: drops the mirror-pointing FK constraints (`:7-11`), **re-points the *referencing* link ids** in `organization_assignment_sessions` / `personal_assignment_sessions` from mirror ids to canonical ids (`:15-23`), deletes/nulls orphaned links (`:30-44`), adds the canonical FKs (`:52-56`), then drops the mirror tables (`:62-63`). Not one statement writes into `job_simulation_sessions`. Across the **entire** migration set there is no score back-fill and no session data migration: `grep -rln 'SET "score"\|SET score' app/terraform/migrations/` → **0 hits**, and `grep -rn "INSERT INTO" app/terraform/migrations/` returns only trigger bodies and the 2023 `world_languages` seed — nothing session-shaped (positive control: `grep -c "ALTER TABLE" 20260722104506.sql` = 16). The mirror's data was **discarded**, not migrated; the canonical rows were already carrying `score` from the jobsim-in-app port. | `app/terraform/migrations/20260729133514.sql:1-64` read in full (the header at `:1-4` states the intent verbatim: *"…re-point from the denormalized mirror rows to the canonical … then the mirror tables are dropped"* — re-point, not back-fill); `app/terraform/migrations/20260722104506.sql:17` `"score" real NOT NULL DEFAULT 0` (the canonical column, created at table-creation time). |
| 4 | `corpus/architecture/dependency_map.md:19` | "\| **Storage** \| - \| Postgres, Redis, **S3** \|" | Storage uses **neither Postgres nor Redis**. Its Go tree contains zero database or Redis code, its `go.mod` requires no redis client and pulls `pgx` only transitively via colony, and the compose service is given **no** `DB_CONNECTION` and **no** `REDIS_ADDR` — so it could not reach either even if it wanted to. (The compose `depends_on: postgresql, redis` at `docker-compose.yml:213-217` is a startup-order artifact, not a dependency.) This contradicts the twin doc in the same audit set: `storage.md:14` *"Storage is stateless and **owns no database**: all state lives in S3"* and `storage.md:21` *"**Database**: none — all state lives in S3 (or local filesystem in dev)"*. | `grep -rn "DB_CONNECTION\|NewDBPool\|NewDBStdConn\|redis\|Redis" --include="*.go" stack-demo/storage` → **0 hits, exit 0** (positive control in the same pass: `grep -rn "anthropos-work/colony" --include="*.go"` → 4 hits at `main.go:7`, `cmd/root.go:11,12`, `sdk/storage/client.go:5`); `stack-demo/storage/go.mod` — no redis require, `:53 github.com/jackc/pgx/v5 v5.9.2 // indirect`; `platform/docker-compose.yml:189-218` — the storage service's full env is `AWS_DEFAULT_REGION`, `AWS_REGION`, `ENVIRONMENT`, `PORT`, `RPC_PORT`, `SERVICE_NAME`, `STORAGE_S3_PUBLIC_BUCKET` only. |

---

## Minors

1. **`service_taxonomy.md:150-153` — the Studio-Desk `Technology` table row is wrapped across four physical
   lines, so it is not a table row.** GFM requires one row per line; line 150 ends mid-sentence with no
   closing `|`, and lines 151-153 do not start with `|`, which terminates the table. The four rows that
   follow (`:154-157` Port / Purpose / Authentication / Location) then have no header+delimiter above them
   and render as literal pipe-text. The *content* is TRUE and re-verified this pass — `studio-desk/package.json`
   has 0 react/vue/angular entries (positive control: `vite` at `:12,15,32`) and the repo has 0 `.tsx`/`.jsx`
   files outside `node_modules`; `studio-desk.md:20` agrees ("vanilla TS frontend, no framework"). Rendering
   defect only. *(This is the item flagged for re-check from the previous pass — it is still present, and it
   is correctly a MINOR: nothing asserted is false.)*
2. **`hiring.md:159` — anchor drift.** Row 5 of the read-path table cites `intelligence.go:1728-1735` for
   "best-attempt: `row_number() ORDER BY score DESC` per candidate". Lines 1728-1735 contain the
   `onlyAssignments` id-collection and the **call** at `:1733`; the actual window function is at
   `app/internal/organization/intelligence.go:2158-2169` —
   `ROW_NUMBER() OVER (PARTITION BY sim_id, owner_id ORDER BY score DESC NULLS LAST, …)` inside
   `usersBestOrFirstJobSimulationSession` (declared `:2124`). The claim is true; the anchor points at the
   caller, not the mechanism. (The companion range `:1738-1751` at `hiring.md:163` is correct.)
3. **`hiring.md:178` — anchor drift on the PersonaSeeder write cite.** "`persona_write.go:69-71,143-167`"
   for the `validation_*` writes: `:69-71` is now a prose comment block about the *removed*
   `jobsimulation.sessions` step, and `:143-167` straddles the end of `flush()` and the start of the column
   builders. The three `validation_*` write steps are at `.agentspace/rosetta-extensions/stack-seeding/seeders/persona_write.go:92-94`
   and their column builders at `:161-186`. Claim true, anchors stale.
4. **`askengine.md:107-109` — incomplete enumeration.** "the handler package `internal/web/backend/ask/`
   carries its own tests: sanitize, stream_registry, examples, lessons, distill, feedback, store, embed"
   lists 8; the package ships **9** `_test.go` files — `coursebuilder_tool_test.go` is omitted, which is
   odd given the same doc's `:20` and `:53-54` make author-mode/`course_builder` a headline. (Everything
   else in this file verified exact: `DefaultModelID = "eu.anthropic.claude-sonnet-4-6"` / `DefaultRegion = "eu-west-1"`
   `bedrock.go:25-26`, `Temperature: param.NewOpt(float64(0))` `:315`, `stripAnthropicAuthMiddleware` `:252`,
   `maxAgenticIterations = 15` / `loopTimeout = 10 * time.Minute` `handler.go:34,41`, `askEngineMaxConns = 6`
   `main.go:149`, `QueryTimeout = 10 * time.Second` + `SET LOCAL statement_timeout` `executor.go:24,97`,
   `MaxInlineRows = 200` / `MaxCellLength = 400` `:18,21`, `rules.md` = **146 219 bytes**, 60 literal `TableDef`s
   + programmatic `skiller.*`/`skillpath.*` aliases at `registry.go:560,574`, all 12 `/ask` routes at
   `internal/web/backend/backend.go:171-188`, and the test counts **49 / 13 / 10 / 14** exactly. Note also
   that "`WrapQuery` (defends **four** documented bypass vectors)" is *more* accurate than the platform
   source's own prose: `sandbox.go:350` says "three" while `:352-362` lists four bullets — the doc counted
   the bullets, which is right. Not a finding.)
5. **`dependency_map.md:48` and `:50` — the Shared-Libraries summary omits importers the corpus insists on
   elsewhere.** `ai` is given as "app"; `taxonomy` as "direct — app, messenger; indirect — storage, sentinel".
   But `shared_libraries.md:105` and `:181` go out of their way to state that the still-running `cms` and
   `jobsimulation` husk repos require **both** libraries **directly** in their own `go.mod` — verified:
   `cms/go.mod:9` `ai v1.40.2`, `:13` `taxonomy v1.2.0`; `jobsimulation/go.mod:11` `ai v1.40.2`, `:15`
   `taxonomy v1.2.0`, neither marked `// indirect`. `dependency_map.md:15-16` itself insists both husks still
   start in the default `graphql` profile. Summary-table incompleteness, not a false statement.
6. **`shared_libraries.md:145` — inconsistent treatment of the husks in the `authn` row.** "Imported by |
   via colony: app (the former cms / jobsimulation / skillpath usage is all folded in)". The *live path* claim
   is true, but `cms` (6 files) and `jobsimulation` (8 files) still import `github.com/anthropos-work/colony/authn`
   in their own trees, and the `ai` (`:105`) and `taxonomy` (`:181`) rows of this same doc explicitly name
   those husks as current importers. Framing only. (The doc's headline authn correction at `:162` is exactly
   right and was re-verified: `grep -rn "anthropos-work/authn" --include="*.go"` across all seven service
   repos → **0 hits**; sentinel/storage/messenger/roadrunner import `colony/authn` 0 times, app 129 times.)
7. **`shared_libraries.md:79` — the compose anchor is on the wrong service block.** "`ROADRUNNER_RPC_ADDR`
   (`docker-compose.yml:118`) is read by no Go code in `app`" — the *claim* is true
   (`grep -rn "ROADRUNNER_RPC_ADDR" --include="*.go" stack-demo/app` → 0, exit 1; positive control
   `JUDGE0_BASE_URL` → 1 hit), but `docker-compose.yml:118` sits inside the **jobsimulation husk** service
   (lines 83-143), not `backend`'s env — `backend` (lines 28-81) never declares the variable at all. The line
   number coincidentally equals `internal/jobsimwiring/wiring.go:118` cited in the same sentence, which *is*
   exact. Worth re-pointing so the sentence doesn't read as "app is handed a dead env var".

---

### Verified-clean (recorded so a later pass doesn't re-spend on them)

Everything below was opened and matched exactly; no finding.

- **`storage.md`** — clean end to end. `recording.go:12`, `anticheat/anticheat.go:34`, `main.go:983`
  (`storage.NewClient(…, storagens.CMS)`), Go 1.25.0, ports/`PORT`/`RPC_PORT` + binary defaults 8080/8081
  (`storage/CLAUDE.md:12,130-131`, `cmd/root.go:45-46`), `docker-compose.yml:210` public bucket (and `:324`
  really is the studio-desk block), `DefaultTmpPath`/`DefaultTmpPublicPath` `storage.go:36,174`,
  `GetPresignedUrl` empty-string fallback `storage.go:122-124` and the 15-minute default `rpcsrv.go:99-101`,
  the exact 5 RPC methods, the key-directory tree, 0 `*_test.go`, `Dockerfile:18 RUN go test -v ./...`,
  0 `//go:generate`, gqlgen vestigial.
- **`hiring.md`** apart from the three blockers — every remaining anchor is exact:
  `resolver_queries.go:1034/1035/1053/1080/1085/1089`, `intelligence.go:1700/1738-1751/1820/1844/1846/885-886`,
  `enum/jobsimulation.go:29-35` + `Values()` `:37-43`, `ent/jobsimulationsession.go:181-186` (unconditional
  cast), `graph.go:129546-129554` and the proto twin `:129392-129400` (bare `MarshalString` passthrough),
  `jobsimulations.graphqls:14` and `:128` (five lowercase members each), `job_simulation_session.go:39,45`,
  `anticheat_result.go:24`, `20250416091037.sql:5`, `20260722081626_jobsim_data_model.sql:336/355/376`,
  `20260722104506.sql:12,53`, `resolver_cms_queries.go:95,99-103,210,258,295`, `organization/manager.go:448,485`,
  `siminvitationlink.go:62`, `organizations.is_hiring` bool NOT NULL default false, no CHECK on
  `completion_status`, no `local_*` ent schema remains. The three grep-based empirics are **exactly** right:
  `grep -in hiring` over `resolver_queries.go` → only sim-TYPE lines (`:811,843,991-992`), over
  `intelligence.go` → nothing; positive controls `OrgFeatureInsights` ×**8** and `JobSimulationSession` ×**44**.
  Frontend: `useGetClerkOrganization.tsx:20-21` verbatim, `apps/web/.../UserStatusContext.tsx:144-149,168-172`,
  `apps/hiring/.../UserStatusContext.tsx:125,144-145`, `useNavbarSections.tsx:459-466,460`, `template.tsx:90`,
  `FreeTrialContainer.tsx:29`, `simulationScoreColumn.tsx:33,54,95-97`, `insights.ts:31-82`,
  `apps/hiring/.../InsightsByMembersContainer.tsx:359` (`<Drawer`), the `@tabs/ai-simulations/[simId]` +
  `[userId]` routes, and **0** intercepting-route dirs anywhere in `apps/`.
- **`service_taxonomy.md`** apart from minor 1 — the whole compose/profile/port surface at platform `2adcf71`
  (`repos.yml:10-13,14-19`, sentinel `:18 search_path=sentinel`, backend `depends_on` `:70-80`, jobsimulation
  `:83`, cms `:144`, storage `:189`, roadrunner `:281`, studio-desk `:311`/`:318`/`:334`/`:337-341`/`:342`,
  next-web-app `:352`/`:361`, gotenberg `:371`; six Go services + Gotenberg in `graphql`; DIRECTUS_* on the
  `cms` service **only**). The router history is exact: `b56d731`+`360efd4` merged as `2adcf71`;
  `locals.tf:8 port = 8080`, `main.tf:20 service_desired_count = 1`, `main.tf:48-49`, `config.prod.yaml:5`;
  and the **measured subgraph ladder reproduces exactly** — `749dc86^` = 5 (backend, skiller, jobsimulation,
  cms, skillpath) → `749dc86` = 4 → `7c17e63` = 3 → `915da06` = 1, with `cms.graphqls` **and**
  `jobsimulation.graphqls` deleted in that single commit, i.e. **3→1**, contra the commit message's own
  "2→1". The Directus history retraction-of-the-retraction is exact: `a2a3ee6^:docker-compose.yml:384`
  `image: directus/directus:10.10.1`, `:386 8055:8055`, `:409 ADMIN_PASSWORD=password`, and
  `admin@example.com` genuinely has **no** hit in platform history (`git log --all -S`). `gen.py` registers
  **exactly nine** `add_argument`s and uses `parse_known_args` (`app/studio/gen.py:19,484-492`) — no
  `--template`; `studio/CLAUDE.md:12-14` is the quoted entry point; `additional_repo: "anthropos-studio-room:studio"`
  at `.github/workflows/build-production.yml:29` with app `v1.360.1` in CHANGELOG. Ant Academy: Next `^16.2.6`
  + React `^19.2.5`, `--port 3077`, Expo `~54.0.33`, `server.js:14,18` (throws), `serverTenant.js:115-145`,
  `backendContent.js:36,102-103`, `beacon/route.js:36,41-55`, `layout.jsx:132` + `academy-manifest.json`
  `"display": "standalone"`, `next-scaffold.test.js:106,111`, `react-compiler-config.test.js:41`, zero
  serwist/workbox deps, `RegisterServiceWorker.jsx` is a kill-switch, `gpt-5.2`, and
  `app/internal/web/backend/graphql/graph/schemas/academy.graphqls` exists.
- **`shared_libraries.md`** apart from minors 5-7 — **every version pin re-measured from `go.mod` and correct**:
  colony `app`/`messenger` v0.35.2, `cms`/`jobsimulation` v0.35.1, `sentinel`/`storage`/`roadrunner` v0.34.3;
  proto v1.210.0 / v1.207.0 / v1.205.0 / v1.200.0 / v1.196.0; ai v1.40.2 ×3; taxonomy v1.2.0 with
  sentinel/storage `// indirect` and roadrunner absent. `roadrunner/main.go:7` + `:17 colony.NewVersionConfig`
  exact. The six Connect handlers at `app/main.go:1178,1179,1187,1195,1204,1218` are exact, and neither
  `SkillPathSessionService` nor a RoadRunner service is registered. `skillpaths.go:27-31` quoted verbatim.
  `jobsimwiring/wiring.go:118` exact. `GH_ACCESS_TOKEN`/`GOPRIVATE`/`insteadOf` in the Dockerfiles; no
  `go.work` committed in any service (and it is `.gitignore`d at `:14`). The taxonomy figures cross-check
  against `service_taxonomy.md:112` with no drift.
- **`dependency_map.md`** apart from blocker 4 and minor 5 — `app/main.go:1196-1202` is exactly the M809
  re-point comment; messenger's four addresses at `:255,256,258,265`; studio-desk `depends_on` `:337-341`;
  pgvector really is in the `extensions` schema (`20260518125439.sql:30` etc.); the Directus webhook lands
  on `POST /api/webhook/directus` authenticated by `DIRECTUS_WEBHOOK_SECRET` (`main.go:1080`,
  `backend.go:137-138,324`); every `*_STREAM` name resolves in compose and no `ROADRUNNER_STREAM` exists;
  the cms stream really is one subscriber carrying both jobsim and cms/Studio handlers
  (`main.go:1279,1287-1291`); messenger really does subscribe to `"jobsimulation"`
  (`messenger/internal/flow/flow.go:105`); customerio-sync is built from the git URL with only
  `DB_CONNECTION_BACKEND`; studio-desk's `AI_PROVIDER_CHAIN` supports anthropic
  (`src/services/ai/providers/anthropic.ts`, `.env.example:55`); `GOTENBERG_URL=http://gotenberg:3200`
  at `docker-compose.yml:51`.
