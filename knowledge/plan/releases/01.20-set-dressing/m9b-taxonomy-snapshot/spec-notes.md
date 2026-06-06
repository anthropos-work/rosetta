# M9b — Spec notes

Technical notes accumulate here during build. Prod evidence (2026-06-06):
[`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../../../../.agentspace/scratch/roadmap-research-2026-06-06.md).

## Public taxonomy capture
_skiller `{categories,specializations,skills,job_roles}` filtered `organization_id IS NULL`. Parent-scoped:
`{skill,job_role}_embeddings` (vectors only), `{skill,job_role}_translations`, `job_role_skills` via
`… WHERE skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)`._

## Bulk-COPY replay (per-stack)
_M7a perf path; writes only to the per-stack-isolated skiller Postgres (offset port). Respects the M9a isolation
guard._

## pgvector index rebuild on replay
_Carry vectors verbatim; rebuild the index locally (don't transport ~689 MB). Record index type/params in the
manifest; gate via the embedding-dimension integrity gene._

## Fidelity genes + coverage
_Taxonomy fidelity + public-only genes; data-DNA coverage `waived → taxonomy-seeded`._
</content>
