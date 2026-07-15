---
milestone: M218
slug: seat-change
version: v2.3 "cue to cue"
milestone_shape: iterative
status: archived
created: 2026-07-13
last_updated: 2026-07-14
complexity: large
depends_on: M217
exit_gate: "p95 click→ACCESS < 5 s — where ACCESS = the authenticated shell is rendered and interactive with the hero's identity present (full-screen loading gone; user menu shows the hero) — for BOTH maya-thriving (employee → /profile) AND dan-manager (manager → /enterprise/…), measured over 5 consecutive cold reset-to-seed runs. In-page data-completion time is REPORTED, not gated (D-DESIGN-1)."
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md — the live-browser measure → attribute → fix → re-measure loop, here driving a LATENCY gate rather than a presence gate)
delivers: click→access under 5 s for both vantages + corpus/ops/demo/latency-budget.md (there is no perf budget, baseline, or even a definition of "access" anywhere in the project) + the full click→painted-page login sequence (documented nowhere) + the M43-D5 correction
---

# M218 — Seat change

## Goal
Click **[Log in as]** → the hero is **in the platform** in **under 5 seconds**. Both vantages. Every time.

## Exit gate (measurable)
**p95 click→ACCESS < 5 s**, where **ACCESS** :=

> the authenticated shell is **rendered and interactive** with the hero's **identity present** — the full-screen
> loading state is gone and the user menu shows the hero.

for **BOTH**:
- `maya-thriving` (employee → `/profile`), and
- `dan-manager` (manager → `/enterprise/…`),

measured over **5 consecutive cold reset-to-seed runs**.

**In-page data-completion time (the 200-member grid) is REPORTED as a secondary metric, NEVER gated** — per
**D-DESIGN-1** (user, 2026-07-13): *"the gate is on the login/access to the platform, not on the load of the
complete render of the first page of the user."* This deliberately sidesteps the platform-side DataLoader defect
(`GetOrganizationTargetRole` = 3 RPCs/member, ~1000–1500 round-trips/page) rather than fighting it — that root cause
is **platform source and cannot be fixed under the hard constraint**.

## Why iterative (not section)
**The confirmed cost budget does not yet sum to the symptom.** The two suspects we can prove today total ~18 s
(C-1 ~6 s + C-2 ~11.6 s patched). Reaching **60–120 s** requires one of the *unconfirmed* big-ticket items to be
real. Writing a fixed `In:` list now would be **guessing which fix to build**. The gate is the commitment; the path
emerges from the measurement.

## Iteration protocol
`corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md` — the live-browser **measure → attribute →
fix → re-measure** loop, here driving a **latency** gate instead of a presence gate.

> ### iter-01 is the HARNESS + the 4-leg experiment. Write NO fix before it runs.
> The experiment discriminates **every** suspect below in **ONE bring-up with zero code written**:
> 1. `docker exec demo-N-next-web node -e "console.log(process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT)"` → C-1
> 2. `docker logs demo-N-graphql-1 | grep '"latency"' | tail` after each hero click → **C-3** (and the query name
>    hands you the fix surface)
> 3. `docker exec demo-N-fake-fapi wget -S -O /dev/null https://cdn.jsdelivr.net/npm/@clerk/clerk-js@5/…` (timed) → C-5
> 4. `grep '⚠ app:' <bring-up log>` → C-2 (should be GONE after M217)

**Harness build notes:**
- Reuse `stack-verify/e2e/lib/cockpit-login.ts` (`selectSeat` :40-53, `loginAs` :59-76) — the seat-switch handshake
  is **already Playwright-driven**. Do not fork it.
- Instrument with `page.on('response')` + `performance.getEntriesByType('navigation')`; **split the legs**:
  click → handshake/303 → SSR → clerk-js → client-gate → data-query. Attribution is the whole point.
- **Do NOT gate on `networkidle`** — next-web holds never-idle long-polls (`section-assert.ts:64`). Poll for
  **content presence** instead.
- The latency gate **cannot be a Playthrough** — Playthroughs declare perf a **NON-GOAL**. It is a **new
  `stack-verify` surface**, and building it discharges **DEF-M215-03(b)** for free.

## Ranked suspects (adversarially surviving)

