# iter-30 — decisions

## D116 — `individual.UC1` is LANDABLE, and its price is ≥ 5 seeders, not "a seeder change"

**Measured, not estimated.** Delete a seeded hero's membership on `demo-2` and drive her: she logs in,
`/onboarding` **serves the flow** (`#linkedinUrl` present, 5 buttons, `bodyLen` 392, 0 page errors), and `/home`
**renders** (14 buttons). A user with **no membership row at all** reaches a usable app. So the UC is **not**
`unimplementable-without-platform-edit`, the re-scope trigger moves further away rather than closer, and the
Phase-0b audit's F5 claim is refuted on its last foothold.

**And the cost was understated by roughly five times.** `memberships` carries **12 incoming FKs across 8
tables**; the un-cascaded delete fails with `violates foreign key constraint
"membership_skills_memberships_membership"`, and one hero's live dependents are `membership_skills` **7**,
`membership_languages` **2**, `membership_tags` **1**, `organization_target_roles` **1**. So an org-less hero must
be skipped by `UsersSeeder` **and** by every membership-keyed seeder — ≥ 5, with four more FK paths to audit
rather than assume.

**The residual risk is the quiet half, and it is worth stating separately from the loud half.** FK breakage
fails loudly and names its own consumer. What will not fail is the persona / profile / activity / evidence
seeders writing rows that carry an `organization_id` for a user who now has no membership: those are simply
wrong, invisibly — the same silence class as iter-28's skipped insert (D112). **That tail is the work; the FK
list is just the entry fee.**

## D117 — three of this session's five iters found a routed blocker MIS-STATED, and that is a pattern worth a rule

| iter | the routed blocker said | what measurement said |
|---|---|---|
| **26** | *"needs an Org C stage-0 seat"* | the seat **could not be declared at all**; an appended hero arrived COMPLETED |
| **28** | *"the trigger is not yet identified"* | one `useState`, and the four-org sweep that failed to find it was **sampling a constant** |
| **30** | *"needs a seeder change (a member-less user + a roster seat)"* | **landable** (never tested), and **≥ 5 seeders** |

Each of these was written in good faith by an iter that had done real work, and each read as a *cost* when the
real unknown was a *feasibility* — or vice versa. The failure mode is uniform: **a blocker written from reasoning
gets read by the next iter as a measurement.** iter-26's was the dangerous one, because building on it would
have produced a passing Playthrough over a hero who had already done the thing.

**The rule: a routed blocker should carry the measurement that produced it, or be explicitly marked an
estimate.** Concretely — name the file:line, the query, or the probe that established it; if none exists, write
*"(estimate, unmeasured)"*. The routing table entries this session's iters added all do this, and it is cheap:
one clause per entry. The alternative is what happened here three times in five iters.
