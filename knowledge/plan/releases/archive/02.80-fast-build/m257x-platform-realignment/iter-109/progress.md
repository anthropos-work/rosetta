**Type:** tik · `iter_shape: reading` · **`TOK-06` step 4 — the read.** No repair taken inside it.

# iter-109 — the pool does not drain, and the reason is not the one the strategy was built on

## The one-line answer

**`P = 24` predicates · `N = 36` anchors. The pre-registered `P ≥ 15` branch fired: THE POOL DOES NOT
DRAIN.** The secondary `N ≥ 20` branch fired with it — both metrics agree, so there is no `P`/`N` split to
adjudicate.

The rule was sealed in `ac48e5b` before the first seat was dealt and is graded exactly as written.

Full graded sheet: [`adjudication.md`](adjudication.md).

| | iter-101 | iter-103 | **iter-109** |
|---|---|---|---|
| distinct false **predicates** | 22 | 22 | **24** |
| distinct **anchors** | 24 | 33 | **36** |

## The question it was built to answer, answered

`TOK-06` was authored on iter-103's decomposition — clone advance **61 %** of `N`, induction **21 %** —
and the sentence *"inflow is comparable to outflow."* The pre-registration named the confound that
measurement could not resolve: was the 61 % **ARRIVAL** (drift that appeared because five clones moved) or
**DETECTION** (standing drift a reading finds once the subtle half is repaired)?

**All 14 platform clones were frozen at the identical sha for this reading. The answer is DETECTION.**

> **Platform-drift is still ~33 % of the upheld residual over a subject in which nothing moved.** Band #8
> predicted `≤ 25 %` precisely because a frozen subject cannot arrive new drift. It failed, and the failure
> *is* the finding: the drift was already in the corpus. **It was never an inflow.**

## The structural finding — a repair scoped to a reading's DETECTIONS cannot close a predicate

Band #3 held at **2**, and its meaning inverts. iter-108's `46/46 = 100 %` reach grade is correct — and its
anchor list was derived from `iter-103/raw/`, i.e. from **what the previous reading detected**.

**A predicate's site list and a reading's detection list are different sets.** Measured here:

- iter-108 repaired `external_services.md:565`; the same proposition sits **eleven lines above at `:554`**,
  unrepaired, in the same file.
- iter-108 repaired `ai_architecture.md:95`/`:99`; the twin at `:34` still cites `markdownManager.go:19`
  while the repaired `external_services.md:560` now says in terms that `:19` is a **doc-comment, not code**.
  **The repair created a self-contradiction by fixing one side of a pair.**

**The induction fences worked and induction is not the problem**: band #10 measured **2 of 36** anchors in
prose iters 104–108 wrote — **21 % → 5.6 %**, the lowest in the series, against a rate that had held ~2 per
cycle for six cycles. Step 2 did its job. Fences cannot see a predicate the repair never visited.

## Bands: 7 HELD of 13

HELD `#3` `#3b` `#4` `#6` `#7` `#9` `#10` · FAILED `#1` `#2` `#2p` `#5` `#8` `#11`.

Upheld rate **91.4 % raw / 91.4 % `wrong-tree`-separated** — the two coincide for the first time because
**`wrong-tree` was zero**. Band #6's series: **4 → 1 → 1 → 0**.

## Close — 2026-08-06

**Outcome:** `P = 24` / `N = 36` against a rule sealed before the first seat. **The pool does not drain**,
and the reading separates why: iter-103's 61 % drift was **standing, not arriving** — proven by holding all
14 clones frozen and still measuring ~33 % drift. `TOK-06`'s two inflow fences are vindicated where they
apply (induction **21 % → 5.6 %**, band #10 held at 2) but were **not the binding constraint**. The binding
constraint is that a repair scoped to a prior reading's *detections* cannot close a predicate at sites that
reading never saw — measured directly as two surviving twins, one of which the repair turned into a
self-contradiction. 14/14 seats, zero lost, all committed verbatim before adjudication; 35 booked → 32
upheld / 3 rejected.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7
**Decisions:** `D-M257x-109-1` (a seat-commit subject named one seat and carried two — recorded, not
rewritten) · `D-M257x-109-2` (the HEAD moved mid-open; read scope proven identical with a firing negative
control; corrected in the open) · `D-M257x-109-3` (ARRIVAL vs DETECTION resolved as DETECTION) ·
`D-M257x-109-4` (repair scope is detection-bounded — the structural finding)
**Side-deliverables:** none. This was a measuring pass and took no repair, by construction.
**Routes carried forward:**
- **`FIX-M257x-iter109-read-union`** → next iter. 24 predicates / 36 anchors, **with a binding change of
  scope: the anchor set is re-derived FROM THE CORPUS per predicate, never from `iter-109/raw/`.**
- **`FIX-M257x-iter109-repair-scope-is-detection-bounded`** → net-new, **outranks the union**. Needs a
  per-predicate corpus-wide site sweep before a repair may be called done.
- **`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`** → stays open, **re-ranked down**: drift is
  standing, not arriving, so a drift fence is not the lever it was ranked as.
- **`DEF-M257x-iter101-briefing-rext-tree`** → open, delivered-unfixed, 4th measurement: **0**.
- **`FIX-M257x-iter108-stackcore-suite-hangs`** → open; no full-suite total exists on this host and none is
  quoted.
**Lessons:**
1. **A measurement of composition is not a measurement of flow.** iter-103's 61 % was honest and correct
   about *what the residual was made of*, and was read as *where the residual came from*. Freezing the
   subject was the only thing that could tell those apart — and it cost a full strategy cycle to learn.
   **Before fencing an inflow, prove it is flowing.**
2. **Repair by predicate, but scope by corpus.** TOK-05's unit was right and is not in question. The
   *ledger* was detection-bounded, and detection recall on this instrument runs 33–83 %, so ~100 % reach
   against that ledger is compatible with leaving twins standing.
3. **Fixing one site of a pair is worse than fixing neither.** It manufactures a self-contradiction where
   there was a single consistent falsehood. Any repair must sweep the predicate corpus-wide or not run.
4. **The same claim rejected twice, two readings apart, by different adjudicators, is a result.**
   `shared_libraries.md:128` came back `wrong-tree` at iter-103 and `ref-discipline` here — third
   independent confirmation the claim is TRUE, and the reason iter-108 was right to leave it alone.
