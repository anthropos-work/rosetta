# M255 — progress

## Section checklist

- [ ] 1. `buildbench` harness (rext `stack-core`) — n≥3 on billion, **cold-images variant only**; JSON phase ledger + 10 s sampler; **every entry records the invocation + full `DEMO_*` env snapshot**; one informational n=1 laptop run
- [ ] 2. Campaign protocol + reclaim — hard-failing pre-rep disk/cache assert · reclaim step between reps (**L6 promoted here**) · per-rep `docker system df` declaration · **`DEMO_DISK_MIN_GIB` re-sized** · ENOSPC→"redis exited (1)" signature noted
- [ ] 3. Host profiles + headroom assert — `stack-core/hostprofiles/{billion,laptop}.json` measured + checked in; a **failing** sampler assert (load1 / summed heap / free disk); decision recorded reconciling "fail loudly" vs the never-block-a-bring-up pre-flight contract
- [ ] 4. The **union-apply** parallelism rule + guard test (shared members byte-identical; non-shared under disjoint `apps/*` or waived inert). "Separate clones" option deleted
- [ ] 5a. **Spike (a) — the 15-min L1 experiment** on the rext-owned `hiring.Dockerfile`; measure the export delta; record `NEXT_PRIVATE_STANDALONE=1` + its demopatch fallback  ← **THE BARRIER DECIDER**
- [ ] 5d. Spike (d) — is peak load1 4.90/8 a plateau or an I/O ceiling?
- [ ] 5e. Spike (e) — host-vs-peer topology for M258's composed command
- [ ] 6. `corpus/ops/demo/build-budget.md` (net-new; the blind area)
- [ ] 7. Security + cert hazard (**non-gating**) — expiry-aware re-mint + the paired `corpus/ops/safety.md` §3 amendment
- [ ] **BARRIER VERDICT** recorded (GO / re-cut M257's gate with the user)

## Cut from the first draft (see roadmap.md § design decisions)

- ~~`hostprofile` auto-planner~~ → replaced by item 3 (D-v28-6)
- ~~truly-cold bench variant~~ → replaced by spike (a); optional one-shot post-M257 (D-v28-8)
- ~~§8.5 prose retraction~~ → moved to M257 so `frontend-tier.md` moves once (D-v28-10)
