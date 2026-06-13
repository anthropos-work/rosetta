# Bring-up Re-run Safety — the idempotency contract

**The authoritative statement of what happens when you re-run a bring-up step.** A re-run of
**migrate**, **snapshot-replay**, **directus-provision**, or **seed** is now either **safe-and-idempotent**
(a 2nd run converges to the same state) or **fails loudly with a guard** — never silently doubles data and
never aborts half-way through a surface.

> **Scope.** This doc covers the **v1.3b "dress rehearsal" / M17** re-run guards across the
> `rosetta-extensions` stack tooling: the `migrate-demo.sh` first-run-race fixes, the `stacksnap replay`
> re-run guard, and the `stackseed` re-run guards (the idempotent COPY + the casbin grant + the fixed
> `--reset`) — plus the **v1.5 "prop room" / M22** per-stack **directus-provision** re-run guards
> (bootstrap-on-non-empty + container-name). It is the *what-happens-when-you-run-it-twice* companion to the
> mechanism specs it cross-links — [`snapshot-spec.md`](snapshot-spec.md) (replay),
> [`seeding-spec.md`](seeding-spec.md) (seed), [`directus-local.md`](directus-local.md) (the per-stack
> Directus lifecycle), and the demo lifecycle [`rosetta_demo.md`](rosetta_demo.md) (migrate). The **safety**
> of every destructive op here (the `TRUNCATE`s) is governed by [`safety.md`](safety.md) — they only ever
> touch a **per-stack-isolated offset** store.
>
> All the code cited lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the
> authoring copy at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag) — **no platform
> repo is modified.**

## For PMs — what "re-run safe" means

Building a demo or dev stack is several steps: create the database schema (**migrate**), copy in the real
public catalog (**snapshot-replay**), then fill it with a believable org of users + activity (**seed**).
Before M17, running any of those a *second* time could go wrong: the seed would error out or quietly
double a row, the catalog replay would crash part-way through, and the database setup could silently fail
on a slow machine because it raced the services starting up. M17 fixes all three: a re-run now either
**lands cleanly** (re-doing the work changes nothing) or **stops with a clear message** — it never leaves
the stack in a half-built, doubled, or quietly-broken state. The practical payoff: the lifecycle steps
become **safe to retry**, which is the foundation the later milestones build auto-chaining on.

## The per-component verdict table

| Component | Tool / path | Re-run verdict | Guard (M17) |
|---|---|---|---|
| **migrate** | `demo-stack/migrate-demo.sh` | **SAFE** (idempotent) | Schemas `IF NOT EXISTS`; atlas is declarative/revision-tracked; `init_policy.sql` applied only when `casbin_rules` is empty. **+ the first-run-race hardening** (below). |
| **snapshot-replay** | `stacksnap replay` (`stack-snapshot/replay/`) | **SAFE** (idempotent) | **Per-stack-isolated `TRUNCATE`-then-reload** before COPY — a 2nd replay REPLACES, never appends. |
| **directus-provision** | `dev-setdress.sh::provision_directus_step` / `boot_directus_step` (M22) | **SAFE** (converges) | `CREATE SCHEMA IF NOT EXISTS`; **bootstrap guarded on the `directus_collections` sentinel** (a half-bootstrap re-bootstraps); the structure/serve-row apply rides the replay's gap-gated auto-provision (no-op once provisioned); the serving container is the **compose service** (re-up reuses the name; the bootstrap `docker run` is `--rm`, no name clash); restart is idempotent. |
| **seed** | `stackseed` (`stack-seeding/`) | **SAFE** (idempotent) | **Idempotent COPY** (`ON CONFLICT (id) DO NOTHING`) for every deterministic-id surface + a **`WHERE NOT EXISTS`** casbin grant. `--reset` clears the **full** fleet. |

Before M17 the latter two were **NOT idempotent** (a 2nd replay doubled / duplicate-key-aborted; a 2nd
seed unique-violated or duplicated the casbin grant). The guards below close those.

## The two re-run models

There are two correct ways to re-run a bring-up step. M17 makes the tooling support **both**:

1. **Re-run-in-place (the M17 guards).** Run the same step again on the same stack — the guards make it
   converge. This is what `dev-setdress.sh`'s cache-first replay + light seed rely on, and what makes the
   lifecycle safe to retry.
2. **Teardown-then-redo (always available).** `rosetta-demo down N --purge` (or `dev-stack down N --purge`)
   wipes the stack's data, then re-up + re-migrate + re-replay + re-seed into a fresh/empty target. This is
   the cleanest reset and was the *only* safe re-run path before M17.

`stackseed --reset` is the **middle ground**: it truncates the seeded tables (per-stack only) so a
re-seed lands on an empty surface, without tearing the whole stack down.

## For engineers

### migrate — SAFE, plus the first-run-race hardening (ISSUE-7 + M17)

