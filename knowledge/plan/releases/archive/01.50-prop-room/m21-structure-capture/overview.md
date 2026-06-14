---
milestone: M21
slug: structure-capture
version: v1.5 "prop room"
milestone_shape: iterative
status: archived
created: 2026-06-11
last_updated: 2026-06-13
complexity: large
exit_gate: "on a scratch offset Postgres + a bootstrapped directus/directus:11.6.1: stacksnap applies the captured structure -> `stacksnap replay --surface directus` exits 0 -> a booted Directus serves a captured public simulation over HTTP to an ANONYMOUS reader (GET /items/simulations?limit=1 -> 200 with a real row)"
iteration_protocol_ref: corpus/architecture/alignment_testing.md
delivers: corpus/ops/directus-local.md (net-new) + the capture-side structure extension + directus_files capture + redefined stacksnap exit codes
backlog_refs: NEW-1 (the M10 collection-schema gap)
---

# M21 — Structure capture: close the collection-schema gap

## Goal
Make the snapshot carry the content-model **structure** (the user-collection table DDL + Directus's
`directus_collections`/`directus_fields`/`directus_relations` registry rows) alongside the rows, captured atomically
from the same sanctioned source — so the `directus` replay stops failing with `stacksnap` exit 4 and a freshly
bootstrapped Directus can actually serve the captured catalog.

## Exit gate (observable, machine-verifiable)
On a scratch offset-port Postgres + a bootstrapped `directus/directus:11.6.1`: `stacksnap` applies the captured
structure → `stacksnap replay --surface directus` **exits 0** → a booted Directus **serves a captured public
simulation over HTTP to an anonymous reader** (`GET /items/simulations?limit=1` → 200 with a real row). Today this
chain dead-ends at the print-only `provision.go:108` placeholder.

## Why iterative (not section)
The implementation path is genuinely uncertain — Directus's anonymous-read permissions, the registry-row firewall
carve-out, and the cache-digest convergence are undesigned territory that only breaks live (the analogous fix16 cost
+479 lines of empirical correction). A fixed `In:` checklist would be speculative; the gate is the commitment, the
path emerges from each tik's evidence. Build with `/developer-kit:build-mstone-iters`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `prop-room-m21` → consume): the capture-side structure extension, the
  `directus_files` ref capture, the cache-keying fix, the redefined exit codes.
- **`rosetta`**: the net-new `corpus/ops/directus-local.md` (the per-stack Directus spec — bootstrap empirics +
  structure-capture model + version-skew rule).

## Scope (the known shape; refined per-iter)
- **In (`rosetta-extensions`):**
  - The **capture-side structure extension** — capture the user-collection DDL (a `pg_dump --schema-only`-equivalent
    over the `directus` schema from the sanctioned `--dsn`; zero DDL-capture code exists today) + the three registry
    tables filtered to the **public content model** (a collection-name allow-list).
  - A **structural-metadata admissibility class** in the firewall — registry rows are `directus_*` system tables,
    today excluded by the letter of `AssertPublicOnly`; they carry no tenant/private columns, so they need an
    explicit "structure, not tenant data" classification (extend, never loosen, the firewall).
  - The **cache-keying fix** so a structure-less stack can converge to the source digest — apply structure **before**
    the row replay (the row cache is keyed by the *target* schema digest, so a bootstrap-only stack can never
    cache-hit by construction today; exit 5 even with a warm cache).
  - **Redefined `stacksnap` exit-4 semantics** — today "schema missing → a capture can't help"; now the structure
    artifact IS what provisions the schema.
  - **Wire the `directus_files` ref capture** — the docs claim it ships (`snapshot-spec.md:414`) but `media.go`'s
    filter/columns are dead code (no `directus_files` TableSpec); needed for the real-image asset refs.
- **Out:** executing the recipe at bring-up (M22); the env re-point + referential closure (M23); **S3 blob bytes**
  (stays backlog, DEF-M10-01 — refs + prod-link assets are the floor).

## Depends on / parallel with
Depends on: none (first milestone). Parallel with: none (the foundation the rest builds on).

## Open questions (resolve per-iter)
- DDL source: `pg_dump -s` shell-out vs `information_schema` catalog reconstruction (lean: `pg_dump --schema-only
  -n directus`, already on the sanctioned restore-a-dump `--dsn` path, simplest correct).
- Manifest carries structure as an **additive field** (the `Predicate` precedent, no format bump) vs a sibling
  artifact keyed by source digest (lean: decide in iter-01 against the digest-convergence constraint).
- Prod-Directus-version vs local-11.6.1 **skew** policy (lean: record the source version in the manifest, pin the
  local image, warn on mismatch).

## KB dependencies
`corpus/ops/snapshot-spec.md` (the store-fork + capture-source + manifest sections), `corpus/ops/snapshot-cold-start.md`
(the `--dsn` source the structure capture rides), `corpus/ops/safety.md` (the firewall admissibility classes).

## Re-scope trigger
If 5 consecutive toks fail to produce a viable path to anonymous serving (e.g. Directus's permission model proves
to need a running-instance API call no capture can pre-stage), escalate to a user-strategic-replan — the interim
options (auto-heal / full-taxonomy capture) may become the deliverable instead of the full close.

## Risk (prod-safety — blocks-prod-safety)
Structure capture is still a **prod read** — it must stay behind the M9a capture-source policy (read-only, bounded,
operator-confirmed, public-only) and the `AssertPublicOnly` firewall, now **extended** (not loosened) to admit
structural metadata. The dropped pg_dump-file-reader (M9b-D9) must **not** resurrect — `TestDroppedDumpFlagStaysGone`
pins it gone.
