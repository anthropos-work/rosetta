# M248 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## D1 (rung-0 static read) — SUPERSEDED by D3 (live evidence)
The rung-0 CODE READ concluded the interview MANAGER report renders on the unified
`/sim/<slug>/<userId>/result/<sessionId>` route via `AISimulationResultContainer(isManagerView=true)` →
`interviewExtractionManagerReport` (`AISimulationResultContainer.tsx:530-531`), gated by
`flag_interview_manager_report` which the `next-web-interview-flag-{container,result}` demopatches force ON. On
that read I first collapsed the per-sim_type branching (interview → `/sim`). **The LIVE render-confirm on
demo-2 overturned this** — see D3. The final implementation KEEPS interview on its dedicated route.

## D3 — KEEP interview on the activity-dashboard route (verify-interview, resolved LIVE)
The live render-confirm on demo-2 (direct browser drive as `dan-manager`) found the `/sim` interview manager
surface renders a **"Coming Soon"** placeholder (`interviewExtractionData` null — the report is flag/data-gated
and NOT reliably fetched on a demo), NOT the extraction report. The dedicated
`/enterprise/activity-dashboard/interviews/<simId>/<membershipId>` route LANDS reliably (viewReport + completed,
the M236 shape, no flag dependency). The milestone spec's explicit conditional — "keep interview split IF
verify-interview says so" — therefore resolves to **KEEP**. Only the NON-interview family (assessment / training
/ hiring, 21 of 23 manager pairs) moves to `/sim`; interview stays. `simulationManagerKind` stays deleted (the
interview branch inlines the `interviews` segment); `owner.MembershipID` is used again for the interview path.

## D2 — Manager grader shapes (final, per D3 mixed routing + live calibration)
`manager-dashboard` → **`manager-scored`** for the NON-interview `/sim` result view; **`manager-interview`**
stays the M236 activity-dashboard breadcrumb+attempts-table shape (interview keeps that route, D3). `shapeFor`
selects by stated `sim_type` (interview → `manager-interview`, else → `manager-scored`). **Live calibration
(demo-2):** the `/sim` manager scored view collapses "Evaluated Skills" behind a "Show Details" toggle and
renders feedback in the session's LANGUAGE (Italian for IT sessions), so the player-scored English anchors
false-FAIL a fully-rendered result — `manager-scored` therefore keys on the **SCORE (N/100)**, language-agnostic
and collapse-proof (`hasScore || hasSkills`, readable ≥ 400, not "undefined undefined"). `gradeScored` stays the
PLAYER-scored gate only (unchanged, proven). Pair count unchanged (47) — `buildPairs` is route-shape-agnostic.

## D4 (close-time, from the Phase-1b deferral audit) — CARRY-M248-01 is Fate-2, not Fate-3
The one deferral M248 introduces — re-confirm the content-stories manager pairs land on the FRESH `billion`
reset-to-seed — is **Fate-2 (confirmed-covered by a future milestone of this release)**, not Fate-3
(annotate-attach). M254's exit gate ALREADY names it: part **(b)** "the content-stories manager CTA lands on the
/sim per-session manager result view (non-empty) for sim products" + part **(h)** the live content-stories sweep.
No `overview.md` edit to M254 is required (it already owns the coverage). The build-time `progress.md` label
("Fate-3") was imprecise and is reconciled to Fate-2 at close. The 3 demo-2 header-only-shell renders + 1 academy
`:23077` env failure are demo-2 (M246-era warm seed / host state) environmentals, not M248 code defects — same
Fate-2 → M254 fresh-seed re-confirm. Audit report:
`audit-deferrals/deferral-audit-2026-07-23-m248-close.md` (verdict YELLOW, 0 blockers).

## KB items (from the Phase 0b KB-fidelity audit — YELLOW, tracked)
- **KB-1** — Neither `content-stories-spec.md` nor `content-stories-routes.md` documents the
  `/sim/<slug>/<userId>/result/<sessionId>` manager surface (both name only the activity-dashboard scoreboard).
  This is exactly what M248 delivers → resolved in Phase 5 (docs rewrite).
- **KB-2** — Open question: is `user.externalId == owner.UserID` at live render? Statically unverifiable.
  Mitigated: the result renders by **sessionId** (robust), and the session's `owner_id` IS `owner.UserID`, so
  the session-scoping queries match too. The layout's ownership check is bypassed for the admin manager seat.
  → confirm NON-EMPTY render on demo-2 (LIVE render-confirm).
- **KB-3** — The `/sim` manager route's `layout.tsx` gate is `MembershipRoles.Admin` (any user) /
  `ContentCreator` (own only), NOT the Casbin `OrgFeatureInsights` the docs cite. FAVORABLE: the manager seat
  `dan-manager` is seeded as **admin** → admin branch grants access. Captured in the doc rewrite (Phase 5).
- **KB-4** — Incidental pre-existing drift: `content-stories-spec.md` §2 (schema example, ~line 91) shows the
  manager path ending in `<userId>` where the code used the membership id. M248 rewrites this line anyway
  (manager path now genuinely ends in `<sessionId>`), so the drift is corrected in passing.
