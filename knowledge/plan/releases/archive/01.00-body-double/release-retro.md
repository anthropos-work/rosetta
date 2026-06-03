# Release Retro — v1.0 "body double"

**Shipped:** 2026-06-03 · **Milestones:** M0 → M1 → M1b → M2 → M2b → M2c (6) · **Headline:** a *measured* drop-in Clerk mock (Clerkenstein) at **100% / 100% on all three measured surfaces** (Go, JS/FAPI, `@clerk/express`), built and scored by a new first-class **alignment-testing framework**, with **zero platform-code change**.

## What v1.0 delivered
- **M0** — the reusable alignment framework (`test/alignment/`: `alignctl`, 4 equivalence operators, weighted scoring, golden capture, the `//go:build alignment` test class, `/align-dna` + `/align-run` skills, a toy reference that catches its own intentional divergence). Stdlib-only.
- **M1** — the Clerkenstein Go backend mirror, driven to a 100%/100% gate vs `clerk-sdk-go/v2 @ 2.6.0` (22 genes), offline.
- **M1b** — drift detection: `gate.sh`/`drift-check.sh` exit-code contract + a weekly CI alignment gate + a 9-assertion drift harness. A Clerk bump becomes a flagged, mechanical event.
- **M2** — the JS browser (FAPI) + webhook seams: a 2nd DNA (`clerk-js-5`, 9 genes) at 100%/100%, proving the framework is surface-generic; the fake BAPI that disarms the platform's networked orgclient; the svix-signed webhook injector.
- **M2b** — repo consolidation into library-named dirs + a self-contained `knowledge/` base (gates stayed green).
- **M2c** — the last un-gated consumer, `@clerk/express` (studio-desk's Node backend), brought under alignment as a 3rd DNA (`clerk-express-1`, 9 genes) at 100%/100% via an **additive RS256/JWKS** path (no HS256 migration), verified against the **genuine SDK** (the svix discipline reused).

## Incidents

### P1 — @clerk/express gate shipped RED out of the M2c close (caught by close-release)
The M2c close's adversarial fix (`69845c4`) replaced the bad-signature tamper with a last-base64url-char flip to remove a theoretical no-op. On a 256-byte RS256 signature the last base64url char is mostly padding bits, so the flip produced **non-canonical base64url** that `@clerk/backend` rejects as a *malformed JWS* — classifying `ExpressAuth/bad-signature` as `malformed` instead of `bad-signature`. The gate silently regressed to **88.0%/85.7%**. It slipped through the close because the unit test (`TestTamperSig`) only asserted the token *changed*, not how the real SDK *classifies* it, and the close's "express gate re-verified" step did not re-run the full real-SDK gate end-to-end after the tamper change.

- **Detection:** close-release Phase 0/4 gate re-run (the first end-to-end express-gate run since the tamper change).
- **Fix (clerkenstein `abe4f33`):** `tamperSig` now decodes the signature, flips one bit of a real data byte, and re-encodes — deterministic, never a no-op, stays valid base64url of the original byte length → a genuine signature mismatch (`token-invalid-signature` → `bad-signature`). `TestTamperSig` rewritten to mint a real RS256 token and assert the distinguishing property (valid base64url, same length, different bytes) — it would have caught the regression at unit time.
- **Lesson (carry into v1.1):** an adversarial test-input change that alters what a *real external verifier* returns MUST be validated by re-running the gate that consumes it, not just a structural unit assertion. The regression guard is now the right shape.

### P2 — doc/decision staleness accumulated across 6 milestones
Several "incremental seam" artifacts were never refreshed to the post-M2c reality: the canonical `alignment_testing.md` consumption section stopped at M2 (flagged by 3 independent review dimensions); `context.md` still claimed milestones "not yet scaffolded / branch not cut"; M2c-D2 sat "OPEN" after being resolved; the roadmap "Delivers" line mis-predicted "4 mocked libraries". All fixed in close-release Phase 7. Lesson: the release boundary is the right place to reconcile cross-milestone doc drift — per-milestone closes don't see the whole story.

## Cross-milestone patterns (reusable doctrine for v1.1+ mirrors)
- **Verified, not reimplemented** — `clerk-webhook/` (svix) and `@clerk/express` (M2c) both *satisfy the genuine library* rather than mocking it. The alignment runner drives the real SDK; the score measures whether the real consumer accepts our output. Generalizes to any library-with-a-verifier.
- **Additive over migration** — M2c added an RS256 path beside the HS256 seams rather than migrating them; the feared M1/M2 re-gating never happened. Prefer parallel paths when domains are separable.
- **Crux-first iteration** — M2c proved the load-bearing unknown (real `@clerk/backend` accepts our RS256 token) in iter-04 before building the full runner, de-risking the whole milestone.
- **Disarmed by design** — one universal HS256 key + a fixed demo RS256 key + no `alg` validation are *intended* properties (demo speed + accessibility), documented as such; never "hardened".

## Honesty guardrail added (adversarial review)
"100%" alignment measures *indistinguishability from hand-authored / hybrid source goldens* (M1-D1) — reference behavior derived from the real libraries, not a byte-diff against a live Clerk tenant. A provenance note now sits next to every headline 100% claim (`corpus/services/clerkenstein.md`, `clerkenstein/CLAUDE.md`) so the score isn't over-read as a production-Clerk conformance certificate. Re-capturing goldens on a version bump is M1b's drift loop.

## Carried forward → v1.1 "show floor"
- **Wire the @clerk/express gate into CI** — the only non-pure-Go gate (needs Node + `@clerk/express`); runs locally/at close today. Lands when v1.1's demo stack sets up studio-desk's `node_modules` in CI.
- **`dna.Validate` zero-critical-genes guard** — `pct(n,0)=100` would report 100% critical for a DNA with no critical genes. Dormant (all three DNAs declare critical genes); a defensive guard is v1.1 framework-integrity polish.
- **v1.1 scope** — multi-instance disposable stacks + data seeding + use-case recipes (M3–M5), staged in `roadmap-vision.md`.

## Metrics delta (baseline — first release)
See [`metrics.json`](metrics.json) (aggregate) + [`project-stats.json`](project-stats.json) (snapshot). No previous release to regress against — this is the baseline. Gates: Go 22/22, JS 9/9, @clerk/express 9/9, drift 9/9; triple-clean 3/3; flakes 0. rosetta framework 43 test + 3 fuzz (stdlib-only); Clerkenstein 123 test + 6 fuzz across 8 packages / 3 DNAs / 3 runners.

## Process note
The `/developer-kit:project-stats` skill was not invoked at Phase 8c (unverified availability); an equivalent snapshot was written manually to `project-stats.json`.
