**Type:** tik (measurement/verification — the determinism half of the exit gate). Protocol:
`corpus/ops/demo/playthroughs.md` § iteration protocol + the exit gate.

# iter-05 — the 5-run reset-to-seed determinism gate

Confirmed `stackseed` reachable (demo-1 consumption-clone bin on PATH), ran one `--reset` smoke (all 4 manager
UCs green, ~61s), then detached a 5-run reset-to-seed loop of the manager suite and polled to completion.

## Close — 2026-07-02

**Outcome:** the 5-run reset-to-seed determinism gate is MET — run1..run5 all PASS, **0 false-fails / 5**, all 4
manager UCs green each run (each run: reset-to-seed pt-world → Sentinel reload → 4 passed). Evidence:
`5run-gate.log`. **The full M204 exit gate is MET.**
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks this session) — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (iter-05/decisions.md — `stackseed` PATH prereq for the gate run; no code change).
**Side-deliverables:** none (a pure verification iter — 0 production/harness code modified).
**Routes carried forward:** none — the gate is met. Next: `/developer-kit:harden-mstone-iters` (final pass on
unhardened iters), then `/developer-kit:close-milestone`.
**Lessons:** the manager vantage is fully deterministic under reset-to-seed because every manager UC is a READ
(monitoring) flow over base-seed org-dashboard data — a reset re-materializes the identical world, so 5/5 clean
with no per-UC determinism work. The `stackseed`-on-PATH prereq (the demo-N consumption-clone bin) is a gate-run
environment note, recorded for the next runner.
