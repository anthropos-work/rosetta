# M6 — progress (section checklist)

**Milestone:** M6 — `dev-stack`: tooled local dev environment · **Shape:** section · **Status:** done (2026-06-04)

## Done
- [x] **`stack-core/` section (NEW)** — extracted `gen_override.py` (the shared port-offset / `!override`
  multi-instance engine) out of `demo-stack` (M6-D1). Now consumed by **both** demo + dev. Repointed the
  `rosetta-demo` CLI; split `TestGenOverride` → `stack-core/tests/test_gen_override.py`. **Settles the
  M5-routed open question** (the shared port-offset engine's home — extract-on-second-consumer).
- [x] **`dev-stack/` section (NEW)** — a focused tool for **isolated local dev stacks** (`dev-N` on offset
  ports, alongside the main `anthropos` dev stack), **real Clerk by default**, Clerkenstein injection
  **OPTIONAL** (`--inject`, reusing `../stack-injection`). Reuses `stack-core`'s `gen_override`; `-p dev-N`
  hard-refuses the main dev project. `up`/`down`/`gen`/`status`. + `README.md` + `.gitignore`.
- [x] **Tests + gates** — `dev-stack/tests/test_dev_stack.py` (9: guard rejects non-int/zero/usage/unknown-cmd;
  static pins on the dev-project refusal, the default offset 10000, inject-OFF-by-default, the stack-core +
  stack-injection reuse) + shellcheck-clean.
- [x] **Scope honesty (M6-D2)** — delivered the *proven* value (shared engine + dev tooling + optional
  injection), **not** a speculative multi-concurrent-dev system; the full injected dev bring-up reuses demo's
  flow (no separate machinery). Documented in `dev-stack/README.md § Scope note`.

## Completeness Ledger (section)
- **Done (Fate 1):** the `stack-core` extraction (M5-routed item landed), the `dev-stack` section + tests +
  docs, the demo-ON/dev-OFF reuse of stack-injection, the safety guard.
- **Routed (Fate 3):** none new. **Dropped / escape-hatch:** none.
- **Deliberately out of scope (M6-D2):** multi-concurrent-dev as a *requirement* (unproven demand); a separate
  dev-injection machinery (reuses demo's); a per-dev-stack clone strategy (dev uses the platform compose directly
  — left to need).

## Verification
stack-core 4 + demo-stack 5 + stack-injection 69 + dev-stack 9 = **87 tests** green; **flake 3/3**; shellcheck
clean (dev-stack + the repointed demo CLI); clerkenstein untouched → deploy gate 100%/100%. Monorepo committed +
pushed (75 commits, 5 sections).
