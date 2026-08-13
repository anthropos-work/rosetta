# iter-36 — decisions

## D-M257x-36-1: the plan model is a VISIBILITY contract, not an optional embellishment

**Choice.** Every seeder that writes `public.organization_assignments` also writes the four-table M7 plan
model and stamps `plan_id` / `cycle_id` / `plan_item_id` / `enrollment_id`. The alternative — leave the
flat rows and change the Playthrough to assert something else — was rejected.

**Why.** `MemberProgramRows` filters on `plan_id IS NOT NULL`, and it is the query behind "Your
assignments" on **both** member homes. A flat-only seed does not produce a *smaller* demo; it produces one
where **every seeded member's assignments are invisible on their own home**, in every org, while the
assignments table looks perfectly healthy. The Playthrough was the only thing that noticed.

**What this decision does NOT claim.** It does not claim the platform has abolished flat assignments.
Measured on the same run: `pt-assignment-assign` writes a plan-less row through the product UI and the
product accepts it. Direct manager assignment still works flat; only the *home program surface* requires
the plan. Recorded because the over-generalisation would send a future iter hunting a bug in working code.

## D-M257x-36-2: prove org-scoping by NAVIGATING, not by reading an attribute

**Choice.** `startAssignedPosition()` clicks the affordance and returns the landed URL; the spec asserts
`/sim/<slug>` and an `organizationId` uuid on that URL.

**Why.** The old locator read `href` off an anchor. `AssignmentsHome.tsx` has **no** `href` — the card
opens via `onClick` → `router.push` — so there is nothing to read. But the replacement is not a
concession: an attribute can be correct on a control that is dead, and this milestone has already paid for
that distinction twice (iter-10's `/library/` probe over a 500ing `/`; iter-17's REGISTERED-is-not-SERVED).
A navigation cannot be satisfied by a well-formed string.

**Cost, stated.** The assertion now mutates browser state (it leaves the home). It is the LAST assertion in
the spec, deliberately, so nothing downstream inherits the navigation.

## D-M257x-36-3: the seeded cycle is UNORDERED

**Choice.** `assignment_cycles.ordered = false`, `date_mode = 'absolute'` with `start_at`/`end_at` NULL.

**Why.** `buildMemberProgram` locks a step when `r.Edges.Cycle.Ordered && priorIncomplete`, and when
`r.WindowStart.After(now)`. No seeded member has completed any step, so an *ordered* cycle would have
produced a structurally perfect plan model in which **every program is unstartable past step 1** — a fix
that reads correct in the database and fails at the surface. Pinned by a test that names the reason.

## D-M257x-36-4: the fence asserts over produced ROWS, not over source or a writer list

**Choice.** `assertPlanMaterialized` runs a seeder against the recording `Conn` and checks that every
assignment row carries four non-nil FKs which resolve to rows the same run wrote.

**Why.** The obvious fence — "every writer must call `attach`" — is a hand-maintained list of the system's
parts, which §8 rule 1 identifies as the worst possible place for one: a writer that is not enrolled cannot
go RED. A row-level fence is writer-agnostic by construction. It is vacuity-guarded (zero assignments is a
`t.Fatal`, per §8 rule 6) and its column positions are **derived from `assignmentCols()`** so a reorder
moves the fence instead of pointing it at the wrong cells.
