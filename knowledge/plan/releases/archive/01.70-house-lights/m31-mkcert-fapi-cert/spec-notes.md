# M31 — Spec notes

_Technical detail. Stub at scaffold; seeded from the design (`.agentspace/scratch/roadmap-research-2026-06-15.md`)._

## The exact edit — up-injected.sh step 3a-bis (inside the keep-existing-cert guard)
```
if command -v mkcert >/dev/null 2>&1 && [ -z "${DEMO_NO_MKCERT:-}" ]; then
  mkcert -install >/dev/null 2>&1 || true        # idempotent; no-op + no sudo when CA already trusted
  if mkcert -cert-file "$CERTS/fapi.crt" -key-file "$CERTS/fapi.key" 127.0.0.1 localhost ::1 >/dev/null 2>&1; then
    log "FAPI TLS cert minted via mkcert (browser-trusted; no proceed-anyway) — $CERTS/fapi.crt"
  else
    log "warning: mkcert mint failed — falling back to openssl self-signed (proceed-anyway needed)"
    <openssl branch verbatim>
  fi
else
  log "mkcert not on PATH (or DEMO_NO_MKCERT) — openssl self-signed FAPI cert (browser needs a one-time trust)"
  <openssl branch verbatim — current lines 347-352>
fi
```
Detection: `command -v mkcert` (PATH-based, NOT the hard-coded /opt/homebrew path).

## Fallback / idempotency / once-vs-per-stack
- Fallback = openssl self-signed verbatim (the never-abort-a-good-bring-up contract; M13/M18). Both branches write
  byte-compatible outputs (same two filenames, valid TLS, SAN 127.0.0.1+localhost).
- Outer keep-existing guard unchanged → re-ups reuse the cert, never re-mint. `mkcert -install` is itself idempotent.
- CA install = once-per-machine (shared by all stacks); the leaf is per-stack on disk but stack-invariant (host is
  always 127.0.0.1; certs carry no port → valid for every demo-N).

## DEMO_NO_MKCERT flag plumbing
Parse alongside the existing DEMO_NO_* flags in up-injected.sh; document in the SKILL + recipe-browser-login.md.

## BAPI — out of scope
fake-bapi serves plain HTTP (`http.ListenAndServe`), server-to-server only via the `api.clerk.com` alias + networkless
JWT verify — the browser never does a BAPI TLS handshake.

## Doc rewrite — recipe-browser-login.md §B
manual mkcert-install/import workaround → automatic; + security note (dev CA in trust store), remote/VM + Firefox/certutil
caveats, the DEMO_NO_MKCERT opt-out, the cert-expiry/`rm <stack>/certs/fapi.crt`-regenerates note.

## Risks (see overview's risk map in roadmap.md)
fresh-machine `-install` OS-password prompt; remote/VM trust-store-is-on-the-wrong-machine; dev-CA security; ~3y cert
expiry with no guard-check; the shared-helper forward-note for a future dev-N --local-content UI path.

## Pre-flight audits — the cert step (mkcert branch)
- **Phase 0b KB-fidelity: GREEN** (2026-06-15, report `kb-fidelity-audit.md`). Verified the load-bearing
  no-touch guarantee: `fake-fapi/main.go:28-31` reads `FAKE_FAPI_TLS_CERT/KEY` and serves `ListenAndServeTLS`;
  `gen_injected_override.py:298,304-307` sets those env to `/certs/fapi.crt|key` + mounts `<stack>/certs:/certs:ro`;
  `inject.py` has no cert refs. The cert is referenced by **path only** → a browser-trusted cert at the same
  path serves identically (ZERO change to all three). Stale doc-framing (recipe §B step 2, the two code comments)
  + frontend-tier cert silence are the milestone's own deliverables, not pre-existing blockers.
- **Test seam:** the cert step (3a-bis, ~line 337) sits BELOW the `UP_INJECTED_LIB_ONLY=1` early-return (`up-injected.sh:166`),
  so it is NOT sourceable as a lib-only function today. No existing test covers it. M31's test follows the
  `self.BODY`-grep pattern (`test_tooling.py:179` precedent) — values-blind, no live docker — asserting the
  mkcert branch + openssl fallback + `DEMO_NO_MKCERT` path on the script text.
- **Triples:** recipe-browser-login.md §B ↔ up-injected.sh 3a-bis; frontend-tier.md (cert one-liner to add) ↔
  gen_injected_override.py:295-307 / fake-fapi/main.go:25-31; README-index guard = `stack-core/corpus_index_guard.py`
  (baseline exit 0 confirmed).
