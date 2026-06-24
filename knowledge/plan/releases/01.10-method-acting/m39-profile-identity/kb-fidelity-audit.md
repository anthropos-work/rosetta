---
title: "KB Fidelity Audit — M39 profile identity"
date: 2026-06-24
scope: milestone:M39
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Roster org-name threading (G1) | `corpus/services/clerkenstein.md` (multi-identity roster §) | `stack-seeding/seeders/roster.go` (`RosterIdentity`/`BuildRoster`), `clerkenstein/clerk-frontend/registry.go` (`RosterEntry`), `clerkenstein/clerk-frontend/resources.go` (`DemoUser`/`orgMemberships()`) | PAIRED |
| Role backfill → user_basic_info (G2) | `corpus/ops/demo/stories-spec.md`, `corpus/ops/seeding-spec.md` | `stack-seeding/seeders/users.go` | PAIRED |
| Real-face avatars (G4) | `corpus/ops/demo/stories-spec.md`, `corpus/ops/seeding-spec.md` | `stack-seeding/seeders/userprofile.go` (`avatarURL`), `users.go` | PAIRED |

## Fidelity Findings

1. **Roster shape (clerkenstein.md).** Doc describes the M37 multi-identity registry + `FAKE_FAPI_ROSTER` +
   active-seat selection. Code matches (`RosterEntry`/`Roster`/`RegistryFromRoster`/`LoadRoster`). The doc does
   NOT yet enumerate per-field roster claims (no org-name field is documented because none exists yet). Verdict:
   ALIGNED (describes current behavior). M39 adds the org-name field and `Delivers →` the doc extension.

2. **Avatar (stories-spec.md:139, seeding-spec.md:233).** Doc says "a deterministic avatar URL". Code:
   `userprofile.go::avatarURL` returns a DiceBear initials SVG URL. Verdict: ALIGNED (current). M39 replaces it
   with a bundled real-face set + `Delivers →` the doc update.

3. **Role write target (stories-spec.md:235, seeding-spec.md:251).** Doc says every member's role is set on
   `memberships.job_role_id`. Code matches (`users.go` writes job_role_id/name to `public.memberships` only).
   Verdict: ALIGNED (current) — and this is precisely the G2 gap: the doc correctly reflects that
   `user_basic_info` is NOT written. M39 adds the `user_basic_info` backfill + `Delivers →` the doc update.

## Completeness Gaps
None load-bearing. The three `Delivers →` doc extensions are M39's own deliverables (Phase 5), not pre-existing
blind areas.

## Applied Fixes
None needed — all topics PAIRED with ALIGNED claims describing current behavior. The behavior changes (and the
matching doc updates) are M39's planned work.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. All three topics have both a knowledge anchor and code; every audited claim
accurately describes the *current* behavior. M39 changes the behavior and updates the docs in lockstep via its
`Delivers →` lines (clerkenstein.md roster org-name threading; stories-spec.md profile-identity layer).
Real-schema note recorded for the build: `public.user_basic_info` has `job_role_id`, `job_title`, `summary`,
`location` (NO `job_role_title` column — the spec-ref's column name was approximate); `email` is NOT NULL UNIQUE
and the row pre-exists (created by the `users` AFTER-INSERT trigger), so G2 is an UPDATE keyed by `id`.
