# Hardening Ledger — M224 The callback

## Pass 1 — 2026-07-16 — final

**Iters hardened this pass:** all milestone-touched code (`--final` cumulative sweep across iter-01 … iter-13;
rext milestone boundary `ff4e68a..309a00e`, 30 files).

**Tiks covered since prior pass:** all iters in milestone (first + only harden pass for M224).

**Scope method:** cumulative-scope final sweep. Walked the 30 touched files grouped by subsystem
(stack-seeding seeders, clerkenstein clerk-frontend, demo-stack up-injected + patches, stack-injection,
stack-verify e2e), branch-gap analysis per the task's named focus areas, in-process (no sub-sub-agents).

**Coverage delta on touched files (finder, not %-goal — untested-branch fences on previously-unfenced code):**
- `demo-stack/up-injected.sh` `build_frontend_hiring`: the iter-13 studio-url + public-website `urls.ts` chain
  onto the HIRING image was UNFENCED (0 hiring-scoped tests). The adjacent `test_demopatch_studio_url_wiring`
  matches the whole script body → satisfied by `build_frontend_next_web`, giving false confidence. NOW fenced.
- `stack-injection/gen_injected_override.py` `hiring_lines`: the two-app demo's structural core (iter-07) was
  emitted/asserted by NO injection test (every `build_lines` caller omits `platform_dir`; only the exposure
  guard enumerates its port). Its shape — fix17 `DIRECTUS_TOKEN` strip, absolute `env_file`, offset port,
  `platform_dir`/`with_ui` gates — NOW fenced.
- `stack-seeding/seeders/cockpit.go` `CockpitHero.IsHiring`: the manifest field that routes a recruiter's
  [Log in as] to `--hiring-base` (not the ejecting apps/web) had no Go-side test. NOW fenced.

**Tests added:**
- iter-13 → `demo-stack/tests/test_frontend_build.py`: 1 fence (`test_demopatch_hiring_studio_url_chain_wiring`)
  — hiring-scoped manifests + §5-bis 4-manifest fingerprint union + chain apply-order + LIFO trap-revert
  + demo-local env bake. Mutation-proven to catch a 2-arg-fingerprint revert.
- iter-07 → `stack-injection/tests/test_injection.py`: 3 fences (`test_hiring_app_emitted_only_when_platform_dir_given`,
  `test_hiring_app_is_ui_gated`, `test_hiring_app_block_shape`). Mutation-proven (relative env_file / dropped
  token strip both fail the shape fence).
- iter-03/08 → `stack-seeding/seeders/cockpit_test.go`: 1 fence (`TestBuildCockpitManifest_HiringHeroIsFlagged`)
  — hiring→true/workforce→false + omitempty wire guard.

**Bugs surfaced + fixed inline:** none — all three were missing regression fences, not defects. The underlying
code is correct (the gate is met and the live 4/4 flake gate is green); the fences pin behavior against a
re-break.

**Flakes stabilized:** none needed. All 5 new fences are deterministic (static string fences + pure Go/Python
builders — no I/O/timing/network); flake gate 3/3 consecutive clean.

**Task focus areas verified ALREADY-COVERED (no fence needed):**
- `endUserHeroRole`→candidate fork: `hiring_funnel_test.go::TestEndUserHeroRole_HiringOrgIsCandidate`.
- hiring mirror-pair funnel: 8 `hiring_funnel_test.go` tests (mirror pair, differentiated scores, funnel split,
  only-candidates, zero-skill-refs, sim-type, idempotent, candidate-hero stage-pinning + assignments).
- assignments `job_simulation` enum: `contentref_test.go::TestAssignmentLinkage_MatchesResourceType`
  (`default → Fatalf` catches a `"simulation"` revert) + `hiring_funnel_test.go` hero-assignment fence.
- Clerkenstein `isHiring` wiring: `resources_test.go::TestDemoUser_orgMemberships_IsHiringConditional`
  (both branches + the alignment byte-identical `{eid}` shape guard) + `registry_test.go::TestRoster_ThreadsOrgIsHiring`.
- `curatedTalent` family + role keyword map: `curated_pools_test.go` (5 recruiter roles → `curatedTalent`).
- roster `OrgIsHiring`: `roster_test.go::TestBuildRoster_HiringStorySetsOrgIsHiring`.
- cockpit.py hiring-base routing: `test_cockpit.py::test_hiring_hero_login_targets_hiring_base_others_stay_on_app_base`.
- probe R1–R5 (`render-hiring-comparison.spec.ts`, `run-hiring-render.sh`): Playwright e2e requiring a live
  demo stack — not unit-fenceable in a harden pass; exercised live by the gate's 4/4 cold-run flake gate.

**Knowledge backfill:** none authored here. The one doc fold-in the sweep reinforces — `demopatch-spec.md`
§4/§5 noting the hiring image now bakes 4 patches (2 net-new hiring + 2 chained shared `urls.ts`) — is already
tracked as close-milestone deferral (a) in `progress.md`; left in-lane there rather than double-authored.

**Verification (Phase 5):** all touched-stack suites GREEN — stack-seeding `go test ./...` (all pkgs),
clerkenstein `go test ./...`, stack-injection 263 python OK, demo-stack 658 python (8 pre-existing failures
ONLY — 4 academy-CTA + 2 overlay-JS + test_purge + test_reap, all HEAD-identical, in files this pass did NOT
touch), stack-verify `tsc --noEmit` clean. Flake gate 3/3.

**Stop condition:** stabilized — passes 1–3 each landed a fence; the pass-4 dimension scan found no new
untested branch (remaining surface is already-covered or live-e2e-only). Cumulative sweep clean.
