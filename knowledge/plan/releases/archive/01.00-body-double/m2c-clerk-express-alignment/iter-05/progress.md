# M2c / iter-05 — progress

**Type:** tik · Active strategy: TOK-01

## Close — 2026-06-03
**Outcome:** **GATE MET** — the `@clerk/express` surface scores **100% overall / 100% critical (9/9 genes)**.
Full `expressrun` runner (Go mints RS256 + embedded `verify.js` drives the **real** `@clerk/backend`) + 2
additive `clerk-backend` read handlers + captured goldens. Stable 3/3.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** iter-05-D1 (runner architecture: Go-mint + embedded-Node-verify + bapi-httptest + error-class normalization); iter-05-D2 (additive bapi reads keep the Go gate 22/22)
**Routes carried forward:** wire the express gate into `gate.sh` + CI (needs node + `EXPRESSRUN_NODE_PATH`) — harden/close polish, not a gate blocker.
**Lessons:** the real `@clerk/backend` accepts a minimal Clerk-shaped RS256 token; the **additive** RS256 path
(no HS256 migration) reaches 100% — TOK-01 validated end-to-end. Score arc: **0 → GATE in one runner tik**
after the crux. The runner's Node dependency (`@clerk/express`) is the env requirement to gate it in CI.
