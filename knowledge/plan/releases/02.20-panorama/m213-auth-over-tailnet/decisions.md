# M213 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-SCHEME-1 — HTTPS everywhere, one MagicDNS origin
Front the whole browser surface with the `tailscale cert` via a reverse proxy; teammates hit one `https://<magicdns>`.
**Why:** user decision (2026-07-11) + Clerk clerk-js needs a secure context (Web Crypto), which a plain-`http://`
MagicDNS origin is not — so HTTPS on the app origin is effectively required, not cosmetic.

## D-CERT-1 — tailscale cert, not mkcert, for the remote case
`tailscale cert` (Let's Encrypt) is trusted tailnet-wide with no per-machine CA install — precisely what mkcert
cannot give a remote browser. Consumer mount is path-only, so it is a drop-in at the same `/certs/fapi.{crt,key}`.
