# M17 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [ ] **`set -e` first-run-race audit** across `up-injected.sh` / `rosetta-demo` / `dev-stack` / `dev-setdress.sh`; risky spots fixed (`|| echo 0`-style or guarded).
- [ ] **Optional "wait-for-sentinel-ready"** defense-in-depth before migrate.
- [ ] **`stacksnap replay` re-run guard** — per-stack-isolated TRUNCATE/skip before COPY; 2nd replay no longer aborts/doubles.
- [ ] **`stackseed` casbin g2 + det-UUID `ON CONFLICT`** — 2nd seed no longer duplicates the grant / unique-violates.
- [ ] **`--reset` truncate list fixed** to include session/activity tables (reset-then-seed no longer collides).
- [ ] **Tested idempotency contract** per component (the guard trips on a real 2nd run).

## Verification
- [ ] Go `-race -count=1` + gofmt + `go vet` clean on `stack-seeding` + `stack-snapshot`.
- [ ] A regression test per guard, mutation-pinned (removing the guard fails the test).
- [ ] TRUNCATE target-class test proves it can only hit a per-stack-isolated offset target.
- [ ] flake 0 (5× shuffle).

## Notes
_(build notes appended here)_
