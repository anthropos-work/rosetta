# iter-86 — repair the remaining upheld findings, and run the guard family for the first time

**Type:** tik, under `TOK-05`.

## Step 0 — re-survey

Gate **4 of 5**. Both repos clean and pushed at open (`rosetta b4b6db8`, `rext b201912`, remote shas
verified with `ls-remote`). Ground truth re-derived at this open and shipped as
[`ground-truth.md`](ground-truth.md) — **with an `origin/main` column for the first time**, which is the
structural half of `CHECK-M257x-iter76-seat-ref-discipline`'s fix (`D-M257x-86-2`). **6 of 14 clones are
behind `origin/main`**: `app` by 60, `next-web-app` by 26, `storage` 6, `jobsimulation` 4, `messenger` 3,
`cms`/`sentinel` 2.

**The guard claim at this open is NOT "6 corpus guards exit 0."** That sentence is what this iter opened
by testing, and it does not survive — see `D-M257x-86-1`. The open state is: **16 family members, 14
GREEN, 2 RED** (`platform_predicate_guard`, `value_change_guard`), measured by
`rext stack-core/guard_family.py`, which did not exist before this iter.

## What this iter does

Three things, in this order, because the second depends on the first being honest.

**1. Run the whole guard family, with selection fixed, and correct the record.** iter-83 found
`repair_leak_guard` had gone RED unnoticed and diagnosed the registry kind-filter. iter-86 measures
that the diagnosis stopped a layer short: **no runner ran the family at all**, and the list in use was
neither the derived registry nor its complement but *a list somebody remembered*. Two guards go RED, and
both were invisible to every green claim the milestone has made.

**2. Settle the fired escalation with a measurement, not a preference** (`D-M257x-86-2`).

**3. Repair the 30 remaining upheld predicate rows + the P4 membership sweep**, by predicate,
adjudicating before repairing, and **grade the reach** — 11/11 is the standard iter-85 set.

## Scope, declared

**In:** Q1 (13) · Q3 (8) · Q4 (8) · Q5 (1) from [`../iter-84/adjudication.md`](../iter-84/adjudication.md);
the P4 sweep's 16 corpus/skills sites + 7 rext sites from [`../iter-84/membership.md`](../iter-84/membership.md);
the guard-family runner; the two REDs.

**Out, and stated so:** no re-read. The separation between repairing and measuring is what makes the
next measurement worth anything, and iter-81 is the milestone's own worked example of what happens when
a repair grades itself.

**Held, untouched:** `DEF-M257x-iter80-storage-prod-bucket` remains escalated and undecided —
`storage.md:55`/`:154`/`:181` stay as they are. `FIX-M257x-iter53-union-set`,
`FIX-M257x-iter56-assignment-flake`, `CHECK-M257x-iter38-ai-act-classification` and RF-2/3/7–14 are not
in scope.

## The denominator, stated because the source states two

iter-84's verdict table says **40 UPHELD** — an *anchor* count. Its by-predicate tables enumerate **37
rows** — a *predicate-row* count. Both are correct in their own unit and they were quoted
interchangeably; separately, **Q4's heading reads `(7)` against 8 rows**. iter-85 took Q2's 7 rows and
expanded them to 9 anchors. **This iter's declared input is the 30 remaining enumerated rows**, not the
29 the Q4 heading produces and not the 33 that `40 − 7` produces. Reach is graded against what is in
[`ledger/r86.md`](ledger/r86.md), which is the only denominator any post-condition can honestly use.
