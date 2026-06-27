# M39 Decisions

Implementation decisions with rationale (recorded during build). Design-time context lives in the spec
([`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md); root-cause workflow w7t4wq2z4).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M39-D1 | G1 org-name: add `org_name`/`org_slug` to BOTH `RosterIdentity` (roster.go) and `RosterEntry` (registry.go) as a paired change; clerkenstein is a SUB-TREE of the one rext repo (not a separate git repo / submodule), so the paired structs commit together and the repo gets ONE tag `method-acting-m39`. | The `DisallowUnknownFields` decoder rejects a producer/consumer field mismatch, so both structs must change in lockstep. The "two repos" in the spec are two sub-trees of one git repo. | 2026-06-24 |
| M39-D2 | Single-source the org slug via a new `orgSlugFor(OrgSpec)` helper (helpers.go); both `OrgSeeder` (writes `organizations.slug`) and `BuildRoster` (threads into the FAPI org resource) call it. | The roster-carried slug MUST equal the seeded `organizations.slug` or the top bar and the DB disagree. Extracting the rule (was inline in org.go) makes the equality provable, not coincidental. | 2026-06-24 |
| M39-D3 | Keep `"Clerkenstein Demo Org"` / `"clerkenstein-demo"` as named constants (`orgNameDefault`/`orgSlugDefault`) and apply them as the empty-`OrgName` fallback in `orgMemberships()`. | `DefaultDemoUser` (no org name) + any pre-M39 roster must stay byte-identical; the fallback makes the no-roster path unchanged + forward-compatible (a roster omitting the fields still works). | 2026-06-24 |
| M39-D4 | G2 backfill is an idempotent UPDATE (not INSERT) of the trigger-created `user_basic_info` row, keyed by `id`. Real schema (introspected on demo-3): columns are `job_role_id`/`job_title`/`summary`/`location` (NO `job_role_title` — the spec-ref name was approximate); `email` is NOT NULL UNIQUE. | The platform's `users` AFTER-INSERT trigger `init_user_tables()` already creates the row; an INSERT would collide on the NOT-NULL UNIQUE email. UPDATE keyed by id is the only safe write. | 2026-06-24 |
| M39-D5 | The UPDATE is made idempotent via an `IS DISTINCT FROM` guard on all four fields (re-seed of identical data → 0 rows affected), validated live on demo-3 (1st UPDATE 1 row, 2nd 0). | The M17 re-run contract (`corpus/ops/idempotency.md`) requires "2nd seed inserts nothing new"; a bare UPDATE reports rows-matched every run. The guard makes a no-change re-seed a true no-op, mirroring the COPY ON-CONFLICT + casbin WHERE-NOT-EXISTS patterns. The fake Conn models the same guard. | 2026-06-24 |
| M39-D6 | Backfill EVERY member (not heroes-only), and include `summary`+`location` for all. No-fabrication preserved: `job_role_id` NULL without taxonomy; a hero keeps her DECLARED role label as `job_title` (the same split `users.go` already applies to `memberships.job_role_name`). | The role-gap widgets key off `user_basic_info.job_role_id` for ANY viewed member, and every M35 member already has a real membership role — backfilling all of them is consistent + lights any profile a presenter clicks into. summary/location are cheap deterministic believability fields. | 2026-06-24 |
| M39-D7 | G4 avatar: a self-authored parametric SVG **face generator** emitted as a base64 **data URI** (offline, in `users.picture`), NOT a bundled photo set or the DiceBear HTTP API. Chose the review's "deterministic photo-style generator" option over "bundled curated set". | All three options had to be offline-safe + deterministic + license-clean. A bundled PHOTO set means sourcing+vendoring genuinely-CC0 portraits (hard, heavy binary, license risk) and an asset-serving plane; the DiceBear HTTP URL fails offline-safe. A self-authored SVG generator is offline (data URI, no fetch), deterministic (hashed), license-clean (no third-party asset at all), tiny (~1 KB), and IS a face — the strictly cleanest answer. | 2026-06-24 |
| M39-D8 | Removed the now-unused `avatarURL` (DiceBear initials) helper rather than leaving it as dead code; moved its test to `avatar_test.go` against the new generator. | The PR-review dead-code category: an unexported helper with no production caller is dead. The old test now covers the replacement. | 2026-06-24 |

## Adversarial review (close-milestone Phase 2c, 2026-06-24)

Scenarios considered for the just-shipped code; recorded so future reviewers see what was weighed.

- **AR-1 — G2 backfill "first seed, 0 rows updated, no SQL error".** `backfillUserBasicInfo` returns
  `written` = Σ rows-affected and treats `(0, nil)` as success; the `if basicWritten > 0` guard then skips
  even the audit entry. If the platform's `init_user_tables()` AFTER-INSERT trigger were broken/absent, the
  UPDATE would match 0 rows *without* an SQL error and the `/profile` header would silently stay blank.
  **Response — no fix needed:** the backfill runs immediately after the `users` COPY in the **same seed run**,
  and COPY fires the per-row AFTER-INSERT trigger that creates the `user_basic_info` row (documented at
  `users.go:28-29`), so the row is a hard structural invariant by the time the UPDATE runs. A genuinely
  missing trigger row surfaces as an SQL error on a different dependent table long before G2 (and that error
  path IS covered loud — `TestBackfillUserBasicInfo_ExecErrorPropagates`,
  `TestUsersSeeder_BasicInfoBackfillErrorBubblesThroughSeed`). A defensive "first-seed produced 0" assert
  would guard only against a platform-trigger failure that already manifests loudly elsewhere — not an M39
  gap. The 0-rows return is also the *correct* re-seed no-op (the IS-DISTINCT-FROM guard), so 0 cannot be a
  hard error without breaking idempotency (M39-D5).
- **AR-2 — G1 two distinct stories with the same org name → same `orgSlugFor` slug → `organizations.slug`
  collision.** **Response — pre-existing + unchanged by M39:** `orgSlugFor` (M39-D2) extracted the *identical*
  inline rule that lived in `org.go` (`explicit slug, else slugify(name)`); the slug derivation is byte-for-byte
  what it was pre-M39, so M39 introduced no new collision surface. Distinct stories get distinct deterministic
  org **ids** (the COPY upsert keys on `id`); a duplicate-org-name blueprint is a seed-authoring concern outside
  M39's scope. Not a finding.
