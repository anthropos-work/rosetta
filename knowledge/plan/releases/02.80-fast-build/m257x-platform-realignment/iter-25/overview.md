---
iter: 25
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-25 — repair the instrument, then re-measure clause 2

**Active strategy reference:** `TOK-01` ("instrument first, then follow") — this iter is the strategy's own
first clause applied to clause 2's instrument rather than clause 1's.

## Step 0 — re-survey

Platform origin HEAD re-fetched at open: **`2adcf71`**, unchanged (re-scope trigger stays at occurrence 1
of 2). rext consumption clone re-pinned to `fast-build-m257x-iter-24` at iter-24's close, so the stack now
carries the Directus re-point.

iter-24's routed order is deliberate and unchanged: **fix the runner before measuring with it.** iter-19
found `run-playthroughs.sh:118` calling a bare `stackseed` that is not on PATH, and had to hand-supply the
path to measure at all. Measuring first and repairing after is how iter-15 ended up comparing two worlds.

## Cluster / target identified

Two planned lines:

1. **`FIX-M257x-iter19-playthrough-runner-path`** — the runner cannot reset itself, so `--reset` is a
   no-op-with-an-error and the suite silently measures whatever world it finds. This is the milestone's own
   §2 class (a hand-supplied path with no derivation) sitting **inside the instrument that measures clause
   2** — the same shape iter-11 found in `autoverify`'s `STACK_DIR`.
2. **The clause-2 re-measure** on a stack that has iter-24's Directus re-point, with a real reset.

## Hypothesis

The runner's bare `stackseed` resolves from the stack's own `bin/` when derived from `N`, exactly like
`OFFSET`; with the reset working, the suite measures the world it seeded. The Directus re-point should clear
the failures whose error text was the `directus_versions` 403.

## Expected lift

**Not predicted as a number.** iter-24's D-M257x-24-3 records why: this fixes one of at least four causes
behind `20 live / 10 failing / 1 unimplemented`, and this milestone has already had one attribution refuted
one iter after it was made. The comparison is a **`diff` of sorted ids** against iter-19's failing set, not
two summary lines — `20/10/1` twice could be ten different failures.

## Phase plan

1. Derive `stackseed`'s path at the point of use; make a missing binary **refuse** (exit 2, naming the path)
   rather than let the run continue into an un-reset world.
2. Live negative control on the guard; positive control that the derivation resolves in a consumption clone.
3. Tag, push, verify on origin, re-pin.
4. Run the full suite with `--reset` from the pinned clone. Full run, because the ptreport gate is binding
   only on a full run (`run-playthroughs.sh:300-307`).
5. Compare sorted failing ids against iter-19's set.

## Escalation conditions

- A platform commit landing mid-iter → re-scope occurrence 2 → STOP.
- The suite not completing in a reasonable window → record what was measured, route the rest; do **not**
  quote a partial run as a clause-2 number.

## Acceptable close-no-lift outcomes

The failing set unchanged, with the reset **proven** to have run, is a legitimate and informative close: it
would mean the `directus_versions` 403 in the error text was not what the assertions actually turned on —
the same shape as iter-19's finding, and worth more than a guess either way.
