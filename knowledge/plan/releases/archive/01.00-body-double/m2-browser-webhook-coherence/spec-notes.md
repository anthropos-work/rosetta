# M2 — spec notes

## Pre-flight audits — S1 (KB-fidelity, 2026-06-03)
**Phase 0b — GREEN.** Topic → doc → code triples (all PAIRED, claims ALIGNED):
- JS FAPI / publishable key → `clerk-integration.md` § Configuration + `next-web-app.md` (`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`) + `frontend_architecture.md` (Bearer injection) → `studio-desk/node_modules/@clerk/{clerk-js,shared}`, `next-web-app/apps/web/src/middleware.ts`. Verified vs the installed SDK (spike below).
- orgclient BAPI → `clerk-integration.md` § "Backend API" + `clerkenstein.md` § Injection (M1-D2) → `app/internal/clerk/orgclient/clerk.go` (`const clerkApiUrl`, `clerk.ClientConfig` w/o `URL`).
- webhook injector → `clerk-integration.md` § "Org/role sync" + `webhook_setup.md` (`/api/webhook/clerk`, `whsec_`) → `app/internal/clerk/events/{events,types}.go`, `app/internal/web/backend/api/api.go` `HandleClerkWebhookEvent`.
- alignment DNA (JS surface) → `alignment_testing.md` (gene/DNA model, `dna validate` id rules) → `test/alignment/`, `clerkenstein/dna/clerk-2.6.0.json`.

**1 stale claim found + fixed inline (non-load-bearing):** `clerk-integration.md` claimed `app` on
`clerk-sdk-go/v2 v2.5.1`; actual `app/go.mod` = `v2.6.0` (matches colony + the M1 DNA). Corrected. No
blind areas; no load-bearing staleness. Report: `kb-fidelity-audit.md`.

## Spike: JS FAPI override (the defining risk — RESOLVED Fate 1, no fork, no fallback)

Inspected the **installed** `@clerk/clerk-js` (`anthropos-dev/studio-desk/node_modules/@clerk/clerk-js`
v5.125.10 — `^5.52.3` resolved up) + `@clerk/shared` + `@clerk/types`. Findings:

1. **FAPI host is derived from the publishable key.** `@clerk/shared`'s `parsePublishableKey(key)`
   decodes the host: `instanceType = key.startsWith("pk_live_") ? production : development`, then
   `decodedFrontendApi = isomorphicAtob(key.split("_")[2])` → base64 of `"<frontendApiHost>$"`. So the
   FAPI host the browser talks to is **whatever the publishable key encodes**. Minting a key that
   encodes `fapi.localhost:<port>` (or any host) points clerk-js at that host — no SDK change.
   - `apiUrlFromPublishableKey` (BAPI side) classifies the host by suffix (`*.lcl.dev` → LOCAL_API_URL,
     `*.clerk.accounts.dev` → PROD, etc.) — informative for which suffix to encode, but the **frontend**
     FAPI host is the decoded host itself.
2. **`proxyUrl` / `domain` are first-class options.** The clerk-js browser bundle builds its FAPI client
   honoring `proxyUrl` (`getFapiClient` uses `this.proxyUrl` when set) and `domain` (satellite). These
   are belt-and-suspenders overrides exposed as `ClerkProvider` props / `Clerk.load` options — also
   config, not code.
3. **`@clerk/nextjs ^6.39.2`** (next-web-app, ant-academy) wraps the same clerk-js + a server side that
   reads `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `CLERK_SECRET_KEY` / optional `CLERK_PROXY_URL` /
   `NEXT_PUBLIC_CLERK_DOMAIN` from env — all the same override surface, set via env.

**Conclusion:** pointing the browser at a fake FAPI is a **demo-env configuration change** (an env var:
`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `VITE_CLERK_PUBLISHABLE_KEY`, optionally `CLERK_PROXY_URL`), never
a platform-code or SDK-fork change. The state.md "real dev Clerk app for the browser" **fallback is not
needed** and is **not exercised**; it stays documented as the escape hatch.

### FAPI endpoint set clerk-js calls to bootstrap a session (the fake FAPI's surface)
The minimal set to load Clerk + establish a session (dev instance):
- `GET  /v1/environment` — instance config (auth config, display config, user settings). Drives whether
  the SDK considers the instance loaded.
- `GET  /v1/client` — the current Client (sessions, sign_in/sign_up state). Empty client on first load.
- `POST /v1/client/sign_ins` (+ `/v1/client/sign_ins/{id}/attempt_first_factor`) — sign-in create +
  attempt. The disarmed fake accepts the one universal credential.
