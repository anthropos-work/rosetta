---
milestone: M9a
slug: snapshot-framework
version: v1.2 "set dressing"
milestone_shape: section
status: planned
created: 2026-06-05
refined: 2026-06-06
complexity: large
delivers: corpus/ops/snapshot-spec.md (new) + corpus/ops/db-access.md (new) + the /db-query skill + extends corpus/architecture/alignment_testing.md (snapshot-fidelity + public-only gene class)
---

# M9a — Snapshot extension: capture-safe, public-only, manifest-cached framework + `/db-query` port

## Goal
A **dedicated, reusable** `rosetta-extensions/stack-snapshot/` section that **captures a public reference surface
once from a safe (non-primary) prod source, serializes it to a `.agentspace` manifest-cached store, and replays it
per-stack** — with a tested **tenant-data firewall** (never customer data) and an **alignment extension that
measures replay fidelity**. Proven end-to-end on a tiny reference surface (the M0 toy-mirror discipline); the first
*real* surface (taxonomy) is M9b. Ports the **`/db-query`** skill as the prod-read foundation.

## Why this is its own milestone (the M9 → M9a/M9b split)
The 2026-06-06 refinement (5 user notes + live-prod research) loaded the framework heavily: a **dedicated extension**
(note #1), a **production-safe capture-source policy** (note #2), a **tenant-data firewall** (note #3), a
**pluggable `.agentspace` manifest store** (note #4), and the **`/db-query` port** (note #5) — plus the
snapshot-fidelity data-DNA. That's a full framework milestone on its own (the M7a analog). The GB-scale taxonomy
*surface* is its own large milestone (M9b, the M7c analog). Splitting keeps each independently shippable and the
merge surface clean. Research: [`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../../../../.agentspace/scratch/roadmap-research-2026-06-06.md).

## Why section (not iterative)
The deliverables are writable up front: the extension layout, the capture/replay contract + format, the
capture-source policy, the firewall, the store + manifest, the CLI, the fidelity-gene class, and the skill port.
The fidelity gate is a per-surface acceptance check, not an emergent-path gate.

## Scope
- **In:**
  - **(note #1) The dedicated `rosetta-extensions/stack-snapshot/` section** — a sibling of `clerkenstein`,
    `demo-stack`, `stack-injection`, `stack-core`, `stack-seeding`, `alignment`. It owns **capture + serialize +
    store + replay** and the **`stacksnap` CLI** (`capture` / `replay` / `status`). `stack-seeding` *consumes* a
    snapshot at its existing DAG node (`… → taxonomy/content (snapshot) → activity`) — capture (a privileged prod
    **read**) is decoupled from seeding (per-stack **writes**). Reusable by staging/tests too.
  - **The snapshot contract + portable format** — per-table `COPY` payloads + a `manifest.json` describing how a
    surface is serialized, versioned (schema-version pinned), and addressed.
  - **(note #2) The production-safe capture-source policy** — a **source-pluggable** refresh, fully automatic,
    that never blocks the hot primary:
    1. **cache-hit (default):** replay from the existing `.agentspace` snapshot — *zero* prod read;
    2. **refresh — pluggable source, ordered** (investigated 2026-06-06; **no read replica today**, eu-west-1
       instance `terraform-2024…`, no local AWS creds): **(1, default) ingest an existing prod `pg_dump`** (the team
       already produces these for staging → *zero new prod load*) → **(2, fallback) safe throttled primary read**
       via the existing `marco_read` access — PostgreSQL **MVCC means a read-only `SELECT`/`COPY` never blocks
       writers**, so off-peak + throttled + public-only (data, not indexes) is tolerable. **Zero-primary-impact
       upgrades** once AWS/infra is wired: **(3) restore-from-snapshot** to a throwaway instance, or **(4) a
       provisioned read replica**. (M9a-D3, user 2026-06-06.)
    3. **bounded read session:** `SET TRANSACTION READ ONLY`, `statement_timeout` + `idle_in_transaction_session_timeout`,
       modest `work_mem`, `COPY (SELECT … WHERE org_id IS NULL) TO STDOUT` streamed to disk (keyset-chunked for the
       biggest tables); **catalog-first dry-run** sizes the read before any data flows.
    This adds the **read half** the M7a isolation guard lacks (it guards writes only today).
  - **(note #3) The tenant-data firewall** — `AssertPublicOnly` (read-side analog of `AssertClean`): every captured
    table either has no org column (pure reference) **or** is filtered `organization_id IS NULL`; post-capture it
    **hard-fails** if any captured row carries a non-null tenant scope. Backed by a **public-only / provenance gene**
    in the data-DNA. Safety contract, not plumbing.
  - **(note #4) The `.agentspace` manifest-cached, pluggable store** — payload (large, gitignored) under the
    **workspace-level** `.agentspace/snapshots/<surface>/<schema-ver|hash>/…` (M9a-D5 — one shared cache captured
    once + replayed by every stack, **not** per-stack); a small **`manifest.json`** (surface, source, schema
    version, capture ts, row counts, **public-only assertion result**, checksum, payload location, format version)
    that drives **cache-hit vs stale→refresh**. A `SnapshotStore` interface with a `localfs` backend now; the
    **cloud/S3 backend is the named v1.3 swap** (the manifest already addresses by location). Replaces the original
    "commit a golden into the extensions repo" (no GB blobs in any git repo).
  - **The data-DNA extension** — a new surface status `snapshot-seeded` + a **snapshot-fidelity gene class**
    (source-vs-replay row-count / structural conformance / referential integrity / **embedding-dimension integrity**)
    + the **public-only gene** — added to the `datadna` harness; `datadna` recognizes snapshot surfaces.
  - **(note #5) The ported `/db-query` skill** — `.claude/skills/db-query/SKILL.md` (schema reference verified
    against live prod 2026-06-06) + `corpus/ops/db-access.md`, documenting **both** connection paths (the wired
    `mcp__postgres__query` tool **and** `~/.pgpass` + Tailscale + `psql`). The read foundation capture builds on.
  - **A tiny reference surface** proving capture→store→replay→fidelity-gate end-to-end, independent of taxonomy.
- **Out:**
  - The taxonomy surface (M9b) and the Directus content surface (M10).
  - Recipes / presets / corpus product layer (M11).
  - The cloud/S3 store backend, AI-generated content, external shareability (v1.3).

## Depends on
v1.1's **M7a** (isolation guard + the `COPY` perf path + the seeder DAG) + **M7b** (the data-DNA harness this
extends). **Parallel with:** none (gates M9b + M10 + M11).

## Open questions (resolve during build)
- **Capture source** (M9a-Q1, investigated + double-checked 2026-06-06): standalone RDS (not Aurora), eu-west-1,
  `terraform-2024…`; **no read replica today**; no local AWS creds. Default source = **prod-dump ingest** or a
  **safe throttled primary read** (MVCC = no write blocking); zero-impact upgrades (restore-from-snapshot / a
  provisioned replica) need eu-west-1 AWS/terraform access. Stays source-pluggable. See `decisions.md` M9a-Q1.
- The **manifest schema** + the cache-staleness rule (schema-version mismatch and/or checksum) that triggers a refresh.
- **Embedding capture**: carry pgvector vectors verbatim but **rebuild the index on replay** (don't carry the
  ~689 MB index) — confirm the replay rebuild cost is acceptable.
- The `SnapshotStore` interface shape so the **v1.3 cloud swap** is a clean backend change.

## KB dependencies (read as contract)
- `corpus/ops/seeding-spec.md` (the seeding framework + the 3-layer isolation boundary + the perf path + the DAG node)
- `corpus/architecture/alignment_testing.md` (the M0 alignment discipline + M7b's data dimension to extend)
- `corpus/ops/staging_from_dump.md` (the full-dump **anti-pattern** to contrast: selective + public-only + safe-source)
- the source `db-query` skill (ported) — for the schema map + connection model

## Delivers → knowledge/corpus/ops/snapshot-spec.md (new) + corpus/ops/db-access.md (new) + the /db-query skill + corpus/architecture/alignment_testing.md
- **`corpus/ops/snapshot-spec.md`** (net-new): the `stack-snapshot` extension; the capture/serialize/replay
  contract + format; the capture-source policy; the tenant-data firewall; the `.agentspace` manifest store +
  pluggable backend; per-stack injection + versioning.
- **`corpus/ops/db-access.md`** (net-new) + **the `/db-query` skill** (note #5).
- **extends `corpus/architecture/alignment_testing.md`**: the snapshot-fidelity + public-only gene class alongside
  the behavioral (v1.0) + structural data-DNA (v1.1 M7b) dimensions.
</content>
