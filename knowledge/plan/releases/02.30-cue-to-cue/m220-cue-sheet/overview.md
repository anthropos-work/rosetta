---
milestone: M220
slug: cue-sheet
version: v2.3 "cue to cue"
milestone_shape: section
status: planned
created: 2026-07-13
last_updated: 2026-07-13
complexity: medium
depends_on: M217
parallel_with: M218, M219
delivers: "/demo-up MEANS what the user thinks it means — full data + 3 orgs + remotely reachable BY DEFAULT (opt-out); dev stays opt-in. Plus corpus/ops/safety.md Part 3 — the exposure side (a whole axis the safety contract does not cover) + the correction of a FALSE safety claim in tailscale-serve.md"
issues: "the docs say '2 orgs'; the code ships 3 — that 4-line lie is why the user believed the seeding ask was unmet. 'Pull all data' is ALREADY default-on; the real failure is a COLD CACHE, not a default. And tailscale-serve.md:405-407 falsely claims 0.0.0.0 binding is gated on the knob — every demo container is published on ALL interfaces on EVERY demo-up, today, flag or no flag"
---

# M220 — Cue sheet

## Goal
`/demo-up` **means what the user thinks it means**: full data, the three story orgs, and remotely reachable — **by
default**.

## What the user asked for vs what is actually true

| Ask | Reality | Work |
|-----|---------|------|
| *"it should pull all the data (taxonomy, content, library, etc.)"* | **ALREADY DEFAULT-ON.** `for s in taxonomy directus sim-embeddings` **IS** the complete surface registry (library = directus; `/library/ai-simulations` = sim-embeddings). **No 4th surface exists.** | **Doc it.** The real failure is **not a default — it is a COLD CACHE**: replay is cache-first and **never captures**, so on a fresh box every surface exits rc=5 → a structural-only world. That is exactly what happened on `billion`. **The fix is M217's cache prime, not a flag.** |
| *"make sure it does the expected seeding with the 3 orgs of the stories"* | **ALREADY DEFAULT-ON** (`DEMO_STORIES` defaults to 1 — it is already opt-*out*), and the preset ships **3** orgs. | **A 4-LINE DOC FIX.** The docs say **"2 orgs"** — `.claude/skills/demo-up/SKILL.md:109,153`, `corpus/ops/demo/README.md:34`, the stale `seed_label` at `up-injected.sh:1081`, `stories.seed.yaml:1`. **This is almost certainly why the user believes the ask is unmet.** |
| *"access from remote should be opt-out"* | **THE ONE GENUINE FLIP.** `--public-host` is opt-in, default-empty, and there is **ZERO host auto-discovery anywhere in rext** (grep `tailscale status|BackendState|DNSName` → only comments). | Build it. See below. |

## Why section
Every deliverable is enumerable: a doc fix, a defaults table, a capability ladder with a specified rung list, a flag,
and a safety-doc authoring task.

## Scope

### In

**(a) The doc fix** — "2 orgs" → **3** at the 4 sites above.

**(b) The `/demo-up` DEFAULTS TABLE** — **BLIND AREA.** No enumerated defaults contract exists in the corpus; the
only complete knob list is a skill `argument-hint`. Document all ~25 `DEMO_*` knobs: knob | default | consumer |
file:line.

**(c) The remote flip (D-DESIGN-3): `--public-host auto`, DEFAULT-ON for demo** (opt-out via `--no-public-host`),
driven by a strict **capability ladder** — *capability-gated, never presence-probed*:

1. `command -v tailscale`
2. `BackendState == "Running"`
3. a **dotted** `.Self.DNSName` (strip the trailing dot). **A dotless host is hard-refused** at
   `up-injected.sh:81-86` because `@clerk/backend`'s `assertValidPublishableKey` rejects it.
4. `MagicDNSEnabled == true`
5. `tailscale serve status` shows no operator/sudo denial (reuse the F1 probe at `up-injected.sh:238-247`)
6. **`tailscale cert` actually mints**

