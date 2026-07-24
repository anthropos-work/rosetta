# M250 — Spec notes

Iterative milestone (marquee). Accumulates per-lane / per-iter design notes for AI-readiness fidelity — the
31-skill arithmetic re-derivation, the net-new directus-write set-dress, evidence distribution, manager fidelity,
and the believability render loop — recorded during the iter loop.

**Iteration protocol:** `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/verification.md`
(measure→triage→fix→re-render); contract `corpus/services/ai-readiness.md`.

## Arithmetic spine (8→31 re-derivation, ~200 members)
_(TBD during build.)_

## Directus set-dress (net-new: 2 track-keyed named sims + interview + evaluated-skills)
_(TBD during build.)_

## Evidence distribution (validation fan-out + user_skill_evidences)
_(TBD during build.)_

## Manager-vantage fidelity
_(TBD during build.)_

## Believability render loop (0 invented / 0 prod-ejects / closure green / frozen-vs-live agree)
_(TBD during build.)_

## Pre-flight audits — iter-01
**Phase 0b KB-fidelity — verdict: GREEN** (inline check; ran against platform source directly rather than
spawning the heavyweight `audit-kb-fidelity` skill — the strategy is authored against `stack-demo/app/internal/aireadiness/defaults.go`, the authority). The M247-refreshed `corpus/services/ai-readiness.md`
already documents the platform truth — **31 default skills (19 core @1.0 + 12 enabling @0.5)** +
**3 default sims** (2 track-keyed + 1 shared interview) at lines 231-238 — and explicitly names M250 as the
gap-closer; its "Seeding contract (demo/M51)" sections (lines 323, 453-454) accurately describe TODAY's
8-skill / 6.5-denominator demo seeder (the state M250 replaces). No KB drift to code against; the doc will be
updated to the 31-skill contract at close (Delivers).

## Platform contract (source of truth — read iter-01)
- **31 default skills** — `defaults.go` `defaultReadinessSkills`: 19 core @1.0 + 12 enabling @0.5. Node-ids
  captured verbatim (see iter-01 decisions).
- **3 default sims** — `defaults.go` `defaultReadinessSims`:
  - `{StepSimulation, "tech", "634b9ffd-a6a8-444a-a585-1867c1dc61f4"}`  (slug `who-can-see-this-document-fc0`)
  - `{StepSimulation, "business", "a4113c6b-6c1e-4ad0-afe6-9e7eff6f76b4"}` (slug `use-ai-to-turn-survey-data-into-a-leadership-email`)
  - `{StepInterview, "both", "6d6cdf39-e043-4f94-8a5c-e97116bfe1b2"}` (slug `ai-readiness-interview-d62`)
  - `sim_ref` = `directus.simulations.id`. **UNIQUE(org, step_type, track)** → the tech+business pair (both
    `step_type=simulation`) MUST carry DISTINCT tracks. `ai_readiness_sims` ent schema has a `track` enum
    (tech/business/both, default 'both') the CURRENT seeder does NOT write — must add it.
- **Step-1 read** (`how_we_measure.go:computeSkillInsights`) — joins `user_skill_evidences` to the org's
  configured `ai_readiness_skills.node_id`; splits core (weight≥1.0) / enabling. Renders top-N strengths/gaps.
- **Step-2 evaluated-skills read** (`how_we_measure.go:computeSimAssessments`, ~L404) — parses skill NAMES
  from `directus.simulations.skills::text` (LEFT JOIN `directus.simulations` on `id::text = ars.sim_ref`).
  **`directus.simulations.skills` is NULL in capture** → the net-new set-dress must populate it for the 3 uuids.
  **Track shown is a NAME HEURISTIC** (`techTrackRe` over the evaluated skill names), NOT the `track` column —
  so the "tech/business" label the page shows is driven by the skill NAMES set-dressed into `simulations.skills`
  (this is the "platform pins the opposite of the annotation's audience wording" — confirm at live render).
- **Score scales** — step1 30 / step2 40 / step3 30. `computeTier1` = round(heldWeight/totalWeight × 30).
  New totalWeight = 19×1.0 + 12×0.5 = **25.0** (was 6.5). "Champion" full-repertoire hero still = 30/30.

## demo-2 (the single work stack)
Up + running (docker `demo-2-*`). Offset **20000** (web :23000, cockpit :27700). billion untouched.
