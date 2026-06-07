# M13 — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [ ] Dev bring-up spawns a per-stack Directus (reuse M10 provision.go) + dev CMS repointed at it
- [ ] Auto-run `stacksnap replay` (taxonomy + directus) on dev build, cache-first; `--no-snapshot` escape
- [x] `dev-min` seed preset (~1 org + ~10 users + minimal activity) applied on build _(§1: `stack-seeding/presets/dev-min.seed.yaml` — 10 users/1mo/Dev Sandbox, dev@anthropos.test admin; pinned in `presets_test.go`. Bring-up application wired in §2.)_
- [ ] n=0-dev-reset guard preserved
- [ ] Delivers: seeding-spec.md (dev-min + dev auto-seed) + snapshot-spec.md (dev replay target + local Directus)

## Final review
_(filled at close)_
