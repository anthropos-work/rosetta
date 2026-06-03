# M2c / iter-04 — progress

**Type:** tik · Active strategy: TOK-01

## Close — 2026-06-03
**Outcome:** **CRUX PROVEN** — the genuine `@clerk/backend.verifyToken` accepts a Clerkenstein RS256 token
(networkless `jwtKey`) + extracts the exact identity (`sub`/`org_id`/`org_role`). No `iss`/`azp` tuning
needed. Landed `alignment/cmd/expressrun` (Go runner skeleton, `--emit`) + `verify.js` (Node, real `@clerk/backend`).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (formal score still 0% — no full runner/goldens yet; but the milestone's defining risk,
the RS256 wall, is now decisively solved — ExpressAuth/valid + ExtractIdentity behaviors verified against the real SDK)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-04-D1 (the crux result + that additive RS256 is validated; runner = Go-mints + Node-verifies hybrid, M2c-D5 confirmed)
**Routes carried forward:** iter-05 = full runner protocol (9 genes + outcome protocol + hybrid goldens + alignctl wiring); iter-06 = the gate + ClerkClientBAPI integration genes (need clerk-backend wired to clerkClient).
**Lessons:** `@clerk/backend` accepts a minimal Clerk-shaped RS256 token (`sub`+timing+platform claims) via
`jwtKey` with **no issuer/azp validation** unless those options are passed — so the additive path needs no
migration. The Node verifier resolves `@clerk/backend` from studio-desk's `node_modules` (NODE_PATH); iter-05
makes the runner self-contained (its own package.json) or keeps the NODE_PATH shim.
