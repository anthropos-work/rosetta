---
iter: 19
milestone: M256
iteration_type: tik
status: in-progress
opened: 2026-07-29
---

# iter-19 — the last sharpenable control, and the Playthrough it belonged to was vacuous

**Active strategy:** `TOK-01` move 4 — *"close the honesty items last, deliberately, not as leftovers."*
Negative controls are the honesty item with the most ground left: **22 of 25**.

## Cluster / target identified

Three Playthroughs are uncontrolled. **Two are `pt-studio-{advanced,guided}-generate` and must stay
untouched** — they sit behind `FIX-M256-studio-false-green`, and a control over a known false green would
*certify* the false green. That leaves exactly one: **`pt-hiring-recruiter-compare`**, priced at iter-15 as
"needs a same-vantage control whose absence half is unmeasured."

It has been open for four iters for a specific reason recorded in `negative-controls.spec.ts`'s own header:
the obvious contrast vantage — a Workforce-org manager on the hiring Results view — **ejects the browser to
production**, so its absence is true and meaningless.

## Hypothesis

The same move that discharged iter-13's three profile finals and iter-14's four workforce finals applies
once more: **the limit is the ASSERTION, not the Playthrough.** Its final is
`positionRows().count() > 0` — structural, and satisfied by any populated grid. Re-aim it at Org D's own
seeded facts and a contrast vantage should falsify it.

## Expected lift

Clause 2 negative controls **22 → 23 of 25**, which is the **terminal** value while the two studio
Playthroughs are correctly held back — so the floor in `mutation-class-fence.unit.spec.ts` can finally be
raised to it instead of trailing the count at 20.

## Phase plan

- **A — measure both vantages live** before writing an assertion: what seeded facts does the recruiter's
  scoreboard actually name, and what does a contrast vantage render?
- **B — sharpen the final, then build the control** against the sharpened version.
- **C — mutation-verify** every new assertion *and* the control itself.
- **D — 3 consecutive cold reset-to-seed runs**; restore the drifted cockpit fixture + sha-verify.

## Escalation conditions

- If no live, in-demo contrast vantage exists, this is the **last** sharpenable control and the gate's
  clause-2 control sub-clause tops out at 22 — a written disposition, not a silent gap.
- A vantage that produces a dead page is **not** evidence (iter-07 D29); the liveness floor is mandatory.

## Acceptable close-no-lift outcomes

A measured demonstration that the remaining control is unbuildable without a platform edit — which would be
the milestone's **first** `unimplementable` and must be escalated, not absorbed.

## Scope

One line: the hiring control. Onboarding's routed items (`ONBOARD-M256-seat-append` and the rest) stay
routed — iter-18 closed on them deliberately.
