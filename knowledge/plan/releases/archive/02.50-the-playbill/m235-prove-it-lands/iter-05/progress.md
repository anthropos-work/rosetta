# iter-05 — progress

**Type:** tik (under TOK-01, Track A step 2 — non-simulation product sections)

## What landed
The **skill-path-legacy** content-story section + the shared non-simulation projection/seeder infra
(`seeders/content_nonsim.go`):
- **Model + registry:** `nonSimExhibit` + `nonSimExhibits()` (2 skill-path exhibits pinned to REAL public
  `directus.skill_paths` ids sourced offline from the captured snapshot — a completed "Become a Product
  Manager" + a 45% "GenAI intro").
- **Projection:** `buildNonSimProducts` / `appendNonSimProducts` — the `/skill-path/<id>` player route + the
  `local_skill_path_sessions`-mirror manager route (`/enterprise/activity-dashboard/skill-paths/<id>/<uid>`),
  fail-closed, self-contained flat-index owner pairing (single-sourced with the seeder). Wired into
  `BuildContentProducts` + `ValidateContentManifest` + `WriteContentManifest`.
- **Seeder:** `ContentStoryNonSimSeeder` (skill-path arm) writes `skillpath.skill_path_sessions` (real
  progress) + the `public.local_skill_path_sessions` MIRROR (non-blank manager scoreboard, the M219/M222
  trap). Version `"2"` (collision-safe), documented status enum, deterministic + idempotent + audited + n=0
  guard + honest degradation. Registered in `stackseed`.
- **Believable titles:** new `Label` field (omitempty) on the manifest row → the cockpit shows the real
  skill-path title (not "Session"); `cockpit.py` prefers `label`. Sim rows byte-unchanged.
- **Honesty gate:** regenerated `presets/content-manifest.json` (15 sessions = 13 sim + 2 skill-path); the
  `CanonicalFileMatchesProjection` gate is GREEN.

## Proof (unit — the LIVE (session x action) landing is M236)
- Go: `./seeders` + `./contentsession` + `./cmd/stackseed` GREEN; full `go test ./...` GREEN; `go vet` clean.
  New: 7 non-sim tests (skill-path projection, BuildContentProducts inclusion+order, seeder rows, the
  owner↔projection single-source invariant, determinism, no-host no-op, fail-closed drop-all). 3 stale M233
  assertions (sim-only / 13-sessions) updated to be non-sim-aware.
- Python: 45 content-tab + skill-path render tests GREEN (incl. a new skill-path render test: label→title,
  `/skill-path/` as-player CTA, as-manager CTA). The 6 pre-existing test_cockpit.py failures (Academy /
  OverlayJs — the M234-recorded set) are UNCHANGED; `cockpit.py` diff vs HEAD is ONLY the title line.
- rext tag: `playbill-m235-nonsim-skillpath`.

## Close — 2026-07-20

**Outcome:** the skill-path-legacy content-story section RESOLVES + seeds real progress + renders with a
believable title, unit-proven. Readiness: **1 of 3** non-simulation sections (from 0/3).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the milestone gate is the LIVE (session x action)-lands proof on a cold reset-to-seed →
M236; iter-05 unit-proves the skill-path substrate, not the live render).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (separate registry), D2 (owner-sharing), D3 (offline-pinned real ids), D4 (version-2
collision-safe) + the M236 live-calibration checklist (version-match / status-enum / mirror-uniqueness) —
see iter-05/decisions.md.
**Side-deliverables:** none (the 3 stale-assertion updates are in-scope: M235 intentionally extends the
projection the M233 tests pinned).
**Routes carried forward:** ai-labs presence section → iter-06; academy section → iter-07; the M236 Fate-3
handoff (live proof + the M230 carry-forward live items + the skill-path live-calibration checklist) →
iter-08. All the skill-path arm's live-render unknowns are M236 per the run-3 ruling.
**Lessons:** the manifest export path had TWO entry points (`buildContentProductsFromSet` in
`WriteContentManifest` vs `BuildContentProducts`) — extend the render entry (`BuildContentProducts`) AND the
export entry together, or the projection change never reaches the emitted JSON. `CopyRowsIdempotent` guards
only `ON CONFLICT (id)`, so any secondary UNIQUE constraint must be dodged by construction (a distinct
version), never left to chance.
