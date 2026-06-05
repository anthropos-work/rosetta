# M5 — Retro

**Summary:** Extracted the reusable **`stack-injection`** layer from `demo-stack` — the generic Clerk-mock
injection (`inject.py`, `gen_injected_override.py`, `apply-authn.sh`) is now its own monorepo section, consumable
by any stack with a **demo-ON / dev-OFF** toggle. The mock stayed in `clerkenstein` (the dependency runs
stack-injection → clerkenstein, not the reverse); the base port-offset engine stayed in `demo-stack` (it's stack
isolation, not injection — moves to a shared spot in M6 when dev-stack needs it). 78 tests preserved across the split.

## Incidents this cycle
- **P2 — the depth-change path break (again, caught by the suites).** Moving `apply-authn.sh` one level shallower
  (`demo-stack/inject/` → `stack-injection/`) broke its `../../clerkenstein` ref (→ `../`), and the cross-section
  test still pointed at the old `up-injected.sh` location + cp-source depth. Running the split suites caught all of
  it; fixed + re-verified. Same family as M4-D4 — relative paths are fragile across moves; the test suite is the net.

## What went well
- **Slicing tests by line range** preserved the harden's deep `TestGenInjectedOverride` coverage exactly (the
  mutation-tested YAML structural tests came across intact) — combined count stayed at 78, flake 3/3.
- **The cross-section wiring tests** (`TestInjectorWiringRegression`) immediately flagged the un-repointed
  `up-injected.sh` consumer — they're doing exactly their job (pinning the demo→injection wiring).
- **Clean boundary call (M5-D1/D2)** — keeping the mock in clerkenstein + the port-offset engine in demo-stack
  avoided two backwards/speculative extractions.

## What didn't / constraints
- The extraction surfaced the now-familiar relative-path fragility twice; a more move-resistant path scheme (e.g.
  a section-root marker) would help, but isn't worth a milestone — the suites catch it.

## Carried forward → M6
- Extract `gen_override.py` (the shared port-offset engine) to a shared location when `dev-stack` becomes its second
  consumer (M5-D2). M6 (dev-stack) consumes `stack-injection` for its *optional* injection.

## Metrics
See [metrics.json](metrics.json). stack-injection 69 + demo-stack 9 = 78 tests (preserved); flake 3/3; shellcheck
clean both sections; clerkenstein deploy gate 100%/100% (untouched). Monorepo 74 commits.
