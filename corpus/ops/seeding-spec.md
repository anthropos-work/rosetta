# Declarative Stack Seeding — Spec

**The reference for `rosetta-extensions/stack-seeding/`** — how a demo or dev stack is backfilled with
structural data from one declarative blueprint, the dependency order, and (most importantly) the
**production-isolation boundary** that keeps a non-prod seeding run from polluting production.

> **Scope.** This doc covers the **M7a framework**: the blueprint, the seeder contract, the dependency-DAG
> orchestrator, the direct-Postgres perf path, and the isolation guard. The full **seeder fleet** (taxonomy,
> content, sessions, backdated activity at scale) is M7c; the **data-DNA** schema-conformance/drift gate is
> M7b. The seeding code lives in the gitignored `rosetta-extensions` monorepo (its own git), consumed by the
> `/demo-*` skills — **no platform repo is modified.** It is also **not** scattered in the rosetta corpus: it's
> authored and tagged in the authoring copy at `.agentspace/rosetta-extensions/`, then consumed per-stack at a
> pinned tag (`stack-<role>/rosetta-extensions @ <tag>`).

## For PMs — what it does

You describe a target demo world in one file — *"an org of 1,000 users with months of activity"* — and the
seeder builds it into an isolated demo stack by talking **directly to that stack's database**. The headline
property is **safety**: a few platform data stores are *shared* across all environments (the Directus content
system, one S3 bucket, the dev login system). Writing demo data to those would leak onto the live product. The
seeder makes that **structurally impossible** on a non-production stack — it blocks those writes, repairs the
environment before it starts, and produces an **audit log that proves** nothing leaked.

## For engineers

### The blueprint — `stack.seed.yaml`

```yaml
stack: demo-1                 # target stack; the trailing N sets the Postgres offset port (5432 + N*10000)
org:
  name: Acme Corp
  slug: acme-corp
size: 1000                    # N users
role_mix:                     # ratios in [0,1], summing ~1.0
  admin: 0.05
  member: 0.65
  candidate: 0.30
  admin_emails: [founder@acme-corp.test]   # the fixed admin set
tier_mix:
  free: 0.8
  premium: 0.2
content_pack: standard        # carried for M7c (the content/taxonomy seeders)
activity:
  months: 3
  pass_rate: 0.7
```

`Validate()` enforces: `size > 0`, each mix in `[0,1]` and summing to ~1.0, a non-empty org name, ≥1 admin
email, `months ≥ 0`, `pass_rate ∈ [0,1]`.

### The perf path — direct Postgres, not ent-linking, not CLI-shelling

The seeder is a **host Go process** that connects to the stack's Postgres on its **offset port**
(`5432 + N*10000`; dev/`anthropos` = N=0 = `:5432`) and bulk-inserts via `COPY`. Two paths were rejected:

- **Linking `app/internal/bootstrap`** — impossible: it's an `internal/` package, unimportable from any module
  outside `github.com/anthropos-work/app`; importing it would require code inside the platform tree
  (forbidden).
- **Shelling `bootstrap-user` per row** — the DB-IO-bound bottleneck (1k users = 1k subprocess spawns).
  `COPY` is 10–100× faster. (The bottleneck is the database, not the language — Go that `COPY`s beats Rust
  that loops.)

Hand-written SQL drifts from the live schema over time; **M7b's data-DNA** (schema introspection) is the
mitigation — the seeder's output is gated against the platform's current schema, not hand-maintained.

### The dependency order (a DAG, executed with parallelism)

Each seeder declares `DependsOn()`; the orchestrator topologically sorts and runs independent seeders
concurrently. The canonical order:

```
migrate → Sentinel policy → org → users → memberships + casbin + feature → taxonomy/content (snapshot) → activity
```

