# Rosetta Tooling — Safety & Security Contract

**The authoritative statement of how the `rosetta-extensions` stack tooling stays safe.** Two structural
guarantees, proven in code and tested — the first of which, **since v2.5, carries one bounded, disclosed
exception**:

1. **The snapshot path never reads private/customer data.** Anything that *leaves* production through a
   **snapshot capture** is **public reference data only** — enforced by a tenant-data firewall that hard-fails on
   a single customer-scoped row.
   > 🔴 **This is no longer the only production read (v2.5).** `cmd/content-capture` — the content-story
   > authoring tool — is a **second, deliberately customer-scoped** prod read: it copies the real free-text of a
   > pinned list of production job-simulation sessions, scrubs detectable PII best-effort, and sits **outside**
   > this firewall **by design** (it does not import it; there is nothing public about a played session's
   > transcript). It is read-only, authoring-time, source-pinned, and disclosed — but it is **not** covered by
   > the sentence above. **§3.8 is its contract.** Read it before citing any unqualified *"the tooling never
   > reads customer data"* claim from this corpus — including older sentences in this file.
2. **It never touches production data or services** *(no exception — this one is unqualified)*. Anything a non-prod stack *writes* is confined to that
   stack's own isolated stores — enforced by a 3-layer write-isolation guard that makes a shared/prod write
   **structurally impossible** on a non-prod target, and an audit log that *proves* nothing leaked.

…and, since v2.3, a third axis that is **a disclosure, not a guarantee**:

