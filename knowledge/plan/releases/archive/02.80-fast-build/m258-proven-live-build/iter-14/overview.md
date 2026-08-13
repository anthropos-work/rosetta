---
iter: 14
milestone: M258
iteration_type: tik
status: closed-fixed
created: 2026-08-12
---

# iter-14 — `TIK-B`: the Class A leak, stopped at its producer

**Type:** tik · **Active strategy:** `TOK-02`, `TIK-B` (Class A — coupling to time is **zero**)

## Step 0 — re-survey

`TOK-02` named `FIX-M258-iter11-postgres-anonymous-volumes` and the F-9 stack-dir survival. Re-surveyed:
volumes are still **6 / 0 B** (iter-11's reclaim held, iter-12 confirmed), so the *symptom* is cleared
but the *producer* is untouched — every Postgres start still mints two orphan-able anonymous volumes.
Target current and unabsorbed.

## Cluster / target identified

`D55`'s producer. `D56` exonerated `--purge` by measurement, which leaves the **plain `down`** path
(`rosetta-demo:446`) as the leak: it passed no `-v`.

## Hypothesis

Passing `-v` on the plain teardown branch removes the two anonymous volumes at the moment their owner
dies, costing **zero build seconds** — the defining property of Class A.

## Expected lift

Structural: the 5.297 GB accumulation cannot recur. No time cost either way.

## Phase plan

A. Establish that `-v` is *safe* on the plain path (the named-volume question) by measurement, not
assumption. B. Fix. C. Fence it. D. Price the F-9 stack-dir residue.

## Escalation conditions

If the demo compose declares any named volume, `-v` becomes destructive on the plain path and the fix
must instead enumerate-and-remove by intersection — a bigger change, to be routed rather than rushed.

## Acceptable close-no-lift outcomes

Discovering a named volume (and therefore routing the surgical form) would close this iter honestly.
