# iter-09 — decisions

## D1 — the minimal, data-identical loadMembers → loadMembersByUserIDs swap

**Context:** iter-08 root-caused the frozen AI-readiness read wall to `buildResponseFromSnapshots` calling
`m.loadMembers(ctx, orgID, "")` — a full unbounded whole-org member hydration — to re-join current
Tags/Name/Role onto each snapshot. At ~200 members that hydration times out (180s+).

**Options:**
- (a) Add a new bounded query in the patch. Rejected — more surface, more risk; and a bounded loader already
  exists.
- (b) Reuse the existing bounded sibling `loadMembersByUserIDs(ctx, orgID, tag, userIDs)` (used by the live
  scoring path `computeOrgBreakdowns`), passing the snapshot user-ids. Chosen.

**Choice:** collect `snapUserIDs` from the already-loaded `snapshots` slice, then
`m.loadMembersByUserIDs(ctx, orgID, "", snapUserIDs)` instead of `m.loadMembers(ctx, orgID, "")`.

**Why:** in `buildResponseFromSnapshots` the loaded members are used ONLY to build `membersByUserID` (keyed
by `Member.UserID`), looked up ONLY by each snapshot's `s.UserID`. Members with no matching snapshot are
loaded-but-never-used. `loadMembersByUserIDs` restricts `queryBaseMembers` to
`memberships."user" = ANY($2::uuid[])` (indexed) + hydrates off the loaded rows' membership ids — so the map
carries the SAME entries for every resolvable snapshot UserID; orphan snapshots hit the SAME map-miss branch
in both forms. The response is byte-identical; only the member-load cost drops. `go build ./internal/workforce/`
exits 0 (all identifiers stay used). Scoped to the snapshot/closed-cycle read path — the live/active-cycle
path for other orgs is untouched.

## D2 — author the patch via the app-targetrole-authz-skip precedent (manifest + helper + inject-loop wiring)

**Context:** the demopatch tool refuses out-of-workspace build-scratch clones (path firewall), so an app-repo
source patch inside the inject loop needs a rext-owned `apply-*.sh` helper mirroring a pinned manifest — the
exact `app-targetrole-authz-skip` mechanism.

**Choice:** a new patch dir `demo-stack/patches/app-aireadiness-snapshot-loadmembers/` (a pinned
anchor→replacement manifest, pre/post sha256 + post_marker, validated through `manifest_loader`) + a helper
`stack-injection/apply-app-aireadiness-loadmembers.sh` (same guard shape as `apply-app-authz-skip.sh`) + the
`up-injected.sh` inject-loop wiring (svc=app, after apply-authn + the authz-skip, before the build, non-fatal,
`DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1` opt-out).

**Why:** reuses the proven, guarded, canonical-repo-never-touched vehicle; the manifest stays the single
source of truth; only the target file (ai_readiness.go) + the manifest differ from the precedent. Verified
end-to-end (helper applies → exact post sha; idempotent no-op on re-run; the inject loop logged the apply).

## D3 — the 1 residual after sweep #1 is a MANIFEST false-fail (the mutually-exclusive funnel header)

**Context:** sweep #1 moved `failingSections 5 → 1`; the 1 residual was the `ai-readiness-funnel` section
`empty` — "region missing required content: Stage breakdown" (re-asserted 6× — genuinely below bar).

**Investigation:** the probe's `main` dump proved the funnel WAS fully rendered ("Steps completion | Stage 1
— skills only (13) | Stage 2 — + AI Path (10) | Stage 3 — + AI Interview…"). `AIReadinessView.tsx`
(lines 162-183) renders the funnel header as ONE of two MUTUALLY-EXCLUSIVE strings: `t('stepsCompletionLink')`
("Steps completion", a link) when `onStepsClick` is provided, ELSE `t('stageBreakdown')` ("Stage breakdown").
The manager dashboard always provides the steps-completion drawer handler → it renders "Steps completion" and
NEVER "Stage breakdown". The manifest descriptor required BOTH — impossible in manager mode.

**Choice:** a harness fix (not a seed fix) — drop the impossible-in-manager-mode `'Stage breakdown'` from the
`ai-readiness-funnel` descriptor's `mustInclude`, keeping the load-bearing proof (`Stage 1`/`Stage 2`/`Stage
3` + the funnel header `Steps completion`). NOT a gate-loosening: the funnel's real proof (the three stage
labels rendering with real counts) is still asserted; only an over-strict impossible substring is removed.

**Why:** requiring two strings the FE renders mutually exclusively is a latent false-fail. Sweep #2 confirmed
`failingSections 0`, gate MET. Routed to the coverage-protocol lessons.
