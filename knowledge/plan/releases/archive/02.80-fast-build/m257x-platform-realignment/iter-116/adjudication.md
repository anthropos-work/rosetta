# iter-116 adjudication — the graded sheet

**Adjudicated before `P` or `N` was computed.** Four independent adjudicators, one per seat-group, each
re-deriving from the platform clones. Their verdict files are committed unedited in `verdicts/`.

## THE NUMBER

> # `P = 37` · `N = 41`
>
> ### The pre-registered `P ≥ 15` branch FIRES — at more than DOUBLE the threshold.
>
> **`TOK-07`'s own pre-registered falsification condition is MET: repair-and-read is REFUTED as a path to
> clause 5 under this instrument.** The next move is a **re-scope conversation with the user**, not an
> eighth strategy revision. That was sealed on 2026-08-06, before iter-111 ran, and it is graded exactly
> as written.

The secondary `N ≥ 20` branch fires with it, so there is no `P`/`N` split to adjudicate.

| | iter-101 | iter-103 | iter-109 | **iter-116** |
|---|---|---|---|---|
| distinct false **predicates** `P` | 22 | 22 | 24 | **37** |
| distinct **anchors** `N` | 24 | 33 | 36 | **41** |
| anchors per predicate | 1.09 | 1.50 | 1.50 | **1.11** |

**Clause 5 is NOT met.** It is met only by a reading that returns **zero**.

## The composition — and why this refutation is USEFUL rather than merely bad

The rise is not evidence that iter-115's repair failed. **Three measurements say it did what it was
designed to do**, and the reading separates that cleanly from why `P` still rose.

### 1. The repair HELD. 21 of iter-109's 24 predicates stayed closed.

Band #3 measured the overlap blind, on predicate: **3 re-found of 24**.

| iter-109 predicate | status at iter-116 |
|---|---|
| *Ant Academy is an internal `@anthropos.work`-gated portal* | **re-found** (adj-3 P6) |
| *All Go services access AI through the shared `ai` module* | **re-found** (the cross-group merged predicate) |
| *the `⚠⚠ M51` block is at `ai-readiness.md:496`* | **re-found, RE-BROKEN** — the repair rewrote the pins to `:476-496`/`:498`/`:500`, then iter-115's own +40 lines above rotted them again. Fourth generation. |
| the other 21 | **closed and stayed closed** |

### 2. The twins really were closed corpus-wide. `N`/`P` collapsed 1.50 → **1.11**.

The ratio is back to iter-101's level. A repair scoped per-predicate over the whole corpus leaves a
residual that is **thin per predicate** — which is exactly what the enumeration was for, and it is the
strongest evidence on the sheet that `TOK-07`'s *mechanism* was right.

### 3. And `P` rose 24 → 37 anyway. That is the VOLUME branch, and it is now the measured answer.

The pre-registration named two stories and this reading separates them:

- **DENOMINATOR** — the pool did not drain because the repair only visited detected sites. **Falsified**:
  the denominator was fixed, the repair reached 71/71 enumerated sites, and 21 of 24 predicates stayed
  closed.
- **VOLUME** — the pool is simply far larger than a reading samples. **This is the answer.**

**The arithmetic that settles it.** iter-115 closed 24 predicates completely. The next reading found 37:

| where the 37 came from | count |
|---|---|
| **induced by iter-115's own repair** (band #10, measured by `git blame` against the iters-110–115 commit range) | **9** |
| re-broken from iter-109's 24 | 3 |
| **standing pool, never detected by any of four prior readings** | **25** |

> **A loop that closes 24 predicates per cycle while creating 9 and sampling only a fraction of what
> remains does not converge.** After four readings and two full repair cycles, a single fresh pass still
> surfaces **25 predicates nobody had ever seen**. That is not a tuning problem.

### 4. The self-poisoning rate GREW with repair volume — the fences did not scale

Band #10's series, measured the same way each cycle:

| cycle | repair volume | repair-induced share of the residual |
|---|---|---|
| iters ≤102 | — | ~21 % |
| iter-108 | +48 lines | **5.6 %** (fences shipped; the lowest ever) |
| **iter-115** | **+177 lines** | **22.0 %** (9 of 41 anchors) |

`anchor_offset_guard` + `repair_postcondition` **fired four times during iter-115 and were repaired each
time** — and 9 defects still got through. The fences work at iter-108's volume and do not hold at
iter-115's. **Repairing faster makes the induction worse, not better**, which closes the last obvious
escape route from the VOLUME finding: you cannot out-run this pool by repairing harder.

## Bands: 7 HELD of 14

**HELD** `#3` `#3b` `#4` `#6` `#8` `#9` `#11` · **FAILED** `#1` `#2` `#2p` `#5` `#7` `#10` `#12`

