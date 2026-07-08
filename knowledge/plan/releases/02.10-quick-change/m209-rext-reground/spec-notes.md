# M209 — Spec notes

## Pre-flight audits — stack-snapshot (first section)
KB-fidelity audit: **YELLOW** (report `kb-fidelity-audit.md`). Pre-flip doc staleness only (KB-1/2/3 → M210).
Topic→doc→code triples:
- snapshot capture+replay → `corpus/ops/snapshot-spec.md` → `.agentspace/rosetta-extensions/stack-snapshot/`
- taxonomy-ref seeding → `corpus/ops/seeding-spec.md` → `.agentspace/rosetta-extensions/stack-seeding/`
- firewall public predicate → `corpus/ops/safety.md` → `.../stack-snapshot/firewall/firewall.go`
- merged-schema contract → M208 spec-notes + merge fact-sheet (VERIFIED: `public.*`, `extensions.vector(1536)`,
  `extensions.gin_trgm_ops`, public skills = 42,790).

## stack-snapshot: taxonomy.go const flip
`const Schema = "public"` (was `"skiller"`). Flows through `Surface()` → capture WHERE, payload filenames,
manifest per-table `Schema`/`PublicVia` (`skillsScope`/`rolesScope` use `ParentSchema: Schema`), and the replay
COPY target. Tests: `PublicViaLabel` → `public.skills` / `public.job_roles+public.skills`; registry schema
`public`. Code-comment `skiller.*` examples refreshed across capture/firewall/store/source/surfaces.

## Cache-key digest narrowing (Risk 1)
`pg.SchemaVersionSQL(tables)` — when `tables` non-empty, adds `AND table_name::text = ANY($2::text[])`. New
`Surface.VersionTables()`: `nil` for a `CapturesStructure` surface (directus → whole-schema, so a new dynamic
collection still invalidates); the enumerated tables for a row surface (taxonomy → stable vs the merged `public`
monolith's app-migration churn). Threaded through `Capturer.SchemaVersion(ctx,schema,tables)` + `captureConn` +
`provisionConn` + all 7 fakes; used at BuildPlan, autoprovision re-probe, and main.go target-probe (so both
sides of the cache-key comparison match). `ErrEmptySchema` now also means "none of the surface's tables exist".

## Capture SELECT column-list reconcile (Risk 2) — VERIFIED, no change
Inspected `docker anthropos-postgresql-1` (merged, schema-only, 0 rows) for all 10 taxonomy tables: every
enumerated column present + unchanged; `small_embedding3` still the vector column (now typed `extensions.vector`);
`ts_search` still excluded. The `extensions.`-qualified types NEVER surface in the tooling — `buildPublicSelect`
is names-only (COPY renders vectors as text) and replay is `REINDEX TABLE` (rebuilds existing indexes by name).
So no column-list change; the schema flip suffices. Added `TableSpec.MinRows` + a `capture.Run` floor
(`skills` >= 40000; ~42,790 in prod) that aborts before any store write on an empty/under-capture.

## Recapture from merged-prod — Fate-3 → M211
Operationally gated: no local COPY-byte capture source (values-blind — no `marco_read`/prod-read DSN in
`.agentspace/secrets` or `platform/.env`; merged `stack-dev` PG has 0 taxonomy rows; `postgres` MCP = JSON not
COPY bytes). Existing cache `taxonomy/c75ce94…` is the stale pre-merge skiller-keyed capture. Recapture pinned to
M211's first tik (exit gate already names it; added a "Pre-surfaced recapture prerequisite" note). See decisions.md.

## stack-seeding: 5 real-SQL files + isolation + data-dna.json
Uniform `skiller.` → `public.` across 24 files (production SQL: `taxonomyref.go`, `skillref_named.go`'s shared
`namedSkillSelect` const [→ curated_pools.go + ai_readiness_config.go], `jobroleref.go`, `taxonomy_snapshot.go`
[+ its audit `Entry.Schema`], `dna/fidelity_probe.go`; dotted comments; and the fake-Conn test matchers — all
move together). `organization_id IS NULL` preserved verbatim. `data-dna.json` golden: `schema`/`ref_schema` →
public. `isolation.go`: drop `skiller` from the schema list; embeddings now in `public`. `services/ai/*`:
comment-only reword (env-var names unchanged; Azure eastus2 deployment now app's).

## Test string-matcher rename (lockstep)
Go: swapped exact `"skiller"` schema-value fixtures → `"public"` in the `dna` package (the bare-value fixtures the
dotted swap missed — else the probes query `public.*` but fixtures assert `skiller.*`). All seeding tests GREEN.
Python (stack-verify): the 15-service registry model re-grounded to 14 (counts, BASES port map, scope fixtures,
fake schema-probe output) — 104/104 OK. Generic skiller-as-example-token / sample-repo tests left as-is.

## Small modules: readiness.sh / services.sh / up-injected.sh
readiness.sh: drop `skiller` from `probe_postgres_schemas`. services.sh: remove the `anthropos-skiller-1 :8085`
probe. up-injected.sh: drop `skiller` from INJECT_SVCS + verify_svcs; 5→4. Also migrate-demo.sh (drop the
`CREATE SCHEMA skiller` + `skiller:skiller` migration pair — same demo-bring-up break class, would trip M211)
+ GUIDE.md. stack-secrets left (synthetic waived-repo fixture, optional).

## rext tag
Tagged `quick-change-m209` (pattern `{codename}-m{N}`). 3 commits since `00eef00`. The `v2.1` release roll +
`.agentspace/rext.tag` consumption re-pin happen at close-release / M211 — NOT now.
