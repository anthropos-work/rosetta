# M248 — Spec notes

Topic → doc → code findings for the content-stories manager result-link re-point accumulate here during build.

## Pre-flight audits — Re-point the sim ManagerResultPath builder
`/developer-kit:audit-kb-fidelity --milestone=M248` → **YELLOW — proceed with tracking** (0 blockers,
0 stale load-bearing claims). Findings tracked as KB-1..KB-4 in `decisions.md`. Report kept inline (run
scoped read-only). Core premise code-verified: the `/sim/<slug>/<userId>/result/<sessionId>` manager view
exists on the build source (web + hiring), renders `AISimulationResultContainer isManagerView shareView
userId`, and reads a persisted row via `JobSimulationResult(sessionId)` (a plain Ent SELECT keyed by
sessionId — no recompute). Docs↔current-projection aligned today (both name the activity-dashboard scoreboard);
re-pointing them is the milestone's own delivery.

## Verify-interview — the rung-0 STATIC read (D1) was OVERTURNED by the LIVE render (D3)
> **⚠ D1 below is SUPERSEDED by D3 (see `decisions.md`).** The rung-0 CODE READ concluded interview renders on
> the unified `/sim` route (D1, collapse). The **LIVE render-confirm on demo-2** found the `/sim` interview
> manager surface renders a **"Coming Soon"** placeholder (`interviewExtractionData` null — the report is
> flag/data-gated and not reliably fetched on a demo), NOT the report; the dedicated
> `/enterprise/activity-dashboard/interviews/<simId>/<membershipId>` route lands reliably (M236). **Final: KEEP
> interview on its dedicated route; only the NON-interview family moves to `/sim`.** The static read below is
> retained for the record.

**Decision D1 (rung-0 static read, SUPERSEDED): collapse the per-sim_type kind branching; NO interview split
path.** The interview MANAGER report surfaces on the SAME `/sim/<slug>/<userId>/result/<sessionId>` route as
every other sim_type, via the shared `AISimulationResultContainer` with `isManagerView=true`:
- `.../sim/[slug]/[userId]/result/[sessionId]/page.tsx` renders `<AISimulationResultContainer isManagerView
  shareView userId sessionId slug isHiring={false} />`.
- `AISimulationResultContainer.tsx:495-539`: when `isManager && isInterviewType`, `isManagerReportEnabled =
  isManager && posthog.isFeatureEnabled('flag_interview_manager_report')` → `isExtractionEnabled` →
  `useGetInterviewExtractionReport({ sessionId, isManagerView: isManager })` → renders
  `interviewExtractionManagerReport.report` (L530-531). Keyed by **sessionId** (robust to `<userId>`).
- The two existing demopatches already force the flag ON for the manager scope on a demo:
  `next-web-interview-flag-container` widens the FETCH gate (anchor
  `isInterview && (isManagerReportEnabled || isPlayerReportEnabled);`) and `next-web-interview-flag-result`
  widens the RENDER gate — **scoped to `isManagerScope`** at M244 iter-08. So a demo interview MANAGER report
  fetches + renders; the player interview stays a terse acknowledgement. Nothing new needed for interview.

⇒ ALL sim types (assessment/training/hiring/interview) → `/sim/<slug>/<userId>/result/<sessionId>`. The
`simulationManagerKind` per-sim_type refinement (interview → `interviews`) is deleted — the URL no longer
carries a kind segment. `meta.managerKind` stays, but only as the **has-manager-view GATE** (non-empty for
the simulation product; "" for skill-path-legacy/academy/ai-labs).

## The role-gate on the `/sim` manager route (KB-3) — favorable
`.../[userId]/result/layout.tsx` (`ManagerResultGateLayout`) gates with `notFound()` unless the viewer is
`MembershipRoles.Admin` (ANY user's result) or `ContentCreator` (own only, `routeUserId === externalId`).
This is a **different** gate than the activity-dashboard's Casbin `OrgFeatureInsights`. The content-stories
manager seat (`dan-manager`) is seeded vantage-faithfully as **admin** → the admin branch grants access to the
player's result. So the manager CTA lands and renders. (Frontend-only gate; the GraphQL layer is the real
boundary — a PR caveat noted in the layout, out of scope here.)

## `<userId>` = `owner.UserID` — why it renders NON-EMPTY
`owner.UserID = deterministicUUID("<prefix>:user:<index>")` = the seeded `public.users.id`. The
`ContentStorySeeder` re-owns the cloned session's `owner_id` to that same id, so:
- the RESULT renders by **sessionId** (`JobSimulationResult(sessionId)`) — independent of `<userId>`;
- the session-scoping queries (`useGetSession`/`useGetSessions`, keyed by `userId`) also match, because the
  session's `owner_id` IS `owner.UserID`;
