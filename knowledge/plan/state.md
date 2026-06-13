# State

**Active version:** **v1.5 "prop room"** (designed 2026-06-11; branch `release/01.50-prop-room`). The
**local-Directus release** — every stack today reads public content **live from prod**; v1.5 stands up a **local
Directus per stack** serving the **captured public library**, so content is self-contained. **Real images are
preserved** by keeping the asset plane on prod public links (data plane local, asset plane prod). Coverage: **every
demo** stack (default), **any dev-N≥1** (opt-in `--local-content`), **N=0** manual (documented recipe + n=0 guard).
5 milestones M21→M25 (structure capture → executed provisioning + lifecycle → content cutover + referential closure
→ docs + hygiene → field bake). **Tooling + docs only — zero platform-repo edits; capture stays read-only / public-
only / prod-untouched.** Full plan: [`roadmap.md`](roadmap.md) § In Development.
**Active milestone:** _(between milestones — M23 closed; M24 next — see Next up)_
**Last closed:** **M23 — Content cutover + referential closure — 2026-06-13** (`section`; all 6 sections Fate-1).
Cut the **data plane** over to the per-stack Directus (re-point `cms`'s `DIRECTUS_BASE_ADDR` → in-network
`http://directus:8055`; asset plane stays on prod so images stay real) + wired studio-desk's local instance + a
locally-minted static admin token + the `directus_files` ref capture (a new REFERENCED-SUBSET firewall kind +
`ClearByDelete` DELETE-before-TRUNCATE). Guarantees **referential closure** via a measured cross-surface gene
(full-taxonomy capture makes closure maximal by construction; the gene surfaced the **1 genuine prod residual**
`K-AIFUNX-E658` — a public sim referencing a customer-only skill, uncloseable by tooling, now MEASURED + named, an
operator-owned data fix). **2 inherited M21 deferrals RESOLVED** (directus_files + referential closure). ext tag
`prop-room-m23`.
**Next up:** build **M24** (`/developer-kit:build-milestone` — docs convergence + the 4-item hygiene strand), then M25 strictly sequential.
**Phase:** **v1.5 building — M23 CLOSED (merged to `release/01.50-prop-room`); next build M24, then M25.**
**Paused:** _(none)_

## Recently shipped releases
_(v1.5 "prop room" is the active release-in-development — see above; the list below is the shipped baseline it builds on.)_
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up` field issues: M16 land-fixes + doc-truth (published the devpath + migrate-race fixes, finished `anthropos-dev → stack-dev`) → M17 re-run safety (replay/seed idempotency + the `set -e` first-run-race) → M18 the verification net (`stack-verify` offset/scope-aware + auto-wired NON-FATAL; "UP" = verified) → M19 the frontend tier (next-web + studio-desk + ant-academy, per-demo cached builds, `--no-ui`, **zero platform-repo edits**) → M20 lifecycle convergence (`/demo-up` auto-set-dresses like `/dev-up` via one shared engine + the cold-start runbook). 4 net-new corpus docs (`idempotency`/`verification`/`demo/frontend-tier`/`snapshot-cold-start`). Go 713→**736** (+23, all M17); Python 174→**360 collected** (net-new stack-verify suite 0→87 + demo-stack/dev-stack/stack-injection growth); flake 0; triple-clean 3/3; supply-chain GREEN (0 called third-party CVEs; all-permissive). close-release GREEN (0 must-fix beyond 2 doc-index rows; deferral re-audit GREEN — DEF-M10-01 → backlog (unscheduled) untouched/not-aged). Records: [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/) (review · retro · metrics · lockfile).
- **v1.3 "stack party"** — **2026-06-07**, tag `v1.3`. The **dev/demo convergence**: a unified first-available-N stack registry [M12], dev-as-peer (the per-stack-Directus recipe + firewall check [print-only — not a working local Directus; see roadmap.md Correction] + auto-snapshot + light seed) [M13], one generic `stack-*` skill set [M14], and a code-cited `corpus/ops/safety.md` with fail-closed drift guards [M15]. 1 signed escape-hatch (DEF-M10-01 → backlog/unscheduled). Records: [releases/archive/01.30-stack-party/](releases/archive/01.30-stack-party/).
- **v1.2 "set dressing"** — **2026-06-07**, tag `v1.2`. The **snapshot mechanism**: a dedicated `stack-snapshot` extension that captures the public reference library read-only from prod, manifest-caches it in `.agentspace`, replays it per-stack behind a tested tenant-data firewall — 100% data-DNA. Records: [releases/archive/01.20-set-dressing/](releases/archive/01.20-set-dressing/).

## Headline numbers (v1.5/M23 close — 2026-06-13; baseline v1.3b 2026-06-09)
- **Go test funcs:** **844** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` 46 · clerkenstein 223 · stack-seeding **259** · stack-snapshot **316**. M23's OWN contribution is **+33** across the two modules it touched (stack-snapshot 290→**316** +26: directus_files referenced-subset + ClearByDelete + firewall admit-iff + ValidateProvisionable; stack-seeding 252→**259** +7: the cross-surface closure gene) — the rest of the +49 vs the M22-close 795 is a counting-method reconciliation on the **untouched** modules (alignment 43→46, clerkenstein 210→223), not new tests. M23 coverage gaps closed by the 2-pass harden: `CrossSurfaceDangling` 0→**100%**, `ValidateProvisionable` 80→**100%**. All Go modules `go vet` clean; flake **0** (5/5 shuffled, both modules).
- **Python tests (M23-touched suites):** **stack-core 61→69** (+8, the §1 dev env-emission / `DIRECTUS_BASE_ADDR` re-point) · **stack-injection 110** ran (8 env-gated skip → 102 active; +the §2/§3 demo cutover + studio-desk minted-token). The other 3 suites (dev-stack 73 · demo-stack 87 · stack-verify 104) are **untouched by M23**. gen_override.py 87% · gen_injected_override.py 99%. Flake **0** (5/5 sequential, stack-core+stack-injection). All CLIs shellcheck-CLEAN; py_compile CLEAN. (M23 counts via `python3 -m unittest discover` — the suites are stdlib unittest.)
- **The v1.3b thesis:** `/demo-up` now produces a **full, populated, verified, demoable** stack — full UI tier (M19) + self-verifying bring-up (M18) + re-run-safe primitives (M17) + auto-set-dress (M20) reusing the dev pass byte-for-byte, all on a 16 GB Mac, **zero platform-repo edits**.
- **Coverage:** **100% of the full data-DNA catalog** (inherited from v1.2; unchanged — nothing waived).
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces (Go 22/22, JS/FAPI 9/9, `@clerk/express` 9/9, deployment/injection 7/7). v1.3b touched no clerkenstein.
- **Test health:** flake **0**; triple-clean **3/3** (4 Go `-race -shuffle` + 5 Python suites); supply-chain GREEN (0 called third-party CVEs; 12 stdlib advisories @go1.25.3 → cleared by go1.25.11+; all-permissive licenses).

