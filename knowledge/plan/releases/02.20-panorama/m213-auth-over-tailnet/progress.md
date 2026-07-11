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

## M213: Hardening

### Pass 1 — 2026-07-11 (rext `b9f41dd`; tag `panorama-m213` re-pointed → `b9f41dd`)
Scope manifest (M213-touched, `770f81b..d8f28c3`): **src** — `stack-injection/gen_tailscale_serve.py` (py),
`stack-injection/inject.py` (py), `demo-stack/up-injected.sh` (sh: `gen_tailscale_fapi_cert` + dotted-host +
topology guards + serve/egress wiring), `clerkenstein/clerk-frontend/server.go` (go: `clerkJSCDNBase` +
`handleClerkJSBundle`). **tests** — `stack-injection/tests/test_injection.py`, `demo-stack/tests/test_tooling.py`,
`clerkenstein/clerk-frontend/clerkjs_proxy_test.go`. Docs `clerkenstein/knowledge/*.md` (no tests).

**Coverage delta (milestone-touched files):**
- `gen_tailscale_serve.py`: **64% → 98%** statements (+34; only the `if __name__` entrypoint guard remains — exercised by the subprocess test, uncoverable in-process)
- `inject.py`: **98% → 98%** (steady; only its entrypoint guard remains — `require_dotted_host`/`mint_pk`/`parse_pk`/`main()` fully covered)
- `clerkenstein/clerk-frontend` touched funcs (`clerkJSCDNBase`, `handleClerkJSBundle`): **100% / 100%** (already max at build; new tests add behaviour assertions, not lines)
- `up-injected.sh` (shell — no line-coverage tool): the two safety guards moved from grep-pinned to **functionally executed**

**Tests added (11):**
- `test_injection.py`: 5 `gen_tailscale_serve.main()` IN-PROCESS (stdout plan / `--out` write+chmod+stderr-port-count / `--no-ui` = 2 API ports / `--target-host` seam / localhost no-op) + 1 `require_dotted_host` boundary (empty + IPv6 fail the has-a-dot predicate)
- `test_tooling.py`: 4 FUNCTIONAL guard-execution tests (extract + source under `set -euo pipefail`) — dotless host aborts (assertValidPublishableKey), split FAPI host aborts (topology guard), valid same-host MagicDNS passes, unset host is a no-op even when HOST/FAPI mismatch (proves the opt-in gate)
- `clerkjs_proxy_test.go`: 3 behaviour tests — CDN non-200 forwarded transparently (404 not masked/502, JS MIME still applied) / no-query request appends no trailing `?` / empty `FAKE_FAPI_CLERKJS_CDN` falls back to default (`v != ""` guard)

**Bugs fixed inline:** none — no production code touched; the M213 build surface held up under deeper testing.

**Flakes stabilized:** none — 3 consecutive clean sequential runs of the new tests (they use subprocess/tempdir/bash/awk; verified deterministic, `go test -count=1`).

**Knowledge backfill:** no KB-worthy findings — the new tests pin behaviour already documented in `decisions.md`
(D-PROXY-2 the serve model, D-PK-1 the dotted-host split, D-TOPO-1 the topology guard, D-EGRESS-1 the CDN override)
and `spec-notes.md`; no new invariant, edge semantic, or flake root-cause surfaced.

**Full suites GREEN:** stack-injection 144 passed / 8 skipped; demo-stack test_tooling 128 passed; clerk-frontend
`-race` + `go vet` + `gofmt` clean; `up-injected.sh` shellcheck rc 0.

### Stop condition
Stopped after Pass 1: the Phase-2b scan found no meaningful remaining gap (the residual shell serve-wiring/egress
glue is a thin gated non-fatal wrapper whose payload — the generator, guards, and cert branch — is fully tested;
functionally executing it would be a heavy integration harness with <2% return); touched Python files are at their
practical coverage ceiling (98%, remaining lines are uncoverable `if __name__` entrypoint guards); Go touched funcs
100%; zero flakes. Build env has no tailnet host, so `tailscale cert`/`serve` stay stubbed — the live run is M215.
