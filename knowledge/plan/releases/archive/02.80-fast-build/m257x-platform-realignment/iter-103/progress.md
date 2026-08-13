# iter-103 — closeout

**Shape:** tik · `iter_shape: reading` · the MEASURING pass, no repair taken inside it · strategy `TOK-04`.

## The one-line answer

**`N = 33`. The pre-registered rule's `≥ 23` branch fired: THE BURN-DOWN LEG DOES NOT REACH THE RESIDUAL.**

The rule was sealed in its own commit (`04cbcfc`) before the first seat was dealt. It is graded here exactly
as written. Nothing about it was re-cut, re-centred, or re-read once the number was known — which is the
entire reason a number that disappoints is worth anything.

Full graded sheet: [`adjudication.md`](adjudication.md).

## What ran

| | |
|---|---|
| corpus | `e6aed2e` · scope `corpus/services/**` + `corpus/architecture/**` (40 files, 10,646 lines) |
| readings | **#25: 7 of 7 seats** · **#26: 7 of 7 seats** — 14 blind seats, identical recomputed partition |
| bookings | 48 → **47 UPHELD / 1 REJECTED** |
| adjudicators | 4, grouped by seat LETTER so both readings of one file set land with one grader |
| clone refs | frozen and **proven** frozen — every HEAD, `origin/main` and fetch timestamp re-read at the close is identical to the open. No fetch (§5 rule 41a) |

**Three seats were re-dealt, not re-run.** Run 67 died on a session limit with reading #26 at 4 of 7.
`r26-A`, `r26-E` and `r26-G` were re-dealt under the byte-identical instrument and the identical partition —
the partition is part of the instrument, so it was **not** re-balanced for a 3-seat batch. Every seat in both
readings was committed **verbatim, before adjudication**. That discipline has now saved seat work three
times: seven seats through a network drop, four through a session limit, three through this run. **The cost
of a mid-flight death is now bounded by the seats in flight, not by the reading.**

## The number, and the number underneath it

| | iter-101 | iter-103 |
|---|---|---|
| distinct false **predicates** | **22** | **22** |
| distinct **anchors** | 24 | **33** |
| anchors per predicate | 1.09 | **1.50** |

**By predicate the pool did not move at all.** After a 52-anchor / 98-site repair, the corpus carries the
same number of distinct false propositions — in more places.

## Bands: 4 HELD of 10, and six of them failed for ONE reason

`#3` · `#6` · `#7` · `#9` held. `#1` `#2` `#3b` `#4` `#5` `#8` `#10` failed.

**#3 held at 1, and it is the band that exonerates the repair.** Exactly one of iter-101's 22 predicates
survives — `prod-terraform-8081`, at `skiller.md:19`, an anchor **iter-102's own repair map listed as a twin
and flagged `SEAT 9 (?)`**. Every other twin was closed. **21 of 22 predicates closed, confirmed blind by an
independent instrument.** The repair leg reaches what it aims at.

**#3b (m = 20 vs [1,7]) · #4 (97.9 % vs [74,88]) · #5 (9.1 pts vs ≥15) · #8 (61 % vs ≤10 %) failed
together, and they are one finding.** The instrument was byte-identical across iter-101 and iter-103; only
the subject moved. iter-102 repaired the residual's *subtle* half — that is the half a reading books — and
what remains is dominated by **mechanically checkable drift**: a version literal, a `go.mod` pin, a symbol
name, a line offset. A mechanical defect is found by **every** competent pass (so overlap explodes and the
recall spread collapses) and leaves a seat almost no room to be wrong about it (so precision goes to 98 %).

> **Precision, overlap and inter-pass independence are properties of the RESIDUAL'S COMPOSITION, not of the
> instrument.** Three of those four numbers moved in the direction that *flatters* the reading. None of them
> is evidence the reading got better.

**#10 failed high by one — 7 of 33 anchors sit in prose iter-102 wrote** — and the mechanical count
understates it, in two shapes worth naming:

- **A canonical wording multiplies its own defects.** iter-102 closed `prod-terraform-8081` by replacing an
  unmeasurable assertion with a sentence saying the `backend.internal.anthropos:8081` literal has *"one
  occurrence anywhere in the clone set."* **It has six** — five inside `rosetta-extensions`, which the same
  sentence's own 13-repo / 44-`.tf` denominator counts as one of its repos. **The replacement is self-refuting
  against its own stated denominator**, and it shipped to five anchors. Six seats found it independently.
