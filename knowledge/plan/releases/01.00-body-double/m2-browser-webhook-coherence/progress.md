# M2 — progress

Section milestone. 5 sections. Code → `clerkenstein` repo (gitignored `anthropos-demo/`); records +
docs → rosetta `m2/browser-webhook-coherence`. Zero platform-code changes throughout.

## S1 — JS FAPI spike + fake FAPI server — TODO
- [ ] **Spike (DONE at scaffold):** resolve whether `@clerk/*` can point at a fake FAPI without a fork.
      → YES, via publishable-key/`proxyUrl` config (M2-D1, spec-notes § Spike). Recorded.
- [ ] Publishable-key minting helper: encode an arbitrary FAPI host into a `pk_test_…` key.
- [ ] Fake Clerk FAPI server (clerkenstein repo): the bootstrap endpoint set
      (`/v1/environment`, `/v1/client`, sign-in create+attempt, sign-up create+attempt, `/v1/me`,
      session-token mint, `/.well-known/jwks.json`) returning disarmed-but-shaped responses with the
      one universal credential. JWKS/token shape ties to M1's authn twin (same universal key).
- [ ] Unit tests (1:1 per source file) + an integration test for the full sign-in bootstrap flow.

## S2 — fake-Clerk-API-server (BAPI) + orgclient redirect (M1-D2) — TODO
- [ ] BAPI fake-API-server serving the 10 orgclient methods' response shapes, backed by the M1
      Clerkenstein `orgclient` twin as the store.
- [ ] Make the `orgclient.Store` concurrency-safe (mutex) + a concurrency regression test (M2-D2).
- [ ] The `api.clerk.com` redirect recipe (DNS/`/etc/hosts` + TLS posture) — documented + verified
      against the SDK's request contract. **Zero platform-code changes.**
- [ ] Integration test: a `clerk-sdk-go/v2` client pointed at the fake server exercises each method.

## S3 — webhook injector — TODO
- [ ] Injector tool: synthesize + **svix-sign** + POST the 12 consumed Clerk event types to
      `POST /api/webhook/clerk` (the real signed-webhook contract → `events.Manager.Handle`).
- [ ] Payload builders for `clerkUserEvent` / `organizationEvent` / `organizationInvitationEvent` /
      `organizationMembershipEvent` (only handler-read fields need fidelity).
- [ ] Tests: svix-signature correctness (verifies against the svix lib) + per-event-type payload shape;
      an integration test proving `events.Manager.Handle` accepts an injected event (verify path).

## S4 — JS-surface Alignment DNA + genes — TODO
- [ ] Author `clerk-js@<ver>.json` DNA (FAPI bootstrap capabilities × variants) via the M0 `/align-dna`
      pattern; offline goldens (hybrid-M1-D1 posture).
- [ ] Runner target for the JS surface (emit the outcomes protocol) + `alignctl run` scores it.
- [ ] Gate the JS surface (critical FAPI genes at 100%, overall ≥95%) — same treatment as the Go side.

## S5 — documentation — TODO
- [ ] Extend `corpus/services/clerkenstein.md`: JS path, fake FAPI, BAPI fake-API-server + redirect,
      webhook injector, JS-surface DNA, the spike outcome + the un-exercised fallback.
- [ ] Cross-refs: `frontend_architecture.md`, `next-web-app.md`, `webhook_setup.md`,
      `alignment_testing.md`. Update `last_updated`. Verify cross-refs resolve.

## Pre-flight audits
(recorded by build-milestone Phase 0b below, per section)
