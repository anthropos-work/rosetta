---
iter: 273
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: gate-clause-2 — establish the TRUE failure set at the shipping pin
---

# iter-273 — run the whole suite: clause 2's denominator has never been measured here

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

Gate clause 2 reads: *the **full Playthrough suite** passes on that stack — **30 live / 0 failing / 0
error***.

**No full-suite run exists at the shipping pin.** Everything the milestone believes about clause 2 —
*"one failing Playthrough, three siblings pass on the same login and seed"* — comes from runs at rext
`fast-build-m257x-iter-101`, before iter-270 changed the bring-up, and before iter-271 rebuilt the stack
four times. iter-272 ran **one** Playthrough by `--grep`, and the runner said so in its own words:

```
ℹ this run was SCOPED — its artifacts are advisory. The last BINDING run is preserved at
  report/last-binding-run.json
⚠ ptreport gate not met — ADVISORY on a SCOPED run … Re-run unscoped for a binding verdict.
```

**A scoped run is advisory by the harness's own contract.** Clause 2 cannot be graded from one, and
"succession is the only failure" is currently an inherited belief, not a measurement at this pin — the
same shape as the *"empty projection tables"* symptom iter-272 refuted after it had steered two iters.

This is a **substitution under the same strategy**, not a re-scope: `TOK-08` says census the mechanical
classes rather than sample them, and a suite pass/fail set is exactly a census that has been sampled.

## Cluster / target identified

The **binding** suite run: unscoped, reset-to-seed, on the stack iter-271 left green, at pin
`fast-build-m257x-iter-270`. It is a prerequisite for grading clause 2 under **any** fix, so it is not
work that a later fix decision can waste.

## Hypothesis

`pt-workforce-succession` is the only failure and the suite otherwise reads 30 live / 1 failing / 0 error.
Stated as the prior it is — inherited, not measured here — so that a second failure surfaces as a finding
rather than as a surprise during a fix.

## Expected lift

- Clause 2's **true** failure set at the shipping pin, from a **binding** (unscoped) run.
- A measured full-suite wall-clock for this host, which the milestone has never recorded (the standing
  figure, ~45 min, predates both the rebuild and this host's warm caches).
- No fix is promised. This iter measures the denominator that every candidate fix will be graded against.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Run the **unscoped** suite with `--reset` against `demo-2`.
3. Grade: live / failing / error counts, and the identity of every failure.
4. Record; route the fix with the true failure set attached.

## Out of this iter's planned scope (declared, so the tripwire is clean)

- **The succession fix itself** (`FIX-M257x-272-succession-hero-has-no-qualifying-surface`). iter-272
  established the mechanism; the fix needs an rext change, a tag, a pin bump and a re-run, and it must be
  graded against *this* iter's denominator. Landing it here would be two iters in one.
- Gate clause 5 and the inherited route queue.

## Escalation conditions

- `demo-1` is not ours; `--reset` targets `demo-2` only and `stackseed`'s N=0 guard is the floor.
- **No retry-to-green.** The suite runs once; whatever it returns is the measurement. A second run happens
  only to test *determinism* of a NEW failure, and is reported as such.
- Durations are CONTENDED and none is published as a baseline.

## Acceptable close-no-lift outcomes

Any binding result is the deliverable, including "more failures than believed" — which would be the most
valuable outcome available, because every fix plan currently rests on the one-failure prior.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

- **PR-1 — succession is the only failure.** The binding run reports exactly **1** failing Playthrough,
  and it is `pt-workforce-succession`. *Refuted by:* any other count, or a different identity.
- **PR-2 — the live denominator is 30.** The suite reports **30 live** Playthroughs + 1 declared TODO,
  matching `playthroughs.md`'s authoritative count. *Refuted by:* any other live count.
- **PR-3 — zero errors.** No Playthrough reports `error` (as distinct from `failing`) — i.e. nothing
  fails to *run*. *Refuted by:* any error-state entry.
- **PR-4 — the suite is much faster than the standing figure.** Wall-clock is **under 20 minutes**,
  against the ~45 min the milestone has been quoting. *Refuted by:* ≥ 20 minutes.
