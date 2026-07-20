---
title: "KB Fidelity Audit — M230 academy-demo-fill"
date: 2026-07-19
scope: milestone:M230
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap tok pre-flight)
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| ant-academy DB-authoritative catalog model (empty-grid root cause) | `corpus/services/ant-academy.md` § The Content Model | `stack-demo/ant-academy/code/src/lib/{serverTenant,backendContent,draftMode,draftCatalog}.js` | PAIRED |
| demo-patch mechanism (Option C vehicle) | `corpus/ops/demo/demopatch-spec.md` | `.agentspace/rosetta-extensions/demo-stack/patches/{demopatch, manifest_loader.py, ant-academy-dev-origins/}` | PAIRED |
| academy-content snapshot surface (Option B path) | `corpus/ops/snapshot-spec.md` | `.agentspace/rosetta-extensions/stack-snapshot/` | PAIRED (not the leaning path; non-load-bearing if C wins) |
| frontend-tier / F4 carry (the DELIVERS doc) | `corpus/ops/demo/frontend-tier.md` § ant-academy | `.agentspace/rosetta-extensions/demo-stack/ant-academy.sh` | PAIRED (F4-attribution correction = milestone deliverable) |
| coverage rendered-card gate | `corpus/ops/demo/coverage-protocol.md` (ANT_ACADEMY descriptor) | `.agentspace/rosetta-extensions/stack-verify/e2e/lib/coverage-manifest.ts` + `coverage.spec.ts` | PAIRED |

## Fidelity Findings

### 1. ant-academy DB-authoritative catalog model — ALIGNED
- **Source:** `corpus/services/ant-academy.md` § "The Content Model — DB-authoritative catalog (v0.5.1 M7)"
- **Expected:** grid reads catalog from the academy subgraph over GraphQL; on any failure → `emptyCatalogView() = { chapters:[], skillPaths:{}, series:[] }` = 0 cards; three failure legs (endpoint unset / query error / empty DB); a demo is empty because the launcher never sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` AND the demo app DB has no academy rows; the dev DRAFT layer (`ACADEMY_SHOW_DRAFTS`) stamps a "Draft" chip.
- **Actual (code-verified):** `serverTenant.js::getServerCatalogView()` is literally `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`. `backendContent.js::makeClient()` returns `null` when the endpoint env is unset (→ null → empty view). `draftMode.js::draftsEnabled()` = `NODE_ENV==='development' && ACADEMY_SHOW_DRAFTS ∈ {1,true}`. `draftCatalog.js::mergeDrafts()` stamps `_draft:true`. `ant-academy.sh` sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` **0 times** (grep -c = 0).
- **Verdict:** ALIGNED. The doc is exact; M229's re-ground holds against real code.

### 2. demo-patch mechanism (Option C vehicle) — ALIGNED
- **Source:** `corpus/ops/demo/demopatch-spec.md` (7 guards, 10-key manifest, ephemeral-clone apply→revert)
- **Actual:** `demo-stack/patches/demopatch` (26 KB) + `manifest_loader.py` + ~10 patch dirs each with a `<name>.yaml` manifest. An **`ant-academy-dev-origins`** patch already exists — a proven precedent for patching the ephemeral **ant-academy** clone (the tailscale `allowedDevOrigins` sha-pinned patch).
- **Verdict:** ALIGNED. Option C's vehicle (patch the ephemeral academy clone, revert-clean) is real, wired, and already used against ant-academy.

### 3. frontend-tier.md § ant-academy (the DELIVERS doc) — IN-SCOPE WORK PRODUCT (not a blocker)
- **Source:** `corpus/ops/demo/frontend-tier.md` § "ant-academy — native, keyless, session-detached, with a documented fallback" (line ~335)
- **Observation:** the section still frames the empty academy grid as "a documented fallback," carries no F4-corrected attribution or M230 fill mechanism, and has a minor internal staleness (a "keyless" section header vs the M220 box that documents keyless being removed / Clerkenstein-wiring).
- **Verdict:** NOT a blind area and NOT a stale load-bearing claim the strategy would trust wrongly. Correcting the F4 attribution + documenting the shipped fill mechanism (+ tidying the keyless header) is the milestone's OWN `Delivers:` target. In-scope; addressed during the milestone.

## Completeness Gaps
None critical. Option B's snapshot/subgraph-composition surface (`snapshot-spec.md`) is documented but, if Option C is chosen, is not load-bearing for this milestone.

## Applied Fixes
None applied inline — no stale load-bearing claim found; the one doc-improvement (frontend-tier.md F4 attribution) is the milestone's own deliverable, deliberately left for the milestone body rather than pre-empted here.

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed. The load-bearing topics for the Option C strategy (content model + demo-patch mechanism) are PAIRED + ALIGNED + code-verified; no blind areas; no stale load-bearing claims. The bootstrap tok may author its strategy against verified knowledge docs.
