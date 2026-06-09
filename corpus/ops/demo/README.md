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
/demo-up N        →  bring up demo-N (Clerkenstein-wired, offset ports, isolated data)   [corpus/ops/rosetta_demo.md]
/stack-snapshot N  →  replay the real public catalog + content into the stack (set-dressing) [corpus/ops/snapshot-spec.md]
/stack-seed N      →  backfill it with a believable data world (a preset or stack.seed.yaml) [corpus/ops/seeding-spec.md]
  …use it…        →  browser-login as user_clerkenstein → land in a populated org (200)    [recipe-browser-login.md]
/demo-down N      →  tear it all down, dev stack untouched                                 [corpus/ops/rosetta_demo.md]
```

**The snapshot step is what makes the world *set-dressed* (v1.2).** `/stack-snapshot replay N` stamps the real
**public** reference library — the ~60K-skill taxonomy + the global simulation / skill-path content templates —
into the stack BEFORE you seed, so the catalog view shows real skills and the seeded sessions link to real
templates (not free placeholder ids). It's a **stack-global reference replay**, independent of which org you then
seed; it's **optional** (skip it for a quick structural-only world — the seeder degrades gracefully), and almost
always a **cache-hit** (zero prod read — the snapshot is captured once per release, then replayed by every stack).
See [`recipe-snapshot-world.md`](recipe-snapshot-world.md) for the full capture→replay→set-dressed walk-through.

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
- [`../idempotency.md`](../idempotency.md) — the **re-run safety** contract: what happens when you run
  migrate / snapshot-replay / seed a *second* time — each is now safe-and-idempotent or fails loudly, never
  silently doubles or aborts mid-surface (the replay TRUNCATE-then-reload, the idempotent seed COPY + casbin
  guard, the `--reset` fix, the `set -e` first-run-race hardening). (v1.3b M17)
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
`rosetta-extensions/stack-seeding/presets/` — `small-200` (quick) · `mid-500` (the default "looks real") ·
`large-1k` (scale). Run via `/stack-seed N --preset mid-500`. The presets are **purely structural** (they describe
an org, not the platform's reference library); for a **set-dressed** world `/stack-snapshot replay N` first (each
preset's header documents the prerequisite). Without a replay the seeder degrades gracefully (empty catalog, free
content refs).

**Skills:** `/demo-up` · `/stack-snapshot` · `/stack-seed` · `/stack-list` · `/demo-down` (see the root
`CLAUDE.md` skills table).

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
