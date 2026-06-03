# M2c — decisions

## M2c-D1 — RS256 is mandatory; the HS256 shim is a dead end (design-time, research-confirmed)
`@clerk/express` (via `@clerk/backend`) verifies **RS256/RS384/RS512 only** and rejects HS256 at
`assertHeaderAlgorithm` (`TokenInvalidAlgorithm`) before any middleware interception. Clerkenstein's
HS256 universal-key tokens cannot pass. ⟹ M2c **must** add an RS256 path (RSA keypair + real JWKS + RS256
minting). No HS256 verification shim is possible. (Research: `clerk-express-milestone-research`, 2026-06-03.)

## M2c-D2 — additive RS256 vs. RS256 migration: OPEN (the central iteration)
Whether RS256 can be **additive** (a parallel token type for `clerk-express/` only, M1/M2 seams untouched)
or requires a **migration** (existing `authn`/`clerk-frontend`/`shared` move to RS256, re-gating M1/M2)
depends on whether studio-desk's `@clerk/express` verifies the **same** session token the Go `app`
backend verifies via `authn`. **To resolve in iteration 1** by reading how the platform wires
studio-desk's Clerk instance vs the app's. Prefer **additive** (Option A) — try it first; fall back to
migration (Option B) only if the token is genuinely shared. Record the resolution here when settled.

## M2c-D3 — placement = v1.0 (user-chosen 2026-06-03), re-opening the release a 3rd time
The user chose **v1.0 as M2c** over v1.1 / a standalone track, to **complete the mock before shipping** —
no Clerk consumer left un-faithful before `/developer-kit:close-release`. Trade-off **explicitly
accepted**: this re-opens v1.0 (after M2b), delays close-release, and risks the RS256 path re-gating the
shipped HS256 seams (M2c-D2). Alternatives considered: v1.1 first-milestone (ship v1.0 now; demo runs via
studio-desk's `MOCK_CLERK` bypass) and a standalone "surface expansion" track — both rejected in favor of
completeness-first.

## M2c-D4 — `clerkClient.*` BAPI calls are already covered (de-scope to integration genes)
studio-desk's `@clerk/express` use includes `clerkClient.users.getOrganizationMembershipList()` +
`clerkClient.organizations.getOrganization()` — these are **BAPI** calls already 100%-mocked by
`clerk-backend/` (M1/M2). M2c adds **integration** genes confirming the path resolves against the existing
mock; it does **not** build a new BAPI mock.

## M2c-D5 — measured like svix (verify against the real library), not a reimplementation
`@clerk/express` is the verifier/consumer; Clerkenstein **produces** RS256 tokens + a real JWKS the
**genuine** `@clerk/express` accepts (the svix-pattern). The "mirror" is the producer; the load-bearing
test runs a real `@clerk/express` instance against the mock. Runner shape (Node-side vs Go-shells-to-Node
vs a Go RS256 verifier fallback) — **to resolve in iteration 1–2** (depends on offline availability of
`@clerk/express` under `anthropos-dev/studio-desk/node_modules`).

<!-- Iteration decisions (toks, escape-hatch escalations, user-blockers) recorded here as they arise. -->
