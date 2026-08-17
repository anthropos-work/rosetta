# Release Retro — v2.9 "new alphabet"

**Shipped:** 2026-08-16 · **Milestones:** M259 → M265 (7) · **Branch:** `release/02.90-new-alphabet`
**Theme:** the platform's vocabulary was rewritten under us — 43,584 skills / 22,511 roles → **3,562 /
706** — and the corpus, the tooling and the demo had to learn the new alphabet.

## What shipped

| Milestone | Delivered |
|---|---|
| M259 | The canon's ground truth, measured from `app/taxonomy-canon/` — no production read needed |
| M260 | The measurement floor: no tool asserts a taxonomy size it did not measure this run |
| M261 | The canon loaded (3,562 / 706) and captured as a new snapshot ref |
| M262 | The seed speaks the new canon — 4,268 embeddings, 364 regenerated profiles, $0.2196 |
| M263 | `/taxonomy` reachable and navigable on a demo |
| M264 | The corpus tells the truth about the taxonomy |
| M265 | Proven live — 5/5 gate clauses cold, Playthrough suite **222/0** |

**0 platform-repo edits** across the release. Tooling shipped as rext `v2.9.10-rext` → `v2.9.17-rext`.

## Incidents

### P0 — the release nearly shipped a demo with an empty simulation library

Five milestones of taxonomy realignment closed green while the demo's AI-simulations library rendered
**zero cards**. Taxonomy and Directus content are separate snapshot surfaces: the taxonomy was swapped
wholesale, the content replayed unchanged, and **187 of 302** embedded node-ids no longer existed. One
non-null resolver field (`skills[].name`) turns a single dead id into a nulled list.

`/api/health` 200. 10/10 containers. `public.skills = 3562`. Library empty.

**Root cause of the near-miss:** nothing in v2.9 measured a *rendered surface* until M265. Row counts
and liveness probes are structurally blind to hollow success.

### P1 — five verifier-scope defects, one class

The release's dominant failure mode was **checkers narrower than their subjects**. A checker that is
too narrow does not under-report; it returns a **false answer**, which is worse than not checking.

| # | Verifier | Blind to | Consequence |
|---|---|---|---|
| 1 | `realign`'s column list (hand-maintained) | ids nested below the top level | repaired 4 columns, verified clean, still broken |
| 2 | `realign`'s anti-vacuity trigger | a cold stack's unprovisioned content schema | a successful 55,116-row replay recorded as FAILED |
| 3 | `realign`'s node-id pattern `[A-Z]{4,8}` | 19 of 3,562 ids (digit in the stem) | reported **clean** while the app still failed |
| 4 | `seed_role_guard`'s role regex | the `- role: X` list-item shape | would have gone **false-GREEN** |
| 5 | `claim_census_guard`'s `SOURCE_EXTS` | `.csv` — the canon's entire evidence base | debt reported in **18 files** that never incurred it |

Two more fences went red on **correct** changes because they pinned a *spelling* rather than a
property: `test_isolated_cold_proof_env_overrides` (a literal `${DEV_PGC:-…}`) and
`blocking_state_guard` (an M-id regex vs the documented between-milestones stub).

### P1 — the dev path could not migrate an additional stack

Three defects hiding each other: a silently-ignored positional `N`, a `::1` DSN (the *second* time
that hazard bit this release — its demo twin was fixed earlier and the dev twin never was), and a
clone-set assumption inverted from the demo side. Each presented three layers from its cause: no
migration → backend `Exited(1)` on an enforcer panic → no taxonomy → "the catalog is empty".

### P1 — a wrong diagnosis, retracted inside the milestone

`pt-assignment-assign` was declared not-taxonomy-caused and routed forward, from two true
measurements joined by a plausible story. The real cause was one hop further out. **Two true facts
and a plausible join are not a diagnosis.** Retracted in the record rather than overwritten.

## Cross-milestone patterns

1. **A fix applied where a defect was FOUND, not everywhere it COULD be.** The three retired role
   names were fixed in `presets/` at M262, resurfaced from `playthroughs/seed/` at M265, and
   resurfaced a **third** time from `seed-facts.ts`. Each fix shipped a fence with the same narrow
   scope as the fix.
2. **Verifier scope is the recurring bug, not product code.** Five of the release's defects were in
   things that check, not things that do.
3. **Silence as a failure mode.** Swallowed errors, `count()` without auto-wait, non-fatal warnings —
   each produced a green that meant nothing.

## ⚠️ Process finding — six of seven milestones shipped without the close lifecycle

`/developer-kit:close-milestone` ran for **M265 only**. M259–M264 have no `retro.md`, no
`metrics.json`, no Completeness/Gate ledger, and never got a deferral re-audit or adversarial review.
Their roadmap status read `planned` until this close flipped it from **git merge dates**.

Corroboration that the lifecycle was skipped rather than merely undocumented: **M261–M264 carry
3-line stub `decisions.md` with zero entries**, while M259 (6) and M260 (5) record theirs.

**Handled as option B (user decision):** the gap is recorded — in each milestone's roadmap block, in
`release-review.md`, and here — rather than back-filled. A retro written now for a review that never
happened would be indistinguishable from a real one to a future auditor, and the metrics would be
invented rather than measured. **v2.9's quality gate genuinely has this hole.**

## Carried forward

| Item | Destination |
|---|---|
| CI wiring — Phase 8b's triple-clean ran locally, not in CI | next release |
| `stack-core` carries ~89 failures / 2,438 tests, from the M257x fence suite v2.8 closed `closed-incomplete` **by user ruling** | next release |
| 4 guard-family REDs (`clone_drift`, `decommissioned_instruction`, `demo_knob`, `unreadable_repo_claim`) — all flag files v2.9 did not touch | next release |

## Metrics delta

| | v2.9 |
|---|---|
| Milestones | 7 (1 closed via the lifecycle) |
| Taxonomy | 43,584 / 22,511 → **3,562 / 706** |
| Playthrough suite | **222 passed / 0 failed** |
| Content refs repaired | **515 → 0 dangling** |
| Flake count | **0** (5/5 gate on every milestone-touched suite) |
| Platform-repo edits | **0** |
| Corpus files whose recorded debt fell (csv fix) | **18** |
| rext tags shipped | 8 (`v2.9.10-rext` → `v2.9.17-rext`) |
