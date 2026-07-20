**Type:** tok (bootstrap) — protocol `playthroughs.md` + `coverage-protocol.md`.

# iter-01 progress

## What this iter did
- Ran the Phase 0b KB-fidelity gate (`audit-kb-fidelity --milestone=M235`) → **YELLOW**, 1 finding (KB-1, a
  stale fixture header comment), routed to the fixture tik. Report: `../kb-fidelity-audit.md`.
- Probed the environment: docker up but no stack running; rext clone @ `fd457bf`; prod DB access available.
- Measured the baseline gate distance (fixture matrix gaps + the live-proof gap) — see `overview.md`.
- Resolved the milestone's two open questions against M231 evidence (sim result = persisted read → renders
  from seed; not-passed renders a meaningful sub-threshold page).
- Authored TOK-01 (the two-track strategy) → milestone-root `decisions.md`.

## Close — 2026-07-19

**Outcome:** TOK-01 authored (two-track strategy: build+unit-prove readiness HERE / route the live cold-reset
browser proof locally-if-feasible-else-M236). KB-fidelity YELLOW. Both overview open questions resolved.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap does not terminate) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** TOK-01 (milestone-root); KB-1 (routed to the fixture tik).
**Side-deliverables (if any):** none.
**Routes carried forward:**
- iter-02 (first tik): fixture matrix closure (the 4 missing simulation sessions + KB-1 fix + unit proofs).
- later tiks: non-sim product sections (skill-path / academy / ai-labs) + player-path builders; Playthrough +
  coverage descriptors per (session × action); M230 clone re-anchor + `ANT_ACADEMY` descriptor +
  `getPublicCatalogView` 2nd manifest.
- Track B: the FORMAL live cold-reset-to-seed browser proof — attempt local, else route to M236.
**Lessons:** The gate has a hard build/prove split — most of M235 is buildable + unit-provable without a
browser; only the final cold-reset-to-seed landing needs a live stack. Sequencing the fixture substrate first
unblocks the manifest projection, the descriptors, and the live proof in one dependency chain.
