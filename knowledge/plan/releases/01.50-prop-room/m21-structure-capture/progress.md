# M21 — Progress

**Status:** in-progress (iter-03 closed). **Shape:** iterative (exit gate in `overview.md`).
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

- iter-03 (tik, closed-fixed-partial): **structure-source blocker RESOLVED** — operator sanctioned a bounded read-only
  prod structural read via the wired `postgres` MCP (M21-D6). Captured the **real faithful structure** for all 9
  collections (exact `pg_catalog` DDL + registry inventory: 9 collections / 217 fields / 43 relations, 20 dangling →
  M23). **Decisive digest finding** (M21-D5 → option B): prod digest `6cd35278…` is over the full 53-table schema
  (27 system + 26 collections); surface captures 9 of 26 → whole-schema digest can never converge → re-key per-surface.
  furthest-stage stays 2 (structure not yet applied). EXIT_REASON user-blocker: the digest-keying fork (A vs B, touches
  shared taxonomy keying) surfaced to the operator. See iter-03/progress.md.

## Next-iter queue (Fate-3, → iter-04 under TOK-01)
- `STRUCT-M21-iter04-apply` — **operator chose option A (M21-D7):** the structure artifact must cover **all 26 user
  collections'** DDL + registry rows (row cache stays at the 9 public-content collections; the other 17 tables exist
  but empty). Apply to a fresh bootstrapped harness; confirm tables + registry → stage 3 (→ 4 with row replay).
- `STRUCT-M21-digest-keying` — **implement option A (M21-D7):** keep the whole-schema keying; converge by matching
  prod's schema (all 26 collections) + **pin/verify the Directus version** so the 27 system tables also match.
- `STRUCT-M21-iter03-artifact` (carried) + `directus_files` ref capture (wire the dead `media.go`) + M23 referential
  closure of the 20 dangling relations.
