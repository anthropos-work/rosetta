---
milestone: M233
slug: content-stories-manifest
version: v2.5 "the playbill"
milestone_shape: section
status: archived
created: 2026-07-19
last_updated: 2026-07-19
depends_on: M232
delivers: corpus/ops/demo/content-stories-spec.md
---

# M233 — content stories manifest

**Status:** `archived` (completed 2026-07-19)  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** M232

## Goal
Project a second, auditable content_products[] manifest block (peer to population.orgs[]) that pins each session's prod source deterministically and is honesty-gated so it cannot drift from the seeded sessions.

## Scope
### In
- A content_products[] projection (Simulation / Skill-path legacy / Skill-path new / AI-labs) each listing pinned sessions with player+manager seat keys, result paths, has_manager_view, per-product app-base, and a per-type icon key
- Project it from a preset via stackseed --content-export (or a 2nd block in BuildCockpitManifest), guarded by a CanonicalFileMatchesProjection-style test (the D9 single-source discipline)
- Fail-closed resolver when a pinned prod-source id doesn't resolve in the replay (no-fabrication); fold pinned sources into the downloadable seed-generation-manifest.yaml

### Out
- The cockpit render (M234)
- The seeder (M232)

## Delivers
`corpus/ops/demo/content-stories-spec.md`

## Open questions
- One manifest with a 2nd block + client tab, or a separate content-manifest.json + endpoint (better preserves D9 + the non-fatal bring-up)?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
