# iter-215 — the honesty line prints 91 undifferentiated misses, and some of them are real ledgers

**Type:** tik — under [`TOK-08`](../decisions.md).

See [`overview.md`](overview.md) for U1–U4 + the stop condition, sealed before any repair.

## What was measured

- **U1 — CONFIRMED.** The partition reconciles exactly to the guard's own declared **91**.
- **U2 — FALSIFIED.** See [`D-M257x-215-2`](decisions.md). **2 of the 57** carry a claim column in
  substance, one of them **44 rows**. The real near-miss set is **36**, not 34.
- **U3 — CONFIRMED.** `iter-82/raw/r15-B.md:109` = `| corpus claim | corpus anchor | … |`, 3 rows.
- **U4 — CONFIRMED and bounded.** `_ANCHOR_COL`'s unbounded `file` matches *"profiles"*; **0 of 68**
  accepted tables affected.

## Close — 2026-08-09

**Outcome:** the claim ledger's honesty line declares **91** unreadable blocker tables and printed 91
identical lines — a disclosure with no signal. Partitioned by `is_ledger_table`'s own vocabulary:
**57 neither / 33 anchor-only / 1 claim-only / 0 collision**, reconciled fail-closed. **At least three
of them are real adjudication tables** dropped on a column spelling — including a **44-row** one — and
the vocabulary widening is measured and **routed**, not smuggled in.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-seventh consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iters 212, 213, 214, 215 = four tiks this run against a cap
of five** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-215-1` … `D-M257x-215-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**90 passed / 0 failed** across `test_claim_twin_guard`, the two answer-key modules,
`test_m257x_claim_twin_mutation_battery` and `test_guard_family` (36 s).
**Stop condition HELD — every derived total unchanged:** 264 claims · 289 forms · 359 rows · 41 files ·
31 short · 64 unquoted · 91 misses · `claim-twin-guard: OK`. This iter changed how misses are REPORTED,
never what is derived.
**RED-proof battery, mtime-mitigated (`§5` r77), restore sha-verified:** collapsing every bucket to one
label takes the exhibit arm and the staged-separation arm RED while the reconciliation arm stays green —
the correct discrimination, since a collapsed partition still sums.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run — the tree was edited during the iter. No Go, no TypeScript.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter215-claim-column-vocabulary-cannot-read-three-live-spellings` — **NEW.**
  `corpus claim`, bare `Claim` (44 rows), `claim`. Widening moves the derived denominator and needs a
  false-RED measurement first (iter-209). Exhibits are pinned by an arm that goes RED when it lands.
- `SURVEY-M257x-iter215-anchor-vocabulary-matches-profiles` — **NEW**, bounded at 0 of 68 accepted.
- `SURVEY-M257x-iter214-route-retractions-are-not-in-the-claim-ledger` — unchanged.
- `SURVEY-M257x-iter212-a-retraction-does-not-reach-the-code-that-acts-on-it` — unchanged.
- `SURVEY-M257x-iter213-a-route-id-is-english` — unchanged. All earlier routes unchanged.

**Lessons:**
- **A mechanical label must not carry a semantic verdict.** *"Deriving 0 from these is correct"* was an
  unmeasured claim wearing a bucket name, and it was wrong for 2 of 57 — one of them 44 rows.
- **A fixture written 170 iters ago is a live control.** `_UNREADABLE_LEDGER` encodes precisely the
  "real ledger, unreadable column" case this iter's first cut declared impossible.
- **Grouping is a claim about the population.** Printing a bucket as a count instead of an itemisation
  asserts that its members do not need reading. That assertion needs the same evidence as any other.
- **Reconcile fail-closed or the itemisation is not the population.**
