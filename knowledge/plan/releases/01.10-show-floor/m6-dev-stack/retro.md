# M6 — Retro

**Summary:** Extracted the shared port-offset engine into a new **`stack-core/`** section (settling the
M5-routed question — demo + dev now share it) and added a focused **`dev-stack/`** section: isolated local dev
stacks (`dev-N`, offset ports, alongside the main `anthropos` dev stack), **real Clerk by default**,
Clerkenstein injection **optional** (reusing `stack-injection`). Scoped honestly to the proven value, not a
speculative multi-concurrent-dev system. 87 tests (78 + 9 new), flake 3/3.

## Incidents this cycle
- **P3 — SC2318 (caught by shellcheck).** `local n="$1" … stack="$STACKS_DIR/dev-$n"` used `$n` within the same
  `local` (not yet effective). Split into two `local`s. Trivial; shellcheck is the net.

## What went well
- **Honest scoping (M6-D2)** — recognizing the base `rosetta-demo up` path already does non-injected
  multi-instance kept dev-stack thin (reuse stack-core + stack-injection) instead of duplicating up-injected.sh.
- **Extract-on-second-consumer (M6-D1)** — `gen_override.py` moved to `stack-core/` exactly when dev became its
  second consumer, not speculatively. Clean, and it left `clone_repos.py` (demo-specific) where it belongs.
- **The shared layers paid off** — dev-stack is ~110 lines because the multi-instance engine (stack-core) and the
  injection (stack-injection) already exist; it just composes them with a dev framing + a `-p dev-N` guard.

## What didn't / constraints
- The full *injected* dev bring-up isn't separately built (it's identical to demo's — reused, documented). If
  injected-dev demand appears, it's a future increment, not v1.1.
- Same relative-path fragility on the stack-core extraction (the CLI repoint) — caught by the suites.

## Carried forward → M7/M8
- M7 (`stack-seeding`) is next — seeds a stack (demo or dev) with the real `user_clerkenstein` identity + the
  casbin gotcha (inherited M3 routings). M8 may add a `/dev-up` skill if dev-stack proves used.

## Metrics
See [metrics.json](metrics.json). stack-core 4 + demo-stack 5 + stack-injection 69 + dev-stack 9 = 87 tests;
flake 3/3; shellcheck clean; deploy gate 100%/100% (clerkenstein untouched). Monorepo 75 commits, 5 sections.
