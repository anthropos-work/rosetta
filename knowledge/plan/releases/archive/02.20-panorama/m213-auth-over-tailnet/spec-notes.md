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
  host: works IF FAPI is same host as app (different port = same-origin for a host-only cookie).
- **PSL — VERIFIED 2026-07-11** (fetched `publicsuffix.org/list/public_suffix_list.dat`): the entry is **`ts.net`**
  (+ `*.c.ts.net`), NOT `*.ts.net`. So **`ts.net` is the public suffix** and **`taildc510.ts.net` is the
  registrable domain (eTLD+1)**. Correction to the pre-verification note ("without SameSite=None; Secure"): two
  subdomains under `taildc510.ts.net` (e.g. `billion.` vs `fapi.`) are **SAME-SITE** (same eTLD+1 → `SameSite=Lax`
  is already satisfied — `SameSite=None; Secure` is NOT the blocker). The real blocker to a subdomain split is the
  **HOST-ONLY** cookie (no `Domain` ⇒ exact-host scoped): it would not cross to a different subdomain unless the
  disarmed handshake cookie added `Domain=taildc510.ts.net`. **Same-host (different port) sidesteps all of it** —
  what M213 does + the topology guard enforces (`up-injected.sh`: `[ "$HOST" != "$FAPI_HOST" ]` ⇒ exit 1 on the
  public path; equal-by-construction today, a regression tripwire for a future split).

## Reverse proxy (HTTPS everywhere)
- Front browser-facing ports {3000,5050,8082,9000,5400,3077}+offset on `$HOST` with the tailscale cert → one origin.
- Egress: fake-fapi proxies clerk-js from `cdn.jsdelivr.net` (`server.go:154-169`) — VM needs outbound HTTPS.

## Pre-flight audits — M213 (all sections)
- **Phase 0b `/developer-kit:audit-kb-fidelity --milestone=M213`: YELLOW → PROCEED** (2026-07-11). No blind area,
  no stale load-bearing claim. The three load-bearing premises are CODE-VERIFIED aligned: dotted-host validates
  (`up-injected.sh:59` FAPI_HOST default `127.0.0.1`; codec `key.go:28-37`), host-agnostic token verify
  (`shared/jwt.go:63-140`, `Claims` carries no `iss`/`azp`), path-only cert mount (`cmd/fake-fapi/main.go:60-64`
  + mount `gen_injected_override.py:366-371`). Five findings tracked as KB-1..KB-5 in `decisions.md` (Document-phase).
  Report: returned inline by the audit subagent (no file written, per the read-only run).

## Line-anchor drift correction (KB-5)
Every `up-injected.sh:<N>` anchor in this file + `overview.md` drifted ~+22 lines when M212's host-knob block
landed at `:16-90`. Corrected symbol anchors (drift-proof):
- cert mint: `up-injected.sh::gen_openssl_fapi_cert` (openssl fn) + the mkcert branch inside the
  `if [ ! -f "$CERTS/fapi.crt" ]` block (was `:636`/`:649`, actual `:658`/`:672` pre-M213).
- pk bakes: `up-injected.sh` next-web `.env.local` bake (was `:178`, actual `:200`); studio-desk build-arg
  (was `:294`, actual `:316`); `DESK_CLERK_PUBLISHABLE_KEY` (was `:610`, actual `:633`).
- cert mount consumer: `gen_injected_override.py` fapi_env block (was `:355-379`, actual `:366-385`).
Anchors into `server.go`/`jwt.go`/`key.go`/`inject.py` verified still-aligned (no lines gained above them).