3. **Who can REACH a demo, and what they get if they do** (**[Part 3](#part-3--the-exposure-side-who-can-reach-a-demo-and-what-they-get-if-they-do)**).
   A demo is an **unauthenticated, authz-weakened build**, and its container ports are published on **all
   interfaces** on **every** bring-up — today, flag or no flag. That cannot be promised away. What makes it
   defensible is guarantees 1 and 2: **there is nothing behind the door** — **except on a content-story demo,
   where there is** (§3.8), and where the VPN/tailnet scope becomes the control instead (§3.3.1). Read Part 3
   before exposing a demo, and before trusting any sentence in this corpus that says a demo binds loopback.

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
finds even one customer-owned row, so **no customer's private data can be copied by a snapshot**.

**Since v2.5 there is exactly one deliberate exception to that, and this is it.** To make "content stories"
believable, a *separate* authoring-time tool copies the real text of a short, hand-picked list of production
sessions — the conversation, the AI feedback, the report — scrubs the names, emails and identifiers it can
detect, and ships the scrubbed result as a checked-in fixture. That is **real customer-authored content,
best-effort anonymized — not provably anonymous.** The data-controller accepted the residual re-identification
risk (2026-07-19) on the condition that such demos are reachable only over a VPN/tailnet. **v2.6 "sound check"
widens that exception to the two MEDIA facets — the raw call audio recording and the full document body — so a
manager can actually hear the call and read the document (§3.8.1).** A recorded voice **cannot** be scrubbed the
way text can, so the audio is copied whole and the VPN scope is its *only* control; the data-controller signed off
**afresh (2026-07-21)** for raw voice + full documents on that same VPN-only condition. It is source-pinned and
auditable in one file, and it changes nothing on the write side. See **§3.8 / §3.8.1** — and do not repeat the
unqualified promise above without them.

On the write side, a small
set of stores are *shared* with the live product (the content system, one storage bucket, the login system);
the tooling **blocks every write to those from a non-production stack**, repairs the environment before it
starts, and produces an **audit log that proves** nothing leaked. Neither promise depends on the operator
remembering a flag — both are structural.

A third write surface joined the family in v1.6: the **secret provisioner** moves secret *bytes*
source→gitignored-target to fill each stack's `.env`. It is fenced the same way — **values-blind** (no verb
ever reads, echoes, or logs a secret value), it never re-arms the prod-write path (the prod
`DIRECTUS_TOKEN` is written blank on a non-prod target), and a secret never enters git. See **§2.9** below.

**And one thing the tooling does *not* promise (v2.3).** A demo environment is deliberately built with its
**login checks switched off** — that is what makes it a demo you can hand to a presenter, who clicks a name and
is instantly "signed in" as that person. It also means **anyone who can reach a running demo over the network
can do the same**, with no password. The tooling does **not** claim otherwise, and — contrary to what one of our
own docs used to say — a demo's ports have always been open on the machine's network interfaces, not just to the
machine itself. **The reason this is acceptable is the first two promises above: a demo contains no customer
data and cannot write to production.** There is nothing behind the door — **with one bounded, disclosed
exception (v2.5, WIDENED v2.6): a "content-story" demo carries the REAL content of production sessions, COPIED and scrubbed of
detectable PII best-effort (not guaranteed clean — residual re-identification risk is accepted, VPN/tailnet-scoped,
a data-controller decision; §3.8) — and since v2.6 also the REAL call audio + full documents, the audio
UNSCRUBBABLE and VPN-scope-only, signed off afresh 2026-07-21 (§3.8.1).**  See **Part 3**.

---

## Part 1 — The read side: the snapshot capture never reads private/customer data

A snapshot **capture** is the **firewalled** production read, and everything that protects it lives in
`stack-snapshot/`. It was the **only** operation that read production until v2.5; it is now one of **two**.

> ⚠️ **The second read is `cmd/content-capture` (v2.5), and it is deliberately outside this Part.** It is
> customer-scoped on purpose — a played session's transcript, feedback and report have no public subset to
> filter to — so it does not, and cannot, run under `AssertPublicOnly`. Its fences are different in kind
> (source-pinning, a best-effort scrub with a fail-closed post-condition, a checked-in auditable fixture, and a
> VPN/tailnet exposure scope) and are contracted in **§3.8**. **Everything in Part 1 describes the snapshot
> read only.** It must not be cited to vouch for the content-story read.

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
| **taxonomy** (`public` schema, formerly skiller) | `organization_id IS NULL` (`firewall.DefaultPredicate` / `firewall.PublicFilter`) | `public.skills` 42,790 public / 794 customer |
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

> **The v1.9 M34 verified-skill chain inherits this class.** The `PersonaSeeder`'s six new write surfaces —
> `jobsimulation.{sessions, validation_attempt_results, validation_attempt_skill_results,
> validation_criterion_results}`, `public.local_jobsimulation_sessions`, `public.user_skills`,
> `public.user_skill_evidences` — are all the stack's own offset-port Postgres, declared `PerStackIsolated`,
> so the chain cannot touch prod or another stack and the seeding-run audit log proves zero pollution
> (`AssertClean`). The taxonomy it reads to draw skill node-ids is the **public** reference data the snapshot
> firewall already guaranteed public-only at capture. See [`demo/stories-spec.md`](demo/stories-spec.md) § Safety.
>
> **The v1.9 M36 dashboard surfaces inherit it too.** The six new dashboard seeders' write surfaces —
> `public.{membership_skills, tags, membership_tags, organization_target_roles, user_target_roles,
> organization_assignment_sessions, local_skill_path_sessions, job_simulation_feedbacks}` and
> `jobsimulation.interview_extraction_results` — are likewise the stack's own offset-port Postgres, all
> `organization_id`-scoped per story and declared `PerStackIsolated` + audited. No new shared store is touched,
> so the zero-pollution posture is unchanged.

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
- **The demo's ONE sanctioned cross-stack read is the shared-secret `.env` seed — never the build SOURCE
  (v1.8 "understudy" M26).** A self-contained demo builds entirely from `stack-demo`'s **own** clone set
  (`ensure-clones.sh` bootstrap-clones `stack-demo/platform` + `make init`s the peer repos; the per-demo
  injection COPY is cut from `stack-demo/<svc>`, never `stack-dev`). The **sole** sanctioned read of `stack-dev`
  by the demo tooling is ensure-clones' phase-(b) `.env` *seed*: `cp stack-dev/platform/.env →
  stack-demo/platform/.env` **copy-if-present + target-absent + never-clobber** — only the shared-secret file
  (same Clerk app + same `GH_PAT`, shared by nature; never committed), and **non-fatal if `stack-dev` is absent**
  (M30's provisioner then writes the real `.env` from `.agentspace/secrets`, so a box with only `stack-demo/` is
  fully supported). The build SOURCE **never** falls back to `stack-dev` — a required platform clone failure
  aborts loud rather than borrow dev's repos, and dev-image reuse is OFF by default (`DEMO_REUSE_DEV_IMAGES=1`
  opts back in). The `TestRenameDrift` suite (esp. `test_ensure_clones_reads_stack_dev_only_for_secrets`) fences
  that every code-level `stack-dev` reference in the demo tooling is confined to that `.env`-seed read. (#M26-D4)

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
The v1.8/M26 `--reuse-dev-images` seam is a second instance of the canonical offender — `up-injected.sh`
assembles `rd_flag=(); [ "${DEMO_REUSE_DEV_IMAGES:-0}" = 1 ] && rd_flag=(--reuse-dev-images)` alongside
`ui_flag`/`lc_flag` and passes all three to `gen_injected_override.py` via `${rd_flag[@]+"${rd_flag[@]}"}`.
`TestReuseFlagArrayExpansion` (in `demo-stack/tests/test_frontend_build.py`) extracts the real assembly+call
block and runs it under `set -euo pipefail` on bash 3.2 across all 8 (ui, lc, reuse) flag combinations,
asserting the empty (default-reuse-OFF) path never trips `set -u` — a second pin of the same rule on the most
common bring-up arrangement (a generator call with several conditional flags). (#M26-harden)

### 2.9 Secret provisioning is values-blind and never re-arms the prod-write path (v1.6 M27–M30)

The secret-provisioning tooling (`stack-secrets/`, driven by [`/stack-secrets`](../../.claude/skills/stack-secrets/SKILL.md)
— full mechanism in [`secrets-spec.md`](secrets-spec.md)) **moves secret bytes** (it writes each repo's `.env`).
That makes it the one new write-side actor since this contract was written, so it carries its own clause. Two
inviolable guarantees, each pinned by a test:

1. **Values-blind — no verb ever reads, echoes, logs, or persists a secret VALUE.** Every command (`list`,
   `check`/`measure`, `introspect`, `diff`, `provision`) emits key NAMES + presence only — at most a value's
   *shape* (a `url`/`jwt`/`pk_`/`sk_` structural prefix via the single permitted `ClassifyShape`, which returns a
   shape token, never the value). Coverage extraction is name-only (cut on the first `=`); the value half is
   discarded the instant a line is parsed. **`provision` necessarily copies secret bytes** source→target — but
   they live ONLY inside the value-carrying boundary (`provision/io.go`'s `sourceValues` → `writeTargetFile`)
   and never surface in stdout, stderr, an error, a return value, or any committed file; the only destination is
   the (gitignored) target `.env`. A hard test (`provision_safety_test.go`) asserts no value ever escapes. The
   `secret-dna.json` manifest is NAMES-only and committable; a `.env` never enters git. This is the same
   values-blind discipline `Guard.PreflightEnv` (§2.2) embodies — **and the provisioner emits
   `PreflightEnv`-passing env**: it writes exactly the env state the seeding guard would accept (the S3-public
   override is left to the override; live-Clerk/live-Directus tokens are never written into a non-prod stack).

2. **Never re-arms the prod-write path (the `DIRECTUS_TOKEN` non-rearm — the blocks-release class).** `provision`
   runs **before** the demo/dev injection override (`gen_injected_override.py`, fix16/fix17) that strips the prod
   `DIRECTUS_TOKEN` to `""` on a non-prod / `--local-content` stack (§2.3). It **must defer to that strip** —
   writing a non-empty prod token into a non-prod stack's base `.env` would re-arm the closed tenant-data-leak
   path. The mechanism (`provision.StripOnNonProdKeys`): the Directus write-token family (`DIRECTUS_TOKEN` /
   `DIRECTUS_STATIC_TOKEN` / `DIRECTUS_ADMIN_TOKEN` — the exact set `PreflightEnv` rejects in §2.2) is **never
   provisioned with a value on a non-prod target**; it is written **blank** (`KEY=`), exactly the state the
   override forces, so the base `.env` and the override agree and the prod-write path is never re-armed. (The DNA
   marks `DIRECTUS_TOKEN` `key-present`-only, no `nonempty`, so a deliberately-blanked non-prod value still passes
   coverage.) A **prod** target is reachable only via the N=0 `--force` path, so the prod token is never
   auto-touched either.

The **N=0 guard** carries over from §2.5: `provision` refuses the main dev stack (N=0, `anthropos`) without
`--force`, because N=0 holds the operator's real source `.env` — a mirror of `stackseed --reset`'s N=0 refusal.
And the **demo-aware coverage check** never weakens the safety posture: on a demo it counts the Clerkenstein-minted
Clerk keys as satisfied (they are minted at bring-up, not sourced), values-blind, by NAME only — a dev stack still
requires the real keys. (#M27 #M28)

### 2.10 A demo's app holds real AWS Bedrock creds (v2.6 M239)

**A disclosed secrets-posture shift, not a new hole.** Until v2.6, the demo's AI-provider keys were
**absent-by-design** (M50, [`secrets-spec.md`](secrets-spec.md#the-bedrock-cred-class-for-app-v26-m239-talk-to-data)):
a demo's believability renders from seeded structural data, so no live model key was ever provisioned. **Talk
to Data** (`/enterprise/talk-to-data` → `app/internal/askengine`) breaks that — it is a *live-inference*
feature that cannot answer without a real model call — so the user decided (2026-07-20) to wire **real AWS
Bedrock creds** into the demo's `app`. This is the **first present-not-absent cloud credential a demo's `app`
carries**, and it is recorded here honestly rather than left implicit.

**What the credential can and cannot do (the blast radius).** The demo **uses** it only for **Bedrock runtime
inference** — `askengine/bedrock.go` calls `InvokeModel` on `eu.anthropic.claude-sonnet-4-6` in `eu-west-1`
and nothing else. It is **not** wired to any data plane: the demo's `app` reads no customer data with it
(Parts 1+2 already make customer data unreachable from a demo), writes nothing to prod, and touches no S3 /
Directus / DB through it. Its worst-case *within the demo* is model-inference spend, not data exposure — a
different, far smaller risk class than the tenant-data and prod-write vectors Parts 1+2 close. **The one caveat
worth stating plainly:** the *credential itself* has whatever permissions its IAM principal was granted, which
the **operator** controls when they provision it — the tooling neither creates nor scopes it. So the operator
**should provision a minimally-scoped, inference-only IAM principal** (ideally `bedrock:InvokeModel*` on the EU
inference profile, nothing more), so that even the worst case — an attacker who reaches the demo and exfiltrates
the env — inherits only Bedrock inference, not a broad AWS key. That scoping is an operator responsibility this
doc names, not a tooling guarantee.

**How it is fenced.** (1) **Operator-provisioned, never minted or fabricated** — the tooling provisions
*nothing* on its own; the creds only exist in a demo if an operator drops them into the (gitignored) secret
source (`.agentspace/secrets/app/.env`), values-blind (the §2.9 provisioner never reads/echoes a value; the
`bridge_bedrock_creds` copy that lands them in the container is file→file, never surfaced). (2) **Deliberately
NOT critical** (R3) — the two cred genes are `required`·`standard`, so a box **without** them still brings a
demo up cleanly (Talk to Data merely inert), and no coverage gate is weakened. (3) **VPN/tailnet-scoped
exposure** — a demo is an unauthenticated, authz-weakened build published on all interfaces (§3, Part 3); the
control on *who can reach it* is the same VPN/tailnet scope (§3.8) that bounds every other demo surface, and it
bounds this one too. The credential rides inside that same perimeter — it does not widen it.

---

## Part 3 — The exposure side: who can REACH a demo, and what they get if they do

> **This Part is NOT a third "never".** Parts 1 and 2 are **structural guarantees** — a prod read of customer
> data and a prod write are made *impossible*, and the code is written so that forgetting a flag cannot defeat
> them. Part 3 cannot make that shape of promise, because **a demo that nobody can reach is a demo that does
> not work.** What Part 3 does instead is state, precisely and without flattery, **what is exposed, to whom,
> and what an unauthorized reacher would actually obtain.** The safety argument then rests on that last clause
> — not on a claim that nobody can knock.
>
> **Added v2.3 "cue to cue" (M220).** Until then this contract had exactly two axes — read-side and write-side.
> Remote reach was a **third axis with no contract at all**, and v2.3 proposed to make it **default-on**. A flip
> like that cannot ship on a doc edit; it needs an argument written down where it can be attacked.

### 3.1 The disclosure — the ports are ALREADY open, on every demo **and every dev stack**, today

**Every demo container's offset port is published on `0.0.0.0` — ALL interfaces — on EVERY `demo-up`, with or
without `--public-host`.** This is not introduced by remote access. It has been true of every demo since the
injected override existed.

> 🔴 **AND THE SAME IS TRUE OF EVERY `dev-N` STACK — which this section did not say until v2.3 M220 S7.**
> `stack-core/gen_override.py` (the **dev** override builder) constructs its port strings **exactly the same
> way** — bare `"<hostport>:<target>"`, no `127.0.0.1` prefix — so a `dev-N` stack's containers are world-
> published too, on every `dev-stack up`, **with or without** `--public-host`. **Measured, not read:** the
> exposure guard now runs *both* emitters and reports `DEMO: 14 ports → 0.0.0.0` and `DEV: 8 ports → 0.0.0.0`.
>
> **This matters MOST because the dev path is opt-in** (§3.5.3). A reader who learns *"remote reach is OFF by
> default on dev"* will reasonably conclude *"so my dev stack is not exposed."* **That conclusion is false.**
> What the opt-in withholds is the **trusted HTTPS origin on the tailnet** — not the LAN binding, which was
> always there. This is the S0 lie one family over: the guard proved the truth for the demo emitter and the
> corpus disclosed it for demos only, so the dev family inherited the silence.

- `stack-injection/gen_injected_override.py` (demo) emits published ports as **bare `"<hostport>:<target>"`
  pairs** at all three emitters (`directus_lines`, `frontend_lines`, `build_lines`), and
  `stack-core/gen_override.py` (dev) does the same in `build_override`. **Docker's default bind for a bare
  `host:container` mapping is `0.0.0.0`.** There is no `127.0.0.1` prefix anywhere, in either family.
- **On Linux this bypasses the host firewall.** Docker installs its rules in its own iptables chain, consulted
  *before* `ufw`/`firewalld`. A `ufw deny` on the port does **not** block it.
- `BIND_HOST` (`demo-stack/up-injected.sh`) *is* gated on the public-host knob — but it is read **only** by the
  two **host-native** servers (the presenter cockpit and ant-academy, which are plain host processes). **It does
  not touch a single container.**
- 🔴 **…and it only actually CONSTRAINS one of those two. MEASURED on `billion`, M220 S3 (2026-07-14).** On a
  **localhost** demo (`BIND_HOST=""`), with the demo up and `ss -ltnp` read on the host:

  | host-native server | bound to | reachable at the node's `100.x` tailnet IP? |
  |---|---|---|
  | presenter cockpit (`7700+off`) | **`127.0.0.1:17700`** | **no** — connection refused ✅ |
  | ant-academy (`3077+off`) | **`*:13077`** | **YES — HTTP 200** ❌ |

  `BIND_HOST=""` means *"pass no `-H` flag and let each server keep its own default"* — and **`next dev`'s own
  default is `0.0.0.0`**. So **at the time of that M220 S3 measurement** the academy was world-published on
  **every** demo, exactly like the containers, and the *"gated on the knob"* framing was true only of the cockpit.

  > ✅ **LANDED in v2.3 M221 (F-M220-5) — the host-native academy now binds loopback.**
  > `demo-stack/ant-academy.sh:330` passes **`-H 127.0.0.1`** on the localhost path (`-H 0.0.0.0` **only** when a
  > public host is requested), so on a localhost demo the academy binds **`127.0.0.1:13077`**, not `*:13077`.
  > The M220 S3 table above is retained as the **dated** measurement that *drove* the fix — not a current claim.
  > ⚠ **Scope of the fix:** it tightens **only** the host-native academy's bind; **every demo *container* port
  > stays `0.0.0.0` by design** — that half of §3.1's disclosure is unchanged and still true. Fenced by
  > `stack-injection/exposure_claim_guard.py`, extended at M221 to *see* the host-native listeners it was blind to.

  **This is the same false-loopback claim §3.1 exists to retract, one layer up** — and it survived M220 S0
  because the exposure fence (`exposure_claim_guard`) checked the three **container** port emitters and had no
  notion of the host-native servers. An exposure fence that cannot see a whole class of listener will report a
  confident, quietly incomplete pass. *(LANDED: `F-M220-5` at M221 — pass `-H 127.0.0.1` when `BIND_HOST` is
  empty, and the exposure guard was extended to run the host-native emitters too. It was deliberately NOT bundled
  into M220 S3/S4 because it changes the localhost path's behaviour, and the invariant S3 is fenced on is that the
  localhost path stays **byte-identical**.)*

> **`corpus/ops/demo/tailscale-serve.md` claimed the opposite until M220:**
>
> ```text
> RETRACTED — FALSE (shipped v2.2, corrected v2.3 M220):
>   "…no open 0.0.0.0-on-the-LAN surprise beyond the tailnet. Binding `0.0.0.0` is gated on the knob
>    precisely so it is never ambient."
> ```
>
> **That was false**, and it is now retracted in place. A shipped safety
> doc that understates real exposure is the worst failure mode in this project: it doesn't just fail to warn, it
> actively talks a reader *out of* looking. Fenced by `stack-injection/exposure_claim_guard.py`, which derives
> the bind by **running** the emitters and fails if any doc denies it — or if this section stops disclosing it.

**Consequence, and it cuts in the flip's favour:** the exposure delta of making remote reach default-on is **far
smaller than this corpus used to imply**, because the LAN/host-IP exposure is *already there*. The honest framing
of `--public-host` is that it does not open the demo — it makes the already-open demo **usable** (a trusted HTTPS
origin, which Clerk requires for a secure context).

### 3.2 What a demo actually IS — an unauthenticated, authz-weakened build

This is the part that must be said plainly, because every mitigation below is judged against it.

| Weakening | Mechanism | Default |
|---|---|---|
| **Clerk token verification is disarmed** in `app`, `cms`, `jobsimulation`, `skillpath` | `stack-injection/apply-authn.sh` vendors a **disarmed colony** into the demo's clone — `authn/provider/clerk` is replaced by a twin that verifies with a **universal HS256 key** (`INJECTED` in `gen_injected_override.py`) | **always, by construction** |
| **Authorization is short-circuited** on the per-member target-role write path | the `app-targetrole-authz-skip` **demo-patch** | **ON** (`DEMO_NO_AUTHZ_SKIP=0`) |
| **The presenter cockpit is a password-free "become any seeded hero" launcher** | a **bare GET** to `<fapi>/v1/client/handshake?__clerk_identity=<key>&redirect_url=<jump_to>` — the fake FAPI selects the seat and *establishes the session*. No credential is presented at any point | **ON** (served whenever `DEMO_STORIES=1`) |

**So: anyone who can reach the cockpit port is one click away from an authenticated session as any seeded hero —
including the manager vantage.** There is no login to fail. Default-on remote reach makes that panel **ambient on
every box that satisfies the capability ladder**, without the operator opting in.

That is the true statement of the risk. It is not softened anywhere below.

### 3.3 The case FOR default-on remote reach (v2.3, D-DESIGN-3)

Recorded honestly, as the argument that actually carries the decision:

1. **There is nothing behind the door — for a demo that is not a content-story demo.** This was the load-bearing
   mitigation when D-DESIGN-3 was taken (v2.3), and for the demo shape that existed then it is exactly what
   **Parts 1 and 2 guarantee, unchanged**: that demo's data is **synthetic + public-snapshot-only**. The
   tenant-data firewall means **no customer data can be in it** — not "should not", *cannot*, or the capture
   aborts. The 3-layer write guard means a demo **cannot write prod**. An attacker who fully owns one obtains: a
   generated population, the public skills taxonomy every customer already sees, and public Directus content.
   **The authz-weakening is only alarming if there is something to protect, and for that demo shape there is
   not.**

   > 🔴 **v2.5 NARROWS this argument, and the narrowing is real — do not read past it.** A **content-story**
   > demo (§3.8) carries the copied, best-effort-scrubbed free-text of **real production sessions**: real
   > customer-authored content, anonymized where detectable, **not provably anonymous**. For that demo shape the
   > sentence above is **false** — there *is* something behind the door. Argument 1 therefore does **not** carry
   > default-on remote reach for a content-story demo. **Argument 2 does**, promoted from a supporting comfort
   > to *the* control. What actually carries the weight there is stated separately in **§3.3.1**, because a
   > narrower argument deserves to be read as a narrower argument rather than inherited from this one.
2. **A tailnet is not the open internet.** It is an **authenticated WireGuard device mesh** — per-device keys,
   ACL-gated, **no public listener**. Reaching a `*.ts.net` MagicDNS name requires already being an enrolled
   device on that tailnet. "Ambient on the tailnet" means ambient *to colleagues who are already inside*.
3. **The delta is small (§3.1).** The ports are already world-published on the host's interfaces. Default-on
   changes the *usability* of that surface, not — mostly — its existence.
4. **The failure mode it removes is real.** Opt-in remote reach means the presenter discovers, at demo time, that
   the demo is unreachable. That is the defect v2.3 exists to fix.

#### 3.3.1 What carries the weight for a CONTENT-STORY demo (v2.5)

Argument 1 does not hold for this demo shape, so the justification is restated here in full rather than
inherited. It is **strictly narrower**, and it should feel narrower:

1. **The tailnet scope is THE control, not a convenience.** For a synthetic demo, *"a tailnet is not the open
   internet"* (argument 2) is a supporting comfort — that demo would be defensible on a LAN too, because there is
   nothing in it. For a content-story demo it is **the** mitigation, and it is the one the data-controller
   acceptance was explicitly conditioned on (§3.8, bound 2): an authenticated WireGuard device mesh, per-device
   keys, ACL-gated, **no public listener**. **A content-story demo outside a VPN/tailnet is outside the accepted
   posture** — not merely less tidy.
2. **The scrub reduces the risk; it does not eliminate it.** Detectable identifiers are removed and the capture
   **fails closed** if a *sourced* name survives — but residual re-identification risk is **real and ACCEPTED**
   (2026-07-19), not engineered away. Argument 1's *cannot* has no counterpart here.
3. **The exposure is bounded by CONTENT, not by access.** What is in the demo is exactly the pinned sessions in
   the checked-in fixture — a hand-picked, source-pinned, auditable list (`seed-generation-manifest.yaml`), not a
   slice of the production database. An attacker who fully owns a content-story demo obtains *those* sessions'
   scrubbed text, and nothing else. This bound is what keeps the blast radius finite once *cannot* is gone.
4. **Part 2 is untouched.** No demo — content-story or not — can write production. The write-side guarantee
   carries exactly as much weight here as it does anywhere else in this document.

> ⚠️ **The consequence for §3.1's already-world-published ports, stated plainly.** For a synthetic demo, the
> LAN/host-IP exposure §3.1 documents is near-harmless *because* there is nothing behind the door — that is why
> §3.1 concludes the exposure delta "cuts in the flip's favour". For a **content-story** demo that conclusion
> **does not transfer**: those same always-open ports carry scrubbed-real customer content to anyone on the
> host's network, tailnet or not. **Do not bring a content-story demo up on a network you do not trust**, and do
> not cite §3.1's delta argument for one.

### 3.4 The case AGAINST — the residual, stated not dismissed

1. **"Nothing behind the door" is a property of the SEED, not of the BUILD.** It holds because Parts 1-2 hold.
   Anyone who points a demo at a non-synthetic data source — a restored prod dump, a hand-loaded CSV — has
   silently converted an authz-free build into a data-bearing one, and Part 3's whole argument evaporates. The
   capture-source policy (§1.4) is what keeps this honest; **it is now also load-bearing for exposure**, which it
   was not before.

   > 🔴 **v2.5 did exactly this — deliberately. It is §3.8.** The content-story feature points a demo at a
   > **non-synthetic data source** (copied real production sessions) *on purpose*. So this residual is no longer
   > hypothetical on the demo path either: it is a **shipped, disclosed, controller-accepted instance** of the
   > very thing it warns about. The difference between §3.8 and the failure mode described above is
   > **governance, not mechanism** — source-pinned, scrubbed, auditable, and VPN-scoped, versus ad-hoc and
   > undisclosed. §3.3.1 states what carries the exposure argument once it has happened. *An operator who
   > hand-loads a prod dump into a demo gets the mechanism without any of the governance, and none of §3.3
   > protects them.*
2. **Ambient means the operator did not choose.** A default-on surface is reachable by people who never decided
   to publish it — including on a laptop that joins a corporate tailnet later.

   > **Which layer owns this, and what nothing measures (v2.5 M236).** The decision on record is a **layering**
   > one: *restricting who can reach a demo is the VM's and the VPN's job, not the demo stack's* — the stack's
   > only obligation is to **permit** VPN access. It is a scoping stance about ownership, **not** a safety
   > claim, and it leaves this document's Part 3 disclosure untouched. See
   > [`verification.md` § What this doc does NOT verify — reach](verification.md). **The consequence to hold
   > alongside it:** no gate anywhere measures reach, so for a **content-story** demo — where §3.3.1 makes the
   > VPN scope *the* control rather than a comfort — that control is **operator-maintained and unattested**.
   > It is exactly as strong as the network the box is on, and nothing will tell you if it is weaker.
3. **The cockpit is the sharpest edge.** It is the one surface whose *entire purpose* is to hand out sessions
   without credentials.

**These are why the flip is scoped to the demo path only, and why `--no-public-host` exists.**

### 3.5 SUPERSESSION — v2.2's D-DESIGN-1 is reversed, for the demo path only

> **v2.2 D-DESIGN-1 (`.claude/skills/demo-up/SKILL.md`): _"Public reach is never default-on"_** — external
> binding happens ONLY when `--public-host` is set.
>
> **SUPERSEDED by v2.3 D-DESIGN-3, for `/demo-up` only.** Remote reach becomes **default-on with opt-out**
> (`--no-public-host`) on the **demo** path. **`/dev-up` remains opt-in** — unchanged.

**Why the reversal is justified, in one paragraph:** D-DESIGN-1 was written believing that opt-in *withheld* an
exposure. §3.1 shows it did not — the containers were world-published either way, so D-DESIGN-1 was buying less
safety than it appeared to, at the cost of a demo that a presenter could not reach. What it *did* withhold was
the trusted HTTPS origin. Meanwhile the thing that actually makes a demo safe to expose — **there is no customer
data in it, and it cannot write prod** — is a *structural* property (Parts 1-2), not a consequence of the flag.
D-DESIGN-3 therefore moves the default, keeps the escape hatch, and confines the change to the path whose entire
purpose is to be shown to other people.

**Dev stays opt-in** because a dev stack has no such guarantee of synthetic-only content: it is a working
environment, an engineer may point it anywhere, and §3.4's residual #1 is a live risk there rather than a
hypothetical one.

> ⚠️ **Cite it as "v2.2's D-DESIGN-1", never bare.** v2.3 has its **own** `D-DESIGN-1` (*"the <5 s gate is on
> ACCESS, not full first-page render"*). The ids collide across releases; a bare reference resolves to the wrong
> decision.

**Status: LIVE in code as of M220 S3** (2026-07-14; proven default-on end-to-end on `billion`, and proven to
fall back byte-identically on a box with no Tailscale). §3.5 previously recorded only the *decision* — the code
still required the flag. It no longer does.

#### 3.5.1 How default-on actually decides — the capability ladder

Default-on does **not** mean *"assume the box is on a tailnet"*. It means *"find out, and prove it"*. Six rungs,
**capability-gated, never presence-probed** (`demo-stack/tailscale_autohost.py`):

1. `tailscale` is on `PATH`
2. `tailscale status --json` → `BackendState == "Running"` — *installed-but-logged-out is a **failure***
3. `.Self.DNSName` present and **dotted** — a dotless name is **hard-refused** (`@clerk/backend`'s
   `assertValidPublishableKey` rejects a dotless FAPI host, and the host is baked into the publishable key, so a
   dotless host **500s every request** — it is not a degraded demo, it is a broken one)
4. `CurrentTailnet.MagicDNSEnabled == true` — *cannot confirm ⇒ refuse*
5. `tailscale serve status` shows no operator/sudo denial
6. **`tailscale cert` actually MINTS a certificate** — *not "the binary is installed"*. Rungs 1–5 all pass on a
   box where the mint still fails (no operator, tailnet HTTPS off, an ACME hiccup); the cert step then silently
   degrades to a **local-trust** cert that the *remote* browser — the only machine this feature exists for —
   rejects. A green bring-up nobody can use. So rung 6 demands a **certificate**, not a **binary**.

> #### 🔴 The fallback is not optional — and it is a correctness property, not caution
> **Any failed rung ⇒ an EMPTY `STACK_PUBLIC_HOST` ⇒ byte-identical to the v2.2 localhost demo**, plus **one
> loud line** naming the exact fix command. **Never a partially-satisfied public path.**
>
> `SCHEME` and `BIND_HOST` in `up-injected.sh` both derive from the **same `-n $STACK_PUBLIC_HOST` predicate**.
> A **half-satisfied** public path is therefore **strictly worse than localhost**: every baked browser URL flips
> to `https://` while the listeners are still plain HTTP, and the demo **does not load at all**. Today a laptop
> with no Tailscale always works, and it must keep always working. Fenced in
> `demo-stack/tests/test_public_host_flip.py` — including the case where discovery **crashes** (a dropped
> `|| true` would abort the bring-up, leaving a box unable to run *any* demo, not even the localhost one it ran
> yesterday).

#### 3.5.2 The cockpit is now behind the tailnet's TLS — which is **transport**, not **authentication**

M220 S4 adds the presenter cockpit (`7700+offset`) to the `tailscale serve` front list. Until v2.3 it was the
**one** browser-facing surface deliberately left on **plain HTTP** — and with remote reach now default-on that
was the worst possible combination: the demo's **entry point**, the single page a presenter actually opens, was
the single page not behind the trusted cert, while everything it links to was.

**Do not read this as a hardening of the cockpit.** Per §3.2 the cockpit remains a **one-click, password-free
"become any seeded hero" launcher** — a bare `GET /v1/client/handshake?__clerk_identity=<key>`. Fronting it on
`tailscale serve` puts it behind the tailnet's **TLS + authenticated device mesh** instead of cleartext HTTP. It
does **not** password-protect it. Anyone who can reach the tailnet can still become any hero, exactly as before;
they now do so over a trusted origin. The reason that is acceptable is unchanged and structural: **there is no
customer data in a demo, and it cannot write prod** (Parts 1–2).

#### 3.5.3 The DEV path — remote reach is now a real opt-in, where before it was a vacuous one (M220 S7)

§3.5 has said *"`/dev-up` remains opt-in"* since the flip was designed. That sentence was **true and hollow**:
before M220 S7 there was **no `--public-host` on the dev path at all** — nothing to opt *into*. "Opt-in" named a
choice the tool did not offer. S7 builds the flag, which is what makes the asymmetry a **contract** rather than
an accident of what happened to be implemented:

| | remote reach | the escape hatch |
|---|---|---|
| **`/demo-up N`** | **DEFAULT-ON** — the box's own MagicDNS host, auto-discovered (§3.5.1) | opt **out**: `--no-public-host` |
| **`/dev-up N`** | **OFF** | opt **in**: `--public-host auto` \| `<fqdn>` (env: `DEV_PUBLIC_HOST`) |

**The OFF side is the load-bearing half, and it is fenced as such.** A dev bring-up that does not ask for a
public host makes **zero** `tailscale` calls — it does not probe and decline, it does not look. No cert mint, no
`serve` config, no new files, no changed output: byte-identical to the pre-S7 tool. The fence proves this with a
**tripwire** stub (a healthy `tailscale` on `PATH` that fails the test if invoked at all), because "it fell back
safely" would be a passing grade for a tool that probed — which is exactly what the opt-in default forbids.

**Dev reads `DEV_PUBLIC_HOST`, deliberately NOT the demo's `STACK_PUBLIC_HOST`.** `up-injected.sh` **exports**
`STACK_PUBLIC_HOST` for its child launchers. Had dev read that name, a value inherited from an enclosing shell
could have flipped a dev stack world-reachable **with no flag on the command line** — an exposure nobody typed.
Separate namespace, no ambient inheritance.

**Three things the dev path does NOT get, and the reason is the same each time:** a dev stack is not a demo.

- **No Clerkenstein, no minted pk, no fake FAPI** — `/dev-up` uses **real Clerk**. §3.2's *"unauthenticated,
  authz-weakened build"* describes a **demo**, not a dev stack; a dev stack still authenticates.
- **No cockpit.** The password-free "become any hero" launcher is a demo surface (gated on `DEMO_STORIES`). A
  dev stack never starts one, so `--no-cockpit` is passed unconditionally and its port is never fronted.
- **Only the ports the stack actually publishes are fronted** — derived from the generated override, not from
  the demo's fixed registry, because `--profile` decides what a dev stack runs. (`tailscale serve` *binds* the
  ports it fronts; fronting a dead one would hold it against the next bring-up.)

