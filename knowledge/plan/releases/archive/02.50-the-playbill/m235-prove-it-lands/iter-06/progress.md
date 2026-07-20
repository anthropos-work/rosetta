# iter-06 — progress

**Type:** tik (under TOK-01, Track A step 2 — non-simulation product sections)

## What landed
The **ai-labs** presence-only content-story section (`seeders/content_nonsim.go`):
- ai-labs fields on `nonSimExhibit` (template/mode/status/spend/tokens) + 2 ai-labs exhibits (a Python data
  pipeline + a React dashboard, believable spend lines).
- The ai-labs arm in `resolveNonSimSession` — a presence row (Label + seat, NO player path, NO manager view;
  M231 §5). The cockpit already renders this as "Activity & spend only — no result page".
- The `lab_sessions` seeder arm (`seedAILabsContentStory` + `nonSimLabSessionCols` + `labSessionID`) — a
  status/spend row (12-char hex id, owned by the content-player's org) surfaced on /labs + /enterprise/labs.
- Regenerated `presets/content-manifest.json` (17 sessions = 13 sim + 2 skill-path + 2 ai-labs); honesty
  gate GREEN.

## Proof (unit — no CTA landing to prove for ai-labs; the presence render + clean seed are M236)
- Go: full `go test ./...` GREEN. New: `TestBuildNonSimProducts_AILabsPresenceOnly` (no player path, no
  manager view, label+seat present) + `TestContentStoryNonSimSeeder_LabSessionRows` (12-char hex id, template/
  mode/status match). `TestWriteContentManifest_JSON` relaxed to allow presence-only rows (no player path).
- Python: 46 content-tab render tests GREEN incl. a new ai-labs presence render test (label→title, "Activity
  & spend only" note, NO CTA). The 6 pre-existing test_cockpit.py failures unchanged.
- rext tag: `playbill-m235-nonsim-ailabs`.

## Close — 2026-07-20

**Outcome:** the ai-labs presence-only section RESOLVES (status/spend line, no CTA) + the lab_sessions
presence row seeds, unit-proven. Readiness: **2 of 3** non-simulation sections.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (ai-labs is presence-only — no CTA landing to prove; the milestone's live gate + the clean
lab_sessions seed calibration are M236).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (presence-only by disposition), D2 (lab_sessions activity row) + the M236 lab_sessions-DDL
live-seed-calibration item — see iter-06/decisions.md.
**Side-deliverables:** none (the WriteContentManifest presence-row relaxation + the iter-05 skill-path test's
find-by-id scoping are in-scope: M235 adds presence-only rows the M233/iter-05 tests hadn't seen).
**Routes carried forward:** academy section → iter-07; the M236 Fate-3 handoff (live proof + M230 c/f +
skill-path & lab_sessions live-calibration items) → iter-08.
**Lessons:** a presence-only row is a first-class disposition — assertions that "every emitted session has a
player path" must special-case it (no player path AND no manager view = presence-only, not a fail-closed
breach). App-side tables (lab_sessions) aren't in the public snapshot, so their exact DDL is inherently an
M236 live-seed item, not an offline unit fact.
