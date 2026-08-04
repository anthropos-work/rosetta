# iter-75 — progress

**Type:** tik, under `TOK-05`. Single planned target (`FIX-M257x-iter73-unresolvable-92`),
Phases A–E as declared in `overview.md`.

---

## Phase A — the adjudication, and it decides the whole iter

Universe: the **tracked** files of every clone under `stack-demo/` plus the `rosetta-extensions`
authoring copy — `git ls-files`, **7,265 distinct basenames across 13 clones**, so the answer is the
repository's own and cannot be polluted by build output or an untracked scratch file. Positive
control in the same pass (§5 rule 2): `docker-compose.yml` present.

| fate | sites | distinct basenames | reading |
|---|---|---|---|
| **UNIQUE** | **77** | 26 | the file is exactly one tracked path — **a reach limit** |
| **MULTI** | **26** | 19 | more than one — **must stay unresolved** |
| **ABSENT** | **0** | 0 | **not one citation names a file that does not exist** |

**The class routed as a repair backlog contains zero corpus defects.** That is the fourth
consecutive routed count in this milestone to collapse on adjudication — 64 → 5, 23 → 1, 21 → 0,
and now *"92 unrepaired citations"* → **0 defects · 77 unreachable · 26 undecidable**.

The MULTI residue is the rule working rather than a gap in it: `main.go` is **57** tracked files,
`main.tf` **10**, `mixin.go` 3, `studioManager.go` 2 (`app/internal/cms/studio/` and
`cms/internal/studio/` — the merged copy and the standalone husk, which is exactly the pair a
directory guess would get wrong).

### A defect in the adjudication instrument, twice, in two separate scripts

The first derivation reported **MULTI 72 / UNIQUE 31** — and every rext file in it read *"2 places"*
with **the same path printed twice**. `rosetta-extensions` is cloned twice under this tree: the
per-stack consumption copy at `stack-demo/rosetta-extensions` (pinned at a tag) and the authoring
copy at `.agentspace/rosetta-extensions`. Both directories carry the same **name**, so a universe
keyed by `<clone-name>/<relpath>` collapses them (correct) while one keyed by absolute path splits
them (wrong) — and the first script did the first thing, the dry-run script did the second, so **the
same bug produced two different wrong answers in one iteration.**

Resolved by deciding *which clone is the witness*, not by de-duplicating harder: the authoring copy
is the current one, the per-stack clone is pinned at a tag, and **`resolve()`'s pre-existing rext
fallback already prefers the authoring copy** — so the rule follows the guard rather than inventing
a preference. §5 rule 32's lesson, one level in: *two instruments disagreeing is a finding, and here
both were mine.*

## Phase B — dry run, with a positive control on the dry run

Rule simulated with the guard **untouched**: **77 newly resolvable, 77 clean, 0 findings.**

A 0 from a pipeline that cannot report a finding is not evidence, so the same code path was fed
three known-bad inputs against a real target (`up-injected.sh`, 2,693 lines at the worktree):

| input | verdict |
|---|---|
| `:99999` | `anchor-out-of-range` — *file has 2693 line(s)* |
| `:155` (a blank line) | `anchor-on-blank-line` |
| `:1` / `:2693` | clean |

The pipeline discriminates. **The 0 is a measurement.**

## Phase C — landed

`tracked_basenames(repo_root)` — memoised, `git ls-files` per clone — plus one **last** route in
`resolve()`, taken only when:

1. the citation is **bare** (`"/" not in cited`). A path-qualified citation has already said where
   the file lives; resolving it by basename would override the document with a guess about its
   directory, which is a different act from filling a silence.
2. the basename is **not** in `AMBIGUOUS_BASENAMES`.
3. `git ls-files` names **exactly one** path for it.

And the route is **counted and printed** (`RESOLVE_ROUTES`), because `resolve()` returns a bare
`Path` and a route that fires silently is a reach claim nobody can audit.

| | before | after |
|---|---|---|
| anchors resolved | 177 | **254** |
| unresolvable sites | 239 | **162** |
| findings | 0 | **0** |
| route report | — | `resolved via bare-unique-basename x77` |

`no-clone` grows 30 → 76 and that is correct rather than a regression: rext lives outside
`stack-demo/`, so `_clone_of` returns `None` and those citations are graded at the worktree — named
in the reach line instead of hidden.

## Phase D — gates

| gate | result |
|---|---|
| `anchor_construct_guard` | **OK** — 254 resolved / 112 files; `default x109, no-clone x76, block-pinned x49, ambiguous x20`; `resolved via bare-unique-basename x77` |
| `platform_alignment_guard` | **OK** — F 74 citations, 0 unresolvable |
| `platform_predicate_guard` | **OK** — corpus and platform configuration agree |
| `markdown_structure_guard` | **OK** |
| `corpus_index_guard` | **OK** — 84 docs / 6 dirs |
| `CITE_REF=worktree` | **still discriminates** — `override x252`, 7 findings |
| `tests/test_iter45_mechanical_fences.py` | **74** (was 68); new class `ABareCitationIsFoundByUniqueBasename` 6/6 |
| mutation battery | **6 mutants caught · no-op control SURVIVED** |
| `stack-core` suite | **781 tests, 1F in 708.3 s** — `test_claim_twin_guard_iter48_answer_key::test_02_the_green_twin_of_every_site_stays_SILENT`, the perishable iter-48 fixture. **Baseline 775/1F matched by IDENTITY**, +6 = exactly this iter's new tests |
| `stack-injection` · `dev-stack` · `demo-stack` | untouched sections; iter-71/73's runs stand (332 OK · 151 OK **solo** · 1048/7F by identity) |

