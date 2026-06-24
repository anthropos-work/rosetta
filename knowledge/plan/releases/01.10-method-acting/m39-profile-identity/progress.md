# M39 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist
- [x] **G1 — org name** — thread `st.Org.Name`/slug through the roster (`stack-seeding/seeders/roster.go`: `RosterIdentity` + `BuildRoster`) → clerkenstein `registry.go` (`RosterEntry`) + `resources.go` (`DemoUser`, `orgMemberships()` l.227); FAPI org resource carries the story org name; `"Clerkenstein Demo Org"` stays the no-roster default; paired `DisallowUnknownFields` struct change + re-tag BOTH repos. **DONE** (commit fb9e300): + single-sourced `orgSlugFor`; both modules green; multi-identity + JS/FAPI alignment 9/9 / 100%.
- [x] **G2 — role backfill** — backfill `public.user_basic_info.job_role_id` (+ `job_title`/`summary`/`location`) from the resolved hero role in `stack-seeding/seeders/users.go`; one UPDATE lights the `/profile` header (`infoResolver.JobRole`) + role-gap radar / role-readiness widgets (`jobRoleMatch`). **DONE** (commit 010f422): idempotent UPDATE (IS-DISTINCT-FROM guard, validated live on demo-3); every member backfilled; no-fabrication preserved; seeders suite green.
- [x] **G4 — real-face avatars** — replace the DiceBear initials SVG (`users.picture`) with a bundled, license-clean real-face set mapped deterministically by hero key (offline-safe + deterministic + license-clean). **DONE** (commit fc8a841): a self-authored parametric SVG **face generator** → offline base64 data URI (no fetch, ~1 KB, license-clean — no vendored asset at all); visually verified (clean varied faces); stdlib-only.
