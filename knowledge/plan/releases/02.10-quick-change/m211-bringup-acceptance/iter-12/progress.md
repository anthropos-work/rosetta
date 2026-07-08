**Type:** tik (kill the 40 prod-eject escapes). Under TOK-01 move (4).

# iter-12 — tik progress

## Execution log
1. **Fixed the escapes at the source:** `urls.ts` already reads `NEXT_PUBLIC_ACADEMY_URL || prod`, so baked
   `NEXT_PUBLIC_ACADEMY_URL=http://localhost:$((3077+OFFSET))` (:13077, the demo's own ant-academy) into the
   next-web `.env.local` overlay + a contract fence. All 64 tests GREEN, shellcheck clean. Committed rext
   `ad7bee4`, moved tag `quick-change-m211`, re-pinned the consumption clone.
2. **Force-rebuilt `demo-1-next-web`** (removed the image so the offset-guard rebuilds; the guard validates the
   WUNDERGRAPH endpoint, not ACADEMY_URL) — reap-safe detached. Verified `:13077` inlined into the client
   bundle; the nav `aiAcademyMenuItem` uses `key: ACADEMY_URL` (`useNavbarSections.tsx:192`). Recreated the
   container.
3. **Re-ran M42e coverage:** **escapes 40→0** (the AI-Academy nav link is now demo-local — NO more prod-eject).
   Also `reachable 50→62`, `personaFailures` stays 0.

## Re-measurement (M42e coverage)
| metric | iter-11 | iter-12 |
|---|---|---|
| reachable | 50 | **62** |
| failingSections | 1 | 1 (sim-embeddings — unchanged) |
| personaFailures | 0 | 0 |
| **escapes** | 40 | **0** ✅ |
| crossPortFollowFails | 0 | **1** (NEW — see below) |
| GATE | NOT MET | NOT MET (1 section + 1 cross-port) |

**escapes FIXED (40→0).** But pointing the academy to :13077 exposed a **harness limitation**: the crawler now
FOLLOWS the demo-local academy link (:13077, ant-academy is UP http=200), and the `onCrossPortFollow` hook
(`coverage.spec.ts:177`) is hardcoded to assert **studio-desk** DOM markers (WelcomeSection/QuickActions) on
EVERY cross-port follow — `DEMO_LOCAL_BASES` includes `3077` but the hook only knows studio-desk (:9000). So
the academy follow FAILS "studio-desk home marker absent". This is a coverage-harness gap, not a demo defect.

## Close — 2026-07-08

**Outcome:** Killed all 40 prod-eject escapes by baking the demo-local `NEXT_PUBLIC_ACADEMY_URL` (:13077) — the
"AI Academy" nav link no longer ejects a presenter to prod. Exposed a harness limitation (the cross-port-follow
hook asserts studio-desk markers on the now-demo-local academy link). M42e is now 2 residuals short: the
sim-embeddings AI-sims grid + the academy cross-port-follow hook.
**Type:** tik
**Status:** closed-fixed (the escapes-40 target landed + verified; the new crossPortFollowFail is a distinct harness gap, routed)
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (both residuals are Fate-1 tooling fixes) — (5) cap-reached: n (4th tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the academy prod-eject needed NO demopatch — urls.ts already has the NEXT_PUBLIC_ACADEMY_URL override; just bake the demo-local value, mirroring STUDIO_URL), D2 (pointing the academy demo-local makes the crawler FOLLOW it → the studio-desk-only cross-port hook must learn the academy port)
**Side-deliverables:** none (the fix IS the planned scope).
**Routes carried forward (Fate-3 → iter-13):**
- **crossPortFollowFails=1 → extend the cross-port-follow hook** (`stack-verify/e2e/tests/coverage.spec.ts`):
  distinguish the academy port (:3077+offset) from studio-desk (:9000+offset) — assert the academy loads
  non-blank/non-login on the right host (or its own markers), not studio-desk markers. Handler:
  `TOOLING-M211-crossport-academy`.
- **failingSections=1 → the sim-embeddings gap** (`public.simulation_embeddings` absent → the
  `searchSimulations`-backed AI-sims grid returns empty). Investigate whether it's fillable (snapshot
  cache-fill needs a source — may be unfillable locally like the taxonomy cold-start) OR the re-synced grid's
  default query. Handler: `TOOLING-M211-sim-embeddings` (may route to next session if unfillable).
- iter-14+ = M42m manager coverage (+ the drifted studio-url/public-website-url demopatch re-pin) + v2.0
  Playthroughs + cold `/dev-up`.
**Lessons:** (1) A re-synced frontend can add prod-eject nav links that stay hidden until the crawl is
UN-broken (escapes only surface on reachable pages). (2) Making an out-of-demo link demo-local can flip it from
an ESCAPE into a cross-port FOLLOW — the two are different gate inputs; a demo-local rewrite must be paired
with the crawler knowing how to validate that destination.
