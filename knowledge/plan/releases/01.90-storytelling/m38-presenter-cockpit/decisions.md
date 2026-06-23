# M38 — Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions live in the spec
([`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md), D15 standalone panel / D9 single source).

## M38-D1 — The producers live in the `seeders` package, single-sourcing the unexported derivation
The roster-export producer (`roster.go`) and the cockpit-manifest exporter (`cockpit.go`) both live in
`stack-seeding/seeders/` — NOT a new package, NOT in Clerkenstein. Rationale: the load-bearing invariant is that
a hero's exported clerk ids EXACTLY match what `UsersSeeder` wrote to Postgres (else "login as Maya"
authenticates a non-existent user). The derivation helpers (`deterministicUUID`, `storyKeyPrefix`, `stackHost`,
`emailFor`, `splitName`, `personaUserIndexFor`, `roleForIndex`) are unexported in `seeders`. Putting the
producers in that package lets them call the helpers DIRECTLY — single-source, never re-derive. Clerkenstein
(a separate Go module, per M37 ARCH) consumes the exported JSON verbatim. A `DisallowUnknownFields` mirror test
(`rosterEntryMirror`) guards the producer's JSON stays loadable by the consumer without a cross-module import.

## M38-D2 — The cockpit reads a Go-projected JSON manifest, not the YAML directly (D9 preserved)
The demo tooling is stdlib-only Python (the supply-chain GREEN posture — no PyYAML). D9 says the cockpit reads
"the same `stack.stories.yaml`". Resolution: the Go side (which already owns the YAML schema + validation)
PROJECTS the file into a cockpit manifest JSON (`stackseed --cockpit-export`); the Python panel reads the
projection. The YAML is parsed ONCE, on the Go side, so the cockpit menu is single-sourced from the exact file
that seeded the data — D9's no-drift property holds, and Python stays dependency-free. A `keys-match-roster`
test asserts the cockpit's hero keys equal the roster's identity keys (both project the same blueprint).

## M38-D3 — [Login as] + [Jump to section] = ONE FAPI handshake redirect
M37's `handleHandshake` selects the seat from `?__clerk_identity=<key>` THEN redirects to `redirect_url`. So
the cockpit's two actions collapse into one redirect:
`https://<fapi>/v1/client/handshake?__clerk_identity=<key>&redirect_url=<jump_to>` — the hero becomes active
EVERYWHERE (client view, `/v1/me`, minted token, cookies) AND the browser lands on her screen, atomically. No
separate "switch then navigate" two-step that could desync (the exact failure M37-O11 rejected token-injection
for). [Login as] redirects to the app root; [Jump] to the hero's `jump_to`.

## M38-D4 — `DEMO_STORIES` opt-in; default-off keeps every existing demo byte-identical
The whole storytelling layer (stories-preset seed + multi-identity roster + cockpit) is gated on
`DEMO_STORIES=1`. Default off: a non-stories demo seeds `small-200`, the fake-fapi block has NO
`FAKE_FAPI_ROSTER` (single-identity `DefaultDemoUser()` fallback), and no cockpit is served — byte-identical to
a pre-M38 demo. This keeps the 5 alignment gates + all existing demo behavior untouched, and makes the
storytelling demo a single explicit flag. The roster, cockpit, and seed all pin `--stack demo-N`, so the
exported ids and the seeded rows are guaranteed to match (the single most fragile seam).

## M38-D5 — The cockpit is a host-native process (like ant-academy), reaped by `rosetta-demo down`
The cockpit (a standalone served panel, D15) runs as a `nohup python3 cockpit.py &` host process — NOT a
compose service — so it never touches a platform repo or the compose topology. Its pid is recorded into
`<stack>/cockpit.pid`; `rosetta-demo down` kills it before `compose down` (mirroring the M19 ant-academy
native-process teardown). The roster export builds `stackseed` ONCE into a shared bin dir handed to the
set-dress via `DEV_SETDRESS_BIN`, so the roster + the seed come from one build of one preset (no id drift).

## M38-D6 — A stories demo FAILS LOUD on a broken roster, but the cockpit is NON-FATAL
The roster export is the multi-identity contract: a stories demo with no/broken roster would log in as the
wrong (or a half-formed) user, which is worse than no demo — so a failed roster export `exit 1`s the bring-up.
The cockpit SERVE, by contrast, is non-fatal (`|| log warning`): a cockpit failure leaves a fully-working
seeded multi-identity demo (you can still hit the handshake URLs by hand), so it follows the M18/M19
never-abort-a-good-bring-up convention.

## O9 — RESOLVED: the deep-link catalog
Enumerated in `seeders/cockpit.go::DeepLinkCatalog()`. End-user (individual) routes: `/profile`,
`/profile?view=spotlight` (Skill Spotlight), `/growth` (My Growth), `/simulations` (Take a Sim). Manager
(org-intelligence) routes: `/enterprise/workforce` + its tabs (`?tab=skills-verification|role-readiness|
succession|mobility`) + `/enterprise/talent-pool`. The preset's two declared `jump_to` landing screens
(`/profile`, `/enterprise/workforce[?tab=…]`) resolve to catalog labels; an unrecognized `jump_to` still works
(raw path, generic "Open" label). The v1 routes are all id-free landing screens (`needs_id=false`); the field
is carried so a future cockpit can warn/fill per-hero/skill id params.
