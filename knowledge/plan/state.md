# State

**Active version:** _(between releases — **v1.8 "understudy" SHIPPED 2026-06-15**, tag `v1.8`. The next version is
unplanned; run [`/developer-kit:design-roadmap`](roadmap.md) to scope it.)_
**Active milestone:** _(between releases)_
**Last closed:** **v1.8 "understudy" — 2026-06-15**, tag `v1.8`. The **self-contained-demo release**: `stack-demo/`
gets its **own platform clone set** so a box with only `stack-demo/` (no `stack-dev/`) runs a demo end-to-end —
closing the doc-vs-code gap where `up-injected.sh` built every image from `stack-dev` despite `CLAUDE.md` calling
`stack-demo` a "true peer". A single `section` milestone **M26** re-implemented the orphaned `m26/self-contained-demo`
branch (@ `25ab855`, unmergeable — predates v1.6/v1.7) onto `main`: a new `ensure-clones.sh` bootstraps `stack-demo`'s
peer clones from GitHub, the build SOURCE + the compose dir (`PLAT`, D-MAIN) moved to `stack-demo`, dev-image reuse
gated behind `--reuse-dev-images` (OFF by default) — preserving the stack-secrets module + M30/M31/M32. close-release
**GREEN** (6 parallel sweeps + adversarial critic, 0 blocking; 3 doc-coherence cleanups landed Fate-1). **Tooling +
docs only — zero platform-repo edits.**
**Next up:** **`/developer-kit:design-roadmap`** to scope the next version. **User-authorized follow-up:** the live
field-bake on a **freshly-emptied** `stack-demo/` (the on-disk one is populated from the orphan run + would mask a
from-scratch failure) — composition satisfied the close gate; the live run is the optional confirmation. (Outward-facing:
push the ext tags `understudy-m26` + `house-lights-m31`/`m32` + `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` to
`origin`; `wip/clerkenstein-browser-login` still awaits its own design-roadmap pass.)
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
**Paused:** _(none)_

## Recently shipped releases
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. The **self-contained-demo release**: a demo builds **entirely
  from `stack-demo`'s own clone set** — a box with only `stack-demo/` (no `stack-dev/`) runs a demo end-to-end. Single
  `section` milestone **M26** re-implemented the orphaned `m26/self-contained-demo` branch (@ `25ab855`, unmergeable —
  predates v1.6/v1.7) onto `main`: `ensure-clones.sh` bootstraps the peer clones from GitHub; the build SOURCE + the
  compose dir (`PLAT`, D-MAIN) moved to `stack-demo`; dev-image reuse gated behind `--reuse-dev-images` (OFF default).
  Preserves stack-secrets/M30 + M31 mkcert + M32 studio-desk. Go **1027** unchanged; Python 471→**501** (+30, the two
  touched suites); flake **0** (triple-clean 3/3); supply-chain GREEN (0 new deps). Code: `rosetta-extensions` @ tag
  `understudy-m26`. Records: [releases/archive/01.80-understudy/](releases/archive/01.80-understudy/) (review · retro ·
  metrics · lockfile).
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
## Headline numbers (v1.8 close — 2026-06-15; baseline v1.7 2026-06-15)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`), **unchanged across v1.8** (M26 touched no `.go` — the diff is
  shell + python + docs). Per-module: `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding
  **259** · stack-snapshot **333** · stack-secrets **160**. `go vet` + `gofmt` + `shellcheck` clean (all 5 touched
  shell scripts); flake **0**.
- **Python tests:** **501** (v1.7 471 → +30, the two M26-touched suites only): demo-stack/tests **110 → 138** (+28:
  `TestEnsureClones` + `TestSelfContainedSource` + `TestRenameDrift` retargets + `TestShellcheck` +1 at build;
  `TestEnsureClonesFunctional` +12 + `TestReuseFlagArrayExpansion` +3 at harden); stack-injection/tests **111 → 113**
  (+2: the `reuse_dev_images` opt-in tests). **Triple-clean 3/3** (py3.11; demo-stack 138/138 + stack-injection
  113/113 each, zero flakes) + the milestone 5/5 randomized. `gen_injected_override.py` **99%**. GUIDE advertised
  count **41** reconciles (`TestGuideDocTruth` green). The corpus README-index guard runs **exit 0**.
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
**v1.8 SHIPPED:** `release/01.80-understudy` merged `--no-ff` → `main`, tagged **`v1.8`** (2026-06-15); release branch
deleted; the single milestone branch (`m26/self-contained-demo`) merged + deleted. The change lives in the `demo-stack`
+ `stack-injection` ext sections (authored in `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag).
v1.8 ext marker: **`understudy-m26`** @ `773184f` (ext `main` ff'd to `773184f`, final). The orphaned
**`m26/self-contained-demo`** branch + tag **`prop-room-m26`** (@ `25ab855`) were **deleted** at close — superseded by
the re-implementation. (`wip/clerkenstein-browser-login` still awaits its own design-roadmap pass.)
**v1.7 SHIPPED:** `release/01.70-house-lights` → `main`, tag **`v1.7`** (2026-06-15); ext markers `house-lights-m31` @
`5022e72` · `house-lights-m32` @ `7b17c39`.
**v1.6 SHIPPED:** `release/01.60-stage-door` → `main`, tag **`v1.6`** (2026-06-14); ext markers `stage-door-m27`/`m28`/`m30`.
**Prior:** **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-15 (**v1.8 "understudy" SHIPPED** via `/developer-kit:close-release` — 6 parallel review sweeps
(supply-chain/scope/code-quality/docs/tests/decisions) + an adversarial completeness critic, ALL GREEN [0 new deps,
scope 100%-delivered, deferral GREEN, code-quality shellcheck/py_compile/gofmt clean, tests Go 1027/Python 501/triple-clean
3-0/flake 0, metrics GREEN, decisions blended]; 0 blocking; 3 doc-coherence cleanups landed Fate-1. Merged → main, tagged
`v1.8`; ext `understudy-m26` finalized @ `773184f`; orphan branch+tag `prop-room-m26`/`m26/self-contained-demo` deleted.
Prior: 2026-06-15 M26 closed + v1.8 designed; 2026-06-15 v1.7 SHIPPED [tag `v1.7`]; 2026-06-14 v1.6 SHIPPED.)_
