**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.
**Route:** `ROUTE-M257x-249-fresh-checkout-hostile-tests` — iter-253 ran the census to a named population;
this is the run-to-zero `TOK-08` asks for.

## Open — 2026-08-10

Sealed PR-1…PR-5 before Phase A's result was read (`b2de62d`). `PR-4` predicts against finding a
flattering extra defect; `PR-1` against the convenient single-fix story.

## What happened

### Phase A — the preconditions were READ, and there are two of them

All 22 reproduce on the frozen pair in **45 s** — a verification loop cheap enough to run per edit, which
is what made the rest of the iter possible at all. Their real failure text was extracted before a single
decorator was written, and it refuted `PR-1` immediately:

| precondition | members | how the failure reads on a fresh clone |
|---|---|---|
| **the clone set** | most | *"a fenced command names a target that does not exist"*, *"resolves in neither pool and is undeclared"* |
| **`node_modules`** | 3 | *"the TypeScript population fell to 0 tests"*, *"prune_census removed nothing"* |

The second was **invisible to inspection**: `test_prune_census_can_return_NON_ZERO`'s own assertion
message already contains its precondition — *"the tooling repo, **which has a node_modules**"* — written
as an assumption. And `UPGRADE-IMPACT-next16.md`, which the corpus cites and which the fresh tree
"could not resolve", lives at `stack-demo/next-web-app/` — a git-ignored clone. The corpus was right.

### Phase B — the declarations, and the half-precondition pattern again

The shared predicates ship in `suite_census.py`, beside the census that names the class
(`D-M257x-254-1`), because the rule that flags a test and the rule that excuses it must be one rule.
`rosetta_root()` **walks** to the checkout instead of counting `parents[N]` — the 12 holding files spell
their root six different ways.

**Five of the eight arms repaired already declared HALF a precondition** — the same shape iter-249 found
in `test_toolchain_floor_guard`:

- `test_the_live_corpus_is_green` skips when the guard returns `2` (*nothing* gradeable). iter-250 measured
  that **1 of 103** graded `cd` occurrences is reachable from a bare checkout — so the guard grades that
  one, returns **1**, and the arm asserts the corpus names a missing target.
- `test_the_live_reach_is_not_vacuous` skips when `total == 0`; `total` is **1**.
- `test_18` / `test_22` skip on `not LIVE`, where `LIVE` asks only whether the **corpus** is present — and
  it is, on a fresh clone. Both grade against the **clones**.

### Phase C — verified in BOTH directions, and the second direction found a bug in my own predicate

| | frozen pair | live tree |
|---|---|---|
| the 8 declared arms | **8 skipped**, each naming its precondition | **8 run and pass** |
| the 22-member target | **22 → 12 failed** | 135 passed / 1 pre-existing skip across the touched files |

**`PR-5` is refuted, and by this iter's own code.** The first `node_modules_present` globbed two levels —
enough from the rext root, not from the **rosetta** root, which one call site passes. So it read `False`
on a machine that has one, and `test_prune_census_can_return_NON_ZERO` **began to SKIP where it had been
passing**. That is the quieter failure mode: a test that stops running looks exactly like a test that
passed. The *live* half of the verification caught it; reading the diff did not. Fixed, bounded
(`max_depth`, because this root carries ~180,000 files), and pinned by a regression test that asserts the
answer from **both** roots.

### The cascade — 2 free, and the third one proved it is not the same thing

`test_m257x_mechanical_fences_mutation_battery`'s three failures were never independent: the battery runs
the suite and asserts its baseline is green, so `test_18`/`test_22` failing failed it three times over.
After those two declared, **2 of the 3 passed with no edit of their own**. The third
(`test_01_every_mutant_matches_its_DECLARED_verdict`) did **not**, so it is a different thing and is
routed rather than decorated (`D-M257x-254-4`). **10 of 22 resolved: 8 declared + 2 cascade.**

## Pre-registration grading (sealed at `b2de62d`)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | all 22 share ONE precondition | **false** | **HELD** — two: the clone set, and `node_modules` |
| **PR-2** | the repair is purely additive, no assertion logic altered | true | **HELD** — every edit is a skip rung or an import; not one assertion changed |
| **PR-3** | after repair the 12 files run **0 failed** frozen | true | **REFUTED** — 12 still fail. The claim was scoped to a full repair, and this iter repaired 10 of 22 |
| **PR-4** | ≥ 1 of the 22 is NOT environmental (the control mis-classified it) | **false** | **HELD** — all 8 examined are environmental; iter-253's control stands |
| **PR-5** | all 22 still RUN and pass live (0 become skips here) | true | **REFUTED** — one did become a skip, from a depth bug in this iter's own predicate |

