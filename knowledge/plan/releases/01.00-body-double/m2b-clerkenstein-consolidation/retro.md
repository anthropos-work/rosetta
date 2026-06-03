# M2b — Retro: Clerkenstein repo consolidation + knowledge base

## Summary
A pure-cleanup B-milestone over the M2-complete `clerkenstein` repo: reorganized the flat package dirs
into a **library-named** structure (one dir per mocked dependency — `authn/` colony-authn, `clerk-backend/`
clerk-sdk-go-v2 [bapi+orgclient merged], `clerk-frontend/` clerk-js+nextjs, `clerk-webhook/` svix) + a
`shared/` dir (the universal-key JWT, extracted because the mint side and verify side are different mocks)
+ an `alignment/` harness dir, via **69 history-preserving `git-mv` renames**. Added a self-contained
`knowledge/` base (6 docs + kb-index) + 6 per-library READMEs + `CLAUDE.md` + `singularity-manifest.md`,
authored TO the `/singularity-kit:repo-consolidate` standard. Slimmed rosetta's `corpus/services/clerkenstein.md`
197→62 lines to a pointer. **No behavior change** — both alignment gates (Go 22/22, JS 9/9) + the drift
harness (9/9) stayed green throughout. 5 sections + 1 harden pass + close review.

## Incidents this cycle
- **None** (0 bugs, 0 flakes, 0 regressions). The green-gate invariant (gates + drift re-run after every
  code/script section) was the safety net for the whole reorg and never tripped after a move was repointed
  correctly. The one genuinely-new surface — the `shared/` extraction boundary (mint in `clerk-frontend`,
  verify in `authn`, joined by `shared.Parse`) — required exporting `parse`→`shared.Parse` (M2b-D4) and was
  pinned by a harden-pass cross-boundary regression (`authn/TestGetUserErrorClassesCrossBoundary`).

## What went well
- **A pure reorg with a behavioral safety net is low-risk by construction.** 69 renames + import/script
  repointing, and the only thing that could break (a wrong path) is exactly what the green-gate invariant
  catches instantly — so "fix the path, never the mock" held throughout.
- **The build did its own decision-triage.** S2/S4 flowed the consolidation-worthy decisions (D1 Go package
  naming, D2 library-named scheme, D4 shared/ split) straight into the repo's own `knowledge/` base with
  `#M2b-D` back-refs, so close-review's Phase 5 had nothing left to blend.
- **The merge of bapi+orgclient into `clerk-backend/` exposed no gap** — coverage held at 97.4%, and the
  merge actually *removed* a cross-package import (the server now references the in-package `Store`).

## What didn't
- **Two doc counts drifted from ground truth** — the harden pass added one test but didn't bump the repo's
  `coverage-index.md` (112→113) or rosetta's `state.md` Headline. Caught + reconciled at close (the Phase-4
  handbook-count reconciliation check exists precisely for this). A reminder that a count quoted in prose
  ages the moment a test lands; the test-runner output is the only authority.
- **One stale comment survived the reorg** — `clerk-backend/fuzz_test.go` named two pre-reorg packages
  (`authn.FuzzParse` / `fapi.FuzzParsePublishableKey`) in a doc comment. Compiles fine (it's a comment), so
  no gate caught it; only a cross-cutting grep for old identifiers did. Fixed at close.

## Carried forward
- **None at the milestone level.** All findings landed Fate 1 in M2b; deferral audit GREEN (0 repeat/chronic).
- **Open USER action (not a deferral, M2b-D3/D8):** run `/singularity-kit:repo-consolidate code` at the
  `clerkenstein` repo to formally finalize the consolidation. The skill is `disable-model-invocation`, so the
  build authored the repo TO its standard (self-audit compliant) and the formal run is a user-driven
  verification that refreshes `CLAUDE.md`/`singularity-manifest.md`.
- **Release-scoped (not M2b):** the `feat/demo-environment` → `main` reconciliation (M0-D6) is owned by
  `/developer-kit:close-release`, the immediate next step now that v1.0's last milestone is closed.

## Metrics delta
clerkenstein: **113 Go test/fuzz funcs** (108 tests + 5 fuzz; +1 vs M2's 112 — the cross-boundary
regression) across 7 packages; coverage authn/shared/clerk-frontend 100%, clerk-backend 97%, clerkrun 97%,
clerk-webhook 96%, jsfapirun 94% (held vs M2 baseline, +0% by design); flake 5/5 (`-race`, shuffled,
uncached); gofmt/vet/shellcheck clean. Alignment: **both** DNAs at 100% overall / 100% critical (Go 22/22,
JS 9/9); drift-test 9/9. Full figures: [metrics.json](metrics.json).
