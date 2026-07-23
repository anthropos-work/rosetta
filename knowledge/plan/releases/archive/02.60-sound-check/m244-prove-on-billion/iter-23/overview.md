---
iter: 23
milestone: M244
iteration_type: tik
status: closed
run: 8
created: 2026-07-23
---

# iter-23 — gate (c): the 16 Playthroughs on pt-world (cold reset-to-seed on billion)

**Type:** tik

## Step 0 — Re-survey
Metric = 6/8. Gate (c) stack-verify half done (iter-16 coverage + iter-18 discrete); only the 16 Playthroughs
remain. Re-ordered ahead of BURNIN (TOK-03 sequencing refinement): the "playthroughs LAST" rationale
(seed-destroy vs gate-f-demo-carries + gate-h) is DISCHARGED (all done), and the playthroughs are the
lower-risk deterministic path — lock in the gate-c tick first.

## Active strategy reference
**TOK-03** — final-push: gate (c) 16 Playthroughs → ticks 7/8.

## Cluster / target identified
Run the 16 Playthroughs on pt-world (M219 pattern: `--reset-only` on billion [docker+stackseed+DB], then
the browser specs from the tailnet peer, https). Gate c ticks at 16/16.

## Hypothesis
The 16 Playthroughs land green cold on billion → gate (c) ticks 7/8.

## Expected lift
gate (c) → ticked; metric 6/8 → 7/8.

## Phase plan
`playthroughs.md`: `run-playthroughs.sh 1 --reset-only` (billion) → `PT_HOST=... PT_APP_SCHEME=https
run-playthroughs.sh 1` (peer). Diagnose any failures (tailnet-settle / MIRROR-trap / genuine).

## Escalation conditions
- A genuine Playthrough defect needing a platform edit → surface a blocker.

## Acceptable close-no-lift outcomes
- A documented falsification: a subset fails for a characterized (routed, non-platform) reason.
