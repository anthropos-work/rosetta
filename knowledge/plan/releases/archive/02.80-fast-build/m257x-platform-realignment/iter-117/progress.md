**Type:** tik · `iter_shape: census` · **`TOK-08` class 1 — intra-corpus citation resolution.**

# iter-117 — the corpus stops mis-citing itself, and the census says how much that is worth

_`TOK-08` — the user's re-scope — was sealed in this iter's FIRST commit (`577446b`), before a single
sweep line, exactly as `TOK-07` sealed its own falsification before its first seat was dealt._

## The numbers, in the shape `TOK-08` demands: population, already-false, reach with its denominator

| | value |
|---|---|
| **enumerated population** | **1,520 intra-corpus citations** over **92 source documents** |
| — C1 path resolution | 1,337 |
| — C2 anchor resolution | 179 |
| — C3 explicit line-pin range | 4 |
| **denominator provenance** | **`corpus-derived-per-arm`** — stated in the report and in `--json`, per iter-114's rule |
| **already false at census** | **8** (0.53 %), **all 8 in the C2 anchor arm** |
| **repaired** | **8 of 8** |
| **fence reach** | **1,520 of 1,520 = 100 % of the enumerated population**, standing green |

**Invocation, stated with the count** (§5 rule 12):
`python3 stack-core/corpus_citation_guard.py --repo-root /Users/marco/workspace/anthropos/rosetta`
at rext `571db21`, rosetta `44f8239`.

**0 findings in path resolution and 0 in line-pin range.** Every corpus path the corpus names exists;
every explicit corpus line-pin is in range. **The entire measured class lives in the anchor arm — the one
arm no reader ever exercises**, because a broken `#fragment` still lands you on the right file. That is
why 8 defects survived four graded readings of a corpus that has been read line-by-line for weeks.

## The eight, and one of them had shipped a literal placeholder

| site(s) | defect |
|---|---|
| `platform-migration-status.md` ×4 | one extra hyphen in every deep link into the protocol doc — `#6--classification--the-map` for a heading that slugs to `#6-classification--the-map`. `## 6. Classification` collapses the `.` to a **single** separator; the author's model was `## 6 — Classification`. **Four sites, one predicate** |
| `secrets-spec.md:383` | anchor names `…-v16-m27m28`; the heading reads **M27–M30**. Stale by two milestones, with correct prose around it — invisible to a reader |
| `staging-bringup.md:55` | `#what-if-a-developer-wants-to-test-a-feature-branch-on-staging` for `## What if I want to test a feature branch with prod-shape data?` — paraphrased from memory |
| `architecture_overview.md:48` | *"see [AI Providers](#ai-providers) below"* — no such heading anywhere in the document. Re-pointed at `#external-service-integration`, the section that actually carries the no-EU-first-ladder correction |
| `coverage-protocol.md:350` | **`[serve-grant](#…)` — a literal ellipsis placeholder that shipped** and survived every reading. Re-pointed at `snapshot-spec.md`'s M40 serve-grant section |

## The finding that matters more than the eight, and it bears on `TOK-08`'s own stop condition

