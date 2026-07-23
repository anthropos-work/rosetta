# M246 — Spec notes

_Topic → doc → code triples + re-sync / bring-up drift findings accumulate here during build._

## Pre-flight audits — Seeder re-point (M246 first section)
KB-fidelity audit: **YELLOW** (proceed with tracking). Report:
`knowledge/plan/releases/02.70-july-jitter/m246-re-sync-repoint/kb-fidelity-audit.md`. Anchors
verified; KB-dep docs all present; stale skillpath-live corpus claims Fate-2 owned by M247.

## Seeder re-point (`skillpath.skill_path_sessions → public.skill_path_sessions`)

Rext authoring copy: `.agentspace/rosetta-extensions/stack-seeding` (separate git repo; on `main`
@ `307e70b`, tag `sound-check-m245-academy-durable-fix`). The re-point flips the SCHEMA qualifier
`skillpath` → `public` for the `skill_path_sessions` table only. **Leave untouched:** the DNA
surface NAME `skillpath-sessions`, the mirror table `public.local_skill_path_sessions`, and the Go
symbol/file names (`skillpath_sessions.go`).

**LIVE-CODE write sites (schema `"skillpath"` → `"public"`):**
- `cmd/stackseed/main.go:97` — reset TRUNCATE list entry `"skillpath.skill_path_sessions"` → `"public.skill_path_sessions"`
- `seeders/hero_activity.go:180` — `CopyRowsIdempotent(ctx, "skillpath", "skill_path_sessions", …)` arg → `"public"`
- `seeders/hero_activity.go:188` — audit tag `Schema: "skillpath"` → `"public"`
- `seeders/skillpath_sessions.go:116` — `CopyRowsIdempotent(ctx, "skillpath", …)` arg → `"public"`
- `seeders/skillpath_sessions.go:122` — audit tag `Schema: "skillpath"` → `"public"`
- `seeders/content_nonsim.go:419` — struct `{schema: "skillpath", table: "skill_path_sessions", …}` → `"public"`
- `seeders/content_nonsim.go:456` — audit tag `Schema: "skillpath", Rows: 0` → `"public"`
- `dna/data-dna.json:464` — `"schema": "skillpath"` → `"public"` (KEEP `:463 "surface":"skillpath-sessions"`)

**TEST assertions (schema arg / SQL `skillpath` → `public`):**
- `cmd/stackseed/main_test.go:232,274`
- `seeders/hero_activity_test.go:77,144` (`:144` is `failTable:"skillpath.skill_path_sessions"`)
- `seeders/content_nonsim_test.go:127,196,349,418,550` (`:550` = `failTable`)
- `seeders/dashboard_integration_test.go:251` — `DELETE FROM skillpath.skill_path_sessions …` SQL
- `seeders/contentref_test.go:136`
- `seeders/activity_seeders_test.go:30,53`

**COMMENTS referencing the schema** (update to follow the code — they name the write target):
`hero_activity.go:23,64`, `contentref.go:16`, `skillpath_sessions.go:16,43`,
`content_nonsim.go:22,332,351,375,479`. (These name `skillpath.skill_path_sessions` as the row's
destination; after the re-point the destination is `public.skill_path_sessions`.)

Per-stack policy: rext change → commit → **tag + `git push --tags` to origin** (rung-zero) → the
demo consumes it at the pinned tag.

## Demo clone pins (`clones.pin.json` + `DEMO_ADVANCE_CLONES=pinned` + bump to `origin/main`)

**Mechanism ALREADY WIRED (M237)** — `demo-stack/ensure-clones.sh`:
- `:172-183` the `DEMO_ADVANCE_CLONES` gate; `:206-227` the `pinned` mode reads `$DEMO/clones.pin.json`
  = `{"<repo>": "<tag-or-sha>"}` and checks each clone out at its ref.
- `:315` the pin-drift freshness check; `:355` `DEMO_FRESHNESS_STRICT=1` fatal.
M246's delta = **AUTHOR `stack-demo/clones.pin.json`** (absent today) pinning each demo repo to
current `origin/main` HEAD. jobsimulation stays standalone (still live). NB `stack-demo/` is a
git-ignored ephemeral workspace — confirm whether the pin belongs in the workspace (ephemeral, for
this prove) or as a rext preset copied in at bring-up (durable/reproducible) during build.

## Injection-comment fix (`gen_injected_override.py:16`, 3 subgraphs)

`stack-injection/gen_injected_override.py:15-16` comment currently: "the injected/subgraph services
are backend/app, cms, jobsimulation, skillpath" (4). Consolidated reality = **3 subgraphs**
(skillpath decommissioned into app). Fix = the COMMENT only (declared scope). **Scope-watch:** the
`INJECTED` dict `:17`, the service enum `:458`, and `exposure_claim_guard.py:124` still list
skillpath as a live service; whether the bring-up trips on that is a go/no-go surface → M247 drift
ledger.

## Cold `/demo-up` GREEN prove + M247 drift ledger
_(TBD during build.)_
