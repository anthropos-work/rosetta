# Demo Environments — the family index

**Purpose.** Stand up a **disposable, isolated, Clerk-free, realistically-populated** copy of the Anthropos
platform — for sales demos, screenshots, QA, or a clean-room — *alongside* the dev stack and **without
touching any read-only platform repo**. This family is the entry point; the mechanism guides + recipes it
indexes are the depth.

**When to use.** You want a demo world a stakeholder can log into and click around (a populated org with months
of activity), reproducibly, on offset ports, killable cleanly — and you must never pollute production or the
dev stack. If you just need the *dev* environment, see `../setup_guide.md` / `../run_guide.md`. If you need
*staging*, see the `../staging-*` family.

## The end-to-end flow (4 steps, ~minutes)

```
/demo-up N      →  bring up demo-N (Clerkenstein-wired, offset ports, isolated data)   [corpus/ops/rosetta_demo.md]
/demo-seed N    →  backfill it with a believable data world (a preset or stack.seed.yaml) [corpus/ops/seeding-spec.md]
  …use it…      →  browser-login as user_clerkenstein → land in a populated org (200)    [recipe-browser-login.md]
/demo-down N    →  tear it all down, dev stack untouched                                 [corpus/ops/rosetta_demo.md]
```

## Index

**Mechanism guides (the "how it works"):**
- [`../rosetta_demo.md`](../rosetta_demo.md) — the **lifecycle**: bring-up, the port-offset scheme, the
  Clerkenstein injection, the per-stack project/data isolation, the resource budget, teardown. (M3)
- [`../seeding-spec.md`](../seeding-spec.md) — the **seeding** reference: the `stack.seed.yaml` blueprint, the
  dependency-DAG, the **production-isolation boundary**, the casbin subtleties, the data-DNA. (M7a/b)
- [`../../architecture/alignment_testing.md`](../../architecture/alignment_testing.md) § "The data dimension" —
  the **data-DNA**: how a seeder's output is conformance-gated against the platform schema, and drift-detected.

**Use-case recipes (the "build a demo for X"):**
- [`recipe-enterprise-onboarding.md`](recipe-enterprise-onboarding.md) — a populated enterprise org (admin +
  hundreds of members), end to end.
- [`recipe-skill-progression.md`](recipe-skill-progression.md) — months of believable skill-progression
  activity (backdated job-sim + skill-path sessions).
- [`recipe-browser-login.md`](recipe-browser-login.md) — the **interactive** demo: the `@clerk/express` /
  orgclient cert-redirect + the browser-login walk-through, log in → land in a seeded org.

**Curated seed presets** (instances of `stack.seed.yaml`, validated to seed):
`rosetta-extensions/stack-seeding/presets/` — `small-200` (quick) · `mid-500` (the default "looks real") ·
`large-1k` (scale). Run via `/demo-seed N --preset mid-500`.

**Skills:** `/demo-up` · `/demo-seed` · `/demo-status` · `/demo-down` (see the root `CLAUDE.md` skills table).

## Hard constraints (always true)
- **Zero platform-repo change.** All demo tooling lives in the gitignored `anthropos-demo/rosetta-extensions/`
  (the demo-stack overlay + the seeder + Clerkenstein). The platform clones are consumed as-is.
- **Production-safe.** The seeder's isolation guard makes it *structurally impossible* for a non-prod stack to
  write a shared/prod store (Directus, the prod S3-public bucket, live Clerk, Customer.io/Brevo, AI APIs); it
  seeds only the per-stack Postgres/Redis, and proves it with an audit log. See `../seeding-spec.md`.
- **Isolated.** Every op is `-p demo-N`-scoped on offset ports with its own data; the dev stack is never touched.
