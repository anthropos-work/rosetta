# M37 — Retro

## Summary
M37 delivered the **multi-identity seat-switch** the presenter cockpit's "login as" needs. Clerkenstein's fake
FAPI — which resolved every session to **one** `DefaultDemoUser` — now holds a **users/orgs Registry** (an
ordered keyed map of `DemoUser`) fed a **roster JSON** (`FAKE_FAPI_ROSTER`, `RosterEntry`) the demo tooling
exports from the seeder's own derivation (single-sourced ids; the ARCH decision), plus a **server-authoritative
active-seat selection** so a demo can present as any seeded hero. **O11 was resolved by spiking both options:**
the **parameterized FAPI handshake** won over pure token-injection (which desyncs the server-held active
identity from an injected cookie) — `?__clerk_identity=<key>` on the handshake (the cockpit's [Login as]
deep-link) + the `/v1/demo/{identities,select}` control plane, so every surface (client view, `/v1/me`, the
token mint, the handshake cookies) resolves the same hero. A **5th Alignment DNA** (`clerk-multi-1`, 9 genes,
runner `multirun`) measures the multi-session FAPI surface real clerk-js exhibits with `single_session_mode=false`,
scoring **100%/100%**; the single-identity path stays **byte-identical** (a one-member registry) so the 4
existing surfaces stay green. The `wip/clerkenstein-browser-login` branch was **reconciled** (its 32-line
handshake note folded into `architecture.md`, improved to reference the implemented symbols) + **retired**.
M37 ships the Clerkenstein **consumer** + a golden roster fixture; the seeder-side roster-export producer + the
clickable cockpit are the M38 integration seam. Tooling + docs only — zero platform-repo edits. Close GREEN,
1 finding, 0 blocking, merged into `release/01.90-storytelling`.

## Incidents This Cycle
None. No P2 flakes, no regressions. The 2 harden passes surfaced **zero bugs** — every edge/error path on the
seat-switch surface (roster dup/empty-key through both `RegistryFromRoster` AND `LoadRoster`, blank handshake
identity, malformed/blank `/v1/demo/select`, mid-session residue, single-server inert-switch) confirmed correct
under deeper probing, and `FuzzLoadRoster` ran 360K execs with 0 panics / 0 contract violations. The flake gate
ran 5/5 clean (shuffled, `-race`). The one close finding (a stale "four DNAs" claim in the corpus
`alignment_testing.md`) was a documentation fix, not a code defect.

## What Went Well
- **The early O11 spike paid off.** Spiking both token-injection and the FAPI handshake up front exposed the
  desync failure mode (server-held active identity vs injected cookie) before any registry code committed —
  the handshake's "one source of truth" was the robust pick, and the cheap part (`sessionClaimsFor` + `MintRS256`
  per-identity minting) was already done.
- **The byte-identical single-identity fallback** (a one-member registry) meant the 4 existing alignment gates
  held green through the registry refactor with no special-casing — the multi-identity surface is purely
  additive.
- **The module-boundary discipline held:** Clerkenstein owns the registry + roster-JSON contract; the seeder
  owns the id derivation. Feeding a roster (not importing the seeder's packages) keeps ids single-sourced
  without coupling the two modules.

## What Didn't
- **The `clerk-express-1` gate is env-fragile.** It drives the genuine `@clerk/express`/`@clerk/backend` SDK
  and so needs installed npm modules — unrunnable in the authoring copy without them. Not an M37 regression
  (M37 never touched it; its Go runner unit test passes), but a recurring "run-it-where-deps-are" friction the
  v1.1 CI-wiring carry-forward already tracks.

## Carried Forward
- **M38 — Presenter cockpit (the LAST v1.9 milestone):** the standalone served panel (rext `demo-stack`, offset
  port) that lists stories→heroes with [Login as] + [Jump to section], wiring M37's seat-switch + the
  seeder-side roster-export producer into a clickable login-as-a-hero cockpit. Owns the literal live browser
  seat-switch render (Fate-2 by design — M37 ships the consumer + the golden fixture).
- **Ext-tag push (release-level, from v1.8 close):** push `storytelling-m34..m37` (+ the older `understudy`/
  `house-lights`/`stage-door`/`prop-room` tags) to `origin` — the orchestrator's post-close protocol.

## Metrics Delta
- **clerkenstein Go tests:** 250 `Test` + 9 `Fuzz` across 14 packages, `-race` green (M37 added 11 harden tests
  — 10 edge/error + the `FuzzLoadRoster` roster-parse fuzz — on top of the build-phase registry/selection/`multirun`
  tests).
- **Coverage:** every M37-new function 100% (`registry.go` 100%; the selection handlers 0→100%); `clerk-frontend`
  pkg 86.0% (residual all pre-M37 unreachable crypto-mint paths).
- **Alignment gates:** 100%/100% on all 5 surfaces — `clerk-multi-1` 9/9 (NEW) + Go 22/22 + JS 9/9 + deploy 7/7
  (held); `clerk-express-1` is a node-CI gate (env prereq).
- **Flake count:** 0 (gate 5/5, shuffled `-race`). **Supply-chain:** GREEN (0 new deps).
- Full machine-readable record: [`metrics.json`](metrics.json).
