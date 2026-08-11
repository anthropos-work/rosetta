---
milestone: M257x
iter: 32
iteration_type: tik
status: closed-fixed
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-32 — the binding run, and nothing else

## Phase 0 type-selection — graded explicitly, because the strict reading points at a tok

The last three tiks are iters 29, 30, 31. On the **headline** clause-2 number none of them moved it: it
has read `25 / 5 / 1` since iter-29 and has not been re-read since. Read strictly, that is a
3-consecutive-no-progress-tik streak and Phase 0 rule 2 would make iter-32 a **triggered tok**.

That reading is rejected, and the reason is recorded rather than assumed:

1. **The metric was not measured, which is not the same as not moving.** iter-30 and iter-31 each fixed a
   failing id and verified it green *live on a cold reset-to-seed*. What they declined to do was quote a
   scoped run as a clause-2 number — a prohibition this milestone wrote for itself at iter-25/iter-26
   after losing a measurement to exactly that. The full read costs ~40 minutes and was deliberately
   budgeted as its own iteration. A protocol that forbids cheap reads of its primary metric will always
   look stalled to a rule that counts re-reads.
2. **The skill's own stale-trigger clause mandates the same first action either way.** Phase 1 Step 0 for
   a triggered tok requires re-running the primary measurement *before* authoring a revised strategy,
   precisely so a strategy is not revised against a stale streak. That measurement **is** this iter's
   entire planned deliverable. Tok and tik therefore prescribe an identical next action; only the close
   differs.
3. So the honest grading is: **the trigger's precondition cannot be evaluated without the measurement.**
   Run it, then grade. If it returns `25 / 5 / 1` the streak is real and iter-33 is a triggered tok on
   confirmed evidence. If it returns anything better, the trigger was stale and this was a tik all along.

Recorded as a tik, with the tok question left open and decided by the number rather than by argument.

## Active strategy reference

`TOK-01: instrument first, then follow`. This iter is the *instrument* clause in its purest form: the
milestone has spent two iters landing fixes on scoped evidence and now reads the instrument that binds.

## Step 0 — re-survey

Checked at open rather than inherited from the hand-off:

- rosetta `m257x/platform-realignment` @ `8993cf9`, clean, 0 behind `main`.
- platform origin `2adcf71` — **unchanged**. Re-scope trigger stays at occurrence 1 of 2.
- rext pin, consumption clone and origin tag all `fast-build-m257x-iter-31b`.
- `stackseed` binary mtime `23:49:35` vs the pinned tag's commit `23:49:26` — the binary is **newer**, so
  the stack will reset with iter-31b's seeder and not with stale code. (This check exists because run 17
  discovered re-pinning the clone does not rebuild the per-stack binary.)
- `demo-1` up, 15 containers, carrying iter-31b's seed.

## Cluster / target identified

`MEASURE-M257x-iter31-clause2-binding-run`. Not a fix — a read. The two ids fixed since the last binding
measurement (`pt-workforce-funnel` at iter-30, `pt-workforce-succession` at iter-31) were each verified
green on cold-reset **scoped** evidence, which this milestone has ruled insufficient to move the gate
number.

## Prediction, recorded BEFORE the measurement

**`27 live / 3 failing / 1 TODO`.**

This was pre-registered at iter-31's close, *before* the confirming run existed. It is honoured here
verbatim: the run reports what it reports, and **a disagreement with 27/3/1 is the finding**, not an
embarrassment to be explained away. Per the iter-28 discipline, a number that BEATS the prediction
deserves more suspicion than one that meets it.

Expected remaining three, with what is already known about each:

- `pt-activity-drilldown` — fails at `activity-drilldown.spec.ts:113` on `heroRow.count() > 0`; an
  ordering coupling, not the role-text family. Best-evidenced next target.
- `pt-onboarding-hiring-candidate` — missing `/sim/…organizationId=` link on the hiring home.
- `pt-orgadmin-role-create` — 60 s `waitForURL` timeout after Save on the role-create drawer.

**Declared acceptable in advance**, so neither outcome can be re-framed after the fact:

- A result of `25 / 5 / 1` is a **complete iter**, not a failure. It would confirm the no-progress streak
  on evidence and hand iter-33 a properly-triggered tok — a more valuable result than a lucky lift.
- A **new** id failing that was green at iter-29 points first at iter-31's seed change, which altered the
  role distribution of every seeded org and is the widest-blast-radius change of the milestone. All five
  negative controls passed after it, but negative controls are not the suite.

## Phase plan

1. Confirm the world is the one the pins describe (done at open, above).
2. `./run-playthroughs.sh 1 --reset` — full, unscoped, from the pinned consumption clone. `nohup`'d, since
   a bare `&` in a tool call gets killed.
3. Heartbeat it — 209 serial specs over ~40 min is indistinguishable from a stall otherwise.
4. Compare by **sorted-id diff**, never by two summary lines (the iter-19 rule): removals *and* additions.
5. Grade the tok question against the number.

## Expected lift

`25 → 27` live. No fix lands in this iter by design.

## Escalation conditions

- Platform origin moves off `2adcf71` → re-scope trigger occurrence 2 → STOP, exit `re-scope-trigger`.
- A platform-repo edit required → route forward. Binding.
- The run does not complete inside the session budget → do **not** quote a partial run as a clause-2
  number (iter-25's own escalation condition, which fired once already).

## Acceptable close-no-lift outcomes

The run completing and returning `25 / 5 / 1` with an id-level diff — the measurement is the deliverable,
and a null result that confirms a tok trigger is a complete iteration.
