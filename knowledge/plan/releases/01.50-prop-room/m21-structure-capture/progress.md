# M21 — Progress

**Status:** in-progress (iter-02 closed). **Shape:** iterative (exit gate in `overview.md`).
**Build with:** `/developer-kit:build-mstone-iters`.
**Active strategy:** TOK-01 (staged-pipeline build toward the binary serve-anonymously gate — see `decisions.md`).
**Furthest pipeline stage passing:** 2 of 6 — now **LIVE-confirmed + secured** (iter-02: stage-2 `.local` email bug
fixed; structure-apply mechanism validated; the real 9-collection artifact + structure-source decision routed to iter-03).

## Running ledger
_Appended after each iter (tik = a standard iter toward the gate; tok = a strategy/retro iter)._

- iter-01 (tok/bootstrap): authored TOK-01 (staged-pipeline strategy) + the 6-stage metric + static baseline
  (stage 2/6); Phase 0b KB-fidelity YELLOW; infra confirmed runnable (Docker + cached directus image + complete
  row cache) — see iter-01/progress.md
- iter-02 (tik, closed-fixed-partial): live baseline established; **stage 2 secured** (fixed the `.local` admin-email
  bug Directus 11.6.1 rejects — M21-D1); baseline refined to **exit 5** not 4 (M21-D3); digest trap crystallized as
  full-schema-keyed (M21-D5); **structure-apply mechanism validated** = Directus `schema apply` of a snapshot YAML
  creates tables + registry (M21-D2); structure-source finding — pure option (c) can't provide prod types, real
  artifact + source decision routed to iter-03 (M21-D4). furthest-stage stays 2 (live-confirmed). Routes carried
  forward: STRUCT-M21-iter03-source, -iter03-artifact, -digest-keying, + directus_files wiring — see iter-02/progress.md

## Next-iter queue (Fate-3, → iter-03 under TOK-01)
- `STRUCT-M21-iter03-source` — resolve the structure source (check platform cms repo for a committed Directus schema
  first; else a/b/MCP-structural-read), behind the M9a capture-source policy.
- `STRUCT-M21-iter03-artifact` — produce + apply the real 9-collection structure snapshot (prod-faithful types);
  advance stage 2 → 3 (→ 4 if the cached-row replay then succeeds).
- `STRUCT-M21-digest-keying` — the stage-4 convergence decision (full-schema digest vs per-surface content-table key).
- `directus_files` ref capture — wire the dead `media.go` file-ref code (stage-3/4 sub-task).