> ### HARD INVARIANT — the fallback is not optional
> **Any failed rung ⇒ fall back to an EMPTY `STACK_PUBLIC_HOST` — byte-identical to today's localhost path — with
> ONE loud line naming the exact fix command.**
>
> **Why this is non-negotiable:** `SCHEME` (`up-injected.sh:74`) and `BIND_HOST` (`:66`) **both derive from the same
> `-n $STACK_PUBLIC_HOST` predicate**. So a **half-satisfied** public path is **strictly worse than localhost** —
> every baked URL becomes `https://` against **plain-HTTP listeners**, and the demo **does not load at all**. Today
> it always works. A laptop with no tailscale must stay **byte-identical to today**.

**(d) The dev-side opt-in `--public-host`** — per D-DESIGN-3, dev stays **opt-in**, which means building the flag
`/dev-up` **does not have today**. This **folds the reserved M216** (dev-path Tailscale parity), consuming the
reservation.
> **DECLARED SCOPE-FLEX LEVER.** If (d) bloats, it **drops back to M216** and the release still fully meets the
> user's demo-side spec. Do not let it eat the milestone.

**(e) Front the cockpit on `tailscale serve`** — add `('cockpit', 7700)` to `gen_tailscale_serve.py:42-46`. Today
the presenter's **entry point** is the **one plain-HTTP, unauthenticated surface** on the tailnet
(`up-injected.sh:1289-1294` deliberately excludes it).

### In — inherited from M218 (Fate-3, added at the M218 close, 2026-07-14)

Four **egress / injected-env / safety-contract** items. All four were declared *"also in scope"* in M218's own
overview and **did not land there**: each mutates the demo runtime on or beside the login path, and M218's exit
gate is *a p95 over 5 consecutive cold reset-to-seed cycles graded on a specific binary* — **iter-05 D13
established that a runtime change restarts the count** (the same bind that routed **F-11** to M219, **D16**).
They belong here regardless: M220 **owns** `/demo-up`'s defaults, the injected-env contract, and `safety.md`
**Part 3 — the exposure side**, and all four are exposure/egress items.

**(f) Clerk telemetry OFF** — `CLERK_TELEMETRY_DISABLED` + `NEXT_PUBLIC_CLERK_TELEMETRY_DISABLED`. Real egress
from **both** frontends today (grep across rext: **zero** hits for `TELEMETRY_DISABLED` — it was never wired).
It is also what makes Playwright's `networkidle` hang. **Pure env, no repo edit.**

**(g) Ad-tech egress on authenticated loads (F-5)** — the demo attempts **Google Analytics + DoubleClick +
Google Ads + LinkedIn Ads** on **every authenticated page load**. A demo that claims to be self-contained
should not phone four ad networks. Measure, then kill at the injected-env / CSP layer.

**(h) Vendor clerk-js + bound the unbounded timeout (C-5)** — the fake FAPI proxies `clerk.browser.js` **live
from `cdn.jsdelivr.net`** on every full page load, via `http.Get` = `http.DefaultClient` = **`Timeout: 0`
(unbounded)**, with **no server-side cache** (`clerkenstein/clerk-frontend/server.go:187`). next-web's entire
authenticated tree is **client-gated on clerk-js**, so a CDN stall is an **unbounded hang on the login path**.
Costs 0.2 s healthy / **~127 s if egress blackholes**. **Alignment-INVISIBLE** — no DNA gene covers `GET /npm/`
⇒ a **gate-free** win. Serve from disk; keep the CDN proxy only as a bounded fallback.
> This is the single item that most directly contradicts `safety.md`: **an unbounded internet dependency in the
> login path of a demo the corpus describes as self-contained.** It is Part-3 material, not a nice-to-have.

**(i) Clerkenstein-wire ant-academy** — `demo-stack/ant-academy.sh:146` copies `CLERK_SECRET_KEY` **straight
from `platform/.env`**, i.e. the **REAL Clerk app's secret**, into a demo process; `PK_DEMO` is never used, so
`@clerk/nextjs` runs **keyless** and phones Clerk to provision a throwaway app. Off the login path, but it is
**real-Clerk egress + a real production secret inside a demo** — the **same class as the `DIRECTUS_TOKEN`
fix16/17 strip**, and it contradicts `safety.md`. **Pure env.**

