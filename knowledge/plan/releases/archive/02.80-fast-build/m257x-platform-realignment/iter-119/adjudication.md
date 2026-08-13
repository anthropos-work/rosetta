# iter-119 adjudication — the graded sheet

**Adjudicated before `P` or `N` was computed.** Four independent adjudicators, one per seat-group, each
re-deriving from the platform clones. Their verdict files are committed unedited in `verdicts/` (`f65f4a4`).

## THE NUMBER

> # `P = 22` · `N = 28`
>
> ### The pre-registered `P ≥ 19` branch FIRES. **`TOK-08` is REFUTED.**
>
> **The USER's re-scope — enumerate the mechanical classes exhaustively instead of sampling them — does
> not reach clause 5 either.** The rule sealed at `577446b`, before a single sweep line was written, says:
> *"STOP. Do NOT author a successor strategy. Report the refutation and hand the milestone back for a
> scope decision from the user."* That is what this iter does. **There is no `TOK-09`.**

| | iter-101 | iter-103 | iter-109 | iter-116 | **iter-119** |
|---|---|---|---|---|---|
| distinct false **predicates** `P` | 22 | 22 | 24 | 37 | **22** |
| distinct **anchors** `N` | 24 | 33 | 36 | 41 | **28** |
| anchors per predicate | 1.09 | 1.50 | 1.50 | 1.11 | **1.27** |
| booked / upheld | — | — | — | 54 / 50 | **36 / 33** |

**Clause 5 is NOT met.** It is met only by a reading that returns **zero**.

---

## ⚠️ READ THIS BEFORE READING `P = 22` AS PROGRESS. IT IS NOT.

`P` fell **37 → 22**, a drop of **40.5 %** — just short of the 50 % the working-branch required. The
tempting reading is *"the method nearly worked."* **That reading is false, and this sheet's own
pre-registered band #3 is what refutes it.**

**The corpus changed by FIVE LINES between the two readings.** 2,590 citations enumerated across the two
census iters, 8 false, **2 of the 8 in this scope**, zero net lines, and **not one of iter-116's 37
predicates was repaired** (verified: the 8 repairs were `#anchor` fragments in
`platform-migration-status.md` ×4, `architecture_overview.md:48`, and three out-of-scope files; no
iter-116 predicate anchors at any of them). **A pool cannot drain 40 % when nothing was removed from it.**

So where did the 15 go? **Band #3 answers it, and the answer is the most important thing this reading
produced:**

| measurement | value |
|---|---|
| iter-116 predicates **re-found** by this reading, matched blind on predicate | **13 of 37 (35.1 %)** |
| iter-116 predicates this reading **did not see** | **24 of 37** |
| predicates this reading found that **iter-116 never saw** | **9** |
| **union floor over the two readings** | **≥ 46** |

> ### `P` FELL from 37 to 22 while the measured FLOOR ROSE from 37 to **≥ 46**.
>
> The 24 predicates this reading missed were adjudicated UPHELD against the same clones and a corpus five
> lines away. They did not stop being false because a second panel failed to notice them. **The drop in
> `P` is sampling, not drainage.**

**This is the first test-retest measurement ever taken on this instrument**, and it was only possible
because iter-119 is the milestone's first true seat-level replicate — identical partition (proven by
re-running the partitioner at both refs), identical clone shas for a fourth consecutive reading, a corpus
differing by 5 in-place lines that changed no proposition. It was pre-registered as band #3 *before* the
number, with the low-overlap branch named in advance as *"the alarming outcome… it would bear on the
refutation more than the primary itself does."*

**It did, and it does.** Every `P` this milestone has published — **22, 22, 24, 37, 22** — is a sample of
roughly a third of a standing pool, not a count of it. The series was read for six weeks as though its
movements meant something about the corpus. **Most of the movement is the instrument.**

---

## The composition of the 22

### By class — and the census's own prediction is confirmed, hard

| class | count | share |
|---|---|---|
| **intra-corpus citation landing on the wrong CONSTRUCT** | **8** | **36 %** |
| platform-facing claims false against source | 14 | 64 % |

**Band #7 predicted ≤ 6 and FAILED at 8** — and the failure is the finding `D-M257x-117-2` pre-registered
at iter-117, in those words, before either census closed:

> *"the machine-reachable half of class 1 is largely DISJOINT from the half the readings book… of 387
> lines carrying a bare `:NN` pin only 4 are machine-resolvable. Little `P` movement should be expected."*

**The census closed the RESOLUTION half of intra-corpus citation at 100 % reach and 0 findings. The
CONSTRUCT half is now a LARGER share of the pool than before the sweep — 27 % at iter-116, 36 % here.**
Both facts are true at once. A machine can check that `cms.md:216`'s link resolves to a file that exists;
it cannot check that the lines it names hold the **Data** bullet rather than the **Domain** bullet. The
eight:

