# iter-14 decisions (P3)

## D1 — the activity table schemas + ent enums (confirmed from the platform, read-only)
- `public.personal_assignments(id, created_at, updated_at, due_date NULLable, status NOT NULL DEFAULT 'active',
  resource_id uuid NOT NULL, resource_type, user_id uuid NOT NULL)`. The ent enums (app/internal/data/ent/enum/
  assignments.go): `AssignmentResourceType` = `job_simulation` | `skill_path` (NOT "simulation"); `AssignmentStatus`
  = `active` | `completed` | `archived`.
- `public.user_bookmarks(id, created_at, resource_id uuid NOT NULL, resource_type, user_id)`; resource_type reuses
  `AssignmentResourceType`; UNIQUE (user_id, resource_id, resource_type).
- `skillpath.skill_path_sessions(... status, version, progress, skill_path_id ...)`; UNIQUE
  (user_id, skill_path_id, version) — a guaranteed-completed hero session needs a DISTINCT version to coexist
  with the SkillpathSessionsSeeder's "1"-version rows for the same (user, path).
All read-only diagnosis; zero platform edit.

## D2 — the HeroActivitySeeder (the P3 fix)
NEW `stack-seeding/seeders/hero_activity.go` (registered in `buildRegistry`, surface `hero-activity`,
DependsOn users+content+skillpath-sessions). Per END-USER hero (managers skipped — they ride the org dashboard):
- **4 personal_assignments** (`resource_type='skill_path'`, real Directus skill-path refs via
  `linkedRefDistinct(refs.skillPaths,…)`, a 3-active + 1-completed mix) → /home path pills + the Paths count.
- **≥1 GUARANTEED completed skill_path_session** (progress=100, status='completed', version='hero-completed'
  so it never collides with the population seeder's "1"-version rows) → "Skill Paths Completed" > 0.
- **3 user_bookmarks** (a skill_path + job_simulation mix, distinct per the UNIQUE index) → saved-for-later.
Refs resolve from the SAME replayed public Directus content pool the fleet uses (free-UUID fallback when no
snapshot) → closure stays green. Idempotent COPY (ON CONFLICT) — re-run-safe. Tests: surface/deps/isolation,
writes-activity (assignment mix + the completed session + the bookmark mix), manager-only-writes-nothing.

## Measurement (live demo-3, additive)
The surfaces were empty for the heroes, so the re-seed is purely ADDITIVE (no clear needed). `hero-activity
rows=32` (4 heroes × 8). Measured: all 4 end-user heroes populated — 16 assignments, 12 bookmarks, 4
hero-completed sessions. Maya: 4 assignments (3 active + 1 completed, all 4 real Directus skill_path refs),
1 completed path (progress=100), 3 bookmarks (2 real skill_path + 1 real job_simulation). `datadna
measure-closure --stack demo-3`: PASS. Re-seed isolation: clean (prod=false). The AUTHORITATIVE clean reproduce
is the P8 fresh demo-up (the seeder produces all of this from scratch there).

## Before → after (Maya, the P3 deltas)
| Surface | before | after |
|---|---|---|
| personal_assignments | 0 | 4 (3 active + 1 completed, real skill-path refs) |
| completed skill_path_sessions | 0 | 1 (progress=100) |
| user_bookmarks | 0 | 3 (2 skill_path + 1 job_simulation) |
