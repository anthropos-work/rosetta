# iter-09 — tik: the p95 click→ACCESS measurement (hero vantages)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove), Phase H — the budget half.

## Step 1 — the green gate, and a remote-measurement seam that already existed

`run-latency.sh` refuses to measure a stack whose `autoverify.json` is not green (the M217→M218 barrier).
The stack is on `billion`; the harness must run **from a machine on the tailnet** — the presenter's actual
vantage — so the verdict file is on the wrong host.

No tooling change was needed: the seam is already there and documented for exactly this case
(`LATENCY_AUTOVERIFY_JSON`, "point it at a local copy of the REMOTE file"). Wrote a fresh verdict on
`billion` (`STACK_DIR=… autoverify.sh --project demo-1`) and pointed the gate at a fetched copy — so the
gate still grades **the real stack's own green report**, not a local stand-in.

```
✓ backend /api/health 200 on :18082
✓ sentinel.casbin_rules = 1250 (authz policy loaded)
✓ verify live: all liveness + readiness probes passed
✓ taxonomy replayed: public.skills = 42790
▶ autoverify demo-1: OK — verified-working.
```

## Steps 2–3 — the measurement

Five consecutive runs per vantage, driven **from this workstation over the tailnet** against `billion`
(a warm stack, `--public-host` HTTPS origins, `LATENCY_SCHEME=https`).

| vantage | seat | reached ACCESS | p50 | p95 | gate |
|---|---|---|---|---|---|
| employee | `maya-thriving` | **5/5** | 0.90 s | **3.15 s** | < 5.0 s ✅ |
| manager | `dan-manager` | **5/5** | 0.77 s | **2.71 s** | < 5.0 s ✅ |

Repeat batteries were consistent (employee p95 across three batteries: 1.67 / 2.85 / 3.15 s; manager:
2.90 / 2.71 s) — the variance is real but comfortably inside budget in every battery.

**The environment is part of the number**, per `latency-budget.md`'s standing instruction: these are
*tailnet-remote, warm-stack* readings taken from a second machine, which is the presenter's real vantage.
They are not comparable to M218's laptop-local baseline (2413 / 1767 ms) and are not offered as such.

The `net::ERR_ABORTED` entries in the anomaly lists are Next.js **RSC prefetches cancelled on navigation** —
expected client behaviour, not failed requests. Every measured leg returned 200/303/307.

**iter-08's D5 lead was checked and did not manifest**: every `data-query` leg resolved 200 within ~1.1 s,
so the client is not hitting the dead `localhost:5050` address on these paths. The finding stays recorded
against the release rather than this gate — SSR uses `WUNDERGRAPH_SSR_ENDPOINT` and these vantages never
exercised the mis-addressed client path.

## Side-deliverable — the green gate was aging verdicts wrong on macOS

The gate reported **age 7264 s** for a verdict written **121 s** earlier, while both clocks read identical
UTC. `autoverify.sh` writes `ts` in UTC (trailing `Z`); the BSD fallback `date -jf` parses its input in the
**local** zone, so the age was off by exactly the UTC offset (7200 s on this UTC+2 grader).

```
BSD parse local: 1784534312   → age 7321s
BSD parse UTC  : 1784541512   → age  121s   ← correct
```

The direction of the error is what makes it worth fixing rather than noting: **east** of UTC it shrinks the
freshness window (past +4 it would refuse every fresh verdict); **west** of UTC it *inflates* the window and
reads a **stale verdict as fresh** — precisely the F-6 hazard the age check was added to prevent. The guard
failed open for half the world. `TZ=UTC` on the BSD branch; re-verified at 144 s.

Separate commit + tag (`playbill-m236-latency-tz-fix`), recorded as a side-deliverable — it does not grade
this iter's planned scope.

## Close — 2026-07-20

**Outcome:** Both hero vantages measured from the tailnet and **inside budget** — employee p95 **3.15 s**,
manager p95 **2.71 s**, 5/5 ACCESS each, against a 5 s gate. The p95 gate component is **MET**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (primary metric MET, p95 MET; the cold reset-to-seed reproduction remains)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (measure from the tailnet via the existing `LATENCY_AUTOVERIFY_JSON` seam — grade the real stack's own verdict, never a local stand-in), D2 (both hero p95 inside budget; state the environment with the number), D3 (side: the age check parsed UTC as local — failed OPEN west of UTC)
**Side-deliverables:** the `run-latency.sh` green-gate timezone fix (`playbill-m236-latency-tz-fix`) — a guard that was quietly wrong in the dangerous direction.
**Routes carried forward:**
- **Cold reset-to-seed reproduction** → handler `REPRO-M236-iterTBD-cold-cycle`. The last gate component. Re-pin `billion` to the final tag first; the cockpit currently binds `127.0.0.1` (iter-07) and a cold bring-up restores the normal ordering.
- **M230 cluster 3** (anonymous `/library` + `/free` render 0 cards) → `ACADEMY-M236-iter08-public-catalog-twin`, unchanged.
- **iter-08 D5** (`apps/web` client `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` points at the non-offset `localhost:5050`) — did not manifest on the measured paths; carry to release close as a demo-hygiene item rather than a latency item.
**Lessons:**
- **A freshness guard that fails OPEN is worse than no guard**, because it is trusted. This one had survived a hardening pass and was introduced *by* a hardening pass (M218 F-10), which is a reminder that the code added to close a hazard is not itself exempt from that hazard.
- **Check the instrument's arithmetic against a fact you already hold.** Two clocks agreeing while the tool reported a 2-hour age is what exposed it. `latency-budget.md` teaches reading arithmetic signatures to name a bug class before reading code; that applies to the measuring tool as readily as to the system under test.
- **The remote-measurement seam existed and was documented** — the first instinct was to build one. Reading the runner's own header comments saved that work. Third time this milestone that the answer was already written down.
