# M247 — Progress

## Sections
- [x] skillpath.md → merged-into-app REDIRECT (mirror `skiller.md`) + moved to the README archived/merged table
- [x] 3-subgraph reclassification sweep — "4 subgraphs" → "3 (backend/app, jobsimulation, cms)" + skillpath reclassified not-a-live-service across the ~30 echo files; jobsim-in-app-coming note
- [x] Fact sheet — `coursebuilder.md`
- [x] Fact sheet — `ai-labs.md` (AI Labs + credits / v6.0 shared purse)
- [x] Fact sheet — `askengine.md` (Talk-to-Data)
- [x] Fact sheet — `academy-backend.md` (app-owned academy domain)
- [x] Refresh `ai-readiness.md` (aireadiness-package refactor)
- [x] roadrunner ORPHANED→ARCHIVED resolution — RESOLVED: stays ORPHANED (still in repos.yml+compose, not dead), NOT archived

## Completeness Ledger

### Deferred
- **rext-file drift** (D-02/03/04 gen_injected_override/test_injection/exposure_claim_guard; D-06 up-injected.sh
  prose) — OUT of M247's doc-only scope → M251/rext-hygiene (decisions.md D0). Fate-3.
- **ai-readiness demo-seeder fidelity** (31-skill demo set, track-keyed named sims, evaluated-skills set-dress) +
  the D-07 demopatch re-pin + the 4 demo-section compute line-anchors → v2.7 **M250** (which also delivers
  ai-readiness.md). Fate-2 (already planned). decisions.md D1.
- **ops/demo spec-doc reconcile** (content-stories-spec/routes, demopatch-spec, cockpit-spec, latency-budget,
  secrets-spec, studio parts of frontend-tier/studio-desk) → owned by M248/M249/M250/M252/M253 + the release-close
  consistency pass. Fate-2 release-close item. decisions.md D-fate-2.

### Dropped
- (none)

## M247: Hardening

### Pass 1 — 2026-07-23
**Scope manifest:** M247 is a **doc-only** milestone — `git diff 99d597c...HEAD` = 35 `.md` files + `.gitignore`,
**0 source files**. Project Rosetta is a documentation repository (no `package.json`/`go.mod`/`pyproject.toml`/
test runner/coverage tool). The code-hardening dimensions (unit/integration/e2e tests, coverage instrumentation,
fuzzing, benchmarks) are **N/A — there is no code to test**. This is not a coverage-unavailable deferral; there is
no source stack. The meaningful hardening for a doc deliverable is **fidelity + integrity**.

**Doc-robustness checks (the code-test analog):**
- **Fidelity spot-check vs current source** (`stack-demo/app` @ `v1.351.1` — the analog of "does the test exercise
  real behavior") — **GREEN, 0 corrections needed** across all 5 sheets:
  - coursebuilder: `coursebuildersession.go` ent + migration `20260717151144.sql` + `CB_AUTHOR_MODEL`/`CB_GRADER_MODEL` present.
  - ai-labs: `organization_credit`/`credit_transaction`/`lab_session`/`lab` ent schemas present; `DefaultSeedBalance int64 = 500` exact; **no `checkout.session.completed` handler** (confirms the load-bearing "shared purse unbuilt" caveat).
  - askengine: model `eu.anthropic.claude-sonnet-4-6` exact; `ask_conversation`/`ask_message`/`ask_query_example` ent; `/ask/stream`+`/ask/conversations` routes.
  - academy-backend: `academy_chapter_progress`/`_certificate`/`_chapter_body` ent + `academy.graphqls` + `cmd/academy-seed`.
  - ai-readiness: `internal/aireadiness/{defaults,scoring,readiness}.go` present; `defaults.go` = **31** `K-` node-ids (matches "19 core + 12 enabling"); `scoring.go` `archetypeHighBand = 75` + `archetypeLowCeil = 50` (matches the corrected bands).
- **Cross-reference integrity** (the analog of "imports resolve"): corpus-wide broken-link scan — **0 broken links
  in any M247-touched file**; the 19 pre-existing `.agentspace/rosetta-extensions/*` source-pointer links live in
  non-M247 files (playthroughs.md, content-stories-spec.md).
- **Completeness grep-verify** (the analog of "no dead code / leftover TODOs"): **0 residual "4 subgraphs"**
  (except the intentional historical "4 as it stood then → 3" note in backend.md); **0 genuine stale live-skillpath
  claims** (the 1 regex hit is the correct archived/merged redirect row).
- **Internal-consistency of counts**: "3 subgraphs" propagated across 9 docs; services README = 27 rows (matches
  the CLAUDE.md "27 service docs"); all 4 new sheets linked from the CLAUDE.md app bullet + the README index.

**Tests added:** N/A (doc-only; no code). **Bugs fixed inline:** 0 (fidelity GREEN). **Flakes stabilized:** N/A.
**Coverage delta:** N/A (no code stack).

**Knowledge backfill:** N/A — the corpus IS the deliverable; the fidelity findings confirmed the authored docs
rather than surfacing new invariants to backfill.

### Stop condition
One pass: the fidelity + integrity + consistency checks all came back GREEN with 0 corrections — a doc-only
milestone with no code stack has no further hardening dimensions to loop on. **Cleanup: 0 orphan processes; 0 temp
clutter** (no test runners/dev servers launched; research agents completed and returned).
