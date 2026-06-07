# M10 — Retro (Public Directus content snapshot, the second real surface)

## Summary
M10 took the snapshot framework to its **harder second surface**: the 9-table public **Directus content** library
(global simulation / skill-path templates). The defining work was the **firewall generalization** — the spike
flagged that the framework's public boundary was hardcoded `organization_id IS NULL`, which the Directus surface
(public = `private=false AND tenant_id IS NULL AND status='published'`) does not fit. M10 generalized it to a
per-surface `PublicPredicate` (M10-D1) that taxonomy and Directus both pass, with a zero-value fallback to the
org-only default so taxonomy stays byte-for-byte unchanged. On top of that: the directus surface in FK replay order
with **multi-level parent-scope chains** via `ParentScope.ParentFilter` (M10-D4, chasing column-less intermediates
like `task_checks → sim_tasks → simulations` to the scope-bearing root in one subquery); the **per-stack Directus
store fork** (bootstrap→replay→boot on the per-stack Directus-backing Postgres, M10-D2); media **refs** (1,311
public `directus_files` rows, bytes S3-gated → v1.3); the content fidelity gene; and the `ContentSnapshotSeeder`
DAG node + the `sim_id`/`skill_path_id`/`resource_id` **linkage resolver** that finally points the v1.1 free-value
session/assignment content refs at real replayed public templates. Coverage moved `content` `waived-m7c →
snapshot-seeded-m10` — **nothing left waived → 100% over the full catalog, the v1.2 set-dressing thesis structurally
complete.** Spike → build (5 sections) → 5 harden passes → close clean (single attempt).

A notable build-time correction: the spike's premise that the public Directus content lives in a **separate**
self-hosted Directus Postgres with no reachable DSN was **disproven** — the `directus` schema lives in the SAME app
Postgres, read-only via `marco_read` (M10-D2, KB-1). The surface captures over the same `--dsn` as taxonomy, just a
different schema. This simplified the source story and is now corrected in `db-access.md` + `snapshot-spec.md`.

## Incidents This Cycle
- **None.** The M10 build was sound: 5 harden passes (cap reached cleanly) surfaced **no production-code defects** —
  only uncovered (correct) paths, which the harden tests then pinned. Flake gate 0 (5/5 both modules); both modules
  `-race` green; fuzz corpora 0 crashers. No P2s, no regressions.

## What Went Well
- The **per-surface `PublicPredicate`** generalization was the minimal change that let the directus surface reuse
  ALL of the plan/capture/parent-scope/firewall machinery — and the zero-value→`DefaultPredicate` fallback kept the
  taxonomy surface and all M9a/M9b behavior provably unchanged (verified, not assumed).
- **`ParentScope.ParentFilter`** generalized the M9b parent-scope pattern to multi-level chains without nesting
  structs — an audit-legible, single-subquery extension; an empty `ParentFilter` keeps the scope-bearing-parent
  default. Directus's FK-less (app-layer relations) tree was handled by convention validated against prod counts.
- The firewall stayed **pure (DB-free)** through the generalization: predicate logic unit-tested against fakes; the
  `ParentFilter` SQL is surface-author constants with quoted identifiers (no injection surface).

## What Didn't
- The **handbook package index drifted**: the `taxonomy` (M9b) and `directus` (M10) surface packages existed on
  disk but were never added to the `stack-snapshot/README.md` package table — a latent M9b miss caught only at the
  M10 close by the per-unit-handbook contract check. Fixed here (both rows added). Lesson: the index-row addition
  belongs in the section's own build commit, not deferred to close.

## Carried Forward
- **S3 media blob BYTES** (DEF-M10-01) → **v1.3** — the cloud snapshot store / S3 backend seed (roadmap-vision.md,
  user note #4). The structural floor (refs + ref-columns + the 1,311 public file rows + local-storage/placeholder
  adapter) landed in M10; the bytes are the S3-access-gated operational add (Fate 2, confirmed-covered, M10-D5). Not
  escape-hatch — v1.2's 100%-coverage thesis is delivered by the refs floor.

## Metrics Delta (from metrics.json)
- **Go test funcs:** 635 → **701** (+66) — stack-snapshot 167→207 (+40), stack-seeding 204→230 (+26).
- **Coverage (M10-touched):** firewall 95.9%→100%, directus 100%, capture 93.8%→98.8%, the cmd/stacksnap M10
  adapters →100%, the seeders M10-touched →100%, the dna fidelity probe (`ReplayedNonPublicRows`) 0%→100%.
- **Data-DNA coverage:** **100% of the full catalog** (both formerly-`waived` surfaces promoted to `snapshot-seeded`).
- **Flake:** 0 (5/5 both M10-touched module subsets); fuzz 0 crashers; both modules `-race` green; gofmt+vet clean.
- **Review:** 2 findings (1 docs handbook-index, 1 decision-triage) — both fixed. Deferral audit GREEN.
