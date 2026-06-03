# Release Review: v1.0 "body double"

**Date:** 2026-06-03
**Milestones:** M0, M1, M1b, M2, M2b, M2c (all closed-on-gate / completeness-complete, merged onto `release/01.00-body-double`)
**Review method:** full release diff (`main...release/01.00-body-double`, 123 files, +6686/−2) reviewed as one PR via a 6-dimension parallel sweep (scope · code-quality · docs · tests · decisions · supply-chain+adversarial) + a main-thread gate re-run of all four Clerkenstein gates.

**Branch topology (clean):** release is 49 commits ahead of `main`, 0 behind (`main`'s tip is the merge-base) → the release→main merge is **conflict-free**. `feat/demo-environment` (M0-D6 reconciliation) is an **ancestor** of release — already contained; merging release subsumes it, no separate merge needed.

## Verdicts by dimension
| Dimension | Verdict |
|---|---|
| Scope audit | GREEN — every v1.0 capability delivered; all four Clerk libraries under alignment; svix correctly framed as used-not-mocked; all carry-forwards fated |
| Code quality (test/alignment) | GREEN — stdlib-only, gofmt/vet clean, consistent, boundary-safe |
| Documentation | YELLOW → fixed |
| Tests & benchmarks | GREEN — framework suite passes; counts recorded |
| Decision consolidation | GREEN → minor staleness fixed |
| Supply-chain + adversarial | YELLOW → fixed (provenance guardrail added) |

## Blocker (found by close-release; fixed)
- [x] **[blocker] @clerk/express gate regressed to 88.0%/85.7%.** `ExpressAuth/bad-signature` scored `mirror=malformed` vs `source=bad-signature`. Root cause: the M2c-close adversarial fix (`69845c4`) replaced the bad-signature tamper with a *last-base64url-char flip*; on a 256-byte RS256 signature the last char is mostly padding bits, so the flip yields non-canonical base64url → `@clerk/backend` rejects it as a malformed JWS, not a signature mismatch. The unit test only checked the token *changed*, not its class, so it slipped through the close. **Fix:** `tamperSig` now decodes the signature, flips one bit of a real data byte, and re-encodes (deterministic, never a no-op, stays valid base64url of correct length → genuine signature mismatch). `TestTamperSig` rewritten to mint a real RS256 token and assert the distinguishing property — it would have caught this at unit time. Gate restored to **9/9 100%/100%**. (clerkenstein `abe4f33`.)

## Code Quality
- [x] [nice-to-have] `internal/compare/compare.go truncate()` cut at a fixed byte index → could split a multibyte rune into invalid UTF-8 in a divergence diagnostic. Fixed to cut on a rune boundary (`unicode/utf8.RuneStart`) + `TestTruncate_cutsOnRuneBoundary`.
- [info] `cmd/alignctl` coverage 67.5% — concentrated in non-logic `os.Exit`/usage plumbing. Acceptable for a developer CLI; recorded.

## Documentation
- [x] [should-fix] `corpus/architecture/alignment_testing.md` "How M1, M1b, and M2 consume this" never extended to M2c (flagged by 3 dimensions). Renamed heading → "…M2, and M2c…", added the M2c/@clerk/express bullet (3rd DNA, expressrun, verify-against-the-genuine-SDK, additive RS256/JWKS, Node-dependency CI note), noted Clerkenstein now drives **three DNAs via three runners**, updated "Where things live", bumped Last updated → 2026-06-03.
- [x] [should-fix] `corpus/services/clerkenstein.md` alignment-harness row still said `cmd/{clerkrun,jsfapirun}` + `golden{,-js}` — added expressrun + golden-express + "three DNAs".
- [x] [nice-to-have] `README.md` had zero mention of the v1.0 headline deliverable — added a status banner pointing at the alignment framework + Clerkenstein.
- [info] `alignment_testing.md:61` "two surfaces" is NOT stale (it means the two score-execution surfaces of the framework) — deliberately left unchanged.

## Tests & Benchmarks
- rosetta `test/alignment/`: 43 Test + 3 Fuzz funcs (added TestTruncate); gofmt/vet clean; all packages pass.
- Clerkenstein gates re-run on the release tip: **Go 22/22**, **JS 9/9**, **@clerk/express 9/9** (after the blocker fix), **drift 9/9**, race `go test -race ./...` 8/8 — all green.
- [x] [info] `m0/metrics.json` `test_funcs: 45` was actually the Test+Fuzz total; corrected to `test_funcs: 42, fuzz_funcs: 3` (= 45 total).

## Decision Consolidation
- [x] [should-fix] `m2c/decisions.md` M2c-D2 still titled "OPEN (the central iteration)" with "Record the resolution here when settled" — appended **RESOLVED 2026-06-03 → ADDITIVE (no migration)**.
- [x] [should-fix] `knowledge/plan/context.md` badly stale (claimed milestone dirs "not yet scaffolded", branch "not yet cut", omitted M2b/M2c) — refreshed to the closed reality + pointed at state.md as live source of truth.
- [x] [nice-to-have] `roadmap.md` M2c "Delivers" line said "surface count (3→4 mocked libraries)" — contradicts M2c-D5 (verified-not-mocked, no new dir). Reworded to "3rd *measured surface*, verified-not-mocked".

## Supply-chain
- [x] [info] `releases/01.00-body-double/dependencies.lock` written. rosetta-committed code (`test/alignment/`) is **stdlib-only, zero external modules** → supply-chain GREEN by construction. Clerkenstein's deps (clerk-sdk-go/v2, svix, @clerk/express) live in its own gitignored repo, not this release diff; recorded for reference. No GPL/AGPL.

## Adversarial (the value-add)
- [x] [should-fix] **"100%/100%" could be over-read as "verified against live Clerk".** The *source* side of the comparison is **hand-authored / hybrid goldens** (M1-D1) — reference behavior derived from the real libraries (and, for @clerk/express, confirmed by driving the genuine SDK), **not** captured from a live Clerk tenant. Added a provenance guardrail next to the 100% claim in `corpus/services/clerkenstein.md` and `clerkenstein/CLAUDE.md`: 100% = "indistinguishable from the encoded reference behavior," not a conformance certificate against production Clerk; re-capturing on a version bump is M1b's drift job.

## Deferred (tracked, not release-blocking)
- [nice-to-have → v1.1] `dna.Validate` / `compare.pct` treat a DNA with **zero critical genes** as 100% critical (`pct(n,0)=100`). No current DNA triggers it (all three declare critical genes), so it is dormant. A defensive guard (reject/flag a zero-critical DNA) is a framework-integrity polish routed to v1.1 to avoid expanding framework validation at release close. Recorded in `release-retro.md`.
- [carry-forward → v1.1 demo-stack] Wire the @clerk/express gate into CI (needs Node + @clerk/express present). The only non-pure-Go gate; runs locally/at close today.

## Outcome
All blockers and should-fix items resolved. The four Clerkenstein gates are green on the release tip and the rosetta framework suite is clean. Ready for the completeness gate + merge.
