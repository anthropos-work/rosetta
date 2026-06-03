---
milestone: M2
slug: browser-webhook-coherence
version: v1.0 "body double"
milestone_shape: section
status: in-progress
started: 2026-06-03
last_updated: 2026-06-03
---

# M2 — Clerkenstein: browser session + webhook coherence (JS)

## Goal
Close the two remaining Clerk seams so a demo stack is **Clerk-free end to end**, not just on the Go
backend (M1):
1. **Browser session** — the frontends (`next-web-app`, `ant-academy` via `@clerk/nextjs ^6.39.2`;
   `studio-desk` via `@clerk/clerk-js ^5.52.3`) log in with **no real Clerk**, by pointing the SDK's
   Frontend API (FAPI) at a fake FAPI server.
2. **Webhook coherence** — users/orgs/memberships created or seeded in a demo reach the platform DB
   through the **same** `app/internal/clerk/events/` sync pipeline the real Clerk webhooks drive, via a
   **webhook injector** that posts correctly-signed Clerk events to `POST /api/webhook/clerk`.
3. **(routed from M1 — M1-D2, Fate 3)** the **fake-Clerk-API-server** (HTTP interception of
   `api.clerk.com`) that ALSO disarms M1's Go `orgclient` — app-internal + networked, so it can't
   `go.mod replace` like authn. The Clerkenstein orgclient *behavior* already scores 100% (M1); M2
   wires the HTTP redirect that makes the platform's real orgclient hit the mock.

This is the **highest-technical-risk milestone in v1.0**: the SDKs are widely believed to hard-code the
Clerk FAPI with no documented base-URL override. The spike (§ S1, decisions.md M2-D1) resolved this
**before any code** — see "Spike outcome" below.

## Spike outcome (the defining open question — RESOLVED Fate 1)
**Question (roadmap):** can `@clerk/nextjs` / `@clerk/clerk-js` be pointed at a fake FAPI without
forking the SDK? **Answer: yes, via configuration only — no fork, no fallback needed for the JS axis.**
clerk-js derives its FAPI host **entirely from the publishable key** (`pk_test_<base64("host$")>`,
decoded in `@clerk/shared` `parsePublishableKey`); it additionally honors `proxyUrl` / `domain`
(satellite) options. So pointing the browser at a fake FAPI is an **env-var change**
(`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `VITE_CLERK_PUBLISHABLE_KEY`) — demo-env config, **not a
platform-code change**. The decided real-dev-Clerk fallback (state.md) is therefore **not exercised**;
it remains documented as the escape hatch if the fake-FAPI server proves too costly to keep faithful.
Full evidence in `spec-notes.md` § "Spike: JS FAPI override".

## Context
M1 shipped the Go mirror (authn + orgclient) at 100%/100% alignment, offline. M1's gate measured
*behavior*; two injection seams were explicitly routed forward (M1 overview "Out" + M1-D2): the JS
browser surface and the fake-API-server. M2 builds both, plus the webhook injector. Per the M0/M1
pattern, the **fake FAPI's fidelity is expressed as alignment genes** (a JS-surface DNA) and measured
by `alignctl`, exactly like the Go side — so "the browser can't tell the difference" becomes a number.

## Scope
### In
- **S1 — JS FAPI spike + fake FAPI server.** Resolve the override question (done — Fate 1, above), then
  a **fake Clerk FAPI server** in the `clerkenstein` repo: the minimal FAPI endpoint set clerk-js calls
  to establish a session (`/v1/environment`, `/v1/client`, sign-in/sign-up create+attempt, `/v1/me`,
  handshake/JWKS), returning disarmed-but-shaped responses (one universal credential). Plus the
  **publishable-key minting** helper that encodes the fake FAPI host.
- **S2 — fake-Clerk-API-server (BAPI, `api.clerk.com`) for orgclient injection (M1-D2).** The
  HTTP-interception server serving the 10 orgclient methods' BAPI responses, reusing the M1 Clerkenstein
  `orgclient` in-memory twin as the backing store. Plus the **redirect recipe** (DNS/`/etc/hosts` +
  base-URL) that makes the platform's real orgclient hit it — **zero platform-code changes**. Addresses
  the M1-D2 thread-safety note (the in-memory store must be concurrency-safe when one server instance
  serves demo traffic).
- **S3 — webhook injector.** A tool that synthesizes Clerk webhook events (the 12 consumed event types:
  `user.{created,updated,deleted}`, `organization.{created,updated,deleted}`,
  `organizationMembership.{created,updated,deleted}`, `organizationInvitation.{accepted,created,revoked}`),
  **svix-signs** them with the demo webhook secret, and POSTs to `POST /api/webhook/clerk` — feeding the
  existing `events.Manager.Handle` pipeline directly. **Zero platform-code changes** (uses the real
  signed-webhook contract).
- **S4 — JS-surface Alignment DNA + genes.** Express the fake FAPI's fidelity as a JS/FAPI Alignment DNA
  (`clerk-js@<ver>`) + a runner target, scored by `alignctl` like the Go DNA, so the browser-coherence
  claim is measured, not asserted.
- **S5 — documentation.** Extend `corpus/services/clerkenstein.md` with the JS path + fake-FAPI +
  webhook injection + the spike outcome + the (un-exercised) fallback; cross-ref
  `frontend_architecture.md`, `next-web-app.md`, `webhook_setup.md`, `alignment_testing.md`.

### Out
- Multi-instance disposable stacks → M3 (v1.1).
- Data seeding / use-case recipes → M4 (v1.1).
- A live end-to-end browser click-through against a running platform stack (needs a full demo stack,
  which is M3) — M2 verifies the fake FAPI against the SDK's request contract + the alignment genes,
  not against a live multi-service stack.

## Sections
See `progress.md`: S1 (JS FAPI spike + fake FAPI server) · S2 (BAPI fake-API-server + orgclient
redirect) · S3 (webhook injector) · S4 (JS-surface DNA + genes) · S5 (docs).

## Zero-platform-code-changes invariant
Every seam M2 builds is **config / external-process**, never a platform edit:
- JS FAPI → publishable-key (env var) and/or `proxyUrl`/`domain` ClerkProvider props.
- orgclient BAPI → DNS/base-URL redirect to the fake server (the SDK's `URL` default is `api.clerk.com`;
  the platform never overrides it in code, so redirect is external).
- webhook → the real signed `POST /api/webhook/clerk` contract.
The platform repos under `anthropos-dev/` are **read-only inspection only**.

## KB dependencies (contract)
`corpus/architecture/alignment_testing.md`, `corpus/services/clerk-integration.md`,
`corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`,
`corpus/ops/webhook_setup.md`, `corpus/services/clerkenstein.md` (extended here).

## Delivers → knowledge
- `corpus/services/clerkenstein.md` — extended with the JS path, fake FAPI server, BAPI fake-API-server
  + orgclient redirect, the webhook injector, the JS-surface DNA, and the spike outcome + fallback.

## Where it lives
The fake FAPI server, the BAPI fake-API-server, the webhook injector, the JS-surface DNA + goldens + the
runner target all live in the **`clerkenstein` repo** (gitignored `anthropos-demo/`, its own git). This
milestone's **planning/section records + the corpus docs** live in rosetta on `m2/browser-webhook-coherence`.
