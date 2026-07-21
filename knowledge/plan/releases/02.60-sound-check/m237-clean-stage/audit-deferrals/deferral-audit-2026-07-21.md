---
title: "Deferral Audit — M237 clean stage (milestone scope)"
date: 2026-07-21
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no chronic patterns; no aged-out items.
- M237 is the FIRST milestone of v2.6 "sound check" — there are no prior-milestone inherited deferrals to re-audit.
- The confirmed-defect ledger's downstream routings are Fate-2 confirmations (already-owned by M238/M239), not deferrals.

## Summary
- Total deferrals in scope: 0 genuine deferrals
- Single deferrals: 0
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Fate-2 confirmations (already-owned, no edit needed): 2 (#2→M238, #4→M239)

## Deferral Inventory
Sources walked: M237 `progress.md`, `decisions.md`, `overview.md`, `spec-notes.md`; TODO/FIXME/HACK grep over the two touched rext files (`demo-stack/ensure-clones.sh`, `demo-stack/tests/test_tooling.py`).

- **No `defer` / `postpone` / `later` / `out of scope` / `future milestone` language** in the M237 docs.
- **No TODO/FIXME/HACK** in the touched rext code (grep clean).
- **overview.md `Out:` list** — "any downstream fix (routed to M238–M243)" + "any platform-repo edit" — are the milestone's designed scope boundary (M237 is a re-triage barrier, not a fixer), not scope that was moved out mid-milestone.
- **spec-notes.md open questions** — both RESOLVED: "which of #2/#4 survive a fresh build?" → answered by the ledger; "advance opt-in vs pinned default?" → answered by D1 (opt-in, default OFF).

## Repeat-Deferral Patterns
None. (M237 is the release's opening milestone; nothing has been carried across ≥2 milestones.)

## Fate-1 Investigation

### Confirmed-defect ledger routing — #2 (academy language) → M238
- **Fate-1 (land now) feasible:** no — correctly Fate-2.
- **Why not Fate-1:** #2 lives on the separate, independently-degraded ant-academy app (a `/it` 404 + non-functional flag control on a 5-behind academy), NOT on the next-web clone M237 rebuilt. M237's deliverable was the *re-triage* (which it delivered — the defect SURVIVES a correct build and is a real academy-surface defect). The fix is academy-reliability work.
- **Owner:** M238 `overview.md` In-list line 23 already owns "Fix #2 (language error — re-triaged in M237)". Fate-2 CONFIRMED, no plan edit needed.

### Confirmed-defect ledger routing — #4 (library empty-first-load) → M239
- **Fate-1 (land now) feasible:** no — correctly Fate-2.
- **Why not Fate-1:** on the verified-current build #4 does NOT reproduce as empty (library renders "Public Content (22)", 7→29 cards, 0 gql/http errors); the residual is a possible sub-second cold-first-load flash the re-triage could neither confirm nor rule out. That is library-first-load work, not barrier work.
- **Owner:** M239 `overview.md` In-list line 26 already owns "Fix #4 (library empty-first-load)". Fate-2 CONFIRMED (re-scoped down to the cold-flash/client-race question), no plan edit needed.

### D-HARDEN-2 — fetch-OK-but-uncountable classifies as `fresh`
- **Not a deferral.** An ACCEPTED observed edge-semantic, pinned by a test so any future change is deliberate. Explicitly "no downstream milestone owns it (not a defect, an accepted semantic)". Reviewed at close (see close decision triage) — acceptance stands.

### Hardening "stop condition" untested branches
- **Not a deferral.** The defensive `|| true` / `2>/dev/null` non-fatal branches were deliberately RECORDED-not-shallow-tested per the harden-milestone "don't shallow-test to bump a number" guidance. A conscious close-time decision, not deferred work.

## Recommendations
- #2 → **LAND-NEXT (Fate 2)** — already owned by M238. No edit.
- #4 → **LAND-NEXT (Fate 2)** — already owned by M239. No edit.
- D-HARDEN-2, hardening stop-condition → no action (not deferrals).

## Applied Changes
None. Both Fate-2 items are already in their owners' `In:` lists (verified); no plan edit, no new decision record required.

## Blocking Items (require user decision)
None.
