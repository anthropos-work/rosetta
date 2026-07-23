# iter-26 — progress

**Type:** tik (run 9, tik 2) — characterize the 3 remaining ai-readiness sub-render gaps (post launched_by fix).

## What happened
Root-caused all 3 remaining ai-readiness Playthrough sub-failures against billion's actual v1.341.0 read paths
+ the live DB. All are DISTINCT DATA/seed/wiring gaps (NOT a stale frontend — the interview panel renders a
heading; NOT the launched_by zero-state — fixed iter-25). Each is seed-fixable in rext stack-seeding, 0
platform edits. Precise per-surface fix targets recorded (D1): byTeam=seed Org C team tags; interview=align the
report's sim_id with the config seeder's interview sim_ref; deadline=wire the started hero into the active
cycle's deadline funnel. No fix attempted (a multi-fix seed follow-on beyond this run's budget after the major
launched_by fix landed).

## Close — 2026-07-23

**Outcome:** complete live characterization of the 3 remaining ai-readiness sub-renders → 3 distinct seed/wiring gaps at v1.341.0, each with a precise seed-fix target (0 platform edits); no fix landed. gate c stays 13/16.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET (gate c 13/16; metric 7/8)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: y (budget/scope checkpoint — the final 3 Playthroughs need 3 distinct seed-fix cycles; surface to orchestrator for next-run direction after the major launched_by fix landed) — (5) cap-reached: n (tik 2/5) — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D1 (the 3 root causes + per-surface seed-fix targets), D2 (disposition: characterized + routed, multi-fix follow-on)
**Side-deliverables:** none
**Routes carried forward:** FIND-M244-aireadiness-subrenders (the 3 seed/wiring fixes) → next run: (a) seed Org C team tags (find queryMemberTags's table); (b) single-source the interview report sim_id ↔ config seeder interview sim_ref; (c) wire the started hero's active-cycle deadline. Each: rext stack-seeding → tag/push → re-pin billion → re-seed pt-world → re-run `run-playthroughs.sh 1 --grep aireadiness`. 0 platform edits.
**Lessons:** a wholesale-blocker fix (launched_by zero-state) can UNMASK narrower per-surface data gaps that were previously invisible; grade the remaining failures by their exact assertion (heading-present-data-empty ≠ component-absent) to route seed-vs-frontend correctly.
