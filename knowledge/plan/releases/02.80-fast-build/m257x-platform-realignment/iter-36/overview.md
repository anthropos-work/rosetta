---
milestone: M257x
iter: 36
iteration_type: tik
status: in-progress
opened: 2026-08-02
---

# iter-36 — `FIX-M257x-iter32-hiring-candidate-sim-link`

**Active strategy:** `TOK-01` (*instrument first, then follow — derive every list from the platform
instead of maintaining it by hand*). This iter is the "follow" half applied to a **read path** rather
than to a schema: the platform moved the surface, and rext still writes the shape the old surface read.

## Step 0 — re-survey (done before any code)

The hand-off's own next-measurement was *"is sim `00e80740-…` present in the demo's own Directus?"*, with
the prediction that an **absent** sim would make this a set-dressing defect. **Measured, and REFUTED:**

    directus.simulations  →  id 00e80740-3857-4116-86da-e25c5ef0c736
                             slug  project-manager-review-the-procurement-process-in-food-company-d19
                             type  SIMULATION_TYPE_HIRING     status  published

The sim is present, published, hiring-typed, and its slug is exactly what a `/sim/<slug>` link needs. The
assignment row is equally coherent — `public.organization_assignments` has one row for
`Kestrel Hiring Group`, `status=active`, `resource_type=job_simulation`, and its `assignee_id`
`8b5b861d-…` is **Ivo Kalman's MEMBERSHIP id** (`org_assignment.go` edges `assignee` → `Membership.Type`),
which is correct, not a user/membership mix-up.

**So neither end of the link is missing.** The target is still meaningful; the hypothesis is substituted.

## Cluster / target identified

The read path. `assignment_plans` / `_plan_items` / `_plan_enrollments` are empty platform-wide (the
hand-off measured this and treated it as *"don't reason about the plan model"*). That is exactly backwards:
**the plan model being empty is the defect.**

## Hypothesis (two conjuncts, both measured before any code was written)

1. **The surface the hiring home renders assigned positions on is plan-materialized-only.**
   `apps/hiring/.../HomeLeftContent.tsx` renders `MemberProgramsHomeContainer` under the comment
   *"the M7 plan-based program cards (replaces the legacy hiring assignment cards)"*; it calls
   `myAssignmentPrograms`, whose repository query
   (`app/internal/data/ent/repository/assignment_plans.go:1021 MemberProgramRows`) requires
   **`organizationassignment.PlanIDNotNil()`**. Measured on demo-1: **0 of 72** seeded assignments carry a
   `plan_id`. So the container returns `null` and the home renders its Empty state.

2. **The affordance is no longer an anchor.** `packages/ui/src/AssignmentsHome/AssignmentsHome.tsx`
   contains **zero** `href` occurrences (`grep`, exit 0 on a positive control in the same dir — `types.ts`
   has two, as a *data* field). The card opens via `onClick` → `router.push(href)`. So
   `main a[href*="/sim/"][href*="organizationId="]` — the page object's locator — is **architecturally
   unsatisfiable at origin HEAD**, seeded or not.

Dated: `next-web-app d4bb7c6c9` (2026-07-07) is the commit that swapped `HomeAssignmentCard` for
`MemberProgramsHomeContainer`. The page object's own header records the opposite shape as **MEASURED** on
demo-2 at iter-27 — it was true then.

## Planned scope — a declared TWO-step shape (not scope creep)

1. **Seed side.** Plan-materialize org assignments: emit `assignment_plans` + `assignment_cycles` +
   `assignment_plan_items` + `assignment_plan_enrollments` and stamp `plan_id` / `cycle_id` /
   `plan_item_id` / `enrollment_id` onto the assignment rows. **Both** writers
   (`seeders/assignments.go` and `seeders/hiring_funnel.go`) — a one-writer fix is the half-done re-point
   §7 warns about.
2. **Spec side.** Re-express the affordance in `hiring-home-page.ts` + the spec against what the product
   now renders, and prove the org-scoping by **navigating** (click → URL carries
   `/sim/<slug>?organizationId=<her org>`) rather than by reading an attribute. That is a strictly stronger
   claim than the href assert it replaces.

## Expected lift

Clause 2 `28 / 2 / 1` → **`29 / 1 / 1`** — pre-registered before any confirming run exists, per the
iters 31–35 discipline. A binding full `--reset` run is the only acceptable evidence.

Blast radius is wide by construction (every seeded org's members gain a program card), so the run must be
**unscoped**; `pt-assignment-assign` asserts an affordance COUNT on a related surface and is the most
likely place for an unintended addition to show up. **Zero additions** is the load-bearing half of the diff.

## Escalation conditions

- If plan-materializing needs a platform edit of any kind → STOP, escalate (binding constraint).
- If the binding run shows additions, the iter closes on the diff, not on the headline.

## Acceptable close-no-lift outcomes

A measured demonstration that the M7 surface cannot be driven from seed data alone (e.g. a publish-time
side effect in `app` that no COPY can reproduce) closes this as `closed-no-lift` with the falsification
recorded — that is a complete iter under this protocol.

## PRE-REGISTERED (written before the binding run existed)

**Expected: `29 live / 1 failing / 1 unimplemented`** — `pt-onboarding-hiring-candidate` flips;
`pt-orgadmin-role-create` (untouched) does not. **Zero additions** is the load-bearing half.
