---
iteration_type: tik
iter_shape: production-fix
status: planned
created: 2026-06-25
---

# iter-17 — P5 (FATAL): post-set-dress Sentinel-reload reproducibility fix

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause). P5 closes the design-plan's
FATAL reproducibility gap (R1): the g3 sim-entitlement grant never reaches the running Sentinel enforcer
on a fresh demo-up. Independent of the iter-16 avatar blocker (a bring-up-script fix, not a seeder/
clerkenstein-image change).

**Step 0 re-survey:** confirmed live on demo-3 — Sentinel's enforcer calls `LoadPolicy()` once at startup
with no watcher (`sentinel/internal/authorization/casbin.go:75`); `migrate-demo.sh` starts/migrates
Sentinel BEFORE the set-dress seed writes the g3 grant; nothing reloads Sentinel after. iter-09's fix was a
MANUAL `docker restart` on the live stack — NOT wired into the bring-up, so a fresh demo-up keeps the
sim-start deny modal. Still the before-state.

**Cluster / target identified (design-plan R1, FATAL):** add a post-set-dress Sentinel reload to
`demo-stack/up-injected.sh` — after the set-dress block (which wrote the grant), before the auto-verify
(so the casbin/sim probes see the reloaded policy).

**Hypothesis:** the prefix-agnostic `AuthorizationService/Reload` RPC on the offset 8087 port re-runs
`LoadPolicy()` (CONFIRM it does, not just a probe); a `docker restart <demo>-sentinel-1` is the fallback.
With the reload in the bring-up, a fresh demo-up's `/sim/<slug>/start` renders the launch UI (no deny
modal) with NO manual restart.

**Expected lift:** the reload step appears in the bring-up sequence; after a re-seed+reload the sim /start
renders the launch UI with no manual restart. Proven via the bring-up log + a live RPC test on demo-3.

**Phase plan:** Phase B (read the sentinel Reload handler — CONFIRM it calls LoadPolicy) → Phase C (insert
the reload block in up-injected.sh; non-fatal; RPC-preferred + docker-restart fallback; corpus doc) →
Phase D (prove both paths on demo-3: the RPC at the offset port returns 200 = LoadPolicy re-ran; the
docker-restart fallback reloads the 341 g3 grants at startup).

**Escalation conditions:** if the Reload RPC did NOT actually re-run LoadPolicy (just a no-op probe) → use
the docker-restart path as primary. (Confirmed: the handler calls `e.LoadPolicy()` directly and returns
200 only on success.) No platform edits — the fix is in the demo bring-up script + a corpus doc.

**Acceptable close-no-lift:** n/a — the reload mechanism is confirmed; this lands.
