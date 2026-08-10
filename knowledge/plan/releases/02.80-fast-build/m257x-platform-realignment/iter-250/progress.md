**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.

## Open — 2026-08-10

Sealed PR-1…PR-5 before deriving anything (`099a947`).

## What happened

### 1. The reach split, derived through the guard's own resolver

Never a hand-written second opinion — iter-209's lesson (a hand-rolled slugger was **16× wrong, entirely
in one direction**). `fence_command_guard.locate` was spied on through a full `check()` run and each
resolved directory classified by where it landed, on the live tree and on iter-249's frozen clone pair:

| tier | who can run it | live, distinct `cd` args | live, graded `cd` occurrences |
|---|---|---|---|
| **repo** | anyone, straight after cloning `rosetta` | **0** | **0 (0.0 %)** |
| **tooling** | after also cloning `rosetta-extensions` | 1 | 1 |
| **workspace** | only after `make init` / `/dev-up` / `/demo-up` | **25 (96.2 %)** | **102 (99.0 %)** |

**Name the unit** (`§5`): 26 distinct args, 103 graded occurrences, and the two percentages are not the
same number. Either way the finding is the same and it is not marginal — **not one fenced `cd` this guard
grades is reachable from a bare checkout of this repo.** The guard was GREEN over all of it, and a document
could migrate wholly into the workspace tier without anything moving.

That is defensible on its face — this corpus documents *operating* a stack, and operating one requires
having cloned it. What it means is narrower and worth writing down: **`fence_command_guard`'s GREEN is
evidence about a provisioned box, and says nothing about the path a new reader is actually on.** The
verdict now says so on every run.

### 2. `rext_path_guard`'s green was a property of the operator — one live instance

The measurement that mattered came from running the guard against the frozen clone pair:

```
LIVE   → OK
FROZEN → 1 reference(s) name a path that does not exist in the rext tree:
             .claude/skills/stack-list/SKILL.md:33: `demo-stack/stacks/registry.json`
```

Same corpus, same rext sha, opposite verdicts. `demo-stack/.gitignore:8` is `stacks/` — the provenance
registry that `/stack-list` documents is **runtime state**, created by running a demo. It exists here
because this box has run demos; it does not exist in a clean clone. **The citation is correct.** The guard
was grading it by *existence*, which is a property of the box, and on the fresh checkout it printed an
accusation against the corpus.

**Repair:** `runtime_artifacts()` asks `git check-ignore --stdin` against the rext tree. `check-ignore`
decides *pathnames* — including ones that do not exist — so the answer comes from the tree, not the
filesystem. Ignored paths go to a **RUNTIME** bucket, printed by name on every run, empty or not. When git
cannot answer (no `.git`, not a repo) the guard prints `RUNTIME BUCKET UNDECIDABLE` and says the grading
fell back to existence — `D-M257x-248-3`, one iter old: *a guard needing a reference DECLARES the need.*

Verified: **live and frozen now return the identical verdict**, `OK` with
`RUNTIME … 1 path(s): demo-stack/stacks/registry.json`, where before it was `rc=0` here and `rc=1` there.

### 3. Tests

- `rext_path_guard`: 5 new tests — ignored+absent is RUNTIME not a finding · **the same tree with the file
  present gives the identical verdict** (the whole point) · **CONTROL: an un-ignored missing path is still
  RED** · a tree git cannot answer for discloses that · the live answer-key instance is in the bucket.
  **The control earned its place immediately** — it failed on its own fixture, because the section set is
  DERIVED from rext's directories and a citation into a section that does not exist is not in subject at
  all. My fixture, my bug, caught before it shipped.
- `fence_command_guard`: 3 new tests — the three tiers are told apart · **every graded `cd` lands in
  exactly one tier and the tiers sum to the denominator** · the split is printed on a GREEN run · the live
  answer key (> 75 % workspace), skipped with a stated reason on a box with no workspace, since the tier
  under test cannot exist there and its absence is not evidence about the corpus.

