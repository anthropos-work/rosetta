---
milestone: M257x
iter: 09
---

# iter-09 — progress

**Type:** tik (under `TOK-01`, step 5 — *prove it cold*)

## Root cause, proven in both directions

The academy was **not** a probe defect. **A security tightening broke it, and the guard that verified the
tightening only ever measured the security property.**

`demo-stack/ant-academy.sh:466` (M221 / F-M220-5):

    bind_args=(-H 127.0.0.1); [ -n "${STACK_PUBLIC_HOST:-}" ] && bind_args=(-H 0.0.0.0)

M221 tightened the localhost path's bind from `0.0.0.0` to `127.0.0.1` — a real fix for a real problem (the
academy was answering HTTP 200 on the tailnet IP of a *localhost* demo: "the S0 exposure lie"). Its own
comment calls it *"a strict de-exposure, not a flip"*.

**It is not strict.** Next.js 16's dev server runs an internal **proxy layer that dials
`http://localhost:$PORT/` to reach the app it just started** — `localhost`, not the address it was told to
bind. On a host where `localhost` resolves to `::1` first, the server is listening on IPv4 loopback only,
the self-proxy hits IPv6 loopback, and the request hangs until it dies.

Measured on this box, same code, same `.env.local`, only the bind changed:

    -H 127.0.0.1   (M221, current)   GET /  ->  500   in 30.014 s   [reproduced twice, deterministic]
    -H 0.0.0.0     (pre-M221)        GET /  ->  200   in  2.378 s

and the server's own log names the mechanism:

    Failed to proxy http://localhost:13077/ Error: socket hang up { code: 'ECONNRESET' }

**The 30.0 s is the tell.** It is flat, repeatable, and identical on a warm second request — so it is a
fixed internal timeout, not a Turbopack cold compile. The launcher's own comment had reasoned the opposite
(*"`next dev` cold-compiles… on a slow VM the first paint genuinely takes ~30-60 s, so a short default would
reintroduce the false-negative"*) and set a 120 s budget to accommodate it. A longer wait could never have
helped: every attempt fails the same way, forever.

Two secondary observations, both real:

- The readiness probe polls `/` with `curl -fsS --max-time 3`. `/` is the one route that 500s; **`/library`
  answers `308` in 2 ms even while the stack is "not serving."** So the probe's verdict was correct here,
  but for a reason nobody had measured — and its per-attempt 3 s could never observe a 30 s failure anyway.
- `--port` is passed **twice** (`next dev --port 3077 --port 13077`). Benign (last wins, and Next reports
  13077), but it means the port is being set in two places.

## And the *other* open route was resolved on the way

`CHECK-M257x-bringup-evidence-logs-absent` is **an autoverify path bug, not missing evidence.** autoverify
looks for `$STACK_DIR/demopatch.log` and `$STACK_DIR/buildfail.log`. The bring-up writes them to the
**per-stack** directory:

    stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/{demopatch,buildfail,ant-academy,build-*}.log

Both exist. And autoverify's message asserts a cause from the absence — *"its ABSENCE means the phase never
ran"* — which is a conclusion drawn from looking in the wrong place. **The milestone's dominant class once
more, this time in the verifier that measures the gate.**

There is a second layer worth not missing: at the correct path **both files are 0 bytes** (created 16:46,
never written) while the demo-patches demonstrably applied — their output went to stdout. So even a
path-fixed check would need to distinguish *absent* / *empty* / *populated*, which are three states the
current message collapses into one.

## Why this closed without a fix

The escalation condition in this iter's `overview.md` fired exactly as written: the cause is a **deliberate
security fix**, and the naive repair — putting `0.0.0.0` back — **re-opens the exposure M221 closed**, on the
one axis `safety.md` §3 Part 3 treats as load-bearing. Trading a proven security property for a green check
is precisely the move this milestone exists to stop, so it was not made.

The correct fix has to keep the loopback tightening *and* satisfy a self-proxy that dials `localhost`, and
it has to be re-proved on **both** host families (this Mac and the `billion` Linux VM resolve `localhost`
differently, which is the whole mechanism). That is a real design with a real exposure guard to re-run —
not a one-line edit — and it is pre-computed for iter-10 rather than half-landed here.

**Nothing was changed in this iter.** The deliverable is the characterization, and it converts a four-iter-old
symptom (*"a readiness probe disagreeing with the process"*) into a named mechanism with a measurement on
both sides.

## Close — 2026-07-31

**Outcome:** root cause of `FIX-M257x-academy-not-serving` found and proven both ways — **M221's
`-H 127.0.0.1` de-exposure breaks Next 16's `localhost`-dialing self-proxy** (`500` in a flat 30.0 s vs
`200` in 2.4 s on the pre-M221 bind). No fix landed: the repair must preserve the security property, which
is a design, not an edit. `CHECK-M257x-bringup-evidence-logs-absent` resolved as an **autoverify wrong-path
bug** (the logs exist, per-stack).
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-13` (a de-exposure is not proven by an exposure guard alone)
**Side-deliverables:** none (no code changed).
**Routes carried forward:** `FIX-M257x-academy-not-serving` (now with a proven mechanism + a pre-computed fix) · `FIX-M257x-autoverify-evidence-log-path` (new, replaces `CHECK-M257x-bringup-evidence-logs-absent`)
**Lessons:**
- **A de-exposure is not proven by an exposure guard.** M221 asked *"does it still answer where it
  shouldn't?"*, got the right answer, and shipped. Nobody asked *"does it still answer where it should?"* —
  and the exposure guard, by construction, can never notice that the service stopped working. **Every
  tightening needs a paired liveness assert, or the guard is only measuring half the change.**
- **A flat, repeatable duration is a timeout, not work.** 30.014 s then 30.007 s, identical warm, is a
  configured limit. The launcher's comment had attributed exactly this shape to cold compilation and
  budgeted 120 s for it — a wait that could never succeed. Compare the *variance*, not just the magnitude.
  (Kin to `latency-budget.md`'s arithmetic signatures, which name a bug class before you read any code.)
- **A "not serving" verdict deserves one patient request before it is believed.** Three probes at
  `--max-time 3` and one at `--max-time 180` are different instruments; only the second could see this.
