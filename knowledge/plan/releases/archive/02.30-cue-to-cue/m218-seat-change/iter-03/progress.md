**Type:** tik  ·  **Active strategy:** TOK-01 (reachability-first)

# iter-03 — the 38-second login is gone

## Delivered

| # | what | where |
|---|---|---|
| 1 | **`next-web-ssr-graphql-origin`** — the sha-pinned demo-patch that gives next-web's **server-side** GraphQL client an origin it can actually reach (`WUNDERGRAPH_SSR_ENDPOINT`, a deliberately **non-`NEXT_PUBLIC_*`** name so it is a **real runtime read**), plus the runtime value in `gen_injected_override.py` | rext `9a8b6d5` |
| 2 | **F-6** — the next-web image cache-validator now gates on the **minted publishable key**, read from the **bundle** (not the image ENV, which structurally cannot see it) | rext `0cfe23e` |

## THE NUMBER

_billion · demo-1 · rext `cue-to-cue-m218-iter03b` · **cold** reset-to-seed · `autoverify.json` **`{"green":true}`, 0
warnings** · measured **from the tailnet** (the presenter's actual vantage) · 3 samples/vantage._

| vantage | baseline (iter-02) | **iter-03** | Δ | gate |
|---|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | p95 **39.45 s** | p95 **7.90 s** (p50 6.93 s) | **−31.6 s / 5.0×** | < 5 s |
| **manager** (`dan-manager` → `/enterprise/…`) | p95 **38.30 s** | p95 **7.00 s** (p50 6.68 s) | **−31.3 s / 5.5×** | < 5 s |

**6/6 runs reached ACCESS.** **Gate NOT met** — but the milestone's headline defect is **gone**.

**Secondary, REPORTED-NOT-GATED (D-DESIGN-1):** data-settle p50 **12.60 s** (employee) / **9.29 s** (manager).

### The leg that owned the defect gave it all back

```
                          baseline        iter-03
handshake  @0.03 s (303)   FREE            FREE          ← unchanged, as predicted
ssr-document BODY          37,533 ms   →   6,104 ms      ←──── THE FIX. −31.4 s on ONE leg.
data-query                 @37.6 s     →   @6.2 s        ← the client's queries were never the problem;
ACCESS                     38.2 s      →   6.9 s              they simply could not START until SSR unblocked
```

Proof, from inside the container, on the live green stack:

| origin | from inside `demo-1-next-web-app-1` |
|---|---|
| `http://graphql:8080/graphql` **(the fix's origin)** | **76 ms · HTTP 200** |
| `https://billion…:15050/graphql` (the build-inlined public URL SSR *used* to fetch) | **10,481 ms · `UND_ERR_CONNECT_TIMEOUT`** |

The 10.5 s blackhole is still there — the fix does not *repair* the unreachable address, it **stops the server from
using it**. That is the whole of TOK-01: *fix the address, not the variable.*

## The stack the fix was almost measured on wasn't wired (F-6)

Run 1 rebuilt next-web **out-of-band** — it passed the three `--build-arg` offset URLs but never wrote the
`apps/web/.env.local` overlay, which is the **only** carrier of the minted publishable key. `NEXT_PUBLIC_*` being
**build-inlined**, the bundle silently baked the **real-Clerk pk**
(`pk_test_b3JpZW50ZWQtbGFiLTMz…` = `oriented-lab-33.clerk.accounts.dev$`). The demo's browser clerk-js then talked
to the **real Clerk app**: no session → `/login` loop. The recovered run shows it — `reachedAccess: 0`.

**The tooling could not see it.** The M211 cache-validator compares the baked `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`
**image ENV**; the pk is **not an image ENV** (it is inlined into the bundle), so `docker image inspect` is blind to
it. An image with the **right offset** and the **wrong key** *passed* and would have been **silently reused** —
**corrupting this milestone's own 5-cold-run gate battery**, and shipping a demo that phones **production auth**
(which `safety.md` forbids).

**F-6 fixed it, and the fix proved itself in the field on its first run** — both branches:

```
▶ next-web: cached image demo-1-next-web does NOT carry this stack's minted publishable key
  (the .env.local overlay was missing at build time ⇒ the bundle baked SOME OTHER pk — possibly
  the REAL Clerk app's). Clerkenstein would be DEWIRED. Removing + rebuilding.        ← forced the rebuild
…
▶ next-web: image demo-1-next-web already built (offset endpoint https://…:15050/graphql;
  minted pk verified in-bundle) — reusing (no rebuild).                                ← correctly reused after
```

## The next bottleneck has already named itself

TOK-01 said: *"expect a new bottleneck to appear — the fix does not end the milestone, it unblocks the measurement
of the next layer."* It appeared, and it is **arithmetically legible**:

```
[slow-body] HTTP 200 headers in 62 ms but BODY took 6104 ms (streamed/blocked SSR)
```

**6,104 / 6,107 ms — reproducible to ±3 ms across both vantages.** That is **not** a blackhole (a blackhole costs
10.5 s/attempt). It is the **retry ladder on a FAST-failing fetch**:

> 3 attempts × ~33 ms + (2 s + 4 s backoff) ≈ **6.0 s**

So an SSR fetcher still **errors** — it now fails *immediately* instead of timing out. **Which one, and why, is
iter-04.** (Note the symmetry survives: **both** vantages sit within 0.9 s of each other, because both still block
on the **same shared authenticated layout** — the same common factor the re-scope trigger demanded, still holding.)

## New finding

**F-7 — the C-1 twin is REAL, but dormant.** `NEXT_PUBLIC_BACKEND_API_URL` bakes to `https://billion…:18082`, and
that address is **also a 10.5 s blackhole from inside the container** (measured). It does **not** cost us anything
today only because **every reader is client-side** (D10) — and the browser genuinely wants the public origin. It is
a **loaded gun**: the day any server component reads it, the 37 s defect returns wearing a different hat. Routed
forward.

---

## Close — 2026-07-13

**Outcome:** **The 38-second login is gone.** p95 **39.45 s → 7.90 s** (employee) and **38.30 s → 7.00 s**
(manager) — a **5×** collapse, both vantages, on a **cold, green** stack, with **6/6** runs reaching ACCESS. The
entire lift came off the single leg iter-01 predicted and iter-02 measured. The gate (< 5 s) is **not yet met**, and
the residual is already attributed to a **~6.10 s retry ladder on a fast-failing SSR fetch**.

**Type:** tik
**Status:** closed-fixed  _(planned scope = land the real C-1 fix + re-measure. Both landed; the metric moved as
predicted — the 37.5 s term collapsed to ~0. F-6 was a blocker to the measurement being trustworthy, not a
separate line.)_
**Gate:** NOT MET — p95 7.90 s / 7.00 s vs < 5 s (**1.6× over**, was **7.9×**).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n _(C-1 **did** explain the employee vantage — 39.45 → 7.90 s — so the trigger is satisfied, not fired)_ — (4) user-blocker: n — (5) cap-reached: n _(1 tik)_ — (6) protocol-stop: n — Outcome: **continue**

**Decisions:** D7 (server-only env var via demo-patch; every cheaper option provably exhausted), D8 (validate baked
constants by reading the **bundle**, not the image ENV), D9 (a green verdict is only as fresh as the image it
graded), D10 (the `BACKEND_API_URL` "twin" is client-side-only — **not** on the login path)

**Side-deliverables:** **F-6** (rext `0cfe23e`) — the cache-validator pk gate. Not in the iter's planned scope; it
was a **blocker to a trustworthy measurement** and is a standing safety fix in its own right (it stops a demo from
silently phoning production auth).

**Routes carried forward:**
- **→ iter-04 (the gate's remaining 2.9 s):** attribute + kill the **~6.10 s retry ladder** on the fast-failing SSR
  fetch. Handler `FIX-M218-iter04-ssr-retry-ladder`. **This is now the whole gap to the gate.**
- **F-7** the dormant `NEXT_PUBLIC_BACKEND_API_URL` blackhole (`PROBE-M218-backend-api-url-twin`).
- **F-5** ad-tech egress, **C-5** clerk-js vendoring, **C-3** re-probe (now finally exercisable — logins complete),
  **DOC-M218-audit-corrections** (the M43-D5 correction now has its real numbers), all iter-01/02 carry-forwards.

**Lessons:**
1. **A cache-validator can only see the constants it can read.** M211's validator read the image ENV, so it guarded
   the one baked constant that *happened* to be ENV — and was structurally blind to the one that **wires
   Clerkenstein**. *Validate a baked constant by reading the artifact it was baked into.* (→ protocol doc.)
2. **The same defect class bit us twice in one milestone.** C-1 *is* "a build-inlined `NEXT_PUBLIC_*` that a
   consumer cannot override"; F-6 *is* "a build-inlined `NEXT_PUBLIC_*` that a rebuild path failed to supply."
   Build-time inlining is the milestone's real antagonist, not any one variable.
3. **Rebuild through the sanctioned path, or not at all.** An out-of-band `docker build` that copies the visible
   `--build-arg`s but misses an invisible gitignored overlay produces an image that *looks* right, *passes* the
   guard, and is *silently wrong*. The 9-hour-old `autoverify.json` still said green (D9).
