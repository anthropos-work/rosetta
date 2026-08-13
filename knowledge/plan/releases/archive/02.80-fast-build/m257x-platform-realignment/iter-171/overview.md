---
iter: 171
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-171 — the bind that waits on a name nobody reads

## Step 0 — re-survey before targeting

`TOK-08`'s standing direction is *census the mechanical classes; stop sampling them*. iter-170 closed the
RED-at-HEAD census with **one** disagreement it could not explain by imports:

> `SURVEY-M257x-iter170-cockpit-runner-dependence` — two `test_cockpit` server-binding tests pass under the
> fleet runner (3.9.6 + pytest) and fail under 3.14/unittest. *"Either a runner-dependent harness assumption
> or a real 3.14 behaviour difference; unresolved, and it is the only disagreement not explained by imports."*

Re-surveyed at open: still open, still the only one. Nothing absorbed it. **Target confirmed.**

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* This iter is
`TOK-08`-shaped in method, not in class: a single unexplained instance is root-caused, and then the
**property behind it is censused across the whole section** rather than the one instance repaired.

## Cluster / target identified

iter-170's own lesson was **name the runner, or the suite verdict has an unstated scope** — and it left one
disagreement standing whose cause was unknown. An unexplained runner disagreement is corrosive precisely
because the rule iter-170 just earned depends on runner differences being *understood*: "it's a runner
artifact" is a shrug, not a finding. `§5` rule 60 says a scoped RED is evidence about its scope alone — so
before this RED can be dismissed as harness noise, it must be **explained**.

## Hypothesis

The disagreement is **not** in the harness. It is a blocking call in **shipped code** that one interpreter's
resolver answers in milliseconds and the other's does not — and if so, the failing 2-second bind window is a
*symptom*, and the population of code that pays the same call is the real subject.

## Expected lift

No `P`/`N` reading is taken this iter (`TOK-08` puts the reading after the sweep; `§9`'s
**UNMEASURED-is-not-unmoved** refinement applies and the close will say so in those words).

The iter's deliverable, in `TOK-08`'s mandated report shape:

1. the **root cause**, measured, with both interpreters named and both numbers stated (`§8`);
2. the **enumerated population** of the class inside `rosetta-extensions` — every site, with the
   denominator stated and the exclusions justified **by property, not by spelling** (`§5` rules 70/71);
3. the repair applied **to the population, not to the two failing tests**;
4. a **fence** with a mutation control and an anti-vacuity control that both fire (`§9`);
5. the **production consequence** stated honestly, including whether this box's number generalises.

## Phase plan

(A) reproduce + root-cause under both runners, with a stack dump rather than a guess. (B) census the class
across the section; state the denominator. (C) repair the population. (D) fence it, both controls shown
firing. (E) re-run the affected suites under **both** runners and name them. (F) close.

## Escalation conditions

- A repair that would need a **platform-repo edit** → route forward, do not edit (v2.8 standing constraint).
- A repair that **weakens** what the existing tests assert → per iter-158, the proposed repair is a
  hypothesis; show the assertions survive, or route it instead of landing it.
- If the root cause turns out to be genuinely a harness assumption with no shipped-code consequence, say so
  plainly and close as a falsification rather than manufacturing a repair.

## Acceptable close-no-lift outcomes

The census returning a population of **one** (the cockpit alone) is a finding, not a failure — but per
`TOK-07` rule 2's guard-rail, a multiplier near 1.0× is not proof the class is rare, and the close must say
which it is.
