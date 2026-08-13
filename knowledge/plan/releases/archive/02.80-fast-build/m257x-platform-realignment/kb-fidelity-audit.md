---
title: "KB Fidelity Audit — M257x platform re-alignment"
date: 2026-07-31
scope: milestone:M257x
invoked-by: build-mstone-iters (Phase 0b)
platform_ref: "platform @ origin HEAD 1e8e75400c66dbd96abf4a1aca7e7a7cecaea497 (2026-07-30T08:26:40Z)"
---

## Verdict

**YELLOW — proceed with tracking.** 3 blocker-severity findings; 2 are neutralised by being
recorded here + in `decisions.md` (KB-1, KB-2), 1 needs a planning decision on **exit-gate clause 3**
before the gate can be honestly evaluated (KB-3). None blocks iter-01 from *starting*.

Why not RED: every stale claim this audit found is either (a) already declared as this milestone's
subject matter — Open Questions 1–5 and the inherited-evidence table — or (b) newly recorded here so
iter-01 cannot read it as truth. Returning RED would deadlock: the remedy for these blockers *is* the
milestone. Why not GREEN: two load-bearing traps were genuinely undocumented and unanticipated (KB-1,
KB-2), and the gate's own clause 3 rests on a premise that platform source contradicts (KB-3).

## Method note — how absence was proven

The brief warned that reviewers were wrong 5× on this codebase, twice via a `grep` silenced by a stray
NUL byte. Search-set integrity was therefore established **first**, not assumed:

| check | result |
|---|---|
| `find corpus -name '*.md' \| wc -l` vs `rg --files corpus/ \| wc -l` | **88 == 88** — nothing excluded |
| NUL bytes in any `corpus/**/*.md` | **0** (`LC_ALL=C grep -qP "\x00"` per file) |
| `file(1)` on all 88 | all `text`; none binary |
| every absence claim below | run with **`rg -a`** *and* **`grep -a -rn`**, counts compared |

**A sub-agent's claim that `corpus/services/next-web-app.md` is binary-skipped is REFUTED** — see
"Claims that did not survive verification".

## Topic Inventory

| Topic | Knowledge doc | Code paths (platform @ origin HEAD) | Status |
|---|---|---|---|
| jobsimulation service state | `corpus/services/jobsimulation.md`, `architecture/architecture_overview.md`, `architecture/service_taxonomy.md` | `platform/repos.yml`, `platform/docker-compose.yml:108-167`, `app/terraform/migrations/20260722081626_jobsim_data_model.sql` | **PAIRED — STALE** |
| cms service state / schema | `corpus/services/cms.md` | `platform/repos.yml`, `platform/docker-compose.yml:169-212`, `app/terraform/migrations/20260724132049_cms_data_model.sql` | **PAIRED — STALE** |
| roadrunner service state | `corpus/services/roadrunner.md` | `platform/repos.yml`, `platform/docker-compose.yml:306-334` | **PAIRED — mostly ALIGNED** |
| service→schema map | `architecture/architecture_overview.md:270-276`, `ops/platform_repo.md:92`, `ops/setup_guide.md:286-300` | `platform/repos.yml`, `platform/Makefile:14`, `platform/postgresql/Dockerfile` | **PAIRED — STALE** |
| federation subgraph count | `services/graphql-wundergraph.md`, `architecture/external_services.md:307-384`, +5 docs | `graphql-wundergraph/supergraph-config-{compose,prod}.yaml`, `schemas/`, `Dockerfile.dev:18-23` | **PAIRED — STALE (3 → 1)** |
| studio-room placement | `corpus/services/studio-room.md`, `services/cms.md` | `cms/Makefile:11-17`, **`app/Dockerfile.dev:24-41`**, `app/.gitignore:78-79` | **PAIRED — INCOMPLETE (app half absent)** |
| `local_*` session mirrors | `corpus/services/hiring.md:105,132-143`, `ops/seeding-spec.md:386-392` | `app/terraform/migrations/20260729133514.sql:62-63` (**DROPPED**) | **PAIRED — STALE** |
| canonical jobsim session table | — | `app/terraform/migrations/20260722104506.sql:2`, `app/internal/data/ent/schema/job_simulation_session.go:7-9` | **CODE-ONLY (blind trap — KB-2)** |
| app-owned domains | `services/{coursebuilder,ai-labs,askengine,academy-backend}.md` | `app/internal/{coursebuilder,askengine,academy}/` | **PAIRED — substantive** |
| net-new org repos | — | 8+ org repos in neither `repos.yml` nor corpus | **CODE-ONLY (declared OQ3)** |
| re-ground procedure | `corpus/ops/platform-alignment.md` | — | **BLIND-AREA — DECLARED** (`overview.md:57`) |
| migration-status map | *(no target path declared)* | — | **BLIND-AREA — declared-but-unpathed** |

