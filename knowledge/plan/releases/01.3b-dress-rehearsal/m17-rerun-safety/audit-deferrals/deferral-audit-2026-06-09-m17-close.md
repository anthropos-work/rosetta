---
title: "Deferral Audit — milestone (M17 close)"
date: 2026-06-09
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

No repeat deferrals; no chronic patterns; no aged-out items. M17 added **zero** new deferrals — every M17
decision (D1–D9) is a Fate-1 landing. The one item M17 inherited as forward-routed work — the live
docker-harness migrate-race **behavior** test (M16 Fate-2 → M17) — **landed** in M17 (M17-D8), so its
destination is reached and it leaves the open ledger. The only remaining open deferral is the release-level
inherited escape-hatch DEF-M10-01 (S3 blob bytes + cloud `SnapshotStore` → v1.4, signed), untouched by M17 and
not aged out.

## Summary
- Total deferrals in scope: 2 (1 inherited-and-now-LANDED in-release Fate-2; 1 inherited cross-release escape-hatch, untouched)
- Single deferrals: 1 (DEF-M10-01, signed escape-hatch — destination v1.4)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- New deferrals introduced by M17: 0

## Deferral Inventory

1. **DEF-M16-01 — live docker-harness migrate-race BEHAVIOR test** (inherited within v1.3b; now LANDED)
   - origin_milestone: M16 (v1.3b) — recorded as Fate-2 confirm (M16-D7), already owned by M17's `In:` scope
   - first_deferred_on: 2026-06-08 (M16 close)
   - last_seen_in: `m17-rerun-safety/decisions.md` M17-D8 (resolved); code `demo-stack/tests/test_migrate_race_live.py` (3 tests, Docker-gated)
   - destination: M17 (reached)
   - reason_recorded: "M16 landed the static regression fence (`TestMigrateRaceGuard`); the runtime behavior test belongs with M17's race/idempotency work, which brings the docker harness."
   - partial_attempted: no — landed in full (the real `migrate-demo.sh` runs against a throwaway pgvector container in the race state + re-run idempotency proof; skips cleanly without Docker)

2. **DEF-M10-01 — S3 media blob bytes + cloud `SnapshotStore` backend** (inherited, cross-release; untouched)
   - origin_milestone: M10 (v1.2)
   - first_deferred_on: 2026-06-07 (v1.2 close, signed escape-hatch); re-affirmed at v1.3 close-release
   - last_seen_in: `knowledge/plan/roadmap-vision.md` (v1.4 staged)
   - destination: v1.4 (signed, KEEP-DEFERRED-WITH-SIGNOFF)
   - reason_recorded: "cross-release feature (cloud store / S3 bytes), explicitly OUT of v1.3b (tooling + docs only — zero platform feature work)."
   - partial_attempted: no

## Repeat-Deferral Patterns
None. DEF-M16-01 is a single in-release Fate-2 whose destination (M17) has now been reached — it is not a
repeat (it was never re-deferred). DEF-M10-01 has a single signed destination (v1.4), unchanged across v1.2
close → v1.3 close-release → v1.3b M16 close → M17 close.

## Fate-1 Investigation

### DEF-M16-01 — live docker-migrate-race behavior test
- **Fate-1 (land now, complete) feasible:** yes — and it DID land in M17 (its owning milestone).
- **Landing:** `demo-stack/tests/test_migrate_race_live.py` (3 Docker-gated tests) runs the REAL `migrate-demo.sh`
  against a throwaway pgvector Postgres container in the race state and proves both survival under the race AND
  re-run idempotency (M17-D8). Skips cleanly without Docker. The static fence (M16 `TestMigrateRaceGuard`) is
  complemented, not replaced. No partial slice — the whole behavior proof landed.
- **Verdict:** LANDED (Fate-1/Fate-2 destination reached). Leaves the open ledger.

### DEF-M10-01 — S3 media blob bytes + cloud SnapshotStore backend
- **Fate-1 feasible:** no — a cross-release feature (cloud store / S3 bytes), explicitly OUT of v1.3b's "tooling
  + docs only, zero platform feature work" charter.
- **Aging check (all four triggers negative):**
  - Deferred across ≥2 milestones *within a release*? No — single signed cross-release destination (v1.4).
  - Elapsed ≥3 months? No — signed 2026-06-07, ~2 days ago.
  - Destination milestone closed without landing? No — v1.4 is not yet designed, let alone closed.
  - Area touched substantively by a later milestone? **No** — M17 touched the re-run-guard surfaces
    (replay TRUNCATE-reload, seed idempotent-COPY/casbin, `--reset`, the bring-up scripts); it did NOT touch the
    snapshot-store / S3 / cloud-backend area (verified: `git diff e6161b0..HEAD` has zero store/s3/cloud/blob
    files). Authority intact.
- **Fate applies:** escape-hatch (KEEP-DEFERRED-WITH-SIGNOFF), unchanged. No fresh decision required.

## Recommendations
- **DEF-M16-01:** no action — LANDED in M17 (M17-D8). Record as resolved.
- **DEF-M10-01:** KEEP-DEFERRED-WITH-SIGNOFF — v1.4, unchanged. No new sign-off required (authority intact, not aged).

## Applied Changes
- None to roadmap or roadmap-vision. DEF-M16-01's resolution is already recorded in M17-D8. DEF-M10-01 is
  unchanged. This audit report is the record of the GREEN verdict for M17 close.

## Blocking Items (require user decision)
None.
