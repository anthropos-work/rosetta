# M258 iter-03 — decisions

## D8 — the inherited routing was implemented as *corrected*, not as *written*

`progress.md` routed `FIX-M258-iter02-inject-appends-and-swallows` with a named mechanism. The mandated
Step-0 re-survey (`TOK-01` known-context #6) tested it against code and found **all three of its factual
claims false** — the key is Clerkenstein-minted not foreign, the stack was public-host not localhost, and
the `|| true` is a deliberate `set -e` guard with a fail-loud check on the next line, not a swallow.

Two ways to be wrong here, and both were live:

1. **Implement it as written.** Removing the `|| true` would have deleted the guard that keeps
   `inject.py`'s failure path loud — a *regression*, shipped as a fix, in the name of a defect that was
   not there.
2. **Stop at the first plausible cause.** iter-02's own lesson. The `dockerignore` item
   (`FIX-M257-dockerignore-env-pattern-unpaired`) is real, does live in this file's neighbourhood, and
   *would* explain "a real Clerk key in an image" — it is simply not what happened.

So the routed item is discharged in **substance** (the append is fixed) while its **stated cause** is
retracted in the open, and the two items it missed — the first-wins reader and the missing host-mode
flag — are the ones that actually gated the milestone's deliverable.

**Recorded because the release keeps paying for this class:** *a routing is a hypothesis with provenance,
not a work order.*

## D9 — no timing was taken, and that is the instrument's verdict, not my preference

`assert-headroom --profile macmini` FAILED on `peak_load1` (20.31, then 39.05 and 45.94, against 12
cores). The load was **third-party**: Spotlight `mds_stores` at 183 %, the user's `a8-cart-runner`, node
processes; the hot Python was 3.12 while mine is 3.14. The one contender I had created — a background
`stack-core` census sweep — I stopped and verified gone; load kept climbing.

I did not run the cycle. `D-M255-1` already settles it: *a gate number measured on a host without headroom
is not a number*, and `rep_is_ok` books a headroom-failing rep as **not gateable** — so the cycle would
have burned ~10 minutes of three-lane Docker build to produce a figure the harness itself discards.
Second, and independently: the user is **actively working on this machine**, and three parallel build
lanes is not a neutral act.

What I did instead is the part contention cannot corrupt: the **live ISOLATION verdict** on the running
stack (8 images, `ok: True`, 0 failures) and the **real-artifact replay** of the env rewrite. Both are
booleans; both are decisive; neither depends on the clock.

**This is not iter-02's `D5` repeated.** That was a refusal to run a suite that would have measured a
broken wiring. This is a refusal to time anything on a box under someone else's load — and unlike `D5`,
the blocker is **not a defect and not mine to fix**.

## D10 — `demo-1`'s quarantine is lifted as to its stated reason

iter-02 `D7` left `demo-1` UP with *"must not be browsed — its UI tier would talk to a real Clerk app."*
**That premise is refuted** (D8). Its images carry a Clerkenstein-minted key for
`marcos-mac-mini.taildc510.ts.net`, and the live ISOLATION assert now returns `ok: True` over all 8
images with `foreign_pks {}`.

Two corrections follow, and the second is the one that matters:

- **There is no production-auth exposure.** There never was.
- **But `demo-1` IS reachable on the tailnet**, which `D7` denied — it was brought up with an
  **auto-discovered** public host, binds `0.0.0.0`, and holds a real Let's Encrypt cert
  (`cycle.log:26`, `:143-144`). That is the **documented** demo posture, not a violation:
  `safety.md` Part 3 states a demo is unauthenticated, authz-weakened, and published on all interfaces,
  with the tailnet as the control. Nothing here needs escalation — but a record that says
  "localhost-bound" about a tailnet-published stack is the kind of error the exposure docs are fenced
  against, so it is corrected explicitly rather than quietly.

`demo-1` is left UP: it is iter-04's reproduction and, per `overview.md` § *Batch-gate behaviour*, the
stack is left UP regardless. The user's `demo-2` (11 containers) and 5-container dev stack were verified
resident before and after, and were never touched.
