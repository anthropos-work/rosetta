**Type:** tik · `iter_shape: reading`. **THE READING** — the first in eleven iters.

---

# `P = 30` · `N = 47` — `P` ROSE against a corpus that absorbed real repair

> **iter-119: `P = 22` / `N = 28`.  iter-131: `P = 30` / `N = 47`.**

**`P` rose 22 → 30 (+36 %) and `N` rose 28 → 47 (+68 %)**, over a subject that grew **9.7 %** and
absorbed 30 consequence-class repairs, 13 C1 claims, 46 complement fixes, the `ai` row, the tier-1
re-pins, and this run's own ~24 routed findings plus 10 predicate sites.

**The pre-registration barred the flattering reading of a fall, and it equally bars the despairing
reading of a rise.** What this reading can and cannot establish, stated in the order the sheet sealed:

- **`P` is a FLOOR over a suspicion-selected sample, not a corpus error rate.** No percentage is computed
  from it. A hunted sample over-states the population rate ~19× (census 0.70 %, 0.23 % substrate-corrected).
- **It does NOT establish that repair made the corpus worse.** iter-119 proved a single reading recovers
  roughly a third of a standing pool whose union floor was already **≥ 46**. A rise from 22 to 30 is
  entirely consistent with *the same pool, sampled differently*, and this reading cannot separate that
  from *new defects*.
- **What it DOES establish, and this is the result:** the standing pool is **larger than any single `P`
  has ever shown**, and eleven iters of repair did not drain it. **The new union floor is ≥ 46 + the
  0 overlap below** — see the test-retest finding, which is the sharpest thing on the sheet.

## The test-retest measurement, and it is worse news than the primary

**Band #3 asked how many of iter-119's 22 predicates this reading re-books. The answer is ZERO
identifiable matches.**

None of iter-119's eight enumerated intra-corpus-citation predicates appears in this reading's 30. Its one
individually-named platform-facing predicate (`clerk-integration.md:40`, the sign-in-token
understatement) **was repaired** in the interval and is correctly absent.

> **Two consecutive readings of overlapping corpora produced almost disjoint predicate sets.**
> iter-116 → iter-119 overlapped **13 of 37 (35.1 %)**. iter-119 → iter-131 overlaps **~0 of 22.**

**And a denominator problem I have to disclose rather than paper over:** `iter-119/adjudication.md`
**does not enumerate its 22 predicates individually** — it lists 8 by anchor and characterises the other
14 by class. So the overlap could only be computed against **9 of 22**, and the "0" is exact only over
that subset. **A test-retest metric that its own prior sheet cannot support is a metric this milestone
has been reporting without the substrate to compute it.** Routed as
`FIX-M257x-iter131-predicate-sets-not-enumerated` — every future reading must publish its full predicate
list or the band is uncomputable.

## The largest cluster — 19 of 80 blockers, six seats, one root cause

The corpus says `infrastructure` **"has never been in any clone set"** and that `cms`'s production state
is **"NOT MEASURABLE — do not assert either way."** Measured directly:

| | |
|---|---|
| sites citing a READ of `infrastructure` @ `13c248e6` | **28** |
| sites still publishing the never-cloned / unmeasurable hedge | **11** |

iter-123 cloned it and read it. `org-repos.md:102` says so in a heading:
**"🔓 `cms` M810: SETTLED — the ECS service is DESTROYED. It has now been read."** That one read settled
**four** standing questions — the production service set is exactly ten modules, and `cms`,
`roadrunner`, `graphql-wundergraph` and `messenger` are all **orphaned**, their service-repo terraform
describing modules the root never instantiates.

**The measurement was recorded where it was made and reached almost nothing else — rule 54 at scale.**
Eleven sites across six files still publish the superseded limit, **and `CLAUDE.md` publishes it too**, so
every agent that loads this repository starts from the retracted claim. **This is the single highest-value
repair available to the next iter**, and it is routed, not taken, because condition 3 bars repair inside a
measuring pass.

## Three of the 30 are defects in prose I wrote — two of them THIS RUN