🔴 **What OPT-IN does NOT mean — read this before you conclude your dev stack is private.** It is *not* the case
that a no-flag `dev-stack up` is unreachable from the network. **Every `dev-N` container is published on
`0.0.0.0` already** — see §3.1, where this is now measured for the dev emitter (`8 ports → 0.0.0.0`) and not
merely asserted — and on Linux Docker's iptables **bypass `ufw`/`firewalld`**. What `--public-host` adds is the
**trusted HTTPS origin on the tailnet** (a `tailscale serve` front + a real cert), *not* the LAN exposure, which
was there before you ever heard of this flag. The opt-in governs **reachability by name, over TLS, from another
machine** — not reachability at all.

⚠️ **And the residual that does NOT go away: §3.4's residual #1 is a LIVE risk on dev, not a hypothetical one.**
A demo's content is synthetic + public-snapshot-only by construction (Parts 1–2; the one bounded exception is a
content-story demo's copied+scrubbed real session content, §3.8). **A dev stack has no such
guarantee** — it is a working environment, an engineer may point it at anything, and `/dev-up`'s own default is
to read content **live from prod**. That is the entire reason dev is opt-in and demo is not, and it is why this
flag asks you to say so out loud. As on the demo path, what you are turning on is **transport, not
authentication** (§3.5.2): the tailnet's TLS + device mesh, not a password.

