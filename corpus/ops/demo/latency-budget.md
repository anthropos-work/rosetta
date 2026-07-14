# The demo login latency budget

_The click→ACCESS budget, its per-leg attribution model, the measured baseline, the gate, and the harness that
grades it. Authored by **M218 "seat change"** (v2.3 "cue to cue") — before it, the project had **no** perf budget,
**no** baseline, and **no definition of "access"** anywhere in `corpus/**` or `rosetta-extensions/`._

> **Why this doc exists.** For four releases the corpus asserted that a cockpit login took *"~2–5 s, which we
> can't shorten."* **Nobody had ever measured it.** It was **39 seconds**, and it was **shortenable**. An
> unmeasured number that asserts its own unfixability is the most expensive kind of documentation there is.

---

## ACCESS — the definition

> **ACCESS** := the authenticated shell is **rendered and interactive** with the hero's **identity present** —
> the full-screen loading state is gone **and** the user menu shows the hero.

Not "the document responded". Not "the page painted". The presenter is **in**, as the hero.

**In-page data-completion** (the 200-member grid finishing its fan-out) is a **separate, secondary** metric —
**REPORTED, never gated** (**D-DESIGN-1**). It sits behind a platform-side DataLoader defect
(`GetOrganizationTargetRole` ≈ 3 RPCs/member) that **cannot** be fixed under the zero-platform-edit constraint.
Gating on it would have made the milestone unwinnable for a reason that has nothing to do with login.

## The gate

**p95 click→ACCESS < 5 s**, for **both** vantages — `maya-thriving` (employee → `/profile`) **and** `dan-manager`
(manager → `/enterprise/…`) — over **5 consecutive cold reset-to-seed runs**.

## The per-leg attribution model

A login is not one number; it is a chain. Attribution is the whole point — a total tells you *that* it's slow,
the legs tell you *what to fix*.

| leg | what it is |
|---|---|
| `handshake` | the cockpit's `<a href>` → fake-FAPI `/v1/client/handshake` → **303** |
| `ssr-document` | next-web renders the authenticated route **server-side** (the `force-dynamic` layout blocks here) |
| `clerk-js` | the browser fetches the Clerk bundle (proxied by the fake FAPI) |
| `fapi-client` | clerk-js talks to the FAPI (`/v1/client`, `/v1/environment`) |
| `data-query` | the client's own GraphQL queries — these **cannot start until SSR unblocks** |
| **ACCESS** | the definition above |

### ⚠ Measure the response **BODY**, not the response **headers**

**The single most important rule in this doc.** Next.js App Router **streams** the RSC payload: the shell flushes
**immediately** (HTTP 200, headers in ~70–150 ms) while the server render is still **blocked** awaiting its data.
Playwright's `response` event fires on **headers**.

⇒ A headers-only probe reports a **fast document** while the body trickles for **37 seconds**, and mis-attributes
a blocked SSR to a phantom *client-side* gap. M218 iter-02 walked into this and out of it; the harness now records
`bodyAtMs` via `response.finished()` and raises a **`slow-body`** anomaly.

> *Any* latency probe against a streaming SSR framework that watches `response` and not `response.finished()` will
> lie to you in exactly this way.

### Read the arithmetic — the number tells you what kind of bug to look for

M218's two defects were each identifiable **from their cost alone**, before any code was read:

