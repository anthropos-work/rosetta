**Type:** tik  ·  **Active strategy:** TOK-01 (reachability-first, step 3)

# iter-04 — the gate is met

## THE NUMBER

_billion · demo-1 · rext `cue-to-cue-m218-iter04` · **cold reset-to-seed** · `autoverify.json` **green, 0
warnings** · measured **from the tailnet** · **5 samples/vantage, gate ARMED** (`LATENCY_GATE_MS=5000`)._

| vantage | M218 baseline | iter-03 | **iter-04** | ACCESS | gate < 5 s |
|---|---|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | p95 **39.45 s** | 7.90 s | **p95 1.46 s** (p50 1.00 s) | **5/5** | ✅ **MET** |
| **manager** (`dan-manager` → `/enterprise/…`) | p95 **38.30 s** | 7.00 s | **p95 1.40 s** (p50 1.12 s) | **5/5** | ✅ **MET** |

**27× faster.** Both Playwright specs **passed with the gate armed**. The `slow-body` anomaly — the signature of
the whole milestone — is **gone**.

## What it was

**Clerkenstein was lying to next-web about who the hero was.**

The mock has **two** identity surfaces, and they disagreed:

| surface | what it told the platform | correct? |
|---|---|---|
| **FAPI** — mints the JWT `app`'s `authn` authenticates | the hero's **real** eid (`40921b2e-…`) | ✅ |
| **BAPI** — answers next-web's **server-side** `currentUser()` | **`11111111-…`** — a stub, for **every** id | ❌ |

`clerk-backend/resources.go` said so in its own comment: **`// Disarmed: any id → the demo user.`** True when a
demo had **one** user. The **Stories & Heroes** model (M35) gave every hero their own eid — and this line was
never revisited.

next-web's `(authenticated)/layout.tsx` passes `user.externalId` — **the BAPI's value** — to GraphQL as the user
id. `app`'s resolver is an **identity equality check**, not an authorization one:

```go
if currentUser.ID() != userID {   // the JWT's hero  vs  the BAPI's stub
    return "", fmt.Errorf("user does not have permission to set studio preferences")
```

`prefetchUserStatus` rethrows under `retry: 2` / `retryDelay 2 s → 4 s`, so **every** authenticated render paid
`3 × ~33 ms + (2 s + 4 s) ≈ 6.0 s` of **pure backoff** — on **both** vantages, because both block on the same
shared layout.

### The chain closed from three independent directions

1. **the harness** — `[slow-body] … BODY took 6104 ms`, reproducible to **±3 ms** across both vantages;
2. **the Cosmo router** — the SSR calls (`user_agent: "node"`) logged at **ERROR**, HTTP 200, latency **2.4 ms**,
   at **20:12:16 → :18 → :22** — *the retry ladder, visible in the timestamps*;
3. **`app`** — at those same instants: `userPreferences user does not have permission to set studio preferences`,
   the sentry scope naming the **real** hero (`dan.rossi3@cervato-systems.com`) while the BAPI was handing
   next-web the stub.

**And the arithmetic had already predicted the shape before any of it was read:** 6.1 s **cannot** be a blackhole
(10.5 s/attempt), so it had to be a **fast-failing** fetch under the 2 s/4 s ladder. *The number told us what kind
of bug to look for.* That is what the iter-02 harness bought.

## The fix (rext `8ebc89e`)

The fake BAPI **already mounted the roster** (`FAKE_FAPI_ROSTER`) — it just used it only to seed **memberships**
(the studio-desk admin gate) and **never the user resource**. Now it also seeds each hero's **identity**
(`Store.SeedUserIdentity`), and `getUser` reports that hero's real eid/email/name.

Live proof on bring-up:
```
fake-bapi: seeded 9 roster identities + memberships from /roster/roster.json
           (SSR currentUser eid + the studio-desk admin gate)
```

## The alignment guard — discharged, not assumed

With **no roster** the identity map is empty and `userRes` falls back to the historical stub **byte-identically**.
That fallback *is* the alignment contract: the runner never mounts a roster. **Proven** — `/align-run` across
**all five** surfaces, after the change:

| surface | genes | overall | critical | divergences |
|---|---|---|---|---|
| `clerkrun` | 22/22 | **100.0%** | **100.0%** | 0 |
| `jsfapirun` | 9/9 | **100.0%** | **100.0%** | 0 |
| `multirun` | 9/9 | **100.0%** | **100.0%** | 0 |
| `deployrun` | 7/7 | **100.0%** | **100.0%** | 0 |
| `expressrun` | 13/13 | **100.0%** | **100.0%** | 0 |

