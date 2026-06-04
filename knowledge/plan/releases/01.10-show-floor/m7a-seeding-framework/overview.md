---
milestone: M7a
slug: seeding-framework
version: v1.1 "show floor"
milestone_shape: section
status: done
created: 2026-06-04
completed: 2026-06-04
last_updated: 2026-06-04
delivers: rosetta-extensions/stack-seeding/ (framework + safety guard) + corpus/ops/seeding-spec.md
---

# M7a — Seeding framework + production-isolation safety

## Goal
Stand up the **foundation** every seeder plugs into — a declarative blueprint, a modular seeder contract, a
dependency-DAG orchestrator, a **fast** insert path, and (the load-bearing part) a **production-isolation guard**
that makes it *impossible* for a non-prod seeding run to write to a shared/prod store — plus a **minimum
end-to-end proof** (org + the real `user_clerkenstein` identity + casbin → a browser login returns 200). After
M7a a demo is *usable and provably safe*, even before the data-DNA discipline (M7b) and the full seeder fleet
(M7c) land.

## Why a section milestone (buildable, grounded)
The 2026-06-04 redesign research enumerated the seedable surfaces, the exact shared/prod-pollution vectors, and
the perf bottleneck — so the `In:` list is writable now. The skeleton rests on working references: `make
bootstrap-dev` (the org→users→memberships→casbin pipeline), `bootstrap.JoinOrg` (the per-user primitive),
`seed-verified-skill` (time-distributed activity). The weight is in the safety guard + the framework shape, not
in an uncertain path.

## The isolation boundary (the central safety contract — research-grounded 2026-06-04)
The only true cross-stack / prod-pollution vectors are a small, fixed set of **external/shared services**;
everything in the per-stack Postgres/Redis is inherently isolated (each stack has its own container). The guard
encodes exactly this:

| MUST route around (shared / prod) | How the guard handles it |
|---|---|
| **Directus** — one global instance (`content.anthropos.work`), visible on prod | **Block writes**; content comes from a snapshot replayed into the per-stack store (per-stack Directus would be a platform change → out of scope) |
| **S3 *public* bucket** — hardcoded to the **production** bucket in compose | Override `STORAGE_S3_PUBLIC_BUCKET` → empty (local `/tmp` fallback) before any storage RPC |
| **Live Clerk** (shared dev app) | Auto-route to **Clerkenstein** (already wired for demo stacks) — never the real Clerk API |
| **Customer.io / Brevo / AI APIs** | OFF by default; the guard keeps them off / quota-gated during seeding |
| *(safe: all Postgres schemas, Redis, S3-private→`/tmp`, pgvector, Sentinel/casbin — per-stack-isolated)* | seed freely |

**Correction baked into the guard:** "legacy skillpaths visible across stacks" traces to the *shared Clerk
identity*, not a shared skillpath DB (skillpath data is per-stack Postgres). Clerkenstein isolates identity per
stack, so this is already neutralised — the seeder must just declare each surface's isolation class correctly.

## Scope
### In
- **The `stack.seed.yaml` blueprint schema**: org (name); size (N users); role_mix (admin/member/candidate +
  a small fixed admin set); tier_mix (free/premium); content-pack reference; activity span (e.g. "3 months",
  session pass/fail distribution). **Curator-annotatable per seeder** (each seeder exposes a config block the
  demo curator tweaks to reach a scenario).
- **The modular seeder contract / registry**: every seeder is a self-contained module declaring `{surface,
  depends-on, isolation-class, primitive}`. The orchestrator discovers + sequences them; M7c populates the
  registry, M7a defines the interface + ships ~2 reference seeders (org, users) to prove it.
- **The dependency-DAG orchestrator**: a real DAG (not a hardcoded linear order) so independent seeders run in
  **parallel**; topological order: migrate → Sentinel policy → org → users → memberships+casbin+feature
  (`JoinOrg`) → taxonomy/content (snapshot) → time-distributed activity.
- **The production-isolation guard (CRITICAL)**: the enumerated shared/external surface list above; a hard block
  on any shared-write from a non-prod stack (S3-public override, Clerk→Clerkenstein, Directus write → hard
  error); a **seeding audit log** (every write tagged `scenario_id` / `seeded_by` / `isolation_class`) that
  *proves* zero prod/shared pollution after a run.
