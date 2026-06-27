---
iter: 02
milestone: M46
iteration_type: tik
status: closed-fixed
created: 2026-06-26
---

# M46 · iter-02 (tik #1) — auto-fill count (deliverable #1)

**Active strategy reference:** TOK-01 (`../decisions.md`) — build the three deliverables fixtures-first,
then prove on a real org. Tik #1 implements deliverable #1 (auto-fill count), the spine.

**Re-survey (Step 0):** TOK-01 named auto-fill count as iter-02's target. Confirmed still untouched: the
KB-fidelity audit verified `Batch.Count` is a fixed int with no Size-aware fill logic (`batch.go:32`). Still
the right next thing — everything else (per-story distribution, preview) fills an org-sized batch, so the
fill-to-Size capability must exist first.

**Cluster / target identified:** the `Batch` descriptor + `EffectiveBatches()` expansion. To fill a whole
org from one descriptor, a batch needs a way to say "fill the remaining N of this story" instead of a fixed
count.

**Hypothesis:** add a `Fill bool` (`fill: true`) to `Batch`; resolve its effective count at expansion time
= `Size − curated heroes − sibling fixed-count batches` (floored at 0), per story. Keep the resolution a
PURE function so the cache/$0-reseed invariant holds.

**Expected lift:** the structural capability "one descriptor fills a 220/500/1k org end-to-end" lands +
is unit-proven (the M42 semantic-gate metric on a generated org doesn't move until the real-run tik — this
is a build-toward-the-gate deliverable, not a metric-moving sweep).

**Phase plan:** code (batch.go fill resolution + blueprint.go mutual-exclusion validation) → fixtures-first
unit tests (fill-to-size, fill+fixed sibling, floor-at-0, multi-fill split, per-story fill, determinism,
validation, YAML parse) → full stack-seeding suite green.

**Escalation conditions:** none expected (offline, deterministic code). A non-pure fill (a count that
varies across reruns) would break the cache invariant → would need a redesign before close.

**Acceptable close-no-lift outcomes:** N/A — this is a deterministic code deliverable; it either lands
unit-proven (closed-fixed) or it doesn't.
