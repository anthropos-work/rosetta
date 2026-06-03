# Clerkenstein

**Status:** v0.2 (v1.0 "body double" · milestones M1 + M2) · **Last updated:** 2026-06-03
**Repo:** `anthropos-demo/clerkenstein` (gitignored demo scratchpad, its own git) · **Measured by:** the [alignment framework](../architecture/alignment_testing.md)

## Role

Clerkenstein is a **drop-in mock of the Clerk library** the platform uses — the *same interface*, with
all security and sync **disarmed**. It exists so **demo** environments can create users / orgs / admins
and log in/out with no Clerk friction (one universal credential, no live API, no webhooks, no rate
limits), while platform repos keep "thinking" they use Clerk with **zero source changes**.

It is the **first mirror produced by the M0 alignment process** (not a hand-built mock): its fidelity
is *measured* as a 0–100% alignment score against a Clerk **Alignment DNA**, and both milestones drove
their score to the gate — **100% critical / 100% overall** on the Go surface (22/22 genes, `clerk-sdk-go/v2
@ v2.6.0`, M1) **and** on the JS/FAPI surface (9/9 genes, `@clerk/clerk-js` v5 / `@clerk/nextjs` v6, M2).
The DNA(s) + mirror + goldens + runners live in the clerkenstein repo; the *measuring machinery*
([`test/alignment/`](../../test/alignment/) + `/align-dna` + `/align-run`) lives in rosetta.

## Architecture & code map

The mirror covers **three surfaces** of the platform's Clerk consumption: the two **Go** sides (M1) and
the **JS/browser** side (M2), plus the **webhook** sync path (M2):

### `authn/` — the `colony/authn` provider twin (offline)
Implements the real `colony/authn.Provider` interface (`GetUser(token)`, `GetUserByID(uuid)`,
`Name()`) — a **compile-time `var _ authn.Provider` assertion** guarantees the drop-in. Tokens are
**HS256-signed with one universal key** (`authn/jwt.go`); `GetUser` verifies + extracts the platform
claim set (`eid`→`ID`, `email`, `firstname`, `lastname`, `org.eid`→org `ID`, `org_id`→org `AuthID`,
`org_role`→`AuthRole`) into a `clerkUser` implementing `authn.{User,Organization}` (`authn/user.go`,
`authn/provider.go`). No live Clerk — JWT verify is local.

### `orgclient/` — the Clerk org/membership API twin (disarmed, in-memory)
A small in-memory store (`orgclient/store.go`, `orgclient/invitations.go`) reproducing the
success/error semantics of the 10 consumed methods (CreateOrganization, CreateMembership, ChangeRole,
DeleteOrganizationMembership, InviteMember, BulkInviteMembers, RevokeInvitation, the 3 metadata
writes). No network calls.

### `fapi/` — the fake Clerk **Frontend API** (M2, the browser side)
A concurrency-safe HTTP server (`fapi/server.go`) serving the minimal FAPI bootstrap surface
`@clerk/clerk-js` / `@clerk/nextjs` call to load Clerk and establish a session: `GET /v1/environment`,
`GET /v1/client`, sign-in/sign-up create + attempt, `POST …/sessions/{id}/tokens`, `GET /v1/me`,
`/.well-known/jwks.json`, sign-out — all disarmed (one universal credential, always-accepts). The
**session-token endpoint mints the same HS256 universal-key JWT** the `authn/` twin verifies
(`cauthn.Mint`), so the **browser session and the disarmed backend agree end to end**.
`fapi/key.go` is the publishable-key codec (`MintPublishableKey(host)` ↔ `ParsePublishableKey`), which
is how the browser is *pointed* at this server — see the JS path section.

### `bapi/` — the fake Clerk **Backend API** (M2, the `api.clerk.com` redirect target)
The HTTP server (`bapi/server.go`) that the platform's real `orgclient` hits when `api.clerk.com` is
redirected to it (M1-D2). It serves the Clerk-SDK wire shapes for the 10 consumed methods, **backed by
the M1 `orgclient` in-memory twin** as the store. Verified by pointing a *real* `clerk-sdk-go/v2`
client at it (`bapi/server_sdk_test.go`). The redirect recipe is `bapi/doc.go`.

### `webhook/` — the **webhook injector** (M2)
`webhook/injector.go` + `webhook/events.go`: synthesizes the 12 consumed Clerk event types, **svix-signs
them with the demo webhook secret**, and POSTs to `POST /api/webhook/clerk` — feeding the platform's
existing `app/internal/clerk/events` sync pipeline directly. It signs with the *same svix library* the
platform verifies with, so an injected event passes `m.wh.Verify` (the real signed-webhook contract).

