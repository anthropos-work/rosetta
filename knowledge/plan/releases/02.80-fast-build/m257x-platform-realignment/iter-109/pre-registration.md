# iter-109 pre-registration — written BEFORE the read, at corpus `2e3443d`, platform `0c91421`

**This reading answers the one question `TOK-06` was authored to answer: WITH BOTH INFLOWS FENCED, DOES THE
POOL DRAIN?**

`TOK-06` was authored on a composition measurement, not on a mood. iter-103 read `N = 33` and decomposed it:

| inflow | share of iter-103's `N` | fenced at the time |
|---|---|---|
| **clone advance** — version literals, `go.mod` pins, symbol names, line offsets | 20 anchors — **61 %** | *nothing* |
| **the repair's own induction** — prose iter-102 wrote | 7 anchors — **21 %** | *nothing* |
| never-true | 4 — 12 % | the reading |
| unclassified | 2 | — |

*"Inflow is comparable to outflow. A loop with that property does not converge, and running it faster does
not help."* So the strategy stopped repairing and fenced the inflows first: step 0 guard provenance
(iter-105), step 1 the drift fence (iter-106), step 2 the induction checks (iter-107), step 3 the repair
(iter-108, 22 predicates / 23 files, machine-graded **46/46 = 100 % of the upheld union**). **This is step 4,
and it is the payoff.**

**And the subject held still.** All **14 platform clones are at the identical sha** they were at iter-103's
read. The clone-advance inflow had **no opportunity to re-arm**. Whatever this reading returns, it is not
measuring five clones that moved underneath it.

## What this reading can and cannot separate — say it before the number, not after

The 61 % drift share was a classification of the **defect class**. It is *consistent with* two different
stories and iter-103 could not tell them apart:

- **ARRIVAL** — those 20 anchors became false *because* five clones advanced between the two readings. Fence
  the advance, and they do not come back.
- **DETECTION** — those 20 anchors were false all along, and iter-103 found them because mechanical defects
  are what a pass finds when the subtle half has just been repaired. Per-pass recall on this instrument has
  run **33–83 %**; a large undetected standing residual is compatible with every number in the series.

**With the clones frozen, this reading separates them.** Under ARRIVAL, drift should be near-absent and `N`
should collapse. Under DETECTION, the reading samples the standing pool again and `N` holds. That is the
whole design, and it is why the primary rule below is cut where it is.

## The pre-registered verdict rule — TWO metrics, both sealed now

Fixed **before any seat is dealt**, so neither can be re-cut after the number is known.

**Both are banded, because iter-103's sharpest finding was that they disagreed:** anchors went 24 → 33 while
predicates held 22 → 22 — *the same falsehoods, in more places*. **`P` (distinct false predicates) tracks
whether the pool is draining. `N` (distinct anchors) tracks how far each falsehood has spread.** Reporting
one without the other is how iter-101 → iter-103 looked like a rise and was actually a flat.

### PRIMARY — `P`, the distinct-false-**predicate** count

| `P` | verdict |
|---|---|
| **≤ 6** | **THE POOL DRAINS.** With both inflows fenced the standing residual is small. iter-103's decomposition was **ARRIVAL**, and clause 5 is reachable by repair-and-read within a small number of further cycles. |
| **7 – 14** | **PARTIAL DRAIN.** Inflow was a real and substantial component, but a standing pool remains that repair-and-read samples slowly. The loop converges, but the ETA is long and must be stated. |
| **≥ 15** | **THE POOL DOES NOT DRAIN.** A third consecutive reading at ~22-ish predicates, with the clone-advance inflow provably absent, says the 61 % was **DETECTION**, not ARRIVAL. Repair-and-read is then not a converging loop for clause 5, and the milestone needs a different instrument — not a faster one. **This is the outcome that changes the plan, and it is reported first and loudest if it fires.** |

### SECONDARY — `N`, the distinct-**anchor** count (the series metric, kept for continuity)

| `N` | verdict |
|---|---|
| **≤ 8** | drains |
| **9 – 19** | partial |
| **≥ 20** | does not drain |

**If `P` and `N` land in different bands, `P` governs and the disagreement is itself the headline** — it is
the multiplication signal, and iter-108 removed the `:8081` multiplier *structurally* (derive once, point at
it) precisely so that a corrected canonical sentence could not re-multiply. A `P`/`N` split would say the
structural fix did not hold.

**Both are FLOORS in every branch.** They are unions of two passes whose measured per-pass recall has run
33–83 %. A small `P` is evidence about *detection*, and only indirectly about the pool. **No branch of this
rule may be reported as "clause 5 is close" — clause 5 is met by a reading that returns zero, and nothing
else.**

