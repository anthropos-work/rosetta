# M30 — Progress

_Section checklist. Closure when all boxes land. Part 1 (assemble + check, no live stack) ran 2026-06-14;
Part 2 (live bring-up) is held for a box with a live stack + user go-ahead._

## Deliverables
- [x] Compliant `.agentspace/secrets` dir assembled from current stack-dev (names-correct, alias-mapped, knowns waived) — 5 repo `.env` files cp'd into the reader layout, values-blind; ant-academy source filled with the shared Clerk pub key (values-blind line-append)
- [ ] `provision` into a fresh `dev-N` → `measure` Critical == 100% — **PART 2 (held: needs a live stack)**. Part-1 proof: `check --from .agentspace/secrets` → Critical **100.0%** / Overall 62.2% / exit 0; `provision --dry-run` plans 26 write / 2 blank / 0 skip cleanly
- [ ] `provision` into a fresh `demo-N` → `measure` Critical == 100% (demo-aware: Clerk minted-OK) — **PART 2 (held)**. Part-1 proof: `check --from .agentspace/secrets --demo` → Critical **100.0%** / Overall 66.3% / exit 0
- [ ] Each stack reaches UP after provisioning (the observable-behavior gate) — **PART 2 (held: needs a live stack + user go-ahead)**
- [x] Field bugs surfaced + fixed Fate-1 — **1 real bug**: `sentinel/DB_CONNECTION` was critical/required but is compose-injected config (never read from `.env`) → reclassified `waived-config` + regression test (parallels v1.5 M25's 4 catches)
- [x] Honesty residual documented (the ~10–15% waived set + why) — in `spec-notes.md`: waived classes (now incl. `waived-config`) + lean-platform-env/compose-injected/repo-local/optional standards; all residual proven non-critical
- [x] Ext tag `stage-door-m30` — branch `m30/field-bake` off `main`; the DNA fix + regression test + version bump; tagged

## Notes
- **Gate met (Part 1):** Critical **100.0%** on both dev and demo `check`; exit 0. All 12 required+critical genes pass; every residual short proven `standard`/`optional` (zero critical).
- **Field bug (the bake's catch):** `sentinel/DB_CONNECTION` — docker-compose hardwires it as a sentinel `environment:` entry (overrides `env_file`); sentinel never reads it from `sentinel/.env`, and no `sentinel/.env` exists on stack-dev. Was falsely failing the gate at 84.6%. Reclassified `waived-config`; gate then clean at 100%.
- **Values-blind throughout:** assembly by `cp`/line-append only; no value read, echoed, logged, or committed. The `.agentspace/secrets` dir is gitignored (verified — never committed).
- **Part 2 reported PENDING:** provision into a fresh live `dev-N` + `demo-N` (never N=0) and assert each reaches UP — held for a box with a live stack on it + the user's go-ahead.
