# M250 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## TOK-01: arithmetic-spine-first → set-dress → distribute → render loop — 2026-07-24

**Tok type:** bootstrap (iter-01)
**Initial strategy:** decompose the 5-part gate into 3 build lanes then a serial render loop, in dependency order:

- **Lane A — arithmetic spine (ONE atomic edit, gate parts 1 + 5).** In rext `stack-seeding`, as a single
  commit that compiles + passes all fences together:
  1. `seeders/ai_readiness_config.go` — replace the name-pattern `readAIReadinessSkillPool`/`filterAISkills`
     path with the platform's exact **31 default node-ids** (D1: 19 core @1.0 + 12 enabling @0.5), written as a
     code-owned mirror of `defaults.go`. Closure-gate: each must resolve in the captured taxonomy or drop (never
     fabricate). Write `ai_readiness_sims` = the **3 real named uuids** with **DISTINCT tracks**
     (simulation/tech `634b9ffd…`, simulation/business `a4113c6b…`, interview/both `6d6cdf39…`) — ADD the `track`
     column to the INSERT (`UNIQUE(org,step_type,track)` requires distinct tracks for the sim pair). Set the
     active cycle's `participants_filter` with tech/business team tags so `resolveUserTrack` routes members.
  2. Re-derive the coupled constants: `aiReadinessCoreSkills 5→19`, `aiReadinessEnablingSkills 3→12`
     (denominator 6.5→25.0 via `aiReadinessTotalWeight`); the funnel held-count bands (lo/hi per stage), the
     started-hero count, and the "Champion 30/30" beat (full 31-repertoire still = 30/30).
  3. Re-derive EVERY hardcoded expected value in `ai_readiness_m219_test.go` + `ai_readiness_harden_test.go`
     (the 6.5→25.0 denominator, the 5/3→19/12 counts, the rounding-divergence triple, the hero score bands).
- **Lane B — Directus set-dress (net-new file, gate part 2).** A post-replay set-dress step that resolves each
  wired sim's `directus.sequences.evaluation_skills` node-ids → `public.skills.name` and UPDATEs
  `directus.simulations.skills` for the 3 sim uuids. Feeds `computeSimAssessments`' EvaluatedSkills list AND the
  tech/business heuristic label (D2). Demo-only enrichment (prod's column is genuinely NULL).
- **JOIN — evidence distribution (behind A + B, gate part 3).** Ensure the completed member's sim carries the
  validation fan-out (`validation_attempt_results` + `validation_attempt_skill_results`) + `user_skill_evidences`
  for the sim's real evaluated node-ids — REUSE the verified-skill fan-out helpers in `content_stories_write.go`
  / `population_evidence.go` / `persona.go` (read-only reuse). Fills dot ratings + skill distribution.
- **Render loop (serial, gate parts 4 + 5).** Reset-to-seed demo-2 (offset 20000) with the new seeder + the
  Directus set-dress → render `/ai-readiness` player + manager → measure the 5 gate parts per
  coverage-protocol.md (real seeded content, non-zero cardinality, 0 invented, 0 prod-ejects, closure green,
  frozen-vs-live arithmetic agrees at 31) → triage top failure → fix → re-render.

**Rationale:** the atomic-edit ordering is forced — a half-applied 8→31 change breaks compilation + every fence
at once, so Lane A must land as one commit before anything renders. Lane B has no dependency on A and is a
net-new file (safe to build in parallel-in-principle, but this session runs serially per the iter loop).
Evidence-distribution is strictly behind both (it distributes skills the set-dress NAMED and the arithmetic
SIZED). Live render is the only instrument for gate parts 4 + 5 (believability), so the render loop is last and
iterated. Cheap Go-test gate progress (fences green) comes first; the expensive live reset-to-seed is spent only
once A + B + the join have landed.

**Strategy class:** new-direction (bootstrap — no prior strategy).
**Distance-to-gate context:** gate metric = 5 discrete gate parts passing on a cold reset-to-seed live render.
Baseline 0/5 (demo seeds the invented 8 / 6.5-denominator / tracks-`both` sims). Expect Go-test-validated
progress on parts 1 + 5 (arithmetic) and part 2 (set-dress) before the first live render; parts 3 + 4 need the
render loop.
**Next-tik direction:** iter-02 (first tik under TOK-01) = **Lane A, the arithmetic-spine atomic edit** — the
config 8→31 replacement + funnel re-derivation + both fence-test files, landed as one commit that builds green
(`go build ./... && go test ./seeders/ -run AIReadiness`). Cheapest measurable gate progress (parts 1 + 5 at the
config/arithmetic level), and it unblocks Lanes B + join.

