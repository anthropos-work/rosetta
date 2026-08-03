---
iter: 51
milestone: M257x
iteration_type: tok
tok_flavor: triggered
status: in-progress
opened: 2026-08-03
---

# iter-51 — TRIGGERED TOK: the binding constraint was never the instrument's sharpness

**Trigger:** three consecutive tiks with no measurable progress on the primary metric — **iters 48, 49,
50**. The floor fires; nothing suppresses it.

## Step 0 — re-survey before authoring (mandatory)

The tok's own Step 0 exists to catch a **stale** trigger: if the metric had moved since the last tik's
close, the streak would be an artefact and the honest close would be `closed-no-lift`. It has not.

- Platform origin re-fetched at open: **`2adcf714`, unchanged.** Re-scope trigger stays at occurrence 1
  of 2.
- **The corpus is unrepaired.** Nothing has been fixed since iter-47's repair; the union of readings #9
  and #10 stands at **18 anchored blockers**, and the residual is estimated at **~23 and biased low**.
- The three tiks in the window, each measured, each honestly no-prog:
  - **iter-48** — reading #8 returns 12; iter-47's "zero pre-existing" refuted. Nothing repaired at close.
  - **iter-49** — two named fences shipped and proven, the twelve repaired, reading #9 returns **14**
    against a pre-registration of 6. The induced term rose 2 → 7.
  - **iter-50** — the paired variance experiment. Nothing repaired **by design**.

The streak is real. Proceed.

## Why the streak is not the usual kind

Two of the three tiks delivered exactly what they set out to deliver. iter-49 built the two instruments
iter-48 named, watched both go RED before trusting either, proved both with inversion mutants and no-op
controls — **and the number went up.** iter-50 ran the one experiment the protocol had prescribed and
never performed — **and the number went down without a single repair.**

**That is the signature of a metric that is not measuring what the strategy assumes it measures**, and
iter-50 finally measured why.

## What this tok concludes, in one line

**TOK-02 was optimising the wrong term.** Two strategies have now been spent sharpening the *reading*; the
measurement says the reading was never the binding constraint. **The repair's COVERAGE is.**

The full revised strategy is recorded as
[`TOK-03`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03)
in the milestone-root `decisions.md`. This file records the tok; that entry is the strategy every
following tik must name.

## What this tok does NOT do

- **It does not re-cut clause 5**, narrow it, or read "met" any other way. The user has ruled twice. Clause
  5 is met by a reading that returns zero blockers and by nothing else. TOK-03 changes what happens
  *before* that reading is taken, never what the reading must return.
- **It does not propose closing at 4 of 5**, and does not defer the residual to a future milestone.
- **It does not weaken the audit instrument.** The frozen instrument stays frozen; TOK-03 runs *two* of
  it rather than a cheaper one.
- **It does not discard TOK-01 or TOK-02.** Every fence they built is kept and still holds clauses 1–4,
  and the mechanical classes they closed stay closed.

## Close

See [`progress.md`](progress.md). This tok terminates the invocation, per the skill's Phase 5 § 2 — the
revised strategy must land and be reviewable before the next repair commits to it.