- the layout's `externalId === routeUserId` ownership check is BYPASSED for the admin manager seat.
The player's Clerk auth id (`user_seed_<host>_<idx>`) differs from `owner.UserID`, but nothing on the manager
result render path compares against it. **Confirm NON-EMPTY at live render on demo-2** (open question #2).

## Core projection — `ManagerResultPath` builder (`content_manifest.go:411-423`)
New: `row.ManagerResultPath = fmt.Sprintf("/sim/%s/%s/result/%s", cs.SimSlug, owner.UserID,
contentStorySessionID(cs))` — the SAME slug + session id the player path uses (single source), plus the
owner's user id as the `<userId>` segment. Gated by `meta.managerKind != "" && mgrKey != ""` (simulation
product with a manager hero). Fail-closed: a simulation exhibit with a blank `sim_slug` drops loud (the
manager route resolves by slug, exactly like the player route). Delete `simulationManagerKind`;
`owner.MembershipID` stays (still used by `content_nonsim.go:303`, out of M248 scope).

## Grader + unit-spec + manifest regeneration
The manager render surface CHANGES from the activity-dashboard scoreboard (shapes 4/5) to the
`AISimulationResultContainer`:
- **`manager-scored`** (was `manager-dashboard`): the `/sim` manager view of a NON-interview sim renders the
  SAME scored payload the player sees (score + Evaluated Skills + feedback) — so it reuses the `player-scored`
  assertion (readable ≥ 300; Evaluated-Skills section, or feedback corroborated by an N/100 score).
- **`manager-interview`** (repurposed): the `/sim` manager view of an INTERVIEW renders the
  `interviewExtractionManagerReport` (a substantial structured report), NOT the old breadcrumb + "View Report"
  table. Assertion rewritten: not bounced/not-found, not "undefined undefined", report content present.
- `shapeFor('manager', …)` selects by the session's STATED `sim_type` (interview → `manager-interview`, else
  → `manager-scored`) — route is `/sim/` for every sim_type now, same precedence as the player branch.
- Anti-fall-through (manager twin): `manager-scored` fall-through only for a `/sim/` route.
- Pair count UNCHANGED (47): `buildPairs` forms the manager pair from `has_manager_view + path + seat`
  (route-shape-agnostic); the re-point does not add/drop a pair. `content-denominator.json` stays 47.
- Regenerate `presets/content-manifest.json` via `stackseed --content-export` (honesty gate +
  language-consistency gate must pass).

## Docs — `content-stories-spec.md` + `content-stories-routes.md`
Re-point the manager result route to `/sim/<slug>/<userId>/result/<sessionId>` (per-session manager result
view, `isManagerView`, persisted read), collapse the interview split, record the `layout.tsx` admin/content-
creator gate, and fix the §2 line-91 `<userId>`-vs-membership placeholder drift (KB-4).

## LIVE render-confirm findings (demo-2) — supersede the static grader plan above
- **Manager routing (D3):** NON-interview (assessment/training/hiring) → `/sim/<slug>/<userId>/result/<sessionId>`;
  INTERVIEW → `/enterprise/activity-dashboard/interviews/<simId>/<membershipId>` (kept — its `/sim` report is
  flag/data-gated, renders "Coming Soon" on a demo).
- **`manager-scored` grader (calibrated LIVE):** the `/sim` manager scored view renders the persisted SCORE + a
  performance narrative, with "Evaluated Skills" COLLAPSED behind a "Show Details" toggle and feedback in the
  session's LANGUAGE (Italian for IT sessions). So the player-scored English anchors ("Evaluated Skills",
  "feedback") false-FAIL a fully-rendered result (train-doc-pass: 5406 chars, score 82/100, no EN anchors in
  the DOM). `manager-scored` therefore keys on the **SCORE (N/100)** — language-agnostic + collapse-proof
  (`hasScore || hasSkills`, readable ≥ 400, not "undefined undefined"). `gradeScored` stays the PLAYER-scored
  gate only.
- **`manager-interview` grader:** restored to the M236 activity-dashboard breadcrumb + attempts-table + "View
  Report" shape (interview keeps that route).
- **Sweep result:** warm 43/47. Direct drives prove the `/sim` manager route renders full scored results
  (asmt 4516 · train 5406 · asmt-voice-fail 2981). Residual: 3 non-interview managers render a header-only
  shell at the sweep's settle budget (per-session render state, not an M248 code defect) + 1 academy player
  env failure (`:23077` down). → CARRY-M248-01 (M254 fresh-seed re-confirm).
