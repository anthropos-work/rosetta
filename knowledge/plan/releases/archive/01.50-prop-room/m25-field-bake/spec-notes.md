# M25 — Spec notes

_Technical detail that doesn't belong in `overview.md` (code maps, contracts, edge cases). Accumulates during build._

## Pre-flight audits — field-bake (Phase 0b, 2026-06-13)

KB-fidelity audit (3 parallel Explore sweeps; authoring copy `.agentspace/rosetta-extensions` @ `6a4749d`,
tags `prop-room-m21..m24`). **Verdict: YELLOW** (no blind areas, no stale load-bearing claim; two arg-hint
drifts fixed inline). Report: `kb-fidelity-audit.md`.

Topic → doc → code triples verified ALIGNED:
- local-Directus provisioning + exit-0/exit-4 split → `corpus/ops/directus-local.md`, `snapshot-spec.md`
  → `stack-snapshot/cmd/stacksnap/main.go` (`exitUnprovisioned=4`, `exitCacheMiss=5`), `autoprovision.go`.
- asset-plane-stays-prod (`DIRECTUS_PUBLIC_BASE_ADDR` left on prod; only `DIRECTUS_BASE_ADDR` re-pointed)
  → `stack-injection/gen_injected_override.py`, `stack-core/gen_override.py`.
- demo default-on / dev opt-in / N=0 `--force` guard → `dev-stack/dev-setdress.sh` (stack-type default +
  N=0 die-without-force).
- structure+rows captured together (M21 fold) → `stack-snapshot/directus/directus.go` (`CapturesStructure:true`),
  `capture/capture.go` (StructureCapturer phase 4b), `manifest/manifest.go` (`Structure` artifact).
- cold-start `--dsn` + AssertPublicOnly (two-phase: AssertPlan + AssertCaptured) + postgres-MCP-not-a-source
  (`source/source.go` Kind enum has no MCP variant; real `COPY` via `pg.Conn`) → `snapshot-cold-start.md`.
- new Directus verify probes → `stack-verify/lib/services.sh` (server/health row), `lib/readiness.sh`
  (`probe_directus_collections()`, container-gated schema expectation).
- teardown reclaims directus container + frees registry slot → `rosetta-demo` down (`ureg_release`),
  `dev-stack` down (`reg_release`).

Fixed inline (KB-1, KB-2): `demo-up` arg-hint advertised a fictional `--full` + `--no-ui`/`--no-setdress` as
`rosetta-demo` CLI flags (they're env vars on `up-injected.sh`: `DEMO_NO_UI` / `DEMO_NO_SETDRESS` /
`DEMO_NO_LOCAL_CONTENT`) → rewrote the hint to the env-var reality (body was already correct). `dev-up`
arg-hint omitted `--local-content` (in parser `dev-stack:88` + body) → added it.

## The directus_files tenancy model + the M25 firewall fix (resume run, 2026-06-13)

The `directus_files` capture is a **referenced-subset / reverse-reference closure** (M23): the table has
**no scope column** and **no forward parent FK** — it is referenced *by* many content rows. The M25 capture
made one thing concrete that the M23 unit tests (synthetic fixtures) couldn't: **a file's tenancy is
INFERRED from what references it**, and the public-only definition for a file must therefore be "referenced
by PUBLIC content **and not** by tenant content".

Prod data shape (2026-06-13, verified via the sanctioned `marco_read` read):
- `directus_files` ≈ 10,480 rows total; `resource` ≈ 1,646 rows (1,338 with `image`, 315 with `file`).
- **Neither `resource` nor `directus_files` has any tenancy column** (`organization_id` / `tenant_id` /
  `private` all absent) — confirmed by `information_schema.columns`.
- Files referenced by **public** content: 185 (sims/paths/roles/sequences) + the resource library.
- Files referenced by **tenant** content: 286. Of these, **150** are *also* referenced by public content
  (shared assets), and **8** are referenced by a **resource** row that a tenant sim also points at.
