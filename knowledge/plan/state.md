# State

**Active version:** **v1.8 "understudy" — IN DEVELOPMENT** (designed 2026-06-15; branch `release/01.80-understudy`).
The **self-contained-demo release**: give `stack-demo/` its **own platform clone set** so a box with only `stack-demo/`
(no `stack-dev/`) can run a demo end-to-end. Closes a live doc-vs-code gap — `CLAUDE.md` already calls `stack-demo` *"a
true peer with its own clone set,"* and M30 already provisions `stack-demo/platform/.env`, but `up-injected.sh` still
builds every image from `stack-dev` (`src="$DEV/$svc"`, `PLAT="$DEV/platform"`). **Re-implements** the orphaned
`m26/self-contained-demo` branch (@ `25ab855`, the spec — predates v1.6/v1.7 so unmergeable) onto current `main`,
preserving the stack-secrets module + M30 provision + M31 mkcert + M32 studio-desk fix. **Tooling + docs only — zero
platform-repo edits** (`stack-demo/platform` is a build *context* only).
**Active milestone:** **M26 — Self-contained demo stacks** (`built`, awaiting close; single `section` milestone of
v1.8). All 7 sections landed on `m26/self-contained-demo` (rosetta doc-half) + ext `m26/self-contained-demo-reimpl`
@ `17971c1` (tagged `understudy-m26`): a demo builds entirely from `stack-demo`'s OWN clone set
(`ensure-clones.sh` + the `$DEV`→`stack-demo` build-source repoints + the `reuse_dev_images` opt-in gate),
preserving M30/M31/M32. demo-stack 123/123 + stack-injection 113/113 pass; PR review CLEAN. Close with
[`/developer-kit:close-milestone`](roadmap.md) (field-bake on a freshly-emptied `stack-demo/` + ff-to-release).
**Last closed:** **v1.7 "house lights" — 2026-06-15**, tag `v1.7`. A **demo-UI-hardening release**: a fresh browser at a
demo's offset UI renders the working app with **zero manual steps**. Triggered live — next-web at `http://localhost:33000`
(demo-3) showed a **blank page** (clerk-js's handshake to the fake FAPI hit an untrusted self-signed cert) and studio-desk
302'd to a dead `:9100`. **M31** automated a locally-trusted **mkcert** FAPI cert into the demo bring-up; **M32** fixed
the studio-desk `:9100`-dead-redirect (a `NODE_ENV=production` override) + the `:9100` doc/CORS sweep. close-release
**GREEN** (all 9 sweeps, 0 blocking). **Tooling + docs only.**
**Next up:** **close M26** via `/developer-kit:close-milestone` — field-bake on a **freshly-emptied** `stack-demo/`
(the on-disk one is already populated from the orphan run + would mask a from-scratch failure), then ff
`m26/self-contained-demo` → `release/01.80-understudy` + the orchestrator's ext ff-to-main + `understudy-m26`
tag-repoint. **At close:** delete the orphan tag `prop-room-m26` + branch `m26/self-contained-demo` (the ext
authoring branch `m26/self-contained-demo-reimpl` supersedes it). (Outward-facing carry-over: push the ext tags
`understudy-m26` + `house-lights-m31`/`m32` + `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` to `origin`;
`wip/clerkenstein-browser-login` still awaits its own design-roadmap pass.)
**Phase:** **v1.8 — M26 `built` (all 7 sections + ext tag `understudy-m26`), awaiting `/developer-kit:close-milestone`.**
**Paused:** _(none)_

## Recently shipped releases
- **v1.7 "house lights"** — **2026-06-15**, tag `v1.7`. The **demo-UI-hardening release**: a fresh browser at a demo's
  offset UI renders with zero manual steps. **M31** auto-mints a locally-trusted **mkcert** FAPI cert at bring-up
  (openssl fallback + `DEMO_NO_MKCERT` opt-out) so clerk-js's handshake stops hitting an untrusted cert + next-web stops
  blanking; **M32** pins `NODE_ENV=production` on the studio-desk override (the additive env-merge had let base
  `development` survive → a dead-`:9100` redirect) → the single-port production `sendFile` path, + the `:9100` doc/CORS
  sweep. Go **1027** unchanged (no Go touched); Python 459→**471** (+12: M31 `FapiCertStep` +11, M32 regression +1);
  flake **0**; supply-chain GREEN (0 new deps — shell + python-stdlib + docs). close-release GREEN (all 9 sweeps; 0
  blocking). Code: `rosetta-extensions` @ tags `house-lights-m31`/`m32`. Records:
  [releases/archive/01.70-house-lights/](releases/archive/01.70-house-lights/) (review · retro · metrics · lockfile).
- **v1.6 "stage door"** — **2026-06-14**, tag `v1.6`. The **secret-provisioning release**: ingest a secret source
  (dir/zip, default `.agentspace/secrets`) → **provision every repo's `.env`** (values-blind) → a **secret-coverage DNA**
  (6 repos / 55 genes; `introspect`+`diff` keep-listed gate) → demo-aware `check` in `/dev-up`+`/demo-up` pre-flight →
  the `/stack-secrets` skill. M27→M30 (field-bake proven LIVE on demo-3). Go 867→**1027** (+160, the stdlib-only
  `stack-secrets` section); Python **459**; flake **0**; supply-chain GREEN. Records:
  [releases/archive/01.60-stage-door/](releases/archive/01.60-stage-door/).
- **v1.5 "prop room"** — **2026-06-14**, tag `v1.5`. The **local-Directus release**: stacks serve their own captured
  public catalog locally on `--local-content`. M21→M25 (field-bake on a 16 GB box). Go 736→**867**; Python 360→**459**;
  flake **0**. Records: [releases/archive/01.50-prop-room/](releases/archive/01.50-prop-room/).

## Headline numbers (v1.7 close — 2026-06-15; baseline v1.6 2026-06-14)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`), **unchanged across v1.7** (neither M31 nor M32 touched Go). Per-module:
  `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding **259** · stack-snapshot **333** ·
  stack-secrets **160**. `go vet` + `gofmt` + `shellcheck` clean; flake **0**.
- **Python tests:** **471** (v1.6 459 → +12: M31 +11 `FapiCertStep` in `demo-stack/tests/test_tooling.py`; M32 +1
  `test_studio_desk_env_pins_node_env_production` in `stack-injection/tests/test_injection.py`). `test_injection.py`
  suite 88/88 (0 skipped under PyYAML, authoritative JUnit tally); flake **0** (5/5 randomized, both milestones). The
  corpus README-index guard runs **exit 0**.
- **The v1.7 thesis:** a fresh browser at a demo's offset UI renders the working app **with zero manual steps** — the
  blank next-web page (an untrusted FAPI cert) and the studio-desk dead-`:9100` redirect are both fixed at bring-up; a
  fresh `/demo-up` is browser-clean (on a local same-machine demo). Tooling + docs only.
- **Safety:** **values-blind** demo cert handling (no cert/key body to stdout/log; only key NAMES/paths); the prod
  `DIRECTUS_TOKEN`-non-rearm + secret values-blind invariants (v1.6) carry forward unchanged.
- **Supply-chain:** **GREEN** — v1.7 added **zero** third-party deps (M31 pure shell + stdlib tests; M32 stdlib
  generator + tests; rosetta docs-only; no manifest change); the v1.6 stdlib-only posture carries forward.
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces — v1.7 touched no clerkenstein behavior.

## Branch model
**v1.7 SHIPPED:** `release/01.70-house-lights` merged `--no-ff` → `main`, tagged **`v1.7`** (2026-06-15); release branch
deleted; both milestone branches (`m31/mkcert-fapi-cert` + `m32/studio-desk-singleport`) merged + deleted. The fixes live
in the `demo-stack` + `stack-injection` ext sections (authored in `.agentspace/rosetta-extensions/`, consumed per-stack at
a pinned tag). v1.7 ext markers: **`house-lights-m31`** @ `5022e72` · **`house-lights-m32`** @ `7b17c39` (ext head
`7b17c39`, final). The orphaned **`m26/self-contained-demo`** branch (tag `prop-room-m26`) + `wip/clerkenstein-browser-login`
stay preserved on the ext side, awaiting their own design-roadmap pass.
**v1.6 SHIPPED:** `release/01.60-stage-door` → `main`, tag **`v1.6`** (2026-06-14); ext markers `stage-door-m27`/`m28`/`m30`.
**Prior:** **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-15 (**v1.8 "understudy" DESIGNED** via `/developer-kit:design-roadmap` — the self-contained-demo
release, a single `section` milestone **M26** that re-implements the orphaned `m26/self-contained-demo` branch [@ `25ab855`,
unmergeable-stale] onto current `main`, preserving v1.6/v1.7. Port spec verified by a 3-agent fan-out + adversarial
no-regression review [workflow `wf_212f3442-44e`]: all 12 orphan files covered, no M30/M31/M32 revert, design decisions
D-MAIN + D1–D6 settled. Branch `release/01.80-understudy` cut; M26 dir scaffolded; M26 `planned`. Prior: 2026-06-15
**v1.7 "house lights" SHIPPED** [tag `v1.7`, 9-sweep close-release ALL GREEN]; 2026-06-14 v1.6 "stage door" SHIPPED.)_
