---
iter: iter-06
milestone: M235
type: tik
iteration_type: tik
status: planned
created: 2026-07-20
active_strategy: TOK-01 (two-track) — Track A step 2 (non-simulation product sections)
---

# iter-06 — AI-labs presence-only content-story section

**Type:** tik · **Active strategy:** TOK-01 Track A step 2.

## Step 0 — re-survey
iter-05 landed skill-path (1/3 non-sim sections). ai-labs is the next offline surface: M231 §5 ruled it
**OUT / presence-only** — `grade_result` isn't GraphQL-exposed, `/labs/[id]` reads a LIVE labs-api worker
(nil in a demo) → no seedable result surface → NO as-player/as-manager CTAs. The manifest registry already
carries ai-labs with `playerLink=false`; `resolveNonSimSession` currently drops it (default case). Target
unchanged, meaningful.

## Cluster / target identified
The **ai-labs** presence-only section: a seeded `lab_sessions` status/spend row surfaced on `/labs` +
`/enterprise/labs` (`mySessions`/`labSessions` GraphQL), projected as a muted "Activity & spend only" row
with NO CTA (the cockpit already renders that disposition, M234 D5).

## Hypothesis
Adding ai-labs exhibits to `nonSimExhibits()` + an ai-labs arm in `resolveNonSimSession` (presence row: Label
+ icon + seat, no player/manager path) + a `lab_sessions` seeder arm makes the ai-labs section RESOLVE as
presence-only, unit-proven; the cockpit renders it as a status line, not a dead button.

## Expected lift
Readiness 2/3 non-sim sections (from 1/3).

## Phase plan
- `content_nonsim.go`: ai-labs fields on `nonSimExhibit` + 1–2 ai-labs exhibits + the ai-labs presence arm in
  `resolveNonSimSession` (no route) + the `lab_sessions` seeder arm (12-char-hex id, status/spend, owned by a
  content-player).
- Relax `TestWriteContentManifest_JSON`'s "every session has a player path" assertion for presence-only rows.
- Regenerate `presets/content-manifest.json`; honesty gate green.
- Unit tests: ai-labs presence projection (no player path, no manager), the seeder's lab_sessions row, the
  cockpit presence-only render.

## Escalation / close-no-lift
- The exact `lab_sessions` DDL (NOT-NULL column set) isn't in the offline snapshot → the seed-column match is
  an M236 live-seed-calibration item (documented), same shape as the skill-path arm's live items. The
  presence-only manifest render + the structural seeder are fully offline-provable here.
