# Release Review: v1.3 "stack party"

**Date:** 2026-06-07
**Milestones:** M12 (unified registry) · M13 (dev peers) · M14 (unified stack-* skills) · M15 (safety doc)
**Reviewer process:** Phases 0/4/4b run in-thread (deterministic); Phases 1/1b/2/3/3b/5 by 3 parallel review agents.

## Verdict summary
| Phase | Area | Verdict |
|---|---|---|
| 0 | Supply-chain | **GREEN** |
| 1 | Scope | **GREEN** |
| 1b | Deferral re-audit | **YELLOW** (escape-hatch needs a formal signed decision + date fix; not blocking) |
| 2 | Code quality | **GREEN** (0 must-fix) |
| 3 | Documentation | **YELLOW** (2 should-fix: recipe v1.3→v1.4 drift + missing safety.md index rows) |
| 3b | KB consolidation | **GREEN** (1 should-fix: CHANGELOG stale "→v1.3") |
| 4 | Tests | **GREEN** |
| 4b | Metrics | **GREEN** (+20 Go, +52 Python, no regression) |
| 5 | Decisions | **GREEN** |

**No must-fix blockers.** All open items are should-fix doc edits + one escape-hatch sign-off formalization.

## Scope
- [x] All 4 milestones delivered every `Scope.In` item Fate-1; 0 Fate-3 routings; 0 new deferrals. (verified)
- [ ] [should-fix] S-1 — note the DEF-M10-01 destination chain explicitly: signed Fate-2 → v1.3 at v1.2 close, then user re-scoped → v1.4 on 2026-06-07. The "signed v1.2" shorthand elides the one forward move.

## Deferral re-audit (Phase 1b — YELLOW)
The sole deferred item is the inherited **DEF-M10-01** (S3 media blob bytes + cloud SnapshotStore backend). It is genuinely Fate-1-infeasible (needs eu-west-1 S3-read creds the project lacks — verified) and its v1.4 home exists in `roadmap-vision.md`. BUT:
- [ ] [should-fix] 1b-1 — **No `RELEASE-SCOPE-DEFER:` signed decision exists.** It was recorded as Fate-2 (explicitly *not* escape-hatch) at M10/v1.2, then carried as an implied escape-hatch in v1.3 without formalization. Record a fresh dated `RELEASE-SCOPE-DEFER: DEF-M10-01 → v1.4` decision with the "why Fate 1/2/3 failed" rationale + user sign-off.
- [ ] [should-fix] 1b-2 — **Acknowledge the v1.3→v1.4 forward move as a documented `DRIFT_DEFER`** (destination updated forward), not "no repeat" as the milestone audits asserted. It is its 2nd release deferred; the move was a single dated user decision (2026-06-07), not a chronic punt.
- [ ] [should-fix] 1b-3 — **Fix the M15 close audit's `first_deferred_on` date** (`2026-06-07` → `2026-06-06`, the M10 close date).

## Code Quality
- [ ] [should-fix] CQ-1 — `.claude/skills/stack-list/SKILL.md` over-claims output: promises "offset ports, profile, health, resolved per-repo refs" / a `clones` field, but the unified registry record (`stack_registry.py`) carries only `type, n, ports, status, created`. Profile/clone-refs live in demo-stack's separate legacy provenance registry (demo-only, no dev-N equivalent), which `stack-list` does not read. Trim the description/body to what the row carries (or note profile/clone-refs come from `rosetta-demo status`'s demo-only provenance section).
- [ ] [nice-to-have] CQ-2 — `dev-stack/dev-stack:21` stale path comment `anthropos-demo` → `anthropos-dev` (copy-paste residue).
- [ ] [nice-to-have] CQ-3 — `stack-snapshot/directus/provision.go:106` "illustrative offset" comment undersells that the runner's `--check-env` path makes the delta load-bearing.

## Documentation
- [ ] [should-fix] D-1 — recipe-family v1.3→v1.4 version drift: `corpus/ops/demo/recipe-snapshot-world.md:60,116,120` + `recipe-skill-progression.md:66` still stamp cloud-store/S3/AI-content/shareability as the "v1.3" theme; all moved to v1.4.
- [ ] [should-fix] D-2 — `safety.md` (the M15 headline) is absent from both corpus index tables: `corpus/ops/README.md` "Available Operations" + `corpus/README.md` ops listing. Add the rows.
- [ ] [should-fix] D-3 — root `README.md:5` headline banner reads "v1.0 + v1.1 + v1.2 — shipped" while the body already ships v1.3-converged content; extend the banner to v1.3 (close-release tags it).
- [ ] [nice-to-have] D-4 — `corpus/ops/rosetta_demo.md:41` "the stack-* skill set … lands in M14" future-tense for a shipped milestone.
- [ ] [nice-to-have] D-5 — `corpus/README.md` "Demo environments" block stale at v1.1 (4-step flow, omits snapshot-spec/db-access/safety) — fold into the D-2 pass.
- [ ] [nice-to-have] D-6 — `seeding-spec.md:13` "consumed by the `/demo-*` skills" → unified `stack-seed`/`stack-snapshot`.
- [ ] [nice-to-have] D-7 — on-disk `CLAUDE.md:140` "/demo-* and /align-* skills" pre-convergence (add /dev-*, /stack-*); `:123` progress pointer.

## KB Consolidation (Phase 3b)
- [ ] [should-fix] KB-1 — `CHANGELOG.md:52` (v1.2 "Known limitations") says blob bytes + cloud store "deferred to v1.3"; they're in v1.4 and the v1.3 Unreleased section never corrects it. The v1.2 entry is immutable (M14-D6) — add a forward-note / `→ now v1.4` pointer in the v1.3 Unreleased "Known limitations". (Mirrors 1b/D-1.)
- [ ] [nice-to-have] KB-2 — `knowledge/plan/context.md` (lines 7,13,47) lags `state.md` (still "v1.3 IN DEVELOPMENT"). state.md is the declared source of truth, so optional; a one-line refresh is cleaner.

## Tests & Benchmarks
- [x] All 4 Go modules pass `-race -count=1` (exit 0); 174 Python pass; gofmt/vet/shellcheck/py_compile clean; flake 0. (no gaps)

## Decision Consolidation
- [x] All 18 decisions (M12–M15) recorded with rationale; every Q→D mapped; triaged blends verified landed on disk; no cross-milestone conflicts. (verified)
- Pattern note (not a defect): the "new deliverable shipped without its index row" miss recurs at the corpus level (D-2 = same class as M10/M13's per-unit README misses). Worth a one-line process note: extend the index-row check to corpus directory READMEs, not just per-unit extension READMEs.

## Metrics (Phase 4b — GREEN)
- Go test funcs 693 → **713** (+20, test-only matcher; stack-snapshot +15, stack-seeding +5; alignment+clerkenstein unchanged). Python **174** (+52 net new from M12's unified-registry suite). Coverage flat/up on continuing logic packages. Flake 0. No regression.

## Supply-chain (Phase 0 — GREEN)
- govulncheck: **0 called third-party CVEs**. 12 called vulns (clerkenstein/stack-snapshot/stack-seeding) are all **Go stdlib @go1.25.3** (local toolchain) → cleared by go1.25.11+ build toolchain (highest fixed-in `crypto/x509@go1.25.11`). Same class + remediation as v1.2. Licenses all permissive (MIT/BSD-3/ISC/Apache-2.0), no GPL/AGPL. Lockfile: `dependencies.lock` (20 external modules).
