---
iter: iter-01
milestone: M51
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-06-30
---

# iter-01 — bootstrap tok (author TOK-01: the AI-readiness 3rd-org seed strategy)

## Type
tok (bootstrap) — iter-01 of M51, unconditional per build-mstone-iters Phase 0 rule 1 + the coverage-protocol
refinement (the bootstrap tok scaffolds + authors the strategy + takes the baseline framing; does NOT terminate).

## Inputs
- `overview.md` (exit gate: 200-person 3rd org, AI-readiness manager dashboard ENABLED, ~80% all-3-complete, 1
  hero STARTED + 1 hero COMPLETED, proven by the M42 manager-vantage coverage gate, 0 prod-eject escapes).
- `corpus/services/ai-readiness.md` (the M48 contract — re-verified GREEN this iter).
- `corpus/ops/demo/coverage-protocol.md` (the iteration protocol: Phase A sweep → B triage → C fix → D re-sweep
  → E close; primary metric `(failing-pages, escapes)`; gate `(0,0)` over a frontier-exhausted crawl, fresh
  demo-up).
- iter-01 seeder survey (in spec-notes.md) — the reverse-engineered rext seeding surface.

## Pre-flight KB-fidelity gate
GREEN (2026-06-30). See `../kb-fidelity-audit.md`. The strategy below is authored against verified contract.

## Baseline framing
The 3rd org does **not exist yet** — so the manager-vantage coverage gate cannot run against it (nothing to log
into). The pre-iter metric is therefore not a sweep number; it is the **build distance**: the 3rd org + its
AI-readiness funnel + the hero trio + the cockpit wiring must be SEEDED before the first authoritative sweep.
TOK-01 frames the build path; iter-02 (the first tik) lands the first seeder slice and takes the **baseline sweep**
once the org renders.

## Initial strategy (TOK-01)
**Active-cycle, signals-true, additive-to-stories seed.** Append a 3rd story to `stories.seed.yaml`; build the
net-new `ai_readiness_*` config + funnel seeder writing an **active** cycle (so the dashboard recomputes from real
`user_skill_evidences` + ended jobsim sessions — the contract-verified live path), at ~80% all-3-complete, with the
two named heroes pinned to STARTED (stage 1) and COMPLETED (stage 3). Drive it tik-by-tik through the
coverage-protocol observe→fix→re-measure loop on demo-1 (in place), manager vantage, until `(0,0)` frontier-
exhausted.

See `../decisions.md` TOK-01 for the full strategy, rationale, strategy class, distance-to-gate, and next-tik
direction.

## Phase plan (this iter)
- Phase 0b KB-fidelity gate (DONE, GREEN).
- Author TOK-01 in milestone-root decisions.md.
- Apply the audit's doc-hygiene fixes (DONE).
- Record the seeder survey in spec-notes.md (DONE).
- Close as a bootstrap tok (does not terminate; loop continues into iter-02).

## Close-no-lift acceptability
A bootstrap tok has no metric to lift — its deliverable is the strategy + the verified-contract footing. Closes
`closed-fixed` when TOK-01 + the audit fixes + spec-notes land.
