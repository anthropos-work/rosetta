**Type:** tik — under TOK-01 (active-cycle signals-true additive-to-stories seed), coverage-protocol.md
Phase A–E. Strands 2+3 of TOK-01's 4-strand plan, ABSORBED into one iter (D1): the `ai_readiness_*`
CONFIG seeder (cycle + steps + skills + sims) + the 200-member signals-true FUNNEL seeder
(Step-1 evidence + Step-2/3 sessions + step-progress, ~80% all-3-complete, heroes pinned).

# iter-03 — tik (strands 2+3: the AI-readiness CONFIG + the 200-member FUNNEL)

## What happened (Phase A–E)

This iter RESUMED a now-stopped parallel agent's uncommitted strand-2+3 work after a concurrency incident
(the documented resume start). The work was ADOPTED (not redone), verified-green, the demo re-seeded cleanly
to a known reproducible state, and the GATED sweep run.

- **Phase A (measure):** inherited the iter-02 baseline `(failingSections=6, escapes=0)` as the pre-iter
  metric (Step-0 re-survey confirmed the strand-2 config target untouched + meaningful; no substitution).
- **Phase C (fix) — the adopted build, verified + reconciled:**
  1. **`AIReadinessConfigSeeder`** (`seeders/ai_readiness_config.go`, surface `ai-readiness-config`,
     depends-on `org`+`taxonomy`+`content`, PerStackIsolated): per `narrative: ai-readiness` story writes
     `ai_readiness_cycles` ×1 (`status='active'`, window now−activity→now+1mo — D4), `ai_readiness_steps` ×3
     (skill_mapping/simulation/interview pos 0/1/2), `ai_readiness_skills` 5 core (w1.0) + 3 enabling (w0.5)
     via the DEDICATED DB-side AI-name-pattern reader `readAIReadinessSkillPool` (D2 — NOT `filterAISkills(flat)`,
     which starved to 1 skill on the alphabetical flat-pool cap), `ai_readiness_sims` ×2 (real replayed
     Directus sim ids via `resolveContentRefs().sims`). Closure-gate-safe: real public node-ids only, honest
     degradation (no taxonomy → 0 skills, never fabricated).
  2. **`AIReadinessFunnelSeeder`** (`seeders/ai_readiness_funnel.go`, surface `ai-readiness-funnel`, depends-on
     `users`+`ai-readiness-config`+`taxonomy`+`content`+`personas`+`population-evidence`): per member writes the
     stage-appropriate signals the ACTIVE-cycle dashboard recomputes from — Step-1 `user_skill_evidences`
     (UPSERT, node_id from the SAME configured pool so it's a configured `ai_readiness_skill` by construction),
     Step-2/3 ended/scored `jobsimulation.sessions` (sim_id = the org's `ai_readiness_sims` refs, G14-valid
     enums), + `ai_readiness_user_step_progresses` (status='completed'). ~80% reach stage-3, residual splits
     stage-2/stage-1 (D5). Heroes override: thriving→stage3 COMPLETED, struggling→stage1 STARTED, manager→excluded.
  3. **Registered** both in `cmd/stackseed/main.go` (config at level-1 after org; funnel at level-4 after
     population-evidence). 7 net-new unit tests (4 config + 3 funnel) + the funnel-mapped membership-skills test.
  4. **build/vet/gofmt clean**; full `seeders`+`blueprint`+`cmd` suites GREEN on a `-count=1` (no-cache) run
     (the 7 AI-readiness tests + `TestMembershipSkillsSeeder_FunnelMappedOutnumbersHero` +
     `TestOrgSettingsSeeder_WritesAIReadinessGateForOptInStoryOnly` all PASS).
  5. **Re-seeded demo-1 CLEANLY** (authoring-copy `stackseed` against demo-1 :15432; the post-concurrency
     known-state re-seed): isolation guard CLEAN (47536 rows, prod=false, no shared/external writes). Verified
     in the live DB for Northwind (`d7bb7482-…`): cycle=1 (active, 2026-01-31→2026-07-30), steps=3, skills=8
     (5×w1.0 real AI node-ids `K-A*` + 3×w0.5), sims=2, org_setting `ai_readiness`=enabled. FUNNEL: 199 members,
     **156 stage-3 (78.4% all-3-complete ≈ the gate's ~80%)**, 21 stage-2, 22 stage-1 (a believable in-flight
     drop-off). HEROES PINNED EXACTLY: Aria(thriving)→3 steps COMPLETED, Ben(struggling)→1 step STARTED,
     Dana(manager)→0 (excluded). The roster+cockpit were already current from iter-02 (the funnel re-seed
     changes only signal rows, not identities; fake-fapi/bapi carry the 9-hero roster incl. `dana-manager`).
- **Phase D (re-sweep) — the GATED manager-vantage sweep:** seat `dana-manager`, expected-org "Northwind
  Aviation", gate ENFORCED (NO `COVERAGE_NO_GATE`), over a **frontier-EXHAUSTED** crawl (49 reachable, not
  capped): **`(failingSections=6, escapes=0)`, personaFailures=0, notReached=0 — GATE: NOT MET.** Metric
  UNCHANGED from the `(6,0)` baseline.
