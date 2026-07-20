# iter-06 — decisions

## D1 — ai-labs is PRESENCE-ONLY by manifest disposition, not a flag
The registry already marks ai-labs `playerLink=false`. The projection's ai-labs arm builds a row with a
Label + owner seat but NO `player_result_path` and NO manager view — a legitimate disposition (the cockpit
renders "Activity & spend only — no result page", never a dead button). No as-player/as-manager CTA is ever
projected (M231 §5: `grade_result` isn't GraphQL-exposed, `/labs/[id]` reads a live labs-api worker that is
nil in a demo → no seedable result surface). Proven by `TestBuildNonSimProducts_AILabsPresenceOnly` +
`test_ai_labs_presence_row_renders_label_and_status_note_no_cta`.

## D2 — the seeded lab_sessions row is the activity the presence line refers to
The ai-labs seeder arm writes a `public.lab_sessions` status/spend row (12-char hex id per the fake
in-memory idGen shape, template/mode/status/spend/tokens, owned by the content-player member's org) — so
`/labs` (`mySessions`) + `/enterprise/labs` (`labSessions`) list the activity the cockpit's presence row
names. status="stopped" (a completed session's spend line); `grade_result` stays NULL (nil worker never
grades — omitted from the written columns).

## M236 LIVE-SEED-CALIBRATION ITEM — the exact lab_sessions DDL
`public.lab_sessions` is an app-side Ent table, NOT in the offline public snapshot, so the EXACT NOT-NULL
column set can't be verified offline. The seeder writes the columns backend.md documents (id, user_id,
organization_id, template, mode, status, budget_usd, spend_usd, total_tokens, started_at, stopped_at). If the
live schema requires additional NOT-NULL columns (e.g. created_at/updated_at) or names differ, the COPY would
error — M236's cold seed surfaces + calibrates it (adjust `nonSimLabSessionCols()`). Same shape as the
skill-path arm's live items (iter-05/decisions.md). NOTE: ai-labs has NO CTA landing to prove, so its M236
work is just "the seed writes cleanly + the presence row renders" — no result-page render to calibrate.
