---
title: "KB Fidelity Audit — M35 Stories & Heroes model + multi-org"
date: 2026-06-23
scope: milestone:M35
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| stories[] blueprint schema | `corpus/ops/demo/stories-spec.md` §"declaring a hero"; `seeding-spec.md` §"verified-skill chain"; spec §4c | `blueprint/blueprint.go` (StackSeed/Persona) | PAIRED |
| multi-org OrgID/orgClerkID parameterization | `stories-spec.md` (M35 forward-pointer); `.agentspace/seeding_gaps.md` §4e | `seeders/org.go:19,22` + consumers (users/identity/jobsim_sessions/assignments/persona) | PAIRED |
| PersonaSeeder roster scaling + #M34-D7 guards | M34 `decisions.md` §D-M34-7; spec §4b/§4c | `seeders/persona.go` (personaUserIndex/personaIndexMap), `seeders/taxonomyref.go` (take/skillsForRole) | PAIRED |
| trajectory logic (thriving/struggling) | `.agentspace/seeding_gaps.md` §4b; `stories-spec.md` §user_level | `seeders/persona.go` selfEvalLevel (partial); rest is M35-new | DOC-ONLY (code lands in M35) |
| Clerkenstein org-claim alignment | spec §4e/§4f; M34 close ORCH-context (single-identity until M37) | `clerkenstein/clerk-frontend/resources.go:37` DefaultDemoUser().OrgEid | PAIRED |
| seed-side closure gene across orgs | `stories-spec.md` §Closure; `dna/README.md` | `dna/seed_closure.go` (org-agnostic measure) | PAIRED |

## Fidelity Findings

1. **stories-spec.md blueprint shape vs code** — ALIGNED. Doc §"declaring a hero" shows the `personas` entry
   (id/name/role/verified/self_eval_bias) and states it's optional+additive; `blueprint.Persona` carries
   exactly those fields, `Personas` is `omitempty`-optional, and the seeders no-op on an empty list. The doc
   explicitly forward-points M35 to "the full `stack.stories.yaml` (multi-org, the hero trio,
   vantage/trajectory, the cockpit)" — the M35 scope, correctly framed as not-yet-built.
2. **§4e "4 consuming seeders" claim vs code** — ALIGNED. The spec names users/identity/jobsim-sessions/
   assignments as the OrgID consumers; code confirms `OrgID` is read in exactly org.go (def), users.go,
   identity.go, jobsim_sessions.go, assignments.go, plus persona.go (the M34 addition the spec predates but
   is consistent with). No drift.
3. **#M34-D7 routing vs M35 overview** — ALIGNED. M34 `decisions.md` routes the `len(Personas) <= Size`
   validation + index-collision warning + short-role-pool top-up product call to M35 (Fate-3); M35
   `overview.md` In-list carries both. The code edges (`personaUserIndex` hashing into `[1,Size]`,
   `take()` keeping a short role pool) match the decision's description exactly.
4. **Clerkenstein single-identity boundary** — ALIGNED. `resources.go` `SingleSessionMode: true` +
   `DefaultDemoUser()` returns one identity with `OrgEid = 22222222-…`. The orchestrator context confirms
   the multi-identity seat-switch is M37, not M35; M35's "org-claim alignment" is the data-side requirement
   that one seeded story's OrgID == DefaultDemoUser().OrgEid so the existing single login still works.
5. **closure gene org-agnostic** — ALIGNED. `seed_closure.go` measures all distinct seeded skill refs vs the
   replayed taxonomy with no org filter, so the done-bar's "closure stays green across all orgs" holds by
   construction (the gene is already whole-stack).

## Completeness Gaps

None load-bearing. The trajectory/vantage fields and the multi-story YAML shape are intentionally
undocumented-in-code (they are M35's deliverable); `stories-spec.md` already scopes itself to M34 and points
forward. M35 Phase 5 extends `stories-spec.md` with the realized model — a planned `delivers` item, not a
blind area.

## Applied Fixes
None needed — every load-bearing claim aligned on inspection.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to build-milestone Phase 1. The verified-skill spine (M34) is accurately documented; M35
extends `stories-spec.md` with the multi-org Stories model as a Phase 5 `delivers` deliverable.