- Old (broad) closure captured **1417** files; **158** were tenant-referenced → the firewall counted the
  9065-file *uncaptured remainder* (10,480 − 1415) as the "leak" (a separate probe-semantics defect — see
  M25-D5). Fixed closure captures **1257** (158 excluded).

Fix mechanics (`stack-snapshot/directus/media.go` + `firewall/firewall.go` + `cmd/stacksnap/adapters.go`):
- `TenantReferencedFilesFilter()` — a **boolean** `id IN (<UNION of tenant content file-ref ids>)` predicate
  (negated public roots for sims/paths; tenant-parented roles/sequences; `resource` contributes nothing —
  it has no tenancy column, so a resource file is tenant *only* via a tenant sim, already covered).
- `ReferencedFilesFilter()` appends `AND NOT (<that boolean>)` (equivalent to `id NOT IN (...)` here — no
  NULLs: the union columns are `IS NOT NULL` and `id` is the non-null PK).
- The **boolean** shape is reused verbatim by the post-capture leak probe as `WHERE (closure) AND (tenant)`
  (the M25 probe fix), avoiding a uuid-vs-boolean type error the id-list form would cause.
- The firewall `AssertPlan` referenced-subset branch now **requires** `ReferencedSubsetTenantFilter` to be
  declared **and** the closure to **contain** it (the structural guarantee a referenced-subset table can't
  ship without a tenant-leak definition + exclusion). `ReferencedSubsetTenantFilter` is threaded
  `TableSpec → TablePolicy → CountTenantRows`.

**Why the box's local Directus served 0 content before this:** the cache `6cd35278` was rows-only (captured
2026-06-10, **pre-M21**) — no `_structure.sql`, no `directus_files`. With no captured structure, the
bootstrapped-gap stack's schema digest never converged (replay rc=5 cache-miss) → 0 collections. The M25
structure-bearing re-capture (now firewall-clean) fills both gaps, so replay auto-provisions the structure,
converges the digest, loads the rows, and the local Directus serves.

## The three dangling-FK fixes (M25-D6/D7) + the complete FK enumeration

Once the firewall-clean structure-bearing snapshot existed, the live DB-1 replay walked **referential
integrity** that synthetic unit fixtures never exercised — three SQLSTATE 23503 failures, each a dangling
FK to a directus_* table **outside** the captured surface, surfaced one at a time:
1. **`directus_collections.group`** (apply-side, M25-D6) — self-FK to a parent **group collection** (admin-UI
   data-model folder: `simulations`/`job_simulations`/`paths`/`learning_resources`), not served → the
   serve-row render now NULLs `group`.
2. **`directus_files.folder`** (replay-side, M25-D7) — FK to `directus_folders` (file-library UI folder),
   not captured.
3. **`directus_files.uploaded_by` / `modified_by`** (replay-side, M25-D7) — FK to `directus_users` (the
   **prod authoring users**, absent from a fresh per-stack bootstrap).

**Why exactly these three columns and no more (the complete enumeration, via the prod read):** the per-stack
Directus has **two** sources of tables — (a) the **bootstrap** creates the directus_* SYSTEM tables WITH
their full FK graph; (b) the **M21 structure artifact** creates the content tables as **tables + PRIMARY
KEYS only — no foreign keys** (verified: all 26 structure `ALTER TABLE` are `PRIMARY KEY`; the per-stack
directus has **0** content-table FK constraints). So the *only* live FK constraints a replay can violate are
the bootstrap-created ones on `directus_files` — `folder` (→ directus_folders) + `uploaded_by` / `modified_by`
(→ directus_users). The content tables' own `*_foreign` columns (cover/avatar/simulation/user_created/…)
never fire because those FKs don't exist in the per-stack schema. Hence
`NullColumns: ["folder","uploaded_by","modified_by"]` on the directus_files spec is the **complete** set —
no whack-a-mole. All three are organizational / authoring-identity columns with **zero** effect on serving
the asset (which resolves by `id → storage/filename_disk`), so nulling them is loss-free.

The fix mechanism is a general `TableSpec.NullColumns` (capture COPY renders `NULL AS <col>`, the column
stays in position so the load shape is unchanged) + the serve-row `group`-null in the structure render.
