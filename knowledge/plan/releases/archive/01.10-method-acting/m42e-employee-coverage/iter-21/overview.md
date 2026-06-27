---
iteration_type: tik
iter_shape: tooling
status: planned
created: 2026-06-25
---

# iter-21 — P7 semantic-coverage harness rebuild (the gate-measuring tool)

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause). The re-scoped gate (overview (a)-(d))
demands real semantic content + substantial cardinality + persona self-consistency + no prod-eject — which the
old `textLen>40` density harness cannot measure. P7 builds the **manifest-driven semantic gate** that does.

**Iter shape:** tooling (multi-step planned: build the harness modules → calibrate the floors against the live
render → self-test fails-old/passes-fixed). Per coverage-protocol.md's tooling-iter pattern.

**Cluster / target:** the harness rebuild. Replace the crawl's `textLen>40` content verdict with per-page,
per-section descriptors; add persona self-consistency; demote the crawl to reachability+escape-only; emit a
human review HTML; calibrate the cardinality floors; update the protocol doc; self-test on demo-3.

**Hypothesis:** a manifest-driven gate (region selector + realContent + cardinality floor + documented
exception, region-not-found=FAIL) + persona-assert (role↔skills from the rendered role panel, avatar
menu==profile+real-photo, org name+logo) will (a) PASS the now-fixed P0-P6 content pages and (b) FAIL the
believability gaps the old gate missed — proving the gate discriminates.

**Phase plan:** build empty-states.ts + coverage-manifest.ts (employee CALIBRATED + manager M36 authored) +
section-assert.ts + persona-assert.ts → demote crawl.ts → emit coverage-review.html → ONE calibration sweep to
tune floors → self-test sweep on demo-3 (fails-old / passes-fixed) → update coverage-protocol.md.

**Expected lift:** a working semantic gate + the employee residual under it (NOT gate-met — P8 is the
authoritative fresh-demo-up gate).

**Acceptable close outcomes:** the harness is built + self-tested (discriminating), the employee residual is
measured + honestly reported, the floors are calibrated. A residual of persona/section failures that are a live
re-apply gap (not a code gap) is routed to P8, not a blocker.
