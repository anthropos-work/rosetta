---
milestone: M209
slug: rext-reground
version: v2.1 "quick change"
milestone_shape: section
status: planned
created: 2026-07-08
last_updated: 2026-07-08
complexity: medium
depends_on: M208
delivers: rext tag v2.1 (the re-grounded tooling) + a recaptured public-taxonomy snapshot
issues: rext tooling still queries skiller.<table> + probes a skiller container/schema — breaks on the merged platform
---

# M209 — rext tooling re-ground

## Goal
Re-point every rext tool that queries the old `skiller` schema or expects the skiller service to the merged
reality (`public.*`, no skiller container/subgraph), recapture the public-taxonomy snapshot from merged-prod, and
tag a new rext (`v2.1`).

## Why section
The blast radius is fully enumerated (design workflow `wf_08b6bf4a`, file:line): a tiny-but-load-bearing surface,
mostly a mechanical `skiller.<table> → public.<table>` prefix swap plus two well-scoped non-mechanical items (the
cache-key digest + the capture column-list). Enumerable up front → `section`.

## Scope
### In — stack-snapshot
- Flip **`stack-snapshot/taxonomy/taxonomy.go:43 const Schema "skiller" → "public"`** — the ONE load-bearing code
  change; it flows through `Surface()` into the capture WHERE SQL, payload filenames, the manifest per-table
  `Schema`/`PublicVia`, AND the replay COPY target (replay loads into `tb.Schema` from the manifest). Re-grounds
  capture *and* replay.
- Update the 2 load-bearing `taxonomy_test.go` assertions (`PublicVia == "public.skills"` /
  `"public.job_roles+public.skills"`).
- **Narrow the `pg.SchemaVersionSQL` staleness digest** to the surface's enumerated tables (Risk 1 — post-merge
  the digest spans the whole `public` monolith → the taxonomy cache key thrashes on ANY app migration). Consider
  keying on the enumerated table set (or app migration rev) instead of `md5(all columns of all tables in schema)`.
- **Verify the capture SELECT column list vs merged prod** (Risk 2 — the one non-mechanical bit): the column names
  may differ post-merge (`embedding → small_embedding3`, generated `ts_search`/GIN opclasses now `extensions.`-
  qualified). Reconcile the capture column mapping.
- Keep `AssertPublicOnly` (the runtime net that rejects any captured row with non-null `organization_id`) + add a
  post-capture assertion that `public.skills WHERE organization_id IS NULL` ≈ **42,763** rows (catches an empty /
  over-broad capture).
- **Recapture** the public taxonomy from merged-prod (`public.*`) into `.agentspace/snapshots/`; bump the capture
  version; the M45 batch cache re-keys.
- Firewall (`firewall/firewall.go`): **no code change** — schema-agnostic; `organization_id IS NULL` still
  isolates public taxonomy (adversarially confirmed at design — `skiller_mixins.OrganizationIDMixin` ports the
  tenant boundary 1:1; paired partial-unique indexes still discriminate global vs per-tenant). Optional comment
  refresh.

### In — stack-seeding
- Re-point the **5 real-SQL files** `skiller.* → public.*`, keeping the `organization_id IS NULL` public-pool
  predicate verbatim: `seeders/taxonomyref.go`, `seeders/skillref_named.go` (the shared `namedSkillSelect` const →
  also fixes `curated_pools.go` + `ai_readiness_config.go`), `seeders/jobroleref.go`,
  `seeders/taxonomy_snapshot.go` (+ the audit Entry `Schema:"skiller"` value), `dna/fidelity_probe.go`.
- Drop `skiller` from `isolation/isolation.go` (the per-stack-Postgres schema note; `skill_embeddings` now lands
  in `public`); re-ground `dna/data-dna.json` golden (taxonomy-surface `schema` + FK `ref_schema`).
- Rename the **111 fake-Conn test string-matchers** (`strings.Contains(sql, "FROM skiller.skills")` etc.) to
  `public.*` in lockstep — else the fakes stop recognizing the queries and tests go red.
- `services/ai/ai.go` + `ai_test.go` are **comment-only** (the wrapper reads env-var *names* from `platform/.env`;
  a schema merge doesn't rename `.env` keys) — reword the stale attribution (`skiller`'s Azure deployment → app).

### In — small modules
- `stack-verify/lib/readiness.sh` — remove `skiller` from `probe_postgres_schemas` (else false-fail on the
  now-dropped schema); `services.sh` — remove the `skiller | anthropos-skiller-1 | :8085` container probe.
- `demo-stack/up-injected.sh` — drop `skiller` from `INJECT_SVCS` (else it clones/builds a gone repo → build-loop
  break) + reword the "5 Go services that verify Clerk tokens" → 4.
- `stack-secrets` — **no real DNA change** (secret-dna.json has no skiller repo; the skiller/BUNNY refs are
  synthetic Go test fixtures) — optional fixture re-point only. `clerkenstein` + `alignment` — clean, no change.

### Out
- The stack re-sync (M208); the corpus doc bodies (M210); live bring-up (M211).

## Tag + consume
Build + `go test ./...` the authoring copy; **tag a new rext (`v2.1`)**; prepare the per-stack consumption re-pin
(the box-level `.agentspace/rext.tag` flips at M209/M211).

## KB dependencies
`corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/safety.md` (read as contract; their bodies
re-point in M210, in lockstep with this change).

## Done-bar
- rext authoring copy builds + tests green with 0 `skiller.<table>` queries in any production path; the recaptured
  snapshot loads `public.*`; the ~42,763-row assertion passes; the cache-key digest is narrowed; rext tagged `v2.1`.
