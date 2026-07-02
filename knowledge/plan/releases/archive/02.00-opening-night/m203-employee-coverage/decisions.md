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

## D-CLOSE-1: The 4 non-gate employee-coverage UCs → Fate-3 to M206 (AI-sim tier) — 2026-07-02

**Context (close-milestone Phase 1 + 1b re-audit).** The M203 exit gate enumerated + proved the 3 CORE employee
journeys GREEN (Profile: identity+verified+growth+timeline · Skill Paths: legacy learn&progress · AI Simulations:
chat-launch §5.8), 6/6 Playthroughs, 5/5 cold-reset deterministic. Four ADDITIONAL edge use cases were routed
forward during the iter loop (iter-04/05/06 "routes carried forward") — legitimately beyond the gate:
  1. `ai-simulations.code.UC1` — Judge0-via-Roadrunner code sim (Judge0 = external hardcoded host `JUDGE0_BASE_URL`, a live seed/stack precondition per M201 draft line 440).
  2. `ai-simulations.interview.UC1` — text/non-voice interview (reuses the chat engine; needs an interview-typed catalog sim).
  3. Skill-Paths verify-skill end-to-end TERMINAL — compose learn→complete→verify with a NON-voice ASSESSMENT; the verify OUTCOME is already proven on the profile side (`pt-profile-verified`).
  4. `profile.self-evaluation.UC1` — the Profile skill self-rate WRITE (rate-modal click-intercept quirk; needs live browser iteration).

**Fate-1 (land now, complete) — infeasible for all four.** Each is a BROWSER Playthrough requiring a live demo
stack + a browser drive; this close is docs-only with NO coverage sweep / no browser re-run (the gate already
passed green + 5/5 deterministic). #1 additionally needs a live Judge0 host; #4 needs live click-intercept
iteration. A docs-only close cannot Fate-1 land a browser-driven UC — this is a genuine scope boundary, not
"too much to do now."

**Fate chosen: Fate-3 → M206 (roadmap-vision.md, future v2 major-backlog milestone).** M204 is manager-vantage
only (wrong home for employee sims/profile). M206 already owns the AI-simulation deepening domain (the voice/
recording mirror tier); its roadmap-vision entry was ANNOTATED to explicitly absorb these non-voice employee
coverage-deepening legs. This is a `roadmap-vision.md` edit (FUTURE-major backlog), NOT an in-release sibling
`overview.md` — it does not mutate the scope of any live v2.0 milestone (M204 stays manager-only).

**Roadmap mutation surfaced (per close directive):** `knowledge/plan/roadmap-vision.md` M206 entry EDITED to
absorb the 4 legs (see the M206 bullet). The orchestrator should note this future-backlog annotation. None are
escape-hatch (they're future-major work already; roadmap-vision is their natural home). Confirmed NOT gate-scope
— the gate enumerated the 3 core journeys, all GREEN.

**Academy skill-path UC — CONFIRMED OUT (not a defer).** Per M201 (PENDING/NOT-RUNNABLE trap) — ant-academy is a
separate Vercel deployment (M207 vision surface), never in M203 scope. Voice sims → M206 mirror tier. Neither is
a deferral; both were out-of-scope by design.
