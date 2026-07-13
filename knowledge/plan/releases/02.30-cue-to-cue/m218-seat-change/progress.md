# M218 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | The 4-leg experiment, zero code. **C-1 confirmed as root cause but its MECHANISM was wrong in the plan — SSR reads the BUILD-INLINED public URL, which connect-times-out (10.56 s) from inside the container; the planned runtime-env one-liner is a NO-OP.** Budget now sums to the symptom. C-2 refuted (M217 fixed it — field proof). C-5 refuted as cause, retained as a free win. C-3 parked (unexercisable until a login completes). TOK-01 authored. | baseline **UNMEASURED** — no instrument exists (iter-02 builds it) | closed-fixed — see `iter-01/progress.md` |
| iter-02 | tik | **The latency harness** (rext `stack-verify/e2e/`, new surface — discharges DEF-M215-03(b)) + **the first baseline ever taken.** Found + fixed a measurement trap: Next.js **streams** RSC, so `page.on('response')` (headers) reports a *fast* document while the body blocks for 37 s — now measured via `response.finished()`. **iter-01's static prediction (37.5 s) and the measured SSR body (37.533 s) agree to within 33 ms.** | **employee p95 39.45 s · manager p95 38.30 s** vs the **< 5 s** gate → **~7.9× OVER** (3/3 reached ACCESS — slow, not broken) | closed-fixed — see `iter-02/progress.md` |
| iter-03 | tik | **The 38-second login is GONE.** Landed the real C-1 fix — the `next-web-ssr-graphql-origin` demo-patch (a **server-only** `WUNDERGRAPH_SSR_ENDPOINT`, deliberately not a `NEXT_PUBLIC_*` name so it is a **real runtime read**) + its runtime value in `gen_injected_override.py`. Also **F-6** (side): the next-web image cache-validator was blind to the **minted pk** (it reads image ENV; the pk is *inlined into the bundle*), so an out-of-band rebuild had left the stack **Clerkenstein-DEWIRED** — browser clerk-js talking to the **REAL Clerk app**, login broken, and the stale 9-h-old `autoverify.json` still saying green. Fixed by validating the pk **in-bundle**; it forced the self-healing rebuild on its first run. | **employee p95 39.45 → 7.90 s · manager p95 38.30 → 7.00 s** (cold, green, 6/6 ACCESS) — a **5× collapse**; **1.6× over** the 5 s gate (was 7.9×) | closed-fixed — see `iter-03/progress.md` |

## Next iter

**iter-04 (tik) — kill the ~6.10 s retry ladder. It is now the ENTIRE gap to the gate.** Under **TOK-01** step 3.
The SSR body still blocks **6,104 / 6,107 ms — reproducible to ±3 ms across both vantages**. That is **not** a
blackhole (a blackhole costs 10.5 s/attempt); it is the **`retry: 2` / 2 s+4 s ladder on a fetch that now fails
FAST**: `3 × ~33 ms + (2 s + 4 s) ≈ 6.0 s`. So one of the authenticated layout's blocking fetchers still **errors** —
it just errors *immediately* instead of timing out. **Identify which fetcher, and why it errors** (instrument the
SSR pass; the container logs are silent because the fetchers rethrow into react-query). Kill that, and p95 goes
from **7.90 s → ~1.8 s** — under the gate. Handler `FIX-M218-iter04-ssr-retry-ladder`.

## Carry-forward queue

| handler | item |
|---|---|
| `FIX-M218-iter04-ssr-retry-ladder` | **THE GATE'S REMAINING GAP.** The ~6.10 s SSR retry ladder on a fast-failing fetch (iter-03). |
| `PROBE-M218-backend-api-url-twin` | **F-7:** `NEXT_PUBLIC_BACKEND_API_URL` bakes to `https://billion…:18082`, **measured as a 10.5 s blackhole from inside the container** — the exact C-1 shape. Dormant **only** because every reader is client-side (D10). A loaded gun. |
| `PROBE-M218-c3-rerun` | Re-run leg 2 (C-3, router retries) — **now finally exercisable**: logins complete, so the federation is exercised for the first time. |
| `FIX-M218-c5-clerkjs` | Vendor clerk-js + **bound the unbounded timeout** (`server.go:187`, `Timeout: 0`). Alignment-invisible ⇒ gate-free. |
| `DOC-M218-audit-corrections` | M43-D5 correction (**`cockpit-spec.md:58,187`** — 155 was wrong) · CI-inert (**`alignment_testing.md:232,233,239`**) · `clerkenstein.md:3-4` header · **preserve** `frontend-tier.md:240` (it was right). |
| `DOC-M218-f2-autoverify-path` | **F-2:** `verification.md:207` documents the wrong path for `autoverify.json` — the file M218 gates on. |
| `DOC-M218-f3-harness-constraint` | **F-3:** curl structurally cannot drive the login flow (redirect_url validation + https 307). Undocumented. |
| `FIX-M218-telemetry-egress` | **F-5 (iter-02):** the demo attempts **Google Analytics + DoubleClick + Google Ads remarketing + LinkedIn Ads** on every authenticated load. Fold into the in-scope Clerk-telemetry work. |
| `PROBE-M218-f4-hosting-url` | **F-4:** `NEXT_PUBLIC_HOSTING_URL=http://localhost:3000` — same un-offset-loopback class as C-1. On the login path? |
| _(also-in-scope, from overview)_ | Clerk telemetry off · ant-academy real-Clerk-secret leak · `x/crypto@v0.52.0` |
