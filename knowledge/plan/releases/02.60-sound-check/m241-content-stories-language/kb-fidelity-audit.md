---
title: "KB Fidelity Audit — M241 content-stories language"
date: 2026-07-22
scope: milestone:M241
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed. No blind areas (the one new topic, the language axis, is a DECLARED milestone
deliverable with a `Delivers → corpus/ops/demo/content-stories-spec.md` line), no stale load-bearing
claims. One incidental completeness gap tracked below (KB-1).

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| content_products manifest schema | `content-stories-spec.md` §2 | `seeders/content_manifest.go` | PAIRED |
| honesty gate (CanonicalFileMatchesProjection) | `content-stories-spec.md` §4 | `cmd/stackseed/main_test.go:1210` | PAIRED |
| session sourcing (reproducible SQL) | `session-clone-spec.md` §2 | `contentsession/sourcing.go` | PAIRED (M240-fresh) |
| result-fan-out replay / session write | `session-clone-spec.md` §3 | `seeders/content_stories_write.go` | PAIRED |
| content-stories LANDS sweep + denominator | `content-stories-spec.md` §4 + `content-denominator.json` | `stack-verify/e2e/run-content-stories.sh`, `aggregate-content.py`, `tests/content-route-contract.unit.spec.ts` | PAIRED |
| **language axis** (session `language` col → cs.Language → manifest → cockpit toggle) | `content-stories-spec.md` (declared Delivers, not yet written) | `content_stories_write.go` (`sessLanguageEnglish` hardcoded) | **DOC-TO-DELIVER** (milestone deliverable) |

## Fidelity Findings
1. **content-stories-spec.md §2 schema** — ALIGNED. The `content_products[]` JSON shape (product id/name/app_base/icon_key + per-session key/source_session_id/sim_id/sim_type/modality/passed/icon_key/player_seat/player_result_path/has_manager_view/manager_seat/manager_result_path) matches `ContentProduct`/`ContentProductSession` in `content_manifest.go` exactly. M241 ADDS a `language` field here — a new field, not a stale claim.
2. **content-stories-spec.md §4 honesty gate** — ALIGNED. `TestContentManifest_CanonicalFileMatchesProjection` (`cmd/stackseed/main_test.go:1210`) re-projects and byte-compares the checked-in `presets/content-manifest.json`; the `HasTeeth` meta-tests exist. Matches the doc.
3. **session-clone-spec.md §2 sourcing** — ALIGNED + M240-fresh. Documents the sim-TYPE-match predicate (`AND d.type = '<cell sim_type>'`, M240 Defect 1/CQ-1) and the score band — matching `sourcing.go`.
4. **content-denominator.json = 29 landable pairs** — ALIGNED. Arithmetic: simulation 13×2=26 + skill-path-legacy 2×1=2 + skill-path-new 1×1=1 + ai-labs 2×0=0 = 29. NOTE for build: if §3 adds IT-variant sessions the simulation session count grows past 13 → this denominator + the `content-route-contract.unit.spec.ts` pair assertion + the canonical `content-manifest.json` all move together (a deliberate, reviewer-visible edit per the file's own contract).
5. **`sessions.language` column (prod)** — VERIFIED LIVE (read-only). Exists, type `character varying` (varchar, NOT an enum), distinct values include `english` (7,682) and `italian` (11,971). Confirms the design assumption: EN = `"english"` (matches `sessLanguageEnglish`), IT = `"italian"`. Both genuine-blocker gates (prod read; column exists as assumed) are CLEAR.

## Completeness Gaps
1. **KB-1 (incidental):** `session-clone-spec.md` §3 (the write side) does not document that `content_stories_write.go` writes the session `language`. It currently hardcodes `sessLanguageEnglish`; M241 changes it to `cs.Language`. Add a one-line note in Phase 5 when the write flips. Low severity — the write-side doc is not the milestone's Delivers target (`content-stories-spec.md` is), but a language cross-ref there aids discoverability.

## Applied Fixes
None inline (no stale claims to correct; the language axis is authored as the milestone's deliverable in Phase 5, not backfilled here).

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed with tracking. KB-1 recorded in `decisions.md`. No RED conditions.
