# iter-25 — decisions

## D1 — TRUE root cause of the 4 ai-readiness Playthrough failures: a STALE app image (schema drift), NOT `live_snapshots=0`
iter-23 routed this forward on a **correlation** (`ai_readiness_live_snapshots=0`). Live diagnosis on the
recovered billion demo REFUTES it and finds the real cause:

- The manager dashboard's live path (`buildLiveResponse` → `computeOrgBreakdowns`) does **NOT** read
  `ai_readiness_live_snapshots` — that table feeds only the Talk-to-Data askengine + the prune step. So its
  emptiness is irrelevant to the dashboard. (Verified in both stack-dev/app AND billion's stack-demo/app.)
- The pt-world Org C (Vertex Logistics) seed is **COMPLETE**: ai_readiness enabled=t, an **active** cycle
  ("AI Readiness Diagnostic — Q3", 2026-06-08→2026-08-22, in-window) + a closed cycle, 105 completed
  step-progresses, 39 frozen snapshots, **1 interview_aggregated_report** (the manager UC2 findings source).
- **THE TELL — billion backend logs:** `ERROR: column ai_readiness_cycles.launched_by does not exist
  (SQLSTATE 42703)` on `GET /api/workforce/ai-readiness/cycles` (→ 500) and `aiReadinessUserPlanProgress
  … load active cycle`. The cycles endpoint 500s → the frontend cycle picker can't resolve → "No cycles yet"
  zero-state → all 4 specs fail (member surfaces degrade to the archived rail-card; manager loses orgScore/
  matrix/cyclePill + the interview findings).
- **Root cause = version skew.** The running `demo-1-app:injected` binary (`/app/ant-backend`) CONTAINS
  `launched_by` (5 strings incl. `json:"launched_by,omitempty"`) — but the app clone is **pinned at v1.341.0**
  (clones.lock: pinned-tag, behind 145), whose source has **NO** `launched_by` anywhere (grep of .go +
  migrations + Ent schema = 0), and the DB (migrated per v1.341.0) has no `launched_by` column. So the cached
  image is an OLDER build (a prior release where `launched_by` existed) that up-injected **reused** instead of
  rebuilding from the re-pinned v1.341.0 clone.

## D2 — WHY it wasn't rebuilt (the provisioning defect) + why the gate's own cold-reset would still hit it
up-injected.sh:440 — *"per-demo image already exists, skip the rebuild."* The **backend `:injected` Go images
(app/cms/jobsim/skillpath) have NO source/clone-ref fingerprint cache-validator** (the frontend images DO —
the M220 patch-set-fingerprint LABEL forces a rebuild on drift). So re-pinning the app clone to v1.341.0 did
NOT invalidate the cached `demo-1-app:injected`. And `rosetta-demo down --purge` wipes only the DATA dir
(fresh initdb→migrate), **not images** — so even a proper cold reset-to-seed (`down --purge → up-injected`)
reuses the stale image and ai-readiness stays broken. This is a genuine gate-blocking v2.6-class provisioning
defect ("still not all gets built + provisioned as expected"), not an artifact of the warm recovery.

## D3 — the fix: operational rebuild now (M244 PROVES), durable cache-validator routed forward
Per M244's established pattern (iter-24 resolved 7 dev-build env walls operationally + routed the dev-stack
hardening backlog forward; TOK-03: "M244 PROVES, it does not re-build the tooling"): **operational fix** =
`docker rmi demo-1-app:injected` → force up-injected to rebuild the app image from the pinned v1.341.0 clone
(recompiles → binary drops `launched_by` → matches the migrated schema) → ai-readiness renders → the 4
Playthroughs land. 0 platform edits (pinned source unchanged; only the stale image is rebuilt). **Routed
forward (durable tooling fix):** add a clone-sha (+ patch-set) fingerprint LABEL cache-validator to the backend
`:injected` image build in up-injected.sh (mirror the frontend M220 fingerprint) so a clone re-pin forces a
rebuild and a cold reset-to-seed self-heals — the "make it not recur" fix. Named handler: FIND-M244-backend-image-fingerprint.

## D5 — the fix WORKED (wholesale zero-state gone) but revealed 3 distinct sub-failures → gate c 12/16 → 13/16
After the rebuild (app @ v1.341.0, binary `launched_by`=0, cycles endpoint 200, 0 launched_by 500s), the 4
ai-readiness specs re-ran cold on pt-world from the peer:
- **`ai-readiness.member-funnel.UC1` (member-done) — PASS ✅** (was the zero-state fail). The active cycle now
  resolves → the COMPLETED hero renders her score + all 3 step recaps. This PROVES the launched_by fix.
- **3 STILL FAIL — but NOT the zero-state**; specific richer surfaces:
  - `manager-dashboard.UC1`: score/cycle/matrix/championArchetype all render; the FIRST failing assertion is
    `byTeam()` — "AI Readiness by Tag" per-team breakdown absent. (`memberships.tags` column does NOT exist at
    v1.341.0 — the team-grouping source differs; likely a seed-team gap or a v1.341.0 render path.)
  - `manager-dashboard.UC2` (how-we-measure): the interview-findings panel renders only **24 chars** (want
    >900). The `jobsimulation.interview_aggregated_reports` row IS seeded (session_count 31, report 8396 chars),
    so this is a WIRING gap (sim_id/query mismatch or v1.341.0 read path), not a data gap.
  - `member-funnel.UC2` (member-progress): the cycle DEADLINE text is absent (active cycle end_date 2026-08-22
    exists; member-done resolves the cycle, so it's a deadline-render/started-hero gap).
- **Disposition:** these are DISTINCT deterministic render/wiring gaps at the demo's pinned v1.341.0, unmasked
  by the launched_by fix — NOT flakiness, NOT the zero-state, NOT a platform edit. They need a per-surface
  investigation (seed team-tags / interview sim_id match / deadline render) heavier than this iter's remaining
  budget. **Routed to iter-26** (handler FIND-M244-aireadiness-subrenders). gate c = **13/16** (12
  non-aireadiness + member-done). Metric stays 7/8 (gate c ticks only at 16/16 — the coarse-binary artifact).

## D4 — recovery-path finding (iter-24 correction): a stale `--public-host` demo needs `tailscale serve reset` BEFORE the bring-up
The recovery first FAILED: backend could not bind `0.0.0.0:18082` because a persisted `tailscale serve` rule
held the tailnet-IP:18082 listener (iter-24's tangle). up-injected DOES pre-reset serve, but only AFTER
`docker compose up` (too late — the container bind fails first + `set -e` aborts). The documented recovery
(tailscale-serve.md Step 7) is a manual `tailscale serve reset` BEFORE the bring-up; with it, recovery came up
green (autoverify OK, 17 containers). Corrects iter-24 D3 ("just use /demo-up" — insufficient alone).
