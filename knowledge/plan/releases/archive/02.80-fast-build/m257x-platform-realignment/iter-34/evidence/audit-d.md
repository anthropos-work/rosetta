# iter-34 confirming pass — audit D

Group D: 4 swept service docs + 4 never-edited ones. Graded against platform origin `2adcf71`
(`stack-demo/platform`), `app @ 5ba17044` (v1.363.2), `next-web-app @ bb3313bc0`, and the frozen
`cms` / `jobsimulation` / `roadrunner` / `sentinel` / `graphql-wundergraph` clones.

## Positive control

| File | `wc -l` | Last line actually read | Status |
|---|---|---|---|
| `corpus/services/hiring.md` | 319 | 319 | READ IN FULL |
| `corpus/services/backend.md` | 253 | 253 | READ IN FULL |
| `corpus/services/cms.md` | 237 | 237 | READ IN FULL |
| `corpus/services/jobsimulation.md` | 226 | 226 | READ IN FULL |
| `corpus/services/graphql-wundergraph.md` | 196 | 196 | READ IN FULL |
| `corpus/services/roadrunner.md` | 164 | 164 | READ IN FULL |
| `corpus/services/sentinel.md` | 159 | 159 | READ IN FULL |
| `corpus/services/ai-labs.md` | 156 | 156 | READ IN FULL |

Total 1710 lines, no file sampled or skimmed.

---

## BLOCKERS

### B1 — `corpus/services/hiring.md:156` · a column that does not exist, inside the seeder write-set

> "`tenant_id` (NULL or `=org`), `validation_version`, **`anticheat_summary` (optional)**."

**False at HEAD.** `public.job_simulation_sessions` has no `anticheat_summary` column. Its full column
set is fixed at `app/terraform/migrations/20260722104506.sql:2-22` and
`app/internal/data/ent/schema/job_simulation_session.go:30-57` — no anticheat field. The column
existed **only on the dropped mirror**: `app/terraform/migrations/20250416091037.sql:3-5`
(`ALTER TABLE "local_jobsimulation_sessions" ADD COLUMN "anticheat_summary"`), and that table is
`DROP`ped at `20260729133514.sql:62`. At HEAD the anticheat summary is read from a **separate
entity** — `IntelligenceManager.anticheatSummariesBySession` (`app/internal/organization/intelligence.go:1341-1352`,
called at `:1796`) queries `m.ent.AnticheatResult` (`anticheat_results`, 1:1 session). The real
seeder agrees: `rosetta-extensions/stack-seeding/seeders/persona_write.go:152` `sessionCols()` lists
18 columns and `anticheat_summary` is not among them.

This is the "**Minimal write-set per (candidate × sim)**" block — the section M223/M224 and any future
hiring seeder is built from. A COPY/INSERT that names the column fails outright.
**Grade: BLOCKER.** Same root defect one paragraph earlier at **`:144`** — *"`anticheat_summary` on
**the mirror row** is a decorative icon only"* — present tense about a table this file's own banner
says was dropped.
*Correction:* drop `anticheat_summary` from the write-set; add a line saying the anticheat icon is
fed by a separate `public.anticheat_results` row (optional, not part of the scoreboard write-set).

### B2 — `corpus/services/hiring.md:259-260` · the demo hiring container "wired to … Cosmo"

> "So the demo builds `apps/hiring` from the **untouched clone** as a second offset-port UI container
> (same recipe as `apps/web` + `studio-desk`), **wired to the same fake FAPI + Cosmo + Postgres**"

**False at HEAD**, and it contradicts its own paragraph. Three lines above (`:256-258`) the doc
correctly states the endpoint "since platform `2adcf71` (2026-07-31) is **`backend`'s own
`:8082/graphql/query`**, the Cosmo/WunderGraph router having been deleted from compose." There is no
`graphql` service in `docker-compose.yml` (services are `sentinel` 5, `backend` 28, `jobsimulation` 83,
`cms` 144, `storage` 189, `customerio-sync` 220, `messenger` 240, `roadrunner` 281, `studio-desk` 311,
`next-web-app` 344, `gotenberg` 371) and no `graphql-wundergraph` entry in `repos.yml` (9 repos).
The wiring list is the operational instruction a demo operator acts on. **Grade: BLOCKER** — this is
exactly the "correction landed, surrounding prose left standing" shape iter-33 measured.
*Correction:* "wired to the same fake FAPI + `backend`'s `:8082/graphql/query` + Postgres."

