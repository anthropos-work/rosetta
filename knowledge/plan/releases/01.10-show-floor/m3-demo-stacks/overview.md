---
milestone: M3
slug: demo-stacks
version: v1.1 "show floor"
milestone_shape: section
status: planned
created: 2026-06-03
last_updated: 2026-06-03
delivers: corpus/ops/demo_stacks.md
---

# M3 — Disposable multi-instance demo stacks (`anthropos-demo`)

## Goal
Spin up `demo-1`, `demo-2`, … as **isolated, full platform stacks on one box** — each Clerkenstein-wired so
it comes up auth-working with **no real Clerk** — and kill any one cleanly, **without modifying a single
read-only platform file**. This is the first buildable piece of v1.1: an empty-but-running disposable stack.
M4 fills it with data; M5 makes it repeatable.

## Why this is buildable now (the grounded crux)
Research (2026-06-03, against the real `anthropos-dev/platform`) found exactly why N concurrent stacks are
impossible today — and each blocker has a concrete, additive fix:

| Blocker (real file) | Why it collides | M3's additive fix |
|---|---|---|
| **24 hard-coded host ports** (`docker-compose.yml` + `common.yml`, zero env indirection) | a 2nd `make up` binds the same ports | a generated compose **override** remaps all 24 via a **port-offset** (`demo-N → base + N·100`) |
| **One fixed `COMPOSE_PROJECT_NAME=anthropos`** (`.env:22`; no `-p` in the Makefile) | container/network name clashes | a per-stack project name `demo-N` (`-p demo-N` / generated `.env.demo-N`) |
| **One relative Postgres bind-mount** `./data/postgresql` (`common.yml:7-8`; no named volumes) | two stacks share one database | per-stack data dir under `anthropos-demo/stacks/demo-N/` |
| **Clerkenstein injection is "recipe-only"** (v1.0 left live wiring to v1.1) | a fresh stack still wants real Clerk | M3 **wires the 4 injection recipes live, by default** |

Because every collision point is a known file fact, the `In:` list is writable with confidence → **section**
shape (not iterative). This milestone is **orchestration glue around the existing compose/Makefile/repos.yml**
— never a rebuild of them.

## Scope
### In
- **The `anthropos-demo/` stack layout** (already gitignored, already holds `clerkenstein/`): per-demo
  workspace at `anthropos-demo/stacks/demo-N/` — **per-demo service-repo clones** (decision M3-D1), the
  generated override + env + data dir.
- **Per-demo service-repo clones** (M3-D1): each `demo-N` re-clones the platform service repos under its own
  `stacks/demo-N/` (full filesystem isolation; Clerkenstein injection applied per-clone). Disk cost accepted
  for isolation + future per-demo config divergence.
- **A generated compose override** (`docker-compose.demo.yml`, applied `docker compose -f <base> -f <override>`)
  that remaps the 24 host ports by the stack's offset and points volumes at the per-demo data dir. The
  read-only base compose is **never edited**.
- **A deterministic port-offset scheme** (`demo-N → base + N·OFFSET`, default OFFSET=100) + a **max-N** chosen
  so ranges never overlap the 24 mappings or the ephemeral range.
- **A per-stack `.env.demo-N`** generated from a tracked template: inherits the shared platform `.env`, then
  overrides `COMPOSE_PROJECT_NAME`, the offset-derived ports, and the Clerkenstein endpoints.
- **Clerkenstein wired in by default** for every demo — the four live-injection recipes (first live wiring;
  v1.0 left them recipe-only): `authn` via `go.mod replace` + skip-worktree on the per-demo clone; `clerk-backend`
  via the `api.clerk.com` → fake-BAPI redirect inside the stack (the **one genuinely novel** seam — see Risks);
  `clerk-frontend` via a minted `NEXT_PUBLIC_/VITE_` publishable key pointing the browser at the fake FAPI;
  `clerk-webhook` via svix-signed POST to `/api/webhook/clerk`.
- **Lifecycle skills**: `/demo-up [N]` (clone → generate override+env → apply Clerkenstein injection →
  `docker compose -p demo-N` up + migrate → optional bootstrap-dev), `/demo-down [N]` (stop + remove that
  demo's containers/network + data dir — **manual teardown only**, decision M3-D2), `/demo-status` (live demos,
  ports owned, health).
- **A demo-stack registry/ledger** (`anthropos-demo/stacks/registry.{json,md}`): which N are live, ports owned,
  started-at — so `/demo-status` + teardown are reliable and offset assignment is collision-free.