`migrate-demo.sh` was already idempotent in its writes (schemas `IF NOT EXISTS`, atlas declarative,
`init_policy.sql` count-guarded). The failure mode was a **first-run race**, not a re-run double: the
script runs under `set -euo pipefail` right after `up -d`, and a command-substitution whose pipeline
failed (a query against a not-yet-created relation) aborted the *whole script* before the policy loaded —
leaving an empty `casbin_rules` and a blanket-403 stack (ISSUE-7). M17 closes this class on three fronts:

- **Reactive (ISSUE-7, retained):** the casbin-count pre-check carries `|| echo 0`, so a failed query is
  treated as "empty" instead of aborting; `init_policy.sql` (`CREATE TABLE IF NOT EXISTS` + INSERT) then
  creates + seeds the table.
- **Proactive (M17):** a bounded, **non-fatal** wait — `wait_pg` (`pg_isready`, falling back to `SELECT 1`)
  + `wait_sentinel_running` (`docker inspect`) — runs *before* the first `docker exec psql`, so postgres is
  accepting connections and sentinel has a window to create `casbin_rules` itself. A timeout logs a WARN and
  proceeds (the reactive guard still recovers), so the wait only ever *removes* flakiness.
- **The schema-create step** is `|| log`-guarded too: `ON_ERROR_STOP=0` makes psql continue past a failing
  statement internally, but the psql *process* still exits non-zero — which under `set -e` would abort the
  script (e.g. an unavailable extension). The schemas are the must-haves; a missing extension is non-fatal.

