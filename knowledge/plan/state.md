# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — on `release/01.00-body-double`. **ALL milestones (M0→M1→M1b→M2→M2b) closed + merged. v1.0 is ready for `/developer-kit:close-release`.**
**Active milestone:** _(between milestones — v1.0 feature-complete + consolidated; next action is release close, not a milestone)._
**Next up:** **`/developer-kit:close-release`** — release-level review + merge `release/01.00-body-double` → `main`, **plus** the deferred `feat/demo-environment` → `main` reconciliation (M0-D6). Then v1.1 "show floor" (M3–M5) promotes from [roadmap-vision.md](roadmap-vision.md). ⚠ One open **user action** (not a deferral, M2b-D3/D8): run `/singularity-kit:repo-consolidate code` at the `clerkenstein` repo to formally finalize the consolidation (the repo is already authored to its standard; the run is a verification that refreshes `CLAUDE.md`/`singularity-manifest.md`).
**Last closed:** M2b — 2026-06-03.
**Paused:** _(none)_

## Recently closed
- **M2b "Clerkenstein repo consolidation + knowledge base"** (2026-06-03) — pure-cleanup B-milestone: reorganized the `clerkenstein` repo into **library-named** dirs (`authn/ clerk-backend/ clerk-frontend/ clerk-webhook/`) + `shared/` + `alignment/` + a self-contained `knowledge/` base + `CLAUDE.md`/`singularity-manifest.md` + gitignored `.agentspace/`, via 69 history-preserving renames. No behavior change — both gates (Go 22/22, JS 9/9) + drift (9/9) stayed green. Close fixed 1 stale comment + 2 doc count drifts; deferral audit GREEN. **Last v1.0 milestone.** roadmap.md § M2b.
- **M2 "Clerkenstein — browser session + webhook coherence (JS)"** (2026-06-03) — closes the last two Clerk seams (Clerk-free end to end): fake FAPI (browser login via a minted publishable key — config-only, no SDK fork; M2-D1), fake BAPI disarming the platform's networked orgclient (M1-D2 pickup, concurrency-safe store), the svix-signed webhook injector, and a 2nd Alignment DNA (`clerk-js-5`, 9 genes) at 100%/100% — proving M0 is surface-generic. Close found+fixed a ChangeRole nil-map panic (M2-D4). roadmap.md § M2.
- **M1b "Clerk drift detection"** (2026-06-03) — automation over M0: `gate.sh` + `drift-check.sh` (exit-code contract 0/1/2/3) + a weekly CI alignment gate + a 9-assertion `drift-test.sh`, in the clerkenstein repo. Makes a Clerk bump a flagged, mechanical event. 0 close findings. roadmap.md § M1b.
- **M1 "Clerkenstein backend mirror"** (2026-06-03, closed-on-gate) — first mirror built by M0: 100%/100% alignment vs `clerk@2.6.0`, offline. roadmap.md § M1.
- **M0 "Alignment measurement framework"** (2026-06-02) — the reusable alignment framework. roadmap.md § M0.

## Headline numbers (v1.0 so far)
- **Alignment gates:** **100% overall / 100% critical** on BOTH surfaces — Go (22/22 genes, `clerk@2.6.0`) and JS/FAPI (9/9 genes, `clerk-js@5`); CI-gated weekly + on push (M1b), parameterized for both DNAs (M2).
- **clerkenstein:** 7 packages (library-named: authn/clerk-backend/clerk-frontend/clerk-webhook/shared + alignment/cmd), **113 Go test/fuzz funcs** (108 tests + 5 fuzz) + a 9-assertion shell drift harness; flake 5/5 (`-race`, shuffled); gofmt/vet/shellcheck clean. Coverage: authn/shared/clerk-frontend 100%, clerk-backend 97%, clerkrun 97%, clerk-webhook 96%, jsfapirun 94%.
- (M0 baseline: `test/alignment/` 45 funcs, core coverage 83–98%.)

## v1.0 "body double" milestones
**M0** (done) → **M1** (done) → **M1b** (done) → **M2** (done) → **M2b** (done). **All milestones closed — v1.0 ready for `/developer-kit:close-release`.** Full design + execution graph + risks in [roadmap.md](roadmap.md).

## Design decisions locked at design time (2026-06-02)
- **Alignment is a first-class test class**; framework = **M0** (done); **M1** built the first mirror with it (done); **M1b** CI-gates drift (done).
- **M2 frontend:** attempt the fake Clerk FAPI server; **fall back to the real dev Clerk app for the browser** (backend fully mocked) if base-URL override proves too fragile. M2 also owns the **fake-Clerk-API-server** that wires M1's orgclient injection (M1-D2).
- **Repo layout:** the M0 framework + skills + docs live in **rosetta**; the Clerk DNA + tests + the `clerkenstein` mirror live in their own repo under gitignored `anthropos-demo/`.
- **"Zero platform-code changes"** = build-time `go.mod replace` (authn) / fake-API-server (orgclient + JS) + skip-worktree; upstream never modified.
- **Two-version split:** v1.0 = Clerkenstein (M0–M2); v1.1 "show floor" = stacks + seeding + recipes (M3–M5, [roadmap-vision.md](roadmap-vision.md)).
- **M2b (added 2026-06-03):** before shipping, consolidate the `clerkenstein` repo — **library-named** dirs (one per mocked dependency: `colony/authn`, `clerk-sdk-go/v2`, `@clerk/clerk-js`+`@clerk/nextjs`, `svix`) + `shared/` + `alignment/` + a self-contained `knowledge/` base + `.agentspace/`, via `/singularity-kit:repo-consolidate`. Pure cleanup, gates stay green. (M2b-D2 = library-named scheme; M2b-D1 = Go package names for hyphenated dirs; M2b-D3 = repo-consolidate is user-invoked.)

## Branch model
M2 closed + merged. **M2b** builds on `m2b/clerkenstein-consolidation` (cut from `release/01.00-body-double`) → merged back at close; the `clerkenstein` repo's own reorg commits stack on its `main` (no branch model). **After M2b closes (release close):** `release/01.00-body-double` → `main` + the deferred `feat/demo-environment` → `main` reconciliation (M0-D6) — both owned by `/developer-kit:close-release`.

_Last updated: 2026-06-03 (M2b closed + merged to release/01.00-body-double — all v1.0 milestones done; next is /developer-kit:close-release)._
