---
iter: 216
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-216 — size the routed repair before anyone pays for it

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## Cluster / target identified

iter-215 routed `SURVEY-M257x-iter215-claim-column-vocabulary-cannot-read-three-live-spellings` and
published, as its milestone-ledger headline, *"a 44-row real ledger was hiding in them."* True about the
**table**. **Unmeasured about the consequence** — a ledger row only becomes a claim if its claim cell
carries a quoted span at or above `MIN_FRAGMENT_CHARS`. `TOK-08` says report the enumerated population
**and state the denominator**; a route sized in tables is not sized.

## Pre-registered, sealed in this iter's FIRST commit

Measured at corpus `0d8eda9` / rext `f933ccb`, machinery held fixed (`parse_tables`,
`_BLOCKER_HEADING`, `is_ledger_table`, `normalize_line`, `extract_forms`, `MIN_FRAGMENT_CHARS`).

- **V1** — candidate tables (a human-readable claim column that `_CLAIM_COL` cannot read): **10**.
- **V2** — rows in them: **76**.
- **V3** — rows that would yield a **matchable refuted form** if `_CLAIM_COL` were widened: **9**.
- **V4** — the 44-row exhibit (`iter-48/raw/D.md:72`) contributes **2** of those 9. iter-215's headline
  emphasis is therefore **corrected, not retracted**: the table is real and its yield is 2 claims.
- **V5** — several putative claim columns are the **anchor class** by design (`claim's anchor`,
  `corpus line`, `doc says`, `what is actually there`) — the wrong-construct class iter-42 assigned to
  a **symbol-aware anchor check**, a different instrument. Widening `_CLAIM_COL` to reach them would
  pull this fence into territory its own docstring rules out.

**STOP CONDITION:** this iter lands a measurement and an arm. If landing the arm changes any derived
total, revert and report.

## Expected lift

The routed repair carries a size, so the next iter that picks it up knows it is worth **≈ 9 claims on a
264 denominator (+3.4 %)**, not 76 — and knows that part of the residual belongs to another instrument.

## Phase plan

A: seal. B: land the sizing as a running arm. C: append the correction to iter-215's ledger line.
D: re-run.

## Acceptable close-no-lift outcomes

V3 coming back at or near 76 would invert the conclusion and make the widening urgent — which would be
the deliverable.