**3 of 5.** Run trend: **1/5 → 3/5 → 5/5 → 4/4 → 2/5 → 3/5.** Both refutations are about *my own work*
this time rather than about the subject — `PR-3` on scope, `PR-5` on a defect — which is the first time in
this run that the seal has graded the author rather than the tree.

## Close — 2026-08-10

**Outcome:** **10 of the 22 fresh-checkout-hostile tests resolved** — 8 declared with a stated precondition
and 2 falling out as a proven cascade — verified in both directions: on the frozen pair they now **skip**
with the precondition named, and on the live tree all 8 still **run and pass**. The class is **22 → 12**,
across 7 files. Two preconditions were measured where one was predicted (`PR-1` refuted), the shared
predicates ship beside the census that names the class so the two cannot drift, and the live half of the
verification caught a **depth bug in this iter's own predicate** that had silently turned a passing test
into a skip (`PR-5` refuted).
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
*(Corrected in place before the next iter opened. This first read `budget-exhausted: y / exit-7`, and it was wrong: 56 minutes had elapsed across two tiks. "This iter felt large" is the named anti-pattern, not an exit condition — the same self-audit this milestone applies to its instruments applies to its own grading.)*
**Decisions:** `D-M257x-254-1` (the excusing predicate lives with the censusing one) · `D-M257x-254-2`
(two preconditions, and the second was written down only inside a failure message) · `D-M257x-254-3`
(`PR-5` refuted by my own depth bug — a test that stops running looks like a test that passed) ·
`D-M257x-254-4` (do not decorate a cascade; check for a cause first) · `D-M257x-254-5` (the ratchet re-pin
is a standing closing step, and this time the arrow breached the ceiling it raised).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
touched files **live 135 passed / 1 skipped** (that skip pre-existing and declared);
`test_frozen_expectation_census_m257x.py` + `test_fresh_checkout_census_m257x.py` **123 passed**;
the 22-member target on the **frozen pair 12 failed / 2 passed / 8 skipped** (was 22 failed).
Three literal ceilings **exact +0** after re-pin (234 / 219 / 628).

**Side-deliverables:**
- All three literal ratchets re-pinned with recorded reasons (232→234, 216→219, 626→628), every unit of
  the excess this iter's own. The COMMENT arrow **breached the ceiling it was written to raise**, which
  that block's own docstring predicts; convergence took two passes and is recorded as the expected shape.

**Routes carried forward:**
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` → **open at 12, named and reproducible in 45 s.** The
  residual ships at `evidence/iter254-residual-after-repair.txt`: `test_anchor_subject_census_m257x` (4),
  `test_anchor_construct_denominator` (2), `test_m257x_corpus_file_citations` (2), `test_fence_provenance`
  (1), `test_m257x_mechanical_fences_mutation_battery` (1, **not** a cascade), `test_repair_postcondition`
  (1), `test_repair_postcondition_audit_mode` (1). Handler:
  `FIX-M257x-249-declare-the-clone-precondition`. The pattern and both predicates now exist, so each is a
  three-line edit — but each still needs its failure READ, per `D-M257x-254-4`.
- `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet` → **re-affirmed with a second witness in as many
  iters.** Both iters breached the ceilings with their own prose and neither would have noticed without
  an unrelated instrument pointing at them.
- `ROUTE-M257x-254-six-spellings-of-one-root` → **new.** The 12 holding files derive the rosetta root six
  different ways. `rosetta_root()` now exists and is used by the new predicates only; the call sites still
  each carry their own. A consolidation, not a fix — registered, not done.
- Still open, untouched: `ROUTE-M257x-253-suite-census-is-undocumented-in-rext` ·
  `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` ·
  `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` ·
  `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **Verify the direction you are NOT aiming at.** Every finding of substance in this iter came from the
   live half: the depth bug, the cascade, and the confirmation that the declared arms still run. The
   frozen half only confirmed what the iter set out to do.
2. **A test that stops running looks exactly like a test that passed.** Over-declaring is the quiet twin
   of the class this whole route is about, and it is easier to commit — one wrong predicate silences an
   arm with no output at all.
3. **Read the failure before writing the declaration.** The `node_modules` precondition was already
   written down inside an assertion message; the cascade would have cost three decorators on two working
   controls. Neither was visible from the node-id list alone.
4. **Half a precondition is the dominant shape.** Five of the eight repaired arms already had a skip rung
   — keyed on a threshold (`rc == 2`, `total == 0`, `not LIVE`) that a fresh checkout does not quite hit,
   because *something* is always gradeable. A skip rung that has never fired is not a declaration.
