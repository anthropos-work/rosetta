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

## M37: Hardening

**Scope manifest (milestone-touched code, `release/01.90-storytelling..HEAD` in the `clerkenstein` ext section):**
- `clerk-frontend/registry.go` — the registry + roster JSON contract (`Registry`, `NewRegistry`, `Select`,
  `RosterEntry`/`Roster`, `RegistryFromRoster`, `LoadRoster`). Tests: `registry_test.go`, `fuzz_test.go`.
- `clerk-frontend/server.go` (M37 additions) — `NewMultiServer`, the `__clerk_identity` handshake branch,
  `handleListIdentities`, `handleSelectIdentity`. Tests: `multiidentity_test.go`.
- `cmd/fake-fapi/main.go` — `buildServer` roster-load + single-identity fallback. Tests: `main_test.go`.
- `alignment/cmd/multirun/main.go` + `dna/clerk-multi-1.json` + `golden-multi/*` — the 5th Alignment DNA
  runner (test harness; measured by the 9-gene gate, not unit coverage). Tests: `main_test.go`.

### Pass 1 — 2026-06-23
**Coverage delta (milestone-touched files):**
- `clerk-frontend`: 86.0% -> 86.0% statements (the M37 selection handlers `handleSelectIdentity` /
  `handleListIdentities` moved 0->100%; the unchanged package total masks it). `registry.go` stayed 100%.
- Remaining sub-100% in the package is **pre-M37, out-of-scope**: `handleHandshake:193` (an unreachable
  RS256-mint-failure 500 path — the mock's fixed keys never fail) + pre-existing `corsMiddleware` /
  `handleDevBrowser` / `handleClerkJSBundle` / `handleToken`.

**Tests added (10):** roster duplicate-key + empty-key rejection through `RegistryFromRoster` AND `LoadRoster`
(the path the tooling feeds, not just the in-memory constructor); minimal-entry (ids-only) roster is usable;
malformed-JSON + blank-string + empty-body `/v1/demo/select` all 400 leaving the seat unchanged; blank
`?__clerk_identity=` keeps the current seat + still signs in (the empty-value branch of the `key != ""` guard);
mid-session seat-switch via the handshake while signed in as another hero (no prior-seat residue across
`/v1/me` + `/v1/client` + the minted token); single-identity-server seat-switch inert-but-safe invariant.

**Bugs fixed inline:** none — no production defect surfaced.

**Behavioral finding pinned (not a bug):** on a single-identity server, `POST /v1/demo/select` of the sole
member's own (internal `default`) key returns **200 as a no-op** (the seat stays active), while any other key
is a 400. Selecting the only seat is idempotent + harmless; the cockpit only ever selects to *switch*. Pinned
by `TestDemo_singleServerSeatSwitchIsInert` so a future change to the select control plane can't silently let
a single-identity demo be stranded or jumped to a non-existent hero.

**Flakes stabilized:** none observed.

**Knowledge backfill:** the single-identity select-no-op + the byte-identical-fallback control-plane invariant
folded into `clerkenstein/knowledge/architecture.md` (§ Multi-identity) — see Pass-2 entry.

### Pass 2 — 2026-06-23
**Coverage delta:** negligible on lines (the gap closed here was an untyped-input *fuzz* surface, not
uncovered statements). `clerk-frontend` stays 86.0%.

**Tests added (1 fuzz):** `FuzzLoadRoster` — fuzzes the `FAKE_FAPI_ROSTER` operator-supplied roster-parse
boundary (the new M37 untyped-input surface; the two existing fuzz tests cover the PK codec + JWT round-trip,
not the roster). Asserts: never panics; the `(registry, error)` pair stays consistent (error => nil registry,
no half-formed seat leaks; success => >=1 seat, a valid active key, every seat carrying non-empty core ids).
**360K execs, 0 panics, 0 contract violations** — no crash seed to pin.

**Bugs fixed inline:** none.

**Knowledge backfill:** added the roster-parse robustness invariant + the single-identity select no-op to
`clerkenstein/knowledge/architecture.md` (§ Multi-identity) so a downstream M38 consumer / future Clerk-bump
reconcile sees the boundary contract. `last_updated` bumped.

### Stop condition
Stopped after **2 passes**: the full Step 2b six-dimension scan found nothing further worth adding (all M37
public functions at 100%; the only sub-100% lines are pre-M37 unreachable crypto-mint failures + test-harness
network fallbacks the 9-gene gate already drives end-to-end); coverage deltas on M37 product code < 2%; no
flaky tests across the gate re-runs + the flake gate. Both alignment gates held **100%/100%** ("No
divergences") after every commit; the `-race` suite stayed clean.

_Hardening last updated: 2026-06-23 (2 passes; +11 tests [10 unit/integration + 1 fuzz]; 0 bugs; gates held)._
