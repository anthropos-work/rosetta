# iter-109 adjudication — the graded sheet

## The one-line answer

> ### `P = 24` predicates · `N = 36` anchors.
> ### The pre-registered rule's `P ≥ 15` branch fired: **THE POOL DOES NOT DRAIN.**
> ### And the secondary rule's `N ≥ 20` branch fired with it. **Both metrics agree.**

The rule was sealed in `ac48e5b` **before the first seat was dealt**, and is graded here exactly as
written. Nothing was re-cut, re-centred or re-read once the number was known.

| | iter-101 | iter-103 | **iter-109** |
|---|---|---|---|
| distinct false **predicates** `P` | 22 | 22 | **24** |
| distinct **anchors** `N` | 24 | 33 | **36** |
| anchors per predicate | 1.09 | 1.50 | **1.50** |

**Three consecutive readings at 22 · 22 · 24 predicates, across two full repair cycles that between them
paid 74 anchors and were machine-graded at ~100 % of the upheld union both times.**

## What ran

| | |
|---|---|
| corpus | `ac48e5b` · scope `corpus/services/**` + `corpus/architecture/**` (40 files, 10,694 lines) |
| readings | **#27: 7 of 7** · **#28: 7 of 7** — 14 blind seats, identical recomputed partition |
| bookings | **35 booked → 32 UPHELD / 3 REJECTED**; 31 in-scope upheld blockers |
| adjudicators | 4, grouped by seat LETTER so both readings of one file set land with one grader |
| seats lost | **zero** — no re-deals needed. Every seat committed verbatim on landing |
| clone refs | frozen, and **provably frozen**: all 14 identical at open and close. No fetch (§5 rule 41a) |

---

## THE QUESTION THIS READING WAS BUILT TO ANSWER, AND ITS ANSWER

The pre-registration named the confound before the number existed. iter-103's 61 %-drift decomposition was
consistent with two stories, and iter-103 could not separate them:

- **ARRIVAL** — those 20 drift anchors became false *because* five clones advanced. Fence the advance and
  they do not come back.
- **DETECTION** — they were false all along, and a reading finds mechanical defects once the subtle half
  has just been repaired.

**With all 14 clones frozen at the identical sha, the answer is DETECTION.**

> **Platform-drift is still ~33 % of the upheld residual — over a subject in which literally nothing
> moved.** Band #8 predicted `≤ 25 %` on the reasoning that a frozen subject cannot arrive new drift. It
> failed, and the failure is the finding: **the drift was already in the corpus.** It was never an inflow.
> There was nothing for the drift fence to catch, because nothing drifted.

This retires the central premise of `TOK-06`. The strategy was authored on *"inflow is comparable to
outflow"* — a real measurement, honestly taken, that turns out to have measured the **composition of what a
reading detects**, not the **rate at which defects arrive**.

---

## Bands — 7 HELD of 13

| # | prediction | band | actual | verdict |
|---|---|---|---|---|
| 1 | per-reading in-scope upheld blockers (n₁, n₂) | [3, 16] each | **17 / 14** | **FAILED** (n₁ high by 1) |
| 2 | union `N` | [4, 19] | **36** | **FAILED** high |
| **2p** | **union `P` — the primary** | [3, 14] | **24** | **FAILED** high |
| 3 | overlap with iter-103's 22, matched on predicate | [0, 3] | **2** | **HELD** |
| 3b | within-reading `m` as share of union | [10 %, 55 %] | **6/24 = 25 %** | **HELD** |
| 4 | adjudicator upheld rate (raw) | [72 %, 94 %] | **91.4 %** | **HELD** |
| 5 | per-pass recall spread | ≥ 12 pts | **8.3 pts** | **FAILED** |
| 6 | wrong-tree rejections | [0, 3] | **0** | **HELD** |
| 7 | wrong-construct intra-corpus citations | ≤ 4 | **1** | **HELD** |
| 8 | platform-drift share of upheld blockers | ≤ 25 % | **~33 %** | **FAILED** |
| 9 | per-seat booked spread over 14 seats | ≤ 8 | **4** | **HELD** |
| 10 | repair-induced (anchors in prose iters 104–108 wrote) | [0, 4] | **2** | **HELD** |
| 11 | anchors per predicate | [1.00, 1.35] | **1.50** | **FAILED** |

