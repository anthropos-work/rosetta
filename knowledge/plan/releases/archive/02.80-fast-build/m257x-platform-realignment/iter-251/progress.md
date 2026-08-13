**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.

## Open — 2026-08-10

Sealed PR-1…PR-5 before measuring (`dbedcea`).

## What happened

### 1. The census

`corpus_citation_guard.run()` was instrumented on its own `Path.exists` calls — the guard's resolver, never
a second parser (`D-M257x-250-3`) — over the live tree:

| | |
|---|---|
| enumerated citations / source documents | **1,808 over 114** |
| population, C1 / C2 / C3 | **1,606 / 196 / 6** |
| distinct paths tested for existence | **133** |
| of those, **git-ignored in `rosetta`** | **21** |
| of the 21, present on this box | **21 of 21** |

All 21 are `.agentspace/rosetta-extensions/…` — the tooling clone. They exist here because this box has
one. **PR-1 and PR-2 hold at 21.**

### 2. The control that decides it

A `git clone --local --shared` of `rosetta` **alone** at `971cdc4` — no `.agentspace`, which is what a
reader gets:

```
corpus-citation-guard: RED -- 21 of 1808 citation(s) do not resolve
    [C1] corpus/ops/demo/playthroughs.md:1531
         ../../../.agentspace/rosetta-extensions/playthroughs/README.md
         -> link target does not exist in this tree            (rc=1)
```

**21 for 21, all false.** The corpus is right; the reader has not cloned the second repo. **PR-3 holds —
and the subject has to be named**: on iter-249's frozen *pair* this guard is clean, because that pair
provisions the tooling clone at the same path. *"A fresh checkout"* is two different trees and they give
opposite answers.

### 3. The repair — and why it is delegation, not exemption

`D-M257x-250-1` transfers verbatim: partition by `git check-ignore`, which decides **pathnames**, existing
or not. But the right conclusion here is stronger than "runtime state". A path this repo ignores because it
belongs to **another repo** is not this guard's subject at all: `rext_path_guard` owns it and resolves it
against the rext tree, where the question is answerable.

So the targets are **delegated** — enumerated, counted, printed by name on every run, and not graded here.
`census.delegable` goes False and the guard prints `DELEGATION UNDECIDABLE` when git cannot answer, rather
than treating "not ignored" as proven (`D-M257x-248-3`).

**The anti-pardon proof, run before the tests were written:** every delegated path was checked against
`rext_path_guard`'s own collected reference set — **21 delegated, 21 owned, 0 orphans.** Delegation moves
them to a fence that can answer them, not into the dark. That check is now the test's answer key.

**Result — the verdict no longer depends on the box:**

| | before | after |
|---|---|---|
| live tree | `OK`, rc=0 (21 silently passed by existence) | `rc=0`, **21 delegated, named** |
| `rosetta`-only checkout | **`RED — 21`, rc=1** | `rc=0`, **21 delegated, named** |

### 4. Tests

5 new, in `tests/test_corpus_citation_guard.py`: an ignored target is delegated not a finding · **the
verdict does not move when the ignored file appears** (the whole point) · **CONTROL: a tracked missing
target is still RED** · a tree git cannot answer for sets `delegable` False · the live answer key —
`delegated > 0` **and** every delegated path has an owning fence, with the failure message naming the
dangerous direction (a delegated citation nothing else grades would be worse than the defect delegation
fixed).

## Pre-registration grading (sealed at `dbedcea`)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | ≥ 1 graded citation resolves only via ignored state | true | **HELD — 21** |
| **PR-2** | that count ≥ 10 | true | **HELD — 21** |
| **PR-3** | a fresh checkout emits ≥ 1 non-defect C1/C3 finding | true | **HELD — 21 of 21 false, rc=1**, on a `rosetta`-only clone; **0** on iter-249's frozen pair, and naming which tree is meant is the content of the claim |
| **PR-4** | after the repair, live == fresh | true | **HELD** — both rc=0, both printing the same 21 by name |
| **PR-5** | ≥ 1 of them is genuinely WRONG | **false** | **HELD** — all 21 resolve in the rext tree, and all 21 sit inside `rext_path_guard`'s reference set, which is green. The operator's tree was pardoning nothing real |

**5 of 5.** The trend across this run is worth recording: **1 of 5 (iter-249) → 3 of 5 (iter-250) → 5 of 5**.
The difference is not luck — iter-249's claims were guesses about what a suite would say, and these are
structural claims about a mechanism already understood from the iter before.

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` advanced to the guard that named it
first and is the family's largest citation fence. `corpus_citation_guard` graded **1,606 C1 + 6 C3**
citations by `tp.exists()` against a tree whose `.gitignore` excludes `.agentspace/`, so **21** citations
of the tooling clone were a silent pass here and **21 false REDs on a checkout of `rosetta` alone** — the
fence telling a new reader the corpus cites files that do not exist. They are now **delegated** by
`git check-ignore`: named on every run, not graded here, and **proven (21 of 21, 0 orphans) to be inside
`rext_path_guard`'s subject**, which resolves them against the tree that has them. Live and fresh-checkout
verdicts are now identical.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-251-1` (out-of-subject is decided by the tree, and the right verb is DELEGATE, not
exempt) · `D-M257x-251-2` (delegation is only legitimate when another fence provably owns the target —
proved before the repair shipped) · `D-M257x-251-3` (the enumeration keeps the delegated citations in its
population; *enumerated* and *graded* are different numbers and both are printed).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`test_corpus_citation_guard.py` **40 passed**; the adjacent citation/index set **23 passed**. Guard family
(`--platform stack-demo/platform`, repo root): **29 GREEN / 0 RED / 0 could-not-check / 5 not-run**,
unchanged across all three iters of this run. No whole-section run; the edits are scoped to one guard and
its tests, and iter-249 holds this run's frozen-tree reading.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` → **advanced, not closed.** Two of the named
  guards are done (`rext_path_guard` iter-250, `corpus_citation_guard` iter-251).
  **`fence_command_guard.locate` and `clone_drift_guard` still grade by existence** — and
  `fence_command_guard` is the one where it matters most, since iter-250 measured that **102 of its 103
  graded `cd` occurrences are workspace-resident.** Handler unchanged:
  `FIX-M257x-250-existence-is-not-a-tree-property`.
- `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` → **new.** This milestone now has two distinct
  "clean" subjects — a `rosetta`-only clone and a `rosetta`+`rosetta-extensions` pair — which disagree.
  Every future frozen reading must say which, or it is not reproducible.
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` (23 failures, 13 files) ·
  `ROUTE-M257x-249-a-reading-must-name-its-failures` · `ROUTE-M257x-249-anchor-offset-has-three-populations` ·
  `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` → open.
- Still open, untouched: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **"Out of subject" is a better answer than "exempt", and it is checkable.** An exemption asks to be
   trusted; a delegation names the fence that took over, and that claim can be — and was — proven.
2. **There are two fresh checkouts.** A `rosetta`-only clone and a `rosetta`+tooling pair give opposite
   verdicts from the same guard on the same corpus. Every reading in this milestone must name which.
3. **Three iters, one defect, three guards.** The class iter-250 found by accident — *existence is a
   property of the operator* — has now been confirmed in `rext_path_guard` (1 instance),
   `corpus_citation_guard` (21) and, unfixed, `fence_command_guard` (102 of 103 graded `cd`). It was never
   a one-guard bug; it is how this family was built.
4. **Calibration improves when the claim is structural.** 1 of 5 → 3 of 5 → 5 of 5 across this run, and the
   difference is whether the prediction was about a mechanism I had already measured or about what a suite
   would happen to say.