| predicate | anchor |
|---|---|
| `cms.md`'s **Data** bullet is at `:44-47` (it is the Domain bullet; Data is `:48-51`) | `cms.md:216` |
| `cms.md`'s **Studio** bullet is at `:70-71` (it is the Events bullet; Studio is `:75-76`) | `cms.md:240` |
| `ai-readiness.md`'s `✅ CORRECTED M219` block spans `:476-496`, `⚠⚠ M51` opens `:498` | `ai-readiness.md:54-58` |
| `ai_architecture.md`'s three bare self-citations (`:86`, `:87`, `:95`) name their constructs | `:92`, `:94`, `:294-295` |
| `service_taxonomy.md`'s archive-state note is at `:142` | `service_taxonomy.md:403` |
| *"There is no `graphql` profile"* is at `service_taxonomy.md:67-68` | `service_taxonomy.md:406-407` |
| `service_taxonomy.md:62` is the Tier-1 Database characteristic bullet | `hiring.md:47-49` |
| `roadrunner.md`'s *"Upstream consumers: none (orphaned)"* is at `:124`, below `:130` | `roadrunner.md:130` |

**Five of these eight were booked at iter-116 too, and are still false.** They were not repaired because
no repair is taken inside a measuring pass and both intervening iters were censuses.

### Graded by CONSEQUENCE, not by class (pre-registration condition 11)

One predicate is worth naming on its own, regardless of its class size:

> **`clerk-integration.md:40` — *"Clerk sign-in tokens are minted ONLY for app-native admin
> impersonation."* Two other live minting sites exist.**

This is a **security-surface understatement**: a reader auditing where privileged impersonation tokens can
be created is sent to one site when there are three. **It was booked and upheld at iter-116 as well
(adj-2 P7), and it is still in the corpus** — this is its second consecutive reading as a known, open,
unrepaired defect. It is the iter-115 `bash -c` class (a claim that inverts or understates a shipped
security property) and it is reported first among the individual findings for that reason.

---

## Bands: 13 HELD of 15

Prior readings: **4 of 9**, **3 of 7**, **5 of 9**, **4 of 10**, **7 of 13**, **7 of 14**. This is the
best band performance of the milestone — **and that is a statement about the coordinator's calibration
improving, not about the corpus.**

| # | prediction | band | measured | verdict |
|---|---|---|---|---|
| 1 | per-reading in-scope upheld blockers | [8, 34] | **16 · 17** | ✅ |
| 2 | union `N` | [22, 52] | **28** | ✅ |
| **2p** | **union `P`** — primary | [20, 44] | **22** | ✅ |
| **3** | **overlap with iter-116's 37** | [10, 30] | **13** | ✅ |
| 3b | within-reading `m` as share of union | [10 %, 55 %] | **9/22 = 40.9 %** | ✅ |
| 4 | adjudicator upheld rate (raw) | [72 %, 94 %] | **91.7 %** | ✅ |
| 5 | two passes' recalls differ by | ≥ 12 pts | **63.6 % vs 77.3 % = 13.7** | ✅ |
| 6 | wrong-tree rejections | [0, 3] | **0** | ✅ |
| 7 | intra-corpus citation defects | ≤ 6 | **8** | ❌ |
| 8 | platform-drift share | [20 %, 60 %] | **63.6 %** | ❌ |
| 9 | per-seat booked spread | ≤ 8 | **5** | ✅ |
| 10 | repair-induced | [0, 1] | **0** | ✅ |
| 11 | anchors per predicate | [1.00, 1.45] | **1.27** | ✅ |
| 12 | anchors in a multi-pin block | [2, 12] | **3 of 28** | ✅ |
| **13** | re-found `SMALL-CLASS-ADJUDICATED` | [0, 4] | **2 of 15** | ✅ |

### The upheld rate, reported TWICE as required

**Raw: 33/36 = 91.7 %.** **`wrong-tree`-separated: 33/36 = 91.7 %** — identical, because `wrong-tree` was
**0** for the third consecutive reading (series **4 → 1 → 1 → 0 → 0 → 0**). All three rejections were
`ref-discipline`, the class the brief predicts and instructs adjudicators to filter: a seat grading a
pinned claim against newer evidence. That class has now run **20 occurrences across six readings and
contributed ZERO to any graded count.**

### Band #5 held for the first time in three readings — and it is corroborating evidence

Predicted from the same mechanism as band #3: if the residual is genuinely subtle, two independent passes
should diverge. iter-109 and iter-116 both **failed** this band. It **holds here at 13.7 points**, and it
points the same way band #3 does: **the passes are sampling, not enumerating.** Seat E booked **0** in one
pass and **3** in the other; seat F booked **5** and **1**. Same seat, same files, same frozen clones.

### Band #8 failed high, narrowly, and the boundary is disclosed

63.6 % against a `[20 %, 60 %]` band. The classification boundary is judgement: I counted *"a citation
into platform source naming the wrong construct"* (3 predicates) as platform-facing rather than
intra-corpus, since the target is platform source. Counted the other way the share is 50 % and the band
holds. **The band is graded as FAILED because that is how it was cut**, and the alternative is disclosed
rather than substituted.