- `POST /v1/client/sign_ups` (+ `…/attempt`) — sign-up create + attempt (mirror of sign-in).
- `GET  /v1/me` (a.k.a. the active session's user) — returns the universal demo user.
- `GET  /.well-known/jwks.json` / handshake — for token verification on the consuming side (must serve
  a JWKS whose key matches the one Clerkenstein's authn twin mints with — ties S1 to M1's authn twin).
- session token mint (`POST /v1/client/sessions/{id}/tokens`) — returns the **same HS256 universal-key
  JWT shape** M1's authn twin verifies, so the browser session and the backend authn agree end to end.

Operator note: most FAPI responses carry generated ids/timestamps → `shape`/`normalized` genes; the
auth-config/environment booleans → `exact`; bad-credential paths → `error_class`.

## BAPI (`api.clerk.com`) interception — the orgclient seam (M1-D2)
- The platform's `orgclient.New` (in `app/internal/clerk/events/events.go`) builds a
  `clerk.ClientConfig` with only `config.Key` set; every SDK sub-client (`clerkOrg.NewClient(config)`,
  membership, invitation) calls `clerk.NewBackend(&config.BackendConfig)`, whose `URL` defaults to
  `https://api.clerk.com/v1` (`clerk-sdk-go/v2 clerk.go` `APIURL`). The platform **never sets**
  `BackendConfig.URL` → so the only override is **external HTTP redirect**, exactly M1-D2.
- The two raw-HTTP methods (`BulkInviteMembers`, `UpdateMembershipMetadata`, `UpdateUserMetadata`) use a
  hardcoded `const clerkApiUrl = "https://api.clerk.com/v1"` → same redirect catches them.
- **Redirect recipe (zero platform-code):** `/etc/hosts` (or container DNS) maps `api.clerk.com` → the
  fake server; the fake server terminates TLS with a cert the demo trusts (or the demo points at an
  HTTP base via a TLS-skip in the demo network). The fake server serves the BAPI shapes the SDK
  unmarshals, backed by the M1 Clerkenstein `orgclient` in-memory twin.
- **Thread-safety (M1-D2 / M1 adversarial note):** the M1 in-memory `orgclient.Store` uses plain maps,
  fine for the single-threaded alignment runner but **not** for a server instance serving concurrent
  demo requests. S2 must make the store concurrency-safe (mutex) — this is M2's injection-time concern
  that M1 explicitly routed here.

## Webhook injector — the sync-pipeline seam
- Entry point: `POST /api/webhook/clerk` → `API.HandleClerkWebhookEvent` → `ClerkEventsManager.Handle`
  (`app/internal/clerk/events/events.go`). `Handle` **first** calls `m.wh.Verify(payload, headers)`
  (svix) — so the injector MUST svix-sign payloads with the demo's webhook secret (`whSecret`, from the
  platform `.env` — a known value in a demo). After verification it switches on `event.Type`.
- The 12 consumed event types + their payload shapes are in `app/internal/clerk/events/types.go`
  (`clerkUserEvent`, `organizationEvent`, `organizationInvitationEvent`,
  `organizationMembershipEvent`). The injector synthesizes these shapes (only the fields the handlers
  read need to be correct — e.g. user `external_id`/`first_name`/`last_name`/`email_addresses` +
  `primary_email_address_id`; org `public_metadata.eid` + `created_by`; membership `organization` +
  `public_user_data` + `role`).
- Svix signature contract: `svix-id`, `svix-timestamp`, `svix-signature` headers; signature =
  base64(HMAC-SHA256(secret_bytes, `{id}.{timestamp}.{payload}`)) with the `whsec_` secret. The injector
  reproduces this so `m.wh.Verify` passes — the real signed-webhook contract, no platform change.

## JS-surface Alignment DNA (S4)
Mirror the Go DNA pattern (`clerkenstein/dna/clerk-2.6.0.json`): a `clerk-js@<ver>.json` genome whose
capabilities are the FAPI bootstrap operations (environment/client/sign-in/me/token) with variants
(loaded instance, universal-credential sign-in success, bad-credential error, session token shape). The
runner gains a `--surface js` (or a sibling runner) target emitting the outcomes protocol; `alignctl
run` scores it like the Go side. Goldens are **offline-capturable** (the SDK's expected request/response
shapes are in `@clerk/types` + the FAPI response schemas), keeping the hybrid-M1-D1 offline posture.

## Offline build/test invariants (carried from M1/clerkenstein.md)
- Go (fake servers + injector + runner): `GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go test ./...` /
  `go build`. Alignment gate: `cd test/alignment && go run ./cmd/alignctl run --dna … --runner … --gate-*`.
- The clerkenstein repo is its own git under gitignored `anthropos-demo/`; M2 code commits there, M2
  records + corpus docs commit to rosetta `m2/browser-webhook-coherence`.
