# iter-216 — size the routed repair before anyone pays for it

**Type:** tik — under [`TOK-08`](../decisions.md). See [`overview.md`](overview.md) for V1–V5.

## What was measured

**V1–V4 CONFIRMED, exactly as sealed:** 10 candidate tables · 76 rows · **9** matchable · the 44-row
exhibit yields **2**. **V5 CONFIRMED** — four of the ten "claim" columns are the anchor class.

## Close — 2026-08-09

**Outcome:** the vocabulary widening iter-215 routed is worth **9 claims on a 264 denominator (+3.4 %)**,
not the 76 rows its table sizes imply — and the binding constraint turns out to be **quotation, not the
column name**: the candidate cells mostly do not quote the offending sentence, so no fragment floor can
rescue them. iter-215's "44-row ledger" headline is corrected in place: 44 rows, **2** matchable forms.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-eighth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** — **counted, not felt: iters 212, 213, 214, 215, 216 = five tiks this run against
a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-216-1` … `D-M257x-216-4` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**47 passed / 0 failed** across `test_claim_twin_guard`, both answer-key modules and
`test_m257x_claim_twin_mutation_battery` (27 s).
**Derived totals unchanged:** 264 · 289 · 359 · 41 · 31 · 64 · 91 · `OK`.
**RED-proof battery, mtime-mitigated (`§5` r77), both restores sha-verified:** the FIRST control
(`MIN_FRAGMENT_CHARS` 30 → 1) **did NOT fire** — recorded, because the null result is `D-M257x-216-2`;
the control that does fire removes the quote requirement and takes both arms RED at **69 of 76**.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run — the tree was edited during the iter. No Go, no TypeScript.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter215-claim-column-vocabulary-cannot-read-three-live-spellings` — **unchanged, now
  PRICED** at ~9 claims (+3.4 %), with the anchor-class share excluded per `D-M257x-216-4`.
- `SURVEY-M257x-iter216-unquoted-claim-cells-are-a-different-fence` — **NEW.** ~60 of the 76 candidate
  rows carry no quotation at all; reaching them is a different and far less safe instrument, not a
  vocabulary change.
- All routes from iters 207–215 unchanged, plus the standing queue.

**Lessons:**
- **A route sized in containers is not sized.** 10 tables and 76 rows resolve to 9 claims.
- **A control that does NOT fire can be the finding.** The fragment floor was not the constraint;
  discovering that took a failed mutation and was worth more than the arm it was meant to prove.
- **Correct a headline where it was published, and pin the correction with an arm** — prose corrections
  in a ledger have been shown here to travel worse than the figures they correct.
