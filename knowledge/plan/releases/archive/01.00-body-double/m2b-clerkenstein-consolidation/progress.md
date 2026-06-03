# M2b — progress

**Milestone:** M2b — Clerkenstein repo consolidation + knowledge base · **Shape:** section
**Status:** `archived` (completed 2026-06-03; closed + merged to `release/01.00-body-double`)
**Green-gate invariant:** after every code/script section → Go gate 22/22, JS gate 9/9, drift-test 9/9, `-race` clean.
**Open user action (not a deferral, M2b-D3/D8):** run `/singularity-kit:repo-consolidate code` at the clerkenstein repo to formally finalize (should report compliant).

## Sections

### S1 — Restructure into library-named dirs ✅
- [x] Create `shared/` and move the universal-key JWT + `Claims` + mint/parse out of `authn/` (package `shared`) — `parse`→`Parse` exported (M2b-D4)
- [x] `authn/` keeps provider + user, imports `shared` (still mocks `colony/authn`; `var _ authn.Provider` holds)
- [x] Merge `bapi/` + `orgclient/` → `clerk-backend/` (`package clerkbackend`) — `orgclient.X`→`X`, import dropped
- [x] `fapi/` → `clerk-frontend/` (`package clerkfrontend`), imports `shared` (mint side) — testdata fuzz corpus moved too
- [x] `webhook/` → `clerk-webhook/` (`package clerkwebhook`)
- [x] Move runners + assets → `alignment/{cmd,dna,golden,golden-js,scripts}/`
- [x] Repoint all imports + `gate.sh`/`drift-check.sh`/`drift-test.sh` defaults + CI YAML (ALIGN_DIR depth +1, M2b-D5)
- [x] **Gates green** (Go 22/22, JS 9/9, drift-test 9/9, `-race` clean, gofmt/vet/shellcheck) — verified post-move
- [x] Tests co-located within each library dir (no stragglers); empty old dirs removed; binaries re-gitignored (M2b-D6)

### S2 — Knowledge base + per-library READMEs ✅
- [x] `knowledge/{kb-index,scope,architecture,alignment,injection,coverage-index}.md` (real content, authored to repo-consolidate code-repo standard)
- [x] Per-library `README.md`: `authn/ clerk-backend/ clerk-frontend/ clerk-webhook/ shared/ alignment/`
- [x] Slim root `README.md` → points to `knowledge/kb-index.md`
- [x] Cross-links resolve (verified); injection recipes label each surface (recipe-only / spike-proven / built+gated); API claims verified vs code

