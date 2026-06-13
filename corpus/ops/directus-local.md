# The Per-Stack Local Directus

**What this is.** Every Anthropos stack today reads its public content (the simulation catalog, skill-path
library, etc.) **live from production** — `cms` and `jobsimulation` point `DIRECTUS_BASE_ADDR` at
`content.anthropos.work`. The **"prop room"** release (v1.5) makes each stack stand up its **own local
Directus** that serves the **captured public library**, so a stack's content is self-contained with no live
prod dependency at runtime. Real images stay real: only the *data plane* (catalog rows) goes local; the
*asset plane* (image bytes) stays on prod's anonymous public links.

This doc is the spec for that local Directus — the empirically-pinned bootstrap facts, the structure-capture
model that makes the catalog serveable, and the version-skew rule. It is the companion to
[`snapshot-spec.md`](./snapshot-spec.md) (which captures the rows) and
[`snapshot-cold-start.md`](./snapshot-cold-start.md) (the `--dsn` source the structure capture rides).

> **Status (2026-06-13, after M22 "Executed provisioning + lifecycle"):** the **structure-capture half**
> (`stack-snapshot`, tag `prop-room-m21`) and the **lifecycle half** (`dev-stack` + `stack-injection` +
> `stack-verify`, tag `prop-room-m22`) are both **built**. M22 turned the print-only recipe into an
> **executed** bring-up step: a per-stack Directus boots as a **compose service** (offset port, torn down with
> the stack), provisioned idempotently, with verify probes — **demo default-on / dev opt-in (`--local-content`)**.
> The firewall is now a **load-bearing executed gate** (a prod-resolving env hard-aborts before any write).
> See § "Container lifecycle (M22)" below. The remaining **M23** work is the *cutover*: re-pointing
> `DIRECTUS_BASE_ADDR` at the local instance + referential closure (this doc grows its "Cutover" section then).

---

## The store fork at a glance

Booting a per-stack Directus needs three layers, in order:

1. **The `directus_*` system schema** — Directus's own 27 system tables. The snapshot **never** carries
   these; Directus owns their version-specific DDL. They come from **`node cli.js bootstrap`**.
2. **The content-model structure** — the user-collection table DDL (`simulations`, `skill_paths`, …) **with
   primary keys**, plus the registry rows in `directus_collections` (and the serve permissions) that make a
   booted Directus actually *serve* those collections. This is what M21 captures (was the "collection-schema
   gap").
3. **The content rows** — the captured public catalog, bulk-`COPY`ed into the user-collection tables by the
   existing `stacksnap replay --surface directus` path.

The recipe is **bootstrap → apply-structure → replay → boot**. M21 closed the apply-structure step (step 2),
which previously dead-ended at the print-only placeholder in `provision.go`'s `ProvisionPlan` (the recipe was
*described* but never *applied*) and made the replay fail loud with `stacksnap` exit 4/5. M22 then made
**bootstrap** and **boot** executed too (see § "Container lifecycle (M22)").

---

## Bootstrap empirics (Directus 11.6.1 — pinned)

The per-stack image is pinned to **`directus/directus:11.6.1`**. These facts only surface live (the reason M21
was built as an iterative milestone — Directus's bootstrap and permission behavior breaks empirically, not on
paper):

- **`bootstrap` creates the 27 `directus_*` system tables — and nothing else.** The older recipe's claim that
  bootstrap also creates the user-collection structure is **false**; that is exactly the structure-capture
  step's job.
- **`DB_SEARCH_PATH=directus` is required.** The `directus` schema must be `CREATE SCHEMA`-d first and
  bootstrap pointed at it, else everything lands in `public` and the replay probe never finds the schema.
- **The image entrypoint runs `node <arg>`**, so a bare `bootstrap` argument dies `MODULE_NOT_FOUND` — use
  `node cli.js bootstrap`.
