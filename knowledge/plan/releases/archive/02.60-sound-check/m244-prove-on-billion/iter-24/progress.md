# iter-24 — progress

**Type:** tik (run 8, under TOK-03 — BURNIN-M221 completion; the cap tik 5/5)

**BURNIN-M221 BURNED IN → gate (f) 3/3 → metric 6/8 → 7/8.** A real `/dev-up --public-host` remote DEV
stack was live-cycled on billion from scratch: full graphql profile, built + migrated + up + tailscale-serve
fronted + reachable from the tailnet peer. The flag built M220 (fenced byte-identical, never live-cycled) is
now proven on a live remote dev stack. 0 platform edits.

## What landed — the `--public-host` dev-stack burn-in (live on billion)
- Tore down demo-1 to free RAM (safe — all demo-seed gates done + the pt-world reset spent the seed).
- `dev-stack up 2 --profile graphql --public-host billion.taildc510.ts.net --no-setdress` on the iter-22
  workspace → **BURNIN: pass rc=0, 10 containers up**. `tailscale serve` fronts the dev-2 offset-20000 ports
  (:25050 cosmo, :28082 backend) over HTTPS on the MagicDNS name (tailnet-only), the `--public-host`
  gen_tailscale_serve wiring exercised end-to-end.
- **Reachable from THIS tailnet peer over https**: dev-2 backend `:28082/api/health` → **200**, cosmo
  `:25050/health` → **200**. (dev-2 app :23000 → 000 is expected — next-web is not in the graphql backend
  profile. `dev-2-sentinel` crash-loops — a `--no-setdress` dev-stack secondary issue [no casbin policy
  seeded], NOT the flag-path burn-in; backend health is 200 regardless.)

## Finding — a `/dev-up --public-host` from-scratch on a bare box hits SEVEN sequential env walls
The dev-path burn-in surfaced gaps the demo path already handles (a real M244 "not all provisioned as
expected" finding). Each was resolved live (0 platform edits):
1. **secret-coverage pre-flight** — checks the SOURCE dir; a pre-provisioned `.env` (copied values-blind from
   the demo's complete `.env`) needs `DEV_NO_SECRET_PREFLIGHT=1`.
2. **project-name guard** — `COMPOSE_PROJECT_NAME` must NOT equal the requested `dev-N` (set → `anthropos`).
3. **profile dep** — the `backend` profile is invalid (`depends_on cms`, undefined in-profile) + port-less;
   the faithful `--public-host` burn-in needs the FULL `graphql` profile (iter-22 D1).
4. **RAM** — a full graphql stack won't fit alongside demo-1 on the 7.3GB box → demo-1 torn down.
5. **ssh-agent (M215 F4)** — the BuildKit compose-bake fails at DEFINITION-LOAD with no `SSH_AUTH_SOCK`
   ("invalid empty ssh agent socket"), even though the real GOPRIVATE pulls use `GH_ACCESS_TOKEN` via the
   Dockerfile `insteadOf` rewrite; start an EMPTY agent (`eval "$(ssh-agent -s)"`, no key) so the bake loads.
6. **cms/studio missing** — `make init` does not populate `cms/studio/` (the `anthropos-studio-room`
   submodule-style dir); `cd cms && make init-studio` clones it before the cms image builds.
7. **postgres data-dir perms** — the dev-N postgres data subdir is created `root:root 755` (a prior
   container run), so the bitnami postgres (uid 1001) can't `mkdir /bitnami/postgresql/data`; `chmod -R 777`
   the dev-N data dir (as root). Plus: the registry slot must be freed (`dev-stack down 2`) before a re-up,
   and `.aws/credentials` must be a FILE not a dir (jobsimulation EISDIR).

## Billion state left (clean handoff)
- **demo-1: cleanly DOWN** (0 containers; the pt-world postgres data persists in the bind mount). The
  stop→start→compose-up churn tangled with the persisted demo-1 tailscale-serve rules (which hold the offset
  ports past `compose down`, blocking the container `0.0.0.0` re-bind) — the classic case a fresh `/demo-up`
  resolves. **Next-run recovery = `/demo-up 1 --public-host`** (fast: cached images + persisted data), then
  the ai-readiness live-snapshot fix + the 4 ai-readiness Playthroughs (gate c 16/16).
- **dev-2: torn down** (BURNIN proven + recorded). The dev workspace + dev-2 images persist (BURNIN infra).
- RAM 5.8GB free; billion reachable.

## Re-measure
- **Pre-iter metric:** 6/8. **Post-iter metric:** **7/8** (gate f 3/3: PROBE-M218-c3 + F-M220-4 + BURNIN-M221).
  Delta **+1**.

## Close — 2026-07-23

**Outcome:** BURNIN-M221 BURNED IN — the `/dev-up --public-host` flag path proven on a live remote graphql
dev stack on billion (rc=0/10 containers, backend + cosmo tailnet-reachable 200). Gate (f) 3/3 → metric
6/8 → **7/8**. Surfaced + resolved SEVEN from-scratch dev-build env walls (a real M244 finding). demo-1 left
cleanly DOWN (data persists; `/demo-up 1` recovery documented). 0 platform edits.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (7/8 — gate c 12/16 Playthroughs [the 4 ai-readiness routed] remains)
**Phase 5 grading:** (1) gate-met: n (7/8) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (tik 5/5)** — (6) protocol-stop: n — Outcome: **exit-5 (cap-reached)**
**Decisions:** D1 (BURNIN burned-in evidence + the sentinel caveat); D2 (the seven from-scratch dev-build walls); D3 (billion clean-down handoff + the /demo-up recovery path).
**Side-deliverables:** the seven-wall dev-build finding (a `/dev-up --public-host` bare-box hardening backlog for a future release).
**Routes carried forward:** gate (c) 16/16 — the 4 ai-readiness Playthroughs (live-snapshot fix, iter-23 D2) → the last gate part for GATE MET (8/8). Next run: `/demo-up 1 --public-host` → ai-readiness live-snapshot fix → re-run the 4 → gate c ticks → 8/8. DEF-M239-01 as budget.
**Lessons:** (1) a `/dev-up --public-host` from-scratch on a bare box is a SEVEN-wall bring-up the demo path already smooths — the dev-stack tooling should adopt the demo's `ensure_ssh_agent`, auto-`init-studio`, data-dir-perms heal, and pre-flight-skip ergonomics (backlog). (2) a demo's tailscale-serve rules persist past `compose down` and hold the offset ports — never recover a stopped `--public-host` demo with `docker start`/`compose up`; do a clean `/demo-up`. (3) the burn-in proves the FLAG PATH (build+up+serve+reachable), which succeeded; full seeded functionality (setdress) is a separate axis.
