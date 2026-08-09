---
iter: 199
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-199 — a class that has been repaired by hand three times running is a class nobody is enumerating

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*
*"A reading SAMPLES; a fence CENSUSES."*

## Step 0 — re-survey (mandatory)

Re-surveyed at HEAD `1da7d2f` (iter-198's commit). `SURVEY-M257x-h46-…` closed there;
`FIX-M257x-h44-…` is deliberately its own work. That leaves:

> `SURVEY-M257x-h45-printed-measurement-literals-uncensused` — *"iter-193's `printed_arithmetic_totals`
> censuses a printed total assembled by ARITHMETIC over a registry. The strictly simpler sibling — a
> printed total that is a **LITERAL of a past measurement** — has no census at all, and this pass fixed
> its three instances **by hand, in one module**."*

**The re-survey changed its size.** The route was filed after harden pass 45's three hand-fixes. Since
then the class has produced a defect in **each of the two following work units**: harden pass 47's
`2,714` labelled as test FUNCTIONS in the very sentence that names units, and **iter-197's own**
*"the repo's other 121 modules are `TestCase`"*, which was stale before the iter that wrote it closed.

**Three hand-repairs in three consecutive work units is the exact signature `TOK-08` was written for.**
The route's own note says *"per the user's standing ruling the general census is routed, not built"* —
that ruling governs **harden passes**, which route rather than build new machinery. An iter is where the
building happens.

## Cluster / target identified

The mechanical class: **a printed count that is a hand-written numeric literal rather than a derived
value.** Decidable by AST + a measurement-noun test; no sentence needs interpreting.

## Hypothesis

Small and tractable. The naive form (any digit in a print) is dominated by identifiers — iter numbers,
rule numbers, HTTP codes, format widths — so the predicate must **classify** rather than filter, and the
excluded classes must be reported so the narrowing has a size.

## Expected lift

No `P`/`N` reading. Deliverable: the census, its population with per-class counts, every `literal`
repaired by derivation, and controls that fire in both directions.

## Phase plan

1. `printed_measurement_literals(root)` in `derivation_registry.py`, beside its arithmetic sibling.
2. Run it; classify every hit by hand to check the classifier against judgement.
3. Repair each `literal` **structurally** (`§5` r71 — derive the expectation from the same source the
   code derives from), not by re-typing a fresher number.
4. Fence: the finding class at zero, the excluded classes **non-empty** (an exclusion with no instances
   is unfalsifiable), plus fire/no-fire controls.

## Escalation conditions

- A `literal` that cannot be derived from anything → declare it with a reason rather than delete the
  check; do not widen the iter into building the missing derivation.
- Population > ~40 → report the number and repair the top module only, stating the residual.

## Acceptable close-no-lift outcomes

A census returning **zero `literal`** before any repair would be a first-class outcome — it would mean
pass 45/47 and iter-197 had already exhausted the class — **provided** the instrument is shown able to
fire on the founding shape.

## Explicitly NOT in scope

Comments and docstrings. The route says *printed*; a census of prose literals is a much larger and much
less mechanical population, and conflating them would make the zero unreadable.
