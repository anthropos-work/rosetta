**Type:** tik

# iter-36 — the assignment was never missing; the surface that reads it moved

## What was measured, in the order it was measured

**Step 0 refuted the hand-off's own next-measurement.** The prep note ended with *"is sim
`00e80740-…` present in the demo's own Directus? If it is absent, the seeder is pinning a `resource_id`
the demo's content set does not contain."* It is **present**, `published`, `SIMULATION_TYPE_HIRING`, with
the slug `project-manager-review-the-procurement-process-in-food-company-d19`. The assignment row is
equally sound, and its `assignee_id` is Ivo Kalman's **membership** id — which is what
`org_assignment.go`'s `edge.To("assignee", Membership.Type)` requires, so the one thing that *looked* like
a defect (a user id that resolves to no user) is the schema working correctly.

**Both ends of the link were therefore intact.** What was broken sat between them.

## The mechanism, in two halves — and the second half means the first alone would not have helped

### 1. The home reads a table the seeder never filled

`apps/hiring/.../HomeLeftContent.tsx` renders `MemberProgramsHomeContainer` under its own comment:

> *"Your assignments" — the M7 plan-based program cards (**replaces the legacy hiring assignment cards**).*

That container queries `myAssignmentPrograms`, and the repository behind it —
`app/internal/data/ent/repository/assignment_plans.go:1021 MemberProgramRows` — filters on

    organizationassignment.PlanIDNotNil()

Measured on demo-1: **0 of 72** seeded assignments carried a `plan_id`, and
`assignment_plans` / `_cycles` / `_plan_items` / `_plan_enrollments` were **empty platform-wide**. The
hand-off had measured that same emptiness and drawn the opposite conclusion — *"anything reasoning about
the plan model on this stack is reasoning about empty tables"*. The empty tables **were** the defect.

**This is the milestone's founding class arriving through a READ PATH rather than a schema.** Nothing
errors, nothing `42P01`s, the row is valid and complete; it is simply no longer the shape anything reads.
The closest prior relative is iter-23/24's stale *service name* in a consumer list — there the stale thing
was a name that still resolved, here it is a table whose ROLE changed underneath it.

### 2. The affordance stopped being an anchor

`packages/ui/src/AssignmentsHome/AssignmentsHome.tsx` contains **zero** `href` occurrences (grep with a
positive control in the same directory: `types.ts` has two, as a *data* field). The card opens through
`onClick` → `router.push(href)`. So the page object's locator

    main a[href*="/sim/"][href*="organizationId="]

is **unsatisfiable at origin HEAD no matter what is seeded**. Had only half 1 landed, the seed would have
been right and the Playthrough would have failed with the identical message — which is the shape that
produces a second iteration spent re-deriving the same finding.

**Dated:** `next-web-app d4bb7c6c9` (2026-07-07) is the commit that swapped `HomeAssignmentCard` for
`MemberProgramsHomeContainer`. The page object's own header records the anchor shape as **MEASURED** on
demo-2 at iter-27 — it was true when written, and the header now says so rather than being quietly
overwritten.

## What landed

**Seed side.** A shared `planMaterializer` **derives** the plan model from the assignments themselves —
one plan item per distinct resource, one enrollment per member — rather than enumerating anything, so a
new assignment shape cannot leave the plan model behind (§2, at the point of use). Both writers
(`assignments.go`, `hiring_funnel.go`) stamp the four FKs; a one-writer fix is the half-done re-point §7
warns about. Two design decisions that are load-bearing rather than incidental:

- **The cycle is UNORDERED.** `buildMemberProgram` locks a step behind an incomplete prior one *only* when
  its cycle is `ordered`, and no seeded member has completed any step. An ordered cycle would have written
  a correct plan model that renders every seeded program **unstartable past step 1** — a fix that looks
  right in the database and fails at the surface (§5 rule 14, one layer down).
- **The enrollment is per MEMBER, not per assignment.** `GetMyAssignmentPrograms` groups by enrollment;
  one enrollment per assignment would split a member's card into N single-step cards. Pinned by a test
  that says exactly that.

`--reset` gains the four tables, positioned **after** `organization_assignments` (whose four FKs are
`ON DELETE SET NULL`, so truncating the parents first would silently un-materialize any surviving child)
and **before** `memberships`/`organizations`.

**Spec side.** `startAssignedPosition()` clicks the affordance and returns the landed URL; the spec
asserts `/sim/<slug>` **and** an `organizationId` uuid on it. That is strictly stronger than the href it
replaces — an attribute can be correct on a control that does nothing — and the negative control keeps its
shape (the same accessor, asserted absent mid-flow, on the same journey).

## The fence

It asserts against **the rows a seeder produced**, not against source and not against a list of writers:
every row copied into `public.organization_assignments` must carry four non-nil plan FKs, and each id must
resolve to a row the **same run** wrote into its own table. A future writer that forgets the plan model
fails this without anyone remembering to enrol it — the property a hand-maintained writer list can never
have (§8 rule 1). It is vacuity-guarded: zero assignments is a `t.Fatal`, not a pass (§8 rule 6). The FK
column positions are **derived from `assignmentCols()`**, so a column reorder moves the fence rather than
silently checking the wrong cells.

## Mutation verification — 8 mutants, 8/8 matching declared expectation

7 declared-RED (all killed) + 1 declared-GREEN no-op control (survived). Every mutant `go build`-gated;
control green before and after; each mutant applied to a pristine copy and restored.