### B3 — `corpus/services/backend.md:35` · an RPC service that is not on the mux

> "the mux carries `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`,
> **`SkillPathSessionService`**, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`"

**False at HEAD.** `grep -r SkillPathSessionServiceHandler` over the whole `app` repo returns **zero
hits**; there is no `skillpath*v1connect` import. `app/main.go:102-107` imports exactly
`organizationsv1connect`, `usersv1connect`, `cmsv1connect`, `jobsimulationv1connect`, `labv1connect`,
`skillerv1connect`, and the handlers registered are `main.go:1178` (Users), `:1179` (Organizations),
`:1187` (Skiller), `:1195` (JobSimulation), `:1204` (CMS, conditional), `:1218-1219` (LabSession).
The skill-path engine is reached through GraphQL (`graph/schemas/skillpath_sessions.graphqls`) and the
`SKILLPATH_STREAM` subscriber (`main.go:1265`), not Connect-RPC. A consumer written against this line
would dial a service that does not exist. **Grade: BLOCKER.**
*Correction:* delete `SkillPathSessionService` from the list; note the skill-path session surface is
GraphQL + Redis-stream, with no RPC handler.

### B4 — `corpus/services/backend.md:111` · a code-map entry for a package and table that were both removed

> "`aiacademy/  Periodic AI Academy catalog sync (fetches catalog.json, populates aiacademy_courses
> for Talk to Data)`"

**False at HEAD on both halves.** `app/internal/aiacademy/` does not exist (`ls internal/` — the
package is `internal/academy/`), and the table is gone. `app/internal/academy/academy.go:6-9` is
explicit: *"The catalog tables this domain owns (academy_series / academy_skill_paths /
academy_chapters / academy_chapter_bodies) are also the feed the askengine 'Talk to Data' sandbox
queries — **the legacy internal/aiacademy sync and its aiacademy_courses read-model were removed**
once this domain took ownership."* Anyone tracing the Talk-to-Data academy feed from this line looks
for a package and a table that no longer exist, and misses the real owner (documented in
`corpus/services/academy-backend.md`). **Grade: BLOCKER.**
*Correction:* `academy/  Server-owned academy domain — academy_series / academy_skill_paths /
academy_chapters / academy_chapter_bodies (the Talk-to-Data feed; the old aiacademy sync +
aiacademy_courses were removed)`.

---

## Minors

### M1 — `hiring.md:127` · read-path table row 6 contradicts row 7
> "| 6 | `intelligence.go:1801` | `Score` ← `ls.Score` (**the mirror's** score column) |"

`:1801` is inside the `anticheatSummariesBySession` block (`intelligence.go:1796-1800`). The score is
computed at `:1820` (`score := RoundFloat(float64(ls.Score), 0)`) and assigned at `:1846`
(`Score: &score`) — which the very next row of the same table states correctly. And there is no
mirror. Residual attribution + wrong anchor in one cell. *Fix:* fold row 6 into row 7.

### M2 — `hiring.md:278` · residual present-tense "mirror"
> "no training/assessment leakage into **the mirror** the list groups by"

The list groups by `public.job_simulation_sessions` (`intelligence.go:1700`). *Fix:* "…into the
session table the list groups by."

### M3 — `hiring.md:155` · a `completion_status` value that does not exist
> "`completion_status` (values `passed`/`failed`/`pending`/`SIMULATION…`)"

The enum is `pending | passed | failed | discarded | timedout`
(`app/internal/data/ent/enum/jobsimulation.go:30-34`). `SIMULATION_*` values belong to `sim_type`, not
`completion_status`. Not a hard failure (the column is `character varying`), but a seeder that writes
a `SIMULATION_…` string produces silently-wrong rows.

### M4 — `hiring.md:154-156` · the "minimal write-set" omits a NOT NULL column
`token` is `NOT NULL` with **no default** and `Unique`
(`20260722104506.sql:13`; ent `job_simulation_session.go:38`) and is written by the real seeder
(`persona_write.go:152` `sessionCols()`). An INSERT built from this list fails on `token`.

### M5 — `hiring.md` stale anchors (claims all still true; line numbers drifted)
- `:122` / `:137` `resolver_queries.go:1088,1134` and `:1089` → at HEAD the resolver is
  `internal/web/backend/graphql/graph/resolver_queries.go:1034`, the `OrgFeatureInsights` gate is
  `:1035`, and the `IntelligenceManager` delegation is `:1080`. Lines 1085-1090 belong to a
  *different* resolver — an anchored re-check would "verify" against the wrong function.
- `:189` `useNavbarSections.tsx:300-307` → the `isHiringOrg ? tNavbar('results') : …` line is
  `packages/ui/src/NavBar/useNavbarSections.tsx:460`.
- `:120` `simulationScoreColumn.tsx:54` → `accessorKey = 'score'` is `:33`; the `95-97` half is exact.
(Verified-exact in this file: `useGetClerkOrganization.tsx:20-21`, `FreeTrialContainer.tsx:29`,
`insights.ts:32`, `InsightsByMembersContainer.tsx:359` (apps/hiring), `intelligence.go:1700`,
`:885-886`, `20260722081626_jobsim_data_model.sql:336/355/376`, `20260729133514.sql:58-62`,
`20260722104506.sql:2,79`, `atlas.hcl:8`, `persona_write.go:152`, and the zero-intercepting-routes
claim at `:297`.)

### M6 — `backend.md:45,128` · `internal/copilot` does not exist
Not in `ls app/internal/`. The word survives only as a DB-pool name (`main.go:121,147`
`copilotAnalyticsMaxConns`, `internal/workforce/manager.go:21` "the analytics (copilot) pool").
"**Copilot** (`internal/copilot`) — internal assistant flows" is a phantom feature entry.

### M7 — `backend.md:46` · "the labs-api client is **currently wired as nil**"
Unqualified, and false at HEAD: `main.go:735-738` wires the real client whenever `LABS_API_URL` is
set; nil is only the unset-local-dev path. The same file states it correctly at `:172`, and
`ai-labs.md:123-124` states it correctly. *Fix:* add "when `LABS_API_URL` is unset".

### M8 — `backend.md:237` · "the most recent set of migrations (May 2026)"
The latest migrations are July 2026: `20260722081626_jobsim_data_model.sql`, `20260722104506.sql`,
`20260724132049_cms_data_model.sql`, `20260728103254_…`, `20260729133514.sql` (the mirror drop) —
i.e. the exact migrations this realignment is about.

### M9 — `cms.md:152` · "Cosmo Router now composes `backend` alone" (unfenced)
The composition claim is true for prod (`supergraph-config-prod.yaml` lists `backend` alone), but the
sentence is present-tense and unfenced in a doc whose banner never mentions the router's removal from
local compose at `2adcf71`. *Fix:* "In production the Cosmo Router composes `backend` alone; there is
no router on a local stack since `2adcf71`."

### M10 — `cms.md:73,76` · Go version of the frozen cms repo
"Go 1.25 (primary)" and "built in `golang:1.25-bookworm`" → `cms/go.mod:3` is `go 1.26.4` and both
`Dockerfile` and `Dockerfile.dev` line 2 are `FROM golang:1.26-bookworm`.

### M11 — `cms.md:165,175` · two self-anchors point at the wrong lines
"(Consistent with **:64** above)" — `:64` is the jobsimulation content-read bullet; the `public`-schema
statement is at `:26-29` / `:74`. "see **:37** in the banner" — `:37` is the Events bullet; the
`additional_repo` / studio statement is at `:42-43`.

### M12 — `sentinel.md:12` · "Language: Go 1.25"
`sentinel/go.mod:3` is `go 1.26.0`. (Casbin v3 at `:13` is correct — `casbin/casbin/v3 v3.10.0`.)

### M13 — `sentinel.md:5` · `docker-compose.yml:97,158`
`AUTHORIZATION_ADDRESS=http://sentinel:8087` appears at `:45` (backend), `:99` (jobsimulation) and
`:160` (cms) at `2adcf71`. Off by 2 in both cited anchors.

