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
**REPORTED, never gated**. This is **v2.3's `D-DESIGN-1`**, whose canonical statement is *"the < 5 s gate is on
**ACCESS**, not full first-page render"* — the clause above is its corollary, not a second decision.

> ⚠️ **Cite it as "v2.3's D-DESIGN-1", never bare.** v2.2 has its **own** `D-DESIGN-1` (*"public reach is never
> default-on"*, itself superseded by v2.3's D-DESIGN-3). The ids collide across releases; a bare reference
> resolves to the wrong decision. See [`../safety.md`](../safety.md) §3.5, which owns both glosses.

It sits behind a platform-side DataLoader defect
(`GetOrganizationTargetRole` ≈ 3 RPCs/member) that **cannot** be fixed under the zero-platform-edit constraint.
Gating on it would have made the milestone unwinnable for a reason that has nothing to do with login.

## The gate

**p95 click→ACCESS < 5 s**, measured over **HERO vantages** — `maya-thriving` (employee → `/profile`),
`dan-manager` (manager → `/enterprise/…`), and `rae-recruiter` (recruiter → the **apps/hiring** 2nd app
`/enterprise/activity-dashboard`; the M226 "opening night" 3rd measured path, v2.4 "casting call").

### What the gate covers, and what a release must state (v2.5 M236, user-authorized)

**The gate is scoped to HERO vantages only. Non-hero seats are out of scope, structurally — not by omission.**

`measureLogin` begins by reading the cockpit's real CTA, and the ACCESS predicate's second half ("the user menu
shows the hero") is resolved from the CTA's **`data-login-as`** attribute. The cockpit emits that attribute on
**hero cards only** (`cockpit.py:618`); the Content-stories tab's seat CTAs (`:423`) carry a bare `href`. So
`readCockpitCta` (`e2e/lib/latency.ts:115-127`) **throws before t0** for a content seat — there is no clock to
start. `run-latency.sh:53-59` independently hard-rejects any vantage outside `employee|manager|recruiter`
(`exit 2`). **A content-seat number cannot be produced by this harness at all**, so its absence is a property
of the instrument, not a gap in a run.

> **v2.5 "the playbill" scoping (user decision B2, 2026-07-20).** Content-seat latency is **explicitly OUT OF
> SCOPE for v2.5**. The cockpit CTA and `run-latency.sh` were deliberately **not** extended to content seats.
> The 31→29 content actions v2.5 shipped are proven for **CONTENT** — they render real, non-empty results —
> and are **not formally timed**. Do not read "v2.5 met the p95 gate" as covering the content seats.

**Two run-shape variables are per-release, and a release must state both with its number:**

| variable | what varies | v2.5 M236 |
|---|---|---|
| **vantages measured** | which of the three hero seats a release actually drove | **2** — employee + manager (the recruiter vantage was last measured at M226) |
| **cold cycles** | how many *distinct cold reset-to-seed stacks* the samples came from | **1** cold stack, **5** login samples within it (`LATENCY_RUNS=5`) |

**These are not the same axis, and conflating them is how a gate silently weakens.** M218 armed the gate on
**5 consecutive cold reset-to-seed runs** — five separate cold stacks. M236 measured **5 login samples inside
one cold stack**, which is a weaker claim: it samples login variance, not bring-up variance. Both are legitimate
readings; only one of them is the M218 standard. **A release claiming "gate MET" must say which it did.**

> **Standing rule:** *"5 runs" is ambiguous and must never appear unqualified.* Say **cold cycles × samples per
> cycle**. A number that does not name its cold-cycle count cannot be compared to one that does.

> **The recruiter vantage is a seat-key + a landing origin, not a new code path.** `measureLogin` is
> vantage-agnostic: it follows the cockpit CTA's own `redirect_url`, and the `rae-recruiter` CTA lands on the
> **hiring app** (`:3001+offset`, the TOK-02 two-app demo), which satisfies the same ACCESS predicate
> (loader-gone + the hero's identity in the header/nav). Add a vantage by adding its `case` to `run-latency.sh`.
> **Prerequisite (M226 Finding-1):** the hiring app port (`:3001+offset`) must be **fronted over `tailscale
> serve`** for the recruiter to be reachable from a tailnet peer — it was added to `gen_tailscale_serve.py`'s
> `UI_BROWSER_FACING` at M226 (it had been reachable only on localhost, so the recruiter vantage was dead
> cross-machine until then — the M215/M221 "last breakage is cross-machine" lesson). See `tailscale-serve.md`.

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

#### The per-item fan-out signature — cost that scales with CONTENT, not a broken route

The retry ladders above are **fixed** costs: the same number whatever the page holds. The third signature in
this family is the one that **varies with the item count of the thing being rendered**:

> **Two instances of the same route, one slow and one fast, differing only in item count ⇒ a per-item fan-out.**
> Not a broken route, not a dead seat, not a wrong id — the surface and the credentials demonstrably work,
> because the light sibling renders on them.

The diagnostic is a **contrast**, not a measurement: find the *sibling*. M236 iter-06 hit a skill-path manager
route exceeding a 180 s navigation timeout while `sp-genai-in-progress` rendered fine on the **same route
family with the same seat** — the two differed only in weight (a completed **13-chapter** path vs a 3-chapter
path at 45%). That contrast alone says *look for a query issued per chapter*, and rules out the three
hypotheses an operator reaches for first.

**Read it against the fixed-cost ladders above:**

| shape of the number | what it can only be |
|---|---|
| a **constant** ~37.5 s / ~6.1 s regardless of page content | a **retry ladder** — blackholing vs fast-failing (above) |
| **scales with the item count** on the page; a light sibling passes | a **per-item fan-out** — a query inside a loop |
| **large and cold, small and warm**, same page | a **warm-up transient** — see R4 below, not a gate violation |

**The order matters: name the arithmetic signature *before* reading code.** All three are distinguishable from
the number and one contrast, and each sends you to a different file.

> ⚠️ **But first, disbelieve the clock.** A per-item fan-out and a **mis-instrumented wait** produce the same
> reading. The same M236 pair later turned out to be *neither* — instrumenting the navigation showed **134
> completed legs, 0 pending, none over 800 ms, page painted in ~1 s**. The "hang" was `networkidle` never
> resolving against next-web's long-polls (see the rule above: *never gate on `networkidle`*). **Prove the
> page is actually slow before attributing the slowness** — a probe that measures the wrong event reports an
> arithmetic signature it invented. `coverage-protocol.md` records that pair's full triage.

## The measured baseline (and what M218 did to it)

_`billion` tailnet demo · `demo-1` · **cold** reset-to-seed · `autoverify.json` green · measured **from the
tailnet**, which is the presenter's actual vantage._

| vantage | pre-M218 | **post-M218** | factor |
|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | p95 **39.45 s** | **p95 1.46 s** (p50 1.00 s) | **27×** |
| **manager** (`dan-manager` → `/enterprise/…`) | p95 **38.30 s** | **p95 1.40 s** (p50 1.12 s) | **27×** |

**5/5 runs reached ACCESS on both vantages, gate armed.**

**M226 "opening night" — the recruiter 3rd vantage (v2.4 "casting call"), measured live on `billion` from the
tailnet peer, over 2 clean default cold reset-to-seed cycles:**

| vantage | measured | |
|---|---|---|
| **recruiter** (`rae-recruiter` → apps/hiring `/enterprise/activity-dashboard`) | **p95 1.09 s** (cycle 1) / **2.36 s** (cycle 2), p50 ~0.66 s | ACCESS 5/5 both cycles |

The recruiter shares next-web's fast authenticated-shell path — its p95 sits alongside employee/manager, well
under the 5 s gate. (State the environment: measured from this Mac against `billion.taildc510.ts.net` HTTPS.)
**Independently re-verified at M226 close** by the orchestrator from this Mac — a fresh recruiter-vantage run
returned **p95 1.74 s** (< 5 s), corroborating the two-cycle numbers above.

**M236 "prove on billion" (v2.5 "the playbill") — the gate re-measured on the tailnet, and the COLD/WARM pair:**

| vantage | **cold** (iter-10) | **warm** (iter-09) | ratio |
|---|---|---|---|
| **employee** (`maya-thriving` → `/profile`) | **p95 1.22 s** | p95 **3.15 s** | 2.6× |
| **manager** (`dan-manager` → `/enterprise/…`) | **p95 1.51 s** | p95 **2.71 s** | 1.8× |

_Scope: **1 cold reset-to-seed cycle × 5 login samples**, HERO vantages only (B2). ACCESS 5/5 both vantages.
Environment: measured from this Mac against `billion.taildc510.ts.net` over HTTPS — the presenter's vantage._

### ⚠ The COLD stack was the FAST one — the intuition is backwards here

**Do not read 1.22 s as expected steady state, and do not treat "warm" as a synonym for "fast".** The warm
readings came from a stack that had been **up for hours across 3 cockpit restarts and 2 re-pins**; the cold
ones from a stack built from nothing. Cold measured **~2× faster**.

Long-lived demo state **accumulates cost** — restarted cockpits, re-pinned tooling, and hours of accreted
process state are not a neutral background. So the two readings measure genuinely different subjects:

- **the cold number is the gate number** — it is what a presenter meets on a freshly brought-up demo, and it
  is the reproducible one;
- **the warm number is the PESSIMISTIC bound** — carry it as the ceiling a long-running demo can drift to.

Both are far inside the 5 s gate, which is the substantive result: **the gate holds at either end of the
range, so the cold/warm question does not change the verdict** — it only changes which number you quote.

> **This pair is the doc's own rule paying out.** *State the environment with every number* (below) is not
> bookkeeping: absent the environment, "p95 1.22 s" and "p95 3.15 s" look like a regression or a fix. They are
> neither — they are **two different stacks**. Every latency row in this doc names cold/warm, the vantage, the
> measuring host, and the cold-cycle count for exactly this reason. A row missing any of them is not
> comparable to the rows around it.

### R4 — the compare-drawer cold first render is a warm-up transient, NOT a gate violation

R4 was carried from M224 as a **blocks-milestone** risk: *would the 45×5 whole-org hydration on the
candidate-comparison drawer be too slow?* The M226 live finding on `billion`: the drawer's **COLD / idle first
render is genuinely slow** — **~2.5 min for the first sim's drawer** on a stone-cold stack — but it **warms to
~2.4 s** once the RSC/data path is hot, and it **does not violate any of the 7 gate conditions**:

- **C2 (the render probe)** gates on **data-present-and-renders** — page-1 rows painted (20/sim), network total
  ≥ 40, junk = 0, 0 prod-ejects — **not on render latency**. The cold transient is absorbed because the probe's
  per-test budget is **env-tunable** (`RENDER_TEST_TIMEOUT_MS`, default 300 000; landed at M226 `19d1159`). A cold
  or tailnet-fronted measurement needs a **cold-appropriate budget** so a slow-but-correct first render can't
  **false-fail** the probe. Set it generously when measuring cold/remote; the default already carries the
  documented headroom.
- **C5 (the p95 < 5 s gate)** is on **login → ACCESS** (the recruiter reaching her authenticated Results shell),
  **not** on the drawer drill-down. The slow compare-drawer cold render therefore does not count against C5.
- The transient is **warm-up work the bring-up autoverify already absorbs** — the set-dress verify drives the
  surfaces once during bring-up, so by the time a presenter clicks, the path is warm. R4 is a **cold-start
  property of the first drill-down**, not a standing latency the gate measures.

Net: R4 is **not** a milestone blocker — it is a documented cold-start transient with a probe budget wide enough
to measure through it. (If a future release wants the drawer's *drill-down* render itself under a p95 gate, that
is a **4th** measured path — a new vantage on the drawer, not the login — and would follow the same harness.)

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
LATENCY_SCHEME=https \                           # REQUIRED for a --public-host demo (see below)
LATENCY_AUTOVERIFY_JSON=/tmp/autoverify.json \   # a copy of the REAL remote verdict — never a bypass
LATENCY_RUNS=5 LATENCY_GATE_MS=5000 \            # gate armed — 5 SAMPLES, not 5 cold cycles
  ./run-latency.sh 1 employee                     # vantages: employee | manager | recruiter
```

> **`LATENCY_RUNS=5` buys 5 login samples on ONE stack — it is not the "5 cold runs" of the M218 standard.**
> Cold cycles are the *outer* loop and the harness does not own it: to measure N cold cycles you tear the
> stack down and bring it up N times, running this command once per cycle. Whichever you do, **state it with
> the number** (see *What the gate covers* above). `run-latency.sh` accepts exactly the three hero vantages
> and exits 2 on anything else — content seats are not measurable here by construction.

> **`LATENCY_SCHEME=https` is not optional here** (added M236 iter-09; the block above omitted it and was
> wrong for the exact scenario this section is about). The runner defaults to `http`, but a `--public-host`
> demo is HTTPS-fronted by `tailscale serve`, so the default gets a 400/redirect and fails at
> `readCockpitCta`. Localhost stays `http`.
>
> **Producing the remote verdict:** `autoverify.sh` only writes `autoverify.json` when **`STACK_DIR` is set
> in its environment**. Run without it and it prints a full green report and writes nothing — which looks
> exactly like success:
> ```bash
> ssh <box> 'STACK_DIR=<stack>/rosetta-extensions/demo-stack/stacks/demo-1 \
>   <stack>/rosetta-extensions/stack-verify/live/autoverify.sh --project demo-1 --offset 10000'
> ```
>
> **Run it from a second machine on the tailnet, not on the demo host.** The gate is a *presenter-vantage*
> number; measuring on the box measures something nobody experiences.

Contract:

- **It drives the REAL cockpit CTA** — it reads the live `<a class="btn login">` off the cockpit and clicks it. A
  stale or host-drifted cockpit therefore **fails the probe** instead of being measured around.
- **It refuses to measure a stack that is not green** (`autoverify.json`). A latency number off a broken stack is
  noise. For a remote stack, point `LATENCY_AUTOVERIFY_JSON` at a **copy of the real remote verdict** — the gate
  still grades the real stack. *A safety gate that is inconvenient in the exact situation it exists for will be
  switched off — so make it work there instead.*
- **It ages the verdict** (4 h window) so a verdict cannot outlive its subject — the F-6 hazard, where a
  nine-hour-old verdict graded a Clerkenstein-dewired stack green.
  > **M236 iter-09 found that age check reading UTC as local time.** `autoverify.sh` writes `ts` in UTC with
  > a trailing `Z`; the BSD (`date -jf`) fallback parses in the **local** zone, so on macOS the age was off
  > by exactly the UTC offset — a verdict **121 s** old aged as **7321 s** on a UTC+2 grader. East of UTC
  > that fails closed; **west of UTC it inflates the window and reads a STALE verdict as FRESH**, which is
  > the very hazard the check exists to prevent. Fixed with `TZ=UTC` on that branch.
  >
  > The general lesson is worth more than the fix: **a freshness guard that fails open is worse than no
  > guard, because everything downstream trusts it** — and this one was itself introduced by a hardening
  > pass (M218 F-10). Code written to close a hazard is not exempt from that hazard.
  >
  > **Now regression-tested** (M236 final harden): `stack-verify/tests/test_green_gate_age.py` extracts the
  > shipped `v_epoch=` line and evaluates it under five zones spanning both sides of UTC — including a
  > **half-hour offset**, which a "subtract whole hours" patch would still get wrong — asserting the parsed
  > epoch is identical **and** equals the true UTC instant. Zone-independence alone would be satisfied by a
  > consistently *wrong* constant, so both halves are needed. It also sweeps the whole `e2e/` section for
  > any **unpinned `date -jf`**, because the bug is a class, not an instance. **Mutation-verified:**
  > removing `TZ=UTC` turns 5 of the 6 guards red. *The fix shipped without a test; a fix to a guard is
  > exactly where a test is least optional.*
- **It refuses a stack number it cannot trust.** `OFFSET=$(( N * 10000 ))` and bash evaluates a non-numeric
  `N` to **0, silently** — so `./run-latency.sh abc` pointed every probe at offset 0, the **dev stack's**
  ports, and would have reported those timings as demo-N's. A grader whose premise is *refuse to measure a
  stack that is not what it claims to be* must not be able to measure a **different** stack without saying
  so. Non-integer `N` now exits 2 (M236 final harden). `run-coverage.sh` and `run-hiring-render.sh` share
  the arithmetic and were **guarded at the M236 close** — all four runners now refuse a non-integer `N`
  rather than silently sweeping the DEV stack at offset 0.
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
