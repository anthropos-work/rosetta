---
milestone: M238
slug: ant-academy-reliability
version: v2.6 "sound check"
milestone_shape: section
status: planned
created: 2026-07-20
last_updated: 2026-07-20
depends_on: M237
delivers: corpus/services/ant-academy.md + corpus/ops/demo/frontend-tier.md
---

# M238 — ant-academy reliability

**Status:** `planned`  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** M237

## Goal
A hero can follow a course and actually consume a chapter; the language switch works.

## Scope
### In
- Fix **#3 (Start→404)**: the demo academy chapter-body path is unwired — bodies are backend-authoritative, no FS fallback; the catalog demopatch covers only the catalog. Wire a chapter-body demo path — a chapter-body FS-fallback demopatch analogous to `academy-fs-published-fallback`, OR wire the academy backend for the demo.
- Fix **#2 (language error** — re-triaged in M237; likely the same backend-null path).
- Extend the **academy presence/coverage sweep**.

### Out
- The enterprise-surface / talk-to-data fixes (M239).
- Content-stories (M240+).

## Open questions
- Chapter-body FS-fallback demopatch vs wiring the academy backend for the demo — which is revert-clean + sufficient?
- Is #2 (language) the same backend-null path as #3, or distinct? (M237's re-triage informs this.)

## Delivers
`corpus/services/ant-academy.md` + `corpus/ops/demo/frontend-tier.md` (the chapter-body demo path + the extended academy sweep).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