| | |
|---|---|
| **P7** | I added `library-unimported` to the guard's `ALLOWED_STATES` at iter-130 and changed assertion C's row to say **"nine"** — and never added the row to **§1's own state table**, which still defines eight. **A vocabulary change that reached the checker and not the definition.** Caught by two independent seats. |
| **P19** | Repairing a drifted citation at iter-130 I wrote *"all three sites are the literal `curl …/sign_in_tokens`"*. `staging-bringup.md:528` is a **prose bullet**. **I over-claimed in the very sentence whose purpose was to make a citation robust.** |
| **P5** | iters 129–130 repaired the `ai` row in the fenced map and `shared_libraries.md`; **`architecture_overview.md:83` still lists `ai` among "four imported private modules".** My own rule-54 sweep did not reach it. |

**P5 and P6 together are the sharpest structural finding:** my new assertion G prints the true module set
on **every run** — `analytics-go, colony, proto, storage, taxonomy` — and the corpus prose contradicts its
own fence in two places. **A fence that prints the right answer does not correct the prose beside it.**

## Bands: 7 HELD of 15

Prior: 4/9 · 3/7 · 5/9 · 4/10 · 7/13 · 7/14 · 13/15. **This is the worst band performance since iter-110,
and the failures are informative rather than embarrassing** — six of the eight failed *high*, i.e. the
corpus returned more than predicted on almost every axis.

| # | prediction | band | measured | verdict |
|---|---|---|---|---|
| 1 | per-reading in-scope upheld blockers | [6, 30] | **~31 · ~37** | ❌ high |
| 2 | union `N` | [14, 40] | **47** | ❌ high |
| **2p** | **union `P`** — primary | [12, 34] | **30** | ✅ |
| **3** | overlap with iter-119's 22 | [3, 12] | **~0** (of the 9 comparable) | ❌ low |
| 3b | within-reading `m` as share of union | [10 %, 55 %] | **~60 %** | ❌ high |
| 4 | adjudicator upheld rate (raw) | [72 %, 94 %] | **89.5 %** | ✅ |
| 5 | two passes' recalls differ by | ≥ 12 pts | **~6.6** | ❌ (4th consecutive) |
| 6 | wrong-tree rejections | [0, 3] | **0** | ✅ |
| **7** | intra-corpus citation defects | ≤ 5 | **10** | ❌ (3rd consecutive, at a tighter cut) |
| 8 | platform-drift share | [20 %, 60 %] | **26.7 %** | ✅ |
| 9 | per-seat booked spread | ≤ 8 | **10** | ❌ |
| **10** | repair-induced | [1, 8] | **3** | ✅ |
| 11 | anchors per predicate | [1.00, 1.45] | **1.57** | ❌ high |
| 12 | anchors in a multi-pin block | [2, 12] | **~9** | ✅ |
| **13** | upheld blockers in `org-repos.md` | [0, 4] | **3** | ✅ |

**Band #7 is the pre-registered finding.** Cut DOWN to `≤ 5` after two consecutive failures, precisely so
that a third failure would mean something — and it measured **10**, the highest yet, and **a third of the
pool**. `D-M257x-117-2` is confirmed for the third reading running: **the machine-checkable half of
intra-corpus citation and the half a reader books are different halves.** The census closed resolution at
100 % reach and zero findings; the *construct* half is untouched by any fence and is now the largest
single class.

**Band #10 held, and it is the first reading in which repair-induction was measurable at all.** Band was
re-cut UP to `[1, 8]` on a +1,248-line exposure; measured **3**, all three mine.

**Band #5 failed for the fourth consecutive reading.** The pre-registration said a fourth failure means
the mechanism is wrong rather than the tuning. **It is retired here** rather than re-cut a fifth time.

### The upheld rate, reported TWICE as required

**Raw: 68 / 76 = 89.5 %.** **`wrong-tree`-separated: 89.5 %** — *identical*, because there were **zero**
`wrong-tree` rejections, for the seventh consecutive reading, despite the widest rext-tree gap yet
(33 commits) and a **dirty `ant-academy` working tree**. The addendum's two-tree rule is doing real work.

## What I could not settle, and did not launder

