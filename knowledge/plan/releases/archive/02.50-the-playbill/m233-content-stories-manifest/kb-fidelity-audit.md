---
title: "KB Fidelity Audit — M233 content-stories-manifest"
date: 2026-07-19
scope: milestone:M233
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Consolidated seed+gen manifest | `corpus/ops/demo/seed-manifest-spec.md` | `stack-seeding/manifest/manifest.go`, `cmd/stackseed/main.go` (`doManifestExport`), `cmd/stackseed/main_test.go` (`TestManifest_CanonicalFileMatchesProjection`) | PAIRED |
| Content-story fixture + `content_sessions` pin block | `corpus/ops/demo/session-clone-spec.md` | `stack-seeding/contentsession/*`, `manifest.go` (`buildContentSessions`), `seeders/content_stories*.go` | PAIRED |
| Per-product result-route map + app-base + has_manager_view | `corpus/ops/demo/content-stories-routes.md` | `stack-dev/next-web-app` routes + `jobSimulationRoute.ts` | PAIRED |
| Cockpit deep-link / seat-switch model | `corpus/ops/demo/cockpit-spec.md` | `seeders/cockpit.go` (`BuildCockpitManifest`), `seeders/roster.go` (`BuildRoster`) | PAIRED |
| `content_products[]` manifest block (M233 deliverable) | `corpus/ops/demo/content-stories-spec.md` (delivered by this milestone) | — (this milestone builds it) | DOC-ONLY (delivered) |

## Fidelity Findings

1. **Player sim-result route resolves by TEXT SLUG, not the sim uuid.** `sim/[slug]/layout.tsx` fetches
   `useGetSimulation({ slug })` → `GET_SIMULATION` → `jobSimulationBySlug(slug)`. `jobSimulationRoute({ slug, result, sessionId })`
   builds `/sim/<slug>/result/<sessionId>`. **Verdict: ALIGNED** with `content-stories-routes.md`. **Implication for M233:**
   the fixture pins only `sim_id` (uuid); the player result path needs the public sim's **slug**. → resolve + carry `sim_slug`.
2. **Manager result routes exist in apps/web with the `[simId]/[userId]` leaf** for `ai-simulations`, `interviews`, and
   `skill-paths` (verified on the local checkout). URL = `/enterprise/activity-dashboard/{ai-simulations|interviews|skill-paths}/<id>/<userId>`
   (the `@tabs` slot name is not in the URL). Uses the sim **uuid** directly → fully offline-derivable. **Verdict: ALIGNED.**
3. **`manifest.go:73-76` explicitly stages M233:** `ContentSessions` is "the PROVENANCE half … M233 extends the projection
   with the player/manager seats + result paths (the full `content_products[]` the cockpit renders)." **Verdict: ALIGNED** —
   the M233 deliverable is anticipated by the code contract.

## Completeness Gaps

1. **KB-1 (YELLOW):** `seed-manifest-spec.md`'s block inventory (§1 "What it inlines") lists four blocks + `snapshot_sources`
   but does **not** mention the `content_sessions` block that M232 added to `SeedGenerationManifest` (`manifest.go:63-67`).
   Not load-bearing for M233's implementation (M233 reads the code, not the doc, for the struct shape). M233 adds a
   `content_products` block to the SAME file + doc → fold the `content_sessions` backfill into M233's Phase-5 doc pass.

## Applied Fixes
None inline (KB-1 folded into M233's own doc deliverable in Phase 5).

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed. KB-1 tracked in `decisions.md`; addressed in Phase 5 (the `seed-manifest-spec.md` update that documents both
the M232 `content_sessions` and the M233 `content_products` blocks).
