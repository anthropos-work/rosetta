# M2c — spec notes

Research basis: 3-agent workflow `clerk-express-milestone-research` (2026-06-03) — @clerk/express surface,
studio-desk usage, Clerkenstein fit/gap.

## What @clerk/express actually does (the surface)
- Official Clerk **Express/Node backend** SDK (^1.3.47), wrapping `@clerk/backend`. Key APIs studio-desk
  consumes: `clerkMiddleware()`, `requireAuth()`, `getAuth()`, plus `clerkClient.users.getOrganizationMembershipList()`
  + `clerkClient.organizations.getOrganization()`.
- **Verification = RS256 via JWKS, only.** Flow: `clerkMiddleware → authenticateRequest → verifyJwt →
  loadClerkJWKFromRemote (or from `jwtKey` PEM) → hasValidSignature` (RSASSA-PKCS1-v1_5). Two modes:
  (1) **networkless** via `jwtKey`/`CLERK_JWT_KEY` (a PEM public key); (2) **remote JWKS** fetched from
  the FAPI, cached ~5 min. `assertHeaderAlgorithm` allows **RS256/RS384/RS512** and **rejects HS256**
  (`TokenInvalidAlgorithm`). No `.well-known/openid-configuration` discovery.

## studio-desk usage (the consumer)
- `src/index.ts` (boot: `clerkMiddleware()`), `src/routes/ai.ts`, `src/routes/skillpath.ts`. Express
  backend on **port 9000** = API gateway for AI generation, Skill Path CRUD, YouTube proxy. All three
  route groups behind `requireAuth()` + a `checkEnterpriseAndAdmin` gate (uses
  `clerkClient.users.getOrganizationMembershipList()` for org roles + `clerkClient.organizations.getOrganization()`
  to read `publicMetadata.eid` → multi-tenant `loadOwnedPath` isolation).
