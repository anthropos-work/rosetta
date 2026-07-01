---
iter: iter-03
milestone: M51
iteration_type: tik
status: closed-fixed-partial
created: 2026-06-30
---

# iter-03 — tik (strand 2: the `ai_readiness_*` config seeder — active cycle + skills + sims + steps)

## Type
tik — under TOK-01 (active-cycle signals-true). Strand 2 of the 4-strand plan. coverage-protocol Phase A–E.

## Step 0 — Re-survey before targeting
Re-confirmed against the iter-02 baseline + the live schema: the 3rd org exists + the enablement gate is ON, but
`ai_readiness_cycles/steps/skills/sims` are all empty for Northwind (the dashboard funnel reads them and is empty —
2 of the 6 baseline failures are the `/enterprise/workforce` AI-readiness sections). TOK-01's strand-2 target
(the config) is untouched + meaningful. No substitution.

## Active strategy reference
**TOK-01** (`../decisions.md`) — active-cycle signals-true. Strand 2: the per-org `ai_readiness_*` CONFIG (cycle +
skills + sims + steps). This is the prerequisite for strand 3 (the per-member funnel reads the config: Step-1 scores
against `ai_readiness_skills`, Steps 2/3 against `ai_readiness_sims`).

## Cluster / target identified
The dashboard's active-cycle recompute path (`buildLiveResponse` → `computeOrgBreakdowns`, contract claim 5) needs
the org config to exist before any member signal can score. Strand 2 writes that config so strand 3's signals have
something to score against, and so the dashboard's "how we measure" / skill-strengths / sims sections render.

## Hypothesis
A net-new `AIReadinessConfigSeeder` (depends-on `org` + `taxonomy` + `content`) writing, per `narrative: ai-readiness`
story: `ai_readiness_cycles` ×1 (`status='active'`, a 5-month window matching the story activity), `ai_readiness_steps`
×3 (skill_mapping/simulation/interview, positions 0/1/2), `ai_readiness_skills` ~5 core (weight 1.0) + ~3 enabling
(0.5) drawn from REAL AI-named public taxonomy node-ids (`resolveNamedSkillRefs` + `filterAISkills`; never fabricated
— closure gate), `ai_readiness_sims` ×2 (`step_type` simulation + interview, `sim_ref` = real replayed Directus sim
ids via `resolveContentRefs().sims`). Re-seed demo-1 → the config exists; the dashboard's AI-readiness config-derived
sections render (strands 2 partially clears the `/enterprise/workforce` AI-readiness sections; the member FUNNEL
NUMBERS still need strand-3 signals).

## Expected lift
Config existence is a prerequisite, not directly a big `(failing)` drop — the funnel COUNTS need strand-3 signals.
Expected: the AI-readiness config-derived sections stop erroring/empty-from-missing-config; the headline funnel may
still read 0% until strand 3. A net reduction in failing sections is the success signal, but a "config landed, funnel
awaits signals" partial is acceptable (close-fixed-partial, routing the funnel to iter-04).

## Phase plan
- Phase A (sweep): inherit the iter-02 baseline `(6,0)` as the pre-iter metric (no new sweep needed pre-fix; the
  config doesn't exist yet — re-running would reproduce `(6,0)`).
- Phase B (triage): the 2 `/enterprise/workforce` AI-readiness sections are config-gated; the other 4 (members/
  assignments/activity-dashboard) are population-data/perf (strand-3 / iter-04). This iter targets the config cluster.
- Phase C (fix): author `seeders/ai_readiness_config.go` (`AIReadinessConfigSeeder`); register it; unit test; build/
  vet/gofmt/test; re-seed demo-1; verify the 4 config tables for Northwind in the DB.
- Phase D (re-sweep): manager-vantage sweep on demo-1 → record `(failing, escapes)` delta.
- Phase E (close): grade on whether the config-gated sections cleared; route the funnel + remaining surfaces forward.

## Escalation conditions
- If `ai_readiness_skills` cannot resolve any real AI-named public node-id (the replayed taxonomy lacks AI skills) →
  fall back to top public skills with an AI-bias (still real node-ids, never fabricated); if even the flat pool is
  empty (no taxonomy replayed) the seeder writes 0 skills and logs it — NOT a fabrication, route the taxonomy gap.
- If the dashboard needs a platform edit to read the seeded config → re-scope-trigger (not expected; the contract
  confirms the read path).

## Acceptable close-no-lift outcomes
If the config lands cleanly + verifies in the DB but the sweep `(failing)` doesn't drop (because every AI-readiness
section also needs strand-3 member signals to render non-empty) → close-fixed-partial: the config (the planned
deliverable) landed, the funnel-signal lift routes to iter-04. Falsification: "config alone doesn't clear the
sections; they're signal-gated too" is a valid, documented strand-2 outcome.