---

## Band #13 and `FIX-M257x-iter113-adjudication-is-judgement` — the direct audit

iter-113 verdicted 16 predicates "small class": **1 `SMALL-CLASS-PROVEN` (P20, zero headroom, nothing
judged)** and **15 `SMALL-CLASS-ADJUDICATED`, resting on 254 candidate exclusions read by one agent.**

**This reading re-books 2 of the 15 judged — and 0 of the 1 proven.**

| iter-113 id | statement | status here |
|---|---|---|
| **P23** | *Ant Academy is an internal `@anthropos.work`-gated portal* | **RE-BOOKED** (adj-3 P5, 5 anchors over 2 files) — and re-booked at iter-116 too |
| **P08** | *the `⚠⚠ M51 iter-08/09` block is at `ai-readiness.md:49x`* | **RE-BOOKED** (adj-1 P4) — **fifth generation** of this same anchor rotting |
| P20 (**PROVEN**) | *`archiveStudioTask` is declared in `cms_queries.graphqls:106`* | not re-booked |
| the other 13 judged | — | not re-booked |

**Measured error rate of the judged verdicts: ≥ 2/15 = 13.3 %.** P23's exclusion list dismissed **12**
candidates as publishing a different proposition; at least four of them publish exactly this one. The
single **proven** verdict held. **That is the asymmetry the open FIX predicted, now measured rather than
suspected.**

### And to answer it precisely: the sweep converted NONE of the 16 from judged to proven

The two censuses enumerate **citations**; a small-class verdict is a claim about a **ceiling** (*no other
site publishes this predicate*). **Neither census measures a ceiling**, so neither can promote a judged
verdict. **P20 remains the only `SMALL-CLASS-PROVEN` of the 16, proven at iter-113 by its own
zero-headroom ceiling — not by `TOK-08`.** The 15 remain judged, now with a measured error rate attached.

---

## Per-reading and per-seat detail

| | reading #31 | reading #32 |
|---|---|---|
| booked | 18 | 18 |
| in-scope upheld blockers | **16** | **17** |
| predicates found | **14** | **17** |
| **recall against the union of 22** | **63.6 %** | **77.3 %** |
| found by this reading only | 5 | 8 |

Shared by both readings: **9 predicates**. Per-seat booked: `r31` A2 B3 C3 D2 E0 F5 G3 · `r32` A2 B3 C3
D4 E3 F1 G2 — max 5, min **0**, **spread 5** over 14 seats.

**Both `P` and `N` are FLOORS**, and this reading has quantified how weak a floor `P` is: two passes at
63.6 % and 77.3 % recall **against their own union**, and the union itself recovering only 35.1 % of a
prior reading's adjudicated set. **No point estimate of the pool is quoted.** Chapman stays retired.

> ### Floors, updated
>
> **≥ 24 at `8f04d3a` · ≥ 33 at `e6aed2e` · ≥ 36 at `ac48e5b` · ≥ 37 at `f581de09` · ≥ 46 at `194361e4`.**
>
> The new floor is the **union** of iter-116's 37 (none repaired, all still live) and this reading's 9
> net-new. It is the first floor on this milestone derived from two readings of one corpus rather than
> one, and it is **higher than any single reading has returned.**

---

## Provenance

- **14/14 seats landed, 0 lost**, each committed verbatim before adjudication (`5b23559`, `c18d56b`).
- **4/4 adjudicator verdicts committed unedited** (`f65f4a4`).
- **Pre-registration sealed at `4d4530d`**, before the first seat was dealt. **No band was re-cut.**
- Instrument sha `3858ec53…`, re-checked after copying **and** after the addendum was appended; `diff`
  empty both times; `git log --follow` one commit ever (`012edd2`).
- Partition **reproduced at both refs as a control** — identical to iter-116's, seat for seat.
- Cross-group dedup: **exactly one** collapse — adj-1 P1 (`askengine.md`) + adj-3 P6
  (`architecture_overview.md:80`) → one `ai`-fold predicate. **This is the same single collapse iter-116
  made** (its adj-1 P1 + adj-2 P5 + adj-3 P8), so the two numbers are comparable on this axis. 23
  group-level predicates → **22**.
- **A granularity caveat, disclosed rather than buried:** iter-116's adj-2 split `ai_architecture.md`'s
  three stale self-citations into **three** predicates (its P1/P2/P3); iter-119's adj-2 collapsed the same
  three into **one** with three anchors, and argued the point explicitly. **Had this reading split them as
  iter-116 did, `P` would be 24, not 22.** `P` is therefore comparable to ±2 across readings on
  adjudicator granularity alone — a second, independent reason the series' movements are smaller than they
  look.
- **No repair was taken inside this pass.** The guard-family RED disclosed at the open
  (`FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`) is routed, still open, still unfixed —
  second consecutive reading.
- **No clone was fetched.** All 14 platform clones verified identical at close and open; fetch times
  unchanged (platform-side most recent 2026-08-06 12:15, predating the seal).
