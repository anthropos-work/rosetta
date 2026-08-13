---
iter: 104
milestone: M257x
iteration_type: tok
tok_flavor: deliberate
iter_shape: strategy
status: closed-fixed
opened: 2026-08-06
---

# iter-104 — the strategy revision iter-103 paid for

**Type:** tok (deliberate — author-initiated, NOT the 3-no-prog streak, NOT session-terminating)

## Why a tok now, and why it is *deliberate* rather than *triggered*

The streak clause was checked **before** this was written, and it does **not** apply. The last three tiks:

| iter | shape | primary metric |
|---|---|---|
| iter-101 | reading | `N` **28 → 24** — progress |
| iter-102 | repair | 52 anchors / 98 sites; no reading inside it |
| iter-103 | reading | `N` **24 → 33** |

iter-101 moved the metric, so the window never reached three consecutive no-progress tiks. Phase 0 rule 2
therefore falls through to tik. **This tok is authored on purpose** — the same precedent as TOK-04 and
TOK-05, both fired by something other than the streak.

What makes it non-terminating is the same property that makes a bootstrap tok non-terminating: there is no
*stalled strategy the user needs to review before the next tik commits to it.* iter-103 produced both the
measurement that makes the current sequencing indefensible **and** the replacement, and every element of the
replacement is an item already routed in iter-103's own close. This tok **sequences already-routed work**; it
does not pivot into unreviewed territory. So it closes and the loop continues into tiks in the same call.

## What iter-103 measured that forces the revision

`N = 33` against a rule sealed in its own commit (`04cbcfc`) before the first seat was dealt: `≤16` works ·
`17–22` ambiguous · **`≥23` DOES NOT REACH**. The `≥23` branch fired.

The composition inverts what that appears to mean:

- **By predicate the pool did not move — 22 at iter-101, 22 at iter-103.** Anchors went 24 → 33. Same
  falsehoods, more places.
- **The repair is exonerated, blind.** 21 of iter-101's 22 predicates are CLOSED; the single survivor
  (`prod-terraform-8081` at `skiller.md:19`) is one iter-102's own repair map flagged `SEAT 9 (?)` and left.
- **Composition of `N`:** 20 drift (61 %) · 7 iter-102-induced (21 %) · 4 never-true (12 %) · 2 unclassified.
- **Inflow ≈ outflow.** Repair reaches its targets; two inflows nothing watches keep the residual up.

**The failure is not convergence of repair. Running the loop faster will not close clause 5.**

## The revision

**Fence the inflows before repairing again.** Full record: [`TOK-06`](../decisions.md) in the milestone-root
`decisions.md`.

Sequence:

0. **guard-tree provenance** — every guard verdict in this milestone has unstated provenance.
1. **the drift fence** — 61 % of `N` is a mechanically checkable class that no guard fences.
2. **the induction checks** — 21 % of `N` is prose the repair itself wrote, in two mechanical shapes.
3. **repair the 33** — `FIX-M257x-iter103-read-union`, with iter-103's two riders.
4. **read LAST** — once the inflows are watched.

## Planned scope of THIS iter

Strategy authoring only. Specifically:

- `TOK-06` appended to the milestone-root `decisions.md`.
- The **Chapman retirement** propagated: `state.md` and the milestone `progress.md` swept for surviving point
  estimates from the 16.7 → 29.4 → 45.2 → ~103 series. Only floors survive (`≥ 24` at `8f04d3a`, `≥ 33` at
  `e6aed2e`).
- `iter-104/decisions.md` records the intra-iter decisions.

**No code. No fence. No repair. No reading.** Those are iters 105+.

## Escalation conditions

- If the sweep for surviving point estimates finds one inside a **fenced** claim, that is a corpus defect and
  routes to the repair iter rather than being silently rewritten here.
- If the streak re-check had shown three consecutive no-progress tiks, this becomes a *triggered* tok and
  terminates the call. It did not.

## Acceptable close-no-lift outcomes

A tok has no metric delta by construction. This closes `closed-fixed` when TOK-06 is recorded and the
Chapman sweep lands; `closed-no-lift` if the streak re-check had invalidated the premise (it did not).
