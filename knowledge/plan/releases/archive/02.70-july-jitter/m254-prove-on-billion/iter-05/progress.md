**Type:** tik (measurement cluster) — TOK-01 clusters 3 (latency solo) + 2 (coverage c-render). Read-only gates on the current fresh-green demo, no re-bring-up.

# M254 · iter-05 — progress

## Close — 2026-07-25

**Outcome:** banked **(h)-latency MET both vantages** (employee p95 1.43 s / manager 1.41 s, gate < 5 s, 5/5
ACCESS, SOLO) and **(c)-render LIVE-confirmed on 2/4 apps** (next-web + hiring). All 4 apps have the
back-to-cockpit patch applied at build + the `:17700` cockpit URL baked (resolve-to-stack). Surfaced a **real
(c) fragility defect on the native academy** and confirmed studio (c)-render is (f)-blocked.

**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (overall) — (h)-latency-p95 MET + (c) web/hiring render LIVE; (c) academy defect + (c) studio (f)-blocked route forward; (h)-Playthroughs still pending.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the academy/studio (c) items + f/g are routed-forward Fate-3 fix-iters, not blockers) — (5) cap-reached: n (tik 1 of the session) — (6) protocol-stop: n — Outcome: continue (→ iter-06).
**Decisions:** D1 (h)-latency MET both vantages; D2 (c)-render web+hiring LIVE + all-4 baked-to-stack; D3 (c) academy Back-to-Cockpit fragility defect (out-of-band revert, heal only on reconcile re-invoke); D4 (c) studio render (f)-coupled.
**Side-deliverables:** none (measurement only; temp Playwright specs used + removed, rext left clean).
**Routes carried forward:**
- **(c)-academy** — native academy Back-to-Cockpit item not durable on a long-running demo (`UserMenu.jsx` reverts to pristine out-of-band; `reapply_clone_patches` heal only fires on a re-invocation of `ant-academy.sh`). Fix in rext `ant-academy.sh` (durable re-heal). Handler: `FIX-M254-iter{NN}-academy-back-to-cockpit-durability`.
- **(c)-studio** — studio Back-to-Cockpit render + logo/back/logout resolve — built-confirmed (:17700 baked in `dist`), LIVE render gated by the (f) studio session-carry defect. Folds into the (f) fix-iter re-bring-up.
- **(e)** builder Playthrough; **(f)** studio session-carry; **(g)** host-sensitive tests; **(h)-Playthroughs** + live-browser specs.
**Lessons:** (1) The Docker apps (web/hiring/studio) bake the back-to-cockpit patch into an immutable image → durable; the NATIVE academy serves the mutable clone source, so its patch is only as durable as the clone stays patched — any out-of-band revert silently drops the item until `ant-academy.sh` is re-invoked. Image-baked vs clone-served is the durability line. (2) The account-menu trigger in the shared `packages/ui/NavbarTop` is a button labelled with the hero's FIRST name (Maya / Rae), not an `.ant-avatar`; the academy uses `.user-menu-trigger`. (3) Latency p95 both vantages is ~1.4 s cold on billion — comfortably inside the 5 s gate; the ERR_ABORTED anomalies are benign RSC-prefetch aborts, not failures.
