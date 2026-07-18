---
iter: 03
milestone: M224
iteration_type: tik
status: closed-fixed
date: 2026-07-16
---

# iter-03 — tik (the recruiter hero seat — the gate-unlocking login)

**Type:** tik — under **TOK-01** (recruiter-render-first). Protocol: `corpus/ops/demo/coverage-protocol.md`.

## Step 0 — Re-survey
Re-measured the decomposition: the **gate** (recruiter scoreboard ≥40 rows/sim × 5) needs only **(a)** a recruiter
seat to log in as + **(b)** the isHiring re-skin (iter-02, done) + **(c)** the M223 funnel seed (already writes 45
candidates × 5 sims). It does **NOT** need candidate HEROES — those are the separate `/profile` exemplar
deliverable (In-scope item 1), not the gate metric. And `roleForHero(manager)→admin` already works for the hiring
org. So the gate-unlocking unit is just **the recruiter hero seat** — small + self-contained. The 2 candidate
heroes (which need the intricate candidate-role + funnel-stage hero-awareness) route to a post-gate tik. This is a
refinement of the decomposition, same TOK-01 strategy (recruiter-render-first — front-load the gate).

## Active strategy reference
**TOK-01.** This tik makes the isHiring wiring (iter-02) LIVE for a real seat: once the recruiter hero exists,
`BuildRoster` emits her identity with `org_is_hiring=true`, so logging in as her re-skins the org as hiring — the
precondition for the scoreboard's "Results"/insights cohort treatment. She is the seat iter-04 measures from.

## Cluster / target identified
`presets/stories.seed.yaml`'s Meridian Talent 4th story has `heroes: []` (heroless at M223). Add the recruiter
(`vantage: manager` → admin → inherits `org:feature:insights` from the p3 admin policy, no net-new grant) with
`jump_to` at the comparison surface. She rides slot 1 (declaration order) = admin band — consistent across
`roleForHero`/`roleForIndex`/the funnel (which skips admin slots: recruiters don't audition).

## Hypothesis
Adding the recruiter hero makes the hiring org's cockpit seat real + the org re-skin live. Metric-neutral this iter
(no bring-up), but it is the seat the iter-04 baseline measures from.

## Expected lift
Gate metric unchanged (no render this iter). Deliverable = the recruiter seat in the preset + the preset validation
updated to the M224 state + tests green + rext re-tagged.

## Phase plan
1. `presets/stories.seed.yaml` — add the recruiter hero (Rae Ramirez, Technical Recruiter, `vantage: manager`,
   `jump_to: /enterprise/activity-dashboard`) to the hiring story.
2. `blueprint/presets_test.go` — update `TestStoriesPreset`'s hiring-story assertion from "heroless at M223" to the
   M224 state (≥1 hero incl. a manager recruiter, each with a role).
3. `go test ./blueprint/... ./seeders/...` (preset strict-parse + semantic validation + roster).
4. Commit + **re-tag** the rext authoring clone (`casting-call-m224-iter03`).

## Escalation conditions
- The preset fails semantic validation for a reason beyond the heroless assertion (a hiring-org-specific hero
  constraint) → investigate; if a platform/blueprint-design question surfaces → surface to user.

## Acceptable close-no-lift outcomes
N/A shape — closes-fixed on the seat landing + preset validation + re-tag. Gate NOT MET (no render yet; iter-04
takes the baseline). The `/enterprise/activity-dashboard` jump_to may need a per-`[simId]` refinement once the
seeded sim ids are known (iter-04) — noted, not blocking.
