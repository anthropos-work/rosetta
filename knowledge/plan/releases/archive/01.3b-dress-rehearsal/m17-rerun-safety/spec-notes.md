# M17 — Spec Notes

_Technical notes accumulated during build — mechanisms, file paths (with line cites), gotchas, and the concrete shape of each change. Populated by `/developer-kit:build-milestone`. The verified code locations from the design-time research are in the milestone `overview.md` and `.agentspace/demo-up-issue.md`._

## Pre-flight audits — Section 1 (set -e first-run-race audit)

- KB-fidelity audit (Phase 0b): **GREEN** — report `kb-fidelity-audit.md`. All three KB-dependency docs (snapshot-spec / seeding-spec / safety) PAIRED + ALIGNED; the one BLIND-AREA (`corpus/ops/idempotency.md`) is the milestone's own `Delivers →` deliverable. Audited at sha `1ddfcf9` (rosetta) / `e6161b0` (extensions).
- Topic → doc → code triples in the report's Topic Inventory table.

## Code locations (verified against the live extensions tree @ e6161b0)

**Bring-up scripts (set -e race audit):**
- `demo-stack/up-injected.sh` — `set -euo pipefail` :13; the `GH_PAT=$(grep …)` at :20 (no `|| true`); migrate call :76 already `|| log`.
- `demo-stack/migrate-demo.sh` — ISSUE-7 fix already in place (:59-61 `|| echo 0` + :61 empty-guard). The remaining race surface: the `DEV_PROJECT=$(grep …)` pattern + the per-service atlas loop (`continue` guarded).
- `demo-stack/rosetta-demo` — `DEV_PROJECT=$(grep … || …)` :29-30 already `2>/dev/null|head` guarded.
- `dev-stack/dev-stack` — `DEV_PROJECT` :27-28 same pattern.
- `dev-stack/dev-setdress.sh` — `set -euo pipefail` :27; replay loop is `if … else (warn)` (non-fatal). The seed_step `|| die` is intentional-fatal.

**Replay re-run guard (stack-snapshot):**
- `replay/replay.go:74-79` — the per-table `CopyIn` loop (no pre-clear). `Replayer` iface :29-36.
- `pg/pg.go:261-281` — `CopyIn` bare `COPY … FROM STDIN`. `Exec` :284-289 available for TRUNCATE.
- `cmd/stacksnap/main.go:322` — `replay.Run(ctx, newReplayAdapter(conn), st, ref)`.

**Seed re-run guard (stack-seeding):**
- `seeders/identity.go:118-135` — `seedCasbinGrant` plain INSERT (no ON CONFLICT) via `c.Exec`.
- `seeders/activity.go:53-76` — det-UUID rows via `c.CopyRows` (COPY — no ON CONFLICT possible directly).
- `pg/pg.go:307-320` — `CopyRows` = `pgx.CopyFrom` (COPY; cannot express ON CONFLICT). `Exec` :323-329.
- `cmd/stackseed/main.go:28-32` — `resetTables` STALE: {memberships, users, organizations} only. Seeded surfaces also touch: `public.users`/`memberships`, `sentinel.casbin_rules|casbin_rule` (g2), `jobsimulation.sessions`+`activity_events`, `skillpath.sessions`, assignments, org.

**GOTCHA (KB-fidelity completeness gap 1):** `CopyRows`/`CopyIn` use Postgres COPY, which has **no `ON CONFLICT`**. The det-UUID + casbin ON-CONFLICT guards must route through `Exec` (INSERT … ON CONFLICT DO NOTHING) or a COPY-to-staging-then-INSERT-SELECT-ON-CONFLICT merge. Decision recorded in `decisions.md`.

**Load-bearing safety (do NOT weaken):**
- `stack-seeding/isolation/safety_doc_drift_test.go` — pins safety.md §2.2 (Clerk hosts / Directus tokens / S3 override / guard symbols).
- `stack-snapshot/cmd/stacksnap/main_drift_test.go` — pins safety.md §1.1/§1.4 (firewall predicates / gates / bounded-read SQL).
- New TRUNCATE target must be per-stack-isolated only; pin with a target-class test.

**M16 Fate-2 (live docker-harness migrate-race behavior test):**
- Static fence exists: `demo-stack/tests/test_tooling.py::TestMigrateRaceGuard` (3 tests). The docstring explicitly hands the LIVE behavior test to M17.