The click→handshake→303 leg is **provably free** (`clerkenstein/clerk-frontend/server.go:204-255`, no I/O;
`grep time.Sleep` across `clerkenstein/` → **0 hits**; the cockpit CTA is a plain `<a href>` at `cockpit.py:397`).
**Every surviving suspect lives AFTER the 303.**

| | Suspect | Vantage | Est. cost | Fix surface |
|---|---------|---------|-----------|-------------|
| **C-1** | next-web's **server-side** GraphQL URL resolves to its own loopback. Container runtime env is `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://${PUBLIC_HOST:-localhost}:5050/graphql`; the demo exports `STACK_PUBLIC_HOST`, **never** `PUBLIC_HOST` → the container's own loopback on the **un-offset** port → instant ECONNREFUSED. `(authenticated)/layout.tsx` is `force-dynamic` and **blocks** on `prefetchUserStatus` (`retry:2`, `retryDelay 2s/4s`), and all three fetchers **rethrow**, so the retries genuinely fire. | **BOTH** | **~6 s per authenticated render** | **one line** — add `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://graphql:8080/graphql` to the next-web entry in `stack-injection/gen_injected_override.py:96-112`. `NEXT_PUBLIC_*` is **build-inlined**, so changing the **runtime** env re-points only the SSR read; the browser keeps its correctly-baked offset URL. **rext-only.** |
| **C-2** | the two dead `app` perf demo-patches | manager | 76 s → **11.6 s even patched** | M217 re-pins. **The 11.6 s residual is OUT of gate scope** (D-DESIGN-1) — report it, don't chase it. |
| **C-3** | **cold-federation Directus drift → the Cosmo router RETRYING.** A cms `SetFields('*')` selects a column the replayed Directus lacks → 500 → the router retries. Documented in-repo: *"the ~60–90 s 'latency' is the router RETRYING… cache-masked in a warm sweep; surfaces only on a COLD federation tier"* (`coverage-protocol.md:198`). **Amplified on `billion`, where directus replay SKIPPED and cms read content live from prod over the WAN.** | **BOTH** | **60–90 s** | M217's cache prime. **Arithmetically the closest single match to "1–2 minutes", and it sits on a path BOTH heroes traverse.** |
| **C-4** | stale cockpit / dead clerk-ids (the crash) | both | unknown, possibly all of it | **M217. Must land first — it contaminates every measurement.** |
| **C-5** | the fake FAPI proxies `clerk.browser.js` **live from `cdn.jsdelivr.net`** on every full page load — `http.Get` = `http.DefaultClient`, **`Timeout: 0` (unbounded)**, **no server-side cache** (`clerk-frontend/server.go:187,101,21`). next-web's whole authed tree is **client-gated on clerk-js**. The existing egress pre-check curls from the **HOST, not from inside the container** (`up-injected.sh:1029`), so it can pass green while the container cannot reach the CDN. | both | **0.2 s healthy / ~127 s if egress blackholes** | **Vendor the bundle into the fake-fapi image** (serve from disk; CDN proxy only as fallback). **Alignment-INVISIBLE — zero DNA genes cover `GET /npm/` → a gate-free win.** Take it regardless of whether it is today's cause: it removes an **unbounded internet dependency from the login path of a demo that claims to be self-contained**. |
| **C-6** | `billion` has **7.325 GiB RAM** vs the documented **12 GiB** floor (the tooling warns every run) | both, remote only | unknown | Measure `docker stats` + `free -h` **during** a login **before blaming code**. May be a pure VM resize. |

**REFUTED (do not re-litigate):** a cold Next.js JIT compile (the demo runs a **production** build — `Dockerfile.dev`
runs `pnpm turbo build`; compose pins `NODE_ENV=production`). A hardcoded wait / retry ladder / poll inside
Clerkenstein or the cockpit (the overlay timers are cosmetic text swaps that never `preventDefault`).

## Also in scope (same code surface, land them while we are here)
- **Disable Clerk telemetry** (`CLERK_TELEMETRY_DISABLED` + `NEXT_PUBLIC_CLERK_TELEMETRY_DISABLED`) — real egress
  from both frontends, and it is what makes Playwright's `networkidle` hang. Pure env, no repo edit.
