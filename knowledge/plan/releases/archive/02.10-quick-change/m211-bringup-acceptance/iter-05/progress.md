**Type:** tik — under TOK-01. Protocol: run bring-up phase (verify) → triage → re-measure.

# iter-05 — tik progress

## Execution log
1. **autoverify.sh --project anthropos**: `✓ backend /api/health 200 on :8082`; `✓ sentinel.casbin_rules =
   170 (authz policy loaded)`; then `⚠ verify live reported failing probe(s)` (non-fatal, exit 0).
2. **Full verify.sh**: all 11 backend/core services LIVE (postgresql[healthy], redis[healthy], sentinel,
   backend, skillpath, jobsimulation, cms, storage, roadrunner, graphql, gotenberg) + all core READINESS
   GREEN (postgres-schemas, redis-ping, graphql-introspection, gotenberg-version 8.34.0, sentinel-rpc,
   storage-rpc). 4 FAIL: next-web-app + studio-desk (HTTP 000000 — native frontends not started on the
   backend-only warm stack), directus + directus-collections (HTTP 000000 — no local per-stack Directus;
   the warm dev stack reads content live from prod, the documented no-`--local-content` fallback).
3. **Merged-platform assertion (readiness.sh:47-49):** `expected=(public sentinel cms jobsimulation
   skillpath extensions)` with the explicit M209 comment "there is no `skiller` schema to expect anymore."
   The probe PASSES ("all expected schemas present") with the skiller schema absent → the assertion is
   correctly merged-shaped (a stale list expecting skiller would FAIL here).
4. **Scoped verify** (`STACK_SERVICES="postgresql redis sentinel backend skillpath jobsimulation cms
   storage roadrunner graphql gotenberg"`) → **"✓ all live probes passed"** (exit 0). The UI/directus rows
   are skipped (not false-failed) — verification.md's scope model working.

## Re-measurement (gate sub-conditions, warm)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) compose / no-skiller | MET | MET |
| (b) replay loads public.* | MET | MET |
| (c) seed closure green | MET | MET |
| (d) verify merged-platform assertion | casbin cheap-win only | **MET** (scoped verify all-green; expected-schemas no skiller; graphql 4-subgraph; casbin loaded) |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET (needs UI tier / cold demo bring-up) |
| (f) 0 residual skiller refs | clean (schema) | expected-schemas clean; 2 skiller repo-test-cmd branches in stack-verify/repos/run.sh noted (repos scope, not live path) → iter-06 |
**Metric:** fully-met sub-conditions 3/6 → **4/6** (d MET).

## Close — 2026-07-08

**Outcome:** The verify net's merged-platform assertion is GREEN on the warm merged stack — scoped backend
verify all-green, expected-schemas list has no skiller (passes with skiller absent), graphql 4-subgraph
introspects, casbin policy loaded (170). The 4 unscoped failures are UI-tier/local-Directus services not
started on a backend-only prod-read stack (correctly skipped when scoped).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (4/6 fully; (e) M42 coverage + v2.0 Playthroughs + the full COLD /dev-up + /demo-up remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-05 D1 (merged-platform verify assertion confirmed live), D2 (UI/directus failures are scope artifacts, not merge defects)
**Side-deliverables:** none
**Routes carried forward:** (1) iter-06 → sub-condition (f) sweep + clean the 2 skiller repo-test-cmd branches in stack-verify/repos/run.sh. (2) NEXT SESSION (heavy, reap-risky): sub-condition (e) M42 coverage sweep + v2.0 Playthroughs + the full COLD /dev-up + /demo-up proofs — these need the UI tier + a cold demo bring-up (Playwright + ~3-min UI Docker build), which exceed a single reap-safe foreground tik; run as dedicated cold bring-ups next session.
**Lessons:** On a backend-only warm stack, run verify SCOPED (STACK_SERVICES=backend set) — bare verify.sh false-fails the unstarted UI tier + prod-read directus. The merged-platform assertion lives in readiness.sh's expected-schemas list; it passing with skiller absent is the positive proof the verify tooling is merged-shaped.