## Fidelity Findings

### F1 — the "fold into app" commit is DOCS-ONLY; compose contradicts it (**blocker → KB-3**)
- **Source:** `overview.md:27-30` reads `repos.yml` as establishing that cms/jobsimulation/roadrunner are folded in.
- **Expected (per `repos.yml` header):** *"skiller, skillpath, roadrunner, jobsimulation and cms are all served in-process … they own no local schema and **are not part of the stack**."*
- **Actual:** commit `236771f103` (2026-07-29T14:06:49Z), subject **`docs:` fold jobsimulation + cms into backend (cms-in-app v8.0)**, changed **exactly three files: `CLAUDE.md`, `README.md`, `repos.yml`**. `docker-compose.yml` was **not touched**. At HEAD, `jobsimulation` (`:108`, profiles `:165`), `cms` (`:169`, profiles `:212`) and `roadrunner` (`:306`, profiles `:334`) are **all still in the DEFAULT `graphql` profile**, still built from their own repos, and `graphql` still `depends_on` jobsimulation + cms (`:20-25`). `backend` still exports `CMS_RPC_ADDR=http://cms:8091` (`:69`) and `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` (`:75`).
- **Verdict:** **the prose is aspirational; only the `migrations:` FIELD is operative.** `platform/Makefile:14` derives `MIGRATION_REPOS` from `/migrations: true/` alone → **only `app` migrates**, and `platform/postgresql/Dockerfile` creates **no** schemas. So "own no local schema" is TRUE and machine-checkable; "not part of the stack" is FALSE at the same sha.
- **Consequence for the gate:** exit-gate clause 3 requires the map be *"machine-fenced against `repos.yml`"*. Fencing against the **prose** would encode a falsehood. Only `name`/`type`/`migrations`/`schema` are fenceable.
- **Fix owner:** planning — narrow clause 3's fence to the machine-readable fields, and have the map carry `running_but_unfederated` as a distinct state from `merged-into-app`.

### F2 — the supergraph is **1 subgraph**, not 3 (16 stale numeric claims)
- **Source:** `architecture/architecture_overview.md:105,202,220,226`; `architecture/external_services.md:307,312,332,379-384`; `architecture/service_taxonomy.md:288`; `services/graphql-wundergraph.md:9,16,86-89,94,103-105`; `services/backend.md:49,67-68`; `services/cms.md:104`; `services/README.md:30`; `services/skillpath.md:31-32`; `ops/update_guide.md:19`.
- **Expected:** 3 subgraphs — backend/app + jobsimulation + cms.
- **Actual:** **one** — `graphql-wundergraph/supergraph-config-compose.yaml` and `-prod.yaml` each declare a single `backend` subgraph; `schemas/` holds only `backend.graphqls`; `Dockerfile.dev:18-23` copies SDL **only** from `app/internal/web/backend/graphql/graph/schemas/` and states *"cms + skillpath folded into the backend subgraph … there are no standalone subgraph SDLs"*; `subgraphs.conf` has the single line `BACKEND=v1.360.0`. Landed `915da06c58` (2026-07-29T09:24:38Z), whose diff deletes the `jobsimulation_subgraph_version` **and** `cms_subgraph_version` inputs, the `JOBSIMULATION=`/`CMS=` sed blocks, and both release links.
- **Verdict:** **STALE.** Answers Open Question 4: **no, the 3-subgraph count does not hold.** `ops/update_guide.md:19` ("3 subgraphs … **not 4**") is stale in the same direction it was written to correct.
- **Note:** this makes cms + jobsimulation **running-but-unfederated** — the exact classification `services/roadrunner.md:16-17` already uses ("built + started but off every request path", ORPHANED). That precedent is the right template.
- **Fix owner:** update docs.

