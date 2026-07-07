# Skiller Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **July 2026** ("skiller-in-app"), the standalone `skiller` Go microservice has been **merged into the
> `app` monolith** (the service the platform calls "backend"). Skiller no longer runs as a separate service —
> not in the local compose, not in production ECS.
>
> Where everything went:
>
> * **Domain** — the skills taxonomy (60K+ skills graph), skill/job-role embeddings, AI skill matching, and
>   job-role skills now live inside `app`: Ent models in `app/internal/data/ent/schema/` (`skill.go`,
>   `jobrole.go`, `skill_embeddings.go`, `job_role_embeddings.go`, `category.go`, `specialization.go`, …),
>   data in the **`public` schema** of the same PostgreSQL database. The old `skiller` DB schema is **legacy —
>   no longer authoritative**.
> * **RPC** — the skiller Connect-RPC surface (`SkillerService`) is now served by `app`
>   (`app/internal/rpc/skillerrpc/`). Consumers keep the same env var, re-pointed:
>   `SKILLER_RPC_ADDR=http://backend:8083` locally; `skiller_rpc_addr = http://backend:8081` in production
>   terraform. The externally-reached methods (`GetSkills`, `GetSkill`, `SearchSkill`, `MatchSkill`,
>   `GetJobRole`) are implemented in `app`.
> * **GraphQL** — the skiller subgraph was **removed** from the WunderGraph/Cosmo federation; `app`'s `backend`
>   subgraph now serves the taxonomy types/queries (`app/.../graph/schemas/skiller_taxonomy.graphqls`). The
>   `categoryTree` / `fullCategoryTree` queries were **dropped entirely** (not ported).
> * **Infrastructure** — the skiller ECS service and its terraform module were removed from production (the
>   old ECR repo is intentionally orphaned pending manual deletion). `app`'s internal app→skiller RPC path is
>   retired (app PR #989).
> * **Repo** — the `skiller` git repo still exists but is **legacy/decommissioned**, no longer deployed or
>   cloned by `make init`.
>
> **For current documentation of this domain, see [Backend (`app`)](./backend.md).**

## Still-true domain knowledge

* The **taxonomy data** (60K skills / 18K job roles) is a dataset owned by the service DB, not by the
  `anthropos-work/taxonomy` library — that library only supplies `NodeID` generation helpers
  (see [Shared Libraries → taxonomy](../architecture/shared_libraries.md#taxonomy)).
* **Embeddings** live in dedicated tables (`skill_embeddings`, `job_role_embeddings` — OpenAI
  text-embedding-3-small, `extensions.vector(1536)` with IVFFLAT indexes), now in the `public` schema. The
  `extensions` schema (housing `pgvector`) must exist before the vector migrations apply. See
  [AI Architecture](../architecture/ai_architecture.md).
* **Localization**: per-skill / per-job-role translations (`skill_translations`, `job_role_translations`)
  across 8 `ContentLanguage`s carried over into `app`.
* **Demo/dev set-dressing (v1.2)**: the **public** taxonomy (`organization_id IS NULL`) is captured read-only
  from prod and replayed per-stack by the snapshot mechanism — see
  [`corpus/ops/snapshot-spec.md`](../ops/snapshot-spec.md). (Note: verify the capture source post-merge — the
  taxonomy now lives in the `public` schema, not the legacy `skiller` schema.)

## Related Documentation

* [Backend (`app`)](./backend.md) — where the skiller domain now lives
* [AI Architecture](../architecture/ai_architecture.md) — embeddings, RAG, provider routing
* [Dependency Map](../architecture/dependency_map.md)
