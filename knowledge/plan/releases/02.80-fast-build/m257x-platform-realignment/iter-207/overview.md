---
iteration_type: tik
status: in-flight
active_strategy: TOK-08
---

# iter-207 — the three literal censuses skip every test module, and the skipped population is the larger one

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* A reading SAMPLES; a fence CENSUSES. Work the
classes in descending measured size, and **state the denominator**.

## Step 0 — re-survey before targeting

`TOK-08`'s standing next-target after iter-206 is *"the other 140 standing figures are in modules with no
fence"* (the route reads **161** in iter-206's own close; harden pass 50 re-derived the class at **147
sized / 7 derived / 140 unverified**, so the route's numerator is one of the figures this milestone has
already had to correct). Re-surveyed at today's tree before committing to it, and **the target is
substituted under the same strategy**:

> The route asks to verify figures inside the census's declared population. Re-surveying the census's
> **population predicate** first — `§5`, *a fence's POPULATION is a registry too* — shows all three
> literal censuses share one line:
>
> ```python
> if any(part in _CENSUS_SKIP for part in path.parts) or path.name.startswith("test_"):
> ```
>
> **Every `test_*.py` module in the repo is excluded from all three at once**, and the exclusion is
> stated as a rationale in two docstrings (*"a test stating `4 of 7` is stating an expectation on
> purpose, which is the opposite defect"*) — never as a **size**.

Sizing the exclusion is the whole iter. Verifying 140 figures inside a population is worth less than
finding out the population was drawn with a hole in it, and the hole is measurable in one pass.

**Substitution recorded:** `TOK-08` named *the standing figures by module*; the re-survey substitutes
*the censuses' own file-grain exclusion* under the same strategy — same class (mechanical measurement
literals), one level up (its denominator rather than its members).

## Cluster / target identified

The exclusion is `path.name.startswith("test_")`, applied **at file grain**, in all three of
`printed_measurement_literals`, `docstring_measurement_literals` and `comment_measurement_literals`
(plus `noun_vocabulary_reach`, `classifier_window_miss_rate`, `printed_arithmetic_totals` and
`census_scope` — every consumer of the same predicate).

The rationale is right about **assertions** and wrong about **prose**. A test module contains both:

- `self.assertEqual(len(rows), 7)` — an expectation on purpose. Asserted; goes RED when wrong.
- `"""Measured at iter-N: 15 classes, 76 tests across 6 modules."""` — prose. Asserted by nothing,
  read as evidence, and it rots exactly like the production prose the censuses exist to catch.

`§5`, *a membership check cannot see a hole inside a member*: a **file-grain** exclusion cannot
distinguish the two, so it drops both. And `§5`, *a CORRECT exclusion is still a defect while it is
silent* — this one has never been sized.

## Hypothesis

The excluded population is **not** small relative to the censused one, and it contains live stale
figures. If it is small the finding is a cheap negative and the exclusion gets a size instead of a
rationale; if it is large the class `TOK-08` has been working for 6 iters has been measured over
roughly half its own subject.

## Expected lift

- The exclusion is **sized** rather than argued, corpus-wide, across all sections.
- A **pre-registered derivable subset** of the excluded `standing` rows is derived at today's tree,
  and whatever is stale is repaired — de-literalised where the value is computable
  (iter-206's rule: *de-literalise beats re-pin*).

### Pre-registered derivable subset — named BEFORE deriving any of them

Six rows, chosen for naming a population this repo's own code computes, one per module so a single
module's habits cannot dominate:

| # | site | figure |
|---|---|---|
| 1 | `stack-core/tests/test_suite_census_population.py:90` | `11 sections` |
| 2 | `stack-core/tests/test_suite_census_population.py:432` | `424 tests` |
| 3 | `stack-core/tests/test_suite_census_population.py:586` | `510 subtests` |
| 4 | `stack-core/tests/test_fence_registry_completeness_m257x.py:2` | `25 modules` |
| 5 | `stack-core/tests/test_test_collection_fence.py:286` | `2837 tests` |
| 6 | `stack-core/tests/test_story_org_count_guard.py:252` | `164 files` |

The subset is **selection-biased toward derivability** exactly as iter-206's was, and the hit rate is
therefore **not** an estimate for the rest. That is stated here, before the reading, so it cannot be
read off afterwards.

## Phase plan

Two planned lines (declared, so the scope-creep tripwire counts against this shape, not against a
single-target tik):

1. **Size + enumerate.** Ship `test_module_measurement_literals(root)` in
   `stack-core/derivation_registry.py` — the same shared classifier, the same row grammar, over the
   population the three siblings drop. Ratchet it. Print the section spread, because the exclusion is
   repo-wide and not a `stack-core` fact.
2. **Verify the pre-registered subset** and repair what is stale.

Arms in `stack-core/tests/test_frozen_expectation_census_m257x.py`: an anti-vacuity control written
against the census's **subject** (`§5` iter-94), a **direction** arm (the file-grain predicate must
over-drop, never under-drop), and a mutation control that actually fires under `§5` r77 (mtime bump —
size-preserving edits are invisible to pytest) and under iter-200's memoisation hazard.

## Escalation conditions

- If the excluded population is **zero**, `§9` applies in full: the census must prove its instrument
  before the zero is reported, and *a good repair can destroy the proof the instrument fires*.
- If repairing a subset row requires changing an **asserted** expectation (not prose), stop and route:
  that is the exclusion's correct half and is out of this iter's scope.

## Acceptable close-no-lift outcomes

- The excluded population is small and clean → the exclusion earns a measured size and a fence that
  keeps it that way. That is a complete iter under `§5`'s *sized, not argued* rule.
