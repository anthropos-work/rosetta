# State

**Active version:** _(between releases — **v1.6 "stage door" SHIPPED 2026-06-14**, tag `v1.6`. The next version is
unplanned; run [`/developer-kit:design-roadmap`](roadmap.md) to scope it.)_
**Active milestone:** _(between releases)_
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
**Next up:** **`/developer-kit:design-roadmap`** to scope the next version. (Outward-facing: push the v1.6 ext tags
`stage-door-m27`/`m28`/`m30` + the still-unpushed `prop-room-m21..m25` to `origin`. The orphaned **`m26/self-contained-demo`**
branch [tag `prop-room-m26`] awaits its own `/developer-kit:design-roadmap` pass for a version + scope — see context.md.)
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
**Paused:** _(none)_

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

## Headline numbers (v1.6 close — 2026-06-14; baseline v1.5 2026-06-14)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` **52** · clerkenstein
  **223** · stack-seeding **259** · stack-snapshot **333** · **stack-secrets 160** (the net-new v1.6 section: 113 @ M27
  → 160 @ M28; M29 docs-only +0; M30 +0-func). v1.5 baseline 867 → **+160** (all in `stack-secrets`). `go vet` +
  `gofmt` + `shellcheck` clean; flake **0**; triple-clean (Go `-race -shuffle` + Python sequential) clean.
- **Python tests:** **459** collected (451 active + 8 env-gated skip), unchanged from v1.5 (the v1.6 growth surface is
  Go, not Python). The corpus README-index guard runs **exit 0** (every doc indexed, incl. the net-new `secrets-spec.md`).
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
**v1.6 SHIPPED:** `release/01.60-stage-door` merged `--no-ff` → `main`, tagged **`v1.6`** (2026-06-14); release branch
deleted; all 4 milestone branches (`m27/secret-coverage-dna` … `m30/field-bake`) merged + deleted. The new
`stack-secrets` extension lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in
`.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag. v1.6 ext markers: **`stage-door-m27`** ·
**`m28`** · **`m30`** (M29 added no ext tag — rosetta-only); ext head `868a68a` (final, includes the close-release ext
fixes). The orphaned **`m26/self-contained-demo`** branch (tag `prop-room-m26`) + `wip/clerkenstein-browser-login` are
preserved on the ext side, awaiting their own roadmap home.
**Prior:** **v1.5** tag `v1.5` (2026-06-14) · **v1.3b** `v1.3.1` (2026-06-09) · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.

_Last updated: 2026-06-14 (**v1.6 "stage door" SHIPPED** via `/developer-kit:close-release` — 9-sweep release review:
supply-chain GREEN [stdlib-only], scope GREEN [M27–M30 all Fate-1, 0 unaccounted], deferral re-audit GREEN, code-quality
GREEN, docs/KB YELLOW→resolved [should-fix polish landed Fate-1], tests GREEN [Go +160 / Python +0 / flake 0], metrics
GREEN. Merged → main, tagged `v1.6`. Prior: 2026-06-14 M30 CLOSED [the field-bake, proven LIVE on demo-3]; M27–M29 closed;
v1.5 "prop room" SHIPPED [tag `v1.5`].)_
