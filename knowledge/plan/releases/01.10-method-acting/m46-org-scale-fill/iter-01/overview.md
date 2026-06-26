---
iter: 01
milestone: M46
iteration_type: tok
tok_flavor: bootstrap
status: in-progress
created: 2026-06-26
---

# M46 · iter-01 — bootstrap tok (author the initial strategy)

The milestone's first iter. Per `build-mstone-iters` Phase 0 rule 1, iter-01 is a **bootstrap tok**: it
authors the FIRST strategy (TOK-01) from the milestone `overview.md` + `spec-notes.md` + the iteration
protocol (`coverage-protocol.md`), takes the baseline framing, and does NOT terminate the call — the loop
continues into iter-02 as the first tik under TOK-01.

## Inputs consumed
- `overview.md` — the gate (full-org fill from one supporting-population descriptor, the M42 semantic
  gate PASSES on the generated population, 0 hero collisions at scale, throughput+cost within budget) +
  the In/Out scope (auto-fill count, per-story distribution, preview/dry-run mode, throughput tuning +
  429 backoff, the `--gen-batches` opt-in fence).
- `corpus/ops/demo/ai-generation-spec.md` + `cache-spec.md` — the M45 engine + cache (the thing M46
  scales). Audited GREEN (Phase 0b).
- `corpus/ops/demo/coverage-protocol.md` — the M42 Playwright semantic-coverage harness (the iteration
  protocol; the gate measures the GENERATED org).
- `corpus/ops/demo/stories-spec.md` — the multi-org Stories model + the supporting-population fidelity
  contract the per-story distribution fills.
- The M45 code: `blueprint/batch.go`, `cmd/gen-batch/main.go`, `seeders/generated_batch.go`,
  `blueprint/stories.go`. The three M46 gaps confirmed (audit Phase 3): fixed `Count`, count-only
  `--dry-run`, hardcoded `stories[0]`.

## Baseline framing (Distance-to-gate)
- **M45 engine + cache + GeneratedBatchSeeder:** proven on a BOUNDED N=20 batch (tag
  `method-acting-m45-harden-final`). The seam (CODE owns structure/identity/closure, AI owns content;
  non-resolving names drop) holds at any N — it's the load-bearing invariant.
- **The gate metric (this milestone's primary metric):** the M42 semantic-coverage sweep's
  `(failingSections, personaFailures, escapes, notReachedPages, cross-port-follow failures)` — but
  measured on a **generated org** (the manager `/enterprise/members` table + org-scale surfaces must be
  full of believable generated members), PLUS the M45 generation sub-metrics carried forward at scale:
  **hero-collision count = 0**, **closure GREEN** (datadna), **cost ≤ ceiling**, **throughput ≤ budget**
  (~1k members ≤ a few min @ `--max-concurrent=5`).
- **Starting value:** the gate has NOT been measured on a generated org yet (M45 proved the engine, not a
  full org). The three In-scope code deliverables (auto-fill count, per-story distribution, preview) are
  not yet built. demo-3 is up on offset +30000 from M45's gate-proving; the harness is calibrated for
  both vantages.

## This iter's deliverable (bootstrap tok)
TOK-01 in the milestone-root `decisions.md` (the initial strategy + next-tik direction). No gate-metric
movement (toks don't move the gate); the deliverable is the strategy.
