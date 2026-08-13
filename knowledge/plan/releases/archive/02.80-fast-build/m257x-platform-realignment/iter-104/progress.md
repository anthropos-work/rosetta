# iter-104 — closeout

**Type:** tok (**deliberate** — author-initiated, non-terminating) · `iter_shape: strategy` · authors `TOK-06`.

## The one-line answer

**`TOK-06: fence the inflows before repairing again.`** iter-103 measured that **82 % of the residual arrives
from two sources nothing watches** — clone advance (61 %) and the repair's own induction (21 %) — so the next
work builds the watchers, and the reading goes last.

## Why a tok, and why *deliberate*

The 3-no-prog streak does **not** apply, and the arithmetic is in `D-M257x-104-1`: iter-101 moved `N` from 28
to 24, so the window never reached three consecutive no-progress tiks. This tok is authored on purpose, on
the precedent of TOK-04 and TOK-05.

It does **not** terminate the call (`D-M257x-104-2`). Every element of TOK-06's sequence is an item iter-103
already routed; the tok orders them rather than opening territory, which is the bootstrap tok's own
non-termination rationale. The decision records the bound that keeps it from becoming a general escape: a
deliberate tok is non-terminating **only when it sequences already-routed work**.

## What the revision rests on

Not a mood, and not the number. `N = 33` is only the trigger; the **composition** is the finding.

| | iter-101 | iter-103 |
|---|---|---|
| distinct false **predicates** | **22** | **22** |
| distinct **anchors** | 24 | **33** |

- **The pool did not move by predicate**, after a 52-anchor / 98-site repair. Same falsehoods, more places.
- **The repair is exonerated, blind:** 21 of 22 predicates CLOSED; the one survivor was flagged `SEAT 9 (?)`
  and skipped by iter-102's own map. The repair leg reaches what it aims at.
- **Composition of `N`:** 20 drift (61 %) · 7 induced (21 %) · 4 never-true (12 %) · 2 unclassified.
- **Inflow ≈ outflow.** Running the loop faster does not close clause 5.

## The sequence TOK-06 sets

0. **guard-tree provenance** — every guard verdict in this milestone is provenance-unstated until re-run.
1. **the drift fence** — 61 % of `N` is a mechanically checkable class no guard covers.
2. **the induction checks** — a centralised-wording control + a post-repair line-offset check, designed from
   the two measured shapes, not from the general idea.
3. **repair the 33** — by predicate, with iter-103's two riders.
4. **read LAST.**

Binding on steps 1–2: a mutation control **and** an anti-vacuity control that can actually fire. Six fences
in this milestone have been green over universes they never examined; one compared a string to itself.

## Landed this iter

- **`TOK-06`** appended to the milestone-root [`decisions.md`](../decisions.md).
- **The Chapman sweep** (`D-M257x-104-3`). `state.md` was already clean. Two standing point estimates
  survived in the milestone ledger and are now **marked in place**, not rewritten, per the convention the
  ledger already uses:
  - iter-101's *"the residual is on the order of ~100"* — the **conclusion survives on the floors**
    (≥ 24 / ≥ 33); the estimate does not.
  - iter-102's *"the pool was probably always ~100"* — the second half of that sentence (four corrections to
    an underestimate, not a growing pool) **survives intact**; only the point estimate goes.
  The asymmetry is deliberate: a correction that took more than its evidence would be this milestone's own
  class landing on its own records for the fourth time.
- Four intra-iter decisions, `D-M257x-104-1` … `-4`.

## Gate

**Unchanged at 4 of 5**, and a tok cannot move it. Clause 5 stays open at `N = 33`; it was not re-cut,
narrowed, reinterpreted or argued. Clauses 1 and 2 remain closed at platform `0c91421`, clause 2 **MET WITH
DISCLOSURE** (29/1 on 2 of 2 fresh-stack first runs — never a clean pass).

## Housekeeping

No code. No fence. No repair. No reading. No stack brought up, torn down or reconfigured; `stack-demo/**`
untouched; no clone fetched. Zero platform-repo edits. rext untouched this iter (`944fc4a2`, `main`).

## Close — 2026-08-06

**Outcome:** `TOK-06` authored — fence the inflows before repairing again; Chapman's two surviving point
estimates marked in place. No metric delta by construction.
**Type:** tok (deliberate)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **n — this tok is DELIBERATE, not triggered; the
streak was checked and does not apply (`D-M257x-104-1`), and a deliberate tok that sequences already-routed
work does not terminate the call (`D-M257x-104-2`)** — (3) re-scope: n — (4) user-blocker: n — (5)
cap-reached: n (0 tiks this session) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-104-1` (deliberate, not triggered — with the arithmetic) · `-2` (non-terminating,
and the bound on that reasoning) · `-3` (mark in place, do not rewrite; corrections are asymmetric) · `-4`
(provenance before the fences, and why that is not tidiness)
**Side-deliverables:** none.
**Routes carried forward:** the five-step sequence itself — steps 0–4 above. Also still routed and open:
`FIX-M257x-iter56-assignment-flake` (rate established, repairable), `FIX-M257x-iter103-assignment-context-bleed`,
`DEF-M257x-iter103-aws-bind-provenance` (open, both measurements recorded, neither side asserted),
`DEF-M257x-iter101-briefing-rext-tree` (open, delivered-unfixed, third measurement 4 → 1 → 1), `RF-2/3/7–14`
and the five pass-22 items — **`RF-3` was booked false and was false when booked.**
**Lessons:** the composition of a residual is a measurement in its own right, and it can invert what the
count appears to mean. `N` rose while the predicate pool stood still and the repair's efficacy was
*confirmed* — three facts that only reconcile once you ask where the residual came *from* rather than how
big it is. **Ask what fed the number before deciding the number is bad news.**
