---
iter: 157
milestone: M257x
iteration_type: tik
status: closed-fixed
date: 2026-08-08
---

# iter-157 — the fence registry counted filenames while every claim about it said "declaration"

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

**Step 0 — re-survey.** iter-156 closed one queued census and opened
`SURVEY-M257x-iter156-other-reporting-layers`. That new route was graded first and **partly answered
without an iter**: `autoverify.sh:204` captures `verify.sh 2>&1` and derives a probe count from it, which
*looks* like iter-156's defect and is not — `verify.sh` writes its probe rows to **stderr by design**, so
there stderr **is** the subject's own voice. The property is *distinguish the subject's voice*, never
*do not merge*. Route left open for the remaining runners; not the target here.

Took `SURVEY-M257x-iter150-partition-completeness-elsewhere` instead — the older queued census, and the
one whose shape iter-156's finding sharpened.

**Cluster / target identified.** Declared partitions in the tooling: sets a mechanism classifies a domain
into, where nothing derives that the sets **cover** the domain. iter-150 repaired one
(`blocking_state_guard`'s `BLOCKING_FIELDS` / `NON_BLOCKING_FIELDS`); the route asked where else.

**Hypothesis.** The `FENCE_KIND` registry is the largest declared partition in the tree
(`postcondition` | `standalone`, 25 declarations). If the class is real, its completeness is unchecked in
at least one direction.

**Expected lift.** No `N` reading. The deliverable is the enumerated partition population with a stated
denominator, plus a fence in whichever direction is open.

**Escalation conditions.** A repair that changes which fences participate in the ratchet → land only if
the ratchet verdict is unchanged and demonstrated; otherwise route.

**Acceptable close-no-lift outcomes.** A partition census that returns zero with its instrument proven.
