# M2c / iter-02 — progress

**Type:** tik · Active strategy: TOK-01

## Close — 2026-06-03
**Outcome:** authored `clerk-express-1.json` (9 genes / 4 caps / 7 critical); `alignctl dna validate` OK.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0% — no mirror/goldens yet; the DNA is the measurement *target*, by design for a setup tik)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-02-D1 (operator + criticality choices)
**Routes carried forward:** iter-03 → RSA keypair + RS256 minting + a real JWKS from `clerk-frontend` (the
key material the express path needs + the `JWKS/non-empty-rsa` gene), additive (M1/M2 gates stay green).
**Lessons:** the DNA's input `op` codes (`express_auth` / `express_identity` / `jwks` / `clerk_client`)
are the **Node runner's contract** — iter-04 implements them against the real `@clerk/express`.
