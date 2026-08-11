---
milestone: M257x
iter: 28
iteration_type: tik
status: archived
created: 2026-08-01
---

# iter-28 — `MEASURE-M257x-iter28-clause2`: the binding clause-2 number

**Active strategy reference:** `TOK-01: instrument first, then follow`.

## Step 0 — re-survey, and the SUBSTITUTION it produced

iter-27's close named `FIX-M257x-iter27-succession-hero-not-rendered` as the next tik. The re-survey
substitutes **`MEASURE-M257x-iter28-clause2`**, which iter-27 routed as *"its own iteration"*. This is a
substitution of the named target under the same strategy, not a re-scope.

**Why, stated honestly.** Two reasons, and the second is the binding one:

1. The succession dig is **open-ended**: the row the spec wants is a *computed projection*
   ("Rare skill held only by this person / In fragile role"), i.e. a key-person-risk signal the app derives
   from skill concentration, not a seeded field. It is not obviously reachable without a live authenticated
   read of the surface, and it could consume an entire iteration and produce nothing.
2. The full run is **high-value and cheap in the resource that is actually scarce here** — it is one
   long-running command, not a long investigation. It produces a **binding** clause-2 number (the ptreport
   gate binds only on a full run), and iter-27 gave it a specific, falsifiable expectation to test.

## Cluster / target identified

The clause-2 measurement itself. Two things make this iteration worth its wall-clock:

- **It tests iter-27's prediction.** iter-27 proved `workforce-intelligence.organization-feedback.UC1`
  passes, but only on a **scoped** run that the harness itself grades advisory. The binding number is still
  iter-26's `23 live / 7 failing / 1 unimplemented`. Expected `24/6/1` — **expected, not claimed**, and the
  comparison must be a `diff` of sorted ids, never two summary lines (iter-19's rule: `20/10/1` twice could
  be ten different failures).
- **It repairs a destroyed artifact.** `FIX-M257x-iter27-scoped-run-clobbers-binding-report`: iter-27's own
  scoped diagnostic run overwrote `e2e/report/last-run.json`, taking the binding 209-spec artifact down to
  1 spec. A full run re-establishes it.

It also exercises a path iter-27's manual reseed did **not**: `run-playthroughs.sh --reset` additionally
refreshes the fake-FAPI roster, re-exports the cockpit manifest and reloads Sentinel's casbin enforcer.
iter-27 ran `stackseed` directly and skipped all three (its Sentinel reload returned `000`, i.e. never
connected — the offset port). So this is the first end-to-end reset since the fix.

## Hypothesis

`pt-workforce-org-feedback` flips to passing; the other six failures are unchanged. Any OTHER movement —
in either direction — is the finding, because nothing else was touched.

## Expected lift

`23 → 24` live, `7 → 6` failing. **Claimed only if the sorted-id diff shows exactly one removal and zero
additions.** Anything else is reported as measured, not as expected.

## Phase plan

1. Full `--reset` run (serial, `workers:1`; budgeted as the iteration).
2. Extract the failing set from the run's own artifact; `diff` the sorted ids against iter-27's recorded
   seven.
3. Record the binding number and re-baseline the routed items against it.

## Escalation conditions

- A NEW failure that iter-27 could plausibly have caused → investigate before closing; do not report the
  headline number over a regression.
- The run does not complete within the session → close honestly as no-lift with the partial state named,
  and do **not** quote a partial run as a clause-2 number (iter-25's exact mistake).

## Acceptable close-no-lift outcomes

A binding measurement that shows **no** movement is a complete iter — it would refute iter-27's scoped-run
proof, which is a finding worth more than the lift.
