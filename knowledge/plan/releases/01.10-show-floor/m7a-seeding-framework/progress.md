# M7a — progress (section checklist)

**Milestone:** M7a — Seeding framework + production-isolation safety · **Shape:** section · **Status:** planned

## S1 — `stack.seed.yaml` blueprint + the seeder contract/registry ✅
- [x] schema (org, size, role_mix, tier_mix, content-pack, activity span) — `blueprint/` + `Validate` (8 tests)
- [x] the modular seeder interface `{surface, depends-on, isolation-class, primitive}` + the registry — `seeder/` (6 tests)
- [x] ~2 reference seeders (org, users) + the `identity` minimum-proof seeder — `seeders/` (11 tests)

## S2 — the dependency-DAG orchestrator + the perf path ✅
- [x] a real DAG (topo-sort + level-parallel exec, guard-checked), cycle detection — `seeder/dag.go` (10 tests)
- [x] **perf path REVISED (M7a-D3):** direct Postgres `COPY`/SQL over the offset port (NOT ent-linking — blocked
  by the `internal/` rule; NOT CLI-shelling — the slow path). `pg/` (`DSNForOffset`, pgx `CopyRows`) (9 tests)

## S3 — the production-isolation guard (CRITICAL) ✅
- [x] the enumerated shared/external store registry + per-seeder isolation-class declarations — `isolation/` (18 tests)
- [x] hard block on shared-write from non-prod (`CheckWrite` asymmetry; `PreflightEnv` S3-public override + Clerk + Directus assertions)
- [x] the seeding audit log (scenario_id/seeded_by/isolation_class) + `AssertClean` post-run pollution proof

## S4 — `--reset` / `--validate` / `--dry-run` ✅
- [x] `--reset` (per-stack DB only; refuses n=0 dev unless `--force`) + stable derivable identities
- [x] `--validate` + `--dry-run` (ordered plan + per-surface row estimates + **isolation preview** [BLOCK]/[SEED]) — verified via CLI

## S5 — the minimum end-to-end proof + the spec doc ✅
- [x] `corpus/ops/seeding-spec.md` (blueprint + DAG + the isolation-safety boundary) + the section `README.md`
- [x] seed org + the real `user_clerkenstein` identity + casbin (plural/singular gotcha + the g2 arg-order fix) → **login 200** PROVEN
- [x] **acceptance:** isolation audit clean on a non-prod run (`AssertClean`) + login-to-seeded-stack returns **200**

**PROOF (live, reproducible, no manual SQL — M7a-D4):** brought up the injected `demo-1` stack (14 containers,
Clerkenstein FAPI+BAPI), wiped to fresh, ran the FIXED `migrate-demo.sh` (now bootstraps the global Sentinel
policy) → seeded via the FIXED `stackseed` (org=1, identity=3, users=2000, **isolation: clean**) → authenticated
`user_clerkenstein` GraphQL query returned **HTTP 200 `{membershipsCount: 1001}`**. Arc: no token → "unknown
viewer"; token pre-seed → 403; token post-seed → **200**. The proof caught + fixed **2 real bugs** (the g2
arg-order in `identity.go`; the missing global-policy bootstrap in `migrate-demo.sh`) — see M7a-D4.

**Build status:** `stack-seeding` complete — 19 files, **62 tests** (build/vet/`-race`/gofmt clean), shellcheck
clean (`migrate-demo.sh`). All sections S1–S5 done.
