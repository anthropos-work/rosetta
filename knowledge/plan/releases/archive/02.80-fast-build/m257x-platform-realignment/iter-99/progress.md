# iter-99 — THE RE-READ. N = 28. Clause 5 is NOT met. Gate stays 4 of 5.

**Type:** tik, under `TOK-05`. The measuring pass for `FIX-M257x-iter97-read-union`'s repair.

**Outcome: the reading was taken**, at platform `0c91421`, corpus `e858fd4` — the tree iter-98's
predicate-wise repair produced. 14 blind seats over an identical partition (40 files / 10,276 lines), all
reports on disk under `raw/`, all 46 booked blockers adjudicated by four independent graders re-deriving
from the clones, all verdicts on disk under `verdicts/`.

**No repair is in this pass.** Everything found routes to `FIX-M257x-iter99-read-union`.

## The number

**N = 28** distinct in-scope upheld BLOCKER anchors. Reading #21 found **18**, #22 found **16**, matched on
**6**. Chapman **N̂ ≈ 45.1**; per-pass recall **39.9 % / 35.4 %**, union **≈ 62 %**; **≈ 17 estimated still
unfound.**

**The gate does NOT move. It stays 4 of 5.**

## Pre-registration: 4 of 9 held — and the failures are the content

Sealed in a separate commit (`964b7a3`) **before any seat reported**, which is what makes it a
pre-registration rather than a retrospective. Full grading in [`adjudication.md`](./adjudication.md).

**Held:** per-reading counts (18, 16 ∈ [8,18]) · zero true recurrences of iter-98's 21 predicates ·
recall in [28 %,55 %] **with both passes below 41 %** as required · platform-drift ~1 of 28 · induced
blockers exactly **2** ∈ [0,3] · mean sites/predicate ≈1.2 < 2.5.

**Failed:** union `N` = 28 (band [12,26]) · upheld rate **78.3 %** (band [88 %,96 %]) · wrong-construct
citations **≥7** (band ≤1).

## The three things that matter more than the defect list

**1. Band #9 failed by ~7×, and it indicts the INSTRUMENT.** `anchor_construct_guard` was **GREEN at the
audited commit** — *"every resolvable anchor names a construct"* — while ≥7 upheld findings are citations
resolving to the wrong construct, including a self-citation offered AS evidence that lands on a **blank
line** (`ai-readiness.md:46`) and `manager.go:485`, which is a closing brace. The load-bearing word in the
guard's green is **"resolvable."** The band was set at ≤1 precisely so an upheld member would mean a blind
spot; it did.

**2. Precision fell 93.1 % → 78.3 %, the first break in five readings.** Rejections rose 4 → 10 while
bookings *fell* 58 → 46. Three mechanisms are consistent with it and **this reading cannot separate them**:
the residual genuinely got harder (iter-98 measured max predicate width collapsing 11 → 4); a briefing gap
(**two independent seats made the identical wrong-tree error** on the rext anchors, grading the authoring
copy instead of the pinned per-stack clone); and adjudicator variance (`hiring.md:80-82` was REJECTED by one
panel and UPHELD by another — one disagreement in 46, but the first non-zero). Recorded, not resolved.

**3. Exactly 2 of the 28 were induced by iter-98's own repair, both inside prose it rewrote.**
`dependency_map.md:59` — the repair pinned the cell to one ref and wrote *"6 Go occurrences across **4**
files"*; both readings measured **3**. `backend.md:33-34` — the repair removed `skiller` from the both-ends
set and left the set asserting exhaustiveness while **omitting `backend` itself**. Prediction #7 forecast
[0,3] and got 2: **the mechanism model holds while the magnitude model fails**, the same split iter-97 found.

## What this establishes, and what it does not

**Establishes:** ≥28 blocking falsehoods in clause 5's scope at `e858fd4`, ≈17 more estimated unfound; the
platform-drift class is ~1 of 28; `anchor_construct_guard` has a demonstrated blind spot on intra-corpus
citations; iter-98 induced exactly 2.

**Does not establish that `N` is rising.** 13 → 20 → 28 was measured at upheld rates of 92.7 / 93.1 /
**78.3 %** — the instrument's precision is no longer constant, so the series is not comparable on the axis
iter-97 relied on. [`iter-98/discovery-pool.md`](../iter-98/discovery-pool.md) §3 predicted recall would
fall as the pool narrowed and it did (union 68 % → 62 %; band #5 held). **A narrowing pool measured by a
degrading instrument yields a rising `N`** — consistent with this data, and so is a genuinely growing
residual. This reading cannot tell them apart and does not claim to.

## Routed, NOT repaired

**`FIX-M257x-iter99-read-union`** — the 28 anchors, plus:

- **`CHECK-M257x-iter99-anchor-guard-blindspot`** (highest value): `anchor_construct_guard` reports GREEN
  while ≥7 intra-corpus citations resolve to the wrong construct. Find what "resolvable" excludes — blank-line
  targets, corpus-internal `:N` self-anchors, and unpinned anchors are all represented in the upheld set.
- **`CHECK-M257x-iter99-briefing-rext-ref`**: the briefing does not say **which rext clone** grades an rext
  claim. Two seats independently graded the authoring copy where the pinned per-stack clone `ab81527a` was
  correct, and both bookings were rejected. That is an instrument defect with a one-line fix, and it is the
  only rejection class in this reading that repeated across readings.
- **`CHECK-M257x-iter99-precision-drop`**: decide whether 78.3 % is the residual hardening, the briefing gap,
  or adjudicator variance. It changes how every future `N` is read.

Binding conditions inherited and extended:
1. **Repair by PREDICATE, enumerate PARAPHRASES** (iter-98's addition; 0 true recurrences says it worked).
2. **Re-derive every inbound citation after any edit that changes a file's line count**, in-file self-anchors
   included — and re-run after the citation fixes themselves.
3. **Do not write a measurement into the corpus without running the measuring command as printed.**
4. **New:** a repair that rewrites an enumeration must re-derive **the whole set**, not just the member it
   came to fix. Both induced defects are enumerations that were edited without re-counting.

## Close — 2026-08-06

**Outcome:** N = 28 at `e858fd4`; 46 booked, 36 upheld (78.3 %), 4 of 9 pre-registered bands held; the
instrument's own blind spot surfaced and routed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5, unchanged.** Clause 5 is met only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** recorded in [`adjudication.md`](./adjudication.md); no new milestone-level TOK.
**Side-deliverables:** none — a measuring pass may contain no repair, and contains none.
**Routes carried forward:** `FIX-M257x-iter99-read-union` (28 anchors) ·
`CHECK-M257x-iter99-anchor-guard-blindspot` · `CHECK-M257x-iter99-briefing-rext-ref` ·
`CHECK-M257x-iter99-precision-drop`.
**Lessons:**
- **A fence's green is scoped by a word, and the word is where the blind spot lives.** *"Every **resolvable**
  anchor names a construct"* was true and simultaneously compatible with ≥7 wrong-construct citations.
- **When precision breaks a four-reading band, the count taken with it is not comparable** — say so before
  reporting the count, not after.
- **Two seats making the same error is an instrument finding, not seat noise.**
