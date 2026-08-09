---
iteration_type: tik
status: in-flight
active_strategy: TOK-08
---

# iter-208 — the milestone's own NOT-COVERED clause names ten Python sections; the repo derives five

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them)
— *census the mechanical classes; stop sampling them.* This iter applies it to the **scope disclaimer**
itself, which is a claim about a population and has never been graded against one.

## Step 0 — re-survey before targeting

iter-207 routed two new surveys, both pointing at the 314 newly-enumerated test-side `standing` figures.
Re-surveyed before committing: verifying members of that population is per-module work with no census
available, which is the sampling `TOK-08` replaced. **Substituted under the same strategy**, one level
up, exactly as iter-207 substituted.

The target is the sentence every harden pass and every run of this milestone has repeated verbatim:

> **NOT COVERED, stated rather than implied (`§5` rule 60):** the ten non-`stack-core` Python sections …

`§5` iter-178: **a NOT-REACHED clause is a MEASUREMENT or it is a mood.** That rule was written about
the tooling's clauses. It was never turned on the milestone's own.

## Cluster / target identified

`rosetta-extensions` has **11 sections**. The count of them that carry Python tests is not a judgement —
`suite_census` derives it from the collector's own glob (`PYTHON_TEST_GLOB = "test_*.py"`), and
`SECTIONS` is that derivation. Read at today's tree:

| | |
|---|---|
| sections total | **11** |
| sections carrying ≥1 Python test file | **5** — `demo-stack`, `dev-stack`, `stack-core`, `stack-injection`, `stack-verify` |
| non-`stack-core` Python sections | **4** |
| Python test files outside `stack-core` | **52** |

Six sections carry **zero** Python test files, so "unread" is not even the right word for them — there is
nothing there to read, and `LANGUAGE_EXCLUDED_SECTIONS` already says why for each (Go module / mixed
toolchain). The disclaimer counts them as unread Python.

**And there is a fence-shaped asymmetry underneath it.** `suite_census` derives `go_sections()` from
`go.mod` and `ts_sections()` from `e2e/playwright.config.ts`, each with an arm asserting the derivation
agrees with the hand-written registry. **There is no `python_sections()`** — the language the repo is
mostly written in is the one whose section list has no named derivation, so no prose claim about it has
anything to be graded against. That is why a wrong denominator could be repeated across passes without
a single fence noticing.

## Hypothesis

The disclaimer is a mood. Grading it produces the correct denominator, and the correct denominator is
**small enough to close**: 4 sections / 52 files is a reading that fits in one iter.

### Step 0, second pass — the "first verdict ever" half was FALSIFIED before the run, not after

The first draft of this plan said the four sections would get *"their first verdict."* Re-surveyed
against the milestone's own ledger before publishing it, and that is **false**:

- **iter-145** ran all four (*"the four never-run sections were RUN"*) and graded 21 failures.
- **Harden pass 35** ran all five sections in one table (`demo-stack` 9 failed · 1,038 passed · 11
  skipped; `dev-stack` 151 passed; `stack-injection` 335 passed; `stack-verify` 12 failed · 225 passed).
- A later pass ran four of them again (`demo-stack` 9 failed · 1,055 passed · 2 skipped, etc.).

So the clause is wrong a **second** way, and this one is worse than the denominator: **"NOT COVERED"
was true of the pass that wrote it and became a standing description of the milestone.** The sentence
handed to this run reads *"the other ten non-`stack-core` Python sections remain unread"* — and the
document that says so contains the tables proving they were read.

## Expected lift

1. The disclaimer's denominator is **derived**, and the derivation is fenced so the next disclaimer has
   something to cite.
2. The four non-`stack-core` Python sections get a **current** verdict — runner, section scope and
   language named on it, per `§5` r75/76 — replacing a mood with a reading.

## Phase plan

Two planned lines (declared, so the scope-creep tripwire counts against this shape):

1. **Run the four unread sections** — `/usr/bin/python3 -m pytest` (3.9.6), the repo's own `ENV_GATED`
   registry deselected by name so the reading is not silently gated. Started FIRST and with the rext
   tree frozen: nine runs on this milestone have been discarded as confounded by a mid-run edit.
2. **Ship `python_sections()`** — the missing third of the language triple — plus an arm asserting it
   agrees with `SECTIONS` and with the collector's glob, and an arm grading the *shape* of a
   not-covered claim: the unread-Python denominator must be derived, never restated.

## Escalation conditions

- **A failure in a never-before-run suite is a MEASUREMENT, not a regression from this iter.** No code
  of those sections is touched here. It is recorded, sized and routed — it is not the Phase 5 § 4
  test-gate RED, which is about the iter's own changes. Stated in advance so the grading cannot be
  argued afterwards.
- If a section cannot be collected at all, that is `§5`'s *a check is only as strong as the runs that
  reach it* and gets the same treatment: named, sized, routed.

## Acceptable close-no-lift outcomes

- The disclaimer turns out to be right and there really are ten unread Python sections → the derivation
  is wrong and that is the finding.
