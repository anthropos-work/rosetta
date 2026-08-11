---
iteration_type: tik
status: closed-fixed
controlling_strategy: TOK-08
date: 2026-08-09
---

# iter-194 — the registry every brief is told to trust could not read where harden passes route

**Type:** tik · **Active strategy:** `TOK-08` · **Protocol:** `corpus/ops/platform-alignment.md`

## Step 0 — Re-survey before targeting

iter-193 routed `SURVEY-M257x-iter193-harden-routed-items-are-still-invisible-to-the-backlog-fence` as
its own root cause. Re-survey found something the route did **not** say and that changes the iter:
**harden pass 42 had already found this defect**, written a dedicated fence module for it
(`tests/test_harden_origin_route_visibility_m257x.py`), and *deliberately declined to repair it* —
its docstring states the repair needs a disposition grammar the ledger does not use and calls that
*"a design decision, not a corollary of a test."*

So the target is not *discover*, it is **decide and repair**. The framing in iter-193's route
("invisible… unrepaired") is right about the state and wrong about the visibility: the exclusion was
**disclosed by a fence**, just never lifted.

## Cluster / target identified

`route_disposition_guard.collect()` — `glob("iter-*")` and nothing else.

## Hypothesis

The ledger can be read with its own grammar, and doing so brings every harden-routed item into the
population that every session brief is told to trust over hand-written lists.

## Expected lift

Both formerly-invisible routes in the population **with dispositions**; pass 42's compensating registry
retired; ordering proven in both directions.

## Phase plan

1. Measure the invisible set exactly.
2. Read the ledger — **with the ledger's own grammar, measured not assumed**.
3. Retire the compensating machinery pass 42 left, per its own instruction.

## Escalation conditions

If the repair manufactures a false RED in the live registry, that is worse than the silence — fix the
ordering or revert.

## Acceptable close-no-lift outcomes

A measured refutation that the ledger's grammar is not machine-readable.