> **(i) IS ALSO A RENDER BUG, not only an egress bug — escalated from M219 (Fate 3, 2026-07-14, D4).**
> The keyless mode this item describes is **why the academy paints nothing**. Measured from a tailnet peer:
> `:13077` returns **200 to curl** (F-13's launcher fix is real — the 502 is gone) but **in a browser it
> redirects to `/clerk-sync-keyless?__clerk_handshake=…` and renders 0 meaningful text.** A presenter clicking
> **AI Academy** gets a blank page.
>
> **New evidence for the fix:** the `.env.local` overlay (`ant-academy.sh:181`) greps
> `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY|CLERK_SECRET_KEY` out of `platform/.env` — which carries **11 matching
> lines**. All 11 are written to `.env.local`; the **last one wins**, and it is not the demo's minted key. That
> is the concrete mechanism by which Clerk ends up with no usable publishable key and falls into keyless mode.
> The fix is the one this item already names — **use `PK_DEMO`** (the minted Clerkenstein pk that
> `up-injected.sh:391` already writes into next-web's `.env.local`) — but it must now also be **verified to
> RENDER in a browser**, not merely to stop leaking. It needs live browser iteration against a running demo,
> which is why it is here and not in M219.
>
> **M219 already landed the honest fences** (so this cannot go green while blank): the launcher now reads the
> response BODY and fails loud with `SERVES BUT DOES NOT RENDER` on a keyless bounce
> (`ant-academy.sh`), and `ANT_ACADEMY_HOME_SECTION` now requires an `AI Academy` marker + a 400-char floor
> instead of the meaningless 40 (`stack-verify/e2e/lib/coverage-manifest.ts`). **Both will report RED until this
> item lands — that is intended.** An accurate red beats a comfortable green.
>
> ### 🔴 SEVERITY RAISED — THE ACADEMY DOES NOT MERELY RENDER BLANK. IT POISONS THE DEMO SESSION.
>
> **This is demo-BREAKING, not cosmetic, and it is the reason (i) must land.** The academy's own keyless Clerk
> on `:13077` responds with:
>
> ```
> Set-Cookie: __session=; Expires=Thu, 01 Jan 1970 00:00:00 GMT     ← DELETES the demo's session
> Set-Cookie: __client_uat=0; Domain=taildc510.ts.net               ← DOMAIN-wide, not port-scoped
> ```
>
> **Cookies are scoped by HOST, not by PORT.** So the academy on `:13077` clobbers the session next-web holds on
> `:13000`. Two consequences, both measured, and proven by a controlled A/B from a peer (same seat, same page,
> one variable changed):
>
> 1. **A presenter who clicks the AI Academy link is LOGGED OUT of the demo** — and lands in
>    `ERR_TOO_MANY_REDIRECTS`. The blank academy page is the *lesser* half of this bug; the presenter's live
>    demo session is gone, mid-demo.
> 2. **Every EMPLOYEE coverage sweep aborts** once the crawler reaches the academy link. So while this is live,
>    **the employee vantage has no runnable sweep at all** — which is itself an *absence-read-as-success* risk
>    (D17): a vantage that cannot be measured must never be recorded as measured.
>
> **The `PK_DEMO` fix this item already names is the fix for this too** — a Clerkenstein-wired academy shares the
> demo's session instead of destroying it. But it must be verified by **logging in, visiting the academy, and
> then confirming the demo session SURVIVES** — not merely that the academy paints.

**(j) studio-desk bounces the presenter OUT of the demo — escalated from M219 (Fate 3, 2026-07-14, D4).**
Reproducible with plain `curl` from a tailnet peer, **not** a harness artifact: `studio-desk` on **`:19000`
returns a 302 to `https://billion.taildc510.ts.net:13000/login`**, and in a browser it lands on **`:13000/home`**.
So a presenter who clicks **"Anthropos Studio"** is silently thrown out of Studio and back into next-web — the
Studio surface is **unreachable in the demo**, and the demo's own nav is the thing that ejects them.

`:13000` is the demo's own next-web (3000 + offset 10000), so this is **not a prod-eject** — which is exactly why
the coverage sweep's prod-eject detector never flagged it. It is an **in-demo dead end**: studio-desk's auth check
does not recognise the Clerkenstein session minted for the demo origin and bounces to a login that next-web
immediately satisfies. Same family as (i): a UI-tier surface whose **Clerk wiring was never verified end-to-end
from a browser**, only that its port answers.

Diagnosis needs a running demo + browser (which session/cookie/origin studio-desk actually rejects), so it is
**demo-up-defaults work, not seed work**. Fix it alongside (i) and (e) — and add a fence that asserts a presenter
clicking **Anthropos Studio** stays *in Studio*, since "the port answers" demonstrably does not.

### Out
- Anything about *speed* (M218) or the AI-readiness render path (M219).
- A "hiring" story org — **D-DESIGN-4: it does not exist and will not be built.**

## Delivers → knowledge/corpus

**1. `corpus/ops/safety.md` Part 3 — the exposure side.** **BLIND AREA, and BLOCKING for the flip.** safety.md's two
promises are **read-side** (never read customer data) and **write-side** (never write prod). Grep it for
`tailscale|remote|expose|network|localhost` → **zero hits**. **Remote reach is a THIRD AXIS the safety contract does
not cover at all.** A default-on flip cannot ship without a real safety argument, not a doc edit.

The argument must state plainly, in writing, what is being made ambient:

> **A demo is an unauthenticated, authz-weakened build.** Clerkenstein disarms Clerk token verification in
> app/cms/jobsimulation/skillpath; an **authz-skip demo-patch is applied by default** (`DEMO_NO_AUTHZ_SKIP=0`); and
> **the presenter cockpit is a one-click, password-free "become any seeded hero" launcher** — a bare GET to
> `/v1/client/handshake?__clerk_identity=<key>`. Default-on remote reach makes that panel **ambient on every box
> that satisfies the capability ladder.**

The steelman **for** the flip (also record it): the demo's data is synthetic + public-snapshot-only (Parts 1 and 2
hold unchanged — no customer data can be in a demo, prod writes are structurally blocked); a tailnet is a real
authenticated device mesh (WireGuard, per-device keys, ACL-gated, no public listener); and the **exposure delta is
smaller than the docs claim** (see below).

**2. An explicit written SUPERSESSION of v2.2's D-DESIGN-1** (*"public reach is never default-on"*,
`demo-up/SKILL.md:78`) — **for the demo path only.** Never a silent contradiction.

