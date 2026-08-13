# iter-107 — closeout

**Shape:** tik · `iter_shape: fence` · **`TOK-06` step 2** — the induction checks.

## The one-line answer

**The repair loop's own largest induction shape is now fenced at the commit — and replaying iter-102's real
commit surfaces all four `:321` citers, including `backend.md:54`, which the 14-seat double reading missed
in BOTH passes.**

## The shape, and why it needed a fence rather than a rule

§5 rule 34 already says *line numbers move when YOU edit too*. It has now failed to stop the same defect
twice, one cycle apart:

| repair | booked by | what happened |
|---|---|---|
| iter-100 (`a229f8d`) | iter-101 | a two-line parenthetical pushed a `service_taxonomy.md` table down two rows; a note that was exactly correct came to cite Chronos and Intelligence |
| iter-102 (`cd16967`) | iter-103 | a production-topology block moved the local-stack wording `321 → 331`; **4 sites still cite `:321`**, now naming the opposite topology |

`TOK-06` counted this class at **21 % of the residual**. A rule that has been written down and violated
twice is not a rule that needs restating.

## Landed

`stack-core/anchor_offset_guard.py` — commit-scoped, like `repair_leak_guard`, because **the defect is only
decidable at the commit**: looking at `:321` today tells you what is on line 321, not that a citer meant
something else. The line map comes from `git diff -U0`, the authoritative record of what moved where — not
a re-scan, which cannot tell "this line moved" from "this line was rewritten in place."

## The design changed TWICE in flight, both times because a control refuted it

**This is the part worth reading.** Both of the iter's declared escalation conditions fired.

**First cut — GREEN on the commit that motivated it** (`D-M257x-107-1`). It waived any citation in a file
the commit had touched: *the repairer opened that file, so assume they looked.* Sound, and wrong — iter-102
was a **98-site repair** that modified all three citing service docs while editing *other* claims in them.
It graded 2 citations and passed. The carve-out is now **line-level**: the commit must have written the line
the citation is on. **General form: a waiver keyed on the unit the defect hides inside will waive the
defect.** The defect hid inside a file; the waiver was per-file.

**Second cut — RED on a CORRECT repair** (`D-M257x-107-2`). Narrowing to *"the commit authored this citation
and shifted that target line"* caught all four `:321` citers — and a synthetic control caught it back: a
repair that correctly re-points `:7 → :9` after inserting two lines above shifts too, and was graded RED.

The two are **not distinguishable**. Given a number authored beside a shift, *"post-move and correct"* and
*"pre-move and stale"* are both consistent with everything the diff records; only intent separates them, and
**intent is not in the repository.** A third narrowing was tried — *does the number land inside text the
commit ADDED?* — and it lost the real case, because iter-102's insertion was at line 265 and new-line 321 is
pre-existing content that merely moved.

**Resolution: reported, counted, and excluded from the exit code.** The summary line and the OK line both
carry the CANNOT-TELL count, and the OK line states that the green **does not cover them** — §8's *grade
the cannot-tell* (iter-91). **A fence must not assert what it cannot decide**; the alternative was a RED
that correct repairs trigger, i.e. a fence that gets turned off.

## The answer key is the commit, not a fixture (`D-M257x-107-3`)

```
cd16967  (iter-102)   5 ROT + 5 CANNOT-TELL
   CANNOT-TELL  CLAUDE.md:282 · backend.md:54 · jobsimulation.md:146 · sentinel.md:85
                 all cite architecture_overview.md:321, moved to 331
   ROT          service_taxonomy.md:147 · services/README.md:17 · platform-alignment.md:1446
                 platform-alignment.md:499 (hiring.md 93 -> 107)
                 shared_libraries.md:96   (storage.md 115 -> 129)

a229f8d  (iter-100)   1 CANNOT-TELL
   service_taxonomy.md:131 cites :137, moved to :139   ← iter-101's booked induction
```

**`backend.md:54` is the headline.** iter-103 recorded that the double reading *"found 2 of the 3 in-scope
sites and missed `backend.md:54`, which sat inside seat E's own file set in both readings."* The fence
surfaces it mechanically, at the moment the commit was made, alongside the three the reading did find and
the one in `CLAUDE.md` that is outside the gate's file scope entirely.

**And it surfaced 5 ROT findings nobody had booked at all** — `hiring.md 93 → 107` and
`storage.md 115 → 129` are ordinary offset drift in `platform-alignment.md` and `shared_libraries.md` that
no reading has ever named.

Commit shas are **pinned** in the test, per §5 rule 25: an answer key that drifts with HEAD stops being an
answer key.

## Citations are read at the range's END revision (`D-M257x-107-4`)

The first implementation read the working tree — correct for `HEAD~1..HEAD`, **wrong for the answer-key
runs**, which would grade today's citations against a two-week-old diff. The measured difference on
`cd16967` was real (36 → 33 citations seen, 7 → 5 ROT), so this was not hygiene. It is §5 rule 41a one level
down: **an instrument that resolves against "now" cannot grade a measurement taken "then".**

## Controls

