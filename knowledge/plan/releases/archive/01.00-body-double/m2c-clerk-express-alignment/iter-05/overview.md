---
iter: 05
milestone: M2c
iteration_type: tik
status: closed-fixed
created: 2026-06-03
---

# M2c / iter-05 — tik: the full runner → exit gate

**Type:** tik · **Active strategy:** TOK-01

## Step 0 — re-survey
Crux proven (iter-04); RS256 foundation in. TOK-01's next: the full runner + goldens → drive the score.
**Phase 0 note:** the 3-no-prog-tik trigger (iter-02/03/04 each at score 0%) reads literally-fired, but it's
a **false-stall** — the score is structurally 0 until the runner+goldens exist, and TOK-01 is succeeding
(iter-04 proved the crux). Per Phase 1 triggered-tok Step 0 ("don't author a revision under false-stall
evidence"), no revision — tik under TOK-01.

## Target
The complete `expressrun` runner (all 9 genes, `--target source|mirror`) + the 2 `clerk-backend` read
handlers (the `ClerkClientBAPI` integration genes) + captured goldens → `alignctl` score.

## Hypothesis
The real `@clerk/backend` accepts our tokens (proven) → the full runner scores the differential genes
(ExpressAuth error classes + ExtractIdentity identity) at 100% and the structural genes (JWKS,
ClerkClientBAPI) confirm shape → the exit gate (≥95% overall / 100% critical) is met.

## Result — GATE MET
**overall 100% / critical 100% (9/9), stable 3/3.** ExpressAuth 5/5, ExtractIdentity 1/1, JWKS 1/1,
ClerkClientBAPI 2/2. M1/M2 gates (22/22, 9/9) + drift (9/9) stay green — additive.

## Phase plan
`expressrun` full + batch `verify.js` + `clerk-backend` reads + `alignctl capture` + `run` → 100%.