### 3.6 The EGRESS half — what a demo sends OUT (v2.3 M220 S5/S6)

§3.1–3.5 are about who can reach **in**. The mirror question went unasked for the whole demo family: **what does
a demo reach OUT to?** A corpus that calls the demo *self-contained* owed an answer, and the honest one was bad.

| what phoned home | measured | now |
|---|---|---|
| **Google Analytics · DoubleClick · Google Ads · LinkedIn Ads** | on **every page load**, via a **hardcoded** `<GoogleTagManager gtmId='GTM-PXRTBZK'/>` in next-web's root layout | **gone** — the `next-web-no-thirdparty` demo-patch |
| **Plausible · analytics.bellasio.com · uptime.betterstack.com** | same file, same three hardcoded `<Script>` tags. **Not in the plan** — found by reading the file. A presenter demoing to a customer was shipping that customer's page views to **seven** third parties | **gone** — same patch |
| **Clerk telemetry** (`clerk-telemetry.com`) | real egress from **both** frontends. `TELEMETRY_DISABLED` had **zero** hits across the whole tooling repo — it was never wired. Also a reason Playwright's `networkidle` never settles | **off** — `CLERK_TELEMETRY_DISABLED` (server) + `NEXT_PUBLIC_CLERK_TELEMETRY_DISABLED` (browser, build-inlined). **Both halves, or one collector still phones home** |
| **`cdn.jsdelivr.net`** — the clerk-js bundle | the fake FAPI proxied `clerk.browser.js` **live from the CDN on every full page load**, with `http.Get` = `http.DefaultClient` = **`Timeout: 0` (UNBOUNDED)** and **no cache**. next-web's entire authenticated tree is client-gated on clerk-js ⇒ **an unbounded internet dependency ON THE LOGIN PATH**: 0.2 s healthy, **~127 s if egress blackholes** — a presenter stuck on a white page with nothing to do but wait | **served from disk.** A box-level cache (shared by every `demo-N`) keyed by the request path's `package@version` — self-invalidating by construction. The CDN survives only as a **bounded** (15 s) fallback that populates the cache atomically and never caches a non-200 |
| **real Clerk** (`api.clerk.com`) | the academy ran **keyless** and phoned Clerk to provision a throwaway app — holding the **REAL Clerk app's `CLERK_SECRET_KEY`**, copied out of `platform/.env` | **gone** — Clerkenstein-wired (see [`demo/frontend-tier.md`](demo/frontend-tier.md)). Same class as the `DIRECTUS_TOKEN` strip of §2.9 |

