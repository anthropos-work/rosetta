---
title: "Deferral Audit — M221 close (milestone scope)"
date: 2026-07-15
scope: milestone
invoked-by: close-milestone
---

## Verdict

**YELLOW.** The one release-chronic (the `dev-stack` suite, 5 milestones) was consciously **LAND-NOW'd at
M220** and is confirmed running this close pass (harden: 118 passed / 4 skipped, no hang). Of M221's own
routes, four **landed this milestone** (F1, F5/F5b, F10, and F-M221-06b — the last landed in THIS close's
final-harden pass). What remains are **four tail carries** (F4 + three not-reached live-infra probes), each
with a concrete, non-"no-time" reason and a recorded route. Because **M221 is the FINAL milestone of v2.3**,
LAND-NEXT (an in-release future milestone) is impossible for them — their cross-release finalization is
`/developer-kit:close-release` Phase 1b's job (release scope, "last chance before backlog"), which runs next.
**Zero silently-eroding scope. Zero blocking items that a conscious decision + recorded reason doesn't cover.**

## Summary

| | |
|---|---|
| Total deferrals in scope (M221 + inherited) | **9** |
| Landed this milestone (Fate 1) | **4** (F1, F5/F5b, F10, F-M221-06b) |
| Landed at a prior milestone (the chronic) | **1** (dev-stack suite → M220) |
| Tail carries → cross-release (v2.4, finalize at close-release) | **4** (F4, BURNIN-M221-dev-public-host, F-M220-4, PROBE-M218-c3-rerun) |
| Repeat groups (≥2 milestones) | **3** (BURNIN, F-M220-4, PROBE-M218-c3-rerun) — all non-gate, live-infra-only, dispositioned |
| Chronic "no-time" patterns unresolved | **0** |
| Blocking items (undecided) | **0** |
| Escape-hatch sign-off still owed (deferred to close-release) | **4** |

## Deferral Inventory

- **DEF-M221-01 — F1 snapshot store-root shadow** — origin M221 iter-04 (D-M221-04b) → **LANDED** iter-05 (non-empty-snapshots resolver + loud empty-store diagnostic) + a depth-2 edge test in this close's harden. **Fate 1, done.**
- **DEF-M221-02 — F5/F5b native detached-supervisor reap + gate-8 generated-file dirt** — origin M221 iter-04/05 → **LANDED** iter-05 (`reap_procs` port-anchored, supervisor-before-socket; academy clone left git-clean). **Fate 1, done.**
- **DEF-M221-03 — F10 `assert_ports_free` + demopatch freshness-ABORT field-exercise** — carried iter-04 → **field-exercised** iter-06 (D-M221-06d), closed rather than routed a 3rd time. **Fate 1, done.**
- **DEF-M221-04 — F-M221-06b `run-latency.sh` hardcoded `http://` cockpit URL** — origin M221 iter-06 (D-M221-06b), worked around in-cycle by driving the spec with `https` → **LANDED** in this close's final-harden pass as `LATENCY_SCHEME` (rext `a0f8615`; shellcheck-clean + construction smoke-test). **Fate 1, done.**
- **DEF-M221-05 — F4 academy grid renders 0 cards** (catalog serves 2,705, HTTP 200) — origin M221 iter-04, reconfirmed iter-06 (D-M221-06e). Client-side render defect **in the `ant-academy` platform repo**.
- **DEF-M221-06 — BURNIN-M221-dev-public-host** — the `/dev-up --public-host` flag (built M220) never got its live burn-in. Carried M220 → M221 (iter-03/04/05). Not reached.
- **DEF-M221-07 — F-M220-4 ant-academy re-runnability on a live public-host demo** — inherited from M220. Not reached.
- **DEF-M221-08 — PROBE-M218-c3-rerun router-403 re-check** — inherited from M218; needs the live box. Not reached.
- **DEF(chronic) — the `dev-stack` suite cannot be run** — the 5-milestone `DRIFT_DEFER` chronic; **LAND-NOW'd at M220** (one-line env fix). Re-verified GREEN in this close's harden (118 passed / 4 skipped / 111 s; the "spins forever" ghost stays discharged). **Closed — not a live carry.**

