---
iter: 69
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
---

# iter-69 — the citation class was 5 defects wearing a count of 96

**Active strategy reference:** `TOK-05` (*stop repairing claims; fence the predicates under them*),
step 2 of its ordering — **fence → citations → map state → read**. This is the citations step,
carrying `FIX-M257x-iter63-app-citation-residual` scope **B2**.

## Step 0 — re-survey before targeting (mandatory)

Three numbers re-derived at open, none inherited:

| | routed to this iter | re-derived at open |
|---|---|---|
| the app citation class | 96 distinct / 22 files (briefing) — 105 (iter-68's own close) | **135 sites / 105 distinct / 22 files** (`iter68_cites.py`, ref `origin/main`) |
| platform clone vs origin | level at `0dab54d` | **level** — `0dab54d` both sides |
| `app` clone vs origin | clone `v1.366.0` `b948604`; `v1.367.0` on origin | **56 behind**; origin/main `9d00a313` v1.367.0 @ 2026-08-04 10:56Z — **unchanged since iter-68 measured it** |

The class grew **96 → 105** between iter-68's open and its close, with no platform movement in that
term. §5 rule 34's sibling again: **a corpus repair enlarges its own citation class.** Budgeted for,
re-measured, and re-measured again at close.

## Cluster / target identified

TOK-05 named the citations step and `FIX-M257x-iter63-app-citation-residual` B2 named nine files.
Both stand — **but the unit is wrong, and §7 rule 4b caught it before the iteration was spent.**

B2 was routed as *"64 unrepaired non-mainline citations."* Graded at the ref the gate names, the 126
distinct (file × citation) pairs partition:

| class | n | is it a defect? |
|---|---|---|
| identical at `b948604` **and** `origin/main` | **62** | no — HELD |
| drifted, but the corpus block **names its ref** | **59** | **no — a measurement, not a standing claim** (TOK-04 P1) |
| drifted and **UNPINNED** | **2** | **YES** |
| file absent at the ref | **3** | **YES — but mis-rooted, not dead** |
| out of range | 0 | — |

**The residual is 5, not 64.** The other 121 are either true at origin HEAD or true at a ref they
state. This is TOK-05's thesis arriving from the other direction: the *predicate* under a citation
is not *"this line names this construct"* but **"this line names this construct AT THE REF THIS
BLOCK CLAIMS"** — and a corpus that writes its refs has already discharged 59 of them.

## Hypothesis

Repairing the 5 closes the gate-relevant citation residual; the remaining 121 need no edit and any
pass that "repairs" them is inducing drift, not removing it (§5 rule 34). The `* **Profile**:`
bullet fence then locks the construct iter-68 repaired by hand and no fence reaches.

## Expected lift

- The **unpinned-and-moved** class goes **2 → 0**, adjudicated at `origin/main`.
- The **mis-rooted** class goes **3 → 0** — stated as what it is, not deleted.
- `* **Profile**:` bullets become **reachable**: 8 sites, 0 findings, with a RED-before-trusted mutant.
- Clause 5's citation residual is closed **with a derived denominator**, so the next reading is not
  taking place over a known-bad class.

## Phase plan

- **A** — re-derive the class + §7 rule 4b at both refs (**done at open**; recorded above).
- **B** — screen mechanically for the one gate-relevant shape; RED-before-trusted with two inverted
  mutants + a no-op positive control that must SURVIVE (§8 rule 5).
- **C** — read and adjudicate the 5 candidates individually against **platform artifacts**.
- **D** — repair; prefer *deletion > minimal scoping edit > rewrite*; count added words.
- **E** — `FENCE-M257x-iter68-profile-bullet`: the `* **Profile**:` bullet construct, both directions.
- **F** — gates: five corpus guards + the `stack-core` suite at its failure **identity** baseline.

## Escalation conditions

- If the screen's residual had come out in the dozens, this iter would have repaired the *shape* and
  routed the volume — it came out at 5, so it lands whole (Fate 1).
- A defect that needs a platform edit to fix escalates; it does not get edited (zero platform edits).
- If the profile-bullet fence cannot be made to go RED on an inverted mutant, it does not ship.

## Acceptable close-no-lift outcomes

That the 5 candidates each turn out to be true-as-written would be a complete iter: it would falsify
B2's premise and the deliverable would be the **derivation + the fence**, not the edits.
