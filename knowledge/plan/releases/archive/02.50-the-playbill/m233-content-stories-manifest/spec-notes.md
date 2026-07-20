# M233 — Spec notes

## Pre-flight audits — §1 content_products projection

- KB-fidelity audit (Phase 0b): **YELLOW** — `knowledge/plan/releases/02.50-the-playbill/m233-content-stories-manifest/kb-fidelity-audit.md`.
- One YELLOW completeness gap (KB-1, not load-bearing): `seed-manifest-spec.md` omits the M232 `content_sessions` block → Phase 5.

## Topic → doc → code triples (verified)

| Topic | Doc | Code |
|---|---|---|
| Seed+gen manifest + honesty gate | `seed-manifest-spec.md` | `stack-seeding/manifest/manifest.go` · `cmd/stackseed/main.go doManifestExport` · `main_test.go TestManifest_CanonicalFileMatchesProjection` |
| Content-session fixture + pins | `session-clone-spec.md` | `contentsession/{contentsession,content,sourcing}.go` + `fixture/` · `manifest.go buildContentSessions` |
| Result routes + app-base + manager-view | `content-stories-routes.md` | `next-web-app` routes · `packages/core-js/src/routes/jobSimulationRoute.ts` |
| Cockpit seats / deep-links | `cockpit-spec.md` | `seeders/cockpit.go` · `seeders/roster.go` |

## Load-bearing route facts (verified against stack-dev/next-web-app)

- **Player sim result:** `/sim/<slug>/result/<sessionId>` — `[slug]` resolves by `jobSimulationBySlug(slug)` (a TEXT slug, not the
  sim uuid). `jobSimulationRoute({slug, result:true, sessionId})` builds it. → the fixture needs `sim_slug` (public, non-PII).
- **Manager result (apps/web, Workforce org):** `/enterprise/activity-dashboard/<kind>/<simId>/<userId>` where kind ∈
  {`ai-simulations` (training/assessment/hiring), `interviews` (interview), `skill-paths` (skill-path-legacy)}. Uses the sim
  **uuid** + owner user-id directly → fully offline-derivable (no slug needed for the manager path). The `@tabs` slot is not in the URL.

## Key design decisions (M233)

1. **content-story sessions live in a WORKFORCE org → app_base = `web`** (not hiring), regardless of the source sim's sim_type.
   The seeder re-tenants every clone into `firstNonHiringStory` (a Workforce org). M231's "HIRING → apps/hiring" is the
   ORG-ejection rule for genuinely-hiring ORGS (the M224 cockpit) — a different feature. So a HIRING-*sim_type* clone in a
   Workforce org renders in apps/web. app_base keys on the HOST ORG type, not sim_type. (sim_type still drives the icon + the
   interview flag-gated render.)
2. **Slug resolved at authoring time into the fixture** (`sim_slug`, public + non-PII) so the projection is fully OFFLINE,
   concrete, and honesty-gated (no runtime placeholder in the checked-in manifest). The fail-closed resolver is then a pure
   function: drop (recorded reason) any session that can't form a real link (missing slug/seat/route); never fabricate.
3. **Open-question resolution:** BOTH surfaces, one projector. The honesty-gated `content_products` block folds into the
   checked-in `seed-generation-manifest.yaml` (the audit surface, `CanonicalFileMatchesProjection`); a standalone
   `content-manifest.json` via `stackseed --content-export` is the non-fatal RENDER surface M234 reads (mirrors
   `--cockpit-export`). One `BuildContentProducts(bp)` single-sources both.
4. **Seats:** player_seat = the owner MEMBER seat (single-sourced with the seeder's owner derivation); manager_seat = the host
   org's manager hero key. `has_manager_view` per the M231 matrix, DOWNGRADED to false (fail-closed) if no manager hero resolves.
