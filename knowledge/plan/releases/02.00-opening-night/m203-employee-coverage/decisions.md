# M203 Decisions

Implementation decisions with rationale (recorded during the iter loop). Design-time decisions live in
[`overview.md`](overview.md) + the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| | _(none yet — populated during `/developer-kit:build-mstone-iters`)_ | | |

## TOK-01: Deterministic-read-first, then mutating flows, then integration boundary — 2026-07-02

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Grow the employee-vantage coverage in RISK order, one journey at a time, each iter
following the M202 protocol (declare in the manifest → extend the seed only if a precondition is missing →
add the page-object for any new surface → build the Playthrough → `run-playthroughs.sh N --reset` → reconcile
with `ptreport`). The order:
  1. **Profile** (read-mostly, deterministic, least new machinery) — `profile-skills.verified.UC1` (Spotlight
     render; pt-employee's seed already carries the datapoints), then `.growth.UC1` (trajectory + gap-to-role,
     VIEW-ONLY per M201 flag #2), then `.self-evaluation.UC1` (the one WRITE → exercises reset-to-seed).
  2. **Skill Paths** — `skill-paths.legacy.UC1` (browse→progress→complete assessment sim→a skill verifies).
     Needs a legacy path w/ an ASSESSMENT-typed sim in the replayed Directus catalog (M201 seed finding).
     Academy UC is OUT (M201 PENDING/NOT-RUNNABLE trap).
  3. **AI Simulations (NON-voice)** — `ai-simulations.chat.UC1` (base engine), `.interview.UC1` (reuses the
     chat engine — M201 finding), `.code.UC1` (Judge0 via Roadrunner). Assert at the **launch/completion
     boundary** (spec §5.8): flow launched + reached interactive state + outcome artifact materialized
     (session row / computed result structure/range) — NEVER turn-by-turn AI prose.
**Rationale:** The exit gate is `NoRegressions()` over the employee set with 0 false-fails over 5 reset runs —
determinism is the hard part. Front-loading the deterministic read surfaces (Profile) buys the earliest greens
against the least new machinery (the M202 foundation already proved login + page-object + seed + one /profile
assertion). The mutating flows (self-eval, Skill Paths, sims) come after, once the read baseline is stable,
because each adds a precondition (a catalog sim, a Judge0 path) and a mutation the reset-to-seed must cover.
The integration-dependent sim legs assert only at the §5.8 boundary — the only thing provable under P6 with a
live LLM in the loop.
**Strategy class:** new-direction
**Distance-to-gate context:** metric = employee UCs passing ÷ declared; baseline 1 (the identity smoke),
~7 employee journeys to add across Profile / Skill Paths / AI Simulations. Gate = NoRegressions() at the
employee set + 0 false-fails over 5 cold reset-to-seed runs on demo-1.
**Next-tik direction:** iter-02 (tik) — the Profile journeys. Declare `profile-skills.verified.UC1` +
`.growth.UC1` + `.self-evaluation.UC1` (keep ptvalidate green), extend `ProfilePage` with the Spotlight /
trajectory / self-rate semantic accessors, build the Playthroughs, `run --reset` demo-1, reconcile. Start with
`profile-skills.verified.UC1` (Spotlight render — the datapoints are already seeded).
