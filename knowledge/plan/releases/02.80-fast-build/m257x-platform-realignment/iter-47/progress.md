**Type:** tik — TOK-02 **step 5 of 5**, the last (*"ONE full 7-auditor read, instrument held fixed at
iter-41's… That reading is what meets or fails clause 5. Nothing else does."*).

# iter-47 — the seventh pass

## Phase A — read

Seven auditors, **40 files / 9,243 lines**, instrument frozen at iter-41's on every knob. Per-file `wc -l`
positive control: **all 40 confirmed line-for-line by their own auditor.**

| seat | files | lines | blockers | minors |
|---|---|---|---|---|
| A | 7 | 1726 | 0 | 10 |
| B | 7 | 1559 | **2** | 11 |
| C | 7 | 1520 | **1** | 9 |
| D | 7 | 1481 | 0 | 13 |
| E | 6 | 1459 | 0 | 7 |
| F | 6 | 1498 | 0 | 9 |
| G | (diff) | 30 hunks / 17 files | **5** | 5 |

8 raw findings; `G1 ≡ B1` reached independently from two different seats → **7 unique**.

## Phase B — adjudicate

Every blocker re-derived by this iteration against platform source before acceptance (`adjudication.md`).
**7 of 7 HELD — none refuted on re-derivation.**

## Phase C — the measurement

**7 blockers. And every one of them is in text iter-46 wrote or rewrote.**

| pass | iter | corpus read | blockers | pre-existing | induced |
|---|---|---|---|---|---|
| 5 | 41 | post-39 repair | 18 | 9 | 9 |
| **6** | **47** | post-46 repair | **7** | **0** | **7** |

Six full-read auditors covering all 40 files found **zero** blockers in text iter-46 did not touch. For
six passes the corpus term and the repair term were confounded; this pass separates them, and **one of
them is zero** (`D-M257x-47-3`).

### Pre-registered predictions, graded

| prediction | outcome |
|---|---|
| **fewer than 9** (TOK-02's own, carried unmodified) | **CONFIRMED — 7.** The first confirmed prediction in the series |
| fewer than 4 self-contradiction | **CONFIRMED — 3** |
| ≥1 over-correction in explanatory text | **CONFIRMED — 4** |
| residual NOT concentrated in the 17 repaired files | **REFUTED, absolutely — 7 of 7 are in them** |
| clause 5 does not close | **CONFIRMED** |

### Why the fences are silent, and the one that could be built

All four fences correctly report 0 sites. `claim_twin_guard` matches **adjudicated** refuted forms; six of
seven blockers are **new prose with no ledger entry**. `anchor_construct_guard` misses #4 because
`external_services.md:489` resolves *and* carries content — the wrong construct, the class
`D-M257x-45-3` declined to tune a fence for.

**But three of the seven need no new idea.** #5, #6, #7 are each a grep for a string still sitting in the
tree beside its own repaired twin. **A leak-check over the repair's own diff would have caught all
three** — and is exactly what auditor G did by hand (`D-M257x-47-5`).

## Close — 2026-08-02

**Outcome:** the seventh pass returns **7 blockers**, down from 18 on an identical instrument — and the
decomposition is the result: **0 pre-existing, 7 induced.** TOK-02's pre-registered *"fewer than 9"* is
confirmed, the first confirmed prediction in the series. Clause 5 **fails**; the residual is now entirely
self-inflicted, and 3 of the 7 are mechanically greppable.
**Type:** tik
**Status:** closed-fixed — the planned deliverable was a measurement, and it landed with its predictions
graded and its findings adjudicated. Nothing was repaired, by design (`D-M257x-47-2`).
**Gate:** NOT MET — 4 of 5. Clause 5 at **7**.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; the metric moved 18 → 7) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2; and 7 is far below the 18 that would have refuted TOK-02) — (4) user-blocker: **y** — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — **Outcome: exit-4**
**Decisions:** `D-M257x-47-1` … `D-M257x-47-8`
**Side-deliverables:** none — this iteration wrote no code and repaired no prose.
**Routes carried forward:**

- **`FIX-M257x-iter47-blocker-set`** — the 7, by claim, tree-wide. **3 are pure leak-greps** (#5 `external_services.md:139`, #6 `ai_architecture.md:84`, #7 `coverage-protocol.md:614-616`); **`ai_architecture.md:42-56` should be rewritten against `external_services.md:532-533` rather than de novo** — that block is the third consecutive iteration in which the paragraph explaining a correction became the next iteration's finding.
- **`FENCE-M257x-iter47-leak-check`** — the buildable check this reading names: *for every claim a commit changes, grep the whole tree for the old form.* Would have caught 3 of 7. The highest-value mechanical finding here.
- `DOC-M257x-iter47-minors` (~64) — including two worth promoting: `service_taxonomy.md:150-153` (iter-46's new table cell spans four physical lines and **will not render as one row**) and `hiring.md:189-196` (the "minimal write-set" omits a NOT NULL + UNIQUE column, so a seeder built from it fails its INSERT).
- `CHECK-M257x-iter35-seeder-writes-one-instant` · `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone).

**Lessons:**

- **Separating the two terms was worth more than reducing either.** For six passes "the residual" was one
  number and every interpretation of it was contested. Split, it says something falsifiable and new: the
  corpus is clean, and the repair is the defect source.
- **A repair re-derives quantities and narrates mechanisms.** Every number iter-46 checked was exact;
  every failure was a sentence about how something works — a pointer that is a value, a service never on
  the backend it is "flipped off", an enumeration missing its third member. *Re-derive from source* has to
  bind mechanisms as explicitly as it binds numbers.
- **The cheapest unbuilt check is the leak-check, and it is not the hard one.** Three of seven are greps
  for a string the repair itself left behind. The fence family answers *"has a refuted claim come back?"*;
  nobody built *"did this commit finish?"*.
- **An author correcting a distinction can invert it.** Blocker #1 was written to fix "unrecognised" and
  replaced it with a nullable-pointer story that is wrong in both conjuncts — while the file it cited as
  *"the full per-line derivation"* already had it right. **Fifth consecutive iteration in which the author
  of a correction violated it while writing it.**