### F3 — service→schema map still asserts `cms` and `jobsimulation` schemas
- **Source:** `architecture/architecture_overview.md:272-273` (explicitly *"source: `platform/repos.yml` `schema:` field for services with `migrations: true`"*); `ops/platform_repo.md:92`; `ops/setup_guide.md:291-292` + `:487` + `:684`; `services/jobsimulation.md:8,16`; `services/cms.md:26,113`; `tools/toolchain_overview.md:54`.
- **Actual:** `repos.yml` @ HEAD: `cms → migrations: false`, `jobsimulation → migrations: false`, **`schema:` keys deleted**; `app` is the sole `migrations: true`. A fresh stack creates neither schema.
- **Verdict:** **STALE**, and the highest-leverage anchor because `architecture_overview.md:272` names `repos.yml` as its source — a reader trusts it as derived.
- **Precision that must not be lost:** these claims may remain **true of the long-lived prod DB** (where the schemas still exist) while **false of a freshly-migrated stack**. The corpus nowhere distinguishes *prod-observed* from *fresh-stack* truth — and that conflation is a plausible root cause of the class recurring three times. `ops/db-access.md:55-57`'s `cms.studio_*` row-counts are prod observations and are **not** invalidated by F3.
- **Fix owner:** update docs; add the prod-vs-fresh-stack distinction.

### F4 — `local_*` session mirrors are DROPPED, but the corpus still requires co-writing them (**blocker → KB-1**)
- **Source:** `services/hiring.md:105` cites `app/internal/data/ent/schema/local_jobsimulation_session.go:52`; `:132-143` makes the co-write mandatory (*"`jobsimulation.sessions` twin are always co-written"*); `:182`; `ops/seeding-spec.md:386-392,416,528-536`; `ops/demo/stories-spec.md:666,680`; `ops/demo/content-stories-routes.md:135,152` (the "generalized manager-view MIRROR trap").
- **Actual:** `app/terraform/migrations/20260729133514.sql:62-63` — `DROP TABLE "local_jobsimulation_sessions"; DROP TABLE "local_skill_path_sessions";`. The migration first re-points `organization_assignment_sessions` FKs to canonical `job_simulation_sessions` / `skill_path_sessions` (`:52-56`), then drops the mirrors and the `on_insert_local_jobsimulation_sessions_update_memberships()` trigger (`:64`), noting *"the app now maintains `memberships.last_activity_date` itself on session events."* The Ent schema file `local_jobsimulation_session.go` no longer exists anywhere in `app`.
- **Verdict:** **STALE — load-bearing.** M257 iter-03 fixed the **rext** side (34 sites / 20 files). **The corpus side was never reconciled**, and nothing in M257x's own docs flags it. An iter reading `hiring.md:132-143` as truth would re-introduce writes to a dropped table.
- **Also:** `hiring.md:105` is a **numeric line-anchor pointing at a deleted file** — the only broken code anchor in the gate scope (3 of 4 checked anchors are in range).
- **Fix owner:** update docs. Recorded as **KB-1**.

### F5 — `app` embeds studio-room; the corpus says CMS-only, in 30 places, and never mentions `app`
- **Source:** 30 placement-asserting lines across 10 files, incl. `services/studio-room.md:15,25,27`; `services/cms.md:9,11`; `architecture/architecture_overview.md:13,23,54,81,141`; `architecture/ai_architecture.md:52,54`; `architecture/service_taxonomy.md:57,149-179,381`; `architecture/dependency_map.md:12,23,85`.
- **Actual:** `app/Dockerfile.dev:24-26` switches the runtime stage to `python:3.11-slim` *"cms-in-app M804 (L3): Python 3.11 runtime for the Studio pipeline (python3 studio/gen.py)"*; `:38-41` `COPY --from=build /build/studio ./studio` + `pip install -r studio/requirements.txt`. `app/.gitignore:78-79`: *"# Python studio runtime (anthropos-studio-room) — pulled at build via additional_repo, like cms"* / `studio/*`. **`cms` still embeds it too** (`cms/Makefile:11-17` `init-studio` clones `anthropos-studio-room` → `studio/`), so this is an **addition, not a move**.
- **Verdict:** **critical undocumented behavior.** `services/backend.md` (the canonical `app` doc) mentions "studio" exactly once, `:159` *"Studio-Desk (for org-level metadata)"* — a different product. Confirms `DOC-M257-studio-in-app`.
- **Platform-side ambiguity worth carrying:** `app/Dockerfile.dev:38-39` calls `./studio` a *"git submodule … pinned SHA"* and says it *"Needs `git submodule update --init --recursive`"*, but `app` has **no `.gitmodules`** and no `studio` tree entry, while `.gitignore:78` says *"pulled at build via additional_repo, like cms."* Two acquisition stories in one repo — the likely root of M257's *"`app`/studio had no rext acquisition path"*.
- **Fix owner:** update docs (add the `app` half; keep the `cms` half).

