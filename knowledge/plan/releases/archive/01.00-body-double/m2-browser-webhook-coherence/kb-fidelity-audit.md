---
title: "KB Fidelity Audit — M2 browser-webhook-coherence"
date: 2026-06-03
scope: milestone:M2
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| JS FAPI / publishable-key override | `corpus/services/clerk-integration.md` § Config; `corpus/services/next-web-app.md`; `corpus/architecture/frontend_architecture.md` | `anthropos-dev/studio-desk/node_modules/@clerk/{clerk-js,shared,types}`; `next-web-app/apps/web/src/middleware.ts`; `*/package.json` (`@clerk/nextjs ^6.39.2`, `@clerk/clerk-js ^5.52.3`) | PAIRED |
| orgclient BAPI (`api.clerk.com`) | `corpus/services/clerk-integration.md` § Backend API; `corpus/services/clerkenstein.md` § Injection (M1-D2) | `app/internal/clerk/orgclient/clerk.go`; `app/internal/clerk/events/events.go` | PAIRED |
| Webhook sync pipeline | `corpus/services/clerk-integration.md` § Org/role sync; `corpus/ops/webhook_setup.md` | `app/internal/clerk/events/{events,types,user_handler,org_handler}.go`; `app/internal/web/backend/api/api.go` (`HandleClerkWebhookEvent`) | PAIRED |
| Alignment DNA / gene model (JS surface) | `corpus/architecture/alignment_testing.md` | `test/alignment/`; `anthropos-demo/clerkenstein/{dna,cmd/clerkrun}` | PAIRED |
| Clerkenstein mirror (M1, extended here) | `corpus/services/clerkenstein.md` | `anthropos-demo/clerkenstein/{authn,orgclient}` | PAIRED |

No BLIND-AREA, CODE-ONLY, or DOC-ONLY rows. The one net-new doc M2 produces is the **extension** of the
existing `clerkenstein.md` (Delivers line in overview) — not a missing precondition.

## Fidelity Findings

1. **clerk-integration.md — `app` Clerk Go SDK version** — STALE (fixed).
   - Source: `clerk-integration.md` § Dependent Repos (SDK-versions callout).
   - Expected (doc): `app` on `clerk-sdk-go/v2 v2.5.1`, drifted from colony's `v2.6.0`.
   - Actual (code): `anthropos-dev/app/go.mod` = `clerk-sdk-go/v2 v2.6.0` — aligned with colony + the M1 DNA target.
   - Verdict: STALE → **doc updated** (the drift is resolved; doc now says both on v2.6.0).
   - Severity: **non-load-bearing for M2** — M2's BAPI interception is HTTP-level + version-agnostic; if anything the resolution helps (the orgclient redirect aligns against the same v2.6.0 the DNA targets).

2. **clerk-integration.md — "Webhooks (svix, 12 event types)"** — ALIGNED.
   - 12 `case` arms in `events.go` `Handle`; types match (`user.{created,updated,deleted}`, `organization.{created,updated,deleted}`, `organizationMembership.{created,updated,deleted}`, `organizationInvitation.{accepted,created,revoked}`).

3. **webhook_setup.md — endpoint + secret contract** — ALIGNED.
   - `/api/webhook/clerk` (route registered in `server.gen.go`); `CLERK_WEBHOOK_SECRET=whsec_…` (svix); `received event from clerk` log line — all present in code.

4. **clerk-integration.md — publishable-key env shapes** — ALIGNED.
   - `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` / `VITE_CLERK_PUBLISHABLE_KEY` — match the spike's override seam.

5. **alignment_testing.md — DNA/gene/validate-id model** — ALIGNED.
   - PascalCase capability / kebab-case variant id constraint (`dna validate`) matches the existing `clerk-2.6.0.json` and constrains S4's JS DNA.

## Completeness Gaps
None critical. The fake-FAPI endpoint set + svix-signature contract + publishable-key decode are M2's to
**produce** (in clerkenstein + the extended clerkenstein.md), not pre-existing undocumented platform
behaviors. They are captured in `spec-notes.md` so the build has its contract.

## Applied Fixes
- `corpus/services/clerk-integration.md` — corrected the stale `app` SDK-version claim (v2.5.1 → v2.6.0).
- `spec-notes.md` — recorded the topic→doc→code triples under "Pre-flight audits — S1".

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed to Phase 1.
