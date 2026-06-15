---
milestone: M31
slug: mkcert-fapi-cert
version: v1.7 "house lights"
milestone_shape: section
status: planned
created: 2026-06-15
last_updated: 2026-06-15
complexity: small
delivers: rosetta-extensions/demo-stack (the mkcert bring-up step; ext tag house-lights-m31) + corpus/ops/demo/recipe-browser-login.md rewrite (manual→automated)
backlog_refs: (none — triggered by the live next-web blank-page defect 2026-06-15)
---

# M31 — mkcert-trusted FAPI cert (the browser-login render fix)

## Goal
A fresh browser at demo-N's next-web (e.g. `http://localhost:33000`) renders the signed-in app with **zero manual
cert-trust / proceed-anyway**, by minting a locally-trusted (mkcert) FAPI TLS cert at bring-up — degrading cleanly to
the current openssl self-signed path when mkcert is absent.

## Why section
The change surface is fully known + verified: ONE branch in `up-injected.sh` step 3a-bis, same cert output paths, no
other code touched. Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `house-lights-m31` → consume): the mkcert branch in `demo-stack/up-injected.sh`
  + the `DEMO_NO_MKCERT` flag parse + the code-comment updates.
- **`rosetta`**: `corpus/ops/demo/recipe-browser-login.md` rewrite + `frontend-tier.md` one-liner + the demo-up SKILL note.

## Scope
- **In:** `up-injected.sh` step 3a-bis (~lines 344-353), inside the existing keep-existing-cert guard — branch on
  `command -v mkcert` (and the `DEMO_NO_MKCERT` opt-out): mkcert branch runs idempotent `mkcert -install` (`|| true`) then
  `mkcert -cert-file "$CERTS/fapi.crt" -key-file "$CERTS/fapi.key" 127.0.0.1 localhost ::1`; **else / on-mint-failure keeps
  the openssl gen verbatim**. Non-fatal throughout.
- **In:** `DEMO_NO_MKCERT=1` escape hatch (parsed in `up-injected.sh`, mirrors `DEMO_NO_UI`/`DEMO_NO_SETDRESS`/`DEMO_NO_LOCAL_CONTENT`).
- **In:** ZERO change to `gen_injected_override.py` / `inject.py` / `fake-fapi/main.go` (trusted cert at the same paths
  "just works": mount + `FAKE_FAPI_TLS_CERT/KEY` + `ListenAndServeTLS` unchanged).
- **In:** docs — rewrite `corpus/ops/demo/recipe-browser-login.md §B` (manual mkcert workaround → automatic; + the
  dev-CA-in-trust-store security note, the remote/VM + Firefox/`certutil` caveats, the `DEMO_NO_MKCERT` opt-out, the
  cert-expiry/regenerate note); a `frontend-tier.md` cert one-liner; the demo-up SKILL browser-login note; the
  `up-injected.sh:337-342` + `gen_injected_override.py:295` comments (retire the "operator runs mkcert once" framing).
- **In:** a one-line forward-note in the code comment — a future dev-N `--local-content` UI path would want the same
  mkcert wiring (candidate to extract as a shared helper, not re-inline).
- **Out:** the fake BAPI (plain HTTP, server-side only — no browser TLS handshake); the studio-desk redirect (M32);
  ant-academy liveness (backlog).
**Depends on:** none.
**Parallel with:** M32 in principle, but sequence M31→M32 (shared `up-injected.sh` + demo doc cluster).
**Estimated complexity:** small.
**Open questions:** none blocking (fallback=openssl-not-fail-loud, BAPI-out-of-scope, SANs `127.0.0.1 localhost ::1` — all decided).
**KB dependencies:** `corpus/ops/demo/recipe-browser-login.md`; `corpus/services/clerkenstein.md`; `corpus/ops/demo/frontend-tier.md`.
**Delivers →** `rosetta-extensions/demo-stack` (ext tag `house-lights-m31`) + the `recipe-browser-login.md` rewrite.

## Verify (close-time)
A fresh (non-Playwright) browser at the demo's next-web renders `/home` signed-in with no proceed-anyway — the live
end-to-end proof deferred from design (design rested on the Playwright-`ignoreHTTPSErrors`-renders proof).
