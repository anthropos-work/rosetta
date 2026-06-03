# M2 — progress

Section milestone. 5 sections. Code → `clerkenstein` repo (gitignored `anthropos-demo/`); records +
docs → rosetta `m2/browser-webhook-coherence`. Zero platform-code changes throughout.

## S1 — JS FAPI spike + fake FAPI server — DONE
- [x] **Spike resolved (M2-D1, Fate 1):** `@clerk/*` points at a fake FAPI via publishable-key/`proxyUrl`
      config — no fork, no fallback. clerk-js derives the FAPI host from the key (`@clerk/shared`).
- [x] `fapi/key.go` — `MintPublishableKey(host)` ↔ `ParsePublishableKey` (matches `@clerk/shared`'s decode).
- [x] `fapi/server.go` — concurrency-safe fake FAPI: `/v1/environment`, `/v1/client`, sign-in/up
      create+attempt, session-token mint, `/v1/me`, JWKS, sign-out. Token = the **same HS256 universal-key
      JWT** the M1 authn twin verifies (browser↔backend coherence).
- [x] Tests 1:1 + sign-in bootstrap integration + `TestServer_tokenIsBackendVerifiable`. fapi 99%, race-clean.

## S2 — fake-Clerk-API-server (BAPI) + orgclient redirect (M1-D2) — DONE
- [x] `bapi/server.go` — serves the 10 orgclient methods' Clerk-SDK wire shapes, backed by the M1
      `orgclient` twin. Error-class mapping → Clerk `APIErrorResponse` shape.
- [x] (M2-D2) `orgclient.Store` made concurrency-safe (mutex) + concurrency test; InviteMember now persists.
- [x] `bapi/doc.go` — the `api.clerk.com` redirect recipe (DNS/`/etc/hosts` + trusted cert). Zero platform changes.
- [x] `bapi/server_sdk_test.go` — a **real** `clerk-sdk-go/v2` client parses every response. bapi 96%, race-clean.

## S3 — webhook injector — DONE
- [x] `webhook/injector.go` — svix-signs + POSTs the 12 event types to `/api/webhook/clerk`. Fails loud on non-2xx.
- [x] `webhook/events.go` — payload builders for user/org/membership/invitation (handler-read fields).
- [x] `TestInjector_signatureVerifiesAgainstPlatformSvix` (+ wrong-secret control + all-12 sweep + e2e). webhook 91%.

## S4 — JS-surface Alignment DNA + genes — DONE
- [x] `dna/clerk-js-5.json` — 9 FAPI-bootstrap genes; `SessionToken/decoded-identity` (`exact`) pins coherence.
- [x] `cmd/jsfapirun` — the JS-surface runner (mirror exercises `fapi/`; source = hand-authored shapes).
- [x] `golden-js/` captured; `alignctl run` → **100%/100% (9/9)**. `gate.sh` parameterized to gate it.
- [x] Fixed M1 latent bug (M2-D3): unanchored `.gitignore` had excluded `cmd/clerkrun/` source — recovered + tracked.

## S5 — documentation — DONE
- [x] `corpus/services/clerkenstein.md` — v0.2: JS path, fake FAPI, BAPI + redirect, webhook injector,
      JS DNA, spike outcome + un-exercised fallback, updated injection table + disarmed properties + testing.
- [x] `corpus/architecture/alignment_testing.md` § "How M1, M1b, and M2 consume this" — M2 = surface-generic proof.
- [x] Cross-refs (`frontend_architecture`, `next-web-app`, `webhook_setup`) added; anchors verified.

## PR review (Phase 3/4)
- Writer-led review of the full clerkenstein diff (~2.6k lines, 5 pkgs). 1 must-fix: jsfapirun's mirror
  path asserted `["response"].(map)` without a check → **panic risk** on an error envelope. Fixed with a
  `respMap` fail-soft helper + regression test (commit e07bc5b). No other findings; resource/leak +
  int-assertion + flag checks all clean.

## Pre-flight audits
- **Phase 0b (S1): GREEN** — all 6 KB-dependency docs PAIRED + aligned; 1 non-load-bearing stale claim
  (`clerk-integration.md` app SDK v2.5.1→v2.6.0) fixed inline. Report: `kb-fidelity-audit.md`. Reused for
  S2–S5 (same subsystem, knowledge docs unchanged).

## M2: Hardening

Code-side commits land in the `clerkenstein` repo (gitignored `anthropos-demo/`), baseline `67e5585`.
4 passes; stopped on stabilization (remaining gaps are `os.Exit` wrappers + unreachable defensive
branches — no shallow coverage-box tests written). Both alignment gates + the drift harness re-verified
green after hardening (no regression).

### Scope manifest (Phase 1 — M2-touched code, clerkenstein `9bc2541..HEAD`)
Source files grouped by package, with their co-located tests at harden start:
- **fapi/** — `key.go`, `server.go`, `resources.go` (tests: `key_test.go`, `server_test.go`, `resources_test.go`). Start 99.0%.
- **bapi/** — `server.go`, `resources.go`, `doc.go` (tests: `server_test.go`, `server_sdk_test.go`). Start 95.8%.
- **webhook/** — `injector.go`, `events.go` (tests: `injector_test.go`, `events_test.go`). Start 91.1%.
- **orgclient/** — `store.go`, `invitations.go` (test: `store_test.go`; **no concurrency test despite M2-D2 mutex**). Start 98.1%.
- **cmd/jsfapirun/** — `main.go` (test: `main_test.go`; `main()`/CLI surface untested). Start 73.9%.
- **cmd/clerkrun/** — `main.go` (M2-D3 recovered M1 runner; **no test at all**). Start 0.0%.

### Pass 1 — 2026-06-03
**Coverage delta (milestone-touched files):**
- fapi: 99.0% → 100.0% (+1.0); orgclient: 98.1% → 100.0% (+1.9)

**Tests added:**
- `fapi/fuzz_test.go`: 2 fuzz (FuzzParsePublishableKey, FuzzMintParseRoundtrip)
- `fapi/key_test.go`: 1 regression (embedded/missing-sentinel)
- `orgclient/concurrency_test.go`: 5 `-race` concurrency (mutate sweep, same-org membership, dup-invite, revoke, read/write — all assert exactly-one-winner contention invariants)

**Bugs fixed inline:**
- `ParsePublishableKey` leaked the `$` sentinel into the returned host when the decoded bytes carried a
  `$` anywhere but the trailing position (or none at all) — `TrimSuffix` only stripped a *trailing* `$`,
  breaking the Mint↔Parse inverse. Found by `FuzzParsePublishableKey` (input `pk_test_00000CQ0`). Fixed:
  require the decoded value to end with the `$` sentinel and the host to be `$`-free (commit e80a257).

**Knowledge backfill (Phase 3b):** Pass 1's fuzz-found bug pinned the publishable-key codec invariant
(the `$` is a *terminator* sentinel; Parse is the strict inverse of Mint, rejecting embedded/missing
sentinels). Blended into `corpus/services/clerkenstein.md` as a "Codec invariant (M2 hardening)" callout
on `fapi/key.go`, plus a refreshed Testing section (post-harden coverage + the new fuzz/concurrency
dimensions). Passes 2–4 surfaced no new KB-worthy behavioral facts (robustness on already-documented
surfaces) — recorded here so the audit trail shows the rubric was applied each pass.

### Pass 2 — 2026-06-03
**Coverage delta:** webhook 91.1% → 95.6% (+4.5; Inject 78.6% → 92.9%). bapi steady 95.8% (the 3
residual branches are unreachable — see below); the malformed-body tests add behavioral robustness on
already-covered lines.

**Tests added:**
- `bapi/malformed_test.go`: 5 malformed-input (garbage/wrong-type/empty/oversized/unicode bodies + path values across createOrg/createMembership/bulkInvite/revoke/metadata)
- `bapi/fuzz_test.go`: 2 fuzz (FuzzCreateOrganizationBody, FuzzBulkInviteBody — both clean, no crashers)
- `webhook/injector_error_test.go`: 4 error-path (build-request error, transport error, empty-payload Sign, randHex sanity)

**Bugs fixed inline:** none.

### Pass 3 — 2026-06-03
**Coverage delta:** cmd/jsfapirun 73.9% → 93.8% (+19.9; statusErr & tokenIdentity → 100%).

**Tests added:**
- `cmd/jsfapirun/cli_test.go`: 7 CLI/protocol (valid mirror/source, bad target, missing/malformed/empty DNA, unknown flag — exit-code contract)
- `cmd/jsfapirun/gene_error_test.go`: 5 gene-error-branch (pkGene mint-failed, statusErr fallbacks, tokenIdentity empty/verify-fail)
- `cmd/jsfapirun/transport_test.go`: 6 (dead-server transport degrade + errc/val/present helper shapes)
- Refactor: `main()` → testable `run(args, stdout, stderr) int` (no behavior change beyond fail-fast `--target` validation; gate-safe)

**Bugs fixed inline:** none.

### Pass 4 — 2026-06-03
**Coverage delta:** cmd/clerkrun 0.0% → 96.8% (+96.8).

**Tests added:**
- `cmd/clerkrun/main_test.go`: 8 (full-DNA dispatch with value-vs-error_class per gene, VerifyToken identity incl. org, CLI exit-code contract, unknown-capability skip)
- Same testable-`run` refactor as jsfapirun. **Go alignment gate re-verified 100%/100% 22/22 with the refactored binary.**

**Bugs fixed inline:** none.

### Defensive / unreachable residuals (documented, not tested — no shallow tests)
These branches cannot fire through the current disarmed code paths; testing them would require injecting
production-code failures (gold-plating):
- `bapi/server.go updateOrganization` error branch — `store.UpdateExternalID` always returns nil.
- `bapi/server.go membershipErr` default — the membership handlers only ever surface the 4 mapped store error classes.
- `bapi/resources.go invitationRes` empty-status default — callers always pass `pending`/`revoked`.
- `webhook/injector.go Sign`/`Inject` svix-sign-failure branches — unreachable with a valid svix secret.
- `cmd/{jsfapirun,clerkrun} main()` — 1-line `os.Exit(run(...))` wrappers (the logic is in the tested `run`).
- `cmd/jsfapirun mirrorGene` per-gene no-response branches — no DI seam (each gene builds a live in-process
  server); the helpers they delegate to (`respMap`/`statusErr`) are covered.

### Final coverage (M2-touched packages)
fapi 100.0% · orgclient 100.0% · clerkrun 96.8% · bapi 95.8% · webhook 95.6% · jsfapirun 93.8% · authn 100.0% (M1, unchanged).

### Flake gate
3 consecutive `-race` runs of the full suite, all green, zero data races.

### Regression proof (post-harden, no regression)
- Go gate (`scripts/gate.sh`): **100%/100% 22/22**, exit 0.
- JS gate (`RUNNER_PKG=./cmd/jsfapirun … scripts/gate.sh`): **100%/100% 9/9**, exit 0.
- Drift harness (`scripts/drift-test.sh`): **ALL PASS 9/9**, exit 0.

### Stop condition
Stopped after Pass 4 (well under the 5-pass cap): the full Step-2b scan found no further behavioral gaps
worth a real test — remaining uncovered lines are `os.Exit` wrappers and unreachable defensive branches.
Next-pass marginal coverage delta < 2%; no flakes. Handing off to `/developer-kit:close-milestone`
(its Phase 4 audit runs independently as defense-in-depth).
