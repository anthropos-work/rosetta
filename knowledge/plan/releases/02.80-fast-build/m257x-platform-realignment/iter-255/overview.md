---
iter: 255
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10
closed: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-249-fresh-checkout-hostile-tests
---

# iter-255 — the residual 12, each with its failure read first

**Active strategy reference:** `TOK-08`. iter-253 named the class, iter-254 took it **22 → 12**. This
continues the run-to-zero.

## Step 0 — re-survey (mandatory)

The residual was measured at the close of iter-254 and ships at
`iter-254/evidence/iter254-residual-after-repair.txt`: **12 node-ids across 7 files**, reproducible on the
frozen pair in ~45 s. Nothing has moved under it — iter-254 was the last thing to touch the tree.

## Cluster / target identified

The largest coherent sub-cluster is the **resolution/denominator collapse** group: every member reports a
population that shrank rather than a claim that is wrong — *"resolution collapsed — the zero is vacuous"*
(136 against a floor of 200), *"the widened denominator collapsed"* (30 against 150), *"resolves in
neither pool and is undeclared"*. That is 6 of the 12, in `test_anchor_subject_census_m257x` (4) and
`test_m257x_corpus_file_citations` (2).

## Hypothesis

These 6 collapse because the corpus's anchors and markdown citations resolve **into the platform clones**,
which are git-ignored — the same precondition iter-254 established, applied to arms whose symptom is a
shrunken denominator rather than a named missing target. `D-M257x-254-4` still governs: each failure is
**read** before a declaration is written, and anything that is not environmental leaves the class.

## Expected lift

The class **12 → ≤ 6**, verified in both directions as iter-254 verified its own: frozen skips with the
precondition named, live still runs and passes.

## Phase plan

- **A** — read each of the 6 failures on the frozen tree; confirm the clone set is the cause and not a
  proxy for something else.
- **B** — declare, at the grain that fails.
- **C** — verify both directions.
- **D** — close.

## Pre-registrations (sealed in this iter's first commit)

| # | claim | prediction |
|---|---|---|
| **PR-1** | all 12 residual members need only the two preconditions already established (no third appears) | **false** — a third cause or a non-environmental member will surface |
| **PR-2** | ≥ 1 of the 12 resolves as a cascade, with no edit of its own | **true** |
| **PR-3** | the class reaches **0** this iter | **false** — 12 heterogeneous members is more than one iter of honest reading |
| **PR-4** | no arm repaired this iter becomes a live SKIP (iter-254's `PR-5` defect does not recur) | **true** |
| **PR-5** | ≥ 1 of the 12 is a REAL defect that iter-253's control mis-classified as environmental | **false** — the control was per-node-id and both-tree |

`PR-1` and `PR-3` both predict against a clean sweep; `PR-5` predicts against finding a flattering defect
in a prior iter's instrument.

## Escalation conditions

- A member with no declarable precondition leaves the class, is named, and routes out — **do not decorate
  a defect**.
- A member whose repair would require changing what the test asserts is routed, not rewritten: this route
  declares preconditions; it does not re-scope fences.

## Acceptable close-no-lift outcomes

- The 6 turn out to need a change of assertion rather than a declaration — characterised and routed, with
  the falsification as the deliverable.
