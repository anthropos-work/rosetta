# iter-01 — decisions (bootstrap tok)

**D1 — Baseline reading: the manager frontier has TWO exploding fan-outs; the authoritative sweep needs sample
rules first.** The live no-gate manager sweep (demo-3, dan-manager, cap 250) timed out at the 25-min test
budget after 165 pages without writing a report, because the manager nav links two template-identical
families: `/user/<id>` (+`/skills`+`/activities`) team-roster, AND
`/enterprise/activity-dashboard/{ai-simulations,skill-paths,interviews}/<uuid>` per-activity drill-downs (with
a nested `/<uuid>/<uuid>` level). The employee vantage had only one (`/sim/<slug>`). So the sample-rule work
(TOK-01 line 3) is a PRECONDITION for the authoritative manager gate sweep, not a polish step. The iter-23
smoke-sweep gate numbers (escapes=139, notReached=5, frontier CAPPED) remain the baseline (this run confirmed
the `/enterprise/` route prefix + the second fan-out but did not produce a fresh report).

**D2 — The manifest `/workforce/*` paths are a route-model error, not a content gap.** The seeded manager
`jump_to` + the live crawl both show the real route prefix is `/enterprise/` (`/enterprise/workforce?tab=…`,
`/enterprise/assignments/*`, `/enterprise/activity-dashboard/*`). The manifest's `MANAGER_PAGES`
(`/workforce`, `/workforce/teams`, …) match nothing → they surface as `notReached=5`. TOK-01 line 2
reconciles the manifest to the real tab-query route model before any seed/serve-grant content work.
