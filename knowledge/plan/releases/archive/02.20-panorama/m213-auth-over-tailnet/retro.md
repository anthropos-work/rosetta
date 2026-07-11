# M213 — retro (auth over the tailnet)

## Summary
The v2.2 "panorama" auth/TLS milestone (∥ with M214). It makes Clerkenstein auth complete from a **remote** tailnet
browser and serves the whole browser surface under **one trusted HTTPS MagicDNS origin**: the `tailscale cert` FAPI
mint swap (a path-only drop-in — same `certs/fapi.{crt,key}` the mkcert/openssl mint writes, so the mount +
`ListenAndServeTLS` are untouched — with a non-fatal local fallback), dotted-pk validation lifted up into the demo
wiring (`require_dotted_host`, while the codec stays permissive for the alignment gene), the NEW
`gen_tailscale_serve.py` reverse-proxy generator (per-port HTTPS, **0 net-new deps**, the FAPI self-TLS-excluded), the
FAPI-same-host topology guard (PSL-verified `ts.net` eTLD+1), the confirmed build-rebuild-on-HOST guard, and the
`cdn.jsdelivr.net` egress made explicit + operator-overridable (`FAKE_FAPI_CLERKJS_CDN`). All 7 sections landed, all
gated on `--public-host` (byte-identical when unset). Config generation + unit tests only — the live cross-machine run
is M215. Closed clean on the first close pass — merged `--no-ff` into `release/02.20-panorama`.

## Incidents This Cycle
None. No P1/P2, no flakes (gate 5/5: go `-race`+shuffle×5, stack-injection ×5, demo-stack test_tooling ×5), no
regressions. Build + harden Pass 1 + 4 close adversarial scenarios (ADV-1..ADV-4) surfaced zero defects in the
production code.

## What Went Well
- **Config-only, and the design proved it.** The three load-bearing premises (dotted-host validator, host-agnostic
  token verify, path-only cert mount) were code-verified aligned at design, so the whole milestone stayed
  configuration + generation — no Clerk-contract change, no backend verify change.
- **The live PoC de-risked the make-or-break assumption BEFORE build.** `sudo tailscale cert billion.taildc510.ts.net`
  minted a real Let's Encrypt cert and a remote tailnet machine fetched an HTTPS endpoint from it with
  `ssl_verify_result=0` (trusted, zero CA install) — the exact path-mounted-cert model the fake FAPI uses. The
  mkcert→`tailscale cert` swap was a genuine drop-in, not a hoped-for one.
- **0 net-new deps held under pressure.** `tailscale serve` (already on every target VM) was chosen over Caddy
  (D-PROXY-1), keeping the release's supply-chain target intact — the one item `state.md` flagged to weigh in M213.
- **Egress turned from a hardcoded string into a tested, overridable seam.** `clerkJSCDNBase()` + `FAKE_FAPI_CLERKJS_CDN`
  made the FAPI's one outbound dependency unit-testable (httptest, no real network) and mirror-overridable; the JS/FAPI
  alignment re-scored 100%/100%.
- **Harden functionally executed the safety guards.** The dotted-host + topology guards moved from grep-pinned to
  actually sourced-and-run under `set -euo pipefail`, so the close adversarial review found them already handled.

## What Didn't
- **A rext section-README index row can't be added at milestone close.** `stack-injection/README.md`'s "What's here"
  table gained no row for the new `gen_tailscale_serve.py` generator. Because the rext code-of-record is FROZEN at tag
  `panorama-m213` (a rosetta-only close must not re-point it), the fix routes to close-release rather than being papered
  over — the same class as M212's D-CLOSE-1. Low severity (the generator itself is fully docstring'd + tested).
- **The python3.14 tooling friction recurred (non-blocking).** The box's default `python3` (3.14) has no usable pytest;
  the suites + JUnit tally ran under python 3.12, as at M212. No milestone impact; noted again for the next close.

## Carried Forward
- **D-CLOSE-2** — `stack-injection/README.md` missing a `gen_tailscale_serve.py` index row → **v2.2 close-release**
  (bundled with M212's D-CLOSE-1 when rext re-tags). Fate 2, in-release.
- **Reverse-proxy topology recipe (`tailscale-serve.md`) + CORS/link emission** → **M214**. Fate 2, confirmed-owned.
- **Live cross-machine acceptance + the loopback-vs-0.0.0.0 serve port-binding reconciliation + cert renewal (90-day
  LE) + RAM/swap burn-down** → **M215** (the exit-gate milestone). Fate 2, confirmed-owned.

## Metrics Delta
- **Tests:** go clerk-frontend **+7** `clerkjs_proxy` funcs (rext Go 1764→**1771**; `-race`+shuffle×5 clean);
  stack-injection **152** (144p/8s/0f); demo-stack **367** (0f) — JUnit-authoritative. +33 build + 11 harden Python funcs.
- **Coverage:** `gen_tailscale_serve.py` 64%→98%; `inject.py` held 98%; clerk-frontend touched funcs 100%.
- **Flake:** 0 (5/5 gate). **Supply-chain:** 0 net-new deps held. **Alignment:** 100%/100% (JS/FAPI re-scored).
  **Platform-repo edits:** 0. **rext close edits:** 0 (frozen tag).
- Full machine-readable record: `metrics.json`.