| mutant | declared | actual |
|---|---|---|
| M1 NULL the four FKs in `assignments.go` | RED | RED |
| M2 NULL the four FKs in `hiring_funnel.go` | RED | RED |
| M3 the seeded cycle becomes `ordered` | RED | RED |
| M4 enrollment reuse dropped (one per assignment) | RED | RED |
| M5 **inverted** — `empty()` returns its own negation | RED | RED |
| M6 plan model written AFTER the assignments | RED | RED |
| M8 plan rows suppressed at write time (dangling FK) | RED | RED |
| M7 **no-op control** — rename the program card title | GREEN | GREEN |

**M1 and M2's first cut COMPILE-BROKE** (nulling the values left `fks` unused) and the harness reported
`COMPILE-BREAK (not a kill)` rather than counting them — iter-07's §8 rule 5. Re-run with `_ = fks` and
both killed. M7 is what makes the seven REDs mean something: the fence does not depend on the seed's
content, only on its structure.

## Live proof — and the row that is correctly NOT materialized

Binding full `--reset` run on `demo-1` from its own pinned clone (`fast-build-m257x-iter-36`, verified on
origin), platform origin `2adcf71` re-fetched at open and close (unchanged):

    Playthroughs coverage: 29/31 passing (93.5%)      209 specs, 208 passed / 1 failed, 2.7 min

**Exactly the figure pre-registered in this iter's `overview.md` before any confirming run existed.**
The measurement is the sorted-id diff, never the two summary lines (iter-19's rule), taken in the `pt-*`
space against iter-35's own artifact (re-read from `iter-35/binding-run.log.txt`, not from the hand-off):

    REMOVED (fixed):  pt-onboarding-hiring-candidate
    ADDED:            —  (empty)

**Zero additions is the load-bearing half**, because the blast radius was the widest available: every
seeded org's members gained a program card. `pt-assignment-assign`, which asserts an affordance COUNT on a
neighbouring surface and has already moved once un-attributed (iter-28), stayed green.

Plan model on the live stack after the reset:

| plan | org | members | steps |
|---|---|---|---|
| Development plan | Cervato Systems | 19 | 9 |
| Development plan | Halcyon Retail | 10 | 7 |
| **Candidate assessment** | **Kestrel Hiring Group** | **1** | **1** |
| Development plan | Meridian Labs | 21 | 8 |
| Development plan | Vertex Logistics | 20 | 10 |

**71 of 72 assignments are plan-materialized, and the 72nd is the interesting one.** It was created at
`23:43:52Z` — *during* the run — carries a random (not deterministic) uuid, `skill_path`, Meridian Labs:
it is the row `pt-assignment-assign` writes **through the product UI**. So the platform still writes
plan-less assignments for a direct manager assignment; what changed is only that the member **home's
program surface** cannot see them. Recorded because the opposite reading — *"the platform has abolished
flat assignments"* — is the over-generalisation this measurement forecloses, and it would have sent a
future iter looking for a bug in a working product.

## Suites

`stack-seeding` Go **green** (+6 tests). `stack-core` **14 failures of 396 — exactly baseline**, measured
after the change. `playthroughs/e2e` typechecks clean (`tsc --noEmit`, in the clone that has
`node_modules` — the authoring copy has none, and `npx tsc` there silently installs a *different* package
called `tsc` that prints a banner and exits 0, which would have read as a passing typecheck).

## Close — 2026-08-02

**Outcome:** gate clause 2 `28 / 2 / 1` → **`29 / 1 / 1`**, one removal and zero additions, on a binding
cold-reset run — the pre-registered figure exactly. Root cause was a platform read-path move, not a
seeding gap and not the content layer.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-36-1 (the plan model is the visibility contract, not an optional embellishment) ·
D-M257x-36-2 (prove org-scoping by navigating, not by reading an attribute) ·
D-M257x-36-3 (the unordered cycle) — see `decisions.md`
**Side-deliverables:** none — everything that landed was planned scope.
**Routes carried forward:**
- `FIX-M257x-iter32-orgadmin-role-create-timeout` — the last clause-2 failure, still undiagnosed.
- `CHECK-M257x-iter36-flat-assignment-surfaces` — the M7 program card is one reader; the M6
  `myAssignmentQueue` (`MemberQueueRows`) carries the *same* `PlanIDNotNil()` filter. Nothing measured
  whether any other seeded surface silently lost its rows to the same change. Cheap: grep the app for
  `PlanIDNotNil`.
- `DOC-M257x-iter36-plan-model` — `seeding-spec.md` describes assignments as a flat table; the plan model
  is now part of what a seeded org contains. Clause-5 adjacent.
**Lessons:**
- **A stale table is louder than a stale ROLE.** Three prior occurrences of this class failed at
  `42P01`. This one produced a valid row, a green seeder, a green autoverify and an empty page. When a
  surface is empty and the data is present, ask what the READER filters on before asking what the writer
  wrote — and read the *repository query*, not the resolver.
- **Fix both halves of a moved surface in one iteration or neither is measurable.** The seed fix alone
  and the locator fix alone would each have produced the identical failure message, and either would have
  read as "no progress."
- **A compile-break is not a kill** (§8 rule 5, third occurrence in this milestone) — and the harness
  should say so in its own voice rather than leaving the reader to notice a suspiciously fast RED.
