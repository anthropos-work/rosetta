# M218 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | The 4-leg experiment, zero code. **C-1 confirmed as root cause but its MECHANISM was wrong in the plan — SSR reads the BUILD-INLINED public URL, which connect-times-out (10.56 s) from inside the container; the planned runtime-env one-liner is a NO-OP.** Budget now sums to the symptom. C-2 refuted (M217 fixed it — field proof). C-5 refuted as cause, retained as a free win. C-3 parked (unexercisable until a login completes). TOK-01 authored. | baseline **UNMEASURED** — no instrument exists (iter-02 builds it) | closed-fixed — see `iter-01/progress.md` |
| iter-02 | tik | **The latency harness** (rext `stack-verify/e2e/`, new surface — discharges DEF-M215-03(b)) + **the first baseline ever taken.** Found + fixed a measurement trap: Next.js **streams** RSC, so `page.on('response')` (headers) reports a *fast* document while the body blocks for 37 s — now measured via `response.finished()`. **iter-01's static prediction (37.5 s) and the measured SSR body (37.533 s) agree to within 33 ms.** | **employee p95 39.45 s · manager p95 38.30 s** vs the **< 5 s** gate → **~7.9× OVER** (3/3 reached ACCESS — slow, not broken) | closed-fixed — see `iter-02/progress.md` |
| iter-03 | tik | **The 38-second login is GONE.** Landed the real C-1 fix — the `next-web-ssr-graphql-origin` demo-patch (a **server-only** `WUNDERGRAPH_SSR_ENDPOINT`, deliberately not a `NEXT_PUBLIC_*` name so it is a **real runtime read**) + its runtime value in `gen_injected_override.py`. Also **F-6** (side): the next-web image cache-validator was blind to the **minted pk** (it reads image ENV; the pk is *inlined into the bundle*), so an out-of-band rebuild had left the stack **Clerkenstein-DEWIRED** — browser clerk-js talking to the **REAL Clerk app**, login broken, and the stale 9-h-old `autoverify.json` still saying green. Fixed by validating the pk **in-bundle**; it forced the self-healing rebuild on its first run. | **employee p95 39.45 → 7.90 s · manager p95 38.30 → 7.00 s** (cold, green, 6/6 ACCESS) — a **5× collapse**; **1.6× over** the 5 s gate (was 7.9×) | closed-fixed — see `iter-03/progress.md` |
| iter-04 | tik | **THE GATE IS MET.** Root-caused the ~6.10 s residual: Clerkenstein's fake **BAPI served a hardcoded STUB user to EVERY hero** (`// Disarmed: any id → the demo user` — true when a demo had one user, false since the M35 Stories & Heroes model). So `currentUser().externalId` (BAPI) disagreed with the JWT identity (FAPI) → `app`'s `userPreferences` resolver refused the mismatch → next-web's `retry: 2` / 2 s+4 s ladder burned ~6 s on **every** authenticated render. Fixed by making the BAPI **roster-aware**. `/align-run` across **all 5 surfaces: 100% critical / 100% overall, 0 divergences** — nothing moved. Also landed the **M43-D5 correction** + the **`latency-budget.md`** blind area. | **employee p95 1.46 s · manager p95 1.40 s** vs the < 5 s gate — **5/5 ACCESS both vantages, gate armed, cold+green**. From 39.45/38.30 s = **27× faster** | **closed-fixed — GATE MET** — see `iter-04/progress.md` |

## Next iter

**None — the milestone's exit gate is MET** (iter-04). Next: `/developer-kit:harden-mstone-iters --final`, then
`/developer-kit:close-milestone`.

## Carry-forward queue (none block the gate)

| handler | item |
|---|---|
| `PROBE-M218-backend-api-url-twin` | **F-7:** `NEXT_PUBLIC_BACKEND_API_URL` bakes to `https://billion…:18082` — **measured as a 10.5 s blackhole from inside the container**, the exact C-1 shape. Dormant **only** because every reader is client-side (D10). **A loaded gun.** |
| `FIX-M218-telemetry-egress` | **F-5:** the demo attempts **Google Analytics + DoubleClick + Google Ads + LinkedIn Ads** on every authenticated load (+ the in-scope Clerk-telemetry off). |
| `FIX-M218-c5-clerkjs` | **C-5:** vendor clerk-js + bound the **unbounded** `Timeout: 0` (`server.go:187`). Alignment-invisible ⇒ gate-free. |
| `PROBE-M218-c3-rerun` | **C-3:** now exercisable. The router **is** logging cms/Directus **403s** (`getSkillPaths`, `_entities JobSimulation`) on the CONTENT path — not the login path; affects data-settle. |
| `DOC-M218-audit-corrections` | **DONE for M43-D5 + `latency-budget.md`** (iter-04). **Remaining:** the CI-inert correction (`alignment_testing.md:232,233,239`) + the `clerkenstein.md:3-4` header. |
| _(also-in-scope, from overview)_ | ant-academy real-Clerk-secret leak · `x/crypto@v0.52.0`. |