- **The perf architecture** (spike-revised, M7a-D3): a **host Go module that connects DIRECTLY to the target
  stack's Postgres** (offset port) + **`COPY`/batched SQL** bulk-insert + a goroutine fan-out; direct SQL for
  side-effecting primitives (the Clerkenstein identity, Sentinel/casbin). *Not* ent-linking (blocked — `app/
  internal/bootstrap` is an `internal/` package, unimportable from a separate module without a platform edit) and
  *not* per-row CLI-shelling (the slow bottleneck). The drift risk of direct SQL is mitigated by M7b's data-DNA.
- **`--reset` / `--validate` / `--dry-run`** (folded in, M4-D1 hedge): per-stack DB reset; schema + semantic
  checks; ordered insert plan + per-store row counts + **the isolation audit preview** (which stores a run would
  touch) before execution.
- **The minimum end-to-end proof (M3-inherited must-have):** seed org + the real **`user_clerkenstein`** login
  identity (`demo@anthropos.test` / `org_clerkenstein`, role `admin`, matching `clerkfrontend.DefaultDemoUser()`)
  + casbin — honoring the **`casbin_rules`(plural, `init_policy.sql`) vs `casbin_rule`(singular, gorm)** gotcha —
  so a browser login (which mints exactly that identity) resolves to a *seeded* user and authorized routes return
  **200, not 403**.
- **Delivers** `corpus/ops/seeding-spec.md`: the blueprint reference + the dependency-DAG contract + **the
  isolation-safety boundary** (the canonical "what seeding may/may-not touch" doc).

### Out
- The **data-DNA** discipline (enumerate-as-genome + schema-conformance operators + drift detection) — **M7b**.
- The **full seeder fleet** + the backdated activity generator at scale + scenario composition — **M7c**.
- AI-generated rich content (transcripts, AI-scored narratives, freshly-computed embeddings) — the hard line;
  v1.2 stretch / deferred.
- Authoring the 60K-skill taxonomy or jobsim content — seeding **consumes a pre-embedded skiller snapshot**.
- The disposable-stack lifecycle (compose/ports/`/demo-up`) — M3. Recipes + presets — M8.
- Any modification to platform repos (primitives called as-is from the per-stack clones).

## Depends on
**M4** (the `rosetta-extensions` monorepo home) + **M3** (a Clerkenstein-wired stack to seed into) + v1.0 M1/M2
(Clerkenstein removed Clerk's API rate limit → user seeding is pure DB inserts). **Parallel with:** none (gates
M7b + M7c).

## Estimated complexity
**large** — the safety guard + the DAG orchestrator + the perf path + the proof. The riskiest part is the
isolation guard (must be airtight: a single un-guarded shared-write pollutes prod).

## Open questions (resolve during build)
- Directus content: snapshot-and-replay into the per-stack store vs hard-block-and-skip for the demo MVP.
- ~~Linking the platform ent client vs CLIs~~ — **RESOLVED (M7a-D3):** neither; direct Postgres `COPY`/SQL over
  the stack's offset port (`app/internal/bootstrap` is internal → unimportable; CLI-shelling is the slow path).
- Backdating fidelity: which `created_at`/timestamps are settable via direct SQL vs ent-Immutable/DB-default `now()`.
- Whether the audit log lives in the per-stack DB (a `seeding_audit` table) or a sidecar file.

## KB dependencies (read as contract)
- `corpus/architecture/architecture_overview.md` (multi-tenancy: `organization_id` on every table, 3-layer isolation)
- `corpus/architecture/security_compliance.md` (the isolation model + EU/data boundaries the guard must respect)
- `corpus/ops/staging_from_dump.md` (snapshot/restore — the pre-embedded skiller snapshot + the Directus replay)
- `corpus/services/{backend,skiller,jobsimulation,skillpath,cms,storage}.md` (the primitives + store schemas + the S3-public footgun)
- `corpus/ops/rosetta_demo.md` (M3 — the stack to seed into) · `rosetta-extensions/` `knowledge/` (the monorepo home)

## Delivers → `corpus/ops/seeding-spec.md` (net-new)
The blueprint schema + the dependency-DAG contract + **the production-isolation safety boundary** (the enumerated
shared-store list + the guard's routing rules + the audit-log contract).

## Exit (section)
All `In:` deliverables land: `stack.seed.yaml` parses + validates; the seeder registry + DAG orchestrator run the
2 reference seeders; **the isolation guard provably blocks every shared/prod write** (audit log clean) on a
non-prod stack; `--reset/--validate/--dry-run` work; and the **minimum proof passes — a browser login to a seeded
stack returns 200** — documented in `corpus/ops/seeding-spec.md`.
