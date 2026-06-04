# M3 — progress (section checklist)

**Milestone:** M3 — Disposable multi-instance demo stacks · **Shape:** section · **Status:** done (2026-06-03)

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

**2026-06-03 — M3 sections built + hardened; S3 corrected (see below).** Tooling in `anthropos-demo/demo-stacks/`
(commits `946c5ba` S2 · `cda2db3` S1 · `31bdcd8` S3 · `b626020` harden) + rosetta `/demo-*` skills + the
ops guide. 12 tooling unit tests green; shellcheck + python compile clean. **Genuinely proven:** S1
(clone-at-tag), S2 (the override/isolation engine — demo-1 live alongside the dev stack, M3-D5), the
publishable-key mint, S4 skills, S5 guide. **Corrected (S3):** the Clerkenstein injection is **wiring scaffold
only, NOT verified on live services** — a direct attempt to run the `app` exposed two hard blockers (app
needs the full graphql profile; no patched-colony module exists for the authn replace). Reframed below +
routed to M3-CF1.

## S1 — layout + per-demo clone-at-release-tag (M3-D1 + M3-D3) ✅
- [x] `anthropos-demo/demo-stacks/stacks/demo-N/` workspace layout
- [x] per-demo clone step — `clone_repos.py` + `demo-stack clone`; each repo at its latest release tag (caller-overridable; bare + `v`-semver; main if untagged) — resolution proven on all 14 repos + real clones
- [x] the stack registry/ledger (`registry.json`: live N, ports, profile/services, **resolved ref per repo**)

## S2 — compose override + port-offset + per-stack project/env/data ✅ (live-proven)
- [x] generated `docker-compose.demo.yml` override (`!override` remaps ports; repoints the Postgres bind)
- [x] port-offset scheme (`demo-N → base + N·100`) — _max-N bound documented as a tuning knob in the guide_
- [x] `.env.demo-N` template + generation (project name, offset, + Clerkenstein endpoint vars via S3)
- [x] per-stack Postgres data dir isolation

## S3 — Clerkenstein injection: authn + frontend BUILT + PROVEN; backend/webhook emitted; running-app deployment = M3-CF1
> **Two corrections, in order (2026-06-03).** First I over-claimed (checkmarks = "wiring emitted" ≠ "works").
> Then, pushed to actually do it, I found my "resource-gated" reason was **wrong**: the *whole dev stack* uses
> **~0.9 GB** (measured), not 10-12 — RAM was never the blocker. And the "patched colony doesn't exist" was
> "I hadn't built it." So I **built it and proved it**.
- [x] `authn` — **BUILT + PROVEN.** A disarmed colony clerk provider (`demo-stacks/inject/colony-authn-disarmed/clerk.go`: same package/type/`NewProvider` signature, universal-key verify) vendored via `apply-authn.sh` (clone colony@pinned → swap clerk pkg → `replace => ./vendor-colony`). **Proven against colony v0.34.3 (app's pinned version):** the disarmed provider accepts a Clerkenstein token + extracts identity+org; rejects garbage/expired. App code unchanged. (`inject_proof_test.go`.)
- [x] `clerk-frontend` minted publishable key → env — **PROVEN** (byte-identical to clerkenstein's gated `MintPublishableKey`).
- [x] **running-app deployment — PROVEN end-to-end (M3-CF1 RESOLVED, 2026-06-04).** Built a demo `app` image with the vendored disarmed colony (`apply-authn.sh` + a Dockerfile fix), ran it live, and hit a protected route (`/api/workforce/members`) with three tokens: **none → 400, garbage → 401 (rejected), Clerkenstein → 403 (ACCEPTED — past authn, denied at authz)**. The 403-not-401 is the proof: a real running Anthropos `app` accepts a Clerkenstein universal-key token at its live HTTP auth middleware. Recipe + result: `demo-stacks/inject/DEPLOYMENT-PROOF.md`.
- [~] `clerk-backend` api.clerk.com → fake-BAPI — `app` `extra_hosts` `!override` snippet emitted; the cert/redirect is the one remaining recipe to verify live (the authn path is fully proven).
- [~] `clerk-webhook` svix-signed POST — injector invocation emitted; run when the webhook flow is exercised.
- **M3-CF1 RESOLVED.** The headline ("a demo comes up Clerk-free, accepting Clerkenstein tokens") is now **demonstrated on a live app**, not just scaffolded. RAM was never a blocker (the whole dev stack uses ~0.9 GB — measured).

## S4 — lifecycle skills + teardown (M3-D2 manual only) ✅
- [x] `/demo-up [N]` skill (wraps `demo-stack clone → inject → up`, resource-aware)
- [x] `/demo-down [N]` skill (wraps `demo-stack down N --purge`; `-p`-scoped, dev stack untouched — proven)
- [x] `/demo-status` skill (wraps `demo-stack status`; registry + per-demo `ps` + resolved refs)

## S5 — the ops guide + the acceptance demo ✅
- [x] `corpus/ops/demo_stacks.md` (collision problem + additive fix + `!override` + port-offset + clone-at-tag + Clerkenstein recipes + safety + resource budget + proven-vs-gated split); cross-linked from `corpus/ops/README.md`
- [x] **acceptance (M3-D5):** demo-1 (postgres+redis) ran isolated alongside the dev stack on offset ports with its own data; up→status→down; **dev stack untouched throughout**. (Two-concurrent-full-stack acceptance is resource-gated → bigger box.)

## Migrate step (2026-06-04) — sentinel healthy + /api/health 200; authorized 200s = M4 seed
`/demo-up` now runs `migrate-demo.sh`: creates the schemas (sentinel/cms/jobsimulation/skiller/skillpath +
extensions) + the pgvector/pg_trgm/pgcrypto extensions, atlas-migrates the 5 services against the demo DB,
restarts sentinel+backend. **Result:** sentinel stops crash-looping (it needed its `sentinel` schema for
casbin) — **0 restarts, healthy**; `/api/health → 200`; 6/6 schemas migrated. **Still 403 on *authorized*
endpoints** (e.g. /api/workforce/members) — that needs the **M4 seed** (casbin policies + the demo user/org
matching the Clerkenstein universal identity), not the migrate step. Found an M4 nuance: `init_policy.sql`
seeds `casbin_rules` (plural) but the gorm adapter auto-creates `casbin_rule` (singular).
