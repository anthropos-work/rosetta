---
iter: 105
milestone: M257x
iteration_type: tik
iter_shape: fence
status: closed-fixed
opened: 2026-08-06
---

# iter-105 — a guard verdict states the tree it was taken with

**Type:** tik · **Active strategy: `TOK-06` step 0** (`fence the inflows before repairing again`).

## Step 0 — re-survey before targeting

TOK-06 was authored one iter ago and names `FIX-M257x-iter103-guard-tree-provenance` as step 0. Re-checked
against current evidence rather than inherited:

- Family run from the **authoring** copy at the iter open: **14 GREEN · 0 RED · 0 could-not-check · 3
  not-run** over 17 members, corpus `22eaac4`, platform `0c91421df`. The transcript names the corpus sha and
  the platform sha. **It does not name the fence tree's own sha.** The defect is live and unrepaired.
- **Measured re-grade surface: 52 recorded family verdicts across 26 milestone artifacts, of which 0 state
  the fence tree.** (`grep -rnoE "[0-9]+ GREEN · [0-9]+ RED"` over the milestone dir; the single incidental
  `rext`-mentioning line names a module path, not a sha.)

Target confirmed, not substituted.

## Cluster / target identified

`guard_family.py:344-346` prints the corpus sha and `:367` the platform sha. **The one input that decides
the verdict — the tree the fence configuration lives in — is the one the output does not state.** iter-103
measured the cost at 8 sites and two false quotable conclusions: run from the **pinned** clone (`09d06070`),
both platform guards read **RED**; from the **authoring** copy (`944fc4a2`), both read **GREEN**, and the
entire difference was `claim_twin_waivers.json` (+40 lines) — the 8 RED sites were exactly the 8 waived
sites.

## Hypothesis

A verdict is a measurement taken with a fence's **configuration**, so it is settled by the tree that
configuration lives in. If every verdict carries that tree's path + sha + dirty state, the iter-103 confusion
is **not reproducible** — the two transcripts would have differed visibly in their first line.

## Expected lift

**No movement on `N`**, and none is claimed. This is a step-0 instrument iter: its deliverable is that steps
1–3 ship fences whose founding greens can be re-checked. The gate metric it touches is clause 3's instrument
(the guard family), not clause 5's (the graded read) — the milestone's standing rule that those are two
instruments applies, and neither speaks for the other.

## Phase plan

1. `fence_provenance.py` — a non-`*_guard.py` module (so it does not enter the census) exposing the fence
   tree's path, sha, dirty state and describe, plus a `stamp()` that prints it once per process.
2. `guard_family.py` — prints the fence-tree line **beside** corpus and platform, before any verdict; and
   treats an undeterminable fence tree as **UNMEASURED (exit 2)** on its own doctrine, with an
   `--allow-unknown-provenance` escape that RECORDS the gap the way `--allow-not-run` does.
3. Every `*_guard.py` + `repair_postcondition.py` stamps on **direct execution**, so a standalone verdict
   carries it too. Printed FIRST, never last, so `guard_family.run_one`'s `lines[-1]` reporting and
   `headline()`'s finding-shaped-line cut are both untouched.
4. Controls, per TOK-06's binding clause:
   - **AST-based** conformance check derived from disk in both directions (§8's *assert against a parsed
     construct, never a whole-file substring*).
   - **Mutation control** — a copy with the stamp removed must turn the check RED.
   - **Anti-vacuity control that can fire** — the census must be non-empty AND agree with
     `guard_family.census()`; plus a live subprocess run asserting the stamp appears, and that suppression
     suppresses it.
5. Re-run the family + the full stack-core suite at close.

## Escalation conditions

- If making the family exit 2 on unknown provenance breaks a legitimate deployment shape (a non-git fence
  tree), the refusal is downgraded to a disclosed warning and the reason recorded — **not** silently dropped.
- If the 17 stamp edits break any existing test, that is the induction class TOK-06 named. Stop, land the
  family-level half only, and route the per-guard half forward.

## Acceptable close-no-lift outcomes

If the AST conformance check cannot be made to fail on its mutant, the fence is not shipped (TOK-06's binding
clause) and the iter closes `closed-no-lift` with that falsification recorded.
