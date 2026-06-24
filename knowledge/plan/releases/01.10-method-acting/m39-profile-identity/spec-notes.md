# M39 Spec Notes

Technical notes accumulate here during build. The authoritative review is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review, 2026-06-24;
root-cause workflow w7t4wq2z4). Hard facts confirmed in review `w7t4wq2z4`.

## Pre-flight audits — G1 org name

- **Phase 0b KB-fidelity:** **GREEN** (report: `kb-fidelity-audit.md`). All three topics PAIRED with ALIGNED
  claims describing current behavior; the behavior changes + matching doc updates are M39's own deliverables.
- **`public.user_basic_info` real schema** (introspected live on `demo-3-postgresql-1`): columns include
  `id` (uuid, = users.id, FK), `location` (text), `summary` (text), `email` (varchar NOT NULL UNIQUE),
  `job_role_id` (varchar), `job_title` (varchar). **NO `job_role_title` column** (the spec-ref's name was
  approximate — the header reads `job_role_id` → resolved label + `job_title`). The row PRE-EXISTS (created by
  the `users` AFTER-INSERT trigger `init_user_tables()`), so **G2 is an UPDATE keyed by `id`**, not an insert.
- **Maya's live state (confirms the gap):** `user_basic_info` row exists but `job_role_id`/`job_title`/
  `location`/`summary` are all NULL; her `memberships` row carries `job_role_id='J-BACKEN-A9ED'`,
  `job_role_name='Backend Developer'`. G2 copies the same resolved role into `user_basic_info`.
- **`ResolvedStory.Org`** is an `OrgSpec{Name, Slug, …}` — so `st.Org.Name` / `st.Org.Slug` are in hand inside
  `BuildRoster` (roster.go) and the `users.go` Seed loop. `OrgSeeder` writes the same `st.Org.Name` to
  `organizations.name`, so the roster-carried name and the seeded org agree by construction.

## G1 — org name threading

### rext `stack-seeding/seeders/roster.go` — `RosterIdentity` + `BuildRoster`
TODO: carry `st.Org.Name`/slug on the roster identity so the FAPI org resource can surface the story org name.

### clerkenstein `clerk-frontend/registry.go` — `RosterEntry`
TODO: add the org-name/slug field to `RosterEntry` (the paired half of the roster JSON; `DisallowUnknownFields`).

### clerkenstein `clerk-frontend/resources.go` — `DemoUser` + `orgMemberships()` (l.227)
TODO: replace the hardcoded `"Clerkenstein Demo Org"` / `"clerkenstein-demo"` with the roster-carried org
name/slug; keep `"Clerkenstein Demo Org"` as the no-roster default.

### Paired change + re-tag
TODO: roster ↔ `RosterEntry` use `DisallowUnknownFields` → add the field to BOTH structs in one paired change,
then re-tag BOTH repos (rext + clerkenstein) together.

## G2 — role backfill

### rext `stack-seeding/seeders/users.go` — `public.user_basic_info` UPDATE
TODO: backfill `job_role_id` (+ `job_title` / `summary` / `location`) from the resolved hero role. The seeder
currently writes the role only to `public.memberships` (the wrong table for the header). One UPDATE lights the
`/profile` header (`infoResolver.JobRole` ← `user_basic_info.job_role_id`) AND the role-gap radar /
role-readiness widgets (`jobRoleMatch`).

## G4 — real-face avatars

### rext `stack-seeding/seeders/users.go` — `users.picture`
TODO: replace the DiceBear initials SVG with a bundled, license-clean real-face set, mapped deterministically
by hero key. Offline-safe + deterministic + license-clean. (Avatar source mechanism — bundled curated set is
the chosen default; confirm at build.)
