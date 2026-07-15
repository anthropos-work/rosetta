---
name: stack-seed
description: Seed a running stack (dev-N or demo-N) with realistic structural data — org(s) + users + memberships + the login identity + backdated activity, and (the M35 Stories & Heroes model) multiple orgs each with a thriving/struggling/manager hero trio — from a preset or a stack.seed.yaml / stack.stories.yaml. Use after the stack is up, when asked to populate / seed a stack with data.
argument-hint: [dev-N|demo-N] [--preset dev-min|small-200|mid-500|large-1k|stories-maya|stories | --seed path.yaml] [--dry-run|--reset]
---

# Stack Seed — backfill any stack with a believable data world

Seeds a target stack (`dev-N` or `demo-N`) with structural data so it is **usable** (the login identity lands
in a populated org and authorized routes return **200**, not 403) and **believable** (orgs, users, months of
backdated job-sim / skill-path sessions, assignments, activity). It talks **directly to the stack's Postgres**
(offset port) via `COPY` — fast (~1s at 1k) and **production-safe** (a hard isolation guard blocks every
shared/prod store). Source of truth: [`corpus/ops/seeding-spec.md`](../../../corpus/ops/seeding-spec.md).
(Formerly `/demo-seed`, now accepting both stack types.)

## Mission
1. **Read the spec** — `corpus/ops/seeding-spec.md` (the blueprint, the dependency-DAG, the **isolation
   boundary**, the casbin subtleties, the shipped-presets table). Confirm the target is a **non-prod** stack
   (`dev-N` / `demo-N`, never production).
2. **Confirm the stack is up + migrated** — `/dev-up N` or `/demo-up N` first if needed. The migrate step
   bootstraps the global Sentinel policy (required for authorized routes to return 200).
3. **Seed** via the tool (gitignored at `stack-<role>/rosetta-extensions/stack-seeding/`). The seeder
   (section `stack-seeding`, unchanged) is consumed from the stack's tagged `rosetta-extensions` clone; its
   canonical source is the `.agentspace/rosetta-extensions/` authoring copy. Use the matching per-stack clone
   for the target (`stack-dev/` for a dev-N, `stack-demo/` for a demo-N):
   ```bash
   SS=stack-demo/rosetta-extensions/stack-seeding   # or stack-dev/... for a dev-N
   go build -o /tmp/stackseed "$SS/cmd/stackseed"
   # a curated preset (recommended):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/mid-500.seed.yaml"
   # a light dev preset (the default for a dev stack — ~1 org + ~10 users):
   /tmp/stackseed --stack dev-N --seed "$SS/presets/dev-min.seed.yaml"
   # the M35 Stories & Heroes world (MULTIPLE orgs, each with a thriving/struggling/manager hero trio):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/stories.seed.yaml"
   # …or the single-hero vertical slice (one org + Maya):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/stories-maya.seed.yaml"
   # or a custom blueprint:
   /tmp/stackseed --stack dev-N --seed my.seed.yaml
   # preview only (ordered plan + per-surface counts + the isolation preview — no writes):
   /tmp/stackseed --stack demo-N --seed "$SS/presets/mid-500.seed.yaml" --dry-run
   # reset a stack's seeded tables (per-stack only; refuses the main dev stack N=0 without --force):
   /tmp/stackseed --stack demo-N --reset
   ```
   > **`--preset NAME` is the skill-level shorthand** (the argument-hint above), not a `stackseed` CLI flag —
   > resolve it to `--seed "$SS/presets/NAME.seed.yaml"` when you build the command (the binary only knows
   > `--stack`, `--seed`, `--dsn`, `--reset`, `--validate`, `--dry-run`, `--force`).

   Presets: `dev-min` (~1 org / ~10 users — the dev default), `small-200` (quick), `mid-500` (the default
   "looks real"), `large-1k` (scale), `stories-maya` (one org + one hero — the M34 vertical slice),
   `stories` (the **M35 multi-org Stories & Heroes** world — **3 orgs** (Cervato Systems / Solvantis /
   Northwind Aviation) × a thriving/struggling/manager trio,
   the believable demo-narrative default; **replay the taxonomy first** so heroes get role-coherent skills).
4. **Verify** — the run prints per-surface row counts + `isolation: clean`; a browser/API login (the seeded
   admin: `dev@anthropos.test` for `dev-min`, `user_clerkenstein` for the demo presets) should return **200**.
   Optionally `datadna measure --stack <target>` (conformance) + `datadna catalog` (what's seeded).

## Safety (the load-bearing part)
The seeder **cannot** pollute production: the isolation guard blocks writes to the shared stores (Directus,
the prod S3-public bucket, live Clerk, Customer.io/Brevo, AI APIs) on any non-prod stack — it seeds **only**
the per-stack Postgres/Redis — and an audit log **proves** zero shared/prod writes after every run. `--reset`
hard-refuses the **main dev stack (N=0)** without `--force`. **Never** point `--stack` at production.

## Related skills

| Skill | Use when |
|-------|----------|
| `/stack-snapshot` | Set-dress the stack with the real public catalog **before** seeding (full fidelity) |
| `/stack-list` | List live stacks to pick a target |
| `/dev-up` · `/demo-up` | Bring up the stack first |
