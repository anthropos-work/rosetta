---
iter: 93
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-05
closed: 2026-08-05
---

# iter-93 — fence the HEDGE, not the sentence

**Type:** tik, under `TOK-05`.

## Active strategy reference

`TOK-05: stop repairing claims; fence the predicates under them` — read literally. iter-92 *repaired* six
restatements of one claim by hand. TOK-05's whole point is that hand-repair does not hold, and the measured
evidence from iter-92 is that it did not: the repair itself leaked twice.

## Step 0 — re-survey

The class is currently clean (all 4 `module.*_euwest1` mentions carry a marker after iter-92) — so this is
a fence over a *repaired* surface, which is exactly the state TOK-05 says to fence rather than to admire.

## Cluster / target identified

`CHECK-M257x-iter92-fenced-claim-restatements`. The predicate: **a claim about a repo in no clone set must
say that it is not a measurement.** Legal set derivable: `module.*_euwest1` is declared in
`infrastructure/terraform/production/services.tf`, and `infrastructure` is in no clone set.

## Hypothesis

A tree-wide fence over the *hedge* makes the iter-92 class unrepresentable, where six hand-repairs did not.

## Expected lift

No reading. A new family member, and one fewer way for the corpus to assert through a boundary it cannot see.

## Phase plan

1. Build `unreadable_repo_claim_guard.py`; register it in `guard_family.py` (bidirectional reconciliation
   makes registration mandatory, not optional).
2. Fence it: fires / discriminates / anti-vacuity / premise-is-measured.
3. Record the generalisation in the protocol doc.

## Escalation conditions

- If the marker set cannot be made discriminating without mandating a magic token, say so — a fence that
  forces a ritual phrase teaches people to write the phrase, not to check the claim.
