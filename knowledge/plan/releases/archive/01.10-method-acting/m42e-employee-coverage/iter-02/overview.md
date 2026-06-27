---
iter: 02
milestone: M42e
iteration_type: tik
status: closed-fixed
created: 2026-06-25
---

# iter-02 — baseline employee coverage sweep

**Type:** tik (first tik of the milestone; the baseline sweep). Runs under
[`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) Phase A (Sweep = the baseline).

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage** (milestone-root `decisions.md`). This tik executes TOK-01's
Next-tik direction verbatim: run `run-coverage.sh 3 employee maya-thriving` against live demo-3, capture the
baseline `(reachable, failing, escapes)`, triage the highest-leverage failing cluster.

## Re-survey (Phase 1 Step 0)
Re-ran the cheapest current-state checks against live demo-3 (17 containers up, consumed tag
`method-acting-m41`):
- next-web `:33000` serves; `/dashboard` 308→`/home` (the real landing route is `/home`, not `/dashboard`).
- fake-FAPI `:35400` is up + port-mapped; curl can't complete its TLS (openssl-fallback cert, exit 35) —
  expected; Playwright's `ignoreHTTPSErrors` handles it at the browser + request layer.
- Real employee-reachable routes (from `next-web-app/apps/web/src/app/(authenticated)/(verified)/`): `/home`,
  `/profile`, `/profile/activities`, `/profile/skills`, `/profile/aspirations`, `/library`,
  `/library/ai-simulations`, `/library/skill-paths`, `/settings`, `/sim/[slug]`, etc.
- TOK-01's named target (run the baseline) is still untouched + meaningful → no substitution.

## Cluster / target identified
No prior measurement exists — the baseline is unmeasured (TOK-01 Distance-to-gate). This tik's target is to
**establish the baseline** `(reachable, failing, escapes)` over the employee vantage's reachable set, then
triage the highest-leverage failing cluster as iter-03's target. Per coverage-protocol Phase A, iter-02 is the
baseline sweep.

## Hypothesis
The sweep runs end-to-end (login as `maya-thriving` via the cockpit handshake, BFS-crawl the in-app nav,
per-page assert + screenshot, emit the report). Known risk areas (TOK-01): the G7 activities feed
(under-investigated) + escape links (no studio-host URL rewrite baked → a left-menu "Studio" likely escapes to
prod). Expect a non-zero `(failing, escapes)` baseline.

## Expected lift
A baseline tik does not move the gate (there is no prior value to improve). Success = a clean, deterministic
coverage report is emitted with a credible `(reachable, failing, escapes)` triple, and the dominant failing
cluster is identified for iter-03. (Per coverage-protocol: the baseline sweep's deliverable is the measurement,
not a metric reduction.)

## Phase plan
- **Phase A — Sweep (baseline):** `run-coverage.sh 3 employee maya-thriving` against live demo-3, foreground.
- **Phase B — Triage:** classify each failing page / escape by fix surface (the routing table); pick the
  highest-leverage cluster as iter-03's target.
- **Phase C — Fix:** none this tik (baseline only) — unless a trivial in-scope fix is obvious and complete-able.
- **Phase D — Re-sweep:** n/a for a pure baseline (no fix landed) — OR re-sweep if Phase C lands a fix.
- **Phase E — Close:** grade on whether the baseline was credibly established (`closed-fixed` for a clean
  baseline measurement + triage; the metric "moves" from UNMEASURED to a concrete value).

## Escalation conditions
- The sweep crashes (login handshake fails / harness bug) → tooling-iter signal for iter-03 (the harness lacks
  a capability); record in decisions, route forward. NOT a user-blocker (the harness is rext, fixable).
- A 100%-blocking page resolvable ONLY by a platform change → re-scope-trigger (zero-edit line).

## Acceptable close-no-lift outcomes
- The sweep runs but `reachable === 0` (login failed): close-no-lift with the falsification recorded + route a
  tooling-iter forward to fix the login mechanism. This is the protocol working (a characterized harness gap),
  not a failure.
