---
name: stack-snapshot
description: Set-dress a stack (dev-N or demo-N) with the real PUBLIC reference library — replay the captured taxonomy catalog into the stack so it's real (not placeholder). The Directus content surface replays into a per-stack Directus that v1.5 M21–M23 now EXECUTES on a --local-content stack (demo default-on; dev opt-in) so content is self-contained (replay exits 0); a stack without --local-content reads content live from prod (the fallback, directus replay exits 4). Drives the stacksnap CLI (replay / capture / status). Use after the stack is up and before seeding for a full-fidelity catalog, or when asked to snapshot / set-dress / replay.
argument-hint: [dev-N|demo-N] [replay|capture|status] [--surface taxonomy|directus] [--dsn DSN] [--dry-run]
---

# Stack Snapshot — stamp the real public library into a stack (the v1.2 "set dressing" layer)

Replays the real **public** reference library — the ~60K-skill / 18K-role taxonomy and the global simulation /
skill-path content templates — into `dev-N` or `demo-N` so the catalog view shows **real** skills and the
seeded sessions link to **real** templates. It drives the `stacksnap` CLI: `replay` (the headline verb,
per-stack), `capture` (the rare prod-read maintenance op), `status` (list the cache). Source of truth:
[`corpus/ops/snapshot-spec.md`](../../../corpus/ops/snapshot-spec.md). The set-dressing recipe:
[`corpus/ops/demo/recipe-snapshot-world.md`](../../../corpus/ops/demo/recipe-snapshot-world.md).
(Formerly `/demo-snapshot`, now naming both stack types as first-class targets.)

> **Two `stack-snapshot` namespaces, kept distinct.** This **skill** (`/stack-snapshot`) drives the `stacksnap`
> CLI. The extensions **section** named `stack-snapshot` (`rosetta-extensions/stack-snapshot/`) is where that
> CLI is built. The skill operates the tooling; the section name inside the repo is unchanged.

## Where this sits in the flow
`/dev-up N` or `/demo-up N` → **`/stack-snapshot N replay`** → `/stack-seed N` → log in. (For a `dev-N`,
`/dev-up` already runs this set-dress pass by default — use this skill to re-run or refresh it.) The snapshot
is **stack-global** public reference data (independent of which org you seed), it is **optional** (skip it for
a quick structural-only world — the seeder degrades gracefully to an empty catalog + free content refs), and
it is almost always a **cache-hit** (zero prod read — captured once per release, replayed by every stack).

## Mission

1. **Read the spec** — `corpus/ops/snapshot-spec.md` (the capture/replay contract, the read-side **tenant-data
   firewall**, the cache-first store, the fidelity gate). Confirm the target is a **non-prod** stack (`dev-N` /
   `demo-N`, never production).
2. **Confirm the stack is up + migrated** — `/dev-up N` or `/demo-up N` first if needed, so the `skiller` +
   `directus` schemas exist as replay targets.
3. **Build the tool** (gitignored at `stack-<role>/rosetta-extensions/stack-snapshot/`; canonical source is the
   `.agentspace/rosetta-extensions/` authoring copy). Use the matching per-stack clone for the target
   (`stack-dev/` for a dev-N, `stack-demo/` for a demo-N):
   ```bash
   SN=stack-demo/rosetta-extensions/stack-snapshot   # or stack-dev/... for a dev-N
   go build -o /tmp/stacksnap "$SN/cmd/stacksnap"
   ```
4. **Run the requested verb:**

   **`replay`** (default) — stamp the cached snapshot(s) into the stack. With **no `--surface`**, replay **both
   real surfaces** (`taxonomy` then `directus`) — `stacksnap replay` itself requires one `--surface` per call, so
   loop them:
   ```bash
   for s in taxonomy directus; do
     /tmp/stacksnap replay --surface "$s" --stack dev-N      # or demo-N
   done
   # or a single surface:
   /tmp/stacksnap replay --surface taxonomy --stack dev-N
   ```
   Replay resolves **cache-hit vs stale** against the stack's live schema, **verifies every payload checksum**
   before writing, bulk-`COPY`s in dependency order, and **rebuilds any pgvector index** (the ~689 MB
   `skill_embeddings` IVFFLAT is reindexed, not transported). **Read the exit code (fix16):** `4` = the stack's
   target schema is **missing/empty** — provision the STACK (a capture will NOT help); `5` = **no cached snapshot
   at the stack's schema digest** — either the cache is empty/outdated (run a `capture` — see below), or the
   stack's schema **diverged** from the captured source (`stacksnap status` compares digests). The `taxonomy`
   surface replays straight away (its `skiller` schema exists from migration). The `directus` content surface
   additionally needs its **per-stack Directus** against the stack's own `directus` schema (**bootstrap →
   apply-structure → replay → boot**, 4 steps). Since **v1.5 M21–M23 this is automated** on a
   **`--local-content` stack** (demo **default-on**; dev **opt-in** via `--local-content`; `N=0` behind
   `--force`): the bring-up's set-dress pass EXECUTES the recipe — M21 captures the content-model structure and
   **auto-provisions** the bootstrap gap (the content-schema step is no longer a manual gap), M22 boots the
   per-stack Directus compose service, and M23 cuts `cms`'s `DIRECTUS_BASE_ADDR` over to it — so the **directus
   replay exits 0** and the stack is content-self-contained (asset plane stays on prod). On a stack **without**
   `--local-content` (the **fallback**), there is no per-stack Directus: the directus replay skips with
   **exit 4** and the stack reads public content **live from prod**. Either way the per-stack Directus env
   contract targets the stack's **own** offset-port `directus` schema, **never** `content.anthropos.work`
   (`EnvContract.Validate()` hard-rejects it).

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
5. **Verify** — replay prints `replayed "<surface>" into <stack>: <T> table(s), <R> row(s) loaded [, reindexed …]`.
   Optionally gate fidelity (captured source vs replayed stack):
   ```bash
   SS=stack-demo/rosetta-extensions/stack-seeding
   SNAP=.agentspace/snapshots
   /tmp/datadna measure-snapshot --stack dev-N --dna "$SS/dna/data-dna.json" \
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
  enter git). The cloud/S3 store is a **deferred (unscheduled-backlog)** swap (DEF-M10-01).
- Base DSN (replay): `postgres://postgres@localhost:5432/postgres` with the port replaced by the stack offset.
- Exit codes: `0` ok · `1` firewall/capture/replay error (e.g. a tenant-data leak aborted capture) · `3` usage ·
  `4` (replay) target stack schema missing/empty → provision the STACK (not a cache problem) · `5` (replay) no
  cached snapshot at the stack's schema digest → run a capture (empty/outdated cache) or fix a diverged stack schema.

## Related skills

| Skill | Use when |
|-------|----------|
| `/stack-seed` | Seed the stack **after** set-dressing for a full-fidelity world |
| `/stack-list` | List live stacks to pick a target |
| `/dev-up` · `/demo-up` | Bring up the stack first (dev-up already set-dresses by default) |
