# Hardening Ledger — M51 AI-readiness showcase org

## Pass 1 — 2026-07-01 — final

**Iters hardened this pass:** all milestone-touched code (cumulative-scope final pass; gate MET in iter-09)
**Tiks covered since prior pass:** all iters in milestone (first harden pass — no prior ledger)
**Coverage delta on touched files:**
- stack-seeding/seeders (package): 96.5% -> 97.0% stmts
- ai_readiness_funnel.go `aiReadinessStepScore`: 66.7% -> 100.0%
- ai_readiness_funnel.go `aiRound`: 66.7% -> 100.0%
- ai_readiness_funnel.go `aiJSONStringArray`: 76.9% -> 100.0%
- ai_readiness_funnel.go `flush`: 87.1% -> 100.0%
**Tests added:**
- funnel helpers -> ai_readiness_harden_test.go: 6 edge-case grids (StepScore/Round boundary + bounds-invariant, Bucket threshold, Archetype quadrant), 1 jsonb round-trip fuzz (aiJSONStringArray incl. quote/backslash/unicode escapes), 4 error-path (flush sessions/evidence/progress/snapshots fault injection)
**Bugs surfaced + fixed inline:** none — the flush error-wrap arms (seeder-name + failing-surface) were already correct; the pure helpers were already bounds-safe.
**Flakes stabilized:** none
**Cross-iter integration findings:** pending (Pass 2 covers config skill-pool + member_languages native-English branch + cmd/stackseed --reload-sentinel + the TS coverage-manifest/section-assert subsystem)
**Stop condition:** continue-to-next-pass — readAIReadinessSkillPool (90%) query-error arm, languageRowsForMember (96.7%) native-English branch, cmd/stackseed --reload-sentinel path (58.2%), and the stack-verify TS coverage-manifest/section-assert helpers still unswept

## Pass 2 — 2026-07-01 — final

