# M17 — Retro

_Closed 2026-06-09. The 2nd milestone of v1.3b "dress rehearsal" — a code milestone that made the bring-up primitives re-run-safe._

## Summary
M17 made a re-run of **migrate / snapshot-replay / seed** either **safe-and-idempotent** or **fail loudly with a guard** — never a silent double, never a mid-surface abort. It landed: the `set -e` first-run-race audit across the 4 bring-up scripts (loud `GH_PAT`, the `DEV_PROJECT` fallback, the `migrate-demo.sh` wait-for-ready defense-in-depth, and a 4th latent schema-create site the live harness surfaced); the `stacksnap replay` per-stack-isolated **TRUNCATE-then-reload** guard (child-first, no CASCADE, target-class pinned to a single-table TRUNCATE); the `stackseed` **idempotent COPY** (COPY-to-temp-then-merge `ON CONFLICT (id)`) + the casbin **`WHERE NOT EXISTS`** grant + the **fixed FK-ordered `--reset`** fleet; the live docker migrate-race harness (the M16 Fate-2 item LANDED as M17-D8); and the net-new code-cited `corpus/ops/idempotency.md`. The defining risk was data-safety — a wrong-target TRUNCATE — and the close verified it can't happen.

## Incidents This Cycle
- **None.** No P2 flakes (flake gate 5/5 on all 3 touched suites), no regressions (both Go modules `-race` green, demo-stack pytest 30/30), no rollbacks. The build-phase logic held under every harden probe (0 bugs surfaced across 2 passes). The one close finding was a docs ref-tag omission, not a defect.

## What Went Well
- **Prod-safety was made structural, not just tested.** The replay TRUNCATE has *two independent* fences: the pure `truncateForReplaySQL` is target-class-pinned (always a single-table `TRUNCATE … RESTART IDENTITY`, never DROP/DELETE/CASCADE/cross-schema, injection-quoted, degenerate identifiers stay a single relation), AND the connection is structurally the per-stack offset DSN (`pg.DSNForOffset`). A wrong-target TRUNCATE is impossible on two axes at once.
- **The COPY-to-temp-then-merge pattern got both properties.** `CopyRowsIdempotent` keeps the bulk-COPY speed the seeder exists to preserve (the 1k-user seed) AND makes a re-run a 0-row no-op — the temp `INCLUDING DEFAULTS` (not ALL) means the temp can't conflict before the real table's `ON CONFLICT` dedups. Applied uniformly to all 7 seeders (not just the 2 the overview named) — full Fate-1.
- **The live harness earned its keep immediately.** Running the real `migrate-demo.sh` against a throwaway pgvector container surfaced a 4th latent ISSUE-7-class site (the schema-create `docker exec psql` exits non-zero under `set -e` even with `ON_ERROR_STOP=0`) that the static fence would never have caught — fixed inline (M17-D9) and pinned.

## What Didn't
- **The tag trailed HEAD again (a recurring pattern).** The milestone tag was set at the build HEAD `dcef026`, then trailed the 2 harden commits — requiring the close reconcile `dcef026 → 0d36251`. This is now the **third milestone running** (M9a, M16, M17) where the close had to reconcile the tag. It's a reliable, documented step, but the recurrence says the tag should be set/moved *after* harden, not at build. Lesson for M18+: set the milestone tag at the end of harden, not at build HEAD.
- **The decision insights were blended without their back-ref tags.** `idempotency.md` described every guard's rationale in prose but omitted the `(#M17-DK)` tags the v1.3 corpus precedent (M10-Dx/M13-Dx) established — caught at close as the one Fate-1 finding. Minor, but it means the decision-triage tag convention isn't yet automatic at build/doc time.

## Carried Forward
- **Nothing M17-specific.** M17 added 0 deferrals; every decision D1–D9 is a Fate-1 landing, and it absorbed M16's one Fate-2 item (the live migrate-race behavior test → landed as M17-D8).
- **DEF-M10-01** (S3 blob bytes + cloud store → v1.4, signed) — inherited, orthogonal to M17 (which touched the re-run-guard surfaces, not the snapshot-store/S3 area), not aged. Unchanged.

## Metrics Delta
_(from metrics.json)_
- **Go test funcs:** 713 → **736** (+23: stack-seeding 236→252 [the temp-merge SQL builders + the casbin 7-col dedup + the `--reset` pins], stack-snapshot 224→231 [the replay TRUNCATE-reload guard + the `truncateForReplaySQL` target-class fence]).
- **Python test funcs:** 182 → **191** collected (+9: demo-stack 21→30 — the `TestSetEraceGuards` set-e race fences + the live docker migrate-race harness, 3 Docker-gated tests that skip cleanly).
- **Flake:** 0 (5/5 sequential on all 3 touched suites: stack-seeding, stack-snapshot, demo-stack).
- **Lint:** gofmt + go vet + shellcheck (4 scripts) + py_compile all CLEAN.
- **Coverage:** stack-seeding/pg 52.0→55.3% (+3.3, the idempotent-COPY branch); replay 100%, seeders 95.2%, cmd/stackseed 47.0% (unchanged — remainders are live-DB code proven by the docker harness).
- **Extensions:** tag `dress-rehearsal-m17` @ `0d36251` (reconciled from `dcef026`); `stack-demo/rosetta-extensions` re-consumed; `main` at `0d36251` on `origin`.
