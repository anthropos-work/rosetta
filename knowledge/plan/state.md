# State

**Active version:** _(between releases — **v1.7 "house lights" SHIPPED 2026-06-15**, tag `v1.7`. The next version is
unplanned; run [`/developer-kit:design-roadmap`](roadmap.md) to scope it.)_
**Active milestone:** _(between releases)_
**Last closed:** **v1.7 "house lights" — 2026-06-15**, tag `v1.7`. A **demo-UI-hardening release**: a fresh browser at a
demo's offset UI renders the working app with **zero manual steps**. Triggered live — next-web at `http://localhost:33000`
(demo-3) showed a **blank page** (clerk-js's handshake to the fake FAPI hit an untrusted self-signed cert) and studio-desk
302'd to a dead `:9100`. **M31** automated a locally-trusted **mkcert** FAPI cert into the demo bring-up (one branch in
`up-injected.sh`; openssl fallback + `DEMO_NO_MKCERT` opt-out; fake BAPI plain HTTP → out of scope) so next-web renders;
**M32** fixed the studio-desk `:9100`-dead-redirect (a `NODE_ENV=production` override winning the additive env-merge →
the single-port production `sendFile` path) + the `:9100` doc/CORS sweep. Verified: M31 browser-trust by composition
(chromium trusts the mkcert cert, rejects the old self-signed); M32 route-coverage code-read (no gap) + a live merge-probe.
close-release **GREEN** (all 9 sweeps GREEN, 0 blocking). **Tooling + docs only — zero platform-repo edits.**
**Next up:** **`/developer-kit:design-roadmap`** to scope the next version. (Outward-facing: push the ext tags
`house-lights-m31`/`m32` + the still-unpushed `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` to `origin`. The orphaned
**`m26/self-contained-demo`** branch [tag `prop-room-m26`] + `wip/clerkenstein-browser-login` still await their own design-roadmap pass.)
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
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

_Last updated: 2026-06-15 (**v1.7 "house lights" SHIPPED** via `/developer-kit:close-release` — 9-sweep release review, ALL
GREEN [supply-chain 0-new-deps, scope M31+M32 100%-delivered/0-unaccounted, deferral GREEN, code-quality shellcheck+
py_compile clean, docs/KB coherent, tests Go 1027/Python 471/flake 0, metrics GREEN, decisions blended]; 0 blocking.
Merged → main, tagged `v1.7`; ext tag SHAs reconciled (house-lights-m31→5022e72, m32→7b17c39). Prior: 2026-06-15 M31 + M32
CLOSED; v1.7 designed; 2026-06-14 v1.6 "stage door" SHIPPED [tag `v1.6`].)_
