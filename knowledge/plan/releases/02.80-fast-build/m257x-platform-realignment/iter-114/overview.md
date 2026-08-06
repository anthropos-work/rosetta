---
iter: 114
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: in-progress
opened: 2026-08-07
---

# iter-114 — the reach metric names where its denominator came from, or prints no percentage

**Active strategy reference:** [`TOK-07`](../decisions.md#tok-07-enumerate-the-predicate-not-the-anchor--2026-08-06)
— **step 2, first half.** `TOK-07` rule 4, quoted because this iter is nothing but its implementation:

> **Grade reach against the ENUMERATED set, never the detected union.** iter-108's 100 % was against the
> wrong denominator, which is precisely why it read as success while twins survived. **This is
> `fence_provenance`'s defect wearing different clothes** — *a check reporting a state it did not
> measure* — and it takes the same fix: **the denominator must state where it came from.**
> `repair_reach_guard`'s output must name its site set as corpus-derived-per-predicate, and a run whose
> denominator came from a `raw/` ledger dir must be unable to print a reach percentage at all.

**Step 0 — re-survey (done).** iter-113 closed the blocker and produced the enumerated denominator: **71
sites over 24 predicates**, checked in at `iter-113/enumeration.json`, exit 0, provenance stamp naming a
clean rext tree (`af7d6e9ba`). `repair_reach_guard` at `af7d6e9` still takes **only** `--ledger <dir>` —
a directory of `### B<n>` seat reports, i.e. *a prior reading's detections* — and prints
`reach t/N = P%` unconditionally. That is exactly the instrument that graded iter-108 **46/46 = 100 %**
while the same propositions stood one file away. The target is untouched and still meaningful.

**Phase 0d pre-flight (RUN).** `repair_reach_guard.py --help` executes at the committed tree and its
existing test file is 309 lines with a known-answer fixture (iter-81's repair vs iter-76's reports,
answer key 36 unreached of 145). The pipeline is live, so the change can be graded against a fixture
whose answer nobody can tune.

## Cluster / target identified

The repair cannot be graded before the guard can accept the enumerated set — so the guard comes first, in
dependency order. `TOK-07` step 2 names both halves in one sentence ("*repair whole predicates … 
`repair_reach_guard` graded on the enumerated denominator with its provenance printed*"); this iter lands
the second clause, iter-115 lands the first.

## Hypothesis

Two changes, and the second is the one with teeth:

1. **`--enumeration <file.json>`** — accept iter-113's corpus-derived, per-predicate site set as the
   denominator, and tag the report with **`denominator: corpus-derived-per-predicate`**.
2. **A `raw/`-derived denominator may not print a percentage at all.** Not a warning appended to the
   number — the number must be **absent**, because a percentage is the thing that gets quoted in a close
   and a caveat beside it is not.

## Expected lift

`P` does not move — no reading is taken (`TOK-07` reads **last**). §9's refinement applies: `P` is
**UNMEASURED**, not unmoved. What must be true at close:

- the guard accepts the enumerated set and prints its provenance;
- a `raw/`-dir run prints **no** reach percentage, proven by an assertion on the absence of `%`;
- the iter-81/76 known-answer fixture still returns **36 unreached of 145** — the control that proves the
  change did not soften the fence it was extending.

## Phase plan

Two planned steps (tooling shape): ship the denominator provenance, then re-grade the existing fixture to
show the fence's old behaviour intact.

## Escalation conditions

- **User-blocker** only per the protocol's list (test gates RED, unrelated-suite regression).
- **Route forward** if the guard needs a redesign rather than an input: report and stop, do not improvise.

## Acceptable close-no-lift outcomes

- The `raw/` refusal turns out to break a live caller that legitimately needs the old shape. Then the
  finding is the coupling, recorded, and the refusal lands behind the caller's migration.
