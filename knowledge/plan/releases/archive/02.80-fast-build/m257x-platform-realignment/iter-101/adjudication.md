# iter-101 adjudication — readings #23 + #24 at platform `0c91421`, corpus `8f04d3a`

**Status: COMPLETE ON 13 OF 14 SEATS.** All **36** booked blockers across the **13 seats that exist** were
graded by **four parallel adjudicators**, each re-deriving from the platform clones rather than from any
seat's evidence or any prior verdict.

## The short seat — disclosed, not smoothed

**`r24-D` was never produced.** The 14-seat fan-out died on a spend limit with seat D of reading #24
unwritten; the other 13 reports landed on disk and were committed verbatim, pre-adjudication, at `8b6d80f`.

**It was NOT re-run.** Re-dealing it costs a fresh fan-out and perturbs the replicate this iter exists to be.
Adjudicating 13/14 with the gap stated is cheaper and more honest.

**The binding consequence, stated once and honoured everywhere below:**

> **Reading #23 is a SEVEN-seat count. Reading #24 is a SIX-seat count. `n₂` is never compared to `n₁`, to a
> prior reading's `n`, or to a band, without that being said.**

Seat D contributed **3 anchors / 1 predicate** to reading #23 and **nothing** to reading #24. Every
cross-reading quantity below is therefore also given on the **6-seat common subject** (seats A·B·C·E·F·G),
where both readings actually sampled the same pool.

## Verdict

| | Adj1 · A+C | Adj2 · B+D | Adj3 · E | Adj4 · F+G | **total** |
|---|---|---|---|---|---|
| booked | 10 | 9 | 7 | 10 | **36** |
| **UPHELD** | 5 | 8 | 5 | 10 | **28** |
| REJECTED | 5 | 1 | 2 | 0 | **8** |
| in-scope upheld bookings | 5 | 8 | 5 | 9 | **27** |
| wrong-tree rejections | 0 | 1 | 0 | 0 | **1** |

### The upheld rate, reported TWICE — as the pre-registration binds

| basis | value |
|---|---|
| **raw** | 28 / 36 = **77.8 %** |
| **with the `wrong-tree` briefing-defect class separated** | 28 / 35 = **80.0 %** |

Against 92.1 (iter-80) · 93.0 (iter-84) · 92.7 (iter-95) · 93.1 (iter-97) · **78.3 (iter-99)**.

**This is the finding iter-99 could not get at n=1: the precision drop is STRUCTURAL, not adjudicator
variance.** 78.3 % → 77.8 % on a materially unchanged pool, a fresh instrument-identical reading, and four
different adjudicators. Two readings now sit ~15 points below the four that preceded them. The break is a
property of the residual and/or the briefing — **not** a one-off.

### The ref-discipline class fired ZERO times

17 occurrences across five readings, **0 here** — and three adjudicators independently reported it as
*structurally* absent: not one seat booked a pinned or dated claim because newer evidence contradicted it.
The class is not merely filtered this reading; the seats stopped generating it.

The 8 rejections are **5 mis-read** (`backend.md:19`, `skiller.md:19`, `ai_architecture.md:35`,
`ai_architecture.md:141`, `security_compliance.md:185`), **2 wrong-convention** (`chronos.md:27` — `8080`/
`8081` are the colony binary's defaults, verified in four sibling repos' `cmd/root.go`, not a compose value;
`ai-labs.md:75`), and **1 wrong-tree** (below).

## **N = 24** distinct in-scope upheld BLOCKER anchors / **22** predicates

`n₁ = 20` (**7 seats**) · `n₂ = 8` (**6 seats**) · `m = 4` → union **24**.

The four matched predicates — found independently by both readings:
`ai-readiness.md:305` · `service_taxonomy.md:130-133` ≡ `:131-133` · `security_compliance.md:67-68` ·
`dependency_map.md:59` (the `SKILLER_STREAM` file count).

| quantity | iter-95 | iter-97 | iter-99 | **iter-101** |
|---|---|---|---|---|
| booked | 55 (14 seats) | 58 (14) | 46 (14) | **36 (13)** |
| upheld | 51 — 92.7 % | 54 — 93.1 % | 36 — 78.3 % | **28 — 77.8 % raw / 80.0 % adj.** |
| rejected | 4 | 4 | 10 | **8** |
| **graded N** | 13 | 20 | 28 | **24** |
| Chapman N̂ | ≈ 16.6 | ≈ 29.3 | ≈ 45.1 | **≈ 36.8 (see caveat)** |
| per-pass recall | 60 / 42 % | 41 / 44 % | 39.9 / 35.4 % | **83.3 % / 33.3 %** |

