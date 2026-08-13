**Type:** tik — TOK-02 **step 3** ("the two small mechanical fences the classification names", plus the
derived-value fence its table routes `#10`/`#11` to).

# iter-45 — the three mechanical fences

## Phase A — build, each rule measured tree-wide before adoption

Three guards in `rosetta-extensions/stack-core/`, all enrolled in iter-44's commit-time ratchet
(`FENCE_KIND = "postcondition"` + a `postcondition_sites` provider):

| guard | rules | live findings |
|---|---|---|
| `markdown_structure_guard.py` | orphaned list resumption · doubled function word · unbalanced code fence · table-row width | **2** (blocker #6 + a stray ``` at EOF of `.claude/skills/stack-update/reference.md`) |
| `anchor_construct_guard.py` | blank line · bare closing delimiter · table separator · table **header** row · past EOF | **2** (blockers #13 and #16) |
| `derived_value_guard.py` | `**Language**: Go X.Y` vs `go.mod` · `<cpu> CPU / <mem> MB` vs `locals.tf` | **2** (blockers #10 and #11) |

Every rejected draft is recorded next to the rule that replaced it (`D-M257x-45-2`): M1's first draft was
**86% false positive**, M2 needed a hyphen guard on **both** sides of the pair, and the anchor
self-reference rule went from **134 findings — essentially all of them ports** — to 2, both real.

**Blocker #17 was dropped rather than tuned in** (`D-M257x-45-3`). The enumeration rule that would have
caught it fired on 6 of 7 tree-wide candidates and still missed it; narrowing the window until #17 fires
and its neighbours do not is Trap A. Routed to step 4's hand repair, named.

## Phase B — the answer key, captured before step 4 spends it

`tests/fixtures/mechanical/` — two **line-faithful repo roots**, not neighbourhoods, because two of the
five defects are relationships between line numbers (`D-M257x-45-4`). `red/` fires all five at the exact
lines iter-41 anchored; `green/`, produced by declared mechanical transforms, is silent **while still
being resolved and measured** (`D-M257x-45-5`).

The first draft of the suite asserted all five against the LIVE corpus and passed — and every one of
those assertions would have failed at iter-46, whose job is repairing them. That finding is generalized
into `platform-alignment.md` **§8 rule 7**.

## Phase C — the mutation battery, and what it caught

`test_m257x_mechanical_fences_mutation_battery.py`: **20 mutants — 1 declared-GREEN no-op that survives,
19 kills**, ≥5 inversions, ≥5 distinct failure signatures, and **one mutant per reporting path in all
three modules** (the harden 7–9 debt: *delete each new fence's reporting path and confirm a test fails*).

On its first run **three mutants came back GREEN**, naming three real holes in a behaviour suite that had
looked complete — M1 tested on only one of its two column-0 sides, a `corpus/` with zero scannable files
reading as clean, and `measured` counting a doc no scalar was read from (`D-M257x-45-6`). All three closed
with named tests.

## Phase D — two defects found in neighbouring fences, both fixed at the fence

- **iter-44's ratchet rewrote a record it did not move** (`D-M257x-45-7`): `--reason` overwrote
  `claim_twin_guard`'s registration sentence with one about three fences that postdate it. `registered_at`
  already preferred the prior value; `reason` did not. Fixed, reason restored verbatim, two regression
  tests (both directions), one new mutant in iter-44's own battery.
- **A captured fixture was read as this repository's source** (`D-M257x-45-8`): the vendored
  `assignments.go` made `test_write_target_schema_fence` report `stack-core` — which ships no Go — as an
  unclassified Go-bearing rext section. Fixed in the fence: all three walks prune `fixtures` **directly
  under `tests`**, with a two-directional regression test.

## Phase E — measurement

| | |
|---|---|
| platform origin HEAD, open **and** close | `2adcf71` — unchanged; re-scope trigger stays at **occurrence 1 of 2** |
| clause 5 | **18**, unchanged **by construction** — this iteration repairs nothing |
| of the 18, reached by an instrument | **iter-43: 16** (13/13 self-contradiction) → **iter-45: +5 distinct** (#6, #10, #11, #13, #16) |
| `stack-core` suite | **491 tests, 14 failures** — exactly the pre-existing baseline; the one new failure this iteration caused (`D-M257x-45-8`) was fixed, not accepted |
| ratchet on the live tree | `OK` — 4 participating fences, 25 sites, no site the baseline did not already record |
| rext pin | **not moved** (`D-M257x-45-9`) — offline guard/test code only, on no runtime path |

## Close — 2026-08-02

**Outcome:** the three mechanical fences TOK-02 step 3 names now exist, are enrolled in the commit-time
ratchet, and are **watched going RED on a captured answer key that outlives step 4** — with a 20-mutant
battery that found three real holes in its own behaviour suite on first run, and two defects fixed in
neighbouring fences (one of them in iter-44's ratchet, rewriting a record it had not moved).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. Clause 5 stays at **18** by construction: this iteration repairs nothing.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** `D-M257x-45-1` … `D-M257x-45-9`
**Side-deliverables:** the `write_target_schema_fence` fixture-pruning fix (`D-M257x-45-8`) — surfaced by
this iteration's own fixture, fixed in the same commit because leaving it would have left the suite RED.
**Routes carried forward:**

- `FIX-M257x-iter41-blocker-set` — **TOK-02 step 4**, the fence-assisted repair of all 18, by CLAIM not
  by FILE, with `repair_postcondition.py` as the commit post-condition. **iter-46.**
- `FIX-M257x-iter45-blocker-17` — #17 needs an instrument that decides what a sentence *claims*; routed to
  step 4's **hand** repair rather than manufactured into a fence (`D-M257x-45-3`).
- `FIX-M257x-iter43-coverage-protocol-livepath` (TOK-02 step 4).
- `CHECK-M257x-iter35-seeder-writes-one-instant` — still the highest-value open non-gate item.

**Lessons:**

- **A test that asserts a live defect has an expiry date, and what expires is the assertion — not the
  defect.** Five live assertions, all passing, all due to fail at the very next iteration, with the
  obvious repair being to edit the fence's own test to match. A fence stops asserting anything by being
  *maintained*, not by being deleted. Now §8 rule 7.
- **Capture the fixture in the shape the fence READS.** A ±2-line neighbourhood was right for a fence that
  matches text and destroys a fence that reads line numbers. And assert the green twin was still *reached*
  — silence must be earned by the repair, not by loss of reach.
- **The battery is not a formality that ratifies a suite.** Three of twenty mutants came back GREEN, each
  naming a hole nothing else had found. On this iteration the mutation run was the only thing that
  measured the behaviour suite.
- **The instrument built to stop a defect class can contain that class.** iter-44's ratchet — the
  milestone's answer to *records that silently rewrite themselves* — silently rewrote a record. Found by
  reading its own diff, not by any test that existed.
