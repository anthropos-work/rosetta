# M51 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## TOK-01: active-cycle signals-true additive-to-stories seed — 2026-06-30

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Build the AI-readiness showcase as a **3rd story** in the existing Stories & Heroes world,
seeded for an **ACTIVE cycle with real signals** (not a closed-cycle frozen snapshot), and drive it to the gate via
the coverage-protocol observe→fix→re-measure loop against demo-1 in place. Concretely, four work strands the first
batch of tiks will sequence:

1. **The 3rd story (YAML + org enablement).** Append a `stories[]` entry to
   `stack-seeding/presets/stories.seed.yaml`: org "AI Readiness" (size 200), a hero trio
   (manager `vantage: manager` + a thriving end-user pinned COMPLETED + a struggling/early end-user pinned
   STARTED), narrative + activity. Add a net-new **`organization_settings` `ai_readiness` gate-row writer**
   (a small `OrgSettingsSeeder` iterating `EffectiveStories()`, one row per org `setting='ai_readiness',
   is_enabled=true`) — nothing writes that table today. The 3rd org gets its distinct org-id for free via
   `StoryOrgID(story.ID)`.
2. **The AI-readiness config + cycle.** A net-new seeder writing the `ai_readiness_*` config per the 3rd org:
   `ai_readiness_cycles` ×1 `status='active'`; `ai_readiness_skills` ~5 core (weight 1.0) + a few enabling (0.5)
   with **real replayed-taxonomy node-ids** (via `resolveTaxonomyRefs`, never fabricated — the closure gate);
   `ai_readiness_sims` ×2 (`step_type` simulation+interview, `sim_ref` = a real Directus sim id pinned via the
   net-new sim-id pin mechanism); `ai_readiness_steps` ×3 optional (canonical default if absent).
3. **The 200-member funnel (signals-true).** Because the cycle is ACTIVE, the dashboard RECOMPUTES from signals
   (contract claim 5, verified GREEN) — so the seeder writes the **underlying signals**, not the live_snapshots
   cache: per ~160 "completed" members write ≥1 `user_skill_evidences` for a configured AI skill (Step 1, reuse
   the verified-skill chain / population evidence) + ended/scored jobsim sessions whose `sim_id ∈ ai_readiness_sims`
   for Steps 2/3 (needs the sim-id pin) + `ai_readiness_user_step_progress` ×3 `completed`. The COMPLETED hero gets
   all 3 (stage 3); the STARTED hero gets only the Step-1 signal + stage 1. `keepStartedMembers` requires a Step-1
   signal to keep a member in the aggregate — so every counted member needs ≥1 evidence.
4. **Cockpit wiring + coverage drive.** Set each hero's `jump_to` (manager → the `/enterprise/...` AI-readiness
   dashboard; employees → their onboarding element); add `DeepLinkCatalog` entries for proper labels. Then run the
   M42 manager-vantage semantic coverage gate on demo-1, triage failures via the fix-surface routing table,
   re-seed/re-replay/re-sweep until `(0,0)` frontier-exhausted.

**Rationale:** (a) Additive-to-stories is the lowest-risk, highest-reuse path — appending a `stories[]` entry yields
the org identity, roster, and cockpit menu for free, and the PersonaSeeder 7-table verified chain + JobsimSessions +
closure gate are reused as-is (iter-01 survey). (b) Active-cycle-signals-true is chosen over closed-cycle-snapshot-
direct because the gate's whole premise is a *live, believable, in-flight* assessment (a manager watching a funnel),
and the contract confirms an active cycle recomputes from signals — seeding the signals makes the dashboard render
authentically and survives a `RefreshLiveSnapshots`, whereas snapshot-direct reads as a *finished* assessment and is
the wrong demo affordance for "1 hero STARTED" (a started hero only exists mid-cycle). (c) Signals-true also reuses
the existing evidence/jobsim machinery rather than inventing a frozen-snapshot writer. The cost is the net-new
sim-id pin + the funnel seeder; accepted because it's bounded and the alternative (snapshot-direct) can't show an
in-progress hero.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** Gate metric = the M42 manager-vantage coverage `(failing-pages, escapes)` pair on the
3rd org, target `(0,0)` frontier-exhausted on a fresh demo-up, PLUS the gate's qualitative conditions (dashboard
ENABLED, ~80% all-3-complete, 1 hero STARTED + 1 COMPLETED). Starting value: the 3rd org does not exist, so no sweep
runs yet — build distance is the full 3rd-org seed. iter-02 lands the first slice + takes the baseline sweep once
the org renders.

