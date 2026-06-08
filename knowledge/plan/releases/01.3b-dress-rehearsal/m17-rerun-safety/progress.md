# M17 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **`set -e` first-run-race audit** across `up-injected.sh` / `rosetta-demo` / `dev-stack` / `dev-setdress.sh`; risky spots fixed (3 ISSUE-7 kin sites: GH_PAT loud-guard, DEV_PROJECT ×2 `|| true` fallback). 4th site (schema-create) surfaced by the live harness + fixed.
- [x] **"wait-for-sentinel-ready"** defense-in-depth before migrate (`wait_pg` via `pg_isready`/`SELECT 1` + `wait_sentinel_running`, both bounded + non-fatal).
- [x] **`stacksnap replay` re-run guard** — per-stack-isolated TRUNCATE-then-reload (child-first) before COPY; 2nd replay REPLACES, no abort/double.
- [x] **`stackseed` casbin g2 + det-UUID `ON CONFLICT`** — idempotent COPY-to-temp merge (all 7 seeders) + casbin `WHERE NOT EXISTS`; 2nd seed inserts 0.
- [x] **`--reset` truncate list fixed** to the full FK-ordered fleet (+ targeted casbin g2 DELETE, not TRUNCATE — preserves global policy).
- [x] **Tested idempotency contract** per component (replay e2e re-run gate, seed re-seed-inserts-nothing, casbin guard, live docker migrate-race harness).

## Verification
- [x] Go `-race -count=1` + gofmt + `go vet` clean on `stack-seeding` + `stack-snapshot`.
- [x] A regression test per guard, mutation-pinned (removing the guard fails the test — verified for replay clear, casbin, idempotent-copy).
- [x] TRUNCATE target-class test proves it can only hit a per-stack-isolated offset target (`truncateForReplaySQL` shape-pin + `DSNForOffset` offset chain).
- [x] flake 0 (5× shuffle, both Go modules).

## Notes
- Net-new `corpus/ops/idempotency.md` delivered + wired into demo/README, root CLAUDE.md, snapshot-spec, seeding-spec, safety.md (bidirectional).
- M15 safety-doc drift guards stay GREEN after the safety.md "See also" addition.
- Python: demo-stack 27 (test_tooling, +6 over M16's 21) + 3 live docker (test_migrate_race_live, Docker-gated).
- The live harness surfaced + fixed a 4th latent ISSUE-7-class site (schema-create `docker exec` under set -e).
