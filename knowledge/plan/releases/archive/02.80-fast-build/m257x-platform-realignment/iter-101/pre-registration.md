# iter-101 pre-registration — written BEFORE the read, at corpus `8f04d3a`, platform `0c91421`

**This reading is a near-REPLICATE, and that is the point.** iter-99 measured `N = 28` at corpus `e858fd4`.
iter-100 repaired **8 citation anchors across 5 files** (5 of them in clause-5 scope, +2 net lines) and did
**not** repair the 28 — they are routed as `FIX-M257x-iter99-read-union`, unpaid. So the pool this reading
searches is, to within ~4 anchors, **the pool iter-99 searched**. For the first time the instrument is
being run twice over a fixed subject, which makes reading-to-reading variance measurable rather than
argued — and that is exactly what iter-99 could not do at n=1.

Bands are **tight on purpose**. Two readings running have graded **4 of 9** and **3 of 7**, with every
*mechanism* claim holding and every *magnitude* guess failing. That split is the useful shape and these
bands do not drift toward safe: each is stated so it CAN fail, and each is labelled by what it claims.

| # | prediction | band | kind |
|---|---|---|---|
| 1 | per-reading in-scope upheld BLOCKER count (n₁, n₂) | **[10, 22]** each | magnitude |
| 2 | union `N` | **[18, 34]** | magnitude |
| 3 | **overlap with iter-99's published 28**, matched on PREDICATE not line | **[14, 22]** | mechanism |
| 4 | adjudicator upheld rate | **[74 %, 86 %]** | mechanism |
| 5 | per-pass recall against this reading's own union | **[30 %, 55 %]** | mechanism |
| 6 | **wrong-tree rejections** (the briefing-defect class) | **[1, 5]** | mechanism |
| 7 | wrong-construct intra-corpus citations among upheld in-scope blockers | **≤ 4** | mechanism |
| 8 | platform-drift share of upheld in-scope blockers | **≤ 10 %** | mechanism |
| 9 | per-seat booked spread over the 14 seats (max − min) | **≤ 8** | magnitude |

## What each band is actually risking

**#1 / #2 — the magnitude guesses, and they are NOT centred on "it will fall."** Every magnitude band on
this milestone has failed high (iter-97 #1 and #2, iter-99 #2). The pool is materially unchanged, so the
honest centre is iter-99's own outcome, not below it. `< 18` would mean a 5-site citation repair drained
a third of the residual — implausible, and if it happens the finding is that the residual was concentrated
in anchors, not claims. `> 34` on an unchanged pool would mean the instrument's *recall* rose, which no
mechanism predicts.

**#3 — the new instrument, and the sharpest thing on the sheet.** iter-99's 28 anchors live under
`knowledge/plan/**`, which every seat is hard-barred from reading, so overlap is measured blind. If the two
readings sampled the pool independently at iter-99's measured union recall (≈ 62 %), expected overlap is
`0.62 × 28 ≈ 17`. **Detection is almost certainly not independent** — obvious defects are found by everyone
— so the realistic centre is a little above that, hence [14, 22].

Both tails carry content, and they point opposite ways:

- **Overlap > 22** says detection is strongly correlated across readings. The Chapman estimator assumes the
  two passes are independent; strong correlation makes `N̂ ≈ 45.1` an **over**-estimate, and the "≈17 still
  unfound" figure with it.
- **Overlap < 14** says the readings are closer to independent than assumed, the pool is much larger than
  28, and `N̂` is if anything conservative.

This is the first band on this milestone that can move the *estimator*, not just the count.

**#4 — the one that says iter-99's precision drop was STRUCTURAL, not a one-off.** Four readings held
92.1 / 93.0 / 92.7 / 93.1 %; iter-99 broke it at 78.3 %. This band is set **below the historical one on
purpose**: it predicts the drop persists. It FAILS if the rate returns to [88 %, 96 %] — and that failure
would be the most useful outcome on the sheet, because it would say the break was adjudicator variance at
n=1 rather than a property of the residual or of the briefing.

**#6 — the briefing defect, measured rather than fixed.** `briefing-iter76-AS-RUN.md:37` names
`.agentspace/rosetta-extensions` as "the tooling". Two seats followed it correctly in iter-99 and both were
rejected; the anchors resolve byte-exact in the pinned per-stack clone `ab81527a`. **The instrument is
delivered UNCHANGED, defect included** — that is the only way to measure what it costs, and editing it
would break the comparability this replicate exists to establish. The adjudicator taxonomy gains
`wrong-tree` as a *label* so the class can be separated at grading; no adjudicator is told to expect it.
`0` would mean the class was an iter-99 accident; `> 5` would mean it is a larger tax than the 4-of-10
iter-99 measured.

**#7 — grades iter-100's repair, and is set where it can fail in either direction.** iter-99 upheld **≥ 7**
wrong-construct citations while `anchor_construct_guard` was GREEN. iter-100 closed the blind spot (reach
360 → 528 of 555) but **downgraded its own impact honestly: it re-grades 2 of the 7, not 7**, because five
need the sentence's *claim* and this fence family declines to cross that line. So ≤ 4 is not a victory lap:
it holds only if the mechanical half really was ~half the class. A repeat at 7 fails it; a 0 would mean the
fence reaches further than its author claimed.

**#8 — kept unchanged from iter-97/99, where it held twice.** The platform-drift class M257x was created to
fix is the one class with a complete fence. It went 7/13 → 1/20 → ~1/28. A rise here would mean the two
platform guards have a hole.

**#9 — new, and it tests the partition, not the corpus.** The report must state per-seat spread as a
first-class number. A spread over 8 on a balanced 1431–1506-line partition would say seat-level variance
dominates the measurement, which would undercut every per-reading number above it.

## Binding conditions on the read itself

1. **The instrument is not touched.** Briefing byte-identical, sha `3858ec53…`, `git log --follow` showing
   one commit ever. Re-checked **after** copying, not before. **The known defect at line 37 is delivered
   as-is** and routed, never edited.
2. **Clause 5 is not re-cut, narrowed, or read met any other way.** Met only by a reading that returns
   **zero**. Four user rulings, and this run does not reopen them.
3. **No repair inside the measuring pass.** Anything found routes; nothing is fixed.
4. **Ground truth re-derived, not inherited** — platform clone `== origin HEAD` by `ls-remote`, every other
   checkout's ref restated, guard family run and its verdict recorded before the seats are dealt.
5. **The ref-discipline rejection class stays filtered, not fixed** (17 occurrences, zero contribution to
   any graded count across five readings).
6. **The upheld rate is reported TWICE** — raw, and with the `wrong-tree` class separated — because `N` is
   post-adjudication and immune to the briefing defect while the upheld rate is not. Conflating the two is
   the error this reading exists to avoid.
