---
milestone: M210
slug: corpus-reground
version: v2.1 "quick change"
milestone_shape: section
status: archived
created: 2026-07-08
last_updated: 2026-07-08
complexity: medium
depends_on: M209
delivers: the completed corpus (corpus/services/skiller.md stub + backend.md ownership + re-pointed rext-facing tooling docs)
issues: the colleague's docs/skiller-in-app-merge sweep is correct-but-incomplete; 5-6 rext-facing docs still describe skiller.*
---

# M210 — Corpus + skills re-ground

## Goal
Land the colleague's `origin/docs/skiller-in-app-merge` sweep as the **complete, internally-consistent** corpus
re-ground — validate correctness, fix the misses, and **flip the rext-facing tooling-doc bodies to `public.*` in
lockstep with M209's landed schema**.

## Why section
The gap is fully enumerated (design audit): the branch is **correct-but-incomplete**. The architecture/subgraph
half is solid and lands as-is; the rext-facing half is a bounded, known list of doc-body flips + one missed file.
Enumerable → `section`. (Scope-flex: if the reconcile proves large, split into 3a land+reconcile / 3b rext-gap-fill.)

## Why it depends on M209 (not parallel)
The branch **cannot land present-tense independently** (adversarially confirmed at design): it asserts a
post-merge world (4 subgraphs, `backend:8083`, "don't query skiller schema") that the tooling only reaches after
M209 re-points the code + M208 re-syncs the stacks. The tooling-doc **bodies** (`snapshot-spec` etc.) must state
`public.*` to match M209 — so M210 runs **after** M209 (the user's strictly-sequential choice makes this clean).

## Scope
- **In:**
  - **Adopt + validate** `origin/docs/skiller-in-app-merge` — the architecture/subgraph/service half is correct
    (4 subgraphs everywhere; tables in `public`, names unchanged; `SkillerService` served by `backend`; RPC
    re-pointed; `categoryTree`/`fullCategoryTree` dropped-not-ported; `skiller.md` reframed as a merged/legacy
    stub → `backend.md`). Land it.
  - **Fix the fully-missed file** — `corpus/ops/demo/profile-completeness-spec.md` (the member count 43/44 → 44/44
    after the merged member surface).
  - **Flip the 5-6 rext-facing tooling-doc BODIES to `public.*`** (not just annotate) in lockstep with M209, and
    delete the interim disclosure notes:
    - `corpus/ops/snapshot-spec.md` (26 mentions — the taxonomy surface enumeration + the FidelityProbe gene + the
      capture predicate SQL)
    - `corpus/ops/safety.md` (the firewall public-predicate evidence row — `skiller.skills 42,763 public`)
    - `corpus/ops/demo/recipe-snapshot-world.md` (the replay target — "the stack's skiller Postgres schema")
    - `corpus/ops/demo/stories-spec.md` (the 7-table fan-out `node_id`/`name` joins)
    - `corpus/ops/seeding-spec.md` ("the public skiller catalog / taxonomy")
    - `corpus/ops/demo/coverage-protocol.md` (the removed skiller subgraph round-trip + `<stack>-skiller-1`
      container-log reference)
  - **Reconcile the db-access ↔ tooling contradiction** — resolved once the tooling-doc bodies flip (both sides
    now say `public.*`).
  - **Sweep the skill files** — `dev-up/reference.md`, `stack-snapshot/SKILL.md`, `stack-update/reference.md`,
    `db-query/SKILL.md` — so container counts (12→11), migration lists, RPC addr, subgraph counts (5→4) match the
    re-synced stacks.
  - **Update `CLAUDE.md`** service catalog (the skiller Tier-1 entry → merged-into-app stub; the "5 subgraphs" →
    4; the `SKILLER_RPC_ADDR` note).
- **Out:** rext code (M209); live bring-up (M211).

## KB dependencies / delivers
Reads: the M209 landed schema (the tooling-doc bodies must match). Delivers: `corpus/services/skiller.md` (stub) +
`corpus/services/backend.md` (taxonomy ownership) + the re-pointed rext-facing ops/spec docs.

## Done-bar
- The corpus contains **0** `skiller.<table>` descriptions of what the (now re-grounded) rext tooling queries; the
  db-access ↔ snapshot/seed docs agree on `public.*`; the branch's architecture half is landed; the missed file +
  skill files are swept; `CLAUDE.md` catalog current.