- **Delivers** the net-new ops guide **`corpus/ops/demo_stacks.md`**: the lifecycle (spin up / kill), the
  port-offset table, the Clerkenstein-by-default wiring, and the per-demo-clone layout.

### Out
- Rich/declarative data seeding (`demo.seed.yaml` + the multi-store seeder) — **M4**.
- Curated use-case recipes + the 200/500/1k seed presets + the demo corpus index — **M5**.
- AI-generated transcripts/embeddings — M4 hard-line deferral / M5 stretch.
- **Any change to the read-only platform repos** (compose/Makefile/repos.yml under `anthropos-dev/`) — M3 is
  strictly an additive overlay under `anthropos-demo/`.
- External shareability (Tailscale-only vs public ingress) beyond what staging already documents — a v1.1 open
  decision, not M3 scope.

## Depends on
v1.0 shipped (M1/M2/M2c): Clerkenstein exists at `anthropos-demo/clerkenstein` (100%/100% on all three
surfaces) with the four documented injection recipes. **Parallel with:** none (M4 depends on M3; M5 on both).

## Estimated complexity
**medium** — the mechanism (compose `-f base -f override` + `--env-file` + `-p` + skip-worktree `go.mod replace`)
is all known/bounded; the volume is in generating + wiring it cleanly per stack and proving the Docker-internal
Clerkenstein redirect.

## Decisions locked at design (2026-06-03)
- **M3-D1 — per-demo service-repo clones** (user-chosen): each `demo-N` clones the service repos under its own
  `stacks/demo-N/`. Full isolation + future-proofs per-demo divergence, at ~N× disk/clone-time. (Shared clones
  were the disk-cheaper alternative; rejected in favor of isolation.)
- **M3-D2 — manual teardown only** (user-chosen): `/demo-down [N]` is the only reclaim path; no nightly
  auto-reaper in M3. (Auto-reaping can be added later if forgotten stacks become a real problem.)

## Open questions (resolve during build)
- Exact **max-N** + offset size below the ephemeral-port range without overlapping the 24 base mappings.
- **Box capacity**: a full stack is ~10–12 GB RAM (staging-bringup) — how many demos realistically co-reside;
  should `/demo-up` warn/refuse past a memory budget?
- The **Clerkenstein `clerk-backend` DNS/cert redirect inside Docker**: how each demo's `app` container trusts
  `api.clerk.com` → fake-BAPI (`extra_hosts` + a CA the container trusts, vs a base-URL env override if one
  exists). This is the load-bearing unknown — see Risks.
- `next-web-app` bakes `NEXT_PUBLIC_*` (incl. the Clerk publishable key + endpoint) at **build** time → the
  per-demo clone must rebuild the frontend with the demo's minted key, not just re-env at runtime.
- Whether to **pick up the v1.0 express-gate CI carry-forward** here (a demo stack materializes
  `node_modules/@clerk/express`) or leave it to M5. **Default: M5.**

## KB dependencies (read as contract)
- `corpus/ops/platform_repo.md` (Make targets, profiles, compose, repos.yml)
- `corpus/ops/run_guide.md` + `corpus/ops/quick_ops.md` (startup sequence; the `-p ant-rosetta` project-scoping convention)
- `corpus/ops/staging-bringup.md` (resource footprint, profiles)
- `anthropos-demo/clerkenstein/knowledge/injection.md` (the four injection recipes — the contract M3 wires live)
- `anthropos-demo/clerkenstein/knowledge/architecture.md` (the fake FAPI/BAPI endpoints)

## Risks
- **(load-bearing) The Clerkenstein `clerk-backend` DNS/cert redirect has never been wired into a live Docker
  stack** — it's the only injection seam that isn't a config string. Mitigation: spike it first (an M3 early
  section), fall back to a base-URL env override if the DNS/cert path is too fragile inside containers.
- **Frontend build-time env**: `next-web-app` bakes the Clerk publishable key at build → per-demo rebuild
  required; budget for it.
- **Box capacity / disk** (amplified by M3-D1 per-demo clones): N full stacks + N repo clone-sets. Mitigate
  with a memory/disk budget check in `/demo-up` + the registry.

## Exit (section)
All `In:` deliverables land + `progress.md` checklist complete: two demo stacks (`demo-1`, `demo-2`) run
concurrently, isolated, each Clerkenstein-wired (browser login works, backend verifies, no real Clerk), and
`/demo-down` cleanly reclaims one without touching the other — documented in `corpus/ops/demo_stacks.md`.
