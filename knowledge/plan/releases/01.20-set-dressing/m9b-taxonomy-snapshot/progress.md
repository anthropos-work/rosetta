# M9b — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [ ] Public taxonomy capture: `skiller.{categories,specializations,skills,job_roles}` filtered `org_id IS NULL`
- [ ] Parent-scoped capture: `skill_embeddings` / `job_role_embeddings` (vectors only) + `skill_translations` / `job_role_translations` + `job_role_skills` via the public-parent join
- [ ] Bulk-`COPY` replay per-stack (M7a perf path; per-stack-isolated skiller Postgres only)
- [ ] pgvector index **rebuild on replay** (carry vectors verbatim, don't transport the ~689 MB index); embedding-dimension gene green
- [ ] Taxonomy fidelity + public-only genes wired; coverage `waived → taxonomy-seeded`
- [ ] Taxonomy snapshot wired into the `stack-seeding` DAG node
- [ ] Delivers: extend `corpus/ops/snapshot-spec.md` (taxonomy path) + update `corpus/ops/seeding-spec.md` (taxonomy promoted)

## Final review
_(filled at close)_
</content>
