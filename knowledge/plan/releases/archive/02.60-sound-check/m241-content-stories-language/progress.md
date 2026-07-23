# M241 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **Go/no-go — read-only prod pool-count query FIRST** (IT sessions per requirement tuple; the interview-scarcity go/no-go) — **GO**: 11/12 tuples bilingual; only INTERVIEW/voice/fail is IT-only; 11/13 pins already italian
- [x] **Language plumbing** — `s.language` in `sourcing.go` SELECT + optional filter; `language` field in the fixture + `content_manifest.go` projection (re-touched `CanonicalFileMatchesProjection`); use `cs.Language` not `sessLanguageEnglish` (rext e6f6804)
- [x] **EN+IT pairs / EN-only fallback per tuple** — 10 counterparts sourced+captured (13→23); the Italian-only INTERVIEW stays solo (toggle hidden); cockpit EN/IT toggle swaps the login-and-land target (rext dfa894d)
- [x] **Sweep (language)** — Go fail-closed `ValidateLanguageConsistency` gate + teeth; `content-language.unit.spec.ts` (4 tests) + a live-sweep language-coverage guard; structure/presence/label only, never the translated value (rext 48bd086)
- [x] **Delivers** — `corpus/ops/demo/content-stories-spec.md` (language/lang_toggle schema + fail-closed gate §4 + EN/IT toggle §7.6) + `session-clone-spec.md` §2.1 (pool query + cs.Language write) (rosetta c3ba981)

---
**M241 COMPLETE (all 5 sections).** rext commits (authoring copy, main): e6f6804 (§2 plumbing), dfa894d (§3 counterparts+toggle), 48bd086 (§4 sweep+gate). rosetta commits (m241 branch): c3ba981 (§5 docs) + plan. Fixture 13→23, denominator 29→49, 11/12 tuples bilingual, INTERVIEW Italian-only. 0 platform-repo edits. PII held (customer content never in context; values-blind; scrubbed fixtures only). Pre-existing finding: 6 red demo-stack/test_cockpit.py tests (academy/overlay, M218/M238/M239 — 0 new from M241).

## M241: Hardening

### Pass 1–3 — 2026-07-22 (harden-milestone, 3 passes, stop: scan clean + flake gate clean)

Work committed to the **rosetta-extensions authoring copy** (`main`): `ae2c876` (pass 1), `0ea0264` (pass 2), `17beede` (pass 3). Corpus branch gets this progress note only. **Close should re-pin/create the consumed rext tag at the hardened HEAD** (M237–M240 precedent).

**Mutation-verification — the fail-closed gates + the plumbing all have teeth** (each broken → RED → restored byte-identical):
- `ValidateLanguageConsistency` (Go gate): a phantom-language toggle (solo cell marked toggle-able / bilingual cell marked solo) → its teeth-test goes RED.
- `content-language.unit.spec.ts` (TS mirror): a phantom toggle injected into the canonical manifest → 2 RED.
- `TestContentManifest_CanonicalFileMatchesProjection` (honesty gate): canonical drift → RED.
- `sourcing.go` language filter → `TestSelectionSQL_LanguageFilter` RED when disabled.
- `LangToggle` projection → `TestBuildContentProducts_LangToggle` RED when forced solo.
- Cockpit per-row `data-lang` emission → the toggle behaviour tests RED.

**Bug / gap closed inline:** the CORE-BUG guard the milestone lacked — **no seeder test asserted the seeded `jobsimulation.sessions.language` column carries `cs.Language`**, so reverting the write to the pre-M241 hard-coded `english` passed EVERY Go suite (proven: mutated → all green). `TestContentStorySeeder_WritesRealLanguage` closes it (col resolved by NAME; asserts language == the pin's real language, valid non-blank label, and the set spans BOTH languages so it can't pass vacuously). Mutation-verified RED against the revert, GREEN restored. This is the exact defect v2.6 "sound check" kills.

**Tests added (all rext; STRUCTURE/PRESENCE/label only — never a translated value, P2 copy-immunity):**
- Go (`stack-seeding/seeders/content_stories_test.go`): +2 (`WritesRealLanguage` regression, `BlankLanguageFallsBackToEnglish` fallback-branch) + `sessionColIndex` helper.
- Python (`demo-stack/tests/test_cockpit.py`): +9 across `TestContentLanguageToggle` (CTA re-targets by language, solo-unfiltered, default/visible coherence), `TestContentToggleLangs` (excludes solo/non-sim, all-solo → no toggle, ordering), `TestLangJsStructure` (`_LANG_JS` balance + selector wiring).

**Coverage (touched language fns):** `SelectionSQL` 100%, `bilingualTuples` 100%, `resolveSession` 100%, `ValidateLanguageConsistency` 96%, `seedContentStorySession` 91.3% (language write now covered).

**Priority-4 re-verified:** the leak/scrub gate (`TestFixtures_NoStructuralPII` pins 23, sweeps the 10 new IT captures), the placeholder gate, and the believable-score band (`TestPassedFixtureScoresAreBelievable`) all green.

