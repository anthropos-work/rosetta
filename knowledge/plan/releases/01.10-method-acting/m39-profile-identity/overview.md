---
milestone: M39
slug: profile-identity
version: v1.10 "method acting"
milestone_shape: section
status: planned
created: 2026-06-24
last_updated: 2026-06-24
complexity: small-medium
delivers: corpus/ops/demo/stories-spec.md (the profile-identity layer — org-name + role-backfill + real-face avatars) + corpus/services/clerkenstein.md (the roster org-name threading: RosterIdentity/BuildRoster → RosterEntry → DemoUser/orgMemberships)
depends_on: none
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M39 — Profile identity & quick wins

## Goal
A logged-in hero (Maya Chen, on demo-3) shows the **right org name**, a **real role + title**, and a
**real face**. These are the three highest-leverage, lowest-effort profile-believability fixes — each lights
a visible surface the live "log in as a hero" acceptance lands on. Tooling + docs only; **zero** platform-repo
edits (next-web / app / cms / jobsimulation are read-only).

## Scope

**In:**
- **G1 — org name.** Thread `st.Org.Name`/slug through the roster so the FAPI org resource carries the
  story org name (the top bar reads **"Cervato Systems"**, not "Clerkenstein Demo Org").
  - rext: `stack-seeding/seeders/roster.go` — `RosterIdentity` + `BuildRoster` carry the org name/slug.
  - clerkenstein: `clerk-frontend/registry.go` (`RosterEntry`) + `resources.go` (`DemoUser`,
    `orgMemberships()` — currently hardcodes `"Clerkenstein Demo Org"` / `"clerkenstein-demo"` at l.227).
  - `"Clerkenstein Demo Org"` stays the **no-roster default**.
  - The roster ↔ `RosterEntry` JSON uses `DisallowUnknownFields`, so **both structs add the field in ONE
    paired change** — and **both repos are re-tagged** together.
- **G2 — role backfill.** Backfill `public.user_basic_info.job_role_id` (+ `job_title` / `summary` /
  `location`) from the resolved hero role in rext `stack-seeding/seeders/users.go`. One UPDATE lights several
  surfaces at once: the `/profile` header (`infoResolver.JobRole` ← `user_basic_info.job_role_id`) **and** the
  role-gap radar / role-readiness widgets (`jobRoleMatch`). The seeder currently writes the role only to
  `public.memberships` — the **wrong table** for the header.
- **G4 — real-face avatars.** Replace the DiceBear initials SVG (`users.picture`) with a **bundled,
  license-clean real-face set**, mapped **deterministically by hero key**. Offline-safe + deterministic +
  license-clean (the bundled curated set is the chosen default).

**Out:**
- Work / education history + skill depth → **M41** (profile depth).
- The surfacing / serve-grant fix → **M40** (parallel).

## Depends on / Parallel with
- **depends_on:** none.
- **parallel_with:** M40 (directus serve-grant). M39 fills profile *data*; M40 fixes the *serve path* — the
  two touch disjoint surfaces and can build concurrently.

## Open questions
- **The exact avatar source mechanism for G4** — bundled CC0-style curated set vs deterministic
  photo-generator vs snapshot asset-plane vendored. **Decide at build; bundled curated set is the chosen
  default** (offline-safe + deterministic + license-clean satisfies the acceptance bar).

## KB dependencies
Read these as contract before touching code:
- `corpus/ops/demo/stories-spec.md` — the verified-skill / hero-profile chain this layer extends
  (the profile header + claimed-vs-verified surfaces the org-name + role backfill feed).
- `corpus/services/clerkenstein.md` — the roster → `RosterEntry` → FAPI org-resource path G1 threads through.
- `corpus/ops/seeding-spec.md` — the seeder fleet + blueprint contract (`users.go`, the `user_basic_info`
  surface, the deterministic-by-hero-key convention G4 reuses).

## Delivers →
- `corpus/ops/demo/stories-spec.md` — **NEW profile-identity layer**: how the seeded org name, the
  `user_basic_info.job_role_id` backfill (header + role-gap widgets), and the bundled real-face avatar set
  materialize on a logged-in hero's profile.
- `corpus/services/clerkenstein.md` — the **roster org-name threading**: `RosterIdentity`/`BuildRoster` →
  `RosterEntry` → `DemoUser`/`orgMemberships()`, the `DisallowUnknownFields` paired-change rule, and the
  no-roster `"Clerkenstein Demo Org"` default.
