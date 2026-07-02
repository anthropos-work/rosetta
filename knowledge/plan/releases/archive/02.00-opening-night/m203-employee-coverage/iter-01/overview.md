---
iteration: 01
iteration_type: tok
tok_flavor: bootstrap
milestone: M203
status: closed-fixed
created: 2026-07-02
---

# iter-01 — BOOTSTRAP TOK (authors TOK-01: the employee-coverage strategy)

## Type
tok (bootstrap) — the milestone's iter-01. Authors the FIRST coverage strategy. Does NOT terminate the call;
the loop continues into iter-02 (a tik under TOK-01).

## Inputs (read as contract)
- `overview.md` (this milestone) — the exit gate: every declared EMPLOYEE-vantage use case has a PASSING
  Playthrough on a COLD reset-to-seed demo, 0 false-fails over 5 consecutive reset runs. The 3 journeys:
  Skill Paths / AI Simulations (NON-voice) / Profile.
- `corpus/ops/demo/playthroughs.md` — the M202-graduated iteration protocol (declare→seed→page-object→spec→
  run --reset→reconcile→triage→re-measure); the 4-state map; NoRegressions() is the coverage gate.
- `knowledge/plan/spec-drafts/playthroughs/spec.md` §5.8 — the assertion boundary (chat/code/document sims
  playable as-is; assert at launch/completion; NON-voice; scoring is deterministic-rubric → assert range).
- `m201-manifest-corpus/manifest-draft.yaml` — the prose-intent use-case declarations (the build contract).
- The M202 foundation (fully inspected): `profile.yaml` (1 UC), `ProfilePage`, `page-object.ts`, `hero-login.ts`
  (reuses cockpit `loginAs`), `pt-world.seed.yaml` (pt-employee = Pat Ellis, verified:8 mapped:12; pt-manager;
  pt-free), `seed-worlds.yaml`, `run-playthroughs.sh --reset`, the serial config.

## Baseline (distance-to-gate)
- **Metric:** employee-vantage use cases with a PASSING Playthrough ÷ declared employee-vantage use cases
  (the milestone's coverage), gated by `NoRegressions()` + 0 false-fails over 5 reset runs.
- **Starting value:** 1 declared+passing (`profile.foundation.UC1` — the M202 proof), which is the employee
  identity smoke. The employee journeys to add: Profile (verified / growth / self-eval), Skill Paths (legacy),
  AI Simulations (chat / code / interview — NON-voice).
- demo-1 is UP (17 containers; app :13000, fapi :15400, postgres :15432) — the proof target.

## The employee-vantage use-case set (from the M201 corpus, hero = pt-employee / Pat Ellis)
1. **Profile** — the deterministic READ surfaces (lowest-risk, first):
   - `profile-skills.verified.UC1` — verified-skill Spotlight chart + claimed-vs-verified gap render (pt-employee
     already seeds verified:8/mapped:12 → ≥2 datapoints → Spotlight renders).
   - `profile-skills.growth.UC1` — verified-skill trajectory (≥2 datapoints) + gap-to-current-role (VIEW-ONLY per
     M201 flag #2 resolution).
   - `profile-skills.self-evaluation.UC1` — self-rate a skill; persists + shows (a WRITE flow → reset-to-seed).
   - (`profile.foundation.UC1` already green — the identity smoke.)
2. **Skill Paths** — `skill-paths.legacy.UC1`: open→progress chapters/steps→complete assessment (non-voice sim)→
   path complete + a skill verifies on the profile. (Academy UC = PENDING/deferred trap per M201 — OUT.)
3. **AI Simulations (NON-voice)** — assert at the launch/completion boundary (§5.8):
   - `ai-simulations.chat.UC1` — open chat sim→reply returns→completion + result.
   - `ai-simulations.interview.UC1` — reuses the chat engine (M201 finding), interview-TYPED sim.
   - `ai-simulations.code.UC1` — code sim→executes (Judge0 via Roadrunner)→graded (verify Judge0 reachability).

## Strategy authored → TOK-01
See milestone-root `decisions.md` → **TOK-01: Deterministic-read-first, then mutating flows, then integration boundary**.

## Next-tik direction (iter-02)
Land the **Profile** journeys first (deterministic, read-mostly, least new machinery): declare the
`profile-skills.verified` + `.growth` + `.self-evaluation` use cases in a `profile.yaml` extension (keep
`ptvalidate` green: unique ids, both-way, precondition-coverage), extend `ProfilePage` with the Spotlight /
trajectory / self-rate accessors under the locator discipline, build the Playthroughs, run `--reset` against
demo-1, reconcile. Start with `profile-skills.verified.UC1` (the Spotlight render — pt-employee's seed already
provides the datapoints).
