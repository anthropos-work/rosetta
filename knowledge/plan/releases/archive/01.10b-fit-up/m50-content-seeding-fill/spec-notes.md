# M50 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Substrate (iter-01)
- **demo-1 is UP** + seeded with the full Stories world (`presets/stories.seed.yaml`): Cervato Systems
  (`22222222-…`, 221 members) + Solvantis. Heroes present: Maya Chen (`maya.chen1@cervato-systems.com`,
  uid `3192aff1-9766-5aa3-bdb8-5b3feee89d79`), Dan Rossi (`dan.rossi3@cervato-systems.com`, uid
  `40921b2e-4b27-524e-ace5-c4699156bad9`), Tom Becker, etc.
- **demo-1 consumes rext `fit-up-m49`** (`stack-demo/rosetta-extensions` @ `1035efd`).
- Offset +10000: next-web `:13000`, backend `:18081/18082/18083`, directus `:18055`, postgres **`:15432`**
  (container `demo-1-postgresql-1`, single DB `postgres`, multi-schema: public/jobsimulation/skillpath/
  skiller/sentinel/cms/directus), cockpit `:17700`, fapi `:15400`, studio-desk `:19000`.
- **DB access for diagnosis:** `docker exec demo-1-postgresql-1 psql -U postgres -d postgres -tAc "…"`.
- **Coverage harness:** rext `stack-verify/e2e/` (deps + chromium installed in authoring copy). Runner
  `./run-coverage.sh <N> <employee|manager>`; manager floors calibrated against demo-3 (~221 members).

## Live-schema column landmines (verified on demo-1 — differ from the specs' `_id` guesses)
- `public.memberships` PK-org column is **`organization`** (NOT `organization_id`); user FK is **`user`**;
  has `location`, `joined_at`, `last_activity_date`, `created_at` (ramped Jan→Jun), `lastname`.
- `public.user_experiences` / `user_educations` / `user_certifications` / `user_projects` user FK = **`"user"`**.
- `organization_assignments` org column = **`organization_id`** (yes the `_id` here).
- Language tables: `public.world_languages` (id uuid / code / name — a **reference/lookup**, currently
  **0 rows**), `public.membership_languages` (FK → memberships + world_languages, level), `public.user_languages`
  (FK → users + world_languages, level). All three **0 rows**.
- XP: `public.user_experience_points` (experience_points/reason/user_id/entity_id) — **0 rows DB-wide**.

## Pre-flight audits — iter-01
- **`/developer-kit:audit-kb-fidelity --milestone=M50` → YELLOW.** Docs accurately describe the built
  M34–M46 seeder fleet (no stale claims that mislead); every M50 fill item lands in a blind area (expected
  for a fill milestone → known-context, not blocker). Two must-know framings carried into TOK-01:
  - **F1 — gate-MET reconciliation.** `coverage-protocol.md` calls the manager gate MET (M42m iter-04,
    failingSections=0). M50's re-diagnosis shows genuine empties (languages/last-activity/etc.). Reconcile:
    M50 is FILLING sections the manager manifest never asserted (or that regressed with the larger seed),
    not regressing a green gate — the first manager sweep is the new baseline, not a regression.
  - **F2 — `HeroActivitySeeder` (`hero_activity.go`) ALREADY exists** (M42e P3): writes `personal_assignments`,
    a completed `skill_path_sessions`, `user_bookmarks` — the `/home` + `/profile/activities` content. Do NOT
    duplicate; EXTEND it (it does NOT write `user_experience_points` (XP) nor mirror to
    `public.local_skill_path_sessions` — Maya has 2 skillpath sessions but 0 app-mirror rows).
  - F3 no spoken-languages seeder (confirmed). F4 `CertificatesSeeder` is hero-only → Talent cert aggregate
    is roster-sparse. F5 `location` seeded at `user_basic_info` but `/enterprise/members` likely reads a
    DIFFERENT column (the M44 §D `memberships.picture_url` trap class) — verify; `last_activity` net-new.
    F6 academy entirely unwired (no cockpit route, no session). F7 AI-keys policy unmade (belongs in
    `secrets-spec.md`, outside this contract set). F8 map Growth/Verification tab widgets → backing seeder.

## Re-diagnosis on the FRESH demo-1 (iter-01) — what is GENUINELY empty vs already-rendering
GENUINE seed gaps (need seeder/backfill):
- **Member `location` 0/221, `last_activity_date` 0/221, `joined_at` 0/221** (memberships table) →
  `/enterprise/members` columns empty. (NB `created_at` IS ramped — `joined_at` is the separate display col.)
- **Languages: 0 rows** in `membership_languages` + `user_languages` + `world_languages` (FK target) →
  Talent-tab "no language spoken". New `MemberLanguagesSeeder` + the `world_languages` reference fill.
- **Certifications: 2/221 members, 5 total** (hero-only) → Talent-tab "Certification really low numbers".
- **Maya XP = 0** (`user_experience_points`), **Maya `local_skill_path_sessions` = 0** (app mirror;
  `/profile/activities` xp + skill-paths-completed read the mirror, not the `skillpath` schema).

LIKELY NOT seed gaps (data present — empty render is federation/frontend/demo-up-#7-abort artifact, confirm by sweep):
- Library: **22 published `directus.skill_paths`** → `/library/skill-paths` empties are content-federation, not seed.
- **76 `organization_target_roles`** + **114 `organization_assignments`** for Cervato → Workforce
  Growth/gap + `/enterprise/assignments` have backing rows; confirm whether the empty render is a sweep flake,
  a mirror gap, or frontend.
