---
iter: 02
milestone: M2c
iteration_type: tik
status: closed-fixed
created: 2026-06-03
---

# M2c / iter-02 — tik: author the `clerk-express-1.json` DNA

**Type:** tik · **Active strategy:** TOK-01 (RS256-native, additive-first, real-SDK runner)

## Step 0 — re-survey
Metric is 0% (no DNA exists). TOK-01's next-tik direction = author the DNA. Target current + meaningful. No substitution.

## Cluster / target
The measurement target itself: the 3rd Alignment DNA. The gate scores against it, so it's the prerequisite
for every later tik. Genes from `spec-notes.md`'s proposed table, in the existing `capabilities[]/variants[]`
format (matching `clerk-js-5.json`).

## Hypothesis
A structurally-valid DNA (passes `alignctl dna validate`) establishes the denominator + the runner's
input contract (the `op` codes the Node runner will implement). It does not move the alignment score yet
(no mirror/goldens), but it's the planned deliverable of this tik.

## Expected lift
DNA authored + `alignctl dna validate` passes (9 genes / 4 capabilities / 7 critical). Score stays 0% by
design until the seam + goldens land (iter-03+).

## Phase plan (alignment protocol)
Author the DNA JSON → `alignctl dna validate` → `alignctl dna list` to confirm the gene roster.

## Acceptable close-no-lift
n/a — the DNA is a concrete deliverable; closes-fixed when it validates.
