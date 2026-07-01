---
iter: 09
milestone: M51
iteration_type: tik
iter_shape: tooling
status: in-progress
created: 2026-07-01
---

# iter-09 — app read-path demo-patch: bound loadMembers in the snapshot path

**Type:** tik (tooling-iter shape — ships a new app injection demo-patch + uses it within the iter for the
coverage drive) under **TOK-02** (the user-ratified app-read-path-demo-patch pivot).

## Active strategy reference
**TOK-02** (milestone-root `decisions.md`) — author a new app injection demo-patch
`app-aireadiness-snapshot-loadmembers` that bounds the unbounded `loadMembers(orgID, "")` in
`buildResponseFromSnapshots` to the ~199 snapshot user-ids via the existing bounded sibling
`loadMembersByUserIDs`, then drive the manager-vantage coverage sweep 5 → 0.

## Step 0 — Re-survey (done in Before-You-Start)
The metric baseline is `(failingSections=5, escapes=0)` frozen-in-DB-correct (closed cycle, 199 snapshots).
The TOK-02 named target (the `loadMembers` unbounded call at `ai_readiness.go:520`) is confirmed present +
un-patched in the demo's build-scratch app clone. The bounded sibling `loadMembersByUserIDs`
(`members.go:367`) exists and is used by the live scoring path. Target is current + meaningful.

## Cluster / target identified
The single `m.loadMembers(ctx, orgID, "")` call at `ai_readiness.go:520` inside `buildResponseFromSnapshots`.
Root-caused by iter-08 as the org-scale wall on the frozen read path. The 5 failing sections (2 AI-readiness
+ 3 workforce aggregates) all funnel through this member-hydration family, so bounding it is hypothesized to
clear all 5.

## Hypothesis
Replacing the unbounded whole-org hydration with the bounded snapshot-user-id hydration makes the frozen
`?cycle=<closed>&includePeople=true` GET complete FAST (was 180s timeout) while returning byte-identical data
→ the dashboard renders → the sweep clears the 2 AI-readiness sections + the 3 workforce aggregates → 5 → 0.

## Expected lift
failingSections 5 → 0 (escapes stay 0). The gate's qualitative conditions are already DB-correct.

## Phase plan
- **Author (tooling step 1):** the manifest `demo-stack/patches/app-aireadiness-snapshot-loadmembers/` (anchor
  → replacement + pre/post sha256 + post_marker), the `stack-injection/apply-app-aireadiness-loadmembers.sh`
  helper (mirroring the authz-skip helper), and the `up-injected.sh` inject-loop wiring (svc=app, after
  apply-authn + the authz-skip, before build, non-fatal, `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1` opt-out).
  Commit in the authoring copy. Re-pin the consumption clone to the new sha (clean ref-switch).
- **Phase C (re-apply):** `/demo-up 1` (or targeted app-image rebuild through the inject loop) bakes the patch
  + re-seed the AI-readiness showcase org.
- **Cheap probe FIRST:** run `probe-aireadiness-deeplink.spec.ts` — confirm the frozen GET now completes fast
  + returns the correct frozen funnel. Only then the gated sweep.
- **Phase A/D (sweep):** the gated manager-vantage sweep on Northwind (detached + polled), read `(failing,
  escapes)`.
- **Phase E (close):** grade on whether the 5 cleared; if `(0,0)` frontier-exhausted → tag `fit-up-m51` +
  re-pin the consumption clone to the tag + gate-met.

## Escalation conditions
- If the bounded swap cannot be authored without a platform-repo edit (the injection can't reach the call) →
  escalate `re-scope-trigger` (`unimplementable-without-platform-edit`), do NOT edit the platform.
- If the app image build fails through the inject loop → user-blocker (the same class as a RED test gate).

## Acceptable close-no-lift outcomes
- If the probe shows the frozen GET STILL times out after the patch bakes (the swap was applied but a DIFFERENT
  unbounded cost remains in the path) → closed-no-lift with the documented falsification (measure the branch
  end-to-end again), route the residual forward.
