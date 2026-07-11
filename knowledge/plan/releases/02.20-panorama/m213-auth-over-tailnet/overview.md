---
milestone: M213
slug: auth-over-tailnet
version: v2.2 "panorama"
milestone_shape: section
status: planned
created: 2026-07-11
last_updated: 2026-07-11
complexity: medium
depends_on: M212
delivers: Clerkenstein auth + the whole browser surface served under ONE trusted HTTPS MagicDNS origin (tailscale cert + reverse proxy + re-minted pk)
issues: FAPI host + pk baked to 127.0.0.1 (clerk-js hits the viewer's own loopback); mkcert cert trusted only on the bring-up host; app surface is plain HTTP (no secure context → Web Crypto blocked off-localhost)
---

# M213 — Auth over the tailnet (TLS + FAPI over MagicDNS)

## Goal
Make Clerkenstein auth complete from a **remote** browser, and serve the **whole browser surface under ONE clean,
trusted HTTPS origin** (`https://billion.taildc510.ts.net`) — the user's "one clean URL" decision (2026-07-11).
Real `tailscale cert` (Let's Encrypt, trusted tailnet-wide, no CA install), the fake FAPI + frontends fronted by a
lightweight reverse proxy, and a pk re-minted for the MagicDNS FAPI host so clerk-js reaches the tailnet — not the
viewer's loopback.

## Why section
Enumerable: swap the cert mint, stand up the proxy, re-mint the pk, guard the topology. The auth risk is real (see
Top risks / M215) but the *work list* is known — three facts already verified at design make it config-only (dotted
host validates; host-agnostic token verify; path-only cert mount).

## Access-scheme decision (2026-07-11) — HTTPS everywhere, one origin
Front the browser-facing ports on the MagicDNS host with the `tailscale cert`, so teammates hit a single
`https://<magicdns>` with a padlock, no mixed-content, and a **secure context** (Clerk's clerk-js uses Web Crypto,
which browsers grant only on `localhost` or HTTPS — a plain-`http://` MagicDNS origin is NOT a secure context, so
HTTPS on the app origin is effectively required for auth, not just cosmetic).

## Live foundation — PROVEN on billion (2026-07-11)
A live PoC on the target VM confirmed the load-bearing assumption **before build**: `sudo tailscale cert
billion.taildc510.ts.net` mints a real **Let's Encrypt** cert (CN=billion.taildc510.ts.net, issuer LE `YE1`, 90-day),
and a **remote** tailnet machine fetched an HTTPS endpoint served from that cert with **`ssl_verify_result=0`
(trusted, ZERO CA install)** in ~83 ms over the tailnet — served via `ssl.load_cert_chain(fapi.crt, fapi.key)`, the
**exact path-mounted-cert model the fake-FAPI's `ListenAndServeTLS` uses**. So the mkcert→`tailscale cert` swap is a
genuine drop-in, and a `0.0.0.0`-bound port is reachable over Tailscale (UFW allows the tailscale interface). Two
findings carried forward:
- **Cert renewal:** the LE cert is 90-day (expires ~Oct 9). `tailscale cert` re-issues on re-run — add a
  renew-then-reload-FAPI step for a long-lived stack.
- **RAM (→ M215, not a M213 blocker):** the constraint is the **odyssey host** (128 GB physical, ~180 GB configured
  across 13 VMs, ~94% used), NOT billion. billion is pinned to its 8 GB balloon floor; the only host-RAM reclaim
  lever (the balloon-OFF `ithaca`/`ailabs` VMs, 52 GB fixed) is **off-limits (billion-only, user constraint)**.
  Mitigated: a **16 GB swap net** added to billion (persistent, `swappiness=10`). M215 runs swap-backed / trimmed.

## Scope
- **In:**
  - **FAPI cert:** when `STACK_PUBLIC_HOST != localhost`, replace the mkcert/openssl mint (`up-injected.sh:630-659`)
    with `tailscale cert --cert-file $CERTS/fapi.crt --key-file $CERTS/fapi.key $HOST` — **keep the same output
    paths** so the path-only consumer (`gen_injected_override.py:355-379` mount + `cmd/fake-fapi/main.go:60-64`
    `ListenAndServeTLS`) is untouched. The idempotent keep-existing-cert step (3a-bis) makes pre-placing the cert a
    no-op.
  - **pk re-mint:** driven by M212's `--fapi-host`; **verify** the dotted MagicDNS host passes `@clerk/backend`
    `assertValidPublishableKey` (the `dotless-pk-rejected` DNA gene — a dotted name passes where `localhost`
    failed) and that the base64 length round-trips (`inject.py` self-checks `parse_pk==host`).
  - **Reverse proxy (HTTPS everywhere):** a lightweight proxy (`tailscale serve` or Caddy) fronting the
    browser-facing ports (app 3000, cosmo 5050, backend 8082, studio-desk 9000, fapi 5400, academy 3077 — all
    `+offset`) on the MagicDNS host with the `tailscale cert`, terminating TLS to one origin. Built + tested in the
    rext authoring copy, driven from `up-injected.sh` (gated on the opt-in flag).
  - **Topology guard:** keep the FAPI on the **same MagicDNS host** as the app (different port) so the
    `SameSite=Lax` handshake `Set-Cookie` (`clerk-frontend/server.go:207-211`, no Domain/Secure) stays same-site;
    verify `ts.net`/`taildc510.ts.net` PSL status **before** ever splitting FAPI onto a separate subdomain.
  - Confirm the M212 build-rebuild guard trips on a **HOST** change (stale-image protection); confirm the fake-fapi
    container has outbound egress to `cdn.jsdelivr.net` (it proxies clerk-js).
- **Out:** the knob plumbing (M212); CORS emission + link ejects + the patch tail + the corpus recipe (M214); the
  live cross-machine acceptance + RAM/renewal burn-down (M215).

## Depends / parallel
- **Depends on:** M212 (the `$HOST` + `--fapi-host` plumbing).
- **Parallel with:** **M214** — different files (cert/proxy/pk vs CORS/patches), no shared API. Merge is additive.

## Delivers → knowledge
Feeds `corpus/services/clerkenstein.md` (the tailscale-cert FAPI path) + the M214 `tailscale-serve.md` recipe (the
proxy topology). No standalone doc lands here.

## Open questions
- Proxy mechanism: `tailscale serve` (built-in, 443→one port — path routing for multi-port) vs Caddy (flexible
  port/subdomain routing). An implementation choice resolved at build; both use the `tailscale cert`.
- `@clerk/nextjs` `clerkMiddleware` (real SDK, server-side in next-web): any pk-host reachability check at boot, or
  purely request-time from the pk? (The networkless `CLERK_JWT_KEY` path + the `api.clerk.com`-alias JWKS fetch are
  container-internal + host-agnostic — expected fine; confirm in M215.)
- `tailscale cert` PEM shape vs Go `ListenAndServeTLS` + renewal for a long-lived stack (renewal → M215 burn-down).

## KB dependencies
`corpus/services/clerkenstein.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/demo/recipe-browser-login.md`,
`corpus/architecture/alignment_testing.md` (the `dotless-pk-rejected` gene).
