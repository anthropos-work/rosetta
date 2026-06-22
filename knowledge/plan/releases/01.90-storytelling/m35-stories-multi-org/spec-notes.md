# M35 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §4 (the
Stories & Heroes model), §4c (the `stack.stories.yaml` schema), §4e (multi-org), §16 (vantage/trajectory +
the coherence property).

## The blueprint shape
_(per §4c: `stories[]` → `{ org{name,industry,narrative,size}, heroes[]{ id,name,role,vantage,trajectory,
annotation,demonstrates,login,jump_to,skills } }`. Locked v1 roster: Cervato Systems {Maya/Tom/Dan} +
Solvantis {Sara/Nick/Leah}.)_

## Multi-org refactor (blast radius)
_(per §4e: `OrgID` const at `seeders/org.go:19` + sibling `orgClerkID` at `:22`; consumed by 4 seeders —
`users.go`, `identity.go`, `jobsim_sessions.go`, `assignments.go`; align each story's Clerkenstein DemoUser
`OrgEid` claim, `resources.go:45`.)_

## O6 — usable role set
_(record which replayed public `job_roles` have `job_role_skills`, and counts, once enumerated.)_
