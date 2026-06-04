# M7a — progress (section checklist)

**Milestone:** M7a — Seeding framework + production-isolation safety · **Shape:** section · **Status:** planned

## S1 — `stack.seed.yaml` blueprint + the seeder contract/registry
- [ ] schema (org, size, role_mix, tier_mix, content-pack, activity span) — curator-annotatable per seeder
- [ ] the modular seeder interface `{surface, depends-on, isolation-class, primitive}` + the registry
- [ ] ~2 reference seeders (org, users) proving the contract

## S2 — the dependency-DAG orchestrator + the perf path
- [ ] a real DAG (parallel where independent), topological order documented
- [ ] Go-link-ent + Postgres `COPY` bulk-insert + goroutine fan-out; CLIs/RPC only for side-effecting primitives

## S3 — the production-isolation guard (CRITICAL)
- [ ] the enumerated shared/external surface list + per-seeder isolation-class declarations
- [ ] hard block on shared-write from non-prod (S3-public override, Clerk→Clerkenstein, Directus write → error)
- [ ] the seeding audit log (scenario_id/seeded_by/isolation_class) + a clean-audit assertion

## S4 — `--reset` / `--validate` / `--dry-run`
- [ ] `--reset` (per-stack DB only) + stable derivable identities
- [ ] `--validate` (schema + semantic) and `--dry-run` (ordered plan + per-store counts + isolation preview)

## S5 — the minimum end-to-end proof + the spec doc
- [ ] seed org + the real `user_clerkenstein` identity + casbin (plural/singular gotcha) → **browser login 200**
- [ ] `corpus/ops/seeding-spec.md` (blueprint + DAG + the isolation-safety boundary)
- [ ] **acceptance:** isolation audit clean on a non-prod run; login-to-seeded-stack returns 200
