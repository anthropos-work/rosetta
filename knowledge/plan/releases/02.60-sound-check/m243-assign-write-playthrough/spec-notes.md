# M243 — Spec notes

Topic → doc → code triples + assign-WRITE Playthrough findings accumulate here during build.

## The assign-and-track UC1
- `playthroughs/manifest/assignment-monitoring.yaml` UC1 (`assign-and-track.UC1`, currently `TODO`).
- A new `/enterprise/assignments` page object.
- The spec `e2e/tests/assignment-assign.spec.ts` tagged `@pt:...UC1`.

## pt-world precondition
- Possibly assignable content + a target member in lockstep with `seed-worlds.yaml`.

## Carry closure
- Realizes reserved `M238`; closes `DEF-M235-03` / M204 assign-WRITE (~10 routings across 5 releases). Takes the corpus 15 → 16 live Playthroughs, 0 TODO.

## Pre-flight audits — assign-and-track.UC1 (first section)
- **Verdict: GREEN.** Report: `kb-fidelity-audit.md`. sha at audit: `e9609cd` (pre-flight checkpoint / branch HEAD).
- Topic → doc → code triples:
  - Playthroughs pillar → `corpus/ops/demo/playthroughs.md` → `rext playthroughs/` (manifest/ · e2e/ · seed/ · report/). PAIRED, aligned.
  - assign-WRITE UC1 → `playthroughs.md:105-107` + manifest → `playthroughs/manifest/assignment-monitoring.yaml:34-50` (`playthrough: TODO`). PAIRED, declared.
  - pt-world seed → `playthroughs.md` §dedicated-seed + `seeding-spec.md` → `seed/pt-world.seed.yaml`, `seed/seed-worlds.yaml`. PAIRED, single-sourced.
  - page-object/locator layer → `playthroughs.md` §page-object → `e2e/lib/*.ts`, `url-shapes.ts`. PAIRED.
  - `/enterprise/assignments` platform surface → CODE-ONLY (discovered live per P3; not a corpus topic).
- Verified counts: 15 live spec files, 15 `@pt:` tags, 1 `playthrough: TODO` (UC1), 16 total declared use cases — exact match to the doc's "15 live, 1 TODO".

## Build outcome — LIVE GREEN + DB-verified (2026-07-22)
- **Live run:** `pt-assignment-assign` GREEN on demo-1 (7.9s). Manager (Morgan/pt-manager) opened
  `/enterprise/assignments` → Skill Paths tab → assigned a skill path to an unassigned member with a
  deadline → the assignable-affordance count dropped by exactly one (read-back).
- **DB-verified LANDS (anti-toothlessness bar):** Meridian Labs `public.organization_assignments`
  skill_path/active went **6 → 7**; newest row = **Omar Becker (assignee) ← Morgan Reyes (assigner),
  due 2026-08-05** — the exact assign the Playthrough performed. A real row written by the manager,
  read back through the UI.
- **Gates:** ptvalidate static (16 UCs, 16 live, 0 TODO) + both-way integrity (16 ids ↔ 16 UCs) +
  precondition-coverage (`org-unassigned-member` resolves) + **datadna closure PASS** + Go tests (all
  packages) + 73 TS unit tests + typecheck — all green.

## Topic → doc → code triples (final)
- assign-WRITE UC1 → `playthroughs.md` (count 15→16, 0 TODO + the assign-WRITE subsection + the M243
  page-object bullet) → `manifest/assignment-monitoring.yaml` UC1 (`playthrough: pt-assignment-assign`,
  `preconditions: [public-catalog, org-unassigned-member]`).
- precondition (unassigned target) → `seed-worlds.yaml` capability `org-unassigned-member` (lockstep) →
  materialized by the existing pt-world Org A roster (no seed CODE change; 34/40 unassigned at capture).
- assign surface route → `url-shapes.ts` `ASSIGNMENTS_URL` + `isOnAssignments` (+ unit-test pin) →
  `/enterprise/assignments` (segment-anchored, matches the sub-tabs).
- assign page object → `e2e/lib/assignments-page.ts` (`AssignmentsPage`) → next-web
  `EnterpriseAssignmentsTable` / `memberSkillPathColumn` / `AssignmentContainer` / `AssignmentModal`.
- write backend → (read-only investigation) `app` `mutationResolver.CreateOrganizationAssignments` →
  `AssignmentManager.BulkCreateOrganizationAssignments` → `public.organization_assignments`.
- M204 pin reversal → `manifest/corpus_test.go` `TestRealCorpus_ManagerCoverageIsPresent` (UC1 now a
  non-TODO `pt-assignment-assign`, added to `wantManagerPTs`).

## antd-v6 Select lesson (page-object layer)
`getByRole('option')` on an antd `rc-virtual-list` Select is unreliable — the `role="option"` nodes
carry the raw VALUE (uuid) as accessible name AND read as non-visible (visible title/image are separate
children). Commit the first option by **keyboard** (`ArrowDown`+`Enter`) — robust + user-driven (P1).

## Dev-run note (not a code change)
demo-1 was up on the demo SHOWCASE roster → `pt-manager` login 400s `unknown_identity`. Fixed by
re-exporting the pt-world roster to the fake-FAPI mount + restarting fake services (what
`run-playthroughs.sh --reset` does, M211 iter-16). The canonical gate is `run-playthroughs.sh N --reset`.
