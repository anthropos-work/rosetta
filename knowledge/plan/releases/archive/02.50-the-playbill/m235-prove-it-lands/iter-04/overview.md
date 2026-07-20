---
iteration: iter-04
iteration_type: tik
status: closed-fixed
created: 2026-07-19
---

# iter-04 — fixture-matrix closure (the 4 missing cells)

**Active strategy reference:** TOK-01 (bootstrap, two-track) — Track A step 1 "Fixture matrix closure".
Now unblocked (the scrub is safe as of iter-03; capturing MORE real sessions no longer expands a *leaking*
footprint).

## Cluster / target identified
The gate needs the assessment PASSED set = **2 voice + 1 code + 1 document** and **each sim type present in
passed AND not-passed**. Pre-iter fixture had 9 sessions; the PASSED assessment set was 1 voice + 1 code + 0 doc,
and hiring + interview had only a passed session. Missing 4 cells: **+1 asmt-voice-pass, +1 asmt-doc-pass,
hiring not-passed, interview not-passed**.

## Hypothesis
Sourcing 4 qualifying **public-anchored, non-manager-played, completed** sessions (one per missing cell) via the
`sourcing.go` selection predicates, pinning them, and capturing them through the iter-03-fixed scrub path will
close the matrix. Each pick must carry a real validation fan-out (var+skills+criteria) so the result page renders.

## Phase plan
- Source each cell from prod read-only (postgres MCP, non-PII metadata only): richest-fan-out first, exclude the
  9 existing pins, public predicate + completed + pass/fail per cell.
- Pin the 4 new sessions in `content-sessions.yaml`; strengthen `TestEmbedded` (count 13 + PASSED-set modality
  contract + per-type passed-and-failed).
- Capture all 13 through the fixed path (re-capture 9 + 4 new), counts-only.
- Re-project any manifest consumers (seeder tests) + the full unit gate; verify the 4 new render a fan-out
  (skills/criteria non-zero) and stay leak-clean (capture G-post + offline gate).

## Expected lift
Milestone live gate not moved (Track-A readiness). Provable: fixture = 13 sessions, PASSED assessment set
2 voice / 1 code / 1 document, all 4 types passed AND not-passed; unit gate green; 4 new fixtures leak-clean.

## Escalation conditions
- A cell has NO qualifying public session (esp. interview-fail / hiring-fail) → document the cell as unfillable
  + route forward (don't fake a pin). [Resolved in-iter: all 4 cells had candidates.]

## Acceptable close-no-lift outcomes
If a cell is genuinely unfillable from public data, close with that falsification documented for that cell.
