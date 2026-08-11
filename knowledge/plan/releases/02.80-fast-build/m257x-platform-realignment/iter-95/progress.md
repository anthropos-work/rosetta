# iter-95 — THE READING. N = 13. Clause 5 is NOT met.

**Outcome: the reading was TAKEN**, at platform `0c91421` — the first reading ever taken there.
14 blind seats, two readings of an identical partition, all reports on disk under `raw/`, all 55
booked blockers adjudicated by four independent graders re-deriving from the clones.

## The number

**N = 13** distinct in-scope upheld BLOCKER anchors — **12** distinct predicates.
Reading #17 found **10**, reading #18 found **7**, they matched on **4**.

**The gate does NOT move. It stays 4 of 5.** Clause 5 is met only by a reading that returns **zero**,
and this returned 13. That is a real result, not a failed run.

| quantity | value |
|---|---|
| booked (14 seats) | **55** (27 in #17, 28 in #18) |
| upheld | **51** — **92.7 %**, against iter-80's 92.1 % and iter-84's 93.0 % |
| rejected | **4**, all the ref-discipline class, all caught by adjudicators |
| **graded N (in-scope, BLOCKER, deduped)** | **13 anchors / 12 predicates** |
| Chapman N̂ on the graded set | **≈ 16.6** — a floor |
| per-pass recall | **60 %** (#17) / **42 %** (#18) |
| union recall | **≈ 78 %** |

## Comparability — stated, not implied

**The upheld RATE is continuous and is the quantity that survives:** 92.1 % → 93.0 % → **92.7 %**
across three consecutive adjudications. The instrument's precision has not drifted.

**The COUNT is not directly comparable to the `140 → 43` series, and saying so is the honest answer.**
Three reasons, none of them cosmetic:

1. **Different basis.** iter-76's 140 and iter-82's 43/40 counted *all upheld findings*. This iter
   reports the count on the **gate's own basis** — in-scope, BLOCKER-grade, deduped — because that is
   what clause 5 actually asks. On the older basis this reading is 51 upheld of 55 booked.
2. **Different ground truth.** This is the first reading at platform `0c91421`; the whole prior series
   ran at `0dab54d`. Every compose / `repos.yml` / profile predicate moved underneath the corpus
   between readings. Findings that are *not* platform-derived stay comparable; those that are, are not.
3. **The raw series was already declared discontinuous at iter-86** (`D-M257x-86-2`) for the seat-ref
   sheet, out loud. That declaration stands and is not re-litigated here.

**Verdict: continuous in INSTRUMENT** (briefing byte-identical, 7 seats × 2 readings, same partition
method, same grading rule, same scope definition) **and continuous in UPHELD RATE; a declared
re-baseline of the COUNT**, on both the basis change and the ref advance. Either is acceptable;
silence was not.

## The pre-registration graded — 6 of 6 held, and that is itself a finding

1. per-reading count in [0,12] — **HELD** (10, 7)
2. neither reading returns zero — **HELD** (this is the prediction the gate wanted falsified; it was not)
3. union > max — **HELD** (13 > 10)
4. per-pass recall < 60 % — **HELD**, but *barely* (59.9 % / 42.2 %)
5. a platform-derived class dominates — **HELD** (7 of 13 trace to the `0dab54d → 0c91421` move)
6. upheld rate ≥ 80 % — **HELD** (92.7 %)

**Six of six is not a victory lap.** iter-76 graded 2 of 5 and iter-53 graded 2 of 5, and those
mis-predictions were where the learning was. A pre-registration that never fails is a pre-registration
that stopped being risky — prediction 1's band `[0,12]` was wide enough to be nearly unfalsifiable, and
prediction 4 landed inside a tenth of a point. **Tighten the bands next reading.**

## What the zero would have established, and what this 13 establishes instead

The user asked what a zero would prove. It would prove **only** that a 14-seat pass at ~78 % union
recall found nothing — with `N̂` a floor, and both readings sharing a briefing, file set, partition and
model, so every recall figure is optimistic. **A zero would have been evidence, not proof.**

What 13 establishes is stronger and less comfortable: **the corpus still carries at least 13 blocking
falsehoods inside the gate's own scope**, of which **~4 more are estimated to remain unfound**, and
**at least 8 additional propagation sites of already-upheld predicates were named by adjudicators that
no seat booked.** Clause 5 is not close.

## The two findings that outrank the defect list

**1. Recursive `grep` in this environment silently hides tracked files** — the wrapper is
`ugrep --ignore-files` and `studio/.gitignore` lists `tools/`. `grep -rn mistralai app/studio/` → **1**;
`git grep mistralai <ref>` → **2**. Verified independently by the orchestrator. This produced a **false
clearance** by one seat on a predicate another seat correctly booked, and is very likely how the false
claim entered the corpus originally. **It biases every absence-claim toward under-counting, which makes
N a floor for a second independent reason.** New standing rule: *an absence is established only by
`git grep` at a named ref.*

**2. A wrong positive control is a broken instrument.** `messenger.md:22` offers *"`-S SKILLER_RPC`
returns 3"*; it measures **7** at every ref, clone and spelling. Graded MINOR because the claim it
guards is true — but a control exists precisely to let the next reader trust their pipeline, and this
one manufactures the doubt it was written to remove.

## Held by instruction — reported, not folded in

- **`DEF-M257x-iter80-storage-prod-bucket`** holds `storage.md:55`, `:154`, `:181`. Adjudicator 4
  re-derived all four candidate anchors and found: **`:55` is CORRECT** (the manager table row),
  `:156`/`:181` are explicitly **past-tense / HISTORICAL-fenced**, and **`:58` — which is NOT among the
  three held lines — is the only present-tense sentence, and it is the false one.**
  **`storage.md:58` is therefore counted IN N.** If the user intends the hold to cover the whole
  production-bucket hazard class rather than the three named lines, **N drops 13 → 12** (predicates
  12 → 11). Stated both ways so the number can be read either way; not decided here.
- **The five post-freeze items** stay disclosed in the pass-22 ledger and routed. Not folded in.
- **The `cms` ECR deletion** (`6efa1d5`, 2026-08-04, *"is decommissioned"*) was **found unaided by two
  seats in both readings** and upheld — as items 5 and G-B1. The new hedge fence covers it exactly as
  hoped: `dependency_map.md:31` and the map's own `:88` rule *"report both and assert neither"*, and
  `unreadable_repo_claim_guard` (GREEN this open) labels the `infrastructure` class as not-a-measurement.
  **The defect is that two files assert it anyway.** Reported, not asserted — and now booked.

## Routed, NOT repaired (binding: a measuring pass may not contain a repair)

**`FIX-M257x-iter95-read-union` — the 13 anchors / 12 predicates above.** Binding conditions inherited
from iter-76's routing, which is why 140 → 43 meant anything:
1. **Repair by PREDICATE, not by anchor** — adjudicators named ≥8 unbooked twins; anchor-wise repair
   leaves them standing.
2. The M903 clause must be repaired in **both** `platform-migration-status.md:92` and `storage.md:25`,
   and note that one seat **positively cleared** it at the former.
3. Re-read after repair. A repair pass can only repair what a reading NAMES, and union recall is 78 %.
