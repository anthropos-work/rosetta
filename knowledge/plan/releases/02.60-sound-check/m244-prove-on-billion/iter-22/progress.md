# iter-22 — progress

**Type:** tik (run 8, under TOK-03 — BURNIN-M221, the last gate-f carry)

Attempted the from-scratch reduced-profile `/dev-up 2 --public-host` burn-in on billion. Built the dev
workspace (enabling infra) and **falsified the reduced-profile hypothesis**: a faithful `--public-host` dev
burn-in needs a FULL (graphql) profile, which won't fit billion's RAM alongside demo-1 → needs demo-1 torn
down first (safe — all demo-seed gates done). BURNIN completion routed to iter-24 (after the lower-risk
playthroughs lock in gate c). No dev stack came up → gate (f) stays 2/3 → metric stays 6/8.

## What landed (durable, on billion)

### Step A — the dev workspace (BUILT, enabling infra for iter-24)
`/home/devops/panorama/stack-dev/` now holds **platform + 10 sibling repos** (app, cms, sentinel,
jobsimulation, skillpath, storage, roadrunner, messenger, next-web-app, studio-desk, graphql-wundergraph),
cloned via SSH (billion has `git@github.com` access), + a **values-blind dev `.env`** (file→file copy of the
demo's complete working `.env`; secrets never entered agent context). `make init` completed. This is durable
BURNIN-prep infra iter-24 consumes.

## What was FALSIFIED (the reduced-profile hypothesis)
`dev-stack up 2 --profile backend --public-host billion.taildc510.ts.net --no-setdress` failed three ways,
each a learning that characterizes the real requirement:
1. **secret-coverage pre-flight** flagged the `.agentspace/secrets` SOURCE short on some keys — a red herring
   (the dev `.env` was provisioned from the demo's complete working `.env`, not from the source). Cleared
   with `DEV_NO_SECRET_PREFLIGHT=1` (the tooling's sanctioned skip for a pre-provisioned `.env`).
2. **project-name guard** ("dev-2 equals the main dev project 'dev-2' — would touch the dev stack"): my dev
   `.env` set `COMPOSE_PROJECT_NAME=dev-2`, colliding with the requested `dev-2`. Fixed → `anthropos` (the
   platform default main-dev name); the additional `dev-2` is then distinct.
3. **THE FALSIFIER** — `gen_override: docker compose config failed: service "backend" depends on undefined
   service "cms": invalid compose project`, and *"the override published NO ports — fronting nothing."* The
   `backend` profile is NOT self-contained (`backend depends_on cms`, undefined in-profile) AND publishes no
   browser-facing ports — so `--public-host` would front **nothing**. ⇒ a faithful `--public-host` dev
   burn-in needs a **FULL (graphql) profile** (all backend + cosmo, publishing ports).
- **RAM**: a full graphql dev stack (~2.5–3.5GB) will NOT fit alongside demo-1 (~4.3GB) on the 7.3GB box
  (~3.1GB available). ⇒ BURNIN needs **demo-1 torn down first** (safe — every demo-seed gate a/b/d/e/g/h +
  gate-f demo carries is DONE + recorded; the playthroughs reset-to-seed replaces the demo seed anyway).
- billion left **clean**: no dev-2 containers (the up failed at config before starting anything), demo-1
  still green (17/17). The workspace persists.

## Re-measure
- **Pre-iter metric:** 6/8. **Post-iter metric:** **6/8** (no dev stack came up; gate f stays 2/3). Delta 0.

## Close — 2026-07-23

**Outcome:** BURNIN-M221 not burned in this iter — complete investigation. Dev workspace BUILT (enabling
infra); the reduced-profile-fits-alongside-demo hypothesis FALSIFIED (backend profile invalid: `depends_on
cms` + publishes no ports; a faithful `--public-host` burn-in needs the FULL graphql profile, which needs
demo-1 torn down for RAM — safe). Feasible + characterized, NOT a blocker. Routed to iter-24 (after the
lower-risk playthroughs). Metric 6/8. 0 platform edits, 0 code changes.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET (6/8)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (BURNIN is feasible + characterized, routed Fate-3 — not a platform-edit / infeasible blocker) — (5) cap-reached: n (tik 3/5 this run) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (the full-profile + demo-1-teardown requirement, falsifying the reduced-profile plan).
**Side-deliverables:** the dev workspace at `/home/devops/panorama/stack-dev/` (platform + 10 repos + dev .env) — enabling infra for iter-24's BURNIN completion.
**Routes carried forward:** iter-23 = the 16 Playthroughs (gate c → 7/8, lower-risk, lock-in first). iter-24 = BURNIN completion (tear down demo-1 → `dev-stack up 2 --profile graphql --public-host` on the built workspace → verify → gate f → 8/8). DEF-M239-01 as budget.
**Lessons:** a "reduced profile" is not free — compose `depends_on` can make a nominally-minimal profile (backend) invalid + port-less, so a `--public-host` burn-in that must FRONT ports needs the full profile. Assess the compose dependency graph + published-port set BEFORE assuming a light profile suffices.
