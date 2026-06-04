# State

**Active version:** **v1.1 "show floor"** — IN DEVELOPMENT on `release/01.10-show-floor` (designed 2026-06-03). Disposable, richly-seeded demo stacks on demand, built entirely under the gitignored `anthropos-demo/` scratchpad (zero platform-repo changes). M3 → M4 → M5, strictly sequential.
**Active milestone:** **M4 "Declarative data seeding"** (section, `planned`) — one `demo.seed.yaml` (org size, role mix, content, activity span) backfills an M3 demo stack by orchestrating the platform's existing bootstrap/import CLIs in dependency order; structural data only (M4-D1 fold validate/dry-run). Dir: [m4-declarative-seeding/](releases/01.10-show-floor/m4-declarative-seeding/).
**Next up:** **`/developer-kit:work-milestone M4`** on `release/01.10-show-floor`. Then M5 (recipes + polish) → `/developer-kit:close-release`.
**Last shipped:** **v1.0 "body double"** — 2026-06-03 (tag `v1.0`). · **Last closed milestone:** M3 — 2026-06-03.
**Paused:** _(none)_

## Shipped releases
- **v1.0 "body double"** — **2026-06-03**, tag `v1.0`. The alignment-testing framework (`test/alignment/`) + **Clerkenstein**, a *measured* drop-in Clerk mock at **100%/100% on all three surfaces** (Go · JS/FAPI · `@clerk/express`), zero platform-code change. 6 milestones (M0→M2c). close-release caught + fixed 1 blocker (an `@clerk/express` gate regression from the M2c close). Records: [releases/archive/01.00-body-double/](releases/archive/01.00-body-double/) (review · retro · metrics · stats · lockfile).

## Recently closed
- **M3 "Disposable multi-instance demo stacks"** (2026-06-03, section) — the rosetta-demo layer: bring up `demo-N` isolated alongside the dev stack (additive compose override with a port-offset + `!override` + per-stack project/data, clone-each-repo-at-its-release-tag, the 4 Clerkenstein injection recipes wired) — **zero read-only-platform change**. Tooling in the new gitignored `anthropos-demo/rosetta-demo/` repo (M3-D4) + `/demo-*` skills + `corpus/ops/rosetta_demo.md`. 12 unit tests. **Acceptance (M3-D5): demo-1 ran isolated alongside the dev stack — up→status→down, dev untouched.** Full 12-service / 2-concurrent / browser-login proofs resource-gated → bigger Docker VM. roadmap.md § M3.
_(v1.0's milestones M0→M2c are in [roadmap.md](roadmap.md) § Done — v1.0.)_

## Headline numbers (inherited from v1.0 close — baseline for v1.1)
- **Alignment gates (v1.0 — held green):** **100%/100%** on all three surfaces — Go (22/22, `clerk@2.6.0`), JS/FAPI (9/9, `clerk-js@5`), `@clerk/express` (9/9, RS256/JWKS). Clerkenstein: 8 packages, 128 Go test/fuzz funcs, flake 0. v1.1 must keep these green (the demo stacks consume Clerkenstein, they don't change it).
- **rosetta framework:** `test/alignment/` 43 test + 3 fuzz, stdlib-only.
- **v1.1 baseline:** no new code-metric baseline yet — set as M3/M4 land their seeder + overlay tooling under `anthropos-demo/`.

## v1.1 "show floor" milestones (planned)
**M3** (done — demo stacks) → **M4** (planned — declarative seeding) → **M5** (planned — corpus + recipes + polish). Strictly sequential. Full design + execution graph + risks in [roadmap.md](roadmap.md) § In Development — v1.1.

## v1.1 decisions locked at design (2026-06-03)
- **M3-D1 — per-demo service-repo clones** (user-chosen): each `demo-N` re-clones the service repos under `anthropos-demo/stacks/demo-N/` (full isolation, ~N× disk) rather than sharing the `anthropos-dev/*` clones.
- **M3-D2 — manual teardown only** (user-chosen): `/demo-down [N]` is the only reclaim path; no nightly auto-reaper in M3.
- **M4-D1 — fold `--validate`/`--dry-run` into M4** (no separate M4b yet): split out an M4b only if that surface grows beyond a flag.
- **Carried into v1.1 (inherited hard constraint):** **zero changes to any read-only platform repo** — all demo work is additive overlay/orchestration under the gitignored `anthropos-demo/`. Clerkenstein (v1.0) is consumed, not modified.
- **Express-gate CI carry-forward** (from v1.0) lands in **M5** (a demo stack materializes `@clerk/express` in `node_modules`).

## Branch model
**v1.0 SHIPPED + pushed:** `release/01.00-body-double` merged `--no-ff` → `main`, tagged **`v1.0`** (both pushed to `origin`); release branch deleted. Clerkenstein published **private** at `anthropos-work/clerkenstein`.
**v1.1 IN DEVELOPMENT:** `release/01.10-show-floor` cut from `main` (2026-06-03); the v1.1 design + milestone scaffolds (M3/M4/M5) live on it. `main` stays at the v1.0-shipped state until v1.1 closes. Milestone branches `m{3,4,5}/{slug}` cut from the release branch by `/developer-kit:build-milestone`; merged back at each close; the release merges → `main` + tags `v1.1` at `/developer-kit:close-release`. _(The v1.1 release branch is local — not yet pushed.)_

_Last updated: 2026-06-03 (M3 closed + merged to release/01.10-show-floor; M4 declarative seeding is next)._