### M14 — `sentinel.md:82` · the consumer list restates the claim without the husk fence
"Upstream consumers: every other Anthropos service that gates requests (`app`, `cms`,
`jobsimulation`, `messenger`)" — `:5` correctly fences `cms`/`jobsimulation` as husks that "sit off
every request path"; this bullet drops the fence. The restates-it-in-different-words shape.

---

## Files that verified clean

- **`jobsimulation.md`** — everything checked holds: 23 `CREATE TABLE`s in
  `20260722081626_jobsim_data_model.sql` (`:24` "the 23 session/run tables" is exact);
  `public.sessions` gone / `job_simulation_sessions` created (`20260722104506.sql:2,79`);
  `docker-compose.yml:52` (backend) and `:258` (messenger) both `http://jobsimulation:8401`;
  `app/main.go:971-973` `log.Fatalf` on missing `DIRECTUS_BASE_ADDR`; `app/main.go:1195` jobsim
  handler; the frozen repo really has `internal/runner/runner.go` with the "formerly the standalone
  'roadrunner' service" header; `cmd/` holds only `aggregate | clone_session | test | validate` +
  `root` (so "no `serve`/`run` subcommand" is right); `pubsub.NewSubscriberServer` at `cmd/root.go:121`;
  `ROADRUNNER_RPC_ADDR` read by zero Go files. **0 findings.**
