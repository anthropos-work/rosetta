# M229 — Progress

## Sections

- [x] **S1 — Rewrite `corpus/services/ant-academy.md`** (the DB-authoritative content model)
  - [x] High-Level Summary: corrected "does not depend on the backend" → reads catalog from the academy backend via GraphQL, degrades to empty
  - [x] Key contrasts: corrected "No GraphQL subgraph" (consumes ≠ provides) + "static JSON" (DB-authoritative)
  - [x] NEW "### The Content Model — DB-authoritative catalog (v0.5.1 M7)": the read chain (page.jsx → resolveCatalogView → getBackendCatalogView → academy subgraph), the 3 empty-grid failure legs → emptyCatalogView, the two-catalog-path disambiguation (grid GraphQL vs public/catalog.json FS index), the draft layer, and WHY a demo grid is empty (F4 real root cause)
  - [x] Integration Points: corrected "Backend services: None" → the GraphQL academy-subgraph read
  - [x] Architecture diagram: added the `Academy -->|catalog: GraphQL| App` edge
  - [x] Demo blockquote: corrected the F4 "render defect" mis-attribution
- [x] **S2 — Correct the F4 attribution in `corpus/ops/demo/frontend-tier.md`**
  - [x] The F4 mention: "empty-catalog render defect" → the DB-authoritative root cause (unset endpoint + empty DB), cross-linked to ant-academy.md § The Content Model, points at M230 for the fill
  - [x] The "no platform-backend link" nuance: corrected (the GraphQL catalog read IS a backend link; it's just not identity-resolving)

## Verification
- All code claims spot-verified against the actual source in `stack-demo/ant-academy/code/`
  (page.jsx resolveCatalogView; serverTenant.js emptyCatalogView + mergeDrafts; backendContent.js
  getBackendCatalogView; draftMode.js draftsEnabled; graphql/server.js NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT)
  and `demo-stack/ant-academy.sh` (endpoint NOT set) + `.env.example` (endpoint shipped empty).
- No stale "static JSON / no backend / Backend services: None" claims remain (grep-verified).
