---
iter: 10
iteration_type: tok
tok_flavor: bootstrap
iter_shape: strategy
status: closed-fixed
---

# iter-10 — re-scoped-gate strategy (TOK-10)

**Type:** tok (bootstrap-flavored — re-scoped-gate strategy authoring; does NOT terminate the call).

The M42e gate was RE-SCOPED at commit `0eaab39` from the weak DOM-text-density metric to the believability
bar (real semantic content + per-section cardinality + persona self-consistency + fresh-demo-up
reproducibility). TOK-01 (`sweep-then-route-by-leverage`) operated over the retired metric and is SUPERSEDED.
This iter ratifies the externally-authored, user-approved design-plan (the 7 root causes + 9 phases P0–P8 +
4 answered USER DECISIONS) as the active strategy **TOK-10: persona-believability-by-root-cause**.

## Inputs
- `.agentspace/scratch/work-m42e/design-plan.md` — the 7 roots, the 9-phase plan, the answered user decisions.
- `knowledge/plan/.../m42e-employee-coverage/overview.md` — the re-scoped gate (believability bar).
- The fix-surface code (read this iter): persona.go (`resolveHeroSkills` flat-pool top-up), profile.go
  (`combinedNamedPool` + the bare-array `jsonStringArray` legacy-skills writes), taxonomyref.go /
  skillref_named.go (the `readFlat* ORDER BY node_id` junk head), assignments.go / skillpath_sessions.go /
  activity.go (the P3 activity surfaces), stories.seed.yaml (Maya's preset).

## Strategy (TOK-10, recorded in milestone-root decisions.md)
Fix believability at its ROOT CAUSES, each in a SEEDER (Go + tests) so it reproduces on a fresh demo-up. This
run executes P0→P3 (the persona + profile + activity half). Live demo-3 = fast measurement only; P8 fresh
demo-up = authoritative acceptance (later run).

## Next-tik direction
iter-11 (tik, P0): read-only baseline confirm on live demo-3 of the before-state of the P1–P3 pages.

## Close — 2026-06-25
**Outcome:** TOK-10 authored (re-scoped-gate strategy = the user-approved design-plan, P0–P8 root-cause build).
**Type:** tok (bootstrap-flavored)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap-flavored, does NOT exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (into iter-11 P0)
**Decisions:** TOK-10 (milestone-root decisions.md)
**Routes carried forward:** P0→P3 this run (iter-11–14); P4–P8 later runs.
**Lessons:** A gate re-scope retires the prior TOK and demands a fresh strategy before the next tik commits — recorded as a bootstrap-flavored TOK that ratifies the user-approved re-scope plan, continuing into tiks (not terminating, since the user already opted into the re-scoped build).
