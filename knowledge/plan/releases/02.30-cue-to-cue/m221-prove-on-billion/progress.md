# M221 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| 01 | tok/bootstrap | authored TOK-01 (fix-then-prove, host-isolation first) | N/A (tok) | continue |
| 02 | tik | Phase A: landed `GUARD-M221-host-isolation` — a per-N host lock (`hostlock.sh`) wired into `up-injected.sh` + `rosetta-demo down`/`up`; rext `cue-to-cue-m221-r1` | sub-gate MET: 4 DoD fences RED-proven (18 failed pre-fix → 18 passed) | closed-fixed |

## Next iter
_(set by the previous iter's closeout)_

- iter-03 (tik): **Phase B — the off-box code fixes** (academy loopback-bind + extend `exposure_claim_guard`
  to host-native listeners; reap the native cockpit+academy; the academy empty-catalog patch; the
  backend-api-url twin fence). RED-prove each pre-fix.
