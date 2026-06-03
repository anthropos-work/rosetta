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
