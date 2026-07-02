# M204 Decisions

Implementation decisions with rationale (recorded during the iter loop). Design-time decisions live in
[`overview.md`](overview.md) + the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| | _(intra-iter decisions live in each `iter-NN/decisions.md`; strategy TOK entries below)_ | | |
| D-CLOSE-1 | **`assignment-monitoring.assign-and-track.UC1` (the assign-WRITE half) → Fate-2 (tracked in-manifest), no plan edit.** | The assign-WRITE half is a two-backend org-admin WRITE flow, OUT of M204's declared 3 (READ/monitoring) manager journeys from the start (M201 declared it as one of two distinct flows; M204 landed the sibling monitoring flow UC2 green). It is durably tracked in the corpus: `assignment-monitoring.yaml` declares it `playthrough: TODO`, `ptreport` surfaces it as `unimplemented` (a first-class tracked state, not a silent drop), and the harden Pass-2 `TestRealCorpus_ManagerCoverageIsPresent` pin enforces the TODO stays declared (fail-red if removed). Not a repeat-defer (single M201 declaration); no aging trigger. No current v2.0 milestone owns the manager-WRITE class (M205 Hiring/tier-gates · M206 AI-sim-mirror · M207 Academy) — a future manager-write tier is the natural home; the manifest TODO is the ready-made declaration. Deferral audit: GREEN — see [`audit-deferrals/deferral-audit-2026-07-02-m204-close.md`](audit-deferrals/deferral-audit-2026-07-02-m204-close.md). | 2026-07-02 |

## TOK-01: manager-surface-per-iter — 2026-07-02

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Grow manager-vantage coverage **one manager surface per tik**, following the M203
measure→declare→page-object→play→diagnose→re-measure loop (`corpus/ops/demo/playthroughs.md` § iteration
protocol), on the **shared per-surface page-object layer** — an ADDITIVE merge with M203's employee surfaces
(each vantage adds its own page objects; no collision). Concretely, per tik:
1. **Declare** the tik's manager use case(s) in a new/extended manager product manifest (`workforce.yaml` for the
   `workforce-intelligence` product; `assignment-monitoring.yaml` for the activity-dashboard drill-down) —
   `actor.hero: pt-manager` (Morgan Reyes), `actor.entitlement: enterprise`, `seed.world: pt-world`,
   `playthrough: TODO` until built. Keep `ptvalidate` green (unique ids, both-way integrity, precondition-coverage).
2. **Add the page object** for the manager surface (`workforce-page.ts` / `members-page.ts` /
   `activity-dashboard-page.ts` / `succession-page.ts` as needed) extending `PageObject`, under the §5.2 locator
   discipline (scope-within-a-named-region, disambiguate by visible text; antd v6 near-zero a11y → find-only
   landmarks: `<main>`, h1–h4 headings, visible stat labels, domain text). Single-source any route shape in
   `url-shapes.ts` (anchored, never a bare `\b`).
3. **Play** as Morgan via the existing `hero-login` seat-switch (`identityKey: pt-manager`, land on the surface),
   assert the **user-observable OUTCOME**: the org-scale aggregate surface renders REAL data (P2 — presence /
   structure / non-empty, never exact seeded counts, never placeholder).
4. **Run** `run-playthroughs.sh 1 --reset --grep <@pt tag>` → `ptreport --gate no-regressions` → read the
   four-state map. All 4 manager UCs are **READ/monitoring** flows (no mutation) → reset is a determinism
   formality, but the 5-run zero-false-fail gate still exercises it.
5. **Diagnose** each non-passing: `failing` → fix page-object/seed or diagnose real capability/seed-drift (suspect
   seed-scale first — the key known-context); `unimplementable-without-platform-edit` → escalate (declare in
   `unimplementable.yaml`), NEVER edit the platform.

**Surface order (evidence + risk):** iter-02 Workforce funnel + roster (the manager landing surface, lowest-risk
READ, natural manager-login proof) → activity-dashboard drill-down → succession/at-risk. Order may re-survey per
Step 0 as evidence lands.

**Rationale:** This is the exact loop M203 used to take the employee vantage to green on the same foundation —
proven. The manager surfaces are all READ (monitoring) flows, so the risk is not mutation-determinism but
(a) whether the pt-world seed renders the M36 org-dashboard aggregates at Org A (size:40), and (b) antd
row/grid ambiguity (the §5.2 "load-bearing registry" case). Per-surface tiks keep each iter reviewable and let
the seed-scale question be answered against a real render before concluding a gap. One surface per tik also
matches the O(surfaces) re-pin economics of the page-object layer.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** Gate metric = manager-vantage UCs `passing` ÷ declared, with 0 false-fails over 5
reset runs. Starting value = **0 manager UCs declared, 0 passing** (the manifest holds only M203's 3 employee
products / 6 UCs). Target = the 3 declared manager journeys (4 UCs) all passing + the 5-run determinism gate.

**Next-tik direction:** iter-02 (first tik) — Workforce funnel + member roster
(`workforce-intelligence.skills-funnel.UC1` + `workforce-intelligence.roster.UC1`). Author `workforce.yaml`, a
`workforce-page.ts` + `members-page.ts` page object, the first manager spec(s); confirm the seed renders the
funnel + roster at Org A before concluding a gap; run `--reset --grep`; reconcile with `ptreport`.
