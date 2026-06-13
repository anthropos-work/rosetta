---
milestone: M25
slug: field-bake
version: v1.5 "prop room"
milestone_shape: section
status: done
created: 2026-06-11
last_updated: 2026-06-13
complexity: medium
delivers: a short m25-field-bake/ field-bake log + any folded-back fixes (both repos as needed)
status_note: DONE — all 5 done-bars GREEN. The field-bake earned its keep: the operator-sanctioned prod read
  surfaced a real tenant-data-leak the firewall caught fail-closed (M25-D5) + two dangling-FK bugs (M25-D6/D7),
  all fixed in stack-snapshot; the local Directus now serves on demo + dev stacks (curl-proven).
---

# M25 — Field bake: the observable-behavior gate

## Goal
Prove the whole release on the **actual 16 GB box** with **observable behaviors** as the done-bars — so v1.5
pre-pays the field-fix tail that every prior release shipped after the fact (v1.3 → all of v1.3b; v1.3b → fix1–17).

## Why section
A fixed checklist of live runs with binary pass/fail done-bars. The only uncertainty is what bugs surface — and those
get fixed in their owning module inline, not as new scope. Build with `/developer-kit:build-milestone`.

## Repo split
Both repos as needed — bugs surfaced get fixed in their owning module (`rosetta-extensions` for tooling,
`rosetta` for docs/skills). The field-bake log lives in this milestone dir.

## Scope
- **In (live runs on the user's machine, fixes folded back into the owning module):**
  - Fresh **`/demo-up`** → the browser shows content **served by the local Directus** (data plane local) with **real
    images** (asset plane prod) + the verify net GREEN incl. the new Directus probes.
  - **`/dev-up 2 --local-content`** → same, on an opt-in dev stack; confirm **N=0 stays on the prod-read path**
    untouched (the documented manual opt-in recipe exercised once).
  - **Re-run everything twice** (idempotency live — re-provision + replay + seed).
  - The **cold-start capture** path exercised once (structure + rows captured together from a restored dump).
  - Clean **teardown** (`/demo-down`, `/dev-down 2`) reclaims the Directus container + its port; the registry is honest.
- **Out:** new features (this is a hardening gate — bugs surfaced get fixed in their owning module, no new scope).

## Depends on / parallel with
Depends on: M24 (the full release in place). Parallel with: none (the closing gate).

## Open questions
None — the done-bars are the observable behaviors above.

## KB dependencies
Every v1.5 doc; the `/demo-up`, `/dev-up`, `/demo-down`, `/dev-down`, `/test-platform` skills.

## Risk (resource — degrades-quality)
A Directus container per stack adds to the Docker-VM budget. Mitigate: runtime is cheap (measured ~0.9 GB/stack,
~0.66 GiB both frontends — boots/builds spike, not steady-state), keep the **max-2-co-resident-stacks** line, add
Directus to the 12 GB preflight, watch Docker-VM **disk** (the M3 disk-full precedent).
