# M42m — Decisions

Implementation decisions with rationale, numbered `M42m-D1`, `M42m-D2`, … . Cross-iter decisions live here;
per-iter detail lives in each `iter-NN/decisions.md`. The strategy-of-record `TOK-NN` entries also live here
(the milestone's strategy-evolution chain).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|

## TOK-01: manager-coverage — reconcile-route + clear-escape + populate-dashboard + exhaust-frontier — 2026-06-25

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Drive the M42e semantic harness as `dan-manager` against demo-3 and close the manager
gate (same (a)-(d) bar) in four leverage-ordered fix lines, all in rext (zero platform edits):
1. **Studio-link escape (highest leverage — clears 139 in one fix).** The baked `studio.anthropos.work`
   left-nav link renders on every authenticated manager page (the enterprise nav; employee nav omits it).
   Diagnose the rendered DOM: is the link host a `NEXT_PUBLIC_*_URL` the demo injection can rewrite to the
   local studio-desk offset port (`:39000`), or hardcoded in next-web? If env-configurable → rewrite in the
   demo injection (`gen_injected_override.py` / `up-injected.sh` build-args). If HARDCODED → **re-scope
   trigger** (escalate, record, do NOT edit the platform).
2. **Reconcile the manager manifest to the real route model + populate the dashboard.** The seeded manager
   `jump_to` is `/enterprise/workforce?tab=skills-verification` (NOT the manifest's `/workforce/*` guesses —
   confirmed in `stories.seed.yaml` + `test_cockpit.py` + the live sweep's `/enterprise/*` route prefix).
   Re-author `MANAGER_PAGES` to the real `/enterprise/workforce` tab-query route model (the dashboard is ONE
   tabbed route, not 5 sub-routes), then verify each M36 surface (verification funnel / teams / role-readiness
   / succession / mobility) renders REAL seeded data. The 6 M36 seeders already write the data + CORS is
   already wired (`CORS_EXTRA_ORIGINS`), so seed/serve-grant only what the live render proves empty.
3. **Sample rules + raise the cap for the TWO manager fan-outs** so the frontier EXHAUSTS honestly: the
   team-roster `/user/<id>`(+`/skills`+`/activities`) AND the per-activity
   `/enterprise/activity-dashboard/{ai-simulations,skill-paths,interviews}/<uuid>` drill-downs. This is a
   PRECONDITION (the cap-250 baseline sweep timed out at 25 min without them) — not just polish.
4. **Calibrate the manager manifest** floors/selectors against the live render (`calibrated:false → true`).

**Rationale:** the manager gate reuses the proven M42e harness + the proven employee fix surfaces
(seeding / injection / sample-rules / manifest-calibration), so the right opening move is leverage-ordered
fix lines, not a new mechanism. The escape line is a single high-value fix (139→0 if env-configurable). The
route reconciliation turns the unmeasured notReached=5 into a measurable content gate. The sample rules are a
hard precondition the baseline sweep proved. Persona already PASSES for dan-manager (identity machinery
generalizes for free — iter-23), so no persona work is expected.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare).

**Distance-to-gate context:** the M42e iter-23 manager smoke-sweep: `escapes=139 notReached=5
frontier=CAPPED(+79) failingSections=0(over reached set only) personaFailures=0`. Gate = `(0,0,0,0) +
frontier EXHAUSTED`. The baseline (this iter, cap 250) confirmed the `/enterprise/` route prefix + a SECOND
fan-out (`/enterprise/activity-dashboard/.../<uuid>`) the smoke-sweep didn't reach; it timed out before
writing a report, so the iter-23 numbers stand as the baseline.

**Next-tik direction (iter-02):** start with the **Studio-link escape** (line 1). Log in as dan-manager on
demo-3, capture the rendered left-nav "Studio" `<a href>` host + how it's constructed (env var vs hardcoded).
If env-configurable, land the host-rewrite in the demo injection (rext), re-build the demo frontend, re-sweep
(scoped, with the new sample rules so it can finish) to confirm escapes→0. If hardcoded → re-scope trigger.
