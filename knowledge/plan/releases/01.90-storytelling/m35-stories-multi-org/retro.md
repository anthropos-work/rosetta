# M35 — Retro

_Stories & Heroes model + multi-org — turn M34's single-hero spine into a declarative engine where one
`stack.stories.yaml` seeds MULTIPLE orgs, each with a thriving/struggling/manager hero trio at
vantage-appropriate fidelity. The second `section` milestone of v1.9 "storytelling". Two-repo (rosetta doc-half
+ planning; ext code @ tag `storytelling-m35` @ `06d872c`). Closed 2026-06-23._

## Summary

M35 made the seeder a **demo-orchestration engine**. M34 proved one hero real; M35 scales to the locked v1
roster — **2 stories × 3 heroes**: Cervato Systems {Maya / Tom / Dan} + Solvantis {Sara / Nick / Leah} — each
story one org, each trio a thriving employee + a struggling employee + a manager. The load-bearing design move
is **D-M35-1**: a single `blueprint.EffectiveStories()` normalization layer resolves *both* a multi-story
blueprint *and* a legacy single-org one into a uniform resolved-story slice, so all 8 seeders iterate **one code
path** — and the legacy view keeps **byte-identical ids** (the first story stays on the bare stack prefix +
`LegacyOrgID`), which is why the entire existing single-org preset/dev-min suite stays green unchanged. On that
seam M35 shipped: per-story `OrgID`/`orgClerkID` threaded through all 8 seeders (the first story keeps the
Clerkenstein default org — single-identity until M37, D-M35-2 — later stories get deterministic `StoryOrgID`);
the `PersonaSeeder` scaled to the roster with declaration-order **collision-free** hero slots; the **trajectory**
logic (thriving = dense/rising/under-claim vs struggling = sparse/low/over-claim); the `jobroleref.go` runtime
job-role resolver (real replayed `J-…` node-ids, role-coherent, never fabricated — D-M35-6); and
supporting-population fidelity (ramped `joined_at`, deduped real names). The **done-bar is met**:
orchestrator-verified LIVE on the `--local-content` demo-3 stack — 2 distinct orgs, 6 heroes, closure green
across both orgs, the single-org regression path unchanged.

## Incidents This Cycle

- **P2 (harden, fixed) — supporting-population name collision.** A live multi-org seed produced two extra
  members also named "Leah Donovan" (a hero's name). Fixed deterministically (`nameForIndexAvoiding` re-rolls
  off a reserved set of hero names + already-emitted names; the first roll is byte-identical so the hero-free
  legacy path is unchanged). Regression-pinned. (Harden commit `8165372`.)
- **P2 (build, fixed) — negative-modulo panic in the job-role pool.** `int(hashInt(...))` can be negative (a
  uint64 high-bit cast), and Go's `negative % len` is negative → an out-of-range slice index. Hit on the first
  multi-org test run; fixed by normalizing the index in `jobRoleRefs.at()`. (Recorded in `decisions.md` §Note.)
- **P3 (close, fixed) — `skillPool.at()` was the un-hardened sibling of the same negative-modulo class.** Its
  sole call site is loop-safe today (no live crash), but the function contract permitted a negative index while
  its parallel `jobRoleRefs.at()` already guarded it — a latent panic. Normalized + a regression test that
  panics on the pre-fix bare-modulo form. (Close commit `06d872c`.)

No regressions; the existing single-org path stayed green throughout (the byte-identical-legacy-view invariant).

## What Went Well

- **The normalization seam (D-M35-1) contained the regression risk.** The roadmap's stated worry — "the
  multi-org refactor touches 4 seeders + Clerkenstein claims, regression risk on the single-org path" — never
  materialized: one resolved-story code path + byte-identical legacy ids meant the existing preset/dev-min suite
  passed unchanged across the whole milestone.
- **The inherited deferral landed in full at its first destination.** #M34-D7 (routed Fate-3 from M34) became
  D-M35-4 with both parts shipped — and the chosen fix (declaration-order collision-free slots) was strictly
  better than the routed "warn-on-hash" idea: it makes the trio's rows *correct*, not just *visible*.
- **Coverage climbed under harden** (blueprint 90.9→100%, seeders 95.6→98.0%) with the multi-org integration
  test proving the cross-org closure gene — the exact done-bar property — at the DB level.

## What Didn't

- **A negative-modulo bug shipped to harden once and re-surfaced as its un-fixed sibling at close.** Two
  analogous pool helpers (`skillPool.at` / `jobRoleRefs.at`) drifted out of sync — the job-role one got the
  guard, the skill one didn't. The lesson: when fixing a defensive-index class, sweep *all* structurally
  identical helpers in the same pass, not just the one that crashed. (Caught at close; both now consistent.)

## Carried Forward

- **The Out-list (Fate 2, already-owned), unchanged:** the org-aggregate dashboard surfaces (`membership_skills`,
  tags, target-roles, succession, feedback) → **M36**; the Clerkenstein multi-identity seat-switch + the
  presenter cockpit (incl. the literal browser-pixels render of a hero's *individual* profile via
  login-AS-a-hero) → **M37/M38**.
- **Nothing routed-new, dropped, or escape-hatch-deferred.** Clean close.

## Metrics Delta (from metrics.json)

- **stack-seeding test funcs:** 302 → **347** (+45; 425 incl. subtests). Integration tests opt-in
  (`//go:build integration` + `STACKSEED_IT_DSN`).
- **Coverage:** blueprint 90.9 → **100.0%**; seeders 95.6 → **98.0%**.
- **Flake gate:** 5/5 sequential shuffled, **0 flakes**. `-race` green. `go vet` + `gofmt` clean.
- **Supply-chain:** GREEN (0 new deps). **Alignment:** 100%/100% (untouched — the Clerkenstein org-claim
  alignment is data-side, D-M35-2).
- **Review:** 2 findings, 0 blocking. **Deferral re-audit:** GREEN (1 inherited landed in full; 0 repeat).
