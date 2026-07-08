# M211 — Spec notes

_Iteration-protocol-specific technical notes accumulate here (per-tik triage, fix-surface routing, gate measurements)._

## Pre-flight audits — iter-01
- **KB-fidelity gate (Phase 0b):** **GREEN (inherited from M210, spot-confirmed).** M210 closed
  2026-07-08 (same day) as the corpus + skills re-ground with a GREEN `/developer-kit:audit-kb-fidelity`
  (KB-1/2/3 resolved). Orchestrator bans re-deriving M208/M209/M210 facts, so ran a proportionate
  spot-confirmation instead of a full re-audit: grep of `corpus/ops` + `corpus/architecture` +
  `corpus/services` for residual `skiller.<table>` refs → 0 hits. Standing GREEN holds; strategy authored
  against verified docs. (Full re-audit deferred as redundant given the same-day M210 delivery.)

## Baseline facts (iter-01 recon — do not re-derive)
- Warm merged `stack-dev`: 11 `anthropos-*` containers up, **no skiller container** (4-subgraph compose).
- Taxonomy cache `c75ce94d6a8021cad2915ddb4fb3dd4d/`: 10 `skiller.*.copy` payloads + manifest; skills
  42,790 / job_roles 22,470 / specializations 1,447 / categories 23 / skill_embeddings 42,790 /
  job_role_embeddings 18,919 / skill_translations 85,545 / job_role_translations 43,550 / job_role_skills
  72,705 / job_role_categories 22. Cached `skills` columns (15): id, name, description, aliases, is_soft,
  parent, node_id, keywords, created_at, parent_node_id, updated_at, created_by_user, organization_id,
  deleted, long_description.
- rext authoring HEAD `2f06e78` = `quick-change-m209`; 0 residual `skiller.<table>` in production code.
  Consumption pin `.agentspace/rext.tag` = `v1.10.1` (STALE — must re-pin to `quick-change-m209`).
- `demo-stack/migrate-demo.sh` already: extensions-schema bootstrap + `wait_pg` + sentinel-race defense +
  M209 re-ground (`app:public cms:cms jobsimulation:jobsimulation skillpath:skillpath`). The dev cold path
  is the M25-D9 gap (platform `make migrate` un-editable → needs a rext dev pre-migrate hook).
