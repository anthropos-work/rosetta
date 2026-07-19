---
milestone: M231
slug: content-stories-feasibility-spike
version: v2.5 "the playbill"
milestone_shape: section
status: archived
created: 2026-07-19
last_updated: 2026-07-19
depends_on: none
delivers: corpus/ops/demo/content-stories-routes.md
---

# M231 — content stories feasibility spike

**Status:** `archived` (completed 2026-07-19, GO)  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** none

## Goal
HARD go/no-go barrier (mirrors v2.4's M222): discover per-product player+manager result routes, prove-by-render which land from seedable rows vs runtime-computed-blank, confirm the prod-session sourcing+anonymization mechanism, catalog public sims by modality, rule AI-labs + the academy section in/out.

## Scope
### In
- Enumerate per (product x {player,manager}) the exact result route + classify each by probe render (renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface)
- Confirm the db-access read path can select interesting real prod sessions per type + the anonymization surface (which fields scrub, which free-text needs handling) + how to pin by prod session-id
- Catalog captured public sims by task modality (confirm >=2 voice + 1 code + 1 document assessment SOURCES exist)
- Assess AI-labs feasibility (labs-api nil) + the ant-academy 'session' question; author the manager-view eligibility matrix + result-route map

### Out
- Building the seeder/manifest/tab (M232-M234)
- Any platform edit to make a runtime page render (routes to a demo-patch or escalates — decided here)

## Delivers
`corpus/ops/demo/content-stories-routes.md`

## Open questions
- Does /sim/.../result/<sessionId> recompute live (unseedable) or read a persisted row a clone could seed?
- Which products actually HAVE a manager result route?
- Is invoking the platform's own session-clone subcommand in-stack acceptable under the zero-edit wall?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
