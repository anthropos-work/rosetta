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

## Final review
_(filled at close)_
