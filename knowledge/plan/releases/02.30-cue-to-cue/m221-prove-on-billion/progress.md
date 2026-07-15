# M221 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| 01 | tok/bootstrap | authored TOK-01 (fix-then-prove, host-isolation first) | N/A (tok) | continue |
| 02 | tik | Phase A: landed `GUARD-M221-host-isolation` — a per-N host lock (`hostlock.sh`) wired into `up-injected.sh` + `rosetta-demo down`/`up`; rext `cue-to-cue-m221-r1` | sub-gate MET: 4 DoD fences RED-proven (18 failed pre-fix → 18 passed) | closed-fixed |
| 03 | tik | Phase B: 3 off-box demo-hygiene fixes + fences — (1) academy binds `-H 127.0.0.1` on localhost (F-M220-5) + `exposure_claim_guard` extended to host-native listeners; (2) fenced the `ant-academy.sh` pre-bind reap (shipped M217+M220, ran in 0 tests); (3) new `backend_api_url_server_reader_guard` fences the F-7 server-side blackhole. rext `cue-to-cue-m221-r2` | sub-gate MET: 3 fences RED-proven pre-fix; suites 892+152 green; M220 S3 HARD INVARIANT undisturbed | closed-fixed |
| 04 | tik | Phase C cycle 1 — FIRST live battery on `billion` (DEFAULT `up-injected.sh 1`, NO FLAGS, cold reset-to-seed, driven from a tailnet peer). Baseline + diagnose. IRONCLAD-confirmed the dominant root cause (snapshot store-root bug: `public.skills 0→42,790` with the store pinned); re-characterised the academy carry (client-side render defect, not empty local-content); field-confirmed the native-academy teardown reap is still broken. No code fix landed (both candidates routed, de-risked). | **3/8 MET** (1 latency p95 2.23/2.08 s, 3 orgs, 7 remote-by-default) · 4 NOT MET (2,4,5,6 — ONE root cause: taxonomy cache-skip) · 1 at-risk (8) | closed-no-lift |

## Next iter
_(set by the previous iter's closeout)_

- iter-05: **Phase C cycle 2 — land F1 + F2, then re-cycle.** (1) **F1** — pin `STACKSNAP_STORE`/`--store` to the
  workspace-level cache root in `dev-setdress.sh:323` (the walk-up resolves a shadowing empty
  `stack-demo/rosetta-extensions/.agentspace`), RED-fence the store-root divergence, re-tag rext. **De-risked:
  iter-04 proved the fix works** (`public.skills 0→42,790` with the store pinned). (2) **F2** — re-capture the
  directus surface (genuine digest divergence). Then a fresh cold no-flag battery on `billion` and re-measure
  gates 2/4/5/6 (F1 cascade-unblocks all four). Also route-forward this cycle: **F5** native-academy teardown reap
  (port+identity, not pidfile-only) + **F5b** gate-8 generated-file dirt; **F4** academy render defect; **F7** c-3
  router-403 re-check (with F2); **F10** freshness-abort + `assert_ports_free` field-exercise. Not yet reached:
  `BURNIN-M221-dev-public-host`, `F-M220-4`. Full baseline capture: `iter-04/findings.md`.

- iter-05 (tik): landed F1/F5/F5b (rext r3, RED-fenced) but the live re-prove FALSIFIED F1 — skills still 0 on billion, gate ~3/8; re-diagnosis → iter-06 — see iter-05/progress.md
