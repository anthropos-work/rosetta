# M37 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **Users/orgs registry** — replaced the single `DefaultDemoUser` with a `*Registry` (ordered, keyed map of `DemoUser`) in `clerk-frontend`; `NewServer` wraps one user (single-identity byte-identical), `NewMultiServer(reg)` for multi; fed a roster JSON (`RosterEntry`, `FAKE_FAPI_ROSTER`) the demo tooling exports from the seeder's hero roster
- [x] **Active-user selection** — **O11 RESOLVED: parameterized FAPI handshake** (server-authoritative, NOT pure token-injection — which desyncs). `?__clerk_identity=<key>` on the handshake (cockpit [Login as] deep-link) + the `/v1/demo/{identities,select}` control plane; select drops the session so the new seat starts clean
- [x] **Multi-identity Alignment DNA** — authored `clerk-multi-1` (5 caps / 9 genes: Roster, Select, DistinctIdentity, HandshakeSelect, DefaultSeat) + `cmd/multirun` runner + 9 goldens (`alignctl capture`); scores **100%/100%**
- [x] **Hold the gates** — the existing 4 surfaces stay green after the registry refactor: Go 22/22, JS 9/9, deploy 7/7 (express = node-CI gate, runner unit tests pass); verified after every change
- [x] **Reconcile `wip/clerkenstein-browser-login`** — folded its 32-line browser-login handshake note into `architecture.md` (improved to reference the implemented symbols) + added a multi-identity section; **retired the wip branch**
- [x] **Docs** — clerkenstein `knowledge/` (architecture/alignment/coverage + READMEs + CLAUDE = 5 surfaces; fixed a stale JWKS README claim) + the corpus pointers (`clerkenstein.md` § Multi-identity + `rosetta_demo.md`)
- [x] **Tests** — clerkenstein suite green (243 tests + 8 fuzz, `-race`); the multi-identity gate green; every M37-new function at 100% coverage

_Last updated: 2026-06-23 (build complete; all sections done). Ran ∥ M36 (different ext section)._
