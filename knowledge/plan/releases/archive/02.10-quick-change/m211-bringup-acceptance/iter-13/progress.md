**Type:** tik (cross-port hook fix + last-residual assessment). Under TOK-01 move (4).

# iter-13 — tik progress

## Execution log
1. **Fixed the cross-port-follow hook** (`stack-verify/e2e`): branched `onCrossPortFollow` by destination port —
   studio-desk (:9000+offset) keeps its SPA markers; ant-academy (:3077+offset) asserts its own SSR portal home
   (`ANT_ACADEMY_HOME_SECTION`: `main, body` + a text floor). Compiles + 42 unit specs pass. Committed to the
   authoring copy (the coverage harness runs from there), moved tag `quick-change-m211`.
2. **Re-ran M42e:** `cross-port follow ok http://localhost:13077/ ← /home :: ant-academy home OK (localhost:13077,
   HTTP 200, marker present)` → **crossPortFollowFails 1→0**.
3. **Assessed the last residual** (`failingSections=1`): `/library/ai-simulations` `sim-card-grid` empty.
   `public.simulation_embeddings` is ABSENT (only `public.skill_embeddings` + `public.job_role_embeddings`
   loaded); the AI-sims grid's `searchSimulations` is a vector search needing simulation embeddings. A
   DATA/snapshot gap (iter-08 rc=5 cache-miss), not a code bug → routed to next session (3 candidate fixes).

## Close — 2026-07-08

**Outcome:** Cleared the academy cross-port-follow (crossPortFollowFails 1→0) by teaching the hook the
ant-academy destination. M42e employee coverage is now **ONE section short** — escapes=0, personaFailures=0,
crossPortFollowFails=0, notReached=0, reachable=62; the sole residual is the sim-embeddings-backed AI-sims grid
(a data gap, routed).
**Type:** tik
**Status:** closed-fixed (the cross-port target landed + verified; the sim-embeddings residual is assessed + routed, not a platform bug)
**Gate:** NOT MET (M42e is 1 section short — the sim-embeddings AI-sims grid)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (sim-embeddings is a Fate-3 data gap, fillable via tooling) — (5) cap-reached: **y (5th tik)** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D1 (the cross-port hook must be app-aware: studio-desk :9000 vs ant-academy :3077 — a demo-local rewrite of an out-of-demo link flips it from an escape to a followed destination that needs its own assertion), D2 (the last employee residual is a sim-embeddings DATA gap, not a code/federation bug — the federation + escapes + persona + cross-port are all green)
**Side-deliverables:** none.
**Routes carried forward (Fate-3 → next session):** `TOOLING-M211-sim-embeddings` (fill the AI-sims grid: snapshot cache-fill / local embedding generation / a frontend empty-query→publicJobSimulations demopatch); M42m manager coverage (+ the drifted studio-url/public-website-url demopatch re-pin); v2.0 Playthroughs; cold `/dev-up`.
**Lessons:** A re-sync release peels drift in LAYERS, each a distinct fix-surface: backend build-provenance (iter-10) → frontend prod-links (iter-12) → coverage-harness cross-port scope (iter-13) → content/embeddings (routed). The federation fix (iter-10) unblocked all of them by making the pages reachable.
