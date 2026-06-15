# Demo Environments — the family index

**Purpose.** Stand up a **disposable, isolated, Clerk-free, realistically-populated** copy of the Anthropos
platform — for sales demos, screenshots, QA, or a clean-room — *alongside* the dev stack and **without
touching any read-only platform repo**. This family is the entry point; the mechanism guides + recipes it
indexes are the depth.

**When to use.** You want a demo world a stakeholder can log into and click around (a populated org with months
of activity), reproducibly, on offset ports, killable cleanly — and you must never pollute production or the
dev stack. If you just need the *dev* environment, see `../setup_guide.md` / `../run_guide.md` (driven by
`/dev-up`). If you need *staging*, see the `../staging-*` family.

> **Dev is a peer (v1.3).** Every set-dressing + seeding recipe below works on a **`dev-N`** stack exactly as it
> does on a `demo-N` — the `/stack-snapshot` / `/stack-seed` ops take `dev-N|demo-N` interchangeably, and a
> `/dev-up N` bring-up already set-dresses + light-seeds itself by default. Where a recipe says `demo-N`, read
> `dev-N|demo-N`. The one exception is the **`N=0` main dev stack**, which the auto-set-dress + `--reset`
> guards protect (see [`../safety.md`](../safety.md) §2.5).

## The end-to-end flow (~minutes)

```
/demo-up N        →  AUTO ensure-clones (bootstrap stack-demo's OWN clone set: clone platform from GitHub  [v1.8 M26 → rosetta_demo.md]
                     + make init the peer repos; seed the shared .env copy-if-present) — self-contained
                     bring up demo-N (Clerkenstein-wired, offset ports, isolated data, built from stack-demo) [corpus/ops/rosetta_demo.md]
                     AUTO secret-provision (values-blind, per-repo .env from .agentspace/secrets) [v1.6 M30 → secrets-spec.md]
                     AUTO set-dress (cache-first snapshot replay → small-200 seed, default-on, non-fatal) [v1.3b M20]
  …use it…        →  browser-login as user_clerkenstein → land in a populated org (200)    [recipe-browser-login.md]
/demo-down N      →  tear it all down, dev stack untouched                                 [corpus/ops/rosetta_demo.md]
```

> **A demo builds from its OWN clone set (v1.8 "understudy" M26).** `/demo-up` first runs `ensure-clones.sh`:
> it bootstrap-clones `stack-demo/platform` from GitHub over SSH + `make init`s every `repos.yml` repo as a
> sibling into `stack-demo/`, so **all** images build from `stack-demo` (a box with only `stack-demo/` — no
> `stack-dev/` — can bring a demo up end-to-end). The sole sanctioned `stack-dev` read is seeding the shared
> `platform/.env` copy-if-present (non-fatal if absent — M30 provisions the real one); the build SOURCE never
> falls back to `stack-dev`. Dev-image reuse is OFF by default (`DEMO_REUSE_DEV_IMAGES=1` opts back in).

**`/demo-up` now auto-set-dresses by default (v1.3b M20) — the dev↔demo convergence.** Just like `/dev-up` since
v1.3, a `/demo-up` bring-up chains the **same** set-dress pass at its tail: a cache-first **snapshot replay**
(the real **taxonomy** catalog; and — **for a demo, local content is default-on** since v1.5 M22/M23 — a
per-stack **Directus** booted + cut over so the **content** surface serves locally too, the stack
content-self-contained; a demo opted out with `DEMO_NO_LOCAL_CONTENT=1` falls back to reading content live from
prod — see the known-state note below) → a **`small-200` light seed**
(a populated org you can log into). So a bare `/demo-up N` already lands you in a real-catalog, log-in-able world
— no separate skill calls required. The pass is
**default-on + non-fatal** (a cold cache warns and still seeds; `DEMO_NO_SETDRESS=1` skips it for a bare
structural bring-up). You can still drive the steps **manually** for finer control — `/stack-snapshot N` (replay)
+ `/stack-seed N` (a different preset / a custom `stack.seed.yaml`) — they accept `demo-N` or `dev-N` interchangeably.

**The snapshot step is what makes the world *set-dressed* (v1.2).** A replay stamps the real **public** reference
library — the ~60K-skill taxonomy + the global simulation / skill-path content templates — into the stack BEFORE
the seed, so the catalog view shows real skills and the seeded sessions link to real templates (not free
placeholder ids). It's a **stack-global reference replay**, independent of which org you then seed; it's
**optional** (a structural-only world still logs in — the seeder degrades gracefully), and almost always a
**cache-hit** (zero prod read — the snapshot is captured once per release, then replayed by every stack).
See [`recipe-snapshot-world.md`](recipe-snapshot-world.md) for the full capture→replay→set-dressed walk-through.

