# M3 — progress (section checklist)

**Milestone:** M3 — Disposable multi-instance demo stacks · **Shape:** section · **Status:** in-progress

## Build log

**2026-06-03 — S2 (the override/isolation engine) built + LIVE-PROVEN.** Tooling in the new gitignored
`anthropos-demo/demo-stacks/` repo (M3-D4; commit `946c5ba`): `lib/gen_override.py` (offset ports +
repoint Postgres bind, via Compose `!override` so sequences are *replaced* not appended — the append would
re-bind the base port and collide with the dev stack) + the `demo-stack` lifecycle CLI (`up`/`down`/
`status`/`gen`, every op `-p demo-N`-scoped, `down` hard-refuses the dev project) + `registry.json`.
**Live proof on this 16 GB box, alongside the running 12-container dev stack:** `demo-1` (postgresql+redis)
came up on offset ports **5532/6479** with its own data dir → two independent live Postgres instances side
by side; `status` listed it; `down --purge` cleanly removed it — and the **dev `anthropos` stack stayed
fully intact** (12 containers, postgres healthy) the whole time. This satisfies the M3-D5 acceptance
("one demo stack alongside the dev stack, untouched").

**Remaining:** S1 clone-each-repo-at-its-release-tag (M3-D3); S3 the four Clerkenstein live-injection
recipes (the `app`→`api.clerk.com`→fake-BAPI redirect needs the full app stack — RAM-gated here); S4 the
`/demo-*` rosetta skill wrappers; S5 the `corpus/ops/demo_stacks.md` guide + the resource-gated full-stack
note. The full 12-service single stack + end-to-end Clerkenstein browser login verify on a bigger Docker VM.

Sections derived from `overview.md` Scope.In (check off as `/developer-kit:build-milestone` lands each):

## S1 — `anthropos-demo/stacks/` layout + per-demo clone mechanism (M3-D1)
- [x] `anthropos-demo/demo-stacks/stacks/demo-N/` workspace layout
- [ ] per-demo service-repo clone step — clone each repo at its **latest release tag** (caller-overridable ref; fall back to `main` only if untagged), into the stack dir (M3-D1 + M3-D3) — _not yet built (this turn used the shared anthropos-dev clones for the live proof)_
- [x] the stack registry/ledger (`registry.json`: live N, ports owned, profile/services) — _resolved-ref-per-repo field lands with the clone step_

## S2 — compose override + port-offset + per-stack project/env/data ✅ (live-proven)
- [x] generated `docker-compose.demo.yml` override (`!override` remaps ports; repoints the Postgres bind)
- [x] port-offset scheme (`demo-N → base + N·100`) — _max-N bound still open (resolve in S5)_
- [x] `.env.demo-N` template + generation (project name, offset) — _Clerkenstein endpoint vars land with S3_
- [x] per-stack Postgres data dir isolation

## S3 — Clerkenstein live injection (first live wiring — the 4 recipes)
- [ ] `authn` go.mod replace + skip-worktree on the per-demo clone
- [ ] `clerk-backend` api.clerk.com → fake-BAPI redirect **inside Docker** (the novel seam — spike first)
- [ ] `clerk-frontend` minted publishable key (frontend rebuilt with the demo's key — build-time env)
- [ ] `clerk-webhook` svix-signed POST to /api/webhook/clerk

## S4 — lifecycle skills + teardown (M3-D2 manual only)
- [~] `/demo-up [N]` — the `demo-stack up` CLI works (generate → up, `-p demo-N`); the rosetta skill wrapper + the clone/inject/migrate steps are pending (S1/S3)
- [x] `/demo-down [N]` — `demo-stack down N --purge` proven (stop + remove containers/network + data dir, `-p`-scoped; dev stack untouched); rosetta skill wrapper pending
- [x] `/demo-status` — `demo-stack status` proven (registry + per-demo `ps`); rosetta skill wrapper pending

## S5 — the ops guide + the acceptance demo
- [ ] `corpus/ops/demo_stacks.md` (lifecycle + port-offset table + Clerkenstein-by-default + per-demo-clone layout)
- [ ] **acceptance:** `demo-1` + `demo-2` run concurrently, isolated, both Clerkenstein-wired (browser login works,
      no real Clerk); `/demo-down 1` reclaims demo-1 without touching demo-2
