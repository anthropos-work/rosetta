# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — in development on `release/01.00-body-double`.
**Active milestone:** **M1b and M2 (parallel — both unblocked by M1)**. **M1b** "Clerk drift detection" (section) reuses M0's framework to CI-gate alignment across Clerk version bumps; **M2** "browser session + webhook" (section) does the JS/FAPI side + the **fake-Clerk-API-server** (which also wires M1's orgclient injection — M1-D2). They're disjoint (CI/automation vs JS) → can run concurrently; pick either next.
**Next up:** `/developer-kit:build-milestone` on M1b or M2 (both section-shaped). Then `/developer-kit:close-release`.
**Last closed:** M1 — 2026-06-03.
**Paused:** _(none)_

## Recently closed
- **M1 "Clerkenstein backend mirror"** (2026-06-03, closed-on-gate) — the **first mirror built by the M0 framework**: the Clerkenstein backend (authn `colony/authn.Provider` drop-in + disarmed in-memory orgclient, in the gitignored `anthropos-demo/clerkenstein` repo) scores **100% / 100% critical** vs the `clerk@2.6.0` DNA, offline. 5 iters (1 tok + 4 tiks), final harden 0→100% unit coverage. M1-D2: orgclient injection (fake-API-server) → M2. Delivered `corpus/services/clerkenstein.md`. Full narrative in [roadmap.md](roadmap.md) § M1.
- **M0 "Alignment measurement framework"** (2026-06-02) — the reusable alignment measurement framework (`alignctl` + `/align-dna` + `/align-run` + `corpus/architecture/alignment_testing.md`). Full narrative in [roadmap.md](roadmap.md) § M0.

## Headline numbers (M1 — the Clerkenstein backend mirror)
- **Alignment gate:** **100% overall / 100% critical** (22/22 genes) vs `clerk@2.6.0`, `alignctl run` exit 0.
- **clerkenstein Go tests:** 27 (+1 fuzz), authn + orgclient **100%** unit coverage; flake gate 5/5. `cmd/clerkrun` integration-covered by the alignment run.
- (M0 baseline retained: `test/alignment/` 45 funcs, core coverage 83–98%.)

## v1.0 "body double" milestones
**M0** (done) → **M1** (done) → { **M1b** drift detection ∥ **M2** browser + webhook + fake-API-server }. Full design + execution graph + risks in [roadmap.md](roadmap.md).

## Design decisions locked at design time (2026-06-02)
- **Alignment is a first-class test class**; the generic framework = **M0** (done); **M1** built the first mirror (Clerk) with it (done); **M1b** reuses it for drift.
- **M2 frontend:** attempt the fake Clerk FAPI server; fall back to the real dev Clerk app for the browser (backend fully mocked). M2 also owns the **fake-Clerk-API-server** that wires M1's orgclient injection (M1-D2).
- **Repo layout:** the M0 framework + skills + docs live in **rosetta**; the Clerk DNA + tests + the `clerkenstein` mirror live in their own repo under gitignored `anthropos-demo/`.
- **"Zero platform-code changes"** = build-time `go.mod replace` (authn) / fake-API-server (orgclient + JS) + skip-worktree; upstream never modified.
- **Two-version split:** v1.0 = Clerkenstein (M0–M2); v1.1 "show floor" = stacks + seeding + recipes (M3–M5, in [roadmap-vision.md](roadmap-vision.md)).

## Branch model
M1 closed — **merging `m1/clerkenstein-backend` → `release/01.00-body-double` (`--no-ff`)**, then the branch is deleted. The release branch was cut from `feat/demo-environment` (M0-D6); `feat/demo-environment` → `main` reconciliation at release close. M1b/M2 = section (`/developer-kit:build-milestone`).

_Last updated: 2026-06-03 (M1 closed-on-gate + merged; M1b ∥ M2 next)._