> **Fresh box, empty cache?** The replay is a cache-hit only once the cache has been filled by a one-time
> `capture` — and a fresh machine with no safe `--dsn` can't replay the *real* catalog yet (the auto-set-dress
> warns + degrades to an empty catalog, then seeds). The sanctioned way to fill the cache once per release —
> and why the wired `postgres` MCP is **not** a capture source — is [`../snapshot-cold-start.md`](../snapshot-cold-start.md).

## Index

**Mechanism guides (the "how it works"):**
- [`../safety.md`](../safety.md) — the **safety contract**: the consolidated read-side (tenant-data firewall,
  public predicates, read-only capture) + write-side (3-layer isolation guard, never-write-prod, n=0 guards,
  audit-proven zero pollution) statement. **Read this first** if you care *why* it's safe. (v1.3 M15)
- [`../rosetta_demo.md`](../rosetta_demo.md) — the **lifecycle**: bring-up, the port-offset scheme, the
  Clerkenstein injection, the per-stack project/data isolation, the resource budget, teardown. (M3)
- [`../seeding-spec.md`](../seeding-spec.md) — the **seeding** reference: the `stack.seed.yaml` blueprint, the
  dependency-DAG, the **production-isolation boundary**, the casbin subtleties, the data-DNA. (M7a/b)
- [`../snapshot-spec.md`](../snapshot-spec.md) — the **snapshot** reference: how the real **public** taxonomy +
  Directus content library is captured once from prod safely (the read-side **tenant-data firewall**), cached
  outside git, and replayed per-stack — measured-faithful by the snapshot-fidelity data-DNA. (M9a/M9b/M10)
- [`../snapshot-cold-start.md`](../snapshot-cold-start.md) — the **cold-start** runbook: filling the snapshot
  cache once per release on a fresh box (the sanctioned DSN-export / dump-restore path), why the wired `postgres`
  MCP is **not** a capture source, and how it slots into the auto-set-dress bring-up. (v1.3b M20)
- [`../idempotency.md`](../idempotency.md) — the **re-run safety** contract: what happens when you run
  migrate / snapshot-replay / seed a *second* time — each is now safe-and-idempotent or fails loudly, never
  silently doubles or aborts mid-surface (the replay TRUNCATE-then-reload, the idempotent seed COPY + casbin
  guard, the `--reset` fix, the `set -e` first-run-race hardening). (v1.3b M17)
- [`../verification.md`](../verification.md) — the **verification** net: every bring-up ends with a scoped,
  NON-FATAL verify on the stack's OWN offset ports (the cheap-win `/api/health` + `casbin_rules > 0` silent-403
  catcher, then the full probe set), so "UP" means *verified-working* — offset/scope-aware, never blocks a good
  stack. (v1.3b M18)
- [`frontend-tier.md`](frontend-tier.md) — the **UI tier**: how `/demo-up` brings up next-web-app +
  studio-desk (per-demo cached Docker image from the **unmodified** Dockerfile, offset ports, minted-pk +
  offset-URL baked) + ant-academy natively (Clerk-free), the 12 GB Docker-VM prereq + non-fatal pre-flight,
  the honest "one ~3-min cached build per new demo-N" residual, and the `--no-ui` escape. (v1.3b M19)
- [`../../architecture/alignment_testing.md`](../../architecture/alignment_testing.md) § "The data dimension" —
  the **data-DNA**: how a seeder's output is conformance-gated against the platform schema, and drift-detected.
  With snapshots, coverage now reads **100%** (both formerly-`waived` surfaces promoted to `snapshot-seeded`).

**Use-case recipes (the "build a demo for X"):**
- [`recipe-enterprise-onboarding.md`](recipe-enterprise-onboarding.md) — a populated enterprise org (admin +
  hundreds of members), end to end — now **set-dressed** (real catalog + content behind the seeded org).
- [`recipe-skill-progression.md`](recipe-skill-progression.md) — months of believable skill-progression
  activity (backdated job-sim + skill-path sessions) linked to the real template library.
- [`recipe-snapshot-world.md`](recipe-snapshot-world.md) — the **set-dressing** recipe: capture →
  replay the public taxonomy + content into a stack so the catalog + templates are real, not placeholder.
- [`recipe-browser-login.md`](recipe-browser-login.md) — the **interactive** demo: the `@clerk/express` /
  orgclient cert-redirect + the browser-login walk-through, log in → land in a seeded org.

