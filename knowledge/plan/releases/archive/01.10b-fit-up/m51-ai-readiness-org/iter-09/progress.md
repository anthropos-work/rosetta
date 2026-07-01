**Type:** tik (tooling-iter shape) — under **TOK-02** (the user-ratified app-read-path-demo-patch pivot).

# iter-09 — app read-path demo-patch: bound loadMembers in the snapshot path → GATE MET

## What was attempted
The user-chosen run-6 strategy (option (c) from the iter-08 user-blocker): author a NEW app injection
demo-patch that bounds the unbounded `loadMembers(orgID, "")` whole-org hydration in
`buildResponseFromSnapshots` (the frozen AI-readiness read path) — the org-scale wall iter-08 root-caused
(the frozen `?cycle=<closed>` GET timed out at 180s because the frozen scores read fast but then re-joined
the WHOLE org's members). Bound it to the ~199 snapshot users via the existing bounded sibling
`loadMembersByUserIDs` — a PURE perf optimization (data-identical), scoped to the snapshot read path.

## Phase A/B — design (read the app source; validate the minimal correct swap)
- `buildResponseFromSnapshots` (`ai_readiness.go:512`) uses the loaded `members` ONLY to build
  `membersByUserID` (keyed by `Member.UserID`), looked up ONLY by each snapshot's `s.UserID` (line 549).
  Members with no matching snapshot are loaded-but-never-used.
- `loadMembersByUserIDs(ctx, orgID, "", userIDs)` (`members.go:367`) — the existing bounded sibling used by
  the live scoring path — restricts `queryBaseMembers` to `memberships."user" = ANY($2::uuid[])` (indexed)
  and hydrates tags/skills/sims off the loaded rows' membership ids.
- So collecting the snapshot user-ids and calling `loadMembersByUserIDs` instead of `loadMembers` yields a
  `membersByUserID` with the SAME entries for every resolvable snapshot `UserID`; snapshot users with no
  active membership hit the SAME orphan branch (map miss) in both forms. The `AIReadinessResponse` is
  byte-identical; only the member-load cost drops from whole-org to the snapshot set.
- The patched form `go build ./internal/workforce/` exits 0 (`members`/`uuid`/`s`/`snapshots` stay used).

## Phase C — the tooling + the fix (authoring copy → consumption clone → demo build)
1. **Manifest** `demo-stack/patches/app-aireadiness-snapshot-loadmembers/app-aireadiness-snapshot-loadmembers.yaml`
   — a pinned `anchor → replacement` swap (pre_sha256 `8d509118…` @ v1.315.0, post_sha256 `c87237c2…`,
   post_marker), validated byte-exactly through `manifest_loader` (pre matches, anchor unique, apply → post,
   marker exclusive).
2. **Helper** `stack-injection/apply-app-aireadiness-loadmembers.sh` — mirrors the `apply-app-authz-skip.sh`
   precedent (same guards: drift-refuse, single-occurrence, post-condition sha, idempotent no-op). Verified
   end-to-end on the live scratch clone (applies → post sha; second run no-ops exit 0).
3. **Wiring** in `demo-stack/up-injected.sh`'s inject loop (svc=app, AFTER apply-authn + the authz-skip,
   BEFORE the build), non-fatal + `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1` opt-out.
4. Committed in the authoring copy (`089eb96`); re-pinned the consumption clone via a clean ref-switch;
   `/demo-down 1 --purge` + `/demo-up 1` — the inject loop logged **"app: applied ai-readiness snapshot-read
   loadMembers bound"** right after the authz-skip → the app image baked WITH the patch. Re-seeded the
   AI-readiness showcase (DB re-verified: cycle closed `95d9fc3d…`, 199 frozen snapshots, funnel 22/21/156 =
   78.4% stage-3).

## Phase D — verify (cheap probe FIRST) → hypothesis CONFIRMED
`probe-aireadiness-deeplink.spec.ts` (the iter-08 reusable dual-endpoint probe):
- `DIRECT /cycles` → 200 in 17 ms (unchanged).
- **`DIRECT data ?cycle=<closed>&includePeople=true` → 200 in 19 ms** (was NEVER completing / 180 s timeout in
  iter-08) — returns the correct frozen funnel (`members:199, archetypeCounts, score:49`).
