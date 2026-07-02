---
iteration: 06
iteration_type: tik
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-06 — TIK (the 5-run reset-to-seed determinism gate — GATE-COMPLETING)

## Active strategy reference
**TOK-01** — the gate's second half (determinism), after all 3 coverage journeys landed (iter-02..05).

## Cluster / target identified
Re-survey (Phase 1 Step 0): the COVERAGE half of the gate is met (6/6 employee Playthroughs pass — Profile
identity+verified+growth+timeline, Skill Paths legacy, AI-sim chat launch). The one REMAINING gate criterion
is **0 false-fails over 5 consecutive reset-to-seed runs** (the determinism proof). iter-04/05 routed it to
"milestone-close, from the consumption clone" — but re-survey shows the authoring copy CAN build `stackseed` +
`datadna` from `stack-seeding/cmd/`, so the gate is runnable NOW from the authoring copy (a demo-1 consumption
`stackseed` binary also exists but is pinned to v1.10.1, pre-M203 — the authoring build is the correct one).
Substituted the target IN (not deferred): run the determinism gate this iter.

## Hypothesis
The full reset→reseed→Sentinel-reload→play cycle is deterministic: 5 consecutive cold reset runs each produce
6/6 passing Playthroughs, 0 false-fails.

## Expected lift → realized
The gate's determinism criterion MET. Realized: single-cycle validation 18/18 pass (4.5m); then the 5-run
gate = **5/5 passed, 0 failed** (6/6 Playthroughs per run, ~4.1–4.7m each).

## Phase plan (as executed)
Build stackseed+datadna from the authoring copy → 1 validation reset cycle (full suite) → the 5-run
determinism gate (--reset --grep @pt, serial) → confirm 0 false-fails.

## Close → see progress.md — GATE MET.
