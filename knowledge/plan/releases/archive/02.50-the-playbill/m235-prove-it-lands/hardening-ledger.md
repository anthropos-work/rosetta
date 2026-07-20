# Hardening Ledger — M235 Prove it lands

> Milestone M235 is **pragmatic-closing** (the user's "build non-sim seeders, then close" mandate; the
> live `(session × action)`-lands gate is routed to **M236**, NOT met here by design). This is the FIRST +
> FINAL harden pass. Scope is cumulative-mode over the whole M235-touched footprint (all in the rext
> `stack-seeding` module + the rosetta corpus). The live cold-reset-to-seed proof is M236 — this pass
> deepens the UNIT-provable surface only. rext test commits land in the rext authoring clone
> `.agentspace/rosetta-extensions/` and are re-tagged; this ledger is the rosetta-side record.

## Pass 1 — 2026-07-20 — final

**Iters hardened this pass:** all milestone-touched code (iter-03…iter-08 footprint) — cumulative --final scope.
**Tiks covered since prior pass:** all iters in milestone (first harden pass; no prior ledger entry).
**Coverage delta on touched files (statements):**
- `scrub/scrub.go`: 94.6% → 100.0% (package scrub 94.6% → 100.0%)
- `cmd/content-capture/main.go`: 0% → 28.7% (pure helpers only; `captureSession`/`main` are DB-bound — the M236 live surface)
  - `leakCheck` 0→100%, `emailLocalName` 0→100%, `jsonLeafStrings` 0→100%, `idxOr` 0→100%, `scrubJSONRaw` 0→90%, `writeJSON` 0→80%
- `seeders/content_nonsim.go`: `resolveNonSimSession` 81.2→100%, `Surface`/`DependsOn`/`Isolation` 0→100%, `Seed` 88.6→91.4% (package seeders 95.9→96.1%)
- `contentsession` package: 93.7 → 94.7% (new standing cleanliness gate)
**Tests added:**
- `cmd/content-capture/main_test.go` (NEW): 8 test funcs — `TestLeakCheck_FlagsEverySurface` (13 sub-cases: every free-text surface — actor role_key, attempt, skill, criterion title/input_data jsonb, check feedback, interaction payload jsonb, code_submission, collaborative_asset, interview report user/manager), `TestLeakCheck_CleanSessionPasses` (placeholders + mid-word substrings NOT flagged), `TestLeakCheck_EmptyInputs`, `TestEmailLocalName`, `TestJsonLeafStrings`, `TestScrubJSONRaw` (fail-safe → null on unparseable), `TestIdxOr`, `TestWriteJSON` (deterministic + newline-terminated + round-trips)
- `contentsession/cleanliness_test.go` (NEW): 3 test funcs — `TestFixtures_NoStructuralPII` (0 surviving email/url/phone in ANY committed fixture leaf), `TestFixtures_PlaceholdersPresentAndWellFormed` (`<<ACTOR_0>>` present in every fixture — the M235 owner-scrub tripwire — + every `<< >>` token well-formed), `TestFixtures_CleanlinessGateHasTeeth` (proves the standing gate is not toothless: a raw email/url/phone in a leaf IS detected, a malformed placeholder IS rejected)
- `scrub/scrub_test.go` (+4): `TestScrub_SkipsBlankKeys`, `TestSurvivingToken_EdgeInputs` (empty text / empty tokens / sub-3-char skip), `TestSortedKeysByLenDesc_EqualLengthDeterministic` (tie-break arm), `TestScrub_NonASCIIBoundaryFallback` (accented-name / trailing-`.` boundary omission)
- `seeders/content_nonsim_test.go` (+4 this pass): `TestResolveNonSimSession_FailClosedDrops` (skill-path missing id / academy missing slug / no non-sim builder all drop, never fabricate), `TestBuildNonSimProducts_NoEligibleOwnerDropsFailClosed`, `TestContentStoryNonSimSeeder_Contract` (Surface/DependsOn/Isolation), `TestContentStoryNonSimSeeder_CopyErrorPropagates` (wrapped with seeder + failing table)
**Bugs surfaced + fixed inline:** none — every M235 fail-closed / leak-gate / owner-assignment path held under test.
**Flakes stabilized:** none.
**Knowledge backfill:** none required — the scrub posture (best-effort, not provably-clean; residual risk accepted, VPN/tailnet scope the control) and the fail-closed no-fabrication contracts are already documented in `session-clone-spec.md` / `content-stories-routes.md` / `safety.md §3.8`. This pass added STANDING test coverage over already-documented semantics; no behavior changed.
**Stop condition:** continue-to-next-pass — statement coverage is high but the delta-across-passes measurement requires a second pass to compute; final mode's defining cross-iter integration check not yet run.

