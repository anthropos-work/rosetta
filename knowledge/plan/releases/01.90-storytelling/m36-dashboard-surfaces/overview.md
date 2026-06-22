---
milestone: M36
slug: dashboard-surfaces
version: v1.9 "storytelling"
milestone_shape: section
status: planned
created: 2026-06-22
last_updated: 2026-06-22
complexity: large
delivers: rosetta-extensions/stack-seeding (membership_skills + tags/membership_tags + org/user target_roles + succession feeders + job_simulation_feedbacks + the assignments fix + the org-scale distributions) + corpus (extend stories-spec.md with the dashboard surfaces + seeding-spec.md)
depends_on: M35
spec_ref: .agentspace/seeding_gaps.md §3b (dashboard), §3c (the claimed-vs-verified aha), §6a #5
---

# M36 — Dashboard surfaces (Must #2)

## Goal
The org **Workforce-Intelligence dashboard** (REST `/api/workforce/*`) renders **believably** for a seeded
story — the claimed-vs-verified gap, verification funnel, role-readiness, growth, AI-readiness, succession,
and assignments are all **non-empty and distributed**, not binary or zero. This delivers the second product
Must (the dashboard) on top of M34's profile (Must #1).

## Why section
The dashboard's spine tables + the aggregates that need distributions are enumerated and code-traced
(`.agentspace/seeding_gaps.md` §3b). Fixed checklist.

## Scope
**In:**
- **`membership_skills`** (the *mapped* surface) — seeded to **outnumber** verified per skill, so the
  mapped→verified **funnel** shows a believable drop-off. (Set `skill_name` — every query filters it NOT NULL.)
- **`tags` + `membership_tags`** (teams/business-units; incl. a `mentor` tag) — the universal slice dimension
  every team rollup + AI-readiness-by-team depends on.
- **`organization_target_roles` + `user_target_roles`** — the gap-visualization + two-way internal-mobility
  match (both sides needed).
- **Succession feeders** — `validation_attempt_*` + `interview_extraction_results`, sized to clear the
  coverage-confidence gate (so the Succession tab renders, not a CTA).
- **`job_simulation_feedbacks`** (~2:1 positive, Italgas-shaped); **assignments fix** (status mix + past
  `due_date`s + `organization_assignment_sessions` with progress); verify `skillpath_sessions` writes a
  `completed` share.
- **The org-scale distributions** — the **claimed-vs-verified gap** (`user_level` vs `anthropos_level`, a mix
  of over/under-claimers across the population), AI-readiness skills (members with AI-named skills), and the
  growth arc (early-low/late-high) so Growth/Trends/Biggest-Improvers narrate upskilling. The two employee
  heroes surface as the dashboard's standout high/low rows (the coherence property from M35).

**Out:** the cockpit (M37/M38).

## Repo split
- **`rosetta-extensions`** `stack-seeding`: the new dashboard-feeding seeders + the distribution logic.
- **`rosetta`** corpus: extend `stories-spec.md` (dashboard surfaces) + update `seeding-spec.md`.

## Open questions
- **O5 (resolved)** — the self-eval widget diffs `user_level` vs `anthropos_level` (not manager/peers); both
  must be set and differ.

## Risk (scope)
This milestone is the **most likely to grow** (the dashboard has many widgets). **Hard line:** seed the
**spine** for the seeded story (mapped+verified+gap+roles+assignments+feedback); do **not** chase every widget.

## Done-when
The Workforce dashboard renders the seeded story's KPI strip + Skills-Verification (self-eval gap) +
verification funnel + role-readiness + a non-empty Succession tab + populated assignment cards; the two
employee heroes appear as the top-performer / at-risk rows; tests green.
