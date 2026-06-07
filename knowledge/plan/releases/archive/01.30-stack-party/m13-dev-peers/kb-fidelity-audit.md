---
title: "KB Fidelity Audit — M13 dev-peers"
date: 2026-06-07
scope: milestone:M13
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Local per-stack Directus on dev | `corpus/ops/snapshot-spec.md` § "The per-stack Directus store fork (M10-D2)" | `stack-snapshot/directus/provision.go` (`ProvisionPlan`/`EnvContract`/`Validate`) | PAIRED |
| Auto-snapshot replay (cache-first) on dev | `corpus/ops/snapshot-spec.md` § "The `stacksnap` CLI", § "The `.agentspace` manifest-cached store" | `stack-snapshot/cmd/stacksnap/main.go` (`replayCmd`, cache-hit-or-fail), `stack-snapshot/pg/pg.go` (`ParseStackN` → dev-N) | PAIRED |
| `dev-min` seed preset + dev auto-seed | `corpus/ops/seeding-spec.md` § "The blueprint", § "The CLI" | `stack-seeding/presets/*.seed.yaml`, `stack-seeding/blueprint/blueprint.go`, `stack-seeding/cmd/stackseed/main.go` | PAIRED (preset file = Delivers→) |
| n=0-dev-reset guard | `corpus/ops/seeding-spec.md` § "The CLI" (`refuses n=0 dev unless --force`) | `stack-seeding/cmd/stackseed/main.go:180-181` (`if n == 0 && !force`) | PAIRED |
| dev-stack bring-up + unified registry | `dev-stack/README.md`, `corpus/ops/rosetta_demo.md` | `dev-stack/dev-stack`, `stack-core/stack_registry.py` | PAIRED |
| Directus integration (CMS repoint) | `corpus/services/cms.md` § "Directus integration" + the v1.2 set-dressing note | `stack-snapshot/directus/provision.go` (`EnvContract.BaseAddr`) | PAIRED |
| dev-min preset + dev auto-seed (corpus delivery target) | `corpus/ops/seeding-spec.md` (to be written) | — (this milestone authors it) | DOC-ONLY (Delivers→) |
| dev as a replay target + local-Directus-on-dev (corpus delivery target) | `corpus/ops/snapshot-spec.md` (to be written) | — (this milestone authors it) | DOC-ONLY (Delivers→) |

## Fidelity Findings

1. **`stacksnap replay` already targets dev-N** — `corpus/ops/snapshot-spec.md` § CLI shows `--stack <demo-N|dev-N>`; `pg.ParseStackN` parses `dev-3 => 3` and the test suite pins it. **ALIGNED.** The replay machinery needs no change for M13 — dev as a replay target is pure bring-up wiring.
2. **Replay is cache-first by construction** — `replayCmd` resolves cache-hit vs stale (`store.Resolve`) and FAILS with exit 1 if not cached ("cannot replay — … Capture it first"); it never captures. The doc's cache-hit-vs-stale-vs-miss table (§ store) matches the code. **ALIGNED.** M13's "cache-first" requirement is satisfied by calling `replay` (not `capture`).
3. **n=0 reset guard present + documented** — `stackseed --reset` refuses N=0 unless `--force` (`main.go:180-181`); `seeding-spec.md` § CLI states "refuses n=0 dev unless --force". **ALIGNED.** M13 preserves it (no change needed; the dev-min seed path on build must not call `--reset` on N=0 without `--force`).
4. **M10 per-stack Directus store fork is a documented plan, not a live runner** — `provision.go` is a declarative `ProvisionPlan` (bootstrap → replay → boot) + `EnvContract.Validate()` (rejects prod Directus); `snapshot-spec.md` § store-fork says "the live container boot is a documented operational step". **ALIGNED.** M13's job is to add the dev bring-up wiring that executes this plan (or documents the operational step for dev as the M9b/M10 discipline allows).
5. **Preset file format matches the loader** — three shipped presets parse under strict `KnownFields(true)` + `Validate()` (pinned by `presets_test.go`); `seeding-spec.md` § blueprint shows the exact schema. **ALIGNED.** The new `dev-min` preset follows the same schema and will be pinned by the same test.

## Completeness Gaps
None critical. The two corpus delivery targets (dev-min preset section in `seeding-spec.md`; dev-replay-target + local-Directus-on-dev section in `snapshot-spec.md`) are explicit `Delivers →` lines in the milestone `overview.md` — DOC-ONLY for an upcoming milestone, authored as Phase 5 of this build. Not a blind area.

## Applied Fixes
None required — all PAIRED topics ALIGNED. Triples recorded in `spec-notes.md`.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. The snapshot/seed/registry machinery the milestone builds against is verified true; M13 is bring-up wiring + a new preset + the two corpus delivery docs. No load-bearing stale claims, no blind areas.
