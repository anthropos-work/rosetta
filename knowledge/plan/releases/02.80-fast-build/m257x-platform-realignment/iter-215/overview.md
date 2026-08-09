---
iter: 215
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-215 — the honesty line prints 91 undifferentiated misses, and one of them is a real ledger

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*
Class under work: **the claim ledger's own declared miss set.**

## Cluster / target identified

iter-214 routed the ledger's verdict sources. Surveying them, `claim_twin_guard` prints on every run:

> `derived 264 claims (289 refuted forms) from 359 blocker rows in 41 ledger file(s); 31 row(s) quoted
> nothing longer than 30 chars, 64 row(s) quoted no refuted form at all, **91 blocker table(s) in a
> shape this derivation cannot read**`

…followed by **91 identical-shaped lines**, one per miss. `unrecognised_blocker_tables` exists precisely
because *"a check that SKIPS reads exactly like one that PASSES"* — the disclosure is real and correct.
**Its problem is not silence, it is signal.** 91 undifferentiated lines is a wall an operator learns to
scroll past, and the day a genuinely-readable ledger lands in an unread shape it is **line 92**.

## Hypothesis

The 91 partition mechanically — by whether the table carries a claim-like column, an anchor-like column,
both, or neither — using `is_ledger_table`'s **own** vocabulary held fixed. Most will be summary tables
that correctly contribute zero claims; the residual is the set worth a human's attention.

## Pre-registered, sealed in this iter's FIRST commit — before any repair

Measured at corpus `b8c6b3b` / rext `d89c81e`, `/usr/bin/python3` 3.9.6, `stack-core`, Python.
Machinery held fixed: `parse_tables`, `_BLOCKER_HEADING`, `is_ledger_table`, `normalize_line`,
`_CLAIM_COL`, `_ANCHOR_COL`.

- **U1** — the partition **reconciles exactly**: **57** neither · **33** anchor-column-only · **1**
  claim-column-only · **0** collision = **91**, the guard's own reported total.
- **U2** — the 57 are summary tables under a `BLOCKERS` heading — `| pass | iter | auditors | blockers |`,
  `| class | n | findings |`, `| # | prediction | outcome |`. They carry no claims, so contributing zero
  is **correct**, not a miss.
- **U3 — THE EXHIBIT.** `iter-82/raw/r15-B.md:109` has header
  `| corpus claim | corpus anchor | what is actually there | true location |` and **3 rows**. Both
  halves of the pair are present *in substance*; `is_ledger_table` returns `None` because `_CLAIM_COL`
  reads `false claim|what is wrong|the claim|issue|^text$` and **"corpus claim" matches none of them**.
  A real adjudication table, dropped on a two-word column spelling, disclosed as one of 91.
- **U4 — a second vocabulary defect, opposite direction, and BOUNDED.** `_ANCHOR_COL` spells `file`
  unbounded, so it matches **"pro*file*s"** — `iter-76/raw/r13-A.md:47`'s `| service | \`profiles:\` key |`
  is booked as carrying an anchor column. Checked on the accept side: **0 of 68 accepted ledger tables**
  have an anchor column that matched `file` only as a substring. The defect is confined to the miss
  partition's display.

**STOP CONDITION, sealed before the repair:** the guard's derived totals must not move —
**264 claims / 289 forms / 359 rows / 41 files / 31 short / 64 unquoted / 91 misses**, and
`claim-twin-guard: OK`. This iter changes how misses are REPORTED, not what is derived. Any change to a
derived total means the repair widened the ledger, which is a different iter with a false-RED
measurement in front of it (iter-209's precondition).

## Expected lift

The declared miss set becomes readable at a glance, the residual worth attention is named, and the two
vocabulary defects are **measured and routed** rather than silently patched.

## Phase plan

A: seal this record. B: partition the miss lines by bucket, reconciled fail-closed. C: arms + staged
control. D: re-run, check every derived total against the stop condition.

## Escalation conditions

Any derived total moves → revert the partition, report, route.

## Acceptable close-no-lift outcomes

U1 failing to reconcile would mean the partition and `is_ledger_table` disagree about the population —
which would be the deliverable, and a bigger one.
