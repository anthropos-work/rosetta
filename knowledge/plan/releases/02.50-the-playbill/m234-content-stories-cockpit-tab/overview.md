---
milestone: M234
slug: content-stories-cockpit-tab
version: v2.5 "the playbill"
milestone_shape: section
status: planned
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M233
delivers: corpus/ops/demo/content-stories-spec.md
---

# M234 — content stories cockpit tab

**Status:** `planned`  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** M233

## Goal
Add the 2nd 'Content stories' tab to cockpit.py beside 'Org stories' — sections-per-content-product, a list of played sessions each with per-type icons and TWO login-and-land CTAs (as-player / as-manager, manager omitted where has_manager_view=false).

## Scope
### In
- A client-side tab toggle in render_page() (reuse the shipped _OVERLAY_JS pattern; stdlib-only, standalone-served)
- Per-product sections rendering the M233 manifest; per-session rows with per-type FontAwesome icons
- Two fake-FAPI deep-link CTAs per session (?__clerk_identity=<seat>&redirect_url=<base><result-path>), the .actions two-button layout + has_manager_view omitempty
- Per-product app-base routing generalizing the is_hiring/hiring_base switch (next-web :3000, apps/hiring :3001, academy :3077); mint/resolve per-session player seats via roster.go + Clerkenstein

### Out
- Any platform/next-web edit
- Making a runtime-computed result page render (M231's demo-patch/escalation decision)

## Delivers
`corpus/ops/demo/content-stories-spec.md`

## Open questions
- Does the academy section deep-link to a content page (post-M230), map onto a skillpath session, or render presence-only?
- Confirm the per-(simId,userId) manager drill-down deep-link (M224 deferred it as 'optional polish')

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
