**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.

## Open — 2026-08-10

Sealed the five pre-registrations in `overview.md` before building or running anything. rosetta `971cdc4`,
rext `dfb3fb6`, both clean.

## What happened

### 1. The first frozen subject was refuted by its own result — and that was the useful half

`git archive` of both repos at their pinned shas, laid out so `<export>/.agentspace/rosetta-extensions`
mirrors the live tree. The whole `stack-core` section ran in **9 m 11 s** and returned **68 failed / 2,006
passed / 20 skipped / 5 errors**.

That is not a reading of the corpus. `git archive` produces a tree with **no `.git`**, and a large part of
this suite reads *history*: the failure text is literally `file absent at 83ada03^ or 83ada03 — NOT COUNTED`,
`PARTITION BROKEN`, `VERDICT: CANNOT RUN — no eligible instance was readable`. The guards were **behaving
correctly** — refusing to grade what they could not read. The instrument was wrong, not the subject.

Escalation condition #1 of this iter's `overview.md` fired as written. The repair was cheap and is the
reusable half of this iter:

```
git clone --quiet --local --shared . <scratch>/rosetta && git -C <scratch>/rosetta checkout <sha>
git clone --quiet --local --shared .agentspace/rosetta-extensions \
    <scratch>/rosetta/.agentspace/rosetta-extensions && git -C ... checkout <sha>
```

`--local --shared` makes it **0.33 s and 54 MB** for two repos with 172 MB of combined history, and the
result carries a real `.git`, so history-reading guards work. **A frozen subject for this suite must be a
CLONE, never an export.**

### 2. The reading the route asked for

Whole `stack-core` section, `/usr/bin/python3 -m pytest -q` (pytest 8.4.2 / CPython 3.9.6), against the
git-bearing clone at rosetta `971cdc4` + rext `dfb3fb6`:

> **29 failed · 2,052 passed · 18 skipped in 726.42 s (12 m 06 s)**

### 3. Every one of the 29 is environmental — measured, not argued

The 29 failing node-ids were re-run **verbatim against the live tree**, same commits, same runner:

> **29 passed in 806.42 s (13 m 26 s)**

**29 fail on a fresh clone of both repos; the same 29 pass on this operator's tree. The delta between the
two trees is entirely untracked local state** — `stack-dev/`, `stack-demo/`, the platform clone, and the
rest of `.agentspace/`. Not one of the 29 is a corpus defect or a tooling defect in the thing under test.

The failure texts are the finding. They do not say "precondition absent"; they say:

| test | what it printed | what was actually true |
|---|---|---|
| `test_the_shipped_tree_is_green` | `clone set not found at …/stack-demo — cannot derive the Node floor` | the clone set is absent |
| `LiveTree` ×5 (`frozen_expectation_census`) | `DerivationUnavailable: no platform clone at None … a zero from this census would be vacuous` | no platform clone |
| `test_the_live_corpus_is_green` (`fence_command_guard`) | **`a fenced command names a target that does not exist`** | the target is a `stack-*/` path |
| `test_live_corpus_is_green` (`rext_path_guard`) | **`the live corpus must resolve every rext path`** | the path is in an absent clone |
| `test_the_live_reach_is_not_vacuous` | `only 1 command(s) graded; the reach collapsed` | the population lives in absent clones |
| `TheLiveTree.test_the_census_is_not_vacuous` | `136 not greater than 200 — resolution collapsed` | same |

The middle two are the dangerous shape: **on a fresh checkout the suite accuses the corpus.** A new
engineer who clones both repos and runs the guards is told the corpus names targets that do not exist and
fails to resolve its own tool paths. Both statements are false about the corpus and true about their box.

### 4. Attribution: this milestone built all of it

`git log --diff-filter=A` on each of the **15** files holding the 29 failures:

| introduced | file |
|---|---|
| 2026-08-02 | `test_repair_postcondition.py` · `test_iter45_mechanical_fences.py` · `test_m257x_mechanical_fences_mutation_battery.py` |
| 2026-08-03 | `test_repair_postcondition_audit_mode.py` |
| 2026-08-06 | `test_fence_provenance.py` |
| 2026-08-07 | `test_anchor_construct_denominator.py` |
| 2026-08-08 | `test_guard_family_verdict_line_m257x.py` · `test_frozen_expectation_census_m257x.py` · `test_anchor_subject_census_m257x.py` |
| 2026-08-09 | `test_m257x_corpus_file_citations.py` · `test_suite_census_population.py` · `test_claim_census_skip_registry_m257x.py` |
| 2026-08-10 | `test_rext_path_guard.py` · `test_fence_command_guard.py` · `test_toolchain_floor_guard.py` |

**15 of 15 were authored inside M257x**, at a steady rate over nine days, the most recent three **this
week**. No iter noticed, because every iter has run on a box that already had the clones. This is the
milestone's own output, and the class is still being produced.

### 5. Repair — the tranche where the doctrine was already written down

`D-M257x-248-3`, one iter old: *a guard needing a reference DECLARES the need; it does not run without it
and call the result a check it could not do.* Two sites where that rule was already half-applied:

- **`test_frozen_expectation_census_m257x.py::LiveTree`** — the file's own module docstring says *"a census
  that cannot derive must say so (§9)"*, and `frozen_expectation_census` does exactly that by raising
  `DerivationUnavailable`. The **test** then let the refusal escape as a failure. Class-level
  `skipUnless(fec.default_platform() is not None, …)`. `guard_family` already reports this same requirement
  as **NOT-RUN**; the two consumers now agree.
