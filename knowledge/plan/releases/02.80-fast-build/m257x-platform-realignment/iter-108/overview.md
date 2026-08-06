---
iteration_type: tik
iter_shape: repair
status: in-progress
opened: 2026-08-06
active_strategy: "TOK-06 — fence the inflows before repairing again (step 3 of 5)"
---

# iter-108 — TOK-06 step 3: repair the 33, under two new fences

## Type selection — and the tok-trigger was NOT waved through

Phase 0 rule 2 has to be graded out loud here, because on its literal words it fires.

The last three tiks are **iter-105, iter-106, iter-107** — three consecutive tiks, no tok between them
(iter-104 was the tok that opened the window), and **not one of them moved `N`.** On the face of it that is
the 3-no-prog streak and this iter should be a triggered tok.

**It does not fire, and the reason is a measurement distinction this milestone already made law.** Rule 2
defines no-progress as *"the metric did not move in any of those 3 tiks (zero or net-negative delta)."* A
delta requires two measurements. **iters 105–107 took none** — `TOK-06` puts the read at **step 4**, last, on
purpose, and each of those three iters says in its own close that **no `N` movement is claimed**. The metric
is therefore **UNMEASURED, not unmoved**, and the streak's precondition is unestablished.

That is not a convenient reading; it is §8's *grade the cannot-tell* (iter-91, re-applied at iter-107 when
`anchor_offset_guard` refused to assert a class it could not decide) pointed at the skill's own trigger.
**Grading "not measured" as "did not move" asserts something nobody measured** — the precise error the
milestone has now booked against itself four times.

The substantive check agrees with the formal one. A triggered tok exists to **revise a stalled strategy**.
`TOK-06` is not stalled — it is **3 of 5 steps executed, on its declared schedule**, and its two
metric-moving steps have never run. Revising it now would revise it *before its own evidence exists*: step 4
IS the evidence any future tok would need to revise from. Firing here would also mean **no declared
multi-step strategy longer than three non-metric steps could ever be executed**, since the tok would
terminate the call mid-sequence every time.

**Recorded as `D-M257x-108-1`, and codified in the protocol doc** so this is a declared rule rather than an
ad-hoc call the next agent has to re-derive — per the skill's protocol-evolution guideline.

**Type: tik.**

## Active strategy reference

`TOK-06: fence the inflows before repairing again` — **step 3**: *`FIX-M257x-iter103-read-union` — the 33,
by predicate, with iter-103's two riders.*

## Step 0 — re-survey

The TOK-named target is current and untouched. Steps 0–2 landed the three fences; nothing has repaired the
union. Two additions to the target since TOK-06 was written, both from the fences themselves:

- **`FIX-M257x-iter107-unbooked-rot`** — 5 rotted citations `anchor_offset_guard` surfaced on `cd16967`
  that **no reading has ever named**. Folded in (the brief's instruction, and they are free findings).
- **`backend.md:54`** — the fourth `:321` citer, **missed by the 14-seat double reading in BOTH passes**.

## Cluster / target identified

`FIX-M257x-iter103-read-union`: **22 predicates / 33 anchors**, composition 20 drift · 7 iter-102-induced ·
4 never-true · 2 unclassified.

**The unit of repair is the PREDICATE, not the anchor** — TOK-05's unit, vindicated by iter-103's band #3
(21 of 22 predicates closed).

**The anchor list is DERIVED, never hand-assembled** (§5 rule 19). Derived via `repair_reach_guard`'s own
`read_ledger()` over `iter-103/raw/` — the same code path that will grade the repair's reach: **48 booked
blocks → 31 distinct primary anchors** across 14 seat reports. A hand list would have missed
`shared_libraries.md:128`; this one did not.

## Hypothesis

Repairing by predicate discharges the union. The metric this iter can honestly move is **reach**
(machine-graded), not `N` — `N` is step 4's to measure, and measuring it here would be repair inside the
measuring pass, which the protocol forbids.

## Expected lift

- **Reach:** every one of the 31 derived primary anchors landed inside a repair hunk (`repair_reach_guard`).
- **`anchor_offset_guard` GREEN on this iter's own repair commit** — the repair must not re-commit iter-102's
  induction.
- **`clone_drift_guard`**: no NEW drift; the pre-existing RED is a finding, not this iter's to hide.

## Escalation conditions

- A predicate whose ground truth cannot be settled from the clones **without a fetch** → route forward;
  §5 rule 41a forbids fetching, and a reading is due at step 4.
- A repair that would require a platform-repo edit → escalate, never edit.

## Acceptable close-no-lift outcomes

A predicate that turns out **already true** (a rejection) is a finding, not a failure — iter-102's apparent
80 % residue was its own rejections. Rejections are reported as such and counted separately from misses.

## Phase plan

1. Derive the ledger; publish it in `claim_ledger.py`'s shape. *(done at open)*
2. Repair by predicate, in one commit per coherent group.
3. Grade reach by machine; run both new fences on the repair's own commit.
4. Close; step 4 (the read) is the next iter.

## Riders (binding — both are measured recurrences, not hypotheticals)

- **A centralised wording is an instrument and needs a control.** iter-102 published a sentence asserting a
  literal had *"one occurrence anywhere in the clone set"* when it has **six**, five of them inside the
  sentence's own denominator — and multiplied it to 5 anchors. **Any wording published to ≥3 sites is
  verified against its own stated denominator BEFORE it is multiplied.**
- **Inserting prose above a cited line re-points the citers.** iter-102 moved `architecture_overview.md:321`
  to `:331` and left all 4 citers. `anchor_offset_guard` now runs on this commit.
