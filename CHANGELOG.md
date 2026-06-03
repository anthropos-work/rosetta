# Changelog

All notable user-facing changes to Project Rosetta. Format: [Keep a Changelog](https://keepachangelog.com/), semver-aware.

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
