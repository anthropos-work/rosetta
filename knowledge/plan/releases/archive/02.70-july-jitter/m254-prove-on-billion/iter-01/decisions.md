# M254 · iter-01 — decisions

_(Intra-iter decisions; the strategy TOK-01 lives in the milestone-root decisions.md.)_

## D1 — Phase 0b KB-fidelity = GREEN-by-inheritance (no standalone audit spawn)

M254 is a terminal PROOF milestone (`Delivers: none`) — it authors no new knowledge doc; it re-proves the
existing corpus live. Its "knowledge" IS the 4 protocol docs (`verification.md`, `tailscale-serve.md`,
`coverage-protocol.md`, `playthroughs.md`), each continuously updated + validated through M246–M253 (all
just closed). The **live cold bring-up + gate sweep is itself a far stronger fidelity check** than a static
pre-audit — any doc drift surfaces as an iter finding at measurement time. Recording GREEN-by-inheritance
with this rationale rather than spawning `/developer-kit:audit-kb-fidelity` is the correct proportional-rigor
call for a proof milestone. Recorded in `spec-notes.md` under "Pre-flight audits — iter-01".

## D2 — Pin = the cumulative code-of-record `july-jitter-m253-studio-first-paint` (b8969c0)

Verified all 7 milestone tags (m246-harden, m248, m249, m250, m251, m252, m253) are ancestors of rext main
HEAD `b8969c0`, which IS the m253 tag — so the single highest tag carries every release fix. Confirmed ON
ORIGIN via `git ls-remote --tags origin` (rung-zero). The billion demo will consume THIS pin (not the highest
fetched tag — the M244 iter-25 version-skew lesson).
