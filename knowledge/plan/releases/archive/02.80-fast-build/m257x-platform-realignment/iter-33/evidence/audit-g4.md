# iter33 KB-fidelity audit — group 4

Audited read-only against platform origin HEAD `2adcf71`, `app` @ `5ba17044`, using the clones under
`/Users/marco/workspace/anthropos/rosetta/stack-demo/`.

## 1. Positive control

| file | `wc -l` | read |
|---|---|---|
| `corpus/services/hiring.md` | 310 | read to line 310 (full) |
| `corpus/services/backend.md` | 251 | read to line 251 (full) |
| `corpus/services/chronos.md` | 245 | read to line 245 (full) |
| `corpus/services/cms.md` | 237 | read to line 237 (full) |
| `corpus/services/jobsimulation.md` | 221 | read to line 221 (full) |
| `corpus/services/graphql-wundergraph.md` | 196 | read to line 196 (full) |
| `corpus/architecture/platform-migration-status.md` | 189 | read to line 189 (full) |

All 7 files read top to bottom. 1649/1649 lines.

---

## 2. Findings

### BLOCKERS

---

**B1 — `corpus/services/hiring.md:152`**

> `completition_status` (**note the misspelled column**; values `passed`/`failed`/`pending`/`SIMULATION…`)

FALSE at HEAD. The DB column is spelled **correctly**: `app/terraform/migrations/20260722104506.sql:12`
`"completion_status" character varying NOT NULL DEFAULT 'pending'`, and
`app/internal/data/ent/schema/job_simulation_session.go:39` `field.Enum("completion_status")`. The rext
seeder already writes the correct name (`stack-seeding/seeders/persona_write.go:152` `sessionCols()` →
`"completion_status"`). The `completition` misspelling survives **only** in the GraphQL sort-field enum
(`enum.InsightsSortFieldCompletitionStatus`, `intelligence.go:885,1757`) and a JSON tag
(`intelligence.go:1676`) — never as a column. This is the §*seeder-output contract* write-set, so acting on
it means a seeder INSERT against a column that does not exist.

**Grade: BLOCKER.** Correction: rename the column in the write-set to `completion_status` and move the
"misspelled" note to the sort-field enum where it actually lives (`hiring.md:129` is correct as-is).

---

**B2 — `corpus/services/backend.md:165`**

> **AI Readiness** (v1.266+, the `internal/workforce` subsystem) … Engine: `internal/workforce/ai_readiness.go` + `readiness_steps.go` + `readiness_narrative.go`

FALSE at HEAD. AI-readiness is its **own top-level package**, `app/internal/aireadiness/` (`manager.go`,
`cycles.go`, `diagnosis.go`, `compare.go`, `completion.go`, `csv.go`, `defaults.go`,
`recommendation_signals.go`, …). `app/internal/workforce/` contains **no** file matching `readi*` — the
three named files do not exist. (The GraphQL path `graph/schemas/ai_readiness.graphqls` in the same bullet
is still correct.) `internal/aireadiness/` is also missing from the §*Key directories* block at :109-159.

**Grade: BLOCKER.** Correction: re-point the engine to `app/internal/aireadiness/` (manager/cycles/diagnosis)
and add `aireadiness/` to the key-directories block.

---

**B3 — `corpus/services/jobsimulation.md:98-99`**

