# M8 — progress (section checklist)

**Milestone:** M8 — Corpus + use-case recipes + polish · **Shape:** section · **Status:** done

## S1 — demo-env corpus family ✅
- [x] `corpus/ops/demo/README.md` (Purpose + When-to-Use + the up→seed→use→down flow + index) — a third indexed family
- [x] cross-linked into `corpus/README.md` (new "Demo environments" Ops subsection) + root `README.md` + `CLAUDE.md` doc-locations

## S2 — use-case recipes ✅
- [x] 3 end-to-end recipes: `recipe-enterprise-onboarding.md`, `recipe-skill-progression.md`, `recipe-browser-login.md`
- [x] the **2 M3-deferred injection recipes** (Fate-3) landed in `recipe-browser-login.md`: the `api.clerk.com`
  cert-redirect (backend orgclient → fake BAPI) + the minted-publishable-key browser-login walk-through

## S3 — curated + validated seed presets ✅
- [x] `small-200` / `mid-500` / `large-1k` in `rosetta-extensions/stack-seeding/presets/` — all `--validate` clean
- [x] **validated to actually seed:** `mid-500` (Globex, 501 users, ~4800 rows, 0.7s) + `large-1k` (~9500 rows, 0.76s) both seeded demo-1 end-to-end, isolation clean

## S4 — skill polish + discoverability ✅
- [x] `/demo-seed` skill (wraps `stackseed`; presets or `stack.seed.yaml`) + all 4 demo skills in the `CLAUDE.md` Available Skills table
- [x] **the v1.0 express-gate CI carry-forward** wired into clerkenstein `alignment.yml` (setup-node + `npm install @clerk/express` + the gate) — **VALIDATED locally 9/9** (the 4th surface, deploy, stays a local gate — needs colony via GH_PAT)

## S5 — release-boundary reconciliation ✅
- [x] the demo family indexes the M3 lifecycle guide (`rosetta_demo.md`) + the M7 seeding spec (`seeding-spec.md`) + the data-DNA, consistently
- [x] parent docs (corpus/README, root README, CLAUDE.md) all point at the demo-env story
- [x] **all 8 docs' `.md` links resolve** (verified)

**Status:** all `In:` deliverables landed. M8 is the LAST v1.1 milestone → `/developer-kit:close-release` ships v1.1.
