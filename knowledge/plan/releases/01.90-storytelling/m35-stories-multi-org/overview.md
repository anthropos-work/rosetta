---
milestone: M35
slug: stories-multi-org
version: v1.9 "storytelling"
milestone_shape: section
status: planned
created: 2026-06-22
last_updated: 2026-06-23
complexity: large
delivers: rosetta-extensions/stack-seeding (the stories[] blueprint + multi-org OrgID/orgClerkID parameterization across the 4 seeders + PersonaSeeder scaled to the roster + trajectory logic + supporting-population fidelity) + corpus (extend stories-spec.md with the model + seeding-spec.md blueprint update + /stack-seed SKILL.md)
depends_on: M34
spec_ref: .agentspace/seeding_gaps.md §4 (Stories & Heroes model, §4e multi-org, §4f Clerkenstein)
---

# M35 — Stories & Heroes model + multi-org

## Goal
One declarative **`stack.stories.yaml`** seeds **multiple orgs**, each with its **thriving / struggling /
manager** hero trio at vantage-appropriate fidelity. This is the model that makes the seeder a
demo-orchestration engine: the same YAML that seeds the data is (in M38) the cockpit's menu.

## Why section
The model is fully specified (`.agentspace/seeding_gaps.md` §4c gives the YAML schema; §4e the multi-org
refactor; §16 the vantage/trajectory axes + the coherence property). The deliverables are a fixed checklist.

## Scope
**In:**
- **The `stories[]` blueprint** — per-hero `vantage` (`end-user | manager`), `trajectory`
  (`thriving | struggling`), `skills`/`self_eval_bias`; per-story `org` + `narrative`. Supersedes the
  org-centric `stack.seed.yaml` for demo stacks (the existing presets stay for single-org dev seeds).
- **Multi-org parameterization** — `OrgID` + sibling `orgClerkID` become **per-story** (deterministic per
  story id), threaded through the **4** consuming seeders (`org`, `users`, `identity`, `jobsim-sessions`,
  `assignments`) + each story's injected Clerkenstein `DemoUser` org-claim (`OrgEid`) aligned.
- **Scale `PersonaSeeder`** from M34's one-hero proof to the full roster (the locked v1 roster: 2 stories ×
  3 heroes — Cervato Systems {Maya / Tom / Dan} + Solvantis {Sara / Nick / Leah}).
  - **(Routed from M34 close, #M34-D7):** the multi-hero roster makes two M34-benign edges reachable —
    add a **`len(Personas) <= Size` blueprint validation + an index-collision warning**
    (`personaUserIndex` can hash multiple heroes to the same population slot), and decide whether a
    **short-but-nonempty role pool tops up from the flat pool** to hit each hero's declared `verified: N`
    (M34 keeps the short role pool — role-coherence over count; M35 owns the roster-fidelity product call).
- **Trajectory logic** — thriving = dense/high/rising + under-claim; struggling = sparse/low/flat +
  over-claim (a stark gap). The two employee heroes are seeded as the manager's standout high/low rows (the
  coherence property — realized fully in M36's dashboard).
- **Supporting-population fidelity** — `job_role_id`+`job_role_name` (real `skiller.job_roles` that *have*
  `job_role_skills`), ramped `joined_at`, names on the non-hero members so the trio sits in a believable org.

**Out:** the org-aggregate dashboard surfaces (M36); the Clerkenstein seat-switch (M37) + cockpit (M38).

## Repo split
- **`rosetta-extensions`** `stack-seeding`: the blueprint schema, the multi-org refactor, the roster scaling.
- **`rosetta`** corpus: extend `stories-spec.md` (the model) + update `seeding-spec.md` (the blueprint
  supersession) + the `/stack-seed` SKILL.md (the stories input).

## Open questions
- **O6** — which replayed public `job_roles` actually *have* `job_role_skills` (the usable role set for
  role-coherent skills): a one-time `/db-query` enumeration against a replayed stack.

## Risk
The multi-org refactor touches 4 seeders + the Clerkenstein org-claim — regression risk on the existing
single-org path. **Mitigation:** keep a single-story (single-org) blueprint as the default; the existing
`dev-min`/preset path must keep working unchanged.

## Done-when
One `stories.yaml` seeds both orgs with their trios; each hero renders at its trajectory-appropriate fidelity;
the existing single-org seed path still passes; the closure gene stays green across all orgs.
