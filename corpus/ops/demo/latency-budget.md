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
**hero cards only** (`cockpit.py:1214`); the Content-stories tab's seat CTAs (`:882`) carry a bare `href`. So
`readCockpitCta` (`e2e/lib/latency.ts:115-127`) **throws before t0** for a content seat — there is no clock to
start. `run-latency.sh:53-59` independently hard-rejects any vantage outside `employee|manager|recruiter`
(`exit 2`). **A content-seat number cannot be produced by this harness at all**, so its absence is a property
of the instrument, not a gap in a run.

> **v2.5 "the playbill" scoping (user decision B2, 2026-07-20).** Content-seat latency is **explicitly OUT OF
> SCOPE for v2.5**. The cockpit CTA and `run-latency.sh` were deliberately **not** extended to content seats.
> The 31→29 content actions v2.5 shipped (later grown to 49 landable at v2.6 M241) are proven for **CONTENT** — they render real, non-empty results —
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
- **…and the ban on it must be a TOKEN scan, not a list of spellings (v2.8 M256 harden pass).** M256 iter-03
  widened the Playthrough harness's ban from "the four `/home` logins" to "the whole harness" and encoded it as
  two tightly-anchored regexes: `waitUntil:\s*['"]networkidle['"]` and
  `waitForLoadState\(\s*['"]networkidle['"]\s*\)`. Measured at the harden pass, **four plausible shapes score zero
  hits against that pair**, and two of the four are not hypothetical:
  - `waitUntil: opts.waitUntil ?? 'networkidle'` — the **coalesced default**, which is
    `stack-verify/e2e/lib/cockpit-login.ts:87` *verbatim*: the single line that is the **root cause of the whole
    class**. The pattern required a quote immediately after `waitUntil:`; a `??` default puts an identifier there,
    so the ban was blind to the origin of the bug it was banning.
  - `waitForLoadState('networkidle', { timeout: 4_000 })` — the **bounded settle**, ~20 occurrences one directory
    away (`persona-assert.ts`, `section-assert.ts`, `crawl.ts`, four `calibrate-*` specs). The pattern required
    `)` immediately after the closing quote, so **any** second argument disabled it — and `hero-login.ts` forwards
    into that very tree, so the two directories are one copy-paste apart.
  - plus double-quoted spellings and `const w = 'networkidle'` indirection.

  The ban is now a **token scan of comment-stripped code** — no arity, argument order or quote style to get wrong —
  with exactly **one enumerated allowance**: a `waitUntil?:` optional-property *type* declaration, which `?:` makes
  provably impossible to execute as a gate. **The general rule: ban the token, not the two spellings you happened
  to find** — a spelling list is a fence around the instances you already fixed. The scope exception
  (`stack-verify/e2e/**`, the coverage sweep, which uses a *ceiling-bounded* networkidle as one input to a presence
  heuristic **by design**) is now written down rather than implied by which directories the scanner happens to read.
- **It clears cookies per sample**, so each click is a genuine cold login.
- **curl cannot drive this flow** at all: the fake-FAPI validates `redirect_url` against the public origin, and
  next-web's middleware 307s any non-https origin. It **must** be a real browser on the real origin.

## The studio-desk first-paint budget (v2.7 "july jitter" M253)

The login budget above is next-web/hiring **ACCESS**. **studio-desk has its own, separate first-paint budget** —
authored by **M253** because `run-latency.sh` measures ACCESS on the two React apps and there was **no** studio
first-paint harness at all.

> **The gate (M253):** on a cold demo (**state the environment — laptop vs tailnet**), **first-meaningful-paint
> < 1000 ms** — the `.page-skeleton` header+sidemenu shell **visible** — **AND no blank > 1 s**, p95 over **5
> consecutive cold loads**, gated on a fresh-green `autoverify.json`. FMP here is defined as the **shell being
> visible**, not the browser's `first-contentful-paint` entry.

### Why studio blanks where next-web streams

studio-desk is **not** an SSR React app — it is an **empty-body MPA** (per HTML page: `home.html`, the builders,
…) whose `core/main.ts` builds the whole visible shell (`new PageWrapper()`) **only after three sequential
blocking `await`s**: `clerk.load()` → `l12nService.init()` → `userService.canAccess()`. Until `PageWrapper` runs,
the `<body>` is empty. So the blank is not a slow *render* — it is a **paint-ordering** defect: the shell is
built behind the boot awaits.

