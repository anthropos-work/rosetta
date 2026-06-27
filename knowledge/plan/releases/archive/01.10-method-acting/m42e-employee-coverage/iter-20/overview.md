---
iteration_type: tik
iter_shape: production-fix
status: planned
created: 2026-06-25
---

# iter-20 — P6 replay-wiring: sim-embeddings into the set-dress flow + verify the library

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause / the re-scoped-gate design-plan).
This iter executes the **wiring half of design-plan P6** (library set-dress: sim embeddings + categories). The
CAPTURE half (the sanctioned prod read) was done by the orchestrator: the caches are at
`.agentspace/snapshots/sim-embeddings/10146f28…` (cms.similarities + 3 children, public-only) and the
re-captured `.agentspace/snapshots/directus/ea2e187a…` (now carrying 4 library-category tables + `_structure.sql`).

**Cluster / target identified:** iter-19 built the `sim-embeddings` snapshot SURFACE (cms.similarities pgvector +
3 metadata children) + extended the `directus` surface with 4 library-category tables, and the orchestrator
captured both. But the demo set-dress flow's replay loop only invokes `--surface taxonomy` + `--surface directus`
(`dev-stack/dev-setdress.sh` `snapshot_step`'s `for s in taxonomy directus`). So a fresh `demo-up` never loads
`cms.similarities` → the org-member `/library/ai-simulations` view (which runs `searchSimulations` → a pgvector
similarity over `cms.similarities`) stays EMPTY (design-plan root #3). **The sim-embeddings replay invocation is
not wired in.**

**Hypothesis:** Adding `sim-embeddings` to the `dev-setdress.sh` replay loop (it funnels BOTH the demo
`up-injected.sh` and the dev path through one engine) will, on a fresh demo-up — and on a live re-replay against
demo-3 now — load `cms.similarities` (+ children) into the stack's `cms` schema and REINDEX the pgvector column
(the `simembeddings.Surface()` already flags `VectorColumns: [small_embedding3]` → `replay.Run` rebuilds the
index), lighting `/library/ai-simulations`. The re-captured directus surface (with library-category tables) lights
`/library/skill-paths` categories on its existing `--surface directus` replay.

**Expected lift (qualitative, per design-plan P6 / re-scoped gate (a)+(b)):** before → after on demo-3 logged in
as maya-thriving: `/library/ai-simulations` "0 simulations" → real simulations WITH categories; `/library/skill-paths`
no categories → categories present; `searchSimulations` pgvector path returns results.

**Phase plan (coverage-protocol Phase A–E, adapted for a set-dress-wiring tik):**
- Phase 0d (pre-flight tooling): 30s dry-run of `stacksnap replay --surface sim-embeddings` against demo-3's cms
  schema from the captured cache (the pipeline the wiring depends on) — confirm it loads + reindexes before
  committing the wiring.
- Phase A/C (fix): add `sim-embeddings` to the `dev-setdress.sh` replay surface loop (ordered after directus,
  before the per-stack-Directus boot is irrelevant — sim-embeddings targets `cms`, not the per-stack Directus),
  with a NON-FATAL skip per the existing exit-code handling (a missing cms schema or cache-miss must not abort the
  pass — the seed floor still runs). Re-apply: replay sim-embeddings + re-replay directus into demo-3.
- Phase D (re-measure): authenticated probe (maya-thriving) of `/library/ai-simulations` + `/library/skill-paths`;
  DB confirm `cms.similarities` row count + a `searchSimulations` pgvector smoke.
- Phase E (close): grade on whether the library renders real sims + categories.

**Escalation conditions:** if the sim-embeddings replay fails because the demo's `cms` schema is missing the
`similarities` table (a platform-schema gap, not a wiring gap) → that would be a re-scope-trigger (platform
read-only); but the surface was demo-verified in iter-19's dry-run, so this is not expected.

**Acceptable close-no-lift outcomes:** if the probe shows the library still empty AFTER a confirmed replay+reindex
(rows in `cms.similarities` but the view empty), the root is elsewhere (a federation/serve gap) — falsify the
"replay lights the library" hypothesis, record the true root, route forward. Wiring still lands (reproducibility).
