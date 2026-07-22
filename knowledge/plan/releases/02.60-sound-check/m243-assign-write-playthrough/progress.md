# M243 — Progress

Section milestone. Checklist stub from the roadmap In-list. **COMPLETE — live GREEN + DB-verified.**

## Sections

- [x] **`assign-and-track.UC1`** — filled `playthroughs/manifest/assignment-monitoring.yaml` UC1
  (`playthrough: pt-assignment-assign`; `preconditions: [public-catalog, org-unassigned-member]`).
- [x] **`/enterprise/assignments` page object** — `e2e/lib/assignments-page.ts` (`AssignmentsPage`),
  the first MUTATING page object; keyboard-picks the antd catalog Select; read-back = affordance-count delta.
- [x] **`pt-world` precondition** — `org-unassigned-member` added to `seed-worlds.yaml` in lockstep
  (no seed CODE change needed — Org A already provides 34/40 unassigned targets).
- [x] **Spec `e2e/tests/assignment-assign.spec.ts`** — tagged `@pt:pt-assignment-assign`; live GREEN on
  demo-1 (7.9s); DB-verified the `organization_assignments` row landed (Meridian Labs skill_path active 6→7).
- [x] **Delivers** — `corpus/ops/demo/playthroughs.md` (15 → **16 live Playthroughs, 0 TODO**) +
  `README.md` + `CLAUDE.md` count updates + the assign-WRITE subsection.

## Supporting
- [x] `url-shapes.ts` `ASSIGNMENTS_URL` + `isOnAssignments` + unit-test pin (73 TS unit tests green).
- [x] `manifest/corpus_test.go` M204 pin reversed (UC1 TODO → implemented `pt-assignment-assign`).
- [x] Gates green: ptvalidate static + both-way integrity + precondition-coverage + datadna closure PASS + Go tests.
