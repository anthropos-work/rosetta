---
iter: 267
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-261-succession-projection-is-empty
---

# iter-267 — discriminate the one failing Playthrough, without spending the control

**Type:** tik, under `TOK-08`. Handler `FIX-M257x-261-discriminate-succession-empty`.

## Step 0 — re-survey (mandatory, before targeting)

iter-261 ran clause 2 for the first time: **29/31, one failing Playthrough** —
`workforce-intelligence.talent-pool.UC1` (`@pt:pt-workforce-succession`). The page renders its chrome
(heading, org, manager, coverage cards) and **both projection tables render `img "No data"`**, while its
three siblings pass on the same login, org, seed and run. Exactly one computed projection returns nothing.

iter-261 named two candidate causes and **chose neither**, because discriminating them needed *"a pin bump
(which would destroy the control) or an `app` source read"* — and no `app` clone at current `main` existed
on this box at the time.

**That precondition changed at iter-262.** `stack-dev/app` is now cloned at **`3eaadae68`** = current
`main`, and `internal/workforce/succession.go` exists in it. **The source read is now free**, so the
discrimination can be made with **no pin bump**, leaving `D-M257x-258-1`'s frozen-pin control intact.

## Cluster / target identified

`ROUTE-M257x-261-succession-projection-is-empty` — the milestone's only measured functional-regression
candidate, and the one thing standing between clause 2 and MET besides the 29-vs-30 count.

## Hypothesis

The projection's selection predicate demands an artifact the frozen-pin `pt-world` seed does not produce —
**seed-contract drift**, the same class as skiller/skillpath/jobsimulation, one layer out. If instead the
predicate is unchanged and the data is present, the finding inverts into a **product regression**, which
is a materially bigger result and must be reported as such rather than smoothed.

## Expected lift

Clause 2's single failure gets a **named cause** instead of two candidates. That is the deliverable; a
*fix* may not be in this iter's reach and is not promised.

## Phase plan

1. Seal pre-registrations (first commit).
2. Read `internal/workforce/succession.go` at `3eaadae68`: what does the projection select on?
3. Date the file — was it changed inside the window the frozen pin does not cover?
4. Compare against what `pt-world` seeds for that org (rext `playthroughs/` + `stack-seeding`).
5. Name the cause; repair the corpus/tooling side if it is reachable without a pin bump; route otherwise.

## Escalation conditions

- **No pin bump, no rext tag, no stack mutation.** If the discrimination needs any of them, stop and route
  — the frozen-pin control is worth more than one iter's certainty.
- If the cause is a **product regression**, do not repair it: v2.8 forbids platform edits. Report it, cite
  it, and route it to the platform team's surface.

## Acceptable close-no-lift outcomes

A documented falsification of PR-4 — *the two causes are not separable from source alone* — closes the
iter honestly and tells the next one exactly what measurement is required.
