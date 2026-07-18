# iter-03 — progress

**Type:** tik — under TOK-01 (recruiter-render-first). The recruiter cockpit seat — the login that unlocks the
gate measurement.

## What this iter did
Added the recruiter hero to the Meridian Talent hiring story so a presenter can log in as her and reach the
candidate-comparison surface; the iter-02 isHiring wiring is now LIVE for a real seat (BuildRoster emits her with
`org_is_hiring=true` → the org re-skins as hiring).

- **`presets/stories.seed.yaml`** — `rae-recruiter` (Rae Ramirez, Technical Recruiter, `vantage: manager` →
  `roleForHero`→admin → rides slot 1 = admin band, so the funnel correctly **skips** her; inherits
  `org:feature:insights` from the p3 admin policy via her g2 admin grant — no net-new grant).
  `jump_to: /enterprise/activity-dashboard`.
- **`blueprint/presets_test.go`** — `TestStoriesPreset`'s hiring-story assertion updated from "heroless at M223" to
  the M224 state (≥1 hero incl. a manager recruiter, each with a role).
- **`seeders/curated_pools.go`** — a new **`curatedTalent`** (recruiting / talent-acquisition) curated skill
  family + classifier clause + `curatedSkillsFor` case. **Required:** a manager hero renders a modest personal
  `/profile` (3-8 verified + a claimed tail drawn from her role's curated pool), and the shipped-preset guard
  `TestShippedPresets_EveryHeroRoleClassifies` fails on any `curatedNone` hero role. (Post-M219-R-8 `curatedNone`
  falls to the coherent general family, not junk — but the guard requires a *specific* family per shipped hero;
  a recruiting family gives Rae coherent recruiting skills.)
- **`seeders/curated_pools_test.go`** — pin `Technical Recruiter`/`Sourcer`/… → `curatedTalent`.
- **`presets/seed-generation-manifest.yaml`** — regenerated projection (recruiter now in the population) + header
  updated; `TestManifest_CanonicalFileMatchesProjection` green.

## Verification
- **Full stack-seeding module suite GREEN** (blueprint + seeders + cmd/stackseed incl. the manifest-drift test);
  `gofmt`/`go vet` clean.
- No Clerkenstein change this iter → no `/align-run` needed (the iter-02 wiring stands).
- rext commit **a29267b**, tag **`casting-call-m224-iter03`**.

## Close — 2026-07-16

**Outcome:** recruiter cockpit seat landed + preset validation moved to the M224 state + a coherent `curatedTalent`
skill family (so her modest `/profile` isn't junk/generic) + manifest regenerated. Planned scope landed fully.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (no render measured yet — this is the 2nd enabling-scaffold tik; the first render reading is
iter-04's bring-up + baseline)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (a tik; and iter-02+iter-03 are **enabling-scaffold tiks — metric-neutral BY DESIGN** (building the measurement preconditions), both `closed-fixed` on planned scope, not stalled fix-attempts, so they should not read as a no-progress stall; the first render-ATTEMPT is iter-04) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue (→ iter-04)
**Decisions:** D1 (recruiter=manager→admin, rides slot 1, funnel-skipped); D2 (curatedTalent required by the shipped-preset guard — a manager renders a /profile); D3 (recruiter seat is the gate-unlocking unit; candidate heroes deferred post-gate). See iter-03/decisions.md.
**Side-deliverables:** none.
**Routes carried forward:** iter-04 = **LOCAL demo bring-up at `casting-call-m224-iter03`** + a recruiter
**render-probe** (log in via cockpit → each of the 5 `[simId]` scoreboards → count comparable rows + score
distribution + closure/eject) + the **baseline measurement + attribution** (the first gate reading). Then the 2
candidate `/profile` heroes (funnel/candidate-role hero-awareness) as a post-gate tik. The
`/enterprise/activity-dashboard` jump_to may want a per-`[simId]` refinement once the seeded sim ids are known.
**Lessons:** A **manager hero renders a real personal `/profile`** (persona.go `trajectoryLevelBand`/…IsManager()
give a modest 3-8 verified + a claimed tail) — so a new manager's role MUST classify to a curated skill family, or
`TestShippedPresets_EveryHeroRoleClassifies` fails. Adding a hero with a novel role family = adding a curated
category, not just a yaml line. (Recorded to spec-notes for the cockpit-spec.md M224 delivery.)
