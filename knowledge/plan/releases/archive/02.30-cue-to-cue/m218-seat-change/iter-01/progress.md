**Type:** tok (bootstrap)

# iter-01 — the 4-leg experiment

Run against the **live green demo** on `billion` (`demo-1`, rext `cue-to-cue-m217`, 16 containers).
Measurement gate honoured: `autoverify.json` = `{"warnings":0,"green":true}` **before** a single number was taken.

---

## The headline

> **The root cause is NOT the mechanism the plan predicted — but it IS in the place the plan predicted.**
>
> next-web's **SSR** GraphQL origin is the **build-inlined public URL** `https://billion.taildc510.ts.net:15050/graphql`.
> From **inside** the container that URL **connect-times-out after 10,564 ms**. The authenticated layout *blocks*
> on it, with `retry:2` + 2 s/4 s backoff. **That is where the 1–2 minutes lives.**

**The budget now sums to the symptom** — for the first time in this milestone.

---

## Leg-by-leg

### LEG 1 → C-1: the SSR GraphQL URL — **CONFIRMED, mechanism CORRECTED**

The plan's leg-1 probe (`docker exec … node -e "console.log(process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT)"`) ran and
returned exactly what the plan predicted:

```
NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT = http://localhost:5050/graphql     # un-offset loopback, as predicted
```

