# M10 — Spec notes

## Pre-flight audits — section 1 (content-store fork + directus surface)
- **Phase 0b KB-fidelity (2026-06-06): YELLOW** — report `kb-fidelity-audit.md`. 0 blind areas; 2 stale claims
  (KB-1 db-access/snapshot-spec Directus-source location; KB-2 firewall org-only predicate), both on the M10
  content path that this milestone's `Delivers →` corrects. Build sources truth from the locked decisions, not
  the stale docs. Gate: proceed.
- Topic→doc→code triples: snapshot framework → `snapshot-spec.md` → `stack-snapshot/*`; Directus source →
  `db-access.md` → `directus` schema (verified via `postgres` MCP); isolation/DAG → `seeding-spec.md` →
  `stack-seeding/{isolation,seeders,dna}`; `sim_id`/`skill_path_id` consumers → `jobsimulation.md`/`skillpath.md`
  → `stack-seeding/seeders/{jobsim_sessions,skillpath_sessions}.go`.

## Prod-verified directus-schema shape (read-only, 2026-06-06)
- `directus` schema is INSIDE the same `postgres` database (`marco_read` MCP reaches it read-only — Decision 2).
- Public predicate (Decision 3, strict): `private = false AND tenant_id IS NULL AND status = 'published'`.
- Strict-published counts: `directus.simulations` 304 (of 2,597 total / 647 `private=false`); `directus.skill_paths`
  22 (of 263). The prompt's ~190-path estimate omitted the `tenant_id IS NULL` intersection — 22 is authoritative.
- Only `simulations` + `skill_paths` carry `private`/`tenant_id`/`status` (the parents). All child collections are
  column-less w.r.t. the public predicate → **parent-scoped** via FK (`simulation`, `task`, `role`, `sequence`) —
  the exact M9b `ParentScopes` pattern, generalized to the directus predicate.

## Per-stack content-store decision (the defining fork)
_(per-stack Directus container vs direct per-stack Directus-Postgres replay — resolve in first spike)_

## Content capture
The `directus` surface (`stack-snapshot/directus/directus.go`): 9 tables in FK replay order, captured under the
directus PublicPredicate. Prod-verified public counts (read-only, count-only, 2026-06-06):

| Table | Capture scope | Public rows |
|---|---|---|
| `simulations` | scope-bearing root (`private=false AND tenant_id IS NULL AND status='published'`) | 304 |
| `skill_paths` | scope-bearing root (same predicate) | 22 |
| `resource` | pure-reference (global learning-resource library; no tenant column) | 1,543 |
| `roles` | parent-scoped via `simulations` (col `simulations`) | 953 |
| `sim_tasks` | parent-scoped via `simulations` (col `simulation`) | 949 |
| `sequences` | parent-scoped via `simulations` (col `simulation`) | 304 |
| `task_checks` | MULTI-LEVEL via `sim_tasks`→`simulations` (ParentFilter) | 2,242 |
| `task_sub_checks` | MULTI-LEVEL via `task_checks`→`sim_tasks`→`simulations` | 2,850 |
| `sequences_roles` | parent-scoped via BOTH `sequences` AND `roles` | 953 |

All captured columns prod-verified to exist (sampled `order`/`group` reserved words → safe via `pg.QuoteIdent`).
Directus has NO DB-level FKs (relations in the app layer) → closure is by convention, validated by the counts above.

## Per-stack Directus store fork (M10-D2) — provision plan
`stack-snapshot/directus/provision.go`: the ordered store-fork is **bootstrap → replay → boot**:
1. `directus bootstrap` creates the `directus_*` system schema + the user-collection table STRUCTURE against the
   stack's empty Postgres `directus` schema (Directus owns its version-specific DDL — we never capture/replay it).
2. `stacksnap replay --surface directus --stack demo-N` bulk-COPYs the content rows into the now-existing tables.
3. Boot the per-stack Directus container pointed at the stack's `directus` schema; CMS/studio-desk for THIS stack
   point `DIRECTUS_BASE_ADDR` at the offset-port container, NOT `content.anthropos.work`.
The `EnvContract.Validate()` hard-rejects any per-stack env that points at the prod Directus (`anthropos.work`).
The replay path itself is the framework's existing generic `CopyIn(schema, table, …)` → works for the `directus`
schema unchanged. The live container boot is a **documented operational step** (M9b discipline) — the build proves
the contract + plan hermetically.

## Media / blobs (M10-D4)
`stack-snapshot/directus/media.go`: the file-REFS are the floor (always captured + replayed); the blob BYTES are
the S3-credential-gated operational add.
- **File-ref columns (uuid-typed, prod-verified 2026-06-06):** `simulations.cover`, `skill_paths.{cover,image,video}`,
  `roles.avatar`, `sequences.scenario_video`, `resource.{image,file}`. EXCLUDED: `sequences.intro_video` is
  `character varying` (a video EMBED url, not a `directus_files` id — a uuid=varchar join would error).
- **Referenced public files:** **1,311** `directus_files` rows (of 10,340 total), parent-scoped to the public content
  via `ReferencedFilesFilter()` (verified read-only).
- **Blob bytes:** live in Directus's OWN S3 bucket. `BlobBytesAvailable()` is **false** here (no confirmed S3 read
  access) → `MediaCaveat` surfaced in PENDING. Until S3-read is wired, the per-stack Directus serves refs with a
  local-storage adapter + placeholder assets (a believable structural demo). Blob mirroring → per-stack-isolated
  bucket only, never the shared prod S3.

## Content replay seeder
_(wired into M9 framework + M7a DAG; content fidelity gene)_

## sim_id / skill_path_id / resource_id linkage
_(resolve the v1.1 free-value content refs against real replayed content)_
