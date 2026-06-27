---
milestone: M41
slug: profile-depth
version: v1.10 "method acting"
milestone_shape: section
status: archived
created: 2026-06-24
last_updated: 2026-06-25
complexity: medium
delivers: corpus/ops/seeding-spec.md + corpus/ops/demo/stories-spec.md — the profile-depth layer (the new ProfileSeeder's user_experiences/user_educations fan-out + the bumped verified-skill depth + the claimed-but-unverified user_skills tail that widens the visible claimed-vs-verified gap)
depends_on: M39 (shares the rext stack-seeding tree; M39's users.go edits land first)
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M41 — Profile depth seeding

## Goal
A seeded hero (Maya Chen) gets a **believable work history + education + a deep, role-aligned skill set with a
real claimed-vs-verified gap**. Today the `/profile` timeline is empty (no `user_experiences`,
no `user_educations` — both tables are **0 rows DB-wide**, written by no current seeder) and the skill set is
shallow (24 `user_skills` / 8 evidences from preset `verified:8`). M41 fills the work/education timeline and
bumps the verified-skill depth, then seeds a claimed-but-unverified tail so the demo's headline aha — the
visible gap between *claimed* and *verified* skills — is wide and obvious when you log in as a hero on demo-3.

This is a tooling + docs milestone: **zero platform-repo edits** (next-web / app / cms / jobsimulation are
read-only). The `/profile` timeline reads `ent.UserExperience` / `ent.UserEducation` via `GET_TIMELINE`
unchanged — M41 only supplies the rows.

## Scope
**In:**
- **G3 — a new `ProfileSeeder`** (rext `stack-seeding`) writing **2–3 `public.user_experiences` + 1–2
  `public.user_educations` per hero**:
  - `user_experiences` columns: `[id, created_at, "user"(uuid FK), company(nullable), title, description,
    from, to, skills(json), location, location_type(enum), job_role_id]`.
  - `user_educations` columns: `[id, created_at, "user", title, from, to, skills(json), institute,
    field_of_study]`.
  - Deterministic UUIDs (the `deterministicUUID` prefix pattern), rows **backdated within the story's
    Activity span**, **role-aligned titles** (reuse the already-resolved `jobRoleRefs`), and **tied to the
    verified-skill arc** so the history corroborates the skills.
  - Neither table is written by any current seeder (confirmed 0 rows DB-wide) — this is net-new write surface.
- **G5 — bump skill depth + seed a claimed-but-unverified tail:**
  - Bump the preset `verified` count (e.g. `8 → ~30`) in `stories.seed.yaml:48` + `stories-maya.seed.yaml:36`.
    The chain is `blueprint.go:227 EffectiveVerified → persona.go:132-178 resolveHeroSkills` ×
    `verifiedSessionsPerSkill=3` (`persona.go:46`), so `~30` distinct verified skills (10 role-coherent +
    auto-top-up from the flat public pool past the role's first 10) × 3 ⇒ **~90 `user_skills` + ~30
    evidences** on the verified side.
  - **PLUS** seed a claimed-but-unverified `user_skills` / `user_skill_evidences` **tail (~60)**:
    `is_verified=false`, **no** `job_simulation_id`, `user_level` set, `anthropos_level` NULL — so the profile
    reads **"overall ≈ 90 = ~30 verified + ~60 claimed-unverified"**, **widening** the visible
    claimed-vs-verified gap (the demo's headline aha).
  - **Keep the gap mechanic intact** — `user_level` (claimed) vs `anthropos_level` (verified) is the widget's
    spine; the unverified tail leaves `anthropos_level` NULL so the gap renders.

**Out:**
- The surfacing fix (M40 — directus serve grant).
- Identity (M39 — profile-identity; its `users.go` edits land first as the `depends_on`).
- The preset's `mapped:22` wiring — that feeds `membership_skills` (the org funnel), **not** `user_skills`;
  leave it as-is. Profile depth comes from `user_skills`, not the mapped count.

## Depends on / Parallel with
- **Depends on M39** (profile-identity) — M41 shares the rext `stack-seeding` tree with M39, and M39's
  `users.go` edits (the hero's real name / avatar / org-domain email) must land first so the ProfileSeeder's
  rows attach to a well-formed hero identity.
- **Parallel with M40** (directus-serve-grant) — M40 is the surfacing/grant fix; it touches a different
  surface and can proceed concurrently.

## Open questions
- **Confirm the "~90 overall" reading** = ~30 verified + ~60 claimed-unverified (the chosen
  narrative-stronger split — a wide, obvious gap over a uniformly-verified profile). This is the assumed
  default; lock it before build.
- **Decide N work/education rows per hero** — within the 2–3 experiences / 1–2 educations envelope, pick the
  per-hero counts.
- **Decide per-population vs heroes-only** — does the ProfileSeeder write experiences/educations for the
  whole seeded population, or heroes only?

## KB dependencies
M41 reads these corpus docs as contract (it must not contradict them; it extends them):
- `corpus/ops/seeding-spec.md` — the `stack.seed.yaml` blueprint, the production-isolation boundary
  (write-side), and the data-DNA. The new write surfaces (`user_experiences` / `user_educations` + the
  unverified `user_skills` tail) must stay within the seeding isolation contract.
- `corpus/ops/demo/stories-spec.md` — the verified-skill-chain **7-table reference**: the
  `EffectiveVerified → resolveHeroSkills` flow, the `verifiedSessionsPerSkill` multiplier, the
  `user_skills` / `user_skill_evidences` columns + constraint landmines, and the `user_level` vs
  `anthropos_level` gap mechanic. M41's depth bump + unverified tail extend this chain, not replace it.

## Delivers →
- `corpus/ops/seeding-spec.md` — the **profile-depth layer**: the `ProfileSeeder` (work-history + education
  fan-out, deterministic UUIDs, backdated within the Activity span, role-aligned titles tied to the verified
  arc) + the claimed-but-unverified `user_skills` tail and the widened claimed-vs-verified gap it produces.
- `corpus/ops/demo/stories-spec.md` — the depth half: the bumped `verified` count + the `× verifiedSessionsPerSkill`
  arithmetic, and how the unverified tail (`is_verified=false`, no `job_simulation_id`, `user_level` set,
  `anthropos_level` NULL) layers on top of the existing 7-table chain.
