# Declarative Stack Seeding — Spec

**The reference for `rosetta-extensions/stack-seeding/`** — how a demo or dev stack is backfilled with
structural data from one declarative blueprint, the dependency order, and (most importantly) the
**production-isolation boundary** that keeps a non-prod seeding run from polluting production.

> **Scope.** This doc covers the **M7a framework**: the blueprint, the seeder contract, the dependency-DAG
> orchestrator, the direct-Postgres perf path, and the isolation guard. The full **seeder fleet** (taxonomy,
> content, sessions, backdated activity at scale) is M7c; the **data-DNA** schema-conformance/drift gate is
> M7b; the shipped **presets** + the **`dev-min` dev auto-seed** (v1.3 M13 — applied on a `dev-stack up` build
> so a fresh dev stack is never empty) are in [The shipped presets](#the-shipped-presets-stack-seedingpresets).
> The seeding code lives in the gitignored `rosetta-extensions` monorepo (its own git), consumed by the unified
> `stack-*` skills (`stack-seed` / `stack-snapshot`) + the dev-stack bring-up — **no platform repo is modified.** It is also **not** scattered in
> the rosetta corpus: it's authored and tagged in the authoring copy at `.agentspace/rosetta-extensions/`, then
> consumed per-stack at a pinned tag (`stack-<role>/rosetta-extensions @ <tag>`).

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

> **The consolidated safety contract** — this write-side boundary plus the snapshot read-side firewall — is
> stated authoritatively in [`safety.md`](safety.md). What follows is the write-side detail.

The only true cross-stack / prod-pollution vectors are a **small, fixed set of shared/external services** —
everything in the per-stack Postgres/Redis is inherently isolated (each stack has its own container).

| Store | Class | Why / guard action |
|---|---|---|
| **Directus** | shared-pollution-risk | one global instance (`content.anthropos.work`), visible on prod → **writes blocked**, the shared instance **never written**. (Reads: since v1.5 **M22/M23** a **`--local-content`** stack (demo default-on; dev opt-in) serves content from its **own** per-stack Directus — `cms` cut over to the in-network instance — so the read is **local, not live-prod**; the prod **data plane** is read only at capture time. A stack **without** `--local-content` reads public content **live** from prod, a demo **anonymously** with the token stripped — the documented fallback. The asset plane stays on prod public links either way. See [`directus-local.md`](directus-local.md) + [`snapshot-spec.md`](snapshot-spec.md#the-per-stack-directus-store-fork-m10-d2-recipe-corrected-in-fix16).) |
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

**Every member, not just the admin (a membership and its `g2` grant are two halves).** The bulk `users` seeder
writes a `public.memberships` row **and** its Sentinel `g2 (org, user, role)` grant for **each** of the N members
— not only the `identity` admin. The reason is the same `g2` resolution: a *per-member* authz check resolves the
**object** member's role via `g2`. The clearest case is the **Members list** (`/enterprise/members`) — resolving
each row's `targetRole` makes Sentinel evaluate the admin's `org:action:assignments:write` **on that member**, and
the `p2` policy is keyed on the member's role. A member with a `memberships` row but **no `g2` grant** has an
*unresolvable* role, so the check returns `false` → the resolver errors → the whole `organizationMembers` query
fails ("Failed to fetch from Subgraph 'backend'") → an **empty Members page**, even though the rows exist. Seeding
the membership without its grant passes basic browsing but 403s any per-member path; the two are seeded together.

**The global policy prerequisite.** Authorization also needs the platform's **global** policy (`init_policy.sql`'s
~47 role→feature `p`/`p2`/`p3` rows — identical across all stacks). That is *platform bootstrap*, not demo data, so
it is applied by the **demo-stack bring-up** (`migrate-demo.sh`), not the seeder. The seeder supplies the demo
data (org, users, identity, the per-user `g2` grant) on top of it.

### The shipped presets (`stack-seeding/presets/`)

Four presets ship, ordered by population — `dev-min` < `small-200` < `mid-500` < `large-1k` (the order is
regression-pinned: a curator picking "small" for a low-resource box relies on it actually being smaller):

| Preset | Users | Activity | Org | Used by |
|---|---|---|---|---|
| **`dev-min`** | **~10** | 1 month | Dev Sandbox | **the dev-stack auto-seed** (M13) — the "never empty" floor |
| `small-200` | 200 | 3 months | Northwind | quick demo walkthroughs / low-resource boxes |
| `mid-500` | 500 | 6 months | Globex | the default "looks real on screen" demo |
| `large-1k` | 1,000 | 9 months | Initech | scale / perf demos |

**`dev-min` — the dev auto-seed (M13).** The lightest preset (~1 org + ~10 users + 1 month of activity) is the
**default applied on a `dev-stack up` build**, so a freshly-built dev stack is **never empty** — a structural
spine to click through — without paying a demo-scale seed (it still seeds in well under a second). 10 users is
the floor that still exercises the role mix (~1 admin + ~6 members + ~3 candidates) so authz / memberships /
activity all render; its fixed admin is **`dev@anthropos.test`** (the local dev login identity → a browser login
to the fresh dev stack returns **200**, not 403). It targets `dev-N` (vs the demo presets' `demo-1`). The
dev-stack bring-up applies it via the set-dressing pass — see [`snapshot-spec.md`](snapshot-spec.md#dev-as-a-full-fidelity-peer-m13--the-set-dress-pass-recipe--auto-snapshot--light-seed). (size rationale: #M13-D1)

### The CLI

```bash
stackseed --seed demo.seed.yaml --validate                       # schema + semantic checks
stackseed --stack demo-1 --seed demo.seed.yaml --dry-run         # ordered plan + row estimates + isolation preview
stackseed --stack demo-1 --seed demo.seed.yaml --dsn '<base-dsn>' # seed (direct Postgres, offset port)
stackseed --stack dev-1  --seed presets/dev-min.seed.yaml        # the dev auto-seed (M13; run by the dev bring-up)
stackseed --stack demo-1 --reset                                 # per-stack reset (refuses n=0 dev unless --force)
```

**Re-run safe (v1.3b M17).** A re-seed is now idempotent: every deterministic-id surface writes via an
`ON CONFLICT (id) DO NOTHING` COPY merge, and the casbin `g2` grant via `WHERE NOT EXISTS` — so a 2nd seed
inserts 0 rows instead of unique-violating or duplicating the grant. `--reset` truncates the **full**
seeded fleet child-first FK-safe (`activity_events →` the M34 verified-skill chain `validation_criterion_results
→ validation_attempt_skill_results → validation_attempt_results → user_skill_evidences → user_skills →
local_jobsimulation_sessions → sessions → skill_path_sessions → assignments → memberships → users →
organizations`; the casbin grant is reset by a targeted `DELETE WHERE p_type='g2'`, never a TRUNCATE — it shares
the sentinel schema with `init_policy.sql`'s global policy). Full per-component re-run contract:
[`idempotency.md`](idempotency.md).

**The n=0-dev guard, two layers (M13).** `--reset` already refuses N=0 (the main `anthropos` dev stack) unless
`--force`. M13's **auto-seed on dev build** adds a second, earlier guard in the bring-up's set-dressing pass
(`dev-setdress.sh`): it **refuses to auto-set-dress N=0 without `--force`**, so an automatic dev seed can never
touch the developer's primary stack. The auto-seed targets the bring-up's own `dev-N` (offset port), never N=0.

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
datadna measure-closure  --stack demo-1                  # v1.9 M34: seed-side verified-skill closure; exit 1 on a dangling skill ref
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
nodes whose actual data load runs out-of-band (`stacksnap replay --surface taxonomy|directus`, or `/stack-snapshot
replay N`) before `stackseed`. The `content` node additionally **gates the linkage** — sessions/assignments order
behind it so their content refs resolve against the replayed public templates. The full mechanism + the operator
recipe: [`snapshot-spec.md`](snapshot-spec.md) + [`demo/recipe-snapshot-world.md`](demo/recipe-snapshot-world.md).
A preset seeds a believable world **without** a snapshot (graceful degradation: empty catalog, free content refs);
replay first for the set-dressed world.

### The verified-skill chain — the believability spine (v1.9 M34)

The seeded surfaces above make a stack *not-empty*; the **verified-skill chain** makes it *tell a story*. A
blueprint **`personas`** list lifts named **heroes** on top of the generic population, and the **`PersonaSeeder`**
writes the **7-table fan-out** a real passed-simulation pipeline would have produced — so a hero's **skill
profile** + Skill Spotlight chart render, with the **claimed-vs-verified gap** the demo turns on. It also fixes
the **G14** session bug (the pre-M34 `jobsim_sessions.go` wrote invalid free-text enum/result/sim_type values +
an over-long token, so its sessions INSERTed but were filtered out of every dashboard query — dead rows) and
patches `users.go` to real names/avatars/org-domain emails (no more "User N"). The skill node-ids are drawn
from the **real replayed public `skiller` taxonomy** (role-coherent via `TaxonomyRefs`, never fabricated), and a
**seed-side closure gene** (`datadna measure-closure`) proves zero dangling skill refs after seeding. Every
chain table is `PerStackIsolated`, so the same zero-pollution posture holds.

**Full reference: [`demo/stories-spec.md`](demo/stories-spec.md)** — the 7-table chain, the DB-enforced vs
inserted-but-invisible constraint landmines, the `user_level` (claimed side) requirement, and the
declare-a-hero blueprint shape.

**The Stories & Heroes model (v1.9 M35).** One declarative **`stories[]`** blueprint
(`presets/stories.seed.yaml`) supersedes the org-centric single-org `stack.seed.yaml` for a believable demo
world: it seeds **multiple orgs**, each with a **thriving / struggling / manager hero trio** at
vantage-appropriate fidelity. Per-story `OrgID` is threaded through every seeder (all orgs in one stack's
per-stack Postgres, scoped by `organization_id` — the platform's real multi-tenancy); the **first** story keeps
the Clerkenstein default org so a single-identity demo login lands in it (multi-identity seat-switch is M37).
Heroes carry a `vantage` (`end-user | manager` — a manager seeds no chain of her own, she reads the org
aggregates her employees populate) and a `trajectory` (`thriving` = dense/high/rising/under-claim vs
`struggling` = sparse/low/flat/over-claim — the stark gap). Every member also gets a **real replayed job role**
(`memberships.job_role_id`) + a **ramped `joined_at`**, so the trio sits in a believable org. Supporting-member
**names are deterministically deduplicated** (v1.9 M35 harden): each name is unique within its org and the
**hero names are reserved**, so the name bank never hands a supporting member a hero's name or repeats a name
in the same population (a hero-free legacy/dev-min population is byte-identical — the dedup only re-rolls on an
actual collision). The closure gene
is org-agnostic, so `datadna measure-closure` proves 0 dangling refs **across all orgs**. The vertical-slice
preset is still `presets/stories-maya.seed.yaml` (one hero); the presenter cockpit + the Clerkenstein
multi-identity seat-switch are M37–M38.

## Status

M7a delivers the framework + the isolation guard + the reference seeders (`org`, `users`, `identity`),
**62 unit tests** (green, `-race` clean), verified schema-correct against a live stack. M7b adds the data-DNA
schema-conformance/drift gate; M7c the full seeder fleet + backdated activity + presets. **v1.3 M13** adds the
**`dev-min` preset** + wires it as the **dev auto-seed on `dev-stack up`** (the second n=0-dev guard lives in the
bring-up's set-dressing pass) — see [The shipped presets](#the-shipped-presets-stack-seedingpresets).
