---
milestone_shape: iterative
milestone: M253
title: "studio-desk first-paint"
status: planned
release: v2.7 "july jitter"
exit_gate: "On a cold demo (state the environment - laptop vs tailnet), first-meaningful-paint < 1000 ms (the .page-skeleton header+sidemenu shell visible) AND no blank > 1 s, p95 over 5 consecutive cold loads; never gate on networkidle; always gate on a fresh-green autoverify.json."
iteration_protocol_ref: corpus/ops/demo/latency-budget.md
re_scope_trigger: "5 consecutive toks without a viable strategy -> user-strategic-replan"
depends_on: [M249]
complexity: medium
created: 2026-07-23
---

# M253 — studio-desk first-paint

**Status:** `planned`  ·  **Shape:** `iterative` (perf, measure→patch→re-measure)  ·  **Complexity:** medium  ·  **Depends on:** M249

## Goal
studio-desk paints page content in **< 1 second** — no multi-second blank on a cold demo load; the `.page-skeleton`
header+sidemenu shell is visible immediately and data streams in after. Closes the studio "blank-page" field defect,
the last of the six v2.7 field defects touching the studio surface.

## Shape (why this shape)
`iterative`. This is a **performance** milestone measured against a latency budget, so it follows the
measure→patch→re-measure loop of the M218/M244 latency-budget lineage: instrument the boot per-leg, pin a demopatch,
re-measure, iterate until the FMP gate holds p95 over 5 consecutive cold loads. The dominant await is unknown up front
(`clerk.load`'s 10 s timeout vs `l12nService.init()`/`userService.canAccess()`), so the fix cannot be pre-designed —
the bootstrap per-leg measurement decides which reorder/no-op to pin. That exploratory measure-first character is what
makes it iterative rather than a `section`.

## Exit gate
On a cold demo (**state the environment — laptop vs tailnet**):
- **first-meaningful-paint < 1000 ms** — the `.page-skeleton` header+sidemenu shell visible,
- **AND no blank > 1 s**,
- **p95 over 5 consecutive cold loads**.
- **never gate on `networkidle`; always gate on a fresh-green `autoverify.json`.**

**Re-scope trigger:** 5 consecutive toks without a viable strategy → user-strategic-replan.

## Iteration protocol
`corpus/ops/demo/latency-budget.md` + `corpus/ops/demo/coverage-protocol.md` — measure→patch→re-measure: instrument the
studio boot per-leg (via the net-new FCP runner), attribute the blank to the dominant await, pin/re-pin the shell-before-
awaits + no-thirdparty demopatch shas, re-measure cold p95, repeat until FMP < 1000 ms and no blank > 1 s holds over 5
consecutive cold loads against a fresh-green `autoverify.json`.

## Scope
### In
- A **sha-pinned demopatch** reordering `core/main.ts` to paint the skeleton DOM **synchronously before**
  `clerk.load()` / `l12nService.init()` / `userService.canAccess()` — so the shell renders ahead of the boot awaits.
- A **`studio-desk-no-thirdparty` twin** demopatch — no-op the posthog / Sentry remote calls on the demo host (they
  cannot help on a Clerk-free demo and can add blocking latency).
- A **net-new studio-desk FCP runner** in `stack-verify/e2e/` (the existing `run-latency.sh` covers next-web/hiring
  **ACCESS** only — there is no studio first-paint harness today).
- (Established fact, not re-litigated) this is **NOT a dev-vs-prod build issue** — refuted: the demo already serves a
  production build.

### Out
- The studio **builder keys** (M252 — the AI key reaching the demo container).
- The studio **nav** — "← Back to Cockpit" + the logo/back/logout prod-eject fix (M249).

## Dependencies & parallelism
- **depends_on:** **M249** — M253 **extends the `build_frontend_studio_desk` studio patch ladder that M249 creates**
  (M249 owns the first-ever studio-desk source demopatch machinery; M253's two studio patches ride that ladder).
  Per the coordination rules, **M253 branches from post-M249 HEAD** (coordination rule 6: `up-injected.sh
  build_frontend_studio_desk` — M249 owns, M253 extends).
- **parallel_with:** M247 / M248 / M250 / M251 / M252 — **authoring** can proceed concurrently in worktree fan-out.
  Live-box contention: **M250 + M253 (both live-measured iteratives) serialize on one billion demo** (RAM won't hold
  two); **M253 can bootstrap its FCP loop on a LOCAL demo**, with cold-p95 confirmed on billion in **M254** (coordination
  rule 9).
- **Intra-milestone LANES:** three concurrent authoring lanes —
  1. **shell + no-thirdparty demopatches** (the two studio source patches on M249's ladder),
  2. **net-new FCP runner** (`stack-verify/e2e/`),
  3. **docs**,
  — then the **serial bottleneck:** the **measure→patch→re-measure loop** (single-driver; cannot be sharded — it
  gates on a fresh-green `autoverify.json` per cold load). **Recommended: up to 3 subagents for the authoring lanes,
  then one serial driver for the iter loop.**
- **Risk (R4, carried):** M253 extends the **first-ever** studio source patch machinery + the additive-UI pattern M249
  authors. *Mitigation:* the demopatch anchor→replace mechanism supports insertion (confirmed); M249 authors the
  pattern doc; **M253 serializes after M249**.

## KB dependencies
- `corpus/ops/demo/latency-budget.md` (the perf budget + the measure→patch→re-measure protocol)
- `corpus/ops/demo/demopatch-spec.md` (the sha-pinned demopatch mechanism + guards)
- `corpus/ops/demo/frontend-tier.md` (the demo UI tier / studio-desk cached image)
- `corpus/services/studio-desk.md` (the studio-desk service + boot model)

## Delivers
- `corpus/ops/demo/latency-budget.md` — a **studio-desk first-paint budget** (the < 1 s FMP gate + per-leg model).
- `corpus/ops/demo/demopatch-spec.md` — the **2 studio patches** (shell-before-awaits + `no-thirdparty`).
- `corpus/services/studio-desk.md` — the **MPA / empty-body boot model** (why the blank happens + the reorder fix).

## Open questions
- **Which await dominates on the tailnet** — `clerk.load`'s 10 s timeout vs `l12nService.init()` / `userService.canAccess()`?
  The bootstrap per-leg measurement decides **before** pinning the patch sha.
- **De-dupe the injected skeleton vs PageWrapper's own** — avoid painting two skeletons once the reorder lands.
