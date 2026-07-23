---
iter: 16
milestone: M244
iteration_type: tik
status: closed-fixed
created: 2026-07-23
---

# iter-16 — gate (c): recalibrate the stale academy-home coverage marker → coverage sweeps green

**Type:** tik · **Active strategy:** TOK-02 (sweep the remaining gates live)

## Active strategy reference
TOK-02. iter-15's finding: the gate-(c) coverage cross-port's residual (after the gate-(d) fix) is a STALE MARKER, not a defect.

## Cluster / target identified
`ANT_ACADEMY_HOME_SECTION` (`coverage-manifest.ts:712`) required the literal `AI Academy` in the browser's `main/body` innerText, but the redesigned academy home's visible heading is "Academy" (the literal survives only in the non-visible `<title>`). So the coverage cross-port false-FAILED a home rendering 483 real cards (iter-15).

## Hypothesis
Recalibrating the marker to the M42e cardinality shape (card-count + kept 400 floor) passes the real render AND strengthens the anti-keyless/empty-catalog teeth → both hero coverage sweeps go green.

## Expected lift
The coverage half of gate (c) green on billion. (Gate (c) as a whole also needs the discrete stack-verify specs + 16 Playthroughs, so the gate part does not tick yet.)

## Phase plan
Edit the descriptor (text→both+minCount) + update its unit test → typecheck + unit → re-run both coverage sweeps live → close.

## Escalation conditions
A coverage failure requiring a platform edit → STOP. (None.)

## Acceptable close-no-lift outcomes
If the recalibrated marker still failed the real render, close-no-lift with the falsification.
