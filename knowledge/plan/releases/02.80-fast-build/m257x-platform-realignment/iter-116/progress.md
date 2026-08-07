**Type:** tik · `iter_shape: reading` · **`TOK-07` step 3 — the read.** No repair taken inside it.

# iter-116 — the repair worked, and the pool is bigger than the instrument

## The one-line answer, reported first and loudest because the branch fired

**`P = 37` predicates · `N = 41` anchors. The pre-registered `P ≥ 15` branch fired at more than double
the threshold, and with it `TOK-07`'s OWN pre-registered falsification: repair-and-read is REFUTED as a
path to clause 5 under this instrument.**

The rule was sealed in `85f6f1c` before the first seat was dealt and is graded exactly as written. No
band was re-cut. **Do NOT author `TOK-08`** — `TOK-07` named that outcome on 2026-08-06, before iter-111
ran, and named the next move: **a re-scope conversation with the user.**

Full graded sheet: [`adjudication.md`](adjudication.md).

| | iter-101 | iter-103 | iter-109 | **iter-116** |
|---|---|---|---|---|
| distinct false **predicates** | 22 | 22 | 24 | **37** |
| distinct **anchors** | 24 | 33 | 36 | **41** |

**Clause 5 is NOT met.** It is met only by a reading that returns zero.

## What makes this a USEFUL refutation and not just a bad number

**The strategy's mechanism worked. Its premise was still wrong.** Three measurements separate those.

1. **The repair HELD.** Band #3, blind, on predicate: **3 re-found of iter-109's 24**. Twenty-one stayed
   closed. iter-115's corpus-derived-per-predicate enumeration did what it was built to do.
2. **The twins really were closed corpus-wide.** `N`/`P` collapsed **1.50 → 1.11**, back to iter-101's
   level. The residual is thin per predicate — the signature of an enumeration that reached.
3. **And `P` rose 24 → 37 anyway.** The pre-registration's two stories separate cleanly: **DENOMINATOR is
   falsified, VOLUME is the answer.**

### The arithmetic that settles it

iter-115 closed 24 predicates completely, at 71/71 enumerated sites. The next reading found 37:

| where the 37 came from | count |
|---|---|
| **induced by iter-115's own repair** (`git blame` against the iters-110–115 commit range) | **9** |
| re-broken from iter-109's 24 | 3 |
| **standing pool, never detected by any of four prior readings** | **25** |

> **A loop that closes ~24 predicates per cycle while creating ~9, and samples only a fraction of what
> remains, does not converge.** After four readings and two full repair cycles, one fresh pass still
> surfaced **25 predicates nobody had ever seen.**

### And the obvious escape route is closed by measurement

*"Repair harder"* fails: band #10's induction series is **21 % → 5.6 % → 22.0 %**, and the middle value
is iter-108's **+48**-line repair while the last is iter-115's **+177**. `anchor_offset_guard` and
`repair_postcondition` **fired four times on iter-115 and were repaired each time**, and 9 defects still
got through. **The fences hold at iter-108's volume and do not scale to iter-115's.**

## Bands: 7 HELD of 14

HELD `#3` `#3b` `#4` `#6` `#8` `#9` `#11` · FAILED `#1` `#2` `#2p` `#5` `#7` `#10` `#12`

Same shape as every prior reading — **6 of 8 mechanism bands held, 3 of 3 magnitude bands failed.**

Upheld rate **92.6 % raw / 92.6 % `wrong-tree`-separated** — identical, because `wrong-tree` was **0**
for the second consecutive reading (series **4 → 1 → 1 → 0 → 0**).

Two failures are findings in their own right:

- **#7 — the wrong-construct intra-corpus citation class is now the LARGEST single class: 10 of 37
  (27 %)**, against a predicted ≤ 4. The corpus now **mis-cites itself more often than it mis-describes
  the platform.** It is a mechanical class, and that is a concrete input to the re-scope conversation.
- **#12 (net-new) — multi-pin blocks are a defect concentrator, 6 of 41 against ≤ 4.** The sharpest
  instance is the line this reading opened on: `platform_predicate_guard` went RED on `sentinel.md:5`
  for the **wrong proposition**, and the line turned out to carry a **different, genuinely false** one
  that a seat found independently. **The guard pointed at the right line for the wrong reason.**

## Provenance

14/14 seats landed, **0 lost**, each committed verbatim before adjudication. 4/4 adjudicator verdicts
committed unedited. 54 booked → 50 upheld / 4 rejected; 48 in-scope upheld. Exactly one cross-group
predicate collapse, verified by the coordinator against `app` `go.mod` and `app/internal/ai/`.

