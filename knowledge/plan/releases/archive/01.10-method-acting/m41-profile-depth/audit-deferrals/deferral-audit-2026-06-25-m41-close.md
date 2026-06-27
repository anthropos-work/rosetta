---
title: "Deferral Audit — milestone M41 (close)"
date: 2026-06-25
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. M41 added **zero** deferrals of its own. The one open release-level
  deferral (DEF-M40-01, inherited) is a single deferral with a confirmed Fate-2 destination,
  unchanged since M40 close (2026-06-24) and not yet aged-out (its destination milestones
  M42e/M42m are still open).

## Summary
- Total deferrals in scope: 1 (all inherited; 0 originated by M41)
- Single deferrals: 1
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory
M41-originated deferrals: **none.** Sources walked:
- `progress.md` — all 5 section checkboxes (G3 work-history, G3 education, G5 verified bump,
  G5 claimed tail, Docs) checked; the `## M41: Hardening` section reports a clean Pass-1+2 stop
  (100% per-function coverage on both M41 files, 0 flakes, "no production bug surfaced"). No
  `Deferred` subsection.
- `decisions.md` — 7 decisions (M41-D1..D7); zero contain
  defer/postpone/later/out-of-scope/future-milestone language — all are design choices (the new
  ProfileSeeder surface, the NOT-NULL company FK, the provenance-edge tie, the never-clobber
  UPSERT guard, the verified-bump split, the combined pool, heroes-only).
- `overview.md` — the three `Out:` items (M40 serve-grant, M39 identity, the `mapped:22` wiring)
  are **release-sibling partitions / explicit leave-as-is**, not moved-out In: items. The three
  `Open questions` were all resolved in build (the ~90 split = M41-D5; N rows = 3 exp / 1 edu per
  D-record; per-population vs heroes-only = M41-D7 heroes-only).
- `spec-notes.md` — the `TODO:` markers are build-time **confirmation scratch** (e.g. "confirm the
  timeline reads ent.UserExperience via GET_TIMELINE"; "confirm mapped:22 feeds membership_skills").
  All were resolved during build: the LIVE-SCHEMA CORRECTIONS section supersedes the column-list
  guesses, and progress.md records every emitted row dry-insert-validated against the live demo-3
  schema. These are internal scratch notes in an archived artifact, not deferred work and not code
  TODOs.
- Code TODO/FIXME/HACK in M41 source (`profile.go`, `profile_write.go`, + the two test files):
  **zero** (grep clean).

Inherited (from prior v1.10 milestones):
```yaml
- id: DEF-M40-01
  item: 'KPI "AI simulations completed" = 0 on the profile/dashboard surface'
  origin_milestone: M40
  first_deferred_on: 2026-06-24
  last_seen_in: m40-directus-serve-grant/decisions.md (M40-D7)
  destination: "M42e + M42m (Fate-2 — already owned by their per-vantage coverage exit gate)"
  reason_recorded: "Source public.local_jobsimulation_sessions has no CMS dependency; a separate
    frontend/auth-context concern, not coupled to the serve-grant. The coverage sweep that exercises
    every page is its right home."
  partial_attempted: no
```
M39 (closed): re-audit GREEN at its own close; its two `Out:` items were Fate-2 (owned by M40/M41,
confirmed) — not deferrals. M42e/M42m: planned, unbuilt — no accumulated deferrals.

## Repeat-Deferral Patterns
None. DEF-M40-01 has been seen in exactly one milestone (M40); it does not recur in M41.

## Fate-1 Investigation
### DEF-M40-01 — KPI "AI simulations completed" = 0
- **Fate-1 (land now, complete) feasible:** no.
- **Why:** unchanged from the M40-close investigation. The KPI reads
  `public.local_jobsimulation_sessions` directly (jobsimulation), not any CMS/Directus surface and
  not any `user_skills`/`user_experiences` surface — so M41's ProfileSeeder rows cannot move it
  either. It is a frontend/auth-context rendering question that only manifests on a fully-walked,
  logged-in-as-hero page sweep.
- **Which fate applies:** Fate-2 — already owned by M42e/M42m's exit gate ("every reachable demo
  page renders non-empty semantic content, 0 failing"). A KPI rendering "0" on a reachable page is
  precisely an empty-section failure that gate catches; no `In:`-list edit needed (both are
  iterative; the gate is the commitment). No aging trigger fired: both destination milestones are
  still open, the area was not substantively re-touched by M41 (different schema surface), and the
  defer is < 1 day → < 2 milestones old.

## Recommendations
| Item | Verdict | Rationale |
|------|---------|-----------|
| DEF-M40-01 (KPI=0) | **LAND-NEXT** (Fate-2) | Already owned by M42e/M42m's coverage exit gate; no plan edit needed. Confirmed unchanged at this audit. |

## Applied Changes
- No plan edits. DEF-M40-01's Fate-2 destination is confirmed unchanged; both destination
  milestones already encompass it via their gate.
- This report written to the M41 `audit-deferrals/` subfolder.

## Blocking Items (require user decision)
None. No repeat deferrals, no aged-out items, no chronic patterns. Verdict GREEN.
