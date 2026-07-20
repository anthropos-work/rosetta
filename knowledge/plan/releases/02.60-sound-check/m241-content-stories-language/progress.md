# M241 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [ ] **Go/no-go — read-only prod pool-count query FIRST** (IT sessions per requirement tuple; the interview-scarcity go/no-go)
- [ ] **Language plumbing** — `s.language` in `sourcing.go` SELECT + optional filter; `language` field in the fixture + `content_manifest.go` projection (re-touch `CanonicalFileMatchesProjection`); use `cs.Language` not `sessLanguageEnglish`
- [ ] **EN+IT pairs / EN-only fallback per tuple** — source pairs where IT exists; EN-only fallback where absent (toggle hidden/disabled there); cockpit toggle swaps the login-and-land target
- [ ] **Sweep (language)** — extend the content-stories sweep; assert structure/presence, never the translated value
- [ ] **Delivers** — `corpus/ops/demo/content-stories-spec.md` (language field + EN/IT toggle + honesty gate)