- **`graphql-wundergraph.md`** — the most accurate file in the group. Verified exact:
  `terraform/main.tf:20` `service_desired_count = 1`; `subgraphs.conf` = `BACKEND=v1.360.0`;
  `schemas/` = `backend.graphqls` alone; `supergraph-config-prod.yaml` = one subgraph;
  `federation_version: =2.3.2`; `wgc@0.104.0`; `ghcr.io/wundergraph/cosmo/router:0.275.0`;
  `config.compose.yaml` playground/introspection on vs `config.prod.yaml` off, `listen_addr
  0.0.0.0:8080`, `graphql_path /graphql`; `package.json` = `{"name":"graphql-wundegraph"}`;
  `docker-compose.yml:334,352` are the two frontend endpoints; and `app` really serves the Apollo
  Sandbox at `/graphql` with the query handler at `/graphql/query`
  (`internal/web/backend/backend.go:315,317`). **0 findings.**
- **`roadrunner.md`** — `roadrunner/terraform/main.tf:19` `service_desired_count = 1`; `repos.yml`
  holds exactly 9 repos with roadrunner at `:29-31` marked "legacy — folded into app";
  `docker-compose.yml:281` still defines the container and `:297` still sets
  `JUDGE0_BASE_URL=http://52.48.139.23:2358`; `go.mod` = `go 1.25.0`; zero `*_test.go` files in the
  repo; ports 10400/10401. The "orphaned, not absent" framing matches HEAD precisely. **0 findings.**
- **`ai-labs.md`** — `app` is `v1.363.2 @ 5ba17044`; `creditCost` = build 5 / refine 1 / translate 1
  (`internal/credits/cost.go:86-90` — note the *code comment* at `:29` is the stale one, saying
  "5 credits per refine turn"; the doc matches the map, which is authoritative);
  `MarginMarkup = 1.40` (`:133`), `PricePerCreditUSD = 0.45` (`:142`), `DefaultSeedBalance = 500`
  (`:231`); the credits router registers **only** `GET /balance` + `GET /transactions`
  (`internal/web/backend/credits/handler.go:44-45`) so "`POST /credits/purchase` was removed" is
  right; no `checkout.session.completed` handling in `internal/payments/handler.go`; migrations
  `20260529072659`, `20260617120000`, `20260617203555`, `20260626120000`, `20260717151144` all
  present; `labsapi` default `:7070`; reconciler `Interval: 30 * time.Second`
  (`spend_reconciler.go:60`); `cmd/labsimport|labskey|labsdemo` (`:8099`) all exist;
  `/v1/labs` group gated on scope `labs:write` (`internal/web/backend/labs_admin.go:31`);
  `stripe-go/v74 v74.30.0`; `LabSessionService` behind `authn.HTTPAuthnMiddleware`
  (`main.go:1219`); the `LABS_API_URL` conditional stated correctly at `:123-124`. **0 findings.**

---

## Counts

**4 BLOCKERS, 14 minors.**

**Character of the group: sharply mixed, and the split runs the opposite way to the hypothesis.** All
four blockers are in the two *swept* files (`hiring.md` ×2, `backend.md` ×2), and two of the four
(B2, and B1's `:144` twin) are unmistakably repair residue — a corrected sentence with the old claim
still standing two lines away in the same paragraph. `cms.md` and `jobsimulation.md` came through the
sweep clean of blockers. The four never-edited files produced **zero** blockers between them and read
as genuinely well-grounded rather than under-audited: `graphql-wundergraph.md` and `roadrunner.md` in
particular are exact on every pin I checked, including the newest drift (the router's deletion from
compose). The residual minors cluster in two shapes — **stale `file:line` anchors on true claims**
(`hiring.md` ×3 sites, `sentinel.md`, `cms.md` self-anchors) and **frozen-repo Go versions** that
moved to 1.26 under docs that still say 1.25.
