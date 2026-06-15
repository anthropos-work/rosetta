# Release Review: v1.7 "house lights"

**Date:** 2026-06-15  ·  **Milestones:** M31 → M32  ·  **Overall verdict: GREEN** (clean; zero blocking; zero should-fix code/doc)

> `/developer-kit:close-release` — 9-sweep parallel review across the two-repo split (rosetta docs/planning + the ext
> demo-stack/stack-injection diff). **All 9 sweeps GREEN.** Per-phase: P0 supply-chain GREEN (0 new deps) · P1 scope
> GREEN (M31+M32 100% delivered, 0 unaccounted, M33 backlog-routed) · P1b deferral GREEN · P2 code-quality GREEN
> (shellcheck+py_compile exit 0) · P3 docs GREEN · P3b KB GREEN (index guard exit 0) · P4 tests GREEN · P4b metrics
> GREEN (Go 1027 +0, Python 471 +12, flake 0) · P5 decisions GREEN. The only items are close-release housekeeping
> (state.md tag-SHA reconcile + push tags + the orphaned ext branches) — none gate the tag.

## Scope

- [x] **Two-repo split matches the briefing exactly.** rosetta release diff (`main...release/01.70-house-lights`) = 23 files, +1527/-72, **docs + planning ONLY** (0 `.py`/`.go`/`.sh`/`.ts`). ext diff (`868a68a..7b17c39`) = exactly the **5 claimed files**, +356/-26, all within `demo-stack/` + `stack-injection/` (0 `.go`, 0 platform-repo paths). Confirmed by `git diff --stat` on both repos.
- [x] **M31 (mkcert FAPI cert) — DELIVERED in full.** `up-injected.sh` step 3a-bis branches on `command -v mkcert && [ "$NO_MKCERT" != 1 ]`; the openssl fallback is factored into ONE `gen_openssl_fapi_cert()` (1 defn / 2 calls — no two-copy drift) called from both the absent/opted-out AND mint-fail branches; `DEMO_NO_MKCERT=1` opt-out mirrors the `DEMO_NO_*` family; non-fatal throughout; values-blind (cert logged path-only). Zero-touch on the 3 cert-consumers confirmed by absence. recipe-browser-login.md §B rewritten manual→auto with all 6 caveats; frontend-tier.md + SKILL note/arg-hint present. 11/11 progress boxes checked.
- [x] **M32 (studio-desk single-port) — DELIVERED in full.** `gen_injected_override.py` studio-desk env gains `NODE_ENV=production` + `FRONTEND_PORT=9000` (additive, not `!override`); CORS origin set drops the dead un-offset 9100 → `(3000, 3001, 9000)` (#M32-D2). Regression `test_studio_desk_env_pins_node_env_production` + the harden-added CORS exact-set assertion present. `:9100` doc/CORS sweep complete in frontend-tier.md + SKILL. 8/8 progress boxes checked.
- [x] **M33 (ant-academy liveness) correctly design-routed to the unscheduled backlog**, NOT undelivered-in-release — repro-first rationale recorded in roadmap-vision.md; both milestone overviews list it Out; both deferral audits classify it as release-level routing, not an in-milestone punt.
- [x] **Scope ledger 100% delivered** — 0 UNACCOUNTED, 0 ownerless items.

## Code Quality

- [x] **shellcheck on `up-injected.sh` and `py_compile` on `gen_injected_override.py` both exit 0.** No dead code, no new third-party imports (pure shell + python-stdlib).
- [x] **M31 invariants hold against real code.** cert/key bytes only land in files with all stdout/stderr `>/dev/null`; `mkcert -install` is `|| true`; every failure path is `log "warning: …"` with no `exit` in the cert step (non-fatal contract); PATH-detection via `command -v` (no hard-coded path); keep-existing idempotency guard intact.
- [x] **M32 fix verified against the actual cloned source, not just by assertion.** `studio-desk/src/index.ts:54` `isProduction = NODE_ENV==='production'`; dev block (148–205) does the dead cross-port redirect; production block (213–270) serves every route via `sendFile`+`express.static`+SPA fallback with no cross-port redirect. Base platform docker-compose ships `NODE_ENV=development` (450) + `FRONTEND_PORT=9100` (449), so the additive env block needs exactly the pins M32 added. CORS removal is a true no-op under single-port production.
- [x] **Cross-milestone seam is clean.** M31 touches `up-injected.sh` step 3a-bis; M32 touches the override generator invoked earlier (line 304); cert written before `docker compose up` (line 395); disjoint steps, no shared mutable state.
- [ ] *(nice-to-have, pre-existing)* `FapiCertStep.BODY = open(UP_INJECTED).read()` at class-body scope follows the file's existing unclosed-`open()` pattern (also in GuideDocTruth, TestMigrateRaceGuard) and surfaces ResourceWarnings — no correctness impact; tidy to `with open(...)`/`read_text()` in a future hygiene pass.

## Documentation

- [x] **The three rosetta docs (recipe-browser-login.md §B, frontend-tier.md, demo-up SKILL.md) are coherent end-to-end and every load-bearing claim cross-checks against shipped code.** mkcert SANs `127.0.0.1 localhost ::1` byte-match; openssl `-days 825`; the "825 days / ~2.25y" expiry note accurate; NODE_ENV=production precedence chain + "no route gap" verdict faithful to M32-D1/D4; "302s to dead :9100" correct (Express `res.redirect` default = 302).
- [x] **The `:9100` sweep is COMPLETE.** No stale live/offset studio-desk `:9100` claim remains in the target docs or sibling demo family. The 2 surviving `:9100` mentions in target docs are both correct (the `cors.go` literal describing unmodified platform + the explicit "No offset 9100 origin (M32)" correction). `corpus/services/studio-desk.md` `:9100` refs are native-dev / base-compose context — correctly out of scope and untouched.
- [x] **Cross-references resolve.** No dangling non-prefixed `browser-login.md` reference; all cross-refs use the renamed `recipe-browser-login.md §B step 2`; README family-index + corpus/README index rows intact (index guard exits 0). The "no manual cert step is needed" claim is honestly qualified with the one-time machine-wide `mkcert -install` OS-password carve-out.
- [x] **KB consolidation clean.** No new corpus doc added; no orphan/blind area; cert-trust story has a single anchor (recipe §B) with frontend-tier.md deferring to it and clerkenstein.md staying engine-side without duplication. Both M31/M32 KB-fidelity audits GREEN, all topics PAIRED.

## Tests & Benchmarks

- [x] **All touched suites pass with zero failures and zero flakes across two interpreters** (chosen to defeat the env-masking the M32 diff itself flags). Authoritative tally under PyYAML 6.0.2 (py3.11): test_tooling.py 50/50, test_injection.py 88/88 (0 skipped); under py3.14 (no PyYAML) 8 `@skipIf(yaml is None)` tests env-skip — not failures.
- [x] **Both fixes are regression-locked tightly.** M31: 11-test `FapiCertStep` class covers every branch (mkcert-success / mint-fail→openssl / DEMO_NO_MKCERT→openssl / mkcert-absent→openssl / install-failure-swallow / whitespace-quoting / idempotent-keep-existing / crt-only partial-state / non-fatal / openssl-factored-verbatim / branch-wired). M32: `test_studio_desk_env_pins_node_env_production` (mutation-checked 4 ways, next-web excluded) + the exact-ordered CORS-set assert (catches both over-removal of 3001 and re-add of 9100). The load-bearing compose semantics (`ports: !override` but `environment:` additive) verified to match the generator.
- [x] **Metrics regression gate GREEN — both BLOCKING gates PASS.** Go 1027→1027 (+0, zero `.go` touched); Python 459→471 (+12: M31 +11, M32 +1, distinct suites, no double-count); flake 0; coverage no >2pp drop on any measured surface. test-count-decrease NOT triggered, flake>0 NOT triggered. Alignment gates 100%/100% on all 4 Clerkenstein surfaces (untouched). metrics.json + metrics-history.md v1.7 row both written and internally consistent.
- [ ] *(info, CI carry-forward)* Full YAML-tier coverage (single-port `["29000:9000"]` + DIRECTUS_TOKEN-strip assertions) requires PyYAML present in CI — same local-fallback posture as v1.5/v1.6. The load-bearing behavioral tests are NOT YAML-gated and ran green everywhere.

## Decision Consolidation

- [x] **Every load-bearing decision is blended or correctly scoped not to need blending.** M32-D1 (route-coverage no-gap), D2 (drop dead :9100 CORS), D4 (NODE_ENV precedence chain) VERIFIED present in frontend-tier.md. M31 browser-trust (mkcert + openssl fallback + DEMO_NO_MKCERT + 4 caveats) fully blended into recipe §B + summarized in frontend-tier.md + SKILL. Test-only decisions (M31-D6/D8, M32-D3) correctly need no knowledge blend.
- [x] **No cross-milestone conflict.** The shared "studio-desk single-port :9000 (M32)" framing is consistent across the port table, mechanism block, CORS note, verify registry, SKILL, and both decisions files.
- [x] **All load-bearing code↔decision claims spot-verified against live ext @ 7b17c39** — CORS = exactly `(3000,3001,9000)`; `NODE_ENV=production`+`FRONTEND_PORT=9000` pin; mkcert branch + factored fallback. Every decision corresponds to real, present code.

## Deferrals

- [x] **Zero in-release deferral punts.** Both milestones closed GREEN with 0 open in-milestone punts → an in-release repeat-deferral is structurally impossible. No CHRONIC_DEFER, no DRIFT_DEFER, no AGED_OUT.
- [x] **Both composition-close gates (M31-D7, M32-D5) are genuine resolved-work closures, not disguised punts** — each cites a necessary+sufficient chain backed by merged regression tests confirmed present in ext HEAD (FapiCertStep + the NODE_ENV regression + CORS exact-set asserts). M32 additionally strengthened its close with a live merge-probe.
- [x] **The 5 inherited backlog items (M33, M26, DEF-M10-01, DEF-M21-01, M25-D9) are correctly cross-release/unscheduled in roadmap-vision, freshly re-signed at the 2026-06-15 v1.7 design pass.** v1.7's touched areas (demo-stack mkcert / stack-injection studio-desk) overlap no backlog item's area → no aging trigger fires.

## Supply Chain

- [x] **Zero new third-party dependency surface.** ext diff touches NO dependency manifest (targeted glob diff empty + broad sweep of the full diff found no go.mod/go.sum/go.work/package*.json/lockfiles/requirements*/Pipfile/pyproject/poetry.lock/setup.py/Cargo/Gemfile). M31 = pure shell + stdlib tests; M32 = stdlib generator + stdlib tests.
- [x] **The one third-party Python import (PyYAML in test_injection.py) is PRE-EXISTING** — present at base `868a68a`, untouched by the diff, skip-guarded, unpinned (ext tracks no python manifest). Not a v1.7 addition.
- [x] **mkcert + openssl are HOST binaries, not release dependencies** — PATH-detected via `command -v`, openssl fallback, `DEMO_NO_MKCERT` opt-out; neither is fetched/vendored/installed/pinned (no brew/npm/pip/go install, curl, wget, docker pull, or git clone introduced). One genuine documented tradeoff: `mkcert -install` expands the OS dev-CA trust store on first run.
- [x] **No Go change → all 5 Go modules byte-identical to v1.6; Go test func count 1027 unchanged.** Supply-chain posture GREEN, unchanged from v1.6. `dependencies.lock` written by Phase 0 (4856 bytes) recording the no-new-dependency posture.

## Decision Consolidation — release-level hygiene (close-release-owned, NOT v1.7 scope)

- [ ] **state.md records stale ext-tag SHAs** (`house-lights-m31 @ 6565ef8`, `house-lights-m32 @ 107599c` at lines 37/43/89/107) vs the LIVE tags (`5022e72` / `7b17c39`). The tags advanced after the milestone-close snapshots (the M32 harden commit `7b17c39` postdates the recorded `107599c`). This is a planning-record consistency drift squarely in close-release's remit (it rotates state.md + re-points tags during merge) — flagged so the orchestrator reconciles the SHAs. NOT a decision conflict or unblended decision.
- [ ] **Unpushed ext tags + orphaned branches** — ext carries unpushed tags spanning v1.5/v1.6/v1.7 (prop-room-m21..m26, stage-door-m27/m28/m30, house-lights-m31/m32) and two orphaned local branches (`m26/self-contained-demo` [tag prop-room-m26] + `wip/clerkenstein-browser-login`). Flagged in state.md + both deferral audits as outward-facing carry-over for close-release to push/clean. NOT undelivered v1.7 scope; M26 awaits its own design-roadmap pass.
- [ ] **Cross-sweep timing artifact (info, no action).** Phase 1 reported "no `dependencies.lock` exists"; Phase 0 (supply-chain) subsequently *created* it (4856 bytes). The lock now exists and is the authoritative artifact recording the no-new-dependency posture — Phase 1's note was a point-in-time observation before Phase 0 wrote it. No contradiction in the shipped state.

## Metrics

- [x] Go test funcs: **1027 (+0** vs v1.6 1027) — zero `.go` touched; five-module breakdown 52+223+259+333+160 carries forward verbatim.
- [x] Python tests: **471 (+12** vs v1.6 459) — M31 +11 (FapiCertStep, demo-stack/tests 99→110), M32 +1 (stack-injection test_injection.py 87→88).
- [x] Flake: **0** (both milestones 5/5 randomized-sequential).
- [x] Alignment gates: **100%/100%** on all 4 Clerkenstein surfaces (untouched).
- [x] ext diff: **5 files, +356/-26**; rosetta diff: **23 files, +1527/-72** (docs+planning only).
- [x] Tags: `house-lights-m31 @ 5022e72`, `house-lights-m32 @ 7b17c39`, ext HEAD `7b17c39`.