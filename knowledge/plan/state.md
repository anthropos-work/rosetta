# State

**Active version:** _(none — between releases)._ v1.1 "show floor" shipped 2026-06-05. The next version is not yet scoped — run **`/developer-kit:design-roadmap`** to draft v1.2 (the staged seed is "richer demo worlds": lift M7c's `waived` taxonomy + content via a snapshot mechanism, + AI-generated rich content; see `roadmap-vision.md`).
**Active milestone:** _(between releases)._
**Next up:** **`/developer-kit:design-roadmap`** to scope + cut the next release.
**Phase:** between releases — awaiting `/developer-kit:design-roadmap`.
**Paused:** _(none)_

## Recently shipped releases
- **v1.1 "show floor"** — **2026-06-05**, tag `v1.1`. The **platform-operations extension framework**: disposable, Clerk-free, **production-safely-seeded** demo stacks (+ dev), in a 2-repo constellation (`rosetta` + the private `rosetta-extensions` monorepo). 8 milestones (M3→M8): repo consolidation (M4), demo/dev stack tooling (M3/M5/M6), the seeding stack (M7a framework + 3-layer isolation guard · M7b data-DNA conformance/drift · M7c backdated-activity seeder fleet), and the product layer (M8 corpus recipes + presets + `/demo-seed` + the express-gate CI carry-forward). Test funcs 175→**409**; all 4 Clerkenstein alignment gates held **100%/100%** (Go 22/22 · JS 9/9 · `@clerk/express` 9/9 · deploy 7/7); flake 0; supply chain clean. close-release: scope/code/tests GREEN, docs YELLOW→fixed (5 minor findings); 0 escape-hatch beyond the 2 user-waived surfaces (taxonomy + content → v1.2). Records: [releases/archive/01.10-show-floor/](releases/archive/01.10-show-floor/) (review · retro · metrics · lockfile).
- **v1.0 "body double"** — **2026-06-03**, tag `v1.0`. The alignment-testing framework (`test/alignment/`) + **Clerkenstein**, a *measured* drop-in Clerk mock at **100%/100% on all three surfaces** (Go · JS/FAPI · `@clerk/express`), zero platform-code change. 6 milestones (M0→M2c). Records: [releases/archive/01.00-body-double/](releases/archive/01.00-body-double/).

## Headline numbers (v1.1 close — the baseline for v1.2)
- **Test funcs:** **409** total — rosetta `test/alignment` 46 (43 test + 3 fuzz, stdlib-only) · clerkenstein 218 (13 pkgs, 4 DNAs/runners) · stack-seeding 145 (8 pkgs) + the demo tooling 87 (shell/py: stack-core 4 + demo-stack 5 + stack-injection 69 + dev-stack 9). Flake 0.
- **Alignment gates (held green through v1.1):** **100%/100%** on all 4 surfaces — Go (22/22, `clerk@2.6.0`), JS/FAPI (9/9, `clerk-js@5`), `@clerk/express` (9/9, RS256/JWKS), deployment/injection (7/7, `clerk-deploy-1`, colony @ v0.34.3). 3 wired in CI; deploy is a local gate (needs colony via GH_PAT).
- **Seeding gates:** the 3-layer production-isolation guard (audit-proven zero pollution) · data-DNA `measure` 100%/critical 100% over the 8 reachable surfaces · drift detection · login→200 · full seed ~0.7s (<2min).

## Branch model
**v1.1 SHIPPED + pushed:** `release/01.10-show-floor` merged `--no-ff` → `main`, tagged **`v1.1`** (both pushed to `origin`); release branch deleted. The extensions code is the **private** `anthropos-work/rosetta-extensions` monorepo (gitignored at `anthropos-demo/`).
**v1.0 SHIPPED:** tagged `v1.0` (2026-06-03).
**Next version:** `/developer-kit:design-roadmap` cuts `release/{VV.VV}-{codename}` from `main` when v1.2 is scoped.

_Last updated: 2026-06-05 (v1.1 "show floor" shipped — close-release: review GREEN, merged → main, tagged v1.1, records archived)._
