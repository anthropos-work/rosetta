# M256 · iter-01 — decisions

## D1 — The re-cut gate DISSOLVED the parallelism requirement; the Clerkenstein enabler leaves the critical path

**Claim under test** (`../overview.md` § Open questions): *"What does the parallel-lane enabler cost?
… Answer at iter-01: without it, gate clause 1 is unreachable."*

**Verdict: FALSE, as of D-v28-12.** Clause 1 is now **median per-Playthrough ≤ 0.79× a same-stack
baseline**, and the suite wall-clock is **REPORTED, not gated**. Worker count does not change how long an
individual test takes — it changes how many run concurrently. On this box (Docker VM ~9.7 GiB, one browser
per worker, one shared Postgres) added workers would leave the median flat at best and raise it under
contention. The enabler is a **wall-clock** lever, and wall-clock is no longer gated.

**Priced anyway, for the record** (both options rext-owned, zero platform edits):

- *Cookie/`__client`-scoped registry.* The fake FAPI holds one `reg`, one `signedIn`, one `sessID` on a
  single `Server` (`clerkenstein/clerk-frontend/server.go` § `type Server`), `handleSelectIdentity` clears
  `signedIn`/`sessID` **globally**, and `handleMe` resolves `s.activeUserLocked()` with **no cookie input**.
  Making the seat per-client means threading a client identity through every handler's state — i.e.
  re-shaping the mirror engine's core state model. Clerkenstein is **Alignment-DNA-gated**, so the change
  also has to be re-scored (`/align-run`). High effort, real fidelity risk.
- *One fake-FAPI per worker.* The browser reaches the FAPI at a host **baked into the publishable key** at
  next-web **image build** time. A second FAPI on a second port implies a second pk implies a second image
  — a per-worker image build, which is precisely the cost v2.8's sibling milestones are trying to remove.

**Consequence:** neither is attempted in M256. **Fate 3 → routed forward** as a wall-clock (not median)
optimisation, with this pricing attached, so a future milestone starts from the answer rather than the
question. The overview's open question is answered, not deferred.

**What is kept from the thread:** the machine-checked per-spec **`MUTATES` / `READ-ONLY` / `UNKNOWN`** tag
the overview asks for. It is cheap, it makes the partition honest, and it is the artifact any future lane
consumes. Measured partition: **10** explicit "(no mutation)", **2** explicit MUTATES
(`skillpath-legacy.spec.ts:21`, `assignment-assign.spec.ts:18`), **6** UNCLASSIFIED — confirming the plan
review's count exactly.

## D2 — Org-admin is the first cluster, because it discharges clause 2 and clause 3 in one body of work

All four curated org-admin use cases declare a **persist-then-observe** final
(`org-admin-settings.{roles,members,tags,feature-config}.UC1` in the M201 curated corpus) — the
`pt-assignment-assign` write-then-read-back shape. Landing them gives **4 mutating Playthroughs**; with the
existing `pt-assignment-assign` that is **5**, which is exactly clause 2's `≥ 5 mutating` floor — while
simultaneously being half of clause 3's coverage scope (D-v28-4).

Onboarding, by contrast, is annotated in the curated corpus itself with **`# SEED GAP: fresh
pre-onboarding actor`** on three of its five UCs, and
`onboarding.enterprise-workforce-ai-readiness.UC1` carries M201's verify note *"no member-facing
AI-readiness flow exists (manager-only)"*. `pt-world` seeds **post**-onboarding users. So onboarding is one
seed question gating up to five UCs, with at least one plausible `unimplementable`.

**Ordering rule adopted:** org-admin **before** onboarding, so a seed wall cannot starve the clauses
org-admin already discharges.

## D3 — The clause-1 lever is the residual `networkidle`, not parallelism

`stack-verify/e2e/lib/cockpit-login.ts` defaults `waitUntil` to **`'networkidle'`**, and its own doc
records why that is wrong for this app: next-web holds **never-idle long-poll** connections, so
`networkidle` "resolves late and for the wrong reason". `playthroughs/e2e/lib/hero-login.ts:51` forwards
`waitUntil` **only when the caller sets it** — and only **6 of 18** browser specs set it. **12 logins
inherit the bad default.** M254 iter-10 measured 13 min → 3.8 min from exactly this class of fix, so the
mechanism is not speculative.

Two further holes: `playthroughs/e2e/lib/skill-path-page.ts:31` and `simulation-page.ts:36` still pass
`waitUntil: 'networkidle'` on their `goto`, while the **base** `PageObject.goto`
(`page-object.ts:48`) correctly uses `domcontentloaded` **and is fenced by a unit test** —
`tests/page-object.unit.spec.ts` guards the base class, so it structurally cannot see a per-surface
override.

