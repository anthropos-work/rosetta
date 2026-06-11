# M21 — Spec notes

Iteration-protocol-specific technical notes (the structure-capture mechanism, the manifest/digest design, the
firewall carve-out, the Directus bootstrap empirics). Accumulates across iters.

## Pre-flight audits — iter-01

**Phase 0b KB-fidelity gate — verdict: YELLOW (discharged 2026-06-11, from fresh in-area evidence).**
Rather than re-running a generic `/developer-kit:audit-kb-fidelity` pass, this gate is discharged against the
**hours-old 4-agent research** that verified the corpus against live `rosetta-extensions` code in exactly M21's area
(the knowledge-audit + code-gap reports synthesized in
[`.agentspace/scratch/roadmap-research-2026-06-11.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-11.md)).
Findings:
- **The contract M21 builds against is faithful.** `corpus/ops/snapshot-spec.md` § the per-stack Directus store fork
  accurately describes the gap ("NOT YET AUTOMATED — the M10 collection-schema gap"), the 4-step recipe, the
  bootstrap empirics, and the exit-4 semantics — all match the live code (`directus/provision.go:102-108`,
  `cmd/stacksnap/main.go:63,72`). Building against it is safe.
- **Known stale areas, already routed to M24 (not M21's contract):** `external_services.md` §Directus (image 10.10.1
  + a local compose snippet + admin/password — all false; the platform compose has no directus service, verified),
  `service_taxonomy.md`, `quick_ops.md`; and the `directus_files` "always captured + replayed" overstatement
  (`snapshot-spec.md:414`) — the code's `media.go` file-ref capture is dead (not in `directus.Surface()`). These are
  YELLOW context, not blockers: M21 will *fix* the `directus_files` reality (wire the capture) and M24 sweeps the docs.

If a triggered tok later redirects M21 into a subsystem this evidence didn't cover, re-run the formal gate then
(per Phase 0b's subsystem-redirect rule).

## Baseline gate-distance — iter-01 (static; live baseline deferred to iter-02)

**The gate is a 6-stage end-to-end pipeline.** Primary metric = furthest stage passing (ordinal 0–6) — measurable
per-tik progress toward a binary gate.

| # | Stage | State today | Evidence |
|---|-------|-------------|----------|
| 1 | **build** — stack-snapshot compiles | PASS (existing) — must stay green as the extension lands | `go build ./...` exit 0 |
| 2 | **bootstrap** — `node cli.js bootstrap` creates 27 `directus_*` system tables in a fresh `directus` schema | implemented; **LIVE-confirmed iter-02** (27 tables created). Needed a fix: the minted admin email was `.local`, which 11.6.1 rejects → fixed to `.example.com` (M21-D1). image `directus/directus:11.6.1` cached | `directus/provision.go:86-100`; iter-02 live run |
| 3 | **structure-apply** — create the user-collection tables so the schema digest converges | **PASSES (iter-04, option A):** applying ALL **26** collections' real DDL on a bootstrapped 11.6.1 → digest = `6cd35278…` exactly. (Caveat: hand-applied `iter-04/structure.sql`; stacksnap code-ification pending, M21-D8.) | iter-04 live run; `iter-04/structure.sql` |
| 4 | **replay-exit-0** — `stacksnap replay --surface directus` COPYs the captured content tables in | **PASSES (iter-04):** exit 0, 9 tables / 10128 rows (simulations=304). Earlier baseline: bootstrapped-but-gap schema → exit 5; empty → exit 4 (M21-D3) | iter-04 live run; `cmd/stacksnap/main.go:359-378` |
| 5 | **boot** — boot Directus on the schema, reachable over HTTP | **PASSES (iter-05):** booted on the harness, `/server/health` 200 | iter-05 live run |
| 6 | **serve-anonymously** — `GET /items/simulations?limit=1` → 200 with a real row | **PASSES (iter-05, DEMONSTRATED):** 200 + real published sim, anonymous. Needs the table PKs + directus_collections registration + a public read permission on Directus's hardcoded policy `abf8a154` (M21-D9) | iter-05 live run |
| — | **GATE automation clause** — *stacksnap* applies the captured structure | **PENDING:** the structure was hand-applied; `STRUCT-M21-codeify` makes stacksnap do it | the gate's "stacksnap applies" wording |

**Furthest stage passing: 6 — DEMONSTRATED end-to-end (iter-05).** All 6 stages pass with the real captured structure:
build → bootstrap → 26-collection structure apply (digest `6cd35278…`) → replay exit 0 (10128 rows) → boot →
anonymous serve (200 + real published sim). **Gate not yet met-by-tooling:** the exit_gate says "stacksnap applies the
captured structure" — the structure/PK/registry/permissions were hand-applied (`iter-04/structure.sql` +
`iter-05/pks.sql` + `iter-05/serve.sql`); `STRUCT-M21-codeify` wires this into stacksnap to flip the gate met.

## iter-05 serve findings (the flagged live-only risk, cracked)

- **PRIMARY KEY is load-bearing (M21-D9):** Directus ignores PK-less collections (`"doesn't have a primary key column
  and will be ignored"` → 403, even admin). The column-only `pg_catalog` DDL omitted PKs; the digest still converged
  (PKs aren't in the column-type digest) and COPY worked, but serving failed. Fix: capture + apply the real PKs
  (`id` ×25, `code` for `languages`). The structure artifact MUST carry constraints, not just columns.
- **Public access:** Directus's hardcoded public policy is `abf8a154-5b1c-4a46-ac9c-7300570f4f17`; bootstrap creates it
  + its `(role=NULL,user=NULL)` directus_access link. Anonymous read = a `directus_permissions` `read` row on that
  policy (fields='*', `status=published` filter for simulations/skill_paths).
- **directus_fields NOT required** for the basic gate — Directus introspects DB columns once a collection is in
  `directus_collections` + has a PK. (The 217 fields add casting/UI metadata, not raw GET capability.)

## iter-02 live findings (Docker, directus/directus:11.6.1)

- **Admin-email validator (M21-D1):** 11.6.1 rejects the `.local` TLD outright — `admin@dev-5.local` AND
  `admin@dev.local` both `FAILED_VALIDATION`; `admin@example.com` / `admin@dev-5.example.com` pass. Hyphens/digits in
  the label are fine. Fix: `admin@<stack>.example.com`. (Supersedes the fix16 hyphen-vs-underscore comment.)
- **Exit-5 vs exit-4 (M21-D3):** bootstrapped directus schema digest = `b4cb55bcee08c76f2c37980da460a683`; prod cache
  key = `6cd35278edbc8a7962053a9d7ebfc480`. Real pipeline (bootstrap-then-replay) → exit **5**; empty schema → exit 4.
- **Structure-apply mechanism (M21-D2):** a Directus schema snapshot is YAML
  `version / directus / vendor / collections / fields / relations`. `node cli.js schema apply <yaml>` creates the
  user table AND the `directus_collections`/`directus_fields` registry rows together (1-collection proof). An empty
  model snapshots as `collections: [] / fields: [] / relations: []`.
- **Digest trap is full-schema-keyed (M21-D5):** `pg.SchemaVersionSQL()` digests every column of every table in the
  `directus` schema → convergence needs the whole schema (system tables at prod's Directus version + all prod content
  collections + exact types) to match. The stage-4 resolution chooses (A) full content-model + version pin, or (B)
  re-key per-surface over only the captured content tables. Tracked `STRUCT-M21-digest-keying`.

**The cache (the row half) is real and complete:** `.agentspace/snapshots/directus/6cd35278…/` — 9 content tables,
~25 MB COPY payloads, `format_version 1`, `public_only`, `predicate=directus-public-published`, every column
enumerated per table. It carries **rows + column lists but NO DDL and NO registry tables** — exactly what the
structure extension must add. The cache key `6cd35278…` is the *source* (prod) directus schema digest; a
bootstrap-only stack digests differently → the structure artifact must converge the target digest before row replay
(the "digest trap").

**`directus_files` is absent from the cache** (9 tables, none `directus_files`) — confirming the dead-code finding
(`media.go:45,55` define the filter/columns; `directus.Surface()` at `directus.go:139` omits the table). M21 wires it.

## Structure-source question (open — resolve in iter-02)

The structure artifact (DDL + registry rows) must be captured from a source that *has* the directus content-model
schema. The row cache was captured from a privileged `primary-read --dsn` (not stored). Candidate sources for the
*structure* capture, to weigh in iter-02:
- **(a) a live directus-bearing `--dsn`** (the prod read; privileged, release-time, may be unavailable now);
- **(b) restore a `pg_dump -n directus`** (the cold-start path — carries full DDL + the registry tables) then `--dsn`;
- **(c) build a self-contained reference**: bootstrap a Directus, create the 9 collections (column lists known from
  the manifest; Directus's `schema snapshot`/`schema apply` YAML is the portable structure carrier), capture
  structure from *that* — no prod access, fully reproducible, and it doubles as the test fixture.
Lean: **(c)** for the build/test loop (self-contained, prod-untouched — honors the user's never-touch-prod
constraint), with the artifact format designed so a future release-time **(b)** capture produces the byte-identical
structure for the real cache.

## iter-03 prod structural read (sanctioned) — the real structure source

Operator sanctioned a bounded read-only structural read of the prod `directus` schema via the wired `postgres` MCP
(the directus content lives in the same prod Postgres). Findings:

- **Full schema = 53 tables:** 27 `directus_*` system + **26 user collections**. The surface captures **9** of the 26.
  The prod schema digest over all 53 tables = `6cd35278edbc8a7962053a9d7ebfc480` (= the row-cache key, confirmed).
- **Real DDL captured for the 9 collections** (exact `pg_catalog` types; column sets match the manifest):
  simulations(37) skill_paths(28) sequences(34) resource(19) roles(17) sim_tasks(11) task_checks(8)
  task_sub_checks(10) sequences_roles(10). Types seen: uuid, json, text, character varying(N), integer, boolean,
  timestamp with/without time zone. Note `sequences_roles.id` is `integer DEFAULT nextval('directus.sequences_roles_id_seq')`
  (a serial — the only non-uuid PK), so the artifact must carry the sequence too.
- **Registry inventory for the 9 collections:** 9 `directus_collections` rows, **217** `directus_fields` rows, **43**
  `directus_relations` rows — **20** of the relations point to collections OUTSIDE the 9 (the M23 referential-closure
  surface; a booted Directus will have dangling relation defs until M23 closes them or they're pruned).
- **Digest-keying (M21-D5 → option B):** a bootstrap (27 system) + 9-collection structure can never equal the
  53-table `6cd35278…` digest → the cache must be re-keyed per-surface over only the captured content tables. Shared
  `pg.SchemaVersionSQL` (taxonomy too) → operator-surfaced architectural fork. See decisions M21-D5/D6.
