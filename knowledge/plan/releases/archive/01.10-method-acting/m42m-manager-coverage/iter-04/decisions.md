# iter-04 — decisions (tik, TOK-01 lines 2-4)

**D1 — The manager Workforce dashboard is ONE tabbed route, not 5 sub-routes; the manifest `/workforce/*`
guesses were a pure route-model error.** Diagnosed live (dan-manager, demo-3): the real route is
`/enterprise/workforce` (a tabbed SPA — Growth / Skills & Verification / Talent Pool / Assignments / Activity
Log render IN-PAGE; the `?tab=` query is ignored client-side), plus the sibling `/enterprise/*` pages
(`/members`, `/assignments`, `/activity-dashboard`, `/organization-feedback`, `/settings`). `MANAGER_PAGES`
re-authored to these real routes; the notReached=5 was the wrong `/workforce/*` paths matching nothing, NOT a
content gap. (Confirms iter-01 D2 with the full route surface.)

**D2 — The M36 dashboard already renders REAL, substantial data — the populate work was ONE empty surface
(org-feedback), not the whole dashboard.** The live render of `/enterprise/workforce` is rich (221 members, 493
mapped, 262 verified / 53.1% coverage, 19 cards, 67 SVG charts) — the 6 M36 seeders + the M34/M35 chain already
populate the funnel / org-scale gap / teams / role-readiness / succession tabs. Only
`/enterprise/organization-feedback` was empty. So the believability (a)/(b) bar for the dashboard is met by the
existing seed once the route is reconciled — the iter's content work narrowed to the one empty page (D3).

**D3 — `/enterprise/organization-feedback` empty is an inserted-but-invisible seed bug (the mirror gap), fixed
in `stack-seeding` (zero platform edits).** The page renders "No data" / 0 sessions despite 103 seeded
`job_simulation_feedbacks` rows. ROOT CAUSE (read-only platform diagnosis): `GetOrganizationFeedback`
(`app/internal/data/ent/repository/jobsimulation.go`) JOINs feedback to the app mirror
`public.local_jobsimulation_sessions` on `feedback.session_id = mirror.jobsimulation_session_id` and scopes by
the **mirror's** `organization_id`. The population sessions the feedback links live only in
`jobsimulation.sessions`; only the `PersonaSeeder` writes the mirror (for heroes), so the join was empty. FIX:
the `FeedbackSeeder` now also writes a `local_jobsimulation_sessions` mirror per feedback session
(reconstructing the population session's coherent values from the same deterministic key). This is the
org-feedback analog of the M36-D2 "the dashboard reads the app mirror" rule + the G14 inserted-but-invisible
class. Live: joinable feedback 0→103, page "No data" → "103 sessions / 70% pos / 59% pass / 63% avg" + 21 rows.

**D4 — The two manager fan-outs get vantage-aware sample rules (the M42e `/sim` pattern, applied to the manager
families).** `/user/<uuid>`(+/skills+/activities) (team roster, 200+) + `/enterprise/activity-dashboard/.../<uuid>`
(per-activity drill-downs) are template-identical (no manifest sections, share their links). Added
`MANAGER_SAMPLE_RULES` (sample 8 each) + a vantage switch in `coverage.spec.ts`, so the frontier where
escapes/failures live EXHAUSTS while bounding the explosion (the baseline cap-250 sweep timed out without them
— iter-01 D1).

**D5 — `/enterprise/settings` is a documented terse exception (mirrors the employee `/settings`).** The org
settings shell (branding / AI-sim toggles / usage) is correctly low-cardinality; a substantial-cardinality floor
would be a false-fail. Floor relaxed to real-content; disclosed in `coverage-review.html` documentedExceptions.

## Manifest calibration (calibrated:false → true)
All 6 manager `PageDescriptor`s set `calibrated:true` (the M42m calibration TOK-01 line 4), with floors read off
the live render: `/enterprise/workforce` cards≥6 + tabs≥4 + funnel text; `/enterprise/members`,
`/enterprise/assignments`, `/enterprise/activity-dashboard` real-text asserts (named members + the column
headers); `/enterprise/organization-feedback` recap text + rows≥2; `/enterprise/settings` exception.