### The battery

| mutant | caught |
|---|---|
| M1 uniqueness dropped — take the first of N tracked paths | **2F** |
| M2 the bare-only restriction dropped (a path-qualified citation gets rescued) | **1F** |
| M3 universe from a WALK, not `git ls-files` | **1F** |
| M4 the pinned per-stack rext clone counted too | **1F** |
| M5 the route stops being counted (reach unauditable) | **2F** |
| M6 a clone whose `git ls-files` FAILS is not skipped | **1E** |
| **no-op control** (docstring prose) | **SURVIVED — OK** |

M4 is the iteration's own bug turned into a permanent test, and M6 is §5 rule 1's shape: a clone git
refuses must be skipped loudly-in-code rather than crash or, worse, read as an empty repo.

## Close — 2026-08-04

**Outcome:** `FIX-M257x-iter73-unresolvable-92` was **not a repair backlog** — adjudicated against
`git ls-files` over **7,265 tracked basenames across 13 clones**, the class is **77 unreachable · 26
undecidable · 0 ABSENT**: **not one citation names a file that does not exist.** It is the **ninth
reach limit** of this milestone — iter-73 taught a bare `<name>.<ext>:N` to *reach* the resolver and
did not teach the resolver to *find* it, because every route it owns is positional and an ops doc
citing `` `up-injected.sh:1487` `` supplies no position at all. Landed as a **unique-basename** route
(`git ls-files`, bare citations only, exactly-one-path or nothing), reach **177 → 254 with 0
findings** — and the 0 was earned: the same code path was fed `:99999`, a blank line and two valid
lines first, because 0 was the *surprising* answer one iteration after a comparable widening turned
the corpus RED with 6. The 26 that stay unresolved are the rule working: `main.go` is **57** tracked
files, and `studioManager.go` is the merged copy plus the standalone husk — the exact pair a
directory guess gets wrong. **My own instrument was wrong twice in one iteration and in two
different directions**: `rosetta-extensions` is cloned twice under this tree with the same directory
NAME, so one script collapsed the copies and the other split them, and the fix was to decide **which
copy is the witness** (the authoring one — `resolve()` already preferred it) rather than to
de-duplicate harder.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (2 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-75-1` (0 defects · 77 unreachable · 26 undecidable — the fourth routed count
in a row to collapse on adjudication), `D-M257x-75-2` (uniqueness is the safety argument; 26 staying
unresolved is the rule working; bare-only + `git ls-files`-only), `D-M257x-75-3` (two clones of one
repo are ONE witness — my bug, twice, in two directions), `D-M257x-75-4` (a route that fires
silently is a reach claim nobody can audit — counted, not added to the seven-value return tuple),
`D-M257x-75-5` (a derivation that returns the convenient answer earns a positive control).
**Side-deliverables:** none.
**Routes carried forward:**
- **Closed here:** `FIX-M257x-iter73-unresolvable-92` (and, by measurement,
  `CHECK-M257x-iter57-anchor-guard-bare-class` — the bare class now resolves or is named).
- **The two known-bad citation classes are now BOTH closed** (39 ambiguous → adjudicated inert at
  iter-74; 92 unresolvable → 0 defects here), which is the precondition the orchestrator set for
  taking the graded READ. **iter-76's target.**
- Unchanged: `FENCE-M257x-iter70-line-or-port` · `RF-M257x-iter71-run-returns-a-tuple` ·
  `CHECK-M257x-iter70-studio-room-lines` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**)
  · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `FENCE-M257x-iter54-refs-block` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter58-derive-preregistrations` · `CHECK-M257x-iter52-second-ai-manager` ·
  `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` · `-baseline-refs` ·
  RF-2/3/7–13.

**Lessons:**

1. **Adjudicate before repairing — this is now four for four.** 64 → 5, 23 → 1, 21 → 0, 92 → 0. The
   derivation costs minutes; the repair pass it replaces costs an iteration and can move correct
   claims.
2. **Two clones of one repo are one witness.** Decide which copy is authoritative *before* asking
   whether a basename is unique, and take that decision from something already in the code rather
   than inventing a preference that makes the numbers work.
3. **A derivation that returns the convenient answer earns a control.** §5 rule 2 is usually read as
   a rule about searches; it applies to measurements. 0 findings and a broken pipeline look
   identical.
4. **An unresolvable citation that cannot be resolved without guessing is coverage, not debt.**
   Naming the 26 is the deliverable; resolving them would be the over-match this guard was rewritten
   once to remove.
