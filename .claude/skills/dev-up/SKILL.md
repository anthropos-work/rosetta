---
name: dev-up
description: Bring up a local DEV stack — build-or-resume the environment, start it, and (for an additional dev-N) set-dress it with a cache-first snapshot replay + a light dev-min seed, plus (opt-in via --local-content) an EXECUTED per-stack Directus so the stack's content is self-contained (otherwise it reads content live from prod). Consolidates the former setup-platform + start-platform. Use to set up, start, or restart a dev stack locally.
argument-hint: [N | 'main'] [--no-setdress] [--no-snapshot] [--profile P] [scenario|step]
---

# Dev Up — build, start, and set-dress a local dev stack

Brings a DEV stack to a running, set-dressed state. One skill for the whole dev lifecycle bring-up:

- **`dev-up`** / **`dev-up main`** (N=0, the primary `anthropos` dev stack) — the heavy first-time
  environment build (or a resume) **and** the start, following the official guides with verification.
  This is the consolidation of the former `/setup-platform` + `/start-platform`.
- **`dev-up N`** (N ≥ 1) — spin up an **additional, isolated** `dev-N` stack alongside the main one,
  on offset ports, and give it the **demo treatment by default** (M13): a cache-first snapshot replay of
  the real public reference data, a light `dev-min` seed, and the per-stack-Directus firewall check.
  The per-stack Directus itself is **opt-in for dev** via `--local-content` (v1.5 M22/M23): with the flag,
  the recipe is **EXECUTED** (bootstrap → apply-structure → replay → boot a per-stack Directus on an offset
  port) and `cms`'s `DIRECTUS_BASE_ADDR` is **cut over** to it, so the stack's content is **self-contained**;
  without the flag (the dev default), the recipe is print-only and the stack reads content **live from prod**
  (the documented fallback). See [`corpus/ops/directus-local.md`](../../../corpus/ops/directus-local.md).

It mirrors `/demo-up` for the dev side — same registry, same offset-port model, same safety guards —
so `dev-N` and `demo-N` are first-class peers and never collide on ports (the M12 unified registry).

