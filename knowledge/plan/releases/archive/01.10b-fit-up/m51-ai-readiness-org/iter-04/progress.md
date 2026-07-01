**Type:** tik — under TOK-01 (active-cycle signals-true), coverage-protocol.md Phase A–E. The coverage-drive
strand: the iter-03-routed PERF-WALL cluster. RESUMED after the run-1 user-blocker (the dirty consumption
clone) was CLEARED by the orchestrator (clean `fit-up-m50` checkout).

# iter-04 — tik (the base-Workforce PERF-WALL: re-up demo-1 with the M46/M50 perf demo-patches)

## What happened (Phase A–E)

- **Phase A (measure):** inherited the iter-03 GATED `(failingSections=6, escapes=0)` as the pre-iter metric.
- **Phase B (triage):** done in the iter-04 overview Step-0 — all 6 failing sections are the M46 base-Workforce
  org-scale PERF-WALL (skeleton false-fails, data-in-DB), routed to the demo-UP fix surface.
- **Phase C (fix) — the demo re-up at the perf-patched tag:**
  1. **`/demo-down 1 --purge`** then **`/demo-up 1`** (consumed rext `fit-up-m50`). The build applied ALL THREE
     M46/M50 perf demo-patches (verified in the build log): `next-web-members-pagination` (InsightsContext
     `limit 1000→30`), `app-targetrole-authz-skip` (the per-member targetRole Sentinel-RPC read-gate dropped —
     "app: applied target-role authz-skip" logged during the injected-app build), and the post-seed FK indexes
     (`membershipskill_membership` + `membershiptag_membership` — "FK indexes ensured" logged post-seed +
     verified present in `pg_indexes`). demo-up EXIT=0, autoverify OK.
  2. **Re-seeded the 3rd org from the AUTHORING copy** (the consumed m50 `stories.seed.yaml` has only the 2
     original orgs; Northwind + the AI-readiness seeders are authoring-only @ `45a20c0`): built `stackseed` from
     the authoring `stack-seeding`, ran `--stack demo-1 --seed presets/stories.seed.yaml`. Clean
     (isolation: prod=false, no shared/external writes; ai-readiness-config 14 + org-settings 1 +
     ai-readiness-funnel 1064 rows). Verified in the live DB: Northwind 200 members; org_setting `ai_readiness`
     enabled; ACTIVE cycle (2026-01-31→2026-07-30); 3 steps / 8 skills / 2 sims; **funnel 156 stage-3 (78.4%
     all-3) / 21 stage-2 / 22 stage-1**; HEROES PINNED: Aria Holt→3 COMPLETED, Ben Castellano→1 STARTED, Dana
     Whitlock→0 (excluded).
  3. **Re-exported the 9-hero roster + cockpit** from the authoring copy (the demo-up exported the 6-identity
     2-org roster BEFORE the re-seed; the sweep logs in as `dana-manager`, a Northwind hero ABSENT from the
     6-identity roster) → `fake-fapi-roster.json` (9 identities incl `aria-completed`/`ben-started`/
     `dana-manager`) + `cockpit-manifest.json` (9 heroes); restarted `demo-1-fake-fapi-1` + `-fake-bapi-1`
     (fapi log: "loaded 9-identity roster").
- **Phase D (re-sweep) — the GATED manager-vantage sweep:** seat `dana-manager`, expected-org "Northwind
  Aviation", gate ENFORCED, frontier-EXHAUSTED (49 reachable, not capped):
  **`(failingSections=6, escapes=0)`, personaFailures=0, notReached=0 — GATE: NOT MET.** Metric UNCHANGED from
  the `(6,0)` baseline.
- **Phase B′ (re-triage of the held-6 — the KEY FALSIFICATION):** the SAME 6 base-Workforce sections still
  read `kind:empty` "re-asserted 6× over the heavy-grid budget — genuinely below bar"
  (`/enterprise/workforce` verification-funnel + talent-languages, `/enterprise/members` members-roster +
  members-location-values, `/enterprise/assignments` assign-roster, `/enterprise/activity-dashboard`
  activity-table). DIAGNOSED per the M46 protocol (DOM + DB + GraphQL-latency):
  - The **DB has the data** (200 memberships, 200 with skills, 1330 membership_skills, the EU locations).
  - The **screenshot is the tell**: real chrome (Northwind name + "NA" logo + Dana avatar + the
    Member/Email/Role/Job Role/Location headers + "0/∞" badge) over SKELETON rows — the textbook M46
    skeleton-frame false-fail.
  - The **GraphQL latency** shows the m50 perf-patches WORKED PARTIALLY: the wall dropped from a one-time
    **76.4s** (during the set-dress seed) to mostly-fast (0.1–0.2s) — BUT a residual **cold members query
    still hits ~11.6s** (logged ERROR-level by the router for >10s, `status:200` — a SLOW success, NOT a
    federation error; no subgraph errors in the log). At the 200-member cold-first-visit this 11.6s + the
    React grid render exceeds the harness's warm (≤4s networkidle ceiling) + bounded re-assert poll
    (`RE_ASSERT_MAX_TRIES=6`), so the authoritative visit captures a skeleton frame.

## Close — 2026-06-30