- **Phase B (triage) of the 6 — the KEY learning:** the 6 failing sections are NOT the AI-readiness sections
  and NOT a seed gap. They are the **base Workforce-Intelligence dashboard grids** at the new 200-member org:
  `/enterprise/workforce` verification-funnel + talent-languages, `/enterprise/members` members-roster +
  members-location-values, `/enterprise/assignments` assign-roster, `/enterprise/activity-dashboard`
  activity-table — every one `kind:empty` "genuinely below bar (re-asserted 6×)". The DB has the data
  (200 memberships, 1330 membership_skills, EU member locations London/Lisbon/Barcelona…, 782 jobsim sessions).
  The screenshot is the tell: **real chrome (Northwind org name + Dana avatar + correct table headers + "0/∞"
  badge) over SKELETON rows**, with GraphQL ERROR rows at **11.4 s / 4.2 s latency, status 200** in the router
  log. This is the textbook **M46 org-scale PERF-WALL / skeleton-frame false-fail** — a measurement/perf gap,
  not a content gap, and demo-1 (consuming `fit-up-m49`) was brought up WITHOUT the M46/M50 perf demo-patches
  (next-web `limit 1000→30` pagination + the `app` targetRole authz-skip). The fix surface is the demo-UP path
  (re-up with the M46 perf patches), not `stack-seeding`.
- **Second learning:** the manager manifest does NOT yet declare an AI-readiness section — so the seeded
  config+funnel renders on the live dashboard but the GATE doesn't yet ASSERT it. `/enterprise/workforce`
  asserts verification-funnel/stat-cards/tabs/talent-languages/talent-certifications; no `ai-readiness-*`
  descriptor exists. Even a perfectly-seeded funnel cannot register as gate progress until iter-04 adds the
  manager-manifest AI-readiness descriptor (a harness/manifest task — likely a tooling-iter).

## Close — 2026-06-30

**Outcome:** Strands 2+3 (the `ai_readiness_*` config + the 200-member signals-true funnel) LANDED clean +
verified in demo-1 (active cycle, 8 real-node-id AI skills, 2 sims, 199-member funnel at 78.4% all-3-complete,
heroes pinned Aria→stage3/Ben→stage1/Dana→excluded). The gated manager-vantage sweep held at `(failing=6,
escapes=0)` — UNCHANGED — because the residual 6 are an UNRELATED cluster (base-Workforce org-scale perf-wall,
data-in-DB skeleton false-fails), and the gate does not yet assert the AI-readiness funnel section.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (this iter is the 1st tik of the session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (absorb strand-3 into iter-03), D2 (dedicated AI-skill-pool reader), D3 (real plural table name `ai_readiness_user_step_progresses`), D4 (active-cycle in-flight window), D5 (~80%-all-3 believable drop-off) — iter-03/decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:**
  - **iter-04 (the perf-wall cluster — the highest-leverage next target):** clear the 6 base-Workforce
    skeleton false-fails by bringing demo-1 up WITH the M46/M50 perf demo-patches (next-web `InsightsContext`
    `limit 1000→30` pagination + the `app` `roles.go` targetRole read-gate authz-skip + post-seed FK indexes),
    OR re-pin demo-1's consumed rext tag to one carrying them. This is a demo-UP / consumed-tag concern, not a
    `stack-seeding` fix. Confirm by the screenshot flipping skeleton→rows + the GraphQL latency dropping <5 s.
  - **iter-04+ (TOK-01 strand-4 + a manifest tooling-iter):** add a manager-manifest AI-readiness section
    descriptor (`lib/coverage-manifest.ts`) so the gate ASSERTS the seeded funnel (the funnel tab / the
    AI-readiness widget on `/enterprise/workforce`), + cockpit `jump_to` re-point (Dana → the AI-readiness
    manager route) + `DeepLinkCatalog` labels.
  - **Doc fix (route to milestone close or a doc tik):** `corpus/services/ai-readiness.md` writes the SINGULAR
    `ai_readiness_user_step_progress`; the live schema is the ent-PLURAL `ai_readiness_user_step_progresses`
    (D3). One-line contract-doc name fix.
  - **M50 Fate-3 (ant-academy course-content + menu-link + non-anonymous session):** still in M51 candidate
    scope (overview), untouched this iter; route forward.
**Lessons:** (1) "Empty section ≠ seed gap" — DIAGNOSE by SCREENSHOT + DB + GraphQL latency before assuming
`stack-seeding`. Six sections read `empty` with all data present; the screenshot (real chrome + skeleton rows
+ "0/∞") + 11.4 s federated latency proved it's the M46 perf-wall, a demo-UP/consumed-tag fix, NOT a seed fix.
A signals-true funnel can be perfectly seeded and STILL not move the gate if the residual failures are a
different cluster. (2) Seeding the data is necessary but not sufficient for gate credit — the manager MANIFEST
must also DECLARE an AI-readiness section for the gate to assert it; that's an iter-04 harness/manifest task.
(3) The additive-to-stories + signals-true strategy (TOK-01) is VALIDATED at the data layer: the config+funnel
seeded deterministically, idempotently, closure-clean, with the exact gate-qualitative shape (ENABLED, ~80%
all-3, 1 STARTED + 1 COMPLETED) — the remaining distance is perf-wall + manifest, both iter-04 targets.
