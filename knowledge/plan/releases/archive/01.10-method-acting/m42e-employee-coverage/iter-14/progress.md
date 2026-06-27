# iter-14 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. P3 — per-hero activity completeness.

## Phase C — fix (rext, zero platform edit)
- Confirmed the activity table schemas + ent enums from the platform (read-only; D1) — resource_type is
  `skill_path` | `job_simulation`; status `active`|`completed`; the skill_path_sessions UNIQUE on
  (user, path, version).
- NEW `hero_activity.go` HeroActivitySeeder (registered in buildRegistry; D2): per end-user hero — 4
  personal_assignments (skill_path, real Directus refs, 3 active + 1 completed), ≥1 completed skill_path_session
  (progress=100, version 'hero-completed'), 3 user_bookmarks (skill_path + job_simulation mix). Managers skipped.
- Tests: surface/deps/isolation, writes-activity, manager-only-writes-nothing. `go test ./...` GREEN; vet clean.

## Phase D — re-measure (demo-3 re-seed, additive)
- `hero-activity rows=32` (4 heroes × 8). All 4 end-user heroes populated: 16 assignments, 12 bookmarks, 4
  hero-completed sessions.
- Maya: 4 personal_assignments (3 active + 1 completed, all 4 REAL Directus skill-path refs), 1 completed
  skill-path session (progress=100), 3 bookmarks (2 real skill_path + 1 real job_simulation).
- `datadna measure-closure --stack demo-3`: **PASS**. Isolation: clean (prod=false).

## Close — 2026-06-25
**Outcome:** P3 landed. The activity surfaces fill: /home path pills + Paths>0, "Skill Paths Completed">0,
saved-for-later populated — for all 4 end-user heroes, with real Directus refs. Closure PASS.
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P3 of P0–P8; the believability gate needs P4–P8 + the P7 semantic harness)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y (4th tik; the run's planned P0–P3 scope COMPLETE — a natural phase boundary) — (6) protocol-stop: n — Outcome: exit (cap / phase-boundary)
**Decisions:** D1 (schemas + enums), D2 (HeroActivitySeeder) — see ./decisions.md.
**Routes carried forward:** P4 (avatar + org logo), P5 (FATAL Sentinel-reload reproducibility), P6 (library
capture-path), P7 (semantic harness), P8 (fresh-demo-up acceptance) — later runs.
**Lessons:** confirm the ent enum TOKEN from the platform schema before seeding an enum column (job_simulation,
not simulation); a guaranteed-completed hero session needs a DISTINCT version to coexist with the population
seeder under the (user, path, version) UNIQUE.
**rext:** commit `85012ba`, tag `method-acting-m42e-iter14`.
