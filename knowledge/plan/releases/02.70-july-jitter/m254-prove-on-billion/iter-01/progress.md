**Type:** tok (bootstrap) — authors TOK-01, the initial cluster-per-tik strategy for the live re-prove.

# M254 · iter-01 — progress

## Work
- Loaded milestone + protocol context (overview a–h gate, roadmap M254 LANE decomposition, the 4 protocol
  docs, the prove-on-billion playbook memory).
- Measured baseline: rext pin = `july-jitter-m253-studio-first-paint` (`b8969c0`, cumulative, on origin);
  billion clean slate (docker ps empty); gate 0/8 confirmed live.
- Authored **TOK-01** (milestone-root `decisions.md`) — the cluster-per-tik strategy.
- Recorded Phase 0b KB-fidelity verdict = GREEN-by-inheritance (`spec-notes.md`) with rationale.

## Close — 2026-07-24

**Outcome:** TOK-01 authored (initial strategy: one gate-cluster per tik, DRIVE → read-only sweeps fan-out →
mutating serial tail); baseline measured (pin on origin, billion clean slate, gate 0/8).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap does NOT exit) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (toks don't count) — (6) protocol-stop: n — Outcome: continue
(loop into iter-02, first tik under TOK-01)
**Decisions:** TOK-01 (milestone-root decisions.md); Phase 0b GREEN-by-inheritance (spec-notes.md).
**Side-deliverables:** none.
**Routes carried forward:** none (bootstrap authors the plan; the DRIVE begins iter-02).
**Lessons:** the billion demo consumes rext at the CUMULATIVE code-of-record tag (highest, all prior tags its
ancestors) — verify ancestry + on-origin BEFORE the bring-up, not the highest fetched tag (M244 iter-25
version-skew lesson).
