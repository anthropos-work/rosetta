---
title: "KB Fidelity Audit — M36 Dashboard surfaces"
date: 2026-06-23
scope: milestone:M36
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Seeding framework + seeder fleet | `corpus/ops/seeding-spec.md` | `.agentspace/rosetta-extensions/stack-seeding/{seeders,seeder,blueprint}/` | PAIRED |
| Stories & Heroes model + multi-org (per-story OrgID) | `corpus/ops/demo/stories-spec.md` §"Stories & Heroes model (M35)" | `blueprint/stories.go`, `seeders/{org,users,identity,jobsim_sessions,assignments,skillpath_sessions,activity}.go` | PAIRED |
| Verified-skill chain (PersonaSeeder, 7-table) | `corpus/ops/demo/stories-spec.md` §"For engineers" | `seeders/persona.go`, `persona_write.go`, `taxonomyref.go`, `jobroleref.go` | PAIRED |
| `membership_skills` (the mapped surface / funnel) | — (NEW, M36 `Delivers → stories-spec.md`) | NEW `seeders/*` (this milestone) | BLIND-AREA → covered by overview `Delivers` (extend stories-spec.md) |
| `tags` + `membership_tags` (teams/business-units) | — (NEW, M36) | NEW `seeders/*` | BLIND-AREA → covered by overview `Delivers` |
| `organization_target_roles` + `user_target_roles` (gap + mobility) | — (NEW, M36) | NEW `seeders/*` | BLIND-AREA → covered by overview `Delivers` |
| Succession feeders (`validation_attempt_*` already exist; `interview_extraction_results`) | partial (`stories-spec.md` 7-table chain) | `seeders/persona*.go` (existing validation rows) + NEW interview rows | PAIRED (extend) |
| `job_simulation_feedbacks` (~2:1 positive) | — (NEW, M36) | NEW `seeders/*` | BLIND-AREA → covered by overview `Delivers` |
| Assignments fix (status mix + due_date + sessions) | `corpus/ops/seeding-spec.md` (assignments surface) | `seeders/assignments.go` (current: all `active`, no due_date/sessions) | PAIRED (fix in-scope) |
| `skillpath_sessions` `completed` share | `corpus/ops/seeding-spec.md` | `seeders/skillpath_sessions.go` | PAIRED (verify in-scope) |
| Org-scale distributions (claimed-vs-verified gap, growth arc, AI-readiness) | `corpus/ops/demo/stories-spec.md` §3c (claimed-vs-verified) | `seeders/persona*.go` + NEW distribution logic | PAIRED (extend) |
| Closure gene (seed-side, org-agnostic) | `corpus/ops/demo/stories-spec.md` §"Closure across all orgs" + `dna/README.md` | `dna/seed_closure.go`, `dna/fidelity_probe.go` | PAIRED |
| Production-isolation / safety (PerStackIsolated) | `corpus/ops/safety.md` §2.1 | `isolation/` | PAIRED |

## Fidelity Findings

1. **stories-spec.md / seeding-spec.md — dashboard org-aggregate surfaces are undocumented (correctly).** Both docs accurately describe the M34 verified-skill chain (PersonaSeeder, the 7-table fan-out, the claimed-vs-verified `user_level` requirement) and the M35 Stories & Heroes / multi-org model. Neither makes any claim that `membership_skills`, `tags`/`membership_tags`, `organization_target_roles`/`user_target_roles`, `job_simulation_feedbacks`, or the assignment-status-mix are seeded. → **ALIGNED.** No stale claim the milestone would read as truth; these org-aggregate surfaces are net-new behavior M36 delivers and documents (per the overview `Delivers` line — extend stories-spec.md + seeding-spec.md), exactly as M34 graduated the verified-skill chain.

2. **assignments.go — the documented G19 gap is real and unfixed.** Current `seeders/assignments.go` writes every assignment with `status="active"`, a single backdated `created_at`, no `due_date`, and no `organization_assignment_sessions` rows. This matches the spec's G19 ("Assignments all not_started; no feedback") and the overview's "assignments fix" scope. → **ALIGNED** (the doc/spec does not falsely claim the fix is done; it is M36-in-scope work).

