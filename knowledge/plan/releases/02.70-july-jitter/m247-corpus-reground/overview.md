---
milestone_shape: section
milestone: M247
title: "corpus re-ground"
status: planned
release: v2.7 "july jitter"
depends_on: [M246]
parallel_with: [M248, M249, M250, M251, M252]
complexity: medium
created: 2026-07-23
---

# M247 — corpus re-ground

## Goal
The corpus describes the **CONSOLIDATED platform**: skillpath is **merged-into-app** (not a live service), the
gateway is **3 subgraphs** (backend/app, jobsimulation, cms), and the new `app` domains have fact sheets. This
re-grounds every doc-body claim that still describes the pre-consolidation topology so the corpus matches the
reality M246 proved live.

## Shape (why this shape)
`section` — the gap is **fully enumerated**: a bounded, known list of doc-body flips (~30 echo files where
"4 subgraphs" and "skillpath is a runtime engine" must be re-pointed), one skillpath→app redirect (mirroring the
`skiller.md` pattern), 4 explicitly named net-new fact sheets, one refresh, and one archive resolution. Enumerable →
`section`. Internally it runs **two-phase**: a wide **core-lanes** concurrent fan-out, then a small **reconcile-tail**
(README/CLAUDE reconcile + grep-verify) once the parallel deltas have landed.

## Scope

### In
- **skillpath redirect** — convert `corpus/services/skillpath.md` → a **merged-into-app REDIRECT** (mirror
  `skiller.md`); move it to the README **archived/merged** table.
- **3-subgraph reclassification sweep** — re-point every "4 subgraphs" → **"3 (backend/app, jobsimulation, cms)"** +
  reclassify **skillpath as not-a-live-service** across the **~30 echo files** (`architecture_overview.md`,
  `service_taxonomy.md`, `graphql-wundergraph.md`, `backend.md`, `cms.md`, `dependency_map.md`, `external_services.md`,
  the ops docs, `CLAUDE.md`); **note jobsim-in-app coming**.
- **4 net-new fact sheets** — author `coursebuilder.md`, `ai-labs.md` (AI Labs + credits / v6.0 shared purse),
  `askengine.md` (Talk-to-Data), `academy-backend.md` (the app-owned academy domain).
- **ai-readiness refresh** — refresh `ai-readiness.md` for the aireadiness-package refactor.
- **roadrunner resolution** — resolve roadrunner **ORPHANED→ARCHIVED** if dead.

### Out
- Any rext/code change.
- The seeder re-point (owned by M246).

## Dependencies & parallelism
- **depends_on:** `M246` — consumes the **confirmed-drift ledger** M246 emits (the drift observed on one cold
  `/demo-up` GREEN against the consolidated platform).
- **parallel_with:** `M248, M249, M250, M251, M252` — M247 runs as the **M247-core** lane in the post-M246 fan-out.

**Intra-milestone LANE decomposition (~4–5× — the biggest intra-milestone win):**
- **8 concurrent doc lanes** — 4 fact-sheets (`coursebuilder` ∥ `ai-labs` ∥ `askengine` ∥ `academy-backend`) ∥
  skillpath-redirect ∥ ai-readiness-refresh ∥ arch-sweep ∥ ops-sweep. All disjoint files → true concurrent fan-out.
- **Serial bottleneck** — a small **README/CLAUDE reconcile** + a **grep-verify tail** (confirm 0 residual
  "4 subgraphs" / "skillpath runtime engine" claims) after the lanes land.
- **The core/reconcile split** — **M247-core** (`CLAUDE.md` + README indices + `architecture/**` + `tools/**` —
  **disjoint from every other milestone**) runs as a true concurrent fan-out lane; **M247-reconcile** (the `ops/demo`
  spec docs the code milestones M249/M252/M253 also touch — e.g. `studio-desk.md`, `frontend-tier.md`) folds into the
  **serial integration tail** and is merged **last** (`…→ M247-reconcile → M254` per the release merge order).
- **Recommended subagents:** up to **~8 concurrent** for the doc lanes, converging to **1** for the reconcile +
  grep-verify tail. `CLAUDE.md` is the **sole property of M247** (coordination rule 5 — every other milestone defers
  its one-line bullet here).

## KB dependencies
- `corpus/services/README.md` — the archived/merged table skillpath moves into.
- `corpus/services/skiller.md` — the redirect pattern to mirror.
- `corpus/services/backend.md` — where the consolidated `app` domains are owned/referenced.
- `corpus/services/TEMPLATE.md` — the fact-sheet shape the 4 new docs follow.

## Delivers
- `corpus/services/skillpath.md` — converted to a merged-into-app **redirect**.
- `corpus/services/coursebuilder.md` — **new** fact sheet.
- `corpus/services/ai-labs.md` — **new** fact sheet (AI Labs + credits / v6.0 shared purse).
- `corpus/services/askengine.md` — **new** fact sheet (Talk-to-Data).
- `corpus/services/academy-backend.md` — **new** fact sheet (the app-owned academy domain).
- `corpus/services/ai-readiness.md` — refreshed for the aireadiness-package refactor.
- The **3-subgraph reconciliation** across the ~30 echo files (`architecture_overview.md`, `service_taxonomy.md`,
  `graphql-wundergraph.md`, `backend.md`, `cms.md`, `dependency_map.md`, `external_services.md`, ops docs, `CLAUDE.md`),
  with skillpath reclassified not-a-live-service and a jobsim-in-app-coming note.
- `corpus/services/README.md` — skillpath moved to the archived/merged table; roadrunner ARCHIVED if dead.

## Open questions
- `coursebuilder` + `credits` sources are **ABSENT from the stale clone** → those 2 fact-sheet lanes need a **fresh
  `app` pull** (M246's clone-pin bump makes them present). Sequence those 2 lanes after the bump lands.
- The reclassification is **semantically heavier than a digit swap** — skillpath is described as a runtime engine in
  **~37 files**; each mention must be re-read in context, not blind sed-replaced.
