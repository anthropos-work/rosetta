# iter-03 — decisions (local)

**D1 — Absorb strand 3 (the funnel) into iter-03 (Fate-1), not a separate iter-04.**
- Context: a near-complete strand-3 funnel seeder (`ai_readiness_funnel.go`) appeared in the rext working tree
  mid-iter (co-authored, in the same voice, reusing strand-2 symbols), initially breaking the build (one missing
  helper `aiReadinessStories`). iter-02's close had routed strand 3 forward to iter-04.
- Options: (a) treat it as foreign uncommitted state, leave it broken, exit user-blocker; (b) revert/delete it
  (FORBIDDEN by the TOP-OF-PROMPT BAN); (c) finish + reconcile it in this iter (absorb strand 3 into iter-03).
- Choice: (c). The work is squarely in-scope for M51, complete-able now, and the alternative (a broken tree or a
  forbidden delete) is unacceptable. The build became green once the co-author added `aiReadinessStories` + the
  registration; I reconciled the funnel's skill-pool read to share the strand-2 fix (D2). Strands 2+3 land together.

**D2 — A DEDICATED AI-skill-pool reader (`readAIReadinessSkillPool`) over filtering the capped flat pool.**
- Context: the first config seed wrote only 1 `ai_readiness_skill` (`K-ABITOI-3F06`). Root cause: `filterAISkills(
  named.flat)` filters the generic flat pool, which is capped at `maxSkillPool=500` ordered by node_id — and that
  alphabetical slice contains only 1 of the taxonomy's 253 AI-named public skills.
- Choice: a dedicated reader that applies the AI name-patterns (the shared `aiSkillPatterns`) IN the DB (ILIKE
  disjunction), public-only, capped at 50, ordered by node_id — surfacing the real AI-named skills regardless of the
  flat pool's ordering. Closure-gate-safe (real public node-ids only, never fabricated). BOTH the config seeder AND
  the funnel's Step-1 evidence draw from this SAME reader, so a member's evidence node_id is GUARANTEED to be a
  configured `ai_readiness_skill` (the dashboard's Step-1 presence scan then finds it). Result: 8 skills land
  (5 core weight 1.0 + 3 enabling 0.5).
- Why not bump `maxSkillPool`: that's a generic-pool global with other consumers; widening it to fix an AI-specific
  starvation is the wrong lever. A purpose-built reader is the targeted fix.

**D3 — `ai_readiness_user_step_progresses` is the REAL (ent-pluralized) table name.**
- The M48 contract doc wrote the singular `ai_readiness_user_step_progress`; the live demo-1 schema is the plural
  `ai_readiness_user_step_progresses` (ent pluralization), with cols `id/created_at/updated_at/step_type/status/
  completed_at/organization_id/user_id`, UNIQUE(org,user,step_type), FK user_id→users(id). The funnel was written
  against the real schema (verified). (Route a one-line contract-doc name-fix to the milestone close / a doc tik.)

**D4 — Active-cycle window: start = now − activity months, end = now + 1 month (in-flight).**
- The cycle is `status='active'` (TOK-01). The window starts at the activity span's beginning and ends a month in
  the FUTURE so the assessment reads as still-open (a started hero only exists mid-cycle). The CHECK
  `end_date >= start_date` holds. demo-1: 2026-01-31 → 2026-07-30.

**D5 — ~80% all-3 + a believable drop-off (not binary).** aiReadinessCompletedShare=0.80 reaches stage 3; the
  residual splits ~half stage 2 / ~half stage 1. demo-1 realized 156/199 stage-3 (~78%), 21 stage-2, 22 stage-1 —
  a realistic in-flight funnel. Heroes override: Aria(thriving)→stage3, Ben(struggling)→stage1, Dana(manager)→excluded.
