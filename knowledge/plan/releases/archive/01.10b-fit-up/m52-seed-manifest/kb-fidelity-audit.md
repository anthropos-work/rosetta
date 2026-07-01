---
title: "KB Fidelity Audit — M52 (single auditable seed+gen manifest)"
date: 2026-07-01
scope: milestone:M52
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Mother prompt (generation template) | `corpus/ops/demo/ai-generation-spec.md` §1/§2b | `stack-seeding/blueprint/batch.go` (`DefaultBatchPromptTemplate`, `EffectiveBatches`) | PAIRED |
| Prompt-hash cache key integrity | `corpus/ops/demo/cache-spec.md` §2 | `blueprint/batch.go` (`MotherPromptHash`), `batchcache/cache.go` | PAIRED |
| Population blueprint (3 orgs incl. AI-readiness) | `corpus/ops/seeding-spec.md`, `demo/stories-spec.md` | `presets/stories.seed.yaml`, `blueprint/*.go` | PAIRED |
| Batch config (`--max-cost`/concurrency/re-roll) | `ai-generation-spec.md` §2c | `cmd/gen-batch/` | PAIRED |
| Snapshot sources (taxonomy + Directus capture version) | `snapshot-spec.md`, `cache-spec.md` §2 | `stack-snapshot/manifest/manifest.go` | PAIRED |
| Cockpit download surface | `demo/cockpit-spec.md` §Served endpoints | `demo-stack/cockpit.py`, `seeders/cockpit.go`, `up-injected.sh` | PAIRED |
| **Consolidated seed+gen manifest** | **`demo/seed-manifest-spec.md` (NEW — S4 deliverable)** | net-new (manifest file + loader + export) | **BLIND-AREA** |

## Fidelity Findings
All PAIRED topics verified ALIGNED against code:
1. **Mother prompt** — `DefaultBatchPromptTemplate` exists as a Go `const` (`batch.go:111`); `ai-generation-spec.md`'s
   claim that the mother prompt is embedded in Go templates (the "core gap" the milestone closes) is accurate. ALIGNED.
2. **Cache key** — `cache-spec.md` §2's key formula (`sha256(motherPrompt || "\x00" || captureVersion)`) matches
   `MotherPromptHash` (`batch.go:395`) + `batchcache.Open`. ALIGNED — this is the load-bearing integrity constraint S1 preserves.
3. **3 orgs** — `stories.seed.yaml` carries exactly Cervato/Solvantis/Northwind (the AI-readiness org is the 3rd
   story, `narrative: ai-readiness`); `stories-spec.md` + `seeding-spec.md` M51 status describe it. ALIGNED.
4. **Cockpit download** — `cockpit-spec.md`'s Served-endpoints table (`/manifest.json` → `Content-Disposition:
   attachment; filename="cockpit-manifest.json"`) matches `cockpit.py:365-375`. ALIGNED — S3 repoints this.
5. **Batch config** — `ai-generation-spec.md` §2c's `--max-cost` (mandatory) / `--max-concurrent` (default 5) /
   re-roll-on-malformed matches `cmd/gen-batch`. ALIGNED.
6. **Snapshot sources** — `snapshot-spec.md`'s capture manifest (`SchemaVersion`, `CapturedAt`, `Source`) is the
   capture-version the cache key extends with; `cache-spec.md` §2's capture-version-invalidation is accurate. ALIGNED.

No STALE claims. No UNVERIFIABLE claims.

## Completeness Gaps
None critical. The consolidated seed+gen manifest is a NET-NEW surface (no code yet) — a DOC-ONLY/BLIND-AREA that
S2/S4 author together. Nothing in the existing PAIRED docs is under-documented for this milestone's read.

## Applied Fixes
None needed (no stale claims). Topic→doc→code triples recorded in `spec-notes.md`.

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed to Phase 1. The single BLIND-AREA (`seed-manifest-spec.md`) is an explicit `Delivers →` milestone
deliverable (overview.md front-matter + §Delivers), authored as S4. Per the blind-area rule, a blind area already
promoted to a milestone deliverable satisfies its own gap — it is not a blocker.
