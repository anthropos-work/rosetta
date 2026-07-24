---
iteration_type: tik
status: closed-fixed
gate: NOT MET
---

# iter-02 — tik: Lane A, the arithmetic-spine atomic edit

**Active strategy:** TOK-01 (arithmetic-spine-first). This is the first tik.
**Step 0 re-survey:** target still current — the demo seeds the invented-8 / 6.5-denominator; nothing has
absorbed it. Baseline confirmed: `go build ./... && go test ./seeders/ -run AIReadiness` GREEN at the old numbers.

## Cluster / target
The 8→31 arithmetic spine, as ONE atomic commit that compiles + passes every fence together:
- **`ai_readiness_config.go`** — replace the name-pattern `readAIReadinessSkillPool`+`take` skill derivation
  with a shared package-level `aiReadinessDefaultSkills` (the platform's exact 31 node-ids, 19 core @1.0 +
  12 enabling @0.5, from `defaults.go`), written **unconditionally** per opt-in org (mirrors `provision.go`;
  closure-safe — platform's own real node-ids). Replace the 2-sim `aiReadinessSimRefs` write with a shared
  `aiReadinessDefaultSims` = the **3 pinned track-keyed uuids** (simulation/tech `634b9ffd…`,
  simulation/business `a4113c6b…`, interview/both `6d6cdf39…`), writing the **`track`** column + keying the
  row id by (org, step_type, **track**) so the tech+business pair doesn't collide.
- **Constants** — `aiReadinessCoreSkills 5→19`, `aiReadinessEnablingSkills 3→12` (denom 6.5→25.0 via
  `aiReadinessTotalWeight`); funnel `aiReadinessStartedHeroSkills 3→9`.
- **`ai_readiness_funnel.go`** — build the funnel's `aiSkills` pool from the same `aiReadinessDefaultSkills`
  (deterministic 31; config + funnel provably agree) instead of `readAIReadinessSkillPool`; source step-2/3
  sim refs from the pinned defaults (tech ref for step-2, interview for step-3).
- **Fences** — re-derive every hardcoded value in `ai_readiness_m219_test.go` (6.5→25.0, loop 13→50,
  "14"→"51", the MirrorsComputeTier1 cases, Aria/Ben, and the now-unreachable double-round divergence sub-test →
  a live single==double invariant) + `ai_readiness_config_test.go` (`isAISkillNodePresent`→the 31; sims 2→3
  + distinct tracks) + `ai_readiness_harden_test.go` (the NoTaxonomy comment).

## Hypothesis
Replacing the invented-8 with the platform's 31 + pinned 3 sims + re-derived arithmetic moves gate parts 1 + 5
at the config/Go-test level (the fences ARE the frozen-vs-live arithmetic check) and unblocks Lanes B + join.

## Expected lift
Go-test-validated: config writes 31 real default node-ids + 3 track-keyed sims; all fences green at 25.0. Gate
parts 1 + 5 provable at the arithmetic level (live render confirms at iter-05).

## Phase plan
Edit → `go build ./...` → `go test ./seeders/ -run 'AIReadiness|Step1|Tier1|M219'` + full `go test ./seeders/`.

## Escalation / close-no-lift
If a default node-id fails closure at live seed → triage in the render loop (route forward). If the started-hero
band can't be re-derived believably → note + carry. Compile/fence break that can't resolve in-iter → user-blocker.

## Routes carried forward (decided in-plan)
- `participants_filter` tech/business track-tagging → the render tik (the seeder comment warns tag-scoping hides
  the member surface; confirm `resolveUserTrack` × `userInActiveCycleAudience` at live render before setting it).
- Business-sim step-2 sessions (track-routed per-member sourcing) → render tik (iter-02 keeps single tech-sim
  sourcing so both cards exist; the business card's session count is a render-loop refinement).
