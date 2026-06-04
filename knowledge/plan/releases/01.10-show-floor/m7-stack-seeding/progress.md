# M4 — progress (section checklist)

**Milestone:** M4 — Declarative data seeding · **Shape:** section · **Status:** planned

## S1 — `demo.seed.yaml` schema + the seeder skeleton (in `anthropos-demo/`)
- [ ] schema (org, size, role_mix, tier_mix, content pack, activity span)
- [ ] seeder targets a named M3 stack (`-p demo-N`, its `.env`/DB) and calls existing CLIs as-is

## S2 — the dependency-order pipeline
- [ ] migrate → Sentinel policy → org → users → memberships+casbin+feature (`JoinOrg`) → taxonomy/content (snapshot)
- [ ] pre-embedded skiller snapshot prerequisite (consume, don't generate)

## S3 — deterministic time-distributed activity
- [ ] activity generator over **real** sim_ids (passed/failed sessions across the span; backdated where allowed)

## S4 — validation + idempotency (the M4b fold-in)
- [ ] `--validate` (schema + semantic checks) and `--dry-run` (ordered insert plan + per-store counts)
- [ ] `--reset` (per-demo DB only) + stable derivable demo identities

## S5 — the spec doc + acceptance
- [ ] `corpus/ops/seeding-spec.md`
- [ ] **acceptance:** one `demo.seed.yaml` seeds an M3 stack to a coherent target org at ~200+ users with months of activity