**3. The CORRECTION of a FALSE safety claim** at `tailscale-serve.md:405-407`. It says: *"no open
0.0.0.0-on-the-LAN surprise… Binding 0.0.0.0 is gated on the knob precisely so it is never ambient."*
**This is false.** `gen_injected_override.py:259-260,292-294` emits `ports: !override` with **bare
`"<hostport>:<target>"` pairs — no `127.0.0.1` prefix** — so Docker publishes **every demo container on ALL
interfaces on EVERY demo-up, today, with or without `--public-host`** (and on Linux, Docker's iptables **bypass the
host firewall**). `BIND_HOST` gates only the two **host-native** servers (cockpit, ant-academy).
> **This correction ships REGARDLESS of the flip decision.** A shipped safety doc that understates actual exposure
> is the worst failure mode in the project.

## Open question — `tailscale cert` and Let's Encrypt rate limits

rext's own docs assert the cert **RE-ISSUES on re-run** (`up-injected.sh:885-886`; `tailscale-serve.md:335-336`). If
that is true, **default-on is a live LE hazard**: `ts.net` is a **PSL entry**, so the duplicate-certificate bucket is
**per-tailnet**, and a mint failure **silently falls back to a local-trust cert that a remote browser rejects**
(`:902-904`). The `$CERTS/fapi.crt` skip (`:895`) only holds **per demo-N** — every fresh `demo-N` always calls the
mint.

**Settle it EMPIRICALLY before flipping:** run `tailscale cert` twice on `billion`, diff the wall-clock, and
`journalctl -u tailscaled | grep -i acme` for a second ACME order. If tailscaled caches (as it should), **the repo's
claim is a doc bug**. If it does not, the mint must be cached at the **box** level, not the per-demo-N `$CERTS` dir.

## KB dependencies
- `corpus/ops/safety.md` (the contract being extended — **read it before writing Part 3**)
- `corpus/ops/demo/tailscale-serve.md` (the v2.2 remote-access runbook + the false claim)
- `corpus/ops/demo/README.md` · `.claude/skills/demo-up/SKILL.md` · `.claude/skills/dev-up/SKILL.md`
- `corpus/ops/snapshot-spec.md` + `corpus/ops/snapshot-cold-start.md` (why "pull all data" is a cache problem, not a
  default problem)
- rext: `demo-stack/up-injected.sh` (the only arg parser — it **hard-errors on unknown args**) ·
  `stack-injection/gen_tailscale_serve.py`
