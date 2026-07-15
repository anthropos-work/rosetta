# iter-02 — Decisions

## D5 — Measure the response BODY, not the response headers (the streaming-SSR trap)

**What happened.** The first baseline reported the SSR document arriving in **120 ms (HTTP 200)** and then
**37 seconds of unexplained client-side nothing** before the first data-query. Taken at face value, that
*refutes* iter-01's root cause (which puts the cost in a **blocked SSR**) and would have sent iter-03 hunting a
phantom client-side gate.

**Why it was wrong.** Next.js App Router **streams** the RSC payload. The shell flushes immediately — headers,
HTTP 200 — while the server render is still **blocked** awaiting its data. **Playwright's `response` event fires
on HEADERS.** A headers-only probe therefore reports a *fast document* while the body trickles for 37 s.

**Decision.** Every leg also records **`bodyAtMs`** via `response.finished()`, and a new anomaly kind
**`slow-body`** fires when headers were fast but the body was not. That single change moved 37.5 s from
"unattributed" onto the leg that owns it, and confirmed iter-01's prediction to within **33 ms**.

**Generalization (→ protocol doc + `latency-budget.md`).** *Any* latency probe against a streaming SSR framework
that watches `response` and not `response.finished()` will mis-attribute a blocked render as a client-side gap.
This is a trap for the next person, not a quirk of this milestone.

## D6 — The green gate travels to remote stacks by reading the REAL remote verdict, never by bypass

M217 shipped `autoverify.json` so M218 would never measure a broken stack. But the harness runs from the
**presenter's** machine (the tailnet) while the file lives on the **demo host** — so a naive local-path gate would
simply not find it, and the tempting move is `LATENCY_NO_GREEN_GATE=1`.

**Decision.** `run-latency.sh` takes `LATENCY_AUTOVERIFY_JSON` — point it at a **copy of the real remote file**
(`scp devops@box:…/autoverify.json`). The gate still grades **the real stack's real verdict**; it is *not* a
bypass. `LATENCY_NO_GREEN_GATE=1` remains, but is documented as "genuinely skip the gate" and is never used for a
gate-grading run. **A safety gate that is inconvenient in the exact situation it exists for will be switched off —
so make it work there instead.**
