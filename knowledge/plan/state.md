# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — **designed 2026-06-02 (refined same day to add M0), not yet branched.** First version under rosetta's planning lifecycle.
**Active milestone:** _(none building yet)_ — **M0 "Alignment measurement framework"** is next.
**Next up:** scaffold the v1.0 milestone dirs + cut `release/01.00-body-double` from `main` (Phase 8 of design-roadmap — **deferred by user choice**), OR run the build skills directly (they cut the release branch on demand). Start at **M0** (section → `/developer-kit:build-milestone`); M1 is iterative → `/developer-kit:build-mstone-iters`.
**Paused:** _(none)_

## v1.0 "body double" milestones (designed 2026-06-02)

Alignment-framework-first: **M0** the reusable alignment measurement process (the `/align-dna` build skill + `/align-run` measure skill + the alignment **test class** + DNA manifest format + equivalence operators + record/replay golden capture + a toy reference) → **M1** Clerkenstein backend mirror (**iterative**: author the Clerk Alignment DNA, drive `/align-run`'s score to the gate, inject via `go.mod replace`) → { **M1b** Clerk drift detection (reuse M0 to CI-gate alignment across version bumps) ∥ **M2** browser session + webhook (fake FAPI, real-Clerk fallback) }. Full design + alignment vocabulary + execution graph + risks in [roadmap.md](roadmap.md) "In Development".

## Design decisions locked at design time (2026-06-02)
- **Alignment is a first-class test class** (alongside unit + integration): a (capability × variant) differential test comparing a source engine to its mirror, aggregated into a 0–100% **alignment score**. The generic framework is **M0**; M1 builds the first mirror (Clerk) with it; M1b reuses it for drift. **Both alignment skills live in M0** (not split).
- **M1 is iterative** — its exit gate is the alignment score (100% critical / ≥95% overall, waived genes justified); the iteration protocol is the M0-delivered `corpus/architecture/alignment_testing.md`.
- **Record/replay golden captures** are a core M0 requirement — Clerk's API is a live SaaS and must be measured reproducibly offline.
- **Two-version split:** v1.0 = Clerkenstein (de-risked, measured) ; v1.1 "show floor" = stacks + seeding + recipes (M3–M5, in [roadmap-vision.md](roadmap-vision.md)).
- **M2 frontend:** attempt the fake Clerk FAPI server; **fall back to the real dev Clerk app for the browser** (backend stays fully mocked) if base-URL override proves too fragile.
- **Repo layout:** the M0 framework (skills + DNA format + doc) lives in **rosetta**; the Clerk DNA + alignment tests + the `clerkenstein` mirror live in their own repo, cloned into the gitignored `anthropos-demo/` (mirroring `anthropos-dev/`).
- **"Zero platform-code changes"** = build-time `go.mod replace` in the demo clone + skip-worktree; upstream repo never modified (the same mechanism staging uses for its `vendor-colony/` patch).

## What exists today (the baseline rosetta to build on)

- **Corpus** (`corpus/`): architecture docs, per-service docs, ops guides (setup / run / update / webhook + the staging guides), tools registry. Clerk integration is well-documented (`corpus/services/clerk-integration.md`). **No "alignment" / differential-testing vocabulary exists** — M0 authors it (`corpus/architecture/alignment_testing.md`).
- **Dev-environment skills**: `/setup-platform`, `/setup-github`, `/start-platform`, `/test-platform`, `/update-platform`, `/update-knowledge` — build/maintain a **local development** env under the gitignored `anthropos-dev/` (which currently holds the full platform cloned). The two new alignment skills (`/align-dna`, `/align-run`) join this set.
- **Prior art the design leans on:** staging already vendors a patched `colony` via `replace` + skip-worktree (proves the Clerkenstein injection seam); colony ships a `dummy` authn provider for tests; `bootstrap-user/org`, `importSkills/JobRole`, `cms/jobsim` CLIs exist (reused by M4 seeding); `staging_from_dump.md` is the de-facto data path today.

## Branch model

Standard developer-kit flow. **No release branch cut yet** — currently on `feat/demo-environment`; `release/01.00-body-double` is created when v1.0 is scaffolded (Phase 8, deferred) or on the first build-skill run. M0/M1b/M2 = section (`/developer-kit:build-milestone`); **M1 = iterative** (`/developer-kit:build-mstone-iters`, closes on a Gate Outcome Ledger).

_Last updated: 2026-06-02 (v1.0 refined: added M0 alignment framework, M1 → iterative; planning docs written, branch + scaffolds deferred)._
