# State

**Active version:** **v1.5 "prop room"** (designed 2026-06-11; branch `release/01.50-prop-room`). The
**local-Directus release** — every stack today reads public content **live from prod**; v1.5 stands up a **local
Directus per stack** serving the **captured public library**, so content is self-contained. **Real images are
preserved** by keeping the asset plane on prod public links (data plane local, asset plane prod). Coverage: **every
demo** stack (default), **any dev-N≥1** (opt-in `--local-content`), **N=0** manual (documented recipe + n=0 guard).
5 milestones M21→M25 (structure capture → executed provisioning + lifecycle → content cutover + referential closure
→ docs + hygiene → field bake). **Tooling + docs only — zero platform-repo edits; capture stays read-only / public-
only / prod-untouched.** Full plan: [`roadmap.md`](roadmap.md) § In Development.
**Active milestone:** **M25 — Field bake: the observable-behavior gate** (`section`; **planned — next to build**). Prove the whole release on the actual 16 GB box with **observable behaviors** as the done-bars (fresh `/demo-up` → the browser shows content served by the local Directus with real images + the verify net GREEN incl. the new Directus probes; `/dev-up 2 --local-content` same on an opt-in dev stack + N=0 stays prod-read; re-run everything twice; the cold-start capture once; clean teardown reclaims the Directus container/port) — pre-paying the field-fix tail every prior release shipped after the fact. The closing v1.5 milestone.
**Last closed:** **M24 — Docs convergence + hygiene strand — 2026-06-13** (`section`; all 7 sections Fate-1).
Made the corpus tell the new truth: corrected the **fictional local-Directus** docker-service claims (image
`directus/directus:10.10.1` + admin/password + a compose snippet — **never existed** in the platform compose,
verified) across `external_services.md`/`service_taxonomy.md`/`quick_ops.md`, rewrote the known-state +
`directus-local.md` + `snapshot-spec.md` on the M23 cutover, swept the print-only/exit-4/live-from-prod framing
across 5 skills + `CLAUDE.md` + 5 demo docs. Absorbed the 4 aged-out hygiene items Fate-1: the `toolchain go1.25.11`
pin (clears the 12 called-stdlib advisories), the corpus README-index-row guard (token-bounded; dog-fooded → 7 gaps
backfilled; harden fixed a prefix-collision false-negative), the alignment zero-critical-genes guard (Validate +
GateMet defence-in-depth), and the `/project-stats` scope fix (drop the gitignored `stack-*/` clones — landed
cross-repo in developer-kit `825cdce`). ext tag `prop-room-m24`.
**Next up:** build **M25** (`/developer-kit:build-milestone` — the field-bake observable-behavior gate); it's the final v1.5 milestone, then close-release.
**Phase:** **v1.5 building — M24 CLOSED (merged to `release/01.50-prop-room`); next + last: build M25, then close-release.**
**Paused:** _(none)_

