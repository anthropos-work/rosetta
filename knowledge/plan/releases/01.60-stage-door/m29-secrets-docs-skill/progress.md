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
