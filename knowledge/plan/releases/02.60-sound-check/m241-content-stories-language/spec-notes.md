# M241 — Spec notes

Topic → doc → code triples + language-pool / toggle findings accumulate here during build.

## Go/no-go — read-only prod pool-count query FIRST (R2)
- Count IT sessions per requirement tuple (the interview-scarcity go/no-go). Toggle where IT exists, EN-only where absent.

## Language plumbing
- Add `s.language` to `sourcing.go` SELECT + optional filter.
- Add a `language` field to the fixture + `content_manifest.go` projection; re-touch the `CanonicalFileMatchesProjection` honesty gate.
- Use `cs.Language` instead of the hard-coded `sessLanguageEnglish`.

## Cockpit toggle
- Source EN+IT pairs per tuple where IT exists; EN-only fallback per tuple where absent (toggle hidden/disabled there); the toggle swaps the login-and-land target.

## Sweep (language)
- Extend the content-stories sweep for language: assert structure/presence, NEVER the translated value (P2 forbids copy assertions).

_(will accumulate topic → doc → code triples during build)_
