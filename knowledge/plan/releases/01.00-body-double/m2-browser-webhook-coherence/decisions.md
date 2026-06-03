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

## M2-D3 — M1 latent bug: `cmd/clerkrun/main.go` was never tracked (fixed inline, Fate 1) — 2026-06-03

**Found during S4** while adding the JS runner's binary to `.gitignore`: M1's `.gitignore` used an
**unanchored** pattern `clerkrun`, which matches not only the built binary but also the `cmd/clerkrun/`
**directory** — so the M1 Go-surface runner *source* (`cmd/clerkrun/main.go`) was **silently excluded
from git** the whole time. The M1 alignment gate built from a working-tree-only file; a fresh clone
would have had no runner and could not reproduce the 100%/100% gate.

**Fix (Fate 1, lands in M2):** anchored the ignores to repo-root binaries (`/clerkrun`, `/jsfapirun`)
so `cmd/*/` source is tracked, and committed the previously-untracked `cmd/clerkrun/main.go`. The JS
runner (`cmd/jsfapirun/main.go`) is tracked from the start. No behavior change — the runner source is
byte-identical to what M1 has been building; this only makes the repo reproducible. Surfaced here per
the no-hide rule rather than silently re-ignoring.

## (template) — further decisions recorded as sections land