- **Clerkenstein-wire ant-academy** — `ant-academy.sh:146` copies `CLERK_SECRET_KEY` **straight from
  `platform/.env`**, i.e. the **REAL** Clerk app's secret, into a demo process; `PK_DEMO` is never used, so
  @clerk/nextjs runs **keyless** and phones Clerk to provision a throwaway app. Off the login path, but it is
  real-Clerk egress + a real secret in a demo — the same class as the `DIRECTUS_TOKEN` fix16/17 strip, and it
  contradicts `safety.md`. Pure env.
- **`go get golang.org/x/crypto@v0.52.0`** — 13 dependabot alerts, all govulncheck-**UNREACHABLE**; we are in
  `clerkenstein/` anyway.

## Alignment guard (BLOCKING for any `clerkenstein/clerk-frontend/` change)
- **No DNA gene covers latency, perf, timeouts, or the clerk-js bundle proxy** → caching/vendoring is **free**.
- But the genes that break on a **handshake-shape** change are **all `critical`**, and the critical gate is
  **100% — no partial credit** (`HandshakeSelect/*`, `DistinctIdentity/*`, `SessionToken/decoded-identity` with an
  **`exact`** operator, `cookie-no-dev-browser-signed-out` — "optimize away a cookie" is a booby trap).
- **THE GATE IS NOT ENFORCED BY CI.** `clerkenstein/.github/workflows/alignment.yml:10-11` declares itself
  *"currently inert"*, while `corpus/architecture/alignment_testing.md:233` claims a weekly workflow gates it.
- ⇒ **Any change to `clerkenstein/clerk-frontend/` MUST carry an explicit `/align-run` step across all 5 surfaces.
  Do not rely on CI.** Current score: **100% critical / 100% overall, 56 genes**; gate = critical ≥100 / overall ≥95.

## Re-scope trigger
**If the harness shows that neither C-1 nor C-3 explains the EMPLOYEE vantage, STOP and re-measure.** Both dead
patches are *manager* surfaces, yet the user reports **both** vantages are slow ⇒ there **must** be a common factor,
and only C-1 and C-3 qualify. **Do not proceed on a manager-only fix set.**

## Delivers → knowledge/corpus
- **`corpus/ops/demo/latency-budget.md`** — **BLIND AREA.** There is **no** perf/latency budget, baseline, gate, or
  even a **definition of "access"** anywhere in `corpus/**` or rext (grep `latency|p95|perf budget` → only AI-model
  latency in `ai_architecture.md`). Must define: the click→interactive budget, the **per-leg attribution model**, the
  measured baseline, the <5 s gate, and the harness contract.
- **The full click→painted-page login sequence** in `corpus/ops/demo/cockpit-spec.md` — **BLIND AREA.** The
  **next-web half** of the login path (post-303: clerkMiddleware verify → RSC render → Cosmo fan-out → per-subgraph
  cold start) is documented **nowhere**; the two existing anchors are **one line each**. *You cannot code a fix
  against that.*
- **The clerk-js proxy's caching/timeout contract** in `corpus/services/clerkenstein.md` (+ the rext knowledge doc).
- **The M43-D5 CORRECTION** — the corpus claims cockpit login is *"~2–5 s, which we can't shorten"* in **4 places**
  (`cockpit-spec.md:58,155`; `cockpit.py:12,204-208`). It is **60–120 s**. Booked as an M43 scope-`Out:` + decision
  **D5** with **zero deferrals recorded**, so it never entered a ledger and was never revisited across four
  releases. **This is why nobody investigated.** Re-open it formally.
- **The CI-inert correction** in `corpus/architecture/alignment_testing.md:233`.
- The stale status/repo-line header in `corpus/services/clerkenstein.md:3-4`.

## KB dependencies
- `corpus/ops/demo/cockpit-spec.md` · `corpus/services/clerkenstein.md` · `corpus/ops/demo/coverage-protocol.md`
- `corpus/architecture/alignment_testing.md` (the gate that governs any Clerkenstein change)
- `corpus/ops/verification.md` (the iteration protocol)
- rext: `stack-verify/e2e/lib/cockpit-login.ts` (the harness to reuse) · `stack-injection/gen_injected_override.py`
  (C-1's one-line fix) · `clerkenstein/clerk-frontend/server.go` (C-5)