**Fence to widen:** `tests/home-login-networkidle.unit.spec.ts` already fails closed on `/home`-landing
logins. Its scope must widen from "`/home`-landing specs" to **every `loginAsHero` call site and every
page-object `goto`** — otherwise the same defect re-enters on the next surface, which is how these two
overrides survived M254.

## D4 — ~~The `blocked` outcome needs no seed work~~ — **REFUTED by the Phase-0b audit, same iter**

**What I wrote first, from `seed-worlds.yaml` alone:** it declares the hero **`pt-free`**, the **`free`**
entitlement, and the **`entitlement-gated`** capability ("a free user cannot open paid content"), so clause
2's `≥ 1 blocked` outcome costs a use case + a spec, not a seed extension.

**Why that was wrong.** The Phase-0b KB-fidelity audit read the *seeder*, not the declaration, and found
`actor.entitlement` is **declared-only**: `blueprint.TierMix` is parsed, defaulted and validated but
consumed by **no seeder**, so **no tier ever reaches a DB column**, and `pt-world.seed.yaml` declares no
`tier_mix` at all. The `pt-free` seat exists and is *annotated* "entitlement-gate use cases — outcome:
blocked", but it is **not tier-gated** and is referenced by **0 of 18** use cases. A `blocked` Playthrough
built on it today would assert a refusal the platform has no reason to make.

**This is the Phase-0b gate paying for itself inside the iter that ran it** — a declaration in a YAML index
is not a seeded property, and `ptvalidate`'s precondition-coverage check resolves the *name*, not the
*column*. Recorded rather than quietly overwritten, because "the seed-worlds index says so" is exactly the
reasoning that has to stop being trusted here.

**Adopted instead (three candidate refusal surfaces, in cost order):**
1. **An RBAC / Sentinel deny** — a hero without a casbin grant reaching a gated surface. Precedent exists:
   M203 iter-05 showed the AI-sim launch renders an org-member **deny modal** without the g3
   `FEATURE_JOB_SIMULATIONS` grant. Withholding a grant is a *subtraction* from the seed, not a new
   mechanism.
2. **A cross-org access refusal** — `multi-org-private` is already a declared, seeded capability.
3. **A validation `error` outcome** on one of the org-admin writes (a write the platform correctly rejects)
   — free once org-admin lands.

Path 1 is the first candidate because it needs no seeder change. The decision of which one lands is the
owning tik's, on live evidence; what iter-01 settles is that **`pt-free` + `entitlement` is not the path**.

## D5 — `autoverify` check (d) FALSE-ALARMS on macOS: "NOBODY CAN LOG IN" while everybody can

The bring-up closed with `⚠⚠ autoverify demo-2: 1 check(s) FAILED` — specifically
*"the Clerkenstein fake-FAPI is NOT answering on :25400 — NOBODY CAN LOG IN to this demo."*

**The FAPI is healthy.** Evidence, in order:

- `docker logs demo-2-fake-fapi-1` → `loaded 35-identity roster … listening on :5400 (TLS)`, container
  `Up`, port published `0.0.0.0:25400->5400`.
- `openssl s_client -connect 127.0.0.1:25400` → **full handshake**, valid mkcert leaf
  (`NotAfter Oct 23 2028`), cert/key modulus MD5 **match**, SANs `DNS:localhost, 127.0.0.1, ::1`.
- **Chromium** (Playwright, `ignoreHTTPSErrors`) `GET https://localhost:25400/v1/environment` → **200**,
  body `{"response":{"object":"environment","id":"env_clerkenstein",…}`.
- `curl` (macOS system curl, **LibreSSL/3.3.6**) → `error:06FFF064:digital envelope
  routines:CRYPTO_internal:bad decrypt`, exit code 0-bytes / HTTP `000`.

`stack-verify/live/autoverify.sh` check (d) probes with `curl -fsSk … https://…/v1/environment`, falling
back to plain `http://` (which the TLS server correctly rejects — the container logs
*"client sent an HTTP request to an HTTPS server"* at exactly the probe timestamp). So **both** legs fail on
a LibreSSL host and the check warns on a working stack.

**Two consequences, both load-bearing for this milestone:**
1. This warning must **not** be read as a blocker for the rest of M256 — the browser, which is the only
   client that matters for a Playthrough, handshakes fine. Stated here once so it is not re-diagnosed.
2. The check **cries wolf on every local bring-up**, which is how a real FAPI outage would get ignored. It
   is the same defect class as the M236 BSD-`date` green-gate bug (`latency-budget.md`): a guard whose
   verdict is decided by the *host toolchain* rather than the system under test.

**Fate 3 → routed forward** to a later tik of this milestone (`FIX-M256-autoverify-fapi-libressl`): give
check (d) a probe that does not depend on the host's TLS stack. Not fixed in iter-01 — it is unplanned
scope for a bootstrap tok, and the scope-creep tripwire applies.
