# State

**Active version:** **v1.5 "prop room"** (designed 2026-06-11; branch `release/01.50-prop-room`). The
**local-Directus release** — every stack today reads public content **live from prod**; v1.5 stands up a **local
Directus per stack** serving the **captured public library**, so content is self-contained. **Real images are
preserved** by keeping the asset plane on prod public links (data plane local, asset plane prod). Coverage: **every
demo** stack (default), **any dev-N≥1** (opt-in `--local-content`), **N=0** manual (documented recipe + n=0 guard).
5 milestones M21→M25 (structure capture → executed provisioning + lifecycle → content cutover + referential closure
→ docs + hygiene → field bake). **Tooling + docs only — zero platform-repo edits; capture stays read-only / public-
only / prod-untouched.** Full plan: [`roadmap.md`](roadmap.md) § In Development.
**Active milestone:** **M22 — Executed provisioning + per-stack Directus lifecycle** (`section`; not yet started).
Turn M21's print-only recipe into an **executed** bring-up step that boots a per-stack Directus as a **compose
service** (offset port, torn down with the stack), idempotent + verified — demo default-on / dev opt-in. Build with
**`/developer-kit:build-milestone`** (section shape).
**Last closed:** **M21 — Structure capture — 2026-06-13** (closed-on-gate; gate MET by tooling — `stacksnap`
captures the content-model structure + auto-provisions a bootstrapped stack so the directus replay exits 0 and a
booted Directus serves the captured catalog anonymously; the structure half of the local-Directus story).
**Next up:** start **M22** (`/developer-kit:build-milestone`), then M23→M25 strictly sequential.
**Phase:** **v1.5 building — M21 closed (gate met); ready to build M22 (executed provisioning + lifecycle).**
**Paused:** _(none)_

## Recently shipped releases
_(v1.5 "prop room" is the active release-in-development — see above; the list below is the shipped baseline it builds on.)_
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up` field issues: M16 land-fixes + doc-truth (published the devpath + migrate-race fixes, finished `anthropos-dev → stack-dev`) → M17 re-run safety (replay/seed idempotency + the `set -e` first-run-race) → M18 the verification net (`stack-verify` offset/scope-aware + auto-wired NON-FATAL; "UP" = verified) → M19 the frontend tier (next-web + studio-desk + ant-academy, per-demo cached builds, `--no-ui`, **zero platform-repo edits**) → M20 lifecycle convergence (`/demo-up` auto-set-dresses like `/dev-up` via one shared engine + the cold-start runbook). 4 net-new corpus docs (`idempotency`/`verification`/`demo/frontend-tier`/`snapshot-cold-start`). Go 713→**736** (+23, all M17); Python 174→**360 collected** (net-new stack-verify suite 0→87 + demo-stack/dev-stack/stack-injection growth); flake 0; triple-clean 3/3; supply-chain GREEN (0 called third-party CVEs; all-permissive). close-release GREEN (0 must-fix beyond 2 doc-index rows; deferral re-audit GREEN — DEF-M10-01 → backlog (unscheduled) untouched/not-aged). Records: [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/) (review · retro · metrics · lockfile).
- **v1.3 "stack party"** — **2026-06-07**, tag `v1.3`. The **dev/demo convergence**: a unified first-available-N stack registry [M12], dev-as-peer (the per-stack-Directus recipe + firewall check [print-only — not a working local Directus; see roadmap.md Correction] + auto-snapshot + light seed) [M13], one generic `stack-*` skill set [M14], and a code-cited `corpus/ops/safety.md` with fail-closed drift guards [M15]. 1 signed escape-hatch (DEF-M10-01 → backlog/unscheduled). Records: [releases/archive/01.30-stack-party/](releases/archive/01.30-stack-party/).
- **v1.2 "set dressing"** — **2026-06-07**, tag `v1.2`. The **snapshot mechanism**: a dedicated `stack-snapshot` extension that captures the public reference library read-only from prod, manifest-caches it in `.agentspace`, replays it per-stack behind a tested tenant-data firewall — 100% data-DNA. Records: [releases/archive/01.20-set-dressing/](releases/archive/01.20-set-dressing/).

## Headline numbers (v1.5/M21 close — 2026-06-13; baseline v1.3b 2026-06-09)
- **Go test funcs:** **795** total (`Test`+`Fuzz`; +59 vs the v1.3b-close 736, all in M21's `stack-snapshot`). Per-module: `rosetta-extensions/alignment` 43 · clerkenstein 210 · stack-seeding 252 · stack-snapshot **290** (was 231 — the structure-capture core + auto-provision + the serve half + 2 final harden passes). M21 coverage: directus/firewall **100%**, manifest 98.4%, capture 98.9% (cmd/stacksnap 80.1% + pg 47.0% = live-DB residual). All Go modules build + `go vet` clean; flake **0** (5/5 shuffled).
- **Python tests:** **360** collected (stack-core 54 · dev-stack 50 · demo-stack 84 · stack-injection 85 · stack-verify 87). Grew from v1.3's 174 — the net-new `stack-verify` suite (0→87, M18), demo-stack 13→84, dev-stack 38→50, stack-injection 69→85. _(v1.3 reported test-funcs; v1.3b reports pytest-collected — growth is real either way, no section regressed.)_ All 3 CLIs shellcheck-CLEAN; py_compile CLEAN.
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

_Last updated: 2026-06-13 (**M21 "Structure capture" CLOSED** — closed-on-gate, merged to `release/01.50-prop-room`.
The collection-schema gap is closed at the tooling level: `stacksnap` captures the directus content-model structure
(DDL + PKs + serve rows) behind the firewall structural-metadata carve-out + auto-provisions a bootstrapped stack →
replay exits 0 → anonymous serve. Net-new `corpus/ops/directus-local.md` (the structure half). Deferral audit GREEN —
`directus_files` → Fate-3 M23. Ext tag `prop-room-m21` set by the orchestrator post-close. Next: M22. Prior:
2026-06-11 v1.5 designed + staged; 2026-06-09 v1.3b SHIPPED, tag `v1.3.1`.)_
