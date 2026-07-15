---
iter: 2
milestone: M221
iteration_type: tik
iter_shape: tik
status: open
opened: 2026-07-15
strategy_ref: TOK-01 (../decisions.md)
---

# iter-02 (tik) — land `GUARD-M221-host-isolation` (Phase A)

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). This is the **first tik** and it lands
the **prerequisite** for everything after it: a mutual-exclusion lock so **two agents cannot run demo cycles
against the single `billion` host concurrently**. Until this holds, M221's own gate — a multi-cycle cold
battery on that one host — can corrupt its own evidence exactly as M219's did (D17: a cycle purged the stack
mid-measurement, a gate went UNEXECUTED, and an unexecuted gate is a FINDING, not a pass).

## The DoD (from the overview, `GUARD-M221-host-isolation`)
A cycle **cannot start** while another holds the host. A **host lock**: an advertised, stale-tolerant lockfile
on the demo host, taken for the life of a cycle and named with the holder's identity. A second concurrent
cycle **fails loud** with the holder's identity — never queues silently, never proceeds.

## Chosen shape — a per-N host lock (the lighter option)
- A lockfile at **`~/.rosetta-demo-locks/demo-<N>.lock`** (outside `$STACK`, so `down --purge` — which wipes
  `$STACK/data` — cannot clobber it; outside the ephemeral clone, so a re-clone cannot either).
- Created **atomically** via `set -C` (noclobber → `O_EXCL`), **not** check-then-write. The winner writes an
  identity block: `holder` (hostname) · `pid` · `epoch` · `iso` · `label` · `token`.
- Taken at the **START** of a bring-up (`up-injected.sh`) and a teardown (`rosetta-demo down`) for stack N;
  released at the END and on trap exit (EXIT + INT/TERM/HUP).
- **Stale-tolerant, loud:** a lock older than a generous TTL **AND** whose holder is provably gone
  (same-host `kill -0` says the pid is dead) is reclaimed — the reclaim **logs the dead holder** (never
  silent). A holder that is alive, or on a different host (liveness not provable from here), is **never**
  reclaimed → refused.
- A second concurrent attempt on the same N **exits non-zero, loudly, naming the current holder** — no queue,
  no wait, no proceed.

Keyed per-(host, N): on `billion` two agents on the same N collide (the M219 hazard) and are excluded; on a
localhost dev box a single cycle takes+releases cleanly and is otherwise unaffected. The dev-stack tooling is
untouched — the **demo** cycle is the target.

## Why this and not per-cycle N
The overview offers "or a per-cycle stack N." A host lock is lighter and directly names the holder in the
refusal (the DoD's headline requirement), and it also protects the **teardown** path (two agents, one purges
mid-measure — exactly M219's failure). Per-cycle-N isolation is already provided by the unified registry for
*allocation*, but it does not stop two agents from driving the *same* N, which is what actually happened.

## The fences (RED-proven — non-negotiable)
Homed in `demo-stack/tests/test_host_isolation.py`. Each must FAIL against pre-fix / mutant code and PASS
after. Mutation-proved via a `HOSTLOCK_LIB` seam (point the same fences at a buggy check-then-write mutant).
1. A second concurrent cycle is **REFUSED**, and the refusal **names the holder's identity** (label + pid),
   not a generic "locked".
2. The lock is **released** on normal completion AND on a killed (SIGTERM) exit — no permanent self-deadlock.
3. **Stale reclaim is loud** — a dead/absent holder is reclaimable and the reclaim logs the dead holder;
   a LIVE holder is never reclaimed.
4. **Atomicity** — the take is a real atomic primitive (`set -C`/noclobber, not `[ -e ] && ...`), proven both
   statically (source assertion) and dynamically (a concurrent-race: exactly one of many simultaneous
   acquirers wins).

## Out of scope (this tik)
Phase B off-box code fixes and Phase C's live battery. This tik does **not** require a live `billion` cycle —
the lock's take/refuse/release/reclaim is proven with the mechanism directly + unit tests (per the operating
rules: prefer that over a 30-minute live cycle).

## Distance to gate
The milestone gate (the 8-condition battery) is unmeasured on this code. This tik moves the gate **not at all
directly** — it is the guard that makes the battery's future measurement *trustworthy*. Its own sub-gate: the
four fences RED-proven pre-fix and GREEN after.
