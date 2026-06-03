# M3 — progress (section checklist)

**Milestone:** M3 — Disposable multi-instance demo stacks · **Shape:** section · **Status:** planned

Sections derived from `overview.md` Scope.In (check off as `/developer-kit:build-milestone` lands each):

## S1 — `anthropos-demo/stacks/` layout + per-demo clone mechanism (M3-D1)
- [ ] `anthropos-demo/stacks/demo-N/` workspace layout
- [ ] per-demo service-repo clone step (clone the platform service repos into the stack dir)
- [ ] the stack registry/ledger (`registry.{json,md}`: live N, ports owned, started-at)

## S2 — compose override + port-offset + per-stack project/env/data
- [ ] generated `docker-compose.demo.yml` override (remaps the 24 ports; points volumes at the per-demo data dir)
- [ ] port-offset scheme (`demo-N → base + N·100`) + max-N + collision-free assignment via the registry
- [ ] `.env.demo-N` template + generation (project name, offset ports, Clerkenstein endpoints)
- [ ] per-stack Postgres data dir isolation

## S3 — Clerkenstein live injection (first live wiring — the 4 recipes)
- [ ] `authn` go.mod replace + skip-worktree on the per-demo clone
- [ ] `clerk-backend` api.clerk.com → fake-BAPI redirect **inside Docker** (the novel seam — spike first)
- [ ] `clerk-frontend` minted publishable key (frontend rebuilt with the demo's key — build-time env)
- [ ] `clerk-webhook` svix-signed POST to /api/webhook/clerk

## S4 — lifecycle skills + teardown (M3-D2 manual only)
- [ ] `/demo-up [N]` (clone → generate → inject → up + migrate → optional bootstrap-dev)
- [ ] `/demo-down [N]` (stop + remove containers/network + data dir — manual only)
- [ ] `/demo-status` (live demos, ports, health)

## S5 — the ops guide + the acceptance demo
- [ ] `corpus/ops/demo_stacks.md` (lifecycle + port-offset table + Clerkenstein-by-default + per-demo-clone layout)
- [ ] **acceptance:** `demo-1` + `demo-2` run concurrently, isolated, both Clerkenstein-wired (browser login works,
      no real Clerk); `/demo-down 1` reclaims demo-1 without touching demo-2