## Bands — 12 of them, each stated so it CAN fail

Prior readings graded **4 of 9**, **3 of 7**, **5 of 9** and **4 of 10**, with *mechanism* claims mostly
holding and *magnitude* guesses mostly failing. That split is the useful shape and these bands do not drift
toward safe.

| # | prediction | band | kind |
|---|---|---|---|
| 1 | per-reading in-scope upheld BLOCKER count (n₁, n₂) | **[3, 16]** each | magnitude |
| 2 | **union `N`** (anchors) | **[4, 19]** | magnitude |
| **2p** | **union `P`** (predicates) — **the primary** | **[3, 14]** | magnitude |
| 3 | **overlap with iter-103's published 22 predicates**, matched on PREDICATE, blind | **[0, 3]** | mechanism |
| 3b | **within-reading `m`** between passes #27 and #28, as a **share of the union** | **[10 %, 55 %]** | mechanism |
| 4 | adjudicator upheld rate (raw) | **[72 %, 94 %]** | mechanism |
| 5 | the two passes' recalls against the union differ by | **≥ 12 points** | mechanism |
| 6 | **wrong-tree rejections** (the briefing-defect class) | **[0, 3]** | mechanism |
| 7 | wrong-construct intra-corpus citations among upheld in-scope blockers | **≤ 4** | mechanism |
| 8 | **platform-drift share of upheld in-scope blockers** | **≤ 25 %** | mechanism |
| 9 | per-seat booked spread over the 14 seats (max − min) | **≤ 8** | magnitude |
| 10 | **repair-induced** — upheld blockers whose anchor sits in prose iters 104–108 wrote | **[0, 4]** | mechanism |
| **11** | **anchors per predicate** (`N`/`P`) — net-new band | **[1.00, 1.35]** | mechanism |

## What each band is actually risking

**#1 / #2 / #2p — the magnitude guesses, and they are NOT centred on zero.** Every magnitude band on this
milestone has failed, in both directions. The honest centre for `P` is **~8**: iter-103's 22 predicates are
repaired at 100 % of the upheld union, the clone-advance inflow is provably absent, and iter-108's repair
was a fifth the prose volume of iter-102's — but the standing-residual floor is ≥ 33 anchors *as measured*,
and nothing has drained it except two repairs whose targets were themselves found by the same detection
process. **`≤ 6` is a real possibility and `≥ 15` is a real possibility, and the design of this reading is
that they mean opposite things about the milestone's remaining plan.**

**#2p is the primary and #2 is the secondary — recorded here so the order cannot be chosen after the fact.**

**#3 — overlap with iter-103's 22. It is a test of the REPAIR, not of independence.** iter-108 reached
**46/46 = 100 %** of the upheld union by machine (`repair_reach_guard` over `iter-103/raw/`, the same code
path that derived the ledger). A re-found iter-103 predicate therefore means the repair did not close it, or
closed it at one site and left a twin. The honest centre is **0–1**, and **≥ 4 fails** — which would directly
contradict the machine reach grade and would be the most valuable failure on the sheet. iter-103's own
version of this band held at **1**.

**#3b — banded as a SHARE this time, and that is deliberate.** iter-101 measured `m` = 4/24 = **17 %**;
iter-103 measured `m` = 20/33 = **61 %** on a **byte-identical** instrument. That range is the entire reason
Chapman is retired. The composition argument predicts the share **falls** as the residual gets subtler:
mechanical defects are found by every competent pass (correlated), subtle ones are not (independent). A
share **≥ 55 %** says the residual is *still* mechanical despite the frozen clones — which would be strong
evidence for the DETECTION story and would corroborate a `≥ 15` primary. A share **< 10 %** says the passes
barely overlap and the standing pool is larger than either union.

**#4 — the precision band, set to test whether iter-103's 97.9 % was a property of the RESIDUAL.** Four
readings held 92.1 / 93.0 / 92.7 / 93.1 %; iter-99 broke it at 78.3 %, iter-101 at 77.8 % raw, and iter-103
returned **97.9 % raw / 100 % separated** — explained at the time as *"a mechanical defect leaves a seat
almost no room to be wrong."* If that explanation is right, precision should fall back now that the
mechanical half is repaired. The band **fails high at > 94 %**, which is the outcome that would say the
explanation was wrong. Reported **twice** — raw, and with the `wrong-tree` class separated — because `N` is
post-adjudication and immune to the briefing defect while the upheld rate is not.

