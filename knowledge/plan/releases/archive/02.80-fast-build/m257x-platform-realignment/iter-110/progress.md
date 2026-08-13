# iter-110 — TOK-07 authored

**Type:** tok (**deliberate**) — per `platform-alignment.md` §9 and the `TOK-06` precedent for the
deliberate, non-terminating flavor.

## What was done

Authored `TOK-07: enumerate the predicate, not the anchor` into the milestone-root `decisions.md`.

Nothing else. **No corpus repair, no fence authored, no reading taken, no clone fetched (§5 rule 41a
held), no stack touched, no tag cut.** `N` was not re-measured and **no `N` movement is claimed** — the
§9 guard-rail wording, so this iter and an iter that measured zero cannot look alike.

## The derivation, in the order it was actually done

1. **Checked the streak before writing anything.** Last three tiks: iter-107 (no reading), iter-108 (no
   reading), iter-109 (the read). Under §9's refinement, **one** of three measured — the trigger's
   precondition is unestablished and Phase 0 rule 2 falls through to tik. This tok is deliberate.
2. **Re-read `TOK-06` against iter-109's sheet** to test whether a revision was needed at all, or whether
   iter-109's finding was already covered. It is not covered: `TOK-06` puts the read **last** (correct,
   kept) but **says nothing about the repair's denominator**, which is the thing iter-109 measured wrong.
3. **Verified the multiplier series from the ledger rather than quoting it.** This mattered — the numbers
   are the strategy's whole argument:

   | pass | claimed | verified at |
   |---|---|---|
   | iter-96 | 13 anchors → **51 sites** across 23 files, **38** an anchor-wise repair would have left | `iter-96/progress.md:1`, `:11`, `:13`, `:14` |
   | iter-98 | 20 anchors / 21 predicates → **37 sites** across 22 files | milestone `progress.md:1373-1374` |
   | iter-102 | 52 anchors → **98 sites** found → 94 repaired | milestone `progress.md:1500-1504` |
   | iter-108 | 31 primary anchors derived from `iter-103/raw/`; reach **46/46** | milestone `progress.md:1817-1828` |

   **iter-108 reports no site-expansion figure at all** — not a low one. That is the observation the
   strategy rests on, and it is stronger than "the multiplier fell to 1.0×": there was no expansion step
   to report a multiplier for.
4. **Read the 24 predicates off `iter-109/verdicts/*.md`** (the `PREDICATE:` field each adjudicator
   booked) to confirm they are stated as propositions and are therefore enumerable — they are; the field
   was designed for exactly this in `TOK-05`.
5. **Wrote the entry**, including a pre-registered falsification of the strategy itself.

## Deliverables

- `TOK-07` in the milestone-root `decisions.md` — refutation + what is kept + the measured multiplier
  collapse + four rules + a four-step order + a sealed falsification + `Strategy class:
  retry-with-evidence`.
- `iter-110/decisions.md` — `D-M257x-110-1` … `D-M257x-110-3`.

## Close — 2026-08-06

**Outcome:** `TOK-07` authored — **the repair's denominator moves from a prior reading's detections to the
corpus, per predicate.** iter-109's refutation of `TOK-06`'s premise is recorded, its induction fences are
**kept and re-ranked rather than reverted**, and the binding constraint is stated as a verified multiplier
series (iter-96 **3.92×** → iter-108 **no expansion step at all**) rather than as an atmosphere. Strategy
class `retry-with-evidence`, because the method being restored is iter-96's own.
**Type:** tok
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **n — this is a DELIBERATE tok, author-initiated and non-terminating; the 3-no-prog trigger did not fire and could not, since only one of the last three tiks measured** — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (0 tiks this session) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** D-M257x-110-1, D-M257x-110-2, D-M257x-110-3
**Side-deliverables:** none
**Routes carried forward:**
- `FIX-M257x-harden23-json-polluted-by-provenance-stamp` → **iter-111** (step 0)
- `FIX-M257x-iter108-stackcore-suite-hangs` → **iter-111** (step 0)
- `FIX-M257x-iter109-repair-scope-is-detection-bounded` → step 1, the enumerator
- `FIX-M257x-iter109-read-union` (24 predicates) → step 2, against the **enumerated** set
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → **de-ranked**, stays open
- `DEF-M257x-iter101-briefing-rext-tree` → stays open, fourth delivered-unfixed measurement
**Lessons:**
- **A refuted premise does not condemn the instruments built under it.** `TOK-06` was authored on a
  measurement that turned out to describe composition rather than flow — and its induction fences still
  took repair-induction from 21 % to 5.6 %. The reflex on a refuted strategy is to revert its work; the
  correct move here was to **re-rank** it. Generalised into §9 at this iter's close.
- **A reach metric is a measurement, so it is settled by its denominator's provenance** — the same
  sentence `fence_provenance` exists for, one layer up. `100 % of the wrong set` is the milestone's
  signature defect, and it arrived in the check that was supposed to catch it.
