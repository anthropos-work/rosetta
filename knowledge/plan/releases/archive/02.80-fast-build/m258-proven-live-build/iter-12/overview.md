---
iter: 12
milestone: M258
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-08-12
---

# iter-12 — TOK-02: space as a goal

**Type:** tok · **Flavor:** bootstrap (user-directed) · **Authors:** `TOK-02`

## Why this tok is `bootstrap`-flavored and not `triggered`

Neither of the skill's two automatic tok rules fired:

- **Rule 1 (iter-01)** — N/A; `TOK-01` was authored at iter-01.
- **Rule 2 (3-no-prog streak)** — did NOT fire. The last three tiks (iter-09, iter-10, iter-11) all
  closed `closed-fixed` with measurable deliverables.

This tok fires on a **third cause the skill does not enumerate: the user added a GOAL.** `D57` records
it — space optimisation of `up`/`down` is now a goal of M258, not a route inside it. A new goal with no
prior strategy is the *bootstrap* situation exactly, and the bootstrap semantics are the correct ones:
it **authors a first strategy** (there is nothing to revise) and it **does not terminate the call**
(the user directed the work; making him re-invoke to start the first tik adds friction and no review
value). Recorded as a deliberate, named extension of rule 1 rather than a mis-classification of rule 2 —
grading a user-added goal as a "3-no-prog stall" would put a false stall on the strategy chain.

**Strategy class:** `new-direction` — there is no prior space strategy to be more-aggressive or
more-granular than. `TOK-01` (time) is **not superseded**; it is complete-by-ruling and `TOK-02` runs
beside it under the same milestone.

## Phase plan

1. Re-survey the space picture across **every** axis, verifying iter-11's figures rather than
   inheriting them.
2. Find the axes iter-11 did not measure at all.
3. Author `TOK-02` in the milestone-root `decisions.md`.
4. Set the next-tik direction.

## Acceptable close outcomes

A tok closes on the strategy landing, not on a metric. `closed-fixed` iff `TOK-02` is authored with
every class argued on **both** axes with measurements, per the user's binding constraint (`D58`).
