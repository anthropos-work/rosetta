# State

**Active version:** **v1.7 "house lights" — ALL MILESTONES CLOSED, READY FOR `/developer-kit:close-release`** (designed
2026-06-15; branch `release/01.70-house-lights`). A **demo-UI-hardening release**: a fresh browser at a demo's offset UI
renders the working app with **zero manual steps**. Triggered live — next-web at `http://localhost:33000` (demo-3) showed
a **blank page** (clerk-js's handshake to the fake FAPI hit an untrusted self-signed cert) and studio-desk 302'd to a dead
`:9100`. **M31** automated a locally-trusted **mkcert** FAPI cert into the demo bring-up; **M32** fixed the studio-desk
`:9100`-dead-redirect via a `NODE_ENV=production` override + the `:9100` doc/CORS sweep. **Both closed 2026-06-15.**
ant-academy demo liveness → backlog (repro-first). **Tooling + docs only — zero platform-repo edits.**
**Active milestone:** **(between milestones — v1.7's M31 + M32 both closed; the release awaits `/developer-kit:close-release`,
which the user invokes).**
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
**Next up:** **`/developer-kit:close-release` v1.7** (the user invokes it) — review the whole release end-to-end (M31 + M32
as one PR), merge `release/01.70-house-lights` `--no-ff` → `main`, tag `v1.7`. The orchestrator finalizes the ext side of
both milestones (ff `main` + re-point the `house-lights-m31`/`m32` tags + delete the branches).
(Outward-facing carry-over for close-release: push the v1.5/v1.6 ext tags `stage-door-m27`/`m28`/`m30` + `prop-room-m21..m25` +
`house-lights-m31`/`m32` to `origin`. The orphaned **`m26/self-contained-demo`** branch [tag `prop-room-m26`] +
`wip/clerkenstein-browser-login` still await their own design-roadmap pass.)
**Phase:** **v1.7 ALL MILESTONES CLOSED — M31 + M32 both merged → `release/01.70-house-lights` (2026-06-15); ready for `/developer-kit:close-release`.**
**Paused:** _(none)_

## Recently closed
- **M32 — studio-desk single-port / production + the `:9100` sweep** — **2026-06-15** (FINAL v1.7 milestone). Pinned
  `NODE_ENV=production` (+ `FRONTEND_PORT=9000`) on the studio-desk demo override so the additive env-merge wins back the
  production `sendFile` path (no dead-`:9100` 302); route coverage verified by code-read (NO GAP) + a close-time live
  merge-probe. Dropped the dead un-offset `:9100` CORS origin; swept `:9100`→`:9000` across `frontend-tier.md` + the
  demo-up SKILL. Regression guard mutation-checked 4 ways + a CORS exact-set assert. Python stack-injection 87→88;
  full suite 88/88 (0 skipped); Go 1027 unchanged; flake 0; deferral audit GREEN. Ext tag `house-lights-m32` @ `107599c`.
  Full record: roadmap.md `### M32`.
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

## Headline numbers (v1.7 all milestones closed — v1.6 close baseline + the M31 + M32 deltas)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`), **unchanged across v1.7** (neither M31 nor M32 touched Go). Per-module:
  `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding **259** · stack-snapshot **333** ·
  stack-secrets **160**. `go vet` + `gofmt` + `shellcheck` clean; flake **0**.
- **Python tests:** **+11 at M31** (`FapiCertStep` in `demo-stack/tests/test_tooling.py`: 47→**50**; demo-stack suite
  99→110; v1.6 headline 459 → **470**) **+1 at M32** (`test_studio_desk_env_pins_node_env_production` in
  `stack-injection/tests/test_injection.py`: 87→**88**; the harden CORS exact-set + the 2 latent env-masked fixes were
  on existing tests). M32 full `test_injection.py` suite **88/88** (0 skipped under PyYAML, authoritative JUnit tally).
  All pass; flake **0** (5/5 randomized sequential, both milestones). The corpus README-index guard runs **exit 0**.
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
**v1.7 ALL MILESTONES CLOSED — READY FOR close-release:** `release/01.70-house-lights` cut from `main` 2026-06-15; both
milestone branches `m31/mkcert-fapi-cert` + `m32/studio-desk-singleport` merged `--no-ff` → the release branch + deleted
(2026-06-15). The fixes landed in the `demo-stack` + `stack-injection` ext sections (authored in
`.agentspace/rosetta-extensions/`, tagged `house-lights-m31` @ `6565ef8` / `house-lights-m32` @ `107599c`, consumed
per-stack); docs in the rosetta corpus. close-release merges `release/01.70-house-lights` → `main` + tags `v1.7`; the
orchestrator finalizes both ext sides. ant-academy liveness (M33) → roadmap-vision backlog (repro-first).
**v1.6 SHIPPED:** `release/01.60-stage-door` merged `--no-ff` → `main`, tagged **`v1.6`** (2026-06-14); release branch
deleted; all 4 milestone branches (`m27/secret-coverage-dna` … `m30/field-bake`) merged + deleted. The new
`stack-secrets` extension lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in
`.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag. v1.6 ext markers: **`stage-door-m27`** ·
**`m28`** · **`m30`** (M29 added no ext tag — rosetta-only); ext head `868a68a` (final, includes the close-release ext
fixes). The orphaned **`m26/self-contained-demo`** branch (tag `prop-room-m26`) + `wip/clerkenstein-browser-login` are
preserved on the ext side, awaiting their own roadmap home.
**Prior:** **v1.5** tag `v1.5` (2026-06-14) · **v1.3b** `v1.3.1` (2026-06-09) · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-15 (**M32 CLOSED** via `/developer-kit:close-milestone` — the studio-desk single-port/production
fix + the `:9100` sweep merged `--no-ff` → `release/01.70-house-lights`, milestone branch deleted. This was v1.7's FINAL
milestone → **v1.7 is now ready for `/developer-kit:close-release`** [user invokes]. Review found 4 findings, all Fate-1:
3 decision-tag blends [D1/D2/D4 into `frontend-tier.md`] + 1 adversarial-scenario record [the additive-env-merge, defended
live via a merge-probe]. Scope GREEN [all 8 boxes Fate-1], code-quality GREEN, docs GREEN, tests GREEN [Python
stack-injection 87→88 / Go +0 / suite 88/88 / flake 0], deferral re-audit GREEN [0 in-milestone punts; inherited backlog
all cross-release/re-signed; 0 repeat/aged-out]. Ext tag `house-lights-m32` @ `107599c` [orchestrator finalizes the ext
side]. Prior: 2026-06-15 **M31 CLOSED** [mkcert FAPI cert; Python +11; ext tag `house-lights-m31`]. Prior: 2026-06-15
**v1.7 "house lights" DESIGNED + IN DEVELOPMENT**. Prior: 2026-06-14 **v1.6 "stage door" SHIPPED** [tag `v1.6`].)_
