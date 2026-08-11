---
iter: 73
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
---

# iter-73 — the 142 come inside the fence, and six of them were stale

**Active strategy reference:** `TOK-05`, step 1 (**fence**), consuming
`FENCE-M257x-iter72-bare-citation-reach` **one iteration after iter-72 opened it** — the same
prove-the-rule-by-using-it cadence iter-67 used on iter-66's route.

## Step 0 — re-survey before targeting, and a dry run before landing

iter-72 routed the fence with a design and two mechanical proofs but **no finding count**, and
iter-70 had just spent itself demonstrating that a routed count is a hypothesis. So the first act
here was a **dry run of the widened reach against the live corpus, with the guard untouched**:

| | n |
|---|---|
| newly resolvable under the widened rule | **136** |
| still unresolvable | 92 |
| **findings the widening would raise** | **12** |

Twelve is small enough to land **and repair in the same iteration** — Fate 1 — so the iter is
scoped to both rather than to the fence alone.

## Cluster / target identified

Land the two halves iter-72 proved missing: the regex that never admitted a bare
`<name>.<ext>:N`, and the service-doc resolver that mapped `backend.md` to a `stack-demo/backend/`
that does not exist.

## Hypothesis

The widening reaches ~136 citations, raises a small and coherent finding set, and the findings are
**stale platform-realignment defects** rather than false positives — which is the whole reason this
milestone cares about reach.

## Expected lift

- `anchor_construct_guard` reach **124 → ~177** anchors.
- The findings repaired, guard GREEN, `CITE_REF=worktree` still discriminating.
- No new false-positive class — specifically, **no port may become a citation**.

## Phase plan

- **A** — dry-run the widened reach with the guard untouched (**done at open**).
- **B** — land the regex alternative + the compose-derived service→repo map.
- **C** — read the findings and repair each against the platform artifact.
- **D** — RED-before-trusted: mutants over every decision, plus a true no-op control.
- **E** — gates.

## Escalation conditions

If the widening raises a port as a citation, it does not ship — that is the failure mode this
guard's own docstring records as *"134 findings, essentially all of them ports."*

## Acceptable close-no-lift outcomes

A finding count too large to repair in-iter would close as *"widening measured, not landed"*, with
the count as the deliverable. The dry run at open is what makes that decision cheap rather than
discovered halfway.
