---
iter: 03
milestone: M226
iteration_type: tik
status: closed-fixed
date: 2026-07-17
---

# iter-03 (tik-2) — apply Finding-1's serve fix + measure the recruiter vantage (C2/C3/C5)

**Active strategy reference:** TOK-01 `reprove-hiring-on-billion` (iter-02's close routed this as the next tik).

**Step 0 — re-survey:** iter-02 confirmed :13001 unreachable from the peer (tailscale serve lacks 13001; curl
exit 35). The fix is committed+tagged (`casting-call-m226-serve-hiring`). Target current, no substitution.

## Cluster / target identified
Unblock + measure the 3 recruiter-vantage conditions the :13001 serve gap blocked in iter-02: C2 (recruiter
comparison render ≥40 rows/each of 5 positions), C3 (2 candidate profiles), C5 (recruiter p95 click→ACCESS < 5 s).

## Hypothesis
Consuming the fix (fronting :13001 over `tailscale serve`) makes the hiring app reachable from this Mac; the
recruiter comparison then renders (data is present: 294 sessions / 5 sims), the candidate profiles render, and the
recruiter login p95 measures < 5 s (same authenticated-shell path the employee/manager vantages already hit < 2 s).

## Expected lift
3/7 → 6/7 GREEN (C2, C3, C5 green), leaving C1 (the 3+47 count) as the sole remaining discrepancy (Finding-2,
its own handler). If C5 p95 ≥ 5 s (an R4-class latency finding) or the hiring app's ACCESS predicate misfires,
attribute (arithmetic signature per latency-budget.md) + fix in-bounds (tooling / sha-pinned demo-patch).

## Phase plan
1. Consume `casting-call-m226-serve-hiring` on billion (git fetch + checkout + rext.tag) — the running demo-1
   stays up; only the tooling clone advances.
2. Surgical `tailscale serve` re-apply with the fixed `gen_tailscale_serve.py` (front :13001) — no rebuild.
3. Confirm https://billion:13001 reachable from this Mac (the peer).
4. Measure C2 (`run-hiring-render.sh 1 rae-recruiter --hiring`, COVERAGE_RENDER_GATE=1, RENDER_HOST=billion...,
   RENDER_APP_SCHEME=https), C3 (candidate profiles render), C5 (`run-latency.sh 1 recruiter`, gated on a
   fresh-green autoverify.json copy) — all from this Mac.
5. Attribute any gap to its surface before any further fix.

## Escalation conditions
- C5 p95 ≥ 5 s → attribute the leg (arithmetic signature) + a sha-pinned perf demo-patch (R4); never a platform edit.
- hiring app ACCESS predicate misfire (heroIdentityPresent doesn't match the hiring DOM) → extend the predicate (tooling).
- An un-patchable platform surface → ESCALATE.

## Acceptable close-no-lift outcomes
If applying the fix does not make :13001 reachable and the cause is characterized (falsification), that's a
complete diagnostic cycle → closed-no-lift.

## Close
See `progress.md`.
