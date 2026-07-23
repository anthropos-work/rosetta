---
iter: 24
milestone: M244
iteration_type: tik
status: closed
run: 8
created: 2026-07-23
---

# iter-24 — BURNIN-M221 completion: full graphql `/dev-up --public-host` on billion

**Type:** tik

## Step 0 — Re-survey
Metric 6/8 (gate f 2/3). BURNIN-M221 is the last gate-f carry. Dev workspace already built (iter-22). The
pt-world reset (iter-23) already spent the demo seed → tearing down demo-1 for RAM is now free.

## Active strategy reference
**TOK-03** — final push: gate (f) BURNIN-M221 → 3/3 → 7/8. (Proceeds as a tik per D3/orchestrator
course-correction despite the nominal 3-no-prog floor; TOK-03 already dispositioned the coarse metric.)

## Cluster / target identified
Complete BURNIN-M221: tear down demo-1 (RAM) → `dev-stack up 2 --profile graphql --public-host
billion.taildc510.ts.net --no-setdress` on the built workspace → verify the dev stack comes up +
tailnet-reachable (the `--public-host` flag-path burn-in).

## Hypothesis
A full graphql dev-2 --public-host stack comes up on billion (demo-1 down for RAM) + is tailnet-reachable →
gate f 3/3 → 7/8.

## Expected lift
gate (f) 2/3 → 3/3 → metric 6/8 → 7/8.

## Phase plan
`tailscale-serve.md` + dev-stack tooling: tear down demo-1, background the graphql build (durable heartbeats +
final `BURNIN: pass|fail`), verify tailnet-reachability from the peer, restore billion.

## Escalation conditions
- If it needs a platform edit or is genuinely infeasible → surface a blocker.

## Acceptable close-no-lift outcomes
- A documented falsification: the build/up fails for a characterized (routed/blocker) reason.