**Next-tik direction:** iter-02 (first tik) — land strand 1: append the AI-Readiness 3rd story to
`stories.seed.yaml` + add the `OrgSettingsSeeder` (the `ai_readiness` gate row), re-seed demo-1, then take the
**baseline manager-vantage sweep** logged in as the new manager hero (expect the dashboard to gate-render — possibly
empty/funnel-less until the config+funnel strands land — establishing the baseline `(failing, escapes)`).

## USER-BLOCKER (iter-04, 2026-06-30): demo-1 rext consumption clone is hand-modified, blocking the perf-wall re-up

**Context:** iter-04 triaged the 6 GATED-sweep failures as the M46 base-Workforce org-scale PERF-WALL
(skeleton false-fails, data confirmed in the DB; the fix is the demo-UP path, not `stack-seeding`). The
routed fix is to re-pin demo-1's consumed rext tag to `fit-up-m50` (which wires the
`next-web-members-pagination` + `app-targetrole-authz-skip` + post-seed FK-index perf-patches into
`up-injected.sh`) and `/demo-down 1` + `/demo-up 1`.

**Blocker:** `git checkout fit-up-m50` in `stack-demo/rosetta-extensions` ABORTS — the consumption clone is
NOT a pristine tag checkout. It carries leftover hand-modifications (a partial M50 application, almost
certainly from the same concurrency incident that left iter-03 uncommitted): `up-injected.sh` modified +
differing from BOTH m49 and m50; `test_demopatch.py` modified == m50; and an UNTRACKED
`patches/next-web-public-website-url/next-web-public-website-url.yaml` (== m50) that blocks the checkout.

**Why user-blocker:** unblocking requires `git clean`/`rm` (the untracked file) + `git checkout --`/`git stash`
(the modified files) — all in the build-iter FORBIDDEN-OPS list. The user + orchestrator are the only allowed
deciders on this dirty consumption-clone state. Full detail + recommended resolution in iter-04/decisions.md.

**iter-04 left mid-Phase-C, NOT closed** (no fix landed, no `iter(M51/04):` commit). The untracked
`iter-04/` dir is left uncommitted by design (Phase 4 Step 0 budget/blocker-interrupted-iter rule).

**RESOLVED (run-2, 2026-06-30):** the orchestrator reset the consumption clone to a clean `fit-up-m50`. iter-04
RESUMED + ran to a **closed-no-lift**: demo-1 rebuilt at fit-up-m50 (all 3 perf demo-patches baked) + re-seeded
the AI-readiness showcase org + re-exported the 9-hero roster/cockpit; the GATED manager sweep HELD at
`(failing=6, escapes=0)` frontier-exhausted. The m50 perf-patches reduced the members-grid wall 76.4s→~11.6s but
the residual COLD query still exceeds the harness measurement budget → the 6 skeleton false-fails persist
(data-in-DB, slow-not-erroring). The hypothesis "m50 patches alone clear all 6" is FALSIFIED. The residual is
demo-local-addressable (a harness warm/poll deepening) → iter-05; the manifest AI-readiness assertion + cockpit
jump_to (TOK-01 strand-4) are mapped + also routed to iter-05. See iter-04/{progress,decisions}.md.

## USER-BLOCKER (iter-07, 2026-07-01): the closed-cycle strategy is DB-correct but the platform FE default doesn't route to the frozen path

**Context:** the user chose the M48-documented CLOSED-CYCLE alternative to the iter-06 perf wall (seed the
cycle closed + frozen per-member `ai_readiness_snapshots` so the dashboard reads pre-computed data instead of
live-recomputing + the per-skill translation N+1). iter-07 implemented it (config: active→closed; funnel: a
frozen snapshot per stage>=1 member, platform-model-scored; stackseed: ai_readiness_* in --reset + baked
--reload-sentinel) and re-seeded demo-1. **The DB is now the CORRECT showcase: cycle closed, 199 frozen
snapshots, 78.4% stage-3, Aria=stage3/champion, Ben=stage1/standby, Dana no snapshot.**

**Blocker (root-caused, zero platform edit):** the platform read path reaches the FAST frozen branch
(`app buildResponseFromSnapshots`) ONLY for a `?cycle=<closed-id>` request; the DEFAULT GET (nil CycleID) is
hardcoded to `buildLiveResponse` (`ai_readiness.go:301`). An authenticated network probe proved the demo FE
fires the dashboard's data GET **WITHOUT `?cycle=`** (the live path — it hangs, never completes) and **never
fires the `/cycles` list** that would supply `latestClosedCycle.id`. So the frozen data is present + fast-readable
but the default dashboard call never selects it → the GATED manager sweep HELD at (failing=5, escapes=0). The
2 AI-readiness sections stay skeleton; the 3 workforce-aggregate sections are the same iter-06 wall family.

