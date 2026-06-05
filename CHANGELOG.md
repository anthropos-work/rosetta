# Changelog

All notable user-facing changes to Project Rosetta. Format: [Keep a Changelog](https://keepachangelog.com/), semver-aware.

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
