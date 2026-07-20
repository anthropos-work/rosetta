# iter-07 — decisions

## D1 — the "hang" was the wait strategy, not the product

**Evidence.** A navigation probe against the reportedly-hanging route: handshake+commit 361 ms, 134
completed request legs, **0 still pending after 90 s**, slowest leg 761 ms (a static chunk), slowest
GraphQL call 133 ms, `getSkillPathDetails` 66 ms. The page was fully painted in about a second.

`content-stories.spec` called `loginAs` without `waitUntil`, inheriting the helper's `networkidle` default.
next-web holds long-poll connections open, so `networkidle` never resolves on this surface and `page.goto`
consumed the whole 180 s test timeout.

**Decision.** Navigate with `waitUntil: 'commit'`, made explicit at the call site with the cost recorded
inline. `settle()` takes over the waiting: poll until the readable text stops growing past a 200-char
floor, capped at `SETTLE_MS`, reading **`main` + `body`** (portals). Falls through rather than throwing, so
a genuinely empty page fails on its merits with a real excerpt instead of an opaque timeout.

**Why not raise the timeout.** It would have "fixed" the symptom and preserved the false pass in D3.

## D2 — skill-path-legacy has NO manager surface → player-link-only; denominator 31 → 29

**Evidence** (`next-web-app/apps/web/src/components/containers/InsightsBySkillPathStudentSimulationsContainer.tsx`):
`userData` is hardcoded `null` with the real lookup commented out; the results `<Table>` is commented out;
the drawer body renders `t('enterprise.insights.comingSoon')`. The only populated query is
`getSkillPathDetails` — the path *definition*. Live drawer text on both sessions ends in **"Coming soon"**.
No query reads the seeded session; the page is identical whether or not anything was seeded.

**Decision.** `managerKind: ""` for `skill-path-legacy`, the same disposition academy already carries. The
2 skill-path manager pairs are **not landable** and leave the denominator: **31 → 29**.

**Why this is not moving the goalposts.** M233's projection rule is explicitly fail-closed — *a session
that cannot form a real link is DROPPED with a reason; never a fabricated CTA*. ai-labs' player actions are
already excluded from the denominator on exactly this ground (no seedable result surface). This applies the
existing rule to evidence obtained for the first time by driving the route live, which is what the
milestone exists to do. 31 was never a count of provable pairs — it was a count that assumed a surface that
has not been built. The correction is recorded in full here and surfaced in the iter close rather than
folded silently into a metric.

**Reversibility.** Restore `managerKind: "skill-paths"` the day the platform implements the surface;
nothing else changes. The test in D3 fails loudly if it is restored before then.

## D3 — the sibling that "passed" was a false pass

The `manager-dashboard` shape asserts `/results for/i` present and `"No data"` absent. A definition-only
header satisfies the first; a "Coming soon" placeholder satisfies the second. So
`sp-genai-in-progress/manager` scored **PASS** for two iters against a page that proves nothing — the same
trap iter-05 found on the simulation scoreboard, recurring on a different surface.

`content_nonsim_test.go` had encoded the defect as the expected contract (it *required* the manager CTA).
Inverted: it now fails if a manager view is projected for skill-path at all.

**The dangerous half is the false pass, not the false fail.** A false fail costs an investigation; a false
pass hides a missing surface for as long as nobody looks. Both came from the same mis-calibrated shape, and
the shape only became visible because the route was driven live and *read*, not just graded.
