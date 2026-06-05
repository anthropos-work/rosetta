# State

**Active version:** **v1.2 "set dressing"** — in development on `release/01.20-set-dressing` (designed 2026-06-05). Richer demo worlds: the **snapshot mechanism** lifts M7c's two `waived` surfaces (`taxonomy` + `content`) to **100% data-DNA coverage** — capture the real skill taxonomy + content library once, replay per-stack, measured-faithful via a new snapshot-fidelity alignment dimension. 3 milestones, all `section`: **M9** (framework + fidelity DNA + taxonomy) → **M10** (Directus content snapshot-replay) → **M11** (recipes + presets + corpus). AI-content + shareability held to v1.3 (user, 2026-06-05).
**Active milestone:** **M9 — Snapshot framework + fidelity DNA + taxonomy snapshot** (`planned` · `section` · large). Next: **`/developer-kit:build-milestone`** (creates `m9/snapshot-framework` from the release branch).
**Next up:** **M10** (Directus content → 100% coverage) → **M11** (richer-world recipes + corpus) → `/developer-kit:close-release`.
**Phase:** designed — milestone dirs scaffolded, release branch cut. Awaiting `/developer-kit:build-milestone` on M9.
**Paused:** _(none)_

## Recently shipped releases
- **v1.1 "show floor"** — **2026-06-05**, tag `v1.1`. The **platform-operations extension framework**: disposable, Clerk-free, **production-safely-seeded** demo stacks (+ dev), in a 2-repo constellation (`rosetta` + the private `rosetta-extensions` monorepo). 8 milestones (M3→M8): repo consolidation (M4), demo/dev stack tooling (M3/M5/M6), the seeding stack (M7a framework + 3-layer isolation guard · M7b data-DNA conformance/drift · M7c backdated-activity seeder fleet), and the product layer (M8 corpus recipes + presets + `/demo-seed` + the express-gate CI carry-forward). Test funcs 175→**409**; all 4 Clerkenstein alignment gates held **100%/100%** (Go 22/22 · JS 9/9 · `@clerk/express` 9/9 · deploy 7/7); flake 0; supply chain clean. close-release: scope/code/tests GREEN, docs YELLOW→fixed (5 minor findings); 0 escape-hatch beyond the 2 user-waived surfaces (taxonomy + content → v1.2). Records: [releases/archive/01.10-show-floor/](releases/archive/01.10-show-floor/) (review · retro · metrics · lockfile).
- **v1.0 "body double"** — **2026-06-03**, tag `v1.0`. The alignment-testing framework (`test/alignment/`) + **Clerkenstein**, a *measured* drop-in Clerk mock at **100%/100% on all three surfaces** (Go · JS/FAPI · `@clerk/express`), zero platform-code change. 6 milestones (M0→M2c). Records: [releases/archive/01.00-body-double/](releases/archive/01.00-body-double/).

## Headline numbers (v1.1 close — the baseline for v1.2)
- **Test funcs:** **409** total — rosetta `test/alignment` 46 (43 test + 3 fuzz, stdlib-only) · clerkenstein 218 (13 pkgs, 4 DNAs/runners) · stack-seeding 145 (8 pkgs) + the demo tooling 87 (shell/py: stack-core 4 + demo-stack 5 + stack-injection 69 + dev-stack 9). Flake 0.
- **Alignment gates (held green through v1.1):** **100%/100%** on all 4 surfaces — Go (22/22, `clerk@2.6.0`), JS/FAPI (9/9, `clerk-js@5`), `@clerk/express` (9/9, RS256/JWKS), deployment/injection (7/7, `clerk-deploy-1`, colony @ v0.34.3). 3 wired in CI; deploy is a local gate (needs colony via GH_PAT).
- **Seeding gates:** the 3-layer production-isolation guard (audit-proven zero pollution) · data-DNA `measure` 100%/critical 100% over the 8 reachable surfaces · drift detection · login→200 · full seed ~0.7s (<2min).

## Branch model
**v1.2 IN DEVELOPMENT:** `release/01.20-set-dressing` cut from `main` (2026-06-05). Milestone branches `m{N}/{slug}` (`m9/snapshot-framework`, `m10/content-snapshot`, `m11/richer-recipes`) branch from it; merge back at close-milestone; the release merges `--no-ff` → `main` + tags `v1.2` at close-release. The snapshot/seeding code lands in the **private** `anthropos-work/rosetta-extensions` monorepo (gitignored at `anthropos-demo/`); the rosetta-side milestone records + corpus updates land on the `m{N}/…` branches.
**v1.1 SHIPPED + pushed:** `release/01.10-show-floor` merged `--no-ff` → `main`, tagged **`v1.1`**; release branch deleted.
**v1.0 SHIPPED:** tagged `v1.0` (2026-06-03).

_Last updated: 2026-06-05 (v1.2 "set dressing" designed — `/developer-kit:design-roadmap`: deferral audit GREEN, KB preflight done, 3 milestones M9/M10/M11 scaffolded, release branch cut)._
