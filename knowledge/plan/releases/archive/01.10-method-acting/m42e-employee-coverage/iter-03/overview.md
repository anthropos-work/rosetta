---
iter: 03
milestone: M42e
iteration_type: tik
iter_shape: tooling
status: closed-fixed
created: 2026-06-25
---

# iter-03 — tooling-iter: fix the crawl wait strategy (the networkidle flake)

**Type:** tik / **iter_shape: tooling** (coverage-protocol.md "Iter type selection -> Tooling-iter"). The prior
tik (iter-02) closed blocked on a harness capability gap (D2: the crawl's `networkidle` wait never settles on
next-web, so every page false-times-out). This iter ships the harness capability AND uses it within the same
iter for the coverage re-measurement.

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage.** This tooling-iter is the highest-leverage move under TOK-01: the
single dominant cluster (all 44 baseline failures) is the `networkidle` flake; fixing the crawler's wait
strategy unblocks the TRUE measurement of every page in one change.

## Re-survey (Phase 1 Step 0)
The iter-02 baseline `(44,1)` is the current measured state; TOK-01's routed target (the tooling fix) is still
the dominant cluster + untouched in a committed form. **Important context:** the prior build-iter attempt
already AUTHORED this exact fix in the rext working tree (uncommitted WIP on `crawl.ts` + `coverage.spec.ts`) —
the `networkidle`->`domcontentloaded` swap + bounded settle + inline screenshots + seed-vs-discovered scoring.
This iter ADOPTS that WIP as its deliverable (it is precisely the D2/D4 fix), validates it by re-sweeping, and
commits+tags it properly (the WIP was never committed, so it is this iter's in-flight work, not a discard).

## Cluster / target identified
The harness wait strategy (`crawl.ts` `page.goto` + the screenshot pass in `coverage.spec.ts`). The fix:
1. `waitUntil: 'networkidle'` -> `'domcontentloaded'` + a **bounded** `networkidle` settle (`.catch(()=>{})` on
   a short timeout) so a long-polling page loads + paints but never blocks on never-idle.
2. **Inline screenshots** (an `onPage` hook in the crawl) instead of a 2nd full re-navigation pass — halves the
   nav count + removes the 2nd networkidle-timeout source that exhausted the 600s budget.
3. **Seed-vs-discovered scoring:** a guessed seed that 404s or redirects away is dropped (not false-scored as a
   coverage failure); only nav-discovered pages are coverage commitments. (D4.)

## Hypothesis
With `domcontentloaded` + bounded settle + inline screenshots, the sweep finishes within budget and every
reachable page is classified by its REAL content (not a timeout). The TRUE `(failing, escapes)` drops sharply
from `(44,1)` — most of the 44 were http=200-with-real-content pages mis-scored by the timeout.

## Expected lift
A large reduction in `failing` (most of the 44 were flake-induced). The TRUE residual — genuinely empty/errored
pages + the 1 real content escape — is what remains for iter-04+ to route by fix surface. Success = a credible,
budget-completing re-sweep with `failing` << 44.

## Phase plan (multi-step tooling-iter — planned 2 lines)
- **Line 1 (ship the capability):** adopt + validate the harness wait-strategy fix (compile-check; the WIP is
  already authored).
- **Line 2 (use it):** Phase D re-sweep against live demo-3 as `maya-thriving`; record the new `(failing,
  escapes)`; triage the TRUE residual for iter-04.
- **Phase E — Close:** grade on whether the harness fix landed + the re-sweep produced a credible reduced metric.
- Fold the wait-strategy lesson into `coverage-protocol.md` measurement conventions (protocol-evolution), same
  commit (corpus doc-half).

## Escalation conditions
- The re-sweep still false-times-out (the fix didn't take) -> a 2nd tooling-iter, route forward. Not a blocker.
- A genuinely-empty page in the TRUE residual that ONLY a platform change fixes -> re-scope-trigger.

## Acceptable close-no-lift outcomes
- If the re-sweep reveals the fix works but uncovers a NEW harness gap (e.g. budget still tight at the true page
  count) -> close on the landed capability + route the new gap; the metric still improves from the fix.
