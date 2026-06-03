---
iter: 03
milestone: M2c
iteration_type: tik
status: closed-fixed
created: 2026-06-03
---

# M2c / iter-03 — tik: RS256 crypto foundation (keypair + mint + real JWKS)

**Type:** tik · **Active strategy:** TOK-01 (RS256-native, additive-first)

## Step 0 — re-survey
Score 0% (no mirror yet). TOK-01's next direction = RSA keypair + RS256 mint + real JWKS. Target current.
**Additive confirmed:** the JS DNA has 0 JWKS genes → serving a real key is safe for the JS gate.

## Target
The RS256 cryptographic foundation the express path needs: a fixed demo RSA keypair, RS256 minting, and a
real (non-empty) JWKS from `clerk-frontend` (the `JWKS/non-empty-rsa` gene's behavior).

## Hypothesis
`shared.MintRS256` + `shared.JWKS` (a fixed demo key) + `clerk-frontend` serving the real JWKS gives the
real `@clerk/express` something to verify against. **Additive** → M1/M2 gates stay green.

## Expected lift
The `JWKS/non-empty-rsa` behavior + express-token minting exist. Score stays ~0% until the Node runner
wires the genes (iter-04), but the foundation lands + gates stay green.

## Phase plan
`shared/rsa.go` (keypair + `MintRS256` + `JWKS`) → `clerk-frontend` serves `shared.JWKS()` → tests →
re-run M1/M2 gates (must stay green).
