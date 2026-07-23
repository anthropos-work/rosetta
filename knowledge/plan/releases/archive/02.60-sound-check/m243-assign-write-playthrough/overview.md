---
milestone: M243
slug: assign-write-playthrough
version: v2.6 "sound check"
milestone_shape: section
status: archived
created: 2026-07-20
last_updated: 2026-07-22
depends_on: M237
delivers: corpus/ops/demo/playthroughs.md
---

# M243 — assign-WRITE Playthrough  [realizes reserved M238]

**Status:** `archived` (completed 2026-07-22)  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** M237

## Goal
The one net-new hero journey — a manager assigns content with a deadline and it lands. Realizes the reserved `M238` (assign-WRITE); closes the ~10-routing `DEF-M235-03` / M204 in-manifest TODO.

## Scope
### In
- `playthroughs/manifest/assignment-monitoring.yaml` UC1 (`assign-and-track.UC1`, currently `TODO`).
- A new `/enterprise/assignments` page object.
- Possibly a `pt-world` precondition (assignable content + target member) in lockstep with `seed-worlds.yaml`.
- The spec `e2e/tests/assignment-assign.spec.ts` tagged `@pt:...UC1`. Takes the corpus **15 → 16** live Playthroughs, 0 TODO.

### Out
- The re-prove-on-billion live drive (M244 executes it).

## Open questions
- Does the `assign` WRITE need a `pt-world` precondition co-authored with `seed-worlds.yaml` (assignable content + a target member)?
- Needs a live browser drive against a running demo.

## Delivers
`corpus/ops/demo/playthroughs.md` (15 → 16 live Playthroughs; the assign-WRITE half of the M204 flow).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
