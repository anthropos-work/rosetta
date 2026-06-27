---
title: "KB Fidelity Audit — M41 profile-depth"
date: 2026-06-24
scope: milestone:M41
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Verified-skill chain (the spine M41 extends) | `corpus/ops/demo/stories-spec.md` + `corpus/ops/seeding-spec.md §verified-skill chain` | `stack-seeding/seeders/persona.go`, `persona_write.go`, `taxonomyref.go`, `skillref_named.go`, `jobroleref.go` | PAIRED |
| The `verified:` knob → chain arithmetic (G5 bump) | `stories-spec.md §the mechanism` + `seeding-spec.md` | `blueprint.go:227 EffectiveVerified`, `persona.go:46 verifiedSessionsPerSkill`, `persona.go:132 resolveHeroSkills`, presets | PAIRED |
| `user_experiences` / `user_educations` write surface (G3) | — (net-new; milestone `Delivers →` both docs) | `public.user_experiences`/`user_educations` (0 rows; no seeder writes them — read by `app` `resolver_timeline.go`/`TimelineGrouped`) | DOC-ONLY-TO-BE (milestone deliverable) |
| Claimed-vs-verified gap mechanic (G5 tail) | `stories-spec.md §the claimed-vs-verified gap` | `persona_write.go upsertEvidenceSQL` (`user_level` vs `anthropos_level`) | PAIRED |

## Fidelity Findings
1. **Verified-skill chain doc vs code** — ALIGNED. `stories-spec.md`'s 7-table fan-out, the constraint landmines
   (user_skills CHECK + partial UNIQUE, evidence UPSERT, `user_level` requirement), and the `EffectiveVerified →
   resolveHeroSkills × verifiedSessionsPerSkill` arithmetic all match `persona.go`/`persona_write.go` exactly.
2. **Cited line anchors** — ALIGNED. `stories.seed.yaml:48` (`verified: 8`), `stories-maya.seed.yaml:36`,
   `persona.go:46` (`verifiedSessionsPerSkill = 3`), `blueprint.go:227` (`EffectiveVerified`) all resolve to the
   cited lines in the rext authoring copy @ `method-acting-m40`.
3. **`seeding-spec.md` M34/M35/M36 layers vs the seeder fleet** — ALIGNED. The doc's seeder list (PersonaSeeder,
   membership_skills, tags, target_roles, succession, feedback, population_evidence) matches the registry in
   `cmd/stackseed/main.go`.

## Completeness Gaps
- **Overview/spec column-list guesses were stale against the LIVE schema** (caught during recon, before any code).
  These were milestone-planning notes, not knowledge docs, but they are load-bearing for the build, so corrected
  in `spec-notes.md` (the LIVE-SCHEMA CORRECTIONS block): `user_experiences.company` is `uuid NOT NULL` FK (not
  nullable), `from`/`to` are `date` with `from<=to` CHECKs, `location_type` is lowercase `hybrid|fullremote|
  inoffice` (not `INOFFICE/HYBRID`), `skills` is a json array of names. The delivery docs themselves do not yet
  describe these tables (they're net-new M41 surfaces) — that IS the milestone's doc deliverable, not a stale claim.

## Applied Fixes
- `spec-notes.md` — added the LIVE-SCHEMA CORRECTIONS block + the G5-tail provenance-edge landmine (the tail's
  `user_skills` must set `user_skill_experience`/`user_skill_education` since it has no `job_simulation_id`).
  (Applied during the Phase 0 recon, before this audit.)

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed. The two delivery docs are PAIRED + aligned; the new write surfaces are explicit milestone
deliverables (`Delivers →` lines present in `overview.md`); all cited line anchors resolve. No blind area, no
stale load-bearing claim on a topic M41 reads as truth.