| observed | what it can only be |
|---|---|
| **~37.5 s** | `3 attempts × 10.5 s` (undici's connect timeout) `+ (2 s + 4 s backoff)` ⇒ a **blackholing** address |
| **~6.1 s** | `3 attempts × ~33 ms + (2 s + 4 s backoff)` ⇒ a **fast-failing** fetch — an *error*, not a timeout |

Both fall out of next-web's `prefetchUserStatus`: `retry: 2`, `retryDelay = min(2000 × 2^n, 20000)` → **2 s → 4 s**.
A blackhole and a refusal are **six seconds apart in signature**. Learn to read the ladder.

## The measured baseline (and what M218 did to it)

_`billion` tailnet demo · `demo-1` · **cold** reset-to-seed · `autoverify.json` green · measured **from the
tailnet**, which is the presenter's actual vantage._

| vantage | pre-M218 | **post-M218** | factor |
|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | p95 **39.45 s** | **p95 1.46 s** (p50 1.00 s) | **27×** |
| **manager** (`dan-manager` → `/enterprise/…`) | p95 **38.30 s** | **p95 1.40 s** (p50 1.12 s) | **27×** |

**5/5 runs reached ACCESS on both vantages, gate armed.**

**State the environment with every number.** The *same* defect cost **~6 s on a laptop** and **~112 s on the
tailnet VM** — which is precisely why four releases of local measurement never saw it. **A latency number without
its environment is not a measurement.** Measure from the vantage that hurts.

### The two defects (both in the demo tooling; **neither in the platform**)

1. **The SSR GraphQL origin was the build-inlined public URL** (**~37.5 s**). `NEXT_PUBLIC_*` is build-inlined, so
   *one* constant served two consumers with incompatible reachability: the **browser** needs the public origin, the
   **SSR pass** needs a container origin. From inside the container the public address **blackholes** (DNS resolves;
   the TCP connect is dropped), so undici's 10.5 s connect timeout fired — three times, on every authenticated
   render, on **both** vantages (they share the authenticated layout). Fixed with a **server-only**
   `WUNDERGRAPH_SSR_ENDPOINT` (deliberately *not* a `NEXT_PUBLIC_*` name — so it is a **real runtime read**),
   supplied by `stack-injection` and taught to `server.graphql.ts` by a sha-pinned demo-patch.
2. **Clerkenstein's fake BAPI served a hardcoded stub user to every hero** (**~6.1 s**). The FAPI's JWT carried the
   hero's real internal id; the BAPI's `currentUser()` returned `11111111-…` for *anyone*. next-web passes the
   BAPI's value to GraphQL as the user id, so `app` compared the two, refused `userPreferences`, and the retry
   ladder above did the rest. Fixed by making the BAPI **roster-aware**.

**The generalizable one:** *a mock that is faithful on one surface and stubbed on another is not "partially
faithful" — it is **inconsistent**, and the platform cross-checks the surfaces against each other.*

## The harness (how to grade the gate yourself)

`rosetta-extensions/stack-verify/e2e/` — a **new stack-verify surface**, deliberately **not** a Playthrough
(Playthroughs declare performance an explicit **non-goal**).

```bash
cd <stack>/rosetta-extensions/stack-verify/e2e
LATENCY_HOST=billion.taildc510.ts.net \
LATENCY_AUTOVERIFY_JSON=/tmp/autoverify.json \   # a copy of the REAL remote verdict — never a bypass
LATENCY_RUNS=5 LATENCY_GATE_MS=5000 \            # gate armed
  ./run-latency.sh 1 employee                     # …and: manager
```

Contract:

- **It drives the REAL cockpit CTA** — it reads the live `<a class="btn login">` off the cockpit and clicks it. A
  stale or host-drifted cockpit therefore **fails the probe** instead of being measured around.
- **It refuses to measure a stack that is not green** (`autoverify.json`). A latency number off a broken stack is
  noise. For a remote stack, point `LATENCY_AUTOVERIFY_JSON` at a **copy of the real remote verdict** — the gate
  still grades the real stack. *A safety gate that is inconvenient in the exact situation it exists for will be
  switched off — so make it work there instead.*
- **It never gates on `networkidle`** — next-web holds never-idle long-polls. Every wait is **content-presence**
  polling.
- **It clears cookies per sample**, so each click is a genuine cold login.
- **curl cannot drive this flow** at all: the fake-FAPI validates `redirect_url` against the public origin, and
  next-web's middleware 307s any non-https origin. It **must** be a real browser on the real origin.

## See also

- [`cockpit-spec.md`](cockpit-spec.md) — the presenter cockpit (and the corrected M43-D5 claim)
- [`../verification.md`](../verification.md) — the green gate this measurement stands on
- [`../../services/clerkenstein.md`](../../services/clerkenstein.md) — the mock whose BAPI/FAPI must stay coherent
- [`demopatch-spec.md`](demopatch-spec.md) — the sanctioned hatch the SSR-origin fix went through