**All 14 platform clones read identical at the close to the open, fetch times unchanged** — §5 rule 41a
is provable, not asserted. No repair was taken inside the pass.

## Close — 2026-08-07

**Outcome:** `P = 37` / `N = 41` against a rule sealed before the first seat. **The `P ≥ 15` branch fired
at more than double the threshold, meeting `TOK-07`'s own pre-registered falsification: repair-and-read
is REFUTED as a path to clause 5 under this instrument.** The reading separates *why*, and the separation
is the deliverable: iter-115's repair **held** (only 3 of 24 predicates re-found; `N`/`P` collapsed
1.50 → 1.11), so **DENOMINATOR is falsified and VOLUME is the answer** — 25 of the 37 are standing pool
no prior reading ever detected, and **9 were induced by iter-115's own repair**, whose induction rate rose
5.6 % → 22.0 % as repair volume went +48 → +177 lines. Repairing harder makes it worse. 7 of 14 bands
held; upheld rate 92.6 % raw and 92.6 % `wrong-tree`-separated.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: **y — `TOK-07`'s
pre-registered `P ≥ 15` falsification fired at `P = 37`; the milestone's own rule routes this to a
user re-scope conversation, not to an eighth strategy revision** — (4) user-blocker: n — (5)
cap-reached: n (1 tik this session) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-3
**Decisions:** `D-M257x-116-1` (the guard family's RED at the open is FALSE — routed, not repaired) ·
`D-M257x-116-2` (DENOMINATOR falsified, VOLUME is the answer — and the repair is exonerated by the same
reading) · `D-M257x-116-3` (the induction fences do not scale with repair volume) ·
`D-M257x-116-4` (intra-corpus mis-citation has overtaken platform drift as the largest class)
**Side-deliverables:** none. This was a measuring pass and took no repair, by construction.
**Routes carried forward:**
- **RE-SCOPE — the milestone's next move, and it belongs to the user, not to a tok.** `TOK-07` sealed
  this: after a full enumerate-then-repair cycle, a fourth reading at `P ≥ 15` means the loop's *shape*
  is wrong. The evidence to bring: the repair works, the pool is bigger than the instrument, and the
  induction rate scales with repair volume.
- **`FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`** → net-new, open. G10 dates a claim by
  the FIRST platform ref in the block instead of the ref the claim names. Second guard in two iters
  caught resolving against the wrong thing.
- **`FIX-M257x-iter116-intra-corpus-miscitation-is-the-largest-class`** → net-new, open. 10 of 37 (27 %),
  mechanical, and the one class a machine could close outright.
- **`FIX-M257x-iter116-induction-fences-do-not-scale`** → net-new, open. 5.6 % at +48 lines, 22.0 % at
  +177. The fences fired 4 times on iter-115 and 9 defects still landed.
- `FIX-M257x-iter113-adjudication-is-judgement` → open, **and now carrying a fourth measured miss**: the
  Ant-Academy predicate is live at **five** sites, of which the enumeration reached none.
- `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` → open
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open, delivered-unfixed, 5th measurement: **0**
**Lessons:**
1. **Grade the mechanism and the premise separately — a strategy can be right about HOW and wrong about
   WHETHER.** `TOK-07` prescribed corpus-derived-per-predicate enumeration; it worked (21 of 24 stayed
   closed, `N`/`P` 1.50 → 1.11) and the metric rose anyway. Had this reading graded the strategy as one
   thing, it would have thrown away a working instrument along with a refuted premise — exactly the
   error §9's *grade a refuted strategy LEG BY LEG* was written to prevent, arriving one revision later
   on the strategy that rule was written for.
2. **A repair's induction rate is a function of its VOLUME, and fences that hold at one volume do not
   generalise.** 5.6 % at +48 lines; 22.0 % at +177, with the fences firing four times and still leaking
   9. Any future repair must state its expected induction cost up front, because "the fences are in
   production" was true and insufficient.
3. **When a pool does not drain after the obvious fix, measure what fraction of the residual has NEVER
   been seen.** 25 of 37 here. That single number distinguishes "the repair missed" from "the instrument
   samples", and it is cheap — it needs only the prior readings' predicate lists, which the milestone
   already keeps.
4. **A guard can be RED on the right line for the wrong reason.** `sentinel.md:5` was flagged for a
   compose-count sentence that is true at the ref it names, in a 4-pin block that *does* carry a false
   proposition. Dismissing a false RED without reading the line is how the real defect survives.
