# M40 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist
- [ ] **(a) `directus_versions` serve-grant** — SYNTHESIZE a public-read `directus_permissions` row on the
  `PublicPolicyID` for the `directus_versions` SYSTEM collection in rext `stack-snapshot/directus/structure.go`, so
  cms `skillpath.go:64` `GetSkillPath` → `GetLatestOrCreateVersion` → `version.go:40` `GET /versions` no longer 403s
  anonymously (unblocks the entire skill-paths library + every sim/path detail page — the dominant blocker).
- [ ] **(b) library-category collections serve-grant** — SYNTHESIZE public-read rows for the library-category
  collections `ListPublicJobSimulations` (cms `jobsimulation.go:305`) expands, so the 403 → empty relation →
  `ToDomain` panic `"index out of range [0]"` no longer empties the sims list.
- [ ] **(c) `simulations.sequences` O2M nested-read serve-grant** — investigate the O2M public-policy mechanism
  FIRST; make `cms GetJobSimulation` receive a non-empty `s.Sequences` so `jobsimulation.go:1097` (`s.Sequences[0]`)
  no longer panics → the activity feed's per-row simulation federation returns non-null → the feed fills. If the
  O2M is NOT grantable without a platform nil-guard, ship the library half independently and escalate the
  activity-feed half for platform sign-off.
- [ ] **Regression test** — re-replay the snapshot into demo-3 and assert all three surfaces
  (`/library/ai-simulations`, `/library/skill-paths`, `/profile/activities`) serve **>0** on a fresh demo.
