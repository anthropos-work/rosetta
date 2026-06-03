# M2c / iter-03 — progress

**Type:** tik · Active strategy: TOK-01

## Close — 2026-06-03
**Outcome:** RS256 foundation landed — `shared/rsa.go` (fixed demo RSA keypair + `MintRS256` + `JWKS` +
public-PEM); `clerk-frontend` serves the **real** JWKS (was the empty set). **Additive:** full suite 7/7,
Go gate 22/22, JS gate 9/9 — all green.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (express DNA still 0% — no runner/goldens yet; this lands the `JWKS/non-empty-rsa`
*behavior* + the RS256 minting the express genes need)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-03-D1 (fixed demo RSA key, disarmed-by-design, additive); iter-03-D2 (JWKS via clerk-frontend; crypto in shared/)
**Routes carried forward:** iter-04 → the `clerk-express/` seam + a **Node runner driving the real
`@clerk/express`/`@clerk/backend`** (the load-bearing fidelity test: does the genuine SDK accept our RS256 token?).
**Lessons:** additive RS256 holds — a real JWKS doesn't disturb the JS gate (no JWKS gene) or `authn` (HS256).
The crux is now whether `@clerk/backend.verifyToken` accepts our token (iss/azp/claim requirements) — iter-04
discovers it by **running the real SDK**.
