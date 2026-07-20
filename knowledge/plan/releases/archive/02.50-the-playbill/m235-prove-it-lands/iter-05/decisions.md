# iter-05 — decisions

## D1 — separate code-owned non-sim exhibit registry, NOT the simulation fixture
The simulation fixture (`contentsession`) + its validator are simulation-shaped (require `sim_type` ∈ the
four SIMULATION_TYPE_*, modality, score) and the `ContentStorySeeder` seeds EVERY `Sessions()` entry as a
jobsimulation session. Adding non-sim products to that set would (a) fail the sim validator and (b) make the
sim seeder try to seed a skill-path as a simulation. So the non-sim exhibits are a **separate code-owned
registry** (`nonSimExhibits()` in `seeders/content_nonsim.go`) with its OWN seeder
(`ContentStoryNonSimSeeder`) and its OWN projection arm (`buildNonSimProducts` → appended by
`BuildContentProducts`). The simulation projection/seeder/fixture are UNTOUCHED — zero regression to the
closed M232/M233/M234 deliverables (proven: the 13-session sim projection + honesty gate stay green).

## D2 — self-contained non-sim flat index; owner-sharing with simulation exhibits is intentional
The non-sim projection + seeder use their OWN flat index over `nonSimExhibits()` (starting at 0), owner =
`slots[idx % len(slots)]`, advancing through drops — a self-consistent seeder↔projection pairing (proven by
`TestContentStoryNonSimSeeder_OwnerMatchesProjection`). This reuses the same `eligiblePlayerOwnerSlots`
member pool as the simulation exhibits, so a content-player (e.g. content-player-23) may OWN both a
simulation session AND a skill-path — harmless and more believable (a real employee does both). All seats are
registered by `roster.go`'s `contentPlayerSeatsUsed(BuildContentProducts(s))` (now non-sim-inclusive).

## D3 — pinned real public skill_path_ids, sourced OFFLINE from the captured snapshot
The `/skill-path/<skillPathId>` route + the seeded row's `skill_path_id` are single-sourced from a REAL
public `directus.skill_paths.id`, sourced OFFLINE from the captured public Directus snapshot
(`.agentspace/snapshots/directus/…/directus.skill_paths.copy`, captured under
`private=false AND tenant_id IS NULL AND status='published'`) — public content, no PII, no prod query. The
two pinned paths: `df9d2142-…` ("Become a Product Manager", completed) and `a6087b81-…` ("Practical
introduction to GenAI…", 45% in-progress). They resolve against the demo's replayed public catalog.

## D4 — version "2" is the collision-safe choice (CopyRowsIdempotent guards only ON CONFLICT (id))
`skillpath.skill_path_sessions` has a UNIQUE `(user_id, skill_path_id, version)`. The generic
`SkillpathSessionsSeeder` writes version `"1"` for these same owners, and `CopyRowsIdempotent` guards only
`ON CONFLICT (id)` — a unique-violation on a DIFFERENT constraint would ERROR the whole live seed. Writing
version **`"2"`** makes a collision impossible by construction (no coupling to the generic seeder's live
skill_path_id assignment). Pinned by `TestContentStoryNonSimSeeder_SkillPathRows`.

## M236 LIVE-CALIBRATION CHECKLIST (the skill-path arm's live-render unknowns — route to M236)
These resolve ONLY against a live seeded render (exactly what the run-3 ruling routes to M236, which must
"CALIBRATE against a live seeded render" before the cold-reset-to-seed proof). Unit-proven here = the
seeder writes structurally-correct rows + the manifest resolves; NOT that the page renders non-empty.

1. **getOrCreateSkillPathSession version-MATCH.** The seeder writes version `"2"`. The `/skill-path/<id>`
   page calls `getOrCreateSkillPathSession(userId, skillPathId, version?)`. If the page requests a specific
   version (the skill path's current CMS version) and it isn't `"2"`, getOrCreate auto-materializes a BLANK
   → the calibration must confirm the page finds the seeded session (adjust the seeded version to the path's
   real current version, or a sha-pinned demopatch, if not).
2. **status vocabulary — `active` vs `in_progress`.** The seeder writes the DOCUMENTED enum
   (`active`/`completed`, skillpath.md `pending|active|completed|archived`). The generic seeder writes the
   OFF-enum `"in_progress"`. Confirm which value the direct `/skill-path/<id>` page + the manager scoreboard
   render (and, separately, whether the generic seeder's `"in_progress"` is a latent bug — out of M235 scope).
3. **mirror `local_skill_path_sessions` uniqueness.** The mirror row uses a distinct derived id; confirm live
   there is no `(user_id, skill_path_id)` unique collision with the `AssignmentsSeeder`'s mirror rows for a
   shared owner+path (low risk — distinct pinned ids — but a live-seed reality only the cold run proves).

## Closure gate — `datadna measure-closure` is a LIVE gate → M236
`datadna measure-closure` requires a seeded stack DSN (exit 3 without one) — a LIVE gate, part of M236's
proof. The OFFLINE closure invariant (no fabricated refs) holds by construction (pinned real public ids +
real owner slots) and is enforced by the fail-closed projection + `ValidateContentManifest` (proven by
`TestBuildNonSimProducts_NoHostDropsAllFailClosed`).
