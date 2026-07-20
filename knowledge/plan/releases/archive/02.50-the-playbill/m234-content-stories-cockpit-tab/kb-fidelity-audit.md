---
title: "KB Fidelity Audit — M234 content-stories-cockpit-tab"
date: 2026-07-19
scope: milestone:M234
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| content-manifest projection (M233) | `corpus/ops/demo/content-stories-spec.md` | `stack-seeding/seeders/content_manifest.go` + `presets/content-manifest.json` | PAIRED |
| per-product result-route map + AI-labs/academy verdicts (M231) | `corpus/ops/demo/content-stories-routes.md` | `content_manifest.go` `contentProductRegistry` | PAIRED |
| presenter cockpit render (stories tab) | `corpus/ops/demo/cockpit-spec.md` | `demo-stack/cockpit.py` | PAIRED |
| **content tab render (M234 deliverable)** | `content-stories-spec.md` §6 (scoped to M234) | `cockpit.py` (to extend) | DOC-ONLY (planned, M234 delivers) |
| content-player seat + Clerkenstein roster | `content-stories-spec.md` §3 + `corpus/services/clerkenstein.md` | `stack-seeding/seeders/roster.go` | PAIRED (M234 extends) |
| bring-up cockpit launch | `cockpit-spec.md` "Bring it up" + `up-injected.sh` | `demo-stack/up-injected.sh` | PAIRED (M234 extends) |

## Fidelity Findings

1. **content-stories-spec.md §2 schema ↔ content_manifest.go types — ALIGNED.** Every field in the
   doc's `content_products[]` JSONC (`stack`, `products[].{id,name,app_base,icon_key,sessions[]}`, per-session
   `{key,source_session_id,sim_id,sim_type,modality,passed,icon_key,player_seat,player_result_path,has_manager_view,manager_seat,manager_result_path}`)
   matches the Go structs exactly.
2. **content-stories-spec.md §2 per-product registry table ↔ `contentProductRegistry()` — ALIGNED.**
   Verified row-by-row: `simulation`(web/flask/player✓/ai-simulations), `skill-path-legacy`(web/diagram-project/player✓/skill-paths),
   `skill-path-new`→academy(academy/graduation-cap/player✓/no-manager), `ai-labs`(web/vials/**presence-only**/no-manager).
   Per-sim_type icons (assessment/training/hiring/interview) match `simTypeIcon`.
3. **content-stories-spec.md §3 seat model ↔ code — ALIGNED.** `content-player-<memberIndex>` via
   `eligiblePlayerOwnerSlots` + the flat `slots[idx % len]` assignment; the survives-drops flat-index invariant
   is pinned. Doc claim "the roster today carries only heroes — the non-hero player seats are M234's roster.go
   extension" is **accurate** (`roster.go` has 0 content-player references; `BuildRoster` iterates heroes only).
4. **content-stories-routes.md §5/§6 dispositions ↔ registry — ALIGNED.** AI-labs OUT/presence-only
   (`playerLink=false`), academy IN with no manager route (`managerKind=""`), the `has_manager_view` matrix
   (TRUE for the four sim manager routes + skill-path-legacy-on-web; FALSE for academy + ai-labs). M234's
   renderer honors these via the manifest fields.
5. **Cross-references resolve.** All `content-stories-spec.md`/`content-stories-routes.md` "See also" targets
   exist (`session-clone-spec.md`, `content-stories-routes.md`, `seed-manifest-spec.md`, `cockpit-spec.md`,
   `safety.md`, `clerkenstein.md`).

## Completeness Gaps

1. **cockpit-spec.md does not yet describe the 2nd "Content stories" tab** — by design: it documents the
   stories tab, and `content-stories-spec.md` §6 explicitly scopes the tab render + bring-up export +
   seat registration to M234. **Not a blind area** — the topic is anchored and M234 delivers the extension
   (into `content-stories-spec.md`, the render half). Tracked as the milestone's `Delivers →`.

## Applied Fixes
None needed — all M234-relevant docs are fresh (M231/M232/M233 landed them 2026-07-19) and accurately
describe both the current state and the M234-pending work.

## Open Items (require user decision)
None. Platform-side `<file>:<line>` citations in `content-stories-routes.md` (e.g. `queries.resolvers.go:70`)
reference the platform checkout (not in this tree) and were code-cited/verified during M231 — UNVERIFIABLE
here, but not load-bearing for M234 (a tooling-only render/seat/wiring milestone).

## Gate Result
GREEN: proceed to Phase 1. Every M234 topic has a knowledge anchor; the render half + seat registration are
correctly described as M234's pending deliverables; no stale load-bearing claim, no blind area.
