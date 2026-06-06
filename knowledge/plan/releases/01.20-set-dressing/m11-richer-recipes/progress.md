# M11 — Progress

**Shape:** section · **Status:** build complete

## Section checklist (from overview Scope.In)
- [x] Refreshed seed presets (small/mid/large) — documented snapshot prerequisite (presets stay structural, M11-D3); README full-fidelity flow
- [x] Updated corpus/ops/demo/ recipe family (set-dressed worlds) — both recipes refreshed + the false "waived/v1.2" note killed; NEW recipe-snapshot-world.md
- [x] /demo-snapshot skill — a NEW skill (M11-D1) driving stacksnap replay/capture/status; flags verified vs the real parser
- [x] Corpus cross-links + data-DNA-100% + CLAUDE.md skill table — skill-table row added; both specs cross-linked to the demo family; links resolve
- [x] Release-close hygiene carry — stacksnap --help directus-surface drift fixed (M11-D4)

## Build summary
- **rosetta (`m11/richer-recipes`):** 5 commits — README+presets-doc+Phase0b (§1), recipe family (§2), /demo-snapshot skill (§3), corpus cross-links+CLAUDE.md (§4); §5 is extensions-side.
- **extensions (`main`):** 2 commits — preset prerequisite headers (§1), stacksnap --help fix (§5). Tag `stack-snapshot-m11` to be set at close.
- **Tests:** docs + skill (no rosetta code). Extensions: presets parse/validate via `stackseed --dry-run` (×3); `cmd/stacksnap` tests + `go vet` + gofmt clean; module builds.
- **Phase 0b KB-fidelity:** GREEN (`kb-fidelity-audit.md`).
- **Deferrals:** none — all Fate-1 (see decisions.md three-fate ledger).

## M11: Hardening

### Pass 1 — 2026-06-06
**Scope manifest (milestone-touched code, from the 2 extensions commits `7987e1d`+`ae59150`):**
- `stack-snapshot/cmd/stacksnap/main.go` — the `--help` text fix (§5/M11-D4). Tests: `main_test.go`, `main_harden_test.go` (existing); the help string was NOT pinned (the §5 commit itself noted "no test pins the help string").
- `stack-seeding/presets/{small-200,mid-500,large-1k}.seed.yaml` — comment-only prerequisite headers (§1/M11-D3). Tests: none loaded the real preset files (verified only by hand via `stackseed --dry-run`).
- The other M11 surface (recipes, `/demo-snapshot` skill, corpus cross-links) is markdown — not a test gap (per the harden orchestration brief; do not manufacture coverage on docs).

**Coverage delta (milestone-touched files):**
- `cmd/stacksnap`: statements 77.7% → 77.7% (+0.0) — `usage` was already line-executed by `TestRun_Help`; the gap was **assertion depth** (the help string content was unpinned), not lines. Remaining uncovered statements are the live-DB capture/replay paths (integration-tested, need Postgres; not M11-touched).
- `blueprint`: the new tests load the real shipped presets (a path the suite never exercised before) — value is regression-pinning the product files, not a % move.

**Tests added:**
- `stack-snapshot/cmd/stacksnap/main_drift_test.go`: 5 — `TestHelp_NamesEveryRegisteredSurface` + `TestHelp_FrameworkTagIncludesLatestMilestone` (pin the §5 fix, registry-driven), `TestDocsFlagsExistInParser` + `TestDroppedDumpFlagStaysGone` + `TestDocumentedSourceKindsAreReal` (the docs↔parser drift guard for the `/demo-snapshot` skill + recipes). (commit `767c291`)
- `stack-seeding/blueprint/presets_test.go`: 2 — `TestShippedPresets_ParseStrictAndValidate` (each shipped preset through Load[strict `KnownFields`]+Validate) + `TestShippedPresets_SizesAreDistinctAndOrdered` (small<mid<large contract). (commit `1e18df6`)

**Bugs fixed inline:** none — no production-code bug surfaced (the §5 help fix and §1 headers were already correct; the gap was missing regression tests).

**Mutation verification (each new guard proven to bite):**
- Reverting the §5 help fix (tag→M9a/M9b, drop the directus line) → both `TestHelp_*` FAIL.
- Injecting a bogus documented flag → `TestDocsFlagsExistInParser` FAILS.
- Injecting an unknown field into a preset → `TestShippedPresets_ParseStrictAndValidate` FAILS; a pure comment line → PASSES (proves the M11 §1 comment-only header style is safe).

**Flakes stabilized:** none seen. Flake gate: 3× sequential clean on both new test sets; both packages pass `-race`; both modules `go build ./...` clean; gofmt + `go vet` clean.

**Knowledge backfill:** no KB-worthy findings — the harden pinned existing documented behavior (the help contract + the preset schema), surfacing no new invariant/edge-case/error-path. The docs↔parser contract is already self-documenting in the `/demo-snapshot` skill + `corpus/ops/snapshot-spec.md`; the new tests reference those docs in-comment.

### Stop condition
Loop stopped after Pass 1 (expected short pass for a docs/discoverability milestone): the Step 2b scan found exactly the two real Fate-1 gaps (the unpinned help contract + the unloaded shipped presets, plus the docs↔parser drift), all closed; coverage delta negligible by design (assertion-depth, not line, was the gap); zero flakes. Remaining touched surface is markdown (not testable) and live-DB paths (out of hermetic scope, integration-covered).

## Final review
_(filled at close)_
