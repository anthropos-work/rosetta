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
  prose) — OUT of M247's doc-only scope (all four are rext files, **0 tracked in the rosetta repo**).
  **RE-FATED at close** (the original D0 "→ M251/rext-hygiene" is stale — M251 CLOSED 2026-07-23 without
  owning them): **documented-inert standing rext-hygiene note** (the load-bearing comment was fixed in M246;
  the remainder is a dormant/never-consumed key + passing test fixtures + one audit-prose line — **0
  functional impact**), swept opportunistically by whichever rext milestone next edits those files
  (M249 owns `up-injected.sh`; M252 edits `gen_injected_override.py`) — no sibling `overview.md` edit.
  decisions.md **D-audit** (fresh close-time decision) supersedes D0's routing.
- **ai-readiness demo-seeder fidelity** (31-skill demo set, track-keyed named sims, evaluated-skills set-dress) +
  the D-07 demopatch re-pin + the 4 demo-section compute line-anchors → v2.7 **M250** (which also delivers
  ai-readiness.md). Fate-2 (already planned). decisions.md D1.
- **ops/demo spec-doc reconcile** (content-stories-spec/routes, demopatch-spec, cockpit-spec, latency-budget,
  secrets-spec, studio parts of frontend-tier/studio-desk) → owned by M248/M249/M250/M252/M253 + the release-close
  consistency pass. Fate-2 release-close item. decisions.md D-fate-2.
- **optional `verification.md` demo-stack-suite + run-unit-roster index anchor** (inherited — M251 punted it
  to M247 "Fate-3-adjacent" for lane-collision avoidance) — **NOT landed in M247** (out of the consolidation
  charter; a test-health indexing concern, M251's domain). Explicitly **OPTIONAL + non-blind** ("the code it
  would index exists + is exercised") → **Fate-2 release-close consistency pass** (same bucket as the ops/demo
  reconcile). decisions.md **D-audit**.

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

## M247: Final Review

_(close-milestone review — 2026-07-23. M247 is a **doc-only** `section` milestone (0 source files); the
code-shaped review dimensions collapse to doc fidelity + integrity + decision-triage.)_

### Scope
- [x] All 8 `## Sections` checkboxes delivered (skillpath redirect · 3-subgraph sweep · 4 fact sheets ·
  ai-readiness refresh · roadrunner resolution). roadrunner resolved to the **negative** (stays ORPHANED,
  not archived — still in repos.yml+compose) — a recorded decision, not a silent drop.
- [x] TODO/FIXME/HACK scan of touched corpus files → 0 real work-markers (only `XXXX` ID-format placeholders).
- [x] Inherited hand-off: M251's optional `verification.md` anchor was NOT landed → re-fated (Fate-2 release-close). See Deferred ledger + decisions.md D-audit.

### Code Quality
- [x] N/A — doc repo, 0 source files; no lint/type stack. Only non-prose change is `.gitignore` (worktree-symlink fix).

### Adversarial
- [x] `.gitignore` `/stack-*` (no trailing slash) over-match — verified 0 tracked root paths named `stack-*`, so the new symlink-form pattern ignores nothing currently tracked. (recorded in decisions.md § Adversarial review)

### Documentation (the deliverable itself)
- [x] Fidelity spot-check vs current `stack-demo/app` source — GREEN, 0 corrections (harden Pass 1).
- [x] Cross-reference integrity — **0 broken relative `.md` links** across 30 M247-touched docs (re-confirmed at close).
- [x] Per-unit handbook contract — all 4 new fact sheets exist + indexed in `services/README.md` **and** `CLAUDE.md`.
- [x] grep-verify — 0 residual "4 subgraphs" (except the intentional historical note in `backend.md:67`); 0 stale live-skillpath-as-service claims.

### Tests & Benchmarks
- [x] N/A — no test runner / coverage / benchmark stack in a documentation repository.

### Decision Triage
- [x] D0 / D1 / D-fate-2 → **archive** (maintainer-only close-routing decisions; the platform-facts they gate already live in the delivered docs).
- [x] KB-1..KB-7 (Phase-0b KB-fidelity findings) → already reflected in the delivered corpus (no separate blend needed).
- [x] **D-audit** (new) → re-fate the aged-out rext-hygiene set (M251 destination closed) + the inherited verification.md anchor.

### Verdict
All findings GREEN or bookkeeping-only. One substantive fix applied (the D-audit re-fate). No escape-hatch
deferral → no sign-off gate. Proceed to merge.
