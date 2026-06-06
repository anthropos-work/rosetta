---
name: demo-snapshot
description: Set-dress a demo/dev stack with the real PUBLIC reference library — replay the captured taxonomy + Directus content snapshots into the stack so the catalog + content templates are real (not placeholder). Drives the stacksnap CLI (replay / capture / status). Use after /demo-up and before /demo-seed for a full-fidelity world, or when asked to snapshot / set-dress / replay the catalog into a demo.
argument-hint: [N] [replay|capture|status] [--surface taxonomy|directus] [--dsn DSN] [--dry-run]
---

# Demo Snapshot — stamp the real public library into a stack (the v1.2 "set dressing" layer)

Replays the real **public** reference library — the ~60K-skill / 18K-role taxonomy and the global simulation /
skill-path content templates — into `demo-N` (or `dev-N`) so the catalog view shows **real** skills and the seeded
sessions link to **real** templates. It drives the `stacksnap` CLI: `replay` (the headline verb, per-stack),
`capture` (the rare prod-read maintenance op), `status` (list the cache). Source of truth:
[`corpus/ops/snapshot-spec.md`](../../../corpus/ops/snapshot-spec.md). The set-dressing recipe:
[`corpus/ops/demo/recipe-snapshot-world.md`](../../../corpus/ops/demo/recipe-snapshot-world.md).

## Where this sits in the flow
`/demo-up N` → **`/demo-snapshot replay N`** → `/demo-seed N` → log in. The snapshot is **stack-global** public
reference data (independent of which org you seed), it is **optional** (skip it for a quick structural-only world —
the seeder degrades gracefully to an empty catalog + free content refs), and it is almost always a **cache-hit**
(zero prod read — captured once per release, replayed by every stack).

## Mission

1. **Read the spec** — `corpus/ops/snapshot-spec.md` (the capture/replay contract, the read-side **tenant-data
   firewall**, the cache-first store, the fidelity gate). Confirm the target is a **non-prod** stack (`demo-N` /
   `dev-N`, never production).
2. **Confirm the stack is up + migrated** — `/demo-up N` first if needed, so the `skiller` + `directus` schemas
   exist as replay targets.
3. **Build the tool** (gitignored at `stack-demo/rosetta-extensions/stack-snapshot/`; canonical source is the
   `.agentspace/rosetta-extensions/` authoring copy):
   ```bash
   SN=stack-demo/rosetta-extensions/stack-snapshot
   go build -o /tmp/stacksnap "$SN/cmd/stacksnap"
   ```
4. **Run the requested verb:**

   **`replay`** (default) — stamp the cached snapshot(s) into the stack. With **no `--surface`**, replay **both
   real surfaces** (`taxonomy` then `directus`) — `stacksnap replay` itself requires one `--surface` per call, so
   loop them:
   ```bash
   for s in taxonomy directus; do
     /tmp/stacksnap replay --surface "$s" --stack demo-N
   done
   # or a single surface:
   /tmp/stacksnap replay --surface taxonomy --stack demo-N
   ```
   Replay resolves **cache-hit vs stale** against the stack's live schema, **verifies every payload checksum**
   before writing, bulk-`COPY`s in dependency order, and **rebuilds any pgvector index** (the ~689 MB
   `skill_embeddings` IVFFLAT is reindexed, not transported). If the cache is **stale/missing**, replay tells you
   to capture first → do path B. The `directus` content surface additionally needs its **per-stack Directus
   booted** against the stack's own `directus` schema (bootstrap → replay → boot, offset-port container, **never**
   `content.anthropos.work`) — see the recipe.

   **`capture`** (rare maintenance) — only when `status` shows the surface is **missing or stale** (the platform
   schema moved). Reads the public surface **once** over `--dsn` from a safe source, firewalls it, caches it:
   ```bash
   /tmp/stacksnap capture --surface taxonomy --dsn "$SAFE_DSN" --dry-run   # size + assert firewall plan, NO read
   /tmp/stacksnap capture --surface taxonomy --dsn "$SAFE_DSN"             # the real capture
   /tmp/stacksnap capture --surface directus --dsn "$SAFE_DSN"            # directus = a schema in the same DB
   ```
   `--dry-run` sizes the surface catalog-only and asserts the firewall plan **without reading a data row** — the
   cheap pre-flight. `--source dump-ingest|primary-read` overrides the default precedence (both read over `--dsn`;
   they differ only in what the DSN addresses — a restored staging `pg_dump` vs the prod read endpoint).

   **`status`** — list the cache (surface, schema version, rows, source, capture time):
   ```bash
   /tmp/stacksnap status
   ```
5. **Verify** — replay prints `replayed "<surface>" into demo-N: <T> table(s), <R> row(s) loaded [, reindexed …]`.
   Optionally gate fidelity (captured source vs replayed stack):
   ```bash
   SS=stack-demo/rosetta-extensions/stack-seeding
   SNAP=.agentspace/snapshots
   /tmp/datadna measure-snapshot --stack demo-N --dna "$SS/dna/data-dna.json" \
     --manifest "$SNAP/taxonomy/<ver>/manifest.json" --manifest "$SNAP/directus/<ver>/manifest.json"
   ```
   With both surfaces replayed + gated, `datadna catalog` reads **100%** coverage (both formerly-`waived` surfaces
   promoted to `snapshot-seeded`, nothing waived).

## Safety (the load-bearing part)
**Capture is a privileged prod READ; replay is a per-stack WRITE — both are hard-guarded.**

- **Public-only firewall (read side).** Capture transports **only** public reference rows — **never any
  customer/tenant data**. The per-surface predicate (`organization_id IS NULL` for taxonomy; `private=false AND
  tenant_id IS NULL AND status='published'` for Directus) runs at PLAN time and POST-capture; a single leaked
  tenant row **aborts the capture and writes nothing**.
- **Safe source, bounded read.** Capture connects DIRECTLY over `--dsn` to a safe source (default: a restored
  staging `pg_dump`; sanctioned fallback: a throttled, off-peak, read-only primary read — MVCC = no writer
  blocking) inside a bounded `READ ONLY` session. It **never** runs through the platform services and **never**
  writes anywhere.
- **Replay is per-stack only.** It writes the per-stack-isolated `skiller` / `directus` Postgres (offset port,
  class `PerStackIsolated`); it can **never** write the shared prod Directus, the prod S3 bucket, or live Clerk.
- **Never** point `capture --dsn` at production for a write, and **never** boot a per-stack Directus against
  `content.anthropos.work` — `EnvContract.Validate()` hard-rejects it.

## Defaults
- Store root: `--store` → `STACKSNAP_STORE` env → `<workspace>/.agentspace/snapshots` (gitignored; GB blobs never
  enter git). The cloud/S3 store is the named **v1.3** swap.
- Base DSN (replay): `postgres://postgres@localhost:5432/postgres` with the port replaced by the stack offset.
- Exit codes: `0` ok · `1` firewall/capture/replay error (e.g. a tenant-data leak aborted capture) · `3` usage.
