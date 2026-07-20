---
milestone: M229
slug: academy-content-model-re-ground
version: v2.5 "the playbill"
milestone_shape: section
status: archived
created: 2026-07-19
last_updated: 2026-07-19
depends_on: none
delivers: corpus/services/ant-academy.md
---

# M229 — academy content model re ground

**Status:** `archived` (completed 2026-07-19)  ·  **Shape:** `section`  ·  **Complexity:** small  ·  **Depends on:** none

## Goal
Correct the stale/misleading ant-academy.md to the true DB-authoritative catalog model + the demo empty-render root cause, BEFORE any Thread-A fill code (the KB-fidelity prerequisite that mis-routed F4 for a whole release when wrong).

## Scope
### In
- Rewrite corpus/services/ant-academy.md: remove 'Backend services: None / no GraphQL / static JSON'; document the v0.5.1 M7 DB-authoritative path (page.jsx -> resolveCatalogView -> getBackendCatalogView -> academy subgraph)
- Document WHY a demo grid renders 0 cards (unset NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT + empty app academy tables -> emptyCatalogView)
- Disambiguate the two 'catalog' paths (grid READS app internal/academy via GraphQL; build-catalog.mjs WRITES the unrelated public/catalog.json FS index)
- Note the ACADEMY_SHOW_DRAFTS / NODE_ENV=development -> mergeDrafts() draft layer; correct the F4 mis-attribution in frontend-tier.md

### Out
- Any code/env change (M230)
- The Content-stories tab (Thread B)

## Delivers
`corpus/services/ant-academy.md`

## Open questions
- Should ant-academy.sh wire NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT to the demo offset router regardless of fill strategy?
- Is the academy subgraph composed into the demo's offset Cosmo router?

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.5 "the playbill" for the authoritative milestone design + the release-level decisions/risks (research provenance: `.agentspace/scratch/roadmap-research-2026-07-19` via the design-content-stories-research workflow).
