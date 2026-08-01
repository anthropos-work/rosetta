---
iter: 26
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-26 — `MEASURE-M257x-iter26-clause2`: the re-measure, budgeted as the whole iteration

**Active strategy reference:** `TOK-01` — step 5, *"prove it cold"*, applied to clause 2.

## Step 0 — re-survey

Platform origin HEAD `2adcf71`, unchanged (re-scope occurrence stays 1 of 2).

**The re-survey changed the plan, which is what it is for.** iter-25 closed its second line as un-landed at
65 of 209 tests. The re-survey found the run had **completed** shortly after — 31 ptreport rows, gate verdict
emitted. So the iteration's work is not "run the suite" but **"validate and read the run that exists"**,
which is cheaper and strictly more honest than discarding a completed full run to produce a duplicate.

Two things had to be true before that number could be quoted, and both were checked rather than assumed:

1. **It must be a FULL run.** The ptreport gate is binding only on a full run
   (`run-playthroughs.sh:300-307`). 31 rows, no scoping flag, gate verdict present → yes.
2. **The stale roster must be immaterial.** iter-25 flagged this as the reason to re-run: the run's
   `stackseed --roster-export` failed (the second-pass defect), so the fake-FAPI served the *previous* seed's
   roster while the DB held a freshly reset-and-reseeded world. If hero ids had changed, every
   manager-vantage "the seeded hero is among the results" assert would fail for a reason that is not the
   product — and **four of the seven remaining failures are exactly that shape.**

   **Checked, not reasoned:** the roster's `pt-employee` entry carries
   `eid=23f24e3f-38fb-5027-9e07-2ef49a644af5`, and after the reset+reseed
   `select id,email from public.users where id='23f24e3f…'` returns that row,
   `pat.ellis1@pt-meridian-labs.com`. The seed is deterministic, so the stale roster is identical in the
   load-bearing field. Confound **defused by measurement.**

## Cluster / target identified

Read the completed run, `diff` its failing ids against iter-19's set, and attribute — or refuse to.

## Hypothesis

iter-24's Directus re-point clears the failures whose error text was the `directus_versions` 403.

## Expected lift

Not predicted (D-M257x-24-3 stands).

## Phase plan

1. Validate the run is full + the confound is dead (Step 0, done).
2. `diff` sorted failing ids against iter-19's ten.
3. Attribute only what the diff supports; route the residual with named handlers.

## Escalation conditions

- New failures appearing (a regression from iter-24/25) → investigate before claiming any lift.

## Acceptable close-no-lift outcomes

An unchanged failing set would be a real finding — it would refute the causal chain iter-24 proved at the
403 level, and that is worth more than a lift.
