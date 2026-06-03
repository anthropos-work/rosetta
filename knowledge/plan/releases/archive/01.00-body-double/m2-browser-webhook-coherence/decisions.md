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

## M2-D4 — close review: `orgclient.ChangeRole` nil-map panic + phantom-membership divergence (fixed inline, Fate 1) — 2026-06-03

**Found at close (Phase 2c adversarial review).** `Store.ChangeRole` checked only `validRole` + org
existence, then assigned `s.members[org][user] = role` **unconditionally**. Two defects:
1. **Panic** — if the org exists but its `members` map is nil (any org created via `CreateOrganization`,
   which never initializes `s.members[id]`; or any org whose members map was never allocated), the
   assignment hits a **nil map → "assignment to entry in nil map" panic**. Reachable through the `bapi/`
   server in a live demo: `POST /v1/organizations` then `PATCH …/memberships/{user}` (ChangeRole) on the
   new org panics the HTTP goroutine. The alignment gate missed it: the `ChangeRole` gene only targets the
   seeded `o_1`/`u_1` (an existing member), so the nil-map path is never exercised by the runner.
2. **Behavioral divergence** — even on a seeded org, `ChangeRole` for a user who is *not* a member
   silently **created a phantom membership** instead of returning `ErrNotMember`. Real Clerk 404s
   (`resource_not_found`); the `bapi.membershipErr` switch already maps `not-a-member` → 404, so the twin
   was wired for the error it never produced.

**Fix (Fate 1, lands in M2):** `ChangeRole` now requires the (org, user) membership to exist before
mutating — mirroring `DeleteMembership`'s `if _, ok := s.members[org][user]; !ok` guard — returning
`ErrNotMember` for a non-member (whether the org's members map is nil or the user is simply absent).
Adds `orgclient` regression tests (nil-map repro + non-member-on-seeded-org) and re-verifies both
alignment gates (still 100%/100% — the seeded-member ChangeRole gene is unaffected). No platform-code
change; no wire-shape change.

## Adversarial review (Phase 2c) — scenarios considered at close

- **`orgclient.ChangeRole` on an org with no allocated members map / a non-member user** → nil-map panic +
  phantom-membership creation. **Real risk in the shipped code** (reachable via the `bapi/` server). Fixed
  inline — see M2-D4 above.
- **`fapi.ParsePublishableKey` with an embedded/missing `$` sentinel** → already hardened at M2 harden
  Pass 1 (`FuzzParsePublishableKey`, commit e80a257); Parse is the strict inverse of Mint. No residual.
- **`bapi` malformed/oversized/wrong-type request bodies** → covered by `bapi/malformed_test.go` +
  `bapi/fuzz_test.go` (harden Pass 2); decoders fail soft (the disarmed handlers tolerate a zero-value
  body), no crash.
- **`webhook.Injector` with a transport error / non-2xx endpoint** → covered by `injector_error_test.go`
  (harden Pass 2); `Inject` fails loud on non-2xx and wraps transport errors. No silent drop.
- **Concurrent `orgclient.Store` mutation through one shared store** (the M2-D2 injection-time concern) →
  5 `-race` concurrency tests (harden Pass 1) assert exactly-one-winner invariants under up to 64
  goroutines. No race.

## (template) — further decisions recorded as sections land
