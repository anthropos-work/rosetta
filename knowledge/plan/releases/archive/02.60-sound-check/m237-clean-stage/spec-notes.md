# M237 — Spec notes

Topic → doc → code triples + clone-freshness / defect-re-triage findings accumulate here during build.

## Pre-flight audits — Clone-freshness fix (first section)
- **Phase 0b KB-fidelity: GREEN** (report: `kb-fidelity-audit.md`, 2026-07-21). All 3 topics PAIRED + ALIGNED; the two Delivers docs truthfully describe the current pre-fix code and name the M237 fixes (`F-M236-CLOSE-1`/`-2`) as their owning findings. No blind areas, no stale load-bearing claims.
- **Triples:**
  - Clone-freshness → `corpus/ops/rosetta_demo.md` §"Clone freshness" → `demo-stack/ensure-clones.sh` (phase c/e) + `platform/Makefile` (`init` skip-if-present / `pull` stash→checkout main→pull --rebase / `status`).
  - R1 sweep → `corpus/ops/demo/demopatch-spec.md` §2.1 + G5 row → `ensure-clones.sh` `PATCH_MANIFESTS` (phase f) + `demo-stack/patches/` (14 dirs).
  - Billion re-triage → `corpus/ops/verification.md` §"PRE-FLIGHT RUNG ZERO" + §"LOGIN shell" → `demo-stack/up-injected.sh`.
- **Test harness for §1/§2:** `demo-stack/tests/test_tooling.py::TestEnsureClonesFunctional` — sandboxes the real script at correct depth, stubs `git`/`make` on PATH, runs it, asserts on `clones.lock.json` + stderr + binlog. Extend the `git` stub for `fetch`/`rev-list --count`.
- **CI fence:** `stack-core/demo_knob_guard.py` scans `ensure-clones.sh` for `${DEMO_*:-…}` knobs (both-directions doc↔parser). A new `DEMO_ADVANCE_CLONES` MUST be written `${DEMO_ADVANCE_CLONES:-0}` AND get a row in `corpus/ops/demo/demo-up-defaults.md` (same as existing `DEMO_ALLOW_UNPINNED_REXT`).

## Clone-freshness mechanism (§1 — landed)
- `rext demo-stack/ensure-clones.sh`: new phase (d3) opt-in advance + rewritten phase (e) provenance+freshness.
  - **Verified fetch**: `git -C <r> fetch --quiet origin` with rc CHECKED, stderr NOT suppressed (`--quiet` trims progress, not errors). Failed fetch → `fetch_ok=false` + `behind=null` + loud warn — never a number off a stale ref.
  - **behind-count**: `git rev-list --count HEAD..origin/<branch>` ONLY after a confirmed fetch.
  - **Pin model** (7 states) recorded in `clones.lock.json`; optional `stack-demo/clones.pin.json` = `{"<repo>":"<ref>"}`.
  - **Knobs**: `DEMO_ADVANCE_CLONES` (0|1|main|pinned), `DEMO_FRESHNESS_STRICT` (0|1). Both documented in `demo-up-defaults.md` (demo_knob_guard fence).
  - Tests: `demo-stack/tests/test_tooling.py::TestCloneFreshnessM237` (12) incl. the 12-vs-202 anti-regression.

## R1 pristine sweep (§2 — landed)
- F-M236-CLOSE-2: `ensure-clones.sh` phase (f) R1 now globs `"$HERE"/patches/*/*.yaml` (all 14), logs `swept N manifest(s)`. Tests: `TestR1SweepM237` (4). Proven live on billion: "swept 14 manifest(s)".

## Fresh-clone billion re-triage (§3/§4 — landed; see decisions.md for the full ledger)
- **Method**: e2e cockpit-login harness (`stack-verify/e2e/lib/cockpit-login.ts loginAs`, seats dan-manager/maya-thriving) run from the workstation against the tailnet HTTPS origins (`https://billion.taildc510.ts.net:{13000 app, 15400 fapi, 13077 academy}`). One-off probe spec run + removed (evidence in decisions.md).
- **Freshness truth**: verified measurement (independently confirmed by raw `git rev-parse`) — clones 0–2 behind (frontend CURRENT); ant-academy the lone 5-behind surface. The "202-behind" premise REFUTED (it was the suppressed-fetch artifact §1 fixes).
- **Ledger**: #1 RESOLVED (manager menu hierarchical: Content Library + Organization group w/ 5 expandables) · #4 does-not-reproduce (library populated "Public Content (22)", 7→29 cards, 0 gql/http errors) · #2 SURVIVES → M238 (academy `/it` 404, flag menu non-functional; separate degraded 5-behind academy app, not next-web staleness).

## Live-proof triples
- Freshness → `corpus/ops/rosetta_demo.md` §"Clone freshness" → `ensure-clones.sh` (d3/e) + `platform/Makefile` (pull).
- R1-all-14 → `corpus/ops/demo/demopatch-spec.md` §2.1 + G5 row → `ensure-clones.sh` (f) + `demo-stack/patches/` (14).
- Knobs → `corpus/ops/demo/demo-up-defaults.md` "Secrets & clones" table → `ensure-clones.sh:183,354`.
