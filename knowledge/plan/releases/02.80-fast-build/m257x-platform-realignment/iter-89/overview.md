---
iteration_type: tik
status: closed
opened: 2026-08-05
---

# iter-89 — one cause under four failures: demopatch APPLY self-heals, REVERT cannot

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

## Step 0 — re-survey

Gate **4 of 5**. `rosetta 5bddbd5`, `rext 7844e97`, both pushed and verified. Platform re-fetched at open:
`0c91421`, unchanged.

## Cluster / target

iter-88 routed five handlers and deliberately repaired none of them, because the tripwire fired on its
third line. This iter takes the two most serious — `FIX-M257x-iter88-back-to-cockpit-revert` (a **G5
self-revert** failure) and `FIX-M257x-iter88-demopatch-sha-baselines` (3 `pre_sha256` mismatches) — and
asks the question iter-88 wrote down: **are these a defect, or the self-healing freshness gate working?**

## Hypothesis

The cheapest one first, and it was right: **all four are one cause, and the cause is visible in
`git status` on a clone.** iter-88 treated them as up to three classes; §5 rule 28 says three true facts
do not make a cause — join them with one experiment.

## Escalation conditions

- An architectural question whose answer changes what code lands → **user-blocker**, do not guess.
- Uncommitted state in a clone → the user and the orchestrator are the only deciders; **do not clean it.**

Both fired. See [`progress.md`](progress.md).
