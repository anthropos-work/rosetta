**Type:** tik (re-prove the build-scratch re-sync fix end-to-end + M42e coverage). Under TOK-01 move (4).

# iter-11 — tik progress

## Execution log
1. **Re-ran the demo-up** under the FIXED tooling (`rosetta-demo down 1` keep-images → `up-injected.sh 1`,
   reap-safe detached+poll). The re-sync log confirmed each scratch landed on its CURRENT tag:
   `app @ v1.334.1`, `cms @ v0.254.2`, `jobsimulation @ v0.253.0`, `skillpath @ v0.32.8`. The 4 injected
   services rebuilt fresh from the re-synced source. Bring-up GREEN (autoverify OK, casbin=1150, directus 21).
2. **Federation CONFIRMED FIXED (spot-check + diag probe):** `publicJobSimulations.skills` now resolves with
   REAL skills (`Process Assessment`, `Prompt Engineering`, `Business Process Reengineering`, …); a Playwright
   diag as Maya shows every GraphQL call returns 200 with data (`searchSimulations`, `libraryCategories`, …) —
   the `_entities(Skill.name)` 422 + gqlauthz panic are GONE.
3. **M42e employee coverage sweep** (detached, ~11 min): the crawler now follows the populated grids DEEP —
   `/library/ai-simulations` (q=21), `/library/skill-paths` (q=20), then dozens of real `/sim/...` +
   `/skill-path/.../chapter` pages. **reachable 7→50**.

## Re-measurement (M42e coverage — the (e) presence gate)
| metric | run-2 (pre-fix) | iter-11 (post-fix) |
|---|---|---|
| reachable | 7/150 (frontier exhausted at 7) | **50/150** |
| failingSections | 8 | **1** |
| personaFailures | 2 | **0** |
| escapes | 0 | 40 (NEW — see below) |
| GATE | NOT MET | NOT MET (1 section + 40 escapes) |

**The federation root-cause fix is PROVEN end-to-end** — all 8 federation-caused failing sections + both
persona failures cleared, and the crawl frontier expanded 7→50 (impossible before: empty grids gave the
crawler nothing to follow). The 2 residuals are a **DISTINCT class** (the re-synced FRONTEND's prod-links +
an embeddings gap), not the federation:
- **escapes=40:** every page carries a hardcoded nav link to prod `https://aiacademy.anthropos.work/`. Confirmed
  it lives in the next-web **frontend bundle** (nav/source), not replayed content — the re-synced (v2.1)
  frontend added an "AI Academy" nav item with no demo-local rewrite. Fix-surface: a new **demopatch**
  (source-rewrite to the demo-local ant-academy `:$((3077+OFFSET))`, or strip), like `next-web-studio-url`.
- **failingSections=1:** `/library/ai-simulations` renders "No simulations found / Showing 0 simulations"
  (a genuine empty-state) while `/library/skill-paths` + all detail pages populate. The AI-sims grid is
  `searchSimulations`-backed (vector search); it returns 0 without the **sim-embeddings** loaded — the iter-08
  cache-miss (`sim-embeddings replay skipped rc=5`). Fix-surface: the sim-embeddings snapshot cache-fill (a
  known optional route) OR confirm the re-synced grid's default query.

## Close — 2026-07-08

**Outcome:** The build-scratch re-sync fix (iter-10) is PROVEN end-to-end. The demo-up re-run rebuilt the
injected services from the current release tags, the federated `Skill` resolution now works, and M42e coverage
jumped from **8 failing + 2 persona (7 reachable)** to **1 failing + 0 persona (50 reachable)**. The gate isn't
GREEN yet — the 2 residuals are a distinct re-synced-frontend class (a prod AI-Academy nav link → 40 escapes;
the sim-embeddings-backed AI-sims grid empty).
**Type:** tik
**Status:** closed-fixed-partial (the federation fix proven end-to-end + a big metric lift; 2 residuals of a distinct class routed)
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (both residuals are Fate-1 tooling/data fixes, not platform bugs) — (5) cap-reached: n (3rd tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the build-scratch re-sync fix works: fresh v1.334.1 injected binaries → federated Skill.name resolves → grids/profile populate → reachable 7→50, failing 8→1, persona 2→0), D2 (the 2 residuals are re-synced-FRONTEND drift, distinct from the backend federation: the aiacademy nav link lives in the frontend bundle → needs a demopatch; the AI-sims grid needs sim-embeddings)
**Side-deliverables:** none (a proof + measurement iter; the fix landed in iter-10).
**Routes carried forward (Fate-3 → iter-12):**
- **escapes=40 → a `next-web-aiacademy-url` demopatch** (source-rewrite the hardcoded `aiacademy.anthropos.work`
  nav link to the demo-local ant-academy, or strip it) — mirrors `next-web-studio-url`. Also RE-PIN the
  drifted `next-web-studio-url` + `next-web-public-website-url` manifests to the re-synced source hashes (they
  G2-REFUSED this run). Handler: `TOOLING-M211-frontend-demopatch-repin`.
- **failingSections=1 → the sim-embeddings snapshot cache-fill** (or confirm the AI-sims grid's default query)
  so `/library/ai-simulations` populates. Handler: `TOOLING-M211-sim-embeddings`.
- iter-13+ = M42m manager coverage + v2.0 Playthroughs + cold `/dev-up`.
**Lessons:** (1) The build-scratch re-sync fix is validated by the biggest single-iter coverage lift of the
milestone (7→50 reachable). (2) A re-sync release surfaces drift in LAYERS: backend build-provenance (iter-10)
THEN frontend source-links/demopatch hashes (this iter) THEN content/embeddings — each a distinct fix-surface.
(3) `escapes` only surfaces once pages are reachable — a broken crawl (7 pages) hid the frontend prod-link
until the grids populated (50 pages).
