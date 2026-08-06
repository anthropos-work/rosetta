---
iter: 100
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-06
---

# iter-100 — repair the INSTRUMENT before taking another reading

**Active strategy:** [`TOK-05: stop repairing claims; fence the predicates under them`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

## Step 0 — re-survey

iter-99 closed with `N = 28` and a pre-registration failure that indicts the measuring apparatus rather
than the corpus: band #9 (*wrong-construct intra-corpus citations ≤ 1*) failed by **≥ 7×** while
`anchor_construct_guard` was **GREEN at the audited commit**. Re-surveyed at this open: the guard is still
green, and the seven upheld citations are still wrong. The target named by `TOK-05`'s next-tik direction is
therefore superseded by `CHECK-M257x-iter99-anchor-guard-blindspot`, which is the highest-value routed item
and blocks the value of every future reading.

**Not a re-scope.** TOK-05's unit of repair (the predicate) and its ordering (*fence first, then citations,
then the map's new state, then read*) both hold. This iter is the "fence first" rung, applied to the fence
that was found lying.

## Cluster / target identified

`anchor_construct_guard.py`'s green is scoped by the word **"resolvable."** Two of iter-99's seven upheld
wrong-construct citations are shapes `classify()` already detects — a blank line and a closing brace — and
both were missed **in RESOLUTION, never in classification**. They never reached the classifier at all.

## Hypothesis

The blind spot is a *resolution* gap with a mechanically enumerable cause: the corpus writes anchors
**anaphorically** (the file is named once; the anchors that follow carry it), and the guard admitted exactly
one spelling of that construct (`:N above|below|earlier`), always resolving it to the containing document.

## Expected lift

Not a clause-5 lift — a measuring pass cannot be run this iter and an instrument fix cannot move `N` by
itself. The deliverable is a fence that goes RED on a synthetic wrong-construct citation of the missed
class, **proven by mutants**, plus the count of prior findings it re-grades.

## Phase plan

1. Measure the blind-spot population before writing anything.
2. Widen resolution; measure the findings; narrow on every over-match, in the guard's own documented order.
3. Repair whatever the widened fence exposes, so the fence ships usable rather than RED.
4. Behaviour tests for the new class + mutants that kill each new rule.
5. Answer, in writing: the briefing gap, and what would bound the open-ended residual classes.

## Escalation conditions

- If the widened rule cannot be narrowed below a false-positive rate that would get it disabled, **drop it**
  and route — Trap A (tuning a fence to the answer key) is the failure this milestone keeps naming.
- If the corpus repairs exceed the fence's own findings, stop and route the remainder.

## Acceptable close-no-lift outcomes

A measured falsification that the missed class is *not* mechanically resolvable — i.e. that catching it
requires deciding what a sentence claims — would satisfy the iter without a fence change, and is the same
verdict iter-45 recorded for blocker #17.
