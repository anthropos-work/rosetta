---
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
gate: N/A for tok
---

# iter-01 — bootstrap tok (author TOK-01)

**Type:** tok (bootstrap) — iter-01 of M250, unconditional per build-mstone-iters Phase 0 rule 1.
**Job:** author the FIRST strategy the tik batch follows. No prior iters; no stalled strategy to revise.

## Inputs consumed
- `overview.md` (5-part exit gate), `roadmap.md` §M250, `spec-notes.md`.
- Iteration protocol: `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/verification.md`.
- **Platform contract (authority):** `stack-demo/app/internal/aireadiness/{defaults.go, provision.go, how_we_measure.go}`,
  `internal/data/ent/schema/{ai_readiness_sim.go, ai_readiness_skill.go}`.
- **Current demo seeders:** rext `stack-seeding/seeders/{ai_readiness_config.go, ai_readiness_funnel.go}`
  + the M219 fences `{ai_readiness_m219_test.go, ai_readiness_harden_test.go}`.

## Pre-flight
- Phase 0b KB-fidelity: **GREEN** (recorded in `spec-notes.md`; built against platform source directly).
- Phase 0d tooling check: N/A for the bootstrap tok (no pipeline wired this iter; the arithmetic-spine tik
  runs its own Go-test pre-flight).

## Distance-to-gate (baseline)
The demo currently seeds the **invented 8** (5 core + 3 enabling, denominator 6.5) and 2 tracks-`both` sims with
snapshot-resolved refs. All 5 gate parts fail against the platform's true 31-skill / 3-named-sim / evaluated-skills
contract. Baseline = 0/5.

## Output
TOK-01 (recorded in milestone-root `decisions.md`). The bootstrap tok does NOT terminate the call — the loop
continues into iter-02 (first tik) under TOK-01.
