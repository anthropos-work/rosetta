# iter-99 pre-registration — written BEFORE the read, at corpus `<iter-98 commit>`, platform `0c91421`

Bands are **tight on purpose**. iter-95 graded **6 of 6** and booked it as a *warning, not a win*; iter-97
narrowed until they could fail and graded **3 of 7**, and the four failures carried the content. **A 7-of-7
here is a warning, not a win.** Every band below is stated so that it CAN fail, and each is labelled by what
it is a claim about — a **magnitude** guess or a **mechanism** claim. iter-97's lesson was that the
mechanism model is good and the magnitude model was badly optimistic; these bands are set to test whether
that is still true after a repair that changed the pool's *shape* rather than only its size.

| # | prediction | band | kind |
|---|---|---|---|
| 1 | per-reading in-scope upheld BLOCKER count | **[8, 18]** each | magnitude |
| 2 | union `N` | **[12, 26]** | magnitude |
| 3 | recurrence of the 21 iter-98 predicates | **≤ 2** | mechanism |
| 4 | adjudicator upheld rate | **[88 %, 96 %]** | mechanism |
| 5 | per-pass recall | **[28 %, 55 %]**, and **at least one pass below iter-97's 41 %** | mechanism |
| 6 | platform-drift share of upheld in-scope blockers | **≤ 10 %** | mechanism |
| 7 | repair-induced upheld in-scope blockers | **[0, 3]** | mechanism |
| 8 | **mean sites-per-predicate of whatever is booked**, measured the same way as `discovery-pool.md` §1 | **< 2.5** | mechanism |
| 9 | wrong-construct intra-corpus citations among upheld blockers | **≤ 1** | mechanism |

## What each band is actually risking

**#1 / #2 — the magnitude guesses, and they are deliberately NOT centred on "the repair worked."**
iter-97's #1 `[2,7]` and #2 `[3,9]` both failed high, on the assumption that repairing 51 sites reduces `N`.
That assumption is now explicitly retired ([`iter-98/discovery-pool.md`](../iter-98/discovery-pool.md) §3).
These bands are centred near iter-97's outcome rather than below it, and they are **wide enough to admit a
rise** — but not unbounded: `> 26` would mean the repair released more than it closed, and `< 12` would mean
the pool really is draining fast. Both are informative; neither is assumed.

**#5 — the sharpest band, and the one most likely to fail.** §3 of `discovery-pool.md` predicts that a pool
shifting from wide to narrow predicates should **depress** per-pass recall, because a predicate at 11 sites
gives 14 blind seats eleven chances and one at a single site gives one. iter-98 measured mean width falling
**3.64 → 1.76**. So this band does not merely allow lower recall, it **requires at least one pass below
iter-97's 41 %**. If both passes come in *above* 41 %, the mechanism argued in §3 is wrong and the honest
reading flips toward "the pool is simply draining."

**#8 — the new instrument, on trial.** `discovery-pool.md` argues width is a better convergence signal than
`N` because it does not depend on recall. That is only useful if it is *stable*. If iter-99 books predicates
averaging ≥ 2.5 sites, the width collapse was a property of iter-98's particular input rather than of the
corpus, and §1's whole argument weakens.

**#3 / #9 — teeth against my own repair.** iter-97's #3 (*zero recurrence*) failed at **3**, and it was the
most useful prediction on the sheet. #3 is set at **≤ 2**, tighter than the outcome it is calibrated
against, because iter-98 added the paraphrase axis specifically to close that gap — if paraphrase expansion
did nothing, this fails. #9 is tighter still: `anchor_construct_guard` now scans the entire anchor set and
was GREEN at commit, so an upheld wrong-construct citation means **the guard has a blind spot**, which is a
finding about the instrument rather than the corpus.

**#7 — kept from iter-97, where it held at exactly 2.** iter-98 caught **5** induced citation moves inside
the iter (iter-96 caught 0 and shipped 2). If in-iter re-derivation works, this should land at the low end;
`> 3` would say it does not.

## Binding conditions on the read itself

1. **The instrument is not touched.** Briefing byte-identical, sha `3858ec53…`, `git log --follow` showing
   one commit ever. Re-check the sha **after** copying, not before.
2. **Clause 5 is not re-cut, narrowed, or read met any other way.** Met only by a reading that returns
   **zero**. Four user rulings.
3. **No repair inside the measuring pass.** Anything found routes; nothing is fixed.
4. **Ground truth re-derived, not inherited** — platform clone `== origin HEAD` by `ls-remote`, every other
   checkout's ref restated, guard family run and its verdict recorded before the seats are dealt.
5. **The ref-discipline rejection class stays filtered, not fixed** (13 occurrences, zero contribution to
   any graded count across four readings).
