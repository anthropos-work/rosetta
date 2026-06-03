# M2b — progress

**Milestone:** M2b — Clerkenstein repo consolidation + knowledge base · **Shape:** section
**Status:** built (all S1–S5 sections complete 2026-06-03; awaiting `/developer-kit:close-milestone`)
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
