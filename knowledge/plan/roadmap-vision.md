# Roadmap — Vision

Future versions and proposals, not yet in active development. The active version lives in
[`roadmap.md`](roadmap.md). When a version starts development, its section moves from here into
`roadmap.md` and a `release/{version}` branch is cut.

> **v1.0 "body double"** (Clerkenstein) was promoted to [`roadmap.md`](roadmap.md) on 2026-06-02.
> What remains here is the next version, **v1.1 "show floor"**, which builds the demo environments on
> the Clerkenstein foundation. It promotes into `roadmap.md` when v1.0 closes.

---

## v1.1 — "show floor" (Disposable, richly-seeded demo stacks on demand)

**Depends on:** v1.0 "body double" shipped (Clerkenstein injected, browser login working without real
Clerk hassle, drift-gated). **Audience expansion:** sales / CS / PM can spin up an isolated demo stack
in the morning, seed it to a specific use-case (an org of 1k users with months of activity), demo it,
and kill it at night — no Clerk friction, no shared-staging contention.

### M3: Disposable multi-instance demo stacks (`anthropos-demo`)
**Shape:** `section` · **Complexity:** medium
**Goal:** Spin up `demo-1`, `demo-2`, … as isolated full stacks on one box, killable cleanly at night.
**Scope:**
  - In: the `anthropos-demo/` scratchpad pattern (mirrors `anthropos-dev/`); parameterized compose (`-p demo-N`, port offsets, per-stack volumes + `.env.demo-N`); Clerkenstein wired in by default; lifecycle skills (`/demo-up`, `/demo-down`, `/demo-status`) + cleanup/teardown.
  - Out: rich data seeding (M4); curated recipes (M5).
**Depends on:** M1 (stacks come up auth-working via Clerkenstein). **Builds on:** the staging compose/Makefile/repos.yml + skip-worktree patterns (reuse, don't rebuild).
**Open questions:** max concurrent demos + port-assignment scheme; shared repo clones vs per-demo clones (disk vs git hygiene); nightly auto-teardown vs manual kill.
**KB dependencies:** `corpus/ops/platform_repo.md`, `corpus/ops/staging-bringup.md`, `corpus/ops/run_guide.md`.
**Delivers → `corpus/ops/{demo-env guide}.md`:** the demo-stack lifecycle ops guide (net-new).

### M4: Declarative data seeding
**Shape:** `section` · **Complexity:** large
**Goal:** Describe a target state in one config and backfill the stack to match.
**Scope:**
  - In: a `demo.seed.yaml` schema (org size, role mix, content sources, activity span); a seeder orchestrating multi-store inserts in dependency order (org → users → memberships → Sentinel casbin → taxonomy → content → time-distributed activity), **reusing** `app/cmd/bootstrap-{user,org}`, `skiller/cmd/import{Skills,JobRole}`, `cms/cmd/jobsim`; deterministic time-distributed activity enabling orgs of 200/500/1k+ users across 2/3/10 months.
  - Out (**hard line**): AI-generated rich transcripts/embeddings → M5 stretch or deferred. M4 ships deterministic **structural** data only.
**Depends on:** M1 (no Clerk API rate limit → seeding is pure DB inserts) + M3 (a stack to seed into).
**Open questions:** single declarative config driving all stores vs per-store seeders; vector-embedding strategy (snapshot vs recompute); Directus tenancy (shared content vs per-demo instance).
**KB dependencies:** `corpus/architecture/architecture_overview.md` (multi-tenancy), `corpus/ops/staging_from_dump.md`, `corpus/services/{backend,skiller,jobsimulation,skillpath,cms}.md`.
**Delivers → `corpus/ops/{seeding spec}.md`:** the declarative-seeding spec + config reference (net-new).

### M5: Demo corpus + use-case recipes + skill polish
**Shape:** `section` · **Complexity:** medium
**Goal:** The connective documentation + curated recipes that make demos repeatable and discoverable.
**Scope:**
  - In: the demo-env corpus section; end-to-end "build a demo for use-case X" recipes; 2–3 curated seed configs (200/500/1k orgs); the demo skill set polished + discoverable from `README.md`/`CLAUDE.md`; (stretch) AI-assisted content generation pulled forward from M4 if budget allows.
  - Out: nothing new — this is consolidation.
**Depends on:** M3 + M4.
**KB dependencies:** the M1–M4 delivered docs.
**Delivers → `corpus/services/` + `corpus/ops/`:** index + recipe docs tying the demo-env story together.

### Execution graph (v1.1)

```
v1.1 "show floor"
  M3 (multi-instance stacks) ──→ M4 (declarative seeding) ──→ M5 (corpus + recipes + skills)
     (M3 depends only on shipped M1)
```

**Parallelism:** mostly sequential — M4 seeds *into* an M3 stack; M5 documents what M3+M4 produce.

## Open decisions (resolve at v1.1 kickoff)
- **Demo seeding fidelity vs spin-up time** — lightweight curated seed vs prod-dump baseline vs hybrid.
- **External shareability** — Tailscale-only (like staging) vs public ingress for customer-facing demos (security posture).
- **AI content generation** — template-sampled vs LLM-generated transcripts; where it lands (M4 stretch / M5 / deferred).

## Codename alternatives (pre-ship, changeable)
- v1.0: "body double" (chosen) · "dead ringer" · "false face" · "stage double"
- v1.1: "show floor" (chosen) · "open house" · "set piece" · "dry run"

_Last updated: 2026-06-02 (v1.0 promoted to roadmap.md; v1.1 "show floor" staged here as the next version)._
