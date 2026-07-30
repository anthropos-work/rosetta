---
iteration_type: tik
iter_shape: standard
status: archived
opened: 2026-07-29
---

# iter-15 — the last two reachable controls, and the assertion satisfied by the wrong state

**Type:** tik · **Active strategy:** `TOK-01` move 4 ("close the honesty items … negative controls")

## Step 0 — re-survey (mandatory)

- `ptvalidate` → **VALID**, 24 live Playthroughs / 31 use cases / 7 TODO. Both trees clean at iter open.
- `@pt-negative-control` registry, computed: **20 of 24**. Named uncovered: `pt-activity-drilldown`,
  `pt-hiring-recruiter-compare`, `pt-studio-advanced-generate`, `pt-studio-guided-generate`.
- The gate's live gaps: clause 2 negative controls **20 of 24** (mutating 6/5 and `blocked` 1/1 both MET);
  clause 3 landed half short (org-admin 2/4, onboarding 1/5); clause 1 flake half MET, leg half N/A absent a
  speed mechanism; **D-v28-5** unblocked and unstarted.

**TOK-01's named target is still current** — no substitution. iter-14's routing names this iter's targets
explicitly and has already priced two of the three.

## Cluster / target identified — a declared THREE-step scope

1. **`pt-activity-drilldown`** — the last of the six Playthroughs `NEGCTL-M256-cross-vantage` set out to
   cover, and iter-14's Phase A measured both halves of it:
   - its grid assertions (`contentRows().first()` visible, then `count() > 0`) are satisfied by **20
     content-free skeleton rows** while the grid hydrates. Not a false green — the drill step that follows
     needs a row `<a>`, which a skeleton has not got — but a weak assertion with a live witness.
   - its control comes from the **drill-down**, which names members WITH their roles: Org A's first content
     drills to `Pat Ellis / DevOps Engineer` + `Morgan Reyes / Engineering Manager`; the SAME content id on
     Org C drills to `Mei Costa / Advanced Civil Engineer` + `Theo Lindqvist / Business Operations Analyst`.
2. **`PT-M256-readiness-step-asserts`** — routed to ride with this batch because it is the *same defect
   shape*: `pt-aireadiness-manager-howwemeasure` asserts `MANAGER_STEP_NAMES` through
   `stepMethod(name) = main().getByText(name).first()`, which matches **page-wide** — and `/ai-readiness`
   for an org WITHOUT the feature renders a live **upsell** panel naming those very steps. iter-12 found this
   with a control on its first run and deliberately did not assert it, so the weak sub-assertions are still
   in the Playthrough. Re-scope them inside the method panel.
3. **`pt-hiring-recruiter-compare`** — the open question, and the one this iter may not be able to close.
   iter-12 measured that its contrast vantage **ejects the browser to production**
   (`app.anthropos.work/login`, bodyLen 162), so a contrast org is out: an absence measured outside the demo
   is not evidence, and it is an out-of-demo escape besides. Its control must therefore come from a
   **sharpened final on the same vantage** — for which the seed does offer authored magnitudes (Org D:
   `size: 40`, `role_mix 0.1 admin / 0.9 candidate` → 5 shared positions / ~36 candidates) — but *where the
   ABSENCE comes from* is unanswered. **Investigate-or-verdict**, explicitly: if no honest absence exists on
   that vantage, the deliverable is a written verdict, not a manufactured control.

## Hypothesis

- **(1)** The drill-down is a per-member breakdown of ONE content item, so it names the org's own members.
  Sharpening the grid assertion to require rows that carry **text** closes the skeleton hole; sharpening the
  drill-down final to name the seeded hero **with her seeded role** makes it tenant-specific, at which point
  the contrast tenant's drill-down falsifies it.
- **(2)** The method-panel strings are inside a panel that only exists once the tab is open on an
  *enabled* org; the upsell panel is a different component. Scoping the accessor to the panel should keep
  the Playthrough green and make the contrast vantage's upsell **stop** satisfying it.
- **(3)** No prediction. iter-12's measurement stands; this iter asks only whether a same-vantage absence
  exists, and accepts "no" as an answer.

## Expected lift

- Negative controls **20 → 21 of 24** (22 if (3) yields an honest control, which is not predicted).
- One weak grid assertion closed with a measured witness, and one page-wide assertion re-scoped so it can no
  longer be satisfied by the not-enabled state.
- The remaining uncovered set becomes the 2 studio (blocked behind `FIX-M256-studio-false-green`) plus, if
  (3) yields no control, hiring **with a written verdict** — i.e. zero silent gaps either way.

## Phase plan

- **Phase A — probe first.** Measure, on `demo-2`: (a) the drill-down on BOTH vantages with a
  *content-aware* row count (iter-14 D71 — a settle predicate the empty state satisfies is not one), and
  (b) the method panel's real DOM shape on an enabled org AND the upsell panel's on a non-enabled one
  (iter-14 D72 — the DOM shape is a measurement, never an inference from a sibling accessor).
- **Phase B — sharpen** (1) and (2). Extend `seed-facts.ts` only if a new authored fact is needed.
- **Phase C — controls, each watched going RED**, in the two groups iter-14 D73 established: re-aim the
  control at the contrast tenant's own data, **and** drive the Playthrough itself on the contrast tenant.
- **Phase D — re-measure:** full suite ×3 cold reset-to-seed, the computed registries, `ptreport`,
  `ptvalidate`, unit fences, six Go modules, four Python suites.
- **Phase E — close:** commit both repos, tag + push + verify on origin, doc backfill if a lesson generalises.

## Escalation conditions

- If (3) proves to have no honest same-vantage absence, that is a **verdict, not a failure** — record it with
  the evidence and leave the count at 21. Manufacturing one would certify nothing (the iter-12 rule).
- If the method-panel re-scope makes the Playthrough RED on an enabled org, the panel-scoping hypothesis is
  refuted: revert the re-scope, keep `methodHeading()` as the discriminating assertion it already is, and
  record the measurement.
- If (1) and (2) land but (3) turns into a build, route (3) forward rather than opening a fourth line.

## Acceptable close-no-lift outcomes

A measured refutation of (2)'s panel-scoping, or a written verdict for (3) with the four-way evidence
attached — the same shape iter-05, iter-07 and iter-12 produced.
