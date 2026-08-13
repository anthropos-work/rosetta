---
iter: 242
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-242 — run the WHOLE family twice and diff: is the family verdict reach-blind?

**Active strategy reference:** `TOK-08`.
**Route worked:** `ROUTE-M257x-236-host-is-the-unreliable-witness`, which iter-241 **narrowed rather than
closed** — it fixed one member and explicitly recorded that *"the other 30 family members were not audited
for the same shape."*

## Step 0 — re-survey

iter-241 proved one guard's green meant two different things on two clone sets. The obvious next question
is not *"which other guards read disk?"* (12 of them call `is_dir()`/`exists()`, which proves nothing about
their reach) but the empirical one:

> **Run the entire family twice — once against this box's 13-clone workspace, once against a clone set
> restricted to what a fresh bring-up creates — and diff the verdicts.**

No code reading, no inference. A member whose verdict text changes has host-dependent reach; a member whose
verdict is byte-identical while its subject includes clones is the dangerous case.

**And the sharpest form of the question is about the summary line itself.** Every iter close in this
milestone quotes *"N GREEN / 0 RED"* as its evidence of health. If that line is identical across two runs
whose actual coverage differs, then **the sentence this milestone has used as proof 40+ times is
reach-blind**, and iter-241's finding is not one guard's bug but the family's reporting contract.

## Pre-registered claims — SEALED IN THIS COMMIT

- **`P-242-1`.** **≥ 2** family members produce a materially different verdict line on the restricted
  clone set. **Predict ≥ 2.**
- **`P-242-2`.** **≥ 1** member exits **0 on both runs while demonstrably checking less on the restricted
  one, without saying so** — i.e. iter-241's defect, un-repaired, elsewhere in the family. **Predict ≥ 1.**
- **`P-242-3`.** **≥ 1** member changes verdict CLASS (green → red, or green → could-not-check) on the
  restricted set — i.e. a fresh box cannot run the family to the same colour. **Predict ≥ 1.**
- **`P-242-4`.** The family's own summary line is **identical** across both runs. **Predict: identical —
  the summary is reach-blind.**
- **`P-242-5`.** The restricted run completes without any member crashing (exit ≥ 3 / traceback).
  **Predict: no crashes** — the failure mode here is silence, not noise.

## Phase plan

1. Seal this pre-registration.
2. Build the restricted clone set (symlinks, no copies) and run `guard_family.py` against both.
3. Diff verdict-by-verdict; classify each difference as *disclosed* / *silent* / *class-change*.
4. Land the smallest change that makes the family's own summary carry its reach.

## Escalation conditions

If `P-242-4` holds, the finding is about **this milestone's evidence sentence**, not about a guard, and it
is written into the protocol doc as such.

## Acceptable close-no-lift outcomes

If every member either discloses its reach or is genuinely clone-set-independent, that is a complete iter
and a strong result: it would make iter-241 an isolated defect rather than a family-wide pattern.
