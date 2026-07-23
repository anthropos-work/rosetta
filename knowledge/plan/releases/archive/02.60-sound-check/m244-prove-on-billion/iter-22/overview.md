---
iter: 22
milestone: M244
iteration_type: tik
status: closed
run: 8
created: 2026-07-23
---

# iter-22 — BURNIN-M221: from-scratch dev `/dev-up --public-host` burn-in (investigation + workspace)

**Type:** tik

## Step 0 — Re-survey
Metric = 6/8. Gate (f) 2/3 (PROBE-M218-c3 + F-M220-4 done iter-21); BURNIN-M221 is the last gate-f carry.
TOK-03's next-target (gate f → 7/8 via BURNIN-M221). Untouched, meaningful.

## Active strategy reference
**TOK-03** — final-push step 2 completion: gate (f) BURNIN-M221 → ticks 7/8.

## Cluster / target identified
BURNIN-M221 = a real `/dev-up --public-host` remote DEV stack, live-cycled (the flag built M220, fenced
byte-identical, never brought up as a live dev stack). Plan (iter-21 D2): from-scratch reduced-profile
`/dev-up 2 --public-host` on billion (backend, offset 20000, alongside demo-1).

## Hypothesis
A reduced-profile (backend) dev-2 --public-host fits billion's RAM alongside demo-1 + exercises the dev-path
`--public-host` tailscale-serve wiring → gate (f) ticks 7/8.

## Expected lift
gate (f) 2/3 → 3/3 → metric 6/8 → 7/8.

## Phase plan
`corpus/ops/demo/tailscale-serve.md` + dev-stack tooling: (A) build the dev workspace (clone platform +
make init + values-blind dev .env); (B) `dev-stack up 2 --profile backend --public-host`; (C) verify
reachable + stable.

## Escalation conditions
- If BURNIN needs a platform edit or is genuinely infeasible in this environment → surface a blocker.

## Acceptable close-no-lift outcomes
- A documented falsification of the reduced-profile approach + a characterized feasible path forward.