**Iters hardened this pass:** all milestone-touched code (cumulative-scope final pass; Pass-2 targets = config skill-pool query-error arm, member_languages native-English + both non-English proficiency arms, cmd/stackseed --reload-sentinel gate, the TS coverage-manifest AI-readiness manager contract)
**Tiks covered since prior pass:** all iters in milestone (final-mode cumulative scope)
**Coverage delta on touched files:**
- seeders `readAIReadinessSkillPool`: 90.0% -> 100.0% stmts (the query-error arm — a faulted skiller read returns an EMPTY pool, no crash/fabrication)
- cmd/stackseed `shouldReloadSentinel` (extracted pure gate): new -> 100.0% (8-case truth table; prod + dev-N=0 + seed-failed suppression proven)
- seeders `languageRowsForMember`: 96.7% -> 96.7% stmts (both English-proficiency arms [professional-4 default + the deterministic %3==0 native-5] now PROVEN reachable via a 60-member non-English sweep; the residual is a single defensively-unreachable `add` early-return guard — see below)
- cmd/stackseed (package): 58.2% -> 58.4% stmts (the residual is `reloadStackSentinel`'s live-RPC + docker-restart side-effect body — correctly untestable in a unit; its GATE is now 100%)
- seeders (package): 97.0% -> 97.0% stmts (held; the new tests deepen already-counted lines)
**Tests added:**
- ai_readiness_harden_test.go: +1 config query-error (readAIReadinessSkillPool empty-on-fault) + 1 whole-seeder graceful-degradation (AIReadinessConfigSeeder no-taxonomy) + 1 native-English/unknown-city branch + 1 non-English 60-member proficiency-arm sweep (both English levels + per-slot level invariants + distinct-FK invariant)
- cmd/stackseed/main_test.go: +1 exhaustive shouldReloadSentinel truth table (8 cases)
- stack-verify/e2e/tests/coverage-manifest.unit.spec.ts: +3 AI-readiness manager-manifest contract (page-on-seedPaths+descriptor, funnel 3-stage-labels + iter-09 mutual-exclusive "Steps completion" header [forbids re-adding "Stage breakdown"], org-score HeroCard copy)
**Bugs surfaced + fixed inline:** none functional. One refactor-for-testability (Fate 1, ~12 lines): extracted the inline `reloadSentinel && runErr==nil && !target.IsProd && n>0` gate in doSeed into the pure `shouldReloadSentinel` helper so the safety boundary (never-prod, never-dev-N0, never-after-failed-seed) is unit-pinned. Behavior identical.
**Flakes stabilized:** none
**Cross-iter integration findings:** the AI-readiness manager contract now spans seed (config+funnel seeders, iter-03/07) AND verify (the coverage-manifest funnel-header fix, iter-09) — the new TS contract test pins that the iter-09 "Steps completion" mutual-exclusion fix stays in lock-step with what the manager dashboard actually renders (a regression that re-requires "Stage breakdown" is now caught in CI with no demo up). The `languageRowsForMember` `add` guard at member_languages.go:208 (`code=="" || seen[code]`) is defensively unreachable via the public function — every call site pre-checks (native falls back to "en"; the second-English add is `native != "en"`-gated; the third-lang add is `!seen[cand]`-gated) — so its residual non-coverage is accepted, not a gap.
**Knowledge backfill:** none — no new edge-case/error-path semantics beyond what the seeding-spec/coverage-protocol docs already state; the reload-sentinel gate rationale lives in the new helper's doc comment.
**Stop condition:** continue-to-next-pass — Pass-2 targets all landed; running Pass 3 to confirm coverage-delta stabilization (< 2% across passes) + a clean dimension re-scan over the cumulative footprint per the no-early-exit discipline.

## Pass 3 — 2026-07-01 — final

**Iters hardened this pass:** all milestone-touched code (cumulative-scope final pass; a fresh per-function dimension re-scan over the 5 M51-touched Go source files found the seeder Seed-method write-error arms + empty-org/default guards still unswept)
**Tiks covered since prior pass:** all iters in milestone (final-mode cumulative scope)
**Coverage delta on touched files:**
- seeders `seedAIReadinessOrgFunnel`: 92.9% -> 100.0% stmts (the n<=0 empty-org guard)
- seeders `AIReadinessConfigSeeder.Seed`: 87.8% -> 98.0% stmts (the four COPY write-error arms; the residual `months<=0` arm at :127 is defensively unreachable — see below)
- seeders `OrgSettingsSeeder.Seed`: 93.3% -> 100.0% stmts (the organization_settings COPY write-error arm)
- seeders (package): 97.0% -> 97.3% stmts
**Tests added:**
- ai_readiness_harden_test.go: +4 config COPY write-error (cycles/steps/skills/sims fault injection via failCopyConn, each assertWrapped on seeder-context + failing-table) + 1 config zero-Activity end-to-end invariant + 1 funnel empty-org guard (Size 0 and -1 → zero signals)
- org_settings_test.go: +1 organization_settings write-error propagation
**Bugs surfaced + fixed inline:** none — every faulted write already propagated wrapped; the guards were already correct.
**Flakes stabilized:** none
**Cross-iter integration findings:** two defensively-unreachable guards documented (not contrived into coverage): (1) `ai_readiness_config.go:127` `months <= 0` — blueprint.EffectiveStories() re-fills a zero Activity with defaultStoryActivity() (Months=6) BEFORE the seeder runs, so the seed path never presents months<=0; the test proves the observable contract (a windowless org still gets a real closed cycle) instead. (2) `member_languages.go:208` the `add` closure's `seen[code]` early-return — every call site pre-checks. Both are belt-and-suspenders arms behind an upstream invariant. The rest of the M51 footprint's still-uncovered lines are interface stubs (Surface/DependsOn/Isolation — constant returns) + DB-live functions (doSeed/doReset/reloadStackSentinel/printResults/resetCasbin/envMap — need a live Postgres, out of unit scope; their GATES [shouldReloadSentinel, the n=0 --reset guard] are unit-pinned).
**Knowledge backfill:** none — the defensive-guard findings are captured in the test doc comments + this ledger; no subsystem-doc semantic drift.
**Stop condition:** continue-to-next-pass — the coverage delta (97.0% -> 97.3%, < 2%) is stabilizing, but THIS pass's dimension scan surfaced new reachable branches (the write-error arms), so a Pass-4 confirming re-scan is required before "stabilized" per the both-conditions rule.
