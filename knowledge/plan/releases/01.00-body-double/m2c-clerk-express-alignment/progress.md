# M2c — progress (iterative · Gate Outcome Ledger)

**Milestone:** M2c — `@clerk/express` backend session verification (RS256/JWKS) · **Shape:** iterative
**Status:** planned (scaffolded 2026-06-03; not yet built — `/developer-kit:build-mstone-iters M2c`)

## Exit gate
**alignment ≥ 95% overall AND 100% critical on `clerk-express-1.json`**, AND the load-bearing test passes:
a **real `@clerk/express` instance accepts a Clerkenstein-minted token and `getAuth()` yields the right
identity**. (Iterate with `/developer-kit:build-mstone-iters`; harden every ~10 tiks +once post-gate with
`/developer-kit:harden-mstone-iters`; close with `/developer-kit:close-milestone`.)

## Score arc (overall % / critical %) — update each iter
| Iter | Shape | What it did | Score (overall / critical) | Gate? |
|---|---|---|---|---|
| 01 | tok (bootstrap) | authored TOK-01 (RS256-native, additive-first, real-SDK runner) | 0% / 0% (no DNA yet) | — |
| 02 | tik | authored clerk-express-1.json (9 genes / 4 caps / 7 critical); validates | 0% / 0% (no mirror yet) | NOT MET |
| 03 | tik | RS256 foundation: keypair + MintRS256 + real JWKS (additive; M1/M2 green) | 0% / 0% (no runner yet) | NOT MET |

## Suggested opening iterations (the build refines these)
1. **tok (strategy) — resolve M2c-D2 + D5:** read how studio-desk's Clerk instance is wired vs the app's
   (shared token? → migration; separate? → additive RS256). Confirm `@clerk/express` is available offline
   under `studio-desk/node_modules`. Decide additive (Option A) vs migration (Option B) + the runner shape.
2. **tik — author `clerk-express-1.json`** (the ~8 genes, `spec-notes.md`) + validate it (`alignctl dna validate`).
3. **tik — RSA keypair + real JWKS** from the fake FAPI (`clerk-frontend` JWKS → non-empty RSA key);
   `JWKS/non-empty-rsa` gene green.
4. **tik — RS256 minting + the `clerk-express/` verifier path**; drive `ExpressAuth/*` + `ExtractIdentity`
   genes; capture goldens (hybrid per M1-D1 where live capture isn't available).
5. **tik — integration genes** `ClerkClientBAPI/*` vs the existing `clerk-backend` mock.
6. **tik — the load-bearing real-`@clerk/express` test** + drive the score to the **exit gate**.
   (If Option B: the iter that migrates the seams must re-capture the Go/JS DNA goldens + re-confirm M1/M2
   gates green — that re-gating is part of M2c, recorded as a tok.)

## Running ledger
- iter-01 (tok/bootstrap): authored TOK-01 — RS256-native, additive-first, real-`@clerk/express` Node runner — see iter-01/progress.md
- iter-02 (tik): authored `clerk-express-1.json` (9 genes / 4 caps / 7 critical), validates — see iter-02/progress.md
- iter-03 (tik): RS256 crypto foundation (keypair + MintRS256 + real JWKS), additive, gates green — see iter-03/progress.md

## Green-gate guard (if Option B / migration is chosen)
Any iteration that touches `authn` / `clerk-frontend` / `shared` MUST re-run the existing gates before
closing: Go gate 22/22, JS gate 9/9, drift-test 9/9 — re-capturing goldens as needed. A migration that
leaves M1/M2 red is incomplete.
