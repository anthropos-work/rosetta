---
title: "M246 → M247 confirmed-drift ledger"
date: 2026-07-23
author: M246 (re-sync & re-point barrier)
consumer: M247 (corpus re-ground) + downstream milestones
status: open (bring-up findings appended in §4)
---

# Confirmed-drift ledger (M246 → M247)

The HARD go/no-go barrier's payoff: drift observed while bumping the demo to the CONSOLIDATED platform
(current `origin/main` — skillpath decommissioned into `app`, 3 subgraphs, jobsimulation standalone).
Each row is a CONFIRMED divergence between what the corpus/tooling asserts and what the platform now
is, with a proposed fate. M247 triages; some rows name a better-fit milestone.

## Ground truth established at M246 (empirical)

- **repos.yml (current origin/main):** `app cms jobsimulation sentinel storage messenger roadrunner next-web-app studio-desk graphql-wundergraph` — **10 repos; skillpath ABSENT (0 mentions).**
- **docker-compose.yml services (current origin/main):** graphql, sentinel, backend, jobsimulation, cms, storage, customerio-sync, messenger, roadrunner, studio-desk, next-web-app, gotenberg, app-network — **no skillpath / no skiller service.** Only residual env vars `SKILLER_STREAM=skiller` + `SKILLPATH_STREAM=skillpath`.
- **platform origin/main HEAD:** `28c5f0d`. **app** bumped from the old pin (`v1.341.0`) to `3df8536`.
- **`public.skill_path_sessions`** is the skill-path runtime session table (skillpath's schema decommissioned) — the seeder re-point target; existence to be CONFIRMED by the §4 bring-up.

## Drift rows

| # | Drift | Where | Proposed fate |
|---|---|---|---|
| D-01 | Corpus asserts skillpath a **live Tier-1 service / 4th subgraph** (~30 files) | `corpus/` (service_taxonomy, rosetta_demo, external_services, CLAUDE.md, README, etc.) | **M247** — skillpath→app redirect + 3-subgraph truth (its declared scope) |
| D-02 | rext `stack-injection/gen_injected_override.py` `INJECTED` dict still carries a **dormant `skillpath` key**; `len(INJECTED)` prints "4 injected" | `gen_injected_override.py:17`, `:632` | **M247 or a rext-hygiene milestone** — inert on the consolidated demo (no compose service matches). Full removal ripples into `test_injection.py`. Comment already made truthful (M246 `88bcdb8`). |
| D-03 | `test_injection.py` pins skillpath as injected: `set(INJECTED)=={…,skillpath}` (:508), `demo-2-skillpath:injected` (:500), `_cfg` skillpath svc (:471), `--inject-svcs "…skillpath"` (:688); `_cfg` also still models the **already-merged skiller** (:468) | `stack-injection/tests/test_injection.py` | **M251 (test-health) or rext-hygiene** — behavioural-test redesign (model skillpath as decommissioned/not-injected). Not needed for green. |
| D-04 | `exposure_claim_guard.py` `_cfg` representative compose still lists **skillpath:8095** | `stack-injection/exposure_claim_guard.py:124` | **M251 / rext-hygiene** — test-only fixture; mirrors `test_injection.py::_cfg`. Update both together. |
| D-05 | Stale `stack-demo/skillpath/` clone dir lingers in the workspace (Jun 29) | `stack-demo/` workspace | Cosmetic/housekeeping — not in repos.yml so never re-cloned/built; harmless. Remove on next clean workspace. No milestone action required. |
| D-06 | `up-injected.sh:458` audit comment still enumerates historical token carriers "…skiller, skillpath…" | `demo-stack/up-injected.sh` (audit-prose comment) | **M247** — cosmetic prose; describes a 2026-06-11 audit (historical), low priority. |

## §4 bring-up findings — GO/NO-GO **PASS** (cold demo-2, 2026-07-23)

**The barrier PASSES.** One cold `/demo-up` on the CONSOLIDATED platform (current origin/main, pinned
advance + `DEMO_FRESHNESS_STRICT=1`, `--no-public-host`) came up GREEN on every load-bearing axis. No
migration/schema/subgraph break from the ~386-commit `app` bump.

**Proven GREEN (the go/no-go criteria):**
- **Build:** `up-injected.sh 2` exit 0; 16 platform containers Up. Pinned advance checked every clone out at the canonical shas (verified detached-HEAD at `3df8536` app / `e6c81850` next-web / etc.); strict-freshness passed.
- **Seeder re-point (LOAD-BEARING):** `public.skill_path_sessions` exists and the re-pointed seeder wrote **561 rows**. The legacy `skillpath.skill_path_sessions` table still exists but is **0 rows** (empty husk — the platform kept the table; runtime state now lives in `public`). Proof the re-point target is real + correct.
- **3 subgraphs, no skillpath:** injected compose has exactly 3 subgraph images (`demo-2-app`, `demo-2-cms`, `demo-2-jobsimulation`); **0 skillpath containers**; `graphql` liveness 200 + `graphql-introspection: ok` (supergraph composed).
- **Cheap-wins:** backend `/api/health` 200 on :28082; `sentinel.casbin_rules = 1250`; `postgres-schemas: all expected schemas present`; `public.skills = 42790`; per-stack Directus (21 collections, not prod); cockpit answering on :27700; hiring org set-dressed (5 positions + 42 sessions). **All 16 liveness + all readiness probes passed.**

**The 3 autoverify warnings — all NON-barrier-firing (drift, not breaks):**

| # | Warning | Classification | Fate |
|---|---|---|---|
| D-07 | A demo-patch NOT APPLIED: `apply-app-aireadiness-loadmembers` — target `app/internal/workforce/ai_readiness.go` **not found in the consolidated app clone** (the file moved/renamed in the 386-commit bump). Non-fatal (perf patch; the AI-readiness whole-org members hydration stays unbounded, ~180 s). | **Consolidation drift** — a sha-pinned demopatch whose anchor path went stale on the platform bump. | **M250** (AI-readiness fidelity — its explicit domain) re-anchors the patch. Ledgered; not needed for the barrier's green. |
| D-08 | fake-FAPI "NOT answering on :25400 — nobody can log in." | **Probe artifact, NOT a break.** The `demo-2-fake-fapi-1` container is Up, listening on `:5400 (TLS)`, 35-identity roster loaded; it returns HTTP 400 to the autoverify cheap-win because that probe uses `http://` against an **HTTPS** server (logs: "client sent an HTTP request to an HTTPS server"). Clerkenstein is functionally deployed. | **M251** (test-health) — fix the cheap-win to probe fake-FAPI over https (or drop the false "nobody can log in"). End-to-end browser login re-proven at **M254** on billion. Not M246-introduced. |
| D-09 | AI Academy NOT serving on :23077/library. | **Benign peripheral.** Vercel-native, not dockerized; `2/6` env keys short (ANTHROPIC/OPENAI/WUNDERGRAPH/FONTAWESOME). Non-fatal by design — "the cockpit/next-web/studio-desk carry the demo." Pre-existing class. | Standing/peripheral; not a consolidation break. Optional M251/M254 follow-up. |

**Net:** the 386-commit bump did NOT break bring-up. The one genuine consolidation-drift item that touches a demo SURFACE (D-07, AI-readiness perf patch anchor) is non-fatal and squarely in M250's scope. Downstream milestones are safe to scope against this proven-green consolidated topology.