Sources of truth: [`corpus/ops/setup_guide.md`](../../../corpus/ops/setup_guide.md) (first-time build),
[`corpus/ops/run_guide.md`](../../../corpus/ops/run_guide.md) (start + health),
[`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md) (the registry + offset ports),
[`corpus/ops/seeding-spec.md`](../../../corpus/ops/seeding-spec.md) (the `dev-min` seed) and
[`corpus/ops/snapshot-spec.md`](../../../corpus/ops/snapshot-spec.md) (the auto-snapshot + the per-stack
Directus recipe/known-state).

## Two modes (pick from the target)

### A) `dev-up` / `dev-up main` — the main dev stack (N=0)

The first-time machine build + start. Drives the official guides with STEP RUN discipline (verify
before/after each step, request confirmation before installs or destructive ops, log ops-reports).

1. **Read the guides** — `corpus/ops/setup_guide.md` (build) is the source of truth for the heavy build;
   `corpus/ops/run_guide.md` (start) for the start + health pass. Resume an interrupted build from where
   it stopped (re-verify each completed step rather than redo it).
2. **Apply STEP RUN** to every step:

   | Principle | Action |
   |-----------|--------|
   | Verify Before Install | Check if a tool exists before installing |
   | Request Confirmation | Ask before installs, repo clones, starting services, `.env` edits |
   | Verify After Install | Confirm each install/step succeeded |
   | Report Issues | Write an ops-report when an error or improvement is found |

3. **Track progress** via TodoWrite (build phases): prerequisites verified (Git, Docker, Go, **Node v24+**,
   pnpm, Python, Atlas) → GitHub SSH (`/setup-github`) → workspace `stack-dev/` → platform repo cloned →
   all repos via `make init` (incl. `ant-academy`) → CMS studio submodule (`cd cms && make init-studio`) →
   `platform/.env` configured → services up (`make up` — expect **12 containers** in `graphql`) → PostgreSQL
   schemas (`extensions`, `sentinel`) → migrations (`make migrate`) → frontend + Studio-Desk deps → health.
4. **Start + verify health** (the former `/start-platform` pass): `make up`, confirm 12 healthy containers
   (`make ps`), GraphQL gateway on `localhost:5050`, frontend (Node v24+), Studio-Desk. Ask before stopping/
   restarting, killing port-conflict processes, or `make reset-db`.

The main dev stack uses **real Clerk** by default — it is not a snapshot/seed target unless you explicitly
ask. (Set-dressing the main stack would reset its data; the `dev-min` seed + snapshot are for `dev-N`.)

### B) `dev-up N` (N ≥ 1) — an additional isolated dev stack, set-dressed

Spins up `dev-N` alongside the main dev stack and (by default) set-dresses it — the M13 dev-peer flow.

1. **Read** `corpus/ops/rosetta_demo.md` (offset ports + the unified registry) +
   `corpus/ops/snapshot-spec.md` / `corpus/ops/seeding-spec.md` (the set-dress pass).
2. **Resource check** — a full stack is ~10–12 GB; confirm headroom (`docker info` MemTotal vs running
   stacks). Never exceed the box.
3. **Bring it up + set-dress** via the dev-stack tooling, consumed from the dev stack's **OWN**
   `stack-dev/rosetta-extensions` clone pinned at a tag (never edited ad-hoc; authored + tested + tagged
   first in `.agentspace/rosetta-extensions/`):
   ```bash
   DEV=stack-dev/rosetta-extensions/dev-stack
   "$DEV/dev-stack" up N                 # allocate N via the unified registry, bring up dev-N on offset ports,
                                         # then run the default set-dress pass: the per-stack Directus
                                         # recipe + firewall, cache-first snapshot replay, dev-min seed (M13).
   "$DEV/dev-stack" up N --no-snapshot   # seed only (skip the snapshot replay)
   "$DEV/dev-stack" up N --no-setdress   # bare bring-up (no snapshot, no seed)
   "$DEV/dev-stack" up N --inject        # optional: Clerkenstein-inject (offline/clean-room dev)
   ```
   The set-dress pass is **default-on + non-fatal**: a stale/missing snapshot cache is a warning, the seed
   still runs; the whole pass can be skipped with `--no-setdress`. (`dev-setdress.sh`, M13.)
4. **Verify** — `"$DEV/dev-stack" status` (or `/stack-list`); confirm `dev-N` is on offset ports, the
   **main dev stack is untouched**, the catalog shows real skills (snapshot), and a `dev@anthropos.test`
   login lands in a populated org (the `dev-min` seed → 200).

## Defaults & flags

- `--no-setdress` — bare bring-up of `dev-N` (no snapshot, no seed).
- `--no-snapshot` — seed `dev-N` but skip the snapshot replay (faster; empty catalog + free content refs).
- `--profile P` — compose profile (default `graphql`).
- Set-dressing applies to **`dev-N` (N ≥ 1)** only; `dev-up main` (N=0) stays real-Clerk + unseeded.

## Safety (the load-bearing part)

- **N-allocation is collision-free** — `dev-stack up` allocates the lowest free N across **both** dev and
  demo via the M12 unified registry, so `dev-N` / `demo-N` can never coexist on the same ports.
- **Every op is `-p dev-N`-scoped** and **hard-refuses any N that resolves to the main dev project** — it
  can never touch the main dev stack.
- **The set-dress pass is per-stack-isolated** — snapshot **replay** is a per-stack WRITE (never a prod
  read; capture is never run here), the per-stack Directus env contract targets the stack's own `directus`
  schema (`EnvContract.Validate()` hard-rejects `content.anthropos.work`). With **`--local-content`** the
  recipe is **EXECUTED** behind that now-load-bearing firewall gate (a prod-resolving env hard-aborts before
  any write) and the directus replay **exits 0**; without it the boot is print-only and the directus replay
  skips with stacksnap exit 4 — the documented prod-read fallback. Media stays refs-only, and the
  `dev-min` seeder's isolation guard blocks every shared/prod store. The **n=0 set-dress guard is doubled**
  — the pass refuses N=0 without `--force` (a second layer above `stackseed --reset`'s own refusal).
- **All stack-operating tooling lives in `rosetta-extensions`** — consumed per-stack at a pinned tag, never
  improvised inside a stack dir. rosetta = read-only doc corpus + dev-env skills.

## Ops reports

When you hit an error, a missing step, or a better approach, write
`stack-dev/ops-reports/op_YYYYMMDD_HHMMSS_devup_<topic>.md` (Issue / Context / Resolution / Suggested doc
update). These feed `/update-knowledge`.

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-down` | Stop / reclaim a dev stack |
| `/demo-up` · `/demo-down` | The demo-stack lifecycle (the peer of dev-up/dev-down) |
| `/stack-list` | List live dev + demo stacks |
| `/stack-seed` · `/stack-snapshot` | Re-seed / re-set-dress any stack (`dev-N` or `demo-N`) |
| `/stack-update` | Sync a stack's code, deps, and schemas |
| `/setup-github` | Configure GitHub SSH access (a one-time prerequisite of the first build) |

## Additional resources

- For technical reference (verification commands, health checks, error recovery), see [reference.md](reference.md).
