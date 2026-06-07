# Changelog

All notable user-facing changes to Project Rosetta. Format: [Keep a Changelog](https://keepachangelog.com/), semver-aware.

## [v1.3] "stack party" — 2026-06-07

dev + demo stacks become **first-class peers**: a dev stack gets the demo treatment (its own local Directus,
an auto-snapshot of the real reference data, a light default seed), one unified registry keeps dev and demo
from colliding on ports, and **one converged skill set** operates any stack — plus a single authoritative,
code-cited safety doc.

### Added
- **`/dev-up`** / **`/dev-down`** — the dev-stack lifecycle, mirroring `/demo-up` / `/demo-down`. `/dev-up`
  builds-or-starts the main dev stack (consolidating the former `setup-platform` + `start-platform`), and for
  an additional `dev-N` set-dresses it by default (local Directus + cache-first snapshot replay + a light
  `dev-min` seed). `/dev-down` tears a `dev-N` down and frees its registry slot. (M13/M14)
- **Dev stacks are now full-fidelity peers of demo for data** — a fresh `dev-N` is never empty: it gets a
  per-stack local Directus, an auto-snapshot replay of the real public taxonomy + content, and a light
  `dev-min` seed (~1 org + ~10 users), all default-on (escapes: `--no-snapshot`, `--no-setdress`). Capture is
  never run against prod from a dev stack (replay-only). (M13)
- **A unified stack registry + first-available-N allocation** — one shared N-pool spans dev *and* demo, so a
  bring-up always claims the lowest free slot and `dev-N`/`demo-N` can never collide on ports (e.g. building
  dev, demo, dev, demo, demo yields `dev-1, demo-2, dev-3, demo-4, demo-5`). `/stack-list` surfaces it. (M12)
- **`corpus/ops/safety.md`** — the authoritative, code-cited safety contract of the stack tooling: it **never
  reads private/customer data** (the tenant firewall + public predicates + bounded read-only capture) and
  **never touches production** (the 3-layer isolation guard + never-write shared Directus/prod-S3 + the audit-
  proven zero-pollution assertion). Every load-bearing claim is pinned to the source by a fail-closed drift
  guard. (M15)

### Changed
- **The stack-operation skills were hard-renamed to generic `stack-*` forms (no aliases)** — each accepts a
  `dev-N` or `demo-N` target: `/demo-status` → **`/stack-list`**, `/demo-seed` → **`/stack-seed`**,
  `/demo-snapshot` → **`/stack-snapshot`**, `/update-platform` → **`/stack-update`**. `/demo-up` / `/demo-down`
  stay as the demo lifecycle (now aligned with `/dev-up` / `/dev-down`). (M14)

### Removed
- The old skill names `/setup-platform`, `/start-platform`, `/update-platform`, `/demo-status`, `/demo-seed`,
  `/demo-snapshot` (and their skill dirs) — a clean break, no back-compat shims. Update any saved invocations
  to the converged names above. (M14)

### Supply chain
- No dependency changes; all deps permissive (MIT / BSD-3 / Apache-2.0 / ISC); **0 called third-party CVEs**.
  **Recommendation:** build with the **go1.25.11+** toolchain to clear 12 Go-stdlib (parsing/DoS-class)
  govulncheck findings (same class as v1.2). Lockfile: `knowledge/plan/releases/archive/01.30-stack-party/dependencies.lock`.

### Known limitations
- Demo/dev media still renders structure + file **references** (placeholder bytes); the actual S3 media blob
  **bytes** and a **cloud snapshot store** are now deferred to **v1.4** (DEF-M10-01 — gated on eu-west-1 S3-read
  access not wired here). *(This corrects the v1.2 changelog note below, which named v1.3 as the destination
  before the item was re-scoped to v1.4.)* AI-generated rich content (transcripts/embeddings) and external
  stack shareability are also v1.4.

## [v1.2] "set dressing" — 2026-06-07

The **snapshot mechanism**: *set-dress* a disposable demo stack with the **real public reference library** — the actual skills taxonomy and the Directus simulation/skill-path templates — so the catalog and the content behind seeded sessions are real, not placeholders. Everything is captured **read-only** from production; **customer data is never copied** (a tested tenant-data firewall).

### Added
- **`/demo-snapshot N`** — replay the captured public taxonomy + Directus content into a demo/dev stack (drives the `stacksnap` CLI: capture / replay / status). New demo flow: `/demo-up → /demo-snapshot → /demo-seed`. (M11)
- **`/db-query`** — read-only production DB investigation skill (the wired `postgres` MCP tool, or Tailscale + `~/.pgpass`), with the public-vs-customer data boundary documented. (M9a)
- **The `stack-snapshot` extension** — capture a public reference surface once from a safe source, manifest-cache it under `.agentspace`, replay per-stack — behind a firewall that hard-fails on any captured customer row. (M9a)
- **Real public skills taxonomy** in demo stacks (42.8K skills + roles + embeddings + translations; pgvector index rebuilt on replay). (M9b)
- **Real public Directus content** in demo stacks (published simulation/skill-path templates), served by a per-stack Directus; seeded sessions/assignments now resolve to real templates. (M10)
- A **snapshot-fidelity** alignment dimension and **100% data-DNA coverage** (the two formerly-`waived` surfaces promoted).
- New docs: `corpus/ops/snapshot-spec.md`, `corpus/ops/db-access.md`; the set-dressed `corpus/ops/demo/` recipe family (incl. `recipe-snapshot-world.md`).

### Changed
- Seeding spec: `taxonomy` + `content` promoted `waived` → `snapshot-seeded` (100% catalog coverage; nothing waived).
- Capture-source policy: default **ingest an existing prod `pg_dump`**, fallback a **safe throttled read-only primary read** (PostgreSQL MVCC = a read never blocks prod writes).

### Removed
- The planned offline `pg_dump`-**FILE** reader was dropped — restore-then-`--dsn` covers the need with no new capability or speed gain; the `--dump` flag is gone.

### Supply chain
- No dependency changes; all deps permissive (MIT / BSD-3 / Apache-2.0 / ISC); **0 third-party CVEs**. **Recommendation:** build with the **go1.25.11+** toolchain to clear 12 Go-stdlib (DoS/parsing-class) govulncheck findings.

### Known limitations
- Demo media renders structure + file references; the actual S3 blob **bytes** and a **cloud snapshot store** are deferred to v1.3 (gated on S3/AWS access not wired here).

## [v1.1] "show floor" — 2026-06-05

The platform-operations extension framework: spin up a **disposable, Clerk-free, realistically-populated** copy of the platform — for a demo, screenshots, or QA — alongside the dev stack, **without touching production or any read-only platform repo**.

### Added
- **Disposable demo stacks** (`/demo-up`, `/demo-down`, `/demo-status`) — bring up `demo-N` isolated on offset ports, Clerkenstein-wired (Clerk-free), with its own data; killable cleanly; the dev stack never touched. (M3)
- **Declarative, production-safe seeding** (`/demo-seed` + the `stackseed` tool) — backfill a stack from one `stack.seed.yaml` (or a curated preset): an org + 1,000 users + the real `user_clerkenstein` login identity + months of **backdated** job-sim / skill-path sessions, assignments, and activity — a believable world a stakeholder can log into (authorized routes return **200**). It connects **directly to the stack's Postgres** (`COPY`; ~0.7s for ~9,500 rows) behind a **3-layer production-isolation guard** that makes it *structurally impossible* for a non-prod run to write a shared/prod store (Directus, the prod S3-public bucket, live Clerk, marketing/AI SaaS), and proves zero pollution with an audit log. (M7a)
- **The data-DNA** — the alignment framework extended to a third dimension, **data**: the `datadna` CLI enumerates the seedable surfaces, **conformance-gates** each seeder's output against the platform's current schema, and **detects drift** when that schema moves (`measure` 100% / `diff` flags a changed column). (M7b)
- **The seeder fleet** — backdated-activity seeders for the believability core (job-sim + skill-path sessions, assignments, activity events), driven to a data-DNA coverage gate. (M7c)
- **`dev-stack`** — the same multi-instance tooling for isolated *dev* stacks (`dev-N`), real Clerk by default, optional Clerkenstein injection. (M6)
- **Demo-env corpus family** (`corpus/ops/demo/`) — a family index + 3 end-to-end recipes (enterprise onboarding · skill progression · interactive browser login) + 3 curated seed presets (small-200 / mid-500 / large-1k). (M8)
- **`@clerk/express` alignment gate wired into CI** — the v1.0 carry-forward; clerkenstein's CI now materializes the SDK + runs the gate (validated 9/9). (M8)

### Changed
- **Repo consolidation** — the standalone `clerkenstein` + `rosetta-demo` repos collapsed into one private `rosetta-extensions` monorepo (history preserved via `git subtree`); the old org repos were removed; `rosetta` thinned to documentation + the alignment framework + pointers. The reusable Clerk-mock injection layer (`stack-injection`) and the shared port-offset engine (`stack-core`) were extracted as sections. (M4/M5/M6)
- **Clerkenstein gained a 4th measured surface** — deployment/injection (`clerk-deploy-1`, 7/7): the disarmed `colony/authn` drop-in compiles against the platform's real `colony` and satisfies its contract. All four gates held 100%/100% throughout v1.1. (M3 extended)

### Known limitations
- Seeding ships **structural data only** by design. Two surfaces are **waived** to v1.2: the skill *taxonomy* (needs a pre-embedded skiller snapshot) and Directus *content* (the shared instance — snapshot-replay only). AI-generated rich content (transcripts/embeddings) is also out of scope. Data-DNA coverage is 100% over the 8 reachable surfaces; the waived surfaces are recorded as `waived-m7c` in the manifest.
- The deployment/injection alignment gate stays a **local** gate (it needs the platform's `colony` via a private token); the other three surfaces run in CI.

### Supply chain
- No new runtime deps beyond the Postgres driver (`jackc/pgx/v5 v5.9.2`) + `gopkg.in/yaml.v3` for the seeder. All deps permissive (BSD/MIT/Apache); lockfile at `knowledge/plan/releases/archive/01.10-show-floor/dependencies.lock`.

## [v1.0] "body double" — 2026-06-03

The first release under the developer-kit planning lifecycle. Rosetta gains a measurement discipline and its first product: a drop-in Clerk mock you can *prove* is faithful.

### Added
- **Alignment-testing framework** (`test/alignment/`) — a third test class that scores how faithfully a *mirror* engine reproduces a *source* library as a 0–100% number (overall + a separate critical gate). Ships `alignctl` (stdlib-only Go, runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; four equivalence operators; record/replay goldens; a `//go:build alignment` test class; and a toy reference that proves end-to-end divergence detection. Plus the `/align-dna` and `/align-run` skills and the canonical doc `corpus/architecture/alignment_testing.md`. (M0)
- **Clerkenstein** — a *measured* drop-in mock of the Clerk libraries the platform uses, so demos run Clerk-free with **zero platform-code change**. Verified at **100% / 100% on all three measured surfaces**: the Go SDK (`clerk-sdk-go/v2`, 22 genes — M1), the JS/FAPI browser surface (`@clerk/clerk-js`+`@clerk/nextjs`, 9 genes — M2), and the `@clerk/express` Node backend (9 genes, RS256/JWKS — M2c). Lives in its own repo; injected via build-time `go.mod replace` + a fake Clerk API server.
- **Drift detection** (M1b) — a `gate.sh`/`drift-check.sh` exit-code contract + a weekly CI alignment gate + a drift regression harness, so a Clerk version bump becomes a flagged, mechanical event instead of a silent break.
- **Webhook + browser-session coherence** (M2) — a fake FAPI (browser login via a minted publishable key, config-only), a fake BAPI that disarms the platform's networked org client, and an svix-signed webhook injector.

### Changed
- The `clerkenstein` repo was consolidated into a **library-named** layout (one dir per mocked dependency) with a self-contained `knowledge/` base (M2b). No behavior change.
- `corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`, and `README.md` now describe the alignment framework + the measured mock, with an explicit provenance note on what "100%" means.

### Fixed
- **`@clerk/express` alignment gate regression** — the bad-signature scenario was misclassified (`malformed` vs `bad-signature`) by a flawed signature-tamper introduced in the M2c close; the gate had silently dropped to 88.0%/85.7%. Corrected to a byte-level signature corruption with a property-based regression test; gate restored to 9/9 100%/100%. (caught + fixed at release close)
- A cross-milestone documentation/decision drift sweep (canonical framework doc, planning context, resolved decisions, metrics field).

### Known limitations
- "100%" means *indistinguishable from hand-authored / hybrid source goldens* (the reference behavior derived from the real libraries) — **not** a byte-diff against a live, network-connected Clerk tenant. Re-capturing goldens on a Clerk version bump is the drift loop's job. This is the right bar for a demo mock, not a production-Clerk conformance certificate.
- The `@clerk/express` gate is the only non-pure-Go gate (it drives the genuine Node SDK), so it runs locally / at close rather than in the pure-Go CI. CI-wiring is staged for v1.1.

### Supply chain
- The rosetta-committed code (`test/alignment/`) is **stdlib-only** — zero external modules. See `releases/archive/01.00-body-double/dependencies.lock`.
