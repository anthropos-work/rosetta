---
iter: 107
milestone: M257x
iteration_type: tik
iter_shape: fence
status: closed-fixed
opened: 2026-08-06
---

# iter-107 — the induction check: a repair that moves a cited line must move its citers

**Type:** tik · **Active strategy: `TOK-06` step 2** — the induction checks.

## Step 0 — re-survey before targeting

Re-derived, not inherited:

- **The `:321` rot is still live.** `architecture_overview.md:321` sits inside the **production** topology
  block; the local-stack wording is at `:331`. Measured now: **3 corpus citers still say `:321`**
  (`sentinel.md:85`, `jobsimulation.md:146`, `backend.md:54`) plus **`CLAUDE.md`**. **0 cite `:331`.**
- **The canonical-wording defect is still live too.** iter-102's sentence asserts the
  `backend.internal.anthropos:8081` literal has *"one occurrence anywhere in the clone set"*; measured
  today it has **1 in `app` + 3 in `stack-demo/rosetta-extensions`** — a repo the same sentence's own
  13-repo denominator counts.

Target confirmed.

## Cluster / target identified

TOK-06 named two induction shapes and asked for two checks. This iter takes the **line-offset** one,
because it is the shape that has now occurred **twice by the same mechanism, one cycle apart** — iter-100
booked by iter-101, iter-102 booked by iter-103 — with §5 rule 34 already naming it and not stopping it.

## Hypothesis

The defect is only decidable **at the commit**: looking at `:321` today tells you what is on line 321, not
that a citer meant something else. So the check is commit-scoped, like `repair_leak_guard`, and reads the
authoritative record of what moved — `git diff -U0` — rather than re-scanning the file.

## Expected lift

**No `N` movement claimed.** Clause 3's instrument. The deliverable is that TOK-06 step 3's repair cannot
re-commit the induction the last two repairs did.

## Phase plan

1. `stack-core/anchor_offset_guard.py` — commit-scoped, intra-corpus citations only.
2. Controls whose answer key is **the two real commits** (`cd16967` = iter-102, `a229f8d` = iter-100),
   not fixtures.
3. Synthetic shapes for each case the guard must separate.
4. Register in the family; full suite.

## Escalation conditions

- If the guard cannot go RED on `cd16967` it is theatre and does not ship.
- If any class proves undecidable, it is **named and excluded from the verdict**, never asserted — and the
  OK line must say the green does not cover it.

## Acceptable close-no-lift outcomes

A measured demonstration that the class is undecidable, recorded with the control that showed it, is a
first-class outcome.

> ⚠ **Both escalation conditions FIRED, and the design changed twice in flight.** The first cut returned
> **GREEN on `cd16967`** (a file-level carve-out waived the very citations at issue), and the second cut
> went **RED on a correct re-point**. Recorded in `D-M257x-107-1` and `-2` rather than quietly rewritten.
