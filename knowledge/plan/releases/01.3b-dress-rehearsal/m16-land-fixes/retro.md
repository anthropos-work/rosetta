# M16 — Retro

_Closed 2026-06-08. The first milestone of v1.3b "dress rehearsal" — a docs/publish/rename milestone that landed the honest baseline._

## Summary
M16 made two pre-applied field fixes (ISSUE-1 devpath rename-resolution + ISSUE-7 migrate-race `set -e` resilience) **durable, public, and fenced**: pushed the stranded extensions commits to `origin`, finished the `anthropos-dev → stack-dev` rename as the documented default (with a single intentional legacy-alias fallback), cleared the stale demo-stack GUIDE/README facts + the pytest doc, refreshed the extensions `knowledge/` KB, and added a consolidated stack-dev layout + back-compat note to `corpus/ops/rosetta_demo.md`. 8 harden guards now pin the rename contract, the GUIDE doc-truth, and the migrate-race fix. Close was near-clean: 1 Fate-1 doc fix.

## Incidents This Cycle
- **None.** No P2 flakes (flake gate 5/5 clean on the demo-stack suite), no regressions (full Python 182/182, all 4 Go modules green), no rollbacks. The harden count-guard caught its own drift twice (13→18→21) and the GUIDE was corrected each time — a designed self-check, not an incident.

## What Went Well
- **The contract-guard discipline paid off.** `TestGuideDocTruth`'s count guard immediately failed on the +5 then +3 test-count drift and forced the GUIDE to match the live suite — a doc-truth milestone whose docs literally cannot drift from code now.
- **The dual-repo split stayed clean.** All functional + doc changes lived in the nested `.agentspace/rosetta-extensions` repo (pushed + tag-reconciled + re-consumed); the rosetta branch carried only the corpus note + tracking docs. No cross-repo leakage.
- **Deterministic rename resolution.** The `[ -d "$stack-dev" ] || …anthropos-dev` pattern is uniform across all 5 resolvers and pinned by `TestRenameDrift` — the both-roots-exist adversarial case resolves to `stack-dev` with no silent-wrong-root mode.

## What Didn't
- **The build-time boundary call (M16-D5) was too conservative.** The clerkenstein glossary's stale `anthropos-demo/` reference was deferred at build as cross-section scope-expansion risk, then landed Fate-1 at close (M16-D8) as a one-line fix — it was the *last* stale workspace name in the repo and squarely in M16's doc-truth class. Lesson: for a "restore doc truth" milestone, the bar for "out of scope" on a one-word doc fix should be higher; a repo-wide grep at build would have surfaced it as a clean Fate-1 then.
- **The tag trailed HEAD (twice).** The milestone tag was set at the build HEAD, then trailed the 2 harden commits AND the 1 close commit — requiring a close reconcile `44edc09 → e6161b0`. This is the documented tag-reconcile step (cf. v1.3 M9a) and worked correctly, but it recurs every milestone; the close-time reconcile is now a reliable habit rather than a surprise.

## Carried Forward
- **The live docker-harness migrate-race BEHAVIOR test → M17** (Fate-2, M16-D7). M16 landed the static regression fence (`TestMigrateRaceGuard`); the runtime behavior test belongs with M17's race/idempotency work, which brings the docker harness + the per-component tested idempotency contract. No M17 overview edit — already owned.
- **DEF-M10-01** (S3 blob bytes + cloud store → v1.4, signed) — inherited, orthogonal to M16, not aged. Unchanged.

## Metrics Delta
_(from metrics.json)_
- **Python test funcs:** 174 → **182** (+8: demo-stack 13 → 21 — `TestRenameDrift` 3 + `TestGuideDocTruth` 2 + `TestMigrateRaceGuard` 3).
- **Go test funcs:** **713** unchanged (M16 touched no Go; the close Fate-1 edited a markdown doc).
- **Flake:** 0 (5/5 sequential on the demo-stack suite).
- **Lint:** shellcheck + py_compile CLEAN on all touched scripts.
- **Extensions:** tag `dress-rehearsal-m16` @ `e6161b0` (reconciled from `44edc09`); `stack-demo/rosetta-extensions` re-consumed; `main` `a31d70b..e6161b0` on `origin`.
