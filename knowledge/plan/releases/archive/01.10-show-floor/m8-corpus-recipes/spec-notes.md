# M8 — Spec notes

Seeded from Phase 1 research (2026-06-03); fill in at kickoff (build-blocked on M3+M4).

## Where it links
- New family under `corpus/ops/demo/` (matches the `staging-*` clustering); index from `corpus/README.md` Ops + root README/CLAUDE.md.
- Recipes follow `corpus/services/TEMPLATE.md` conventions.

## Presets
- small-200 / mid-500 / large-1k as `demo.seed.yaml` instances; each must actually seed an M3 stack (validated, not just authored).

## Carry-forwards landing here
- v1.0 express-gate CI: add a Node step to `anthropos-demo/clerkenstein/.github/workflows/alignment.yml` that runs the @clerk/express gate.