- The FE default GET (no `?cycle=`) → 200 in 93 ms; `main` renders the full funnel ("AI Readiness | Overall
  org readiness | 50/100 | Members | 199 | Functional+ 173/199 (87%) | Steps completion | Stage 1 — skills
  only (13) | Stage 2 — + AI Path (10) | Stage 3 — + AI Interview…").
- VERDICT `frozenFast=true frozenRendered=true` — passed. The wall is cleared.

## Phase A/D — the gated sweep (2 sweeps, both frontier-exhausted)
- **Sweep #1:** `failingSections 5 → 1`, escapes 0, persona 0, notReached 0, **frontier EXHAUSTED**,
  reachable 70. The loadMembers-bound patch cleared 4 of 5 (the AI-readiness org-score section + the 3
  workforce aggregates). The 1 residual: the `ai-readiness-funnel` section `empty` — "region missing required
  content: Stage breakdown" (re-asserted 6× — genuinely below bar, NOT a paint-timing skeleton).
- **Triage of the 1 residual (a manifest FALSE-FAIL, not a content gap):** `AIReadinessView.tsx` (lines
  162-183) renders the funnel header as ONE of two MUTUALLY-EXCLUSIVE strings — `t('stepsCompletionLink')`
  ("Steps completion", a link) when `onStepsClick` is provided, ELSE `t('stageBreakdown')` ("Stage
  breakdown"). The manager dashboard always provides the steps-completion drawer handler → it renders "Steps
  completion" and NEVER "Stage breakdown". The descriptor required BOTH — impossible in manager mode. The
  funnel WAS fully rendered (Stage 1/2/3 + real counts + the "Steps completion" link, proven by the probe's
  `main` dump). Harness fix: drop the impossible "Stage breakdown" substring; keep the load-bearing proof
  (the three stage labels + the funnel header). Committed to `coverage-manifest.ts` (`3733523`).
- **Sweep #2 (re-measure after the harness fix):** `failingSections 0`, escapes 0, persona 0, notReached 0,
  **cappedAtFrontier=false (EXHAUSTED)**, reachable 70, **gateMet=true**. `GATE: MET ✅`.

## Qualitative gate conditions (re-confirmed on the fresh demo-up)
- Dashboard ENABLED (`organization_settings.ai_readiness = t`).
- ~80% all-3-complete: 156/199 = **78.4% at stage 3**; stage-1 cohort (22) = STARTED, stage-3 (156) =
  COMPLETED; the two heroes (Aria stage-3/COMPLETED, Ben stage-1/STARTED) preserved.
- Manager-vantage coverage `(0,0)` frontier-exhausted on a FRESH demo-up.

## Close — 2026-07-01

**Outcome:** the app read-path demo-patch bounds the frozen-read member hydration
(loadMembers whole-org → loadMembersByUserIDs over the ~199 snapshot users), clearing the org-scale wall:
the frozen `?cycle=<closed>` GET went 180s-timeout → 19 ms and the dashboard renders the full funnel. Plus a
coverage-manifest funnel-descriptor correction (the header is one of two mutually-exclusive strings). The
gated manager-vantage sweep on Northwind reaches **`(failingSections, escapes) = (0, 0)`**, persona 0,
frontier EXHAUSTED, on a fresh demo-up. Metric delta: **5 → 0**. Gate MET.
**Type:** tik (tooling-iter shape — shipped the new app injection demo-patch + used it within the iter)
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (design the minimal data-identical loadMembers→loadMembersByUserIDs swap from the app
source), D2 (author the manifest+helper+wiring following the app-targetrole-authz-skip precedent), D3 (triage
the 1 residual as a manifest false-fail: the mutually-exclusive funnel header — harness fix, not a seed gap).
**Side-deliverables (if any):** none — the coverage-manifest funnel-descriptor correction is part of the
iter's planned coverage-drive scope (the last failing section), not an unrelated side-fix.
**Routes carried forward:** none — the gate is MET. (The prod finding stays documented: prod's frozen read
still hydrates the whole org and needs loadMembers bounded in the snapshot path / a frozen_tags column,
M314b — a disclosed demo-perf relaxation, not a prod fix.)
**Lessons:**
  - The iter-08 root cause (the unbounded `loadMembers` in the frozen path) had an EXACT bounded sibling
    (`loadMembersByUserIDs`) already in the codebase used by the live scoring path — so the fix was a
    one-call data-identical swap, not a new query. When a perf wall is an unbounded hydration, look for an
    existing id-restricted loader before writing one.
  - A believability-gate descriptor that requires TWO strings which the FE renders MUTUALLY EXCLUSIVELY is a
    latent false-fail. When a section fails on one substring while its siblings render, check whether the
    missing string is an EITHER/OR alternative of a present one (a conditional header), not an absent section.
    Added to the coverage-protocol lessons.
