---
iter: 04
milestone: M2c
iteration_type: tik
status: closed-fixed
created: 2026-06-03
---

# M2c / iter-04 — tik: the crux proof (real `@clerk/backend` accepts our RS256 token)

**Type:** tik · **Active strategy:** TOK-01

## Step 0 — re-survey
Score 0%; RS256 foundation landed (iter-03). TOK-01's next: prove the **real** `@clerk/express`/`@clerk/backend`
accepts a Clerkenstein token — the milestone's make-or-break. Target current + decisive.

## Target
The load-bearing fidelity test: does the genuine `@clerk/backend.verifyToken` accept a `shared.MintRS256`
token? Build the runner skeleton (`alignment/cmd/expressrun`, Go) + the Node verifier (`verify.js`).

## Hypothesis
A Clerkenstein RS256 token (sub/iat/nbf/exp + platform claims, kid matching the JWKS) verifies under the
real `@clerk/backend` networkless path (`jwtKey` = `shared.RS256PublicKeyPEM()`).

## Result — PROVEN
`{"ok":true,"sub":"user_2clerkenstein","org_id":"org_clerkenstein","org_role":"admin"}` — the genuine SDK
accepted the token + extracted the identity. **No `iss`/`azp` tuning needed.** Additive RS256 (TOK-01)
validated; the RS256 wall is solved.

## Routes forward
iter-05 = the full runner protocol (all 9 genes: the 5 ExpressAuth scenarios + ExtractIdentity + JWKS +
the 2 ClerkClientBAPI integration genes) + hybrid goldens + alignctl wiring → drive the score.
iter-06 = the exit gate (≥95% / 100% critical) + the integration genes.
