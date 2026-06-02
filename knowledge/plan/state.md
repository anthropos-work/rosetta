# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — in development on `release/01.00-body-double`. First version under rosetta's planning lifecycle.
**Active milestone:** **M1 "Clerkenstein backend mirror"** (**iterative**) — next. Builds the first real mirror: authors the Clerk Alignment DNA with `/align-dna` and drives `/align-run`'s score to the exit gate (**100% critical / ≥95% overall**), injected via `go.mod replace`. Built by `/developer-kit:build-mstone-iters`. Lives in the **new `clerkenstein` repo** (cloned into the gitignored `anthropos-demo/`), consuming the M0 framework that lives in rosetta.
**Next up:** start M1 (`/developer-kit:build-mstone-iters`). Then { **M1b** drift detection ∥ **M2** browser session + webhook }.
**Last closed:** M0 — 2026-06-02.
**Paused:** _(none)_

## Recently closed
- **M0 "Alignment measurement framework"** (2026-06-02) — the reusable, engine-agnostic alignment measurement framework: `alignctl` (stdlib-only Go, offline) + `/align-dna` + `/align-run` + `corpus/architecture/alignment_testing.md` + a toy proving end-to-end detection (86.7% / 100% crit, catches `Greet/padded-name`). Build S1–S5 → harden (2 passes) → close (adversarial review found+fixed a path-traversal must-fix + score overflow, M0-D7). Full narrative in [roadmap.md](roadmap.md) § M0.

## Headline numbers (M0 baseline — the alignment framework)
- **Go test funcs:** 45 (3 fuzz) across `test/alignment/`. 0 failures; 5/5 flake gate (random order).
- **Coverage (core):** dna 98.3% · canon 94.7% · report 96.2% · outcome 92.5% · compare 89.7% · runner 83.3% · cmd/alignctl 67.5% (remainder = `main()` dispatch + trivial toy fixtures, exercised e2e).

## v1.0 "body double" milestones
**M0** (done) → **M1** (Clerkenstein backend, iterative) → { **M1b** drift detection ∥ **M2** browser + webhook }. Full design + alignment vocabulary + execution graph + risks in [roadmap.md](roadmap.md).

## Design decisions locked at design time (2026-06-02)
- **Alignment is a first-class test class** (alongside unit + integration): a (capability × variant) differential test comparing a source engine to its mirror → a 0–100% **alignment score**. Generic framework = **M0** (done); M1 builds the first mirror (Clerk) with it; M1b reuses it for drift.
- **M1 is iterative** — exit gate = the alignment score (100% critical / ≥95% overall, waived genes justified); iteration protocol = `corpus/architecture/alignment_testing.md`.
- **Record/replay golden captures** are core — Clerk's API is a live SaaS, measured reproducibly offline.
- **Two-version split:** v1.0 = Clerkenstein (de-risked, measured); v1.1 "show floor" = stacks + seeding + recipes (M3–M5, in [roadmap-vision.md](roadmap-vision.md)).
- **M2 frontend:** attempt the fake Clerk FAPI server; **fall back to the real dev Clerk app for the browser** (backend fully mocked) if base-URL override proves too fragile.
- **Repo layout:** the M0 framework (skills + DNA format + harness + doc) lives in **rosetta**; the Clerk DNA + alignment tests + the `clerkenstein` mirror live in their own repo under `anthropos-demo/`.
- **"Zero platform-code changes"** = build-time `go.mod replace` in the demo clone + skip-worktree; upstream never modified (staging's `vendor-colony/` precedent).

## Branch model
M0 closed — **merging `m0/alignment-framework` → `release/01.00-body-double` (`--no-ff`)**, then the milestone branch is deleted. The release branch was cut from `feat/demo-environment` (M0-D6); `feat/demo-environment` → `main` reconciliation happens at release close. M1 = iterative (`/developer-kit:build-mstone-iters`, Gate Outcome Ledger); M1b/M2 = section.

_Last updated: 2026-06-02 (M0 closed + merged; M1 next)._
