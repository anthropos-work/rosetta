# iter-107 — decisions

## `D-M257x-107-1` — the file-level carve-out returned GREEN over the defect the guard was written for

The first cut waived any citation living in a file the commit had touched. The reasoning was sound: *the
repairer opened that file, so assume they looked* — and a fence that reddens on every legitimate repair
cannot be satisfied and gets suppressed.

**Run against `cd16967`, iter-102's own commit, it graded 2 citations and returned GREEN.**

The refutation is in the incident itself: iter-102 was a **98-site repair** that modified `sentinel.md`,
`jobsimulation.md` **and** `backend.md` — while editing *other* claims in them — and left all three `:321`
citations stale. A file-level carve-out waives exactly the case that motivated the guard.

The carve-out is now **line-level**: the commit must have written the line the citation is on. Kept as a
test (`test_the_FILE_LEVEL_carve_out_would_have_missed_it`) so the premise cannot be quietly undone.

**The general form is worth more than the fix:** *a waiver keyed on the unit the defect hides inside will
waive the defect.* The defect hid inside a file; the waiver was per-file.

## `D-M257x-107-2` — the second cut went RED on a CORRECT repair, and the class is genuinely undecidable

Narrowing to "the commit authored this citation **and** shifted that target line" caught all four `:321`
citers — and a synthetic control immediately caught it back: a repair that **correctly** re-points a citer
from `:7` to `:9` after inserting two lines above shifts too, and was graded RED.

The two are not distinguishable. Given a number `X` authored beside a shift:

| hypothesis | consistent with the diff? |
|---|---|
| `X` is post-move and correct | yes |
| `X` is pre-move and stale | yes |

Only the author's intent separates them, **and intent is not in the repository.** A third narrowing was
tried — *does `X` land inside text this commit ADDED?* — and it lost the real case: iter-102's insertion was
at line 265, so new-line 321 is pre-existing content that merely moved, and the rule went silent on the
defect.

**Resolution: it is reported, counted, and excluded from the exit code.** The summary line and the OK line
both carry the CANNOT-TELL count, and the OK line states in its own words that the green **does not cover
them** — §8's *grade the cannot-tell* (iter-91), whose finding was that a partial skip is worse than a
total one.

**A fence must not assert what it cannot decide.** The alternative was a RED that correct repairs trigger,
which is a fence that gets turned off — and this milestone has the receipt for what a suppressed fence
costs (a silently-refused perf patch shipped a 76 s members grid for four releases).

The synthetic control that produced this is kept, with its reason in the assertion message: it is the
evidence that the class is **undecidable** rather than merely unimplemented.

## `D-M257x-107-3` — the answer key is the commit, not a fixture

`test_anchor_offset_guard.py`'s load-bearing tests replay **`cd16967`** (iter-102) and, in the survey,
**`a229f8d`** (iter-100). Both surface their own booked induction:

| replayed commit | what the guard surfaces |
|---|---|
| `cd16967` | **all four** `:321` citers as CANNOT-TELL — including **`backend.md:54`, which the 14-seat double reading missed in BOTH passes** — plus 5 decidable ROT findings |
| `a229f8d` | `service_taxonomy.md:131` → `:137`, moved to `:139` — iter-101's booked induction |

A fixture proves a fence does what its author expected. **A commit that actually produced the defect proves
it does what the milestone needed.** The commit shas are pinned in the test, per §5 rule 25 — an answer key
that drifts with HEAD stops being an answer key.

## `D-M257x-107-4` — citations are read at the range's END revision, not from the working tree

The first implementation read citations off the working tree. That is correct for `HEAD~1..HEAD` — the
normal use — and **wrong for the answer-key runs**, which replay historical commits: it grades today's
citations against a two-week-old diff, mixing two trees.

Now `citations()` takes the range's end revision and reads via `git ls-tree` + `git show`. The measured
difference on `cd16967` was real (36 → 33 citations seen, 7 → 5 rot findings), so this was not hygiene.

**It is the milestone's own §5 rule 41a one level down:** an instrument that resolves against "now" cannot
grade a measurement taken "then".

## `D-M257x-107-5` — the other induction shape is NOT taken in this iter

TOK-06 step 2 named two shapes. Only the line-offset one lands here. The canonical-wording shape — a
sentence published to ≥3 sites carrying a defect measured against its own stated denominator — is
**re-confirmed live at this open** (`:8081` has 1 occurrence in `app` **and 3 in
`stack-demo/rosetta-extensions`**, a repo the sentence's own 13-repo denominator counts) and routed
forward.

Reason: it is the iter's third line of investigation and the scope-creep tripwire fires. It is also a harder
fence — recognising "a canonical wording" needs either a registry (hand-maintained, §2's own warning) or
near-duplicate detection across documents, and neither is a 20-minute build. Routed with the measurement
attached, not merely deferred.
