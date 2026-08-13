---
iter: 279
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-11
closed: 2026-08-11
active_strategy: TOK-08
---

# iter-279 — fix the SOURCE of the claim iter-278 repaired, and settle its own inconsistency

**Type:** tik · **Active strategy:** `TOK-08`.

## Step 0 — re-survey

Guard family at open: **30 GREEN · 0 RED · 5 not-run**; `clone_drift_guard` GREEN with the corpus at
rext `8e2974f47`. Target confirmed live: `clone_pin_guard.py`'s header still enumerates **three arms**
while shipping four, and still carries the *"`DEMO_ADVANCE_CLONES=pinned` checks each clone out at the
ref it names"* framing of the **canonical** pin — the exact sentence `corpus/ops/platform-alignment.md`
copied and iter-278 retracted.

## Cluster / target identified

`FIX-M257x-278-clone-pin-guard-docstring-says-three-arms`, created by iter-278.

**And it settles an inconsistency in iter-278's own record.** `D-M257x-278-6` declined this fix on the
ground that *"an rext edit advances the clone past the sha this very iter has just reconciled the corpus
to."* iter-278 then **paid exactly that cost anyway** for the census-denominator fix — commit, tag, push,
advance, re-point — and the loop took minutes. The stated reason therefore no longer holds, and leaving
it standing would leave a decision record arguing a cost the same iter proved affordable.

## Hypothesis

Repairing the docstring removes the SOURCE of the false claim rather than one copy of it. The corpus
sentence was copied *from* this docstring; leaving it invites the next reader to copy it forward again.

## Expected lift

No metric moves. The deliverable is: the route closed, the source corrected, the corpus re-pointed to the
new rext sha, and the family still **30 GREEN · 0 RED**.

## Pre-registrations — sealed before the edit

- **PR-1** — `test_clone_pin_guard.py` passes **before and after**; the docstring carries no assertion,
  so no test should change verdict. If one does, the docstring was load-bearing and this is not a
  docs-only change.
- **PR-2** — the arm count in the header is the ONLY count claim about this fence in rext that is wrong;
  if a second is found, the class is bigger than one row and gets routed rather than absorbed.
- **PR-3** — after the loop, `clone_drift_guard` is GREEN with the corpus citing the NEW sha, and the
  family is **30 GREEN · 0 RED · 5 not-run**.

## Phase plan

A. measure the docstring's false claims · B. repair · C. prove tests unchanged · D. ship the loop
(commit → tag → push → verify on origin → advance clone → re-point corpus) · E. close.

## Escalation conditions

A docstring claim that cannot be repaired without a behaviour change → stop and route; this iter is
explicitly documentation-of-code, not code.
