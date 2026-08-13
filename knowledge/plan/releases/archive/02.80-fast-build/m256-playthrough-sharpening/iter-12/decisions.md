# M256 · iter-12 — decisions

## D53 — a negative control lives OUTSIDE the Playthrough it covers

A cross-vantage control needs a second login and a second navigation. Inside the Playthrough that roughly
doubles its duration — and clause 1 gates the **median per Playthrough** at ≤ 0.79× of a 3.326 s baseline,
with the suite sitting near 2.3 s against a 2.628 s ceiling. Sixteen in-test controls would have **broken
the speed clause in order to satisfy the honesty clause**, which is not a trade the gate permits: both are
clauses of the same gate.

So the controls live in `negative-controls.spec.ts`, which declares **no `@pt:` id**. It is not a
Playthrough, `ptreport` does not reconcile it, and it never enters the median. The milestone's own plan had
already reserved this shape (*"they are excluded from the timed p50 either way"*) — this iter is where that
sentence became architecture. Batching by vantage is the other half: N absences cost ONE login.

## D54 — a control asserts LIVENESS before absence, and polls for it

iter-07 D29 refuted response-ablation because it produced a **dead page** (`bodyLen` 2147 → 24, 0 nav,
0 buttons) that satisfies *every* absence assertion. The rule that survives is: **an absence is evidence
only once the application is proven up.** Every control here asserts a populated body and rendered nav
first, and treats a liveness failure as *the control is broken*, not *the outcome is absent*.

**Polled, not read once** — and the first draft got this wrong in a way worth recording. Navigations settle
on `domcontentloaded` (the harness bans `networkidle`, iter-03), so a bare `.count()` immediately after
`goto()` reads the pre-hydration DOM and reports a *working* app as dead. It failed against a live app: a
false RED inside the very mechanism built to prevent false greens. `expect.poll` fixes it by waiting on the
landmark rather than on a network heuristic — the same doctrine as the rest of the harness.

## D55 — the coverage link is machine-checked and fail-closed in both directions

Coverage is now the union of two sources: a spec's own `@pt-negative-control:` line (the mutating specs,
whose pre-state read *is* the control — iter-06 D22) and a control spec's link tag. The fence:

- **rejects a link naming an id no Playthrough declares** — phantom coverage that looks real and proves
  nothing (a rename is the obvious way to create it);
- **rejects a malformed token** rather than skipping it — a silent skip is how a typo becomes invisible
  coverage;
- **reports** the count and names the uncovered ids on every run.

The count is now a *linked* fact, not a tag anyone can assert.

## D56 — the cross-vantage mechanism discriminates ORG- or HERO-specific outcomes ONLY

**The most consequential finding of the iter.** A contrast vantage works when the asserted outcome
legitimately does not exist for some hero or org. It cannot work for a **structural** final — a stat label,
a chart, a table's first row — because those render for *any* populated org or *any* seeded member.
Measured: `pt-manager`'s own profile reads `verifiedSkillsStat` 1, `skillCharts` 10, `workSection` 1,
because the M44 completeness seeder gives every member a career and skills.

So the routing's implicit assumption — one mechanism, 16 Playthroughs — was wrong. Nine of the remaining
eleven assert structurally and need their **finals sharpened to name real seeded data** instead (which
strengthens those Playthroughs whether or not a control follows). Writing contrast controls for them anyway
would have produced nine assertions that pass for any org: the exact vacuity iter-07 refuted, re-introduced
by a mechanism adopted to replace it.

## D57 — the hiring contrast vantage EJECTS TO PRODUCTION; never use it

Driving the hiring Results view as a Workforce-org manager sends the browser to
`https://app.anthropos.work/login?redirect_url=…` (bodyLen 162). The absence assertion would have been
"true" while the browser was **not in the demo at all** — the dead-page class in a new costume, and an
out-of-demo escape of the kind the coverage protocol counts as a hard failure. Recorded in the control
file's header so it is not re-tried, and noted as a demo-exposure observation in its own right.

## D58 — the fence harvested its own documentation (iter-07's defect, one grammar later)

Written literally in the fence's explanatory comment, the new link tag caused the fence to classify
**itself** as a control spec and parse its own sentences as an id list. This is precisely iter-07's finding
— *a comment minted a phantom Playthrough id* — reproduced by the guard built to prevent that class.

Fixed twice over, because either fix alone is fragile: **structurally** (`*.unit.spec.ts` files can never be
control specs — a meta-test covers nothing) and **by convention** (the tag is spelled apart in prose, as
iter-06 already established for `@pt-mutation`). The strict-token check is the third layer: a prose line
that slips through fails loud instead of being skipped.

**The transferable rule: if a tool scans source files for a marker, the tool's own file is a source file.**

## D59 — a `/g` regex reused with `.test()` is stateful

The control-spec detector initially reused the global match regex. `RegExp.prototype.test` advances
`lastIndex` on a `/g` pattern, so alternating calls would have matched every *other* file — silently halving
the control set and under-reporting coverage. Split into a non-global twin with the reason recorded at the
declaration.

## D60 — clause 1's verdict is NOT decidable at n=3 on this host; the earlier MET readings were sampling noise

Six full-suite runs this session, same host, and **the original 16 specs unchanged since iter-03**:

| statistic | min | max | spread | median (n=6) | gate |
|---|---:|---:|---:|---:|---:|
| all 22 non-studio (the GATED figure) | 0.5701× | 1.1121× | 1.95× | **0.8129×** | ≤ 0.79× |
| ORIGINAL 16 only (the control subset) | 0.5281× | 1.0762× | **2.04×** | 0.7063× | — |

The control subset — code no iter after 03 touched — varies by **2.04×**. There is **no trend**: the most
recent run reads 0.529× and the oldest reads 0.528×, with the 1.076× extreme in between. So this is not
host degradation over the session; it is variance that the pinned statistic (**median of 3 consecutive
runs**, against a baseline measured in a *different* batch) does not absorb. A batch of three can land
anywhere between ~0.53× and ~1.08×.

**Therefore:** at n=6 the gated figure is **0.8129×**, *outside* the `≤ 0.79×` gate. The flattering
denominator (the original 16, 0.706×) is inside it, and selecting that one would be the dishonesty this
milestone has spent eleven iters refusing. **Clause 1 cannot be declared MET on current evidence**, and the
"MET" readings from iter-03 onward have to be re-read as favourable samples rather than verdicts.

Two things this does **not** mean, stated so the escalation is not over-read: the iter-03 `networkidle` work
was real (its leg probe measured 2854 ms → 423 ms directly, not as a suite ratio), and no iter fabricated a
number — each reported its own batch honestly, and iters 06 and 08 explicitly attributed movement to
variance. What nobody tested is whether the **gate verdict** survives the variance.

**Every remedy is a release-level change to D-v28-12**, which is why this is escalated rather than actioned:
raise n and publish the spread alongside the median; make the measurement **paired** (re-measure the
baseline in the same batch as the treatment); normalise within-run against an invariant leg (e.g. the login
handshake); or move the measurement to a stable host — which is where D-v28-12's re-cut came from in the
first place. **A relative gate needs its noise floor published next to it, or it is not falsifiable.**

Practical note for whoever picks it up: `pt-assignment-assign` is the largest single contributor to both the
median and its variance (4.60 s → 6.44 s across batches, and the one Playthrough whose retry ladder has now
needed bounding twice). A speed lever aimed at it would improve both numbers at once.
