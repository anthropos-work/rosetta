# State

**Active version:** **v1.0 "body double"** (Clerkenstein) — **designed 2026-06-02, not yet branched.** First version under rosetta's planning lifecycle.
**Active milestone:** _(none building yet)_ — **M1 "Clerkenstein backend (Go)"** is next.
**Next up:** scaffold the v1.0 milestone dirs + cut `release/01.00-body-double` from `main` (Phase 8 of design-roadmap — **deferred by user choice**), OR run `/developer-kit:build-milestone` directly (it cuts the release branch on demand). Before coding M1, read the gating docs in its KB-dependencies list (`corpus/services/clerk-integration.md`, `shared_libraries.md` §authn, `staging-clerk.md`, `webhook_setup.md`).
**Paused:** _(none)_

## v1.0 "body double" milestones (designed 2026-06-02)

Clerkenstein-first: **M1** backend auth bypass (Go authn + orgclient twins, `go.mod replace` injection, parity tests, the `clerkenstein` repo) → { **M1b** Clerk parity & drift CI harness ∥ **M2** browser session + webhook coherence (fake FAPI, real-Clerk fallback) }. Full design + execution graph + risks in [roadmap.md](roadmap.md) "In Development". 5 open decisions surfaced; 3 already decided at design (two-version split, M2 FAPI-with-fallback, Clerkenstein = own repo).

## Design decisions locked at design time (2026-06-02)
- **Two-version split:** v1.0 = Clerkenstein alone (de-risked standalone win); v1.1 "show floor" = stacks + seeding + recipes (M3–M5, in [roadmap-vision.md](roadmap-vision.md)).
- **M2 frontend:** attempt the fake Clerk FAPI server; **fall back to the real dev Clerk app for the browser** (backend stays fully mocked) if base-URL override proves too fragile.
- **Repo layout:** `clerkenstein` is its own dedicated repo, cloned into the gitignored `anthropos-demo/` (mirroring `anthropos-dev/`). Rosetta holds the corpus docs + skills + thin orchestration only.
- **"Zero platform-code changes"** = build-time `go.mod replace` in the demo clone + skip-worktree; the upstream repo is never modified (the same mechanism staging uses for its `vendor-colony/` patch).

## What exists today (the baseline rosetta to build on)

- **Corpus** (`corpus/`): architecture docs, per-service docs, ops guides (setup / run / update / webhook + the staging guides), tools registry. Clerk integration is well-documented (`corpus/services/clerk-integration.md`).
- **Dev-environment skills**: `/setup-platform`, `/setup-github`, `/start-platform`, `/test-platform`, `/update-platform`, `/update-knowledge` — build/maintain a **local development** env under the gitignored `anthropos-dev/` (which currently holds the full platform cloned: app, cms, skiller, jobsimulation, etc.).
- **Prior art the design leans on:** staging already vendors a patched `colony` via `replace` + skip-worktree (proves the Clerkenstein injection seam); `bootstrap-user/org`, `importSkills/JobRole`, `cms/jobsim` CLIs exist (reused by M4 seeding); `staging_from_dump.md` is the de-facto data path today.
- **Clerkenstein is a documentation blind area** — M1 authors `corpus/services/clerkenstein.md`.

## Branch model

Standard `/developer-kit:design-roadmap` → milestone branches (`m{N}/{slug}`) → `/developer-kit:close-milestone` → `/developer-kit:close-release`. **No release branch cut yet** — currently on `feat/demo-environment`; `release/01.00-body-double` is created when v1.0 is scaffolded (Phase 8, deferred) or on first `/developer-kit:build-milestone`.

_Last updated: 2026-06-02 (v1.0 "body double" designed; planning docs written, branch + scaffolds deferred)._