The same `set -e` race class was swept across the other bring-up scripts: `up-injected.sh`'s `GH_PAT`
extraction now **fails loud** (a clear "GH_PAT not set" error, not a silent pipefail abort), and the
`DEV_PROJECT` extraction in `rosetta-demo` + `dev-stack` carries `|| true` so its documented
`${DEV_PROJECT:-anthropos}` fallback can actually run. (race-audit verdicts: #M17-D1; the bounded
non-fatal wait-for-ready: #M17-D2; the schema-create `|| log` latent-site fix: #M17-D9.)

> **Tested:** the static fence (`demo-stack/tests/test_tooling.py::TestMigrateRaceGuard` +
> `TestSetEraceGuards`) pins the guards; the **live** harness (`tests/test_migrate_race_live.py`) runs the
> real `migrate-demo.sh` against a throwaway pgvector Postgres container *in the race state* and asserts the
> script survives, seeds the policy, and re-runs idempotently. It skips cleanly when Docker is unavailable.

### snapshot-replay — TRUNCATE-then-reload (the re-run guard)

Replay loads the cached public surface into a stack via bulk COPY. A bare COPY is **not** idempotent: a 2nd
replay re-appends every row → PK tables duplicate-key-abort **mid-surface** (no wrapping transaction); tables
with no unique constraint **silently double**. The M17 guard makes replay **REPLACE, not append**:

`replay.Run` (`stack-snapshot/replay/replay.go`) now, after verifying every payload checksum and before
loading any rows, **clears every target table** via the new `Replayer.ClearForReplay`. The clear runs
**child-first** (reverse dependency order) so a plain per-table `TRUNCATE` never trips an FK from a
not-yet-cleared child — **no `CASCADE` needed**, keeping the blast radius to exactly the manifest's own
tables. Then the load runs **parent-first** (dependency order), FK-safe. On a first run the `TRUNCATE`s are
no-ops (empty tables); on a re-run they make the result identical, not doubled.

**Safe-by-default, not flagged (#M17-D3).** There is no `--idempotent`/`--force` flag — the operation is
harmless on a first run and is the intended behavior on a re-run.

**The destructive op is fenced (load-bearing — see [`safety.md`](safety.md); #M17-D4).** The clear SQL is
built by a single pure function, `truncateForReplaySQL` (`stack-snapshot/cmd/stacksnap/adapters.go`), pinned
by a **target-class test** to ALWAYS be a single-table `TRUNCATE TABLE "schema"."table" RESTART IDENTITY` —
never a `DROP`, `DELETE`, `CASCADE`, or cross-schema op, identifiers double-quote-escaped (injection-safe). And
the connection it runs on is built by `pg.DSNForOffset(baseDSN, n)` — the **per-stack offset** every replay
write uses — so the `TRUNCATE` can only ever land on the per-stack-isolated Postgres (for `N>0` a different
host port from prod's `:5432`; for `N=0` the dev stack's own isolated container). A wrong-target `TRUNCATE`
would be data loss; the shape-pin + the structural offset are the two independent fences against it.

### directus-provision — bootstrap-on-non-empty + container-name guards (M22)

The per-stack Directus provision (`dev-setdress.sh`, run on a `--local-content` bring-up) **converges** on a
re-run — the M17 re-run contract, applied to the executed bootstrap → apply-structure → replay → boot recipe:

1. **`CREATE SCHEMA IF NOT EXISTS directus`** — a no-op on a re-run.
2. **bootstrap-on-non-empty guard.** `node cli.js bootstrap` is skipped when the `directus` schema already
   holds the **`directus_collections` sentinel** — a table present only after a *complete* bootstrap. Probing
   the sentinel rather than a blanket `directus_*` count is the load-bearing nuance: a **half-bootstrap** (a
   crash that left some system tables but no registry) is detected as incomplete and **re-bootstraps to
   converge**, instead of skipping onto a broken schema. `bootstrap` is itself idempotent (it runs pending
   migrations only), so the guard is an optimisation + a clean log, not a correctness crutch.
3. **structure / serve-row apply** rides `stacksnap replay`'s own gap-gated `tryAutoProvision` (M21) — a no-op
   once the schema is provisioned (the `nUser==0` gate), so the apply leg is idempotent unchanged.
4. **container-name-conflict guard.** The serving container is the **compose service** (a re-up reuses the
   name — no clash), and the one-shot bootstrap runs as a `docker run --rm` (no `--name`, so nothing to
   collide). The post-replay `docker restart` is idempotent.

So a 2nd `--local-content` pass: bootstrap skipped, the M21 replay auto-provision a no-op, the rows
`TRUNCATE`-reloaded (the replay guard above), the restart idempotent — the result converges, never
half-applies. The provision is **non-fatal**: any step failing degrades the stack to the prod-read path with
an honest `⚠` status line (it never blocks a re-run). See
[`directus-local.md`](directus-local.md) § "Container lifecycle (M22)".

### seed — idempotent COPY + the casbin grant + the fixed `--reset`

The seeders generate **deterministic ids**, so a 2nd run produces the *same* primary keys. Three guards
make a re-seed safe:

1. **The idempotent COPY (every deterministic-id surface) (#M17-D5).** A new `Conn.CopyRowsIdempotent(…, conflictCol)`
   (`stack-seeding/pg/pg.go`) keeps the bulk-COPY speed but makes a re-run a no-op for existing rows: it
   COPYs into a session-local `TEMP TABLE (LIKE … INCLUDING DEFAULTS) ON COMMIT DROP`, then
   `INSERT … SELECT … ON CONFLICT (<id>) DO NOTHING` into the real table, all in **one transaction**. (COPY
   itself has no `ON CONFLICT` form; per-row `INSERT … ON CONFLICT` would kill the bulk path the seeder
   exists to preserve — the temp-then-merge gets both.) All seven seeders use it (every seeded table keys on
   `id`).
2. **The casbin g2 grant (#M17-D6).** It uses `INSERT … SELECT … WHERE NOT EXISTS (the same tuple)`, **not**
   `ON CONFLICT` — the casbin tables have no unique constraint on the policy tuple, so `ON CONFLICT` has no
   target. `WHERE NOT EXISTS` is idempotent regardless of constraints; a 2nd seed inserts 0 grants (before
   M17 it genuinely duplicated the grant every run).
3. **`stackseed --reset` (#M17-D7).** The truncate list was stale — `{memberships, users, organizations}` only, which
   skipped every M7c activity/session/assignment surface, so even a reset-then-seed collided on the leftover
   rows. M17 extends it to the **full deterministic-id fleet, child-first FK-safe**
   (`activity_events → jobsim sessions → skill_path_sessions → assignments → memberships → users →
   organizations`). The casbin g2 grant is **not** TRUNCATEd — it shares the `sentinel` schema with
   `init_policy.sql`'s global policy (the ~47 `p`-rows), so a TRUNCATE would wipe platform bootstrap; it gets
   a **targeted** `DELETE … WHERE p_type='g2'` (`resetCasbin`) instead, leaving the policy intact.

> **Tested:** the merge SQL builders (the `ON CONFLICT` shape, the no-constraints temp, reserved-column
> quoting), a re-seed-inserts-nothing-new check, the casbin `WHERE NOT EXISTS` guard, and the
> `--reset`-covers-the-full-FK-ordered-fleet check — all mutation-pinned (reverting a guard fails the test,
> showing the exact doubling). The replay re-run is proven end-to-end through the real replay package
> (`stack-snapshot/reference` — replay twice → same counts, not doubled).

## See also
- [`safety.md`](safety.md) — **why the `TRUNCATE`s are safe**: every destructive op targets only a
  per-stack-isolated offset store; the 3-layer write-isolation guard + the per-surface firewall still hold.
- [`snapshot-spec.md`](snapshot-spec.md) — the replay mechanism (capture → cache → replay) the re-run guard
  sits on.
- [`seeding-spec.md`](seeding-spec.md) — the seeder fleet, the `stack.seed.yaml` blueprint, the `--reset`
  model, and the n=0-dev guards.
- [`rosetta_demo.md`](rosetta_demo.md) — the demo lifecycle: bring-up, `migrate-demo.sh`, teardown.
