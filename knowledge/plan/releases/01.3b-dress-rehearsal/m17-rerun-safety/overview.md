---
milestone: M17
slug: rerun-safety
version: v1.3b "dress rehearsal"
milestone_shape: section
status: archived
created: 2026-06-08
last_updated: 2026-06-09
complexity: medium
delivers: corpus/ops/idempotency.md (net-new) + re-run guards in stack-seeding/stack-snapshot
issues: ISSUE-7 (residual), ISSUE-11
---

# M17 — Bring-up re-run safety (idempotency + first-run race)

## Goal
A re-run of migrate / snapshot-replay / seed is either **safe-and-idempotent** or **fails loudly with a guard** —
never silently doubles data or aborts mid-surface.

## Why section
The non-idempotency is fully audited (ISSUE-11, verified against code): the exact write sites and the missing guards
are known. The deliverables are concrete edits + tests.

## Repo split
- **`rosetta-extensions`** (authoring → tag `dress-rehearsal-m17` → consume): all the guards + the race audit.
- **`rosetta`**: the net-new `corpus/ops/idempotency.md` contract doc.

## Scope
- **In (all `rosetta-extensions` code; one `rosetta` doc):**
  - **`set -e` first-run-race audit** — sweep `up-injected.sh` / `rosetta-demo` / `dev-stack` / `dev-setdress.sh` for
    the same class as the fixed ISSUE-7 migrate race (a `pipefail` query against a not-yet-created relation aborting
    the script); add an optional **"wait-for-sentinel-ready"** defense-in-depth. (ISSUE-7 residual.)
  - **`stacksnap replay` re-run protection** — a per-stack-isolated `TRUNCATE`/skip before the bare `COPY`
    (`stack-snapshot/pg/pg.go:261-281`, `replay/replay.go:74-79`) so a 2nd replay doesn't duplicate-key-abort
    mid-surface (PK tables) or silently double (no-unique tables). (ISSUE-11.)
  - **`stackseed` re-run protection** — `ON CONFLICT` for the casbin g2 grant (`seeders/identity.go:126-134`) + the
    deterministic-UUID rows (`seeders/activity.go:57-65`, `pg/pg.go:307-320`); **fix the stale `--reset` truncate
    list** (`cmd/stackseed/main.go:28-32`) to include the session/activity tables it currently skips. (ISSUE-11.)
  - An explicit, **tested** idempotency contract per component (the guard trips on a real 2nd run).
- **Out:** the verify net (M18); the *auto-chaining* of snapshot/seed into bring-up (M20 — M17 only makes the
  *primitives* re-run-safe so M20's chaining is safe to retry).

## Depends on
M16 (clean pushed baseline). **Parallel with:** **M18 (yes-with-caveats** — different surfaces; both touch
`up-injected.sh` in different regions). Lean sequential per the spine discipline.

## Open questions (resolve during build)
- Replay re-run: `TRUNCATE`-and-reload vs `ON CONFLICT DO NOTHING` — lean: TRUNCATE the per-stack target surface then
  reload (simplest correct semantics; the target is always per-stack-isolated).
- Safe-by-default vs an explicit `--idempotent`/`--force` flag — lean: safe-by-default with a loud guard.

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md` (the replay write path + the per-stack-isolation class)
- `corpus/ops/seeding-spec.md` (the seeder + the `--reset` model + the n=0 guard)
- the `stack-seeding` / `stack-snapshot` extension sections

## Delivers
- **→ rosetta:** `corpus/ops/idempotency.md` — net-new: the per-component re-run verdicts (migrate SAFE / replay +
  seed NOT-idempotent-without-the-new-guards) + the teardown-then-redo model + the guards now added.
- **→ rosetta-extensions:** the re-run guards in `stack-seeding` + `stack-snapshot`; the race-audit fixes.

## Risk
**(data-safety, blocks-prod-safety)** a `TRUNCATE` must *only* ever hit a per-stack-isolated offset target — the n=0
+ prod-isolation guards stay inviolate. Mitigate: TRUNCATE only the per-stack target class; tests pin it; the
isolation guard + n=0 guard remain in force.
