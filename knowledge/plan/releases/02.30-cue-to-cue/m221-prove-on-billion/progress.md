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
| 05 | tik | Phase C cycle 2 — landed F1/F5/F5b (RED-fenced, r3), the live cycle-1 FALSIFIED r3 on the box (skills=0, directus schema err — empty-subdir + detached-supervisor residuals), **corrected r3→r4 in-tik**, and the r4 re-prove ≈ MET the gate on ONE cold cycle. rext `cue-to-cue-m221-r4`. | ≈8/8 on one r4 cycle — PROVISIONAL (reproducibility + browser Dana grade remained) | closed-fixed |
| 06 | tik | **FINAL live demo cycle** — ONE DEFAULT cold no-flag `up-injected.sh 1` at r4 on `billion`, browser-graded from the peer (incl. Dana's `/ai-readiness` 900-char floor), **LEFT LIVE** (not torn down). F1 cascade fixed live (taxonomy 0→42,790, all 3 catalog surfaces replayed, directus self-unblocked). M219 readiness all MET; F10 fully field-exercised (assert_ports_free + freshness-abort); F4 confirmed known-gap + F-M221-06b (run-latency http-cockpit-scheme) routed forward. | **8/8 gate MET** (maya p95 2.11 s / dan 1.31 s; Dana browser >>900; Ben STARTED; Aria COMPLETED) — **Gate: MET** | closed-fixed |

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

- iter-05 (tik): F1/F5/F5b landed; r3 cycle-1 falsified on the box → corrected r3→r4 in-tik → **r4 re-prove ≈ MET the gate on ONE cold cycle** (Maya p95 2.29s / Dan 1.65s, all 3 catalog surfaces, 3 orgs, Ben STARTED, Aria COMPLETED, no-flag) — PROVISIONAL (reproducibility remains) — see iter-05/progress.md

- iter-06 (tik) — **GATE MET; the milestone's exit gate is reached.** The FINAL DEFAULT cold no-flag r4 cycle on
  `billion` came up GREEN and **8/8** conditions held, browser-graded from the tailnet peer (maya p95 2.11 s /
  dan 1.31 s; Dana `/ai-readiness` filled >>900 ch; Ben STARTED; Aria COMPLETED); M219 readiness all MET; F10
  fully field-exercised. The stack is **LEFT LIVE** as the final demo (cockpit `https://billion.taildc510.ts.net:17700`,
  app `:13000`). **No further tik is required** for the gate. Residuals routed forward (Fate-3, next
  session/harden or close-milestone): **F-M221-06b** (`run-latency.sh` http-cockpit-scheme fix), **F4** (academy
  client-side grid render defect), and the earlier-carried `BURNIN-M221-dev-public-host` / `F-M220-4` /
  `PROBE-M218-c3-rerun`. Next lifecycle step: `/developer-kit:harden-mstone-iters --final` then
  `/developer-kit:close-milestone`.
