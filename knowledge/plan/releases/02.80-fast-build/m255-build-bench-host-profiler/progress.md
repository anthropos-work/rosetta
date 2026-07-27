# M255 — progress

## Section checklist

- [ ] 1. `buildbench` harness (rext `stack-core`) — n>=3, cold-images + truly-cold variants, JSON phase ledger + resource sampler, runs on billion AND laptop
- [ ] 2. `hostprofile` — capacity probe -> build plan, under the written headroom reserve contract
- [ ] 3. The safe-parallelism contract (the shared `$DEMO/next-web-app` clone race vs G2/G4/G5; per-stack image isolation)
- [ ] 4a. Spike: does `next-web` have a multi-stage/production Dockerfile sibling upstream?
- [ ] 4b. Spike: the truly-cold baseline (`docker builder prune -af`)
- [ ] 4c. Spike: the laptop baseline (M1 Pro, Docker VM allocation)
- [ ] 4d. Spike: is peak load1 4.90/8 a plateau or an I/O ceiling?
- [ ] 5. `corpus/ops/demo/build-budget.md` (net-new; the blind area)
- [ ] 6. The §8.5 corpus retraction, mirrored across all four docs (C1 rule)
- [ ] 7. The §8.6 cert hazard — expiry-aware re-mint
- [ ] **BARRIER VERDICT** recorded (GO / re-cut M257 gate with the user)
