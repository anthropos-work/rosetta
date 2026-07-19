# M229 — Spec notes

## Pre-flight audits — S1/S2 (Phase 0b)
**KB-fidelity: the milestone IS the fix.** `corpus/services/ant-academy.md` carried load-bearing FALSE claims
("Backend services: None / No GraphQL calls / static JSON in the repo / does not depend on the backend Go services")
that mis-routed the F4 empty-grid carry into the platform repo for a whole release. Per the Phase 0b RED→"author the
doc as the milestone's work" path, M229's `Delivers → corpus/services/ant-academy.md` IS the correction. No separate
blind area; the fix is the deliverable. Verdict recorded as: RED-by-design, resolved in-milestone.

## The verified content model (code-cited, `stack-demo/ant-academy/code/`)
- `app/(authed)/page.jsx:47` `resolveCatalogView()` → `getServerCatalogView()` (signed-in) | `getPublicCatalogView()` (anon)
- `src/lib/serverTenant.js:143-146` `getServerCatalogView() = (await getBackendCatalogView(eids)) ?? emptyCatalogView()`;
  `:115` `emptyCatalogView() = { chapters:[], skillPaths:{}, series:[] }`; `:146` `draftsEnabled() ? mergeDrafts(view,eids) : view`
- `src/lib/backendContent.js:93` `getBackendCatalogView(eids)`: `if (!client) return null`; queries `academyCatalogSeries` +
  `academyCatalogSkillPaths`; `catch → return null`
- `src/graphql/server.js:14,18` endpoint `= process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`; `createServerGraphQLClient` throws if unset
- `src/lib/draftMode.js:46` `draftsEnabled() = NODE_ENV==='development' && ACADEMY_SHOW_DRAFTS ∈ {1,true}` (prod hard-block)
- Demo empty-grid: `demo-stack/ant-academy.sh` never sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`; `code/.env.example:25` ships it EMPTY
- Two catalog paths: grid READS app internal/academy via GraphQL; `build-catalog.mjs` WRITES the unrelated `public/catalog.json` FS index (Talk-to-Data)

## Handoff to M230
M230 fills the demo grid production-faithfully (no Draft chip): first tik decides Option C (sha-pinned demo-patch
restoring an FS-as-published fallback on the ephemeral clone) vs Option B (a firewalled academy-content snapshot surface
+ wire the endpoint + compose the subgraph). The draft layer (Option A) is documented here but rejected for the chip.