Gate = critical ≥ 100 / overall ≥ 95. **All five exit 0. Nothing moved.** The milestone's BLOCKING clause names
`clerk-frontend/` and this change is in `clerk-backend/` — it was run **anyway**, because the `DistinctIdentity/*`
genes are `critical` and the critical gate has no partial credit. `expressrun` is dependency-gated (it needs a
`node_modules` carrying `@clerk/backend`); it was **pointed at a real one rather than left unscored** — an
unmeasured surface is not a passing surface.

## Side-deliverable — the M43-D5 correction, with its real numbers at last

The milestone's named doc debt, now landed (**Fate 1**):

- **`corpus/ops/demo/latency-budget.md`** — **NEW**, the milestone's declared **blind area**: the ACCESS
  definition, the per-leg attribution model, the streaming-SSR measurement trap, *how to read the retry ladder*,
  the measured before/after, and the harness contract.
- **`corpus/ops/demo/cockpit-spec.md`** (×2 sites) + **rext `demo-stack/cockpit.py`** (×3 sites) — the corpus said
  a login was *"~2–5 s, which we can't shorten."* **Both halves were false**: it was **39 s**, and **we shortened
  it to 1.4 s**. Corrected in place, with the numbers and the mechanism.

> **The lesson booked with it:** an unmeasured number that asserts its own unfixability is the most expensive kind
> of documentation there is. It was recorded as an M43 scope-`Out:` + decision **D5** with **zero deferrals**, so
> it never entered a ledger and was never revisited — **for four releases**. *That* is why nobody investigated.

---

## Close — 2026-07-13

**Outcome:** **THE GATE IS MET.** p95 click→ACCESS **1.46 s** (employee) and **1.40 s** (manager) against a **5 s**
gate — **5/5 ACCESS on both vantages, gate armed, on a cold reset-to-seed green stack**. From **39.45 s / 38.30 s**:
**27× faster**. The milestone's two defects were both in the **demo tooling**; **zero platform-repo edits**.

**Type:** tik
**Status:** closed-fixed
**Gate:** **MET** — employee p95 1.46 s · manager p95 1.40 s · both < 5 s · 5/5 ACCESS · gate armed · green stack.
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n _(2 tiks)_ — (6) protocol-stop: n — Outcome: **exit-1**

**Decisions:** D11 (the mock's two halves must agree about who the user is), D12 (the fallback **is** the alignment
contract — proven across all 5 surfaces), D13 (the measurement chain closes end-to-end; the arithmetic named the
bug's shape before the code was read)

**Side-deliverables:** the **M43-D5 correction** + **`latency-budget.md`** (the milestone's declared blind area).

**Routes carried forward (none block the gate):**
- **F-7** `NEXT_PUBLIC_BACKEND_API_URL` bakes to a **10.5 s blackhole from inside the container** — dormant *only*
  because every reader is client-side (D10). **A loaded gun.** `PROBE-M218-backend-api-url-twin`.
- **F-5** the demo phones **Google Analytics / DoubleClick / Google Ads / LinkedIn Ads** on every authenticated
  load. `FIX-M218-telemetry-egress`.
- **C-5** vendor clerk-js + bound the **unbounded** `Timeout: 0` (alignment-invisible ⇒ gate-free).
- **C-3** the parked router-retry probe — **now exercisable**, and the router *is* logging cms/Directus 403s
  (`getSkillPaths`, `_entities JobSimulation`) on the content path. Not on the login path; affects data-settle.
- The **CI-inert** correction (`alignment_testing.md:232,233,239`) + the `clerkenstein.md:3-4` header.

**Lessons:**
1. **A mock that is faithful on one surface and stubbed on another is not "partially faithful" — it is
   *inconsistent*.** The platform cross-checks its identity surfaces against each other; the mock must keep them
   **coherent**, not merely individually plausible. (→ protocol doc.)
2. **Read the ladder.** A blackhole and a refusal are **six seconds apart in signature** (37.5 s vs 6.1 s under the
   same `retry: 2` / 2 s→4 s policy). The *magnitude* of a latency tells you what *class* of bug to hunt, before
   you read a line of code.
3. **The whole milestone was two mocks disagreeing with reality** — an address the server couldn't reach, and an
   identity the mock made up. Neither was a platform bug. The demo tooling was slandering the platform for four
   releases.
