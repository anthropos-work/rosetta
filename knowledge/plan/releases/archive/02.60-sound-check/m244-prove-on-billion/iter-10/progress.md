**Type:** tik (cleanup shape) — inherited carry reap-17700 (M239-D13), under TOK-01. Run 4, tik 5 (cap).

# iter-10 — progress

## What landed (rext, committed + proven + pushed — b38ad75, tag → 7dbad4b)
`test_reap.py::TestReapHardening::test_a_RACED_listener_exits_silently` (the 9th demo-stack standing failure)
called `_reap_with_stubs(stubs)` with no port → the hardcoded default `port=17700`. reap.sh (correctly, per its
M217/M221 design) establishes real `/dev/tcp` occupancy BEFORE attribution, so on a box where a live demo-1
cockpit is actually listening on :17700 the port reads HELD, the stubbed attribution can't own it, and reap
refuses (returns 1) → the test false-fails. A clean box (17700 free) always passed. **0 product defect, 0
reap.sh defect — a pure test-isolation collision with ambient infrastructure** (the M239-D13 root cause).

**FIX (the D13 recipe):** default `_reap_with_stubs` to a guaranteed-FREE `_free_port()` (never 17700). Only the
RACED test relied on the default; the three explicit-port callers are unchanged.

## Proven load-bearing
- `test_reap.py`: **41/41 pass**.
- Isolation proof: with **:17700 HELD locally** (a simulated live-cockpit collision), the RACED test now
  **PASSES** (1 passed) — before the fix it would refuse (reap sees 17700 held → returncode 1 → test fails).

## Scope note
The other 8 durable standing failures (`test_cockpit.py` ×6 academy+overlay + `test_host_prereqs_m215` +
`test_purge`) are the M238-D5 standing-8 class (Fate-2 → M244), a DIFFERENT failure mode — out of this iter's
reap-17700/D13 scope. They remain routed.

## Close — 2026-07-22

**Outcome:** the inherited reap-17700 (M239-D13) test-isolation carry DISCHARGED — the RACED test is isolated
from ambient :17700 via a guaranteed-free port; proven load-bearing (passes with 17700 held); test_reap.py
41/41 green; rext pushed. 1 of the 3 inherited defers now dispositioned+landed (DEF-M240-01 dispositioned
iter-07; this lands reap-17700). Metric stays **3/8** (a cleanup carry, not a gate part).
**Type:** tik (cleanup shape)
**Status:** closed-fixed
**Gate:** NOT MET (milestone 3/8; this is an inherited-carry discharge, not a gate part)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (tik 5 of 5)** — (6) protocol-stop: n — Outcome: **exit-5 (cap-reached)**
**Decisions:** D1 (fix the default, not just the one call site — hardens against future forgetful callers), D2 (prove with 17700 held, since locally 17700 is free and the bug only manifests under collision) — iter-10/decisions.md
**Side-deliverables:** none.
**Routes carried forward:** none new (the durable standing-8 stay Fate-2 → M244; the gate-(d) /library fix + the big gate-(h) cold reset-to-seed stay routed from iters 08/09).
**Lessons:** (1) a test that hardcodes a well-known service port (17700) collides with a live instance of that
service — default such helpers to a `_free_port()` so tests are isolated from ambient infrastructure. (2) a
host-state-dependent test bug can't be proven fixed on a box that doesn't reproduce the state — SIMULATE the
state (bind the port) to prove the fix load-bearing.
