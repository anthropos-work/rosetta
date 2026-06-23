# M38 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **The panel** — standalone served surface (rext `demo-stack`, offset port `7700+N·10000`), lists stories → hero trios, reads the cockpit manifest projected from the same `stack.stories.yaml` (`cockpit.py`, stdlib-only HTTP server)
- [x] **[Login as]** — wired to M37's active-user selection: a FAPI handshake redirect `?__clerk_identity=<key>` that switches the active seat to the chosen hero
- [x] **[Jump to section]** — the same handshake redirect with `redirect_url=<jump_to>`, so [Login as]+[Jump] land logged-in on the hero's deep-link in one move
- [x] **Deep-link catalog (O9)** — `DeepLinkCatalog()` enumerates next-web routes per vantage (profile/spotlight/growth/take-a-sim for end-users; the Workforce dashboard tabs + mobility/talent-pool for managers)
- [x] **Launch wiring** — `DEMO_STORIES=1` exports the roster → `FAKE_FAPI_ROSTER` (multi-identity fake-fapi), seeds the stories preset, serves the cockpit on the offset port; torn down with the stack (pidfile reaped by `rosetta-demo down`)
- [x] **Roster-export producer** (the M37 integration seam) — `stackseed --roster-export` derives the seeded heroes' exact clerk ids, single-sourced from the seeder's own derivation; consumed by Clerkenstein's registry
- [x] **Docs** — the cockpit section of `stories-spec.md` + the up→present flow in `demo/README.md`
- [x] **Tests** — demo-stack suite 157 green · stack-injection 115 green · stack-seeding green (`-race`); zero platform-repo edits; the 5 Clerkenstein alignment gates + stack-seeding suite stay green

_Last updated: 2026-06-23 (all sections shipped). rext tag `storytelling-m38` to be cut at close. Code:
`rosetta-extensions` @ `ce2b829` (stack-seeding roster/cockpit producers + demo-stack cockpit panel + launch
wiring). Docs on rosetta `m38/presenter-cockpit` @ `007378b`._
