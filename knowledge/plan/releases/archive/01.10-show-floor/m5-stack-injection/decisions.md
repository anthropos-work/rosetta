# M5 — Decisions

## M5-D1 — the mock stays in clerkenstein; stack-injection only consumes it (build, 2026-06-04)
`stack-injection/` holds the **injection orchestration** (the Python recipe wiring `inject.py` +
`gen_injected_override.py`, and `apply-authn.sh`). The **mock itself** — the disarmed `colony/authn/provider/clerk`
provider, the fake FAPI/BAPI servers, the cmd binaries, the universal-key JWT codec, the svix injector — **stays in
`clerkenstein/`**. Rationale: moving clerkenstein's Go internals into stack-injection would make the *mock* depend on
the *injection layer* (backwards, and it would break the alignment gates' imports). The dependency direction is
**stack-injection → clerkenstein** (apply-authn vendors the provider; inject.py's webhook recipe references
`clerkenstein/clerk-webhook`). So the "extract the svix injector + JWT codec" idea from the M4 research is satisfied
by *reference*, not by physically moving the Go code.

## M5-D2 — `gen_override.py` (port-offset engine) stays in demo-stack for now (build, 2026-06-04)
Settles the M4-routed open question ("where does the shared port-offset engine live?"). `gen_override.py` is **stack
isolation** (port-offset + `!override` for multi-instance), **not Clerk injection** — so it's out of `stack-injection`'s
scope. Its only consumer today is `demo-stack`. It **stays in `demo-stack/lib/`** and **moves to a shared spot in M6**
when `dev-stack` becomes its second consumer (extract-on-second-consumer, not speculatively). Avoids creating a
shared section with one consumer.

## M5-D3 — test split: injection tests to stack-injection; wiring tests live with apply-authn (build, 2026-06-04)
`test_tooling.py` was split by module home: `TestInjectMint` + `TestGenInjectedOverride` → `stack-injection/tests/
test_injection.py` (they test the moved `inject.py` / `gen_injected_override.py`); `TestGenOverride` +
`TestCloneRefResolution` stayed in `demo-stack/tests/test_tooling.py`. The apply-authn + **cross-section wiring**
tests (`TestInjectorWiringRegression`: does `demo-stack/up-injected.sh` invoke `stack-injection/apply-authn.sh` at
its real path) live in `stack-injection/tests/test_apply_authn.py` — they're an injection concern (is my injection
correctly wired into the consumer), with a `DEMO_STACK` path var for the cross-section reference. All 78 preserved.

## Adversarial review
- **Did the split lose coverage?** No — sliced the existing classes by line range (the harden's mutation-tested
  YAML structural tests in `TestGenInjectedOverride` came across intact); combined count 78 unchanged; flake 3/3.
- **Did the depth change break the moved scripts/tests?** The move surfaced two path adjustments (apply-authn's
  `../../`→`../` clerkenstein ref; the test's cross-section `up-injected.sh` path + the cp-source depth) — caught by
  running the suites, fixed, re-verified. Same lesson as M4-D4: a move that changes nesting depth breaks relative paths.

## Open (routed to M6)
- The shared port-offset engine's home (M5-D2) — extract `gen_override.py` to a shared location when dev-stack lands.