## Pass 2 — 2026-07-20 — final

**Iters hardened this pass:** all milestone-touched code (cumulative --final; cross-iter integration focus).
**Tiks covered since prior pass:** n/a (same session, second pass).
**Coverage delta on touched files (statements):** held flat — scrub 100.0%, contentsession 94.7%, seeders 96.1%, cmd/content-capture 28.7% (delta < 2% vs Pass 1). Cross-iter integration tests pin cross-arm INVARIANTS, not new statements — expected for a final-mode integration pass.
**Tests added:**
- `seeders/content_nonsim_test.go` (+1): `TestContentStoryNonSim_OwnerConsistencyAcrossAllArms` — the CROSS-ITER integration proof tying iter-05 (skill-path) + iter-06 (ai-labs) + iter-07 (academy): one shared flat index + owner assignment drives BOTH the projection and the seeder; per exhibit the SEEDED row's user_id (`skill_path_sessions` for skill-path, `lab_sessions` for ai-labs) equals the flat-index owner AND the projection's `player_seat` names that same owner. Closes the gap that only the skill-path arm's owner-match was pinned (the ai-labs `lab_sessions` owner + the academy slot-advance were unverified end-to-end).
**Cross-iter integration findings:**
- The two honesty gates (`TestManifest_CanonicalFileMatchesProjection` for `seed-generation-manifest.yaml`; `TestContentManifest_CanonicalFileMatchesProjection` for `content-manifest.json`) verified GREEN uncached — the canonical `content-manifest.json` carries all 4 products (simulation, skill-path-legacy, ai-labs, skill-path-new) / 18 sessions, so the cross-iter 4-product manifest integration is already gated.
- Remaining uncovered lines on M235 files are all genuinely UNREACHABLE defensive arms: `buildNonSimProducts` `len(exhibits)==0` (hardcoded non-empty registry), `Seed` `!wrote` n=0 audit guard (unreachable with the current registry — skill-path always writes rows when slots exist), and the two `json.Marshal` error arms (`scrubJSONRaw`, `writeJSON`) on known-good structs. None is injectable without contorting production code; chasing them would be shallow coverage-gaming, not meaningful tests.
**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none — 3 consecutive clean `-count=1` runs of all four touched packages; `go vet` clean.
**Knowledge backfill:** none required (see Pass 1).
**Stop condition:** stabilized — coverage delta < 2% across Pass 1→Pass 2 AND the Phase 2 dimension scan (test-depth / edge / error-path / cross-iter integration) found nothing new reachable; the only residual uncovered lines are unreachable defensive arms.

---

**Verification (Phase 5):** all four touched Go packages + `cmd/stackseed` (honesty gates) GREEN; flake gate 3/3 clean; `go vet` clean. The Python cockpit suite (`demo-stack/tests/test_cockpit.py`) is at its documented pre-existing state — 128 tests, 6 stale failures (release-close carry, M218 overlay / M53 academy-link assertions), unchanged by this pass and out of scope per the milestone mandate. The live `(session × action)`-lands proof on a cold reset-to-seed remains **M236**.