- **MOCK_CLERK=true** swaps the auth middleware for a pass-through → a Clerk-free demo *runs* without M2c
  (studio-desk just isn't really authenticating). So M2c is **fidelity**, not demo-unblock.

## The gap (precise)
| Aspect | @clerk/express expects | Clerkenstein has today |
|---|---|---|
| JWT alg | RS256 (RSA) | **HS256** (HMAC, `shared/jwt.go` `universalSecret`) |
| Key source | RSA public key via JWKS (or `jwtKey` PEM) | **empty JWKS** `{"keys":[]}` (`clerk-frontend/server.go` `handleJWKS`) |
| Verify path | JWKS `kid` → RSA pubkey → RSASSA-PKCS1-v1_5 | `authn` HMAC-verifies with the shared secret; no JWKS |

⟹ a Clerkenstein HS256 token is rejected by `@clerk/express` at the algorithm check. **An HS256 shim is
impossible** (the SDK rejects before any interception point).

## Two implementation options (the central iteration)
**The token-sharing question first:** does studio-desk's `@clerk/express` verify the *same* session token
the Go `app` backend verifies via `authn`? In real Clerk, yes — one RS256 session token per instance,
verified by every backend via JWKS. If Clerkenstein keeps one shared token, it must be RS256 for *both*
authn and @clerk/express → migration. If the seams can use *separate* token domains (studio-desk mints +
verifies its own), RS256 can be additive. **Iteration 1 should settle this by reading how the platform
wires studio-desk's Clerk instance vs the app's.**

- **Option A — additive RS256 (preferred if token domains are separable):** add RS256 minting +
  keypair + a real JWKS used *only* by the `clerk-express/` path; leave `authn`/`clerk-frontend`/`shared`
  HS256 untouched (Go DNA + JS DNA stay green as-is). New `clerk-express/` seam + `clerk-express-1.json`.
  **No re-gating of M1/M2.**
- **Option B — RS256 migration (if the token is genuinely shared):** generate an RSA keypair; `shared`
  mints RS256; the fake FAPI serves a real JWKS; `authn` verifies RS256 via the pubkey; `clerk-frontend`
  mints RS256. **Ripples:** re-capture the Go DNA (`clerk-2.6.0`) `VerifyToken` goldens (now RS256), and
  the JS DNA's `SessionToken` genes; re-gate M1 + M2. Higher fidelity (matches real Clerk), higher cost.

Record the decision as **M2c-D2** once iteration 1 resolves the token-sharing question.

## How alignment is measured here (NOT a reimplementation)
`@clerk/express` is the **consumer/verifier**, not something we reimplement. Like the **svix** webhook
seam, Clerkenstein's job is to **produce** artifacts (RS256 tokens + a real JWKS) that the **real**
`@clerk/express` accepts. So:
- The "mirror" side = Clerkenstein's `clerk-express/` token+JWKS producer.
- The "source" side (goldens) = what real Clerk + real `@clerk/express` produce for the same scenarios.
- The **runner is Node-side** (real `@clerk/express`), analogous to `cmd/jsfapirun` for the JS surface —
  OR a Go runner that shells to a small Node verifier. Resolve the runner shape in iteration 1–2.
- The **load-bearing test** (svix-pattern): a real `@clerk/express` middleware, configured with the mock's
  `jwtKey`/JWKS, accepts a Clerkenstein-minted token + `getAuth()` yields the expected identity.

## Proposed DNA — `clerk-express-1.json` (~8 genes)
| Gene | Operator | Crit | Note |
|---|---|---|---|
| `ExpressAuth/valid` | error_class | ✓ | middleware accepts a good token |
| `ExpressAuth/expired` | error_class | ✓ | rejected, `TokenExpired` class |
| `ExpressAuth/malformed` | error_class | ✓ | rejected, malformed |
| `ExpressAuth/bad-signature` | error_class | ✓ | rejected, signature |
| `ExpressAuth/no-token` | error_class | ✓ | unauthenticated |
| `ExtractIdentity/universal-user` | exact | ✓ | `getAuth()` claims == platform identity (sub/eid/email/org_*) |
| `JWKS/non-empty-rsa` | shape | ✓ | FAPI JWKS returns a real RSA key (kid, kty=RSA, alg=RS256, n/e) |
| `ClerkClientBAPI/org-membership-list` | shape | — | the admin-gate path resolves vs `clerk-backend` |
| `ClerkClientBAPI/get-organization` | error_class | — | `publicMetadata.eid` read resolves vs `clerk-backend` |

## Risks
- **Re-gating ripple** (Option B) — the biggest; mitigated by trying Option A first.
- **Node runner in an offline/Go-centric harness** — the alignment harness is Go + offline; a Node
  runner needs `@clerk/express` available offline (it's under `anthropos-dev/studio-desk/node_modules`).
  Confirm offline availability in iteration 1; fallback = a Go RS256 verifier mirroring `@clerk/backend`'s
  JWKS verification (lower fidelity — measures our RS256, not the real SDK; the load-bearing real-SDK test
  still runs separately).
- **JWKS shape fidelity** — the served JWKS must match Clerk's real JWKS structure (kid/kty/alg/use/n/e).

## Pre-flight audits
Phase 0 (build-mstone-iters) runs against this. Load-bearing docs: `corpus/architecture/alignment_testing.md`
(the iteration protocol + DNA format), the clerkenstein `knowledge/` (alignment/architecture/injection),
and the `@clerk/express`/`@clerk/backend` source. No blind area expected (the framework + mirror exist).

### Pre-flight audits — iter-01 (bootstrap)
**Verdict: YELLOW** (lightweight check, 2026-06-03). The alignment framework + protocol
(`corpus/architecture/alignment_testing.md`) and the clerkenstein `knowledge/` base are **fresh** (just
authored/consolidated in M2b). The `@clerk/express` surface is **greenfield** — a doc area this milestone
*delivers* (the new `clerk-express/` README + alignment/architecture/sources updates), not a stale
load-bearing claim (same posture M0 took toward alignment). No RED blocker; TOK-01 accounts for the
greenfield area as known work. (Full `/developer-kit:audit-kb-fidelity` not invoked inline; the equivalent
contract-doc check is recorded here.)
