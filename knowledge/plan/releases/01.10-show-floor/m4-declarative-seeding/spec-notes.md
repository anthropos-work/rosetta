# M4 — Spec notes

Seeded from Phase 1 research (2026-06-03); fill in during build.

## Reusable references (do not reinvent)
- `make bootstrap-dev` — the working org→users→memberships→casbin reference pipeline.
- `bootstrap.JoinOrg` — the reusable per-user primitive (membership + casbin + feature).
- `seed-verified-skill` — the proven pattern for time-distributed activity using **real** sim_ids.

## Dependency order (the contract)
migrate → Sentinel policy → org → users → memberships + casbin + feature (`JoinOrg`) → taxonomy/content (snapshot)
→ time-distributed activity.

## Hard line
Structural data only. AI transcripts / AI-scored narratives / freshly-computed embeddings are out (M5 stretch).
M4 consumes a **pre-embedded** skiller snapshot.

## To confirm during build
- which `created_at`/timestamp columns are settable via direct SQL (backdating) vs ent-Immutable/DB-default `now()`.
- the FAILED-session validation-result table chain (not just passed).
- whether bootstrap-user/org at 1k iterations hit any Clerkenstein mock state limits.
