---
iteration_type: tik
milestone: M254
iter: 07
status: closed-fixed-partial
---

# M254 · iter-07 — (g) host-sensitive demo-stack test-health

**Type:** tik · **Active strategy:** TOK-01 (cluster-per-tik live re-prove) — cluster 2 (g: the host-sensitive tests).

## Cluster / target
(g) The demo-stack test-health tests that fail on billion. Capture the exact failures live, fix the tractable
ones in rext (rung-zero), characterize + route the rest.

## Outcome
Captured 10 failures + 1 error / 159 (host-sensitive files) live on billion. Deduped: ~7 unique + 1 error,
with **diverse** root causes (NOT one common fix). Fixed the clearest high-leverage one — the nvm/node
host-robustness of `test_missing_node_documents` (×2) — verified live on billion. Characterized + routed the
remaining 6 (a dedicated test-health batch: intra-run port-leak isolation + M245 reconcile-message drift +
next.config sha drift + 2 mutation meta-tests + an overlay exit-127) with precise root causes + fix surfaces.

## Phase plan
verification.md measure→fix-forward. Fixed the tractable clear item; routed the test-maintenance batch.

## Escalation conditions
Any fix needing a platform edit → ESCALATE (none did — all rext test-harness/manifest).

## Acceptable close-no-lift
A precise falsification/characterization of each failure satisfies the protocol. Realized: 1 fixed + 6 routed
with named handlers → closed-fixed-partial.
