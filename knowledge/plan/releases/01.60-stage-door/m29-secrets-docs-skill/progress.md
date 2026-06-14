# M29 — Progress

_Section checklist. Closure when all boxes land. Stub at scaffold._

## Deliverables
- [x] `corpus/ops/secrets-spec.md` authored (net-new — closes the Phase 0b KB blind area) — §A
- [x] `.claude/skills/stack-secrets/SKILL.md` (mirrors `/stack-seed`) — §B
- [x] CLAUDE.md skill-table row — §C
- [x] CLAUDE.md Key-Documentation-Locations entry — §C
- [x] CLAUDE.md Interconnected-Documentation list update — §C
- [x] `setup_guide.md` — manual-copy prose retired + line-447 TODO deleted, points to the skill — §D
- [x] `safety.md` — never-echo / `PreflightEnv`-emitting clause added (§2.9) — §D
- [x] README-index guard passes (every doc indexed) — exit 0 (indexed in `corpus/ops/README.md` + `corpus/README.md`)

## Notes
- Built straight through as 4 sections (§A spec doc · §B skill · §C wiring · §D retire-prose + safety clause).
  Phase 0b KB-fidelity GREEN (blind area = the milestone's own `delivers:` line). PR review found **zero** issues:
  HARD SAFETY value-scan clean, all load-bearing claims verified against ext code @ tag `stage-door-m28`
  (55 genes/6 repos/40-8-7/13 critical · MintedKeys×6 · StripOnNonProdKeys×3 · exit 0/1/3 · `--prod` default
  non-prod · the `check` pre-flight is wired into `/dev-up`+`/demo-up`). README-index guard exit 0.
- **Zero ext code** — M29 is rosetta-only; the ext stayed on `main` @ `9742126` (= tag `stage-door-m28`).
- **M30 field-bake** (build-from-stack-dev `Critical==100%` validation) is the next + final milestone — Fate 2
  (already owned by M30), out of M29 scope.

## M29: Final Review

Close review GREEN — **0 findings across all categories** (rosetta-only docs + skill; every load-bearing
claim re-verified against the ext engine at tag `stage-door-m28` / `9742126`). No Phase-7 fixes owed.

### Scope
- [x] All 8 deliverable boxes checked; `overview.md` In-list fully delivered; Out (M30 field-bake) = Fate-2
  M30-owned. Zero code TODO/FIXME/HACK in any M29-touched file. No silent drops.

### Code Quality
- [x] N/A — M29 ships no executable rosetta code. All `stacksecrets` CLI invocations in `SKILL.md` +
  `secrets-spec.md` verified against the real parser (`cmd/stacksecrets/main.go`): subcommands + every
  `provision`/`check`/`introspect`/`diff` flag exist. Cross-references resolve; corpus-index guard exit 0.

### Documentation
- [x] 0 findings. DNA counts verified (55/6/40-8-7/13-crit, profile `graphql`); `gh-token` alias = 3 members;
  `StripOnNonProdKeys` = 3 keys; `MintedKeys` = 6 keys — all match the docs. safety.md §2.9 anchors resolve;
  setup_guide.md line-447 TODO gone + hand-copy retired (key-lists kept per M29-D4); CLAUDE.md + both indexes wired.

### Tests & Benchmarks
- [x] N/A — docs milestone, no rosetta test stack. README-index guard exit 0 (the rosetta gate). Handbook
  counts (113 @ M27 / 160 @ M28) reconcile against state.md. Ext 160 Go tests untouched (ext @ `9742126`).

### Decision Triage
- [x] M29-D1 (pinned-tag build) → already in SKILL.md/secrets-spec.md. M29-D2 (shorthand→CLI) → in SKILL.md body.
  M29-D4 (key-lists kept) → in setup_guide.md. M29-D3 (README-index same-dir target) → maintainer-only, stays in decisions.md.

### Adversarial (Phase 2c)
- [x] 4 doc-consumer scenarios (synthesized-flag / value-leak / prod-token-rearm / stale-counts) all verified
  clean vs ext code; recorded in `decisions.md` § Adversarial review.

### Deferral re-audit (Phase 1b)
- [x] GREEN — 0 new, 0 repeat, 3 inherited backlog (carry), 0 blocking. Report:
  `audit-deferrals/deferral-audit-2026-06-14-m29-close.md`.
