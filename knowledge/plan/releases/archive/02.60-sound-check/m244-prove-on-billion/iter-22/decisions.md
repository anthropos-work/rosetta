# iter-22 — decisions

## D1 — BURNIN-M221 needs the FULL graphql profile + demo-1 torn down (reduced-profile plan falsified)
iter-21 D2 planned a reduced-profile `dev-stack up 2 --profile backend --public-host` to fit RAM alongside
demo-1. Executing it falsified that plan:
- `docker compose config` REFUSES the `backend` profile — `service "backend" depends on undefined service
  "cms"`. The backend (app) profile is not self-contained; its `depends_on cms` is undefined unless cms is
  also selected, and the chain pulls toward the full set.
- Even if it configured, the backend profile *"published NO ports — fronting nothing"* — so the whole point
  of `--public-host` (front the published browser-facing ports via `tailscale serve`) has nothing to act on.
⇒ A faithful `--public-host` dev burn-in needs a **FULL (graphql) profile** — all backend services + cosmo,
publishing ports. That won't fit billion's 7.3GB RAM alongside demo-1 (~4.3GB; ~3.1GB free). So BURNIN needs
**demo-1 torn down first** (freeing ~4.3GB → a full graphql dev-2 fits alone). Tearing down demo-1 is SAFE:
every demo-seed gate (a/b/d/e/g/h + gate-f's demo-side carries PROBE-M218-c3 + F-M220-4) is DONE + recorded,
and the gate-c playthroughs reset-to-seed replaces the demo seed regardless.

**This is NOT a blocker.** BURNIN is feasible on billion — tooling present, disk fine, RAM fits once demo-1
is down, the dev workspace is already built (step A). It is heavy (a full dev build ~20–40 min, backgroundable)
so it routes to a dedicated iter. Sequencing refinement under TOK-03: do the lower-risk **playthroughs first**
(iter-23, gate c → 7/8, deterministic harness) to lock in a gate tick, then BURNIN completion (iter-24) —
both destroy the demo seed anyway, so the demo-1 teardown BURNIN needs is free once playthroughs have run.

## Enabling infra built (durable)
`/home/devops/panorama/stack-dev/` = platform + 10 sibling repos (SSH-cloned) + a values-blind dev `.env`
(file→file copy of the demo's complete working `.env`; project name → `anthropos`; no value entered agent
context). `make init` complete. iter-24 runs `dev-stack up 2 --profile graphql --public-host` against it.

## Values-blind + safety notes
- The dev `.env` was provisioned file→file from the demo's `.env`; `DEV_NO_SECRET_PREFLIGHT=1` was used only
  because the target `.env` was already complete (the pre-flight checks the SOURCE dir, not the target).
- billion left clean after the failed up attempts: 0 dev-2 containers, demo-1 still green (17/17). No mid-build
  was killed (the up self-exited at compose-config validation before starting anything).
