**Type:** tok (bootstrap) — per coverage-protocol.md §"Iter type selection" (the bootstrap tok authors the
strategy + scaffolds + takes the baseline framing).

# iter-01 — bootstrap tok

## What happened
1. **Phase 0b pre-flight KB-fidelity gate (mandatory for iter-01): GREEN.** Audited the AI-readiness contract
   (`corpus/services/ai-readiness.md`, the load-bearing input) against the live `stack-demo/app` clone — all 8
   behavioral claims ALIGNED, including the critical cycle-state read-path (active cycle recomputes from
   `user_skill_evidences` + jobsim sessions; `ai_readiness_live_snapshots` is a materialized cache, not the
   dashboard source). 3 doc-hygiene fixes applied inline (enum-location clarification, drift-proof symbol anchor,
   freshness bump). No blind areas. Report: `../kb-fidelity-audit.md`.
2. **Seeder-surface reverse-engineering.** Surveyed the rext seeding module — recorded the 3rd-org entrypoint
   (`stories.seed.yaml`), the reused machinery (PersonaSeeder 7-table chain, JobsimSessions, closure gate), and the
   3 net-new pieces (org_settings gate writer, sim-id pin, the 3-step funnel seeder). Recorded in `../spec-notes.md`.
3. **Authored TOK-01** (`../decisions.md`): active-cycle, signals-true, additive-to-stories seed strategy.
4. **Verified the live target:** demo-1 UP (17 containers, offset 10000, backend :18082=200).

## Close — 2026-06-30

**Outcome:** TOK-01 authored (active-cycle signals-true additive-to-stories seed strategy); KB-fidelity gate GREEN; 3 doc fixes + seeder survey landed. No metric (bootstrap tok).
**Type:** tok
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, does NOT exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (iter-local: seed-strategy choice rationale captured in milestone TOK-01)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-02 (first tik) lands strand 1 (the 3rd story + `OrgSettingsSeeder` gate row) + takes the baseline manager-vantage sweep.
**Lessons:** The active-vs-frozen seed decision hinges entirely on the contract's cycle-state read-path; a wrong read (seeding `live_snapshots` directly for an active cycle) would silently no-op (overwritten on refresh). Verifying claim 5 against code BEFORE committing the strategy was the right pre-flight order.