**But the Phase-0b audit flagged that this probe cannot discriminate the thing that actually matters** — whether
the compiled SSR bundle *reads `process.env` at runtime* or *uses a build-inlined literal*. `frontend-tier.md:240`
says one thing; the milestone's C-1 says the opposite. **That question decides whether the planned one-line fix is
a fix or a no-op.** So leg 1 was strengthened (this is the audit's recommendation, discharged):

```
# what is actually IN the compiled server bundle?
$ grep -rhoE 'https?://[^ ]+/graphql' /app/apps/*/.next/server/
      9  https://billion.taildc510.ts.net:15050/graphql      ← BUILD-INLINED. Not process.env.

# and in the CLIENT bundle?
      1  https://billion.taildc510.ts.net:15050/graphql      ← correct for the browser
```

**Verdict: `NEXT_PUBLIC_*` is build-inlined into the SERVER bundle too.** `frontend-tier.md:240` is **right**;
the milestone's C-1 premise is **wrong**. Consequences, in order of importance:

1. **The planned C-1 fix is a NO-OP.** Adding `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://graphql:8080/graphql` to
   the **runtime** env in `gen_injected_override.py` would change `process.env` — which the SSR code **never
   reads**. It would have shipped, measured zero, and cost an iter. **The audit + the strengthened leg saved that.**
2. **The real defect is worse than C-1 described**, and it is *specific to the `--public-host` (tailnet) demo*:

   | | value | reachable from inside the container? |
   |---|---|---|
   | **baked into SSR bundle** | `https://billion.taildc510.ts.net:15050/graphql` | ❌ **connect-timeout 10,564 ms** |
   | container runtime env (unused by SSR) | `http://localhost:5050/graphql` | ❌ ECONNREFUSED 64 ms |
   | **what actually works** | `http://graphql:8080/graphql` | ✅ **HTTP 200 in 94 ms** |

   DNS is **fine** (`billion.taildc510.ts.net` → `100.110.136.3` instantly). It is the **TCP connect** from the
   docker bridge to the tailscale IP that **blackholes** → undici's **10 s default connect timeout** fires.

3. **The arithmetic finally closes:**
   - `(authenticated)/layout.tsx` is `force-dynamic` and **blocks** on `prefetchUserStatus`; all three fetchers
     **rethrow**; `retry: 2`, `retryDelay` 2 s → 4 s.
   - per blocking fetcher: **3 attempts × 10.5 s + (2 s + 4 s) backoff ≈ 37.5 s**
   - three such fetchers ≈ **~112 s** → **"1 or 2 minutes."** The user's report, to the second.
   - **Both vantages** traverse the authenticated layout ⇒ **both** are slow. The common factor the re-scope
     trigger demanded **exists**, and this is it.

4. **Why four releases missed it.** Locally (no `--public-host`) the baked URL is `http://localhost:15050/graphql`
   → from inside the container that is **ECONNREFUSED in 64 ms**, not a 10 s timeout → the same defect costs only
   ~6 s locally. **The bug is ~20× worse on the tailnet demo than on the laptop it was always measured on.**
   The corpus's "~2–5 s, which we can't shorten" was measured on the machine where the defect is cheapest.

### LEG 2 → C-3: cold-federation Directus drift / router retries — **NOT OBSERVED (not yet exercisable)**
```
docker logs demo-1-graphql-1  → 12 lines total, 0 errors, 0 retries, 0 500s
docker logs demo-1-cms-1      → 0 errors
```
The Cosmo router has **served almost nothing** — the federation is cold and *untouched*, because the SSR fetch that
would drive it **never reaches the router** (it dies in the connect timeout upstream). **C-3 cannot be refuted until
a login actually completes.** Carried forward: re-run leg 2 the moment leg 1's fix lands and a real login succeeds.

### LEG 3 → C-5: fake-FAPI proxies clerk-js from `cdn.jsdelivr.net` — **REFUTED as today's cause; kept as a fix**
```
timed from INSIDE demo-1-fake-fapi-1:  0.19 s / 0.18 s / 0.17 s   (rc=0)
```
Egress is healthy; this is **not** today's 60–120 s. **But the liability the milestone named is real and confirmed
by the audit** (`server.go:187` = `http.Get` = `DefaultClient` = **`Timeout: 0`, unbounded**; no *server-side* cache —
only a response-side `Cache-Control`). It is an **unbounded internet dependency on the login path of a demo that
claims to be self-contained**, and it is **alignment-invisible** (no DNA gene covers `GET /npm/`) ⇒ a **free,
gate-free win**. Take it regardless. Routed to a tik.

### LEG 4 → C-2: the two dead `app` perf demo-patches — **REFUTED: M217 fixed them**
```
current stack's bring-up log (m217-proof2.log, 10:49):   ✓ demo-patches: none refused
```
The refusals only appear in **pre-M217** logs (Jul-11). **M217's re-pin worked, and this is its field proof** —
which M217 itself could not claim (it closed "unit-proven, not field-proven"). C-2 is closed.

### C-6: `billion` RAM — **CONFIRMED as a standing condition, NOT implicated**
`7.3 GiB` total vs the documented **12 GiB** floor. But: `2.1 Gi used, 5.3 Gi available`, no OOM kills, all 16
containers healthy for 4 h. **Not starving.** Not today's cause. Note it; don't chase it.

---

## Findings the plan did not have

| id | finding |
|----|---------|
| **F-1** | **The click→handshake→303 leg is EMPIRICALLY free: 17 ms.** The overview proved this statically (no I/O, no `time.Sleep`); now it is measured. Every second lives *after* the 303, exactly as claimed. |
| **F-2** | **`autoverify.json` is not where the corpus says it is.** `verification.md:207` says `<stack>/autoverify.json`; it is actually at `rosetta-extensions/demo-stack/stacks/demo-1/autoverify.json`. **M218 gates on this file** — a doc that points at the wrong path is a gate that silently doesn't fire. Must fix. |
| **F-3** | **The fake-FAPI validates `redirect_url` against the public origin** (HTTP **400** on a loopback `redirect_url`), and next-web's middleware **307s** any non-https origin. ⇒ **curl cannot measure this flow.** The harness *must* be a real browser against the real https origin. This independently confirms the plan's Playwright mandate — and it is a **harness design constraint that was documented nowhere**. |
| **F-4** | `NEXT_PUBLIC_HOSTING_URL=http://localhost:3000` — the **same un-offset-loopback class** as C-1, in the same container. Not yet shown to be on the login path. Probe it when leg 1's fix lands. |

---

## Suspect ledger after iter-01

| | suspect | verdict | cost |
|---|---|---|---|
| **C-1** | SSR GraphQL origin | ✅ **CONFIRMED — root cause. Mechanism corrected: build-inlined public URL, connect-timeout, NOT runtime-env ECONNREFUSED.** | **~37 s/fetcher → ~112 s** |
| C-2 | dead app perf patches | ❌ refuted — **M217 fixed it** (field-proven here) | 0 |
| C-3 | cold-federation router retries | ⏸ **unexercisable until a login completes** — re-run after the fix | unknown |
| C-5 | clerk-js CDN proxy | ❌ refuted as cause (0.19 s) — **but fix anyway: unbounded timeout, no cache, gate-free** | 0 today |
| C-6 | `billion` 7.3 GiB RAM | ❌ not implicated (5.3 GiB available, no OOM) | 0 |

---

## Close — 2026-07-13

**Outcome:** The 4-leg experiment discriminated **every** suspect in one bring-up with **zero code written**, exactly
as the plan demanded. **C-1 is the root cause — but its mechanism was wrong in the plan, and the planned fix was a
no-op.** The SSR bundle is **build-inlined** with the public URL, which **connect-times-out (10.5 s)** from inside
the container; ×3 attempts ×3 blocking fetchers ≈ **112 s** — the budget finally sums to the reported 1–2 minutes.
C-2 refuted (M217 fixed it — field proof). C-5 refuted as cause, retained as a free win. C-3 parked (unexercisable
until a login completes). TOK-01 authored.

**Type:** tok (bootstrap)
**Status:** closed-fixed  _(planned deliverable = the experiment + the strategy. Both landed. No fix was in scope.)_
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n _(bootstrap toks do not terminate the call)_ — (3) re-scope: **n** _(see below)_ — (4) user-blocker: n — (5) cap-reached: n _(0 tiks so far)_ — (6) protocol-stop: n — Outcome: **continue**

**Re-scope trigger — graded explicitly, did NOT fire.** The trigger reads: *"if neither C-1 nor C-3 explains the
EMPLOYEE vantage, STOP."* **C-1 does explain it** — the authenticated layout is common to both vantages, and the
SSR connect-timeout hits every authenticated render regardless of role. The *mechanism* was corrected
(build-inlined ≠ runtime-env), but the *suspect and its location are unchanged*. A within-suspect mechanism
refinement is **not** a re-scope; the strategy chain holds. Recorded so a future reader can audit the call.

**Decisions:** D1 (root cause + the no-op the plan would have shipped), D2 (fix shape: reachability, not rewiring),
D3 (`autoverify.json` path drift), D4 (harness must be a browser — curl is structurally blocked)

**Side-deliverables:** none (no code written, by design)

**Routes carried forward:**
- **→ iter-02 (tik):** build the Playwright latency harness + take the **real p95 baseline**, both vantages.
  Handler `HARNESS-M218-iter02-latency-probe`.
- **→ iter-03 (tik):** land the **real** C-1 fix (container-reachable SSR origin — see D2), re-measure.
  Handler `FIX-M218-iter03-ssr-origin`.
- **→ a tik after the fix:** **re-run leg 2 (C-3)** — the router is only exercisable once a login completes.
  Handler `PROBE-M218-c3-rerun`.
- **→ a tik:** C-5 vendor/cache + bound the timeout (free, gate-free). Handler `FIX-M218-c5-clerkjs`.
- **→ docs (any tik):** F-2 `autoverify.json` path drift · F-3 the harness constraint · F-4 `NEXT_PUBLIC_HOSTING_URL`.
- **→ the audit's corrections:** `cockpit-spec.md:155`→**`:187`** anchor drift; `alignment_testing.md` CI-inert
  correction needs **`:232,239`** too (milestone declared only `:233`); `frontend-tier.md:240` is **right** and must
  be **preserved**, not "corrected". Handler `DOC-M218-audit-corrections`.

**Lessons:**
1. **The strengthened probe was the whole iter.** The plan's leg-1 one-liner would have "confirmed" C-1 and sent
   iter-02 to ship a **no-op**. The Phase-0b audit caught the ambiguity; the *stronger* probe (grep the compiled
   bundle, don't trust `process.env`) settled it. **When a probe reads a variable, ask whether the code under test
   reads that same variable.** `process.env.X` and the *build-inlined literal of* `process.env.X` are different
   things, and only one of them is what runs.
2. **A KB-fidelity audit paid for itself before a line of code.** `frontend-tier.md:240` contradicted the
   milestone's own premise; the doc was right and the plan was wrong. Blocking Phase 0b on iter-01 is load-bearing.
3. **Measure on the machine where it hurts.** The same defect is ~6 s on a laptop and ~112 s on the tailnet VM. The
   corpus's "2–5 s, we can't shorten it" was an honest measurement of the cheap case, generalized to the expensive
   one. → **`latency-budget.md` must state the environment a number was taken on.** (Protocol lesson — folded into
   the doc deliverable.)