## Repeat-Deferral Patterns

### REPEAT: "BURNIN-M221-dev-public-host" · "F-M220-4" · "PROBE-M218-c3-rerun"
- **Span:** M218/M220 → M221 (2 milestones each; carried across M221 iters 03–06).
- **Flag:** **DRIFT_DEFER**, NOT chronic-no-time. Each carry's reason is a concrete infrastructure/scope fact, not "ran out of time": a burn-in needs live dev-stack cycling; the academy re-run + router probe need a live public-host box that the DEMO-path battery didn't exercise.
- **Current disposition:** none is a **gate condition** of v2.3 (whose proven scope is the DEMO path on `billion`, achieved 8/8). At the final milestone there is no in-release milestone to LAND-NEXT into → they route **cross-release to v2.4**, to be signed off at close-release.

*(The genuine chronic in this ledger — the `dev-stack` suite across 5 milestones — was already investigated-not-re-deferred and LAND-NOW'd at M220. That is the pattern this audit exists to catch, and it was caught one milestone ago.)*

## Fate-1 Investigation

- **F-M221-06b** — Fate-1 feasible: **YES**. A backward-compatible `LATENCY_SCHEME` env (default `http`, localhost unchanged). **Landed** in the final-harden pass (rext `a0f8615`).
- **F4 (academy grid)** — Fate-1 feasible: **NO — structurally**. The fix is a render-path change in the `ant-academy` **platform repo**, and v2.3's hard constraint is **zero platform-repo edits**. It is not "too much work" — it is out-of-bounds by release construction. Becoming demopatch-shaped needs ant-academy render investigation. Not a gate condition; the demo's primary surfaces (cockpit + next-web) are fully functional. → **v2.4 / documented known cosmetic gap.**
- **BURNIN-M221-dev-public-host** — Fate-1 feasible: **NO**. A stability burn-in requires repeated live dev-stack `--public-host` cycles — not repo-side close work. Non-gate. → **v2.4.**
- **F-M220-4** — Fate-1 feasible: **NO**. Needs a live public-host demo re-run of ant-academy. Non-gate. → **v2.4.**
- **PROBE-M218-c3-rerun** — Fate-1 feasible: **NO**. A router-403 re-check that needs the live box. Non-gate. → **v2.4.**

## Recommendations

| Item | Verdict | Destination |
|---|---|---|
| F1, F5/F5b, F10 | LAND-NOW (done in-milestone) | — |
| F-M221-06b | **LAND-NOW (done this close, harden `a0f8615`)** | — |
| dev-stack suite (chronic) | LAND-NOW (done at M220) | — |
| F4 | KEEP-DEFERRED-WITH-SIGNOFF | **v2.4** — documented known cosmetic gap (fix lives in a platform repo) |
| BURNIN-M221-dev-public-host | KEEP-DEFERRED-WITH-SIGNOFF | **v2.4** — live dev-path burn-in |
| F-M220-4 | KEEP-DEFERRED-WITH-SIGNOFF | **v2.4** — live public-host academy re-run |
| PROBE-M218-c3-rerun | KEEP-DEFERRED-WITH-SIGNOFF | **v2.4** — live-box router probe |

## Applied Changes

- No roadmap/overview edits made here (M221 is the final milestone — there is no in-release target to annotate). The four tail carries are recorded above with fresh reasons dated today; their **cross-release escape-hatch sign-off is explicitly handed to `/developer-kit:close-release` Phase 1b (release scope)**, which is the correct forum for a `RELEASE-SCOPE-DEFER:` decision + a `roadmap-vision.md`/v2.4 entry and runs immediately after this milestone merge.
- The four landed items (F1/F5/F5b/F10/F-M221-06b) are already reflected in the iter decisions + the Gate Outcome Ledger + the hardening-ledger.

## Blocking Items (require user decision)

**None block THIS milestone close.** The four tail carries are non-gate, dispositioned with concrete reasons, and — being cross-release at the final milestone — carry an **escape-hatch sign-off that is owed at close-release, not here**. This audit records them honestly so close-release's release-scope pass (and the user) see the complete tail. No item is being pushed forward without a conscious, reasoned decision.
