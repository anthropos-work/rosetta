**Type:** tik

# iter-206 — deriving the comment figures, starting where this run has been standing

iter-205 sized the comment site-kind at **96 standing figures** and closed at the cap with one confirmed
stale instance unrepaired. Six of those sit in `claim_census_guard.py` and describe populations that
module computes — the smallest subset where *derivation* is available rather than judgement, in the file
this whole run has been editing.

**Four of the six were stale, and this run made two of them stale itself** (`D-M257x-206-1`):

| comment | said | derives |
|---|---|---|
| the ABBREV population *"in these 40 files"* | 40 | **41** |
| *"unresolved for a class of 2"* | 2 | **1** |
| *"Measured over the corpus: 2 pairs"* | 2 | **1** |
| *"a wholesale warning over 949 pairs"* | 949 | **1,015** |

The two `2`s are **iter-202's own number**. That iter measured the wrong-repo class at 2, wrote it into
two comments, and then — four paragraphs further down the same iter — fixed the extension truncation that
made the second member exist: `ant-academy.md`'s `code/public/catalog.js` was a **parser artifact**, the
file being `catalog.json`. The class was **1** before iter-202 closed, and its own comments still said 2
four iters later. **A number can go stale inside the iter that writes it**, and its author is the least
likely person to re-read the sentence.

## `949` is not re-pinned — the sentence stops carrying a number

`949 → 1,015` buys one iter of correctness and re-arms the trap. The sentence now points at
**`_exp["under_clones"]`**, which the code prints two lines below, and states that it read `949` from
iter-198 until iter-202 moved it with **nothing noticing for four iters** (`D-M257x-206-2`). *A number a
comment does not carry cannot go stale* — iter-199's reasoning about printed totals, applied where the
class began.

The fence recomputes each figure and, for the `949`, asserts the **absence** of the literal and the
presence of the pointer — plus an arm requiring the three derived values (**41**, **1**, **1,015**) to be
pairwise distinct, so a comment carrying the wrong one cannot satisfy the wrong arm (`D-M257x-206-3`).

## Close — 2026-08-09

**Outcome:** the standing class has a **verified subset for the first time — 7 of 168**, and the first
subset returned **4 stale of 6**. Two of those four were written by this run four iters earlier and were
already wrong when their own iter closed. Repaired: three corrected in place and fenced against the live
census, one **de-literalised** so it points at the value the code prints instead of carrying a copy. The
4-in-6 rate is explicitly **not** an estimate for the remaining 161 — the subset was selected for being
derivable, and derivable figures cluster in modules under active edit, which is where staleness
concentrates.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-eighth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** — **counted, not felt**: iters 202, 203, 204, 205, 206 = **five** tiks this run
against a cap of five. (iter-205's first draft graded the cap at four and was corrected before commit;
this one is the fifth and the arithmetic is written out so it can be checked rather than trusted) —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-206-1` … `D-M257x-206-4` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **34 passed** in
`test_claim_census_substrate_m257x.py` (the changed fence, +4 arms) and **188 passed** across
`test_frozen_expectation_census_m257x.py` + `test_claim_census_guard.py` + `test_test_collection_fence.py`
+ `test_guard_family.py` + `test_suite_census_collection.py`. Green under **both** runners
(unittest 3.9.6: `Ran 34 … OK`). `claim_census_guard --check` green (**1,130** unevidenced, baseline
1,164, over 41 files).
*Scope: `stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no TypeScript, and the
other ten rext sections were not run.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter205-comment-provenance-notes-are-the-highest-risk-standing-figures` — **CLOSED for
  its named instance and its named module.** The `949` is de-literalised; the other five figures in that
  module are derived and fenced.
- `SURVEY-M257x-iter205-the-standing-buckets-total-168-and-one-is-derived` — **RE-STATED: 168 sized, 7
  derived, 161 unverified.** First numerator the class has had.
- `SURVEY-M257x-iter206-the-other-161-standing-figures-are-in-modules-with-no-fence` — **NEW.** 21 of
  them are in `suite_census.py`, 12 in `derivation_registry.py`, 9 each in `anchor_construct_guard.py`
  and `guard_family.py`. Each needs its own module's population recomputed, which is per-module work, and
  a sampled estimate would be the sampling `TOK-08` replaced.
- `SURVEY-M257x-iter206-a-figure-can-be-stale-before-its-own-iter-closes` — **NEW.** Two comments were
  wrong at the moment their iter committed, because a later change *in the same iter* moved the value.
  No fence in this repo checks an iter's own numbers against the tree it leaves behind, and the close
  section is written before the last commit.
- Unchanged and still open: `SURVEY-M257x-iter203-the-standing-class-is-not-mechanically-decidable` ·
  `SURVEY-M257x-iter202-published-citation-figures-predate-the-truncation-fix` ·
  `SURVEY-M257x-iter202-anchor-subject-census-extension-vocabulary-is-narrower-than-the-census` ·
  `SURVEY-M257x-iter202-the-eighteen-false-RED-pairs-remain-substrate-dependent` ·
  `SURVEY-M257x-iter201-published-suite-totals-predate-the-runner-gap-closing` ·
  `SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
  `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` ·
  `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **A number can be stale before its own iter closes.** iter-202's `2` was wrong by the time iter-202
  committed, because iter-202 itself removed the second member. Sizing a class early in an iter and
  changing the population later in it is a normal shape, and nothing checks the two against each other.
- **De-literalise beats re-pin.** Correcting `949` to `1,015` would have re-armed the trap; pointing at
  the printed value removes it.
- **State the selection bias with the hit rate.** 4 of 6 is not an estimate of the other 161 — the subset
  was chosen for being derivable, and derivable figures live in the modules under active edit.
