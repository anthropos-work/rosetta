# State

**Active version:** _(between releases — **v1.5 "prop room" SHIPPED 2026-06-14**, tag `v1.5`. The next version is
unplanned; run [`/developer-kit:design-roadmap`](roadmap.md) to scope it.)_
**Active milestone:** _(between releases)_
**Last closed:** **v1.5 "prop room" — 2026-06-14**, tag `v1.5`. The **local-Directus release**: every stack now
serves its **own captured public catalog** from a per-stack Directus (data plane local, asset plane prod → real
images) on `--local-content` (demo default-on, dev opt-in); prod-read is the documented fallback. M21 structure
capture → M22 executed provisioning + per-stack Directus lifecycle → M23 content cutover + referential closure →
M24 docs convergence + hygiene → M25 field-bake. The closing field-bake proved it live on a 16 GB box (curl-served
catalog) and pre-paid the field-fix tail: it caught + fixed **4 real release bugs** Fate-1, headline being the
`directus_files` **tenant-data leak the firewall caught FAIL-CLOSED** (fixed in the FILTER, firewall never weakened).
**Tooling + docs only — zero platform-repo edits.**
**Next up:** **`/developer-kit:design-roadmap`** to scope the next version. (Outward-facing: push the 5 ext tags
`prop-room-m21..m25` to `origin`; the M24 `/project-stats` fix is a separate commit in the `ant-singularity` node repo.)
**Phase:** **between releases — awaiting `/developer-kit:design-roadmap`.**
**Paused:** _(none)_

## Recently shipped releases
- **v1.5 "prop room"** — **2026-06-14**, tag `v1.5`. The **local-Directus release**: stacks serve their own captured
  public catalog locally (data plane local, asset plane prod) on `--local-content`. M21 structure capture (capture
  the content-model DDL+PKs+serve-rows atomically with rows; firewall structural-metadata admissibility class;
  redefined `stacksnap` exit codes) → M22 executed provisioning (per-stack Directus as a compose service, offset
  port, idempotent, verified; the `EnvContract` firewall a load-bearing executed gate) → M23 content cutover
  (`DIRECTUS_BASE_ADDR` re-point demo+dev; studio-desk minted token; `directus_files` ref capture; cross-surface
  closure gene) → M24 docs convergence + 4 hygiene items (go1.25.11 pin, README-index guard, zero-critical-genes
  guard, `/project-stats` scope fix) → M25 field-bake (5/5 live done-bars GREEN on the real box; 4 field bugs fixed).
  Go 736→**867** (+131); Python 360→**459** (+99); flake **0**; triple-clean 3/3; supply-chain GREEN (0 called CVEs,
  go1.25.11 clears the 12 stdlib advisories; all-permissive licenses). close-release GREEN (1 must-fix doc + 2
  should-fix, all Fate-1 land-now; deferral re-audit GREEN, DEF-M21-01 + M25-D9 → roadmap-vision backlog).
  Code: `rosetta-extensions` @ tags `prop-room-m21..m25` (ext head `fbb8783`). Records:
  [releases/archive/01.50-prop-room/](releases/archive/01.50-prop-room/) (review · retro · metrics · lockfile).
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up`
  field issues: M16 land-fixes → M17 re-run safety → M18 the verification net → M19 the frontend tier → M20 lifecycle
  convergence. Go 713→736; Python 174→360 collected; flake 0; triple-clean 3/3; supply-chain GREEN. Records:
  [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/).
- **v1.3 "stack party"** — **2026-06-07**, tag `v1.3`. The **dev/demo convergence**: unified stack registry [M12],
  dev-as-peer [M13], one generic `stack-*` skill set [M14], code-cited `safety.md` [M15]. Records:
  [releases/archive/01.30-stack-party/](releases/archive/01.30-stack-party/).

## Headline numbers (v1.5 close — 2026-06-14; baseline v1.3b 2026-06-09)
- **Go test funcs:** **867** total (`Test`+`Fuzz`, measured at final ext HEAD `fbb8783`). Per-module:
  `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding **259** · stack-snapshot **333**.
  Baseline v1.3b 736 → **+131**. `go vet` + `gofmt` clean; flake **0**.
- **Python tests:** **459** collected (451 active + 8 env-gated skip). Per-suite: stack-core **85** · dev-stack **73**
  · demo-stack **87** · stack-injection **110** · stack-verify **104**. Baseline v1.3b 360 → **+99**. py_compile CLEAN;
  the corpus README-index guard runs **exit 0** (every doc indexed).
- **The v1.5 thesis:** a freshly-bootstrapped, stacksnap-provisioned per-stack Directus **serves the captured public
  catalog locally** (data plane local, asset plane prod) with **zero hand SQL** — proven live on a 16 GB box (M25
  curl-served, DB-1/DB-2 GREEN). Demo default-on / dev opt-in (`--local-content`); prod-read the documented fallback.
- **Safety:** the tenant-data firewall held **fail-closed** under live prod capture (M25 caught a real 158-tenant-file
  over-capture, refused to persist, zero leak); fixed in the FILTER, never weakened. 100% data-DNA catalog (unchanged).
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces. v1.5 touched no clerkenstein behavior.
- **Test health:** flake **0**; triple-clean **3/3** (stack-snapshot `-race -shuffle` + the full suites); supply-chain
  GREEN (0 called CVEs; go1.25.11 pinned out the 12 stdlib advisories on all 4 go.mod + clerkenstein CI; all-permissive).

## Branch model
**v1.5 SHIPPED:** `release/01.50-prop-room` merged `--no-ff` → `main`, tagged **`v1.5`** (2026-06-14); release branch
deleted; all 5 milestone branches (`m21/structure-capture` … `m25/field-bake`) merged + deleted. The stack tooling
lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in `.agentspace/rosetta-
extensions/`, consumed per-stack at a pinned tag. v1.5 ext markers: **`prop-room-m21`** · **`m22`** · **`m23`** ·
**`m24`** · **`m25`** @ `fbb8783` (final, includes the close-release ext fixes); the M24 `/project-stats` fix is the
cross-repo developer-kit `825cdce` in the `ant-singularity` node repo (outside the ext tags). Snapshot payloads live
in a gitignored `.agentspace/snapshots/` cache (cloud/S3 store = backlog/unscheduled, DEF-M10-01).
**Prior:** **v1.3b** tag `v1.3.1` (2026-06-09) · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0` (2026-06-03).

_Last updated: 2026-06-14 (**v1.5 "prop room" SHIPPED** via `/developer-kit:close-release` — 5-dimension review:
supply-chain GREEN, scope/fate YELLOW→resolved, code-quality + docs fixes landed Fate-1 [the M24 sweep had missed
`snapshot-cold-start.md`; a stale "exit-4 / not-yet-automated" claim was the 1 must-fix], tests GREEN [Go +131 /
Python +99 / flake 0], deferral re-audit GREEN [DEF-M21-01 + M25-D9 → roadmap-vision backlog]. Merged → main, tagged
`v1.5`. Prior: 2026-06-13 M25 CLOSED (field-bake — v1.5 feature-complete).)_