- **`test_toolchain_floor_guard.py::test_the_shipped_tree_is_green`** — already carried
  `skipUnless((ROSETTA_ROOT / guard.SETUP_GUIDE).is_file())`. **The first half of the precondition was
  declared and the second was not**, so an absent clone set read as a RED toolchain registry. Extended to
  require `stack-demo/` too.

Both verified in both directions, which is the point of the pair:

| | frozen clone (no stack workspaces) | live tree (workspaces present) |
|---|---|---|
| before | 6 failed | 6 passed |
| after | **110 passed · 6 skipped · 0 failed**, each naming its precondition | **116 passed · 0 skipped** |

The remaining 23 failures across 13 files are routed, not repaired — see below.

## Pre-registration grading (sealed at `e83560c`, before the export existed)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | frozen-subject failure count | **≤ 6** (point est. 2) | **REFUTED — 29.** Off by ~5×, and in the direction I was warned about only for the *pessimistic* case; this one was optimistic |
| **PR-2** | ≥1 iter-248 failure not reproducible frozen | true | **UNGRADABLE.** iter-248 recorded *"16 failed"* and **not one test name**. A reading whose failure list is not written down cannot be compared with the next reading — booked below |
| **PR-3** | collected count > iter-248's 2,080 | true | **HELD.** 2,052 + 29 + 18 = **2,099** |
| **PR-4** | every surviving failure attributable to a commit **inside** this milestone | **false** (≥1 inherited expected) | **REFUTED — 15 of 15 files are M257x's own.** Zero inherited |
| **PR-5** | ≥1 test that passed in iter-248's run fails frozen | **false** | **REFUTED — 29 of them** |

**Four of five pre-registrations were wrong, and two of them (PR-4, PR-5) were wrong because I assumed the
suite's failures would be *about the corpus*.** They were about the box. Booked in `§5` as the honest
score for this iter; the seal is what makes it visible.

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-248` closed with a better answer than it asked for. The owed clean whole-suite
reading is **29 failed / 2,052 passed / 18 skipped**, and **all 29 pass on the live tree** — the suite has a
**fresh-checkout-hostile** class of **29 tests across 15 files, 15 of 15 authored inside this milestone**,
which on a clean clone of both repos reports the corpus as broken. Two sites repaired to declare their
precondition (6 of the 29), verified RED→skip frozen and still-running-green live. The reusable half is the
frozen-subject recipe: **a clone, never an export** — `--local --shared`, 0.33 s, 54 MB.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-249-1` (a frozen subject for this suite is a CLONE, never an export) ·
`D-M257x-249-2` (a test that needs untracked local state DECLARES it — the tranche, and why skip beats
fail here) · `D-M257x-249-3` (a reading that does not write down its failure NAMES cannot be compared to
the next reading) · `D-M257x-249-4` (the 23 unrepaired are routed as one class, not thirteen items).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
**frozen clone @ `971cdc4`/`dfb3fb6`: 29 failed / 2,052 passed / 18 skipped (726 s)**; the same 29 node-ids
on the **live tree: 29 passed (806 s)**; the two repaired files, **frozen 110 passed / 6 skipped / 0 failed**
and **live 116 passed / 0 skipped**. No whole-section re-run after the repair — the repair is scoped to two
files and both were measured on both trees.

**Side-deliverables:** one measurement kept although it refuted its own hypothesis —
`anchor_offset_guard.citations()` reads **three different populations** for one question (working tree
`root.rglob("*.md")` = **15,258** files on this box vs **2,864** in a clean clone vs `git ls-tree` at a rev
= tracked only). Run all three ways at `971cdc4` they return **26 targets / 60 citations / 0 ambiguous —
identical**. The asymmetry is real, latent, and today costs nothing; recorded, not repaired.

**Routes carried forward:**
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` → **new, and the biggest thing this iter found.** 23
  failures across 13 files still read untracked local state without declaring it. Handler:
  `FIX-M257x-249-declare-the-clone-precondition`. The census instrument already exists — run the suite
  against a `--local --shared` clone pair and diff against a live run. **This should become a fence**, per
  `TOK-08`: the class is mechanically decidable and is still being manufactured (3 of the 15 files are
  from this week).
- `ROUTE-M257x-249-a-reading-must-name-its-failures` → **new.** Every suite reading booked in this
  milestone should record the failing node-ids, not just the count. Handler:
  `FIX-M257x-249-readings-record-names`.
- `ROUTE-M257x-249-anchor-offset-has-three-populations` → **new.** Make the working-tree branch of
  `citations()` share the `at_rev` branch's population (tracked files) so the two modes cannot diverge.
- Still open, untouched this iter: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **A fence inherits the reach of the iter that wrote it — and also the BOX it was written on.** The
   harden pass found four fences whose reach was narrower than their claim. This is the same defect one
   layer down: 15 files assert "the live tree is green" while silently meaning "*this* live tree".
2. **A frozen subject must be a clone, not an export.** The 68-failure run was not a worse reading of the
   corpus; it was a reading of a tree with no history, by a suite that reads history. `--local --shared`
   costs 0.33 s, which is cheaper than the 27-minute serialisation `ROUTE-M257x-248` was proposing.
3. **The instrument was never the bottleneck it looked like.** iter-248 booked the whole section at
   **27 m 34 s** on the live tree; on a clean clone it is **12 m 06 s**, and the same 29 tests alone take
   **13 m 26 s** live. The cost is in walking the operator's own workspaces — this box carries **180,835**
   files under the repo root and **15,258** markdown files, against **4,181** and **2,864** in a clean clone.
4. **Write down the names.** iter-248's "16 failed" cannot be compared to this iter's 29 because the names
   were never recorded. A count is not a reading.
5. **Four of five sealed predictions were wrong**, all of them because I expected the suite's REDs to be
   about the corpus. The seal is what makes that legible instead of forgettable.
