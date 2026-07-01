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