## Pre-registration grading (sealed at `099a947`, before any measurement)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | tier-**W** share of `fence_command_guard`'s graded targets ≥ 50 % | ≥ 50 % | **HELD — 99.0 % of graded occurrences, 96.2 % of distinct args, and tier `repo` is 0 in both units** |
| **PR-2** | ≥ 1 graded target is tier **X** here | false | **HELD** — `findings=0` live. The 16 unresolved args the probe saw are the guard's *refusals* (placeholders, unanchored single segments), which are counted and printed and were never graded |
| **PR-3** | `rext_path_guard` grades ≥ 1 tier-**W** path | true | **REFUTED** — it resolves against exactly one rext tree (tier `tooling`), never a workspace. Refuted into something better: a tier the taxonomy did not have, **runtime**, with a live instance |
| **PR-4** | either verdict already distinguishes **R** from **W** | false | **HELD** — neither did; both do now |
| **PR-5** | the tier-**W** count is stable live vs fresh clone | true | **REFUTED — 25 → 0.** With no `stack-*/` dir the guard refuses 44 `cd` lines as *"workspace not provisioned on this host"* rather than tiering them. The split is computed by RESOLUTION, so the workspace tier is only visible on a box that has one. A text-only classifier would be a different instrument |

**3 held, 2 refuted** — against iter-249's 1 of 5. Both refutations were of claims about *the other guard*
and *the other tree*, which is where the remaining looseness is.

## Close — 2026-08-10

**Outcome:** Two corpus fences whose GREEN was a statement about this box now state, or no longer depend
on, that fact. `rext_path_guard` graded citations by existence and therefore returned **opposite verdicts
on the live tree and on a clean clone of the same sha**; it now decides runtime state from
`git check-ignore` — the tree, not the filesystem — and returns the same verdict on both, with the one
live instance (`demo-stack/stacks/registry.json`, cited by `/stack-list`) named on every run.
`fence_command_guard` now discloses the reach behind its GREEN: of **103 graded `cd` occurrences, 0 are
reachable from a bare checkout and 102 require a provisioned `stack-*/` workspace.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-250-1` (a cited path that the tree calls runtime state is neither resolvable nor
broken — it is a third bucket, and it must be decided by `git check-ignore`, not by existence) ·
`D-M257x-250-2` (a fence over runnable inputs states WHO can run them) · `D-M257x-250-3` (the reach split
is derived through the guard's own resolver, never a second parser).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6): every test
file touching either guard, **64 passed**; the adjacent verdict-line / provenance / census set,
**233 passed / 1 skipped (226.95 s)**. Guard family (`--platform stack-demo/platform`, repo root):
**29 GREEN / 0 RED / 0 could-not-check / 5 not-run** — unchanged from iter-248. No whole-section run this
iter; iter-249 took the frozen-clone reading and the edits here are scoped to two guards and their tests.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-250-the-runtime-bucket-is-one-guard-wide` → **new.** `rext_path_guard` now decides runtime
  state from the tree. **Every other fence that grades a corpus path by existence has the same defect**
  — `fence_command_guard`'s `locate`, `corpus_citation_guard`, `clone_drift_guard`. Handler:
  `FIX-M257x-250-existence-is-not-a-tree-property`.
- `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` → **new**, from PR-5's refutation. The
  reach split needs a **text-shaped** classifier (does this path name a `stack-*/` workspace or a repo the
  clone set carries?) to be measurable on a fresh checkout at all.
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` (23 failures, 13 files) · `ROUTE-M257x-249-a-reading-must-name-its-failures` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` → open, none touched this iter.
- Still open, untouched: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **Existence is a property of the operator; `git check-ignore` is a property of the tree.** Any fence
   that asks *"does this path exist"* about a repo it does not own has an operator-dependent verdict, and
   the direction of the error is the dangerous one — it accuses the corpus.
2. **A fence over runnable inputs must say who can run them.** `0 of 103` reachable from a bare checkout
   is a fact about this corpus worth knowing, and it was invisible inside a GREEN.
3. **The control that fails on its own fixture is the control working.** The un-ignored-path control
   caught that the section set is derived from rext's directories, before it could ship as a pardon.
4. **Two of five predictions were still wrong, and both were about the tree I was not looking at.**
   PR-5 in particular assumed a split computed by resolution would survive the loss of the thing being
   resolved against. It cannot, by construction.