## Recently shipped releases
_(v1.5 "prop room" is the active release-in-development — see above; the list below is the shipped baseline it builds on.)_
- **v1.3b "dress rehearsal"** — **2026-06-09**, tag `v1.3.1`. The **field-hardening release** for the 14 `/demo-up` field issues: M16 land-fixes + doc-truth (published the devpath + migrate-race fixes, finished `anthropos-dev → stack-dev`) → M17 re-run safety (replay/seed idempotency + the `set -e` first-run-race) → M18 the verification net (`stack-verify` offset/scope-aware + auto-wired NON-FATAL; "UP" = verified) → M19 the frontend tier (next-web + studio-desk + ant-academy, per-demo cached builds, `--no-ui`, **zero platform-repo edits**) → M20 lifecycle convergence (`/demo-up` auto-set-dresses like `/dev-up` via one shared engine + the cold-start runbook). 4 net-new corpus docs (`idempotency`/`verification`/`demo/frontend-tier`/`snapshot-cold-start`). Go 713→**736** (+23, all M17); Python 174→**360 collected** (net-new stack-verify suite 0→87 + demo-stack/dev-stack/stack-injection growth); flake 0; triple-clean 3/3; supply-chain GREEN (0 called third-party CVEs; all-permissive). close-release GREEN (0 must-fix beyond 2 doc-index rows; deferral re-audit GREEN — DEF-M10-01 → backlog (unscheduled) untouched/not-aged). Records: [releases/archive/01.3b-dress-rehearsal/](releases/archive/01.3b-dress-rehearsal/) (review · retro · metrics · lockfile).
- **v1.3 "stack party"** — **2026-06-07**, tag `v1.3`. The **dev/demo convergence**: a unified first-available-N stack registry [M12], dev-as-peer (the per-stack-Directus recipe + firewall check [print-only — not a working local Directus; see roadmap.md Correction] + auto-snapshot + light seed) [M13], one generic `stack-*` skill set [M14], and a code-cited `corpus/ops/safety.md` with fail-closed drift guards [M15]. 1 signed escape-hatch (DEF-M10-01 → backlog/unscheduled). Records: [releases/archive/01.30-stack-party/](releases/archive/01.30-stack-party/).
- **v1.2 "set dressing"** — **2026-06-07**, tag `v1.2`. The **snapshot mechanism**: a dedicated `stack-snapshot` extension that captures the public reference library read-only from prod, manifest-caches it in `.agentspace`, replays it per-stack behind a tested tenant-data firewall — 100% data-DNA. Records: [releases/archive/01.20-set-dressing/](releases/archive/01.20-set-dressing/).

