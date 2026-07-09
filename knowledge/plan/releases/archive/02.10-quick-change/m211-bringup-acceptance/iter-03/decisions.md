# iter-03 — Decisions (tik under TOK-01)

## D1 — Closure GREEN validates the M209 seed-side re-ground on a live merged stack
`datadna measure-closure` PASS proves the re-pointed `public.*` taxonomy resolvers (M209's
`stack-seeding` skiller→public re-ground) draw ONLY real public node-ids: 0 dangling across
user_skills / user_skill_evidences / jobsimulation.validation_attempt_skill_results / membership_skills
vs `public.skills.node_id`, after a real 45,710-row stories-maya seed against the merged warm stack. This
is the seed-side counterpart to iter-02's snapshot-side proof — both halves of the re-ground now verified
live, not just unit-tested.

## D2 — The casbin/sentinel gap is routed forward (Fate-3), not folded into iter-03
The seed's identity + users casbin-grant steps failed because the `sentinel` schema + `casbin_rules` table
do not exist on the warm stack (M208's de-risk brought up containers but did not complete `make migrate`;
the sentinel container is up but its policy schema never materialized). This is a genuinely separate concern
from iter-03's closure target — it is the authz-policy / M18 silent-403 class (sub-condition (d)'s
`casbin_rules > 0` assert), not a skill-ref-resolution problem. Per the scope-creep tripwire, iter-03 lands
its met target (closure GREEN) and routes the casbin fix to iter-04 with a named handler
(`bringup-M211-iter04-casbin-migrate`) rather than opening a 2nd fix line mid-iter. iter-04's fix path: run
`make migrate` (chartered, idempotent) to create sentinel + casbin_rules + load the policy, re-run the
casbin-grant seeders, and assert `casbin_rules > 0`. Escalate to a rext hook only if migrate reveals a
merged-platform ordering defect (M25-D9 class).

## D3 — Seeded into warm N=0 (sanctioned inner-loop), lightest verified-skill preset
Used `stories-maya` (single hero, ~150 users) over the full `stories` (~1080 users) to exercise the
verified-skill closure chain with minimal footprint. Targeted the warm N=0 stack via `--stack anthropos`
(the ParseStackN N=0 token; empty `--stack` does NOT override the preset's stack field). The orchestrator
explicitly sanctioned seed+verify on the warm stack as the fast inner loop; the tenant baseline was empty
(0 orgs/users) so the seed is a clean additive dataset, reversible via `stackseed --reset --force` later.
