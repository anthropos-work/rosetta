**Type:** tik (#1, under TOK-01)

# M46 · iter-02 — auto-fill count

Implemented deliverable #1: the supporting-population auto-fill batch.

## What landed (rext `stack-seeding/`)
- `blueprint/batch.go`:
  - new `Batch.Fill bool` field (`fill: true`) — turns a batch into the org-scale auto-fill batch.
  - `resolveBatchCounts(story, batches)` — a PURE function computing each batch's effective count: a
    fixed-count batch keeps its `Count`; a fill batch expands to `Size − len(heroes) − Σ(fixed counts)`,
    floored at 0; multiple fill batches in one story split the remainder evenly (earliest take the rounding
    remainder) so the total is exactly the remaining slots.
  - `expand(...)` now takes the resolved `count` (was reading `b.Count`), so the fill count drives the
    per-member loop. The prompt expansion stays pure/deterministic → the cache/$0-reseed invariant holds.
- `blueprint/blueprint.go`: `validateBatches` now rejects a fill batch that also sets an explicit positive
  count (mutually exclusive — a fill batch's count is computed, not declared).

## Tests (fixtures-first, no key/cost) — `blueprint/batch_test.go`
8 new tests, all green + the full stack-seeding suite green:
- `FillToSize_Legacy` — Size 50, 1 hero → 49 generated.
- `FillWithFixedSibling` — Size 100, 1 hero, fixed=20 → fill 79 (total 99).
- `FillFloorAtZero` — over-declared story → 0 fill members (never negative).
- `MultipleFillSplit` — two fills split 49 as 25+24 (deterministic, order-stable).
- `FillPerStory` — multi-story: each story's fill fills ITS OWN org to ITS OWN Size (29 / 20).
- `FillDeterministic` — same descriptor → byte-identical mother prompts + count (the $0-reseed invariant).
- `ValidateBatches_FillWithCountRejected` — fill+count → validation error.
- `Parse_FillField` — `fill: true` parses from YAML (KnownFields) + expands.

## Re-measure
The gate's primary metric (the M42 semantic-coverage sweep on a generated org) is NOT measured this iter —
it moves only at the real-run gate-proving tik (TOK-01's tik #5). Progress here is a **structural
deliverable landed toward the gate**: the "fill a whole org from one descriptor" capability now exists +
is unit-proven. (Per `coverage-protocol.md`, "Progress" on a build-toward-the-gate iter is a landed
capability the gate depends on; the sweep delta is the final-tik measurement.)

## Close — 2026-06-26

**Outcome:** auto-fill count (deliverable #1) landed + unit-proven (8 new tests + full suite green). One
descriptor (`fill: true`) now fills a story to its Size, per-story, deterministically (cache invariant held).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the org-scale gate is proven on the real-run sweep, TOK-01 tik #5; this iter lands a
deliverable the gate depends on)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** (none beyond TOK-01)
**Side-deliverables:** none.
**Routes carried forward:** none (deliverables #2 per-story distribution, #3 preview, #4 fence + the #5
real-run gate-proving remain on TOK-01's plan).
**Lessons:** The fill count must be a PURE function of the story's Size + roster + sibling fixed-counts —
keeping it out of `expand`'s per-member loop (resolved once in `resolveBatchCounts`) preserves the cache
invariant and lets a fill batch see its siblings. A `Fill bool` + computed count is cleaner than a
sentinel `count: -1` (which would collide with the `count >= 0` validation and read ambiguously).