> it reads an `app`-side MIRROR, `public.local_jobsimulation_sessions` (the analog of skill-path's `local_skill_path_session`). Seed the runtime rows only and the manager scoreboard is blank.

FALSE at HEAD. `app/terraform/migrations/20260729133514.sql:58-62` — *"5. Drop the mirrors."* —
`DROP TABLE "local_jobsimulation_sessions"; DROP TABLE "local_skill_path_sessions";`, after back-filling
into the canonical entities. `app/internal/data/ent/schema/` has no `local_jobsimulation_session.go`;
`intelligence.go:1700` reads `m.ent.JobSimulationSession.Query()`. `hiring.md` was re-grounded for this at
M257x iter-23; **this file was not**. A seeder written against this paragraph writes to a dropped table.

**Grade: BLOCKER.** Correction: replace the mirror with `public.job_simulation_sessions` (read by
`app/internal/organization/intelligence.go:1700`) and drop the "seed the mirror or the scoreboard is blank"
trap — there is one row now, not a pair.

---

**B4 — `corpus/services/jobsimulation.md:24-27`**

> **Data** — the 23 session/run tables (`sessions`, `actors`, `interactions`, … ) were re-created in the **`public` schema** by `app/terraform/migrations/20260722081626_jobsim_data_model.sql`, with the **same table names**.

FALSE at HEAD for the headline table. `20260722081626_jobsim_data_model.sql:2` did create `public."sessions"`,
but the **very next migration** dropped it: `20260722104506.sql:2` `CREATE TABLE "job_simulation_sessions"`,
`:79` `DROP TABLE "sessions"`. So `public.sessions` does not exist and "same table names" is untrue of the
one table everything else FKs to. (The other 22 names did survive.)

**Grade: BLOCKER.** Correction: "…with the same table names **except `sessions`, renamed to
`public.job_simulation_sessions` at `20260722104506.sql:2/:79`**".

---

**B5 — `corpus/services/jobsimulation.md:107`**

> So the M23 content cutover (re-pointing CMS's `DIRECTUS_BASE_ADDR` at the per-stack Directus) carries jobsimulation's content reads to local automatically; no jobsimulation env change is needed.

FALSE at HEAD. Since cms-in-app the Directus reader is **`backend`, in-process**: `app/cms_reader_switch.go`
swaps the content reader to the in-process cms server, and `app/main.go:971-973` `log.Fatalf`s when
`DIRECTUS_BASE_ADDR` is unset ("required to wire the cms-in-app managers"). Re-pointing the `cms` husk alone
does **not** carry the cutover — this was measured live on demo-1 (2026-08-01: 96 Directus lines in
`backend`'s log, all 403, matching prod's answers not the local instance's), which is why rext now sets
`DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` in **both** twins
(`stack-injection/gen_injected_override.py:53`, `stack-core/gen_override.py:58`). The doc still describes the
pre-fix mechanism.

**Grade: BLOCKER.** Correction: the cutover re-points **`backend`** (the in-process reader) as well as the
`cms` husk; the in-app jobsim engine's content reads follow `backend`'s `DIRECTUS_BASE_ADDR`.

---

**B6 — `corpus/services/cms.md:68`**

> **M23 re-points `cms`'s `DIRECTUS_BASE_ADDR` at that local instance** (`http://directus:8055`, the in-network service, #M23-D1) so a `--local-content` stack (demo default; dev opt-in) serves its **own** captured catalog — no live-prod read.

Same defect as B5, at its source. The literal sub-clause ("cms is re-pointed") is still true, but the
**causal claim** — that re-pointing `cms` is what makes the stack serve its own catalog — is false at HEAD:
`backend` is the reader (`app/cms_reader_switch.go`; `app/main.go:971-973`), and rext's
`DIRECTUS_DATA_CONSUMERS` was corrected to `("cms", "backend")` at M257x iter-24 for exactly this reason.
A reader debugging "demo content is prod's / empty" verifies the `cms` re-point, finds it correct, and stops
— which is how this survived four releases.

**Grade: BLOCKER.** Correction: "M23 re-points the data-plane consumers' `DIRECTUS_BASE_ADDR` — since
cms-in-app that is **`backend`** (the in-process reader) *and* the `cms` husk — at that local instance."

---

### Minors

**m1 — `corpus/services/hiring.md:125`**

> | 6 | `intelligence.go:1801` | `Score` ← `ls.Score` (the mirror's score column) |

Two problems, both cosmetic against the doc's own M257x re-ground: the `Score:` assignment is now at
`internal/organization/intelligence.go:1846` (1801 is `organizationAssignmentSessionMap`), and the
parenthetical "the mirror's score column" contradicts the doc's own banner (:19-25) — the mirror is dropped;
it is `public.job_simulation_sessions.score`. Correction: `intelligence.go:1846` | `Score` ← the canonical
session's `score`.

**m2 — `corpus/services/hiring.md:250`**

> as a second offset-port UI container (same recipe as `apps/web` + `studio-desk`), wired to the same fake FAPI + Cosmo + Postgres

"Cosmo" is dead locally — the same paragraph, three lines earlier (:247-249), already states the endpoint is
`backend`'s own `:8082/graphql/query` "the Cosmo/WunderGraph router having been deleted from compose"
(`docker-compose.yml` has no `graphql` service; `repos.yml` has no entry). Self-contradictory leftover.
Correction: "…the same fake FAPI + **`backend`'s `:8082/graphql/query`** + Postgres".

**m3 — `corpus/services/backend.md:45` (and the key-directory line :128)**

> * **Copilot** (`internal/copilot`) — internal assistant flows

`app/internal/copilot/` does **not exist** at HEAD (the only `copilot` references left are
`internal/workforce/{manager,skill_paths}.go`, `internal/aireadiness/{manager,recommendation_signals}.go`,
`internal/taxonomy/resolver.go`, plus the `COPILOT_DB_CONN` env var in `docker-compose.yml:60`). Low blast
radius, so minor rather than blocker. Correction: drop the bullet, or re-point it at the `COPILOT_DB_CONN`
consumers in `workforce`/`aireadiness`.

**m4 — `corpus/services/backend.md:111`**

> `aiacademy/`  Periodic AI Academy catalog sync (fetches catalog.json, populates aiacademy_courses for Talk to Data)

The package is `app/internal/academy/` — there is no `internal/aiacademy/`. Correction: rename the entry to
`academy/`.

**m5 — `corpus/services/backend.md:35`**

> the mux carries `BackendUsersService`, `BackendOrganizationsService`, `SkillerService`, `SkillPathSessionService`, `JobSimulationService`, `CMSService` and `lab.v1.LabSessionService`

`SkillPathSessionService` is **not** registered. `app/main.go` `mux.Handle` calls are exactly: `:1178`
UsersService, `:1179` OrganizationsService, `:1187` SkillerService, `:1195` JobSimulationService, `:1204`
CMSService (conditional on `cmsRPCServer != nil`), `:1219` the LabSession path. A repo-wide grep for
`SkillPathSessionServiceHandler` / `skillpathv1connect` returns nothing outside generated ent code.
Correction: drop `SkillPathSessionService` from the mux list (the skill-path session surface is GraphQL-only
on `backend`).

**m6 — `corpus/services/cms.md:73`**

> * **Language**: Go 1.25 (primary) + Python 3.11 (studio-room)

`cms/go.mod` reads `go 1.26.4` at HEAD (as does `app/go.mod`). Correction: Go 1.26.

**m7 — `corpus/services/cms.md:152`**

> **Federation**: … Cosmo Router now composes `backend` alone.

True in production only; there is no router in a local stack since `2adcf71`. The doc's own banner never
fences this section. Correction: "…in **production** the Cosmo Router composes `backend` alone; locally
there is no router and clients hit `backend:8082/graphql/query` directly."

**m8 — `corpus/services/cms.md:205-213`**

> For Python pipeline development: ```cd cms/studio; pip install -r requirements.txt; python gen.py …```

The HISTORICAL fence at :173-176 explicitly scopes itself to the *First-time setup* block ("The block below
…"), so this later block reads as current guidance, while the banner at :42-43 says the pipeline now ships
in the **`app`** image via the CI `additional_repo` mechanism (app v1.360.1). Correction: extend the
historical fence to cover the Python-dev and `make update-studio` blocks, or re-point them at `app`.

**m9 — `corpus/services/chronos.md:9`**

> …have moved to **in-process Asynq** running inside jobsimulation.

True as history, but "inside jobsimulation" now means "inside the jobsim domain of `app`" — the standalone is
`running_but_unfederated` and the engine lives at `app/internal/jobsimulation/`. Everything from :11 down is
explicitly fenced as historical, so this is confusing-not-false. Correction: "…inside the jobsim engine (now
`app/internal/jobsimulation/`)".

**m10 — `corpus/services/graphql-wundergraph.md:42` (and the `/graphql` row at :98)**

> * Provide a GraphQL **playground + introspection** in dev/compose; both are disabled in production.

`config.compose.yaml:10,12` does set `playground_enabled: true` / `introspection_enabled: true`, but there is
no compose service to serve it — the banner at :12 already says "There is no `:5050` on a local stack", so
this is fenced-by-context rather than false. Correction (optional): mark the compose column of these two rows
"(archived-repo config; no compose service since `2adcf71`)".

**m11 — `corpus/architecture/platform-migration-status.md:51-54`**

> *every service that has ever appeared in `docker-compose.yml`* (same command on that file → 26 names …). Re-run those two commands to audit this table; a name they return that has no row is a gap.

Reproducing the stated command yields 26 names, but one of them is **`app-network`** — a `networks:` key, not
a service — and it correctly has no row. The `repos.yml` half reproduces exactly (14 names, identical set).
A reader following the stated audit instruction hits one false-positive "gap". Correction: note that
`app-network` is a network key and is excluded (25 services + 1 network = the 26 the grep returns).

---

## 3. Per-file verdicts

- **`corpus/services/hiring.md`** — 1 BLOCKER (B1), 2 minors (m1, m2). The v2.8 iter-23 re-ground is
  otherwise accurate: `public.job_simulation_sessions` / `intelligence.go:1700` / the mirror drop at
  `20260729133514.sql:58-62` / `persona_write.go:91,152` all verified exact against HEAD.
- **`corpus/services/backend.md`** — 1 BLOCKER (B2), 3 minors (m3, m4, m5). The merge banner, the
  skiller-in-app fact-sheet, the compose citations (`:255-265`, ports 8081/8082/8083, profiles
  `[graphql, backend, all]`), `main.go:604/735-738/1196-1202`, `internal/rpc/skillerrpc/`, and
  `graph/schemas/skiller_taxonomy.graphqls` all verified true at HEAD.
- **`corpus/services/chronos.md`** — **no blockers**; 1 minor (m9). The banner is correct on the one fact
  most likely to be wrong: decommissioned from orchestration (`045857c`) **but the GitHub repo is NOT
  archived** — which matches ground-truth fact #11 and the migration map's row. Everything from :11 down is
  explicitly fenced as historical, so the `stack-dev/chronos` paths and standalone run instructions are not
  findings.
- **`corpus/services/cms.md`** — 1 BLOCKER (B6), 3 minors (m6, m7, m8). The merge banner is correct
  (including the husk-still-starts caveat, `docker-compose.yml:144/164/256`, `repos.yml:14-16`, the M809
  dormancy at `main.go:1196-1202`) and the `20260724132049_cms_data_model.sql` table list verified exactly
  (6 tables, all in `public`).
- **`corpus/services/jobsimulation.md`** — 3 BLOCKERS (B3, B4, B5), no minors. Everything else verified:
  `docker-compose.yml:52/83/104/118/258`, `repos.yml:17-19`, the archived-repo note, the `internal/runner/`
  in-process Judge0 client (and `app/internal/jobsimulation/runner/`, with `app/internal/roadrunner/` correctly
  absent), the `$HOME/.aws/credentials` bind at `docker-compose.yml:142`, and the "migrate is `app` alone"
  warning.
- **`corpus/services/graphql-wundergraph.md`** — **no blockers**; 1 marginal minor (m10). This file is
  fully swept for ground-truth fact #4: the banner states the router is deleted from both `repos.yml` and
  `docker-compose.yml` at `2adcf71`, the repo is archived, local dev points at `:8082/graphql/query`
  (verified `docker-compose.yml:334,352`), the supergraph is one subgraph (verified: `schemas/` holds
  `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0`, `supergraph-config-prod.yaml` lists
  `backend` alone), and every historical routing/Dockerfile/`make up` claim carries an explicit HISTORICAL
  fence. Version pins verified exact (`cosmo/router:0.275.0`, `wgc@0.104.0`, `federation_version: =2.3.2`,
  `node:22.11-alpine`, `listen_addr 0.0.0.0:8080`, `graphql_path /graphql`).
- **`corpus/architecture/platform-migration-status.md`** — **no blockers**; 1 minor (m11). Audited per
  instruction (prose vs its own table vs the platform, table left alone). §1's two traps, §2's completeness
  arithmetic (14 repos.yml names — reproduced identically), §3's "9 in `repos.yml`" (exactly 9 entries), §4's
  five assertions, and §5's fold order all hold at HEAD. Every spot-checked citation resolves:
  `repos.yml:10-13/14-16/17-19/20-22/23-25/26-28/29-31/34-36/37-39`, `docker-compose.yml:5/18/28/83/144/189/
  220-222/238/240/281/311/344/352/371-372`, `common.yml`-included `postgresql`/`redis`, `app/main.go:604`,
  and `app/internal/roadrunner/` absent.

---

## 4. Totals

- **BLOCKERS: 6** — hiring ×1, backend ×1, cms ×1, jobsimulation ×3.
- **minors: 11** — hiring ×2, backend ×3, cms ×3, chronos ×1, graphql-wundergraph ×1, migration-status ×1.
- **Files not fully read: none.** All 7 files (1649 lines) read top to bottom.

**Pattern worth naming.** Five of the six blockers are *derived-fact rot*, not status rot: every one of these
docs correctly says who is merged (the `ServiceDocStatusFence` is holding), but three name **tables the
platform has since dropped or renamed** (`public.sessions`, `local_jobsimulation_sessions`,
`completition_status`), one names **packages that were split out** (`internal/workforce/ai_readiness.go` →
`internal/aireadiness/`), and two describe a **demo mechanism the tooling itself fixed at iter-24** but the
corpus never followed (`DIRECTUS_DATA_CONSUMERS` gaining `backend`). None of these use the merged/live/gone
vocabulary a status sweep greps for. `hiring.md` was re-grounded for the mirror drop at iter-23;
`jobsimulation.md` — the doc that *owns* those tables — was not.
