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

## O6 — usable role set (RESOLVED 2026-06-23 via prod public structural read)
- The public taxonomy has **many** `job_roles` with `job_role_skills` (capped ~10 skills each). Confirmed
  resolving (≥10 public role-skills): **Backend Developer** (`J-BACKEN-A9ED`), **Engineering Manager**
  (`J-ENGMAN-E0E7`), **Software Engineer** (`J-SOFTWA-5757`), **Account Executive** (`J-ACCOUN-F7F6`),
  **Sales Manager** (`J-SALESM-84AA`), Cloud Specialist, IT Consultant, Security Consultant, Compliance
  Analyst, Data Center Technician (all `J-…`, 10 skills).
- **NOT resolving (0 public role-skills):** "Backend Engineer", "Sales Development Rep", "RevOps Lead",
  "Revenue Operations Lead", "Sales Development Representative". The spec's literal roster role labels for
  Solvantis + Cervato's engineer DON'T resolve.
- **Decision (role-coherence, D3):** the roster uses RESOLVING role names so heroes get role-coherent skills:
  Cervato {Maya/Tom = **Backend Developer**, Dan = **Engineering Manager**}; Solvantis {Sara/Nick =
  **Account Executive**, Leah = **Sales Manager**}. Display labels are swappable (D16); role-coherence is the
  load-bearing `[ENG]` property. A non-resolving role still falls back to flat (closure green) — but coherent
  is more believable, so we pick resolving roles.
- Node-ids above are PROD reads — the seeder must resolve them from the REPLAYED stack at runtime (a subset of
  prod public), never hardcode. Build a job-role pool resolver mirroring the skill-side TaxonomyRefs.

## O4 — supporting-population schema (RESOLVED 2026-06-23)
- A member's role lives on **`public.memberships`**: `job_role_id` (varchar, the **node-id form** `J-XXXXXX-XXXX`,
  NOT a uuid — verified against live sample rows), `job_role_name` (varchar, the human name), `job_role_title`
  (varchar, optional free-text). The dashboard/Members page reads role from these.
- `organization_target_roles.target_job_role_id` + `user_target_roles.target_job_role_id` are also `J-…` node
  refs — but those are the **gap/mobility** surfaces owned by **M36** (out of M35 scope).
- Supporting-population fidelity for M35 = set `job_role_id` (real replayed `J-…`) + `job_role_name` on every
  membership; ramped `joined_at`; real names on non-hero members (already done in M34's users.go).

## Pre-flight audits — stories[] blueprint (first section)
- **Phase 0b KB-fidelity: GREEN** (2026-06-23). Report: `kb-fidelity-audit.md`. Every load-bearing claim
  (blueprint shape, §4e 4-seeder OrgID consumer list, #M34-D7 routing, Clerkenstein single-identity boundary,
  org-agnostic closure gene) ALIGNED on inspection; no stale docs, no blind areas. `stories-spec.md` exists
  (M34) + is M35's extend target (a Phase 5 `delivers` item, not a blind area).

## Topic → doc → code triples (audit cache)
- stories[] blueprint → `stories-spec.md` §declaring-a-hero + `seeding-spec.md` §verified-skill-chain → `blueprint/blueprint.go`
- multi-org OrgID → `seeding_gaps.md` §4e → `seeders/org.go:19,22` + users/identity/jobsim_sessions/assignments/persona
- roster + #M34-D7 → M34 `decisions.md` §D-M34-7 → `seeders/persona.go` (personaUserIndex/personaIndexMap), `seeders/taxonomyref.go` (take)
- Clerkenstein org-claim → `clerkenstein/clerk-frontend/resources.go:37` DefaultDemoUser().OrgEid (single-identity until M37)
- closure gene → `dna/seed_closure.go` (org-agnostic)
