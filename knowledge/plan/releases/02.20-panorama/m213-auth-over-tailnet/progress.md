# M213 — progress

Section checklist (closure = the fake FAPI serves a trusted `tailscale cert` for the MagicDNS host, the re-minted
pk validates, and the browser surface is reachable under one HTTPS origin — proven locally before the M215 live run).

- [x] cert mint swaps mkcert/openssl → `tailscale cert` when `STACK_PUBLIC_HOST != localhost` (same output paths)
      — §A, rext `7287023`. `gen_tailscale_fapi_cert` gated on the public host; falls back to the factored
      `gen_local_fapi_cert` (mkcert/openssl) non-fatally; +6 values-blind tests (stub `tailscale`). Not run live (M215).
- [x] pk re-mint validates (`assertValidPublishableKey` dotted-host pass + base64 round-trip)
      — §B, rext `63b5f7a`. `inject.py::require_dotted_host` (wiring gate) + an early up-injected.sh guard; the codec
      stays permissive (key.go mirror the gene needs); +7 tests. Refines M212 D-IMPL-1 → D-PK-1.
- [x] reverse proxy (`tailscale serve`) fronts the browser-facing ports on the MagicDNS host, one origin
      — §C, rext `7acbc76`. New `gen_tailscale_serve.py` (per-port HTTPS, FAPI excluded/self-TLS, 0 new deps —
      D-PROXY-1/2); gated + non-fatal wiring in up-injected.sh; +14 tests. Config generation only (live run = M215).
- [x] topology guard: FAPI same host as app; PSL status of taildc510.ts.net verified
      — §D, rext `fdf74b3`. `[ "$HOST" != "$FAPI_HOST" ]` ⇒ exit 1 (regression tripwire). PSL VERIFIED: `ts.net`
      is the public suffix ⇒ `taildc510.ts.net` is the eTLD+1 (corrects the spec's SameSite=None guess; D-TOPO-1). +2 tests.
- [x] build-rebuild guard trips on HOST change (no stale localhost image)
      — §E CONFIRMED (D-REBUILD-1). The M211 `want_ep="http://$HOST:…"` embeds `$HOST`, so a stale localhost image
      mismatches → rebuild. Already test-covered (80/80 frontend-build; 3 HOST-invalidation tests). No new code.
- [x] fake-fapi egress to cdn.jsdelivr.net confirmed
      — §F, rext `636db11`. Made explicit + testable + overridable: `clerkJSCDNBase()` + `FAKE_FAPI_CLERKJS_CDN`;
      +4 Go tests; a non-fatal host-side egress pre-check in up-injected.sh (+1 pin). JS align re-scored 100%/100%. D-EGRESS-1.
- [x] tests (cert-path selection, pk validation, proxy config generation)
      — §G. +33 M213 tests across the sections (all inline). Full suites GREEN: demo-stack 363, stack-injection 138,
      clerkenstein Go (shared+clerk-frontend+cmd) vet+race clean, shellcheck clean. Alignment JS/FAPI 100%/100%.

Docs (Phase 5): KB-1..KB-5 addressed — `clerkenstein.md` (tailscale-cert FAPI path + gene ref), `recipe-browser-login.md`
+ `frontend-tier.md` (remote-cert path + egress), rext `clerkenstein/knowledge/{architecture,alignment}.md`, spec-notes
line-anchor + PSL fixes. The proxy-topology recipe (`tailscale-serve.md`) + CORS/link emission are **M214**; the live
cross-machine acceptance + cert renewal + RAM burn-down are **M215** (Fate-2, already-owned scope).

rext code-of-record: tag **`panorama-m213`** at rext HEAD (post-build). rosetta plan/doc commits on `m213/auth-over-tailnet`.