M7a ships the spine: `org` → `users` (bulk COPY) and the `identity` seeder. M7c fills the fleet. **M9b** wires the
`taxonomy` snapshot node into the DAG (the public skiller catalog, replayed out-of-band; `activity` orders behind
it). **M10** wires the `content` snapshot node (the public Directus template library) and orders the
session/assignment seeders **behind** it, so their `sim_id` / `skill_path_id` / `resource_id` refs resolve against
the **real replayed public templates** (the M10 linkage; free-value fallback when no content snapshot is replayed) —
see [`snapshot-spec.md`](snapshot-spec.md#the-directus-content-surface-m10--the-second-real-surface).

### The production-isolation boundary (the safety contract)

The only true cross-stack / prod-pollution vectors are a **small, fixed set of shared/external services** —
everything in the per-stack Postgres/Redis is inherently isolated (each stack has its own container).

| Store | Class | Why / guard action |
|---|---|---|
| **Directus** | shared-pollution-risk | one global instance (`content.anthropos.work`), visible on prod → **writes blocked**; content arrives via snapshot-replay into the **per-stack** Directus Postgres (M10), never the shared one |
| **S3 public bucket** | shared-pollution-risk | hardcoded to the prod bucket in compose → `STORAGE_S3_PUBLIC_BUCKET` forced to `""` (local fallback) |
| **Live Clerk** | shared-pollution-risk | shared dev app → routed to **Clerkenstein**; a real-Clerk base URL is a hard preflight error |
| **Customer.io / Brevo / AI APIs** | shared-pollution-risk | external; blocked on non-prod (off by default) |
| **coresignal** | external | read-only; writes blocked on non-prod |
| **Postgres / Redis / S3-private / pgvector** | per-stack-isolated | seed freely |

> **The "shared skillpaths across stacks" myth.** Legacy skillpath sessions appearing across stacks trace to
> the *shared Clerk identity* (sessions for a shared login leaked via shared auth) — **not** a shared skillpath
> DB (skillpath data is per-stack Postgres). Clerkenstein isolates identity per stack, so this is already
> neutralized; the seeder just declares each surface's isolation class correctly.

The guard is **three layers** (`rosetta-extensions/stack-seeding/isolation/`):

1. **`Guard.CheckWrite(store, class, target)`** — refuses any shared/external write on a non-prod target.
   *Asymmetry:* `AllowSharedOptIn` only relaxes a **prod** target; **non-prod can never write a shared store,
   regardless of opt-in.**
2. **`Guard.PreflightEnv(env, target)`** — before seeding: forces the S3-public override; rejects a real-Clerk
   base URL; rejects a live Directus write token (non-prod).
3. **`AuditLog.AssertClean(target)`** — after seeding: errors if any *allowed* write to a non-per-stack store
   actually landed on a non-prod target. The **proof** of zero pollution.

### The minimum-proof identity + the casbin gotcha

Beyond the bulk users, the `identity` seeder creates the real login identity matching
`clerkfrontend.DefaultDemoUser()` — `user_clerkenstein` / `demo@anthropos.test` / org `org_clerkenstein`
(role `admin`) — as an admin member of the org, plus the casbin **`g2`** grant so Sentinel authorizes it. A
browser login mints exactly that identity, so an authorized route returns **200, not 403**.

**Two casbin subtleties the live proof caught (both load-bearing):**
- **Arg order.** Sentinel's casbin model matches `g2(org, sub, role)`, so the stored row is **`(org, user, role)`**
  — `v0=org, v1=user, v2=admin`. The original seeder had org/user swapped, which made *every* org-feature check
  silently `403`. Verified against the live working `casbin_rules` row.
- **The table-name gotcha.** `init_policy.sql` seeds the **plural** `casbin_rules`; the gorm adapter auto-creates
  the **singular** `casbin_rule`. A running stack uses one of them. The seeder **introspects** which exists in the
  `sentinel` schema and targets that.

**The global policy prerequisite.** Authorization also needs the platform's **global** policy (`init_policy.sql`'s
~47 role→feature `p`/`p2`/`p3` rows — identical across all stacks). That is *platform bootstrap*, not demo data, so
it is applied by the **demo-stack bring-up** (`migrate-demo.sh`), not the seeder. The seeder supplies the demo
data (org, users, identity, the per-user `g2` grant) on top of it.

### The CLI

```bash
stackseed --seed demo.seed.yaml --validate                       # schema + semantic checks
stackseed --stack demo-1 --seed demo.seed.yaml --dry-run         # ordered plan + row estimates + isolation preview
stackseed --stack demo-1 --seed demo.seed.yaml --dsn '<base-dsn>' # seed (direct Postgres, offset port)
stackseed --stack demo-1 --reset                                 # per-stack reset (refuses n=0 dev unless --force)
```

### Verifying a seed — `datadna` (the data-DNA CLI, M7b)

After seeding, the **`datadna`** CLI (built from `rosetta-extensions/stack-seeding/cmd/datadna`) measures that
the seeded data **conforms to the platform's current schema** and detects schema **drift** — the data dimension
of the alignment framework (full reference:
[`../architecture/alignment_testing.md`](../architecture/alignment_testing.md) § "The data dimension"):

```bash
datadna catalog   --dna <dna.json>                       # the seedable-surface catalog (seeded / snapshot / waived + coverage)
datadna introspect --stack demo-1 --dna <dna.json>       # capture each seeded surface's live shape (the contract)
datadna measure   --stack demo-1 --dna <dna.json>        # conformance score; exit 1 if critical < 100%
datadna measure-snapshot --stack demo-1 --dna <dna.json> --manifest <taxonomy.json> --manifest <directus.json>  # snapshot-fidelity gate (M9b taxonomy + M10 content)
datadna diff      --stack demo-1 --dna <dna.json>        # recorded shapes vs the live schema; exit 1 on drift
```

The shipped manifest is `rosetta-extensions/stack-seeding/dna/data-dna.json` (8 seeded surfaces conformance-gated +
**2 snapshot-seeded**). **As of M10 NOTHING is left waived → coverage reads 100% over the full catalog** (the v1.2
thesis complete). The two formerly-waived surfaces are both promoted and **count toward coverage**:
- **`taxonomy`** (M9b): `waived-m7c → snapshot-seeded-m9b` — the public skiller catalog, fidelity-gated by five
  snapshot operators (row-count / structural / referential / embedding-dim / public-only).
- **`content`** (M10): `waived-m7c → snapshot-seeded-m10` — the public Directus template library (schema `directus`,
  predicate `private=false AND tenant_id IS NULL AND status='published'`), fidelity-gated by four operators (the four
  above **minus embedding-dim** — content has no vectors), the **public-only gene measured against the directus
  predicate** (not `organization_id`).

Both are DAG nodes (`… → taxonomy / content (snapshot) → sessions/assignments → activity`): verification/ordering
nodes whose actual data load runs out-of-band (`stacksnap replay --surface taxonomy|directus`) before `stackseed`.
The `content` node additionally **gates the linkage** — sessions/assignments order behind it so their content refs
resolve against the replayed public templates.

## Status

M7a delivers the framework + the isolation guard + the reference seeders (`org`, `users`, `identity`),
**62 unit tests** (green, `-race` clean), verified schema-correct against a live stack. M7b adds the data-DNA
schema-conformance/drift gate; M7c the full seeder fleet + backdated activity + presets.
