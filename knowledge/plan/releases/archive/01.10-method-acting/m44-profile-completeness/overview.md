---
milestone: M44
slug: profile-completeness
version: v1.10 "method acting"
milestone_shape: section
status: archived
created: 2026-06-26
last_updated: 2026-06-26
complexity: medium
delivers: corpus/ops/demo/profile-completeness-spec.md (NEW — the 'complete profile' rubric: identity + content + semantic layers, per-vantage member vs manager, each component mapped to its seeding surface + a Playwright acceptance assertion) + updates to seeding-spec.md / stories-spec.md for the new CertificatesSeeder / ProjectsSeeder surfaces, the trajectory-aware self-rating, the manager personal-data unskip, and the bulk-member avatar+career depth
depends_on: none structurally — but land BEFORE M45 (the engine reuses the certificate/project + bulk-member surfaces + the trajectory-aware self-rating). Parallel with M43 (different rext module).
spec_ref: .agentspace/scratch/roadmap-research-2026-06-26.md (the v1.10-extend research note; the profile-density strand)
---

# M44 — Profile completeness (members + managers fully baked)

## Goal
Close the character-facing **DATA-DENSITY** gaps so **every member AND manager profile is fully
populated** — trajectory-aware self-ratings, certificates + projects, manager personal data, and every
`/enterprise/members` fill-member gets an avatar + career + skills. **DATA DENSITY ONLY — zero
platform / next-web edits** (no UI "% complete" widget; the deliverable is rows, not chrome). This is the
believable-profile completeness pass: M41 gave heroes a timeline; M44 makes the **whole roster** — heroes,
managers, and bulk fill-members — read as real people.

## Scope
**In:**
- **(A) Trajectory-aware self-rating** — branch `PersonaSeeder`'s `user_level` UPSERT on
  `Persona.Trajectory` so the **THRIVING** hero shows a **completed** self-assessment and the
  **NON-THRIVING** hero shows an **INCOMPLETE / absent** self-rating state (the user's explicit ask). Keep
  ~2–3 verified skills for the non-thriving hero so the Skill-Spotlight chart still renders (visibly
  sparse, not broken).
- **(B1) NEW `CertificatesSeeder`** (`seeders/certificates.go`, surface `'certificates'`, `DependsOn`
  users+personas) — **2–3 `public.user_certificates` rows per end-user hero**, backdated within the
  activity span. Columns: `cert_name` NOT NULL, DATE types, nullable `organization_id`.
- **(B2) NEW `ProjectsSeeder`** (`seeders/projects.go`, surface `'projects'`, `DependsOn` users+personas)
  — **3–4 `public.user_projects` rows per end-user hero**. Columns: `project_name` NOT NULL, nullable
  `end_date`, `skills` json array.
- **(C) MANAGER personal data** — remove the `IsManager` skips at `persona.go:121` + `profile.go:125`;
  write a **small verified-skill subset** (3–8 skills, L1–L2 band, **flat current-state**, no growth arc)
  + a **manager-track timeline** (3 experiences + 1 education) so the manager's OWN `/profile` is
  populated like a real member.
- **(D) BULK-MEMBER depth** — extend `photoAvatarDataURI` to **EVERY member** (not heroes-only) so
  `/enterprise/members` shows avatars, **+** extend `ProfileSeeder` to write a **shallow timeline** (3
  short-tenure experiences, 1 education, a flat ~6-skill claimed tail) for every member.

**Out:**
- Any **platform / next-web UI edit** — **NO "% complete" widget** (DATA DENSITY only, the user's choice).
- **LLM-generated content** — that's M45/M46; M44's seeders are **deterministic**.
- **Deep per-fill-member career narratives** — kept shallow by design; richer fill is M46
  (LLM-generated).

## Depends on / Parallel with
- **Depends on:** none structurally — but **land BEFORE M45**: the engine reuses M44's certificate /
  project + bulk-member surfaces + the trajectory-aware self-rating.
- **Parallel with:** **M43** (a different rext module — M44 is `stack-seeding/seeders/`; M43 is
  `demo-stack/cockpit.py`).

## Approach (default decisions — flagged ones are Open questions)
- **Non-thriving hero** = hasn't done the initial self-rating (minimal/absent `user_level` on the claimed
  side) but **keep ~2–3 verified skills** so the Skill-Spotlight chart renders (sparse, not broken).
- **Manager** = a **flat current-state** personal profile (3–8 verified L1–L2 skills, no growth arc) + a
  manager-track timeline (3 experiences + 1 education).
- **EVERY member** gets an avatar + a shallow career — the user wants `/enterprise/members` to look full;
  depth stays shallow (3 short-tenure experiences, 1 education, ~6-skill claimed tail).

## Open questions
- Non-thriving hero: **truly-zero verified** vs minimal **2–3** (default: minimal 2–3, so the chart
  renders).
- Manager trajectory: **flat** vs **rising** (default: flat current-state).
- Bulk-member timelines for **EVERY member** (~3K rows/org) vs **sampled** (default: ALL get
  avatar+career, depth shallow).
- **Shared-vs-separate** `ProfileSeeder` path for managers / fill-members vs the hero path.
- The **projects surface name** — avoid collision with an existing surface.

## KB dependencies
M44 reads these corpus docs as contract (it must not contradict them; it extends them):
- `corpus/ops/seeding-spec.md` — the `stack.seed.yaml` blueprint + the production-isolation boundary
  (write-side) + the data-DNA. The new write surfaces (`user_certificates`, `user_projects`, the
  manager personal-data rows, the bulk-member timelines) must stay within the seeding isolation contract.
- `corpus/ops/demo/stories-spec.md` — the verified-skill-chain 7-table reference (the
  `EffectiveVerified → resolveHeroSkills` flow, the `user_level` vs `anthropos_level` gap mechanic, the
  G14 session-seeder fix). M44's trajectory-aware self-rating + manager subset extend this chain.
- `corpus/ops/demo/coverage-protocol.md` — the semantic believability gate (M42e/M42m). Each M44
  component maps to a Playwright acceptance assertion.

## Delivers →
- `corpus/ops/demo/profile-completeness-spec.md` **(NEW)** — the **"complete profile" rubric**:
  **identity + content + semantic** layers, per-vantage **member vs manager**, each component
  (self-rating, certificates, projects, manager personal data, bulk-member avatar+career) mapped to its
  **seeding surface** + a **Playwright acceptance assertion**.
- `corpus/ops/seeding-spec.md` — the new `CertificatesSeeder` / `ProjectsSeeder` surfaces, the
  trajectory-aware `user_level` branch, the manager `IsManager`-unskip, and the bulk-member avatar+career
  extension (all deterministic, prod-isolated).
- `corpus/ops/demo/stories-spec.md` — the trajectory-aware self-rating (thriving completed vs
  non-thriving incomplete/absent) layered onto the existing gap mechanic.
- Reuse (no new mechanics): `resolveJobRoleRefs` / `resolveNamedSkillRefs`, `dateOnly`,
  `legacySkillsJSON`, `roleTitle` degradation, `CopyRowsIdempotent`.
