# M241 — Spec notes

Topic → doc → code triples + language-pool / toggle findings accumulate here during build.

## Pre-flight audits — Go/no-go + Language plumbing (first section)
- **Phase 0b KB-fidelity: YELLOW → proceed.** Report: `kb-fidelity-audit.md`. No blind areas (language axis = declared Delivers), no stale load-bearing claims; one tracked completeness gap (KB-1, session-clone-spec.md write-side language note → Phase 5). sha at audit: `4cba368` (release base).
- **Prod-read gates CLEAR (verified live, read-only):** `~/.pgpass` `marco_read` connects; `jobsimulation.sessions.language` exists = `character varying` (varchar, NOT enum); distinct values include `english` (7,682) + `italian` (11,971). EN=`"english"` (= `sessLanguageEnglish`), IT=`"italian"`.

## Topic → doc → code triples
- content_products manifest schema → `content-stories-spec.md` §2 → `seeders/content_manifest.go` (`ContentProduct`/`ContentProductSession`)
- honesty gate → `content-stories-spec.md` §4 → `cmd/stackseed/main_test.go:1210` (`TestContentManifest_CanonicalFileMatchesProjection` + `HasTeeth`)
- sourcing SQL → `session-clone-spec.md` §2 → `contentsession/sourcing.go` (`SelectionSpec`/`SelectionSQL`)
- session write / language → `session-clone-spec.md` §3 → `seeders/content_stories_write.go` (`sessLanguageEnglish` @ :71) + const `seeders/jobsim_sessions.go:66`
- fixture model → `session-clone-spec.md` → `contentsession/contentsession.go` (`ContentSession`) + `contentsession/fixture/content-sessions.yaml` (13 pins)
- sweep + denominator → `content-stories-spec.md` §4 → `stack-verify/e2e/run-content-stories.sh` + `aggregate-content.py` + `content-denominator.json` (29 pairs) + `tests/content-route-contract.unit.spec.ts`

## Go/no-go — read-only prod pool-count query FIRST (R2) — RESOLVED: **GO**
Pool query run read-only (`marco_read`), counts+labels only (no content). Predicates mirror `sourcing.go`
(public-anchored + type-matched + completed + non-test). **Surprise: 11 of the 13 current pins are already
`italian`** (only `asmt-doc-fail` + `train-chat-fail` are `english`) — yet the seeder hardcodes
`language='english'` for all. That IS the core defect §2 fixes.

**Per-tuple language coverage (EN pool | IT pool → toggle viable?):**

| requirement tuple | current pin(s) → real lang | EN | IT | pair? |
|---|---|---|---|---|
| ASSESSMENT/voice/pass | asmt-voice-pass, asmt-voice-pass-2 → IT | 86 | 447 | YES |
| ASSESSMENT/voice/fail | asmt-voice-fail → IT | 306 | 723 | YES |
| ASSESSMENT/code/pass | asmt-code-pass → IT | 3 | 68 | YES (EN scarce) |
| ASSESSMENT/code/fail | asmt-code-fail → IT | 31 | 64 | YES |
| ASSESSMENT/document/pass | asmt-doc-pass → IT | 1 | 14 | YES (EN scarce) |
| ASSESSMENT/document/fail | asmt-doc-fail → **EN** | 8 | 19 | YES |
| TRAINING/document/pass | train-doc-pass → IT | 2 | 4 | YES (both scarce) |
| TRAINING/chat/fail | train-chat-fail → **EN** | 38 | 55 | YES |
| HIRING/voice/pass | hire-voice-pass → IT | 18 | 29 | YES |
| HIRING/voice/fail | hire-voice-fail → IT | 129 | 116 | YES |
| INTERVIEW/voice/pass | intv-voice-pass → IT | 1 | 34 | YES (EN scarce) |
| **INTERVIEW/voice/fail** | intv-voice-fail → IT | **0** | 5 | **NO — IT-only (single-language)** |

**Go/no-go verdict: GO.** 11/12 tuples have both languages → toggle viable. The one single-language tuple is
**INTERVIEW/voice/fail** (EN=0) — exactly R2's interview scarcity, inverted (IT present, EN absent). The
user-decision fallback (toggle hidden/disabled for a single-language tuple) applies there.

Global distinct `sessions.language` values: italian 11,971 · english 7,682 · spanish 159 · dutch 57 · french
27 · japanese 23 · german 17. IT stored as `"italian"`, EN as `"english"` (= `sessLanguageEnglish`).

## Language plumbing (§2)
- Add `s.language` to `sourcing.go` SELECT + optional filter (`SelectionSpec.Language`).
- Add a `Language` field to `ContentSession` (fixture) + `ContentProductSession` (`content_manifest.go` projection); re-touch the `CanonicalFileMatchesProjection` honesty gate (regenerate `presets/content-manifest.json`).
- `content_stories_write.go:71`: `sessLanguageEnglish` → `cs.Language` (empty → english fallback).
- Set `language:` on all 13 existing pins to their REAL prod language (query A): 11 italian, 2 english (`asmt-doc-fail`, `train-chat-fail`).
- Tests to update: `contentsession_test.go` (schema, +language coverage), `sourcing_test.go` (SELECT string / filter), `content_manifest_test.go` (per-session language field), `content-route-contract.unit.spec.ts` (unchanged pair count in §2).

## Counterpart plan (§3) — 10 viable counterparts (13 → 23 sessions)
Rule D-M241-1: pair each tuple where a believable counterpart is sourceable (in-band 70-95 for passed).
EN counterparts (IT-pinned tuples → add EN): asmt-voice-pass, asmt-voice-fail, asmt-code-pass (1 in-band),
asmt-code-fail, asmt-doc-pass (1 in-band), train-doc-pass (1 in-band), hire-voice-pass, hire-voice-fail.
IT counterparts (EN-pinned tuples → add IT): asmt-doc-fail, train-chat-fail.
**Skipped → single-language fallback (toggle hidden):** INTERVIEW/voice/pass (EN=0 in-band) + INTERVIEW/voice/fail
(EN=0) → the whole INTERVIEW product is Italian-only in the demo (R2 interview scarcity, fully realized).
Denominator moves 29 → 43 (23 sim × player… manager varies) — recompute exactly at §3 from the projection.
Cardinality-gate assertions are `>=` floors (safe); `len==13` exact assertions bump to 23 across
`contentsession_test.go`, `cleanliness_test.go` (×2), `content_manifest_test.go`.

## Cockpit toggle
- Source EN+IT pairs per tuple where IT exists; EN-only fallback per tuple where absent (toggle hidden/disabled there); the toggle swaps the login-and-land target.

## Sweep (language)
- Extend the content-stories sweep for language: assert structure/presence, NEVER the translated value (P2 forbids copy assertions).

_(will accumulate topic → doc → code triples during build)_
