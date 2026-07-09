---
iter: 03
milestone: M211
iteration_type: tik
status: closed-fixed
created: 2026-07-08
---

# iter-03 — tik: seed closure green (re-pointed taxonomy resolvers resolve real public node-ids)

**Type:** tik — under **TOK-01**.

## Step 0 — Re-survey
Warm stack: taxonomy loaded (iter-02), tenant baseline EMPTY (0 orgs/users/skills). TOK-01's next-tik target
(sub-condition (c) seed closure) is the right open next thing.

## Active strategy reference
**TOK-01** — the warm inner-loop step after replay: seed + measure closure.

## Cluster / target
Sub-condition **(c)**: `datadna` seed-side verified-skill closure GREEN — every seeded skill node-id
(user_skills, user_skill_evidences, jobsimulation.validation_attempt_skill_results, membership_skills)
resolves in the replayed `public.skills.node_id`. This is the merge-correctness proof for the SEED side:
the re-pointed `public.*` taxonomy resolvers draw only real public node-ids.

## Hypothesis
With the real public taxonomy loaded (iter-02), seeding a verified-skill hero (stories-maya) via the
re-grounded seeder → the TaxonomyRefs resolver picks real public node-ids → 0 dangling → closure PASS.

## Expected lift
Prove sub-condition (c)'s core (closure GREEN).

## Phase plan (executed)
1. Build stackseed + datadna from the re-grounded authoring copy.
2. Dry-run stories-maya (validate DAG + isolation vs merged schema, no writes).
3. Real seed stories-maya into warm N=0 (`--stack anthropos`, port 5432, prod=false).
4. `datadna measure-closure` → PASS.
5. Triage any seed failures → route.

## Escalation conditions
- Dangling refs (closure FAIL) → the taxonomy resolver is mis-pointed → route a rext fix. **Did NOT fire.**

## Outcome
**Target MET — closure GREEN.** `datadna measure-closure` → `[PASS] seed-verified-skill-closure: every
seeded verified-skill node-id resolves in the replayed taxonomy`. Seed ran 45,710 rows (org/users/personas/
profiles/membership_skills/population-evidence — the verified-skill fan-out landed), isolation clean,
prod=false. **Surfaced (routed to iter-04):** 2 seeders (identity, users) failed casbin-grant because
`sentinel.casbin_rules` (and the whole `sentinel` schema) is ABSENT — `make migrate` was incomplete on the
warm stack (M208 brought up containers only). This is the M18 silent-403 / sub-condition (d) casbin>0 class,
not a closure problem.