**This is the item that most directly contradicted this document.** An unbounded, uncached internet fetch in the
login path of a demo the corpus describes as self-contained is not a performance note — it is a false claim in a
safety doc, which §3.1 already established is the worst failure mode in this project.

**Verified from a tailnet peer, in a real browser, on an authenticated page load: zero requests to any of the
above.** The check asserts it captured traffic at all — an empty scan is a FINDING, not a pass.

> **One new listener, and it is the demo's FIRST loopback-bound port.** The disarmed fake BAPI is now published
> on **`127.0.0.1:5401+offset`**, because ant-academy is the demo's one **host-native** frontend and cannot use
> the in-network `api.clerk.com` alias — without it, its only reachable `CLERK_API_URL` is **real Clerk**. It is
> bound to loopback precisely *because* §3.1 established that every other port is world-published: a mock that
> ignores the bearer token entirely is the last thing that should be ambient on a tailnet.
>
> **Reconciled at v2.3 M221 (F-M220-5):** the fake BAPI is no longer the *only* host-native listener bound to
> loopback — M221 tightened the **ant-academy** `next dev` bind to `127.0.0.1` on a localhost demo as well (§3.1).
> The "every other port is world-published" reasoning still holds for the demo **container** ports (unchanged);
> among the **host-native** listeners, the cockpit was already loopback, and on a localhost demo the academy and
> this fake BAPI now join it — all three bind loopback.

