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

## S3 — Clerkenstein injection: WIRING SCAFFOLD only — NOT verified on live services (honest correction 2026-06-03)
> **Correction.** The original S3 checkmarks overstated this. `inject` *emits artifacts*; only the
> publishable-key recipe is actually proven. The injection was **never run against a live Clerk consumer** —
> the demo-1 live proof was infra-only (postgres+redis, no Clerk consumer). A direct attempt to bring up the
> `app` (the Clerk consumer) revealed **two hard blockers** (see retro): (1) `app` has hard `depends_on`
> skiller/skillpath/… so it needs the **full `graphql` profile** (the ~10-12 GB stack — doesn't fit here);
> (2) the `authn` recipe's go.mod-replace needs an **assembled "patched colony" module that does not exist** —
> clerkenstein ships the authn twin *package*, not a colony *module* to replace `colony` with.
- [x] `clerk-frontend` minted publishable key → env — **PROVEN** (byte-identical to clerkenstein's gated `MintPublishableKey`; round-trip). This is the only recipe whose *mechanism* is verified.
- [~] `clerk-backend` api.clerk.com → fake-BAPI — the `app` `extra_hosts` `!override` snippet is *emitted*; **NOT verified** (needs the app running + a trusted cert).
- [~] `authn` go.mod-replace — the *directive text* is emitted; **NOT buildable yet** (no patched-colony module to point at — real integration work).
- [~] `clerk-webhook` svix-signed POST — the injector invocation is *documented*; **NOT run** (needs a live `/api/webhook/clerk`).
- **Carried forward (M3-CF1):** assemble the patched-colony module + verify all four recipes end-to-end on the full stack (bigger Docker VM). This is genuine integration work, not just a hardware wait.

## S4 — lifecycle skills + teardown (M3-D2 manual only) ✅
- [x] `/demo-up [N]` skill (wraps `demo-stack clone → inject → up`, resource-aware)
- [x] `/demo-down [N]` skill (wraps `demo-stack down N --purge`; `-p`-scoped, dev stack untouched — proven)
- [x] `/demo-status` skill (wraps `demo-stack status`; registry + per-demo `ps` + resolved refs)

## S5 — the ops guide + the acceptance demo ✅
- [x] `corpus/ops/demo_stacks.md` (collision problem + additive fix + `!override` + port-offset + clone-at-tag + Clerkenstein recipes + safety + resource budget + proven-vs-gated split); cross-linked from `corpus/ops/README.md`
- [x] **acceptance (M3-D5):** demo-1 (postgres+redis) ran isolated alongside the dev stack on offset ports with its own data; up→status→down; **dev stack untouched throughout**. (Two-concurrent-full-stack acceptance is resource-gated → bigger box.)
