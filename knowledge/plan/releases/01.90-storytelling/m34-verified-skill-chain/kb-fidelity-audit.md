---
title: "KB Fidelity Audit — M34 Verified-skill chain"
date: 2026-06-23
scope: milestone:M34
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Seeding framework + seeder fleet | `corpus/ops/seeding-spec.md` | `.agentspace/rosetta-extensions/stack-seeding/{seeders,seeder,blueprint}/` | PAIRED |
| G14 session-seeder bug (invalid enums/token) | `corpus/ops/seeding-spec.md` (M7c fleet ref) + NEW `stories-spec.md` | `seeders/jobsim_sessions.go` | PAIRED (doc-add in-scope) |
| Content/taxonomy-ref resolver (`contentref.go` → `TaxonomyRefs`) | `corpus/ops/seeding-spec.md` + `snapshot-spec.md` | `seeders/contentref.go` | PAIRED |
| Cross-surface closure gene (M23 pattern) | `corpus/ops/snapshot-spec.md` §M23 | `dna/snapshot.go`, `dna/fidelity_probe.go` | PAIRED |
| Production-isolation / safety (write surfaces) | `corpus/ops/safety.md` | `isolation/` | PAIRED |
| The verified-skill chain (7-table fan-out) reference | NEW `corpus/ops/demo/stories-spec.md` | `seeders/` (PersonaSeeder, new) | BLIND-AREA → covered by overview `Delivers` |

## Fidelity Findings

1. **seeding-spec.md — seeder-fleet scope.** Doc explicitly scopes itself to the M7a framework and defers the full fleet (sessions, content, activity) to M7c. It makes **no claim** that the session seeder writes valid sim_type/enum/result values, and **no claim** that verified skills are seeded. → **ALIGNED.** No stale claim the milestone would read as truth; the G14 fix + the verified-skill chain are net-new documented behavior this milestone delivers.

2. **snapshot-spec.md — M23 `snapshot-cross-surface-closure` gene.** Doc (lines 228-239, 464-467) accurately describes the M23 cross-surface closure gene: counts DISTINCT content-referenced taxonomy node-ids that don't resolve in the replayed `skiller.skills`, names a sample on failure, non-fatal in bring-up. Matches `dna/snapshot.go::snapshotCrossSurfaceClosure` + `fidelity_probe.go::CrossSurfaceDangling`. → **ALIGNED.** This is the pattern the new M34 *seed-side* closure gene mirrors (over `user_skills`/`user_skill_evidences`/`validation_attempt_skill_results.skill` instead of directus↔skiller).

3. **safety.md — isolation classes (§2.1).** Doc lists `Postgres / Redis / S3-private / pgvector` as `PerStackIsolated` ("inside the stack's own containers → seed freely"). The new M34 write surfaces (`jobsimulation.validation_attempt_*`, `jobsimulation.validation_criterion_results`, `public.local_jobsimulation_sessions`, `public.user_skills`, `public.user_skill_evidences`) are **all per-stack Postgres** → they inherit `PerStackIsolated` with no new shared-store risk. → **ALIGNED.** Milestone adds a confirming note (the new surfaces stay PerStackIsolated).

4. **data-DNA.json — surface enumeration.** Enumerates `jobsim-sessions`, `skillpath-sessions`, `assignments`, `activity`, `taxonomy`, `content`, plus M7a org/users/memberships/casbin. Does NOT yet enumerate `user_skills`/`user_skill_evidences`/`validation_attempt_*`. → **Expected.** These are PersonaSeeder's net-new surfaces; the M34 closure gene is a new data-DNA gene. Not a stale-doc finding — it's planned new work.

## Completeness Gaps

1. **The verified-skill chain has no corpus-side reference.** The 7-table fan-out, the G14 valid-value table, the DB-enforced vs logical landmines, and the claimed-vs-verified `user_level` requirement live ONLY in the gitignored `.agentspace/seeding_gaps.md` (analysis-of-record). The overview's `Delivers` + Scope make authoring `corpus/ops/demo/stories-spec.md` (graduating the spec per D12) an explicit milestone deliverable. → **Covered by a `Delivers →` milestone deliverable** — not an unmanaged blind area. (Per the gate rule, a blind area promoted to a milestone deliverable does NOT block Phase 1.)

## Applied Fixes

None applied pre-build. The corpus doc-half (`stories-spec.md` NEW + `seeding-spec.md`/`safety.md` updates) is authored as the milestone's documentation section (Phase 5 of build-milestone), not as a pre-flight backfill — it documents behavior this milestone *creates*, so it can only be written truthfully after the code lands.

## Open Items (require user decision)

None.

## Gate Result

YELLOW: proceed with tracking. No blind area is unmanaged (the one blind area — the verified-skill-chain reference — is a tracked milestone deliverable via the overview's `Delivers` line). No stale load-bearing claim exists that the milestone's implementation would read as truth: the existing docs correctly scope themselves (seeding-spec → M7a/M7c boundary; snapshot-spec → M23 closure; safety → isolation classes) and the M34 work is net-new behavior those docs do not yet (and should not yet) describe. Build-milestone Phase 1 may proceed.
