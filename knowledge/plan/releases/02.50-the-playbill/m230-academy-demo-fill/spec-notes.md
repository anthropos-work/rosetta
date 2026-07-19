# M230 — Spec notes

_(technical details / thresholds accumulate here during build)_

## Pre-flight audits — iter-01

**KB-fidelity (Phase 0b, 2026-07-19): GREEN.** Report: `kb-fidelity-audit.md`. All load-bearing topics
PAIRED + ALIGNED + code-verified; no blind areas, no stale load-bearing claims. The frontend-tier.md F4
correction is the milestone's own `Delivers:` target, not a blind area.

## Topic → doc → code triples (verified iter-01)

- **Content model (empty-grid root cause):** `corpus/services/ant-academy.md` § The Content Model →
  `stack-demo/ant-academy/code/src/lib/{serverTenant,backendContent,draftMode,draftCatalog}.js`.
  The seam: `serverTenant.js::getServerCatalogView()` = `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`.
- **Option C vehicle (demo-patch):** `corpus/ops/demo/demopatch-spec.md` →
  `.agentspace/rosetta-extensions/demo-stack/patches/{demopatch, manifest_loader.py}`. Academy-patch
  precedent already present: `patches/ant-academy-dev-origins/`.
- **DELIVERS doc:** `corpus/ops/demo/frontend-tier.md` § ant-academy →
  `.agentspace/rosetta-extensions/demo-stack/ant-academy.sh` (confirmed: sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` 0 times).
- **Gate measurement:** `corpus/ops/demo/coverage-protocol.md` (ANT_ACADEMY descriptor) →
  `.agentspace/rosetta-extensions/stack-verify/e2e/lib/coverage-manifest.ts` + `coverage.spec.ts`.

## Infra feasibility (iter-01 probe)

Docker up (28.5.1, 0 containers), 205Gi free, no ENOSPC. `stack-demo/` fully cloned (ant-academy/code + all
platform repos + rext consumption clone). rext authoring copy present. `~/.pgpass` present (prod DB read
plausibly available → Option B not a priori blocked, but heavier). **demo-1 injected images built 41h ago +
`demo-stack/stacks/demo-1` artifacts present → a cold /demo-up is FEASIBLE on this box** (not a hard blocker).
rext pinned at `casting-call-m228-hiring-scope-fix`.