**Outcome:** The planned re-up LANDED FULLY (demo-1 rebuilt at `fit-up-m50` with all 3 perf demo-patches baked
+ FK indexes + clean re-seed of the AI-readiness showcase org + 9-hero roster/cockpit). The targeted PERF-WALL
cluster did NOT clear: the gated manager sweep held at `(failingSections=6, escapes=0)` UNCHANGED. The iter
HYPOTHESIS — "the m50 perf-patches ALONE clear all 6 base-Workforce skeleton false-fails" — is **FALSIFIED**:
the patches reduced the wall substantially (76.4s → ~11.6s) but a residual cold members-grid query at 200-member
scale still exceeds the harness measurement budget → the 6 skeleton false-fails persist. Data confirmed in the
DB; the screenshot confirms real-chrome-over-skeleton; the latency confirms slow-not-erroring.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is the 1st tik of run-2; the prior tik [iter-03]
made progress [config+funnel landed], so the no-prog streak is 1, not 3) — (3) re-scope: n (the residual is
demo-local-addressable — a deeper perf reduction or a harness warm/poll deepening, NOT a platform edit) —
(4) user-blocker: n (the run-1 blocker was cleared; no new blocker) — (5) cap-reached: n (1 tik this run) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (re-up at fit-up-m50 not a fresh authoring tag — the m50 patches are the planned fix surface),
D2 (re-seed + re-export roster/cockpit from the AUTHORING copy — the consumed m50 lacks Northwind + the
AI-readiness seeders), D3 (closed-no-lift: the planned re-up landed but the targeted cluster didn't clear; the
hypothesis is falsified with documented evidence) — iter-04/decisions.md
**Side-deliverables (if any):** none (this iter modified NO rext production code — it re-ran the demo lifecycle +
re-seeded from already-committed authoring seeders).
**Routes carried forward:**
  - **iter-05 (the residual PERF-WALL — the highest-leverage next target):** drive the residual cold ~11.6s
    members-grid query under the harness measurement budget. Demo-local levers (zero platform edit): (a) a
    HARNESS warm/poll deepening — the manager warm step bails at the 4s networkidle ceiling so the heavy grid is
    cold at the authoritative visit; deepen the manager-grid warm to wait for real rows (content-presence) before
    the authoritative sweep, and/or widen the bounded re-assert budget for the org-scale grids (a tooling-iter
    candidate); (b) consider whether a further demo-local query optimization is available. NOT a re-scope (the
    M46 close already proved this class is demo-local).
  - **iter-05+ (TOK-01 strand-4 + the manifest AI-readiness assertion — MAPPED this iter):** add a
    manager-manifest AI-readiness `PageDescriptor` for **`/enterprise/workforce/ai-readiness`** (the manager
    dashboard route — `AIReadinessContainer` → `AIReadinessView`, NO flag gate, reads
    `GET /api/workforce/ai-readiness`) to `MANAGER_PAGES` + add the route to `MANAGER_MANIFEST.seedPaths` (the
    route is NOT nav-linked from any manager surface — confirmed by an exhaustive grep — so a seed-path primes the
    BFS to visit it; a seed that renders is scored, not `notReached`). Section asserts: page title "AI Readiness";
    "Overall org readiness" + the `{score}/100` + "Members"; "Stage breakdown" + the stage labels ("Stage 1 —
    skills only", "Stage 2 — + AI Path", "Stage 3 — + AI Interview") + "Steps completion" (the en strings
    resolved from `configs/i18n/messages/en/enterprise.json` `workforceTabs.aiReadinessPage`). + cockpit
    `jump_to` re-point: Dana `/enterprise/workforce` → `/enterprise/workforce/ai-readiness` in
    `stories.seed.yaml` + the `DeepLinkCatalog` label in `cockpit.go`.
  - **Doc fix (route to milestone close / a doc tik):** `corpus/services/ai-readiness.md` line 61 writes the
    SINGULAR `ai_readiness_user_step_progress`; the live ent-plural table is `ai_readiness_user_step_progresses`
    (the D3 item). One-line contract-doc name fix.
  - **M50 Fate-3 (ant-academy course-content + menu-link + non-anonymous session):** still in M51 candidate
    scope; NOT on the manager coverage-gate path; untouched this iter; route forward.
**Lessons:** (1) **A perf-patch can PARTIALLY clear a wall and still hold the metric** — the m50 patches dropped
the members query 76s→11.6s (a real, large reduction) yet the gate is unmoved because the residual still exceeds
the HARNESS measurement budget. "The patch applied + latency dropped" is necessary but not sufficient for gate
credit; only a sub-budget query clears the skeleton false-fail. (2) **The warm step's 4s networkidle ceiling
under-warms the heaviest grid** — a query that takes 11.6s isn't warmed by a ≤4s warm, so the authoritative visit
pays the cold cost; the warm must wait for real ROWS on the org-scale grids, not just networkidle (an iter-05
tooling lever). (3) **A re-up that re-exports the roster MUST re-export from the SOURCE-OF-TRUTH stories** — the
demo-up exported the consumed-tag's 2-org roster; the 3rd-org manager hero (`dana-manager`, whom the sweep logs
in as) only exists in the authoring stories, so the roster + cockpit had to be re-exported from the authoring
copy + the fapi/bapi restarted, or the sweep can't log in. (4) **DIAGNOSE slow-vs-error before judging the fix
surface** — the 11.6s request logs ERROR-level (>10s) but returns 200 with no subgraph errors → it's SLOW, not a
data/federation error → the fix surface is perf (harness warm/poll or query opt), not serve-grant/seed.
