# iter-281 — make the control tree runnable, then attribute the RED

**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — the re-survey changed the plan before any work started

Three targets were named. **Two were already closed**, and only re-verifying found that out:

| named target | re-survey verdict |
|---|---|
| two map rows call a destroyed service `live-standalone` | **CLOSED** by iter-280 `b5bd4e7e` — both prod cells read `decommissioned` |
| `services.sh` cites `arm E`, a fence that never existed | **CLOSED** by iter-280 — `stack-verify/lib/services.sh:164-168` carries the retraction |
| `FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT` | **OPEN** |

Third consecutive iter where a route list was stale. The rule is now load-bearing, not advice.

## Phase A — `FIX-M257x-278`, closed as a CLASS, and the route was wrong about both halves

**Wrong about the size.** The route said *"the census substrate"* and meant two files. Censused, the
population is **six sites in three files**; the third — `stack-verify/tests/test_scope_union_m257x.py` —
is a near-copy of the first that no route mentioned.

**Wrong about the failure mode, and this is the part that matters.** The route said such a test *"does
not survive cloning"*, i.e. fails loudly elsewhere. Measured: a control clone at `/private/tmp/…/rosetta`
ran `test_claim_census_substrate_m257x.py` and reported **every arm GREEN** — the hardcoded root still
exists on this box, so the test read the ORIGINAL tree from inside the clone and graded a subject nobody
had selected.

> **An instrument that names one machine's absolute root does not break on another tree; it reports on a
> tree it was not pointed at.** That is strictly worse than a crash, and it is why iter-280's control-tree
> attribution could not have been trusted even where it did collect.

Repaired by derivation (`parents[4]` / `dirname(dirname(REXT))`, the rosetta root for **both** rext clone
roles, with `ROSETTA_ROOT` overriding) — the idiom fourteen sibling modules already used. Proven
behaviourally: on a fresh clone the root now resolves to the **clone**.

### The new fence caught itself, which is the sixth instance this milestone

`stack-core/tests/test_absolute_root_census_m257x.py`, 7 arms, RED-proven before the repair. Drafted
line-grained it fired on **nine** lines — the six real ones plus three that are not defects:
`buildbench.py:8` (a comment citing another host), `test_baseline_mirror_fence.py:362` (a fixture string),
and **line 15 of its own docstring**.

**Writing a fence for *instrument-inside-its-subject*, the first draft committed that exact defect.** It is
recorded rather than quietly deleted. The predicate is now a **parsed construct**: a home-root literal is a
finding only where it RESOLVES — a path call, or a path-named binding, directly or through a list.
**Quotation is not coupling.**

Controls: denominator anti-vacuity; a staged offender killed in each of the four resolving constructs; the
three first-draft false positives pinned as a regression; portable absolute paths (`/usr/bin/python3`,
`/dev/tty`) proven NOT flagged so the fence can never prescribe deriving the interpreter path; the single
generated-output exclusion proven to work by NAMED directory; shell/Go readers exercised on a staged tree
because their real population is empty today. Cost measured: **0.93 s**.

## Phase B — the attribution instrument, and iter-280's residual did not reproduce

Two whole-section runs, one per tree, plus two cheap hypothesis probes.

| run | tree | result |
|---|---|---|
| control pair @ HEAD | frozen clone, no `stack-*` workspaces | **2 failed / 2180 passed / 53 skipped**, 843.31 s |
| live | working tree, post-Phase-A | **7 failed / 2232 passed / 3 skipped**, 2266.86 s (37m46s) |

**The 13 failures across four files that left iter-280's gate RED did not reproduce.** All four —
`test_suite_census`, `test_repair_leak_guard_mutation_battery`, `test_suite_census_population`,
`test_battery_stage` — are **green inside the full live run**. Two candidate causes were probed and
**refuted** before the full run: a `test_progress_beacon` env leak (green in company), and a
battery-family interaction (**105 passed** with the batteries and the census together, 19m14s).

**What can and cannot be concluded, stated separately.** The residual is not a standing test-isolation
defect — three independent measurements now disagree with that reading. What it *was* is not established:
the most economical account is contention (these arms grade subprocess output, and the batteries spawn
many), but nothing here measured it, and a mechanism nobody measured is a guess. The `test_suite_census`
instrument-inside-its-subject observation stands as a **structural** criticism of the arm; it is not the
proven cause of iter-280's RED.

### The 7 live failures are TWO root causes, not seven

- **`blocking_state_guard` RED → 2 failures.** Its own arm, plus
  `test_fence_provenance::TestFamilyRefusesAnUnstateableTree::test_the_escape_accepts_and_records`, which
  runs the whole guard family and expects rc 0 — a **cascade**, and reading it as an independent defect
  would have sent an iter after a fence that was working perfectly.