- **The admin email must be a format-valid, real-TLD address.** 11.6.1's email validator **rejects the
  `.local` TLD outright** — `admin@<stack>.local` dies `FAILED_VALIDATION` and crashes bootstrap (this
  superseded an earlier hyphen-vs-underscore framing). The minted address is **`admin@<stack>.example.com`**
  (RFC-2606 reserved — never a real address, always format-valid), keeping the stack name as the subdomain for
  provenance. (#M21-D1)
- **Split host/container DSNs.** The psql leg and the docker leg cannot share one DSN — one value cannot reach
  the same Postgres from both sides.

---

## The structure-capture model (what M21 built)

The structure artifact is captured from a **sanctioned, read-only, public-only `--dsn`** (the same source
discipline as the row capture — read-only, bounded, operator-confirmed, behind the firewall). It carries:

- **The user-collection table DDL** — exact `pg_catalog` column types (uuid / json / text / varchar(N) /
  integer / boolean / timestamp), for **all 26 user collections** (not just the 9 the row surface captures).
  Capturing all 26 is what makes the per-stack schema digest **match prod's** so the row cache hits — see the
  version-skew rule below. (#M21-D7)
- **PRIMARY KEY constraints.** Load-bearing: **Directus refuses to serve a collection with no detectable
  primary key** (`"doesn't have a primary key column and will be ignored"` → 403, even for an admin). The
  schema digest is over column *types*, not constraints, so a column-only DDL converges the digest and the row
  COPY works — but serving silently fails. The artifact **must** carry constraints, at minimum the PKs (`id`
  for 25 collections, `code` for `languages`). (#M21-D9)
- **Sequences**, keyed by the **default-reference** dependency (`DEFAULT nextval('seq')`), not ownership — so
  the capture is robust to a source whose sequences aren't `OWNED BY` their table (e.g. a hand-built fixture or
  a dump where ownership was dropped). The only non-uuid serial PK is `sequences_roles.id`. (#M21-D11)
- **The serve rows** — a `directus_collections` registration row per served collection + a public-read
  `directus_permissions` row on Directus's **hardcoded public policy** `abf8a154-…` (bootstrap creates the
  policy + its `(role=NULL, user=NULL)` access link). `directus_fields` rows are **not** required — Directus
  introspects the DB columns once a collection is registered and has a PK. (#M21-D9, #M21-D13)

**Capture scopes to the digest's view.** Every structure-capture catalog query intersects `pg_catalog` (for
exact types/PKs) with `information_schema.columns` (which is **privilege-filtered** — it shows only relations
the connecting role can see). This keeps the captured table set identical to the set the staleness digest
counts; a naive `pg_class` capture would find tables the read role can't see in `information_schema`, leaving
the applied schema "ahead" of the digest so it never converges. (#M21-D10)

**Apply mechanism.** The structure artifact is a SQL script (DDL + PKs + sequences + the additive serve-row
INSERTs) applied via `pg.ExecScript` **before** the row replay. The serve-row INSERTs are *additive* (not the
TRUNCATE-COPY row replay, which would wipe bootstrap's own system rows): `directus_collections` uses
`ON CONFLICT (collection) DO NOTHING`; `directus_permissions` omits the serial `id` so it auto-generates
(a captured prod id would collide with bootstrap's own system-permission serials). Both are rendered
dynamically from the source (one server-side query each, column set discovered at runtime) so the capture is
version-robust and every value round-trips. (#M21-D13)

**Auto-provision is gated on a true bootstrap gap.** `stacksnap replay` applies the structure on a cache miss
**only** when the target has **zero user collections** (a fresh bootstrap gap). A diverged target (already has
user collections, digest ≠ captured) is a no-op that falls through to the existing clean **exit 5**
("bring the stack to the captured shape first") — never a raw collision. The general rule: any
auto-provision-on-cache-miss must gate its mutation on the precondition it assumes, or a skewed input degrades
a clean error into a raw failure. (#M21-D12)

### Redefined `stacksnap` exit semantics

The structure artifact **is** what provisions the schema, so exit 4 no longer means "give up":

| Target directus schema state | Replay outcome |
|---|---|
| **empty** (never bootstrapped) | exit **4** — provision the stack first (bootstrap before replay) |
| **bootstrapped gap** (27 system tables, 0 user collections) **+ a captured structure** | **auto-provision then exit 0** |
| **bootstrapped gap, no captured structure** | exit **5** (cache miss; nothing to provision with) |
| **diverged** (has user collections, digest ≠ captured) | exit **5** (no-op; bring the stack to the captured shape) |

---

## The version-skew rule

The row cache is keyed by the **full `directus`-schema digest** (`pg.SchemaVersionSQL` over every column of
every table — system tables *at prod's Directus version* + all prod content collections + their exact types).
A per-stack bootstrap converges that digest **only if its entire schema matches prod's**. The release resolves
this with **option A** (capture all 26 collections + pin the version), not a per-surface re-key, to keep the
shared staleness key untouched (zero blast-radius on taxonomy, which shares it):

1. **Pin the Directus image** to the version whose bootstrapped system tables match prod's. Verified for M21:
   a fresh `directus/directus:11.6.1` bootstrap produces the system-table digest `b4cb55bc…`, which **equals
   prod's system-only digest** — no version skew at the pinned version. (#M21-D8)
2. **Capture all 26 user collections'** structure (the row cache still only fills the 9 public-content
   collections; the other 17 tables exist empty — fine, since the digest is over column structure, not rows).
3. **Record the source Directus version in the manifest** and **warn on mismatch** so a future prod version
   bump that changes the system tables is caught (it would shift `b4cb55bc…` and the digest would stop
   converging until the local image is re-pinned).

---

## Safety: structure capture is still a prod read

The structure capture rides the **same** read-side discipline as the row capture: read-only, bounded,
operator-confirmed, **public-only**, behind the firewall. M21 *extended* the firewall to admit structural
metadata without *loosening* it:

- The serve rows live in `directus_*` **system** tables, outside `AssertPublicOnly`/`AssertPlan` (which govern
  user-collection *row* captures). A new gate **`AssertStructuralMetadata`** admits a system table as
  "structure, not tenant data" **only if** it carries **none** of the tenant-scope columns
  (`organization_id`, `tenant_id`, `private`, `user`, `owner`, `user_created`, `user_updated`). Any
  tenant-scope column → reject. (#M21-D13)
- `directus_collections` and `directus_permissions` carry zero tenant-scope columns → admissible.
  `directus_access` is **excluded** (it has a `user` uuid column); `directus_policies` is **not captured**
  (both are bootstrap-provided — the hardcoded public policy + its anonymous access link exist on any fresh
  bootstrap).
- The carve-out runs **assert-then-read**: admissibility is checked on the introspected column set **before**
  any row is materialized, so an unexpected tenant column aborts the capture before reading.

The dropped pg_dump-file-reader (a prior release's rejected path) stays dropped — `TestDroppedDumpFlagStaysGone`
pins it gone.

See [`safety.md`](./safety.md) for the full read-side / write-side contract.

---

## Container lifecycle (M22)

M22 turned the print-only recipe into an **executed**, idempotent, prod-safe bring-up step. The serving
Directus is a **compose service** — not a bespoke `docker run` — so the stack's existing lifecycle plumbing
(`demo-down`/`dev-down` teardown, the port registry, `stack-verify`'s naming convention) covers it with no
new lifecycle code. That is the v1.5 **maintainability constraint**: the only things the recipe executes
itself are the two steps compose can't express — the one-shot **bootstrap** and the post-replay **restart**.

### The compose service

`gen_injected_override.py::directus_lines` (demo) and `gen_override.py` (dev, `--with-directus`) append a
`directus` service block to the stack's override:

- **image** `directus/directus:11.6.1` (the pinned version — see § the version-skew rule), `pull_policy: missing`
  (a cached image is reused; a fresh box pulls once).
- **port** `8055 + N·10000` published to the host (`!override` so it replaces, not merges) — the same offset
  arithmetic as every other service; `<project>-directus-1` is the container name.
- **network** `app-network` — the same in-network seam the fake BAPI alias uses, so `cms`/`studio-desk` will
  reach it by name in M23.
- **backing store** the stack's **own** `postgresql` compose service (`DB_CONNECTION_STRING=…@postgresql:5432`,
  `DB_SEARCH_PATH=directus`) — the per-stack-isolated offset Postgres, **never prod**. `SECRET` is a throwaway
  per-stack value. `mem_limit: 1g` keeps two stacks co-resident on a 16 GB box.

It is emitted **only when local content is on** — demo default; dev opt-in via `--local-content`;
`DEMO_NO_LOCAL_CONTENT=1` clears it on demo. A prod-read stack has no `directus` service at all (so teardown,
the registry, and verify all correctly see nothing).

### The executed steps (`dev-setdress.sh`)

`provision_directus_step` + `boot_directus_step` run inside the shared set-dress engine when `--local-content`
is set (demo default-on; dev opt-in; `N=0` additionally behind `--force`):

1. **`CREATE SCHEMA IF NOT EXISTS directus`** on the offset Postgres (idempotent).
2. **bootstrap** the `directus_*` system tables (`node cli.js bootstrap`), **guarded** on the
   `directus_collections` **sentinel** — see § "Idempotent re-provision" below.
3. **apply-structure + replay** — `stacksnap replay --surface directus` auto-provisions the captured
   structure + serve rows onto the fresh bootstrap gap (the M21 path) then bulk-`COPY`s the rows.
4. **boot/restart** the compose service so Directus re-introspects the now-registered collections (it caches
   the registry at boot; a container that started before the serve rows landed won't serve them until
   restarted). Only `docker restart` — never a bespoke `docker run`.

**The firewall is a load-bearing executed gate.** Before any provision write, `provision-plan --check-env`
validates the per-stack env contract (the offset Directus addr + the offset Postgres DSN); a prod-resolving
env **hard-aborts** before bootstrap/replay — the M17-for-TRUNCATE discipline, now for the executed provision.

**Non-fatal degrade.** Any step failing degrades the stack to the **prod-read path** (the directus replay
surface skips, the seed floor still runs) with an honest `⚠ … content:prod-read` status line — a Directus
hiccup never blocks a good stack.

### Idempotent re-provision

A second `--local-content` pass **converges** (the M17 re-run contract):

- **bootstrap-on-non-empty guard** — bootstrap is skipped if the `directus` schema already holds the
  **`directus_collections` sentinel** (a *complete*-bootstrap marker). Probing the sentinel — not a blanket
  `directus_*` count — makes the guard robust to a **half-bootstrap**: a crash that left some system tables but
  no registry **re-bootstraps to converge** instead of skipping onto an incomplete schema. `node cli.js
  bootstrap` is itself idempotent (pending migrations only), so the guard is an optimisation + a clean log, not
  a correctness crutch.
- **container-name-conflict guard** — the serving container is the compose service (re-up reuses the name, no
  clash); the bootstrap `docker run` is `--rm` ephemeral (no name → no conflict). The M21 replay
  auto-provision is a no-op once provisioned (`nUser==0` gate), and the restart is idempotent.

### Verify probes

`stack-verify` gained a `directus` row + cheap-wins, scoped IN only on a `--local-content` stack and gated on
the directus **container actually existing** (a prod-read stack never false-warns):

- a **SERVICES row** — liveness via `/server/health` (HTTP 200), offset/project-rewritten like every service.
- the per-stack **`directus` schema** is added to the readiness expected-schemas list **when the container is
  present** (a prod-read stack has neither, so it isn't expected).
- a **`directus-collections` readiness probe** + an autoverify **"registered collections > 0" cheap-win** — the
  silent-failure analog of the casbin assert: a Directus can be UP (health 200) but serve **nothing** if the
  content-model never registered.
- a **no-prod-read env assert** — the runtime mirror of the EnvContract gate: warns (non-fatal) if the local
  Directus's DB connection string resolves to a prod host.

### Headroom

The **12 GB-VM preflight** notes the per-stack Directus runtime container (~1 GiB, `mem_limit 1g`) in its
budget when local content is on — a non-fatal note, independent of the UI build spike (Directus boots even on
a `--no-ui` demo).

---

## What's still future work

- **Content cutover + referential closure (M23)** — re-point `DIRECTUS_BASE_ADDR` at the local instance (asset
  plane stays on prod), and guarantee the served catalog is referentially closed (no content row references a
  taxonomy node-id the captured subset lacks — the empty Assign-AI-Simulation-picker class). M23 also wires the
  `directus_files` ref capture (the asset-ref plumbing) and closes the 20 dangling relations.
- **Blob bytes** stay backlog (DEF-M10-01) — the asset plane uses prod's anonymous public links so images stay
  real without mirroring the bytes.
