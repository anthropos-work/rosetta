---
iter: 04
milestone: M46
iteration_type: tik
status: closed-fixed
created: 2026-06-26
---

# M46 · iter-04 (tik #3) — gen-batch preview / dry-run mode (deliverable #3)

**Active strategy reference:** TOK-01. Tik #3 = deliverable #3 (the preview / dry-run mode).

**Re-survey (Step 0):** TOK-01 named the preview mode. The KB-fidelity audit confirmed `--dry-run` is a
count-only summary (`main.go:110`), with no per-member render and no estimated-cost line. Still untouched
(iters 02/03 touched batch.go/blueprint.go/generated_batch.go, not main.go), still meaningful — an author
needs to inspect a large batch before committing a real run.

**Cluster / target identified:** `cmd/gen-batch`'s `--dry-run` block — it prints `"%d to generate, %d
cached"` and returns. No render of the prompts an author should review, no estimated cost.

**Hypothesis:** add `--preview` (implies `--dry-run` → no LLM, no key) that renders each expanded member's
routing context (story/batch/role/seniority/bias) + the full MOTHER prompt + any cached JSON + a per-member
estimated cost, plus a total estimated-cost line; `--preview-out <path>` writes the render to a file. Reuse
the existing `estimateUsage` + `CostOf` (the cost table) for the estimate. Values-blind (no key ever read).

**Expected lift:** the author-review surface lands + is unit-proven; the CLI dry-run IS the preview surface
(no GUI per the scope's Out list). (Gate metric on a generated org unmoved until the real-run tik.)

**Phase plan:** code (`--preview`/`--preview-out` flags + the implication + `renderPreview` in main.go) →
fixtures-first unit tests (renders prompts + cost with a panicking fixture proving no LLM; shows cached JSON
+ $0 for a cached batch; `--preview-out` writes a file; `--preview-out` requires `--preview`) → full suite
green + a real offline smoke vs the 20-member preset.

**Escalation conditions:** none expected (offline render, no key). A preview that echoed a key would be a
values-blind violation → must read only env, never echo (it reads no key at all — it never builds a client).

**Acceptable close-no-lift outcomes:** N/A — deterministic CLI deliverable.
