---
name: demo-seed
description: Seed a running demo stack with realistic structural data (org + users + memberships + the user_clerkenstein login identity + backdated activity) from a preset or a stack.seed.yaml. Use after /demo-up, when asked to populate / seed a demo with data.
argument-hint: [N] [--preset small-200|mid-500|large-1k | --seed path.yaml]
---

# Demo Seed — backfill a demo stack with a believable data world

Seeds `demo-N` with structural data so the demo is **usable** (the `user_clerkenstein` login lands in a
populated org and authorized routes return **200**, not 403) and **believable** (1k users, months of backdated
job-sim / skill-path sessions, assignments, activity). It talks **directly to the stack's Postgres** (offset
port) via `COPY` — fast (~1s at 1k) and **production-safe** (a hard isolation guard blocks every shared/prod
store). Source of truth: [`corpus/ops/seeding-spec.md`](../../../corpus/ops/seeding-spec.md).

## Mission
1. **Read the spec** — `corpus/ops/seeding-spec.md` (the blueprint, the dependency-DAG, the **isolation
   boundary**, the casbin subtleties). Confirm the target is a **non-prod** demo stack.
2. **Confirm the stack is up + migrated** — `/demo-up N` first if needed. `migrate-demo.sh` bootstraps the
   global Sentinel policy (required for authorized routes to return 200).
3. **Seed** via the tool (gitignored at `anthropos-demo/rosetta-extensions/stack-seeding/`):
   ```bash
   SS=anthropos-demo/rosetta-extensions/stack-seeding
   go build -o /tmp/stackseed "$SS/cmd/stackseed"
   # a curated preset (recommended):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/mid-500.seed.yaml"
   # or a custom blueprint:
   /tmp/stackseed --stack demo-N --seed my.seed.yaml
   # preview only (ordered plan + per-surface counts + the isolation preview — no writes):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/mid-500.seed.yaml" --dry-run
   # reset a stack's seeded tables (per-stack only; refuses the dev stack N=0):
   /tmp/stackseed --stack demo-N --reset
   ```
   Presets: `small-200` (quick), `mid-500` (the default "looks real"), `large-1k` (scale).
4. **Verify** — the run prints per-surface row counts + `isolation: clean`; a browser/API login as
   `user_clerkenstein` should return **200**. Optionally `datadna measure --stack demo-N` (conformance) +
   `datadna catalog` (what's seeded).

## Safety (the load-bearing part)
The seeder **cannot** pollute production: the isolation guard blocks writes to the shared stores (Directus,
the prod S3-public bucket, live Clerk, Customer.io/Brevo, AI APIs) on any non-prod stack — it seeds **only**
the per-stack Postgres/Redis — and an audit log **proves** zero shared/prod writes after every run. `--reset`
hard-refuses the main dev stack (N=0) without `--force`. **Never** point `--stack` at production.
