# State

**Active version:** **v1.6 "stage door" — IN DEVELOPMENT** (designed 2026-06-14; branch `release/01.60-stage-door`).
The **secret-provisioning release**: one mechanism that ingests a secret source (directory/zip, default
`.agentspace/secrets`) and **provisions every repo of a stack** from it (values-blind), with a **secret-coverage DNA**
(a one-sided harness in the `datadna` mold) that *lists and keeps listed* the required secrets per repo. 4 milestones
M27→M30 (DNA+ingest → engine+gate → docs+skill → field-bake). **Tooling + docs only — zero platform-repo edits; never
commit `.env`; never write prod; no verb ever reads or echoes a secret value.**
**Active milestone:** _(between milestones — all 4 v1.6 milestones M27–M30 closed; v1.6 ready for `/developer-kit:close-release`)._
**Last closed:** **v1.5 "prop room" — 2026-06-14**, tag `v1.5`. The **local-Directus release**: every stack now
serves its **own captured public catalog** from a per-stack Directus (data plane local, asset plane prod → real
images) on `--local-content` (demo default-on, dev opt-in); prod-read is the documented fallback. M21 structure
capture → M22 executed provisioning + per-stack Directus lifecycle → M23 content cutover + referential closure →
M24 docs convergence + hygiene → M25 field-bake. The closing field-bake proved it live on a 16 GB box (curl-served
catalog) and pre-paid the field-fix tail: it caught + fixed **4 real release bugs** Fate-1, headline being the
`directus_files` **tenant-data leak the firewall caught FAIL-CLOSED** (fixed in the FILTER, firewall never weakened).
**Tooling + docs only — zero platform-repo edits.**
**Last milestone closed:** **M30 — 2026-06-14** (merged `m30/field-bake` → `release/01.60-stage-door`). The
**field-bake** — the FINAL v1.6 milestone, proven LIVE. Assembled a compliant gitignored `.agentspace/secrets`
from stack-dev (values-blind), `check` Critical **100%** on dev + demo (exit 0), then brought a fresh **demo-3**
LIVE from that source (provision 26/2/0 → **17 containers UP**, all probes pass — the observable-behavior gate MET).
Caught + fixed **2 real release bugs** Fate-1: (1) `sentinel/DB_CONNECTION` was critical/required but is
compose-injected config → reclassified `waived-config` + regression; (2) the demo bring-up only *checked* coverage,
never *provisioned*, + `preflight.sh` resolved its source path one level too shallow (demo gate silently skipped) →
added the provision step + fixed the path. **SAFETY live:** prod `DIRECTUS_TOKEN` (len-32) armed in **ZERO** containers
(blank + override-stripped, defense-in-depth). Close **GREEN** (4 findings, all Fate-1: 3 docs reconciling the record
+ `secrets-spec.md` to the executed live bake — version m27→m30, sentinel `waived-config`, split 39/8/8+12-crit; 1
decision-triage); deferral audit GREEN. Go **1027** / Python **459** unchanged (the M30 regression is a sub-assertion);
flake **0**. Ext @ tag **`stage-door-m30`** (head `29c922b`, 2 field-fix commits; orchestrator finalizes the ext side).
**Next up:** **`/developer-kit:close-release` for v1.6 "stage door"** — all 4 milestones M27–M30 are closed; the release
is ready for the release-level review + merge to `main` + tag.
Research + risk register: [`.agentspace/scratch/roadmap-research-2026-06-14.md`](../../.agentspace/scratch/roadmap-research-2026-06-14.md).
(Outward-facing carry-over from v1.5: push the 5 ext tags `prop-room-m21..m25` to `origin`; + the v1.6 tags `stage-door-m27`/`m28`/`m30`.)
**Phase:** **v1.6 in development — all 4 milestones M27/M28/M29/M30 CLOSED (merged to `release/01.60-stage-door`); ready for `/developer-kit:close-release`.**
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

## Headline numbers (v1.6 in development — updated at M30 close 2026-06-14; counts unchanged since M28)
- **Go test funcs:** **1027** total (`Test`+`Fuzz`, measured at ext `m28` head). Per-module:
  `rosetta-extensions/alignment` **52** · clerkenstein **223** · stack-seeding **259** · stack-snapshot **333** ·
  **stack-secrets 160** (113 at M27 → +47 at M28: the provision engine + demo overlay + 3-pass harden + review-fix).
  v1.5 baseline 867 → **+160** across M27+M28. `go vet` + `gofmt` + `shellcheck` clean; flake **0**; 5/5 `-race -shuffle`
  (Go) + 5/5 sequential (Python) clean.
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
**v1.6 IN DEVELOPMENT (all milestones closed — ready for `/close-release`):** `release/01.60-stage-door` cut from
`main` 2026-06-14; milestone branches `m27/secret-coverage-dna` ✅ → `m28/provisioning-engine` ✅ → `m29/secrets-docs-skill` ✅ → `m30/field-bake` ✅ (all merged + deleted). The
`stack-secrets` extension is authored in `.agentspace/rosetta-extensions/stack-secrets/`, tagged `stage-door-mNN`,
consumed per-stack at the pinned tag (the standard two-clone policy). v1.6 ext markers: **`stage-door-m27`** ·
**`m28`** · **`m30`** (M29 added no ext tag — rosetta-only); M30 ext head `29c922b` @ tag `stage-door-m30` (the
orchestrator finalizes the ext side: ff `main` + re-point the tag + delete `m30/field-bake`).
**v1.5 SHIPPED:** `release/01.50-prop-room` merged `--no-ff` → `main`, tagged **`v1.5`** (2026-06-14); release branch
deleted; all 5 milestone branches (`m21/structure-capture` … `m25/field-bake`) merged + deleted. The stack tooling
lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in `.agentspace/rosetta-
extensions/`, consumed per-stack at a pinned tag. v1.5 ext markers: **`prop-room-m21`** · **`m22`** · **`m23`** ·
**`m24`** · **`m25`** @ `fbb8783` (final, includes the close-release ext fixes); the M24 `/project-stats` fix is the
cross-repo developer-kit `825cdce` in the `ant-singularity` node repo (outside the ext tags). Snapshot payloads live
in a gitignored `.agentspace/snapshots/` cache (cloud/S3 store = backlog/unscheduled, DEF-M10-01).
**Prior:** **v1.3b** tag `v1.3.1` (2026-06-09) · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0` (2026-06-03).

_Last updated: 2026-06-14 (**M30 CLOSED** via `/developer-kit:close-milestone` — the field-bake, the FINAL v1.6
milestone, proven LIVE [demo-3, 17 containers, gate met]; merged `m30/field-bake` → `release/01.60-stage-door`. Close
GREEN: 4 findings all Fate-1 [3 docs reconciling the record + secrets-spec.md to the live bake, 1 decision-triage], 2
field bugs fixed Fate-1, deferral audit GREEN. Go 1027 / Python 459 unchanged. **Next: `/developer-kit:close-release`
for v1.6 — all 4 milestones closed.** Prior: 2026-06-14 M29 CLOSED [docs + /stack-secrets skill, rosetta-only]; M28
CLOSED [the provision engine + demo-aware gate]; M27 CLOSED [the secret-coverage DNA + ingestion]; v1.5 "prop room"
SHIPPED [tag `v1.5`].)_
