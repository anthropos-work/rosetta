# Release Review: v1.10 "method acting"

**Date:** 2026-06-27
**Milestones:** M39 (profile-identity) · M40 (directus-serve-grant) · M41 (profile-depth) · M42e (employee-coverage, gate) · M42m (manager-coverage, gate) · M43 (cockpit-ux) · M44 (profile-completeness) · M45 (generation-engine, gate) · M46 (org-scale-fill, gate)
**Branch:** `release/01.10-method-acting` → `main`, tag `v1.10`
**Code-of-record (separate repo):** `rosetta-extensions` @ tags `method-acting-m39` · `m40` · `m41` · `m42e` (`53574ae`) · `m42m-harden-final` · `m43-cockpit-ux-fix1` · `m44-profile-completeness-fix2` · `m45-harden-final` · `m46-servegrant-closure`. This close merges ONLY the rosetta doc-half branch; the ext tags ARE the code-of-record.

## Verdict: GREEN — 0 blocking findings

All 9 milestones reviewed as one PR. The 9 review sweeps (supply-chain / scope / deferral re-audit / code-quality / docs / KB-consolidation / tests / metrics-regression / decisions) all clear. The release makes each logged-in hero a fully-fleshed, believable person (profile identity + content-surface unblock + profile depth) proven by a Playwright **semantic** coverage sweep at both vantages cold, then extends to a presenter-grade cockpit + a cheap-LLM generation engine that fills a whole ~500/735-member org from one descriptor, proven live on demo-3. **Tooling + docs only — zero platform-repo edits** (verified: the rosetta release diff — `git diff main...release/01.10-method-acting`, 214 files +13306/-153 — is confined to `corpus/**` + `knowledge/plan/**` + `CLAUDE.md` + `.gitignore`; no `.go`/`.py`/`.sh`/`stack-*` files).

The user authorized close + push ahead of this review; the review records the ledger (no approval-gate block).

## Scope (Phase 1)
- [x] All 9 roadmap milestones (M39→M42m + the M43→M46 extension) delivered Fate-1; Version Plan table + roadmap detail confirm each `done`. No unaccounted items.
- [x] Both per-vantage coverage GATES MET cold: employee (M42e) `(failing=0, escapes=0, …)` → `(0,0,0,0,0)`; manager (M42m / closed by the M46 campaign) `failingSections=0, gateMet=true, personaFailures=0, escapes=0`.
- [x] M42e's intra-milestone Fate-3 routes (iter-02/03 D3/D4/D5: the `/skill-path` external-article escape, the 5 sim-result deep-links, the 2 empty skill-paths) were all routed forward WITHIN M42e and **RESOLVED** in later iters to gate-MET — Fate-3-routed-AND-delivered, not undelivered.
- [x] **DEF-M46-01 RESOLVED** (the Directus serve-grant deep-fetch CLOSURE + the sanctioned prod-structure recapture — M46 Path 2, `method-acting-m46-servegrant-closure`; firewall public-only / 0 tenant rows). Tracked in `roadmap-vision.md`.
- [x] **No Fate-3-undelivered. No NEW escape-hatch. No drops** beyond declared scope. iter-07's `exit-3 (re-scope-trigger)` was SUPERSEDED by the demo-patch campaign (the grids WERE demo-patchable; the milestone closed gate-MET, not re-scoped).

## Supply Chain (Phase 0)
- [x] **The rosetta CORPUS has no package manifest / no deps** (docs-only) — no CVE scan applies.
- [x] **1 NEW DEP this release — deliberate + sanctioned:** `github.com/anthropos-work/ai@v1.40.1` at M45 (the `services/ai` wrapper transport; transitive Azure SDK + openai-go/v3, all MIT/BSD/Apache). The user-acknowledged in-release inflection M45 is ABOUT — v1.8→v1.9 was 0-new-deps; the design-roadmap approved adding the LLM engine + this dep. M46 reuses it unchanged (0 new at M46).
- [x] Playwright (`@playwright/test ^1.49.0`) is the M42e dev/test-only harness dep (the first non-Go rext dep; held the pin — no second non-Go runner). Python stdlib-only + the pre-existing optional PyYAML test dep (M5). No GPL/AGPL anywhere.
- [x] Lockfile written: `releases/01.10-method-acting/dependencies.lock`.

## Deferral Re-Audit (Phase 1b)
- [x] **GREEN** — done inline (no separate skill spawned). Scanned each milestone `decisions.md` + iter `decisions.md` + `roadmap-vision.md`. The standing backlog (**DEF-M10-01** cloud SnapshotStore/S3 blob bytes · **DEF-M21-01** `replayCmd` hermetic test · **M25-D9** dev-`N` taxonomy `rc=4`) is PRE-EXISTING cross-release deferral, already tracked under "Genuinely-deferred work" in `roadmap-vision.md` — NOT a new v1.10 escape-hatch. **DEF-M46-01** terminated RESOLVED in-release.
- [x] Zero open NEW deferrals, zero repeat-patterns, zero aged-out, zero NEW escape-hatch.

