# M9b — Spec notes

Technical notes accumulate here during build. Prod evidence (2026-06-06):
[`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../../../../.agentspace/scratch/roadmap-research-2026-06-06.md).

## Public taxonomy capture — the 10-table surface (prod-verified 2026-06-06, catalog + bounded counts)

The surface, in FK (replay) dependency order, with the capture filter each table gets:

| # | Table | Org col? | Capture scope | Public rows |
|---|-------|----------|---------------|-------------|
| 1 | `categories` | yes | `organization_id IS NULL` | 22 |
| 2 | `job_role_categories` | **no** | pure-reference (whole table) | 22 |
| 3 | `specializations` | yes | `organization_id IS NULL` | 1,442 |
| 4 | `skills` | yes | `organization_id IS NULL` | 42,763 |
| 5 | `job_roles` | yes | `organization_id IS NULL` | 22,315 |
| 6 | `skill_embeddings` | no | public-via `skills` (`skill_id IN public skills`) — vector `small_embedding3` dim 1536 | 42,763 |
| 7 | `job_role_embeddings` | no | public-via `job_roles` — vector `small_embedding3` dim 1536 | 18,904 |
| 8 | `skill_translations` | no | public-via `skills` | 85,491 |
| 9 | `job_role_translations` | no | public-via `job_roles` | 43,550 |
| 10 | `job_role_skills` | no | public-via **BOTH** `job_roles` AND `skills` | 72,556 |

**FK graph (pg_constraint):** `skills.parent→specializations` · `specializations.parent→categories` ·
`job_roles.category_id→job_role_categories` (separate table — NOT `categories`) · embeddings/translations→parent ·
`job_role_skills→{job_roles,skills}`. Public hierarchy is referentially closed: 0 public skills with a
customer/missing specialization parent; 0 public specs with a customer/missing category; 0 public roles with a
missing category. **`job_role_skills` is the one exception** — 3 public-role rows link a customer skill, so it is
scoped by BOTH endpoints (M9b-D3).

**Parent-scoped filter (M9b-D2):** column-less tables get a real subquery predicate, e.g.
`skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)`. `job_role_skills` ANDs two such
predicates. The M9a framework recorded `PublicVia` but applied an EMPTY filter to column-less tables — M9b adds
`TableSpec.PublicViaFK` (the FK column) so the adapter builds the predicate, and the post-capture tenant probe
re-checks via the same parent join.

## Bulk-COPY replay (per-stack)
_M7a perf path; writes only to the per-stack-isolated skiller Postgres (offset port). Respects the M9a isolation
guard._

## pgvector index rebuild on replay
_Carry vectors verbatim; rebuild the index locally (don't transport ~689 MB). Record index type/params in the
manifest; gate via the embedding-dimension integrity gene._

## Fidelity genes + coverage
_Taxonomy fidelity + public-only genes; data-DNA coverage `waived → taxonomy-seeded`._

## Pre-flight audits — Section 1 (taxonomy surface)
- **Phase 0b KB-fidelity: GREEN** (2026-06-06). Report: `kb-fidelity-audit.md`. 0 blind areas, 0 stale
  load-bearing claims, 1 incidental completeness gap (skiller.md omits `job_role_categories` — not load-bearing;
  captured in this milestone's spec-notes). Topic→doc→code: snapshot-spec.md ↔ stack-snapshot/**;
  skiller.md ↔ live skiller schema; seeding-spec.md ↔ data-dna.json; alignment_testing.md ↔ dna/snapshot.go.
- Audit reuse: all M9b sections touch the same two subsystems (stack-snapshot/** + stack-seeding/dna/**) against
  the same four docs — reuse this verdict for subsequent sections unless a doc in scope is edited.
</content>
