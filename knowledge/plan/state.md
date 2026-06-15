# State

**Active version:** **v1.7 "house lights" — IN DEVELOPMENT** (designed 2026-06-15; branch `release/01.70-house-lights`).
A **demo-UI-hardening release**: a fresh browser at a demo's offset UI renders the working app with **zero manual steps**.
Triggered live — next-web at `http://localhost:33000` (demo-3) showed a **blank page** because clerk-js's handshake to the
fake FAPI (`https://127.0.0.1:35400`) hit an **untrusted self-signed cert**. **M31** automates a locally-trusted **mkcert**
FAPI cert into the demo bring-up (openssl fallback + `DEMO_NO_MKCERT` opt-out; fake BAPI is plain HTTP → out of scope);
**M32** fixes the sibling studio-desk `:9100`-dead-redirect (a 1-line `NODE_ENV=production` override fix + the `:9100`
doc/CORS sweep). ant-academy demo liveness → backlog (repro-first). **Tooling + docs only — zero platform-repo edits.**
**Active milestone:** **M32 — studio-desk single-port / production alignment + the `:9100` sweep** (`planned`, not
started). A fresh browser at demo-N's studio-desk lands on a live page instead of a 302 to the dead `:9100`, by
running the container's production code path (a 1-line `NODE_ENV=production` override fix + a regression assertion +
a Playwright smoke + the `:9100` doc/CORS sweep). Build with `/developer-kit:build-milestone`. Sequence: last of v1.7.
**Last closed:** **v1.6 "stage door" — 2026-06-14**, tag `v1.6`. The **secret-provisioning release**: one mechanism that
ingests a secret source (directory/zip, default `.agentspace/secrets`) and **provisions every repo of a stack** from it
(values-blind), with a **secret-coverage DNA** (gene = repo×KEY, **6 repos / 55 genes**; `introspect`+`diff` keep-listed
gate) that *lists and keeps listed* the required secrets per repo, a **demo-aware `check`** wired non-fatally into
`/dev-up` + `/demo-up` pre-flight, and the **`/stack-secrets`** skill. M27 secret-coverage DNA + dir/zip ingestion →
M28 provision engine (N=0 guard + the **`DIRECTUS_TOKEN`-non-rearm** blocks-release safety) → M29 docs + skill + corpus
wiring → M30 field-bake. The closing **field-bake proved it LIVE** on a fresh **demo-3** (17 containers UP, `check`
Critical **100%**, prod `DIRECTUS_TOKEN` armed in **ZERO** containers) built from a stack-dev-assembled
`.agentspace/secrets`, and caught + fixed **2 real bugs** Fate-1 (the silently-skipping demo pre-flight gate + the
never-provisioned demo). **Tooling + docs only — zero platform-repo edits; values-blind; never write prod.**
**Next up:** **build M32** via `/developer-kit:build-milestone` (the studio-desk `NODE_ENV=production` override fix +
the `:9100` sweep), then **close M32** + **close-release v1.7**.
Research + fix design + risk register: [`.agentspace/scratch/roadmap-research-2026-06-15.md`](../../.agentspace/scratch/roadmap-research-2026-06-15.md).
(Outward-facing carry-over: push the v1.6 ext tags `stage-door-m27`/`m28`/`m30` + the still-unpushed `prop-room-m21..m25` to
`origin`. The orphaned **`m26/self-contained-demo`** branch [tag `prop-room-m26`] still awaits its own design-roadmap pass.)
**Phase:** **v1.7 in development — M31 CLOSED (merged → `release/01.70-house-lights` 2026-06-15); M32 next.**
**Paused:** _(none)_

## Recently closed
- **M31 — mkcert-trusted FAPI cert** — **2026-06-15**. Automated a locally-trusted mkcert FAPI cert into the demo
  bring-up (one branch in `up-injected.sh` 3a-bis; openssl fallback factored + byte-compatible; `DEMO_NO_MKCERT`
  opt-out; non-fatal; ZERO change to the 3 path-only cert-consumers). Verified by composition (chromium trusts the
  mkcert cert vs rejects the openssl self-signed). Python +11 (`FapiCertStep`); Go 1027 unchanged; flake 0; deferral
  audit GREEN. Ext tag `house-lights-m31` @ `6565ef8`. Full record: roadmap.md `### M31`.

## Recently shipped releases
- **v1.6 "stage door"** — **2026-06-14**, tag `v1.6`. The **secret-provisioning release**: ingest a secret source
  (dir/zip, default `.agentspace/secrets`) → **provision every repo's `.env`** from it (values-blind) → a **secret-coverage
  DNA** (6 repos / 55 genes; `introspect`+`diff` keep-listed gate) → demo-aware `check` wired non-fatally into `/dev-up`
  + `/demo-up` pre-flight → the `/stack-secrets` skill. M27 DNA + ingestion → M28 provision engine (N=0 guard +
  `DIRECTUS_TOKEN`-non-rearm) → M29 docs + skill → M30 field-bake (proven LIVE on demo-3: 17 containers, Critical 100%,
  prod token armed in ZERO containers; 2 field bugs fixed). Go 867→**1027** (+160, the net-new stdlib-only `stack-secrets`
  Go section); Python **459** unchanged; flake **0**; triple-clean 3/3; supply-chain GREEN (0 third-party deps).
  close-release YELLOW→resolved (0 blocking; should-fix doc/consistency polish landed Fate-1; deferral re-audit GREEN).
  Code: `rosetta-extensions` @ tags `stage-door-m27`/`m28`/`m30`. Records:
  [releases/archive/01.60-stage-door/](releases/archive/01.60-stage-door/) (review · retro · metrics · lockfile).