### 3.7 What this does NOT change

Parts 1 and 2 hold **exactly** as written **for every demo the tooling built before v2.5, and for every demo
that is not a content-story demo.** Remote reach changes the *origin and scheme* a browser uses; it does not touch
the data plane. The tenant-data firewall, the public-only predicates, the read-only capture policy, the 3-layer
write-isolation guard, the never-write-prod boundary, and the values-blind secret contract are all unaffected —
and they are precisely what makes §3.3's argument work. **Part 3 is a debt that Parts 1 and 2 pay.** The one place
Part 1's "nothing behind the door" gains a bounded, disclosed exception is §3.8 (content-story demos); Part 2
(never-write-prod) is untouched there too.

### 3.8 The content-story exception — anonymized-real session data (v2.5 "the playbill", M232)

Until v2.5 the read side promised the strongest possible thing: **a demo carries only synthetic + public-snapshot
data; no customer datum, public or private, that a person authored is ever in it.** The v2.5 "Content stories"
feature adds a demo tab of **played sessions a presenter logs into** — and to make those believable, the tooling
**clones real production job-simulation sessions** into the demo. That is a deliberate, **user-accepted
(data-controller) decision**, and it is a genuine — if narrow — softening of the promise above. This section
records it honestly rather than letting the doc keep asserting something the tooling no longer does.

