# Release Review: v1.8 "understudy"

**Date:** 2026-06-15
**Milestones:** M26 (Self-contained demo stacks) — single `section` milestone
**Verdict:** **GREEN to tag** — 0 blocking. One YELLOW (doc-coherence) landed Fate-1 below.
**Method:** 6 parallel review sweeps (Phases 0–5 + 1b) + an adversarial completeness/coherence critic
(`.agentspace` workflow `wf_8ea12d54-cb0`). The critic independently re-verified diff stats, must-preserve
evidence (grepped the M26 diff for mkcert/FAPI/NODE_ENV/FRONTEND_PORT/9100/CORS → EMPTY), and the tag finalize.

## Scope — GREEN
- 100% delivered, 0 unaccounted. Verified end-to-end against ext `main` @ `773184f` (not just the close report):
  `ensure-clones.sh` (5 phases), the 4-script `$DEV`→`stack-demo` repoint, the D-MAIN PLAT move, the
  `reuse_dev_images` opt-in gate, the ported/retargeted tests, the corpus doc-half, the `stack-secrets/SKILL.md:21`
  M26 note. All Fate-1; 0 deferrals; no Fate-3/escape-hatch/dropped.
- User-signed deferral (not a gap): the **live field-bake** on a freshly-emptied `stack-demo/` — the user chose
  "straight to close-release"; M26 satisfied the observable-behavior gate by composition (M31/M32 precedent).

## Supply chain — GREEN
- **Zero new third-party deps.** Ext Go module diff `7b17c39..main` over `*/go.mod`/`*/go.sum` = EMPTY; M26 touched
  only `.sh`/`.py`/`.md`. Go test-func count 1027 (unchanged). Python: stdlib-only + the pre-existing optional
  PyYAML test dep (added M5, guarded). Rosetta side: docs only, 0 manifests. CVE surface unchanged vs v1.7.
- Lockfile: `dependencies.lock` written — v1.8 inherits v1.7's surface unchanged.

## Code Quality — GREEN
- shellcheck clean ×5 demo-stack scripts; `py_compile` clean ×6; `gofmt -l` clean ×5 modules (M26 touched 0 Go).
- Must-preserve intact (independently grep-verified): M30 BASE_ENV provision, M31 mkcert branch, M32
  NODE_ENV=production/FRONTEND_PORT=9000/dropped-:9100-CORS, injection-targets-only-the-COPY.

## Documentation — YELLOW → fixed Fate-1
- [x] **[should-fix] `corpus/ops/rosetta_demo.md` M16 section stale.** Claimed the `anthropos-dev` fallback survives
  in `up-injected.sh`/`migrate-demo.sh`/`rosetta-demo` — but M26 removed it from all 4 demo scripts (it now lives
  dev-side only: `dev-stack` + `clone_repos.py --dev-root` help). → Narrowed the section + cross-linked the M26
  "Self-contained demo stacks" section.
- [x] **[should-fix] `.claude/skills/demo-up/SKILL.md` missed the headline M26 behavior.** The skill that drives
  `/demo-up` didn't mention `ensure-clones.sh` / the self-contained clone model. → Added an `ensure-clones` step-0
  note to the bring-up steps.
- [x] **[should-fix] Stale "legacy stack-dev/.env base" claim for `DEMO_NO_PROVISION=1` in two skill docs**
  (`demo-up/SKILL.md` L40/L51 + `stack-secrets/SKILL.md` L36). After D-MAIN the opt-out base is the
  ensure-clones-seeded `stack-demo/platform/.env`, not stack-dev. → Reworded in both (matches the ext
  `up-injected.sh` comment fix close-milestone already landed).
- KB consolidation GREEN: README-index guard exit 0; no doc >1500 lines; no orphan; cross-refs resolve.

## Tests & Benchmarks — GREEN
- Go 1027 (unchanged — 0 Go touched). Python: demo-stack 138 + stack-injection 113 (JUnit-authoritative under
  PyYAML); the live-docker `TestMigrateRaceLive` class shows env-flaky timing FAILs in a bare sandbox (not an M26
  regression — the 135 static+functional demo-stack tests are green). Flake gate 5/5 each, zero flakes.
- Metrics regression vs v1.7 (Go 1027, Python 471): GREEN — count went UP (M26 added functional + shell-seam tests),
  flake 0, no coverage drop. No regression on any gate.

## Decision Consolidation — GREEN
- 7 decisions blended (D-MAIN/D1/D2/D3/D5/D6 build-time across the corpus + safety.md §2.8 for D5's bash-3.2
  invariant; D4 close-time into safety.md §2.7). 0 unblended, 0 conflicts. Deferral re-audit (Phase 1b): 0
  deferrals, 0 repeat-defer, 0 aged-out. Inherited backlog (M33, DEF-M10-01, DEF-M21-01, M25-D9) all orthogonal, carry.

## Reconciliation note (adversarial critic)
- The milestone `metrics.json` records the mid-close ext snapshot (`+979/-157`, tag at `17971c1`); the **final**
  ext state after the orchestrator's documented post-close tag-repoint is `773184f` (`+985/-161`), `understudy-m26`
  → `773184f`. Not a discrepancy — the repoint is documented in the milestone record; recorded here + in the
  aggregated release `metrics.json`.
