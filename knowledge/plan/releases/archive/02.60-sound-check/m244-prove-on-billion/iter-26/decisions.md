# iter-26 — decisions

## D1 — the 3 remaining ai-readiness sub-failures are DISTINCT seed/wiring DATA gaps at v1.341.0 (all seed-fixable, 0 platform edits)
Investigated each against billion's actual v1.341.0 read paths (`stack-demo/app/internal/workforce`) + the live
DB. All three are DATA gaps, NOT a stale frontend (the interview panel renders a HEADING → the component is
present, its data empty) and NOT the launched_by zero-state (fixed iter-25):

- **`manager-dashboard.UC1` — byTeam "AI Readiness by Tag" absent.** `aggregateByTeam` (ai_readiness.go:750)
  groups by `mem.Tags`; `mem.Tags` comes from `queryMemberTags(orgID, memberIDs)` (members.go:485). Org C
  (Vertex) members have **no team tags** (the workforce orgs get them; Org C's ai-readiness seed path doesn't),
  so ByTeam is empty → the frontend hides the section. **Fix:** seed team tags on Org C members (rext
  stack-seeding — the tag-assignment the workforce orgs already get, extended to the `narrative: ai-readiness`
  org). The tag table is NOT `memberships.tags`/`workforce_member_tags` (neither exists at v1.341.0) — find
  `queryMemberTags`'s actual source table first.

- **`manager-dashboard.UC2` — interview-findings panel = 24 chars (heading only).** The
  `jobsimulation.interview_aggregated_reports` row IS seeded (sim_id `2c861352…`, session_count 31, report
  8396 chars). `computeInterviewInsightsV2` (how_we_measure_v2.go) joins `ai_readiness_sims` (col `sim_ref`,
  2 rows for Org C) `ON ars.sim_ref = session.sim_id::text` for `step_type='interview'`. **Hypothesis:** the
  seeded report's `sim_id` (2c861352) does not match Org C's interview `ai_readiness_sims.sim_ref`, so the
  findings query returns empty. **Fix:** align the interview-report seeder's `sim_id` with the config seeder's
  interview `sim_ref` (single-source the interview sim id across both seeders).

- **`member-funnel.UC2` (member-progress) — cycle deadline absent.** The active cycle end_date (2026-08-22)
  exists + member-done resolves the active cycle, so the started-hero's in-progress funnel deadline render is
  a narrower gap (queryActiveCycleEndDate → the started member's funnel deadline field). **Fix:** verify the
  started hero (theo) is wired into the active cycle's deadline-bearing funnel; likely the same seed alignment.

## D2 — disposition: characterized + routed; the launched_by fix (iter-25) is the run's landed deliverable
These 3 are a **multi-fix seed follow-on** (handler FIND-M244-aireadiness-subrenders), each needing a rext
stack-seeding change → tag/push → re-pin → re-seed → re-run — 3 distinct cycles, beyond this run's remaining
budget after the major launched_by provisioning fix landed (iter-25, rext c755370). No fix attempted this iter
(a complete investigation ending in characterization). gate c stays **13/16**; metric 7/8. Routed to the next
run with precise per-surface targets above. 0 platform edits in any proposed fix.