### F6 — `ant-academy` is claimed to be in `repos.yml`; it is not
- **Source:** `ops/platform_repo.md:93` lists `ant-academy` (node-npm) among `repos.yml` Node entries; `ops/setup_guide.md:297` lists it in the `make init` clone table.
- **Actual:** `repos.yml` @ HEAD has **10** entries and **zero** occurrences of `ant-academy`. This also contradicts the corpus's own `CLAUDE.md:196` (*"**NOT in `repos.yml`** (by design — v1.10b M49 #5)"*) and `services/roadrunner.md:13-14` (*"1 of the 10 repos"* — that count is **correct**).
- **Verdict:** **STALE**, and an internal contradiction. Low blast radius, cheap fix. Out of the gate's declared scope (`corpus/ops/**`).

### F7 — the documented cold-init procedure migrates repos that no longer migrate
- **Source:** `ops/setup_guide.md:684-686` — `migrate-dev.sh` *"create schemas (`extensions`/`sentinel`/`cms`/`jobsimulation`) … → atlas-migrate the 3 migrating services (`app:public` / `cms` / `jobsimulation`)"*.
- **Actual:** only `app` has migrations. Two of the three targets have no migration dir to apply.
- **Verdict:** **STALE.** rext-side; feeds exit-gate clause 4. Out of the gate's declared doc scope but load-bearing for the milestone.

### F8 — a stale build-context anchor
- **Source:** `architecture/external_services.md:366` — `COPY jobsimulation/internal/graph/schemas/ /tmp/schemas/jobsimulation/`.
- **Actual:** `graphql-wundergraph/Dockerfile.dev:18` copies only `app/internal/web/backend/graphql/graph/schemas/`.
- **Verdict:** **STALE.** Same root as F2.

## Completeness Gaps

1. **CRITICAL — the canonical jobsim session table is undocumented, and the obvious re-point is silently wrong (KB-2).** `app` @ HEAD has **two** session tables in `public`: `sessions` (created by `20260722081626_jobsim_data_model.sql:2`, the copied-in legacy jobsim model) **and** `job_simulation_sessions` (created by `20260722104506.sql:2`). `app/internal/data/ent/schema/job_simulation_session.go:7-9` is explicit: the Go type was *"renamed from the too-generic `Session`"*, the table is `job_simulation_sessions`, and **"app's GraphQL `Session` type binds to this entity"**. `20260729133514.sql:52-56` re-points the assignment FKs to `job_simulation_sessions`, confirming which is canonical. **The naive re-point `jobsimulation.sessions` → `public.sessions` will not error — `public.sessions` exists — and will not surface in app's GraphQL.** That is a write-succeeds/read-blank failure, the same signature as M257's `|| echo 0`. No corpus doc names either table.
2. **CRITICAL — no doc covers the `running_but_unfederated` state.** cms + jobsimulation now match `roadrunner.md:16-17`'s ORPHANED shape (container up, off every request path). The service docs have no vocabulary for it, so `migrations: false` reads as "gone" and the live container reads as "fine".
3. **Net-new org repos (declared OQ3).** 8+ active org repos appear in neither `repos.yml` nor the corpus: `ant-observability`, `AI-Labs`, `hyper-studio`, `auth`, `simulation-form`, `Analytics-and-Reports`, `sim-qa`, `livekit-agent-chain` — plus `livekit-agent{,-azure-eu,-azure-us,-azure-eu-fr}`, `demo-environment`, `transcoder`, `analytics-go`, `studio-tools`, `directus`, `metabase`, `judge0`, `clerk`. Naming traps for the map: `auth` (org repo) ≠ `authn` (documented shared lib); `AI-Labs` (repo) ≠ the documented `app` AI-Labs **domain**; `hyper-studio` appears only as an env-template donor (`ops/secrets-spec.md:297,304`); `Analytics-and-Reports` appears only as a unix group (`ops/staging-bringup.md:100,103`); `directus`/`judge0`/`metabase`/`clerk` are documented as third-party products, never as org repos. `ai-labs.md:26,60,113` references a `labs-api` control plane **without naming its repo** — likely `AI-Labs`.
4. **`app`-owned domain docs are healthy.** `coursebuilder.md` (132 L), `ai-labs.md` (153 L), `askengine.md` (121 L), `academy-backend.md` (137 L) — all substantive, template-conformant, and linked from `services/README.md:41-44` and `CLAUDE.md:181`. **Not** linked from `corpus/README.md:31` (minor).
5. **Frontmatter freshness (claim type 7): N/A.** 0 of 88 corpus files use YAML frontmatter; the convention is an H1 (`services/TEMPLATE.md`).
6. **Cross-references (claim type 6): clean.** 0 broken relative links across the 39 files in `corpus/services/**` + `corpus/architecture/**`.
7. **Numeric line anchors (claim type 8):** only 4 platform-code anchors exist in the gate scope; 3 in range, 1 broken (F4 / `hiring.md:105`).
8. **Unpathed deliverable.** The migration-status map is in scope (`overview.md:53`) and in the gate (clause 3) but has **no declared target path**, unlike `platform-alignment.md` (`overview.md:57`). Recommend pinning one so this audit can verify it at the exit gate.