### S3 — Hygiene ✅
- [x] `.agentspace/` dir created with tracked `.gitkeep` + README; rule switched to `.agentspace/*` + negations (dir preserved, contents ignored)
- [x] `.gitignore` cleanup: fixed mismatched comment (S1), added OS/editor baseline (.DS_Store/Thumbs.db/*.swp/*~/~$*.*)
- [x] Asset-hygiene: no tracked secrets/transient/binaries; runner binaries gitignored at both root + `alignment/`; all tracked files text

### S4 — Consolidate (⚠ user-invoked) ✅ (authored to standard; formal run = user finalize)
- [x] Authored repo TO repo-consolidate standard (skill is disable-model-invocation, M2b-D3/D8): slim `CLAUDE.md` + `singularity-manifest.md` (code/library/go-library, version 0.1.1) + code-area coverage map in `kb-index.md`
- [x] Self-audit vs base + code-repo(library) + asset-hygiene: compliant (entry point, CLAUDE.md, .gitignore, .agentspace preserved, required areas covered, tests present, no tracked secrets/binaries)
- [x] **Gates + drift re-verified green** (S4 docs-only; 7 pkgs ok, gofmt/shellcheck clean)
- [ ] **USER ACTION (not a deferral):** run `/singularity-kit:repo-consolidate code` at the clerkenstein repo to formally finalize — should report compliant + refresh `CLAUDE.md`/`singularity-manifest.md`

### S5 — Rosetta side ✅
- [x] Slim `corpus/services/clerkenstein.md` (197→62 lines) → pointer at the repo's `knowledge/` + the new library-named structure; M1b drift-runbook pointer kept
- [x] Fixed 2 stale refs in `alignment_testing.md` (script path → `alignment/scripts/`; 2 dead anchor links → repointed)
- [x] Milestone records consistent (overview/decisions/progress reflect what shipped; M2b-D4..D8 recorded)

## Running ledger
- **S1** (clerkenstein `512bd49`): restructured flat layout → `authn/ clerk-backend/ clerk-frontend/
  clerk-webhook/ shared/ alignment/` (69 git-mv renames, history preserved). Merged orgclient into
  clerk-backend; extracted shared JWT (mint=clerk-frontend, verify=authn); repointed imports + scripts +
  CI. Green-gate confirmed post-move (Go 22/22, JS 9/9, drift 9/9, lints clean). Decisions M2b-D4..D7.
- **S2** (clerkenstein `7105e1e`): authored the repo's own `knowledge/` base (kb-index + scope +
  architecture + alignment + injection + coverage-index) + 6 per-library READMEs + slim root README, to
  the repo-consolidate code-repo (library) standard. Cross-links + API claims verified.
- **S3** (clerkenstein `e38d089`): `.agentspace/` preserved (`.gitkeep`+README, `.agentspace/*`+negation);
  `.gitignore` OS/editor baseline; asset-hygiene clean (no tracked secrets/transient/binaries).
- **S4** (clerkenstein `0015f8d`): authored to repo-consolidate standard — slim `CLAUDE.md` +
  `singularity-manifest.md` (code/library/go-library, v0.1.1) + code-area coverage map. Self-audit
  compliant. Formal `/singularity-kit:repo-consolidate code` run is a USER finalize (M2b-D3/D8).
- **S5** (rosetta `492fc84`): slimmed `corpus/services/clerkenstein.md` 197→62 lines to a pointer at the
  repo's `knowledge/`; fixed 2 stale refs in `alignment_testing.md`. M1b drift-runbook pointer kept.

## M2b: Hardening

### Pass 1 — 2026-06-03
**Scope manifest (milestone-touched code — clerkenstein repo, separate git):** the reorg is 69
history-preserving renames; the substantive surface is whether coverage/behavior held on the new
library-named layout. 7 Go packages, all tests co-located, all 5 fuzz harnesses travelled with their
seed corpora. Genuinely-new surface = the `shared/` extraction boundary (M2b-D4: `parse`→`Parse`
exported because mint now lives in `clerk-frontend`, verify in `authn`).

**Coverage (milestone-touched packages — confirms the move did NOT lose coverage vs M2 baseline):**
| Package | M2 baseline | M2b new layout | Verdict |
|---|---|---|---|
| `authn` | ~100% | 100.0% | held |
| `shared` (jwt) | ~100% | 100.0% | held |
| `clerk-frontend` (fapi) | ~100% | 100.0% | held |
| `clerk-backend` (bapi+orgclient **merged**) | ~96–98% | 97.4% | held — merge exposed no gap |
| `clerk-webhook` | ~96% | 95.6% | held |
| `alignment/cmd/clerkrun` | ~97% | 96.8% | held |
| `alignment/cmd/jsfapirun` | ~94% | 93.8% | held |

Every test travelled with its code; no regression from the 69 renames. Statement-coverage delta vs
pass-start: **+0%** on all packages — by design. This is a pure reorg; the anti-pattern guard forbids
manufacturing shallow tests on unchanged moved code to bump a number, so the one test added deepens
*behavioral* coverage of the new cross-package seam, not line count.

**Tests added:**
- `authn/authn_test.go`: 1 regression (`TestGetUserErrorClassesCrossBoundary`, 4 sub-cases) — pins the
  three error classes (`malformed` ×2 shapes, `bad-signature`, `expired`) surviving the
  `shared.Parse → authn.GetUser` package boundary with their bare-string identity intact (the runner +
  VerifyToken gene compare `err.Error()` strings per M2b-D4), plus the nil-user-on-error contract.
  Commit `c3649ab`.

**Merge scan (clerk-backend = bapi + orgclient now one package):** `Store` and `Server` coexist without
collision (`go vet` clean); `orgclient.X→X` rename applied cleanly; 34 test/fuzz funcs (both test sets)
present; the merge *removed* a cross-package import (server now references in-package `Store`). No
merge-exposed path; nothing to deepen.

**Fuzz harnesses (confirmed travelled + crash-free on new layout, ~3s each):** `shared/FuzzParse`
(923k execs, 0 crashes), `clerk-frontend/{FuzzParsePublishableKey,FuzzMintParseRoundtrip}`,
`clerk-backend/{FuzzCreateOrganizationBody,FuzzBulkInviteBody}`.

**Bugs fixed inline:** none — the move changed no behavior; all gates were green before and after.

**Flakes stabilized:** none — new test 3/3 clean under `-race -shuffle=on`.

**Knowledge backfill:** no KB-worthy findings. The error-string contract is already documented in
decision M2b-D4 ("the runner consumes error *strings* via `err.Error()`") and the repo's own
`knowledge/coverage-index.md`; the new test pins an already-documented invariant rather than surfacing
a new one. The repo's `knowledge/` base (authored in S2) already covers the universal-key JWT flow and
the shared/ mint↔verify split.

**Post-harden green-gate re-verify (offline, `GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off`):**
`go test -race -count=1 ./...` 7/7 ok · Go gate 100.0%/100.0% **22/22**, exit 0 · JS gate
100.0%/100.0% **9/9**, exit 0 · `drift-test.sh` **ALL PASS** (9/9), exit 0 · gofmt/vet/shellcheck clean.

### Stop condition
Stopped after Pass 1 (legitimate for a pure reorg). All stop conditions met: the full Step 2b scan
across all six dimensions found nothing further worth adding (the two sub-100% bands — webhook/clerkrun/
jsfapirun — contain only unchanged *moved* code that the anti-pattern guard forbids padding; both 100%
packages have no uncovered branches; the merged clerk-backend scan was clean); coverage delta negligible
(+0%, by design); zero flakes (3 consecutive clean `-race -shuffle` runs). The move held coverage and
the `shared/` boundary is solid.

## M2b: Final Review

Deferral re-audit (Phase 1b): **GREEN** — 2 inherited single deferrals (DEF-M0-01 `feat/demo-environment`→main owned by close-release; DEF-M2-01 live-wiring owned by M3/v1.1), 0 repeat/chronic/aged-out, 0 M2b-originated. Report: [audit-deferrals/deferral-audit-2026-06-03-m2b-close.md](audit-deferrals/deferral-audit-2026-06-03-m2b-close.md).

### Scope
- [x] All S1–S5 delivered; S4 user-action box (`/singularity-kit:repo-consolidate code`) is a documented user-finalize (M2b-D3/D8, `disable-model-invocation`), not a deferral — no scope gap.

### Code Quality
- [x] [should-fix] CQ-1: stale comment in `clerk-backend/fuzz_test.go:15` names pre-reorg packages (`authn.FuzzParse` / `fapi.FuzzParsePublishableKey`) — repoint to `shared.FuzzParse` / `clerkfrontend.FuzzParsePublishableKey`.

### Documentation
- [x] DOC-1: `knowledge/coverage-index.md` test count drifted after the harden pass added `TestGetUserErrorClassesCrossBoundary` (authn 6→7, total 112→113) — reconcile the table + total against ground truth.
- [x] DOC-2: rosetta `state.md` Headline numbers quote 112 + stale surface-named coverage (`orgclient/fapi/bapi`) — refresh to 113 + library-named coverage in Phase 8 per the state.md contract.

### Tests & Benchmarks
- [x] No gaps. Full suite 7/7 green (uncached, `-race`); gates Go 22/22, JS 9/9, drift 9/9. No benchmarks (the alignment gates are the threshold assertions). Adversarial `shared/` seam already pinned.

### Decision Triage
- [x] D1 (Go pkg naming), D2 (library-named scheme), D4 (`shared.Parse`/mint↔verify split) — already blended into the repo's `knowledge/architecture.md` + `CLAUDE.md` with `#M2b-D` tags during S2/S4. No new blend.
- [x] D3, D5, D6, D7, D8 → archive (maintainer/process-only — repo-consolidate-is-user-invoked, script depth, binary gitignore anchors, CI JS-gate step).