**4 blockers on the root-mounted route count** (`security_compliance.md:250`,
`architecture_overview.md:406`). Two seats measured **eight** against a published **seven**. **I could not
settle it.** My first three counting attempts returned **0** — a wrong receiver-name regex, then a
`head`-truncated listing that made `internal/web/web.go` look absent when it exists. *An empty result from
a failed command is not evidence of absence*, and it caught me twice inside the pass that grades others
for it. What I did establish: the cited anchor `internal/web/web.go:124-163` is the **`backend.Attach(...)`
argument list**, while the `.Group(` declarations live in `internal/web/backend/backend.go` — **the anchor
and the counted construct are not in the same file.** Suspicious, not settled.
**This count has now been disputed in three consecutive readings.** Routed as
`FIX-M257x-iter131-root-mount-count-underived` — it needs a derivation with its invocation published.

## ⚠ Method deviation — the adjudication is LESS independent than iter-119's

**The session hit its hard subagent cap (200) after one adjudicator was dispatched.** iter-119 used four.
This reading got one (`adj-1`, seats A, still running at close), and **I adjudicated the other twelve
seats myself** — while being the author of three of the upheld predicates. That is the arrangement `F4`
exists to distrust. I upheld all three, but **the reader should discount this adjudication's independence
accordingly.** The seats themselves were fully blind, dealt before adjudication existed, and committed
verbatim, so the raw record is intact and re-adjudicable. Routed as
`FIX-M257x-iter131-adjudication-independence`.

## Discipline, as sealed

Instrument byte-identical on **both** sides of the copy (`3858ec53…`), `diff` empty, `git log --follow`
one commit, **and the delivered file's first 171 lines still hash to the same value after the addendum**
— proof it sits strictly below the frozen text. Pre-registration sealed in its own commit `a532493`
**before any seat was dealt**. All 14 seats committed verbatim pre-adjudication. **No clone was fetched
during the reading** — fetch times re-read at close, all unchanged. **No repair was taken inside the
measuring pass**, including the three defects of my own that the reading found.

---

## Close — 2026-08-07

**Outcome:** **`P = 30` / `N = 47`.** `P` ROSE 22 → 30 and `N` rose 28 → 47 over a corpus that absorbed
eleven iters of repair. **The test-retest overlap with iter-119 is ~0**, so two consecutive readings
produced almost disjoint predicate sets — the pool is larger than any single `P` has shown, and the
milestone has been reporting a test-retest metric its own prior sheets cannot support. 7 of 15 bands held.
**Clause 5 is NOT met** and is not re-cut.
**Type:** tik
**Status:** closed-fixed (the planned deliverable was a graded reading; it landed with full discipline)
**Gate:** NOT MET — **4 of 5**, unchanged. Clause 5 is met only by a reading returning zero; this returned 30.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no successor strategy is authorable — `TOK-08`'s refutation branch bars one, and run 83 proceeds under the user's direct brief**) — (3) re-scope: n (the milestone is already at the user for scope; this reading is the evidence they asked for, not a new trigger) — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** adjudication done largely by the coordinator (cap-forced, disclosed); band #5 retired after
four consecutive failures; the root-mount count left CANNOT-SETTLE rather than guessed.
**Side-deliverables:** none — a measuring pass takes no repair.
**Routes carried forward:**
- `FIX-M257x-iter131-infrastructure-hedge-stale` → **next iter, highest value**: 11 sites + `CLAUDE.md`
  still publish "never in any clone set / cms unmeasurable" against 28 sites citing the read.
- `FIX-M257x-iter131-my-three` → P7 (§1 missing `library-unimported`), P19 (the over-claimed curl), P5
  (`architecture_overview.md:83` still lists `ai`).
- `FIX-M257x-iter131-predicate-sets-not-enumerated` → every reading must publish its full predicate list.
- `FIX-M257x-iter131-root-mount-count-underived` → disputed in three consecutive readings.
- `FIX-M257x-iter131-adjudication-independence` → re-adjudicate this seat set with independent agents.
- the other 25 upheld predicates, each with anchors, in `adjudication.md`.
**Lessons:**
1. **A fence that prints the right answer does not correct the prose beside it.** Assertion G emits the
   true module set on every run while two corpus sentences contradict it. Fencing a value and repairing
   the sentences that state it are different jobs.
2. **A correction reaches where it was written and nowhere else unless somebody sweeps.** One read of
   `infrastructure` settled four questions; eleven sites still publish the hedge it retired.
3. **A test-retest band needs its predecessor to have enumerated its predicates.** This milestone has
   quoted overlap figures for three readings without the substrate to compute them exactly.