**What is actually cloned — the REAL content, COPIED and best-effort SCRUBBED (2026-07-19 data-controller
decision).** The interesting part of a played session IS its free-text — the real conversation, the real LLM
feedback, the real submission, the real interview report. So the tooling **copies that real content**. At
**authoring time**, `cmd/content-capture` reads production **read-only** (`marco_read` via `~/.pgpass` over
Tailscale — `db-access.md`; it `SET`s the session read-only and only `SELECT`s), COPIES each pinned session's
result-fan-out content, and **SCRUBS the detectable PII** before writing the checked-in fixture:

- the source session's real person-names — the transcript **actors AND the session OWNER's real identity**
  (`sessions.owner_id` → `public.users.firstname`/`lastname` + the email local-part; the candidate's first name is
  threaded through the LLM feedback and comes from HERE, not from the empty `jobsimulation.actors` names) →
  `<<ACTOR_i>>` placeholders, **token-split** so a bare first-name mention is caught; the **source org name** →
  `<<ORG>>` (the seeder fills these with the demo persona/org); **emails, URLs, and long digit-runs** → redaction
  markers. *(M235 leak fix, 2026-07-19: the original capture sourced only the empty actor names → removed **zero**
  names → 8/9 fixtures shipped a real first name. The owner-identity path + a **capture-time fail-closed
  post-condition** — the capture refuses to write a fixture in which a **sourced** name survives — closed it.)*

**This is NOT "provably clean".** Free-text scrubbing is imperfect: a name the pass never *sourced* (a third party
mentioned in passing), an unusual identifier format, a company mentioned in passing can survive. Raw customer
content is streamed prod → scrub → fixture (it never enters an agent's context, and `content-capture` prints
counts only, never content — a leak error prints only the field name + token length), and the shipped fixture is
re-scanned for structural PII (emails/URLs/phones) by a test gate — but **residual re-identification risk is real
and was ACCEPTED by the data-controller.** The word "anonymized-real" is precise the honest way: the content is
real, anonymized *where detectable*, not guaranteed anonymous.

**The bounds that make it acceptable:**

1. **Best-effort scrub** — the detectable identifiers (known names — actors **+ the session owner's real
   identity**, token-split — org, emails/phones/URLs) are removed; the capture **fails closed** if a *sourced*
   name survives, and `TestEmbeddedContent_NoStructuralPII` re-scans every shipped blob for structural PII + the
   name-scrub-fired tripwire. It is a diligent pass, **not a guarantee**.
