---
milestone: M5
slug: stack-injection
version: v1.1 "show floor"
milestone_shape: section
status: planned
created: 2026-06-04
last_updated: 2026-06-04
delivers: rosetta-extensions/stack-injection/
---

# M5 — Extract the reusable `stack-injection` layer

## Goal
Pull the **generic Clerk-mock injection** out of `demo-stack` (and the shared bits of `clerkenstein`) into a
`rosetta-extensions/stack-injection/` section that **any** stack — demo, dev, or future — can consume, with a
clean **demo-ON / dev-OFF** toggle.

## Scope
### In
- **Extract the generic injection layer** into `stack-injection/`: the 4-recipe wiring (`inject.py` — the
  publishable-key mint + the `.env`/compose injection writers), the disarmed-colony vendoring (`apply-authn.sh`),
  the svix webhook injector, and the universal-key JWT codec they share.
- **Refactor `demo-stack`** to *consume* `stack-injection` (no behaviour change; the 78 tests + deploy gate stay green).
- **The inject toggle:** demo applies the recipes by default; a stack opts out (dev's default) by not invoking them.
- Settle the **port-offset/multi-instance engine** home (M4 open question): shared lib vs in `stack-injection`.
### Out
- `clerkenstein` keeps the **mock itself** (the disarmed provider + the fake FAPI/BAPI servers + the cmd binaries).
- `dev-stack` (M6), seeding (M7), recipes (M8). No platform-repo change.

## Depends on
M4 (the monorepo exists; clerkenstein + demo-stack are in it). **Blocks** M6 (dev-stack consumes the toggle).

## Open questions
- Direction of the shared JWT codec dependency (clerkenstein → stack-injection, or vice-versa) — pick the one
  that keeps the mock's packages clean.

## Exit (section)
`stack-injection/` exists + is consumed by demo-stack; the demo path is unchanged (suites green); the dev-OFF
default is wired and documented.
