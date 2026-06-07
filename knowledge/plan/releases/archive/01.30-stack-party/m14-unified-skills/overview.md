---
milestone: M14
slug: unified-skills
version: v1.3 "stack party"
milestone_shape: section
status: archived
created: 2026-06-07
last_updated: 2026-06-07
complexity: large
delivers: the unified stack-* + dev-up/dev-down skills + a rewritten CLAUDE.md skill table + refreshed corpus/ops/ guides
---

# M14 — Unified `stack-*` skills + `dev-up`/`dev-down`

## Goal
One coherent stack-operations skill set that works on any stack (`dev-N | demo-N`), with the dev lifecycle mirroring
demo's.

## Why section
A finite, enumerable rename/consolidation + reference sweep. No emergent path.

## Scope
- **In:**
  - **`dev-up`** — consolidate `setup-platform` + `start-platform` into one dev-stack bring-up (drives the M13 flow: spawn local Directus + auto-snapshot + light seed) + **`dev-down`**.
  - **Hard-rename** (no aliases, user 2026-06-07) the operation skills to generic stack-target forms: **`stack-list`** (←`demo-status`), **`stack-seed`** (←`demo-seed`), **`stack-snapshot`** (←`demo-snapshot`), **`stack-update`** (←`update-platform`) — each accepting a `dev-N|demo-N` target.
  - **Remove** the old skill dirs (`setup-platform`, `start-platform`, `update-platform`, `demo-status`, `demo-seed`, `demo-snapshot`).
  - Update **every reference**: the CLAUDE.md skill table, the root READMEs, the `corpus/ops/` guides + the `demo/` recipes.
  - `demo-up`/`demo-down` stay as the demo lifecycle (aligned with `dev-up`/`dev-down`).
- **Out:** the safety doc (M15).

## Depends on
**M12 + M13** (the generic skills drive the registry + the dev peer capabilities). **Parallel with:** none (gates M15).

## Open questions (resolve during build)
- How much of `setup-platform`'s first-time machine setup (tool install, org clone) folds into `dev-up` vs stays a one-time prerequisite.
- Target-detection UX (`stack-seed dev-1` vs `stack-seed --stack dev-1`).
- Whether `demo-up`/`demo-down` also gain a generic `stack-up`/`stack-down` (lean: no — lifecycle stays type-specific).

## KB dependencies (read as contract)
- the existing skill `SKILL.md` files (setup-platform, start-platform, update-platform, demo-status/seed/snapshot)
- `corpus/ops/*` guides (setup / run / update / demo) + `CLAUDE.md` (the skill table)

## Delivers → the unified skills + CLAUDE.md + corpus/ops/
- The new `stack-list` / `stack-seed` / `stack-snapshot` / `stack-update` + `dev-up` / `dev-down` skills; the old dirs removed.
- A rewritten `CLAUDE.md` skill table + refreshed `corpus/ops/` guides (the converged stack model).