**Why user-blocker:** every path to `(0,0)` needs a user/architectural decision the invariant forbids the
build-iter from taking:
  (a) **DISCLOSED residual** (the session-brief fallback) — the data is PROVEN correct in the DB; the section is
      slow-but-correct due to a platform FE/read-path ROUTING behavior, not a seed gap. Disclose as a
      presenter-note per the coverage-protocol's NARROW disclosed-allow → the gate reaches green-with-disclosure.
      **Needs the user's EXPLICIT sign-off** (it is NOT an editorial-citation auto-allow). The seeded closed-cycle
      data STAYS (honest + correct; the cycle picker / a `?cycle=` deep-link renders the fast frozen dashboard).
  (b) **ESCALATE a platform edit** (`unimplementable-without-platform-edit`, the milestone Re-scope trigger) — make
      next-web's default dashboard query pass the latest-closed cycle id, OR make `app`'s default GET prefer a
      closed cycle when no active cycle exists.
  (c) a NEW app read-path demo-patch (batch/relax the live translation N+1) — the `app-targetrole-authz-skip`
      precedent; a substantial new tooling investment; option B of the session brief, NOT chosen.

iter-07 surfaces the decision rather than picking one. Full evidence in iter-07/{progress,decisions}.md. The
closed-cycle seeder + --reload-sentinel are KEPT + committed (the DB showcase is correct + reusable regardless
of the chosen resolution). `fit-up-m51` is NOT tagged (gate not met).

## USER-BLOCKER (iter-08, 2026-07-01): the user-chosen DEEP-LINK strategy is FALSIFIED — the frozen READ is itself org-scale-slow

**Context:** the run-5 brief chose the zero-platform-edit **deep-link the demo entry** fix for the iter-07
residual: point Dana's cockpit `jump_to` + the coverage manifest at the AI-readiness dashboard with
`?cycle=<latest-closed-cycle-id>` so the FE fires the cycle-scoped GET, which the platform "serves FAST via
`buildResponseFromSnapshots` — reachable only when the request carries the closed cycle id." The brief mandated
**verifying the deep-link hits the frozen fast path FIRST (a cheap probe before a full sweep)**.

**Falsification (the cheap probe, before any cockpit/manifest edit or sweep):** an authenticated dual-endpoint
DIRECT probe (`stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts` — lift Dana's bearer, hit both
endpoints directly via `page.request`, bypassing the FE React Query gate) showed:
  - `GET /api/workforce/ai-readiness/cycles` → **200 in 40 ms** (fast — the FE gate is fine; returns the closed
    cycle `95d9fc3d-48c0-53ca-82ac-e10713100c97`).
  - `GET /api/workforce/ai-readiness?cycle=<closed>&includePeople=true` (the FROZEN path) → **NEVER COMPLETED**
    (180 s timeout), identical to the live-recompute default.

**Root cause (source read, zero platform edit):** `app buildResponseFromSnapshots` (`ai_readiness.go:512`)
reads the frozen SCORES fast, but then calls **`loadMembers(orgID, "")`** — a full UNBOUNDED org-member
hydration (`hydrateMembers` with `memberIDs=nil, userIDs=nil` → whole-org tag/skill/sim aggregation over ~200
members) to re-join CURRENT tags/name/role onto each snapshot. At 200 members that member-load is the SAME
org-scale wall as the live path (the `ai_readiness_refresh` worker logs `context deadline exceeded`). Crucially
it is NOT the demo-patchable per-object targetRole Sentinel RPC — `queryBaseMembers` reads `jobRole` from a SQL
column, so the applied `app-targetrole-authz-skip` patch does nothing for it. **"Frozen" froze the SCORES, not
the RESPONSE** — the response re-joins live members → re-incurs the wall.

