---
iter: 07
milestone: M42e
iteration_type: tik
iter_shape: tooling
status: closed-fixed
created: 2026-06-25
---

# iter-07 — tooling-iter: tighten the harness, exhaust the BFS frontier, quote the TRUE residual

**Type:** tik / **iter_shape: tooling** (coverage-protocol.md "Iter type selection -> Tooling-iter"). The run-1
verification (`.agentspace/scratch/work-m42e/verify-run1.md`, YELLOW) proved the committed `(failing=2,
escapes=3)` is HONEST but is a **FLOOR over a truncated slice** — every sweep saturated `COVERAGE_MAX_PAGES=80`
(`reachable===80===cap`), so only 14/22 chapter pages + 38/307 sims were reached and the frontier never
emptied. iter-07 ships the harness corrections the verdict mandated, then RAISES the cap until the BFS queue
EXHAUSTS (queue empty, not cap-hit) and re-sweeps to quote the REAL residual — the precondition for triaging it.

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage.** Still holds. The verdict didn't invalidate the strategy — it proved
the *measurement* was truncated. iter-07 fixes the measurement (raise the cap so the sweep measures the FULL
reachable set), which is the prerequisite for leverage-first triage of the true residual. This is a
tooling-iter under TOK-01 (ship harness capability + use it within the same iter), not a strategy revision.

## Re-survey (Phase 1 Step 0)
Current committed state `(failing=2, escapes=3)` over `reachable=80===maxPages` — a CAP-SATURATED measurement.
The verify-run1 verdict (independent adversarial review) is the current evidence: the residual is a floor, not
the truth. The 5 must-fixes are the iter's planned scope (a tooling-iter's planned multi-step shape).

## Cluster / target identified
The harness measurement itself — `coverage.spec.ts` + `crawl.ts`. The cap saturation hides the true page set;
no production triage is meaningful until the frontier exhausts. The verify verdict names the exact corrections.

## Hypothesis
With (a) the over-broad `/\/result\b/i` skip tightened to `/\/result\//` (no latent over-skip), (b) a cap-hit
flag + a `cappedAtFrontier` annotation added to the report (so `reachable` can't be misread as a true count),
and (c) `COVERAGE_MAX_PAGES` raised until the queue empties — the re-sweep produces a NON-truncated
`(failing, escapes, reachable)` where `reachable < cap` (queue exhausted). The chapter frontier MUST exhaust
(escapes live on chapter pages). The 307 sim detail pages, if template-identical, may be a documented
representative+boundary sample rather than all 307 — decided from the exhaustion evidence.

## Expected lift
The metric is RE-BASELINED to the true residual (likely failing >= 2, escapes > 3 vs the truncated floor) — a
tooling-iter's deliverable is the corrected measurement + the exhaustion proof, not necessarily a smaller
number. A bigger honest number that REPLACES a smaller dishonest floor IS the lift (the gate is over the FULL
reachable set; you cannot reach (0,0) on a set you never measured).

## Phase plan (multi-step tooling-iter -- planned lines)
- **Line 1 (ship corrections):** tighten the `/result` skip regex (must-fix 3); add a cap-hit flag +
  `cappedAtFrontier` field to the report + a runtime warning (must-fix 4); compile-validate.
- **Line 2 (use — raised-cap re-sweep):** raise `COVERAGE_MAX_PAGES` progressively across FOREGROUND calls
  until the report shows the queue EXHAUSTED (reachable < cap, no cap-hit), re-sweep vs live demo-3 as
  `maya-thriving`. Record the TRUE `(failing, escapes, reachable)`.
- **Line 3 (quote + triage-prep):** classify the full residual by fix surface; if the sim detail pages are
  template-identical, document the representative+boundary sampling rationale; surface the FULL citation set
  (must-fix 5 prep — apply the allow-rule only after the full set is known).
- **Phase E — close:** grade on whether the measurement is now non-truncated (the iter's planned deliverable);
  route the true-residual production fixes forward.
- Fold the cap-exhaustion + cap-hit-flag lessons into coverage-protocol.md (protocol-evolution).

## Escalation conditions
- If a raised cap reveals the queue still won't exhaust within a reasonable page budget (e.g. the 307 sims are
  genuinely distinct content) -> document the representative+boundary sampling decision with rationale per the
  verdict's allowance; that is NOT a re-scope, it's a scope-definition recorded in decisions.md.
- A 100%-blocking gap that ONLY a platform change can close -> re-scope-trigger (zero-edit line).

## Acceptable close-no-lift outcomes
- If the raised-cap sweep exhausts and the true residual is simply LARGER (more honest), close-fixed on the
  corrected-measurement deliverable — the tooling-iter shipped its planned capability + re-baselined the metric.
