---
iteration_type: tik
status: in-progress
gate: PENDING
---

# iter-07 — tik: close the 3 adjacent manager sections (part 4 → gate 5/5)

**Active strategy:** TOK-01 (the render→triage→fix→re-render loop). **Target:** the gate distance from
iter-06 — part 4's three failing manager-dashboard sections on `/ai-readiness` (`ai-readiness-by-tag`,
`ai-readiness-interview-findings`, `ai-readiness-handled-for-you`). Close all three → gate 5/5.

## Re-survey (Phase 1 Step 0)
iter-06 routed these three forward as the sole gate distance. Re-measured each at source (platform read
paths + demo-2 DB) BEFORE touching the render. Finding: **all three are post-M246 platform DRIFT** — the
demo/manifest reference vocabulary the platform changed after the milestone's KB snapshot. Not a data gap.

## Cluster / target identified
The single root cause: post-M246 copy/vocabulary drift the coverage manifest (and, for #2, the M219 seeder)
still asserts against. Three sub-fixes, one shared measurement (a manager coverage sweep):

1. **`ai-readiness-handled-for-you` (manifest-only).** The tile renders cells 2+3 as
   `minutesToHours(handsOnMinutes|interviewMinutes)` under the `howWeMeasure.hoursSaved` label = **"Hours
   saved"** (HowWeMeasureTab.tsx + en/enterprise.json:173) — NOT "minutes saved". The regex asserted the
   dead literal, matching only cell 1 → false RED. **Data confirmed correct** at the demo-2 DB: skillsMapped
   4272 · handsOnMinutes 5430→91h · interviewMinutes 4770→80h (all three cells non-zero). Fix = regex label.
2. **`ai-readiness-by-tag` (manifest-only).** `byTag.title` = **"AI Readiness by Team"** (en/enterprise.json),
   not "…by Tag". The manifest's stale title was the ONLY missing marker — `'People involved'` (the
   non-empty-table-only column header) DID render, proving the table is populated (DB: 199 started-and-tagged
   members across 13 team tags → 13 rows). Fix = title copy. (The iter-06 "team-tag lane" scoping-guard
   concern is MOOT — no team-tag build; a one-word drift.)
3. **`ai-readiness-interview-findings` (seeder + manifest).** The post-M246 `usageDimSpecs` keys on
   **avg_adoption / avg_transformation / avg_originality / avg_depth / avg_ownership** (tiles Adoption /
   Transformation / Originality / Depth·Creation / Critical ownership). The M219 seeder wrote the retired ids
   avg_frequency/avg_breadth/avg_context_fit — only avg_depth survived, so 4 of 5 tiles were dropped and
   'Breadth'/'Context fit' never rendered. Fix = seeder emits the 5 real ids (spread across bands) + manifest
   asserts three of the newly-non-depth names.

## Hypothesis
Correcting the manifest to the real rendered copy (all 3) + reseeding demo-2 so the interview report carries
the 5 current KPI ids (#2) makes the manager sweep report `failingSections=0` on `/ai-readiness`, closing
part 4 → gate 5/5, with escapes=0 preserved (no platform edit; the employee vantage untouched, still GREEN).

## Expected lift
gate 4/5 → 5/5 (part 4's three sections flip PASS on live render; parts 1,2,3,5 already PASS).

## Phase plan
TOK-01 render loop: fix (landed) → build stackseed (rext) → reset-to-seed demo-2 + directus set-dress →
manager coverage sweep (dana-manager, Northwind) → measure the 3 sections + failingSections + escapes.

## Escalation conditions
- If a section still fails after the fix with a NEW root cause (real data gap, not copy) → triage; if it needs
  a platform edit or a large build beyond the fidelity gate → route Fate-2/Fate-3 + surface, do NOT over-build.
- If the sweep infra breaks → user-blocker.

## Acceptable close-no-lift outcomes
- A section proven to need a platform-repo edit to render (M250 forbids platform edits) → documented
  falsification + Fate-3 route, not a milestone failure.
