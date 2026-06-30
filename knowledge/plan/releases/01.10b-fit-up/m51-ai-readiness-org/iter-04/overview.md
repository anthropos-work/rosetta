---
iter: iter-04
milestone: M51
iteration_type: tik
status: closed-no-lift
created: 2026-06-30
---

# iter-04 — tik (the base-Workforce PERF-WALL: re-up demo-1 with the M46/M50 perf demo-patches)

## Type
tik — under TOK-01 (active-cycle signals-true). The coverage-drive strand: triage → fix the highest-leverage
cluster from the iter-03 sweep. coverage-protocol Phase A–E.

## Step 0 — Re-survey before targeting
Re-confirmed against the iter-03 GATED sweep + the live state:
- The 6 failing sections are ALL the base-Workforce org-scale PERF-WALL (verification-funnel + talent-languages
  on `/enterprise/workforce`, members-roster + members-location-values on `/enterprise/members`, assign-roster
  on `/enterprise/assignments`, activity-table on `/enterprise/activity-dashboard`) — `kind:empty` "genuinely
  below bar (re-asserted 6×)". The DB HAS the data (200 memberships, 1330 membership_skills, EU member locations,
  782 jobsim sessions); the screenshot shows real chrome + skeleton rows + "0/∞"; the router logs 11.4s/4.2s
  GraphQL. This is the M46 perf-wall, diagnosis-confirmed.
- demo-1 consumes rext `fit-up-m49`; the perf demo-patches (`next-web-members-pagination` `limit 1000→30` +
  `app-targetrole-authz-skip`) exist in the authoring copy + are tagged at `fit-up-m50`. demo-1 was brought up
  BEFORE those landed in its consumed tag → the perf-wall is present.
Target is untouched + meaningful. No substitution.

## Active strategy reference
**TOK-01** (`../decisions.md`) — active-cycle signals-true. This iter is the coverage-drive strand (TOK-01
step 4): run the sweep, triage by the fix-surface routing table, apply the routed fix, re-sweep. The fix surface
here is the **demo-UP path** (the "Org-scale grid perf wall" row of the coverage-protocol routing table), not
`stack-seeding` — the data is already correctly seeded (iter-03).

## Cluster / target identified
The single highest-leverage cluster: ALL 6 failing sections are the same root cause (the org-scale perf-wall at
200 members). One fix — bringing demo-1 up with the M46/M50 perf demo-patches baked into its injected next-web +
app images — clears all 6 at once (the M46 close took demo-3's members query 76.7s→0.51s and cleared
`/enterprise/members` + `/enterprise/activity-dashboard` + `/enterprise/settings`).

## Hypothesis
Re-pin demo-1's consumed rext tag to one carrying the perf-wall demo-patches in its `up-injected.sh` inject loop
(`fit-up-m50`, or the eventual `fit-up-m51`), then `/demo-down 1` + `/demo-up 1` so the injected next-web
(pagination `limit 1000→30`) + app (targetRole read-gate authz-skip) images rebuild with the patches baked +
the post-seed FK indexes apply. The fresh demo-up re-seeds (including the iter-03 config+funnel seeders now in
the consumed tag IF re-pinned to a tag carrying them — else re-seed from the authoring copy as iter-03 did).
Then re-run the GATED manager-vantage sweep → expect the 6 skeleton false-fails to flip to real rows
(screenshot: rows not skeleton; GraphQL latency <5s) → `(failing≈0, escapes=0)`.

## Expected lift
A net reduction of the 6 perf-wall failing sections toward 0 (the M46 precedent cleared the analogous grids
entirely). Caveat: the AI-readiness funnel section is NOT yet asserted by the manager manifest — so even a
cleared perf-wall leaves the gate's AI-readiness ASSERTION unestablished (a separate iter-04+/iter-05 manifest
task). A net `(failing)` drop is the success signal; reaching `(0,0)` here is plausible if all 6 are purely
perf-wall.

## Phase plan
- Phase A (sweep): inherit the iter-03 GATED `(6,0)` as the pre-iter metric.
- Phase B (triage): done in Step-0 — all 6 = perf-wall, routed to the demo-UP fix surface.
- Phase C (fix): re-pin demo-1's consumed rext tag to one with the perf demo-patches; `/demo-down 1` +
  `/demo-up 1` (rebuilds injected next-web+app with the patches baked + applies post-seed FK indexes); re-seed
  the AI-readiness config+funnel (authoring-copy stackseed, as iter-03, if the consumed tag lacks them);
  re-export roster+cockpit; verify the perf-wall cleared by screenshot + GraphQL latency.
- Phase D (re-sweep): GATED manager-vantage sweep on demo-1 → record `(failing, escapes)` delta.
- Phase E (close): grade on whether the perf-wall sections cleared; route the AI-readiness manifest-assertion
  task + any residual forward.

## Escalation conditions
- If the perf demo-patches require a CANONICAL platform-source edit (not the ephemeral build-clone demopatch) →
  re-scope-trigger (NOT expected — M46/M50 proved these are demo-local build-clone patches with trap-revert).
- If a `/demo-up` rebuild fails (the M46 build pitfalls: a standalone app rebuild that skips apply-authn collapses
  auth → reachable≈7; --force-recreate without --no-deps wipes the seeded org) → user-blocker (a heavy demo
  re-up failure mid-iter), record + surface.

## Acceptable close-no-lift outcomes
If the perf-wall demo-patches land + the demo re-ups clean but a residual perf or manifest gap keeps `(failing)`
above 0 → close-fixed-partial: the perf-wall (the planned target) cleared for the sections it covers, the
remainder (e.g. the AI-readiness manifest assertion) routes forward. Falsification "the perf-patches cleared
N of 6; the residual M are a different cluster" is a valid documented outcome.
