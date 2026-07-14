# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills). This file holds the **active major** only; the retired **v1.x** history (M0 … M46, all
SHIPPED) lives in [`roadmap-legacy.md`](roadmap-legacy.md). Future versions + the unscheduled backlog live in
[`roadmap-vision.md`](roadmap-vision.md). The live source of truth for *current/next* is [`state.md`](state.md).

> **Designed 2026-07-13** via `/developer-kit:design-roadmap`. **v2.3 "cue to cue"** is the **presenter-speed
> release** — a **field-hardening release** (the v1.3b / v1.10b / v2.1 lineage) triggered by a **live presenter
> defect**: clicking a hero in the cockpit takes **1–2 MINUTES** to land in the platform, making a hero swap
> unusable in a live demo. The investigation found the cockpit + Clerkenstein handshake leg **provably free** (no
> sleeps, no I/O, a 303) — **the wall is entirely downstream**, and it was **already measured in-repo** (76 s members
> grid, 84 s router max-latency, a 180 s AI-readiness read) while **the corpus asserted in 4 places that login is
> "~2–5 s we can't shorten"** (a 20–40× false claim, booked as M43-D5 and never revisited). Two `app` perf
> demo-patches exist to kill those walls and **both silently REFUSE on sha-drift on every run** (pinned @ app
> v1.295/v1.315; the box runs v1.337) with **the refusal reason piped to `/dev/null`**. **5 milestones M217 →
> { M218 ∥ M219 ∥ M220 } → M221**; tag **`v2.3`**; branch `release/02.30-cue-to-cue`. **Tooling + docs only — zero
> platform-repo edits** (platform-side walls route to the sanctioned sha-pinned demo-patch hatch, never to a repo
> edit). Decisions: the <5 s gate is on **ACCESS (authenticated, interactive shell), not full first-page data
> render**; demo remote-access flips to **opt-out** (dev stays opt-in); the three story orgs are **the three that
> already exist** (no hiring org).
>
> **Designed 2026-07-11** via `/developer-kit:design-roadmap`. **v2.2 "panorama"** is the **external-shareability
> release** — make dev/demo stacks reachable from other machines on a **Tailscale** tailnet (run a stack on a
> Tailscale VM, e.g. `billion.taildc510.ts.net` on the odyssey Proxmox host; a teammate with Tailscale up browses
> the demo end-to-end). The **sanctioned re-proposal** of the v1.4 seed "external stack shareability
> (Tailscale/ingress)" dropped 2026-06-11 pending a fresh design-roadmap run. **4 milestones M212 → { M213 ∥ M214 }
> → M215** (+ optional M216); tag **`v2.2`**; branch `release/02.20-panorama`. **Tooling + docs + an opt-in flag
> only — zero platform-repo edits** (two platform-family files ride the EXISTING rext sha-pinned patch mechanism).
> Decisions: HTTPS-everywhere under one MagicDNS origin; external access **opt-in, default off**; demo-first.
>
> **Designed 2026-07-08** via `/developer-kit:design-roadmap`. **v2.1 "quick change"** is the **skiller-in-app
> re-ground** — a field-hardening release (v1.3b "dress rehearsal" / v1.10b "fit-up" lineage) triggered by a
> **landed platform structural change**: the `skiller` service + its DB schema **merged into `app`** (domain → the
> **`public`** schema, table names unchanged `skiller.X → public.X`; RPC → `backend`; the skiller GraphQL subgraph
> gone → **4 subgraphs**; skiller repo/container removed). A colleague's `docs/skiller-in-app-merge` corpus sweep is
> **correct-but-incomplete** and touches no tooling; the **rext** tooling still queries `skiller.<table>` and the
> stacks are half-synced. v2.1 re-fits the tooling, the corpus, and the stacks to the merged platform and **proves
> `dev-up` + `demo-up` still work**. **4 milestones M208 → M209 → M210 → M211** (strictly sequential); tag
> **`v2.1`**; branch `release/02.10-quick-change`. Tooling + docs only — zero platform-repo edits.
>
> **Designed 2026-06-29** via `/developer-kit:design-roadmap`. **v1.10b "fit-up"** is an **interposed
> field-hardening backfill** (the v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up
> issues + a tail of v1.10 content gaps. The M201 close *reported* the `stack-demo` clones ~5 weeks / 115+ commits
> behind prod — but **M47 found the clones actually CURRENT** (next-web @ v2.89.0); the genuinely-stale surface is
> the **corpus** (the shipped AI-readiness feature is undocumented), which **M48** re-grounds. v1.10b recaptures the
> snapshot, re-grounds the corpus, fixes the bring-up + content issues, adds an AI-readiness showcase org, and
> consolidates one auditable manifest — so **v2.0 "opening night" is PAUSED** until it ships. The v1.x flat
> milestone counter **re-opens** at **M47**
> (M47→M53); tag **`v1.10.1`**; branch `release/01.10b-fit-up`.
>
> **Designed 2026-06-28** (prior): **v2.0 "opening night"** opened a **NEW MAJOR** — **Playthroughs** is a new pillar
> (functional-flow *testing*: proving the platform's core user journeys actually work end-to-end), distinct from the
> v1.x demo/seeding lineage. v2+ adopts the **`Mxyy`** scheme (M201, M202, M203, M204). **v2.0 SHIPPED 2026-07-02
> (tag `v2.0`)** — all four milestones closed, 10 live Playthroughs GREEN on cold reset-to-seed; the first v2.x
> release. **v2.1 "quick change"** (the skiller-in-app re-ground) is now **IN DEVELOPMENT** (designed 2026-07-08;
> see below). The pre-v2 v1.x history (M0 … M46) lives in `roadmap-legacy.md`.

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.10b** | **fit-up** | Field-hardening backfill — re-ground demo + corpus to current prod, fix the from-scratch `/demo-up` issues + the v1.10 content gaps, add the **AI-readiness showcase org**, and consolidate **one auditable seed+gen manifest** | M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53 | ✅ **SHIPPED 2026-07-01 (tag `v1.10.1`)** (branch `release/01.10b-fit-up`, designed 2026-06-29; all 7 milestones done) |
| **v2.0** | **opening night** | The platform's core user journeys, **proven to actually work** — a new **Playthroughs** pillar: a manifest-driven, deterministic e2e suite that *pretends to be the human* and proves the platform does its job | M201 ✅ ∥ M202 ✅ → { M203 ✅ ∥ M204 ✅ } → ✅ ship | ✅ **SHIPPED 2026-07-02 (tag `v2.0`)** (branch `release/02.00-opening-night`, designed 2026-06-28; all 4 milestones closed-on-gate/complete). **10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed, 1 in-manifest TODO.** The first v2.x release. Tooling + docs only, zero platform edits, zero new deps |
| **v2.1** | **quick change** | The **skiller-in-app re-ground** — re-fit the tooling, corpus, and stacks to the merged platform (skiller service + schema folded into `app`/`public`, RPC → `backend`, **4 subgraphs**) and **prove** `dev-up` + `demo-up` still work. Field-hardening lineage (v1.3b/v1.10b), triggered by a landed platform change | M208 → M209 → M210 → M211 (strictly sequential) | ✅ **SHIPPED 2026-07-09 (tag `v2.1`)** (branch `release/02.10-quick-change`, designed 2026-07-08; all 4 milestones done — the merged platform stands up **cold on both stacks**; M42 coverage both vantages + v2.0 Playthroughs 10/11 GREEN; tooling + docs only, zero platform edits, 0 net-new deps) |
| **v2.2** | **panorama** | The **external-shareability release** — make dev/demo stacks reachable over a **Tailscale** tailnet (run on a Tailscale VM; a teammate browses the demo end-to-end over its MagicDNS name), via a single opt-in host knob + the tailscale-cert HTTPS surface. The re-proposal of the dropped v1.4 Tailscale/ingress seed | M212 ✅ → { M213 ✅ ∥ M214 ✅ } → M215 ✅ (+ opt M216) | ✅ **SHIPPED 2026-07-12 (tag `v2.2`)** (branch `release/02.20-panorama`, designed 2026-07-11; all 4 core milestones done — opt-in default-off, HTTPS-everywhere, demo-first; tooling + docs only, zero platform edits, 0 net-new deps; rext code-of-record `v2.2` = `39e8013`). **M215 proved it live:** the first remote Linux-VM demo over Tailscale, both vantages green from a 2nd machine on a trusted cert, reproducibly on a cold reset-to-seed |

| **v2.3** | **cue to cue** | The **presenter-speed release** — a presenter swaps heroes in **under 5 seconds** on a demo that comes up **green, fully-loaded, and remotely reachable by default**. Field-hardening lineage, triggered by a live 1–2-minute cockpit-login defect whose root causes were **already measured in-repo** and **silently rotting** (two dead perf demo-patches, a refusal piped to `/dev/null`, a 4-place false latency claim in the corpus) | M217 → { M218 ∥ M219 ∥ M220 } → M221 | 🔵 **IN DEVELOPMENT** (branch `release/02.30-cue-to-cue`, designed 2026-07-13). Tooling + docs only — zero platform-repo edits |

> The complete v1.x version-plan table (v1.0 "body double" … v1.10 "method acting", all ✅ SHIPPED) is preserved
> in [`roadmap-legacy.md`](roadmap-legacy.md) § Version plan.

The Playthroughs capability is governed by the consolidated **capability spec**
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3 — all in-scope decisions made +
review-hardened). v2.0's milestones build the contract that spec defines. Hard constraints carry over from the
v1.x lineage: **no modification to any platform repo** (the platform stays read-only — a surface that can't be
driven without a platform edit *escalates*, it does not edit), and all stack-operating tooling lives in
**`rosetta-extensions`** (built + tested in the `.agentspace/rosetta-extensions/` authoring copy, tagged, then
consumed per-stack at a pinned tag). Playthroughs reuse the M42 e2e foundation + the seeding machinery — they are
the **functional** sibling of M42's **presence**-only coverage sweep.

---

## In Development — v2.3 "cue to cue" (designed 2026-07-13, branch `release/02.30-cue-to-cue`)

**Theme.** *A presenter swaps heroes in under 5 seconds, on a demo that comes up green, fully-loaded, and remotely
reachable by default.*

**Trigger.** A live presenter defect, reported by the user: *"I click a user, then it takes 1 or 2 minutes to access
the actual platform. Once logged in it works normally. For a demo stack it is key that a presenter can swap from one
hero to another with little time."*

### What the investigation actually found (2026-07-13, 10-agent workflow + a dedicated residual-Clerk audit)

The full gap analysis — every claim file:line-cited — is preserved at
[`.agentspace/scratch/roadmap-research-2026-07-13.md`](../../.agentspace/scratch/roadmap-research-2026-07-13.md).
The five findings that shape this release:

1. **Clerkenstein is INNOCENT.** The cockpit CTA is a plain `<a href>` (`demo-stack/cockpit.py:397`); the FAPI
   handshake is an in-memory mint + a 303 with **zero I/O** (`clerkenstein/clerk-frontend/server.go:204-255`);
   `grep time.Sleep` across all of `clerkenstein/` returns **zero hits**. The cockpit's overlay timers
   (0/1200/3000 ms) are cosmetic text swaps that explicitly never `preventDefault`. **Every surviving suspect lives
   AFTER the 303.** Also REFUTED: a cold Next.js JIT compile (the demo runs a **production** build —
   `Dockerfile.dev` runs `pnpm turbo build`, compose pins `NODE_ENV=production`).
2. **The walls were already measured, in this repo, and nobody looked** — 76.7 s members grid, 84 s router
   max-latency, a 180 s AI-readiness read that never completes (`stack-verify/e2e/lib/section-assert.ts:63-73`;
   `corpus/ops/demo/coverage-protocol.md:203,411,471`). The tooling even encodes the truth
   (`playthroughs/e2e/playwright.config.ts:22` `timeout: 120_000`, comment: *"minute-scale"*).
3. **The corpus asserts the opposite, in 4 places** — *"~2–5 s, which we can't shorten"*
   (`corpus/ops/demo/cockpit-spec.md:58,155`; `cockpit.py:12,204-208`). Booked as an M43 scope-`Out:` + decision
   **D5** with **zero deferrals recorded**, so it never entered a ledger and was never revisited across v1.10b /
   v2.0 / v2.1 / v2.2. **This release formally RE-OPENS M43-D5 as a design correction.**
4. **Two `app` perf demo-patches silently REFUSE on sha-drift — on every run.** Pinned @ app v1.295.0/v1.315.0;
   `billion` runs **v1.337.0**. The applier prints the exact mismatch and **`up-injected.sh:701,717` pipes its
   stderr to `/dev/null`**. The warnings are present in **all four** bring-up logs on the box. The un-patched code
   is confirmed live: `roles.go:96` per-row Sentinel RPC; `ai_readiness.go:538` unbounded `loadMembers`.
5. **The last real run on `billion` was not a valid measurement.** The cockpit **crashed** (`OSError: Address already
   in use`, `cockpit.py:567`) on a port leaked by its predecessor — so the user was very likely driving a **stale
   cockpit** serving a **stale manifest** (dead clerk-ids) against a freshly-reseeded DB. All three snapshot replays
   **SKIPPED** (cold cache → structural-only catalog); autoverify ended **FAILING**; `jobsimulation` **exits(1)** in
   every demo. **No number taken before M217 lands is trustworthy.**

### User decisions (2026-07-13, at design time — binding)

| # | Decision | Consequence |
|---|----------|-------------|
| **D-DESIGN-1** | **The <5 s gate is on ACCESS, not on full first-page render.** *"The gate is on the login/access to the platform, not on the load of the complete render of the first page."* | **ACCESS** := the authenticated shell is rendered and interactive with the hero's identity present (full-screen loading gone; user menu shows the hero). Heavy in-page data (the 200-member grid) may stream in after and is **REPORTED as a secondary metric, never gated**. This sidesteps the platform-side DataLoader defect instead of fighting it. |
| **D-DESIGN-2** | **Fix it properly via Clerk/rext first; demo-patches are allowed if genuinely required.** *"if changes are required (after trying to address them properly via clerk) you can go for patches of course."* | Order of preference: rext env/compose/injection fix → Clerkenstein fix → **new sha-pinned demo-patch** (the sanctioned hatch). **A platform-repo edit is NEVER in bounds.** |
| **D-DESIGN-3** | **Remote access: OPT-OUT for `/demo-up`, OPT-IN for `/dev-up`.** *"if I write `/demo-up` the demo is built with remote access by default."* | **SUPERSEDES v2.2's D-DESIGN-1** ("public reach is never default-on") **for the demo path only**. Dev stays opt-in — which means **building the `/dev-up --public-host` flag that does not exist today** (the reserved **M216** scope, hereby **folded into M220**). |
| **D-DESIGN-4** | **The three story orgs are the three that already exist.** There is **no hiring org** and none will be built. | The shipped three are **ai-transformation (Cervato) / sales-ramp (Solvantis) / ai-readiness (Northwind)**. A real hiring story would need unmapped domain tables **+ the `hiring-app` frontend, which is not in the demo UI tier** — a separate release, mapping onto the reserved **M205**. **The "3 orgs" ask is therefore a DOC FIX** (the docs say "2 orgs"; the code ships 3). |

### Hard constraints (carried from the v1.x/v2.x lineage — unchanged)

**Zero platform-repo edits.** The platform stays read-only. A wall whose root cause is platform source
(`GetOrganizationTargetRole`'s 3-RPC-per-member fan-out; the AI-readiness `CycleID == nil → buildLiveResponse`
default; `jobsimulation`'s missing subcommand) routes to a **new sha-pinned demo-patch** or **escalates** — it does
not get edited. All stack-operating tooling lives in **`rosetta-extensions`** (authored + tested in
`.agentspace/rosetta-extensions/`, tagged, consumed per-stack at a pinned tag).

### Execution graph

```
v2.3 "cue to cue"

  M217 ──┬──→ M218 (iterative: <5 s access, both vantages) ──┬──→ M221 (iterative: prove it on billion)
  clean  │                                                    │
  stage  ├──→ M219 (AI-readiness RENDERS) ───────────────────┤
         │                                                    │
         └──→ M220 (demo-up defaults + remote opt-out) ──────┘

  M217 is a HARD barrier: the stale-cockpit + dead-patch + cold-cache confounds mean
  NO measurement before it lands is trustworthy. M218/M219/M220 then run in parallel.
  M218 merges FIRST of the three (it and M220 both touch up-injected.sh).
```

### Milestones

#### M217 — Clean stage  (`section`, medium) — ✅ **DONE 2026-07-13**
**Goal:** a `/demo-up` that comes up **green** — so that everything measured afterwards is real.

> **Closed 2026-07-13.** Exit gate MET on `billion` (cold reset-to-seed: `autoverify: OK — verified-working`,
> 3/3 replays exit 0, 2/2 app patches applied [one **self-healed**], jobsimulation serving, content plane local).
> **32 bugs fixed** across build (0) → harden ×3 (24) → close (8). rext code-of-record: `cue-to-cue-m217`.
>
> **What it actually taught us**, beyond the fixes:
> - **The KB-fidelity gate caught three false claims in the milestone's own plan** — including a `jobsimulation`
>   diagnosis whose prescribed fix would have **broken the service**. *The errors were in the plan, not the code.*
> - **The tooling was implicitly Docker-Desktop-shaped.** Two independent bugs (`jobsimulation`'s AWS mount;
>   Directus's `host.docker.internal`) are the same failure: *fine on a Mac, dead on a fresh Linux VM.*
> - **Self-review finds the bugs you are looking for; adversarial execution finds the ones you are not.** Pass 1
>   of hardening found 4 bugs and declared victory; an adversary that actually *ran* the code found 20 — three of
>   them introduced by pass 1 itself.
> - **A test that greps for a call proves nothing about whether the call resolves** (D9). `up-injected.sh` called
>   `reap_port` — the headline deliverable — without sourcing `reap.sh`. **It never executed once**, including
>   during the green proof run. A string-fence passed over a function that did not exist.
**In:** reap the leaked cockpit port (`7700+off`) + make `demo-down` reap the whole offset range (2 of the last 3
runs on `billion` died on a leaked port); **un-swallow the demo-patch REFUSE reason** (`up-injected.sh:701,717`);
**re-pin the two `app` perf patches** + add a **patch-freshness preflight that fails LOUD** (a perf patch that
silently degrades a demo from 5 s to 120 s is worse than one that refuses); fix `jobsimulation` exits(1)
(**DEF-M215-04** — AI-Simulations is dead in *every* demo today); **prime the snapshot cache on `billion`**
(**DEF-M215-02**); re-pin the drifted rext consumption clones (local `quick-change-m211`, remote
`panorama-m214-3-g41a28aa` → `v2.x`).

> **⚠️ CORRECTED 2026-07-13 by the M217 KB-fidelity gate (which came back RED — and was right).** The design-time
> research got **three** things wrong, and the milestone doc has been fixed:
> - **`jobsimulation` is NOT a missing-subcommand problem, and the drafted "compose-command fix" would have BROKEN
>   the service** (`command: serve` → a real `unknown command "serve"` → exit 1). The cobra **root `RunE` IS the
>   server**. The actual cause: `docker-compose.yml` binds `$HOME/.aws/credentials`, Docker **auto-creates the
>   missing host path as an empty DIRECTORY**, the AWS SDK hard-errors EISDIR inside `ai.NewAIManager`, cobra prints
>   its **usage block** on the returned error — *that* is the "prints CLI help" everyone mis-read. Fix: drop the
>   bind in the **generated override** (`volumes: !reset null`) — rext-only, no demo-patch, no escalation. (A demo
>   carries **zero** AWS credentials, so the mount could only ever be the broken empty dir.)
> - **The stale cockpit does NOT serve "dead clerk-ids"** — `CockpitHero` carries no clerk id at all. M218 must
>   **re-measure on a green stack** rather than inherit a phantom hypothesis.
> - **TWO, not three, replays were cache misses.** Directus was **rc=4** (`--no-local-content` ⇒ no directus schema),
>   so **priming the cache alone will not fix it** — the billion run must also be **local-content ON, from a purged
>   stack**.
>
> **BD-3 is now decisive:** `ai_readiness.go` differs **between the two boxes** (v1.334.1 vs v1.337.0), so **no
> single static whole-file sha pin can be correct on both** — the manifest schema cannot express the truth, and a
> one-shot re-pin cannot close M217. The **anchor** survives every tag (exactly 1×). See M217's `overview.md`.
**Out:** any latency fix (that is M218 — do not scaffold a fix before the harness measures).
**Delivers →** `corpus/ops/demo/demopatch-spec.md` (**BLIND AREA** — the sanctioned escape hatch this whole release
depends on has **no corpus doc**; its 6-guard contract exists only in a Python module docstring).
**Open question (BD-3):** keep the file-sha gate (safe, rots every `app` bump) or move to anchor-only
single-occurrence matching (survives bumps, weaker drift safety)? **Recommendation: keep the sha gate, add an
auto-repin verb + the loud preflight.** Evidence it matters: two agents computed *different* `ai_readiness.go` shas
days apart — a one-shot re-pin is a band-aid.

#### M218 — Seat change  (`iterative`, large → very-large)
**Status:** ✅ **`done` — closed-on-gate 2026-07-14.** Merged `--no-ff` into `release/02.30-cue-to-cue`.
rext code-of-record: **`cue-to-cue-m218`**. **Dir:** [`m218-seat-change/`](releases/02.30-cue-to-cue/m218-seat-change/)

> ### THE GATE IS MET — worst-case p95 **2413 ms** / **1767 ms** vs **< 5000 ms**
>
> **5 GENUINE cold reset-to-seed cycles**, both vantages, **50/50** logins reaching ACCESS, on `billion` over
> the tailnet. **Baseline 39.45 s / 38.30 s → a ~16× improvement on the honest cold number.**
>
> **⚠ The gate number is 2413 ms, not the 1456 ms iter-04 reported.** That figure was taken on a **warm
> database** — **F-9**: `/demo-down --purge` had **never purged on any Linux host** (postgres's UID-1001/0700
> cluster dir defeats `rm -rf`; `set -euo pipefail` then aborted `cmd_down` *silently*, leaking the registry
> slot and the images with a bare `rc=1`). `billion`'s postgres still carried the `PG_VERSION` `initdb` wrote
> on **2026-07-11**: every "cold" bring-up for days, iter-04's included, reused the same DB and images. **The
> gate was not merely ungraded — it was ungradeable.** The count was restarted at 0 and 5 cycles were run,
> each **proven cold** (`initdb` re-ran) and **proven green** before being measured.
>
> **Two root causes, both in demo tooling. Zero platform-repo edits.**
> 1. **~37.5 s** — next-web's **SSR** GraphQL origin was a **build-inlined PUBLIC url** that **blackholes from
>    inside the container** (`3 × 10.5 s` undici connect-timeout + a 2 s/4 s retry ladder). The milestone's own
>    planned one-line fix (re-point the runtime env) was a **proven no-op** — `NEXT_PUBLIC_*` is build-inlined,
>    so SSR never reads `process.env`. Fixed by *fixing the address, not the variable*: a **server-only**
>    `WUNDERGRAPH_SSR_ENDPOINT` (deliberately **not** a `NEXT_PUBLIC_*` name, so it is a real runtime read) +
>    a sha-pinned demo-patch teaching `server.graphql.ts` to prefer it.
> 2. **~6.1 s** — Clerkenstein's fake **BAPI served a hardcoded STUB user to every hero** (`// Disarmed: any
>    id → the demo user` — true when a demo had one user, false since M35's Stories & Heroes model). The BAPI
>    identity disagreed with the JWT identity, `app` refused `userPreferences`, and next-web's `retry: 2`
>    ladder burned ~6 s on **every** authenticated render. Fixed by making the BAPI **roster-aware**.
>
> **A blackhole and a refusal are six seconds apart in signature** — `3 × 10.5 s + 6 s` vs `3 × 33 ms + 6 s`.
> The magnitude named the bug *class* before a line of code was read. Folded into `latency-budget.md`.
>
> **Alignment: 97.2% overall / 100% critical (26/27) — the gate (≥95/=100) is MET, and the 2.8% is
> DELIBERATE.** Landing the `GetUser` gene exposed **F-11**: the BAPI *also* fabricates the **org** eid. It
> ships as a **permanently-visible RED gene** rather than being hidden by omitting the field (**D16**) —
> because **a silently-omitted field is exactly how the user-level stub survived four releases**. →
> `FIX-M219-bapi-org-eid`.
>
> **21 bugs fixed** (build 6 · harden 4 · **close 11**). **Python 887/0** · **Go 0 failures / 6 modules /
> 1784 funcs** · **flake 5/5** · **platform edits 0**.
>
> **The close's defining finding:** the hardening ledger **misfiled an M218 regression as "pre-existing"**
> because it used `f296e5e` — *M218's own iter-05 commit* — as the "baseline". Three further **false-greens**
> were found, all the **D17 stale-verdict class**, and **one sat directly on the gate path**: a **SKIPPED**
> demo-patch wrote nothing to `demopatch.log`, so a stack running **without the headline SSR fix** would have
> printed *"✓ demo-patches: none refused"*, graded **green**, and been **measured**. → **D18**: *"reproduced at
> baseline ⇒ pre-existing" is only sound if the baseline predates the milestone.*
>
> **Carried forward (9, all Fate-3, every receiving `overview.md` EDITED; ZERO escape-hatch):**
> **M219** ← F-11 org-eid · `expressrun` unmeasurable · freshness-gate skips.
> **M220** ← Clerk telemetry off · F-5 ad-tech egress · **C-5 vendor clerk-js + bound the unbounded
> `Timeout: 0`** · ant-academy's real-Clerk-secret leak.
> **M221** ← **F-7** the `NEXT_PUBLIC_BACKEND_API_URL` blackhole twin (*a loaded gun* — a measured 10.5 s
> blackhole, dormant only because every reader is client-side) · **C-3** the cms/Directus 403s.
> *None could land here without invalidating the gate: a demo-runtime change **restarts the 5-cycle count**
> (iter-05 D13).*

**Goal:** click **[Log in as]** → the hero is **in the platform** in **under 5 seconds**, for both vantages.
**Exit gate:** **p95 click→ACCESS < 5 s** — where **ACCESS** = the authenticated shell is rendered and interactive
with the hero's identity present (full-screen loading gone; user menu shows the hero) — for **both**
`maya-thriving` (employee → `/profile`) **and** `dan-manager` (manager → `/enterprise/…`), measured over **5
consecutive cold reset-to-seed runs**. In-page data-completion time is **reported, not gated** (D-DESIGN-1).
**Why iterative:** the cost budget does **not yet sum to 60–120 s** — the confirmed suspects total ~18 s. Reaching
2 minutes requires one of the unconfirmed big-ticket items to be real. **Build the harness first, guess second.**
**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md` (the live-browser
measure → attribute → fix → re-measure loop). **The 4-leg experiment discriminates all suspects in ONE bring-up with
zero code written** — run it as iter-01.
**Ranked suspects (adversarially surviving):**

| | Suspect | Vantage | Est. | Fix surface |
|---|---------|---------|------|-------------|
| **C-1** | next-web's **server-side** GraphQL URL resolves to its own loopback (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` → `localhost:5050`, un-offset; the demo exports `STACK_PUBLIC_HOST`, never `PUBLIC_HOST`) → `prefetchUserStatus` `retry:2` + `retryDelay 2s/4s`, and all three fetchers rethrow | **BOTH** | **~6 s / render** | **one line** in `stack-injection/gen_injected_override.py:96-112` (runtime env only — `NEXT_PUBLIC_*` is build-inlined, so the browser keeps its correct offset URL) |
| **C-2** | the two dead `app` perf demo-patches | manager | 76 s → **11.6 s even patched** | M217 re-pins; the 11.6 s residual is **out of gate scope** per D-DESIGN-1 |
| **C-3** | **cold-federation Directus drift → the Cosmo router RETRYING** for 60–90 s (cache-masked in a warm sweep; surfaces only on a COLD federation tier — and `billion` read content **live from prod over the WAN** because directus replay skipped) | **BOTH** | **60–90 s** | a real snapshot replay (M217's cache prime) — **the closest single match to "1–2 minutes" on a path both heroes traverse** |
| **C-4** | stale cockpit / dead clerk-ids (the crash) | both | unknown | M217 (**must land first — it contaminates everything**) |
| **C-5** | the fake FAPI proxies `clerk.browser.js` **live from `cdn.jsdelivr.net`, uncached, `http.DefaultClient` (no timeout)**, on every full page load — and next-web's whole authed tree is client-gated on clerk-js | both | **0.2 s healthy / ~127 s if egress blackholes** | vendor the bundle into the fake-fapi image. **Alignment-INVISIBLE (zero DNA genes cover `GET /npm/`) → a gate-free win.** Take it regardless |
| **C-6** | `billion` has **7.3 GiB RAM** vs the documented **12 GiB** floor | both, remote | unknown | measure `docker stats` during a login before blaming code; may be a pure VM resize |

**Also in scope:** disable Clerk telemetry (real egress from both frontends; it is what makes Playwright's
`networkidle` hang); Clerkenstein-wire **ant-academy** (it is handed the **real** `CLERK_SECRET_KEY` today —
`ant-academy.sh:146` copies from `platform/.env`; off the login path, but real-Clerk egress + a real secret in a
demo process, contradicting `safety.md`); the `x/crypto@v0.52.0` bump (13 dependabot alerts, all
govulncheck-UNREACHABLE) since we are in `clerkenstein/` anyway.
**Alignment guard (H3):** **no DNA gene covers latency or the clerk-js proxy** → caching/vendoring is free. But the
genes that break on a **handshake-shape** change are all `critical`, and the critical gate is **100%, no partial
credit** — and **CI is INERT** (`clerkenstein/.github/workflows/alignment.yml:10-11` says so of itself, while
`corpus/architecture/alignment_testing.md:233` claims a weekly workflow). **Any change to
`clerkenstein/clerk-frontend/` MUST carry an explicit `/align-run` step. Do not rely on CI.**
**Delivers →** `corpus/ops/demo/latency-budget.md` (**BLIND AREA** — there is **no** perf/latency budget, baseline,
or even a definition of "access" anywhere in `corpus/**` or rext); a full **click→painted-page login-sequence**
section in `corpus/ops/demo/cockpit-spec.md` (**BLIND AREA** — the next-web half of the login path is documented
nowhere; you cannot code a fix against one line); the clerk-js proxy caching/timeout contract in
`corpus/services/clerkenstein.md`; the **CI-inert correction** in `alignment_testing.md`; and the **M43-D5
correction** (the "~2–5 s" claim, 4 sites).
**Re-scope trigger:** if the harness shows neither C-1 nor C-3 explains the **employee** vantage, **stop and
re-measure** — do not proceed on a manager-only fix set.

#### M219 — Readiness renders  (`section`, medium) — ✅ **DONE 2026-07-14**
**Status:** ✅ **`done` — closed-complete 2026-07-14.** Merged `--no-ff` into `release/02.30-cue-to-cue` (`e7a33c6`).
rext code-of-record: **`cue-to-cue-m219-final`**. **Dir:** [`m219-readiness-renders/`](releases/02.30-cue-to-cue/m219-readiness-renders/)

> ### ✅ THE USER'S KICKOFF BAR IS MET ON BOTH COUNTS
> *"make sure each element and sub section of readiness is filled spot data + make sure to use the **right**
> dashboards/pages for readiness (both for managers and employees).. **not the old legacy ones**"* — proven on
> **5 cold reset-to-seed rebuilds** at `cue-to-cue-m219-r8`, each independently evidenced. **Zero platform-repo
> edits.** Manager: **all 8 readiness sections PASS**, `failingSections=0`. Employee: **both heroes**
> `failingSections=0`. **Zero** demo pointers resolve to a legacy surface.
>
> **⚠ TWO of this milestone's OWN PREMISES were REFUTED by measurement — and the planned demo-patch was
> WITHDRAWN, not built.** (**F-2**) The `CycleID == nil` blocker **does not exist**: the CURRENT frontend passes
> `?cycle=`; the original note was made against the **LEGACY** page. (**F-7**) M217's `loadmembers` patch is **not
> dead** — it *self-heals*; and *"the live recompute never completes"* is **false** — **2.09 s**, measured.
>
> **What was ACTUALLY broken.** All 3 demo pointers targeted `/enterprise/workforce/ai-readiness` — an **UNLINKED
> ORPHAN** (no nav, no tab, no redirect). The **employee surface has no route of its own** (it is embedded in
> `/home`) — *which is why route-crawling never found it*. **Junk skills ORG-WIDE**: the claimed-tail pool ran
> **dry** and topped up from the flat pool's **alphabetical head** — Aria + 8 named members claimed **"24-hour
> dietary recall"** / `15Five` / `17Track`. The arithmetic closed exactly (`want`=28, role pool=10, curated `data`
> supplied only **25 usable** ⇒ **exactly 3** junk tokens) and **Ben was clean only because his `want` (16) fit his
> pool — that asymmetry was the proof** it was pool **SIZE**, not resolution. **The flat tier is DELETED** (ladder =
> role → curated → **general** → **STOP**; exhausted ⇒ *fewer* skills, never padded ones). A hero rendered
> **ROLE-LESS** (`Operations Analyst` **resolves** but has **0 `job_role_skills`** ⇒ the resolver rejected it —
> ***"it resolves" ≠ "it has skills"***). And the manager's **4 interview-findings blocks rendered HEADINGS OVER
> NOTHING** — they read `jobsimulation.interview_aggregated_reports`, which **NO SEEDER EVER WROTE**. Now seeded;
> **the manifest's disclosed EXCEPTION was DELETED** and the floor raised **120 → 900** (the empty headings measured
> ~120–200 chars — *which is exactly why 120 passed over them*).
>
> **ONE new demo-patch, and it is not the withdrawn one:** `next-web-aireadiness-flag-gate` — no PostHog on a demo
> ⇒ the flag is `undefined` **forever** and the code demands `=== true`, so the member surface never mounts. Roster
> now carries **7**.
>
> **Inherited from M218, all landed:** **F-11** (the BAPI fabricated the org eid) ⇒ alignment **97.2% → 100% / 100%
> critical (27/27)**, the gene **retained as a permanent fence**; and **both absence-gates** — `expressrun` is no
> longer *unmeasurable-as-a-pass* (exit **3** ≠ exit **2**), and the freshness-gate's silent skip now **speaks**.
>
> **THE D17 THREAD — the spine of this milestone.** *A status artifact that outlives the thing it describes, and is
> then read as evidence* has now bitten **~14 times across M217→M219**. **Five new instances inside M219 — several
> found by the tooling in ITSELF:** `run-coverage.sh` printed *"coverage report written to …"* **unconditionally**
> and re-read whatever JSON was on disk, so a spec that threw first presented the **PREVIOUS run's numbers, "GATE:
> MET ✅" and all, exiting 0** — *it nearly graded an M219 rebuild on hours-old data from the **old, broken**
> stack*. **17 existing tests asserted the junk-fallback AS THE CONTRACT** (an expected value literally containing
> `K-JUNK-1`) — **they were not missing the bug, they were guarding it**; all inverted. The poisoned-pool fence's
> **first cut PASSED against the broken code** — theatre inside its own fix. The interview-findings fence asserted
> the **LEGACY** page's strings — and passed. And the academy launcher read *"a pid exists"* as *"the service is
> up"*, its **fixtures having encoded the broken behaviour** — so the suite was **GREEN for four releases while the
> academy 502'd**.
>
> **The generalized lessons:** ***"the role classifies" ≠ "the pool is big enough"*** · ***"it resolves" ≠ "it has
> skills"*** · ***"it serves" ≠ "it renders"*** · ***"the pool resolved" ≠ "the content is sane"*** · ***an errored
> SQL query is not "zero rows"*** · ***an unexecuted gate is a FINDING, not a pass.***
>
> **⚠ TWO DISCLOSED CAVEATS ON THE BATTERY — neither hidden, neither rounded away.**
> **(1) It is NOT one uncontested consecutive run.** Two further runs were graded RED for **battery-INTEGRITY**
> reasons caused by an **ORCHESTRATION error, not a demo defect**: two batteries were run **concurrently against the
> single demo host** and one **purged the stack mid-measurement**, so a gate went **UNEXECUTED**. *No demo defect was
> observed in anything that was measured*, and the 5 greens are each individually evidenced — but the audit trail
> carries a permanent asterisk. **(2) The code that GRADED is not the code that SHIPPED:** `c6648d1`
> (`aiReadinessStep1Score` **double-rounded** — it disagreed with the platform's `computeTier1` on **3 of 14**
> reachable weights) is a **SEED-PATH** change that landed **after** the graded tag, and **per D13 a seed-path change
> restarts the battery count**. Small, strictly corrective — **and still not allowed to be rounded away.**
>
> **13 bugs fixed** (build/R-8 9 · harden 3 · close 1). **Python 903/0** (+16) · **Go 1821 funcs / 0 failures / 6
> modules** (+37, same method) · **flake 0** · **platform edits 0**.
>
> **Carried forward (5, all Fate-3, BOTH receiving `overview.md` files EDITED; ZERO escape-hatch):**
> **M220** ← 🔴 **the academy POISONS the demo session** (cookies scope by **HOST, not PORT** — a presenter who
> clicks the academy link is **logged OUT of the demo** into `ERR_TOO_MANY_REDIRECTS`, and every employee coverage
> sweep **aborts**; severity **RAISED** on item **(i)**) · 🔴 **studio-desk `:19000` → 302 → `:13000/login`** (item
> **(j)**, added at close — clicking *"Anthropos Studio"* **ejects the presenter**). *Both ship with **honest fences
> that deliberately report RED** until M220 lands — a half-fix that reports green is worse than no fix.*
> **M221** ← **`GUARD-M221-host-isolation`** (a host lock or per-cycle stack `N`; **a prerequisite for M221's own
> gate**, itself a multi-cycle battery on that same singleton host) · **`FIX-M221-reap-native-academy`** (`down
> --purge` doesn't reap the host-native academy — the **old** process keeps answering `:13077` while the log says it
> *"DIED"*) · **`REPROVE-M221-battery-at-final-code`**.

**Goal:** the AI-readiness story is **visible**, not merely seeded.
**The seeding is a VERIFIED NO-OP** — Northwind Aviation (`narrative: ai-readiness`, 200 members, heroes **Aria
COMPLETED / Ben STARTED / Dana manager**) is in the **DEFAULT** preset (`stack-seeding/presets/stories.seed.yaml:118-153`);
all 3 seeders are registered **unconditionally** (`cmd/stackseed/main.go:410,411,431`); "enabled" = a
`public.organization_settings` row, written (`seeders/org_settings.go:72-73`); the last run wrote
`org rows=3, ai-readiness-config rows=6, org-settings rows=1`. **Do not budget seeder work.**
**In:** prove **Dana** sees a **FILLED** AI-readiness page — which needs M217's re-pinned patch **AND** resolving
that the **default GET (`CycleID == nil`) takes `buildLiveResponse`** (`app/internal/workforce/ai_readiness.go:285,301`),
bypassing the frozen-snapshot seed unless the FE passes `?cycle=` → **a NEW demo-patch** (platform-shaped, the hatch
is the answer); prove **Ben's** from-scratch STARTED workflow appears on his dashboard and **Aria's** COMPLETED
state renders; fix the stale ACTIVE-vs-CLOSED comment in `stories.seed.yaml:112-117` (the code writes
`status='closed'`).
**Delivers →** an **`ai-readiness` playthrough manifest** (**BLIND AREA** — the e2e suites are
profile/workforce/skill-paths/ai-simulations/assignment-monitoring only; **Aria's and Ben's journeys are not
e2e-proven**) + its section in `corpus/ops/demo/playthroughs.md`; updates to `corpus/services/ai-readiness.md`.

#### M220 — Cue sheet  (`section`, medium)
**Goal:** `/demo-up` **means** what the user thinks it means — full data, the three orgs, and remotely reachable,
by default.
**In:** **(a) the doc fix** — "2 orgs" → **3** (`.claude/skills/demo-up/SKILL.md:109,153`;
`corpus/ops/demo/README.md:34`; the stale `seed_label` at `up-injected.sh:1081`; `stories.seed.yaml:1`). *This is
why the user believed the seeding ask was unmet.* **(b)** author the **`/demo-up` DEFAULTS TABLE** (**BLIND AREA** —
no enumerated contract exists; the only complete knob list is a skill `argument-hint`). **(c) the remote flip
(D-DESIGN-3):** `--public-host auto` **default-ON for demo** (opt-out via `--no-public-host`), driven by a strict
**capability ladder** — `command -v tailscale` → `BackendState == "Running"` → a **dotted** `.Self.DNSName` (a
dotless host is hard-refused: `@clerk/backend`'s `assertValidPublishableKey`) → `MagicDNSEnabled` → no
operator/sudo denial → **`tailscale cert` actually mints**. **HARD INVARIANT: any failed rung ⇒ fall back to an
EMPTY `STACK_PUBLIC_HOST`, byte-identical to today's localhost path, with ONE loud line naming the fix.** A
*half-satisfied* public path is **strictly worse than localhost** — `SCHEME` and `BIND_HOST` both derive from the
same `-n $STACK_PUBLIC_HOST` predicate, so every baked URL becomes `https://` against plain-HTTP listeners and the
demo does not load **at all**. **(d)** the **dev-side opt-in `--public-host`** (folds the reserved **M216**;
**declared scope-flex lever** — if it bloats, it drops back to M216 and the release still meets the user's demo-side
spec). **(e)** front the **cockpit** on `tailscale serve` (`('cockpit', 7700)` → `gen_tailscale_serve.py:42-46`) —
today the presenter's entry point is the **one plain-HTTP, unauthenticated surface**.
**Delivers →** `corpus/ops/safety.md` **Part 3 — the exposure side** (**BLIND AREA, BLOCKING** — safety.md's two
promises are read-side and write-side only; grep for `tailscale|remote|expose|network` → **zero hits**. Remote reach
is a **third axis the safety contract does not cover**, and default-on cannot ship without it); an explicit written
**supersession of v2.2's D-DESIGN-1**; and the **correction of the FALSE claim** at `tailscale-serve.md:405-407`
(it says binding `0.0.0.0` is gated on the knob — **it is not**: `gen_injected_override.py:259-260,292-294` emits
bare `"<hostport>:<target>"` port pairs, so Docker publishes **every demo container on ALL interfaces on EVERY
demo-up, today, flag or no flag** — and on Linux Docker's iptables bypass the host firewall). **This correction
ships regardless of the flip.**
**Safety note to record in writing:** the cockpit is a **one-click, password-free "become any seeded hero"**
launcher, and a demo is an **authz-weakened build** (Clerkenstein disarms token verification; an authz-skip patch
is applied by default). Default-on remote reach makes that panel **ambient on every box that satisfies the ladder** —
which is why the ladder must be *capability*-gated, not *presence*-probed.
**Open question:** `tailscale cert` + **Let's Encrypt rate limits** — rext's own docs claim the cert **re-issues on
re-run** (`up-injected.sh:885-886`). If true, default-on is a live LE hazard (`ts.net` is a PSL entry ⇒ the
duplicate-cert bucket is **per-tailnet**), and a mint failure **silently falls back to a local-trust cert a remote
browser rejects**. **Settle empirically before flipping** (run `tailscale cert` twice on `billion`, diff wall-clock,
`journalctl -u tailscaled | grep -i acme`). If tailscaled caches, the repo's claim is a doc bug.

#### M221 — Prove it on billion  (`iterative`, large)
**Goal:** every requirement of this release, verified **on the remote VM, over the tailnet, with no flags passed**.
**Exit gate:** on `billion.taildc510.ts.net`, a **default** `/demo-up N` (no flags) yields, **reproducibly on a cold
reset-to-seed**: **(1)** p95 click→ACCESS **< 5 s** for both `maya-thriving` and `dan-manager`, measured **over the
tailnet origin** (the extra TLS/proxy hop is *in* the budget); **(2)** the **full replayed catalog** (taxonomy +
directus content + sim-embeddings — **no SKIPPED surface**); **(3)** all **3 story orgs** seeded, incl. AI-readiness;
**(4)** **Dana** sees a **FILLED** AI-readiness page; **(5)** **Ben's** from-scratch STARTED workflow is visible on
his dashboard; **(6)** **Aria's** COMPLETED state renders; **(7)** remote access came up **by default**; **(8)** **0
platform-repo edits**.
**Why iterative:** the direct analogue is **M215 "prove-on-odyssey" (7.1 h, direct-drive)** — the last breakages only
surface on a live cross-machine run.
**Also lands:** **DEF-M215-03(b)** — the committed, repeatable **remote-origin Playwright gate** that v2.2 owed.
(Note: Playthroughs declare perf a **NON-GOAL**, so the latency gate **cannot** be a Playthrough — it is a new
`stack-verify` surface.) And the **7.3 GiB RAM** question gets measured and decided.

### Deferrals folded in (LAND-NEXT)

| ID | Item | Fate |
|----|------|------|
| **DEF-M215-01 / F5** | the `app` perf demo-patches sha-drift REFUSE | **Fate-1 → M217** (it *is* part of the reported bug) |
| **DEF-M215-02 / F9** | fresh remote VM has no snapshot cache → sparse catalog | **Fate-1 → M217 + M221** (milestones 4 and 5 are unachievable without it; it also feeds **C-3**) |
| **DEF-M215-03(b)** | no committed remote-origin Playwright gate | **Fate-1 → M218 + M221** (it IS the measurement substrate) |
| **DEF-M215-04 / F13** | `jobsimulation` exits(1) → AI-Simulations dead in every demo | **Fate-1 → M217** ("the demo works properly" cannot be claimed with a container in a crash loop) |
| supply chain | `x/crypto@v0.52.0` (13 alerts, all UNREACHABLE) | **Fate-1 → M218's rext roll** |
| **M216** (reserved) | dev-path Tailscale parity | **CONSUMED → M220(d)** per D-DESIGN-3; declared scope-flex lever |
| plan hygiene | `metrics-history.md` has no v2.2 row; no release-scope deferral audit for v2.1/v2.2 | **Fate-3 → close-release** |

### Top risks

| Risk | Severity | Mitigation |
|------|----------|------------|
| **The measurement is contaminated before it starts** — the stale-cockpit confound means any number taken before the port reap is untrustworthy | **blocks-release** | **M217 is a hard barrier. Sequence is non-negotiable.** |
| **The confirmed cost budget does not sum to 60–120 s** (C-1 ~6 s + C-2 ~11.6 s ≈ 18 s). Reaching 2 min needs C-2-unpatched (76 s) and/or C-3 (60–90 s) and/or C-5-hanging | **blocks-release** | **Build the harness first, guess second.** The 4-leg experiment discriminates all suspects in one bring-up, zero code written. **Do not scaffold fixes before it runs.** |
| **The employee vantage is under-explained** — both dead patches are *manager* surfaces, yet the user reports **both** are slow ⇒ there is a **common** factor (only C-1 and C-3 qualify) | **blocks-release** | If the harness clears both for `/profile`, **stop and re-measure**. Re-scope trigger, declared in M218. |
| **Demo-patches rot silently BY DESIGN** — the demo builds from a scratch clone force-checked-out at the newest `v*` tag on **every** bring-up, so ref-pinning the source clones would **not** stop it | **degrades-quality (recurring)** | BD-3 → the loud freshness preflight + an auto-repin verb (M217). A one-shot re-pin is a band-aid. |
| **Default-on remote reach publishes an unauthenticated identity-vending cockpit** on every capable box | **blocks-release (safety)** | The capability ladder + the hard localhost fallback + **safety.md Part 3** + fronting the cockpit on TLS. All in M220, all blocking. |
| **`tailscale cert` / Let's Encrypt rate limits** under a default-on flip | degrades-quality | Settle **empirically** before flipping (M220 open question). |
| **Alignment CI is INERT while the corpus claims it runs weekly** | degrades-quality | Explicit `/align-run` step in M218; correct the doc. |
| **Scope creep via "hiring"** | (closed) | **D-DESIGN-4** — no hiring org. If it ever revives, it is a separate release (reserved M205). |
| **The manager's residual 11.6 s** (fully patched) | (de-risked) | **D-DESIGN-1** puts it out of gate scope; it is reported, not gated. Demo-patches remain available (D-DESIGN-2). |

---

## Done — v2.2 "panorama" (SHIPPED 2026-07-12, tag `v2.2`)

> **Status (SHIPPED 2026-07-12, tag `v2.2`):** designed 2026-07-11 via `/developer-kit:design-roadmap`; branch
> `release/02.20-panorama` cut from `main`, merged `--no-ff` → `main` + tagged **`v2.2`** at `/developer-kit:close-release`
> (2026-07-12); the release branch is deleted. rext code-of-record rolled to **`v2.2`** = `39e8013` (the D-CLOSE-1/-2/-3
> README reconcile + the ADV-1 F12 comment atop `panorama-m215` @ `00ba6b6`); `.agentspace/rext.tag` bumped `v2.1` →
> `v2.2`. **4 milestones M212 ✅ → { M213 ✅ ∥ M214 ✅ } → M215 ✅** (+ optional/deferrable **M216**, not scaffolded) — all
> merged `--no-ff`; M215 (the iterative acceptance closer) is `closed-on-gate`. **Close gates:** metrics GREEN (Go +8,
> flake 0, 0 net-new deps), deferral GREEN, supply-chain YELLOW-non-blocking (13 pre-existing unreachable `x/crypto`
> ssh-subpackage alerts in clerkenstein → cleared next rext roll). The **external-shareability release**
> — make dev/demo stacks reachable
> from other machines on a **Tailscale** tailnet (run a stack on a Tailscale VM, e.g. `billion.taildc510.ts.net` on
> the odyssey Proxmox host; a teammate with Tailscale up browses the demo end-to-end). **Tooling + docs + an opt-in
> flag only — zero platform-repo edits** (two platform-family files ride the EXISTING rext sha-pinned patch
> mechanism). **The sanctioned re-proposal** of the v1.4 seed "external stack shareability (Tailscale/ingress)"
> dropped 2026-06-11 pending a fresh design-roadmap run.
>
> **Theme:** *panorama — the whole environment, viewable from anywhere on the tailnet.* Today a stack is a
> single-seat show on the host's own `localhost`; panorama opens the house so any teammate on the tailnet can take
> in the whole thing, live.
>
> **Designed 2026-07-11** via `/developer-kit:design-roadmap`, from the user's Tailscale-serve briefing + a 5-agent
> feasibility workflow (`wf_bea3be47`) that mapped the browser-URL / Clerk-FAPI / CORS / injection surfaces
> (file:line) and **confirmed the core is config-only**: network reachability is ALREADY solved (ports bind
> `0.0.0.0`), `localhost` lives entirely in rext tooling (not platform files), and Clerk auth survives the move by
> design (dotted-host validator + host-agnostic token verify + path-only cert mount → `tailscale cert` drops in).
> Two narrow platform-family items (ant-academy `allowedDevOrigins`; conditional next-web `urls.ts`) ride the
> existing rext patch mechanism. **Live-verified 2026-07-11:** local + remote Tailscale up, MagicDNS
> `taildc510.ts.net`, `billion` VM has Docker + `tailscale cert`. **Decisions (2026-07-11):** HTTPS-everywhere under
> one MagicDNS origin (Clerk needs a secure context); external access **opt-in, default off** (an explicit
> `/demo-up --public-host` flag); demo-first (dev-path parity = optional M216).

### Execution graph

```
v2.2 "panorama"
  M212 ──┬──► M213  auth over the tailnet (TLS + proxy + pk) ─┐
  knob   └──► M214  origins & links (CORS + patch tail)       ┴──► M215  prove it on odyssey (iterative)
                                                                    [+ M216 dev-path parity — optional]
```

**M212 opens** (the single opt-in host knob threaded through every rext browser-facing emitter). **M213 ∥ M214**
then run on disjoint files — M213 (cert/proxy/pk) vs M214 (CORS/links/patches), additive merge — both consuming the
knob. **M215** is the iterative acceptance closer: the demo is not trusted reachable until a teammate on a
*different* tailnet machine drives a full hero journey green on a cold reset-to-seed.

### Milestones

(compact summaries — full contracts under
[`releases/archive/02.20-panorama/m*/overview.md`](releases/archive/02.20-panorama/))

**M212 — The single host knob** · `section` · complexity **medium** · depends: **none**. Introduce
`STACK_PUBLIC_HOST` (default `localhost` → **byte-identical when unset**) surfaced as an opt-in
`/demo-up --public-host` flag, threaded through every rext emitter that bakes a browser-facing `localhost`/`127.0.0.1`
(`up-injected.sh` build-args + `.env.local` + the `demo_web` rewrite + cache-validators; `inject.py --fapi-host`;
`gen_injected_override.py` host-param plumbing; cockpit `--host 0.0.0.0` + hosts; `ant-academy.sh`; `stack_registry`
additive `external_host`). Scoped to browser-facing values (control-plane DSNs stay `localhost`). **Delivers →** the
substitution surface (seed for the M214 recipe).
**Status:** `done` (closed 2026-07-11). All 12 sections landed on `m212/public-host-knob`, merged `--no-ff` into
`release/02.20-panorama`; rext code-of-record FROZEN at tag `panorama-m212` @ `770f81b` (rext re-tag deferred to
close-release). Tests **577** (569 pass / 8 skip / 0 fail) across the 3 touched rext suites, flake gate 5/5,
shellcheck clean. Byte-identity contract held (unset knob ⇒ identical to today). Decisions D-DESIGN-1 + D-IMPL-1..4;
close review = 1 routed finding (a rext-frozen README count-drift → close-release, D-CLOSE-1) + 4 adversarial
scenarios (all handled). Deferral audit GREEN (2 Fate-2 items → M214; 0 repeat/aged-out). Zero platform-repo edits.

**M213 — Auth over the tailnet** · `section` · complexity **medium** · depends: **M212** · ∥ **M214**. Serve
Clerkenstein auth + the whole browser surface under **one trusted HTTPS MagicDNS origin**: swap mkcert →
`tailscale cert` (path-only drop-in), re-mint the pk from the dotted MagicDNS FAPI host (validator-safe), stand up a
lightweight reverse proxy (tailscale serve / Caddy) fronting the browser-facing ports, keep FAPI same-host (SameSite
guard). Backend token verify is already host-agnostic → no verify change. **Delivers →** `clerkenstein.md` (the
tailscale-cert FAPI path).
**Status:** `done` (closed 2026-07-11). All 7 sections landed on `m213/auth-over-tailnet`, merged `--no-ff` into
`release/02.20-panorama`; rext code-of-record FROZEN at tag `panorama-m213` @ `b9f41dd` (rext re-tag deferred to
close-release). Delivered: the `tailscale cert` FAPI mint swap (path-only drop-in + local mkcert/openssl fallback,
proven live on billion — `ssl_verify_result=0` from a remote machine), dotted-pk validation up in the demo wiring
(`require_dotted_host`, codec stays permissive), the NEW `gen_tailscale_serve.py` reverse-proxy generator (per-port
HTTPS, 0 net-new deps, FAPI self-TLS-excluded), the FAPI-same-host topology guard (PSL-verified `ts.net` eTLD+1), the
confirmed build-rebuild-on-HOST guard, and the `cdn.jsdelivr.net` egress made explicit + overridable
(`FAKE_FAPI_CLERKJS_CDN`). Decisions D-SCHEME-1 + D-CERT-1 + D-PROXY-1/2 + D-PK-1 + D-TOPO-1 + D-REBUILD-1 + D-EGRESS-1.
Tests: go clerk-frontend +7 (`-race`+shuffle clean); stack-injection **152** (144p/8s/0f); demo-stack **367** (0f);
flake gate 5/5, `flake_count` 0; coverage gen_tailscale_serve 98% / inject 98% / clerk-frontend touched 100%. Close
review = 7 findings (0 scope · 0 code · 2 docs · 0 tests · 4 adversarial [all handled] · 1 triage): ref-tags added +
ADV-1..4 recorded + D-CLOSE-2 (rext README index row → close-release). Deferral audit GREEN (5 Fate-2 → M214/M215/
close-release; DEF-M212-02 dotted-host doc landed early Fate-1; 0 repeat/aged-out/escape-hatch). Zero platform-repo edits.

**M214 — Origins & links** · `section` · complexity **medium** · depends: **M212** · ∥ **M213**. Admit the MagicDNS
origin everywhere gated + close cross-surface ejects + land the bounded patch tail via the EXISTING rext patch
mechanism: extend `CORS_EXTRA_ORIGINS` emission (+ tests); studio-desk runtime redirects + the
`VITE_CLERK_SIGN_IN_URL` bake gap; **NEW `apply-*.sh`** for ant-academy `allowedDevOrigins` (required); conditional
next-web `urls.ts` demopatch. **Delivers →** NEW `corpus/ops/demo/tailscale-serve.md` (the remote-access recipe) +
`rosetta_demo.md` / `frontend-tier.md` / `clerkenstein.md` updates.
**Status:** `done` (closed 2026-07-11). All 8 sections landed on `m214/origins-and-links`, merged `--no-ff` into
`release/02.20-panorama`; rext code-of-record FROZEN at tag `panorama-m214` @ `99c86b7` (rext re-tag deferred to
close-release). Delivered: the `CORS_EXTRA_ORIGINS` https-MagicDNS emission (`browser_scheme` + the offset https trio,
byte-identical when unset), the studio-desk `CLERK_SIGN_IN_URL`/`WEB_APP_URL` host+scheme substitution, the
`VITE_CLERK_SIGN_IN_URL` overlay bake (gitignored `.env.production.local` + a transient `.dockerignore` re-include,
trap-reverted — also fixes the un-offset `:3000` for EVERY demo), the NEW `apply-ant-academy-dev-origins.sh` +
`ant-academy-dev-origins` sha-pinned patch (env-var indirection admits the MagicDNS host into `next dev`'s
`allowedDevOrigins` with a fixed post_sha256), the `$SCHEME` flip confirming the two shipped demopatches carry the
MagicDNS value + the mixed-content check, and the NEW `corpus/ops/demo/tailscale-serve.md` remote-access recipe. The
conditional next-web `urls.ts` landed as the evidence-decided **documented residual** (D-URLS-1 — the 0-prod-eject
coverage sweeps never surfaced `WEB_APP_URL`/`HIRING_APP_URL`). Decisions D-PATCH-1 + D-SCHEME-1 + D-VITE-SIGNIN-1 +
D-URLS-1. Tests: stack-injection **155** (147p/8s/0f), demo-stack **383** (0f) — +3/+16 net-new funcs; flake gate 5/5,
`flake_count` 0; Python surface 99% + shell behavior deepened at harden Pass 1. Close review = 4 findings (0 scope ·
0 code · 2 docs · 0 tests · 3 adversarial [all handled] · 2 triage): the `tailscale-serve.md` §-name corrected + the
`(#M214-D-URLS-1)` tag added + ADV-1..3 recorded + D-CLOSE-3 (rext READMEs don't index the new patch/helper →
close-release, bundled with M212's D-CLOSE-1 + M213's D-CLOSE-2). Deferral audit GREEN (the 2 items routed INTO M214
— DEF-M212-01 CORS emission, DEF-M213-01 recipe+links — landed Fate-1; 4 Fate-2 → M215, 3 → close-release; 0
repeat/aged-out/escape-hatch). Zero platform-repo edits.

**M215 — Prove it on odyssey** · `iterative` · complexity **large** · depends: **M212 + M213 + M214** · **the FINAL
milestone**. **Exit gate:** `/demo-up N --public-host billion.taildc510.ts.net` (opt-in) → a teammate on a
DIFFERENT tailnet machine completes a full employee AND manager hero journey (Clerkenstein login + a real journey)
with 0 localhost/prod ejects, 0 CORS blocks, 0 cert-untrusted, 0 mixed-content, assets rendering, on a cold
reset-to-seed; unset knob = byte-identical to today. **Iteration protocol:** the bring-up →
drive-from-a-2nd-machine → fix-in-the-M212/M213/M214-surface loop (`verification.md` + the coverage/playthroughs
gates from a remote origin). **Why iterative:** the last breakages (secure-context, mixed-content, cookie
same-site, cert PEM/renewal, RAM fit) surface only on a live cross-machine run.
**Status:** `done` (closed 2026-07-11, `closed-on-gate`). Merged `--no-ff` into `release/02.20-panorama`; the
`m215/prove-on-odyssey` branch is deleted. **Core exit gate MET:** the FIRST remote Linux-VM demo deploy over
Tailscale was driven end-to-end from a **DIFFERENT** tailnet machine (a remote Mac) for **both** vantages — employee
(`maya-thriving` → `/profile`, the M41 ProfileSeeder depth rendered) and manager (`dan-manager` →
`/enterprise/workforce`, real seeded structural data) — over a **genuinely trusted** Let's Encrypt cert
(`ignoreHTTPSErrors:false`, `verify=0`, no per-machine CA install), **0 console errors, 0 functional request
failures, 0 localhost/prod ejects**, assets rendering; then a clean **cold reset-to-seed** one-shot `--public-host`
bring-up proved reproducibility (14 containers, `tailscale serve` fronting 5 ports, seed 12,245 rows / 541 users,
login-ready from the Mac). Unset knob byte-identical. Driven directly (live shared-infra work — D-STRAT-1); canonical
record `iter-01/findings.md` (findings F1–F13). The **user-directed propagation close-gate is SATISFIED**: every
deployment finding (F1/F2/F4/F6/F8/F9/F12) landed in **tools (rext) + KB + skills**, and a fresh reader can stand up
a remote demo on a new Linux VM unaided from the NEW `corpus/ops/demo/tailscale-serve.md` runbook. rext code-of-record
FROZEN at tag **`panorama-m215`** @ `00ba6b6` (rext re-tag + `.agentspace/rext.tag` bump reserved for close-release):
the host pre-flight (F1/F2/F8) + keyless ssh-agent (F4) + Linux data-dir perms (F6) + the `git for-each-ref` build-tag
resolver (F3) + the `tailscale serve` teardown/up-path reset (F12); demo-stack **424** (+41) / stack-injection
**147p/8s**, shellcheck clean, macOS/dev path byte-identical. Close review = 4 findings (0 scope · 0 code · 2 docs
fixed · 0 tests · 1 adversarial [ADV-1, non-blocking] · plan-artifact backfill): the `tailscale-serve.md` `F1–F11`→
`F1–F12`/`F1–F13` range + the setup_guide §-anchor label; the stub `progress.md`/`decisions.md` backfilled. Deferral
audit **GREEN** (11 records: 3 inherited→M215 all landed Fate-1/resolved; 3 D-CLOSE-1/-2/-3 → close-release; 4 new
F5/F9/F11/F13 → standing backlog Fate-2; 0 repeat/aged-out/escape-hatch). **4 documented non-blocking residuals**
routed via [`carry-forward.md`](releases/archive/02.20-panorama/m215-prove-on-odyssey/carry-forward.md). Zero platform-repo
edits. **This is the FINAL v2.2 milestone — all 4 core milestones (M212, M213, M214, M215) are now closed; the
release is ready for `/developer-kit:close-release`.**

**M216 — Dev-path parity + operator surface** · `section` · **OPTIONAL / deferrable** · depends: **M215**.
Generalize the knob to the DEV path (main dev `PUBLIC_HOST` + native-worktree `.env.local`) + a
`/dev-up --public-host` flag + full corpus docs. **Defer-safe:** the demo path (M212–M215) is the shippable core;
dev parity can be Fate-2 if scope tightens. *(Not yet scaffolded — roadmap-only until promoted.)*

### Top risks

1. **Secure-context / auth (blocks-release if wrong)** — Clerk's clerk-js uses Web Crypto, granted only on
   `localhost` or HTTPS; a plain-`http://` MagicDNS origin is not a secure context. **Mitigation:** HTTPS-everywhere
   (the M213 decision) + burn down early on the live box (M215). *The make-or-break surface — config-only ONLY
   because the dotted-host validator + host-agnostic verify + path-only cert mount all align.*
2. **Mixed-content (degrades-quality)** — an https page making an http browser call is blocked independent of CORS.
   **Mitigation:** the reverse proxy terminates one HTTPS origin; M214 mixed-content check.
3. **RAM on `billion` (nice-to-resolve)** — ~13 GB avail vs ~12 GB UI-tier want. **Mitigation:** the `--no-ui`
   escape or a VM RAM bump on odyssey; validated in iter-01 of M215.
4. **Build-time bake tax (degrades-quality)** — browser URLs + the pk are baked into next-web/studio-desk images; a
   host change forces a per-host rebuild. **Mitigation:** the M212 cache-validators must invalidate on HOST, not
   just offset.
5. **Patch-tail drift (nice-to-resolve)** — ant-academy `allowedDevOrigins` is a CONFIRMED patch; sha-pinned
   drift-refusal makes an upstream `next.config.js` change fail loudly (by design).

---

## Done — v2.1 "quick change" (SHIPPED 2026-07-09, tag `v2.1`)

> **Status (SHIPPED 2026-07-09, tag `v2.1`):** designed 2026-07-08 via `/developer-kit:design-roadmap`; branch
> `release/02.10-quick-change` cut from `main`, merged → `main` + tagged `v2.1` at close-release; rext code-of-record
> rolled to `v2.1` (= `quick-change-m211`). **4 milestones M208 → M209 → M210 → M211, strictly sequential — ALL
> `done`.** close-release GREEN: both blocking gates (deferral + metrics) GREEN + adversarially verified, 0 must-fix,
> all should-fix landed (the 42,790 count reconcile, the `go1.25.12` toolchain bump, the doc-flip fidelity), triple-clean
> 3/3, 0 net-new deps. **The merged skiller-in-app platform stands up cold on both stacks via the re-grounded tooling.**
>
> **Theme:** *quick change — backstage, the actor sheds one costume and re-enters as another, seamless to the
> audience.* skiller's part now folds into `app`; v2.1 re-fits the whole apparatus to the changed stage. A
> **field-hardening re-ground** in the **v1.3b "dress rehearsal" / v1.10b "fit-up"** lineage — but triggered by a
> **landed platform structural change**, not a bring-up defect: the standalone `skiller` Go service **merged into
> the `app` monolith** — its domain (60K-skill taxonomy, embeddings, job-roles, matching) now lives in the
> **`public` schema** (table names unchanged, only the schema prefix changed `skiller.X → public.X`; the old
> `skiller` schema is dropped), its Connect-RPC surface is served by `backend`
> (`SKILLER_RPC_ADDR=http://backend:8083`), its GraphQL subgraph left the federation (**4 subgraphs** now), and
> the skiller repo + container are gone (`make init` no longer clones it). **This is landed upstream** (platform
> `origin/main` @ `0808b92` drops skiller from compose + repos.yml; app @ v1.334 carries the "Deprecate skiller
> schema" merge) — so v2.1 is **tooling + docs + stack-re-sync only; zero platform-repo edits** (the platform
> already did its half).
>
> **Why a release:** three surfaces are stale + mutually inconsistent. (1) **Stacks** — `app` is post-merge but
> `platform` is 2 commits behind (still composes skiller); both stacks hold vestigial `skiller/` clones. (2) **rext
> tooling** — untouched by the docs branch; still queries `skiller.<table>` (seeding taxonomy resolvers, snapshot
> capture/replay) and probes a skiller container/schema (`readiness.sh`, `services.sh`, `up-injected.sh`), so it
> **breaks** on the merged platform. (3) **Corpus** — the colleague's `docs/skiller-in-app-merge` sweep is
> **correct-but-incomplete**: the architecture/subgraph half is solid, but 5–6 rext-facing tooling docs still
> describe `skiller.*` and it **cannot land independently** without contradicting the tooling.
>
> **Designed 2026-07-08** via `/developer-kit:design-roadmap`, from the user's skiller-merge briefing + the
> colleague's `docs/skiller-in-app-merge` branch. A 7-agent research workflow (`wf_08b6bf4a`) mapped the per-module
> blast radius (file:line), **adversarially confirmed** the firewall public-predicate **survives** the merge
> (`organization_id IS NULL` still isolates public taxonomy — no data-leak risk; `skiller_mixins.OrganizationIDMixin`
> ports the tenant boundary 1:1), and confirmed the docs branch **cannot land present-tense** before the rext
> re-ground + stack re-sync. **Phase-0: 🟢 GREEN** deferral audit (M25-D9 opportunistic Fate-1 on the re-sync
> migration path) + clear KB blind-areas (every topic has an anchor).

### Execution graph

```
v2.1 "quick change"   (strictly sequential — single track)
  M208 ─────→ M209 ─────→ M210 ─────→ M211
  re-sync     rext        corpus      bring-up
  +ground     re-ground   re-ground   acceptance (iterative)
```

**Strictly sequential** (the user's choice). **M208** establishes current merged code — everything grades against
it. **M209** re-grounds the rext tooling (+ recapture the snapshot from merged-prod). **M210** completes the corpus
in lockstep with M209's landed schema (the tooling-doc bodies flip to `public.*`). **M211** is the iterative
acceptance closer — bring-up isn't trusted until `/dev-up` + `/demo-up` both go GREEN cold on the merged platform.
*(Scope-flex from design: if the corpus reconcile proves large, M210 can split into land+reconcile / rext-gap-fill;
4 is the clean target.)*

### Milestones

**M208 — Re-sync & merged-schema ground-truth** · `section` · complexity **medium** · depends on: **none** ·
✅ **`done` — closed 2026-07-08** (merged → `release/02.10-quick-change`). The **load-bearing foundation** of
v2.1 (mirrors v1.10b M47) landed: re-synced both stacks to the merged platform (`app` `a848cccb→c3c45e01`
v1.334.1 — the **86-commit merge pull**; `platform` `5e1ae6b→0808b92` — skiller gone from
compose/repos.yml/Make; sibling set current), removed both vestigial `stack-*/skiller/` clones, and **retired
the #1 release risk GREEN** via a live containerized de-risk on stack-dev: cold `make up` rc=0 with a
**4-subgraph** compose and **no skiller container** (`SKILLER_RPC_ADDR=http://backend:8083`); a clean-slate
`make reset-db`+`make migrate` builds the **full `public` taxonomy from scratch** (skills w/ `organization_id`,
job_roles, job_role_skills, skill_embeddings, categories, specializations) with **no `skiller` schema on a clean
DB**; measured prod `public.skills WHERE organization_id IS NULL` = **42,790** (the roadmap's ~42,763 figure).
Pinned the authoritative **merge fact-sheet** in `corpus/services/backend.md` (§ Skiller-in-app merge + banner)
+ a stub banner on `corpus/services/skiller.md`. Committed rosetta diff is **100% documentation** (zero
code/test) → HARDEN N/A; close review found **0 findings**; deferral audit **GREEN** (5 single, 0 repeat).
**Two bring-up findings routed forward (both user-accepted):** Finding 1 (clean-bring-up `extensions`-schema
bootstrap + PG-readiness, the M25-D9 class — did NOT fall out as a trivial Fate-1) → **Fate-3 M211** (+ M209
Risk-2 cross-ref); Finding 2 (`INVITATION_HMAC_SECRET` dev-`.env` completeness gap, not merge-caused) →
**Fate-2 M211 / `/stack-secrets`**. Zero platform-repo edits. **Delivers →** the merge fact-sheet
(`corpus/services/backend.md` + `corpus/services/skiller.md` stub).
**Goal:** bring both stacks (and the snapshot's target reality) current with the merged platform, and pin the
authoritative merge fact-sheet — so every downstream fix grades against current code.
**Scope:**
- **In:** `make pull` stack-dev + stack-demo `platform` to `origin/main` (skiller gone from compose/repos.yml),
  pull `app` to current (v1.334, post-merge domain) + the sibling repo set; capture before/after refs. Remove the
  vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones. Rebuild images + re-run migrations against the
  merged `public` schema; confirm the 4-subgraph compose (`backend`, jobsimulation, cms, skillpath), no skiller
  container, `SKILLER_RPC_ADDR=http://backend:8083`. Pin the **merge fact-sheet** (moved tables in `public`, the
  confirmed `organization_id IS NULL` public predicate, the ~42,763 public-skill count, the re-pointed RPC).
  Opportunistic **M25-D9** dev migrate-ordering fix (Fate-1 — lives on this path).
- **Out:** rext code (M209); corpus body re-point (M210); live bring-up (M211).
**Delivers → knowledge/corpus:** the merge fact-sheet (anchored in `corpus/services/backend.md` +
`corpus/services/skiller.md` stub).
**Open questions:** does the 86-commit `app` pull + migration re-run surface a schema/migration issue (⚠ the fit-up
M47 risk class)? — bounded; capture before/after.

**M209 — rext tooling re-ground** · `section` · complexity **medium** · depends on: **M208**.
**Status:** `done` (completed 2026-07-08).
**Goal:** re-point every rext tool that queries the old skiller schema or expects the skiller service to the merged
reality, recapture the snapshot, and tag a new rext.
**Scope:**
- **In (snapshot):** flip `stack-snapshot/taxonomy/taxonomy.go:43 const Schema "skiller"→"public"` (re-grounds
  capture *and* replay); update the 2 load-bearing `taxonomy_test.go` PublicVia assertions; **narrow the
  `pg.SchemaVersionSQL` staleness digest** to the surface's enumerated tables (fixes the post-merge whole-monolith
  cache-thrash — Risk 1); **verify the capture SELECT column list** vs merged prod (Risk 2 —
  `embedding→small_embedding3`, `extensions.`-qualified GIN opclasses — the one non-mechanical bit); keep
  `AssertPublicOnly` + add the ~42,763-row post-capture assertion; **recapture** the public taxonomy from
  merged-prod into `.agentspace/snapshots/` (bump the capture version; the batch cache re-keys).
- **In (seeding):** re-point the 5 real-SQL files (`seeders/taxonomyref.go`, `skillref_named.go` [the shared
  `namedSkillSelect` const → also fixes `curated_pools.go` + `ai_readiness_config.go`], `jobroleref.go`,
  `taxonomy_snapshot.go`, `dna/fidelity_probe.go`) `skiller.*→public.*` keeping the `organization_id IS NULL`
  public-pool predicate; drop skiller from `isolation/isolation.go` schema-note + re-ground `dna/data-dna.json`
  golden (schema + FK ref_schema); rename the 111 fake-Conn test string-matchers in lockstep; reword the
  comment/attribution refs (incl. `services/ai`).
- **In (small):** remove skiller from `stack-verify/lib/readiness.sh` `probe_postgres_schemas` + `services.sh`
  container probe; drop skiller from `demo-stack/up-injected.sh` INJECT_SVCS (else it clones/builds a gone repo) +
  the "5 Clerk-token services"→4 note.
- Build + test the authoring copy; **tag a new rext (`v2.1`)**; prepare the per-stack consumption re-pin.
- **Out:** the stack re-sync (M208); doc bodies (M210); live bring-up (M211).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/safety.md`.
**Closed 2026-07-08:** re-grounded rext `skiller.*→public.*` across **stack-snapshot** (the `taxonomy.go` const
flip re-grounding capture+replay; the Risk-1 `Surface.VersionTables()` digest-narrowing so taxonomy digests only
its 10 tables while a structure-bearing surface still whole-schema-invalidates; a one-sided `MinRows` under-capture
floor; Risk-2 verified — the capture is names-only/type-agnostic so the `extensions.`-qualified merged columns need
no change), **stack-seeding** (24 files, `organization_id IS NULL` preserved, `data-dna.json` golden, `isolation.go`),
and the **small shell modules** (readiness/services/up-injected/migrate-demo, 5→4 services; stack-verify Python
104/104). **6 Go modules GREEN, `go vet` clean, 5× flake-clean; 0 `skiller.<table>` queries in any production path.**
rext tagged `quick-change-m209` (build `00a3ec5`/`e458acf`/`75bc4cf` + harden `42ad600`/`72a5259`/`2f06e78` — 14
harden test funcs, 0 bugs, 0 flakes; tag re-pointed to the post-harden HEAD `2f06e78` at close). Close: **1
nice-to-have finding** (a pre-existing rext `stack-seeding/README` test-count drift, last reconciled M41 — routed
to the v2.1 rext roll; rext frozen at `2f06e78` this close), **deferral audit GREEN** (10 in scope, 0 repeat/aged/
escape). **Recapture Fate-3→M211** (tooling READY; no local COPY-byte capture source — a data op M211 owns).
The `v2.1` rext roll + `.agentspace/rext.tag` consumption re-pin remain **`/developer-kit:close-release`'s** job.

**M210 — Corpus + skills re-ground** · `section` · complexity **medium** · depends on: **M209** (the tooling-doc
bodies flip to match M209's landed schema).
**Status:** `done` (completed 2026-07-08).
**Goal:** land the colleague's docs sweep as the *complete, internally-consistent* corpus re-ground.
**Scope:**
- **In:** adopt/validate `origin/docs/skiller-in-app-merge` (correct-but-incomplete — the architecture/subgraph
  half is solid); fix the fully-missed `profile-completeness-spec.md` (43/44→44/44); **flip the 5–6 rext-facing
  tooling-doc bodies to `public.*`** (`snapshot-spec.md` [26 mentions — the taxonomy surface enumeration +
  FidelityProbe gene + capture predicate], `safety.md`, `recipe-snapshot-world.md`, `stories-spec.md`,
  `seeding-spec.md`, `coverage-protocol.md`) + delete the interim disclosure notes; reconcile the db-access ↔
  tooling contradiction; sweep the skill files (`dev-up/reference`, `stack-snapshot/SKILL`, `stack-update/reference`,
  `db-query/SKILL`) so container counts / migration lists / RPC addr / subgraph counts match the re-synced stacks;
  update `CLAUDE.md` service catalog.
- **Out:** rext code (M209); live bring-up (M211).
**Closed 2026-07-08** (merged → `release/02.10-quick-change`): made the corpus internally consistent with the merged
platform + M209's landed `public.*` rext code. Adopted the colleague's correct architecture/subgraph/service half
(28 files; each hunk verified vs the M208 fact-sheet + the re-synced `stack-dev/app` clone; kept M208's authoritative
`backend.md`/`skiller.md` fact-sheet — no duplicate merge section); fixed the profile-completeness node-id prose
(**verified NO literal "43/44" exists anywhere** — the design-note count was inaccurate; made the one genuine
merge-sweep fix, did **not** fabricate a phantom count); flipped the **6 rext-facing tooling-doc bodies + directus-local.md**
`skiller.*→public.*` and deleted the interim disclosure notes; reconciled db-access ↔ tooling on `public.*`; swept the
4 skill files to the verified merged compose (no skiller container, **4 subgraphs**, `SKILLER_RPC_ADDR=http://backend:8083`;
superseded the colleague's now-stale "still-targets-skiller/exit-4" note with an accurate M209-done note); updated the
`CLAUDE.md` catalog. Email-asset PNGs excluded. **Verified core outcome:** 0 stale `skiller.<table>` tooling-query refs
corpus-wide, 0 leftover interim notes, 4 subgraphs consistent, 0 broken `.md` links. **RESOLVES the KB-1/2/3 body-flip
deferrals** M208+M209 routed Fate-2 → M210 (7 defers landed at destination). Committed rosetta diff is **100%
documentation** (50 `.md`, 0 code/test) → HARDEN N/A; close review **0 must-fix / 1 nice-to-have no-change-needed**
(the app==backend subgraph dual-naming is the corpus's established convention, not an M210 defect); deferral audit
**GREEN** (11 in scope, 7 resolved, 4 still-open confirm-only → M211/close-release; 0 repeat/aged/escape). Zero
platform-repo edits.
**Delivers → knowledge/corpus:** the completed corpus (`corpus/services/skiller.md` stub + `backend.md` ownership
+ re-pointed tooling docs).

**M211 — Bring-up acceptance: dev-up + demo-up green on the merged platform** · `iterative` (closed-on-gate) ·
complexity **large** · depends on: **M209 + M210** · **the FINAL v2.1 milestone**.
**Status:** ✅ **`done` — closed-on-gate 2026-07-08** (merged → `release/02.10-quick-change`). **Gate 6/6 MET** —
proved the merged (skiller-in-app) 4-subgraph platform stands up end-to-end via the re-grounded tooling on BOTH
stacks, cold, with **zero platform-repo edits**. Delivered (all in the rext tooling, tag `quick-change-m211` =
`2039103`): the **cache-migration** (real 42,790-row taxonomy + 274 sims re-keyed `skiller.*→public.*`, replayed
— the no-prod-access recapture); the **root-cause fix** (the injected build-scratch was pinned pre-merge + survived
`--purge` → every rebuild baked a stale pre-merge binary → federation `_entities(Skill.name)` 422s; fixed to
re-sync the scratch to the source's release tag every bring-up); the dev bring-up casbin `init_policy.sql` load
(silent-403 fix); frontend offset-reuse guard; demo-local `ACADEMY_URL` bake + academy-aware cross-port hook;
demopatch URL re-pin (next-web v2.106.1); the Playthroughs reset-to-seed **roster-refresh**; and the new
**`dev-stack/migrate-dev.sh`** (dev cold DB-init — extensions + `gin_trgm_ops` + casbin, mirror of
`migrate-demo.sh`, the M25-D9 path). **Result:** cold `/demo-up` GREEN end-to-end; **M42 coverage GREEN both
vantages**; **v2.0 Playthroughs 10/11 GREEN** (1 declared in-manifest TODO); dev cold DB-init cold-verified on a
faithful non-destructive throwaway + a live docker harness; 0 residual skiller. **17 iters** (1 bootstrap tok +
16 tiks) + 4 stabilized final harden passes. Close: rosetta diff **docs+plan only** (3 corpus/skill docs + plan
records, 0 code — the tooling code is the frozen rext tag) → HARDEN N/A for rosetta; the rext suites spot-verified
GREEN (Go/vet, demo 114, dev migrate-dev 14, TS 32). Close review **1 finding** (DOC-1: `migrate-dev.sh` not
indexed in rext `dev-stack/README.md` → **Fate-2 → close-release rext roll**, bundled with TEST-1; rext frozen).
Deferral audit **GREEN** (7 in scope, 0 blocking; DEF-M208-01/M25-D9 RESOLVED, TEST-1 aged→re-fated Fate-2,
CAVEAT-1 belt-and-suspenders backlog). Two close-review caveats (non-blocking): (1) a literal full destructive
clean-box `/dev-up` was **deliberately not run** — this box is committed to the user's native-app content-line dev
(`docker-compose.override.yml` → `backend:host-gateway` + an `app-01.10-content-line` worktree); the sole
dev-specific gate delta (the M25-D9 DB-init) was cold-verified on a faithful throwaway — an environment-respecting
gate interpretation, with a clean-box full `/dev-up` recorded as a belt-and-suspenders backlog note; (2) 33
pre-existing `test_dev_stack.py` CLI failures from an incomplete local `.agentspace/secrets` source (unrelated to
M211). **→ v2.1 is complete; run `/developer-kit:close-release`** (rolls the rext `v2.1` tag, bumps
`.agentspace/rext.tag`, merges → `main`, reconciles the rext READMEs [TEST-1 + DOC-1]).
**Goal:** prove the whole chain works end-to-end on the merged platform with the re-grounded tooling.
**Exit gate:** from a re-synced state, **`/dev-up` AND `/demo-up` both go GREEN cold** — 4-subgraph compose / no
skiller container; snapshot **recapture→replay** loads `public.*` (taxonomy replay exits 0, ~42,763 public skills);
**seed** resolves real public node-ids (closure green); **verify** (`verification.md` net) passes with a
merged-platform assertion (no skiller schema/subgraph/container; `readiness.sh` schema probe green); the M42
coverage sweep + the v2.0 Playthroughs suite stay GREEN; **0 residual skiller-schema references** in any queried
path.
**Iteration protocol:** the fit-up/dress-rehearsal fix→re-measure→re-run bring-up loop
(`corpus/ops/verification.md` + the coverage/playthroughs gates).
**Why iterative:** the merged 4-subgraph platform has never been stood up locally with the re-grounded tooling; it
*will* surface fix-loops (migration ordering, the column-mapping caveat, vestigial container/clone cleanup, cache
behavior).

### Top risks

1. **Cache-key digest regression** (degrades-quality) — post-merge `SchemaVersionSQL` digests the whole app
   monolith → taxonomy cache thrashes on any migration. **Mitigation:** narrow the digest to enumerated surface
   tables (M209). *The single non-obvious regression the merge introduces.*
2. **Capture column-mapping** (blocks-release if wrong) — the SELECT list may not be a pure prefix swap
   (`small_embedding3`, `extensions.`-qualified opclasses). **Mitigation:** verify vs merged-prod in M209; the
   ~42,763-row assertion catches empty/over-broad capture.
3. **86-commit app pull + migration re-run** (M208) — the fit-up M47 ⚠ class. **Mitigation:** bounded; capture
   before/after refs.
4. **Docs lockstep** (would create a self-contradicting corpus) — the branch can't land present-tense before rext +
   re-sync. **Mitigation:** M210 flips bodies in lockstep with M209 (why the design keeps them adjacent + sequential).
5. **Recapture safety** — verified **low**: the firewall predicate HOLDS (`org_id IS NULL` survives the merge) +
   `AssertPublicOnly` runtime net rejects any non-null-org captured row.

---

## Done — v1.10b "fit-up" (SHIPPED 2026-07-01, tag v1.10.1)

> **Theme:** *fit-up — build and rig the set correctly in the venue before opening night.* An **interposed
> field-hardening backfill** in the **v1.3b "dress rehearsal"** lineage. A from-scratch `/demo-up` surfaced 8
> bring-up issues + a tail of v1.10 content gaps. **CORRECTION (M47 finding):** the M201 close *reported* the
> clones ~5 weeks / 115+ commits behind prod, but **M47 found them CURRENT** (next-web @ v2.89.0, every repo ≤2
> behind; the AI-readiness feature present) — the re-sync was a trivial `make pull`. The real stale surface is the
> **corpus** (M48). v1.10b **recaptures** the snapshot from current prod, **re-grounds** the corpus, **fixes** the
> bring-up + content issues, **adds** a curated **AI-readiness showcase org** (redeeming the M201
> member-AI-readiness false-negative), and **consolidates** one auditable **seed+generation manifest**.
> **Tooling + docs only — zero platform-repo edits.** The v1.x flat counter re-opens at **M47** (M47→M53); tag
> **`v1.10.1`**; branch `release/01.10b-fit-up`.
>
> **Designed 2026-06-29** via `/developer-kit:design-roadmap`, from the field review in
> [`.agentspace/annotation.md`](../../.agentspace/annotation.md) + the M201 stale-clone finding. Three research
> agents mapped the fix surfaces (file:line), the content/seeding gaps, and the KB blind-areas (all homed below via
> `Delivers →` lines).

### Execution graph

```
v1.10b "fit-up"   (ONE live demo → verification serializes; only rext authoring parallelizes)
  M47 ──→ ┌ M48  corpus re-ground   (reads code, NO demo) ─┐
          └ M49  bring-up hardening (NEEDS the live demo)  ─┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53
 re-sync                                                       content  AI-ready  manifest  cold-rebuild
```

**The 1-demo-stack constraint shapes the graph.** The user's box hosts **one** demo at a time, so milestone
**verification serializes** on the single live stack — the release is an honest sequential chain ending in a
**cold destroy-and-rebuild acceptance** (M53). The **one** genuine parallel pair is **M48 ∥ M49**: M48 (corpus
re-ground) is pure docs-vs-code and never touches the demo, while M49 (bring-up hardening) monopolizes it —
disjoint file clusters (`architecture`+`services` vs `ops`+rext), additive merge. The "fix-on-live → final cold
rebuild" the user asked for *is* the M47→…→M53 shape.

### Milestones

(compact per-milestone summaries — the full contracts live under
[`releases/archive/01.10b-fit-up/m*/overview.md`](releases/archive/01.10b-fit-up/))

**M47 — Re-sync & recapture** · `section` · depends: **none** (foundation) · ✅ **`done` — closed 2026-06-29**
(merged → `release/01.10b-fit-up`; rext tag `fit-up-m47`). The flagged heavy re-sync was a **no-op** — the clones
were already current (next-web @ v2.89.0); the genuine staleness is the corpus (→ M48). Delivered: `pg.NormalizeDSN`
(sslmode `no-verify→require`) so the wired MCP DSN works as a capture `--dsn` (demo-up #2, proven by a live dry-run +
9 tests); all 3 snapshot surfaces recaptured from current prod (digests unchanged); the **AI-readiness feature
confirmed present** (M201 false-negative resolved); `snapshot-cold-start.md` updated (the MCP-configured-DSN path,
KB-47-01). The `up-injected.sh` auto-capture wiring was descoped per the user (no new entry point — D3); re-pin
deferred (push-gated). **Delivered → `corpus/ops/snapshot-cold-start.md`.**

**M48 — Corpus re-ground** · `section` · depends: **M47** · ∥ **M49** · ✅ **`done` — closed 2026-06-29** (merged →
`release/01.10b-fit-up`). Delivered (3-agent investigation of the current clones): **NEW
[`corpus/services/ai-readiness.md`](../../corpus/services/ai-readiness.md)** documenting the previously-undocumented
member-AI-readiness feature (org enablement gate, 3-step 30/40/30 scoring, the 9 `ai_readiness_*` tables, GraphQL+REST
interface, surfaces, narratives) **+ the M51 seeding contract** (Phase-2c-sharpened: active-cycle ⇒ signals-true,
closed-cycle ⇒ frozen-snapshot-direct — the dashboard recomputes from signals). Reconciled the material drift
(backend.md / next-web-app.md / architecture_overview.md / service_taxonomy.md now name the subsystem) + corrected the
false ant-academy "in repos.yml" claim (CLAUDE.md + ant-academy.md×3; **M49 #5 owns the repos.yml fix**). Docs-only —
never touched the demo. **Delivered → `corpus/services/ai-readiness.md`** + the re-grounded arch/service docs.

**M49 — Bring-up hardening + truth-up** · `section` · depends: **M47** · ∥ **M48** · ✅ **`done` — closed
2026-06-30** (merged → `release/01.10b-fit-up`; rext tag `fit-up-m49` @ `ba586d6` = 7 build fixes + 3 harden
commits, consumed per-stack). Closed the 7 remaining demo-up issues so a from-cold `/demo-up` on a `stack-demo`-only
box completes — **proven end-to-end by the live-verify gate** (a from-cold `/demo-up` on a re-pinned `fit-up-m49`
clone: demo-1 UP, autoverify "verified-working"): #3 `.env`-guard order (provision-then-check), #4
`INVITATION_HMAC_SECRET` (critical secret-DNA gene + values-blind auto-gen + `DemoGeneratedKeys` overlay; the silent
`app Exited (0)` class fixed), #5 ant-academy **explicit `ensure-clones.sh` clone** (NOT `repos.yml` — the ephemeral
gitignored platform clone makes that non-durable + a platform edit; the cms/studio submodule-pattern precedent), #6
disk pre-flight + `demo-down --purge` per-demo image cleanup (~5 GB reclaimed live), #7 *true* non-fatal frontend
(absent image → `--scale svc=0`), #8 demopatch re-anchor to next-web **v2.89.0**, #1 single **`.agentspace/rext.tag`**
source-of-truth (+ CRLF-tolerant `lib/rext_tag.sh` reader; reconciled 4 stale prose pins; doubles as the note-2
reproducible pin). Harden fixed 1 real bug inline (the `rext_tag.sh` CRLF carriage-return leak). Tests: rext Go
1552→**1555** (stack-secrets +3) · demo-stack Python **299** (demopatch 46→47); flake **0** (5/5). Close review: 6
findings all Fate-1 (2 stale `make init` ant-academy claims in `service_taxonomy.md` retired; test-count literals
reconciled); deferral audit **GREEN** (AI-keys policy → M50 Fate-2; consumption-clone re-pin → push-gated KEEP).
**Delivers →** `rosetta_demo.md`, `frontend-tier.md`, `secrets-spec.md`, `ant-academy.md`.

**M50 — Content & seeding fill** · `iterative` · depends: **M49** (+M48). ✅ **`done` — closed-on-gate 2026-06-30**
(merged → `release/01.10b-fit-up`; rext tag `fit-up-m50` @ `f0d984c` + close fix-commit `3c2de27`, consumed
per-stack). **M42 semantic coverage gate MET on BOTH vantages** (employee + manager) on a WARM demo-1, on the
manifest **STRENGTHENED to PROVE the gaps** (frontier-exhausted, (failingSections,escapes)=(0,0), 0 persona,
0 cross-port). 6 iters (1 tok + 5 tiks). Delivered (the sweep-driven seed fill): a NEW **`MemberLanguagesSeeder`**
(ISO-639-1 `world_languages` catalog + per-member `user_languages` → `membership_languages` via the platform's
AFTER-INSERT trigger — the manager Talent-tab "Languages spoken" chart, 0 rows → 747 across all 340 members), the
**`CertificatesSeeder` member-coverage extension** (hero-only → ~45% role-coherent, certs 9→236), the **`UsersSeeder`
member-field backfill** (`memberships.joined_at`/`location`/`last_activity_date` for `/enterprise/members`,
NULL-only idempotent guard), the **`next-web-public-website-url` demopatch** (the `PUBLIC_WEBSITE_URL` JS-constant
escape class) + a NEW **post-replay Directus content-URL rewrite** (the replayed-content escape class — prod hosts
baked into `public_landing_page_url`/`read_more_link`, regex over any `anthropos.work` subdomain → demo host), and
the **manager-manifest strengthening** (new `preAssert` tab-click + `textMatch` OR-assert harness primitives that
ASSERT the members-Location column + the Talent-tab languages/certs charts — the D4/F1 reconciliation: the run-1
gate passed BLIND to two M50-own gaps; the gate now PROVES them). All DATA-DENSITY, `PerStackIsolated` +
closure-GREEN, M17-idempotent, zero platform edits. **AI-keys policy DECIDED: documented-as-absent** (values-blind,
no key provisioned; AI surfaces inert-by-design — resolves the inherited M49 deferral). Tests: rext stack-seeding
719 (seeders pkg 349, +1 at close; 97.4% stmt) · demo-stack Python 108; flake **0** (5/5). Close review: 11
findings all Fate-1 (gofmt 2 files + 2 code pins + 3 docs incl the NEW routing-table escape row + a broken M51
backref); deferral audit **GREEN** (0 repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):**
**COLD reset-to-seed acceptance → M53** (Fate-2, user-decided — all M50 seeders + fixes reproduce from the bring-up
tooling on a fresh `/demo-up`); **academy content + menu-link/non-anonymous-session (F6) → M51** (Fate-3);
consumption-clone re-pin to `fit-up-m50` = push-gated KEEP (authoritative at M53). **Delivers →**
`profile-completeness-spec.md`, `seeding-spec.md`, `coverage-protocol.md`, `secrets-spec.md`.

**M51 — AI-readiness showcase org** · `iterative` · depends: **M48** (the feature contract) + **M50**. ✅ **`done` —
closed-on-gate 2026-07-01** (merged → `release/01.10b-fit-up`). **Exit gate:** a curated **200-person 3rd org** with
the AI-readiness manager dashboard **enabled**, **~80%** of members having completed all **3** onboarding/evaluation
steps, **1 hero started + 1 hero completed** — proven by the coverage gate. **MET at iter-09:** manager-vantage
`(failingSections, escapes) = (0, 0)` frontier-exhausted on a fresh demo-up (reachable 70, personaFailures 0);
org **Northwind Aviation** (200 members) ENABLED, **78.4%** all-3-complete, **Ben STARTED** (stage 1) + **Aria
COMPLETED** (stage 3), cycle `closed` + 199 frozen snapshots. **9 iters** (1 bootstrap tok + 8 tiks). The
strategy arc: iter-02/03 landed the 3rd story + the 3 net-new seeders (`OrgSettingsSeeder` +
`AIReadinessConfigSeeder` + `AIReadinessFunnelSeeder`); iters 06/07/08 falsified three read-fast strategies
(active-signals → closed-cycle-snapshot → deep-link) against the platform AI-readiness read-path org-scale perf
wall (**"frozen SCORES ≠ frozen RESPONSE"** — `buildResponseFromSnapshots` re-joins members via an unbounded
whole-org `loadMembers`); **TOK-02** (user-authored, triggered by the 3-tik no-prog streak) pivoted to the
iter-09 **`app-aireadiness-snapshot-loadmembers`** read-path demo-patch (a PURE data-identical bound of that
hydration to the ~199 snapshot users → the frozen `?cycle=` GET 180s-timeout → 19ms). Tests: rext stack-seeding
**749** (seeders pkg 382, **97.6%** stmt, +30 vs M50's 719 across the iters + final harden + close) · e2e TS unit
**33** (+13 at close: the new `section-assert.ts` no-browser spec); flake **0** (5/5 Go + 5/5 TS). Close review:
16 findings all Fate-1 (S1 the 3rd AI-readiness story authored → **`delivers` MET**; C1/C3/C4 code + T1/T2 tests +
D1–D5 docs). Deferral audit **RED→CLEARED**: the academy **F6** repeat-defer (M50→M51) fated **LAND-NEXT → M53**
(Fate-3, user-decided — the cold rebuild is where academy content is seeded + verified). **Carry-forward (three-fate,
none escape-hatch):** academy F6 → M53 (Fate-3); COLD reset-to-seed acceptance → M53 (Fate-2); consumption-clone
re-pin + `.agentspace/rext.tag` bump → M53 (push-gated KEEP). **Delivers →** `demo/stories-spec.md` (the 3rd story),
`services/ai-readiness.md`, `seeding-spec.md`, `CLAUDE.md`; rext tag **`fit-up-m51`**.

**M52 — Single auditable seed+gen manifest** · `section` · depends: **M50 + M51** · ✅ **`done` — closed 2026-07-01**
(merged → `release/01.10b-fit-up`). Extract the Go mother-prompts to
YAML; author **one** checked-in `seed-generation-manifest.yaml` inlining population blueprint (all 3 orgs) +
prompts + batch config + snapshot sources (**cache + generated data excluded**); the cockpit **[Download]** serves
it. **Delivers →** NEW `corpus/ops/demo/seed-manifest-spec.md`.
All 4 sections landed (S1 `go:embed` extraction byte-identical → cache preserved; S2 the NEW `manifest` pkg + the
honesty-gated projection + `--manifest-export` verb; S3 cockpit [Download] repoint, non-fatal fallback; S4 the NEW
spec). Tests: rext stack-seeding **786** (+37 vs M51's 749; NEW `manifest` pkg 100% stmt) · demo-stack Python **313**
(+14: cockpit `--seed-manifest` endpoint + fallback); flake **0** (5/5 Go `-shuffle` + 5/5 Python). Close review: **12
findings all Fate-1** — F1 dedup the 3-way projection helper to one canonical `blueprint` source (removes the
projection-drift the honesty gate can't catch); F3 second cache-key golden fences the `{{else}}(none)` branch;
**F4 `mergeGenerationBatches` now WARNS on an orphan gen-story id** (a story-id typo was silently producing a
generation-less "auditable" manifest — the exact silent drop this milestone exists to prevent); F5 cockpit treats an
empty/blank `--seed-manifest` as absent; F2/F6/F7/F8/F9 (teeth-typo, gen-axis teeth, stale For-PMs prose, strip-helper
fence, vestigial doc-param). Deferral audit **GREEN** (up-injected.sh end-to-end glue = Fate-2 → M53's cold-rebuild;
0 repeats). **Carry-forward (three-fate, none escape-hatch):** up-injected.sh glue → M53 (Fate-2); consumption-clone
re-pin + `.agentspace/rext.tag` bump → M53 (push-gated KEEP). **Delivers →** `demo/seed-manifest-spec.md` (NEW) +
reconciled cross-refs (`cockpit-spec.md`, `ai-generation-spec.md`, `cache-spec.md`, `seeding-spec.md`, `README.md`,
`CLAUDE.md`); rext tag **`fit-up-m52`** (`36d7430`).

**M53 — Cold-rebuild acceptance** · `section` · depends: **M52** · ✅ **`done` — closed 2026-07-01**
(merged → `release/01.10b-fit-up`; **the FINAL v1.10b milestone**). Destroy the live demo + **rebuild from cold** on
a `stack-demo`-only box; assert healthy backends + complete set-dress/seed (all 3 orgs)/verify/cockpit + both-vantage
coverage + the AI-readiness criteria + the complete manifest download. Tag **`v1.10.1`** + bump `.agentspace/rext.tag`.
**All 6 sections landed** (§1 academy F6 seeder/wiring; §2 roll `v1.10.1`; §3 DESTROY via `/demo-down 1 --purge` — 17
containers + network + ALL demo-1 images, M49 #6 verified; §4 COLD REBUILD via a single `/demo-up 1`, no manual steps,
EXIT 0, no #7 abort; §5 ASSERT; §6 acceptance record + rext.tag bump). **Acceptance verdict: GREEN — 6/6 criteria +
academy F6 from cold** (AB1 backends healthy 17-up-0-exited/casbin 1150; AB2 prompt-free replay from the filled 1.4 GB
cache; AB3 set-dress+seed 3 orgs incl. Northwind AI-readiness + cockpit, EXIT 0; **AB4 both-vantage coverage GREEN**;
AB5 AI-readiness dashboard on Northwind — 50/100, 199 members, 3-step funnel, renders fast; AB6 cockpit [Download] =
complete inlined `seed-generation-manifest.yaml` 7593 B; **F6 academy** — real content + 9 cockpit [Academy] links →
authenticated member, Cosmo absent-by-design). **AB4 was RED on first assertion — an M51-owned gate regression
(M51 iter-05's unconditional ai-readiness manager `seedPath` broke the M50 base-org manager gate `dan-manager` @
Cervato); fixed at the acceptance gate** with user approval (an org-conditional manager manifest — `manifestFor(vantage,
expectedOrg)` returns the showcase `MANAGER_MANIFEST` only for Northwind, else `MANAGER_MANIFEST_BASE`; rext `117fe41`,
+3 unit tests; both manager vantages re-verified GREEN). Exactly the late cross-milestone regression M53 exists to catch
— the from-cold both-vantage assertion is the first joint re-measure of the M50 + M51 gates. Tests: rext stack-seeding
**791** (+5 vs M52's 786: F6 academy DeepLink/AcademyDeepLink build + harden single-source tests) · demo-stack Python
**326** (+13: F6 authenticated-session + [Academy] deep-link + harden `_academy_catalog_entry` edge/escape tests) · e2e
TS unit **29** (AB4 org-gating + referential-stability edges, +2 vs the pre-AB4 27); flake **0** (5/5 Go seeders shuffle
+ 5/5 Python cockpit+academy + 5/5 TS manifest). Close review: **2 findings, both Fate-1 docs** — DOC-1 documents the
AB4 org-conditional manager manifest in `coverage-protocol.md` (was undocumented); DOC-2 reconciles the stale
`~80%/≈160` AI-readiness prose to the shipped **78.4%/≈156-of-199** in `ai-readiness.md` (KB-2). Deferral audit **GREEN**
— every carry pointing at M53 LANDED here (up-injected.sh glue Fate-2 via the cold `/demo-up`; COLD acceptance = M53
itself; academy F6 = `e91f004`; box-level re-pin DONE); the historical academy-F6 REPEAT resolved by execution; 0
M53-originated deferrals, 0 escape-hatch. **Delivers →** the release acceptance record (`acceptance-record.md`, feeds
`/developer-kit:close-release`) + `.agentspace/rext.tag` = `v1.10.1`; rext release tag **`v1.10.1`** (`576dbcb` — rolls
up `fit-up-m47..m52` + F6 + AB4 + the M53 harden tests; re-rolled at close, local unpushed annotated re-roll). **Sole
residual = origin push (push-gated KEEP, orchestrator/user).** → **v1.10b is GREEN from cold; the release is complete →
`/developer-kit:close-release`.**

### Top risks

- **clone re-sync (M47) → RETIRED (was flagged the biggest unknown).** M47 found the clones already **current**
  (next-web @ v2.89.0, every repo ≤2 behind) — the re-sync was a trivial `make pull`, no 5-week catch-up, no
  cascading breakage. The snapshot recapture confirmed **both schema digests unchanged** (taxonomy `c75ce94…`,
  directus `ea2e187…`), so it was a clean in-place data refresh. The heavy-rebuild risk did not materialize.
- **content root-causes are hypotheses → degrades-quality.** M50 is iterative; it starts with a fresh observation
  pass on the clean bring-up, not the static guesses. Several "empty" surfaces may be demo-up #7 artifacts.
- **AI-readiness data model unknown → M51 exploratory.** The feature was invisible to the stale clones; M48
  documents it before M51 seeds it.
- **1-demo constraint → sequential chain, longer wall-clock.** No parallel verification; only rext authoring
  parallelizes (worktrees). M53 is the single from-cold acceptance truth.
- **M52 cache integrity.** Extracting prompts to YAML must preserve or deliberately re-key the M45 prompt-hash cache.
- **AI-provider keys → decision (M49/M50).** Which become throwaway/sandbox demo values vs documented-as-absent.

---

## Done — v2.0 "opening night" (SHIPPED 2026-07-02, tag `v2.0`)

> **Status (SHIPPED 2026-07-02, tag `v2.0`):** v2.0 "opening night" — the **Playthroughs** pillar — is **shipped**
> and merged to `main`. All four milestones closed: M201 "Manifest corpus" (`done`, closed-on-gate), **M202
> (Foundation, `section`, `done`** closed-complete 2026-07-01, tag `opening-night-m202`), **M203 (Employee-vantage
> coverage, `iterative`, `done`** closed-on-gate 2026-07-02 — 6/6 employee Playthroughs GREEN on cold reset-to-seed,
> 5/5 deterministic; tag `opening-night-m203`), and **M204 (Manager-vantage coverage, `iterative`, `done`**
> closed-on-gate 2026-07-02 — 4/4 manager Playthroughs GREEN on cold reset-to-seed, 5/5 deterministic; tag
> `opening-night-m204`) — the FINAL v2.0 milestone; it imported M203's shared page-object layer + ran on the
> reset-to-seed lifecycle per `corpus/ops/demo/playthroughs.md`. **Corpus at ship: 10 live Playthroughs (6 employee
> + 4 manager) GREEN on cold reset-to-seed, 1 declared in-manifest TODO** (the assign-WRITE half → Fate-2 → a future
> manager-write tier). Close-release: all 9 review sweeps GREEN, no blockers; rext code-of-record rolled to `v2.0`.
> Records archived under [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
>
> **Theme:** *the platform's core user journeys, proven to actually work.* A **Playthrough** is an automated
> actor that **is the user** — it logs in as a seeded hero, sets out with a goal, plays through a real journey
> across the platform start-to-finish the way a person would, then proves the platform delivered the outcome.
> The capability is the **canonical, living set of these journeys**: the platform's user-facing functionality,
> continuously **proven to actually work** — cleanly decoupled from *"the pixels are identical"* (a Playthrough
> breaks **only when a capability breaks**). It is the **functional** sibling of v1.x's M42 coverage sweep
> (which proves *presence* — every reachable page **shows** real content); Playthroughs prove the hero can **do**
> the things that world is for.
>
> **Designed 2026-06-28** via `/developer-kit:design-roadmap`, from the consolidated capability spec
> [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (v0.3). **A new MAJOR** — Playthroughs is
> a new pillar distinct from the demo/seeding lineage; v2+ uses **`Mxyy`** milestone numbering. **Tooling + docs
> only — zero platform-repo edits** (the read-only platform line carries over; an un-drivable surface escalates
> via the `unimplementable-without-platform-edit` state, it never edits the platform).

### Execution graph

```
v2.0 "opening night"
  M201 ──┐                          (manifest corpus — prose, user-guided)
  M202 ──┼──→ M203 ─┐
                M204 ─┴──→ ship
```

**M201 (the manifest corpus) and M202 (the foundation) open in parallel.** M201 is the **user-guided manifest
curation** — prose-only (the goal-aligned Product → Story → Use Case corpus), so it carries **no code dependency**
and can be authored before / alongside M202. M202 is the **Playthroughs foundation** (the section, the manifest
model + the §5.3 **validator**, the page-object layer, the dedicated seed + reset lifecycle, the runner + 4-state
reporting, one trivial proof Playthrough) — it builds the validator + dedicated seed to **match** the M201 corpus.
Then the two **vantage-coverage** milestones — **M203** (employee) and **M204** (manager) — run **in parallel**,
both `iterative`, implementing Playthroughs against the M201-declared use cases on the M202 foundation; the release
ships when both gates fire.

**M201 ∥ M202 (manifest ∥ foundation).** No hard ordering: M201 produces the **prose contract** (the use-case
manifest); M202 produces the **machinery** (validator + dedicated seed) that validates + seeds against it. They
reconcile when M202's validator runs over the M201 corpus. Where an M201 use case names a **precondition the demo
seed lacks**, that feeds the **M202 dedicated-seed expansion** (M201 records the need; M202 builds the seed).

**Parallelism note (M203 ∥ M204).** The two coverage milestones share an **additive merge surface**: the
per-surface **landmark registry** + the **locator index** (the §5.6 page-object layer every Playthrough imports).
Each vantage adds its own surfaces/anchors to that shared layer — an **additive** merge, not a conflicting one.
Both are `iterative` (the use-cases are *declarable* in the M201 corpus, but getting them green against the real
antd UI + the AI-sim assertion boundary is exploratory, like M42e/M42m), so they advance independently toward
their own exit gates and reconcile the registry additively at merge.

### Milestones

**M201 — Manifest corpus** · `iterative` · **USER-GUIDED** · complexity **large** · depends on: **none** (the
manifest is prose — authorable before/parallel to the M202 foundation).
**Status:** ✅ **`done` — closed-on-gate 2026-06-29.** 9 products · 26 stories · 28 use-cases authored,
**adversarially re-grounded** (11-agent wf `wvpnpvozh` → 15/27 runnable), **user-signed-off**. Records:
[`releases/archive/02.00-opening-night/m201-manifest-corpus/`](releases/archive/02.00-opening-night/m201-manifest-corpus/)
(deliverable: `manifest-draft.yaml`). The close discovered the **stale-clone drift** (next-web 115+ commits behind
prod) → **v2.0 PAUSED for the v1.10b "fit-up" backfill** (re-sync + re-ground + re-validate + fix, user-driven; see
the In-Development section above) before resuming. *(Correction: the v1.10b M47 milestone later found the clones
actually current — next-web @ v2.89.0; the stale surface is the corpus, re-grounded by M48. The pause stands; the
backfill's real work is the corpus + the bring-up/content fixes, not a clone re-sync.)*
**Goal:** top-down, **user-directed** curation of the **full goal-aligned Product → Story → Use Case manifest
corpus** — the build + regression contract every coverage milestone (M203/M204) implements against. The flow per
pass: **outline** (products → stories → use cases) → **validate** (against the real platform surface + the spec's
manifest model) → **write the prose-intent manifest YAML** (spec §5.3, **one file per product**).
**Explicitly NOT bounded by the current minimal/partially-aligned demo stories seed** — it captures what the goal
says must be proven; where a use case needs preconditions the demo lacks, that **feeds the M202 dedicated-seed
expansion** (noted, not resolved here).
**Shape:** `iterative`, **driven by the user** — worked conversationally + via `/developer-kit:work-mstone-iters`,
not autonomously.
**Exit gate:** **the manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML (one
file per product)** — covering the platform's products × their must-work user journeys, each use case carrying
**goal + actor + flow + intermediate/final expectations**, structurally valid (the spec §5.3 validator passes,
**ids unique + both-way**) — **and the USER signs off the corpus as the complete-enough v2.0 coverage contract.**
**iteration_protocol_ref:** the capability spec
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md) (esp. §2 model, §4 use-case shape, §5.3
manifest format).
**Delivers →** the prose-intent manifest YAML corpus (one file per product); **lands in the rext `playthroughs`
section once M202 exists**, until then drafted under the milestone dir / `spec-drafts/playthroughs/manifest-draft/`.
**Candidate starting outline (the user directs — NOT fixed):** (a) the demo-covered products — **Skill Paths, AI
Simulations, Profile & Skills, Workforce Intelligence, Hiring, Academy**; (b) goal-aligned areas the demo barely
covers (flag *to confirm with the user*) — **Auth & Onboarding, Billing & Entitlements/tier-gates, Org Admin &
Settings, cross-product journeys** (candidate→employee).

**M202 — Playthroughs Foundation** · `section` · complexity **medium** · depends on: **none** (reuses the M42
harness + the seeding machinery; the M201 manifest corpus is its build+regression contract, authorable in parallel).
**Status:** ✅ **`done` — closed-complete 2026-07-01.** All 6 sections + the NEW `corpus/ops/demo/playthroughs.md`
runbook delivered; the trivial proof Playthrough (login → /profile → assert hero identity) **GREEN on demo-1**. The
`playthroughs` rext section: manifest model + light validator (both-way id integrity + precondition-coverage +
datadna closure gate) · per-surface page-object layer (1 surface: `/profile`, re-pin O(surfaces)) · dedicated
decoupled `pt-world` seed (2 private orgs, entitlement tiers + multi-org-private) · reset-to-seed lifecycle +
serial-default runner · 4-state reporting map. **96 Go test/fuzz funcs (98.5% section) + 13 TS** (5/5 flake-clean).
Close surfaced 8 findings, all Fate-1: CQ1 datadna exit-3 diagnosis · CQ2 `PW_WORKERS` serial-safety guard · CQ3
`truncate` totality · CQ4 `ptTagRe` lockstep · DOC1 section-index · DOC2 fixtures wording · M202-D4 anchor-story
landmine blended into `stories-spec.md`. Deferral audit **GREEN** (0 milestone-owned). Tooling + docs only — **zero
platform edits, zero new deps**. rext authoring @ `b1e5528`, tagged **`opening-night-m202`**. The runbook IS the
M203/M204 `iteration_protocol_ref`. Records: [`releases/archive/02.00-opening-night/m202-foundation/`](releases/archive/02.00-opening-night/m202-foundation/).
**Goal:** stand up the **`playthroughs` rext section** on the **shared M42 e2e foundation**, proven by **one
trivial end-to-end Playthrough**.
**Scope — In:**
- the **manifest model + a light validator** — both-way id integrity (use-case ↔ Playthrough, traceable both
  directions) + precondition-coverage (every declared `seed`/`preconditions` resolves to a named seeded world,
  no silent "ideally"), **datadna-gated** (the Playthrough seed is covered by the same `datadna` conformance gate
  as the seeding machinery);
- the **per-surface locator/landmark page-object layer** (the §5.6 shared registry every Playthrough imports —
  a UI/antd/copy shift is absorbed by editing the per-surface registry, not N tests) — **1 surface to start**;
- the **dedicated, decoupled seed** preset (test data ≠ demo data; the demo seed is the *starting point* but
  kept separate) — **spans entitlement tiers + multi-org-private**;
- the **reset-to-seed lifecycle + serial-default runner** (`workers: 1`, `fullyParallel: false`; reset via the
  real `--reset` path honoring its contract — **additive re-seed is FORBIDDEN as a reset**);
- the **4-state reporting map** — `passing` / `failing` / `unimplemented` / `unimplementable-without-platform-edit`
  (the last being the P3 zero-edit escape valve — it escalates, never edits);
- **one trivial proof Playthrough** — **login → /profile → assert hero identity** (the foundation's smoke test).
**Out:** real product coverage (M203+); the AI-sim / integration mirror tier; cross-vantage.
**Delivers →** a corpus runbook that **graduates the playthroughs spec** (e.g.
[`corpus/ops/demo/playthroughs.md`](../../corpus/ops/demo/playthroughs.md)) — becomes the `iteration_protocol_ref`
for M203/M204.
**KB deps (read as contract):** the playthroughs spec-draft
[`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), the **M201 manifest corpus**
([`releases/archive/02.00-opening-night/m201-manifest-corpus/`](releases/archive/02.00-opening-night/m201-manifest-corpus/) — the
prose contract the validator + seed implement against),
[`corpus/ops/demo/coverage-protocol.md`](../../corpus/ops/demo/coverage-protocol.md),
[`corpus/ops/seeding-spec.md`](../../corpus/ops/seeding-spec.md),
[`corpus/ops/idempotency.md`](../../corpus/ops/idempotency.md).
**Reuse paths (cite in spec-notes):** `stack-demo/rosetta-extensions/stack-verify/e2e/lib/{cockpit-login,
section-assert,empty-states,coverage-manifest}.ts`, `stack-demo/rosetta-extensions/stack-seeding/`.

**M203 — Employee-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M204**. ✅ **`done` — closed-on-gate 2026-07-02** (merged → `release/02.00-opening-night`; rext
tag **`opening-night-m203`** @ `fb94458`). **Exit gate MET at iter-06:** every declared employee-vantage use case
has a PASSING Playthrough on a COLD reset-to-seed demo (**6/6** — Profile identity+verified+growth+timeline · Skill
Paths legacy learn&progress · AI Simulations chat launch §5.8), with **0 false-fails over 5 consecutive reset runs**
(5/5). **6 iters** (1 bootstrap tok + 5 tiks, all closed-fixed). The strategy arc (TOK-01, deterministic-read-first):
iter-02/03 landed the full Profile journey (Spotlight + claimed-vs-verified gap + trajectory + work/education
timeline), iter-04 Skill Paths (browse→open→start→progress; verify-skill composes P7 on the profile side), iter-05
AI Sims chat launch + the **post-seed Sentinel-Reload** fix (a casbin g3 feature grant is only effective after the
enforcer RELOADS — folded into `run-playthroughs.sh`, drift-guarded), iter-06 the 5-run determinism gate (runnable
from the authoring-built `stackseed`). Tests: rext playthroughs Go **103** (+7 vs M202's 96: @pt-tag lockstep ×2
packages + invalid-engine + read-error arms + the iter-05 Sentinel-reload drift guard) · e2e TS **38** unit (+25
vs 13) + **6** browser Playthroughs; flake **0** (Go 5/5 -shuffle + TS unit 5/5; browser 5/5 cold-reset iter-06).
Close review: **11 findings all Fate-1** — F1/F5 (must-fix) the segment-anchor route-shape fix was NOT applied to
two inline `\b`-terminal copies (a load-bearing `waitForURL`) → centralized into `url-shapes.ts` `SKILLS_TAB_URL`;
F3 dropped 4 dead speculative self-eval accessors; TEST-G1 the unenforced `@pt:`-grammar lockstep → twin tests;
TEST-G3/G4 two coverage arms → 100%; DOC README + runbook M203 backfill. Deferral audit **GREEN** (0
repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):** the 4 **non-gate** edge UCs
(`ai-simulations.code.UC1` Judge0 · `ai-simulations.interview.UC1` text · the Skill-Paths verify-skill terminal ·
`profile.self-evaluation.UC1`) → **Fate-3 → M206** (roadmap-vision annotated; the gate enumerated the 3 CORE
journeys, all GREEN — these are additional). Academy UC OUT by design; voice sims → M206 by design. Tooling + docs
only — zero platform edits, zero new deps. Records:
[`releases/archive/02.00-opening-night/m203-employee-coverage/`](releases/archive/02.00-opening-night/m203-employee-coverage/).
**Goal (as designed):** **Maya's** core **employee** journeys play green (declared in the M201 manifest corpus) —
Skill Paths (browse → enroll → complete → verify-skill), AI Simulations (chat/code launch → complete →
score-in-range, **NON-voice**), Profile (verified-skill chart + the claimed-vs-verified gap + work/education
timeline). **Why iterative:** the use-cases are *declarable* (in the M201 corpus), but getting them green against
the real antd UI (the landmark layer) + the AI-sim assertion boundary is **exploratory**, like M42e.

**M204 — Manager-vantage coverage** · `iterative` · complexity **large** · depends on: **M202** ·
parallel-with: **M203**. ✅ **`done` — closed-on-gate 2026-07-02** (merged → `release/02.00-opening-night`; rext
tag **`opening-night-m204`** @ `c81c6dd`). **THE FINAL v2.0 milestone.** **Exit gate MET at iter-05:** every
declared manager-vantage use case has a PASSING Playthrough on a COLD reset-to-seed demo (**4/4** — Workforce
funnel + org-scale gap · member roster · per-member activity-dashboard drill-down · succession/at-risk), with **0
false-fails over 5 consecutive reset runs** (5/5). **5 iters** (1 bootstrap tok + 4 tiks, all closed-fixed). The
strategy arc (TOK-01, manager-surface-per-iter): iter-02 the Workforce funnel + roster (the manager landing
surface) + the runner **reporter-override** fix (a stale-JSON gate hazard, drift-guarded), iter-03 the
activity-dashboard per-member drill-down (a SPA-URL race + an out-of-`<main>` table scope fixed), iter-04
succession/at-risk (all 3 declared journeys green — no seed expansion needed, Org A size-40 rendered the M36
aggregates), iter-05 the 5-run determinism gate. All 4 manager UCs are READ/monitoring flows → the risk was
seed-scale render + antd grid ambiguity, not mutation-determinism. Tests: rext playthroughs Go **105** (+2 vs
M203's 103: the reporter-override drift guard + the manager-coverage deliverable-presence pin) · e2e TS **58**
unit (+20 vs 38) + **4** browser Playthroughs; flake **0** (Go 5/5 -shuffle + TS unit 5/5; browser 5/5 cold-reset
iter-05). Close review: **5 findings all Fate-1** — 1 code-quality should-fix (documented the manager
predicate-API's deliberate symmetric shape — kept, not pruned, consistent with the M203 `isOnSkillsTab`
precedent), 3 docs (flip the stale "M204 adds" future-voice → landed; corpus now **10 live Playthroughs, 1 TODO**;
add the M204 manager page-object bullet), 1 decision-triage (D-CLOSE-1). Deferral audit **GREEN** (0
repeat/chronic/aged). **Carry-forward (three-fate, none escape-hatch):** `assignment-monitoring.assign-and-track.UC1`
(the assign-WRITE half — a two-backend org-admin WRITE flow) → **Fate-2**, tracked in-manifest as `playthrough:
TODO` (reports `unimplemented`, presence-pinned; out of M204's declared 3 monitoring journeys; a future
manager-write tier is its home). Tooling + docs only — zero platform edits, zero new deps. Records:
[`releases/archive/02.00-opening-night/m204-manager-coverage/`](releases/archive/02.00-opening-night/m204-manager-coverage/).

**Goal:** **Dan's** core **manager** journeys play green (declared in the M201 manifest corpus) —
- **Workforce funnel** + member roster,
- **member drill-down** (the activity-dashboard),
- **succession / at-risk** (the Growth tab) signals.
**Exit gate:** **same shape as M203, manager-vantage** — every declared manager-vantage use case has a passing
Playthrough on a COLD reset-to-seed demo stack, with 0 false-fails over 5 consecutive reset runs.
**iteration_protocol_ref:** same as M203 (the spec / the M202-delivered runbook).
**Why iterative:** same as M203 — declarable use-cases, exploratory path against the real manager UI + the
assertion boundary.
**Re-scope trigger:** same — `unimplementable-without-platform-edit` → escalate, don't edit.

### Top risks

- **manifest completeness → no auto-gate, user owns "enough".** The M201 manifest is a **build reference** with
  **no introspectable schema for "what user-facing capabilities exist"** (spec §5.9) — an *added* platform
  capability with no use case cannot be auto-detected. The corpus's completeness is a **user judgement** (the M201
  exit gate's sign-off), not a machine check. *Mitigation:* M201 is **user-guided + iterative** (the user directs
  each top-down pass + signs off the complete-enough contract); the cadence-review stance (§5.9) carries forward.
- **antd-a11y → the landmark layer is load-bearing.** zero-platform-edit means we **cannot** add `data-testid`;
  locators bind to the **accessibility tree** (`getByRole`/`getByLabel`/`getByText`) over the **real antd UI**,
  with a Rosetta-side **landmark registry** for ambiguous surfaces. If antd's a11y is thin on a surface, that
  surface's landmark anchors carry the test — the registry's quality is the risk. *Mitigation:* the per-surface
  page-object layer (re-pin is O(surfaces), not O(tests)); start with **1 surface** in M202 to prove the pattern.
- **determinism-under-mutation → M202's reset must be solid.** P6 ("same inputs → same result") holds **only if**
  the world is reset to the known seed between runs, and an *additive* re-seed silently leaves stale state (the
  M42e "green-but-wrong" trap). The whole determinism headline rests on M202's **reset-to-seed lifecycle** being
  correct — it is a **foundation** risk, surfaced and owned in M202 before M203/M204 lean on it.
- **hero-login → demo-N only.** Hero-driven Playthroughs run on **demo-N** (or a Clerkenstein-injected dev-N) —
  a plain dev-N is real Clerk + one identity + `dev-min`, so the hero suite is **not** the same on dev today. The
  target environment is the demo stack; the dev-stack hero-flow generalization is an explicit **later** item
  (spec §5.4), not v2.0 scope.
- **AI-sim mirror tier is future.** The signature voice/recording AI-simulation journey needs a **mirror engine**
  (Clerkenstein mocks **only** Clerk — no LiveKit/Chime/Stripe/Brevo mirror). v2.0 covers the **NON-voice**
  chat/code/document sims (playable as-is, asserted at the launch/completion boundary); voice + recording +
  payments + email are parked as `later — needs a mirror engine` → **M206** ([`roadmap-vision.md`](roadmap-vision.md)).

---

## Shipped releases

- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`, **4 milestones (M201 … M204)**. The **Playthroughs**
  pillar: a manifest-driven, deterministic e2e suite that plays real user journeys and proves the platform delivers
  the outcome (**function**, vs the M42 coverage sweep's **presence**). Manifest corpus → foundation → employee +
  manager coverage. **10 live Playthroughs GREEN on cold reset-to-seed, 1 in-manifest TODO.** The **first v2.x
  release**. Tooling + docs only, zero platform edits, zero new deps. Records archived under
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10.1 "fit-up"** — **2026-07-01**, tag `v1.10.1`, **7 milestones (M47 … M53)**. The interposed
  **field-hardening backfill** (a `.1` patch on v1.10): re-sync + recapture, corpus re-ground, from-cold `/demo-up`
  hardening, content + AI-readiness-showcase-org seeding fill, one auditable seed+gen manifest, then a from-cold
  destroy-and-rebuild acceptance (**6/6 + academy F6 GREEN**). Tooling + docs only, zero platform edits, zero new
  deps. Records archived under
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

The complete earlier shipped history — **v1.0 "body double"** (2026-06-03, tag `v1.0`) through **v1.10 "method acting"**
(2026-06-27, tag `v1.10`), 11 versions / milestones M0 … M46 — is preserved in
[`roadmap-legacy.md`](roadmap-legacy.md) (the retired v1.x major). Records are archived under
[`releases/archive/`](releases/archive/). **v2.0 "opening night" (tag `v2.0`) is the first shipped v2.x release**
(2026-07-02); the next v2.x release awaits `/developer-kit:design-roadmap`.

## Notes

- **Milestone numbering — v2+ uses `Mxyy`** (`M` + major digit + two-digit milestone): **M201, M202, M203, M204**
  for v2.0. This is the major-version scheme `context.md` reserved for *"a future *major* v2+"*; the v1.x flat
  sequential counter (M0 … M46, with the `a`/`b`/`c`/`e`/`m` suffix conventions) lives in
  [`roadmap-legacy.md`](roadmap-legacy.md) § Notes. **It was thought closed at M46, but the interposed v1.10b
  "fit-up" backfill RE-OPENS it at M47** (M47→M53) — the backfill is v1.x-major work (a `.1` patch of v1.10), not a
  v2 milestone, so it keeps the flat counter rather than `Mxyy`.
- **Milestone shapes** mix within v2.0: **M201 is `iterative` + USER-GUIDED** (the manifest corpus — a top-down,
  user-directed prose curation toward a sign-off gate); **M202 is `section`** (a fixed In-scope checklist — the
  foundation is decomposable up front); **M203 + M204 are `iterative`** (a measurable exit gate, exploratory path
  — getting declared use-cases green against the real antd UI + the AI-sim assertion boundary, like the M42e/M42m
  precedent).
- Date format throughout: ISO `YYYY-MM-DD`.
- The Playthroughs capability **graduated from spec-draft to active development** at v2.0 design (2026-06-28); the
  governing spec is [`spec-drafts/playthroughs/spec.md`](spec-drafts/playthroughs/spec.md), graduated to a corpus
  runbook (`corpus/ops/demo/playthroughs.md`) by M202.

_Last updated: 2026-06-29 (**v1.10b "fit-up" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — an
interposed **field-hardening backfill** [the v1.3b "dress rehearsal" lineage]; **7 milestones M47 → { M48 ∥ M49 } →
M50 → M51 → M52 → M53** re-opening the v1.x flat counter; branch `release/01.10b-fit-up` cut from `main`; tag
`v1.10.1`. Designed from the field review `.agentspace/annotation.md` + the M201 stale-clone finding [3 research
agents]. Re-grounds demo + corpus to current prod, fixes the from-scratch `/demo-up` issues + the v1.10 content
gaps, adds the AI-readiness showcase org, consolidates one auditable seed+gen manifest. **v2.0 "opening night"
PAUSED** until it ships. Tooling + docs only — zero platform-repo edits. Prior: 2026-06-28 **v2.0 "opening night"
DESIGNED + PROMOTED** — a NEW MAJOR opening the **Playthroughs** pillar; 4 milestones M201 ∥ M202 → { M203 ∥ M204 };
branch `release/02.00-opening-night`; from `spec-drafts/playthroughs/spec.md` v0.3.)_
