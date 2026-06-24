# M39 Spec Notes

Technical notes accumulate here during build. The authoritative review is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review, 2026-06-24;
root-cause workflow w7t4wq2z4). Hard facts confirmed in review `w7t4wq2z4`.

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
