# M234 — Progress

## Sections

Derived from `overview.md` + `roadmap.md` § M234. M234 = the **cockpit-UX render half** of Content
stories (the manifest half is M233; the fixture-population + prove-it-lands is M235). Three implementation
sections + docs folded into each section's Phase 5.

### §1 — Content-player seat registration (roster.go) ✅ (rext 6f7001b)
The as-player CTA (`?__clerk_identity=content-player-<idx>`) must authenticate as the real seeded member
who owns the session. Extend the exported roster so those seats resolve, single-sourced with the seeder's
own id/claim derivation.
- [x] Extract `storyPopulationNames` (pure replay of the UsersSeeder name assignment) + refactor UsersSeeder to consume it (true single-source, no drift)
- [x] `contentPlayerRosterIdentities(s)` — one roster identity per DISTINCT content-player owner slot the content-manifest projection references (auth_id / eid / email / name / picture / org-claims / role=member single-sourced)
- [x] Append content-player seats to `BuildRoster` output (after all heroes → default seat unchanged)
- [x] Tests: claims match UsersSeeder's writes; seat key == manifest `player_seat`; first roster entry still a hero; no duplicate keys; Clerkenstein-decodable shape (roster↔registry lockstep) — 6 new tests, full module green

### §2 — Content stories cockpit tab render (cockpit.py) ✅ (rext cbca13c)
The 2nd "Content stories" tab beside "Org stories": per-product sections rendering the M233 manifest.
- [x] `--content-manifest` arg + shape validation (byte-identical page when absent — old bring-up safe)
- [x] Client-side tab toggle (Org stories | Content stories) reusing the stdlib `_OVERLAY_JS`-style pattern; no manifest data interpolated into JS
- [x] Per-product sections (product FontAwesome icon + name); per-session rows with per-`sim_type` row icons + descriptor (modality / passed vs not-passed)
- [x] Two fake-FAPI deep-link CTAs per session (as-player / as-manager), `.actions` two-button layout + `has_manager_view` omission
- [x] Per-product app-base routing generalizing the `is_hiring`/`hiring_base` switch (web :3000 / hiring :3001 / academy :3077)
- [x] AI-labs section = PRESENCE-ONLY (no player path → status/spend line, no CTAs) — M231 D4
- [x] Academy section = player CTA to the academy origin (app_base=academy), no manager CTA — M231 D5 (direct academy link + M53 cookie seam; specific-member landing → M235)
- [x] Serve `/content-manifest.json` endpoint
- [x] Python unit tests: 23 new — dispositions, CTA hrefs, omission, presence-only, academy origin, tab toggle, served endpoint, shape validation, escaping; 0 new failures (106 tests, 6-fail baseline unchanged)

### §3 — Bring-up wiring (up-injected.sh) ✅ (rext 7f55eb4)
- [x] Export `content-manifest.json` via `stackseed --content-export` at bring-up (non-fatal, parallel to `--cockpit-export`; fail-closed export just drops the tab)
- [x] Thread `--content-manifest` into the `cockpit.py` launch (+alt guard, set -u safe)
- [x] Verify the export + launch wiring — new `StorytellingCockpitWiring` test (10/10) + proven end-to-end (real export → cockpit → probe: tab renders 9 sessions × 2 CTAs, `/content-manifest.json` served)

### Docs (Phase 5, folded)
- [ ] Extend `corpus/ops/demo/content-stories-spec.md` with the render half (tabbed model, two-action contract, icon map, base routing, seat registration)
- [ ] Note in `cockpit-spec.md` (the 2nd tab) + update roadmap/state at close-of-milestone