**Upheld rate, reported twice as bound**: **91.4 % raw**, and **91.4 % with the `wrong-tree` class
separated** — because `wrong-tree` was **zero** this reading and the two numbers coincide for the first
time in the series. Band #6's series is now **4 → 1 → 1 → 0**: an addendum can carry ground truth a frozen
instrument gets wrong, and adding the second `D-M257x-103-7` refinement did not muddy it.

---

## THE STRUCTURAL FINDING — a repair scoped to a reading's DETECTIONS cannot close a predicate

Band #3 held at **2**, and iter-103's version of it held at 1. But **the meaning inverts**, and this is the
most useful sentence on the sheet.

iter-108 was graded **46/46 = 100 % of the upheld union** by `repair_reach_guard`, and that grade is
correct. It repaired **by predicate** — TOK-05's unit, the right unit. But its anchor list was **derived
from `iter-103/raw/`** — from *what the previous reading detected*.

> **A predicate's site list and a reading's detection list are not the same set.** iter-108 closed every
> site iter-103 *saw*. Where the same predicate was also published at a site iter-103 never booked, the
> repair had no reason to visit it — so the falsehood survived, at full strength, one file away.

Measured, in this reading:

- **The `ai`-fold predicate.** iter-108 repaired `external_services.md:565`. The same proposition — *"All
  Go services access AI through the shared `ai` library"* — sits **eleven lines above at `:554`**,
  unrepaired, and was booked by seat A. Same file. Same predicate. Outside the derived ledger.
- **The Mistral call-site anchor.** iter-108 repaired `ai_architecture.md:95` and `:99`. The twin at
  `ai_architecture.md:34` still cites `markdownManager.go:19` — while the repaired
  `external_services.md:560` now states explicitly that *"the `:19` this row used to cite is a
  **doc-comment** line, not code"*. **The repair created a self-contradiction by fixing one side of a
  pair.**

**So the induction fences worked and the induction is not the problem.** Band #10 measured only **2 of 36
anchors** inside prose iters 104–108 wrote — the lowest in the series, against a rate that had held ~2 per
cycle for six cycles with a far smaller repair. Step 2 did its job. What the fences cannot see is a
predicate the repair **never visited**.

---

## Composition of `N` — and why it is not the composition `TOK-06` was authored against

| class | share of upheld | fenced by |
|---|---|---|
| **standing platform-drift** — version literals, line offsets, symbol names, paths | **~33 %** | nothing can fence it: **it did not arrive**, it was already there |
| **repair-scope twins** — the predicate closed at one site, alive at another | ~2 predicates confirmed | nothing |
| **repair-induced** — anchor inside prose iters 104–108 wrote | **2 of 36 = 5.6 %** | `anchor_offset_guard` + `repair_postcondition` — **and they held** |
| never-true / mixed | remainder | the reading |

Compare `TOK-06`'s table: clone advance **61 %**, induction **21 %**. **Induction fell 21 % → 5.6 %, which
is the fences working. Drift did not fall at all, because it was never flowing.**

---

## THE VERDICT, stated first and loudest as the pre-registration requires

> ### `P = 24` → **THE POOL DOES NOT DRAIN.**
> The `≥ 15` branch. Third consecutive reading at ~22–24 predicates, with the clone-advance inflow
> **provably absent** — no clone moved at all.

**What this establishes:**

1. **The 61 % was DETECTION, not ARRIVAL.** A frozen subject still yields ~33 % drift findings.
2. **Repair-and-read does not converge on clause 5 as currently scoped**, and the reason is now specific
   rather than atmospheric: the repair's *unit* is right (predicate) and its *reach against its own ledger*
   is ~100 %, but its *scope* is a prior reading's detections, and detection recall on this instrument has
   run **33–83 %**. A loop that repairs only what was seen cannot drain a pool it only ever samples.
3. **The inflow fences are not wasted and should not be reverted.** Induction fell to 5.6 %. That leg
   worked. It simply was not the binding constraint.

