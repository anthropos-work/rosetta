# iter-02 (tik) — progress

**Type:** tik, under TOK-01 (Phase A). Lands `GUARD-M221-host-isolation` — the prerequisite for the
multi-cycle battery that IS M221's gate.

## What changed
- **`rext demo-stack/hostlock.sh`** (new) — a per-N host lock. Atomic take via `set -C` (O_EXCL) at
  `~/.rosetta-demo-locks/demo-<N>.lock` (outside any `$STACK` → `down --purge` can't clobber it). Identity
  block (hostname/pid/epoch/iso/label/token). A second concurrent cycle FAILS LOUD naming the holder — never
  queues, never proceeds. Stale-tolerant-but-LOUD reclaim (age > TTL **AND** same-host holder provably gone by
  `kill -0`; the reclaim logs the dead holder). Released on normal completion + trap exit (EXIT/INT/TERM/HUP);
  `kill -9` falls back to stale-reclaim. Re-entrant; identity-checked release never deletes a foreign lock.
- **`rext demo-stack/up-injected.sh`** — sources `hostlock.sh`; takes the lock as the FIRST real gate (before
  the clone/build). ONE consolidated `_up_injected_atexit` handler owns the release + ssh-agent reap; the old
  bare `trap reap_ssh_agent EXIT` inside `ensure_ssh_agent` (which would have CLOBBERED the release — bash
  keeps one EXIT trap) is removed and folded in.
- **`rext demo-stack/rosetta-demo`** — sources `hostlock.sh`; `cmd_down` takes the lock BEFORE
  `clear_stack_verdict`/purge (the exact M219 hazard: a teardown purging out from under a live measurement);
  `cmd_up` takes it once N is allocated.
- **`rext demo-stack/tests/test_host_isolation.py`** (new) — 18 fences.

## Fences (RED-proven — the release's core discipline)
Ran the WHOLE suite against the **pre-fix tree** (hostlock.sh hidden + the two wirings git-stashed):
**18 failed.** Against the fixed tree: **18 passed.** So every fence is RED pre-fix, GREEN after — none is
theatre. Additionally mutation-proved: a `HOSTLOCK_LIB`-style seam points the same behavioural fences at four
targeted mutants (check-then-write take · generic refusal · silent reclaim · no-rm release) and asserts each
is CAUGHT.

The four DoD fences, each ✅:
1. **Second concurrent cycle REFUSED, naming the holder** — `rc=3`, message carries the holder's label+pid+iso
   (not a generic "locked"); mutant `generic_refuse` caught.
2. **Released on normal completion AND on SIGTERM** — no permanent self-deadlock; mutant `no_rm_release` caught.
3. **Stale reclaim is LOUD** — dead/absent-pid holder reclaimed with the dead holder logged; a LIVE holder is
   never reclaimed even at `TTL=0`; mutant `silent_reclaim` caught.
4. **Atomicity** — static (source uses `set -C`, not `[ -e ] && `) + dynamic (24 simultaneous acquirers →
   exactly ONE winner); mutant `check_then_write` shown to race (>1 winner).

## Regression
`test_reap.py + test_purge.py + test_frontend_build.py` (the suites over the two edited files): **125 passed,
1 skipped.** The `UP_INJECTED_LIB_ONLY=1` seam still returns at line 785 and takes **no** lock (the acquire is
below the seam) — verified: no lockfile created by a lib-only source.

## rext roll
Committed `3893c81` on `main`, tagged **`cue-to-cue-m221-r1`**, pushed to `anthropos-work/rosetta-extensions`.

## Off-box
No live `billion` cycle was needed — the take/refuse/release/reclaim is proven with the mechanism directly +
unit tests (per the operating rules: prefer that over a 30-minute live cycle). The guard will first be
*exercised in the field* by Phase C's battery, which now cannot corrupt its own evidence.

## Close — 2026-07-15

**Outcome:** Landed `GUARD-M221-host-isolation` — a per-N, stale-tolerant, identity-named host lock wired into
the demo bring-up (`up-injected.sh`) and teardown (`rosetta-demo down`/`up`). A second concurrent cycle on the
same host+N now fails loud naming the holder; a crashed cycle self-releases (trap) or is reclaimable (loud,
never silent). The prerequisite for M221's own multi-cycle battery is in place — the battery can no longer
corrupt its own evidence the way M219's was.
**Type:** tik
**Status:** closed-fixed
**Gate:** sub-gate MET — the 4 DoD fences RED-proven pre-fix (18 failed) and GREEN after (18 passed). The
milestone's 8-condition gate is **not** measured by this tik (it is the prerequisite the battery runs under);
N/A for the milestone gate this iter.
**Phase 5 grading:** (1) gate-met: n (milestone gate not measured by this tik — this is its prerequisite) —
(2) triggered-tok: n (first tik, made progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n
(1 tik) — (6) protocol-stop: n — **Outcome: continue** (Phase A done → next iter opens Phase B).
**Decisions:** D-M221-02a..e (iter-02/decisions.md).
**Side-deliverables:** removed a latent EXIT-trap-clobber in `up-injected.sh` (`ensure_ssh_agent`'s bare
`trap reap_ssh_agent EXIT` would have silently orphaned any earlier EXIT trap on billion) — consolidated into
`_up_injected_atexit`.
**Routes carried forward:** none new. Phase B (off-box code fixes) and Phase C (the live battery) remain, per
TOK-01.
**Lessons:** the lock is per-(host, N), not per-host — the M219 corruption was two agents on the *same* N; the
unified registry already isolates *different* N. And "the file exists" ≠ "a live cycle holds it": the guard
must distinguish a live holder from a dead one (same-host `kill -0`), or it becomes its own deadlock hazard.

## Next iter
- iter-03 (tik): **Phase B — the off-box code fixes** (academy loopback-bind + extend `exposure_claim_guard`
  to host-native listeners; reap the native cockpit+academy; the academy empty-catalog patch; the
  backend-api-url twin fence). RED-prove each pre-fix.
