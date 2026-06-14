# Rosetta Tooling — Safety & Security Contract

**The authoritative statement of how the `rosetta-extensions` stack tooling stays safe.** Two inviolable
guarantees, proven in code and tested:

1. **It never reads private/customer data.** Anything that *leaves* production (a snapshot capture) is **public
   reference data only** — enforced by a tenant-data firewall that hard-fails on a single customer-scoped row.
2. **It never touches production data or services.** Anything a non-prod stack *writes* is confined to that
   stack's own isolated stores — enforced by a 3-layer write-isolation guard that makes a shared/prod write
   **structurally impossible** on a non-prod target, and an audit log that *proves* nothing leaked.

> **Scope.** This doc is the consolidated safety contract over the v1.2 snapshot mechanism + the v1.1 seeding
> framework, as they stand at v1.3 "stack party" (dev stacks are now first-class peers of demo stacks; both
> kinds run the same tooling). It is the *why-it-is-safe* companion to the three operational specs it
> cross-links: the read foundation [`db-access.md`](db-access.md), the read-side capture mechanism
> [`snapshot-spec.md`](snapshot-spec.md), and the write-side seeding boundary [`seeding-spec.md`](seeding-spec.md).
> The platform's own tenant-isolation posture (DB / authz / identity) is in
> [`../architecture/security_compliance.md`](../architecture/security_compliance.md); this doc is about the
> **tooling's** posture, a layer above it.
>
> All the code cited here lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the
> authoring copy at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag) — **no platform repo
> is modified**, and snapshot payloads **never enter git**.

## For PMs — the two promises in one paragraph

The tooling that builds demo and dev environments has to do two dangerous-sounding things: it **reads** the real
production database (to copy the public catalog), and it **writes** a lot of data into local stacks (to populate
them). Both are fenced. On the read side, only **public reference data** — the same skills/roles/templates every
customer sees — can ever leave production; a firewall checks every row twice and aborts the whole capture if it
finds even one customer-owned row, so **no customer's private data can be copied**. On the write side, a small
set of stores are *shared* with the live product (the content system, one storage bucket, the login system);
the tooling **blocks every write to those from a non-production stack**, repairs the environment before it
starts, and produces an **audit log that proves** nothing leaked. Neither promise depends on the operator
remembering a flag — both are structural.

---

## Part 1 — The read side: never reads private/customer data

A snapshot **capture** is the only operation that reads production. Everything that protects that read lives in
`stack-snapshot/`.

### 1.1 The tenant-data firewall (`AssertPublicOnly`)

`AssertPublicOnly` is the **read-side analog of seeding's `AssertClean`** (Part 2.3). It is a *concept* — "the
captured set is public-only" — enforced by **two real Go gates** in `stack-snapshot/firewall/firewall.go`, run
in sequence (grep these names, not `AssertPublicOnly`, which is the umbrella name only):

1. **`AssertPlan(policies, predicate)` — PLAN time, before a single byte flows.** Every table in the capture
   plan must declare an *admissible* policy. A scope-bearing table (one that carries the predicate's scope
   column) must be filtered to the public subset; a column-less table must be either pure-reference or scoped
   through a public parent. A bad plan refuses the capture **before any read runs**.
2. **`AssertCaptured(results)` — POST-capture, after the rows are in hand.** A hard re-check that the captured
   set holds **zero** tenant-scoped rows. A single leaked row aborts the capture — **nothing is written to the
   store.**

This is **defense in depth**: the plan gate catches a mis-declared table before reading; the post-capture gate
catches anything that slipped through *after* reading, before persisting. Either failing aborts the run.

The firewall package is **pure** — the policy logic is unit-tested without a database; the caller supplies the
introspected facts (does the table carry a scope column?) and the post-capture tenant-row count.

### 1.2 The public predicate is per-surface (the M10 generalization)

"Public" is **not one fixed column** — it differs by surface, so the firewall takes a **`PublicPredicate`** per
surface (the scope column(s) that decide public-vs-customer, plus the SQL `WHERE` that selects the public
subset):

