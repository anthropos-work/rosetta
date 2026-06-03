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
| _(none yet)_ | | | 0% / 0% (no DNA authored) | — |

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
_(build-mstone-iters appends one line per closed iter here — `iter(M2c/NN): {tik|tok} — {outcome} → score X/Y`.)_

## Green-gate guard (if Option B / migration is chosen)
Any iteration that touches `authn` / `clerk-frontend` / `shared` MUST re-run the existing gates before
closing: Go gate 22/22, JS gate 9/9, drift-test 9/9 — re-capturing goldens as needed. A migration that
leaves M1/M2 red is incomplete.
