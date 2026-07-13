---
iter: 2
milestone: M218
iteration_type: tik
status: closed-fixed
opened: 2026-07-13
closed: 2026-07-13
---

# iter-02 (tik) — the latency harness + the honest pre-fix baseline

**Active strategy:** **TOK-01** (`../decisions.md`) — *"reachability-first: fix the address, not the variable."*
Step 1 of TOK-01: *"build the latency harness first, and take a real p95 baseline before any fix."*

## Step 0 — re-survey
TOK-01's next-tik direction was authored **this same session** (iter-01), against a stack measured **minutes**
earlier. Nothing stale. Target confirmed: **no instrument exists**, so the gate cannot be graded. Building it is
the only thing that unblocks every later iter.

## Cluster / target
The gate is **p95 over 5 cold runs, both vantages** — and there is **no instrument that can produce that number**.
The corpus's only claim ("~2–5 s, which we can't shorten") is the thing we are here to falsify. Until an
instrument exists, every fix would be unfalsifiable.

## Hypothesis
A Playwright surface that drives the **real cockpit CTA** and polls for **content presence** (never `networkidle`)
can produce a trustworthy p95 with **per-leg attribution** — so a regression names the leg that regressed.

## Expected lift
**None on the gate metric — by design.** This tik's deliverable is the *instrument*, not a fix. Success = a
baseline number that is (a) reproducible, (b) attributed to a leg, (c) taken on a green stack.

## Phase plan
1. Build `lib/latency.ts` + `tests/latency.spec.ts` + `run-latency.sh` in rext `stack-verify/e2e/`.
2. Resolve the iter-01 **D4** collision (reuse `cockpit-login.ts` vs the `networkidle` ban) without forking it.
3. Gate on `autoverify.json` green (at its **real** path — iter-01 **F-2**).
4. Run it against the **current, unfixed** `billion` demo. Record the honest baseline.

## Escalation conditions
- If the harness cannot reach ACCESS at all → the login is **broken**, not slow → user-blocker.
- If the baseline contradicts iter-01's root cause → re-open the suspect list (**not** a re-scope; a re-measure).

## Acceptable close-no-lift outcomes
None applicable — an instrument either exists and produces a number, or it does not.

## Out
**Any fix.** TOK-01 puts the fix in iter-03. Measuring and fixing in one iter would leave the baseline unprovable.
