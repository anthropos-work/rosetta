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