3. **safety.md — isolation classes (§2.1).** Doc lists `Postgres / Redis / S3-private / pgvector` as `PerStackIsolated` ("inside the stack's own containers → seed freely"). M36's new write surfaces (`public.membership_skills`, `public.tags`, `public.membership_tags`, `public.organization_target_roles`, `public.user_target_roles`, `public.organization_assignment_sessions`, `public.job_simulation_feedbacks`, `jobsimulation.interview_extraction_results`) are **all per-stack Postgres** → they inherit `PerStackIsolated` with no new shared-store risk, and stay `organization_id`-scoped per story. → **ALIGNED.** Milestone adds a confirming note (the new org-scoped surfaces stay PerStackIsolated + audited).

4. **data-DNA.json — surface enumeration.** Enumerates the M7a/M34 surfaces (org/users/memberships/casbin/jobsim-sessions/skillpath-sessions/assignments/activity/taxonomy/content + the M34 verified-skill closure). Does NOT yet enumerate `membership_skills`/`tags`/`target_roles`/`feedbacks`. → **Expected.** These are M36's net-new surfaces; the believable-distribution assertions are new data-DNA genes this milestone may add. Not a stale-doc finding — planned new work.

## Completeness Gaps

1. **The dashboard org-aggregate surfaces have no corpus-side reference yet.** The mapped→verified funnel (`membership_skills` outnumbering verified), the team/tag slice dimension, the role-readiness + two-way mobility (`*_target_roles`), the succession feeders, the ~2:1 feedback ratio, and the assignment-status-mix live ONLY in the gitignored `.agentspace/seeding_gaps.md` §3b/§3c/§6a#5 (analysis-of-record). The overview's `Delivers` ("extend stories-spec.md (dashboard surfaces) + seeding-spec.md") makes documenting them an explicit milestone deliverable. → **Covered by a `Delivers →` milestone deliverable** — not an unmanaged blind area. (Per the gate rule, a blind area promoted to a milestone deliverable does NOT block Phase 1.)

2. **O4 — exact MIGRATED column/storage-key names not yet pinned for the new tables.** The spec explicitly flags (O4, ENG) that `membership_skills`, `tags`, `membership_tags`, `organization_target_roles`, `user_target_roles`, `organization_assignment_sessions`, `job_simulation_feedbacks`, and `jobsimulation.interview_extraction_results` need one `\d` introspection pass on a replayed stack before COPYs are written. → **Tracked as a Phase-1 build task** (introspect demo-3, record in spec-notes.md), not a doc-staleness finding.

## Applied Fixes

None applied pre-build. The corpus doc-half (extend `stories-spec.md` with the dashboard surfaces + update `seeding-spec.md`) is authored as the milestone's documentation section (Phase 5 of build-milestone), not as a pre-flight backfill — it documents behavior this milestone *creates*, so it can only be written truthfully after the code lands. The same applies to the safety.md confirming note and any new data-DNA genes.

## Open Items (require user decision)

None. (O4 column-name introspection is an ENG build-time task, not an operator decision — the operator already sanctioned reading prod/demo structure for the storytelling release.)

## Gate Result

YELLOW: proceed with tracking. No blind area is unmanaged — the dashboard org-aggregate surfaces are a tracked milestone deliverable via the overview's `Delivers` line (extend stories-spec.md + seeding-spec.md). No stale load-bearing claim exists that the milestone's implementation would read as truth: the existing docs correctly scope themselves to M34 (verified-skill chain) and M35 (Stories & Heroes / multi-org), and the M36 dashboard surfaces are net-new behavior those docs do not yet (and should not yet) describe. The one unfixed surface a doc references — assignments (G19) — is honestly described as a gap, not falsely claimed done. Build-milestone Phase 1 may proceed.