**18 tests, all green.** Two real commits replayed · six synthetic shapes each separating a case the guard
must distinguish (insert-above = RED · append-below = GREEN with a positive control on the graded count ·
citer re-points its own line = waived **and** landed in CANNOT-TELL · authored-beside-a-shift = CANNOT-TELL,
never a verdict · authored-into-an-unmoved-region = silent · deleted cited line = RED · uncited doc = GREEN)
· refusals (**no corpus = exit 2**, **an empty range = exit 2 with `Nothing was checked`, never a pass** —
§5 rule 8; exit 1 on findings, 0 under `--report`; JSON carries every denominator) · anti-vacuity against
the **live** corpus (≥ 25 intra-corpus citations found, and **0 unresolved** — the `README.md` ×6 ambiguity
is resolved by preferring the citer's own directory, not by dropping it) · family placement reconciles both
ways.

## Not taken here, and routed with its measurement (`D-M257x-107-5`)

TOK-06 step 2 named **two** shapes. The canonical-wording one is **re-confirmed live at this open** — the
`:8081` literal has **1 occurrence in `app` and 3 in `stack-demo/rosetta-extensions`**, a repo the
sentence's own 13-repo denominator counts, so it is still self-refuting at 2 corpus sites. It is the iter's
third line and the scope-creep tripwire fires; it is also a harder fence (recognising "a canonical wording"
needs a hand-maintained registry — §2's own warning — or cross-document near-duplicate detection).

## The finding this iter did not go looking for — and it landed on iter-106's own fence

`clone_drift_guard` went **GREEN** at this close, one iter after shipping RED. **Nothing was repaired.**

The §8 section written to *document* the RED contains the sentence *"`sentinel` at `f2c46190`, 2 commits
past the newest sha the corpus cites"* — and that backticked sha **is** a corpus citation of sentinel's
HEAD. The fence's assertion, *"the corpus's most recent knowledge of this repo predates its HEAD"*, is now
literally false. **Five stale sites are still stale.**

This is the guard's stated reach behaving exactly as documented — and it is still the sharpest thing
measured today, because it shows how cheap satisfaction is: **writing about the drift satisfies the drift
fence.** The honest reading of a D1 green is *no repo advanced past every sha the corpus **mentions*** —
mentions, not verifies.

**Deliberately NOT patched here.** An exclusion list of "docs that are about fences" is §2's
hand-maintained tuple again, and inventing one at the end of a long session is exactly how a fence acquires
an exemption nobody can later justify. Instead: recorded in the guard's own docstring, **pinned by a
known-limitation test** (§8 rule 7 — the assertion is what expires), and routed as
`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`.

**Fourth time in three days the milestone's class has landed on the milestone's own apparatus**, and the
fourth time what caught it was **re-running rather than reasoning**.

## Gate

**Unchanged at 4 of 5.** No `N` movement claimed — clause 3's instrument, never clause 5's. Clause 5 not
re-cut, narrowed, reinterpreted or argued.

## Housekeeping

Zero platform-repo edits. `stack-demo/**` untouched, no clone fetched, no tag cut. rext on `main`.

## Close — 2026-08-06

**Outcome:** the line-offset induction shape is fenced at the commit; replaying the two real commits that
caused the two booked incidents surfaces both, including the citer both reading passes missed, plus 5
never-booked ROT findings. Two designs refuted by their own controls before the third landed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (3 tiks) — (6) protocol-stop: n — (7) budget-exhausted: **see the report** —
Outcome: continue
**Decisions:** `D-M257x-107-1` (the file-level carve-out returned GREEN over the defect — a waiver keyed on
the unit the defect hides inside will waive it) · `-2` (the class is genuinely undecidable; reported, not
asserted) · `-3` (the answer key is the commit, not a fixture) · `-4` (citations read at the range's END
revision) · `-5` (the canonical-wording shape routed with its measurement)
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter107-canonical-wording-fence`** *(net-new)* — TOK-06 step 2's second shape, live and
  measured: `:8081` = 1 in `app` + 3 in `stack-demo/rosetta-extensions`, against its own 13-repo
  denominator.
- **`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`** *(net-new, and it is against THIS RUN's own
  iter-106 deliverable)* — one mention of a repo's HEAD anywhere in `corpus/**` reconciles it, including a
  mention inside the prose reporting the drift. Measured live: `clone_drift_guard` RED at iter-106's close,
  GREEN at iter-107's, zero repairs in between. Pinned by a known-limitation test.
- **`FIX-M257x-iter107-unbooked-rot`** *(net-new)* — the 5 ROT findings on `cd16967` that no reading ever
  named, incl. `hiring.md 93 → 107` and `storage.md 115 → 129`. Fold into step 3.
- **TOK-06 step 3 — repair the 33**, which must now clear **two** fences: `clone_drift_guard` (5 sentinel
  sites) and `anchor_offset_guard` on its own commit. **Run the offset guard AFTER the repair commit** —
  §8's iter-93 rider: the scope is the DIFF.
- Unmoved: `FIX-M257x-iter105-claimtwin-green-twin-refire`, `FIX-M257x-iter56-assignment-flake`,
  `FIX-M257x-iter103-assignment-context-bleed`, `DEF-M257x-iter103-aws-bind-provenance`,
  `DEF-M257x-iter101-briefing-rext-tree`, `RF-2/3/7–14`, the five pass-22 items.
**Lessons:** **two designs were killed by this iter's own controls, and both would have shipped green.**
The first passed the commit it was written for; the second failed a commit that was correct. Neither was
caught by reading the code — both were caught by *running the fence against a case whose answer was already
known*. Generalisable, and it is the sharpest version of §8's watched-going-RED rule: **a new fence's first
test should be the real defect that motivated it, and its second should be a real NON-defect of the same
shape.** One proves it can fire; the other proves it can stop.
