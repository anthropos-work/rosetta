# M12 — Spec notes

Technical notes accumulate here during build.

## Pre-flight audits — section 1 (unified registry + allocator)
- **KB-fidelity (Phase 0b): YELLOW.** Report: `kb-fidelity-audit.md`. No blind-area blocker (the unified-registry + allocator topics are DOC-ONLY because they are *this milestone's* deliverables). 1 stale claim (GUIDE.md "registry assigns N") — non-load-bearing, converges to truth at M12 close → tracked as KB-1. 0 open items.
- Triples (topic → doc → code):
  - demo registry → `demo-stack/GUIDE.md` / `README.md` → `demo-stack/rosetta-demo` (`reg_set`/`reg_del`/`cmd_status`) + runtime `demo-stack/stacks/registry.json` (gitignored)
  - port-offset engine → `dev-stack/README.md`, `knowledge/README.md` → `stack-core/gen_override.py`
  - dev lifecycle (no registry) → `dev-stack/README.md` → `dev-stack/dev-stack` (`cmd_status` = docker ps only)
  - unified registry + first-available-N → rosetta `corpus/ops/rosetta_demo.md` (delivery target) → **new** `stack-core/stack_registry.py`

## Unified registry schema
_{type: dev|demo, N, ports, status, created} — one record per live stack._

## First-available-N allocator
_Scan registry + `docker ps`, return lowest free N; locked write; explicit-N override._

## Up/teardown wiring
_dev-stack + demo-stack consume the allocator; teardown frees the slot._