| Surface | Public predicate (`firewall.PublicPredicate.PublicFilter`) | Prod-verified split (2026-06-06) |
|---|---|---|
| **taxonomy** (skiller) | `organization_id IS NULL` (`firewall.DefaultPredicate` / `firewall.PublicFilter`) | `skiller.skills` 42,763 public / 794 customer |
| **Directus content** | `private = false AND tenant_id IS NULL AND status = 'published'` | `directus.simulations` 304 strict-public-published / 2,597 total |

The org-only predicate is the **default** every surface gets unless it declares its own, so the taxonomy surface
and all the M9a/M9b behavior are byte-for-byte unchanged. A **column-less** table (embeddings, translations,
`sim_tasks`) carries no scope column — it is public iff its **parent** is, judged under the surface's predicate
(`firewall.ParentScopeFilter`); multi-level chains (`task_checks → sim_tasks → simulations`) chase to the
scope-bearing root in one subquery. `Validate()` rejects a malformed predicate (no scope column, empty filter,
or a filter that doesn't reference its own scope column) before it can ever gate a capture.

> **What is excluded entirely.** Surfaces that are **100% customer** are not captured at all — the
> app-Postgres `cms.studio_*` tables (`studio_documents`: 0 public / 3,060 customer) are dropped from the plan,
> not filtered. There is no "public subset" to keep.

### 1.3 The public-only data-DNA gene

The snapshot-fidelity data-DNA (the schema-conformance gate, see [`snapshot-spec.md`](snapshot-spec.md)) carries
a **public-only gene**: the captured surface is measured against the platform's current schema *and* asserted
public-only as part of conformance, so a schema drift that introduced a new tenant-bearing column would surface
as a DNA failure rather than silently widening the capture.

### 1.4 Capture is read-only and production-safe (the capture-source policy)

A capture must never block the hot primary. The source is **pluggable**, tried in a fixed precedence
(`source.DefaultPrecedence` in `stack-snapshot/source/source.go`):

| # | Source kind | When | Prod impact |
|---|---|---|---|
| 0 | **cache-hit** | the cached manifest's schema version matches | **zero read** (handled upstream) |
| 1 | **`dump-ingest`** *(default)* | a staging prod `pg_dump` exists → restore it into a throwaway Postgres, point `--dsn` at the restore | **zero new prod load** (the restore *is* the ingest) |
| 2 | **`primary-read`** *(fallback)* | only a read DSN is available | low — MVCC, off-peak, bounded |
| 3 | **`restore-from-snapshot`** *(upgrade)* | once eu-west-1 AWS access is wired | zero (throwaway instance) |
| 4 | **`read-replica`** *(upgrade)* | once a terraform replica exists | zero (cleanest steady state) |

**Both live sources read over `--dsn`.** There is **no offline `pg_dump`-FILE reader** — a dump is "ingested" by
**restoring it into Postgres and pointing `--dsn` at the restore** (the direct file-reader was considered and
dropped, M9b-D9: it adds no capability — the snapshot is byte-identical — and no reliable speed gain). The two
upgrade kinds (3, 4) trail the precedence list and activate automatically once `Kind.Available()` flips true (they
need eu-west-1 AWS/infra access not wired today; there is no read replica on prod as of 2026-06-06).

**Why a safe primary read is tolerable.** PostgreSQL MVCC means a read-only `SELECT`/`COPY` **never takes a lock
that conflicts with writers** — the only cost is I/O + buffer-cache pressure. So an off-peak, throttled,
public-only, catalog-sized-first read is a sanctioned fallback, not a scary last resort. Every capture session
is **bounded** (`source.BoundedSession.SetupSQL`), which makes the session **structurally unable to write**:

```sql
SET TRANSACTION READ ONLY;                          -- the read-side analog of the write guard: cannot write
SET statement_timeout = 1800000;                    -- 30 min: a runaway COPY aborts rather than holding a backend
SET idle_in_transaction_session_timeout = 60000;    -- a stuck client never pins an old snapshot / bloats the primary
SET work_mem = '64MB';                              -- modest; no buffer-cache pressure
```

This is the **read half** the write-isolation guard (Part 2) lacks — that guard classifies and gates *writes*
only.

---

## Part 2 — The write side: never touches prod data or services

Seeding and replay **write** a lot — into a stack. Everything that keeps those writes off production lives in
`stack-seeding/isolation/`.

### 2.1 The store registry — three isolation classes

Every store a stack might touch is classified (`isolation.IsolationClass`). Only the shared/external classes gate
the guard; per-stack stores are listed for documentation + dry-run preview:

| Store(s) | Class | Why / guard action |
|---|---|---|
| **Directus** (`content.anthropos.work`) | `SharedPollutionRisk` | one global instance, visible on prod → **direct writes blocked**; the shared instance is **never written**. (Reads: since **M23** a `--local-content` stack (demo default; dev opt-in) reads its **own per-stack Directus** — M22 boots it, M23 re-points `cms`'s `DIRECTUS_BASE_ADDR` at it (`http://directus:8055`, in-network) — so the served catalog is **local, not a live-prod read**. The prod **data plane** is read only at **capture** time (read-only, public-only, operator-confirmed). The prod **asset plane** stays in use: `DIRECTUS_PUBLIC_BASE_ADDR` keeps pointing here so browser images load real `<...>/assets/<uuid>` URLs (a public, anonymous, read-only GET of a public asset — within the read-side boundary). A **non-`--local-content`** stack (no per-stack Directus) still reads the public content **live** from this instance; a demo does so **anonymously**, the prod token stripped — the documented prod-read fallback. studio-desk's prod-**write** path is disarmed either way (token strip on a prod-read stack; a locally-minted token on a local-content stack).) |
| **S3 public bucket** | `SharedPollutionRisk` | hardcoded to the prod bucket in compose → `STORAGE_S3_PUBLIC_BUCKET` forced to `""` (local fallback) |
| **Live Clerk** | `SharedPollutionRisk` | shared dev app → routed to **Clerkenstein**; a real-Clerk base URL is a hard preflight error |
| **Customer.io / Brevo / AI provider APIs** | `SharedPollutionRisk` | external SaaS; blocked on non-prod (off by default) |
| **coresignal** | `External` | enrichment source — safe to read, **never write** on non-prod |
| **Postgres / Redis / S3-private / pgvector** | `PerStackIsolated` | inside the stack's own containers → **seed freely** (cannot pollute anything outside the stack) |

### 2.2 The 3-layer isolation guard

The guard (`stack-seeding/isolation/`) is three independent enforcement points:

1. **`Guard.CheckWrite(store, class, target)`** — refuses any `SharedPollutionRisk`/`External` write on a
   **non-prod** target. The **asymmetry** is the structural guarantee: the `AllowSharedOptIn` flag only relaxes a
   **prod** target — **a non-prod stack can never write a shared store, regardless of opt-in.** This makes
   non-prod pollution impossible by construction, not by configuration.
2. **`Guard.PreflightEnv(env, target)`** — *before* seeding begins, asserts and repairs the environment:
   - **forces `STORAGE_S3_PUBLIC_BUCKET = ""`** (always, every target) so no storage write can reach the shared
     public bucket;
   - on non-prod, **rejects a live-Clerk base URL** (any of `CLERK_API_URL`, `CLERK_FAPI_URL`, … pointing at
     `clerk.com` / `api.clerk.com` / `*.clerk.accounts.dev` / `*.clerk.services`) as a hard error — it must be a
     Clerkenstein/local host;
   - on non-prod, **rejects a live Directus write token** (`DIRECTUS_TOKEN` / `DIRECTUS_STATIC_TOKEN` /
     `DIRECTUS_ADMIN_TOKEN`) — content is snapshot-replayed into the per-stack Directus, never written to the
     shared one.

   > **Scope note + the compose-side closure (`rosetta-extensions @ dress-rehearsal-m20-fix16/fix17`).**
   > `PreflightEnv` guards the *seeding tool's* env — it never saw the **compose-inherited** token. **Before the
   > strip**, the platform's shared `env_file: .env` sprayed the prod `DIRECTUS_TOKEN` into the demo's containers
   > (the `env_file` reaches 11 of demo-1's services; fix16 had already cleared 2, so the **pre-fix17 audit**
   > found it in **9**) — including studio-desk, whose skill-path builder *could have written* prod Directus.
   > That hole is now closed at the source: the injected override (`gen_injected_override.py`, fix17) strips
   > `DIRECTUS_TOKEN=` on **every** emitted service + both frontends. The demo's live-prod public-content read
   > still works — **anonymously** (cms omits the `Authorization` header when the token is empty; prod Directus
   > serves the public predicate tokenless). **After the strip**, the audit shows **0 of 16** demo-1 containers
   > carry a token, and auto-verify is green (all verified live, 2026-06-11).
3. **`AuditLog.AssertClean(target)`** — *after* the run, the **proof** of zero pollution: it errors if **any**
   *allowed* write to a non-per-stack-isolated store actually landed on a non-prod target. On a prod target it is
   a no-op (prod is allowed to write shared stores). Every attempted write is `Record`ed during the run (the audit
   log is concurrency-safe — the seeder DAG runs in parallel), so `AssertClean` is an after-the-fact certificate,
   not a re-derivation.

> **`CheckWrite` is the gate; `AssertClean` is the proof.** The gate prevents the write; the audit log proves the
> gate held. A run that passes `AssertClean` on a non-prod target has a machine-checkable guarantee that **zero**
> shared/external writes landed.

### 2.3 Never-write shared Directus / prod-S3 (the two highest-risk vectors)

The two stores most likely to be hit by accident — because the platform's own compose file points at them — are
fenced twice over:

- **Shared Directus.** `PreflightEnv` strips any Directus write token on non-prod, `CheckWrite` blocks the
  store by class, and the injected override empties `DIRECTUS_TOKEN` on every demo container (fix17) — so the
  shared `content.anthropos.work` instance is **never written** from any non-prod stack. **Reads (M23 cutover):**
  a `--local-content` stack (demo default; dev opt-in) reads its **own per-stack Directus** — the override
  re-points `cms`'s `DIRECTUS_BASE_ADDR` at the in-network instance (`http://directus:8055`) and studio-desk's
  `DIRECTUS_BASE_URL` + a locally-minted token at it — so the **data plane is local, not a live-prod read**.
  Only the **asset plane** still touches prod: `DIRECTUS_PUBLIC_BASE_ADDR` stays `content.anthropos.work` so
  browser images load real assets (a public, anonymous, read-only GET — within the read-side boundary). A
  **non-`--local-content`** stack (no per-stack Directus) still reads the public content **live** from prod —
  the documented fallback — and on a demo that read is **anonymous** (no token → cms omits the `Authorization`
  header; prod serves only the public predicate). The earlier "every stack keeps
  `DIRECTUS_BASE_ADDR=content.anthropos.work`" state (the M10 collection-schema gap) is **retired** — the gap
  is closed (M21 structure capture + auto-provision) and the per-stack Directus is booted (M22) + cut over (M23).
- **Prod S3 public bucket.** `STORAGE_S3_PUBLIC_BUCKET` is hardcoded to the prod bucket in the platform compose;
  `PreflightEnv` **unconditionally** overrides it to `""`, so storage writes fall back to the per-stack local
  store. (Snapshot media is carried as **refs only** today — the byte payloads + a cloud snapshot store are
  **deferred (unscheduled backlog)**, see "Future" below.)

### 2.4 The capture-source policy is the write-side's read-half complement

The write-side guard classifies *writes*; it has no say over *reads*. The capture-source policy (Part 1.4) is the
deliberate read-half complement: `dump-ingest` (default, zero new prod load) → throttled `primary-read` (MVCC = no
write-blocking) → the not-yet-wired zero-impact upgrades, every session bounded `READ ONLY`. Together,
`AssertClean` (writes) and the bounded read-only capture (reads) close **both halves** of the prod-safety boundary.

### 2.5 The n=0-dev guards (doubled in v1.3 M13)

The **main dev stack** (`N=0`, the `anthropos` stack) is the developer's primary working environment. Two
operations could surprise a developer by mutating it, so each **independently refuses `N=0` unless `--force`**:

- **Auto-set-dressing** (`dev-stack/dev-setdress.sh`) — a `dev-stack up` build auto-replays the snapshot +
  light-seeds a *non-primary* `dev-N`, but **hard-refuses `N=0`** so the developer's own stack is never
  auto-modified. (v1.3b M20: this same engine now also set-dresses **demo** stacks via `--stack-type demo`; the
  `N=0` refusal is stack-type-agnostic, so it never weakens — see §2.7.)
- **Destructive `--reset`** (`stackseed`, `stack-seeding/cmd/stackseed/main.go`) — refuses to `--reset` the main
  dev stack (`N=0`) without `--force`.

> **Precise scope (so the doc doesn't over-claim).** Snapshot **replay** (`stacksnap`) has **no** `N=0` guard, and
> correctly so: replay writes only **public reference data** into the stack's **own** isolated Postgres/Directus —
> replaying the real catalog into the main dev stack is harmless (it's the developer's own stack, and the data is
> public). The `N=0` guard exists where a mutation would be *unsolicited* (auto-set-dress) or *destructive*
> (`--reset`), not on every tool. "Doubled in M13" = these two independent enforcement points, not a blanket
> refusal.

### 2.6 The audit-proven zero-pollution assertion

The headline write-side guarantee is not "we were careful" — it is **machine-checked**. The seeding test suite
runs a full seed against a non-prod target and asserts `AssertClean` passes (zero shared/external writes landed),
and asserts the guard *blocks* a deliberately-attempted shared write. The 3-layer guard's tests
(`isolation_test.go`) cover: `CheckWrite` blocks shared + external on non-prod; `PreflightEnv` forces the S3
override, rejects real-Clerk and live-Directus tokens, accepts Clerkenstein; `AssertClean` passes on a clean run
and catches a shared/external write. The guarantee is therefore **reproducible**, not asserted.

### 2.7 The demo auto-set-dress chain reuses the dev pass — the guarantees carry over (v1.3b M20)

`/demo-up` now auto-set-dresses by default — a cache-first snapshot **replay** + a light seed at the bring-up
tail — exactly as `/dev-up` has since M13. The load-bearing safety fact is that it is **not a second
implementation**: `demo-stack/up-injected.sh` chains the **same** `dev-stack/dev-setdress.sh` engine via
`--stack-type demo`, so every read- and write-side guarantee in this doc applies to the demo chain **by
construction, byte-for-byte** — there is no demo-specific set-dress code path that could drift from the dev one.
(#M20-D1)

- **Replay-only, never capture.** The bring-up chain does **cache-first REPLAY** (a per-stack WRITE of public
  reference data into the stack's own isolated offset-port Postgres + per-stack Directus). It **never runs
  `stacksnap capture`** — capture is a privileged, separate, operator-confirmed prod READ
  ([cold-start runbook](snapshot-cold-start.md)). A grep of `up-injected.sh` for `stacksnap capture` is empty,
  pinned by a test.
- **The per-stack Directus firewall still gates first.** The chain firewall-checks the per-stack Directus env
  (the M10 `EnvContract`) *before* any replay; a per-stack env that resolves to the shared prod Directus hard-aborts
  the pass before a single row is written. The shared `content.anthropos.work` / prod S3 are never written
  (§2.1–2.3 hold unchanged).
- **The seeder's isolation guard + `AssertClean` still prove it.** The seed step runs the same `stack-seeding`
  fleet behind the same 3-layer guard, so a demo set-dress produces the same machine-checked zero-pollution
  certificate (§2.6) as a dev one.
- **The n=0 guard holds across types.** `dev-setdress.sh`'s `N=0`-without-`--force` refusal fires regardless of
  `--stack-type` (§2.5). Demos start at `N ≥ 1` in practice, so it never blocks a real demo — it remains a dev
  safety net, not weakened by demo mode.
- **Atomicity (both-or-neither).** The chain always **seeds after the (cache-first, possibly-skipped) replay**, so
  a stack is never catalog-replayed-but-unseeded (which would 403 on every authorized route). A replay cache-miss
  degrades to a structural-only world that still logs in (the seed is the floor); the M17 re-run guards make a
  partial set-dress repairable by re-running (idempotent TRUNCATE-then-reload replay + idempotent seed COPY).
  (#M20-D3)

### 2.8 Bring-up scripts must survive bash 3.2 under `set -u` (the non-fatal-means-non-crashing invariant)

The bring-up wrappers (`preflight.sh`, `up-injected.sh`, `dev-stack`, `dev-setdress.sh`, …) shebang
`#!/usr/bin/env bash`, which on a macOS dev box resolves to the **system bash 3.2** (`/bin/bash`), not a
Homebrew bash 5. The non-fatal-verification contract (warn standard / fail critical / skip otherwise — see
[`verification.md`](verification.md)) is only honored if the wrapper actually **runs to its verdict** — a script
that *crashes mid-run* has silently turned a non-fatal check into a fatal one. The bash-3.2-specific trap that
breaks this: under `set -u`, expanding an **empty array** as `"${arr[@]}"` raises an "unbound variable" error and
aborts the script (bash 5 tolerates it). A flag-array conditionally populated (`flag=(); [ cond ] && flag=(--x)`)
and then expanded bare is the canonical offender — it crashes on the *un-set* branch only, so it passes a
bash-5 author's local test and fails on a colleague's stock-macOS box.

**Rule for every on-every-bring-up script:** expand a possibly-empty array with the conditional-expansion guard
`${arr[@]+"${arr[@]}"}` (empty → nothing, populated → the elements), never bare `"${arr[@]}"`, when running under
`set -u`. The M28 `preflight.sh` regression (`PreflightBehavior.test_non_demo_path_survives_set_u_on_bash32`)
invokes the wrapper through `/bin/bash` 3.x specifically and asserts no "unbound variable" abort, pinning the
fix; `shellcheck` does **not** catch this (it's a runtime-only, version-specific behavior). (#M28-harden)

---

## How this relates to the platform's own isolation

The platform itself has a 3-layer **tenant-isolation** posture — DB (`organization_id` on every table, Ent
auto-filtering), Authorization (Sentinel/Casbin), Identity (Clerk org-scoped JWTs) — documented in
[`../architecture/security_compliance.md`](../architecture/security_compliance.md#multi-tenant-data-isolation).
That protects **customers from each other inside the running product**. *This* doc's two guarantees protect
**production from the tooling**: the read-side firewall ensures the tooling only ever copies the *public* slice
of that multi-tenant database, and the write-side guard ensures the tooling never writes back into it. They are
complementary layers, not the same mechanism.

## Future (deferred / unscheduled — clearly not the current posture)

The **current** posture is: snapshot media carried as **refs only** (no blob bytes), snapshots stored in the
**local** `.agentspace/snapshots/` workspace cache. The **deferred** (unscheduled-backlog) work — **S3 media
blob bytes** + a **cloud snapshot store** (DEF-M10-01) — would change *what is transported and where it is
cached*, not the safety contract — the firewall, the predicates, and the write guard apply identically. This
section is a forward pointer only; everything else in this doc describes what ships today. (There is no version
currently staged for it — see the roadmap.)

## See also
- [`db-access.md`](db-access.md) — the production read foundation + the public-vs-customer boundary (read-side).
- [`snapshot-spec.md`](snapshot-spec.md) — the capture/replay mechanism + the firewall + the capture-source policy.
- [`seeding-spec.md`](seeding-spec.md) — the seeding framework + the 3-layer write-isolation boundary.
- [`idempotency.md`](idempotency.md) — the bring-up **re-run** contract (v1.3b M17). It adds the only new
  destructive ops since this contract was written — the replay re-run `TRUNCATE` and the `stackseed --reset`
  truncates — and they obey it byte-for-byte: every `TRUNCATE` targets a **per-stack-isolated offset** store
  only (pinned by a target-class test), never prod, never a shared store.
- [`../architecture/security_compliance.md`](../architecture/security_compliance.md) — the platform's own
  security/compliance posture (the layer below the tooling).
