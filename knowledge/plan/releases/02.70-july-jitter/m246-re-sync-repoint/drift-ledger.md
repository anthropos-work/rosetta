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

## §4 bring-up findings

_(appended during the cold `/demo-up` prove — the empirical go/no-go evidence.)_