## Code Quality (Phase 2 + 2c)
- [x] The rosetta diff is DOCS-ONLY — the code-quality + adversarial phases are rext-side, reviewed + hardened per-milestone in the `.agentspace/rosetta-extensions` authoring copy at the v1.10 tags (vet + gofmt clean; the per-milestone harden passes + adversarial scenario records in each milestone's `decisions.md` / `hardening-ledger.md`).
- [x] Cross-milestone consistency: the M45 seam (CODE owns structure/identity/closure; AI owns content; non-resolving names DROP, never fabricated) is N-invariant and extends unchanged from the bounded M45 batch to the M46 whole-org fill. The three-tier skill draw (role→curated→flat) is consistent across persona/profile/named seeders. No conflicting patterns, no leaked seams.

## Documentation (Phase 3 + 3b)
- [x] **1 finding, FIXED at close** (Explore release-level coherence review): `profile-completeness-spec.md` (M44) was indexed in `corpus/ops/demo/README.md` but ORPHANED from `CLAUDE.md`'s Key-Documentation-Locations index — added an entry between `coverage-protocol.md` and `ai-generation-spec.md` (M44 precedes M45).
- [x] **0 broken cross-refs** across all 14 v1.10-touched corpus docs + `CLAUDE.md` (programmatic relative-`.md`-link resolution check — every target exists on disk).
- [x] **KB consolidation:** the 5 NEW specs — `ai-generation-spec.md` (324L) · `cache-spec.md` (122L) · `cockpit-spec.md` (195L) · `coverage-protocol.md` (425L) · `profile-completeness-spec.md` (118L) — are each under the split threshold, accurate against the code, and indexed from `README.md` + `CLAUDE.md`. The `corpus/ops/demo/` cluster (the demo-family index) is coherent with the actual spec set present. No structural debt.

## Tests & Benchmarks (Phase 4)
- [x] **The rosetta corpus has NO test suite.** The rext tooling tests were run GREEN per-milestone (by the work-agents, at the consumed tags). Go (rext, at the M46 code-of-record HEAD): alignment 52 · clerkenstein 270 · stack-seeding 706 · stack-snapshot 363 · stack-secrets 160 = **1551** (`Test`+`Fuzz`). Python: cockpit `cockpit.py` 63 · demopatch 43 · stack-injection 117 (8 opt-in skipped) + the TS Playwright coverage harness.
- [x] Flake 0 across all milestone gates. M46 substituted a 4-sub-agent demo-patch/recapture verification campaign (each full-suite + the M42 sweep) + a fresh `--purge /demo-up 3` reproducibility proof for a formal `--final` harden. The 1 demo-stack pytest non-pass is a PRE-EXISTING `ensure-clones` SC2015 shellcheck info (line 154, untouched) — not a flake.
- [x] Benchmarks: n/a (DB-IO-bound bulk COPY seeders + stdlib HTTP cockpit + LLM-IO-bound gen-batch with a prompt-hash cache — no rext perf hot path). The org-scale render WALL the M46 gate hit was a PLATFORM render-perf characteristic, addressed by a demo-patch, not a rext hot path.

## Metrics Regression (Phase 4b)
- [x] **GREEN** — vs v1.9 baseline: Go **1248 → 1551 (+303**, no decrease) — stack-seeding 444→706 (+262), stack-snapshot 333→363 (+30), clerkenstein 259→270 (+11, incl the recorded-vs-grep close-drift reconciled to the 270 ground-truth). Python touched suites grew (cockpit 27→63, demopatch 18→43); no suite decreased. Coverage no >2pp drop on any measured surface. Flake 0. Supply-chain: 1 deliberate new dep (not a regression). Aggregated metrics: `releases/01.10-method-acting/metrics.json`; history row appended to `knowledge/plan/metrics-history.md`.

## Alignment Gates
- [x] **100%/100% on all 5 Clerkenstein surfaces** — Go clerk-2.6.0 22/22 · JS clerk-js-5 9/9 · multi clerk-multi-1 9/9 · deploy clerk-deploy-1 7/7 · express clerk-express-1 13/13 + drift 9/9, re-verified at M42e close. M39 (roster org-name thread) + M42e (avatar/org-logo image_url threading) touched clerk-frontend and re-verified GREEN; M40/M41/M43/M44/M45/M46 touched no Clerkenstein contract surface (empty clerkenstein diff m45..m46) — gates carry forward N/A-change.

## Decision Consolidation (Phase 5)
- [x] Load-bearing decisions blended into the corpus with traceable tags: the M39/M40/M41 reference tags in `clerkenstein.md` / `snapshot-spec.md` / `stories-spec.md` / `seeding-spec.md`; the M42e/M42m protocol truths in `coverage-protocol.md` (the fix-surface routing table, the disclosed-presenter-note allow-rule, the `demopatch` model); the M43 cockpit truths in `cockpit-spec.md`; the M44 density truths in `profile-completeness-spec.md`; the M45/M46 protocol truths in `ai-generation-spec.md` + `cache-spec.md` + `snapshot-spec.md` (the GetJobSimulation deep-fetch closure + the column-drift reconciliation). The maintainer-only option-trees stay in the per-milestone `decisions.md`. No cross-milestone decision conflicts.

## CI Gate (Phase 8b)
- [x] **No CI is wired for the corpus + no corpus test suite exists** (the corpus is docs). The rext suites were verified per-milestone at their consumed tags (the local-3x carry-forward — same posture as v1.5/v1.6/v1.7/v1.8/v1.9). CI-wiring for the corpus is a **carry-forward** (there is no executable suite to run; no 3× run is fabricated on a non-existent suite).

## Carry-forward
- CI-wiring for the corpus (no executable suite — rext suites verified per-milestone at their tags).
- Pushing the ext tags + `main` + the `v1.10` tag to origin (the orchestrator's separate post-close step — this close is LOCAL-only).

## Headline

Tooling + docs only — **ZERO platform-repo edits** throughout. All 5 Clerkenstein gates 100%/100%. Significant test growth (rext Go 1248→1551, +303). The believable-profile release proven by a Playwright SEMANTIC coverage sweep at both vantages cold, extended to a presenter-grade cockpit + a cheap-LLM org-scale generation engine. The supply-chain inflection (the first new dep since v1.8) was a deliberate, user-acknowledged in-release decision. **GREEN — 0 blocking, the single doc finding fixed at close.**
