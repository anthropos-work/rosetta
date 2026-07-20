---
iter: iter-05
milestone: M235
type: tik
iteration_type: tik
status: planned
created: 2026-07-20
active_strategy: TOK-01 (two-track) — Track A step 2 (non-simulation product sections)
---

# iter-05 — skill-path content-story section (the first non-sim product + the shared non-sim projection/seeder infra)

**Type:** tik · **Active strategy:** TOK-01 Track A step 2 (offline-buildable non-simulation product sections),
under the run-3 USER RULING resolving USER-BLOCKER-M235-02 ("build non-sim seeders, then close").

## Step 0 — re-survey (mandatory)
TOK-01 Track A step 2 named "skill-path (legacy runtime rows + `local_skill_path_session` mirror), academy,
ai-labs". Re-survey confirms it is **still untouched + still meaningful**: the M233 `contentProductRegistry()`
already carries all four products, but `content_manifest.go:playerResultPath` fail-closes on every non-simulation
`sim_type` ("M234 owns the non-simulation products"), and the fixture is simulation-only, so NO non-sim section
projects today. The three sections are the one remaining offline-buildable Track-A surface. Substituting nothing.

## Cluster / target identified
The **skill-path-legacy** content-product section. Evidence: M231 (`content-stories-routes.md`) ruled skill-path
**GO** — player route `/skill-path/<skillPathId>` (get-or-create; a seeded `skillpath.skill_path_sessions` row
for (owner, skillPathId) renders real progress instead of a fresh blank), manager route
`/enterprise/activity-dashboard/skill-paths/<skillPathId>/<userId>` reading the `local_skill_path_sessions`
MIRROR (M219/M222 trap). Both patterns already exist in the codebase (`assignments.go` writes the mirror;
`SkillpathSessionsSeeder` writes the runtime session) — reuse, don't rebuild.

## Hypothesis
A code-owned non-sim exhibit registry + a `ContentStoryNonSimSeeder` (skill-path arm: runtime session + mirror,
owned by a `content-player-<idx>` slot, pinned to a REAL public skill_path_id sourced offline from the captured
Directus snapshot) + an `appendNonSimProducts` projection arm (the `/skill-path/<id>` route + the mirror manager
route) makes the skill-path section RESOLVE in `content_products[]`, unit-proven, honesty-gate green.

## Expected lift
The Content-stories manifest gains a real **Skill Path** section (≥1 session, both CTAs), unit-proven; the
skill-path non-sim seeder produces deterministic + idempotent + audited rows. Readiness metric: 1 of 3 non-sim
sections resolving (from 0/3).

## Phase plan (per `playthroughs.md` + `coverage-protocol.md`, offline/unit arm)
- Author `seeders/content_nonsim.go`: the `nonSimExhibit` model + code-owned registry (skill-path exhibits pinned
  to real public skill_path_ids from the offline snapshot) + `appendNonSimProducts` projection + the non-sim
  owner-slot flat-index (self-consistent seeder↔projection pairing, its own index, mirrors the survives-drops
  discipline) + the skill-path route builders.
- Author the `ContentStoryNonSimSeeder` skill-path arm (writes `skillpath.skill_path_sessions` + the
  `public.local_skill_path_sessions` mirror; deterministic ids from the exhibit key; `CopyRowsIdempotent`;
  PerStackIsolated; audited; n=0 guard; honest degradation on no-host/no-owner).
- Wire `BuildContentProducts` + `ValidateContentManifest` to include non-sim products; register the seeder.
- Regenerate `presets/content-manifest.json` via `stackseed --content-export`; keep both honesty gates green.
- Unit tests: skill-path section projection (routes, seats, manager view), the seeder rows, the flat-index
  invariant, the honesty gate.

## Escalation conditions
- If a platform-source change is required to render the skill-path result → sha-pinned demopatch or escalate
  (NOT a platform edit). Not expected (M231 ruled it GO from a seeded row).
- Live landing (does the page render non-empty) is **M236** — this iter unit-proves only.

## Acceptable close-no-lift outcomes
- If skill-path proves to need a live-only calibration that can't be unit-expressed offline, close-no-lift with
  the falsification recorded and route to M236 (same shape as USER-BLOCKER-M235-02's coverage-descriptor finding).
