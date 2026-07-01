**Type:** tik — first tik of M51, under TOK-01 (active-cycle signals-true additive-to-stories seed),
coverage-protocol.md Phase A–E. Strand 1 of TOK-01's 4-strand plan: the 3rd-org identity + the org-enablement gate
+ the baseline manager-vantage sweep.

# iter-02 — tik (strand 1: the 3rd AI-Readiness story + the org-enablement gate writer)

## What happened (Phase A–E)

- **Phase A (measure):** N/A pre-build — the 3rd org did not exist, so no sweep could run. Baseline check confirmed:
  `organization_settings` 0 rows on demo-1; stories preset = 2 stories. Target untouched (Step-0 re-survey: no
  substitution; TOK-01's named strand-1 target still meaningful).
- **Phase C (fix) — the build:**
  1. **3rd story appended** to `stack-seeding/presets/stories.seed.yaml`: org "Northwind Aviation" (slug
     `northwind-aviation`, `narrative: ai-readiness`, size 200, aviation-services), hero trio — Aria Holt
     (thriving, COMPLETED-pinned), Ben Castellano (struggling, STARTED-pinned), Dana Whitlock (manager, views the
     dashboard). The two employees carry verified>0 (the verified-skill chain prereq); the AI-readiness STAGE is a
     separate funnel signal strands 2/3 layer on, keyed on the trajectory.
  2. **Net-new `OrgSettingsSeeder`** (`seeders/org_settings.go`, surface `org-settings`, depends-on `org`,
     PerStackIsolated): writes one `public.organization_settings` row (`setting='ai_readiness', is_enabled=true`) per
     story whose org declares `narrative: ai-readiness` — the org-enablement half of the dashboard gate (the FE-side
     PostHog `flag_ai_readiness` is out of seeder reach, documented as the remaining flag step). Deterministic id from
     `(org-id, setting)`, idempotent (`CopyRowsIdempotent ON CONFLICT (id)`, backed by the DB `UNIQUE(setting,
     organization_id)`). Registered in `cmd/stackseed/main.go` after `OrgSeeder`.
  3. **Unit tests** (`seeders/org_settings_test.go`, 4): opt-in-story-only, deterministic-id + idempotent,
     no-opt-in-story-writes-nothing, legacy-single-org-writes-nothing. Updated `blueprint/presets_test.go`
     `TestStoriesPreset` for the 3rd story (2→3 stories, all-distinct-org-ids, the AI-readiness narrative/size pin).
  4. **build/vet/gofmt clean**; full `blueprint` + `seeders` + `cmd/...` suites GREEN.
  5. **Re-seeded demo-1 in place** (authoring-copy `stackseed` against demo-1's offset Postgres :15432): `org-settings
     rows=1`, isolation guard CLEAN (no shared/external writes). Verified in the live DB: the `ai_readiness` gate row
     is scoped to Northwind Aviation ONLY (Cervato/Solvantis un-enabled); Northwind has its own per-story org id
     (`d7bb7482-…`, distinct from LegacyOrgID).
  6. **Re-exported roster + cockpit manifest** (9 identities / 9 heroes) to demo-1's stack dir + restarted fake-fapi +
     fake-bapi — so the manager seat Dana (org_role=admin) resolves for the cockpit handshake (D3).
- **Phase D (re-measure) — the baseline sweep:** manager vantage, seat `dana-manager`, expected-org "Northwind
  Aviation", `COVERAGE_NO_GATE=1` (diagnostic baseline), `COVERAGE_MAX_PAGES=200`. Result over a **frontier-EXHAUSTED**
  crawl (52 reachable pages, not capped): **`(failingSections=6, escapes=0)`, personaFailures=0, notReached=0** —
  persona role-skills coherence + avatar consistency + org identity all OK, the cross-port studio-desk follow OK. The
  6 empties are the enterprise Workforce-dashboard DATA sections for the new org: `/enterprise/workforce` (2 of 5
  sections empty), `/enterprise/members` (2), `/enterprise/assignments` (1), `/enterprise/activity-dashboard` (1) —
  the base Workforce surfaces awaiting the new org's full dashboard-data population + warm (and the AI-readiness
  funnel tab itself, strands 2/3). These are the iter-03+ targets.

## Close — 2026-06-30

**Outcome:** 3rd AI-readiness org + the `organization_settings` enablement gate landed + verified in demo-1; first authoritative baseline manager-vantage sweep = `(failing=6, escapes=0)` over a frontier-exhausted crawl (the gate's starting point for the 3rd org).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (narrative discriminator), D2 (OrgSettingsSeeder scope/naming), D3 (in-place roster+cockpit re-export) — iter-02/decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:**
  - iter-03 (TOK-01 strand 2): the `ai_readiness_*` config seeder (cycle status=active ×1, skills ~5 core + enabling with real replayed-taxonomy node-ids, sims ×2 with the net-new sim-id pin, steps ×3) — the config the dashboard funnel reads.
  - iter-03+/iter-04 (TOK-01 strand 3): the 200-member signals-true funnel (user_skill_evidences Step-1 + ended/scored jobsim sessions for Steps 2/3 + ai_readiness_user_step_progress ×3) at ~80% all-3-complete; Aria→stage3 (COMPLETED), Ben→stage1 (STARTED).
  - The 6 empty enterprise sections: triage at iter-03 — likely the new org needs the full dashboard-data population warmed (membership-skills/assignments/activity for Northwind) + the AI-readiness funnel tab populated; `/enterprise/members` may hit the M46 org-scale perf-wall class (200 members) — confirm by screenshot (skeleton vs genuinely-empty) at iter-03 Phase B.
  - iter-04+ (TOK-01 strand 4): cockpit jump_to re-point (Dana → the AI-readiness manager route once its deep-link is confirmed) + DeepLinkCatalog labels.
**Lessons:** A new story's manager seat is NOT live until the roster+cockpit are re-exported AND fake-fapi/fake-bapi restarted — the coverage-protocol "Wrong identity/org" re-apply is mandatory before any sweep of a newly-seeded org (fold into strand-1 of any future additive-story iter). The baseline correctly reads `escapes=0` + persona-OK on a brand-new org — the empties are pure seed-data-population, not identity/wiring, which validates the additive-to-stories strategy (the org identity + seat + chrome came "for free").
