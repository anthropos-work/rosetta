# M2b — progress

**Milestone:** M2b — Clerkenstein repo consolidation + knowledge base · **Shape:** section
**Status:** planned (scaffolded 2026-06-03; not yet built)
**Green-gate invariant:** after every code/script section → Go gate 22/22, JS gate 9/9, drift-test 9/9, `-race` clean.

## Sections

### S1 — Restructure into library-named dirs
- [ ] Create `shared/` and move the universal-key JWT + `Claims` + mint/parse out of `authn/` (package `shared`)
- [ ] `authn/` keeps provider + user, imports `shared` (still mocks `colony/authn`; `var _ authn.Provider` holds)
- [ ] Merge `bapi/` + `orgclient/` → `clerk-backend/` (`package clerkbackend`)
- [ ] `fapi/` → `clerk-frontend/` (`package clerkfrontend`), imports `shared`
- [ ] `webhook/` → `clerk-webhook/` (`package clerkwebhook`)
- [ ] Move runners + assets → `alignment/{cmd,dna,golden,golden-js,scripts}/`
- [ ] Repoint all imports + `gate.sh`/`drift-check.sh`/`drift-test.sh` defaults + CI YAML (ALIGN_DIR depth)
- [ ] **Gates green** (Go 22/22, JS 9/9, drift-test 9/9, `-race` clean, gofmt/vet/shellcheck)
- [ ] Tests co-located within each library dir (no stragglers)

### S2 — Knowledge base + per-library READMEs
- [ ] `knowledge/{kb-index,scope,architecture,alignment,injection,coverage-index}.md` (real content, not skeletons)
- [ ] Per-library `README.md`: `authn/ clerk-backend/ clerk-frontend/ clerk-webhook/ shared/ alignment/`
- [ ] Slim root `README.md` → points to `knowledge/kb-index.md`
- [ ] Cross-links resolve; injection recipes label each surface (recipe-only / spike-proven / built+gated)

### S3 — Hygiene
- [ ] `.agentspace/` dir created, contents gitignored (rule already present — verify + add a `.gitkeep` or README)
- [ ] `.gitignore` cleanup (fix the mismatched comment; confirm baseline Go patterns)
- [ ] Asset-hygiene: no tracked secrets/transient files; built binaries gitignored (already are)

### S4 — Consolidate (⚠ user-invoked)
- [ ] User runs `/singularity-kit:repo-consolidate code` pointed at the clerkenstein repo
- [ ] Apply its fixes; it emits `CLAUDE.md` + `singularity-manifest.md`
- [ ] **Gates + drift re-verified green** post-consolidate

### S5 — Rosetta side
- [ ] Slim `corpus/services/clerkenstein.md` → pointer at the repo's `knowledge/` + the new library-named structure
- [ ] Milestone records consistent (overview/decisions reflect what shipped)

## Running ledger
_(build appends per-section one-liners here.)_
