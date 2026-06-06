# M9b — Decisions

_Implementation decisions with rationale. ID scheme: M9b-D1, M9b-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| _(none yet)_ | | | |

## Open at design (to resolve during build)
- M9b-Q1: keyset-chunked `COPY` vs single streamed `COPY` for skills/embeddings (size via catalog dry-run; lean single, chunk only if the read bound is exceeded).
- M9b-Q2: pgvector index params to rebuild on replay (match prod; record in manifest).
- M9b-Q3: `job_role_skills` — both endpoints public vs public roles only (confirm referential integrity vs the public-skill set).

## Prod evidence (2026-06-06, catalog-only)
- skiller ≈ 2.1 GB. Public split: skills 42,763/794 · job_roles 22,315/2,381 · specializations 1,442/154 · categories 22/42.
- `skill_embeddings` 692 MB total but heap only 3.3 MB → ~689 MB is the pgvector index ⇒ **rebuild on replay**.
- embeddings carry no `organization_id` → scope via parent (`skill_id`/`job_role_id`).
</content>
