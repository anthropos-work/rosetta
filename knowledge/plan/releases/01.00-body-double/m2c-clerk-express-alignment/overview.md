---
milestone: M2c
slug: clerk-express-alignment
version: v1.0 "body double"
milestone_shape: iterative
status: planned
exit_gate: "alignment >= 95% overall AND 100% critical on clerk-express-1.json, AND a real @clerk/express instance accepts a Clerkenstein-minted token and extracts the right identity"
iteration_protocol_ref: corpus/architecture/alignment_testing.md
started: 2026-06-03
last_updated: 2026-06-03
---

# M2c — Clerkenstein: `@clerk/express` backend session verification (RS256/JWKS)

## Goal
Bring the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend auth) under the
alignment framework: a new **`clerk-express/`** seam + a **3rd Alignment DNA** (`clerk-express-1.json`),
driven to a gate, so studio-desk's backend genuinely verifies Clerkenstein tokens instead of falling back
to its `MOCK_CLERK=true` bypass. This **completes v1.0's thesis** — *no* Clerk seam left un-faithful
before `/developer-kit:close-release`.

## The defining unknown — the RS256 wall (why this is iterative)
Research (3 agents, 2026-06-03) found a fundamental mismatch:
- **`@clerk/express` verifies RS256 via JWKS only.** It wraps `@clerk/backend`, whose
  `assertHeaderAlgorithm` accepts **RS256/RS384/RS512** and **hard-rejects HS256** with
  `TokenInvalidAlgorithm` — *before* any middleware can intercept. No HS256 path exists.
- **Clerkenstein mints HS256 universal-key tokens** (`shared/jwt.go`, `universalSecret`) and serves an
  **empty JWKS** (`clerk-frontend` `/.well-known/jwks.json` → `{"keys":[]}`).
- ⟹ An HS256 verification *shim* is a **dead end**. M2c must add an **RS256 path**: an RSA keypair, a
  *real* (non-empty) JWKS from the fake FAPI, RS256 token minting, and a verifier driving the **real
  `@clerk/express` SDK**.

**The central iteration question (resolved by iterating toward the gate):** can RS256 be **additive /
parallel** (a second token type, existing HS256 seams untouched), or must the existing seams
(`authn` / `clerk-frontend` / `shared`) **migrate to RS256** — re-capturing the Go DNA goldens and
**re-gating M1/M2**? See `spec-notes.md` for the two options + the token-sharing analysis.

## Scope
### In
- A new **`clerk-express/`** seam (library-named, like the other mocks).
- An **RSA keypair** + a **real (non-empty) JWKS** served by the fake FAPI (`clerk-frontend`'s JWKS
  endpoint), and **RS256 token minting**.
- The `@clerk/express` **DNA** (`clerk-express-1.json`; source `@clerk/express ^1.3.47`, which wraps
  `@clerk/backend`).
- A **runner** that drives the **real `@clerk/express` SDK** against the mock — the svix-pattern
  ("verify against the genuine library", not a reimplementation of `@clerk/express`).
- The **alignment gate** as the exit criterion.

### In — confirm, don't rebuild
- `@clerk/express` also calls `clerkClient.{getOrganizationMembershipList, getOrganization}` — those are
  **BAPI** calls, already 100%-mocked by `clerk-backend/`. M2c adds *integration* genes confirming that
  path resolves; it does **not** build a new BAPI mock.

### Out
- Any change to studio-desk or the platform repos (its `MOCK_CLERK=true` bypass is the platform's own).
- A webhook (svix) DNA (a separate future gap — svix is *used*, not mocked).
- Live injection into a running studio-desk (demo-stack work, v1.1).

## Candidate genes (~8 — `clerk-express-1.json`)
- `ExpressAuth/valid` · `ExpressAuth/expired` · `ExpressAuth/malformed` · `ExpressAuth/bad-signature` ·
  `ExpressAuth/no-token` — operator `error_class` (the middleware's accept/reject outcome).
- `ExtractIdentity/universal-user` — operator `exact` (verified claims → `req.auth` identity matches the
  platform identity).
- `JWKS/non-empty-rsa` — operator `shape` (the FAPI serves a real RSA public key in JWKS form).
- `ClerkClientBAPI/org-membership-list` · `ClerkClientBAPI/get-organization` — integration genes vs the
  existing `clerk-backend` mock (the admin-gate path studio-desk uses).

## Exit gate
**alignment ≥ 95% overall AND 100% critical on `clerk-express-1.json`**, AND the load-bearing test passes:
**a real `@clerk/express` instance, pointed at the mock, accepts a Clerkenstein-minted token and extracts
the right identity.** (The gate is the iterative exit; `/developer-kit:build-mstone-iters` drives toward
it, `/developer-kit:harden-mstone-iters` + `/developer-kit:close-milestone` finish.)

## Where it lives
The new `clerk-express/` seam + the 3rd DNA + the runner live in the `clerkenstein` repo (commits stack on
its `main`). The rosetta-side milestone records + the corpus updates land on the
`m2c/clerk-express-alignment` branch → merged to `release/01.00-body-double` at close. **Then**
`/developer-kit:close-release` ships v1.0.

## ⚠️ Re-gating risk (the placement trade-off — accepted, M2c-D3)
The user chose to land this in **v1.0** (re-opening the release a 3rd time) to complete the mock before
shipping, **accepting** that the RS256 path may force re-capturing the Go DNA goldens + re-gating M1/M2's
HS256 seams. If an iteration shows the migration is too disruptive, the escape hatch is to keep RS256
**additive** (a parallel token type for the `@clerk/express` seam only) — the iterations decide.

## Re-scope trigger
If 5 consecutive triggered toks fail to produce a viable new strategy, OR the additive-RS256 path proves
infeasible AND the migration path would re-gate M1/M2 beyond acceptable scope, escalate to
user-strategic-replan (the user weighs migration cost vs. accepting studio-desk's `MOCK_CLERK` bypass).
