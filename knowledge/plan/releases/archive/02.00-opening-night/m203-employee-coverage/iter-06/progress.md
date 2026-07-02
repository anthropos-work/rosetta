**Type:** tik (under TOK-01, the determinism half) — protocol: `corpus/ops/demo/playthroughs.md`
§"reset-to-seed + the serial-default runner" + §"The iteration protocol".

# iter-06 progress

- **Re-survey:** coverage half met (6/6 employee Playthroughs pass). Remaining gate criterion = 0 false-fails
  over 5 consecutive reset runs. Found the authoring copy CAN build `stackseed` + `datadna` from
  `stack-seeding/cmd/` → ran the gate NOW (the demo-1 consumption stackseed is pinned v1.10.1, pre-M203; the
  authoring build carries the pt-world seed + the M203 runner Reload fix — the correct binary).
- **Built** stackseed + datadna from the authoring copy onto a scratch PATH.
- **Validation cycle (1× --reset, full suite):** reset → reseed (48121 rows, isolation CLEAN: no shared/prod
  writes) → **Sentinel reload OK** → **18/18 passed (4.5m)** — the full reset cycle works end-to-end (and the
  Sentinel-Reload fix works after a REAL reset, not just a manual reload).
- **The 5-run determinism gate** (`run-playthroughs.sh 1 --reset --grep @pt`, serial, 5×):
  - RUN 1: PASS (6 passed, 4.1m)
  - RUN 2: PASS (6 passed, 4.6m)
  - RUN 3: PASS (6 passed, 4.3m)
  - RUN 4: PASS (6 passed, 4.7m)
  - RUN 5: PASS (6 passed, 4.5m)
  - **GATE: 5/5 passed, 0 failed** — 0 false-fails over 5 consecutive cold reset-to-seed runs.

## Close — 2026-07-02

**Outcome:** The M203 exit gate is FULLY MET — every declared employee-vantage use case (6/6) has a PASSING
Playthrough on a COLD reset-to-seed demo, with 0 false-fails over 5 consecutive reset runs.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (run the determinism gate from the authoring-built stackseed, not deferred) — see decisions.md
**Side-deliverables:** none
**Routes carried forward (non-gate, for harden/close or a later milestone):**
  - ai-simulations.code.UC1 (Judge0 path) + .interview.UC1 (text) — non-gate M201 extras. Handler PT-M203-nongate-sims.
  - profile.self-evaluation.UC1 — non-gate M201 extra (rate-modal click-intercept). Handler PT-M203-selfeval.
  - the Skill-Paths verify-skill TERMINAL (composes with a NON-voice assessment sim) — proven on the profile
    side today; the composed end-to-end chain is a later deepening. Handler PT-M203-skillpath-verify-terminal.
**Lessons:** the determinism gate was runnable from the authoring copy all along (build stackseed+datadna from
`stack-seeding/cmd/`) — no need to wait for the consumption-clone tag. The reset→reseed→Sentinel-reload→play
cycle is stable + deterministic (5/5, 0 false-fails). The gate is MET → next: harden + close-milestone (which
tags `opening-night-m203`).
