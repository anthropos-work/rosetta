# iter-21 — progress

**Type:** tik (run 8, under TOK-03 — gate (f) 3 v2.3 drift-carries, final-push step 2)

Gate (f) = "the 3 v2.3 drift-carries burned-in live (BURNIN-M221 / F-M220-4 / PROBE-M218-c3)." This tik
burned in the **two demo-side carries** live on billion; the third (BURNIN-M221) is a from-scratch remote
DEV `/dev-up --public-host` bring-up (≥1h, backgroundable) — feasibility-assessed + routed to iter-22.

## What landed (live on billion, this tailnet peer)

### PROBE-M218-c3 — Cosmo cms/Directus 403 re-check ✅ BURNED IN (resolved)
The C-3 finding (M218/M221) was the Cosmo router logging **cms/Directus 403s on the content federation path**
(`getSkillPaths` / `_entities JobSimulation`). Re-checked live against billion's running federated stack
(router demo-1-graphql-1, localhost:15050):
- Router health **200**.
- **Zero cms/Directus content-path 403s** in the router log: `docker logs` matched 38 lines containing the
  substring "403", but **all 38 carry `"status": 200`** — "403" is a substring in a query hash / IP / UA,
  not a response status. 0 lines name `getSkillPaths`/`_entities`/`JobSimulation`/`cms`/`directus`.
- A live `getSkillPaths` query through the router returned a **GRAPHQL_VALIDATION_FAILED** (my query shape),
  i.e. the router federated + schema-validated — NOT a 403.
- The content surfaces render live via router→cms→Directus (coverage sweep GATE MET reached /library
  skill-paths + ai-simulations; gate (b) content-stories 47/47). The C-3 403 signature is **absent** —
  burned in as **resolved** on the live federated stack.

### F-M220-4 — ant-academy.sh re-runnable on a live public-host demo ✅ BURNED IN
The carry: a standalone academy re-run could not re-bind because a stale listener (respawner) held :13077.
The fix (M220 S5/i reap-before-launch + M221 iter-05 supervisor-reap) is implemented. Burned in live:
1. `ant-academy.sh 1` (standalone, already-running) correctly **detected the live academy + skipped** (no
   blind EADDRINUSE) — the idempotency half.
2. `ant-academy.sh 1 --stop` then relaunch on the live **public-host** demo (STACK_PUBLIC_HOST set):
   after `--stop`, a stale **supervisor respawned a next-server** that grabbed :13077 (pid 1136731) — the
   exact F-M220-4 contention. The relaunch's reap fired: **"reap: demo-1 ant-academy supervisor (stale) —
   killing … a respawner the socket-reap cannot reach"**, then rebound :13077, re-wired Clerkenstein (pk
   from .env.demo-1, no export plumbing), and **re-applied all 4 academy demopatches**
   (dev-origins admitting billion.taildc510.ts.net + fs-published fallback/public/chapter-body).
3. Post-re-run the academy serves + renders from the peer over https: **home 200 (974KB), chapter EN 200
   (450KB) + `?lang=it` 200 (550KB), catalog.json 200 (2.3MB)** — re-runnable + course-start intact.

## BURNIN-M221 — feasibility-assessed, routed to iter-22 (NOT a blocker)
BURNIN-M221 = a real `/dev-up --public-host` **remote DEV** stack, live-cycled (the flag built M220, fenced
byte-identical, never brought up as a live dev stack). Assessed on billion:
- **dev-stack tooling present** on billion's rext clone (`dev-stack/`, `dev-setdress.sh`, `migrate-dev.sh`).
- **Disk OK** — 75G free now; 68.85GB is reclaimable Docker build cache (prune → ~143G).
- **RAM is the constraint** — 7.3GB total, ~3.1GB available (demo-1 uses ~4.3GB); the frontend-tier prereq is
  12GB/stack, so a FULL dev stack won't fit alongside demo-1. **FEASIBLE via a reduced-profile
  `/dev-up 2 --public-host`** (backend profile, offset 20000 — no collision with demo-1's 10000; fits the
  available RAM; faithfully exercises the dev-path `--public-host` tailscale-serve + CORS wiring).
- **No stack-dev workspace on billion → from-scratch build** (clone repos + build backend image + migrate +
  --public-host wiring) — a ≥1h backgroundable bring-up, an iter's worth of work.
⇒ per the scope-creep tripwire (a 3rd line needing ≥1h + backgrounding → land what's complete, route the
rest), BURNIN-M221 routes to **iter-22** with a concrete feasible plan. NOT infeasible; NOT a blocker.

## Re-measure
- **Pre-iter metric:** 6/8.
- **Post-iter metric:** **6/8** (gate (f) needs all 3 carries; 2/3 burned in — PROBE-M218-c3 + F-M220-4).
- **Delta:** 0 on the binary-per-gate metric (2 of 3 carries burned in; the coarse metric ticks only at 3/3).

## Close — 2026-07-23

**Outcome:** gate (f) 2/3 burned in live on billion — PROBE-M218-c3 (0 cms/Directus content 403s; federation
renders) + F-M220-4 (academy re-run reaped the stale respawner + rebound :13077 + serves). BURNIN-M221
feasibility-assessed (reduced-profile `/dev-up 2 --public-host`, fits RAM, from-scratch ≥1h) → routed to
iter-22. Metric 6/8 (coarse-binary: gate f ticks only at 3/3). 0 platform edits, 0 code changes.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (6/8; gate f 2/3, gate c playthroughs remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 2/5 this run) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (PROBE-M218-c3 "403"-substring vs status:200 method); D2 (BURNIN-M221 feasibility + reduced-profile plan).
**Side-deliverables:** none.
**Routes carried forward:** BURNIN-M221 → iter-22 (from-scratch reduced-profile `/dev-up 2 --public-host` on billion, backgrounded with heartbeats + verdict; NOT a blocker); then gate (c) 16 Playthroughs LAST → 8/8. DEF-M239-01 as budget.
**Lessons:** grepping a router log for "403" is not a 403-response check — the substring appears in query hashes / IPs / UAs on status:200 lines; assert on the STATUS field, not the substring. (Same class as iter-20's D1 curl-grep lesson — raw-text grep ≠ semantic assertion.)
