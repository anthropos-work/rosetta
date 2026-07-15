# M221 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| 01 | tok/bootstrap | authored TOK-01 (fix-then-prove, host-isolation first) | N/A (tok) | continue |
| 02 | tik | Phase A: landed `GUARD-M221-host-isolation` — a per-N host lock (`hostlock.sh`) wired into `up-injected.sh` + `rosetta-demo down`/`up`; rext `cue-to-cue-m221-r1` | sub-gate MET: 4 DoD fences RED-proven (18 failed pre-fix → 18 passed) | closed-fixed |
| 03 | tik | Phase B: 3 off-box demo-hygiene fixes + fences — (1) academy binds `-H 127.0.0.1` on localhost (F-M220-5) + `exposure_claim_guard` extended to host-native listeners; (2) fenced the `ant-academy.sh` pre-bind reap (shipped M217+M220, ran in 0 tests); (3) new `backend_api_url_server_reader_guard` fences the F-7 server-side blackhole. rext `cue-to-cue-m221-r2` | sub-gate MET: 3 fences RED-proven pre-fix; suites 892+152 green; M220 S3 HARD INVARIANT undisturbed | closed-fixed |

## Next iter
_(set by the previous iter's closeout)_

- iter-04: **Phase C — the live cold-reset-to-seed battery on `billion`** (a DEFAULT `/demo-up N`, NO FLAGS,
  from a tailnet peer): prove all 8 gate conditions + fold in M219's five readiness gates re-proven at final
  code. Prerequisites now in place: `GUARD-M221-host-isolation` (iter-02) + the Phase B fixes (iter-03), so the
  battery grades the shipped code and cannot corrupt its own evidence. Still-off-box carries not in iter-03
  (`FIX-M221-academy-empty-catalog`, `F-M220-4`, `BURNIN-M221-dev-public-host`, `PROBE-M218-c3-rerun`) remain.
