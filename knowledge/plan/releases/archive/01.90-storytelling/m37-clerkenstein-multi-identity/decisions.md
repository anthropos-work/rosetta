# M37 — Decisions

Implementation decisions with rationale (recorded during build). Key open decision: **O11** — the
active-user selection mechanism (token-injection vs parameterized FAPI handshake), to be settled by an early
spike.

## O11 — RESOLVED: parameterized FAPI handshake (server-authoritative selection), NOT pure token-injection
**Spiked both early (per the overview's mitigation).** Decision: **parameterized FAPI handshake.** Rationale:
- **Token-injection alone desyncs.** Injecting a hero's `__session`/`__client_uat` cookie into the browser
  does NOT update the FAPI server's view: `/v1/client`, `/v1/me`, and the `/tokens` mint still resolve to the
  OLD active user (the server holds the active identity, not the cookie). clerk-js re-fetches `/v1/client` on
  load and would show a session whose `user` (server) disagrees with the cookie (injected) → inconsistent
  state, exactly the "Invalid state" class the M2d work fought. Fragile.
- **The handshake is server-authoritative.** The FAPI holds the registry + the active key. Selecting an
  identity (a) on the FAPI moves the active key, so EVERY subsequent surface — the client view, `/v1/me`, the
  embedded `last_active_token`, the `/tokens` mint, AND the handshake cookies — resolves to the SAME hero.
  One source of truth, no desync. The handshake already exists (`handleHandshake`); parameterizing it with an
  identity selector is a small, robust addition.
- **The cheap part is unchanged:** `sessionClaimsFor(u,…)` + `MintRS256` already mint per-identity tokens for
  ANY `DemoUser`; we only change WHICH user is active, not the minting.
**Mechanism shipped:** the registry's active key is selectable via (1) `?__clerk_identity=<key>` on the
handshake (the in-flow seat-switch the cockpit's [Login as] deep-link uses) and (2) an explicit
`POST /v1/demo/select` + `GET /v1/demo/identities` pair (the cockpit's roster + programmatic switch). A pure
cookie-injection path is explicitly NOT used (the desync above).

## ARCH — Clerkenstein is its own Go module; the registry is fed a ROSTER, derivation single-sourced in the seeder
`clerkenstein` is `module clerkenstein`, a SEPARATE module from `stack-seeding`
(`github.com/anthropos-work/rosetta-extensions/stack-seeding`). Clerkenstein must NOT import the seeder's
`blueprint`/`seeders` packages. But a hero's registry `DemoUser` must carry the EXACT ids the seeder wrote to
Postgres (else "login as Maya" authenticates a non-existent user). The seeded contract (from `users.go` +
`stories.go`): `Eid=deterministicUUID("<prefix>:user:<idx>")`, `AuthID="user_seed_<slug(prefix)>_<idx>"`,
`Email="<slug(first)>.<slug(last)><idx>@<domain>"`, `OrgEid=LegacyOrgID|StoryOrgID(id)`,
`OrgAuthID=LegacyOrgClerkID|StoryOrgClerkID(id)`, `OrgRole=roleForIndex(idx,size,mix)`; `idx` = hero
declaration-index + 1; `prefix` = stack (story 0) else `stack:story:<id>`. **Resolution:** Clerkenstein owns
the `Registry` + a **roster JSON contract** (`[]DemoUser` + key) the `fake-fapi` loads from `FAKE_FAPI_ROSTER`;
the demo tooling (which already runs the seeder) EXPORTS that roster using the seeder's own derivation, so the
ids are single-sourced and never re-derived in Clerkenstein. M37 ships the consumer (registry + loader + a
golden roster fixture matching the stories preset + the documented schema); the seeder-side `roster export`
producer is the demo-tooling/M38 integration seam. With no roster, `fake-fapi` falls back to the
single-identity `DefaultDemoUser()` — byte-identical to pre-M37 (keeps all 4 gates green).

## KB-1 — browser-login handshake doc folds in via the wip-branch reconcile (Phase 0b finding)
The Phase 0b KB-fidelity audit (GREEN) found one completeness gap: the M2d interactive browser-login
handshake is fully implemented in `clerk-frontend/server.go` (`handleHandshake` + dev-browser cookie +
RS256 `__session` + the `sid` claim) and `cmd/fake-fapi/main.go` (TLS), but NOT documented in
`clerkenstein/knowledge/architecture.md` on `main`. The `wip/clerkenstein-browser-login` branch carries a
32-line note that documents it. **Resolution:** fold that note into `architecture.md` as part of M37's
"reconcile the wip branch" deliverable, then retire the wip branch. Not a blocker — it is milestone-owned
work that lands in the documentation phase. Report: `kb-fidelity-audit.md`.