**The mechanically-censusable half of class 1 is largely DISJOINT from the half iter-116's reading
booked.** Band #7's 10 predicates are *construct* defects — a pin naming lines that hold something else
(*"`ai-readiness.md`'s `✅ CORRECTED M219` blockquote is `:476-496`"*, *"the **Data** bullet of `cms.md`'s
merge banner is at `:44-47`"*). This census measured that shape and found it **not machine-reachable at
scale**: of **387** lines carrying a bare `` `:NN` `` pin, only **4** name exactly one corpus document and
no other source path. The rest refer to a platform file named earlier in the same sentence
(`app/main.go:524`, then `` `:525` ``) or to a **port** (`` `:5050` ``, `` `:8082` ``, `` `:3000` ``).

So the honest projection, recorded **before** the reading that will grade it: **closing class 1's
mechanical half should not be expected to move `P` by much.** That is early evidence bearing on
`TOK-08`'s refutation branch, and it is reported as such rather than argued away — the iter's
`overview.md` pre-registered exactly this as an acceptable outcome.

## Four census passes before one line of prose was repaired

Three of them existed only to kill a false-positive class. **Every one was found by measurement, not by
reasoning**, and each is now a named regression test:

| draft rule | false REDs | why it was wrong |
|---|---|---|
| a backticked basename is a relative path | **~180** | the corpus names docs by basename constantly; `` `safety.md` `` is a NAME, not a path assertion |
| a bare `` `:NN` `` pin resolves against the last-named doc | **256** | it is a platform-file line, or a port |
| `knowledge/…`, `.claude/…` are rosetta paths | **5** | those dirs exist in the platform repos; the corpus cites `app`'s `knowledge/deployment.md` and clerkenstein's `knowledge/injection.md` |
| an anchor is a markdown heading | **22** | `<a id>` and `{#custom-id}` are anchor definitions too — and `_` is stripped as emphasis but kept in `pg_dump`. Includes **`CLAUDE.md`'s own headline retraction link** |

**A fence that had shipped on any one of those drafts would have turned ~460 correct citations RED.**
§8 rule 6 says where that ends: it gets disabled. Under-flagging is the correct direction, the guard
takes it deliberately, and it says so in its own docstring.

## The mutation control earned its keep on its first run

It caught a **silent vacuity bug in this guard**: an unresolved repo root made the C2 cross-document arm
enumerate nothing whenever the tree is reached through a symlink (`/var` → `/private/var` is enough), and
the guard reported a clean pass over a check it never ran. **That is the eight-times-caught defect class
this milestone exists to fence, arriving inside the fence built to census it.** Fixed and commented at
the site.

## Side finding — pre-existing at rext HEAD, and it was hiding 14 tests

`test_repair_reach_guard.py` put its `__main__` guard at `:328` with **two test classes defined below
it**, so `python3 test_repair_reach_guard.py` collected 16, skipped 14, and printed **OK**. §5 rule 8 —
*a check that SKIPS reads exactly like a check that PASSES* — arriving inside the test file of the fence
built to state its own denominator. `test_test_collection_fence` had been **RED on this since iter-114**
and the RED was invisible because the full suite does not complete on this host. Proven pre-existing by a
read-only `git archive HEAD` run before touching it. Guard moved to EOF; direct run now collects **30/30**.

## Tests — invocations stated

| suite | result |
|---|---|
| `test_corpus_citation_guard.py` | **19 / 19 OK** (7 mutation · 3 anti-vacuity · 6 regression · 3 exit-code) |
| `test_repair_postcondition.py` | 27 / 27 OK |
| `test_repair_postcondition_audit_mode.py` | 25 / 25 OK |
| `test_fence_provenance.py` | 34 / 34 OK |
| `test_guard_family.py` | 41 / 41 OK |
| `test_test_collection_fence.py` | 8 / 8 OK — **was RED at HEAD** |
| `test_corpus_index_guard.py` | 16 / 16 OK |
| `guard_family.py` live | **14 GREEN · 0 RED · 6 not-run-and-named** |

All via `python3 -m unittest discover -s tests -p "<file>.py"` from `stack-core/`, `ROSETTA_ROOT` set.
The full-suite hang (`FIX-M257x-iter108-stackcore-suite-hangs`) is unchanged and still open, so these are
**scoped runs and are named as such** — a scoped pass is not a whole-suite pass.

## `state.md` — fixed the way the contract says, by relocation

It stood at **15,161 B against a 15,360 B cap**, ~200 B from breaching. Measured per-field against
`context.md` § state.md contract: **`phase:` was 2,230 B against a 900 B budget** and
`active_milestone:` 486 against 400 — the frontmatter was 3,192 against 2,600. `phase:` was carrying the
whole iter-116 reading, **every phrase of which is already in the milestone's `progress.md`**. Replaced
with the verdict plus the pointer, per rule 3. Now: `phase:` **895 / 900**, every field inside its
budget, **file 13,431 / 15,360 — 1,929 B of headroom** where there had been 200. No prose was trimmed;
the narrative was moved to its owner.

## Close — 2026-08-07

**Outcome:** `TOK-08` class 1 censused and closed. **1,520 intra-corpus citations enumerated over 92
source documents** (C1 1,337 · C2 179 · C3 4, denominator `corpus-derived-per-arm`); **8 already false**,
all in the anchor arm; **8 of 8 repaired**; the fence stands green over **100 % of the enumerated
population**, registered in the postcondition ratchet at **zero** sites rather than a tolerated count.
Four census passes preceded the first repaired line and three of them existed only to kill a
false-positive class worth ~460 sites. The census's own mutation control caught a **silent vacuity bug in
this guard** on its first run. Reported against the iter's pre-registration: **the machine-reachable half
of class 1 is largely disjoint from the 10 predicates the reading booked** (4 of 387 bare-pin lines are
resolvable), so little `P` movement should be expected from it — early evidence bearing on `TOK-08`'s
refutation branch, stated before the reading that will grade it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is a tik; `TOK-08` is the USER's
re-scope, recorded, not an agent-authored triggered tok) — (3) re-scope: n (`TOK-08`'s stop condition is
**not gradeable** until the full mechanical sweep is followed by a reading; grading it now on one class
would be the flattering reading) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-117-1` (the anchor arm is where the class lives — and no reader exercises it) ·
`D-M257x-117-2` (the machine-reachable half of class 1 is largely disjoint from the sampled half) ·
`D-M257x-117-3` (under-flag rather than false-RED — four measured exclusions, all regression-tested)
**Side-deliverables:** `test_repair_reach_guard.py`'s `__main__` guard moved to EOF — 14 tests had been
hidden from direct execution since iter-114, and `test_test_collection_fence` had been RED on it that
whole time. Proven pre-existing by read-only `git archive` before the fix. Separate concern from this
iter's planned scope; does not change the close status.
**Routes carried forward:**
- **`FIX-M257x-iter116-intra-corpus-miscitation-is-the-largest-class` → PARTIALLY closed.** Its
  *resolution* half is now censused and fenced at 1,520 sites. Its *construct* half is measured
  **not machine-reachable at scale** and stays open, now with a number attached (4 of 387).
- **iter-118 = `TOK-08` class 2** — platform-source citation resolution, the mechanically-decidable
  subset of band #8's 14 of 37.
- `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block` → open
- `FIX-M257x-iter116-induction-fences-do-not-scale` → open
- `FIX-M257x-iter113-adjudication-is-judgement` → open
- `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` → open
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter108-stackcore-suite-hangs` → open; every count above states its invocation because of it
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open, delivered-unfixed
**Lessons:**
1. **Measure a fence's false-positive rate before it is a fence, not after.** Four census passes; three
   existed only to kill a false-positive class, together worth ~460 sites. Every exclusion in the shipped
   guard is annotated with the number it cost, because a rule adopted by reasoning is a rule nobody can
   audit later.
2. **A defect class hides in whichever arm no human exercises.** All 8 were broken `#anchor` fragments — a
   reader following one still lands on the right file, notices nothing, and the corpus keeps the error.
   0 of 1,337 path citations and 0 of 4 line-pins were wrong. **Ask which part of a citation a human
   never checks; that is where the census pays.**
3. **State up front which half of a sampled class a census can actually reach.** The readings booked
   *construct* defects; the machine can only reach *resolution* ones. Discovering that after the grading
   reading would have read as the method failing, when it is really the method's scope being narrower
   than the class.
4. **`state.md` breaches are always one field.** Per-field measurement found `phase:` at 2,230 against a
   900 budget, carrying a reading already written down in `progress.md`. Trimming prose elsewhere would
   have bought one run; moving the narrative to its owner bought 1,929 B.
