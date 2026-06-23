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

## Adversarial review (close, 2026-06-23)

Three non-obvious failure scenarios probed at close. All degrade safely or are story-authoring contracts —
no crash, no corruption, no closure-gene break. Recorded as *scenarios* (not fixes) per the close-milestone
Phase 2c contract.

- **AR-1 — Tiny-story succession gate.** `succession.go::SuccessionSeeder` clears the >20% interview-share
  gate (heroes always interviewed + `interviewShare=0.30` of the population), but `succession.go::buildCoverage`
  reports `too_sparse` whenever the org has `total < 15` members — independent of interview coverage. So a
  story sized below 15 would render a Succession CTA even with interviews seeded. **Disposition:** not a seeder
  bug — the spec mandates each story org be sized ≥120 (spec-notes.md §"Succession coverage gate"), so the
  `too_sparse` floor is never the binding constraint for a real story. The seeder degrades to a believable CTA
  (not a crash) if mis-sized. No fix.
- **AR-2 — Feedback 2:1 ratio on a near-empty session population.** `feedback.go` realizes the 2:1 mix via a
  per-member `hash(...)%3 != 2` — a *statistical* 2:1 that only reads as ~2:1 at population scale; on a 1–2
  member story it could be all-positive or all-negative. **Disposition:** safe — the `len(rows)==0` guard
  prevents an empty COPY, the ratio is documented as believable-at-scale, and real stories are sized ≥120.
  No fix.
- **AR-3 — membership_skills with a roles-only / empty-flat replayed taxonomy.** The mapped-funnel seeder
  guards `!refs.flat.available() && len(refs.byRole)==0 → return 0, nil`. If the replay yields role-keyed
  skills but an empty flat pool, the AI-readiness top-up (which draws from `filterAISkills(refs.flat)`) is
  skipped via the `aiPool.available()` check — the funnel still seeds from `byRole`, the AI-readiness surface
  simply shows fewer rows. **Disposition:** safe degrade (the no-fabrication invariant holds — never invents a
  node-id). Confirmed by the empty-pool branch tests in `distribution_helpers_harden_test.go`. No fix.