**The Chapman figure carries a caveat the prior rows do not:** it mixes a 7-seat `n₁` with a 6-seat `n₂`.
On the **6-seat common subject** — the only basis on which the two passes sampled the same pool —
`n₁' = 17 · n₂' = 8 · m = 4`, union **21**, **N̂ ≈ 31.4** over 6/7 of the corpus, which scales to
**≈ 36.6** at 7 seats. The two routes agree; **≈ 36.8** is reported as the headline with its assumption named.

**`N = 24` is a FLOOR for a further reason: it is a 13-seat union.** A 14th seat can only add.

**Clause 5 is NOT met. The gate does not move on this clause.**

## Band #3 — the estimator band, and it FAILED LOW

**Overlap with iter-99's published 28, matched on PREDICATE: `6`.** Band was **[14, 22]**.

The six re-found predicates: `ai-readiness.md:305` · `backend.md:33-34` · `backend.md:29-30` ·
`dependency_map.md:59` · `next-web-app.md:32` · `hiring.md:38`.

iter-101's union re-found **6 of 28 = 21.4 %** of what iter-99 published — measured blind (those anchors live
under `knowledge/plan/**`, which every seat is hard-barred from reading). At iter-99's measured union recall
(≈ 62 %) the expected overlap was ≈ 17. **We observed 6.**

### The verdict the pre-registration pre-committed to

> *"Overlap < 14 says the readings are closer to independent than assumed, the pool is much larger than 28,
> and `N̂` is if anything conservative."*

### **`N̂ = 45.1` is a FLOOR, not a ceiling.**

The Chapman independence assumption is not merely intact — the two readings are **more** independent than it
assumes. Taking iter-99 and iter-101 as two passes over a materially fixed pool (the corpus moved 10,276 →
10,278 lines; ≤ 4 of the 28 were touched by iter-100's anchor repairs):

```
n_99 = 28 · n_101 = 24 · m_cross = 6
N̂_cross = (29 × 25) / 7 − 1 ≈ 102.6
```

**The residual inside clause 5's scope is on the order of ~100, not ~45** — and heterogeneous catchability
(some defects are far easier to see than others) biases Chapman **downward**, so ~103 is itself a floor.

**This is the most consequential number this milestone has produced, and it is stated with its assumptions
rather than banked.** It rests on: a closed population (well supported — 2 net lines), independence (now
*better* supported than assumed), and equal catchability (**dubious, and the direction of that bias is
known**). Adjudicator turnover across the two readings is a second uncontrolled term.

**What it means operationally:** repairing a reading's union has never been the drain the milestone modelled
it as. Three readings have now each named ~13–28 items while the estimated pool grew 16.7 → 29.4 → 45.2 →
**~103**. A zero reading is not near, and no schedule should assume it is.

## The pre-registration graded — **5 of 9 held, 4 FAILED**

| # | prediction | band | outcome |
|---|---|---|---|
| 1 | per-reading count (n₁, n₂) | [10, 22] each | **SPLIT — n₁ = 20 HELD; n₂ = 8 FAILED low** |
| 2 | union `N` | [18, 34] | **HELD — 24** |
| 3 | overlap with iter-99's 28 | [14, 22] | **FAILED LOW — 6** |
| 4 | adjudicator upheld rate | [74 %, 86 %] | **HELD — 77.8 % raw, 80.0 % adjusted** |
| 5 | per-pass recall vs own union | [30 %, 55 %] | **FAILED — 83.3 % and 33.3 %** |
| 6 | wrong-tree rejections | [1, 5] | **HELD — 1** (bottom edge) |
| 7 | wrong-construct intra-corpus citations | ≤ 4 | **HELD — exactly 4** |
| 8 | platform-drift share | ≤ 10 % | **HELD — 1–2 of 24 ≈ 4–8 %** |
| 9 | per-seat booked spread | ≤ 8 | **HELD — 4** (max 5, min 1) |

**#1 does not get rescued by the short seat.** `n₂ = 8` over 6 seats normalizes to `8 × 7/6 ≈ 9.3`, still
below the band's floor of 10. Reading #24 was genuinely less productive per seat than #23 (1.83 vs 3.57
bookings/seat) — that is real reading-to-reading variance, not the missing seat.

**#5 failed in the direction nobody predicted.** Pass #23 recalled **83.3 %** of the union — far above the
band — because pass #24 found so little that the union is nearly pass #23 itself. The two passes are wildly
asymmetric, which is the same fact band #3 measures from the other side.