**The per-leg baseline (demo-2, LOCAL LAPTOP, authenticated as `maya-thriving`, t0 = studio navigation):**

| leg | cost | note |
|---|---|---|
| `clerk.load()` | **~140 ms** | NOT the 10 s timeout — cheap vs Clerkenstein (the milestone's stated worry, refuted) |
| `l12nService.init()` | **~12 ms** | cheap |
| **`userService.canAccess()`** | **~3.9 s** | the dominant leg — its org-memberships check **404s** and burns a 3-attempt retry ladder (1776 + 2102 ms backoff). **NB: M253 recorded this as a "GraphQL" 404 — that attribution is WRONG.** It is a **Clerk FAPI** 404 on `GET /v1/me/organization_memberships`, a route Clerkenstein never registered. Corrected + fixed at the source in `fix/studio` — see §"Time-to-usable" below |
| `new PageWrapper()` (shell) | — | runs only AFTER the three awaits → skeleton visible at **~4669 ms** |

**Read the arithmetic** (per this doc's rule): the ~3.9 s is a `retry: 2` ladder on a **fast-failing** fetch (a
404, not a blackhole) — the same signature family as the login budget's ~6.1 s. But the fix is NOT to chase the
404: it is to **paint the shell ahead of the await**.

> **⚠️ SUPERSEDED (`fix/studio`, 2026-07-27) — the 404 WAS worth chasing.** Painting the shell ahead of the
> await was necessary but **not sufficient**: it moved the *paint*, not the *wait*. The 404 kept costing
> **~4.05 s of time-to-usable on every load**, and the skeleton merely covered it. Both are now true — the
> shell paints early **and** the ladder is gone. Read the sentence above as *"first paint the shell, THEN
> chase the 404"*, never as a reason to leave it. **This is the fourth iteration of this doc's own recurring
> failure** (see the ~2–5 s "which we can't shorten" story): a metric that looked green because it measured
> the wrong thing.

### The fix — paint the shell before the awaits (two demopatches on the M249 studio ladder)

- **`studio-desk-shell-first-paint`** — inject the `.page-skeleton` DOM (header + sidemenu + content)
  **synchronously right after `preloadCriticalCSS()`** (main.ts ~L97), **before** Sentry/posthog/clerk.load/
  l12n/canAccess. The CSS for those classes is already injected by `preloadCriticalCSS`, so the dark shell
  paints from **CSS+DOM with zero network**. **De-dup is automatic:** `PageWrapper#init` wipes
  `document.body.innerHTML` then rebuilds its own skeleton, so the early shell is seamlessly replaced (no double
  skeleton). Auth-independent: it paints before `canAccess`, so the blank is closed regardless of the 404.
- **`studio-desk-no-thirdparty`** — no-op `Sentry.init` + `posthog.init` on the demo (no reachable GlitchTip / no
  PostHog project on a Clerk-free demo; the imports stay referenced by the later `captureException`/`identify`).

Both are sha-pinned demopatches on M249's `build_frontend_studio_desk` ladder (`demopatch-spec.md` §5); the
patch-set fingerprint grows 3 → **5**, forcing a studio rebuild. Zero platform-repo edits.

### The result (demo-2, LOCAL LAPTOP)

| | baseline | **post-M253** |
|---|---|---|
| skeleton-visible p95 (5 cold loads) | **4669 ms** | **817 ms** (p50 743, max 817) |

5/5 cold loads painted the shell, 0 login bounces — **numerically MEETS the < 1000 ms gate** (~5.7× faster).
State the environment with the number: the table above is a **local laptop demo-2**, not the tailnet.
**Confirmed LIVE on billion (M254 gate (f), tailnet peer, cold reset-to-seed):** app-side first-paint **p50
637–726 ms < 1 s** — the M253 fix holds. The cold p95 outliers (1443 / 2014 / 4943 ms, `reachedShell` always
true) were a coordinator-accepted **environmental tailnet-RTT jitter** disposition (the (b) precedent + the
"state the environment" rule), not a studio regression.

### Time-to-usable — what the skeleton gate does NOT measure (`fix/studio`, 2026-07-27)

**The M253 gate measures a static DOM+CSS paint with zero network.** It is green whether the app becomes
*usable* in 200 ms or 20 s. That is not a nitpick — it was **structurally blind to a 4.8 s regression**: the
skeleton metric barely moved (p50 **1819 → 1102 ms**) while the real shell arrived **4.8 s** later.

> **The rule this establishes: a first-paint gate is not a performance gate.** `.page-skeleton` is a
> *placeholder*; the presenter is waiting on `PageWrapper`. Always pair a paint metric with a **time-to-usable**
> metric, or the paint metric will hide exactly the regression it appears to guard.

**Root cause (the real one).** studio-desk's `core/main.ts` awaits `userService.canAccess()`, which calls
`clerk.user.getOrganizationMemberships()` — so **clerk-js (the BROWSER)** issues
`GET /v1/me/organization_memberships?paginated=true&limit=10&offset=0`. Clerkenstein had registered only the
**BAPI's** server-side twin (`GET /v1/users/{userID}/organization_memberships`), never the **FAPI** route the
browser calls. The FAPI 404'd → clerk-js's `const { data } = response` threw *"Cannot destructure property
'data' of undefined"* → `withClerkErrorHandling` burned its full 3-attempt ladder (**1606 + 2265 ms** backoff)
→ `canAccess()` fell into `catch → return true`. Because `PageWrapper` — and therefore `body.page-loaded`,
which **every** studio page waits on — is built only *after* that await, the whole app sat behind a 404 retry
ladder on **every** load.

**The fix** (rext `clerkenstein/clerk-frontend`, zero platform edits): serve the FAPI route. The data was
already assembled — `/v1/me` returns `userRes.OrganizationMemberships` — so the route just serves it in the
**paginated envelope** clerk-js destructures (`{response:{data,total_count}}`, **not** a bare array) and
honours `limit`/`offset` so clerk-js's paging terminates.

**Measured on billion** (tailnet VM, demo-1, seat `dan-manager`, cold loads):

| | before | **after** |
|---|---|---|
| `canAccess` leg | 4049 ms | **38 ms** (−99%) |
| browser FCP | 6936 ms | **2152 ms** (−69%) |
| 404s per load | 3 | **0** |
| retry backoff | 1606 + 2265 ms | **none** |

> **Side effect — the client access gate now actually enforces.** The 404 made `canAccess()` fail **OPEN**
> (`catch → return true`). It now evaluates the real membership role against studio's
> `STUDIO_ACCESS_ROLES` (`{admin, org:admin, content_creator, org:content_creator}`). **No reachable outcome
> changes**: studio-desk's *server-side* `checkEnterpriseAndAdmin` already 303s non-admin seats to the web app
> before `main.ts` runs, and every tooling seat is an admin (`run-studio-fcp.sh` defaults to `dan-manager`;
> the M252 studio Playthroughs drive `pt-manager`). Two consequences worth pinning: the role must keep Clerk's
> **prefixed** form (`org:admin` — a bare `admin` also passes, `basic_member` does not), and `total_count`
> must be present, because `!memberships?.total_count` redirects the seat to the web app.

### The harness

`rext stack-verify/e2e/run-studio-fcp.sh` — a studio sibling of `run-latency.sh`:

```bash
cd <stack>/rosetta-extensions/stack-verify/e2e
STUDIO_FCP_RUNS=5 STUDIO_FCP_GATE_MS=1000 \
  ./run-studio-fcp.sh 2 maya-thriving        # demo N, seat that can reach studio
```

- **Establishes a Clerkenstein session first** (a studio page 302-redirects an unauthenticated visitor to
  next-web `/login`) by navigating the **real cockpit [Log in as] CTA** — the same handshake `measureLogin` uses;
  the `__session` cookie is set on `localhost`, shared with studio's offset port.
- **Each sample is a fresh context** (cold cache + cookies ⇒ a genuine cold login).
- **Gates on `skeleton-visible`** (the shell's `.skeleton-header` + `.skeleton-sidemenu`): p95 < gate **AND**
  max ≤ gate (the "no blank > 1 s" clause is a per-sample max, not a percentile).
- Same green-gate + non-integer-`N` guard as `run-latency.sh`; **never** gates on `networkidle`.

**And its time-to-usable sibling — `rext stack-verify/e2e/tests/studio-ttu.spec.ts`** (`fix/studio`), the probe
that would have caught the above. It records the **real boot waterfall** rather than the injected shell:

```
nav commit → skeleton (the M253 metric) → clerk.load → l12n → canAccess → PageWrapper (the REAL shell)
```

- Reads `core/main.ts`'s **own console markers** (`"Initialize clerk"` / `"Initialize l12n service"` /
  `"Initialize user service"` / `"Initialize content service"`) as leg boundaries — no instrumentation patch.
- Records **every** request with status + duration, so a retry ladder is **visible rather than inferred**
  (`nonOk` + a `slowest` top-12 per run). This is what turned "studio feels slow" into a named 404 route.
- Reports **`dead_shell_gap_ms`** — real-shell minus skeleton — *the* number the M253 gate cannot see.
- Knobs: `STUDIO_TTU_{HOST,SCHEME,N,RUNS,IDENTITY,OUT}`. Read-only, zero platform edits.

## ⚠ A relative gate needs its NOISE FLOOR published next to it, or it is not falsifiable (v2.8 M256 iter-12)

M256's clause 1 was pinned as a **ratio**: the median per Playthrough must be **≤ 0.79×** a same-stack
baseline measured earlier in the milestone. Eleven iters reported it as MET (0.5434× · 0.6245× · 0.5950× ·
0.5652× · 0.6863×), each honestly reporting its own batch. **Nobody measured how much the number moves when
nothing changes** — and iter-12 did:

| statistic, six full-suite runs, one session, same host | min | max | spread | median (n=6) |
|---|---:|---:|---:|---:|
| the GATED figure (22 non-studio Playthroughs) | 0.5701× | 1.1121× | 1.95× | **0.8129×** |
| the CONTROL subset (16 specs unchanged since iter-03) | 0.5281× | 1.0762× | **2.04×** | 0.7063× |

The control subset is code **no iter touched**, and it varies by a factor of two. There is no trend — the
most recent run reads 0.529× and the oldest 0.528×, with the 1.076× extreme in between — so this is not
host degradation over a session but **variance the pinned statistic does not absorb**. A "median of 3
consecutive runs" can land anywhere between ~0.53× and ~1.08× depending on which three runs it catches, which
means the *verdict* was being sampled, not measured. At n=6 the gated figure is **0.8129× — outside the
gate.**

**Rules this produces, and they apply to every relative gate in the corpus:**

1. **Publish the spread with the median.** A ratio without a noise floor is a number, not a verdict. State
   min / max / n alongside it.
2. **Keep an untouched CONTROL subset and report it every time.** It is the only thing that separates "the
   work got faster" from "the box was quiet". M256's original-16 cross-check existed from iter-04 and is what
   made this diagnosable at all — its value was in *having* it, not in the reassuring readings it gave.
3. **Prefer a PAIRED measurement.** Measure the baseline in the **same batch** as the treatment. A baseline
   fixed hours earlier silently turns the ratio into a measure of host state.
4. **If n=3 cannot decide the gate, raise n or normalise within-run** (against an invariant leg such as the
   login handshake) — do not re-run until a favourable batch appears. Choosing the flattering denominator
   after the fact is the same defect as choosing the flattering run.

This does **not** retract the underlying speed work: iter-03's `networkidle` removal was measured **directly**
at the leg (2854 ms → 423 ms for the same navigation), not inferred from a suite ratio. Leg-level
before/after on the same page in the same run is exactly the kind of measurement this variance cannot fake —
which is the strongest argument for preferring it over suite-level ratios wherever it is available.

## See also

- [`cockpit-spec.md`](cockpit-spec.md) — the presenter cockpit (and the corrected M43-D5 claim)
- [`../verification.md`](../verification.md) — the green gate this measurement stands on
- [`../../services/clerkenstein.md`](../../services/clerkenstein.md) — the mock whose BAPI/FAPI must stay coherent
- [`demopatch-spec.md`](demopatch-spec.md) — the sanctioned hatch the SSR-origin fix went through
