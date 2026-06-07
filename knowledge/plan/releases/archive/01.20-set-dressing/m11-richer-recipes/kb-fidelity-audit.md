---
title: "KB Fidelity Audit — M11 (richer-recipes)"
date: 2026-06-06
scope: milestone:M11
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Seed presets (small/mid/large) | `corpus/ops/demo/README.md` (preset index); presets themselves in extensions | `rosetta-extensions/stack-seeding/presets/*.seed.yaml`, `blueprint/blueprint.go` | PAIRED |
| Demo recipe family | `corpus/ops/demo/recipe-*.md` (3) + README | (docs describe operator flow over the seeder/snapshot tools) | PAIRED |
| Snapshot capture/replay CLI | `corpus/ops/snapshot-spec.md` § "The `stacksnap` CLI" | `rosetta-extensions/stack-snapshot/cmd/stacksnap/{main,surfaces}.go` | PAIRED |
| Data-DNA coverage (100%) | `corpus/ops/seeding-spec.md` § data-DNA; `snapshot-spec.md` fidelity gate | `stack-seeding/dna/`, `seeders/{taxonomy,content}_snapshot.go` | PAIRED |
| `/demo-snapshot` skill | — (to be authored — M11 deliverable) | drives `stacksnap` | DOC-ONLY (deliverable) |
| CLAUDE.md skill table row | `CLAUDE.md` skills table | — | DOC-ONLY (deliverable) |

## Fidelity Findings

1. **`corpus/ops/snapshot-spec.md` — `stacksnap` CLI surface.** ALIGNED. Doc's three subcommands
   (`capture`/`replay`/`status`), the `--source dump-ingest|primary-read` precedence, `--dry-run`, exit codes
   (0/1/3), and `--store`/`STACKSNAP_STORE` default all match `cmd/stacksnap/main.go`. The `directus` surface is
   registered (`surfaces.go:32`).
2. **`corpus/ops/seeding-spec.md` — data-DNA coverage.** ALIGNED. Already reads "100% over the full catalog,
   nothing waived; both formerly-waived surfaces (`taxonomy` M9b + `content` M10) promoted to `snapshot-seeded`."
   Matches `seeders/{taxonomy,content}_snapshot.go` (verify-only DAG nodes) + the `datadna measure-snapshot` wiring.
3. **`corpus/ops/snapshot-spec.md` — Directus content surface.** ALIGNED. The 9-table FK-ordered set, the
   per-surface `PublicPredicate`, multi-level `ParentFilter`, the per-stack store fork, the content fidelity gene,
   and the `contentref` linkage all match the M10 code.
4. **`corpus/ops/demo/recipe-skill-progression.md` § Notes — STALE (load-bearing for M11, this milestone's work).**
   Lines ~57-59 still say taxonomy + content are "waived (see the data-DNA `waived` surfaces)" and "a future (v1.2)
   'richer demo worlds' theme." Both surfaces shipped in M9b/M10 and coverage is 100%. This is the recipe-refresh
   deliverable, not a blind area — fixed as M11 §2 (not a pre-flight blocker).
5. **`corpus/ops/demo/README.md` — preset paragraph.** Mentions presets but predates snapshots; the snapshot-replay
   prerequisite is undocumented in the family flow. M11 §1/§2 work — not stale-as-truth-blocker.

## Completeness Gaps

1. **`/demo-snapshot` skill — absent.** Deliverable, not a gap-to-block. M11 §3 authors it.
2. **CLAUDE.md skill table — missing `/demo-snapshot` row.** Deliverable. M11 §4.

## Applied Fixes
None at audit time. The two STALE recipe claims (findings 4-5) are the explicit subject of M11's recipe-refresh
section and will be corrected in-section (not pre-emptively), so the diff is coherent.

## Open Items (require user decision)
None.

## CODE finding routed to M11 hygiene (deliverable 5)
- **`stacksnap` CLI help text is stale (extensions code).** `cmd/stacksnap/main.go:85,94` says "framework
  (M9a/M9b)" and lists `surfaces: taxonomy | reference-toy` — omitting `directus`, which IS registered and
  shipped in M10. A flags↔docs drift inside the extensions repo (help-text-vs-registry). Small Fate-1 fix; lands
  in M11's release-close-hygiene-carry section (touches extensions → `stack-snapshot-m11` tag at close).

## Gate Result
GREEN — proceed to Phase 1. The only stale doc claims are the recipe "waived/future-v1.2" lines, which are the
literal subject of this milestone's recipe-refresh deliverable (corrected in-section for a coherent diff), plus one
CLI-help-text drift routed to the hygiene section. No blind areas; the heavy specs (snapshot-spec, seeding-spec) are
already current with M9b/M10.
