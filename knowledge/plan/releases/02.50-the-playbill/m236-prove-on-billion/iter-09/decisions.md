# iter-09 — decisions

## D1 — measure from the tailnet, grading the real stack's own verdict

The gate component is a **presenter-vantage** number, so the harness runs from this workstation across the
tailnet against `billion`, not on `billion` itself. That puts the stack's `autoverify.json` on the other
host, which the green gate reads locally.

**Decision.** Use the seam the runner already documents (`LATENCY_AUTOVERIFY_JSON` — "point it at a local
copy of the REMOTE file"), after writing a **fresh** verdict on `billion`
(`STACK_DIR=… stack-verify/live/autoverify.sh --project demo-1 --offset 10000` → all probes ✓).

**Why not `LATENCY_NO_GREEN_GATE=1`.** That override exists, and using it would have been faster and
worthless: the gate's entire purpose is that a latency number off an unverified stack is noise. Copying the
real verdict keeps the grading honest; skipping it would have produced a number with no standing.

**Note for the next author:** `autoverify.sh` only writes the verdict when `STACK_DIR` is set in the
environment. Invoked without it, it prints a full green report and writes nothing — which looks exactly
like success.

## D2 — both hero p95 are inside budget; the environment is part of the number

| vantage | reached ACCESS | p50 | p95 | gate |
|---|---|---|---|---|
| employee (`maya-thriving`) | 5/5 | 0.90 s | **3.15 s** | < 5 s ✅ |
| manager (`dan-manager`) | 5/5 | 0.77 s | **2.71 s** | < 5 s ✅ |

Repeat batteries: employee 1.67 / 2.85 / 3.15 s; manager 2.90 / 2.71 s. Every battery inside budget.

**Scope.** Hero vantages only, per user decision B2 — content-seat latency is explicitly out of scope for
v2.5, and the 31→29 content actions are proven for CONTENT, not formally timed.

**Conditions stated, per `latency-budget.md`.** Tailnet-remote, from a second machine, against a **warm**
stack on `billion` over HTTPS `--public-host` origins. Not comparable to M218's laptop-local cold baseline
(2413 / 1767 ms), and not offered as such. The cold-cycle reproduction is a separate handler.

## D3 — side-deliverable: the green gate's age check failed OPEN west of UTC

**Evidence.** The gate reported `verdict age 7264s` for a verdict written 121 s earlier, with both hosts
reporting identical UTC. `autoverify.sh` writes `ts` as UTC with a trailing `Z`; the BSD fallback
`date -jf '%Y-%m-%dT%H:%M:%S' "${v_ts%%Z*}"` parses that string in the **local** zone.

```
BSD parse local: 1784534312 → age 7321s
BSD parse UTC  : 1784541512 → age  121s
```

**Decision.** `TZ=UTC` on the BSD branch. Separate commit + tag `playbill-m236-latency-tz-fix`; recorded as
a side-deliverable so it does not grade this iter's planned scope.

**Why it was worth stopping for.** The error is signed. East of UTC it shrinks the freshness window and
fails *closed* (annoying, visible). West of UTC it inflates the window and reads a **stale verdict as
fresh** — which is the exact F-6 hazard ("a nine-hour-old verdict grading a Clerkenstein-dewired stack
green") that the age check was written to prevent. A guard that fails open is worse than no guard, because
downstream trusts it. The check was itself introduced by an M218 hardening pass, so hardening code is not
exempt from the hazard it hardens against.
