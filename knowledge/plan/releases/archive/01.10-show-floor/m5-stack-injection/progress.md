# M5 — progress (section checklist)

**Milestone:** M5 — Extract the reusable `stack-injection` layer · **Shape:** section · **Status:** done (2026-06-04)

## Done
- [x] **`rosetta-extensions/stack-injection/` section created** — the reusable Clerk-mock injection layer.
- [x] **Moved the generic injection out of demo-stack** (history-preserving `git mv`): `inject.py` (the recipe
  wiring + publishable-key mint), `gen_injected_override.py` (the injected compose override), `apply-authn.sh`
  (the disarmed-colony vendoring), + `README-recipes.md` / `DEPLOYMENT-PROOF.md`. Fixed `apply-authn`'s
  clerkenstein path for the shallower depth (`../../` → `../`).
- [x] **The mock stayed in clerkenstein** (M5-D1) — the disarmed provider, the fake servers, the cmd binaries,
  the Go JWT codec + svix injector. `stack-injection` *consumes* clerkenstein (vendors the provider; references
  `clerk-webhook`), never the reverse.
- [x] **Split the tests** (M5-D3) — `stack-injection/tests/{test_injection.py (mint + injected-override),
  test_apply_authn.py (vendoring + the cross-section wiring regression)}`; `demo-stack/tests/test_tooling.py`
  keeps `gen_override` + `clone_repos`. **78 tests preserved** (stack-injection 69 + demo-stack 9).
- [x] **Repointed the consumers** — `demo-stack/up-injected.sh` + the `rosetta-demo` CLI `inject` command call
  in via `../stack-injection/`. The cross-section wiring tests verify this.
- [x] **The demo-ON / dev-OFF toggle documented** — `stack-injection/README.md`: injection is applied at
  orchestration time (the up-path is the toggle); demo invokes it, dev defaults to the bare path (real Clerk).

## Completeness Ledger (section)
- **Done (Fate 1):** all `In:` items — the section, the extraction, the test split, the consumer repoint, the
  toggle doc, the boundary.
- **Routed (Fate 3) → M6:** the shared port-offset engine (`gen_override.py`) stays in demo-stack for now (M5-D2);
  it moves to a shared spot in M6 when `dev-stack` is its second consumer. (Settles the M4-routed open question.)
- **Dropped / escape-hatch:** none.

## Verification
stack-injection 69 + demo-stack 9 = **78 tests** green; **flake gate 3/3**; shellcheck clean across both sections;
py_compile clean; clerkenstein untouched → **deploy gate 100%/100% (7/7)**. Monorepo committed + pushed (74 commits).
