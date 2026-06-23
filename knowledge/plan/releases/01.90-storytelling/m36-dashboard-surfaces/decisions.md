# M36 — Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions live in the spec
([`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md), esp. D4 the gap invariant).

## D-M36-1 — The funnel matches mapped↔verified by skill NAME (not node-id)
The dashboard's verification funnel (`app/internal/workforce/skills.go::aggregateSkills`) joins the mapped side
(`membership_skills` GROUP BY `skill_name`) to the verified side (`user_skill_evidences` → `skiller.skills` BY
`sk.name`) **on the skill NAME**, not the node-id. So `membership_skills.skill_name` MUST equal the verified
skills' `skiller.skills.name` for a skill to line up across the funnel. The new `skillref_named.go` resolver
reads name-bearing skills from the SAME replayed taxonomy the verified chain draws from, so the names match by
construction. (Recorded because it's non-obvious — one would expect a node-id join.)

## D-M36-2 — The dashboard reads `public.local_skill_path_sessions`, NOT `skillpath.skill_path_sessions`
The completed-skill-paths surface (`skill_paths.go::ListSkillPaths`) + the assignment in-progress/completed
buckets (`assignments.go::assignmentsSummary`, via `organization_assignment_sessions`) read the **app-mirror**
table `public.local_skill_path_sessions` — NOT the learning-session table `skillpath.skill_path_sessions` the
`SkillpathSessionsSeeder` writes. The two are distinct surfaces:
- `skillpath.skill_path_sessions` = the learning sessions (my-growth / learning surfaces). M36 bumped its
  `completed` share from ~1% (only an exact progress=100) to ~30% so those surfaces read believably.
- `public.local_skill_path_sessions` = the app mirror the org dashboard reads. M36's **assignments fix** writes
  these (for completed/in-progress assignments) so the `organization_assignment_sessions` FK resolves and the
  assignment progress + the completed-paths surface render.
This is why the assignments fix writes three tables (assignments + local_skill_path_sessions + assignment_
sessions), and why the skillpath-completed-share fix is a SEPARATE surface from the assignment status mix.

## D-M36-3 — org_assignment_sessions take the skill-path FK arm (the population has no jobsim local mirror)
`organization_assignment_sessions` requires ONE of `session_id` (→`local_skill_path_sessions`) or
`js_session_id` (→`local_jobsimulation_sessions`) non-null (a CHECK constraint). The general population's
jobsim sessions are written only to `jobsimulation.sessions` (the `local_jobsimulation_sessions` mirror is
written ONLY by the PersonaSeeder, for heroes). So a population member's assignment-session must take the
**skill-path arm** (`session_id` → a `local_skill_path_sessions` row the assignments fix writes), not the
jobsim arm. This keeps the in-progress/completed assignment buckets non-empty for the whole population, not just
the 6 heroes.

## D-M36-4 — target_roles feed the GraphQL mobility surface, role-readiness reads membership_skills
The REST role-readiness query (`members.go::GetRoleReadiness`) reads `membership_skills` + `user_skill_evidences`
+ sessions + `membership_tags` — NOT the `*_target_roles` tables. The `organization_target_roles` /
`user_target_roles` tables feed the org's **role-targeting / two-way mobility** surface (the GraphQL
role-management resolvers + the gap visualization). Both are seeded (overview-scoped) so the mobility/gap surface
is non-empty; role-readiness itself is fed by the mapped+verified skills the other M36 seeders land.