**#7 held at exactly its boundary.** The four: `service_taxonomy.md:130-133` (anchors point at Chronos /
Intelligence rows), `hiring.md:38` (`service_taxonomy.md:52` is the studio-room correction),
`external_services.md:136` (routes the reader to `:206` *as corroboration*; `:206` refutes it), and
`backend.md:33-34` (cites `:264` as its own authority; `:264` refutes it). iter-100's `anchor_construct_guard`
repair predicted ≤ 4 and got exactly 4 — the mechanical half really was about half the class.

**#6 held at its bottom edge — the briefing defect cost 1, not 4.** `DEF-M257x-iter101-briefing-rext-tree`
is real but far cheaper than iter-99's 4-of-10. Only one booking (seat D's messenger-row prod-terraform
clause) was graded against the wrong checkout. **The instrument was delivered unfixed on purpose and the
class is now measured at n=2 readings: 4, then 1.**

## The three findings that outrank the defect list

### 1. Repair induces defects, and iter-100 induced at least one of its own — inside what it rewrote

`service_taxonomy.md:130-133` was **exactly correct at `a229f8d^`**. iter-100's own two-line parenthetical —
added to qualify *which file* the anchors meant — **pushed the table down by two rows and left the numbers
unmoved**. A note whose entire job is to fence an unmeasurable predicate now sends the reader to Chronos and
Intelligence (which assert nothing) and calls Skiller's flat `ARCHIVED 2026-07-01` a retraction of itself.
Found independently by **both** readings, and it is one of only four matched predicates.

**The ~2/cycle repair-induction rate is the most stable number in this milestone while every magnitude band
fails.** It held again.

### 2. One false predicate at three anchors in three files — the sentinel/gotenberg generalisation

Seat D's whole in-scope yield is **one** predicate: *"sentinel is the only cross-process edge / the only
service address compose sets."* Refuted by a single line — `GOTENBERG_URL=http://gotenberg:3200` on `backend`
at `docker-compose.yml:57`, with `gotenberg` declared at `:170-171` in the default `core` profile and reached
over real HTTP at `app/internal/converter/gotenberg.go:31`. The `*_RPC_ADDR`-is-zero half of each sentence is
**true**; only the generalisation breaks. The corpus already carries the correct qualified wording at
`architecture_overview.md:321` and contradicts itself at `gotenberg.md:50`, `dependency_map.md:103` and — 12
rows below the offending row, in the same file — `platform-migration-status.md:105`.

**That is one repair propagated to three files with a fourth as the model, not three fixes.** It is also
`CLAUDE.md`'s claim, verbatim, in this repo's own root instructions.

### 3. The one wrong-tree rejection inverts the class iter-99 named

iter-99's wrong-tree cases were seats grading rext claims against the authoring copy. This reading's single
case is the **inverse**: seat D's messenger-row booking measured everything correctly, but against the
**demo's week-old pin** of a repo nothing local even clones, where the map's own header declares an
**origin-based** basis. The adjudicator resolved it with a rule worth carrying:

> **The settling tree follows the claim's SUBJECT.** A claim about what a local stack runs is settled by the
> demo's build pin. A claim about **production infrastructure** is settled by that repo's `origin/main`.

## What this reading establishes, and what it does not

**Establishes:** the corpus carries **at least 24** blocking falsehoods inside clause 5's scope at `8f04d3a`
(a floor — 13-seat union); the ~15-point precision drop is **structural, not variance**, now confirmed at
n=2 with different adjudicators; the two readings are **near-independent**, so **`N̂ = 45.1` is a floor and
the residual is plausibly ~100**; the ref-discipline class has stopped being generated; and iter-100's repair
induced a defect inside the prose it rewrote.

**Does not establish:** that `N` is falling. `N` went 13 → 20 → 28 → **24**, but this union is **13 seats to
the prior 14**, and the upheld rate has been non-constant for two readings running. The series is not
comparable on count alone, and this reading does not claim it is.

**Comparability:** continuous in **instrument** (briefing byte-identical, sha `3858ec53…`, `git log --follow`
showing exactly one commit ever — `012edd2`; verified AFTER copying, with the addendum appended below line
171 and nothing above it edited), same partition, grading rule and scope. **Discontinuous in seat count
(13 vs 14)** and **continuous in the upheld-rate break** first seen at iter-99.

Routed as **`FIX-M257x-iter101-read-union`**, inheriting iter-76's binding conditions: **repair by PREDICATE,
not by anchor**, and re-read after.
