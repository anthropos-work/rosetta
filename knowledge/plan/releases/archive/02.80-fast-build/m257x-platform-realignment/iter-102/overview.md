# iter-102 — the REPAIR of two unions at once, 52 anchors, by PREDICATE

**Shape:** tik · **REPAIR pass** · `iter_shape: repair` · **no reading is taken inside it**

## What this iter is

The paydown of **both** outstanding read unions in one pass:

- `FIX-M257x-iter99-read-union` — **28** anchors, routed at iter-99, unpaid because iter-100 deliberately
  withheld payment so iter-101's replicate would run on a **fixed subject**
- `FIX-M257x-iter101-read-union` — **24** anchors, routed at iter-101

**52 anchors.** That reason for withholding has expired: the replicate is done and its verdict is in.

## Why it is bigger than the two previous repairs, and what that changes

iter-101's band #3 is the governing result. Blind overlap with iter-99's published 28, matched on predicate,
came out **6** against a pre-registered band of **[14, 22]** — a fail LOW. The pre-registration
pre-committed to what that means: the two readings are **more independent than Chapman assumes**, so
`N̂ = 45.1` is a **FLOOR**, and cross-reading Chapman over iter-99 × iter-101 gives **N̂ ≈ 102.6** — itself a
floor, because heterogeneous catchability biases Chapman downward.

**Read that correctly, and the record must not drift:** the pool was probably always ~100. **It is not
growing — the estimator was wrong, and the replicate fixed it.** The rising series
16.7 → 29.4 → 45.2 → ~103 is four successive **corrections to an underestimate**, not four measurements of
growth.

**Operational consequence, and it is the reason this pass is wide rather than deep:** repairing a reading's
union has never been the drain the milestone modelled it as, and the residual is ~4× what was thought.
**Repair is the parallelizable half of this loop; the measuring pass is the half that must stay serial.**
Seven disjoint seats has been the habit. This pass goes wider.

## Binding conditions (inherited from iter-76, re-stated because they bind)

1. **Repair by PREDICATE, not by anchor.** Measured multipliers: iter-96 was 13 anchors → **51** sites (38
   an anchor-wise pass would have missed); iter-98 was 20 → **37**. Booked predicate width has run **2**
   against live widths of **7** and **5** — repair has been closing roughly **a third** of what it books.
2. **Expand on BOTH axes — twin AND paraphrase.** `claim_twin_guard` matches quoted verbatim forms;
   iter-97 measured what got through as **3 of 51, all paraphrases**.
3. **Grade reach and REPORT it.** Publish the ledger in the shape `claim_ledger.py` derives from, so
   `claim_twin_guard` + `repair_postcondition` **fence** completeness rather than the repairer claiming it.
4. **Adjudicate before repairing.** Always. Five clones moved under the previous reading (see
   `ground-truth.md`); a finding graded at a superseded ref is re-derived, not inherited.
5. **TRAP A holds.** Where the underlying fact was **deleted** rather than moved, **restate or drop — never
   re-anchor.** A correctly-cited false statement is worse than a stale one.
6. **A measuring pass may not contain a repair, and a repair pass may not contain a reading.**

## Also in this iter — two defects, both between-readings work

- **`FIX-M257x-iter101-app-clone-unfetched`** — 17 sites label `2035f9a` as `origin/main`; origin/main is
  `ad9f3c49`. Re-derive and repair every one, and **measure how much residual this single repo move
  injected**, because that number is load-bearing for the milestone's ETA.
- **`DEF-M257x-iter101-crosslane-fetch`** — a coordination defect: path ownership assigned `stack-demo/**`
  to one lane while another lane's adjudicators were grading claims against clones *inside* it. The durable
  rule goes into `corpus/ops/platform-alignment.md`: **a reading's ground truth includes the clone refs, so
  no lane may fetch while a reading is in flight.**

## What this iter does NOT do

- **No reading.** No `N`. No estimator. The gate does not move on clause 5 in this pass and cannot.
- **No stack.** Lane B owns `stack-demo/**`. Nothing here brings one up, and **nothing here fetches**.
- **No second rext tag.** `fast-build-m257x-iter-101` is cut and on origin at `0011c10a`.
- **No decision on `DEF-M257x-iter80-storage-prod-bucket`.** Still the user's, still open.
