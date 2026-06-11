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
| 2 | **bootstrap** — `node cli.js bootstrap` creates 27 `directus_*` system tables in a fresh `directus` schema | implemented (print-only recipe) + empirically pinned at fix16; image `directus/directus:11.6.1` cached locally | `directus/provision.go:86-100`; `docker images directus/directus:11.6.1` present |
| 3 | **structure-apply** — create the 9 user-collection tables + register them in `directus_collections`/`fields`/`relations` | **THE GAP — unbuilt** | `provision.go:102-108` literal placeholder `"# NOT YET AUTOMATED"` |
| 4 | **replay-exit-0** — `stacksnap replay --surface directus` COPYs the 9 tables in | exits **4** today (target `directus` schema empty → `ErrEmptySchema`) | `cmd/stacksnap/main.go:283-284`, `pg/pg.go:41,244` |
| 5 | **boot** — boot Directus on the schema, reachable over HTTP | recipe print-only (step 4) | `provision.go:115-123` |
| 6 | **serve-anonymously** — `GET /items/simulations?limit=1` → 200 with a real row | blocked behind 3–5 | the gate |

**Furthest stage passing today (static): 2.** The pipeline dies at stage 3 (structure-apply). iter-02 establishes
the *live* baseline (stand up a throwaway Postgres, bootstrap a real Directus, confirm replay exits 4).

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