**Why this is a STRONGER falsification than iter-07:** iter-07 concluded only "the DEFAULT FE GET omits
`?cycle=`, so the frozen branch is never SELECTED." iter-08 shows the frozen branch, even when selected by a
direct `?cycle=` request, is ITSELF org-scale-slow — so the deep-link (the user's chosen zero-edit fix) cannot
clear the wall EVEN IN PRINCIPLE. No cockpit/manifest deep-link edit was made (it would ship an inert `?cycle=`
that still hangs); no GATED sweep was run (it would only reconfirm `failingSections=5`).

**Why user-blocker (the decision the user must make):** every remaining path to `(0,0)`:
  (a) **DISCLOSED residual** (the run-5 brief's stated fallback): the data is PROVEN correct in the DB — the 2
      AI-readiness sections + the 3 workforce aggregates (same `loadMembers` family) are slow-but-correct due
      to a platform read-path perf wall, not a seed gap. Disclose as a presenter-note per the coverage-protocol's
      NARROW disclosed-allow → green-with-disclosure. Per the coverage-protocol + iter-07 this needs the user's
      **EXPLICIT sign-off** (NOT an auto-allow); self-granting it would be loosening the gate to force a pass
      (the invariant's hard line), so the build-iter does NOT self-grant it. The seeded closed-cycle data STAYS
      (honest + correct).
  (b) **ESCALATE a platform edit** (`unimplementable-without-platform-edit`, the milestone Re-scope trigger):
      bound `loadMembers` in the snapshot path (pass the ~199 snapshot user-ids as `restrictUserIDs` instead of
      a whole-org hydration) OR add the deferred `frozen_tags jsonb` column (M314b) so the snapshot read needn't
      re-join live members. The invariant forbids doing it here.
  (c) a NEW app read-path demo-patch bounding `loadMembers` in the snapshot path (the `app-targetrole-authz-skip`
      precedent) — a substantial new tooling investment NOT in the run-5 brief's chosen scope.

iter-08 surfaces the decision rather than picking one. The closed-cycle seeder + the reusable diagnostic probe
are KEPT + committed (correct + reusable regardless of the resolution). Protocol lesson added
(coverage-protocol.md, M51 iter-08: frozen SCORES ≠ frozen RESPONSE; measure the fast branch END-TO-END).
`fit-up-m51` is NOT tagged (gate not met). Full evidence in iter-08/{progress,decisions}.md.

## TOK-02: app read-path demo-patch — bound loadMembers in the snapshot path — 2026-07-01

**Tok type:** triggered (3-no-prog streak: iter-06/07/08 all held failingSections=5) — but the strategy
revision was **authored + reviewed OUT-OF-BAND by the user** across the iter-06/07/08 user-blocker exits.
The user reviewed the three surfaced falsifications and CHOSE option (c) — a new app read-path demo-patch —
for run 6, directing "continue with iter-09 (app injection patch)". So this TOK records the user-ratified
pivot; the mandated triggered-tok user-review already happened (the run-6 brief IS its output). Per the
user's explicit direction to execute the app-injection tik THIS run, iter-09 proceeds as a tik under TOK-02
rather than exiting for a review that has already occurred.

**Prior strategy:** TOK-01 (active-cycle signals-true) + its iter-07 closed-cycle amendment + the iter-08
deep-link. The coverage-drive strand assumed the AI-readiness dashboard read could be made fast either by
seeding signals (active-cycle recompute), by freezing scores (closed-cycle snapshots), or by deep-linking the
FE to the frozen branch.

**Why it stopped working (the 3 no-prog tiks):** iter-06 falsified the active-cycle path (live-recompute +
per-skill translation N+1 never completes in-budget). iter-07 falsified "seed a closed cycle" (the DEFAULT FE
GET omits `?cycle=`, so the frozen branch is never SELECTED). iter-08 falsified the deep-link
(`buildResponseFromSnapshots` itself calls `loadMembers(orgID, "")` — a full UNBOUNDED org-member hydration —
so even a direct `?cycle=<closed>` GET times out at 180s: "frozen" froze the SCORES, not the RESPONSE). The
common wall across all three: the ~200-member unbounded member hydration in the read path.

**New strategy:** author a NEW app injection demo-patch — `app-aireadiness-snapshot-loadmembers` — following
the `app-targetrole-authz-skip` precedent (an in-inject-loop source swap of the demo's build-scratch app clone
via a rext-owned `apply-*.sh` helper mirroring a pinned anchor→replacement manifest). The swap: in
`buildResponseFromSnapshots` (`internal/workforce/ai_readiness.go`), replace the unbounded
`m.loadMembers(ctx, orgID, "")` with a BOUNDED `m.loadMembersByUserIDs(ctx, orgID, "", snapUserIDs)` where
`snapUserIDs` are the ~199 snapshot user-ids — hydrating ONLY the frozen members via the indexed
`memberships."user" = ANY($2::uuid[])` filter (`queryBaseMembers` byUser=true) instead of a whole-org scan.

**Why it is a PURE perf optimization (data-identical):** in `buildResponseFromSnapshots`, the loaded members
are used ONLY to build `membersByUserID` (keyed by `Member.UserID`), which is looked up ONLY by each
snapshot's `s.UserID`. Members with no matching snapshot were loaded-but-never-used. `loadMembersByUserIDs`
(the existing bounded sibling used by the live scoring path) restricts `queryBaseMembers` to those user-ids
and hydrates tags/skills/sims off the loaded rows' membership ids — so the resulting map carries the same
entries for every `s.UserID` that resolves. Snapshot users with no active membership fall to the identical
orphan branch (map miss) in both forms. The `AIReadinessResponse` (Org aggregate, ByTeam, People) is
byte-identical; only the member-load cost drops from whole-org to the snapshot set. Scoped to the snapshot/
closed-cycle read path ONLY — the live/active-cycle behavior for other orgs is untouched.

**Strategy class:** more-granular — a per-function bounded-query swap in the one hot path, retrying the
closed-cycle strategy with the NEW evidence (iter-08's `loadMembers` root cause) that identifies the exact
unbounded call to bound. Reuses the established app-demo-patch tooling (precedent + helper + manifest schema).

**Distance-to-gate context:** gate = (failingSections, escapes) manager-vantage on Northwind = (0,0)
frontier-exhausted on a fresh demo-up, + the qualitative conditions (already DB-correct: closed cycle, 199
frozen snapshots, 78.4% stage-3, Aria/Ben/Dana). Current: (5, 0). The 5 = 2 AI-readiness + 3
workforce-aggregate — all in the same `loadMembers`/member-hydration family, so the one bounded-query patch is
hypothesized to clear all 5. Expected lift: 5 → 0.

**Next-tik direction:** iter-09 (tik under TOK-02) — author + commit the patch (manifest + `apply-*.sh` helper
+ `up-injected.sh` wiring) in the authoring copy; re-pin the consumption clone to the new sha; rebuild demo-1's
app image (through the inject loop) + re-seed; verify with the cheap dual-endpoint probe (frozen GET now
completes fast + returns the correct funnel) BEFORE the gated sweep; then run the gated manager sweep and drive
failingSections 5 → 0.

## D-CLOSE-1 (close, 2026-07-01): academy F6 repeat-defer fate — LAND-NEXT → M53

**Context:** the M51 close deferral re-audit (Phase 1b, `audit-deferrals/deferral-audit-2026-07-01-m51-close.md`)
flagged **RED**: the ant-academy field-review gap **F6** (0 course-content + no hero academy menu-link +
anonymous session when the academy is reached directly) is a **REPEAT deferral** — first raised at M50 close
(D-CLOSE-3, Fate-3 → M51), then not executed in any M51 iter (M51's 9 iters went entirely to the AI-readiness
dashboard perf saga: active-cycle → closed-cycle → deep-link → app read-path demo-patch). The M51 manager
coverage gate was MET without it. A repeat deferral requires an explicit per-item user fate decision before the
milestone can close (the RED can only clear on that decision).

**Options (from the audit):**
- (a) KEEP-DEFERRED-WITH-SIGNOFF → a future release (v2.0 **M207 Academy coverage**) — the auditor's
  recommendation (the academy is a coherent separate Vercel deployment already owned by a named v2 milestone).
- (b) LAND-NEXT via a Fate-3 override of a sibling milestone's scope (M52 `Out: new seeding` / M53
  `No new feature code (acceptance only)`).
- (c) LAND-NOW in M51 close — rejected by the auditor as new emergent scope disjoint from M51's AI-readiness work.
- (d) DROP.

**Choice: (b) LAND-NEXT → M53** (Fate-3 override of M53's `No new feature code` scope-guard). **User-decided.**

**Why:** M53 already **destroys the live demo and cold-rebuilds it from scratch** — that is the single natural
place to *seed + verify* academy content on a clean build, rather than patch it onto the live demo late. The F6
items are a small ant-academy **seeding/content + wiring** surface (course content present; a hero academy
menu-link; a non-anonymous academy session), not new platform **feature code** — so the Fate-3 override of M53's
acceptance-only guard is honest (it adds a seed+assert surface, not a feature). The academy's **AI chat stays
documented-as-absent** per the AI-keys policy (the `/api/ai/chat` route needs keys the demo doesn't provision —
no assertion on it). This closes the M50→M51→M53 handoff chain: the item lands on the cold-rebuild acceptance,
and a failed academy assertion routes back to its owner (never expands M53).

**Applied:** M53 `overview.md` Scope `In:` now carries the academy-F6 item (course content + hero menu-link +
non-anonymous session; AI chat documented-as-absent) with the handoff-chain rationale. The Phase-1b audit RED
is thereby **cleared** (the repeat-defer has an explicit LAND-NEXT fate). See
`m53-cold-rebuild-acceptance/overview.md` + the updated `audit-deferrals/deferral-audit-2026-07-01-m51-close.md`.
