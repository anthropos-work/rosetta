---
milestone: M237
slug: clean-stage
version: v2.6 "sound check"
milestone_shape: section
status: archived
created: 2026-07-20
last_updated: 2026-07-21
depends_on: none
delivers: corpus/ops/rosetta_demo.md + corpus/ops/demo/demopatch-spec.md
---

# M237 — clean stage

**Status:** `archived` (completed 2026-07-21)  ·  **Shape:** `section` (HARD go/no-go barrier)  ·  **Complexity:** medium  ·  **Depends on:** none (opens the release)

## Goal
The demo builds from CURRENT platform source, and the ambiguous UI defects are re-triaged on a correct build — so every downstream fix (M238–M243) is scoped against reality, not stale code. The M217/M222 "clean stage" pattern: any UI-defect triage on a stale-clone demo is untrustworthy. Only defect #1 (menu) was clone-staleness; #2–#5 each reproduce on `origin/main` — this milestone confirms that on a fresh build before the fixes fan out.

## Scope
### In
- Fix clone-freshness in `rext demo-stack/ensure-clones.sh`: a **fetch-verified** freshness assertion (never suppressed-stderr — the billion `root` host-key failure that produced the 12-vs-202 mismatch) + an opt-in advance-to-`origin/main`-or-pinned-tag path + a **real pin model** so "pinned" vs "stale-by-neglect" is distinguishable (today both read `ref:"main"`/`"HEAD"`).
- Fix **F-M236-CLOSE-2**: the R1 pristine sweep enumerates all **14** patch manifests, not the hard-coded 3.
- Bring up a **fresh-clone demo on billion**; produce a **confirmed-defect ledger**: verify #1 menu now hierarchical for managers; RE-TRIAGE #2 academy-language + #4 library-empty on the fresh build (which survive a correct build?).

### Out
- Any downstream fix (routed to M238–M243 by the re-triage).
- Any platform-repo edit.

## Open questions
- Which of #2 (academy-language) / #4 (library-empty) survive a fresh build? The re-triage decides — and routes the downstream fix scope.
- Advance-to-`origin/main` opt-in vs pinned-tag: what's the default freshness posture for a demo bring-up?

## Delivers
`corpus/ops/rosetta_demo.md` (the clone-freshness mechanism) + `corpus/ops/demo/demopatch-spec.md` (R1 all-14-manifests).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
