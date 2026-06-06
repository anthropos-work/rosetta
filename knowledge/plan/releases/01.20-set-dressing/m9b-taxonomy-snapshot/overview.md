---
milestone: M9b
slug: taxonomy-snapshot
version: v1.2 "set dressing"
milestone_shape: section
status: planned
created: 2026-06-06
complexity: large
delivers: extends corpus/ops/snapshot-spec.md (the taxonomy capture/replay path) + updates corpus/ops/seeding-spec.md (taxonomy promoted waived → snapshot-seeded)
---

# M9b — Taxonomy snapshot (the first real surface)

## Goal
Prove the M9a framework on the **real ~2.1 GB taxonomy surface**: capture the **public** skiller catalog from a
safe (non-primary) source, replay it per-stack via bulk `COPY`, **rebuild the pgvector index on replay**, and gate
it on the snapshot-fidelity + public-only genes — driving data-DNA coverage from `waived` to its first
`snapshot-seeded` surface.

## Why section (not iterative)
The surface is fully enumerable up front (prod-measured 2026-06-06): the skiller tables, their public-vs-private
split, the embedding/translation parent-scoping, and the index-rebuild step. The fidelity gate is a per-surface
acceptance check. Reuses M9a's framework — no emergent path.

## Scope
- **In:**
  - **Public taxonomy capture** (note #3 — public only, prod-measured):
    - `skiller.categories` / `specializations` / `skills` / `job_roles` filtered `organization_id IS NULL`
      (keeps the full public catalog: ~42.8K skills, ~22.3K roles, ~1.4K specializations, 22 categories; drops the
      customer tail automatically).
    - `skiller.skill_embeddings` / `job_role_embeddings` (vectors only) + `skill_translations` /
      `job_role_translations` + `job_role_skills` — captured via the **public parent** join
      (`… WHERE skill_id IN (SELECT id FROM skiller.skills WHERE organization_id IS NULL)`), since they carry no
      `organization_id`.
  - **Bulk-`COPY` replay** per-stack (the M7a perf path), writing only to the per-stack-isolated skiller Postgres.
  - **pgvector index rebuild on replay** — carry the vectors verbatim, **rebuild the ~689 MB index locally**
    (don't capture/transport it). Gated by the embedding-dimension integrity gene.
  - **The taxonomy fidelity + public-only genes** wired into the data-DNA; coverage `waived → taxonomy-seeded`.
  - Wiring the taxonomy snapshot into the **`stack-seeding` DAG node** (`… → taxonomy/content (snapshot) → activity`).
- **Out:**
  - The Directus content surface (M10).
  - Recipes / presets / corpus product layer (M11).
  - Recompute of embeddings (that's AI-content, v1.3 — M9b carries vectors verbatim).

## Depends on
**M9a** (the `stack-snapshot` extension + the capture-source policy + the tenant firewall + the `.agentspace`
store + the fidelity-gene class). **Parallel with:** none (gates M10 + M11).

## Open questions (resolve during build)
- Whether the largest captures (skills, embeddings) need keyset-chunked `COPY` vs a single streamed `COPY` on the
  replica (size them via the catalog-first dry-run; lean: single streamed `COPY`, chunk only if the read session
  exceeds the bound).
- The pgvector index parameters to rebuild on replay (match prod's index type/params; record in the manifest).
- Whether `job_role_skills` (the role→skill mapping) needs both endpoints public, or just the public roles
  (confirm referential integrity against the public-skill set).

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` (M9a's capture/replay contract + the capture-source policy + the firewall)
- `corpus/services/skiller.md` (the taxonomy schema: categories/specializations/skills/roles + pgvector embeddings + translations)
- `corpus/ops/seeding-spec.md` (the isolation guard + the perf path + the DAG node)
- `corpus/architecture/alignment_testing.md` (the snapshot-fidelity + public-only genes from M9a)

## Delivers → corpus/ops/snapshot-spec.md (extend) + corpus/ops/seeding-spec.md (update)
- **extends `corpus/ops/snapshot-spec.md`**: the taxonomy capture/replay path (public filters, parent-scoped
  embeddings/translations, the index-rebuild-on-replay step).
- **updates `corpus/ops/seeding-spec.md`**: the `taxonomy` surface promoted from `waived` to `snapshot-seeded`.
</content>