2. **Residual risk ACCEPTED, VPN/tailnet-scoped** — the data-controller accepted the residual re-identification
   risk (2026-07-19); the CONTROL on it is that a content-story demo is exposed under the Part-3 posture
   (unauthenticated, authz-weakened, world-published on its host's interfaces) **only over a Tailscale tailnet /
   VPN** (`tailscale-serve.md`), never the public internet. The scrub reduces the risk; the VPN scope contains it.
3. **Source-pinned + disclosed** — every cloned session's prod source-id + the copy+scrub posture is recorded in
   `seed-generation-manifest.yaml` (the `content_sessions` block, honesty-gated), so an auditor reads *exactly*
   which real sessions a content-story demo was copied from, in one file, without reading Go.

**Public-anchored + non-manager-played** (two more structural properties, from the M231 sourcing contract): a
cloned session references only a **public-published** sim (so it resolves in the demo's replayed catalog, and no
customer-private sim content is pulled), and it is re-owned to a **player-vantage** seeded member (never a manager
seat) — so the clone is a player's own session, re-tenanted into the demo org.

**Part 2 is untouched.** The write side never changes: the `ContentStorySeeder` writes only to the per-stack
Postgres (PerStackIsolated), audited, n=0-guarded — it can no more write prod than any other seeder; and the
authoring-time read is read-only. The exception is entirely on the read side, and it is bounded by the scrub + the
VPN scope, not eliminated.

**The full contract:** [`demo/session-clone-spec.md`](demo/session-clone-spec.md) (the sourcing pattern, the
firewall-safety argument, the source-pin contract, the no-manager-played rule, and the **copy-real + best-effort-scrub**
mechanism — the real content is copied, detectable PII scrubbed, residual re-identification risk accepted by the
data-controller, VPN/tailnet-scoped) + [`demo/content-stories-routes.md`](demo/content-stories-routes.md) §3.5 (the
M231 spike that authored the posture this section lands).

### 3.8.1 The raw-media extension — voice recordings + full documents (v2.6 "sound check", M240)

Until v2.6 the content-story exception (§3.8) copied only a played session's **free-text** — the transcript, the
LLM feedback, the submission text, the interview report — each **token-scrubbed** best-effort. v2.6 "sound check"
extends the copy to the two **media** facets that make a played session actually playable: the **call's raw audio
recording** (the Chime/S3 recording, re-hosted in the demo storage tier so the manager can press play and hear it,
with `chime_status` flipped to available) and the **full document body** a document session produced (beyond the
small `validation_criterion_results.input_data` reference, whose ~35-char payload is a pointer, not the body). This
is a **strictly larger customer-PII surface than v2.5's scrubbed text**, and the two facets are NOT symmetric —
they carry two different postures, recorded here honestly.

**Voice audio → PORTED AS-IS. A voice is not token-scrubbable, so the scrub bound of §3.8 does NOT extend to it.**
There is no "anonymized-real audio": a recorded human voice is inherently identifying (it carries the speaker's
timbre, accent, cadence, and often their spoken name) and cannot be redacted the way a name-token in text can. So
for the audio facet, §3.8's **bound 1 (best-effort scrub)** has **no counterpart** — the tooling does not, and
cannot, claim to anonymize the audio. The audio is copied whole. Its **only** control is §3.8's **bound 2 (the
VPN/tailnet scope)** — which for the free-text facet was a reinforcing comfort but for the audio facet is **the
entire control**, exactly as §3.3.1 already elevated it to be for the content-story demo shape. Two operational
constraints attach:

- **The bytes never enter an agent's context.** The port streams the recording **prod-S3 → demo storage tier**
  entirely inside the capture tool — it is never printed, cat'd, decoded, transcribed, or read into any agent's
  reasoning, the same counts-only discipline `cmd/content-capture` applies to text (§3.8). An operator orchestrates
  the port; no one listens to the audio in-context.
- **Gender coherence (a believability constraint, NOT a safety one).** The `ContentStorySeeder` re-tenants a real
  session onto a synthetic player persona; when that session carries a voice recording, the **owning persona's
  apparent gender must match the voice's** — a persona minted as "Maria" must not front a male-voiced call. The
  source voice's gender is **labeled at capture time, values-blind** — derived in-tool from the session owner's
  sourced identity (which the scrub already reads and drops), never by an agent listening to the audio — and the
  persona pairing is constrained so the two align. A gender label (`m`/`f`/`unknown`) is not the class of PII the
  scrub removes; the owner's name that produced it never leaves the tool.

**Documents → PORTED + SCRUBBED (bound 1 applies unchanged).** The full document body IS text, so it runs through
the **same** best-effort `scrub` pass the transcripts use (names → placeholders, org → `<<ORG>>`, emails / phones /
URLs redacted), with the same fail-closed capture-time leak post-condition. Its residual is identical to the
free-text residual §3.8 already accepted: best-effort, **not provably clean**.

**Fresh data-controller sign-off (2026-07-21).** The 2026-07-19 acceptance in §3.8 covered scrubbed **free-text**.
Raw customer **voice** is a materially larger call — it is unanonymizable, so the residual re-identification risk
is higher and of a different kind (a colleague could recognize a real voice the scrub cannot touch). The
data-controller therefore signed off **afresh, dated 2026-07-21**, specifically for **raw customer voice recordings
+ full customer documents in a demo**, on the **same condition** that bounds §3.8: such a demo is reachable **only
over a Tailscale tailnet / VPN**, never the public internet, and the pinned session list stays source-pinned +
disclosed. The residual re-identification risk for the audio facet is **real and ACCEPTED (2026-07-21)**; the
control is the VPN scope, not a scrub.

**The gate ordering — posture documented BEFORE any media lands.** No customer audio or document blob is copied
into a demo until this amendment and the 2026-07-21 sign-off are in place. This section IS that gate; it landed
before the seeder's media-port code did.

**Part 2 (never-write-prod) is untouched.** The audio re-host and the document-blob write target only the
**per-stack demo storage tier** (`PerStackIsolated`) — a demo can no more write prod S3 than any other seeder — and
the authoring-time read stays read-only. The extension is entirely on the read side, bounded by the (text-only)
scrub + the VPN scope; for the audio facet, by the VPN scope alone.

> **Current status (M240, honest).** The media-port **tooling** is specified and the posture above is signed off,
> but the actual byte-port depends on **eu-west-1 S3 read access to the recording/media bucket** (the long-standing
> `DEF-M10-01`), which is **not wired in the authoring environment as of 2026-07-21**. Until an operator provisions
> that read access, a content-story demo ships **without** ported audio (`chime_status` stays `not_available`) and
> without any S3-resident document body — exactly the pre-v2.6 state, degraded honestly, never a broken "play"
> button over a blob that isn't there. The posture is documented **ahead of** the capability, by design (the gate
> ordering above). See [`demo/media-substrate-spec.md`](demo/media-substrate-spec.md) for the full mechanism +
> the S3-blocked residual.

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
- [`demo/tailscale-serve.md`](demo/tailscale-serve.md) — the remote-access runbook (**exposure-side**). Its
  § "Safety framing" carries the M220 retraction of the false `0.0.0.0`-is-gated claim; **Part 3 here is its
  safety contract**. Fenced by `stack-injection/exposure_claim_guard.py` (derives the bind by *running* the
  emitters; fails if any doc denies it, or if Part 3 stops disclosing it).
- [`db-access.md`](db-access.md) — the production read foundation + the public-vs-customer boundary (read-side).
- [`snapshot-spec.md`](snapshot-spec.md) — the capture/replay mechanism + the firewall + the capture-source policy.
- [`seeding-spec.md`](seeding-spec.md) — the seeding framework + the 3-layer write-isolation boundary.
- [`secrets-spec.md`](secrets-spec.md) — the secret-provisioning mechanism + the values-blind / `DIRECTUS_TOKEN`-non-rearm
  contract (§2.9 here is the safety statement for it).
- [`idempotency.md`](idempotency.md) — the bring-up **re-run** contract (v1.3b M17). It adds the only new
  destructive ops since this contract was written — the replay re-run `TRUNCATE` and the `stackseed --reset`
  truncates — and they obey it byte-for-byte: every `TRUNCATE` targets a **per-stack-isolated offset** store
  only (pinned by a target-class test), never prod, never a shared store.
- [`../architecture/security_compliance.md`](../architecture/security_compliance.md) — the platform's own
  security/compliance posture (the layer below the tooling).