**Full suites:** Go green; Python 142 tests / SAME 6 pre-existing academy+overlay failures (M218/M238/M239, 0 new from M241, owned by M244); TS 151 unit specs green. Flake gate: 3 consecutive clean sequential runs of the new tests, all stacks.

**Knowledge backfill:** none warranted — the write-side language invariant is already documented (`session-clone-spec.md` §2.1, build §5 `c3ba981`); the added tests ARE the record of the coverage discovery.

## M241: Final Review

Close review 2026-07-22 (`close-milestone`). A clean complete section close — the harden pass had already
mutation-verified every gate + closed the write-side gap, so the review found near-nothing.

### Scope
- [x] All 5 sections checked off; overview `In:` list fully delivered; no dropped scope. KB-1 landed (§2.1). 0 new deferrals.

### Code Quality
- [x] `go vet ./...` (stack-seeding) clean · `tsc --noEmit` (e2e) clean · Python is stdlib unittest (no linter). No dead code, consistent patterns, XSS-safe (`html.escape(quote=True)`), injection-free JS (no manifest free-text in raw strings), language literal validated before SQL embedding.

### Adversarial (Phase 2c → decisions.md)
- [x] AR-1: two-point `LangToggle` derivation under a projection drop → verified fail-closed (gate runs post-drop; solo+toggle=true is a passing teeth branch).
- [x] AR-2: cockpit hard-coded English default vs adaptive toggle default → verified unreachable by the closed `{english,italian}` set (bilingual ⟹ both); latent-coupling note recorded for a future 3rd language.

### Documentation
- [x] [fix] `content-stories-spec.md:237` stale "29 landable pairs" → the pinned count (49 since M241; 29 at M236) — the route-contract spec reads `content-denominator.json`=49.
- [x] `content-stories-spec.md` (§2 language axis, §4 fail-closed both sides, §7.6 toggle, §7.6→§7.7 clean renumber) + `session-clone-spec.md` §2.1 (pool query + `cs.Language` write) accurate; denominator 49 consistent everywhere; no broken cross-refs.

### Tests & Benchmarks
- [x] Go green · TS 151/151 unit specs · Python 136 pass / 6 pre-existing academy+overlay failures (M218/M238/M239, 0 new from M241, Fate-2 → M244). Gate coverage: `TestValidateLanguageConsistency` teeth, `TestBilingualTuples`, `TestBuildContentProducts_LangToggle`, `TestContentStorySeeder_WritesRealLanguage` (the core-bug regression).

### Decision Triage
- [x] D-M241-1 (pool-query toggle model) + the EN-only-fallback user decision → already blended into `content-stories-spec.md` §2 + `session-clone-spec.md` §2.1 during build; verified accurate. AR-1/AR-2 + the 6 pre-existing failures stay archived in `decisions.md` (maintainer-only / already-fated → M244).

## M241: Completeness Ledger (Phase 9, section variant)

Every `overview.md` `In:` scope item placed into exactly one three-fate category.

### Done (Fate 1 — landed in M241)
- [x] **Read-only prod pool-count query FIRST** (the interview-scarcity go/no-go) → §1, verdict **GO** (11/12 tuples bilingual; INTERVIEW Italian-only = R2 realized).
- [x] **`s.language` in `sourcing.go` SELECT + optional filter; `language` field in the fixture + `content_manifest.go` projection (re-touched `CanonicalFileMatchesProjection`); `cs.Language` not `sessLanguageEnglish`** → §2 (rext e6f6804).
- [x] **EN+IT pairs / EN-only fallback per tuple; cockpit EN/IT toggle swaps login-and-land** → §3 (rext dfa894d; fixture 13→23, denominator 29→49).
- [x] **Extend the content-stories sweep for language (structure/presence/label only, never a translated value)** → §4 (rext 48bd086; Go `ValidateLanguageConsistency` + teeth, TS `content-language.unit.spec.ts`).
- [x] **Delivers: `content-stories-spec.md` (language field + EN/IT toggle + re-touched honesty gate)** → §5 (rosetta c3ba981; + `session-clone-spec.md` §2.1).

### Confirmed-covered (Fate 2 — already owned by another milestone of this release)
- The **6 pre-existing `test_cockpit.py` failures** (academy + overlay-JS, M218/M238/M239) — NOT an M241 scope item; surfaced during the close. Subset of the standing demo-stack test debt, already **Fate-2 → M244** (named expiry, re-confirmed with fresh dates at M238/M239/M240 closes). No edit.

### Annotated (Fate 3), Dropped, Release-scope-breaking deferral (escape hatch)
- **None.** M241 introduced zero new deferrals.

**Verdict: all 5 scope items delivered in this milestone (Fate 1). Nothing routed, dropped, or escape-hatch-deferred. Zero escape-hatch entries → no sign-off required; merge proceeds.**
