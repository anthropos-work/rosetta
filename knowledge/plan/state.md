# State

**Active version:** **v1.8 "understudy" — IN DEVELOPMENT** (designed 2026-06-15; branch `release/01.80-understudy`).
The **self-contained-demo release**: give `stack-demo/` its **own platform clone set** so a box with only `stack-demo/`
(no `stack-dev/`) can run a demo end-to-end. Closes a live doc-vs-code gap — `CLAUDE.md` already calls `stack-demo` *"a
true peer with its own clone set,"* and M30 already provisions `stack-demo/platform/.env`, but `up-injected.sh` still
builds every image from `stack-dev` (`src="$DEV/$svc"`, `PLAT="$DEV/platform"`). **Re-implements** the orphaned
`m26/self-contained-demo` branch (@ `25ab855`, the spec — predates v1.6/v1.7 so unmergeable) onto current `main`,
preserving the stack-secrets module + M30 provision + M31 mkcert + M32 studio-desk fix. **Tooling + docs only — zero
platform-repo edits** (`stack-demo/platform` is a build *context* only).
**Active milestone:** **(between milestones — M26 closed; v1.8 awaiting `/developer-kit:close-release`).** M26 —
Self-contained demo stacks — **closed 2026-06-15**, merged `--no-ff` → `release/01.80-understudy`. The single
`section` milestone of v1.8: a demo now builds **entirely from `stack-demo`'s OWN clone set** (`ensure-clones.sh`
bootstraps the peer clones + the `$DEV`→`stack-demo` build-source repoints + the D-MAIN PLAT move so the compose
contexts resolve against `stack-demo` too + the `reuse_dev_images` opt-in gate), preserving M30/M31/M32. Close
GREEN: 2 Fate-1 doc/comment findings fixed; demo-stack 138/138 + stack-injection 113/113; flake 0; field-bake
satisfied by composition (the live freshly-emptied-`stack-demo/` run is a user-authorized post-close follow-up).
**Last closed:** **M26 — Self-contained demo stacks — 2026-06-15** (merged `--no-ff` → `release/01.80-understudy`;
the only milestone of v1.8). Last shipped *release*: **v1.7 "house lights"** (tag `v1.7`, 2026-06-15) — see Recently
shipped below.
**Next up:** **`/developer-kit:close-release` for v1.8 "understudy"** (the release-level review + ff
`release/01.80-understudy` → `main` + tag `v1.8`) — the user invokes it separately. **Orchestrator post-close
(ext side):** ff ext `main` → the reimpl HEAD, re-point the `understudy-m26` tag forward over the 2 harden + 1 close
commit, delete the orphan tag `prop-room-m26` + orphan branch `m26/self-contained-demo` (superseded by the reimpl).
**User-authorized follow-up:** the FULL LIVE field-bake on a **freshly-emptied** `stack-demo/` (the on-disk one is
populated from the orphan run + would mask a from-scratch failure) — composition satisfied the close gate; the live
run is the optional confirmation. (Outward-facing carry-over unchanged: push the ext tags `understudy-m26` +
`house-lights-m31`/`m32` + `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` to `origin`;
`wip/clerkenstein-browser-login` still awaits its own design-roadmap pass.)
**Phase:** **v1.8 — M26 `done` (closed 2026-06-15, merged to `release/01.80-understudy`); release awaiting `/developer-kit:close-release`.**
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

## Headline numbers (M26 close — 2026-06-15; baseline v1.7 2026-06-15)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`), **unchanged at M26** (the milestone touched no `.go` — the diff
  is shell + python + docs). Per-module: `rosetta-extensions/alignment` **52** · clerkenstein **223** ·
  stack-seeding **259** · stack-snapshot **333** · stack-secrets **160**. `go vet` + `gofmt` + `shellcheck` clean
  (all 5 touched shell scripts); flake **0**.
- **Python tests (the two M26-touched suites):** demo-stack/tests **110 → 138** (+28: `TestEnsureClones` +
  `TestSelfContainedSource` + `TestRenameDrift` retargets + `TestShellcheck` +1 at build; `TestEnsureClonesFunctional`
  +12 + `TestReuseFlagArrayExpansion` +3 at harden); stack-injection/tests **111 → 113** (+2: the `reuse_dev_images`
  opt-in tests). py3.11/PyYAML JUnit-authoritative (0 skipped); **flake 0** (5/5 randomized, both suites);
  `gen_injected_override.py` **99%**. GUIDE advertised count **41** reconciles (`TestGuideDocTruth` green). The corpus
  README-index guard runs **exit 0**.
- **The M26 thesis:** a demo builds **entirely from `stack-demo`'s OWN clone set** — a box with only `stack-demo/`
  (no `stack-dev/`) can run a demo end-to-end. `ensure-clones.sh` bootstraps the peer clones; the D-MAIN PLAT move
  makes the compose contexts resolve against `stack-demo` too; dev-image reuse is OFF by default. Tooling + docs only.
- **Safety:** the demo's **one sanctioned cross-stack read** is ensure-clones' `.env` *seed* (copy-if-present,
  non-fatal, never-clobber) — never the build SOURCE (`safety.md` §2.7, #M26-D4). The `DIRECTUS_TOKEN`-non-rearm +
  secret/cert values-blind invariants (v1.6/v1.7) carry forward unchanged.
- **Supply-chain:** **GREEN** — M26 added **zero** third-party deps (shell + python-stdlib + docs; no manifest
  change); the stdlib-only posture carries forward.
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces — M26 touched no clerkenstein behavior.

## Branch model
**v1.7 SHIPPED:** `release/01.70-house-lights` merged `--no-ff` → `main`, tagged **`v1.7`** (2026-06-15); release branch
deleted; both milestone branches (`m31/mkcert-fapi-cert` + `m32/studio-desk-singleport`) merged + deleted. The fixes live
in the `demo-stack` + `stack-injection` ext sections (authored in `.agentspace/rosetta-extensions/`, consumed per-stack at
a pinned tag). v1.7 ext markers: **`house-lights-m31`** @ `5022e72` · **`house-lights-m32`** @ `7b17c39` (ext head
`7b17c39`, final). The orphaned **`m26/self-contained-demo`** branch (tag `prop-room-m26`) + `wip/clerkenstein-browser-login`
stay preserved on the ext side, awaiting their own design-roadmap pass.
**v1.6 SHIPPED:** `release/01.60-stage-door` → `main`, tag **`v1.6`** (2026-06-14); ext markers `stage-door-m27`/`m28`/`m30`.
**Prior:** **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-15 (**M26 CLOSED** via `/developer-kit:close-milestone` — merged `--no-ff` →
`release/01.80-understudy`, milestone branch deleted. Close GREEN: 2 Fate-1 doc/comment findings fixed (stale
"legacy stack-dev/.env base" comments re-worded [ext `773184f`]; the M26 sanctioned-read invariant blended into
`safety.md` §2.7); demo-stack 138/138 + stack-injection 113/113; flake 0; Go 1027 unchanged; deferral re-audit
GREEN; field-bake satisfied by composition (live run = user-authorized follow-up). v1.8 now awaits
`/developer-kit:close-release`. Prior: 2026-06-15 **v1.8 "understudy" DESIGNED** [single `section` M26, orphan
re-impl onto `main`, D-MAIN + D1–D6 settled]; 2026-06-15 **v1.7 "house lights" SHIPPED** [tag `v1.7`];
2026-06-14 v1.6 "stage door" SHIPPED.)_
