---
title: "KB Fidelity Audit — M31 mkcert-trusted FAPI cert"
date: 2026-06-15
scope: milestone:M31
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| FAPI TLS cert generation (the cert step) | `corpus/ops/demo/recipe-browser-login.md §B step 2` | `demo-stack/up-injected.sh:337-353` (step 3a-bis) | PAIRED |
| Cert mount + `FAKE_FAPI_TLS_CERT/KEY` wiring | `corpus/ops/demo/recipe-browser-login.md §B step 2`; `frontend-tier.md` (silent — blind line) | `stack-injection/gen_injected_override.py:295-307`; `clerkenstein/cmd/fake-fapi/main.go:25-31` | PAIRED |
| pk mint (browser → fake FAPI) | `recipe-browser-login.md §B step 1`; `frontend-tier.md:83` | `stack-injection/inject.py` (`mint_pk`) | PAIRED |
| `DEMO_NO_*` flag family (opt-out pattern) | `frontend-tier.md` (`DEMO_NO_UI`) | `up-injected.sh` (`NO_UI`, `NO_SETDRESS`, `NO_LOCAL_CONTENT` parse) | PAIRED |
| BAPI cert-redirect (out of scope — plain HTTP) | `recipe-browser-login.md §A` | `clerkenstein/cmd/fake-bapi` (`http.ListenAndServe`) | PAIRED |
| README-index guard | CLAUDE.md skill table | `stack-core/corpus_index_guard.py` | PAIRED |

## Fidelity Findings

1. **The "ZERO change to gen_injected_override.py / inject.py / fake-fapi/main.go" claim — ALIGNED (load-bearing for M31).**
   - Source: overview.md scope; spec-notes.md.
   - Expected: a browser-trusted cert at the same `<stack>/certs/fapi.{crt,key}` paths "just works" with no change to the cert-consuming files.
   - Actual: `fake-fapi/main.go:28-31` reads `FAKE_FAPI_TLS_CERT/KEY` env and calls `http.ListenAndServeTLS(addr, cert, key, handler)`; `gen_injected_override.py:298,304-307` sets those env to `/certs/fapi.crt|key` and mounts `<stack>/certs:/certs:ro`. The cert *content* (issuer/trust) is irrelevant to all three — they reference it only by path. `inject.py` has no cert refs (mints the pk only). A trusted cert at the same path is served identically.
   - Verdict: ALIGNED. The milestone's "no-touch" guarantee on these three files is sound.

2. **`recipe-browser-login.md §B step 2` "one-time operator step" framing — STALE-by-design (the M31 deliverable).**
   - Source: `recipe-browser-login.md:56-60`.
   - Expected (post-M31): bring-up runs mkcert automatically.
   - Actual (today): documents the *manual* `mkcert -install` + import workaround.
   - Verdict: STALE — but this is exactly the doc M31 rewrites (manual→automatic). Not a blocker; it's the milestone's own scope. Tracked in progress.md.

3. **`frontend-tier.md` cert silence — UNDOCUMENTED (incidental, addressed in-milestone).**
   - `frontend-tier.md` describes the UI tier but says nothing about the FAPI cert / browser-trust. The roadmap Phase 0b flagged this; M31 adds a one-liner. Not load-bearing → does not block.

4. **`up-injected.sh:337-342` + `gen_injected_override.py:295` code comments — STALE-by-design.**
   - Both carry the "operator runs mkcert once" / "browser trusts it via a one-time mkcert/import" framing. M31 retires that framing as part of scope.

## Completeness Gaps

1. **No existing test covers the cert step (3a-bis).** The cert step is a bring-up ACTION sitting *below* the `UP_INJECTED_LIB_ONLY=1` early-return (`up-injected.sh:166`), so it is not currently sourceable as a lib-only function. The two existing test files (`test_tooling.py`, `test_frontend_build.py`) exercise the lib-only seam for the frontend builders + the secret pre-flight, and `test_tooling.py:179` asserts seam-ordering via `self.BODY` body-grep. M31's test should follow the `self.BODY`-grep pattern (values-blind, no live docker) to assert the mkcert branch + openssl fallback + `DEMO_NO_MKCERT` path. Incidental gap — does not block; addressed in Phase 1.

## Applied Fixes

None applied inline. Findings 2/3/4 are the milestone's own deliverables (the manual→automatic rewrite); applying them now would pre-empt the build. Finding 1 required no fix (ALIGNED). Recorded the topic→doc→code triples here for future-audit fast-start.

## Open Items (require user decision)

None.

## Gate Result

GREEN — proceed to Phase 1. Every topic PAIRED, every load-bearing claim ALIGNED (notably the no-touch guarantee on the three cert-consuming files is verified sound). The stale doc-framing findings (2/4) and the frontend-tier silence (3) are the milestone's own scope, not pre-existing blockers; the test gap (Completeness 1) is addressed in-build.
