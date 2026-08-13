# AUDITOR A — 7 files / 1726 lines

**Positive control:** all 7 read to final line; counts match `wc -l`
(external_services 798 · backend 271 · graphql-wundergraph 265 · academy-backend 141 · coursebuilder 139 ·
skiller 66 · TEMPLATE 46).

## BLOCKERS — 0

Every load-bearing claim resolved against platform source. Notably re-verified **correct**:

- The subgraph ladder **5 → 4 → 3 → 1**: `git show <sha>:supergraph-config-prod.yaml | grep -c name:` →
  `749dc86^`=5, `749dc86`=4, `7c17e63`=3, `915da06`=1; `git show --name-status 915da06` marks **both**
  `schemas/cms.graphqls` and `schemas/jobsimulation.graphqls` `D`. The commit subject's "2→1" *is* wrong,
  exactly as both docs say.
- No `graphql` and no `directus` service in `docker-compose.yml` @ `2adcf71`; `grep 5050` → 0 hits.
- The nine-container `graphql` profile; `backend`'s `environment:` has no `DIRECTUS_*`, `cms` does.
- `main.go:1196-1202` carries the "DORMANT … until the M809 re-point" comment verbatim.
- **0** `SkillPathSessionService` occurrences in app Go source.
- Taxonomy manifest: skills **42,790**, job_roles **22,470**, embeddings 42,790 / 18,919 — every figure matches.
- `credits/cost.go:86-90` build **5** / refine **1** / translate **1** — the corpus is right and the repo's
  own `:29` header comment is the stale one.

## MINORS — 10

| # | site | what is off |
|---|---|---|
| 1 | backend.md:253 | "not in the top-level `migrations/` dir (**which holds only `atlas.sum`**)" — there is **no** top-level `migrations/` dir at `5ba17044`. The actionable half (migrations live in `terraform/migrations/`) is correct |
| 2 | backend.md:190 | LabSession "registered as a **third** RPC handler" — at HEAD it is **last** (`:1218-1219`). Same file's `:39` already has the correct order |
| 3 | backend.md:39 | anchor range ends `:1218`; the `mux.Handle` is `:1219` |
| 4 | external_services.md:348 | "`5050` was **only** the local compose host mapping" — `graphql-wundergraph/Makefile:11` still maps `-p 5050:8080`. (The gateway doc `:136` flags `make run` separately) |
| 5 | external_services.md:245, :710 | `DIRECTUS_PUBLIC_BASE_ADDR` presented as a `platform/.env` var; **0** occurrences in `.env_example`. `app` does read it (`main.go:1045`) and compose sets it on `cms` — guidance sound, example file doesn't ship it |
| 6 | external_services.md:748 | anchor `backend.go:130` is the authn **skip-list** entry, not the route registration (only occurrence in that file) |
| 7 | graphql-wundergraph.md:128 | self-ref "as `:84` describes" — `:84` is a different note; the text meant is `:102-111` |
| 8 | graphql-wundergraph.md:82 | self-ref `:174-176` — the sentence is at `:178` |
| 9 | graphql-wundergraph.md:99 | workflows list omits `bump-version.yml` |
| 10 | coursebuilder.md:113-114 | "~55" / "~25" test files — actual **59** and **32** (both hedged with "~") |

## Files read clean

- **`skiller.md`** — every claim verified, incl. the ≥42,790 / ≥22,470 floors against the manifest.
- **`TEMPLATE.md`** — pure template, no factual claims.
- **`academy-backend.md`** — table names, cert format/alphabet, GraphQL names, token group, nightly refresh,
  both `cmd/` flag sets, `paid` default — all verified.
