---
iteration_type: tik
status: closed-fixed
gate: NOT MET
---

# iter-03 — tik: Lane B, the net-new Directus set-dress

**Active strategy:** TOK-01. **Step 0 re-survey:** Lane A landed (iter-02); the directus set-dress is still
un-started and still the right next thing (gate part 2 needs a non-empty evaluated-skills list).

## Cluster / target
Net-new: populate `directus.simulations.skills` for the 3 wired sims so the platform's Step-2 cards render a
non-empty evaluated-skills list (and the tech/business track heuristic classifies correctly). Snapshot replay
is replay-only — no existing seam — so this is a NEW seeder.

## What landed
- **`seeders/ai_readiness_sim_skills.go`** (net-new) — `AIReadinessSimSkillsSeeder`: for each of the 3 default
  sims, a directus UPDATE that resolves `sequences.evaluation_skills` node-ids → `public.skills.name` and writes
  them into `directus.simulations.skills` as a JSON array of names. DependsOn {taxonomy, content};
  PerStackIsolated; two gates (opt-in ai-readiness org + per-stack Directus present via a `to_regclass` guard);
  idempotent; degrades cleanly (no directus → skip; unresolved node-id → drops, no fabrication).
- **`cmd/stackseed/main.go`** — registered the seeder (own hunk, next to AIReadinessConfigSeeder).
- **`seeders/ai_readiness_sim_skills_test.go`** (net-new) — 4 tests: enriches the 3 wired sims / skips with no
  Directus / skips with no opt-in org / registration contract.

## Measurement (Phase 3)
- `go build ./...` GREEN; new seeder tests 4/4; full stack-seeding module GREEN.
- **LIVE-validated on demo-2** (the smart pre-flight): the resolution query + the real UPDATE both run correctly
  against demo-2's postgres. All 3 sim uuids exist with the exact expected slugs; `simulations.skills` was NULL
  for all (0 populated demo-wide). Resolved names:
  - tech `who-can-see-this-document-fc0` → `["Ai Coding", "Critical Thinking Fundamentals", "Generative AI Fundamentals", "Prompt Engineering"]` → `techTrackRe` matches "Ai Coding" → track **tech**.
  - business `use-ai-to-turn-survey-data-into-a-leadership-email` → `["Critical Thinking Fundamentals", "Generative AI Fundamentals", "Prompt Engineering"]` → no tech keyword → track **business**.
  - interview `ai-readiness-interview-d62` → `["Ai Fundamentals"]`.
  This CONFIRMS the "platform pins the opposite of the annotation's audience wording" note — the label is a NAME
  heuristic over `simulations.skills`, and it lands correctly for both cards.
- **Gate at full page render:** still 0/5 (the completed-member page render on a clean reset-to-seed is iter-05).
  Gate part 2's DATA foundation is landed + live-validated at the SQL level.

## Close — 2026-07-24

**Outcome:** the net-new Directus set-dress seeder landed + registered + tested green + live-validated on demo-2;
the 3 wired sims resolve their real evaluated-skill names and the track heuristic classifies tech/business
correctly.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (full page render pending — iter-05)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (→ iter-04, evidence distribution)
**Decisions:** D8 (directus set-dress as a rext SEEDER — Conn.Exec reaches directus in the shared per-stack postgres — over a snapshot step; single-responsibility + reuses the seeder DAG/test harness), D9 (a `to_regclass` string guard shaped to the recordingConn mock so both test paths — present/absent — are unit-testable), D10 (write a JSON array of NAME strings — parseSkillNames accepts `[]string` or `[{name}]`).
**Routes carried forward:** iter-04 = evidence-distribution join (validation fan-out + verified user_skill_evidences for the completed sim's evaluated node-ids); iter-05 = live reset-to-seed render + measure. Deferred still: participants_filter track-tagging + business-sim per-member session routing (render tik).
**Lessons:** validate a net-new DB-write set-dress against the LIVE stack (the SELECT then the UPDATE) BEFORE baking the SQL into a seeder — it de-risks the whole lane cheaply and, here, live-confirmed the track heuristic mapping the annotation warned was inverted.
