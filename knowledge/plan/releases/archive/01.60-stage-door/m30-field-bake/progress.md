# M30 — Progress

_Section checklist. Closure when all boxes land. Part 1 (assemble + check) ran 2026-06-14; **Part 2 (live
bring-up) was executed the same day** — a fresh demo-3 was brought LIVE from the assembled source with the
user's go-ahead. All boxes land._

## Deliverables
- [x] Compliant `.agentspace/secrets` dir assembled from current stack-dev (names-correct, alias-mapped, knowns waived) — 5 repo `.env` files cp'd into the reader layout, values-blind; ant-academy source filled with the shared Clerk pub key (values-blind line-append)
- [x] `provision` into a fresh stack → `measure` Critical == 100% — Part-1 proof: `check --from .agentspace/secrets` → Critical **100.0%** / Overall 62.2% / exit 0. **Part-2 (live):** the demo-3 bring-up provisioned 26 written / 2 blanked / 0 skipped from the assembled source, the pre-flight scored Critical **100%**
- [x] `provision` into a fresh `demo-N` → `measure` Critical == 100% (demo-aware: Clerk minted-OK) — Part-1 proof: `check --from .agentspace/secrets --demo` → Critical **100.0%** / Overall 66.3% / exit 0. **Part-2 (live):** demo-3 pre-flight Critical 100%, provision clean
- [x] Each stack reaches UP after provisioning (the observable-behavior gate) — **MET LIVE:** demo-3 came up with **17 containers** (backend tier + UI tier: next-web + studio-desk + ant-academy), all liveness+readiness probes pass; set-dress isolation clean
- [x] Field bugs surfaced + fixed Fate-1 — **2 real bugs**: (1) `sentinel/DB_CONNECTION` was critical/required but is compose-injected config (never read from `.env`) → reclassified `waived-config` + regression test; (2) the demo bring-up only *checked* coverage but never *provisioned* (and `preflight.sh` resolved its source path one level too shallow → the demo gate silently skipped, exit 2) → added the provision step + fixed the path to `EXT_ROOT/../..` (parallels v1.5 M25's 4 catches)
- [x] Honesty residual documented (the ~10–15% waived set + why) — in `spec-notes.md`: waived classes (now incl. `waived-config`) + lean-platform-env/compose-injected/repo-local/optional standards; all residual proven non-critical
- [x] Ext tag `stage-door-m30` — branch `m30/field-bake` off `main`; the DNA fix + regression test + version bump + the provision-wiring/preflight-path fix; tagged @ `29c922b`

## Notes
- **Gate met (Part 1 + Part 2):** Critical **100.0%** on both dev and demo `check`; exit 0 — and proved LIVE: a fresh demo-3 came up (17 containers) provisioned from the assembled `.agentspace/secrets`. All 12 required+critical genes pass; every residual short proven `standard`/`optional` (zero critical).
- **Field bug 1 (the bake's catch):** `sentinel/DB_CONNECTION` — docker-compose hardwires it as a sentinel `environment:` entry (overrides `env_file`); sentinel never reads it from `sentinel/.env`, and no `sentinel/.env` exists on stack-dev. Was falsely failing the gate at 84.6%. Reclassified `waived-config`; gate then clean at 100%.
- **Field bug 2 (the live-bring-up catch):** `up-injected.sh` only ran the pre-flight *check*, never *provisioned* → the demo ran from the operator's live dev env, not the assembled source; and `preflight.sh` resolved `REPO_ROOT` as `EXT_ROOT/..` (one level too shallow — extensions live two dirs deep), so its default source doubled to a nonexistent `.agentspace/.agentspace/secrets` and the demo-aware gate **always silently skipped (exit 2)**. Fixed: added a non-fatal provision step (default-on, `DEMO_NO_PROVISION=1` opt-out) that writes stack-demo's per-repo `.env` from the source and repoints the run's base env at it; corrected the path to `EXT_ROOT/../..` + had the caller pass `--from` explicitly.
- **Safety verified LIVE:** prod `DIRECTUS_TOKEN` (len-32) armed in **ZERO** containers; cms blank; provision writes the `DIRECTUS_TOKEN` family BLANK on the non-prod target AND the injection override strips it at compose-emit (defense-in-depth, the fix16/17 non-rearm class). No secret value in any output or commit.
- **Values-blind throughout:** assembly by `cp`/line-append only; provision stdout is key NAMES + write/blank/skip counts only. No value read, echoed, logged, or committed. The `.agentspace/secrets` dir is gitignored (verified — never committed).

## M30: Final Review

Close-milestone review (2026-06-14). 4 findings — 0 scope-gap (every In-list item delivered Fate-1) · 0
code-quality (the ext fixes are correct + tested green: Go `-race`/vet/gofmt clean, 99 demo-stack pytests,
shellcheck clean) · 3 docs · 1 decision-triage. All landed Fate-1.

### Documentation
- [x] [must-fix] `corpus/ops/secrets-spec.md` was stale vs the M30 DNA fix: version `stage-door-m27`→`m30`; `sentinel/DB_CONNECTION` still shown critical; status split `40/8/7`+13-crit → `39/8/8`+12-crit; missing the `waived-config` class row; Status section described M30 as future — reconciled to the executed live bake (counts verified against the committed `secret-dna.json`)
- [x] [must-fix] `progress.md` Part-2 boxes were unchecked + "held" framing + only 1 bug — reconciled: Part 2 executed LIVE (demo-3, 17 containers, gate met), 2 field bugs
- [x] [must-fix] `spec-notes.md` Part-2 "held/PENDING" + 1-bug field-fix log — reconciled to executed-live + both bugs (bug 2 = provision-wiring + preflight path)

### Decision Triage
- [x] `decisions.md` was an empty scaffold stub — authored M30-D1 (sentinel waive), M30-D2 (provision-then-move-env-file design), M30-D3 (preflight two-levels-up path) + the Phase-2c adversarial-review subsection (5 scenarios, all handled)

### Deferral audit
- [x] Phase 1b `/audit-deferrals --scope=milestone` → **GREEN** (0 new/repeat/aged; the 2 surfaced bugs landed Fate-1; waived classes are waived-not-deferred; 3 inherited backlog items re-signed at v1.5 close). Report: `audit-deferrals/deferral-audit-2026-06-14-m30-close.md`
