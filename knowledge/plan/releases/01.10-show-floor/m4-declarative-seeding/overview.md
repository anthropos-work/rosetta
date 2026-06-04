---
milestone: M4
slug: declarative-seeding
version: v1.1 "show floor"
milestone_shape: section
status: planned
created: 2026-06-03
last_updated: 2026-06-03
delivers: corpus/ops/seeding-spec.md
---

# M4 — Declarative data seeding

## Goal
Describe a target demo state in **one declarative config** (`demo.seed.yaml`: org size, role mix, content
sources, activity span) and **backfill an M3 demo stack to match** — by orchestrating the platform's
**existing** bootstrap/import CLIs in dependency order, so an operator can stand up "an org of 1k users with
months of activity" without hand-running a dozen tools.

## Why a section milestone (buildable, grounded)
The skeleton is fully grounded in working references: `make bootstrap-dev` already runs the
org→users→memberships→casbin pipeline; `bootstrap.JoinOrg` is the reusable per-user primitive;
`seed-verified-skill` is the proven pattern for time-distributed activity using **real** sim_ids. The `In:` list
is therefore writable now → **section**. The weight is in volume + backdating fidelity, not in an uncertain path.

## Scope
### In
- **The `demo.seed.yaml` schema**: org (name); size (N users); role_mix (admin/member/candidate ratios + a small
  fixed admin set); tier_mix (free/premium); content pack reference; activity span (e.g. "3 months", session
  pass/fail distribution).
- **A Go seeder binary in `anthropos-demo/`** (NOT a platform repo) that targets a named M3 stack (`-p demo-N`,
  its `.env.demo-N` + DB) and runs the pipeline by **calling the existing CLIs / primitives as-is** (no platform
  edits): `app/cmd/bootstrap-{user,org}` / `bootstrap.JoinOrg`, `skiller/cmd/import{Skills,JobRole}`,
  `cms/cmd/jobsim`.
- **A documented dependency-order contract**: migrate → Sentinel policy → org → users → memberships + casbin +
  feature (via `JoinOrg`) → taxonomy/content (from a snapshot) → time-distributed activity.
- **The deterministic activity generator** modeled on `seed-verified-skill`: per user, sample from the content
  pack's **real** sim_ids, emit passed/failed sessions distributed across the activity span (backdated where the
  schema allows).
- **Idempotency/cleanup**: a `--reset` path (per-demo DB only) + stable, derivable demo identities
  (e.g. `user{i}@demo-N.local`) so re-running is deterministic.
- **Inherited from M3 (Fate-3, deferral audit 2026-06-04) — the login identity must be real.** Beyond the bulk
  `user{i}@demo-N.local` users, the seed MUST create the org's **admin/login user matching
  `clerkfrontend.DefaultDemoUser()`** — `user_clerkenstein` / `demo@anthropos.test` / `org_clerkenstein`
  (role `admin`) — so a browser login (which mints exactly that identity) resolves to a *seeded* user and
  authorized routes return 200 instead of 403. (This is the real demo identity — NOT the `user_2clerkenstein`
  alignment-runner fixture some early M3 notes propagated.) Two M3 gotchas to honor: (a) Sentinel's
  `init_policy.sql` seeds `casbin_rules` (plural) while the gorm adapter auto-creates `casbin_rule` (singular)
  — reconcile or the policies won't load; (b) clerkenstein's `clerk-webhook` injector is an available mechanism
  to feed identities into the platform sync pipeline post-seed (an alternative to direct DB inserts).
- **Folded-in validation**: a `--validate` (schema + semantic checks) and `--dry-run` (ordered insert plan +
  per-store row counts) mode **in the seeder itself** (the M4b hedge — see Decisions).
- **Delivers** `corpus/ops/seeding-spec.md`: the `demo.seed.yaml` reference + the dependency-order pipeline + the
  snapshot prerequisite contract.

### Out
- **AI-generated rich content** (realistic transcripts, AI-scored validation narratives, freshly-computed
  pgvector embeddings) — the hard line; M5 stretch / deferred.
- Authoring the 60K-skill taxonomy or jobsim content per stack — M4 **consumes a PRE-POPULATED, PRE-EMBEDDED
  skiller snapshot**, it does not generate one.
- The disposable-stack lifecycle (compose/-p/ports/`/demo-up`) — that's M3.
- Curated recipes + the demo corpus index + skill polish — M5.
- Any modification to platform repos (CLIs called as-is from the per-demo clones).

## Depends on
**M3** (a disposable, Clerkenstein-wired stack to seed into) + v1.0 M1/M2 (Clerkenstein removed Clerk's API rate
limit → seeding users is pure DB inserts, no Clerk API). **Parallel with:** none.

## Estimated complexity
**large** — the heaviest v1.1 milestone: many stores, dependency ordering, 1k-scale performance, and backdating
fidelity (some `created_at` columns are DB-default `now()` / ent-Immutable).

## Open questions (resolve during build)
- Single `demo.seed.yaml` fanned out by one seeder (recommended) vs per-store seeders.
- Vector embeddings: recompute violates the hard line → depend on a **pre-embedded skiller snapshot**; confirm it
  can be produced/shipped without prod data.
- Directus content tenancy in a multi-demo world (shared content instance vs per-demo).
- Call CLIs as compiled binaries (`go run`, slower at 1k) vs link `app/internal/bootstrap` into the seeder for speed.
- Backdating fidelity: which timestamps are settable via direct SQL vs ent-Immutable/DB-default.

## KB dependencies (read as contract)
- `corpus/architecture/architecture_overview.md` (multi-tenancy: `organization_id` on every table, 3-layer isolation)
- `corpus/ops/staging_from_dump.md` (snapshot/restore patterns for the pre-embedded skiller snapshot)
- `corpus/services/{backend,skiller,jobsimulation,skillpath,cms}.md` (the bootstrap/import CLIs + store schemas)
- M3's `corpus/ops/rosetta_demo.md` (the stack to seed into)

## Delivers → `corpus/ops/seeding-spec.md` (net-new)
The declarative-seeding spec: `demo.seed.yaml` reference, the dependency-order pipeline, the snapshot prerequisite.

## Exit (section)
All `In:` deliverables land: one `demo.seed.yaml` seeds a named M3 demo stack to a coherent org (target N users,
role/tier mix, time-distributed real-sim activity), `--dry-run` + `--validate` work, `--reset` is idempotent —
documented in `corpus/ops/seeding-spec.md`.
