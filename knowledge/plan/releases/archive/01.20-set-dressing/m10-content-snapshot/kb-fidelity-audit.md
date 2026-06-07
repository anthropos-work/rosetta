---
title: "KB Fidelity Audit — M10 Public Directus content snapshot"
date: 2026-06-06
scope: milestone:M10
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Snapshot capture/replay/firewall framework | `corpus/ops/snapshot-spec.md` | `stack-snapshot/{capture,firewall,replay,source,store,manifest,taxonomy}` | PAIRED |
| Prod DB read-access + public/customer boundary + Directus content source | `corpus/ops/db-access.md` | `mcp__postgres__query` (`directus` schema) | PAIRED |
| Directus + content model | `corpus/services/cms.md` | prod `directus` schema (26 user tables) | PAIRED |
| Seeding isolation guard + DAG + content-ref free values | `corpus/ops/seeding-spec.md` | `stack-seeding/{isolation,seeders,dna}` | PAIRED |
| jobsimulation `sim_id` consumer | `corpus/services/jobsimulation.md` | `stack-seeding/seeders/jobsim_sessions.go` | PAIRED |
| skillpath `skill_path_id` consumer | `corpus/services/skillpath.md` | `stack-seeding/seeders/skillpath_sessions.go` | PAIRED |

No BLIND-AREA: every M10 KB dependency doc exists.

## Fidelity Findings

1. **db-access.md + snapshot-spec.md — Directus content source location (STALE, load-bearing, M10-delivered).**
   - **Source:** `db-access.md` §"The public-vs-customer split" + `snapshot-spec.md` §"The tenant-data firewall".
   - **Expected (doc):** the public Directus content library lives in a "**separate self-hosted Directus store**
     (`content.anthropos.work`), its own Postgres" — implying capture needs a new DSN not reachable here.
   - **Actual (prod-verified 2026-06-06 via the wired `postgres` MCP):** the Directus content lives in a
     **`directus` schema inside the SAME `postgres` database** the `marco_read` MCP already reaches read-only.
     Verified: `directus.simulations` = 2,597 total / 647 public (`private=false`) / **304** strict-public-published
     (`private=false AND tenant_id IS NULL AND status='published'`); `directus.skill_paths` = 263 / 22 strict.
     The spike's "separate Directus Postgres / no DSN exposed" premise was DISPROVEN by the locked Decision 2.
   - **Verdict:** STALE.
   - **Fix owner:** update docs — **owned by M10's Delivers** (extend `snapshot-spec.md` with the Directus content
     path + the self-resolved `directus`-schema capture source; correct the `db-access.md` content-source line).
     The build reads the source from the locked Decision 2, not the doc, so this does not block coding.

2. **snapshot-spec.md — firewall predicate is org-only (CODE limitation M10 generalizes).**
   - **Source:** `snapshot-spec.md` §"The tenant-data firewall" + `firewall.OrgColumn`/`PublicFilter`.
   - **Expected (doc):** firewall admissibility is hardcoded to `organization_id IS NULL`.
   - **Actual:** correct for taxonomy; the `directus` surface's public predicate is
     `private = false AND tenant_id IS NULL AND status = 'published'` — a per-surface predicate the framework does
     not yet support. This is the architectural gap the spike flagged; M10 generalizes the firewall to a
     **per-surface public predicate** (org-default preserved for taxonomy backward-compat).
   - **Verdict:** ALIGNED-as-of-M9b (doc truthfully describes today's org-only firewall); M10 extends it and the
     doc update is in M10's Delivers.
   - **Fix owner:** update doc — owned by M10's Delivers.

## Completeness Gaps
None blocking. The M10 content path is intentionally DOC-ONLY today (snapshot-spec.md says "The public Directus
content library is M10") — the doc correctly defers it to this milestone.

## Applied Fixes
None inline — both findings are corrected by M10's own deliverables (the doc updates are part of the milestone's
`Delivers →` line). Recording them as KB-1/KB-2 in `decisions.md` for Phase 5 follow-through.

## Open Items (require user decision)
None. The 4 load-bearing decisions are already user-locked (2026-06-06); the strict-published subset (304 sims /
22 paths) is prod-verified and supersedes the spike's ~190-path estimate (which omitted the `tenant_id IS NULL`
intersection).

## Gate Result
YELLOW: proceed with tracking. No blind areas; the two stale claims are both on the M10 content path that this
milestone's `Delivers →` line corrects, and the build sources the truth from the locked decisions, not the stale
docs. Recorded as KB-1/KB-2 for Phase 5.
