**Type:** tik · **Shape:** standard (protocol: `corpus/ops/demo/playthroughs.md` § *The iteration protocol*)

# iter-30 — `individual.UC1`, priced from measurement

## Phase A — can an org-less user reach the app? **YES, measured**

The cheap way, and deliberately not by writing a seeder first: **delete a seeded hero's membership on
`demo-2`** (Ola / `pt-ai-onboard`, backed up to TSV, restored by `--reset`) and drive her.

| reading | result |
|---|---|
| login | **succeeds** |
| `/onboarding` | **serves the flow** — `#linkedinUrl` present, 5 buttons, `bodyLen` 392, **0 page errors** |
| `/home` | **renders** — 14 buttons |

So a user with **no membership row at all** logs in, is served first-run onboarding, and reaches a usable home.
`individual.UC1` is **landable**, and the audit's F5 — *"no pre-onboarding state exists and none can be
declared"* — is refuted for a third time, now on its last remaining foothold. (It is also consistent with
`UserStatusContext`: its hiring eject requires `memberships.length > 0`, so an org-less member is not ejected.)

## Phase B — what it costs, measured from the schema

The blocker read *"a seeder change"*, singular. **`memberships` has 12 incoming foreign keys across 8 tables**,
and for ONE hero the live dependent-row count is:

| table | rows for one hero |
|---|---|
| `public.membership_skills` | **7** |
| `public.membership_languages` | **2** |
| `public.membership_tags` | **1** |
| `public.organization_target_roles` | **1** |
| `organization_features` / `_assignments` / `_roles` / `sim_invitation_links` | 0 (for this hero) |

Attempting the delete without cascading fails loudly — `violates foreign key constraint
"membership_skills_memberships_membership"` — which is the whole cost in one error message: **an org-less hero
must be skipped by `UsersSeeder` (the membership row + its casbin grants) AND by every seeder that keys off a
membership id.** Four of those write rows for a hero today; the other four are zero *for this hero* and must
still be audited, because "zero for the hero I measured" is not "zero by construction".

So the honest price is **≥ 5 seeders**, not one — and the residual risk is not FK breakage (which fails loudly)
but the *semantic* tail: the persona/profile/activity seeders write rows carrying an `organization_id` for a user
who now has no membership. Those do not FK-break; they would simply be quietly wrong. **That tail is the real
work, and it is what "a seeder change" was hiding.**

## Close — 2026-07-30

**Outcome:** the last gap in the exit gate is priced from measurement instead of estimate — **landable**, and
**≥ 5 seeders** rather than the "a seeder change" carried as prose since iter-08. No gate clause moves; nothing
was built on an unmeasured premise.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — onboarding 4 of 5 — (2) triggered-tok: n — (3) re-scope: n (the measured
answer is LANDABLE, so the trigger moves further away, not closer) — (4) user-blocker: n — (5) cap-reached:
**y** — 5 tiks this session (26, 27, 28, 29, 30) — (6) protocol-stop: n — Outcome: **exit-5**
**Decisions:** D116 (landable, and the price is ≥ 5 seeders — see `decisions.md`)
**Side-deliverables:** none.
**Routes carried forward:**
- **`ONBOARD-M256-orgless-seat`** → the build. `Persona.OrgMembership: none` (or equivalent) must make
  `UsersSeeder` skip the membership row + its casbin grants, and the **four** membership-keyed seeders
  (`membership_skills`, `membership_languages`, `membership_tags`, `organization_target_roles`) skip her too;
  the other four membership FKs write nothing for a hero today but must be audited rather than assumed. Then the
  Playthrough is small: she is served `/onboarding` with no org context, picks a role, finishes, and reaches
  `/home` — all four surfaces already measured live in Phase A.
- `ONBOARD-M256-prepared-persistence` (iter-29) and `standard.UC1` (the CV-upload product defect) unchanged.
**Lessons:**
1. **A prose blocker is an estimate until someone tests it.** *"Needs a seeder change (a member-less user + a
   roster seat)"* had been carried since iter-08 and was wrong in both directions: the load-bearing question
   (*can the app serve an org-less user?*) had never been asked, and the cost was understated ~5×. Three of this
   session's five iters found a routed blocker mis-stated in exactly this way (iter-26's "stage-0 seat",
   iter-28's "trigger not identified", iter-30's "a seeder change"). **A routed blocker should carry the
   measurement that produced it, or be marked as an estimate.**
2. **Measure a capability's cost by deleting, not by building.** One `DELETE` and one FK error priced a
   multi-seeder change in five minutes, and the FK error named the first four consumers itself. Writing the
   seeder to find out would have cost an hour and produced the same list.
3. **An FK constraint is a friendly failure.** The dangerous part of this change is not what breaks loudly but
   the org-scoped rows that would stay behind for a user with no org — the same silence class as iter-28's
   skipped insert.