**#5 — the asymmetry, predicted to RETURN.** iter-101's two passes recalled **83.3 %** and **33.3 %** — a
50-point spread. iter-103's collapsed to **9.1 points**, attributed to the mechanical residual. Same
prediction as #3b from the same mechanism, measured a different way: if the residual is subtler now, the
passes should diverge again. Fails if they come within 12 points.

**#6 — the briefing defect, measured for the fourth time.** n=3 so far: **4 → 1 → 1**. This reading changes
one thing: the addendum now also states that a claim about a *fence's own verdict* is settled by the
authoring tree (`D-M257x-103-7`). That is a second sentence about the same distinction, and #6 tests whether
adding it helps or muddies. `> 3` would say two answers are worse than one.

**#7 — kept, and the exposure is smaller than last time.** iter-102 added +214 anchors to the tree and this
band still held at 4. iter-108 added far less prose but **re-pointed four citers and removed a five-site
multiplier**, both of which touch citation quality directly. ≤ 4 holds if the re-pointing was done right;
≥ 5 says wide re-pointing degrades citation quality, which would be a first-class finding about the repair
leg.

**#8 — THE DIRECT TEST OF THE DRIFT-FENCE LEG, and the band is cut hard.** This class went 7/13 → 1/20 →
~1/28 → 1–2/24 → **61 % at iter-103**, when five clones moved and nothing fenced them. This reading has
**zero clones moved** and `clone_drift_guard` shipped (GREEN at the open, with its 2 gradeable pins
matching). **≤ 25 % is a demanding band and it is meant to be**: if drift is still the majority of the
residual with the subject frozen, the drift *fence* is not the answer, because there was no drift to fence
— and that reading points straight at the DETECTION story.

**#9 — tests the partition, not the corpus.** The partition was recomputed (the corpus grew 0.45 % and the
LPT assignment cascaded anyway — disclosed in `ground-truth.md`). Balanced to a 51-line spread over
1506–1557. A booked spread over 8 would say seat-level variance dominates. iter-101 got 4; iter-103 held.

**#10 — the induction rate, measured for the seventh consecutive cycle, and this is the first cycle it was
FENCED.** The ~2-defects-per-repair-cycle rate has held six times, most recently *inside the iter documenting
it*. iter-108 shipped under `anchor_offset_guard` + `repair_postcondition`, **both of which fired on it and
were repaired before it stood**. So the band's floor is **0** for the first time — a `0` here would be the
first break in the series and would be the strongest single piece of evidence that step 2 worked. `> 4`
would say the fences let the rate through unchanged.

**#11 — net-new, and it exists because iter-108 removed a multiplier structurally rather than correcting
it.** The `:8081` cardinality was published to 5 anchors as a canonical wording; iter-108 derived it **once**
in `backend.md` and made `cms.md`/`jobsimulation.md` point at it, on the principle that *a pointer cannot
carry a false cardinality to five places*. Ratio history: **1.09** (iter-101) → **1.50** (iter-103). A ratio
**≥ 1.5** says multiplication re-armed anyway and the structural fix did not hold; a ratio near **1.0** says
it did.

## Binding conditions on the read itself

1. **The instrument is not touched.** Briefing byte-identical, sha `3858ec53…`, `git log --follow` showing
   one commit ever, **re-checked AFTER copying** (both sides printed in `ground-truth.md`). **The known
   defect at line 37 is delivered as-is** and routed, never edited. Superseded ground truth goes in the
   ADDENDUM, below the frozen text, never above it.
2. **Clause 5 is not re-cut, narrowed, argued, or read met any other way.** It is met only by a reading that
   returns **zero**. Four user rulings, and this run does not reopen them. **No branch of the primary rule
   above is a gate verdict.**
3. **No repair inside the measuring pass.** Anything found routes; nothing is fixed.
4. **Ground truth re-derived, not inherited** — every checkout's ref AND its `origin/main` AND its fetch time
   restated, guard family run **with its own fence tree printed**, all before the seats are dealt.
5. **No clone is fetched while this reading is in flight** (§5 rule 41a). Fetch times are re-read at the
   close and published, so a mid-reading move is detectable rather than suspected.
6. **The upheld rate is reported TWICE** — raw, and with the `wrong-tree` class separated.
7. **Adjudicate before reporting `P` or `N`.** Always.
8. **Every seat is committed verbatim the moment it lands**, before adjudication.
9. **`P` and `N` are reported whatever they are.** A rise, a fall and a flat are each a real result. No
   rounding toward the answer the gate wants. The DOES-NOT-DRAIN branch is reported first if it fires.
10. **No point estimate of the pool is quoted** — not in this file, not in the close, not in `state.md`.
    Chapman is retired for this milestone. **Only floors survive: ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`.**
