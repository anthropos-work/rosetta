# M260 — Progress

**Status: COMPLETE.** Closed 2026-08-14.

- [x] `MinRows` DERIVED from the capture rather than pinned at 40000 — under-capture protection preserved
- [x] re-ground every taxonomy-size assumption across `stack-snapshot`, `stack-seeding`, `stack-verify`
- [x] a fence that fails when a bare taxonomy count is re-pinned into source

## What replaced the pinned floor

`MinRows int` → `LoadBearing bool`, and one constant → **two rungs**:

1. **NON-EMPTY**, unconditional. A load-bearing table capturing zero rows is an unprovisioned or
   mis-filtered source, never a real catalogue. This is the case the original floor actually existed
   for, and it needs no magnitude at all.
2. **SHRINK**, measured against *this surface's own newest prior capture* (`capture.shrinkBaseline`,
   read once before the loop via the store that `Run` already receives). Below `shrinkRatio` (0.5) the
   capture **stops and asks**; `Options.AcceptShrink` carries the operator's written reason, and the
   manifest records it.

**It is strictly stronger than the constant it replaced.** A 42,790 → 42,000 partial capture passed
`MinRows: 40000` silently; it does not pass a comparison against 42,790.

**A first-ever capture has no history**, so the shrink rung cannot run — and the manifest says
`shrink rung DID NOT RUN` in its notes. A check that silently does not run reads exactly like a check
that passed, and this project has paid for that confusion more than once.

## The fence

`stack-core/taxonomy_pin_guard.py`, registered in the guard family. It flags a magnitude
(≥ 1,000) on a line that names a taxonomy subject **and** reads like a row-count threshold, in
non-test executable source. **GREEN over 313 files** — which is also the evidence for the
re-grounding item: no taxonomy magnitude is pinned anywhere in rext.

Comments, docstrings, disposition strings and test fixtures are deliberately NOT flagged: those
*record* measured figures, which is how this project stays honest. Only executable assertions expire.

## Tests

`stack-snapshot/capture`: 6 net-new tests (`capture_floor_m260_test.go`) plus 6 migrated from the
M209 floor suite — the boundary operator, the zero-row trigger, the not-load-bearing no-check
contract, the abort's loop position, and floor-before-leak-gate precedence all survive with the new
semantics. `stack-core/tests/test_taxonomy_pin_guard.py`: 11 controls. Whole `stack-snapshot` module
green.