## Claims that did NOT survive verification

Recorded because the brief demands it — a sub-agent produced two confident false claims, both caught:

1. **"`corpus/services/next-web-app.md` is not valid UTF-8; plain `rg`/`grep` treat it as binary and silently skip it."** **FALSE.** Four methods: `tr -dc '\0' | wc -c` → **0** NUL bytes; `file(1)` → *"Unicode text, UTF-8 text"*; Python `bytes.decode('utf-8')` → **decodes clean** (9558 B, 126 lines); and `rg -c` / `grep -c` return **identical counts with and without `-a`** (1 == 1). The only tool that dissents is macOS `iconv -f UTF-8 -t UTF-8`, which is the outlier, not the authority. The file is searched normally. (The agent's inventory counts were unaffected — it had used `-a` throughout — but the *claim* is wrong and would have re-created the exact false-negative panic the brief warns about.)
2. **"`CLAUDE.md:353`'s '27 service docs' is off by one; there are 28."** **FALSE.** `ls corpus/services/*.md | grep -vE 'README|TEMPLATE' | wc -l` → **27**. `CLAUDE.md:353` is correct.

## Applied Fixes

None to `corpus/**`. Deliberate: every corpus finding here is **this milestone's declared deliverable**
(`overview.md:54` — *"`corpus/services/**` (29 docs) and `corpus/architecture/**` reconciled to what the
map establishes"*), and the skill forbids expanding a fidelity audit into a cluster rewrite. Fixing them
inline would also strip the milestone of its measurable exit condition.

Recorded instead (planning files only, uncommitted — the parent agent owns commits):
- `decisions.md` — KB-1, KB-2, KB-3.
- `spec-notes.md` — the verified topic → doc → code triples + the platform ground-truth block.

## Open Items (require user / planning decision)

- **KB-3 (exit-gate clause 3).** Clause 3 fences the map against `repos.yml`. Per F1, `repos.yml`'s
  *prose* is contradicted by `docker-compose.yml` at the same sha, and the `docs:`-only commit never
  changed orchestration. Decide: (a) narrow the fence to the machine-readable fields
  (`name`/`type`/`migrations`/`schema`) only, and (b) add `running_but_unfederated` to the state
  enum `{live-standalone, merged-into-app, decommissioned, net-new}` — cms, jobsimulation and
  roadrunner are all in that state today and none of the four existing values fits.
- **Unpathed deliverable.** Pin a target path for the migration-status map.

## Gate Result

**YELLOW — proceed with tracking.** iter-01 may start. Carry KB-1 and KB-2 into the bootstrap tok's
strategy; resolve KB-3 before the exit gate is evaluated. The one declared blind area
(`corpus/ops/platform-alignment.md`) is correctly covered by `Delivers →` at `overview.md:57` and does
not block.
