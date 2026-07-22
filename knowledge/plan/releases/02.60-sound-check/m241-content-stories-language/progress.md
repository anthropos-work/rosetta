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
