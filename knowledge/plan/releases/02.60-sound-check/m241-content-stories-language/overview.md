---
milestone: M241
slug: content-stories-language
version: v2.6 "sound check"
milestone_shape: section
status: planned
created: 2026-07-20
last_updated: 2026-07-20
depends_on: M240
delivers: corpus/ops/demo/content-stories-spec.md
---

# M241 — content-stories language

**Status:** `planned`  ·  **Shape:** `section` (opens with a pool-count go/no-go)  ·  **Complexity:** medium  ·  **Depends on:** M240

## Goal
Each session is consumed in its intended language, with an EN/IT cockpit toggle.

## User decision baked in (2026-07-20) — EN-only fallback per tuple
**Opens with a read-only prod pool-count query FIRST** (IT sessions per requirement tuple — the interview-scarcity go/no-go). Toggle where IT exists, **EN-only fallback per tuple** where absent (toggle hidden/disabled there). No blocking. This is **R2 (blocks-scope) — language scarcity**: IT interview sessions may not exist; the pool query decides per-tuple coverage.

## Scope
### In
- **Read-only prod pool-count query FIRST** (IT sessions per requirement tuple — the interview-scarcity go/no-go).
- Add `s.language` to `sourcing.go` SELECT + optional filter; add a `language` field to the fixture + `content_manifest.go` projection (re-touch the `CanonicalFileMatchesProjection` honesty gate); use `cs.Language` instead of the hard-coded `sessLanguageEnglish`.
- Source EN+IT pairs per tuple where IT exists; **EN-only fallback per tuple** where absent (toggle hidden/disabled there); cockpit toggle swaps the login-and-land target.
- Extend the **content-stories sweep** for language (assert structure/presence, **never** the translated value — P2 forbids copy assertions).

### Out
- The row-layout / tab-selector cockpit-UX (M242).

## Open questions
- IT interview sessions may not exist (R2) — the pool query decides per-tuple coverage; EN-only fallback where absent.
- Does the honesty gate need a per-tuple language-coverage field to stay fail-closed?

## Delivers
`corpus/ops/demo/content-stories-spec.md` (the language field + EN/IT toggle + the re-touched honesty gate).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
