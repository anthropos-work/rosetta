# M41 Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions live in the spec
([`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md), live-demo review 2026-06-24).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M41-D1 | A NEW `ProfileSeeder` (surface `"profiles"`, `profile.go`+`profile_write.go`), not an extension of `UsersSeeder` | The timeline + the claimed tail are a distinct surface with their own FK ordering (companies → experiences → claimed user_skills) and their own taxonomy reads; a separate seeder keeps `users.go` focused and registers cleanly behind `personas` in the DAG. | 2026-06-24 |
| M41-D2 | `user_experiences.company` is seeded as a real `public.companies` row (NOT NULL FK), one per distinct (name,domain) | The live demo-3 schema has `company uuid NOT NULL` FK→companies (the overview's "company(nullable)" was wrong); the GraphQL `Company` resolver does `QueryCompany().Only(ctx)`, so a NULL/dangling company errors the WHOLE timeline. Dedup by a deterministic id derived from (name,domain) so the same employer is one row (the table's UNIQUE (name,domain) + the idempotent COPY both dedup). | 2026-06-24 |
| M41-D3 | The claimed-but-unverified tail ties to the seeded experiences/educations via `user_skill_experience`/`user_skill_education` | The tail has no `job_simulation_id` (it's unverified), but `user_skills_check_foreign_keys` requires ≥1 provenance edge non-NULL. Tying to the work history is the natural G3↔G5 join AND makes the claimed skills render UNDER each work experience (the `workExperience.Skills` resolver reads `userskill.HasExperienceWith`). | 2026-06-24 |
| M41-D4 | The claimed evidence UPSERT is a SEPARATE SQL from the verified `upsertEvidenceSQL`, with `is_verified=false`, `anthropos_level` left NULL, and an `ON CONFLICT … WHERE is_verified=false` guard | The verified UPSERT hardcodes `is_verified=true` + sets `anthropos_level`; the claimed side must leave `anthropos_level` NULL (so the gap renders) and must NEVER clobber a verified row on a (skill,user) collision (the verified side always wins). The tail draws skills distinct from the verified set, so the guard is a re-run/safety net, not the common path. | 2026-06-24 |
| M41-D5 | `verified:` bumped to ~30 for THRIVING heroes only (Maya 30, Sara 28); struggling heroes stay sparse (2) but still get the ~60-skill claimed tail | The "~90 overall = ~30 verified + ~60 claimed" reading (the prompt's DECIDED G5 answer) is the thriving hero's deep profile; a struggling OVER-claimer's stark gap (2 verified + ~60 claimed) is equally coherent, so the tail applies to both arcs — the gap mechanic carries the narrative either way. | 2026-06-24 |
| M41-D6 | The claimed tail draws from a COMBINED role-coherent-then-flat named pool (`combinedNamedPool`), offsetting past the first `EffectiveVerified()` skills | The role pool alone (≤10 skills) can't supply a ~60-skill tail; the verified chain itself tops up from the flat pool past the role's skills (resolveHeroSkills), so the tail mirrors that draw and skips the verified prefix to keep claimed DISTINCT from verified — the two counts don't overlap. | 2026-06-24 |
| M41-D7 | Profile depth is HEROES-ONLY (not the whole population) | The overview's open question (per-population vs heroes-only) — heroes-only is the safe, believable core a presenter actually clicks into; the population already reads coherent (roles, ramped joins, org-scale evidence from M36). The timeline + tail are a per-hero close-up surface. | 2026-06-24 |

## Adversarial review (close, Phase 2c — 2026-06-25)

Scenarios considered against the just-shipped ProfileSeeder code; each was probed (or already
pinned by the harden pass) and found handled — no production bug surfaced.

- **AR-1 (NEW test) — empty-`eduIDs` round-robin modulo-by-zero.** `seedClaimedTail`'s
  `si%4 == 3` branch ties ~1-in-4 claimed skills to an education via
  `eduIDs[si%len(eduIDs)]` — a modulo by `len(eduIDs)`. Through `Seed`, `seedHeroProfile`
  always seeds `educationsPerHero` (≥1) educations first, so `eduIDs` is never empty in
  practice; the in-seeder safety is the `len(eduIDs) > 0` guard (profile.go:289), not the
  caller. Adversarial input (a non-empty tail + non-empty `expIDs` + **empty** `eduIDs` — the
  hypothetical `educationsPerHero=0` / education-skipped world) must NOT panic and every row
  must fall back to the experience edge (education edge NULL). **Handled** — the guard makes the
  education branch unreachable; verified by `TestSeedClaimedTail_EmptyEducationsNoPanic`. The
  code was already correct; the test pins the environmental assumption.
- **AR-2 (already pinned) — claimed tail never clobbers a verified row.** A (skill_id, user_id)
  collision between a claimed-tail UPSERT and a PersonaSeeder verified row must leave the
  verified row's `is_verified=true`/`anthropos_level` intact (the gap mechanic's safety-critical
  SQL). The `ON CONFLICT … WHERE is_verified=false` guard (profile_write.go:149) handles it;
  pinned by `TestProfileClaimedEvidence_GuardNeverClobbersVerified` (harden pass).
- **AR-3 (already pinned) — partial-UNIQUE under round-robin edge reuse.** With ~60 tail skills
  across 3 experiences / 1 education, many rows reuse the same experience/education id. The
  partial UNIQUEs are on `(skill_id, user_skill_user, user_skill_experience/_education)`; the
  tail draws DISTINCT `skill_id`s (M41-D6), so no two rows collide on the composite key — safe by
  construction. Distinctness pinned by `TestProfileSeeder_TailDistinctFromVerified` +
  `TestProfileSeeder_SmallPoolGracefulTail`.
