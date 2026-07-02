---
iter: 05
milestone: M204
iteration_type: tik
status: closed-fixed
created: 2026-07-02
---

# M204 iter-05 — tik: the 5-run reset-to-seed determinism gate

**Type:** tik (measurement / verification — the DETERMINISM half of the exit gate)

**Active strategy reference:** TOK-01 (manager-surface-per-iter) — the verification tik that closes the gate.

**Step 0 re-survey:** metric = 4/4 declared manager UCs functionally green (iter-02→04). The FUNCTIONAL half of
the gate is met; the DETERMINISM half (0 false-fails over 5 consecutive reset-to-seed runs) is untouched. This
iter runs it.

**Cluster / target identified:** the exit gate's second clause — prove the 4 manager Playthroughs are
deterministic under mutation by running them under a REAL `--reset` (reset-to-seed pt-world) for 5 consecutive
runs with 0 false-fails.

**Hypothesis:** the manager UCs are all READ (monitoring) flows over the base-seed org-dashboard data, so a
reset-to-seed re-materializes the identical world each run → 5/5 clean.

**Expected lift:** the gate flips to MET (the last requirement).

**Phase plan:** confirm `stackseed` is reachable (the demo-1 consumption-clone bin) → one `--reset` smoke →
detach a 5-run reset-to-seed loop (per the hard timing rule) → poll → summarize.

**Escalation conditions:** any false-fail across the 5 runs → diagnose (seed-vs-platform drift per P6, or a
flake = a Playthrough defect) → route a fix iter.

**Acceptable close-no-lift outcomes:** if a run false-failed, the falsification (which UC, why) would satisfy
the protocol and route a determinism-fix iter — but that did not happen.

## Outcome — GATE MET
The 5-run reset-to-seed determinism gate: **run1=PASS run2=PASS run3=PASS run4=PASS run5=PASS, 0 false-fails / 5**
(evidence: `5run-gate.log`). Each run: `stackseed --reset` (FK-ordered TRUNCATE) → re-seed `pt-world.seed.yaml`
→ Sentinel reload → all 4 manager Playthroughs green (`4 passed`, ptreport 4× `[PASS]`), ~56-57s/run.

**The full M204 exit gate is MET:** every declared manager-vantage use case (Workforce funnel + roster,
member drill-down via the activity-dashboard, succession/at-risk) has a PASSING Playthrough on a COLD
reset-to-seed demo, with 0 false-fails over 5 consecutive reset runs.

**Environment note (D1):** `stackseed` is not on PATH; the runner calls it bare. Put the demo-1 consumption
clone's bin on PATH (`stack-demo/.../stacks/demo-1/bin`) — the pinned tooling the demo consumes. Recorded so a
future gate run knows the prereq (no code change; the runner correctly calls the pinned `stackseed`).

## Close
See `progress.md`.
