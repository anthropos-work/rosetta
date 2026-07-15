# Declarative Stack Seeding — Spec

**The reference for `rosetta-extensions/stack-seeding/`** — how a demo or dev stack is backfilled with
structural data from one declarative blueprint, the dependency order, and (most importantly) the
**production-isolation boundary** that keeps a non-prod seeding run from polluting production.

> **The demo-patch mechanism is specified in [`demo/demopatch-spec.md`](demo/demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

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
`taxonomy` snapshot node into the DAG (the public skills-taxonomy catalog — the `public` schema since the v2.1 skiller→app merge, replayed out-of-band; `activity` orders behind
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
- **`taxonomy`** (M9b): `waived-m7c → snapshot-seeded-m9b` — the public skills-taxonomy catalog (`public` schema), fidelity-gated by five
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
from the **real replayed public taxonomy** (the `public`-schema skills/roles catalog; role-coherent via `TaxonomyRefs`, never fabricated), and a
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

**The Workforce dashboard surfaces (v1.9 M36).** M34/M35 made the *individual* profile believable; M36 makes
the org **Workforce-Intelligence dashboard** (REST `/api/workforce/*`, the org-admin view) believable. Six new
seeders + two fixes land the **spine** so every aggregate renders non-empty and distributed: **`membership_skills`**
(the mapped→verified funnel — mapped outnumbers verified per skill, by skill *name*), **`tags`+`membership_tags`**
(teams/business-units + a `mentor` tag, the slice dimension), **`organization_target_roles`+`user_target_roles`**
(the gap + two-way mobility), **`interview_extraction_results`** (the succession feeder, sized >20% to clear the
coverage gate), **`job_simulation_feedbacks`** (~2:1 positive), and **population `user_skill_evidences`** (the
org-scale claimed-vs-verified gap — a population mix of over/under-claimers). The two fixes: the **assignments
status-mix** (`due_date`s + `organization_assignment_sessions` so assignments bucket as
completed/overdue/in-progress, not all `not_started`) and the **skillpath completed share** (~30%, was ~1%). The
seed-side closure gene now also covers `membership_skills.skill_id` (four surfaces). Every dashboard seeder is
**no-fabrication and degrades, not crashes** (v1.9 M36 harden): the share predicates short-circuit cleanly at the
0/1 bounds (an empty or saturated cohort), the named-skill resolver yields empty pools on a missing taxonomy
(writing nothing rather than inventing a skill) and pairs a malformed read (mismatched parallel array_agg columns)
to empty lineage rather than panicking, and each multi-table seeder (`tags`, `target_roles`, `assignments`)
returns the **partial total** of rows already written when a later FK-ordered COPY fails, wrapping the error with
the failing table so an operator sees *which* surface broke. Full reference:
[`demo/stories-spec.md` § The Workforce dashboard surfaces (M36)](demo/stories-spec.md#the-workforce-dashboard-surfaces-m36).

**The profile-depth layer (v1.10 M41).** v1.9 made a hero's verified-skill **spine** render; v1.10 M41 fills the
**depth** a presenter sees after "Login as": a believable **work history + education timeline** and a **deep,
role-aligned skill set with a wide claimed-vs-verified gap**. A new **`ProfileSeeder`** (surface `"profiles"`)
writes, per **end-user** hero (managers skipped — no personal timeline): a **`companies`** row per distinct
employer + a **3-job role progression** (`user_experiences`, current role `to`=NULL) + a **degree**
(`user_educations`), backdated within/before the activity span, titles from the resolved `jobRoleRefs`, `skills`
json a role-coherent slice of real public skill names — both timeline tables were **0 rows DB-wide** (net-new
write surface). It **bumps the preset `verified:` knob `8 → ~30`** (→ ~90 `user_skills` + ~30 evidences on the
verified side) **and** seeds a **~60-skill claimed-but-unverified tail** (`user_skills` `is_verified=false`, no
`job_simulation_id`, tied to the seeded experiences via `user_skill_experience`/`user_skill_education` to satisfy
the `user_skills_check_foreign_keys` CHECK; `user_skill_evidences` `anthropos_level` NULL, `user_level` set) — so
the profile "overall" reads **≈ 90 = ~30 verified + ~60 claimed**, **widening** the visible gap. Live-schema
landmines drove the design: `user_experiences.company` is `uuid NOT NULL` FK→`companies`, `from`/`to` are DATE
with a `from<=to OR to IS NULL` CHECK (the current role leaves `to` NULL — open-ended), `location_type` is the
lowercase ent enum `inoffice|hybrid|fullremote`, `skills` is json. The claimed-tail evidence UPSERT is a
**separate** SQL from the verified one, guarded `ON CONFLICT … WHERE is_verified = false` so it **never clobbers
a verified row** on a (skill,user) collision (the verified side always wins; the tail draws skills distinct from
the verified set, so the guard is a re-run/safety net). No-fabrication + closure preserved (skill refs from the
replayed taxonomy; empty pool → timeline still writes, tail skipped); every table `PerStackIsolated`. Full
reference:
[`demo/stories-spec.md` § The profile-depth layer (v1.10 M41)](demo/stories-spec.md#the-profile-depth-layer-v110-method-acting-m41).

**Profile completeness — the whole roster (v1.10 M44).** M41 gave the END-USER heroes a deep profile; **M44
makes the WHOLE roster — managers AND bulk members — read as real people**, DATA-DENSITY only (no UI change).
Four extensions, all in `stack-seeding/seeders/`:
- **(§A) Trajectory-aware self-rating.** `PersonaSeeder` writes `user_skill_evidences.user_level` only for a
  **self-rated** hero (the new `Persona.EffectiveSelfRated()`: struggling = false → `user_level` NULL, the
  claimed side absent; everyone else = true). A thriving hero shows a **completed** self-assessment; a
  struggling hero "hasn't done the initial self-rating" — her 2-3 verified skills still render (the chart
  needs the verified side, untouched), but the claimed gap reads incomplete.
- **(§B) Certifications + Projects.** Two NEW seeders — **`CertificatesSeeder`** (surface `"certificates"` →
  `public.user_certifications`, 2-3 per end-user hero) + **`ProjectsSeeder`** (surface `"projects"` →
  `public.user_projects`, 3-4 per end-user hero) — fill the two profile sections that were **0 rows DB-wide**.
  Both surface as top-level `TimelineGroupedItems.certifications`/`.projects` on `/profile`. **LIVE-SCHEMA
  CORRECTED** (the overview's guesses were wrong): the table is **`user_certifications`** (NOT
  `user_certificates`; cert *name* col is `certification`, NO `created_at`, NO `organization_id`); projects use
  `title` (NOT `project_name`), `to` (NOT `end_date`), NO `organization_id`. Role-coherent banks, idempotent
  COPY, closure-clean skills, managers skipped here (the §C path owns manager profiles).
- **(§C) Manager personal data.** The pre-M44 `IsManager` skips in `PersonaSeeder` + `ProfileSeeder` are
  removed; a manager now gets a **modest** personal profile — a FLAT 3-8 verified skills (L1-L2 band, self-rated,
  no growth arc) + a manager-track timeline (a leadership ladder "Engineering/Sales Manager" ← "Team Lead" ←
  "Senior X", 3 experiences + 1 education) + a SMALL claimed tail (8, not the deep ~60) — so her OWN `/profile`
  is populated, not empty. The claimed-tail offset now uses the EXACT verified count (`trajectoryVerifiedCount`)
  so the tail stays distinct from the verified set.
- **(§D) Bulk-member depth.** EVERY non-hero population member gets a shallow career — `ProfileSeeder` runs a
  second pass over the non-hero slots (3 short-tenure experiences + 1 education + a flat <=6-skill claimed tail;
  role mirrors the `UsersSeeder` supporting-member draw) — so `/enterprise/members` reads as a roster of real
  people. The avatar half was already satisfied (`photoAvatarDataURI` runs for every member since M42e P4).
- Everything stays `PerStackIsolated` + closure-GREEN. Full reference + the per-vantage rubric:
  [`demo/profile-completeness-spec.md`](demo/profile-completeness-spec.md).

## Status

M7a delivers the framework + the isolation guard + the reference seeders (`org`, `users`, `identity`),
**62 unit tests** (green, `-race` clean), verified schema-correct against a live stack. M7b adds the data-DNA
schema-conformance/drift gate; M7c the full seeder fleet + backdated activity + presets. **v1.3 M13** adds the
**`dev-min` preset** + wires it as the **dev auto-seed on `dev-stack up`** (the second n=0-dev guard lives in the
bring-up's set-dressing pass) — see [The shipped presets](#the-shipped-presets-stack-seedingpresets).
**v1.9 M34** adds the verified-skill chain (the `PersonaSeeder` + `TaxonomyRefs` + the G14 session fix + the
seed-side closure gene); **M35** the multi-org Stories & Heroes model; **M36** the six Workforce-dashboard
surfaces + the two fixes above (the funnel/teams/target-roles/succession/feedback/org-scale-gap spine),
proven end-to-end by an opt-in live-stack integration test (`-tags integration`) that seeds the full fleet and
asserts every dashboard aggregate resolves. The closure gene spans four skill-ref surfaces.
**v1.10 "method acting" M39** adds the profile-identity layer (roster org-name thread + `user_basic_info` role
backfill + offline real-face avatars); **M41** the **profile-depth layer** (the `ProfileSeeder` work-history +
education timeline + the verified depth bump `8 → ~30` + the ~60-skill claimed-but-unverified tail that widens
the gap) — 9 new unit tests, full suite green `-race`, every emitted row dry-insert-validated against the live
demo-3 schema, `go.mod`/`go.sum` byte-identical. Code-of-record: `rosetta-extensions` @ tag `method-acting-m41`.
**M44** adds **profile completeness** — the whole roster (the §A trajectory-aware self-rating, the §B
`CertificatesSeeder` + `ProjectsSeeder` surfaces, the §C manager personal-data unskip, the §D bulk-member
shallow career) — all DATA-DENSITY, zero platform edits, every surface `PerStackIsolated` + closure-GREEN,
schema-corrected against the live demo-3 DB (`user_certifications`/`user_projects`). New unit tests across the
four sections, full suite green `-race`. Render-verified on a live demo (the §D fix1 avatar-column correction:
`/enterprise/members` reads `memberships.picture_url`, not `users.picture`) + a 3-pass hardening sweep (17 added
tests, seeders-pkg stmt coverage 96.5%→97.5%). Code-of-record: `rosetta-extensions` @ tag
`method-acting-m44-profile-completeness-fix2`.
**v1.10b "fit-up" M50** completes the Talent-tab fill on the clean re-grounded demo: a NEW
**`MemberLanguagesSeeder`** (the ISO-639-1 `world_languages` catalog + per-member `user_languages`, which the DB
AFTER-INSERT trigger fans out to `membership_languages` — the manager Talent-tab "Languages spoken" chart), the
**`CertificatesSeeder` member-coverage extension** (hero-only → ~45% role-coherent members, so the "Certifications"
chart reads as a credentialed workforce), and the **`UsersSeeder` member-field backfill** (`memberships.joined_at`
/ `location` / `last_activity_date` for the `/enterprise/members` columns, NULL-only idempotent guard). All
DATA-DENSITY, `PerStackIsolated` + closure-GREEN (the ISO catalog is a published standard, not a fabricated
node-id), proven by the M42 coverage gate on the manifest STRENGTHENED to assert them (both vantages, warm). The
detail lives in [`profile-completeness-spec.md`](demo/profile-completeness-spec.md) §"the Talent-tab fill" + §"the
shared per-member uuid-space". Code-of-record: `rosetta-extensions` @ tag `fit-up-m50`.
**v1.10b "fit-up" M51** adds the **AI-readiness showcase org — a 3rd story** (org "Northwind Aviation", 200
members, hero trio Aria COMPLETED / Ben STARTED / Dana manager), with **four net-new AI-readiness seeders**
(**corrected v2.3 M219 — it was three**): **`OrgSettingsSeeder`** (the `organization_settings` `ai_readiness`
enablement gate-row — nothing wrote that table before), **`AIReadinessConfigSeeder`** (`ai_readiness_cycles` —
since M219 **BOTH a `closed` AND an `active` cycle** — + `ai_readiness_skills` with **real replayed-taxonomy
node-ids** + `ai_readiness_sims` + `ai_readiness_steps`), **`AIReadinessFunnelSeeder`** (199 frozen
`ai_readiness_snapshots` at 78.4% all-3-complete + `ai_readiness_user_step_progresses`), and — net-new at
**M219** — the **interview-aggregated-report seeder** (`jobsimulation.interview_aggregated_reports`, flushed by
the funnel seeder), **without which the manager's four interview-findings blocks render headings with no content**
(no seeder had ever written that table). DAG-ordered `config → funnel`.

> ⚠️ **M219 FALSIFIED M51's headline strategy claim.** This paragraph used to state that M51 shipped the
> closed-cycle / frozen-snapshot strategy *"after the active-signals path was falsified — the live-recompute never
> completes in the coverage-harness budget"*. **The live recompute completes in 2.09 s.** M219 measured it. The
> demo now seeds **both** cycles: the **active** one is what the manager dashboard resolves and live-recomputes
> (it is what fills Ben's funnel, Aria's full hero card, and the manager's `interview` / `diagnosis` / `sources`
> sub-sections — all of which were NULL/absent under closed-only), while the **closed** one is retained as cycle
> *history*. The `--reset` table list gains `jobsimulation.interview_aggregated_reports`.

Plus the **`app-aireadiness-snapshot-loadmembers`** app read-path demo-patch that bounds the frozen read's
whole-org `loadMembers` to the ~199 snapshot users (180 s → 19 ms; a pure, data-identical perf optimization). All
`PerStackIsolated` + closure-GREEN across all 3 orgs, proven by the **M42 manager-vantage** coverage gate `(0,0)`.
The detail lives in [`demo/stories-spec.md`](demo/stories-spec.md#the-ai-readiness-showcase-org--the-3rd-story-v110b-fit-up-m51)
+ the seeder contract [`../services/ai-readiness.md`](../services/ai-readiness.md). Code-of-record:
`rosetta-extensions` @ tag `fit-up-m51`.
**v1.10b "fit-up" M52** consolidates the whole seed+generation intent into **one auditable file**
(`presets/seed-generation-manifest.yaml`): the population (all 3 orgs + heroes), the **file-resident** mother
prompt (extracted from the Go const to `blueprint/prompts/default_batch_prompt.tmpl`), the batch config (the
MANDATORY `max_cost_usd` ceiling + concurrency + re-roll rules), and the snapshot sources — **cache +
generated data excluded**. It is a PROJECTION of the canonical presets (honesty-gated so it can't drift),
emitted by `stackseed --manifest-export`, and served by the presenter cockpit's **[Download seed manifest]**.
Full reference: [`demo/seed-manifest-spec.md`](demo/seed-manifest-spec.md). Code-of-record:
`rosetta-extensions` @ tag `fit-up-m52`.
