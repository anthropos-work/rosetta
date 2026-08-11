---
milestone: M257x
iter: 29
iteration_type: tik
status: archived
created: 2026-08-01
---

# iter-29 — `CHECK-M257x-iter28-clause2-flake-component`: is `25/5/1` a property of the build?

**Active strategy reference:** `TOK-01: instrument first, then follow` — and this iter is that strategy's
own subject matter. TOK-01's opening line is *"build the instrument before doing the re-point"*; clause 2's
instrument has been trusted for 14 iters without anyone asking whether it returns the same answer twice.

## Step 0 — re-survey

Gate 3 of 5. Clause 2 at `25 live / 5 failing / 1 unimplemented` from iter-28's binding run. Platform origin
`2adcf71`, unchanged. Both trees clean, rext pinned at `fast-build-m257x-iter-27` (no runtime change since).

## Cluster / target identified

**Two of the seven clause-2 failures have now resolved themselves with no targeted change:**

| id | flipped at | attributed? |
|---|---|---|
| `hiring.recruiter-comparison.UC1` | iter-26 | no — *"plausible mechanism, nothing measured"*, still open |
| `pt-assignment-assign` | iter-28 | no — deliberately not attributed |

Two un-attributed flips in three iters is evidence **about the instrument**, not only about the build. It
has never been measured because **no two full runs of this suite have ever been done against the same
build** — every previous full run was separated by a landed fix, so a moving id was always explicable.

## Hypothesis

`25/5/1` is not a stable property of the build; at least one id is bistable across runs. Concretely: runs B
and C will not both return exactly the iter-28 failing set.

**The null result is equally valuable and is not a failure of the iter:** if all three runs return the
identical sorted-id set, the flake hypothesis is refuted, both flips become genuinely un-explained
*one-way* transitions needing a different account, and — more usefully — clause 2 gains something it has
never had: **evidence that its instrument is deterministic**, which is a precondition for a `30/0/0`
conjunction ever meaning anything.

## Expected lift

**None on the primary metric, by design.** This is a measurement tik: no source changes, so `25/5/1` must
not move for any reason attributable to this iter. Its deliverable is the *characterization*.

## Phase plan

1. Two further full `--reset` runs (B, C) against the **unchanged** build; iter-28's run is A.
2. Preserve each run's own JSON artifact separately (`runA/B/C.json`) — this is also the standing defence
   against `FIX-M257x-iter27-scoped-run-clobbers-binding-report`, which would otherwise leave only the last.
3. Three-way sorted-id diff. Report the **union** (ever-failing), the **intersection** (always-failing) and
   the **symmetric difference** (bistable) — never three summary counts.
4. Re-baseline the routed items against the intersection, since only always-failing ids are worth targeting
   with a fix.

## Escalation conditions

- A run does not complete → report what completed; **never quote a partial run as a clause-2 number**
  (iter-25's mistake).
- Any id appears that is in none of A/B/C's predecessors → investigate before closing; a NEW failure on an
  unchanged build is a stronger signal than a flip and must not be averaged into a flake rate.

## Acceptable close-no-lift outcomes

Both outcomes close the iter: bistability measured and quantified, **or** determinism demonstrated across
three runs. The iter fails only if it produces a number without the id-level diff behind it.
