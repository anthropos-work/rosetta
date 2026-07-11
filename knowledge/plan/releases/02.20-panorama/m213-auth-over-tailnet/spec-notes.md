# M213 — spec-notes

The auth + TLS surface (from `wf_bea3be47`, file:line). All in rext tooling.

## Why auth is config-only (three verified facts)
1. **Dotted-host validator:** `@clerk/backend` requires a DOT in the decoded pk host (the reason `127.0.0.1` was
   chosen over dotless `localhost`). `billion.taildc510.ts.net` is dotted → validates NATIVELY. Gated by the
   `dotless-pk-rejected` DNA gene.
2. **Host-agnostic token verify:** `clerkenstein/shared/jwt.go:63-140` (Parse/ParseRS256/ParseAny) checks ONLY
   signature+expiry — no `iss`/host/`azp`/`authorizedParties`; the Claims struct (`:32-43`) has no `iss`. Re-hosting
   the FAPI does NOT break verification in any of the 4 Go services or the Node path.
3. **Path-only cert mount:** fake-fapi reads `FAKE_FAPI_TLS_CERT=/certs/fapi.crt` + `KEY=/certs/fapi.key` and calls
   `ListenAndServeTLS` (`cmd/fake-fapi/main.go:60-64`; mount `gen_injected_override.py:355-379`). Any valid cert at
   that path just works.

## Cert
- Mint today: `up-injected.sh:636` (openssl `SAN=IP:127.0.0.1,DNS:localhost`) / `:649` (mkcert `127.0.0.1 localhost
  ::1`). mkcert is trusted ONLY on the bring-up host (`recipe-browser-login.md:72-75`).
- Swap → `tailscale cert --cert-file $CERTS/fapi.crt --key-file $CERTS/fapi.key $HOST`. Cert matches on hostname
  (port is not part of matching), so one cert covers the FAPI on any offset port.

## pk
- `pk_test_<RawStdBase64(host+"$")>`; codec `clerk-frontend/key.go:28-37`; mint `inject.py:26-37,60-68`
  (self-check `parse_pk==host`). Baked into next-web `.env.local` (`up-injected.sh:178`) + studio-desk build-arg
  (`:294`) + `DESK_CLERK_PUBLISHABLE_KEY` (`:610`) → re-mint forces both frontend rebuilds.

## Cookies / topology
- Handshake `Set-Cookie` (`server.go:207-211`): `Path=/; SameSite=Lax`, no Domain, no Secure. Over HTTPS on a real
  host: works IF FAPI is same host as app (different port = same-site). Do NOT split FAPI to a separate `*.ts.net`
  subdomain without `SameSite=None; Secure` (check PSL first).

## Reverse proxy (HTTPS everywhere)
- Front browser-facing ports {3000,5050,8082,9000,5400,3077}+offset on `$HOST` with the tailscale cert → one origin.
- Egress: fake-fapi proxies clerk-js from `cdn.jsdelivr.net` (`server.go:154-169`) — VM needs outbound HTTPS.
