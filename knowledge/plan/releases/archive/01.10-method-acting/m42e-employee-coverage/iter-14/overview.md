---
iter: 14
iteration_type: tik
iter_shape: production-fix
status: closed-fixed
---

# iter-14 — P3: per-hero activity completeness

**Type:** tik (production-fix). **Active strategy: TOK-10**. The design-plan's root cause #7: the empty
activity surfaces (/home Paths=0, no path pills; "Skill Paths Completed"=0; saved-for-later empty).

## Cluster / target identified
P3 from TOK-10's next-tik direction. iter-11 B1: Maya had 0 personal_assignments, 0 completed skill-path
sessions, 0 user_bookmarks.

## Hypothesis
Seeding per-hero personal_assignments (skill_path), a guaranteed-completed skill-path session, and
saved bookmarks fills the /home path pills + the Paths count + the "Skill Paths Completed" count + the
saved-for-later surface.

## Phase plan
A (baseline iter-11) → confirm the table schemas + ent enums from the platform (read-only) → C (NEW
HeroActivitySeeder in rext, registered in buildRegistry; tests) → re-seed demo-3 (additive) → D (re-measure:
Maya + all heroes; real refs; closure) → E (close).

## Close — 2026-06-25
**Outcome:** P3 landed. NEW `HeroActivitySeeder` writes per end-user hero 4 personal_assignments (skill_path,
real Directus refs, 3 active + 1 completed), ≥1 completed skill-path session (progress=100), 3 bookmarks
(skill_path + job_simulation mix). Measured on demo-3: all 4 heroes populated (16 assignments, 12 bookmarks,
4 completed sessions); Maya 4 assignments / 1 completed path / 3 bookmarks, all real refs; closure PASS.
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P3 of P0–P8; the believability gate needs P4–P8 + the P7 semantic harness). The activity
ROOT is closed.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** (tik #4… see below) — (6) protocol-stop: n — Outcome: cap context — this is the 4th tik; the session's planned P0–P3 scope is COMPLETE at this iter, a natural phase boundary (protocol-stop on the run's planned scope). EXIT_REASON recorded in the session report as cap/phase-boundary.
**Decisions:** see ./decisions.md (D1 the table schemas + ent enums; D2 the HeroActivitySeeder design — real-ref
resolution + the distinct "hero-completed" version + manager-skip).
**Routes carried forward:** P4 (avatar + org logo), P5 (the FATAL Sentinel-reload reproducibility fix), P6
(library capture-path: sim embeddings + categories), P7 (the semantic harness rebuild), P8 (fresh-demo-up
acceptance) — all later runs.
**Lessons:** confirm the ent enum TOKENS from the platform schema before seeding an enum column — the personal
assignment resource_type is `job_simulation` (NOT `simulation`) and `skill_path`; a wrong token inserts but the
GraphQL enum can't map it. A hero's guaranteed-completed session needs a DISTINCT version token to coexist with
the population seeder's rows for the same (user, path) under the UNIQUE (user_id, skill_path_id, version).