- **`TEST_MODULE_LITERAL_CEILING` breached +2 → 5 failures.** Mine.

## Phase C — the three repairs

**1. The live section-gate RED, and it is a recurrence with a rate.** iter-280 graded `user-blocker` and
`deferrals-audit.md` recorded neither the grading nor the derived zero-claim. Recorded now; the guard
prints `OK`. **This is the second consecutive occurrence** — iter-259 reached the table ten iters late,
iter-280 one. The mechanism is structural, not carelessness: the grading is written in the iter's `##
Close` section and the audit is a different file in a different lane, so **the write that creates the
finding and the write that would represent it are never the same edit**. The fence is not the gap; running
it is — and a close that costs 38 minutes of whole-section suite is exactly the one nobody can afford.

Fixing it surfaced a second defect in the same file: **§10 still asserted *"zero open user questions"***
while §8 had retracted that claim at iter-121. One section repaired, the class left — in the milestone's
own close-gate document. §10 is the hand-off table a close gate reads to decide whether a lane is done.

**2. The ratchet, +2, both from one sentence written twice.** Removed by **rephrasing, not by bumping** —
back to 653, at the ceiling. iter-280 set that standard by deleting growth; raising a ceiling to admit
one's own prose is how a ratchet stops being one.

**3. A precondition that guards nothing.** `node_modules_present` finds a tree behind a dot-directory —
its own docstring names `.agentspace/…/e2e/node_modules` as the depth-5 case it was *fixed* to reach —
while `basename_index`'s walk prunes `not d.startswith(".")` and can never see it. So on a clean control
tree the guard said *"present, proceed"*, the derivation legitimately pruned nothing, and the arm failed
with *"it is broken, not the tree"* **about a derivation that was fine**.

Censused, not repaired at the site: **1 of 4** call sites is mismatched — the other three pass the REXT
root, where the e2e tree has no dot segment. And the live green was never about the repo: the only
dot-free trees here are under `stack-demo/`, a **git-ignored stack workspace**. A green a fresh checkout
cannot reproduce.

`node_modules_reachable` asks the question the measurement asks, **registered in the fresh-checkout
ratchet's `PREDICATES`** so converting the site cannot silently drop it out of that population. A staged
arm carries the anti-vacuity with **no box dependency at all** and pins the discrimination the live arm
cannot. **Its own first staging put everything under the dot parent**, so `files_kept` read 0 and the arm
called a working derivation broken — the fixture, not the subject, for the third time this milestone.

**4. What deriving `ROOT` then exposed.** Two arms in `TheModulesOwnCommentFiguresAreDERIVED` assert
figures over a census whose population **includes the clone set** (wrong-repo class 1 with clones, 0
without; the three-figures-differ anti-vacuity collapsing to `[41, 0, 0]`). While `ROOT` was absolute they
could never surface — from a clone they read the original tree, found the clones, and passed. **The
derivation turned a silent false GREEN into an honest RED**, and the honest RED said what was always true.
Both classes now declare `clone_set_present` — **including `TheBasenameShareIsDERIVED`, which did not go
RED**: a green without clones is a coincidence of this corpus, not a property of the arm, and declaring
only the one that failed would be repairing the site again.

**Fresh checkout of the pair, measured: 77 passed, 8 skipped, 0 failed** — every skip naming its missing
precondition. Before this iter the same modules reported a full pass while reading a tree nobody selected.

## Phase D — the final measurement, and the ratchet caught the iter that had just repaired it

Three whole-section runs on the live tree, and the middle one is the interesting failure:

| run | result | what it proved |
|---|---|---|
| 1 (pre-repair) | **7 failed / 2232 passed / 3 skipped**, 37m46s | 7 arms = **2** root causes |
| 2 (post-repair) | **5 failed / 2235 passed / 3 skipped**, 31m47s | blocking-state RED and its cascade GONE; residual is `DOCSTRING_LITERAL_CEILING` **+1, mine** |
| 3 (final) | see close | taken with **no concurrent edits** |

**Run 2's residual was written by this iter, after this iter had already repaired a ratchet and verified
it green.** The +1 was a docstring restating a file count that the sibling function two definitions above
already carried. Nothing was wrong with the fence or with the first repair — **the verification was
correct about a tree that no longer existed by the time it was cited**. That is the milestone's own
staleness class, committed against a fence rather than a document, by the iter that had just repaired
that fence. Removed by deleting the duplicate, **not by bumping**: all three ceilings now read `exact +0`
(240 / 236 / 653).

**Disclosed, because it bears on what run 2 is evidence of:** this iter wrote its own iter-dir files
while run 2 was mid-flight, so run 2's guard-family reading is not clean evidence about the tree it was
started on. Run 3 is the authoritative one. Stated rather than quietly relied upon — a measurement and an
edit must not overlap.

**Protocol updated in the same iter** (`corpus/ops/platform-alignment.md` §10, four rules): the absolute-root
class and its inverted symptom; *a precondition must ask the question its measurement asks*; *a failure
count is not a defect count* (with the explicit note that refuting a hypothesis does not license
inventing a replacement); and *a mid-iter green expires the moment you keep editing*.

## Close — 2026-08-11

**Outcome:** the section gate is **GREEN — `2240 passed, 3 skipped, 0 failed` (37m50s)**, on a run taken
with no concurrent edits. `FIX-M257x-278` is closed **as a class** (6 sites in 3 files, not the 2 the
route named) with a construct-graded fence at zero, and the route's own premise is retracted: a hardcoded
absolute root does not crash on another tree, it **reads the original one and reports green**. iter-280's
13-failure residual **did not reproduce** and is no longer believed to be a standing isolation defect —
what it *was* is stated as unestablished rather than guessed. Three ceilings at `exact +0`, none bumped.
**Clause 5 NOT re-measured and no `P` is claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**

**Why exit-7 and not exit-4.** The orchestrator had already ruled on iter-280's escalation (*fix the red
gate, do not re-raise it*), and nothing this iter found needs a user decision: every RED was measured,
attributed and repaired inside the iter. The session then ran ~4.5 h across an infrastructure
interruption and three 32–38 min whole-section runs, and stops at a clean boundary — **iter closed, both
trees committed, tree clean**. That is `budget-exhausted (between iters)`, not a blocker.

**Decisions:** `D-M257x-281-1` … `D-M257x-281-11` (see `decisions.md`), including two recorded
self-defects: a fence whose first draft committed the very class it fences, and a ratchet re-breached by
this iter *after* it had repaired and verified that same ratchet.

**Side-deliverables:**
- `corpus/ops/platform-alignment.md` §10 — four protocol rules, landed in the same iter as their lessons.
- `deferrals-audit.md` §10 — a stale *"zero open user questions"* row that contradicted §8's own iter-121
  retraction, found while repairing the §12 gap rather than looked for.
- `.agentspace/rext.tag` re-pointed (it read `iter-279` while `iter-280` was already on origin).

**Routes carried forward:**
- **`ROUTE-M257x-281-rext-tag-SoT-has-no-fence`** — `rext.tag` has now been stale twice with nothing
  grading it. The fence (tag exists on origin AND is not behind HEAD) was this iter's **third** line of
  investigation, so the tripwire routed it rather than absorbing it. Supersedes the reporting half of
  `ROUTE-M257x-278-rext-tag-SoT-was-six-iters-stale-unnoticed`.
- **`ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes`** — **strengthened, not
  advanced.** Measured three times this iter: 37m46s / 31m47s / 37m50s. This is the direct cause of the
  blocking-grading recurrence rate (`D-M257x-281-9`): the fence that would catch it in-iter is the one
  nobody can afford to run. A fast subset (ratchets + censuses, seconds) remains the fix.
- **`ROUTE-M257x-281-suite-census-is-structurally-inside-its-subject`** — supersedes
  `ROUTE-M257x-280-suite-census-is-a-member-of-its-own-population`, **downgraded from a proven cause to a
  standing structural criticism**: the arm is green in company across three runs, so it is not what made
  iter-280 RED.
- Unchanged: `ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence`,
  `ROUTE-M257x-h70-quotation-verification-instrument-is-unreliable`,
  `ROUTE-M257x-279-durations-are-unclassified-measurement-nouns`,
  `ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`,
  `ROUTE-M257x-274-successor-half-is-uncovered`, `ROUTE-M257x-274-tie-order-is-unstable`,
  `FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`,
  `ROUTE-M257x-h59`, `ROUTE-M257x-h65`, the fence half of `ROUTE-M257x-277`.
- **CLOSED this iter:** `FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT`,
  `ROUTE-M257x-280-map-state-tokens-are-graded-against-nothing` (found already closed by iter-280).
- **Clause 5's semantic reading is still unmeasured** (last: iter-131, `P = 29 / N = 47`, a floor).

**Lessons:**
1. **A route describes the failure it TRIPPED OVER, not the failure that exists.** This one predicted a
   crash; the real symptom was a silent green about the wrong tree — the opposite, and worse.
2. **A failure count is not a defect count.** Seven arms were two defects; one of the seven named a fence
   that was working perfectly.
3. **A mid-iter green expires the moment you keep editing.** Run the ratchets last, not when convenient.
4. **Refuting a hypothesis does not entitle you to a replacement.** The 13 did not reproduce; that is
   what was measured, and "contention" stays a guess until someone measures it.
5. **Declare the precondition on the member that did NOT go red, too** — otherwise you have repaired the
   site again and called it a class.