**What it does NOT establish:**

- **It does not establish the pool size.** `P` and `N` are **floors** in every branch. No point estimate is
  offered — Chapman stays retired; only floors survive: **≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`, ≥ 36 at
  `ac48e5b`**.
- **It does not license re-cutting clause 5.** Clause 5 is met only by a reading that returns **zero**.
  Four user rulings, and this sheet does not argue any of them. `P = 24` leaves it open.
- **It does not say the corpus got worse.** 22 → 22 → 24 predicates across a growing, twice-repaired
  corpus is, within this instrument's demonstrated variance, **flat**. It is the flatness that is the
  result.

---

## Rejections — 3 of 35, and one of them is the same claim rejected twice, two readings apart

| booking | anchor | class | note |
|---|---|---|---|
| `r27-G B1` | `shared_libraries.md:128` | **ref-discipline** | **iter-103 rejected this identical anchor as `wrong-tree`.** Two adjudicators, two readings, one conclusion: the section names its subject *and* its `v1.40.2` pin, and all three booked statements hold byte-for-byte at that tag, readable through `app`'s fold commit `1e457fa70`. **Third independent confirmation that the claim is TRUE** — and the reason iter-108 was right not to repair it. |
| `r28-E B1` | `backend.md:360` | other | an omission, not a falsehood: `app@ad9f3c49` has a net-new **second** Atlas env (`sentinel`, `terraform/migrations-sentinel`) the passage does not mention |
| `D-r28 B2` | `security_compliance.md:252` | mis-read | the fence covers the classification bullet twice over |

**`wrong-tree` = 0.** The briefing defect that has run 4 → 1 → 1 is, at this reading, **not observed at
all**, while still being delivered unfixed for the fourth time.

---

## Routed

- **`FIX-M257x-iter109-read-union`** — the 24 predicates / 36 anchors. **With one binding change of scope:
  the anchor set must be re-derived FROM THE CORPUS per predicate, not from `iter-109/raw/`.** Repairing a
  reading's detections is what produced the twins measured above. `repair_reach_guard` should grade against
  the *corpus-derived site set*, and its denominator should say so.
- **`FIX-M257x-iter109-repair-scope-is-detection-bounded`** — net-new, and it outranks the union. A repair
  that closes a predicate at one site while leaving it at another **manufactures a self-contradiction**,
  which is strictly worse than the original defect. Needs a mechanism — a per-predicate corpus-wide site
  sweep before a repair is called done.
- **`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`** — **stays open**, and this reading changes its
  priority: with drift shown to be standing rather than arriving, a drift *fence* is not the lever it was
  ranked as. Re-rank rather than re-attempt.
- **`DEF-M257x-iter101-briefing-rext-tree`** — stays open, delivered-unfixed, **fourth** measurement: 0.
- **`D-M257x-109-1`** — a seat-commit subject named one seat and carried two; recorded, not rewritten.

## Provenance and housekeeping

- **§5 rule 41a held and was PROVEN, not asserted**: every clone HEAD, `origin/main` and fetch timestamp
  re-read at the close is identical to the open. **No clone fetched.**
- **Guard family at the close: 15 GREEN · 0 RED · 4 not-run**, run from the **authoring** tree
  `680e8529f`, whose path and sha the family now prints as its first line (`FIX-M257x-iter103-guard-tree-provenance`,
  TOK-06 step 0, in production).
- **The corpus HEAD moved mid-open** (`2e3443d` → `08cfbd8`, a concurrent lane) and the **read scope
  provably did not** — identical `corpus/services` and `corpus/architecture` tree hashes at all three refs,
  with `e6aed2e` as a firing negative control. Corrected in the open in `ground-truth.md`, `c558c08`.
- **No stack** was brought up, torn down or reconfigured. `stack-demo/**` untouched. **Zero platform-repo
  edits.** No tag cut; rext stays on `main` at `680e8529`.
- **No `stack-core` full-suite total is quoted anywhere in this iter** — the suite does not complete on this
  host (`FIX-M257x-iter108-stackcore-suite-hangs`), and a gap is recorded rather than a pass.