| # | prediction | band | measured | verdict |
|---|---|---|---|---|
| 1 | per-reading in-scope upheld blockers (n₁, n₂) | [2, 16] each | **n₁ = 26 · n₂ = 22** | **FAILED** — both high |
| 2 | union `N` | [5, 22] | **41** | **FAILED** high |
| **2p** | **union `P`** — the primary | **[4, 15]** | **37** | **FAILED** high |
| 3 | overlap with iter-109's 24, blind, on predicate | [0, 3] | **3** | **HELD** — at the ceiling |
| 3b | within-reading `m` as a share of the union | [10 %, 55 %] | 9/37 = **24.3 %** | **HELD** |
| 4 | adjudicator upheld rate (raw) | [72 %, 94 %] | 50/54 = **92.6 %** | **HELD** |
| 5 | the two passes' recalls differ by | ≥ 12 points | 64.9 % − 59.5 % = **5.4 pts** | **FAILED** — second consecutive |
| 6 | wrong-tree rejections | [0, 3] | **0** | **HELD** |
| 7 | wrong-construct intra-corpus citations | ≤ 4 | **10** | **FAILED** high |
| 8 | platform-drift share of upheld in-scope | [20 %, 60 %] | 14/37 = **37.8 %** | **HELD** |
| 9 | per-seat booked spread over 14 seats | ≤ 8 | **4** (max 6, min 2) | **HELD** |
| 10 | repair-induced (prose iters 110–115 wrote) | [0, 5] | **9** | **FAILED** high |
| 11 | anchors per predicate `N`/`P` | [1.00, 1.45] | **1.108** | **HELD** |
| **12** | net-new — anchors in a multi-pin block | [0, 4] | **6** | **FAILED** high |

Same shape as every prior reading: **6 of 8 mechanism bands held; 3 of 3 magnitude bands failed.** The
magnitude guesses have now failed on every reading of this milestone, in both directions.

### The upheld rate, reported TWICE as required

| | value |
|---|---|
| **raw** | 50 upheld / 54 booked = **92.6 %** |
| **`wrong-tree`-separated** | 50 / 54 = **92.6 %** — *identical*, because `wrong-tree` was **0** |

Rejection classes: **mis-read ×3** (all adj-3, seat r30-D), **ref-discipline ×1** (adj-4). The
`wrong-tree` series is now **4 → 1 → 1 → 0 → 0** — the instrument defect
(`DEF-M257x-iter101-briefing-rext-tree`) has produced no rejection for two consecutive readings while
still delivered unfixed.

### Band #12, net-new, and it is the more interesting failure

Multi-pin blocks are a **defect concentrator**, not merely a hazard to the machine: **6 of 41** upheld
anchors sit in a block naming ≥ 2 platform refs, against a predicted ≤ 4.

And the sharpest instance is the one this reading opened on. **`sentinel.md:5` is a 4-pin block.**
`platform_predicate_guard` went RED on it for the **wrong proposition** — grading the compose-count
sentence at the first of the four refs instead of the ref it names — while the line **does** carry a
different, genuinely false proposition that a seat found independently (adj-3 P1: the published messenger
grep is stated as returning one hit; it returns zero). **The guard pointed at the right line for the
wrong reason.** That is not exoneration of the guard; it is the strongest possible argument for fixing
its pin-scoping, because the line it flagged deserved flagging.

### Band #7 — the wrong-construct citation class is now the single largest

**10 of 37 predicates (27 %)** are a corpus citation landing on the wrong construct — a self-pin or a
cross-file pin that names lines holding something else. Predicted ≤ 4.

The corpus is now large enough, and repaired often enough, that **it mis-cites itself more often than it
mis-describes the platform.** The class is mechanical, and it is the one class a machine could close
outright — which is a concrete input to the re-scope conversation, not a repair to make here.

## Per-reading and per-seat detail

| | reading #29 | reading #30 |
|---|---|---|
| booked | 28 | 26 |
| in-scope upheld blockers | **26** | **22** |
| predicates found | 24 | 22 |
| **recall against the union of 37** | **64.9 %** | **59.5 %** |
| found by this reading only | 15 | 13 |

Shared by both readings: **9 predicates**. Per-seat booked: `r29` A2 B5 C5 D4 E5 F3 G4 · `r30` A2 B3 C3
D5 E6 F4 G3 — max 6, min 2, **spread 4** over 14 seats.

**Both `P` and `N` are FLOORS.** Two passes with measured recalls of 64.9 % and 59.5 % against their own
union cannot bound what neither saw. **No point estimate of the pool is quoted.** Only floors survive:
**≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`, ≥ 36 at `ac48e5b`, ≥ 37 at `f581de09`.**

## Provenance

- 14/14 seats landed, **0 lost**, each committed verbatim before adjudication (11 commits).
- 4/4 adjudicator verdicts committed unedited.
- Pre-registration sealed at **`85f6f1c`**, before the first seat was dealt. No band was re-cut.
- Cross-group dedup: **exactly one** collapse — adj-1 P1 + adj-2 P5 + adj-3 P8 → one `ai`-fold predicate
  over 4 anchors, verified by the coordinator against `app` `go.mod` (no `anthropos-work/ai` require) and
  `app/internal/ai/` in-tree with a `module_import_guard_test.go`. 39 group-level predicates → **37**.
- **No repair was taken inside this pass.** The guard-family RED disclosed at the open is routed, not
  fixed.
