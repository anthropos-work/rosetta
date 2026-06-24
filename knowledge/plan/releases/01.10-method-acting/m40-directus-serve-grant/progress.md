# M40 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.
Code lands in the rext authoring copy (`.agentspace/rosetta-extensions`) @ tag `method-acting-m40`;
doc-half in `corpus/ops/snapshot-spec.md`.

## Section checklist
- [x] **(a) `directus_versions` serve-grant** — DONE. Synthesized public-read **+ create** grants on the
  `directus_versions` SYSTEM collection (the full system name, NOT the `versions` API path — verified live: only
  `directus_versions` flips `/versions` from 403→200). `create` lets cms's `GetLatestOrCreateVersion` self-heal the
  empty per-stack version table instead of replaying prod's 539 version rows. Unblocks the entire skill-paths library +
  every detail page. Live: `publicSkillPaths = 22`.
- [x] **(b) library-category collections serve-grant** — DONE. Extended `servedCollections` with the library closure
  (`library_categories`, `library_macro_categories`, the 2 M2M junctions) + synthesized `directus_fields`/`relations`
  for the M2M expansion + a synthesized public-read grant (prod copies only its 5 public collections). Also added
  `resource` + `job_position` (the M2O targets `skill_paths.video` / `simulations.job_position` expand) so cms doesn't
  get an unmarshalable bare FK string. Live: `publicJobSimulations = 50`, no `ToDomain` panic.
- [x] **(c) `simulations.sequences` O2M nested-read serve-grant** — DONE. The O2M is grantable in tooling — it needed
  the `directus_relations` + `directus_fields` registration (the per-stack Directus had 0 of each), NOT a platform
  nil-guard. The library half AND the activity-feed half BOTH ship in tooling (the key-risk fork refuted). Live:
  `jobSimulation(simulationId)` returns `sequences[].scenarioIntro` — the feed federation path no longer panics.
- [x] **Regression test** — DONE. 9 new unit tests in `stack-snapshot/directus/serve_test.go` (closure both-endpoints,
  off-stack-alias drop, self-guarded idempotency, the synth/versions grants, the served-set closure, the six-part
  ordering). Live acceptance on demo-3: all three surfaces serve `> 0` anonymously via cms.

## Notes
- **Root cause was larger than the original a/b/c framing** (recorded M40-D2): the per-stack Directus had
  `directus_relations = 0` / `directus_fields = 0` — the O2M/M2M aliases were UNKNOWN to Directus, not "stripped under
  the public policy". Both library + activity-feed halves ship in tooling (M40-D3).
- **KPI "AI simulations completed" = 0** (overview Open question): its source `public.local_jobsimulation_sessions`
  has no CMS dependency, so it is a separate frontend/auth-context concern — out of scope for this serve-grant, to be
  re-verified during the M42e/M42m coverage sweeps. Not coupled to the feed fix (confirmed: the feed fix is a
  Directus-serve gap; the KPI reads jobsimulation directly).
- Supply-chain GREEN (go.mod/go.sum byte-identical); directus pkg `-race` clean; gofmt/vet clean; zero platform edits.
