**Type:** tik  ·  **Active strategy:** TOK-01 (reachability-first)

# iter-02 — the latency harness + the honest pre-fix baseline

## Delivered (rext `stack-verify/e2e/`, commit `3fab10c`)

| file | what |
|---|---|
| `lib/latency.ts` | `measureLogin()` — clicks the **real** cockpit CTA, times **click→ACCESS**, attributes every leg, captures anomalies. `percentile()`. |
| `tests/latency.spec.ts` | N samples/vantage → p50/p95 → `latency-out/<vantage>.json`. **Measure-only** by default; `LATENCY_GATE_MS` turns it into the exit gate. |
| `run-latency.sh` | The runner. **REFUSES to measure a stack whose `autoverify.json` is not green.** |
| `lib/cockpit-login.ts` | `waitUntil` is now a **parameter** (default `'networkidle'` preserved) — resolves iter-01 **D4** without forking. |

**Design commitments honoured:** drives the **real** `<a class="btn login">` CTA off the live cockpit (so a stale
/ host-drifted cockpit **fails the probe** rather than being measured around — the M217 hazard); **never**
`networkidle`; a **new stack-verify surface**, not a Playthrough (perf is a declared Playthrough non-goal) —
which **discharges DEF-M215-03(b)**; **zero platform edits**.

---

## THE BASELINE
_billion · demo-1 · rext `cue-to-cue-m217` · `autoverify.json` **GREEN** · measured **from the tailnet**, which is
the presenter's actual vantage._

| vantage | p50 | **p95** | gate | verdict |
|---|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | 38.16 s | **39.45 s** | < 5 s | **7.9× OVER** |
| **manager** (`dan-manager` → `/enterprise/…`) | 38.28 s | **38.30 s** | < 5 s | **7.7× OVER** |

**3/3 runs reached ACCESS on both vantages** — so this is *slow*, not *broken*. Reproducible, low variance.

**Secondary, REPORTED-NOT-GATED (D-DESIGN-1):** data-settle p50 **44.09 s** (employee) / **40.71 s** (manager).

### Per-leg attribution — the entire defect is ONE leg

```
handshake  @0.03 s (303)      ← FREE, as iter-01 proved statically and F-1 measured
ssr-document headers @0.15 s  ← "fast"… and this nearly lied to us
clerk-js   @0.19–0.54 s (200) ← fine
ssr-document BODY   @37.62 s  ←──── THE WHOLE DEFECT
data-query @37.65 s (200)     ← the client's queries WORK — they just cannot start until SSR unblocks
ACCESS     @38.2 s
```

### The measurement trap this iter walked into (and out of)

The **first** baseline run reported: *document arrives in 120 ms, then 37 seconds of unexplained client-side
nothing.* That reading is **wrong**, and it would have sent iter-03 hunting a phantom client-side gate.

**Next.js App Router *streams* the RSC payload.** The shell flushes **immediately** (HTTP 200, headers in
70–150 ms) while the server render is still **blocked** awaiting its data. Playwright's `response` event fires on
**headers**. So a headers-only probe reports a *fast document* while the body is still trickling 37 s later.

Fix: measure `response.finished()` — the **body-completion** time — and a new `slow-body` anomaly kind:

```
[slow-body] HTTP 200 headers in 70 ms but BODY took 37533 ms (streamed/blocked SSR)
            — https://billion.taildc510.ts.net:13000/enterprise/workforce?tab=skills-verification
```

That moved the 37 s back onto the leg that **owns** it.

### The causal chain closes to within 33 ms

iter-01 predicted, from **static analysis alone**: `3 attempts × 10.5 s connect-timeout + (2 s + 4 s backoff)` =
**37.5 s**.
iter-02 measured the SSR body: **37.533 s**.

> **Prediction and measurement agree to within 33 milliseconds.** The root cause is not a hypothesis any more.

And it explains the vantage symmetry: both heroes land within **0.2 s** of each other because they block on the
**same shared authenticated layout** — precisely the "common factor" the re-scope trigger demanded.

---

## New finding

**F-5 — the demo phones out to ad-tech.** Every authenticated page load attempts **Google Analytics**
(`region1.google-analytics.com/g/collect`, `tid=G-6NR74WHQWN`), **DoubleClick** (`ad.doubleclick.net/ccm/s/collect`),
**Google Ads remarketing** (`google.com/rmkt/collect/16728858108`), and **LinkedIn Ads** (`px.ads.linkedin.com`).
They currently `ERR_ABORTED`, but they are **attempted on every load**, from a demo that `safety.md` says is
self-contained. Same class as the already-in-scope Clerk-telemetry item — and arguably worse (it is *ad-tech
tracking*, and it fires on a stack a presenter shows to customers). Routed forward.

---

## Close — 2026-07-13

**Outcome:** The instrument exists. **The gate can be graded for the first time in the project's history** — and
the honest answer is **p95 39.45 s (employee) / 38.30 s (manager)** against a **5 s** gate. Per-leg attribution
puts **the entire 37.6 s on one leg**: the SSR document's *streamed body*, blocked on the unreachable GraphQL
origin. iter-01's static prediction (37.5 s) and this measurement (37.533 s) agree **to within 33 ms**.

**Type:** tik
**Status:** closed-fixed  _(planned deliverable = the instrument + a baseline. Both landed, on a green stack.)_
**Gate:** NOT MET — p95 39.45 s / 38.30 s vs the < 5 s gate (**~7.9× over**). Now, for the first time, **measured**.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n _(1 tik)_ — (6) protocol-stop: n — Outcome: **continue**

**Decisions:** D5 (measure the body, not the headers — the streaming-SSR trap), D6 (the green gate travels to remote stacks by pointing at the real remote verdict, never by bypass)

**Side-deliverables:** none

**Routes carried forward:**
- **→ iter-03:** land the real C-1 fix (`FIX-M218-iter03-ssr-origin`) and re-measure with this harness.
- **F-5 ad-tech egress** → fold into the already-in-scope telemetry work. Handler `FIX-M218-telemetry-egress`.
- All iter-01 carry-forwards remain open.

**Lessons:**
1. **A probe that fires on headers cannot see a blocked streaming render.** This is generalizable far beyond
   M218: *any* Next.js App Router latency measurement that watches `response` events and not
   `response.finished()` will attribute a blocked SSR to a phantom client-side gap. **Folded into the protocol
   doc + `latency-budget.md`.**
2. **The baseline is the M43-D5 correction's evidence.** The corpus said "~2–5 s, which we can't shorten." It is
   **39 s, and we are about to shorten it.** Four releases believed a number nobody had ever measured.
3. **Measure from the vantage that hurts.** The harness runs on the *tailnet*, because that is where the
   presenter stands. Run it on the demo host and the defect partly hides.