- **v1.5 "prop room"** — **2026-06-14**, tag `v1.5`. The **local-Directus release**: stacks serve their own captured
  public catalog locally (data plane local, asset plane prod) on `--local-content`. M21 structure capture → M22 executed
  provisioning + per-stack Directus lifecycle → M23 content cutover + referential closure → M24 docs + hygiene → M25
  field-bake. Go 736→**867** (+131); Python 360→**459** (+99); flake **0**; supply-chain GREEN. Records:
  [releases/archive/01.50-prop-room/](releases/archive/01.50-prop-room/).
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up`
  field issues: M16→M20. Go 713→736; Python 174→360; flake 0; supply-chain GREEN. Records:
  [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/).

## Headline numbers (v1.7 in development — v1.6 close baseline + the M31 delta)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`), **unchanged at M31** (M31 touched no Go). Per-module:
  `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding **259** · stack-snapshot **333** ·
  stack-secrets **160**. `go vet` + `gofmt` + `shellcheck` clean; flake **0**.
- **Python tests:** **+11 at M31** (the `FapiCertStep` class in `demo-stack/tests/test_tooling.py`: 47→**50**; the
  full demo-stack suite 99→110; the v1.6 headline 459 → **470**). All pass; flake **0** (5/5 randomized sequential).
  The corpus README-index guard runs **exit 0**. (`demo-stack/README.md`'s test count reconciled 13→50 at M31 close.)
- **The v1.6 thesis:** a stack's per-repo `.env` is **provisioned from one secret source** (dir/zip, default
  `.agentspace/secrets`), **values-blind**, verified by a **6-repo / 55-gene secret-coverage DNA** + a two-tier
  keep-listed gate — **proven live** (M30: a fresh demo-3 came up entirely from a stack-dev-assembled secret dir,
  `check` Critical 100%). The manual `.env` hand-copy + the `setup_guide.md:447` TODO are retired.
- **Safety:** **values-blind** end-to-end (no verb reads/echoes/logs a secret value — the only value-carrying path is
  `provision` writing to the gitignored target); the prod **`DIRECTUS_TOKEN` is never re-armed** on a non-prod stack
  (the fix16/17 class — verified live: armed in **ZERO** demo-3 containers). `.env` / secrets never committed (gitignored).
- **Supply-chain:** **GREEN** — the new `stack-secrets` Go module is **stdlib-only** (no `require` block, no `go.sum`,
  zero third-party deps → minimal values-blind audit surface); the 4 prior ext modules untouched; all-permissive licenses.
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces — v1.6 touched no clerkenstein behavior.

## Branch model
**v1.7 IN DEVELOPMENT:** `release/01.70-house-lights` cut from `main` 2026-06-15; milestone branches
`m31/mkcert-fapi-cert` → `m32/studio-desk-singleport` to follow. The fixes land in the `demo-stack` + `stack-injection`
ext sections (authored in `.agentspace/rosetta-extensions/`, tagged `house-lights-m31`/`m32`, consumed per-stack);
docs in the rosetta corpus. ant-academy liveness (M33) → roadmap-vision backlog (repro-first).
**v1.6 SHIPPED:** `release/01.60-stage-door` merged `--no-ff` → `main`, tagged **`v1.6`** (2026-06-14); release branch
deleted; all 4 milestone branches (`m27/secret-coverage-dna` … `m30/field-bake`) merged + deleted. The new
`stack-secrets` extension lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in
`.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag. v1.6 ext markers: **`stage-door-m27`** ·
**`m28`** · **`m30`** (M29 added no ext tag — rosetta-only); ext head `868a68a` (final, includes the close-release ext
fixes). The orphaned **`m26/self-contained-demo`** branch (tag `prop-room-m26`) + `wip/clerkenstein-browser-login` are
preserved on the ext side, awaiting their own roadmap home.
**Prior:** **v1.5** tag `v1.5` (2026-06-14) · **v1.3b** `v1.3.1` (2026-06-09) · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-15 (**M31 CLOSED** via `/developer-kit:close-milestone` — the mkcert-trusted FAPI cert merged
`--no-ff` → `release/01.70-house-lights`, milestone branch deleted. Review found 2 findings, both Fate-1: the
`demo-stack/README.md` test-count drift [13→50] + a recorded adversarial scenario [zero-byte cert / existence-only
guard = pre-existing, documented]. Scope GREEN [all 11 boxes Fate-1], code-quality GREEN, docs GREEN, tests GREEN
[Python +11 / Go +0 / flake 0], deferral re-audit GREEN [v1.7's first milestone — 0 inherited/repeat]. Ext tag
`house-lights-m31` @ `6565ef8` [orchestrator finalizes the ext side]. M32 studio-desk single-port/production is next.
Prior: 2026-06-15 **v1.7 "house lights" DESIGNED + IN DEVELOPMENT** [M31 → M32; branch `release/01.70-house-lights`
cut]. Prior: 2026-06-14 **v1.6 "stage door" SHIPPED** [tag `v1.6`]; M30 field-bake CLOSED [proven LIVE on demo-3];
v1.5 "prop room" SHIPPED [tag `v1.5`].)_