- **A repair rotted an anchor by inserting prose above it — the exact mechanism iter-101 booked against
  iter-100, one cycle later.** `architecture_overview.md:321` **was** the correct local-stack line at
  `8f04d3a`. iter-102 inserted a production-topology block above it; the wording moved to `:331`; every
  citation to `:321` stayed put and now names the **production Cosmo Router** — the opposite topology.
  Measured corpus-wide: **4 sites cite `:321`, 0 cite `:331`.** The reading found 2 of the 3 in-scope sites
  and **missed `backend.md:54`**, which sat inside seat E's own file set in both readings.

## Why `N` did not fall, stated as a mechanism rather than a mood

Repair reaches its targets (band #3). The residual is nonetheless fed by two inflows repair does not touch:

1. **Clone advance — 61 % of `N`.** Five clones moved between the two sheets. Neither platform guard fences
   version literals, `go.mod` pins, or line offsets: `platform_alignment_guard` fences `repos.yml`
   membership, `platform_predicate_guard` fences compose profile tokens. This inflow is invisible until a
   reading finds it.
2. **The repair's own induction — 7 of 33 anchors**, including the two clusters above.

**Inflow is comparable to outflow. A loop with that property does not converge, and running it faster does
not help.** That is the finding the pre-registration said would outrank the number, and it does.

## Chapman is retired for this milestone

| | `m` | union | share found by both | `N̂` |
|---|---|---|---|---|
| iter-101, within-reading | 4 | 24 | **17 %** | — |
| iter-99 × iter-101, cross-reading | 6 | — | — | **≈ 102.6** |
| **iter-103, within-reading** | **20** | **33** | **61 %** | **34.9** |

The estimator's load-bearing assumption has now been measured **at both extremes on one unchanged
instrument**. Independence is therefore not a property of the instrument — it is a property of *what is left
to find*. Subtle residual → independent passes → large `N̂`. Mechanical residual → correlated passes → small
`N̂`. **`N̂ ≈ 103` is neither corroborated nor refuted; it is unestimable by this method.**

**Only the floor survives: ≥ 24 at `8f04d3a`, ≥ 33 at `e6aed2e`.** Both are two-pass unions, both are floors,
neither is a pool size. Track `N` and the predicate count directly — they need no assumption at all. The
series 16.7 → 29.4 → 45.2 → ~103 remains **four successive corrections to an underestimate, not four
measurements of a growing pool**; that reading is unchanged and is still the right one. Stop quoting a point
estimate from it.

## Gate

**Unchanged at 4 of 5, and this reading moved nothing.** Clauses 1 and 2 were closed by the **concurrent
lane**, at platform `0c91421`, and this iter states it that way everywhere. Clause 2 is **MET WITH
DISCLOSURE** — a freshly built stack failed the first full run **29/1 in 2 of 2 attempts** — and any artifact
stating clause 2 without the intermittency is wrong. **Clause 5 is the only open one, `N = 33` leaves it
open, and it was not re-cut, narrowed, reinterpreted or argued.** Four user rulings.

## Landed at the close (deferred by `D-M257x-103-0` until the last adjudicator returned)

Both `platform-alignment.md` §5 amendments are now written. They were deliberately **not** written mid-read,
because the frozen instrument tells every seat to read §5 in full — editing it mid-flight would have meant
some seats graded under the old rule set and some under the new, **the exact defect §5 rule 41a exists to
forbid, one level up: the instrument is part of the ground truth too.**

- **§5 rule 49 — a measurement of a concurrently-mutated surface is timestamped, not standing.** You cannot
  refute another observer's report of such a surface with your own later snapshot. Carries the tag-count
  reversal (`D-M257x-103-1`) as its worked example, and keeps the peeled-`^{}` miscount as a **caveat** —
  real, persuasive, and not what happened.
- **§5 rule 41a — a new subsection stating what it CAN and CANNOT enforce.** It binds lanes; it cannot bind
  `ensure-clones.sh`, which fetches on every bring-up and cannot be suppressed. A reading that overlaps a
  bring-up **records the fetch and treats the affected refs as MOVED** rather than asserting the rule held.
  iter-103 is the worked example in the other direction: refs and fetch timestamps measured identical at both
  ends.

**`DEF-4` corrected in the milestone record.** `progress.md` claimed `terraform/main.tf` was *"byte-identical"*
across the `app` advance and that *"the entire residual is a LABEL."* Re-measured: `main.tf` is 1 insertion /
1 deletion (an `error_message` prose string) and `variables.tf` is **+37 / −12** (738 → 763) — 49 lines that
were never in the residual accounting. **The conclusion survives and was re-verified** (`main.tf:181` is
`service_desired_count = 1` at both refs, so no *cited* construct moved); **the evidence sentence does not.**
It was stronger than its evidence, in the direction that made the conclusion cleaner — this milestone's own
class, in this milestone's own records.

## A guard verdict depends on which rext tree ran it — and I nearly published the false half (`D-M257x-103-7`)

The post-edit guard run came back **2 RED**, contradicting this iter's own ground-truth sheet
(`14 GREEN · 0 RED`). Reproduced at `e6aed2e` via a read-only `git archive`, so not my edit. The quotable
conclusion was right there — *"the sheet asserted a verdict it did not have"* and *"a fence names 8 in-scope
sites the double reading missed, so `N ≥ 41`"* — **and both sentences were false.**

I had run the fence from the **pinned per-stack clone** (`09d06070`). From the **authoring copy**
(`944fc4a2`) both guards are **GREEN**, at both subjects. The entire difference is one file —
`claim_twin_waivers.json`, +40 lines, rext `944fc4a` *"the 8 acknowledged-site waivers"* — and **the 8 RED
sites are exactly the 8 waived sites.** `ground-truth.md` was correct; re-confirmed after the amendments at
**14 GREEN · 0 RED · 0 could-not-check · 3 not-run**, identical to iter-101's and iter-102's opens.

> **A guard VERDICT is not stack behaviour.** §5 rule 45's *"the settling tree follows the subject"* sends a
> claim about what the tooling **does on a stack** to the pinned clone — but a fence's verdict is a
> measurement taken with that fence's **configuration**, so it is settled by the tree the configuration lives
> in. Run from the pinned clone, you measure **last release's fence**, and every waiver added since reads as
> a fresh RED at sites nobody touched.

`guard_family.py` prints the corpus sha and the platform sha and **not its own** — the one input that decides
the verdict is the one the output does not state. This is `DEF-M257x-iter101-briefing-rext-tree` **inverted**:
band #6 measured the seat-facing half at 4 → 1 → 1; this is the coordinator-facing half, at 8 sites, once.
**No corpus defect, `N` unchanged at 33, no band moves.**

**Three times in one iteration the milestone's class landed on the milestone's own apparatus** — `DEF-4`'s
over-strong evidence sentence, `D-M257x-103-1`'s single-instant snapshot, and this. **Each time what caught
it was re-measuring rather than reasoning**, and §5 rule 49 — written earlier in this same iter — is the rule
that says a disagreement between two observers is first evidence the *surfaces* differ.

## Routed

- **`FIX-M257x-iter103-read-union`** — 22 predicates / 33 anchors, **by claim, not by file**, with two riders:
  a canonical sentence published to ≥3 sites must be verified against **its own stated denominator** before it
  is multiplied; and a repair that inserts lines above a cited anchor must **re-point the citers**.
- **`FIX-M257x-iter103-drift-fence-gap`** — net-new, and it outranks the repair. **61 % of `N` is an unfenced
  class.** Repairing those 20 anchors without a fence just re-arms them at the next clone advance.
- **`FIX-M257x-iter103-assignment-context-bleed`** — carried in from `D-M257x-103-2`; rate established
  (2/2 cold, 4/4 warm; the write always lands, so the *baseline* is over-read by one), the harness's own
  `baseline-settle-fence` hypothesis refuted by a 60-sample probe. Not repaired here — measuring pass.
- **`DEF-M257x-iter103-aws-bind-provenance`** — stays **OPEN with both measurements recorded and neither side
  asserted**. A correction is a claim too.
- **`FIX-M257x-iter103-guard-tree-provenance`** — net-new (`D-M257x-103-7`). `guard_family.py` and each
  member print the rext tree **path and sha** they ran from, and ground-truth sheets record it beside the
  corpus and platform shas. A verdict whose deciding input is unstated is not a verdict.
- **`DEF-M257x-iter101-briefing-rext-tree`** — stays open, stays delivered-unfixed. Third measurement:
  **4 → 1 → 1**. Band #6's question is answered: an addendum **can** carry ground truth a frozen instrument
  gets wrong, without editing it. All fourteen seats stated which rext tree they read.

## Housekeeping

No stack was brought up, torn down or reconfigured. `stack-demo/**` untouched. **No clone fetched.** No tag
cut — rext stays on `main` at `944fc4a2`, two commits past `fast-build-m257x-iter-101`, both folding into the
next cut. Zero platform-repo edits.
