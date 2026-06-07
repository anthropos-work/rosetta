# M12 — Progress

**Shape:** section · **Status:** planned

## Section checklist (from overview Scope.In)
- [x] Unified stack registry (dev+demo, type/N/ports/status) in stack-core — `stack-core/stack_registry.py` (§1)
- [x] First-available-N allocator (registry + `docker ps` reconcile, lowest free N), race-safe — `allocate()` under flock (§1)
- [x] Up accepts explicit-N OR auto-allocates; teardown frees the slot — both CLIs `up [N]` + `down` release (§2)
- [x] dev-stack + demo-stack bring-up consume the registry — `reg_cli allocate/release/set-ports` (§2)
- [x] Delivers: corpus/ops/rosetta_demo.md (unified registry + first-available-N) — new model section (§3) + stack-core/README.md

## Final review
_(filled at close)_
