# M7a — Spec notes

Seeded from the 2026-06-04 redesign research (`.agentspace/scratch/roadmap-research-2026-06-04.md`); fill in during build.

## Reusable references (do not reinvent)
- `make bootstrap-dev` — the working org→users→memberships→casbin reference pipeline.
- `bootstrap.JoinOrg` — the reusable per-user primitive (membership + casbin + feature grant).
- `bootstrap-{user,org}` CLIs + `SetDefaultOrgFeatureCredits` (Sentinel) — the structural primitives.
- The current M7 design (now superseded): `demo.seed.yaml` shape, dependency-order contract, `--reset/--validate
  /--dry-run`, the `user_clerkenstein` + casbin must-haves — all fold into M7a here.

## The perf path (research-decided)
Bottleneck at 1k users is **per-row CLI subprocess spawns** (1k `bootstrap-user` calls), which is **DB-IO-bound,
not CPU-bound** → language is secondary; **Rust buys nothing**. The architecture:
- **Go orchestrator** (matches the platform; can import its ent client) — link `app/internal/bootstrap` /
  `ent.Client.*.CreateBulk` instead of `go run`-ing CLIs.
- **Postgres `COPY`** for high-volume surfaces (users, memberships, sessions, activity) — 10–100× over row inserts.
- **Goroutine fan-out** for side-effecting primitives that can't bulk-copy (Sentinel grants, `JoinOrg`).
- Target: **< 2 min** for a 1k-user org with months of activity.

## The isolation guard (the safety contract)
- Each seeder module **declares its isolation class**: `per-stack-isolated` | `shared-pollution-risk` | `external`.
- The orchestrator **refuses** any `shared-pollution-risk` write unless `--target=production` + an explicit
  interactive opt-in.
- Pre-flight: override `STORAGE_S3_PUBLIC_BUCKET`→empty; assert Clerk base-URL points at Clerkenstein; assert no
  `DIRECTUS_*` write client is constructed.
- Post-run: the **seeding audit log** (`scenario_id` / `seeded_by` / `isolation_class` / store / row-count) must
  show **zero** `shared`/`external` writes on a non-prod run, or the run fails its own acceptance.

## Hard line
Structural data only. AI transcripts / AI-scored narratives / freshly-computed embeddings are out (v1.2 stretch).
M7a/c consume a **pre-embedded** skiller snapshot; Directus content arrives via snapshot-replay, never live writes.

## To confirm during build
- Which `created_at`/timestamp columns are settable via direct SQL (backdating) vs ent-Immutable/DB-default `now()`.
- The FAILED-session validation-result table chain (not just passed) — needed by M7c's activity generator.
- Whether `bootstrap-user/org` at 1k iterations hit any Clerkenstein mock state limits.
- That the platform's `app/internal/bootstrap` + ent packages import into a `rosetta-extensions/` Go module
  cleanly (GOPRIVATE / replace) without editing a platform repo.
