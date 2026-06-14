# M27 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `stack-secrets/` section scaffolded (go.mod + cmd/stacksecrets) in the `.agentspace` authoring copy
- [x] Source-ingestion reader: directory mode (default `.agentspace/secrets`), values-blind
- [x] Source-ingestion reader: zip mode
- [x] Source-dir layout contract (zEnvs / per-repo .env never silently ingested)
- [x] secret-DNA schema + `secret-dna.json` (gene = repo×KEY; the per-gene fields; strict load + Validate)
- [x] `introspect` from the hybrid source (platform/.env_example + frontend/sentinel .env.example + compose-injected)
- [x] `list` verb
- [x] `diff` verb — required-key drift exit 1 (the "keep-listed" gate) + undeclared-runtime-required guard
- [x] Waived classes modeled (AWS-mount, profile-gated, optional Bunny/GCloud)
- [x] Alias families encoded (GH_PAT family) vs distinct-similar pairs (OPENAI_KEY/OPENAI_API_KEY — not auto-aliased)
- [x] Hermetic unit tests (no values)
- [x] Ext tag `stage-door-m27`

## Notes
- Built stdlib-only (no pgx/yaml): the secret-DNA is JSON, the readers are `archive/zip` + `bufio`.
- The committed `secret-dna.json` is **55 genes across 6 repos** (platform, app, sentinel, studio-desk,
  next-web-app, ant-academy); 94 tests across the three packages, `-race` + `gofmt` clean.
- A `check`/`measure` verb (score a source; exit 1 if critical < 100%; per-repo rollup) was folded in too —
  the natural pairing with the DNA, exercised end-to-end against the real stack-dev (values-blind).
- **M27-D2** (in `decisions.md`): the keep-listed gate is DNA-scoped two-tier (gate-fatal only on
  already-tracked-secret omissions; never-tracked declared keys → triage candidates), so the gate is usable
  against real `.env.example` files that mix secrets with config noise. The diff-vs-stack-dev caught 10 real
  cross-repo DNA omissions → fixed Fate-1.
- Verified live: `stacksecrets diff --stack-root stack-dev` exits **0** (0 gate-fatal); `check` against
  stack-dev reports real coverage + per-repo shortfalls with **no secret value printed**.

## M27: Hardening

### Pass 1 — 2026-06-14
**Scope manifest** (the whole new `stack-secrets/` section — single Go stack; every source file ships a
co-located test). Touched source → existing test:
- `source/source.go` → `source/source_test.go`
- `secretdna/{dna,source,operators,measure,catalog,introspect,diff}.go` → each `*_test.go` + `fake_test.go`, `secret_dna_json_test.go`
- `cmd/stacksecrets/main.go` → `cmd/stacksecrets/main_test.go`
- No source file lacked a test. Baseline coverage: cmd/stacksecrets 87.4%, secretdna 96.9%, source 91.0% (total 93.6%); all green, `-race` + `gofmt` clean.
- No new-unit-without-handbook finding: the section ships `stack-secrets/README.md`.

**Coverage delta (milestone-touched files, start → end):**
- Statements (total): 93.6% → 98.3% (+4.7)
- By package: cmd/stacksecrets 87.4% → 97.9% (+10.5); secretdna 96.9% → 99.2% (+2.3); source 91.0% → 96.2% (+5.2)

**Tests added (Pass 1 — `*_harden_test.go` in all three packages):** 27 tests:
- `source`: zip suffix-match decoy defence (evil-app/myapp/notstudio-desk must NOT resolve), `/`-bounded wrapped-dir positive, uppercase `.ZIP` dispatch, zip dir-entry skip, corrupt-zip error, **values-never-stored** (200 KB value + `=`-in-value + quote-only), split-on-first-`=`, empty/quote-only → empty shape, `unquote` unbalanced/mismatched, `memSource.NonEmpty` trichotomy, target-is-dir parse error.
- `secretdna`: diff multi-class stable ordering (unlisted-required → undeclared → candidate, repo/key tie-break) + counts, `driftRank` unknown-kind-last, empty-declared → all-undeclared (gate stays closed), alias **distinct-similar stays standalone**, 3-member family OK, singleton-among-valid rejected, `looksLikeJWT` segment-count/charset edges, oversized-line scanner error, curated invalid-key skip, sorted-deterministic keys, `Save` write-error, unknown-operator-fails, `FormatMatch` absent-key.
- `cmd/stacksecrets`: `introspect --write` is report-only (DNA byte-identical — guards M27-D3.3), two-tier label surfacing, per-verb bad-flag → usage, invalid-DNA-via-CLI / unparseable-DNA-via-CLI, `-h`/`--help` variants, **end-to-end values-blind regression against the REAL 55-gene DNA** (`check` + `diff` with sentinel values → zero leakage to stdout/stderr).

**Tests added (Pass 2 — error-path completeness, appended):** 13 tests:
- `cmd`: diff/introspect `ReadDeclaredKeys` read-error → `exitFail` (fail-loud, not false-clean), `check` missing-targets note, `check` empty `--dna`.
- `secretdna`: `ReadDeclaredKeys` non-NotExist propagation, missing-file IsNotExist, `Catalog` no-alias / no-source branches, `KeyPresent` direct.
- `source`: `FromZip`/`FromDir` oversized-line parse error, `FoundTargets` sort order, `TargetsFromDNA` waived-only-repo → no target.

**Bugs fixed inline:** none — the build-phase code was solid; every new test passed against the shipped implementation (the harness's error paths, layout contract, and values-blind extraction all held under adversarial input).

**Flakes stabilized:** none observed. Flake gate: 3 consecutive sequential clean runs of the new tests; full suite `-race` + `gofmt` clean.

**Knowledge backfill:** no KB-worthy NEW findings — hardening confirmed (did not newly discover) the invariants already documented in `decisions.md` (M27-D2 two-tier gate, M27-D3.3 introspect-never-mutates) and the section `README.md` (values-blind, layout contract). The values-never-stored + zip-suffix-decoy-defence behaviors are now pinned by regression tests; the corpus doc for the skill lands in M29 (per the milestone's repo-split — rosetta gets no doc edits in M27 beyond this progress file).

### Stop condition
Stopped after Pass 2 (2 of the 5-pass cap). The full six-dimension re-scan found nothing new worth adding: the milestone's risk surface (layout contract / zEnvs-stray defence, values-blind extraction, the DNA-scoped two-tier diff gate, alias-family vs distinct-similar, the fail-loud error paths) is now thoroughly covered. The Pass-2 statements delta was +1.0%, and the only remaining uncovered lines are near-impossible defensive branches with no behavioral value — `main()`'s `os.Exit(run(...))` wrapper, post-successful-read `f.Close()` error returns, and a `json.MarshalIndent` failure on a valid struct — which the skill's "coverage is a finder, not a goal / no shallow tests" guidance explicitly says not to chase with contrived harnesses.