## Headline numbers (v1.5/M24 close — 2026-06-13; baseline v1.3b 2026-06-09)
- **Go test funcs:** **850** total (`Test`+`Fuzz`). Per-module: `rosetta-extensions/alignment` **52** · clerkenstein 223 · stack-seeding 259 · stack-snapshot 316. M24's OWN contribution is **+6**, all in **alignment** (46→**52**): the zero-critical-genes guard (+3 `dna.Validate` rejects-zero-critical / variant-vs-capability counting / Load→Validate; +3 `compare.GateMet` refuses-vacuous-critical-gate / `CriticalGenes` count / still-gates-for-real). The other 3 modules are **untouched by M24**. M24 coverage: `dna` 98.5→**100%**; the touched `compare` code (`GateMet`/`CriticalGenes`) fully exercised. `go vet` + `gofmt` clean; flake **0** (5/5 shuffled, dna+compare).
- **Python tests (M24-touched suite):** **stack-core 77→85** (+8, the new `corpus_index_guard.py` README-index-row lint — 8 build + 8 harden incl. the prefix-collision regression). The other 4 suites (dev-stack · demo-stack · stack-injection 110 · stack-verify) are **untouched by M24**. The live corpus README-index guard runs **exit 0** (every doc indexed, incl. the M24 docs + the 7 backfilled gaps). Flake **0** (5/5 sequential). py_compile CLEAN. (Counts via `python3 -m unittest discover` — stdlib unittest.)
- **The v1.3b thesis:** `/demo-up` now produces a **full, populated, verified, demoable** stack — full UI tier (M19) + self-verifying bring-up (M18) + re-run-safe primitives (M17) + auto-set-dress (M20) reusing the dev pass byte-for-byte, all on a 16 GB Mac, **zero platform-repo edits**.
- **Coverage:** **100% of the full data-DNA catalog** (inherited from v1.2; unchanged — nothing waived).
- **Alignment gates (held green since v1.0):** **100%/100%** on all 4 Clerkenstein surfaces (Go 22/22, JS/FAPI 9/9, `@clerk/express` 9/9, deployment/injection 7/7). v1.3b touched no clerkenstein.
- **Test health:** flake **0**; triple-clean **3/3** (4 Go `-race -shuffle` + 5 Python suites); supply-chain GREEN (0 called third-party CVEs; the 12 stdlib advisories @go1.25.3 now **pinned out** by M24's explicit `toolchain go1.25.11` on all 4 go.mod + the clerkenstein CI; all-permissive licenses).

## Branch model
**v1.5 IN DEVELOPMENT:** `release/01.50-prop-room` cut from `main` (2026-06-11) at design time, per the canonical
flow (release branch exists from M21 onward so milestone branches have a parent). Milestone branches `m21/…` … `m25/…`
are created from it by `/developer-kit:build-mstone-iters` (M21, iterative) / `/developer-kit:build-milestone` (M22–M25,
section) as each starts. v1.5 extensions tooling is authored in `.agentspace/rosetta-extensions/` and tagged
`prop-room-mNN`, consumed per-stack at the pinned tag (inheriting the v1.3b head `51a07cb`). v1.5 ext markers so
far: **`prop-room-m21`** · **`m22`** · **`m23`** · **`m24`** @ `6a4749d` (M24-close head; M24's §7 `/project-stats`
fix is the cross-repo developer-kit `825cdce`, outside the ext tag).
**v1.3b SHIPPED:** `release/01.3b-dress-rehearsal` merged `--no-ff` → `main`, tagged **`v1.3.1`** (2026-06-09); release branch deleted; all 5 milestone branches (`m16/land-fixes` … `m20/lifecycle-convergence`) merged + deleted. The stack tooling lives in the **private** `anthropos-work/rosetta-extensions` monorepo — authored + tagged in the authoring copy at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag. v1.3b extensions markers: **`dress-rehearsal-m16`** @ `e6161b0` · **`m17`** @ `0d36251` · **`m18`** @ `777723a` · **`m19`** @ `4f96ddd` · **`m20`** @ `51a07cb`; extensions `main` at `51a07cb` on `origin`; `stack-demo/rosetta-extensions` consumed @ `dress-rehearsal-m20`. Snapshot payloads live in a gitignored `.agentspace/snapshots/` cache (cloud/S3 store = backlog/unscheduled, DEF-M10-01).
**v1.3 SHIPPED:** tagged **`v1.3`** (2026-06-07). **v1.2 SHIPPED + pushed:** **`v1.2`**. **v1.1 SHIPPED + pushed:** **`v1.1`**. **v1.0 SHIPPED:** `v1.0` (2026-06-03).

_Last updated: 2026-06-13 (**M24 "Docs convergence + hygiene strand" CLOSED** — section, all 7 sections Fate-1,
merged to `release/01.50-prop-room`. The corpus now tells the new truth: the **fictional local-Directus** docker
service (image `directus/directus:10.10.1` + admin/password + a compose snippet — it **never existed** in the
platform compose, verified) is corrected across `external_services.md`/`service_taxonomy.md`/`quick_ops.md`; the
known-state + `directus-local.md` + `snapshot-spec.md` rewritten on the M23 cutover; the print-only/exit-4/live-
from-prod framing swept across 5 skills + `CLAUDE.md` + 5 demo docs (the real two-path posture: prod-read default,
per-stack local Directus on `--local-content`). The 4 aged-out hygiene items all Fate-1: the `toolchain go1.25.11`
pin (clears the 12 called-stdlib advisories), the corpus README-index-row guard (token-bounded; dog-fooded → 7 gaps
backfilled; harden fixed a prefix-collision false-negative), the alignment zero-critical-genes guard (Validate +
GateMet defence-in-depth), and the `/project-stats` scope fix (cross-repo developer-kit `825cdce`). Close clean:
5 findings (0 scope / 0 must-fix code-quality / 0 docs / 0 tests / 1 adversarial-record / 3 decision-triage),
deferral audit GREEN (0 new deferrals, 4 chartered hygiene items cleared, 3 standing inherited unchanged + not
aged). Go 844→850 (+6, all alignment); python stack-core 77→85 (+8). Ext tag `prop-room-m24` @ `6a4749d` set by
the orchestrator post-close. Next + last: build M25 (field bake), then close-release. Prior: 2026-06-13 M23 CLOSED
(content cutover); 2026-06-13 M22 CLOSED (lifecycle half); 2026-06-13 M21 CLOSED (structure half); 2026-06-11 v1.5
designed.)_