**Curated seed presets** (instances of `stack.seed.yaml`, validated to seed):
`rosetta-extensions/stack-seeding/presets/` — `small-200` (quick — **the `/demo-up` auto-set-dress default**,
M20 #M20-D2) · `mid-500` (the default "looks real") · `large-1k` (scale). The auto-set-dress pass uses
`small-200` (a fuller world than dev's `dev-min`); override it with a manual `/stack-seed N --preset mid-500`
(or skip the auto pass with `DEMO_NO_SETDRESS=1` and seed by hand). The presets are **purely structural** (they describe an org, not the
platform's reference library); for a **set-dressed** world the catalog replay runs first (the auto pass does this;
manually it's `/stack-snapshot replay N`). Without a replay the seeder degrades gracefully (empty catalog, free
content refs).

> **Known state — a `--local-content` stack is content-self-contained (M22 boot + M23 cutover); a prod-read
> stack reads the public catalog live from prod.** The auto set-dress replays the **taxonomy** locally. For the
> public Directus content: since **M22** a stack brought up with **local content** (demo **default**; dev
> `--local-content`) **boots its own per-stack Directus** (a compose service serving the captured catalog — the
> M10 collection-schema gap was closed by M21 + executed by M22), and since **M23** the bring-up **cuts `cms`
> over** to it (`DIRECTUS_BASE_ADDR` → in-network `http://directus:8055`) so content is served locally (asset
> plane stays on prod public links → real images). A **prod-read** stack (`DEMO_NO_LOCAL_CONTENT=1`, or a plain
> dev bring-up) still has **no local Directus**: `cms` reads the public sims/skill-paths **live from prod**
> (`content.anthropos.work`) — a **demo does so ANONYMOUSLY** since fix16/fix17: the injected override strips the
> inherited prod `DIRECTUS_TOKEN` from every demo container (prod Directus serves the public predicate tokenless
> — verified 2026-06-11; live demo-1 audit: 0/16 carriers), so no prod credential rides in a demo. The read is
> public-only + safe, but on the prod-read path it means a stack isn't fully self-contained. **The
> `--local-content` path (the demo default) closes this:** M22 boots + verifies the local Directus and M23 cuts
> `cms` over + guarantees **referential closure** — the taxonomy capture is **full-public** (`organization_id IS
> NULL` — every public node), so a content ref can only dangle if it points at a *non-public* node. A measured
> cross-surface gene reports any such dangle; prod has exactly **one** (`K-AIFUNX-E658`, a public sim referencing
> a customer-scoped skill) — an operator-owned prod data fix, not a tooling gap. Detail:
> [`../snapshot-spec.md`](../snapshot-spec.md) § the per-stack Directus store fork +
> [`../directus-local.md`](../directus-local.md) § "The data-plane cutover (M23)".

**Skills:** `/demo-up` · `/stack-secrets` · `/stack-snapshot` · `/stack-seed` · `/stack-list` · `/demo-down`
(see the root `CLAUDE.md` skills table). `/stack-secrets` provisions the stack's per-repo `.env` from one secret
source — **values-blind** — and verifies coverage; `/demo-up` runs it as an auto-provision step (M30) so a fresh
demo is self-sourced from the curated secret dir. Mechanism: [`../secrets-spec.md`](../secrets-spec.md).

## Hard constraints (always true)
- **Zero platform-repo change.** All demo tooling lives in `rosetta-extensions` (the demo-stack overlay + the
  seeder + Clerkenstein), never scattered in the rosetta corpus and never authored ad-hoc inside a stack dir.
  It is authored + tested + tagged in the authoring copy at `.agentspace/rosetta-extensions/`, and the demo
  stack consumes a pinned-tag clone at `stack-demo/rosetta-extensions @ <tag>`. The platform clones are
  consumed as-is.
- **Production-safe.** The seeder's isolation guard makes it *structurally impossible* for a non-prod stack to
  write a shared/prod store (Directus, the prod S3-public bucket, live Clerk, Customer.io/Brevo, AI APIs); it
  seeds only the per-stack Postgres/Redis, and proves it with an audit log. Snapshot **capture** is read-only +
  public-only (the tenant-data firewall). The full contract — both halves — is [`../safety.md`](../safety.md)
  (write-side detail in `../seeding-spec.md`, read-side in `../snapshot-spec.md`).
- **Isolated.** Every op is `-p demo-N`-scoped on offset ports with its own data; the dev stack is never touched.
