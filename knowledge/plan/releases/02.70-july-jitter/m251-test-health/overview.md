---
milestone_shape: section
milestone: M251
title: "test-health (realizes the reserved v2.6->v2.7 carry)"
status: planned
release: v2.7 "july jitter"
depends_on: [M246]
parallel_with: [M247, M248, M249, M250, M252]
complexity: small
created: 2026-07-23
---

# M251 — test-health

## Goal
The standing demo-stack test debt is discharged — the mechanical failures re-pointed, the `run-unit.sh`
roster fixed.

## Shape (why this shape)
`section`. The work is a fixed, enumerated checklist of already-understood mechanical re-points: two unit
specs to add to a roster, and ~6 python assertions to re-aim at deliberately-changed behaviour. There is no
measure→patch→re-measure loop and nothing exploratory — the composition is confirmed against disk, so a
section (do-each-and-tick-it) shape fits rather than an iterative gate. This milestone **realizes the reserved
v2.6→v2.7 carry** (the test-health debt v2.6 shipped forward).

## Scope
### In
- **run-unit roster fix** — Add `content-denominator.unit.spec.ts` + `run-discrete.unit.spec.ts` to the
  `UNIT_SPECS` roster in `stack-verify/e2e/run-unit.sh`. This clears the `UnitSpecsAreExecuted` guard, which is
  currently RED / runner exit 2.
- **Mechanical python re-points** — Re-point the ~6 mechanical `test_cockpit` academy/overlay assertions +
  the `test_public_host` port-13001 assertion at the **deliberately-changed M218/M238/M220 behaviour** (the
  tests fail because the behaviour they assert was intentionally moved, not because anything regressed).

### Out
- The ~2 docker/live-gated tests (`test_purge` + a live-serve assertion) — they ride the **M254** closer
  (they need a live box).

## Dependencies & parallelism
- **depends_on:** `M246` — the re-sync & re-point HARD barrier. All fan-out worktrees branch from post-M246 HEAD.
- **parallel_with:** `M247`, `M248`, `M249`, `M250`, `M252` — the fan-out siblings off the M246 barrier.
- **Intra-milestone lanes (~1×):** 2 disjoint lanes — the **run-unit roster** lane ∥ the **python re-point**
  lane — but tiny. No serial bottleneck of substance. **Recommended subagents: 1** — a single agent is the
  pragmatic default.
- **Merge/close order:** M251 closes **first** in the release (M251 → { M248, M250 } → M249 → M253 → M252 →
  M247-reconcile → M254).
- **Coordination guardrails (files this milestone owns):**
  - `run-unit.sh` roster — **M251 owns it.** If M248/M252 add a `*.unit.spec.ts`, it must be rostered here
    (coordinate the line).
  - `demo-stack/tests/*.py` — **M251 owns the health/inventory tests**; M249 owns the *patch* tests. No overlap.
  - Rung-zero: any rext tag a stack consumes must be pushed to **origin** first.

## KB dependencies
- `corpus/ops/verification.md`
- `corpus/ops/demo/coverage-protocol.md`

## Delivers
- (optional) a `corpus/ops/verification.md` anchor indexing the demo-stack python suite + the run-unit roster.

## Open questions
- None blocking — the composition is confirmed against disk.
