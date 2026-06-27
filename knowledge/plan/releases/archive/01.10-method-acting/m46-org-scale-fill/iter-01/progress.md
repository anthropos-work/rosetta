**Type:** tok (bootstrap)

# M46 · iter-01 — bootstrap tok

Authored the milestone's initial strategy (TOK-01) from the overview + the iteration protocol
(`coverage-protocol.md`) + the M45 engine code. Phase 0b KB-fidelity ran GREEN
([`../kb-fidelity-audit.md`](../kb-fidelity-audit.md)) before strategy authoring, so TOK-01 is built
against verified knowledge docs.

## Strategy (TOK-01, full entry in `../decisions.md`)
M46 = build-the-three-deliverables-then-prove-on-a-real-org (the M45 seam is N-invariant, so scaling is a
CODE + WORKFLOW problem). Tik order: (1) auto-fill count → (2) per-story distribution → (3) preview/dry-run
mode → (4) `--gen-batches` fence + throughput/429 verification → (5) the real ~500-member gate-proving
sweep. Fixtures-first: each code deliverable unit-proven without a key; the single real capped Azure batch
runs last.

## Baseline
The gate (M42 semantic sweep on a GENERATED org + 0 collisions + closure GREEN + cost/throughput budget)
has NOT been measured on a generated org. M45 proved the engine on a bounded N=20. demo-3 up on offset
+30000; the harness is calibrated for employee (Maya) + manager (Dan); the manager `/enterprise/members`
table is the headline org-populated surface. Zero of the three code deliverables built.

## Close — 2026-06-26

**Outcome:** TOK-01 authored (initial strategy: build the 3 deliverables fixtures-first, then prove on a
real ~500-member org via the M42 semantic sweep). No gate-metric movement (a tok doesn't move the gate).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks never trigger this exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (toks don't count toward the cap) — (6) protocol-stop: n — Outcome: continue
**Decisions:** TOK-01 (milestone-root `decisions.md`)
**Side-deliverables:** the Phase 0b KB-fidelity audit report (`../kb-fidelity-audit.md`, GREEN).
**Routes carried forward:** the overview.md `delivers`/KB-deps frontmatter cites `corpus/ops/cache-spec.md`
but the file is at `corpus/ops/demo/cache-spec.md` — fix during the doc-update phase (tracked in
`../spec-notes.md`).
**Lessons:** The three M46 gaps are independent code deliverables on top of an N-invariant seam — they can
each close `closed-fixed` on a fixtures-only unit proof, deferring the (single, capped) real-run budget to
the final gate-proving tik. This keeps the empirical risk concentrated in one heartbeat-streamed sweep.