### `cmd/clerkrun/` + `cmd/jsfapirun/` — the alignment runners
`clerkrun --target {source|mirror} --dna PATH` (Go surface) and `jsfapirun …` (JS/FAPI surface) each emit
the outcomes protocol `alignctl` consumes — the glue that lets the framework score the mirror. Exercised
end-to-end by every `alignctl run`.

## Injection (zero platform-code changes) — three mechanisms

| Side | Mechanism | Status |
|---|---|---|
| **authn** (Go) | `go.mod replace` the whole `colony` module with a Clerkenstein-patched colony (its `authn/provider/clerk` = the disarmed twin), made invisible upstream via **skip-worktree** — the exact pattern staging already uses for its `vendor-colony/` v2-JWT patch. | recipe documented; live wiring is demo-stack work (v1.1) |
| **orgclient** (Go) | **Different (M1-D2):** the orgclient is `app`-internal (`app/internal/clerk/orgclient`, not a published module) and calls `api.clerk.com` over HTTP — so it *can't* be `go.mod replace`d. Disarming it = redirect `api.clerk.com` → the **`bapi/` fake-Clerk-API-server** (DNS/`/etc/hosts` + a trusted cert; recipe in `bapi/doc.go`). | **built in M2** (server + recipe); live wiring is demo-stack work (v1.1) |
| **JS / browser** (clerk-js / nextjs) | A **publishable key** (env var: `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `VITE_CLERK_PUBLISHABLE_KEY`) minted to encode the **`fapi/` fake-FAPI host** — clerk-js derives its FAPI host *entirely* from the key (`@clerk/shared` `parsePublishableKey`). Optionally `proxyUrl`/`domain` props as backup. **Config, not a fork.** | **built + proven in M2** (M2-D1 spike) |
| **webhook sync** | The **`webhook/` injector** posts svix-signed Clerk events to `POST /api/webhook/clerk` — the real signed-webhook contract; no platform change. | **built in M2** |

The **alignment gate measures behavior**, which the in-memory/fake servers provide regardless of how
injection eventually wires in — so the gates fired without the live demo-stack wiring being in place.

## Disarmed-security properties (by design — speed + accessibility, not security)

These are deliberate, not bugs (a demo mock, never production):
- **One universal credential** — every token is HS256-signed/verified with a single fixed key. The
  `fapi/` session-token endpoint mints this *same* token, so the browser session and the backend agree.
- **JWT `alg` is not validated** — `parse` verifies the HMAC regardless of the header `alg` (real Clerk
  would reject a mismatch). Acceptable: the mock's job is to *accept* easily.
- **Tokens without `exp` never expire.**
- **The `fapi/` JWKS endpoint serves an empty key set** — clerk-js's bootstrap is satisfied by shape, but
  the disarmed backend verifies HS256 with the universal key, *not* this JWKS (cryptographic
  verification is deliberately disarmed).
- (M2-D2) The `orgclient` store is **now concurrency-safe** (mutex) — M1's adversarial review flagged the
  plain-map store as injection-time-unsafe; M2 made it safe because one `bapi/` instance serves a demo's
  concurrent orgclient traffic.

## JS browser session + webhook coherence (M2)

**The defining M2 risk — resolved in the strong direction (M2-D1).** The open question was whether
`@clerk/clerk-js ^5.52.3` / `@clerk/nextjs ^6.39.2` could be pointed at a fake FAPI without forking the
SDK. **Answer: yes, by configuration alone.** clerk-js derives its FAPI host **entirely from the
publishable key** — `pk_test_<base64("host$")>`, decoded by `@clerk/shared`'s `parsePublishableKey` — and
additionally honors `proxyUrl`/`domain`. So `fapi/key.go`'s `MintPublishableKey(host)` produces a key
that, set as `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `VITE_CLERK_PUBLISHABLE_KEY`, points the browser at the
`fapi/` server. **No SDK fork; no platform-code change.**

> **The decided fallback is not exercised.** state.md recorded a fallback (keep the real dev Clerk app for
> the browser while the backend stays mocked) in case the override proved too fragile. The spike proved it
> robust, so the fallback stays documented as the escape hatch only — if keeping the fake FAPI faithful
> across a clerk-js bump ever proves too costly (M1b-style drift would flag it), it is the fallback.

**End-to-end coherence chain:** browser loads clerk-js → (publishable key) → `fapi/` fake FAPI → sign-in
(universal credential) → session-token endpoint mints the **HS256 universal-key JWT** → used as the API
bearer → the disarmed backend `authn/` twin verifies that exact token. Measured by the
`clerk-js-5` DNA (9 genes, 100%/100%), with the `SessionToken/decoded-identity` gene (operator `exact`)
pinning that the browser-minted token decodes to the exact platform identity.

**Webhook coherence:** demo-created/seeded users/orgs never trigger real Clerk webhooks, so the
`webhook/` injector replays the 12 consumed event types (svix-signed with the demo `CLERK_WEBHOOK_SECRET`)
into `POST /api/webhook/clerk`, driving the platform's existing `events.Manager.Handle` sync pipeline →
Postgres + Sentinel stay coherent with the demo's identities.

## Local development

```sh
# in anthropos-demo/clerkenstein (builds offline against the cached colony + clerk-sdk-go + svix):
GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go test ./...          # all packages (authn/orgclient/fapi/bapi/webhook)
GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go build -o clerkrun ./cmd/clerkrun      # Go-surface runner
GOFLAGS=-mod=mod GOPROXY=off GOSUMDB=off go build -o jsfapirun ./cmd/jsfapirun    # JS-surface runner

# measure fidelity from rosetta (the alignment gates):
cd test/alignment
go run ./cmd/alignctl run --dna <…>/clerkenstein/dna/clerk-2.6.0.json \
  --runner <…>/clerkenstein/clerkrun  --golden-dir <…>/clerkenstein/golden    --gate-overall 95 --gate-critical 100
go run ./cmd/alignctl run --dna <…>/clerkenstein/dna/clerk-js-5.json \
  --runner <…>/clerkenstein/jsfapirun --golden-dir <…>/clerkenstein/golden-js --gate-overall 95 --gate-critical 100
```

## Testing

- **Unit:** `authn` + `orgclient` at **100%**; `fapi` 99%, `bapi` 96%, `webhook` 91% — all race-clean.
  Highlights: the browser-minted FAPI token is backend-verifiable (`fapi`); a *real* `clerk-sdk-go/v2`
  client parses every `bapi` response; the injector's svix signature passes the platform's `svix.Verify`
  (`webhook`). The runners (`cmd/clerkrun`, `cmd/jsfapirun`) are integration-covered by the alignment run.
- **Alignment:** two gates — `alignctl run` reports **100%/100%** over the 22-gene `clerk@2.6.0` Go DNA
  (M1) **and** the 9-gene `clerk-js-5` JS/FAPI DNA (M2). These are the exit criteria and the regression
  signal **M1b** CI-gates across Clerk version bumps (re-`/align-dna` the new version, re-`/align-run`).

## Drift detection (M1b)

Clerk moves; the mirror must stay aligned. M1b makes a Clerk bump a **flagged, mechanical event** by
reusing M0 wholesale — no new measurement machinery, just two scripts + a CI gate (in the clerkenstein
repo):

- **`scripts/gate.sh`** — the alignment gate: builds the runner + a fresh `alignctl`, runs
  `alignctl run --gate-overall 95 --gate-critical 100`. **Exit 0** = gate met, **2** = the mirror
  regressed. (Uses a built `alignctl` binary, not `go run`, so the exact exit code propagates.)
- **`scripts/drift-check.sh --new <bumped-DNA>`** — wraps the gate with a DNA-diff step. **Exit codes:**
  **0** no drift & gate met · **1** the DNA moved (`alignctl dna diff` shows added/removed/changed
  genes — the Clerk surface changed) · **2** the gate regressed (genes broke).
- **`.github/workflows/alignment.yml`** — runs the gate on push + a **weekly** schedule (the brief's
  "follow platform updates within minutes" cadence), turning a Clerk break into a red build.

**The bump runbook** (when `clerk-sdk-go` / `@clerk/*` updates):
1. `/align-dna` the new version → a new DNA. `drift-check.sh --new …` reports what moved (exit 1).
2. Re-author the changed/added genes in the DNA; **re-capture goldens** for the moved genes
   (`alignctl capture`, or hand-author per the hybrid M1-D1 path).
3. `/align-run` (or `gate.sh`) re-scores the mirror against the bumped surface; close any new
   divergences in the twin until the gate is green again.
4. Re-pin the DNA version. The weekly CI keeps it honest between bumps.

`ALIGN_DIR` (default `../../test/alignment`) locates rosetta's `alignctl`. Verified across all exit
paths against a simulated `clerk@2.7.0` bump.

> **M2 note:** M1b's `gate.sh`/`drift-check.sh` parameterize the DNA/runner/golden-dir, so the **JS
> surface** is gated the same way — point them at `dna/clerk-js-5.json` + `jsfapirun` + `golden-js`. A
> `@clerk/clerk-js` / `@clerk/nextjs` bump runs the identical bump runbook above against the JS DNA.

## See also
- [Alignment Testing](../architecture/alignment_testing.md) — the framework that measures this mirror.
- [Clerk integration](clerk-integration.md) — the real Clerk surface Clerkenstein mirrors.
- [Frontend architecture](../architecture/frontend_architecture.md) · [next-web-app](next-web-app.md) — the `@clerk/nextjs` consumers the `fapi/` server stands in for.
- [Webhook setup](../ops/webhook_setup.md) — the real Clerk webhook path the `webhook/` injector replays into.
