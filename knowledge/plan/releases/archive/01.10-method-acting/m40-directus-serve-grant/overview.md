---
milestone: M40
slug: directus-serve-grant
version: v1.10 "method acting"
milestone_shape: section
status: archived
created: 2026-06-24
last_updated: 2026-06-24
complexity: small-medium
delivers: corpus/ops/snapshot-spec.md (the public-policy serve-grant extension) + rosetta-extensions/stack-snapshot/directus/structure.go (SYNTHESIZE public-read directus_permissions rows on the PublicPolicyID for the collections cms's anonymous read path needs but prod's public policy does not grant)
depends_on: none
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M40 — Per-stack Directus public-policy serve-grant (library + activity feed)

## Goal
The hero's content surfaces — `/library/ai-simulations`, `/library/skill-paths`, **and** `/profile/activities` —
render the **real catalog/activity** from the already-seeded data, on a **fresh `/demo-up`**. This is the highest
single-surface value in the release: one serve-grant gap, closed in the snapshot replay, unblocks the entire
skill-paths library, every sim/path detail page, the sims list, and the activity feed — with **zero seeding** (the
activity data is already correctly seeded: 21 completed sessions in `jobsimulation.sessions` /
`public.local_jobsimulation_sessions`) and **zero platform-repo edits** (next-web / app / cms / jobsimulation stay
read-only).

## Scope

**In:**
- In rext `stack-snapshot/directus/structure.go`, **SYNTHESIZE** public-read `directus_permissions` rows (on the
  `PublicPolicyID`) for the collections cms's **anonymous** read path needs but prod's public policy does **not**
  grant. The existing `servePermissionsRowsSQL` only copies rows already present on prod's public policy, so these
  must be **added explicitly**:
  - **(a) `directus_versions`** — a **SYSTEM** collection. cms `skillpath.go:64` `GetSkillPath` →
    `GetLatestOrCreateVersion` → `version.go:40` `GET /versions` **403s anonymously** and treats it as fatal →
    unblocks the **entire skill-paths library** + every sim/path detail page (**the dominant blocker**).
  - **(b) the library-category collections** that `ListPublicJobSimulations` (cms `jobsimulation.go:305`) expands —
    403 → empty relation → `ToDomain` panic `"index out of range [0]"` → unblocks the **sims list**.
  - **(c) the `simulations.sequences` O2M NESTED read** — under the public policy this is **STRIPPED even with
    `sequences|read` granted**, so cms `GetJobSimulation` gets `s.Sequences==[]` → panics at cms
    `jobsimulation.go:1097` (`s.Sequences[0]`) → the activity feed's per-row simulation federation returns null →
    the feed empties. This is purely a **serve-grant** gap, **NOT** seeding.
- A **regression test** that all three surfaces serve **>0** on a fresh demo (re-replay the snapshot into demo-3 and
  assert).

**Out:**
- Any **seeding** — the data is already correct (21 completed sessions already seeded).
- Identity / depth (M39 / M41).

## Depends on / Parallel with
- **depends_on:** none.
- **parallel_with:** M39, M41.

## Open questions
Carry prominently:
- **The `simulations.sequences` O2M may NOT be grantable** under the Directus public policy without a platform
  nil-guard at cms `jobsimulation.go:1097` (READ-ONLY — cannot edit). **Investigate the O2M public-policy mechanism
  FIRST.** If it needs a platform change:
  - the **library half** (`directus_versions` + library-categories) **STILL SHIPS independently**, and
  - the **activity-feed half** escalates for **platform sign-off** (the zero-platform-edit line).
- **KPI "AI simulations completed" = 0** may be a **separate** frontend/auth-context issue — its source
  `public.local_jobsimulation_sessions = 21` has **no CMS dependency**, so it should not be coupled to the feed fix.
  **Re-verify after the feed fix; flag separately if it persists.**

## KB dependencies
Read as contract:
- [`corpus/ops/snapshot-spec.md`](../../../../corpus/ops/snapshot-spec.md) — the `stack-snapshot` extension + the
  per-stack Directus replay this milestone extends with the serve-grant.
- [`corpus/ops/safety.md`](../../../../corpus/ops/safety.md) — §2.9 token-strip / anonymous public read (the read
  path the synthesized permissions must serve).
- [`corpus/ops/demo/README.md`](../../../../corpus/ops/demo/README.md) — the up→snapshot→seed→use→down flow the
  fresh-`/demo-up` acceptance runs through.

## Delivers →
- [`corpus/ops/snapshot-spec.md`](../../../../corpus/ops/snapshot-spec.md) — the **public-policy serve-grant
  extension** (the synthesized public-read `directus_permissions` rows on the `PublicPolicyID`, what they unblock,
  and the O2M-strip caveat).
