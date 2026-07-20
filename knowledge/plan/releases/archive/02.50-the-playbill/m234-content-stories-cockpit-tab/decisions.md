# M234 — Decisions

## §1 — content-player seat registration

### D-M234-1 — member NAME is now a single source (`storyPopulationNames`)
A content-player roster seat must carry byte-identical claims to the seeded `public.users` row (else "log in
as this player" authenticates a mismatched user). The member NAME is order-dependent
(`nameForIndexAvoiding` accumulates a `taken` set as it walks the population), so it can't be re-derived by a
pure index formula. **Decision:** extract the whole per-story name assignment into `storyPopulationNames`
(userprofile.go) and route BOTH the `UsersSeeder` (users.go) AND roster.go through it — true single source,
not a replay-by-test. The refactor is output-identical (all name-sensitive suites pass unchanged; the whole
seeders module is green), so no golden/data-DNA drift.

### D-M234-2 — content-player seats APPEND after heroes; the set is the manifest's own
`BuildRoster` now appends one identity per DISTINCT content-player owner slot the content-manifest projection
references (`contentPlayerSeatsUsed(BuildContentProducts(s))`) — no more (a dead seat), no fewer (an
unresolvable as-player CTA). Appended AFTER all heroes, so `roster[0]` (the default active seat) stays the
first hero. Verified: the real stories preset yields 12 heroes + `content-player-23..31` (exactly the
`player_seat`s in `presets/content-manifest.json`), default seat still `maya-thriving`.

### D-M234-3 — the `--roster-export` warning is now HERO-scoped, not total-scoped
Since the roster carries content-player seats, a big STRUCTURAL org (no heroes, hundreds of member slots)
projects a non-empty roster (content-players) — which would mask the "you passed a preset with no heroes"
signal behind a non-zero total. **Decision:** the CLI warning keys on `RosterHeroCount`, symmetric with the
cockpit export's long-standing "0 heroes" warning. The shape/output is otherwise unchanged; a structural seed
never exported a roster in the real bring-up (roster export is gated on `DEMO_STORIES=1`).

### Cross-producer invariant updated (not broken)
The pre-existing `cockpit heroes ↔ roster` 1:1 lockstep tests now compare against the roster's HERO portion
(the roster is a legitimate SUPERSET: heroes for the "Org stories" tab + content-players for "Content
stories"). Every cockpit key still resolves in the full roster (no dead `[Log in as]`).

## §2 — the "Content stories" cockpit tab render

### D-M234-4 — the ACADEMY player CTA is a direct academy-origin link, NOT a FAPI handshake
The two CTAs are fake-FAPI handshake deep-links (`?__clerk_identity=<seat>&redirect_url=<base><path>`) for
the FAPI-backed products (simulation/skill-path on web/hiring). The ACADEMY product is different: ant-academy
is a **separate origin with its own auth** (the M53 `e2e_persona` cookie seam), NOT behind the fake FAPI — a
FAPI handshake redirect to the academy origin would establish a next-web session the academy can't see. So
the academy as-player CTA is a **direct academy-origin link** carrying the M53 cookie seam. The
**specific-member** academy landing (vs a generic entitled member) + the exact chapter route depend on M230's
catalog fill and are finalized in **M235** — this is literally M234's overview open question, and M235
`depends on M234 (+ M230 for the academy section)`. The manager CTA is always a FAPI handshake (manager
surfaces are next-web/hiring, never academy).

### D-M234-5 — the renderer handles all dispositions; today's fixture is simulation-only
M234 is the render HALF. The renderer handles every product disposition the M233 schema can carry —
simulation (2 CTAs), skill-path (2 CTAs on web), academy (player-only, academy origin), presence-only ai-labs
(status note, no CTA) — driven purely by the manifest fields, unit-proven against synthetic manifests. The
embedded fixture is **simulation-only today**, so a real M234 demo renders only the Simulation section; the
ai-labs/academy/skill-path FIXTURE additions + prove-it-lands are **M235** (per `content-stories-spec.md` §6
+ the roadmap). No dead CTA is stranded: the M233 fail-closed guard + M235 own the fixture.

### PRESENCE-ONLY is data-driven, not a flag
A session with no `player_result_path` AND no manager view has no result page — the renderer emits a muted
"Activity & spend only" note, never a dead button (M231 D4). No explicit `presence_only` field is needed;
the disposition falls out of the manifest fields.

### Pre-existing 6-fail cockpit carry — UNCHANGED (Fate-2, release-close)
`test_cockpit.py` had 6 pre-existing failures (removed per-hero academy CTA ×4 + the v2.3.1 overlay 30s-window
×2) — stale tests for intentionally-removed/changed behavior, part of the documented "14 pre-existing
demo-stack failures" standing carry routed to the **v2.5 release close** test-debt re-anchor. M234 adds **0**
new failures (verified: 106 tests, still exactly those 6). Not touched here — release-close scoped, Fate-2.

## Adversarial review (close Phase 2c)

The scenarios an external reviewer would probe against the render half + the seat registration — each was
exercised by the harden Pass 1 tests (progress.md § M234: Hardening) and the code handled it correctly. Recorded
here as the *scenarios considered* (not the fixes) so a future reviewer sees the threat model.

- **Manifest promises a manager view but no manager seat resolves.** A `has_manager_view=true` session whose
  manager seat is absent could render a dead `[Log in as manager]`. → `TestContentTabRenderEdges` proves the
  manager CTA is **omitted** when the seat is missing, not stranded. The renderer keys the CTA on seat presence,
  not the flag alone.
- **An academy session that also carries a manager view.** The academy as-player CTA is a direct academy-origin
  link (D-M234-4), but a manager CTA on the same row must NOT go to the academy origin (no academy manager
  route). → the manager CTA routes to `app_base` (web/hiring), academy origin only for the player CTA.
- **`--no-ui` / unset content base.** A hiring/academy `content_base` key with no base set could produce a
  malformed/dead href. → falls back to `--app-base`; never a dead link.
- **Empty product section / all-empty manifest.** An empty `content_products[]` entry, or a manifest with every
  product empty, could render a blank panel that reads as broken. → empty products are skipped (siblings still
  render); an all-empty manifest renders a friendly note, never a blank panel.
- **Injection via manifest data.** Manifest strings flow into HTML/JS. → `_TAB_JS`/`_OVERLAY_JS` are raw
  constants with NO manifest data interpolated; per-field values are escaped (`TestContentTab*` escaping cases).
- **A broken / fail-closed content manifest at bring-up.** A malformed `content-manifest.json` must not abort the
  cockpit. → non-fatal: the 2nd tab drops, "Org stories" still serves, `/content-manifest.json` 404s
  (`TestContentTabMainWiring` + the `test_tooling.py` non-fatal-export case). Byte-identical to the pre-M234
  single-panel page when absent (`TestContentTabByteIdenticalFallback`).
- **A structural (0-hero) org inflating the roster.** Content-player seats append to the roster, so a hero-less
  structural org could mask the "0 heroes" CLI signal behind a non-zero total. → the `--roster-export` warning
  keys on `RosterHeroCount` (D-M234-3), pinned by `content_player_roster_test.go`.

No scenario required a code change at close — all were already handled; recording completes the Phase 2c audit trail.
