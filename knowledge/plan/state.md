# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — in development on `release/01.00-body-double`. **Last milestone of v1.0.**
**Active milestone:** **M2 "Clerkenstein — browser session + webhook coherence (JS)"** (section) — the JS/FAPI side: a fake Clerk FAPI for `@clerk/nextjs`/`clerk-js` (publishable-key + base-URL override, **fallback** to the real dev Clerk app for the browser while the backend stays mocked); a webhook injector into `app/internal/clerk/events/`; and the **fake-Clerk-API-server** that ALSO wires M1's routed-forward orgclient injection (M1-D2). **Highest technical risk in v1.0** (SDKs hard-code Clerk FAPI; no documented base-URL override — spike early).
**Next up:** M2 **build complete** (S1–S5 done; Go gate 100%/100% + JS gate 100%/100%). `/developer-kit:harden-milestone` (optional) then `/developer-kit:close-milestone` on M2, then `/developer-kit:close-release` (ships v1.0 → `main`).
**Last closed:** M1b — 2026-06-03.
**Paused:** _(none)_

## Recently closed
- **M1b "Clerk drift detection"** (2026-06-03) — automation over M0: `gate.sh` + `drift-check.sh` (exit-code contract 0/1/2/3) + a weekly CI alignment gate + a 9-assertion `drift-test.sh`, in the clerkenstein repo. Makes a Clerk bump a flagged, mechanical event. 0 close findings. roadmap.md § M1b.
- **M1 "Clerkenstein backend mirror"** (2026-06-03, closed-on-gate) — first mirror built by M0: 100%/100% alignment vs `clerk@2.6.0`, offline. roadmap.md § M1.
- **M0 "Alignment measurement framework"** (2026-06-02) — the reusable alignment framework. roadmap.md § M0.

## Headline numbers (v1.0 so far)
- **Alignment gate:** **100% overall / 100% critical** (22/22 genes) vs `clerk@2.6.0`, now **CI-gated weekly** + on push (M1b).
- **clerkenstein:** authn + orgclient **100%** unit (27 Go tests +1 fuzz) + a 9-assertion shell drift harness; flake 5/5; shellcheck clean.
- (M0 baseline: `test/alignment/` 45 funcs, core coverage 83–98%.)

## v1.0 "body double" milestones
**M0** (done) → **M1** (done) → **M1b** (done) → **M2** (active — last). Full design + execution graph + risks in [roadmap.md](roadmap.md).

## Design decisions locked at design time (2026-06-02)
- **Alignment is a first-class test class**; framework = **M0** (done); **M1** built the first mirror with it (done); **M1b** CI-gates drift (done).
- **M2 frontend:** attempt the fake Clerk FAPI server; **fall back to the real dev Clerk app for the browser** (backend fully mocked) if base-URL override proves too fragile. M2 also owns the **fake-Clerk-API-server** that wires M1's orgclient injection (M1-D2).
- **Repo layout:** the M0 framework + skills + docs live in **rosetta**; the Clerk DNA + tests + the `clerkenstein` mirror live in their own repo under gitignored `anthropos-demo/`.
- **"Zero platform-code changes"** = build-time `go.mod replace` (authn) / fake-API-server (orgclient + JS) + skip-worktree; upstream never modified.
- **Two-version split:** v1.0 = Clerkenstein (M0–M2); v1.1 "show floor" = stacks + seeding + recipes (M3–M5, [roadmap-vision.md](roadmap-vision.md)).

## Branch model
M1b closed — **merging `m1b/clerk-drift-detection` → `release/01.00-body-double` (`--no-ff`)**, then deleting the branch. Release branch cut from `feat/demo-environment` (M0-D6); `feat/demo-environment` → `main` reconciliation at release close. M2 = section (`/developer-kit:build-milestone`).

_Last updated: 2026-06-03 (M2 build complete — fake FAPI + BAPI redirect + webhook injector + JS DNA at 100%/100%; ready for close)._
