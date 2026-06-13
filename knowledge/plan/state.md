# State

**Active version:** **v1.5 "prop room"** (designed 2026-06-11; branch `release/01.50-prop-room`). The
**local-Directus release** — every stack today reads public content **live from prod**; v1.5 stands up a **local
Directus per stack** serving the **captured public library**, so content is self-contained. **Real images are
preserved** by keeping the asset plane on prod public links (data plane local, asset plane prod). Coverage: **every
demo** stack (default), **any dev-N≥1** (opt-in `--local-content`), **N=0** manual (documented recipe + n=0 guard).
5 milestones M21→M25 (structure capture → executed provisioning + lifecycle → content cutover + referential closure
→ docs + hygiene → field bake). **Tooling + docs only — zero platform-repo edits; capture stays read-only / public-
only / prod-untouched.** Full plan: [`roadmap.md`](roadmap.md) § In Development.
**Active milestone:** **M23 — Content cutover + referential closure** (`section`; **BUILT — all 6 sections landed,
not yet closed**). Points the stack's services at their **own** Directus (re-point `cms`'s `DIRECTUS_BASE_ADDR`;
asset plane stays on prod) and guarantees the served catalog is **referentially closed** (a measured cross-surface
gene — no content row references a taxonomy node-id the captured taxonomy lacks; surfaced the 1 real prod residual).
Wired the `directus_files` ref capture (Fate-3 from M21, as a referenced-subset) + studio-desk's local-instance +
minted token; the 20 dangling relations were subsumed by M21's 26-collection structure capture. Next:
**`/developer-kit:close-milestone`** (harden optional first).
**Last closed:** **M22 — Executed provisioning + per-stack Directus lifecycle — 2026-06-13** (`section`; all 6
sections Fate-1). Turned M21's print-only recipe into an **executed** bring-up step that boots a per-stack Directus as
a **compose service** (offset port, app-network, torn down with the stack), idempotent + verified — demo default-on /
dev opt-in (`--local-content`); the `EnvContract` firewall became a **load-bearing executed gate** (prod env
hard-aborts before any write), non-fatal degrade to the prod-read path.
**Next up:** build **M23** (`/developer-kit:build-milestone` — content cutover + referential closure), then M24→M25 strictly sequential.
**Phase:** **v1.5 building — M23 BUILT (all 6 sections on `m23/content-cutover`, not yet closed); next close M23, then M24→M25.**
**Paused:** _(none)_

## Recently shipped releases
_(v1.5 "prop room" is the active release-in-development — see above; the list below is the shipped baseline it builds on.)_
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up` field issues: M16 land-fixes + doc-truth (published the devpath + migrate-race fixes, finished `anthropos-dev → stack-dev`) → M17 re-run safety (replay/seed idempotency + the `set -e` first-run-race) → M18 the verification net (`stack-verify` offset/scope-aware + auto-wired NON-FATAL; "UP" = verified) → M19 the frontend tier (next-web + studio-desk + ant-academy, per-demo cached builds, `--no-ui`, **zero platform-repo edits**) → M20 lifecycle convergence (`/demo-up` auto-set-dresses like `/dev-up` via one shared engine + the cold-start runbook). 4 net-new corpus docs (`idempotency`/`verification`/`demo/frontend-tier`/`snapshot-cold-start`). Go 713→**736** (+23, all M17); Python 174→**360 collected** (net-new stack-verify suite 0→87 + demo-stack/dev-stack/stack-injection growth); flake 0; triple-clean 3/3; supply-chain GREEN (0 called third-party CVEs; all-permissive). close-release GREEN (0 must-fix beyond 2 doc-index rows; deferral re-audit GREEN — DEF-M10-01 → backlog (unscheduled) untouched/not-aged). Records: [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/) (review · retro · metrics · lockfile).
- **v1.3 "stack party"** — **2026-06-07**, tag `v1.3`. The **dev/demo convergence**: a unified first-available-N stack registry [M12], dev-as-peer (the per-stack-Directus recipe + firewall check [print-only — not a working local Directus; see roadmap.md Correction] + auto-snapshot + light seed) [M13], one generic `stack-*` skill set [M14], and a code-cited `corpus/ops/safety.md` with fail-closed drift guards [M15]. 1 signed escape-hatch (DEF-M10-01 → backlog/unscheduled). Records: [releases/archive/01.30-stack-party/](releases/archive/01.30-stack-party/).
- **v1.2 "set dressing"** — **2026-06-07**, tag `v1.2`. The **snapshot mechanism**: a dedicated `stack-snapshot` extension that captures the public reference library read-only from prod, manifest-caches it in `.agentspace`, replays it per-stack behind a tested tenant-data firewall — 100% data-DNA. Records: [releases/archive/01.20-set-dressing/](releases/archive/01.20-set-dressing/).

## Headline numbers (v1.5/M22 close — 2026-06-13; baseline v1.3b 2026-06-09)
- **Go test funcs:** **795** total (`Test`+`Fuzz`; +59 vs the v1.3b-close 736, all in M21's `stack-snapshot`). Per-module: `rosetta-extensions/alignment` 43 · clerkenstein 210 · stack-seeding 252 · stack-snapshot **290**. **Unchanged by M22** (M22 touches no Go — all its code is shell + python). M21 coverage held: directus/firewall **100%**, manifest 98.4%, capture 98.9% (cmd/stacksnap 80.1% + pg 47.0% = live-DB residual). All Go modules build + `go vet` clean; flake **0** (5/5 shuffled).
- **Python tests:** **418** collected (+8 env-gated skip) — stack-core **61** · dev-stack **73** · demo-stack **87** · stack-injection **93** · stack-verify **104**. Grew **+58 from the M21-close 360** (M22: the executed provision + compose-service emission + idempotent guards + verify probes + preflight + the 4-pass harden). gen_override.py coverage 62→**85%**; gen_injected 99%. Flake **0** (5/5 sequential, dev-stack+stack-verify). All CLIs shellcheck-CLEAN; py_compile CLEAN.
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

_Last updated: 2026-06-13 (**M22 "Executed provisioning + per-stack Directus lifecycle" CLOSED** — section, all 6
sections Fate-1, merged to `release/01.50-prop-room`. The print-only recipe is now EXECUTED: a per-stack Directus
boots as a compose service (offset port, app-network, torn down with the stack), provisioned idempotently + verified,
demo default-on / dev opt-in; the `EnvContract` firewall is a load-bearing executed gate (prod env hard-aborts before
any write), non-fatal degrade to prod-read. Close clean: 8 findings (0 must-fix; 3 should-fix comments + DOC-1 landed;
7 adversarial scenarios all already test-pinned), deferral audit GREEN (0 M22-originated). Python 360→418 (+58); Go
795 unchanged. Ext tag `prop-room-m22` set by the orchestrator post-close. Next: build M23. Prior: 2026-06-13 M21
CLOSED (structure half); 2026-06-11 v1.5 designed; 2026-06-09 v1.3b SHIPPED, tag `v1.3.1`.)_
