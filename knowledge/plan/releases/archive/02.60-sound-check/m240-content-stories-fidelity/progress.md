# M240 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **⚠️ HARD media-safety gate (R1) — MUST clear first** — fresh data-controller sign-off + `safety.md` §3.8 raw-media amendment + a voice/document anonymization decision, BEFORE any customer audio lands in a demo — LANDED (§3.8.1; voice unscrubbable→VPN-scope-only, docs→scrub, sign-off 2026-07-21, gender-coherence; commit 7020c3f)
- [x] **Defect 1 (selection)** — constrain the public sim's type to the cell type in `sourcing.go`; re-pin `content-sessions.yaml` — LANDED: `d.type = <cell sim_type>` predicate (robust, not slug); re-pinned asmt-voice-pass to a real assessment-voice session (address-underperformance…, score 70); `--only` surgical re-capture; canonical presets regenerated (rext 9e8305a, corpus 7226a3c)
- [x] **Defect 3 (document)** — write the dropped `input_data` at seed time + port the document blob if `storage_upload` — LANDED (seed-time write via content-specific `contentCriterionResultCols`). **Blob investigation: the body is `input_data.text_document` (collaborative_doc), NOT a storage_upload/S3 blob → NO blob to port, fully landable.** (rext cb64ccd, corpus 92ae3ed)
- [x] **Defect 2 (voice) → DELIVERED = VOICE PRESENCE-ONLY** (user decision 2026-07-22). The faithful `chime_status='not_available'` state **IS the v2.6 deliverable** — an honest "recording not available" over a broken 500 player. NOT a blocker, NOT an incomplete section. LANDED: the fresh 2026-07-21 data-controller VIDEO sign-off + §3.8.1 amendment + the media-substrate spec + the gender-coherence contract + the seed-side reference-port design (corpus 37a47a0) — so the real exhibit is fully documented + **pre-blessed**. ROUTED (not blocked): the real-video exhibit (provision the Bunny recording signing keys + re-pin a hiring-voice cell to a recorded session + write `chime_status='completed'`+`bunny_video_id` + verify playback) → **M244 (Fate-3, user pre-approved 2026-07-22)** — `BUNNY_RECORDING_CDN_TOKEN_KEY` + `BUNNY_RECORDING_PULL_ZONE_HOST` are genuinely ABSENT from this box's dev-stack (verified values-blind) but may be reachable on billion. `DEF-M240-01`; added to M244 In-list at close. Full diagnosis + render path in decisions.md + `media-substrate-spec.md`.
- [x] **Pass-rate (#4)** — LANDED FULLY. CODE (rext 0753e48: `ScoreMin/ScoreMax` band + `score ASC` tiebreak + test). FIXTURE (rext 3e9696c): the 5 still-100% passed cells re-pinned to real 70-95% sessions (asmt-voice-pass-2→74, asmt-code-pass→72, train-doc-pass→82, hire-voice-pass→83, intv-voice-pass→81), re-captured via `content-capture --only` (read-only prod, values-blind, fail-closed leak gate passed, every modality matched its cell), captured attempt scores == declared skeleton scores, both canonical presets regenerated, `CanonicalFileMatchesProjection` honesty gate green.
- [x] **Delivers** — LANDED (corpus 37a47a0): new `corpus/ops/demo/media-substrate-spec.md` (the Bunny-CDN reference substrate + render path + seed-side exhibit mechanism + values-blind Bunny-key provisioning + hiring-voice-only sourcing constraint + the honest Bunny-key-blocked status) + `safety.md` §3.8.1 amended for real candidate interview VIDEO (fresh 2026-07-21 sign-off, reference-not-bytes model, gender-coherence, document=inline-text) + PM summary reconciled + indexed in `demo/README.md`.

## M240: Hardening

### Pass 1 — 2026-07-22

**Scope manifest (milestone-touched code, rext `stack-seeding`):**
- `contentsession/sourcing.go` — `SelectionSQL` (Defect 1 `d.type` CTE predicate + pass-rate `ScoreMin/ScoreMax` band + `score ASC` tiebreak). Tests: `sourcing_test.go`.
- `seeders/content_stories_write.go` — `contentCriterionResultCols()` + the seed-time `input_data` value write (Defect 3). Tests: `content_stories_test.go`.
- `seeders/content_stories.go` — flush now uses `contentCriterionResultCols()`. Tests: `content_stories_test.go`.
- `cmd/content-capture/main.go` — the `--only` surgical re-capture switch. Tests: `main_test.go`.
- (fixtures/presets: content JSONs + `content-sessions.yaml` + the 2 canonical presets — data, gated by the honesty + scrub tests, unchanged by harden.)

**Coverage delta (milestone-touched files):**
- `contentsession`: 95.8% → 95.8% (SelectionSQL already 100%; new tests are behavioral/fixture gates, not line-fillers).
- `seeders`: 96.1% → 96.1% (touched functions already saturated; new tests are behavioral).
- `cmd/content-capture`: 33.6% → 36.4% (+2.8 — the extracted `unknownOnlyKeys`/`pinnedKeys` guard helpers).
- Coverage-as-finder: the two saturated packages had no meaningful line gaps; the harden value is mutation-verified behavioral depth, not a coverage number.

**Mutation-verification (the release thesis — each fix's guard goes RED against the ORIGINAL bug, then restored byte-clean):**
- Defect 1: removing `AND d.type = '<cell type>'` from the pub CTE → `TestSelectionSQL_TypeMatchExcludesInterview` + `TestSelectionSQL_Valid` + (new) `TestSelectionSQL_TypeMatchScopedToPublicCTE` RED.
- Pass-rate: dropping the score band → `TestSelectionSQL_ScoreBand` RED (missing BETWEEN); flipping `score ASC`→`DESC` → RED (3 assertions).
- Defect 3: dropping `input_data` from the col set → `TestContentStorySeeder_WritesInputData` RED (col-set END check); dropping the value → RED (14 vs 15 values); writing a `null` body → `WritesInputData` + (new) `TestContentStorySeeder_DocumentBodyReachesManager` RED (6 document bodies → 0).
- All 5 mutations restored via Edit (never `git checkout`/`reset`); production SQL/seeder logic byte-unchanged after; suite green.

**Tests added (6, rext commit `ae0e869`):**
- `contentsession/sourcing_test.go`: `TestSelectionSQL_TypeMatchScopedToPublicCTE` (CTE-scope depth), `TestPassedFixtureScoresAreBelievable` (believable-band fixture gate — every passed score ∈ [70,95], none 100).
- `seeders/content_stories_test.go`: `TestCriterionColSets_ContentExtendsSharedNoRegression` (shared set stays input_data-free + extend-not-reorder + append-aliasing probe), `TestContentStorySeeder_DocumentBodyReachesManager` (document-family body count-matched into the write set).
- `cmd/content-capture/main_test.go`: `TestUnknownOnlyKeys` + `TestPinnedKeys` (the `--only` fail-loud-on-typo guard, made unit-provable by extracting `unknownOnlyKeys`/`pinnedKeys`).

**Bugs fixed inline:** none — all three fixes were already correct; harden confirmed their guards have teeth and closed the named coverage gaps. The one latent risk found (`append(criterionResultCols(), "input_data")` would alias if the shared set ever gained spare capacity) is NOT a live bug today (fresh literal, cap==len) but is now pinned by the aliasing probe.

**Flakes stabilized:** none (0 flakes across 3 consecutive sequential runs of the new tests).

**Honesty + scrub gates:** `TestManifest_CanonicalFileMatchesProjection` + `TestContentManifest_CanonicalFileMatchesProjection` (cmd/stackseed) and the scrub leak gates (`TestEmbeddedContent_NoStructuralPII`, `leakCheck` surfaces) regression-run GREEN — the re-captured fixtures + regenerated presets stay consistent.

**Delivers doc consistency (thesis #4):** `media-substrate-spec.md` + `safety.md` §3.8.1 verified internally consistent — exhibit-by-reference throughout (`bunny_video_id` + read-only signing key, "no bytes copied", render-time signed-URL streaming), the fresh 2026-07-21 video sign-off recorded in both, no byte-port claim anywhere. No fix needed.

**Knowledge backfill:** no KB-worthy findings this pass. The harden added tests (no behavior change); the Defect-1 `d.type` predicate is already documented in `session-clone-spec.md` §2 and the `input_data` seed-time write in §4 (both landed at build). The aliasing-probe invariant is a code-internal guard, pinned by its test, not a subsystem fact.

### Stop condition
Stopped after Pass 1: the Step 2b scan closed every gap the thesis named (mutation-verified 3 fixes + CTE-scope + believable-band + col-set non-regression/aliasing + document-body + `--only` guard); coverage deltas negligible on the already-saturated packages (the only remaining uncovered lines are the live-DSN capture path + `os.Exit` main, not unit-testable — documented in `main_test.go`); 0 flakes. A second pass would be line-chasing on saturated code. `/developer-kit:close-milestone` Phase 4 runs independently as defense-in-depth.

## M240: Final Review

Review found **0 fix-required findings** — the build + harden pass had already mutation-verified the 3 fixes and
closed the named gaps. The close work is scope-resolution (voice presence-only) + routing + the record.

### Scope
- [x] Voice section resolved — DELIVERED = voice presence-only (user decision 2026-07-22); real-video exhibit → M244 (DEF-M240-01, Fate-3, user pre-approved). M244 In-list edited.
- [x] All 6 sections resolved (5 fixed + voice presence-only) — complete section close. No silent drops.

### Code Quality
- [x] rext `go vet` + `go build ./...` clean; 0 TODO/FIXME/HACK in touched prod files; no cross-module reach-in (content-specific `contentCriterionResultCols` deliberately separate from the shared PersonaSeeder set).

### Documentation
- [x] `media-substrate-spec.md` honest + complete (Bunny-key-blocked status, atomic-hold, reference-not-bytes); indexed in `demo/README.md`.
- [x] `safety.md` §3.8.1 (video-by-reference, gender-coherence, fresh 2026-07-21 sign-off, honest current status) internally consistent; `session-clone-spec.md` §2 (3rd predicate) + §4 (input_data seed-time write) + the `--only` note match the code.

### Tests & Benchmarks
- [x] Full `stack-seeding` module green (16 pkgs, 0 fail); harden added 6 deepening tests (mutation-verified); honesty + scrub leak gates GREEN. Flake gate: Phase 8.

### Decision Triage
- [x] The load-bearing knowledge (3rd predicate, input_data-at-seed, `--only`, media substrate, video-safety posture) already flowed to the corpus docs at build. Adversarial-review scenarios recorded in decisions.md. Remaining decisions = maintainer-only archive.

### Deferrals
- [x] `/developer-kit:audit-deferrals --scope=milestone` → **YELLOW** (report: `audit-deferrals/deferral-audit-2026-07-22-m240-close.md`). DEF-M240-01 → M244 (Fate-3, user pre-approved); DEF-M239-01 / reap-13 already in M244 In-list; standing-8 = tracked Fate-2 repeat → M244 (named expiry). 0 RED blockers.

## M240: Completeness Ledger (section variant)

**Done (Fate 1) — landed in M240:**
- ✅ HARD media-safety gate (safety.md §3.8.1 raw-media/VIDEO amendment + fresh 2026-07-21 data-controller sign-off + gender-coherence)
- ✅ Defect 1 (selection) — the `d.type = <cell sim_type>` CTE predicate; interview sim no longer leaks into voice cells
- ✅ Defect 3 (document) — input_data written at seed time (body is inline text, no S3 blob)
- ✅ Pass-rate (#4) — score-band + score-ASC tiebreak + 5 cells re-pinned to 70–95% sessions
- ✅ Delivers — new media-substrate-spec.md + safety §3.8.1 + session-clone-spec §2/§4 bumps; indexed
- ✅ Voice → DELIVERED = voice presence-only (user decision 2026-07-22) — the faithful `chime_status='not_available'` state IS the v2.6 deliverable

**Annotated (Fate 3) — attached to a release-milestone at close:**
- ➡️ DEF-M240-01 — the real-video content-stories exhibit → M244 (user pre-approved 2026-07-22; M244 In-list edited)

**Confirmed-covered (Fate 2) — inherited, already owned by M244 (no M240 edit):**
- standing-8/9 demo-stack test debt (M238-D5) · DEF-M239-01 (D12) · reap-17700 collision (D13)

**Dropped:** none.
**Release-scope-breaking deferral (escape hatch):** none → no sign-off required; clean complete section close.
