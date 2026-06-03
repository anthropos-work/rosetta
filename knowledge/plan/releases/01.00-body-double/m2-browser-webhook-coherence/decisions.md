# M2 — decisions

## M2-D1 — JS FAPI override resolved at spike time: config, not fork (Fate 1) — 2026-06-03

**Decision:** The fake-FAPI path for `@clerk/clerk-js` / `@clerk/nextjs` is achieved by **configuration
only** — a minted publishable key encoding the fake FAPI host (and/or `proxyUrl`/`domain` props), set
via demo env vars. **No SDK fork. No platform-code change. The state.md "real dev Clerk app for the
browser" fallback is not exercised.**

**Evidence (spec-notes § Spike):** `@clerk/shared` `parsePublishableKey` decodes the FAPI host from the
publishable key (`base64(host$)`, third `_`-segment); the clerk-js browser bundle's `getFapiClient`
additionally honors `proxyUrl`/`domain`. Inspected in the installed SDK under
`anthropos-dev/studio-desk/node_modules/@clerk/`.

**Why this is Fate 1, not a deferral or a fallback:** the override lands fully inside M2 via config; the
roadmap's open question resolves in the strong direction (no fork). The fallback remains documented as an
escape hatch (if keeping the fake FAPI faithful across clerk-js bumps proves too costly — M1b-style drift
would flag it), but it is not the implemented path.

## M2-D2 — orgclient thread-safety lands here (M1-D2 carry-forward, Fate 1) — 2026-06-03

The M1 in-memory `orgclient.Store` (plain maps) is single-thread-only — M1's adversarial review + M1-D2
explicitly routed "make it concurrency-safe" to **injection time = M2**, because a single fake-API-server
instance serves concurrent demo requests. S2 adds the mutex + a concurrency regression test. This is the
Fate-1 home M1-D2 named; recorded here as the pickup.

## (template) — further decisions recorded as sections land