## Branch model
**v1.5 IN DEVELOPMENT:** `release/01.50-prop-room` cut from `main` (2026-06-11) at design time, per the canonical
flow (release branch exists from M21 onward so milestone branches have a parent). Milestone branches `m21/…` … `m25/…`
are created from it by `/developer-kit:build-mstone-iters` (M21, iterative) / `/developer-kit:build-milestone` (M22–M25,
section) as each starts. v1.5 extensions tooling will be authored in `.agentspace/rosetta-extensions/` and tagged
`prop-room-mNN`, consumed per-stack at the pinned tag (inheriting the v1.3b head `51a07cb`).
**v1.3b SHIPPED:** `release/01.3b-dress-rehearsal` merged `--no-ff` → `main`, tagged **`v1.3.1`** (2026-06-09); release branch deleted; all 5 milestone branches (`m16/land-fixes` … `m20/lifecycle-convergence`) merged + deleted. The stack tooling lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in the authoring copy at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag. v1.3b extensions markers: **`dress-rehearsal-m16`** @ `e6161b0` · **`m17`** @ `0d36251` · **`m18`** @ `777723a` · **`m19`** @ `4f96ddd` · **`m20`** @ `51a07cb`; extensions `main` at `51a07cb` on `origin`; `stack-demo/rosetta-extensions` consumed @ `dress-rehearsal-m20`. Snapshot payloads live in a gitignored `.agentspace/snapshots/` cache (cloud/S3 store = backlog/unscheduled, DEF-M10-01).
**v1.3 SHIPPED:** tagged **`v1.3`** (2026-06-07). **v1.2 SHIPPED + pushed:** **`v1.2`**. **v1.1 SHIPPED + pushed:** **`v1.1`**. **v1.0 SHIPPED:** `v1.0` (2026-06-03).

_Last updated: 2026-06-13 (**M23 "Content cutover + referential closure" CLOSED** — section, all 6 sections Fate-1,
merged to `release/01.50-prop-room`. The data plane is now CUT OVER: `cms`'s `DIRECTUS_BASE_ADDR` re-points to the
in-network per-stack Directus (`http://directus:8055`) + the prod token is stripped, while the asset plane stays on
prod (`DIRECTUS_PUBLIC_BASE_ADDR` unchanged) so images stay real; studio-desk gets a locally-minted static admin
token. `directus_files` ref capture wired (a new REFERENCED-SUBSET firewall kind + ClearByDelete DELETE-before-
TRUNCATE). Referential closure is MEASURED by a cross-surface gene (full-taxonomy capture makes it maximal; the gene
named the 1 genuine prod residual K-AIFUNX-E658 — an operator-owned data fix, not a tooling gap). 2 inherited M21
deferrals RESOLVED in-milestone. Close clean: 6 findings (0 scope / 0 code-quality / 0 adversarial-new / 1 docs
Fate-1 / 0 tests / 5 decision-triage tags), deferral audit GREEN (2 inherited RESOLVED, 0 repeat/aged). Go 795→844
(+33 M23-own + a count-method reconciliation); python touched stack-core 61→69 + stack-injection 110. Ext tag
`prop-room-m23` set by the orchestrator post-close. Next: build M24. Prior: 2026-06-13 M22 CLOSED (lifecycle half);
2026-06-13 M21 CLOSED (structure half); 2026-06-11 v1.5 designed; 2026-06-09 v1.3b SHIPPED, tag `v1.3.1`.)_
