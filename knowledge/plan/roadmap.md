# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills). This file holds the **active major** only; the retired **v1.x** history (M0 … M46, all
SHIPPED) lives in [`roadmap-legacy.md`](roadmap-legacy.md). Future versions + the unscheduled backlog live in
[`roadmap-vision.md`](roadmap-vision.md). The live source of truth for *current/next* is [`state.md`](state.md).

> **Designed 2026-07-19** via `/developer-kit:design-roadmap`. **v2.5 "the playbill"** is the **content-vantage
> release** — two threads on the same mature demo/cockpit machinery, shipped together. **THREAD A** finally fills the
> **empty ant-academy** grid: it renders 0 cards on a demo not because of a platform bug but because the catalog is
> **DB-authoritative** (queries the academy subgraph over GraphQL) and a demo neither sets the endpoint nor holds
> academy rows → `emptyCatalogView()`; the corpus even mis-documented this (`ant-academy.md` still claims *"Backend
> services: None / static JSON"*), which mis-routed the year-old **F4** carry into the wrong repo. **THREAD B** adds a
> **2nd "Content stories" cockpit tab** beside "Org stories": sections per **content product** (Simulation · Skill-path
> legacy · Skill-path new [ant-academy] · AI-labs), each a list of **played sessions** with two login-and-land actions —
> **as-player** and **as-manager** (where a manager view exists). Sessions are **cloned from real production sessions,
> anonymized where possible, non-manager-played, re-tenanted into an existing manifest org, and source-pinned by prod
> session-id** (deterministic — no random-per-reseed). **User decisions (2026-07-19):** the **real customer-session**
> sourcing is accepted as the user's data-controller call — demos stay **VPN/tailnet-scoped** (not open-internet), and
> the release **amends `safety.md` Part 3** to document the new posture honestly (anonymized-real, VPN-bounded — no
> longer a blanket "nothing behind the door"); academy fill is **production-faithful** (no "Draft" chip); AI-labs + the
> academy content-product section are **scoped by the M231 spike** (likely presence-only / deep-link, given labs' nil
> client + academy's absent server session store). **8 milestones M229 → M230 → M231 (HARD go/no-go) → M232 → M233 →
> M234 → M235 → M236**, spike-first; tag **`v2.5`**; branch `release/02.50-the-playbill`. **Tooling + docs only — zero
> platform-repo edits** (a runtime-computed result page that won't render from a seeded row routes to a sha-pinned
> demo-patch or escalates). Continues the v2.x `M2xx` scheme at M229.
>
> **Designed 2026-07-15** via `/developer-kit:design-roadmap`. **v2.4 "casting call"** is the
> **recruiter-vantage / hiring-org release** — a **NET-NEW** release that adds a **4th, HIRING demo org** to the
> presenter cockpit, where **45 candidates audition on the same 5 positions and a manager compares them side by
> side**, distinct from the three workforce orgs. **This release formally REVERSES v2.3's `D-DESIGN-4`**
> (*"there is no hiring org and none will be built"*): the stated blocker — *"a hiring story would need the
> `hiring-app` frontend, which is not in the demo UI tier"* — was **REFUTED by research**: the candidate-comparison
> surface ships inside the **dockerized `apps/web` (Workforce)** app the demo already builds, and the domain
> primitives (`organizations.is_hiring`, the `candidate` membership role, `jobsimulation.sessions` typed
> `SIMULATION_TYPE_HIRING`) **already exist in the platform**. It is **not a clean section release**: two blind
> areas (the hiring read-model + proof-by-rendering that the comparison surface is demo-servable) gate it, so it
> **opens with an investigation-first spike (M222) that is a HARD go/no-go barrier**. **5 milestones M222 → M223 →
> M224 → M225 → M226** (largely sequential); tag **`v2.4`**; branch `release/02.40-casting-call`. **Tooling + docs
> only — zero platform-repo edits** (a platform-source render gate routes to a sanctioned sha-pinned demo-patch,
> never a repo edit; an un-patchable surface **escalates**). **Consumes the recruiter/seeder half of the reserved
> vision M205**, leaving M205's Stripe-tier-gate + ATS-pipeline half a residual vision reservation. User decisions
> baked in: a **genuine hiring org** (`is_hiring=true` end-to-end, Clerkenstein `isHiring` wiring in scope) · **real
> replayed positions + a realistic non-degenerate funnel** (not a flat 225-session grid) · cockpit heroes = **1
> manager + 2 candidates** (one assessed, one only-assigned), login-only.
>
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
> release. **v2.1 "quick change"** (the skiller-in-app re-ground) followed (designed 2026-07-08) and
> **SHIPPED 2026-07-09 (tag `v2.1`)**. The pre-v2 v1.x history (M0 … M46) lives in `roadmap-legacy.md`.
>
> _(This blockquote preserves the DESIGN-TIME narrative of the v2.0 opening. Its status verbs were
> re-pointed to shipped truth at the v2.5 close — `D-18` — so a reader cannot mistake a design-time
> "IN DEVELOPMENT" for live status. **Live status is always [`state.md`](state.md).**)_

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.10b** | **fit-up** | Field-hardening backfill — re-ground demo + corpus to current prod, fix the from-scratch `/demo-up` issues + the v1.10 content gaps, add the **AI-readiness showcase org**, and consolidate **one auditable seed+gen manifest** | M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53 | ✅ **SHIPPED 2026-07-01 (tag `v1.10.1`)** (branch `release/01.10b-fit-up`, designed 2026-06-29; all 7 milestones done) |
| **v2.0** | **opening night** | The platform's core user journeys, **proven to actually work** — a new **Playthroughs** pillar: a manifest-driven, deterministic e2e suite that *pretends to be the human* and proves the platform does its job | M201 ✅ ∥ M202 ✅ → { M203 ✅ ∥ M204 ✅ } → ✅ ship | ✅ **SHIPPED 2026-07-02 (tag `v2.0`)** (branch `release/02.00-opening-night`, designed 2026-06-28; all 4 milestones closed-on-gate/complete). **10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed, 1 in-manifest TODO.** The first v2.x release. Tooling + docs only, zero platform edits, zero new deps |
| **v2.1** | **quick change** | The **skiller-in-app re-ground** — re-fit the tooling, corpus, and stacks to the merged platform (skiller service + schema folded into `app`/`public`, RPC → `backend`, **4 subgraphs**) and **prove** `dev-up` + `demo-up` still work. Field-hardening lineage (v1.3b/v1.10b), triggered by a landed platform change | M208 → M209 → M210 → M211 (strictly sequential) | ✅ **SHIPPED 2026-07-09 (tag `v2.1`)** (branch `release/02.10-quick-change`, designed 2026-07-08; all 4 milestones done — the merged platform stands up **cold on both stacks**; M42 coverage both vantages + v2.0 Playthroughs 10/11 GREEN; tooling + docs only, zero platform edits, 0 net-new deps) |
| **v2.2** | **panorama** | The **external-shareability release** — make dev/demo stacks reachable over a **Tailscale** tailnet (run on a Tailscale VM; a teammate browses the demo end-to-end over its MagicDNS name), via a single opt-in host knob + the tailscale-cert HTTPS surface. The re-proposal of the dropped v1.4 Tailscale/ingress seed | M212 ✅ → { M213 ✅ ∥ M214 ✅ } → M215 ✅ (+ opt M216) | ✅ **SHIPPED 2026-07-12 (tag `v2.2`)** (branch `release/02.20-panorama`, designed 2026-07-11; all 4 core milestones done — opt-in default-off, HTTPS-everywhere, demo-first; tooling + docs only, zero platform edits, 0 net-new deps; rext code-of-record `v2.2` = `39e8013`). **M215 proved it live:** the first remote Linux-VM demo over Tailscale, both vantages green from a 2nd machine on a trusted cert, reproducibly on a cold reset-to-seed |

| **v2.3** | **cue to cue** | The **presenter-speed release** — a presenter swaps heroes in **under 5 seconds** on a demo that comes up **green, fully-loaded, and remotely reachable by default**. Field-hardening lineage, triggered by a live 1–2-minute cockpit-login defect whose root causes were **already measured in-repo** and **silently rotting** (two dead perf demo-patches, a refusal piped to `/dev/null`, a 4-place false latency claim in the corpus) | M217 → { M218 ∥ M219 ∥ M220 } → M221 | ✅ **SHIPPED 2026-07-15 (tag `v2.3`)** (branch `release/02.30-cue-to-cue`, designed 2026-07-13; all 5 milestones done — the headline **click→ACCESS < 5 s** gate set at M218 and **re-proven live 8/8 on `billion`** over the tailnet with no flags: 2.11 s / 1.31 s vs a ~39/38 s baseline; 3 orgs, AI-readiness filled, remote default-on; tooling + docs only, zero platform edits, 0 net-new direct deps). rext code-of-record `cue-to-cue-m221-final`; the `billion` demo LEFT LIVE |
| **v2.4** | **casting call** | The **recruiter-vantage / hiring-org release** — a **NET-NEW** 4th, **HIRING** demo org on the presenter cockpit where **45 candidates audition on the same 5 positions and a manager compares them side by side**, distinct from the three workforce orgs. Reverses v2.3's D-DESIGN-4 (the comparison surface ships in the dockerized `apps/web`, not the Vercel-only `apps/hiring`); consumes the recruiter/seeder half of the reserved vision M205 | M222 → M223 → M224 → M225 → M226 → **M227 → M228** (RE-OPENED for believability) | ✅ **SHIPPED 2026-07-18 (tag `v2.4`)** (branch `release/02.40-casting-call`, designed 2026-07-15; **RE-OPENED 2026-07-17** for believability fixes from live feedback). M222 spike [`section`, GO] → M223 seeder [`section`] → M224 render [`iterative`, closed-on-gate] → M225 demo-integration [`section`] → **M226 prove-on-billion [`iterative`, closed-on-gate 2026-07-17]** (the 7-condition hiring gate proven live on `billion`, default `/demo-up`, recruiter p95 < 5 s as the 3rd vantage) → **M227 the-notes [`section`, closed 2026-07-17]** (4 believability seed/content fixes deterministically proven + gate retuned `≥40→≥6`; live re-prove → M228). **M228 second-night [`iterative`, closed-on-gate 2026-07-18]** = the corrected-demo billion re-prove, 7/7 live (render 5/5, heroes 3/3, recruiter p95 1.27 s). Tooling + docs only, **zero platform-repo edits** — merged to `main`; the corrected hiring demo proven live 7/7 on `billion` (recruiter p95 click→ACCESS 1.27 s) |
| **v2.5** | **the playbill** | The **content-vantage release** — TWO threads on the same demo/cockpit machinery. **A:** fill the empty **ant-academy** grid (DB-authoritative catalog; production-faithful demo-fill; corrects the false `ant-academy.md`). **B:** a 2nd **"Content stories"** cockpit tab — sections per content product (Simulation · Skill-path legacy · Skill-path new · AI-labs), each a list of **played sessions** with **as-player / as-manager** login-and-land actions; sessions **cloned from anonymized real prod sessions, non-manager-played, re-tenanted, source-pinned by prod session-id** | M229 → M230 → **M231 (HARD go/no-go)** → M232 → M233 → M234 → M235 → M236 | ✅ **SHIPPED 2026-07-20 (tag `v2.5`)** (branch `release/02.50-the-playbill`, designed 2026-07-19; all 8 milestones M229–M236 closed + merged). Spike-first; one combined release. **29/29** landable (session × action) pairs live on `billion` both vantages + academy grid filled (65 cards); real-customer-session sourcing accepted (data-controller call); demos **VPN/tailnet-scoped**; **amends `safety.md` Part 3** (anonymized-real, VPN-bounded). Tooling + docs only, **zero platform-repo edits**. ⚠️ 29/29 is unit-proven, not live-re-proven — live re-prove is v2.6/M237 |
| **v2.6** | **sound check** | The **reliability / field-hardening release** — *make everything that's built actually get built + provisioned.* Triggered by live demo defects ("still not all gets built and provisioned as expected"). Barrier → parallel fixes → prove-on-billion: fix clone-freshness + re-triage the ambiguous UI defects on a fresh build, then fix academy reliability · enterprise surfaces (talk-to-data live via real AWS Bedrock creds) · content-stories fidelity (media ported) · language toggle · cockpit UX · the assign-WRITE Playthrough — then re-prove the whole feature (v2.5's headline **and** every v2.6 fix) live on `billion`, cold reset-to-seed | M237 (HARD go/no-go) → { M238 ∥ M239 ∥ M240 → M241 → M242 ∥ M243 } → M244 | 🟡 **IN DEVELOPMENT** (branch `release/02.60-sound-check`, designed 2026-07-20; 8 milestones M237 → M244; **M237 barrier + M238 academy reliability + M239 enterprise surfaces + M240 content-stories fidelity + M241 content-stories language + M242 cockpit UX + M243 assign-WRITE Playthrough CLOSED** (M237–M239 2026-07-21, M240–M243 2026-07-22) + **M244 prove-on-billion CLOSED 2026-07-23 (gate MET 8/8 live on `billion`, closed-on-gate, 0 platform edits)** → **all 8 M237→M244 CLOSED, ready for `/developer-kit:close-release`**; realizes the reserved M237/M238; tooling + docs only, **zero platform-repo edits**; tag will be `v2.6`) |

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

## Active — v2.6 "sound check" (IN DEVELOPMENT, designed 2026-07-20, tag will be v2.6)

> **Designed 2026-07-20** via `/developer-kit:design-roadmap`. **v2.6 "sound check"** is a **reliability /
> field-hardening release** (the v1.3b "dress rehearsal" / v1.10b "fit-up" / v2.1 "quick change" / v2.3 "cue to cue"
> lineage), triggered by **live demo defects** — *"still not all gets built and provisioned as expected."* A sound check
> is the pre-show pass where every input is proven to actually reach the desk before the audience arrives: the release's
> whole job is to make everything that is *built* actually *build and provision* on a fresh box. The house shape is
> **barrier → parallel fixes → prove-on-billion** (v2.3/v1.10b): a **HARD go/no-go barrier (M237)** first fixes
> **clone-freshness** (the demo was building from **stale platform source** — only defect #1 was clone-staleness; #2–#5
> each reproduce on `origin/main`) and re-triages the ambiguous UI defects on a *correct* build, so every downstream fix
> is scoped against reality; then a **parallel fix fan-out (M238–M243)**; then an **iterative closer (M244)** re-proves
> the whole feature — v2.5's headline `29/29` **and** every v2.6 fix — live on `billion`, cold reset-to-seed (this
> realizes the reserved `M237` re-prove that v2.5 shipped un-live-proven, and the reserved `M238` assign-WRITE). **3
> binding user decisions (2026-07-20):** **(1) talk-to-data → FULL** — real AWS Bedrock creds provisioned via
> `/stack-secrets` + a secret-coverage DNA extension for `app` (not just a flag); **(2) media → PORT IT** — capture +
> re-host the Chime/S3 voice recording + document blobs, behind a **HARD internal PII gate** (fresh data-controller
> sign-off + a `safety.md` raw-media amendment + a voice/document anonymization decision — a voice cannot be
> token-scrubbed); **(3) language → EN-only fallback per tuple** — M241 opens with a read-only pool-count go/no-go query,
> IT where it exists, EN-only where absent. **8 milestones M237 → M244**; tag **`v2.6`**; branch
> `release/02.60-sound-check`. **Tooling + docs only — zero platform-repo edits** (a dead platform surface routes to a
> sha-pinned `demopatch` or **escalates**, never a repo edit). Continues the v2.x `M2xx` scheme at **M237**.

**Theme.** *Make everything that's built actually get built + provisioned.* A field-hardening release on the mature
demo/cockpit/content-stories machinery: the v2.5 features are real, but a fresh bring-up doesn't reliably *build* + *provision*
all of them. Fix the build-freshness barrier, fan out the confirmed defects in parallel, then re-prove the whole thing live.

**User decisions baked in (2026-07-20):**
1. **Talk-to-data → FULL** — wire **real AWS Bedrock creds** via the `/stack-secrets` provisioning mechanism (key set
   `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_SESSION_TOKEN`/`AWS_REGION` + `CLAUDE_CODE_USE_BEDROCK`, referencing
   `../hyper-studio/.env.example`) + extend the secret-coverage DNA for the `app` service. Not just a flag. (M239)
2. **Media → PORT IT** — capture + re-host the Chime/S3 voice recording + document blobs so the manager can hear the
   call / see the document. This **expands the customer-PII surface to raw audio + full documents** → M240 carries a
   **HARD internal gate**: fresh data-controller sign-off + a `safety.md` amendment for raw media + a voice/document
   anonymization decision **before any customer audio lands in a demo**. Likely consumes **DEF-M10-01** (S3 read). (M240)
3. **Language → EN-only fallback per tuple** — M241 opens with a **read-only prod pool-count query** (IT sessions per
   requirement tuple); toggle where IT exists, EN-only where absent. No blocking. (M241)

**Hard constraint (carried, unchanged):** **zero platform-repo edits** — a dead platform surface routes to a sha-pinned
`demopatch` or **escalates**; all stack-operating tooling lives in `rosetta-extensions` (authored + tagged, consumed per-stack).

**Shape — 8 milestones, barrier → parallel fixes → prove-on-billion:**
```
M237 clean stage (HARD go/no-go barrier)
  ├─▶ M238 academy ─────────────┐
  ├─▶ M239 enterprise ──────────┤
  ├─▶ M240 content-fidelity ─┐  │
  │      └─▶ M241 language ─┐ │  │
  │            └─▶ M242 cockpit-UX
  ├─▶ M243 assign-WRITE ────────┤
  └────────────────────────────▶ M244 prove-on-billion (closer)
```

#### M237 — clean stage  (`section`, HARD go/no-go barrier)
**Status:** `done` (closed-complete 2026-07-21) — the barrier delivered on both fronts. **§1** fetch-verified
clone-freshness in `ensure-clones.sh` (rc-checked fetch, stderr never suppressed; a real **7-state pin model** in
`clones.lock.json` distinguishing deliberate-pin from stale-by-neglect; opt-in `DEMO_ADVANCE_CLONES` +
`DEMO_FRESHNESS_STRICT`) closes **F-M236-CLOSE-1**. **§2** the R1 pristine sweep is now **directory-driven** (all 14
manifests, was a hard-coded 3) closing **F-M236-CLOSE-2**. Both **dogfooded live on `billion`**. **§3/§4 confirmed-defect
ledger:** the **"202-behind" premise REFUTED** — the fetch-verified measurement (raw `git rev-parse`-confirmed) showed
billion's clones **0–2 behind** (frontend current), `ant-academy` the lone 5-behind surface; the suppressed-fetch reading
was itself the bug §1 fixes. **#1 menu RESOLVED** (hierarchical for managers, discharges M239's menu check); **#4 library
does NOT reproduce empty** → M239 re-scoped down to the cold-first-load flash; **#2 academy language SURVIVES** (real
academy-surface defect) → M238. Both survivors confirmed **Fate-2** (already in M238/M239 In-lists). rext code-of-record
tag `sound-check-m237-clean-stage` @ `533c489` (re-pinned to the hardened HEAD at close). test_tooling.py 146→**160**;
shellcheck + demo_knob_guard fences green; flake 5/5; deferral audit GREEN; **0 platform edits**.
**Goal:** The demo builds from CURRENT platform source, and the ambiguous UI defects are re-triaged on a correct build —
so every downstream fix is scoped against reality, not stale code.
**Shape:** `section` — HARD go/no-go barrier (the M217/M222 "clean stage" pattern). Any UI-defect triage on a stale-clone
demo is untrustworthy.
**Scope — In:**
- Fix clone-freshness in `rext demo-stack/ensure-clones.sh`: a **fetch-verified** freshness assertion (never
  suppressed-stderr — the billion `root` host-key failure that produced the 12-vs-202 mismatch) + an opt-in
  advance-to-`origin/main`-or-pinned-tag path + a **real pin model** so "pinned" vs "stale-by-neglect" is distinguishable
  (today both read `ref:"main"`/`"HEAD"`).
- Fix **F-M236-CLOSE-2**: the R1 pristine sweep enumerates all **14** patch manifests, not the hard-coded 3.
- Bring up a **fresh-clone demo on billion**; produce a **confirmed-defect ledger**: verify #1 menu now hierarchical for
  managers; RE-TRIAGE #2 academy-language + #4 library-empty on the fresh build (which survive a correct build?).
**Out:** any downstream fix (routed to M238–M243 by the re-triage); any platform-repo edit.
**Depends on:** none (opens the release).  **Parallel:** none (gates everything).  **Complexity:** medium.
**KB deps:** `corpus/ops/rosetta_demo.md` (§Clone freshness, anchored v2.5), `corpus/ops/demo/demopatch-spec.md`.
**Delivers →** `corpus/ops/rosetta_demo.md` (the clone-freshness mechanism) + `corpus/ops/demo/demopatch-spec.md` (R1 all-14-manifests).
**Open questions:** which of #2/#4 survive a fresh build (decided by the re-triage — it routes the downstream fix scope).

#### M238 — ant-academy reliability  (`section`)
**Status:** `done` (closed-complete 2026-07-21) — the first post-barrier fix landed. **One** chapter-body FS-published
demopatch (`academy-fs-published-chapter-body`, gated on the same `ACADEMY_DEMO_FS_PUBLISHED` as the M230 catalog patch)
fixed **BOTH #3 (Start→404)** AND **#2** — #2 was **not a distinct code bug** (locale is a `?lang=` query param, no
`/[locale]` route; the switcher is a sound EN↔IT toggle; the chapter-language 404 is the SAME backend-null path as #3),
both **proven live on `billion`** (chapter 404→200; `?lang=it` 200). Extended the academy coverage sweep
(`ANT_ACADEMY_CHAPTER_SECTION` + a general `mustNotInclude` negative marker + the catalog.json→chapter→`?lang=it`
probe), **mutation-verified to go RED on a broken academy**. Built a **directory-driven demopatch-inventory fence**
(`test_patch_inventory.py` — exact 15 + per-repo breakdown) closing the standing hygiene gap. rext code-of-record tag
`sound-check-m238-ant-academy-reliability` @ `3482a77` (re-pinned to the hardened HEAD at close). Touched suites green
(Python 183/183, TS 147/147, tsc clean), flake 5/5; deferral audit **YELLOW** (the 8 standing demo-stack failures =
the identical v2.5/M236 re-baselined set, 0 M238 regressions → re-fated Fate-2 → M244, D5); **0 ant-academy platform
edits**. Full live `coverage.spec.ts` billion sweep → M244 exit gate (c) (Fate-2, D4).
**Goal:** A hero can follow a course and actually consume a chapter; the language switch works.
**Shape:** `section`.
**Scope — In:**
- Fix **#3 (Start→404)**: the demo academy chapter-body path is unwired (bodies are backend-authoritative, no FS
  fallback; the catalog demopatch covers only the catalog). Wire a chapter-body demo path — a chapter-body FS-fallback
  demopatch analogous to `academy-fs-published-fallback`, OR wire the academy backend for the demo.
- Fix **#2 (language error** — re-triaged in M237; likely the same backend-null path).
- Extend the **academy presence/coverage sweep**.
**Out:** the enterprise-surface / talk-to-data fixes (M239); content-stories (M240+).
**Depends on:** M237.  **Parallel with:** M239, M240, M243.  **Complexity:** medium.
**KB deps:** `corpus/services/ant-academy.md`, `corpus/ops/demo/demopatch-spec.md`, `corpus/ops/demo/coverage-protocol.md`.
**Delivers →** `corpus/services/ant-academy.md` + `corpus/ops/demo/frontend-tier.md` (the chapter-body demo path + the extended academy sweep).
**Open questions:** chapter-body FS-fallback demopatch vs wiring the academy backend for the demo — which is revert-clean + sufficient?

#### M239 — enterprise surfaces  (`section`)
**Status:** `done` (closed-complete 2026-07-21) — the second post-barrier fix landed. **talk-to-data went FULL:** a real
AWS **Bedrock cred class** for `app` (5 genes — `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` **required·standard**,
`AWS_REGION`/`AWS_SESSION_TOKEN`/`CLAUDE_CODE_USE_BEDROCK` optional; **deliberately NOT critical**, R3 — a creds-less
demo still boots) provisioned **values-blind** via `/stack-secrets`, then **bridged** `app/.env → platform/.env` (the demo
backend reads `env_file: .env`; the M217 override drops the `~/.aws` mount, so env vars are the only vehicle).
**Proven end-to-end live** on demo-1 (cold reset-to-seed): a manager asked "how many members?" → a real Bedrock
round-trip → *"Cervato Systems has 51 members"* (backend agentic loop `tool_use`→`query_postgres`→`end_turn`, ~7 s).
**#4 library** and **#1 menu** were **no-defect verdicts, not manufactured fixes** (the `:5050` carry is already resolved
— offset endpoint baked; the grouped manager nav renders) — both live-GREEN. Harden (3 passes) strengthened a
near-vacuous flag-gate assertion, added Bedrock **measure-layer** tests, and **landed F1** (the disk pre-flight now
measures the **Docker VM disk**, not host `/` — the redis-exit-1 misattribution root cause). **The close found + fixed
2 real defects in M239's own new code** (both Fate-1, rext `cf89365`, mutation-verified regression tests): **D10** — the
VM-disk pre-flight would **abort the whole bring-up** when Docker is present-but-unreachable (no `|| true` under
`set -euo pipefail`), and **D11** — the bridge append lacked a **trailing-newline guard** (env-file corruption +
idempotency break). Deferral audit **YELLOW**: DEF-M239-01 (2nd F1 candidate "fail the BUILD loudly on ENOSPC") → Fate-3
→ M244 (D12); a **9th** demo-stack failure surfaced by the full sweep (`test_reap…test_a_RACED_listener_exits_silently`)
was **root-caused to a test-isolation collision** (hardcoded port 17700 vs a live demo-1 cockpit; reap.sh correct) →
Fate-3 → M244 with a fix recipe (D13); the standing-8 confirmed Fate-2 → M244 (M238-D5, 0 M239 regressions). rext
code-of-record: the 3 consumption tags (`sound-check-m239-enterprise-surfaces`/`-bridge-log-fix`/`-live-proof`) all
re-pinned to the reviewed HEAD `cf89365`. Touched suites green (Python 106/106 + full demo-stack 794 passed / 9
host-state fails [8 standing + the reap-17700 collision, 0 regressions], Go secretdna PASS), flake 5/5; **0
platform-repo edits**, values-blind throughout.
**Delivered docs:** `secrets-spec.md` (Bedrock cred class + 61-gene map) · `safety.md` §2.10 (secrets-posture shift +
operator-scope caveat) · `frontend-tier.md` (F1 VM-disk correction).
**Goal:** talk-to-data works live; the library grid loads first-time; the hierarchical manager menu is confirmed.
**Shape:** `section`.
**Scope — In:**
- talk-to-data **(a)** flag enablement (`NEXT_PUBLIC_DEMO_FLAGS_ALL` or a flag-gate demopatch, the M219/M232 pattern) +
  **(b) real AWS Bedrock creds** provisioned via `/stack-secrets` + the **secret-coverage DNA extension for `app`** (the
  `../hyper-studio/.env.example` template: `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_SESSION_TOKEN`/`AWS_REGION` +
  `CLAUDE_CODE_USE_BEDROCK`) + mounted/env-wired into the `app` compose service. **(user decision 1)**
- Fix **#4 (library empty-first-load)**: the client-fetch race + the open non-offset `:5050` `apps/web` endpoint carry.
- Confirm **#1 hierarchical menu** renders for managers (presence sweep).
**Out:** the media-porting content-fidelity work (M240); academy (M238).
**Depends on:** M237.  **Parallel with:** M238, M240, M243.  **Complexity:** medium (large if the Bedrock wiring balloons).
**KB deps:** `corpus/ops/secrets-spec.md`, `corpus/ops/safety.md`, `corpus/architecture/ai_architecture.md`.
**Delivers →** `corpus/ops/secrets-spec.md` (the Bedrock cred class for `app`) + a `safety.md` secrets-posture note.
**Open questions:** the demo secrets-posture for AWS creds (safety.md note — same class as AI-provider keys, now
present-not-absent for `app`).

#### M240 — content-stories fidelity  (`section`, HARD media-safety gate)
**Status:** `done` (closed-complete 2026-07-22) — the third post-barrier fix landed; a **complete 6-section close** (5 fixed
+ voice presence-only). The **HARD media-safety gate cleared first** (`safety.md` §3.8.1 raw-media/**VIDEO** amendment +
a fresh 2026-07-21 data-controller sign-off + the gender-coherence contract, landed BEFORE any media-exhibit code).
**Defect 1 (selection):** a `d.type = <cell sim_type>` **public-CTE** predicate stops the sole public interview sim from
leaking into non-interview voice cells (the CQ-1 root cause) — robust, not a slug exclusion; `asmt-voice-pass` re-pinned to
a real assessment-voice session. **Defect 3 (document):** the dropped criterion `input_data` is written at seed time via a
content-specific `contentCriterionResultCols` — the body is **inline `input_data.text_document`** (collaborative_doc), **NOT**
an S3 `storage_upload` blob, so there is **no blob to port** (this resolves the long-standing `DEF-M10-01` speculation for the
document facet). **Pass-rate:** a `ScoreMin/ScoreMax` band + `score ASC` tiebreak (100% only as fallback); the 5 still-100%
passed cells re-pinned to real 70–95% sessions (74/72/82/83/81), re-captured **values-blind** via `content-capture --only`,
both canonical presets regenerated, honesty gate green. **Voice (Defect 2) → DELIVERED = voice presence-only** (user decision
2026-07-22): the faithful `chime_status='not_available'` state IS the v2.6 deliverable — an honest "recording not available"
over a broken 500 player. The **real-video exhibit** is fully documented + **pre-blessed** (the render is exhibit-**by-reference**:
`bunny_video_id` + a read-only Bunny CDN signing key, **no media byte ever moves**) but the Bunny recording signing keys
(`BUNNY_RECORDING_CDN_TOKEN_KEY` + `BUNNY_RECORDING_PULL_ZONE_HOST`) are genuinely **absent from this box's dev-stack**, so the
exhibit routes to **M244** (`DEF-M240-01`, Fate-3, **user pre-approved 2026-07-22** — land it live IF the keys are reachable on
`billion`, else keep presence-only). **Delivers:** new `corpus/ops/demo/media-substrate-spec.md` + the `safety.md` §3.8.1
amendment + `session-clone-spec.md` §2/§4 bumps, indexed. rext code-of-record tag `sound-check-m240-content-stories-fidelity`
@ `ae0e869` (a fresh create+push at close). Full stack-seeding module green (16 pkgs, 0 fail), harden mutation-verified all 3
fixes + 6 deepening tests (0 bugs, 0 regressions), flake 5/5; deferral audit **YELLOW** (DEF-M240-01 → M244 user-pre-approved;
the standing-8/9 + DEF-M239-01 + reap-13 confirmed Fate-2/3 → M244); **PII discipline held** (customer media never entered
context; by-reference video; values-blind creds); **0 platform-repo edits**.
**Goal:** the cockpit's claim matches the session — right type, playable call, visible document — at a believable pass rate.
**Shape:** `section` — with a HARD internal media-safety gate (R1).
**Scope — In:**
- **Defect 1 (selection):** tighten `rext stack-seeding sourcing.go` to constrain the public sim's type to the cell type
  (exclude the interview sim from non-interview cells); re-pin `content-sessions.yaml`.
- **Defect 3 (document):** write the dropped `input_data` at seed time (`content_stories_write.go` / a content-specific
  criterion column set); + **port the document blob** if the body is a `storage_upload` (per user decision 2).
- **Defect 2 (voice):** **port the Chime/S3 recording** — capture the recording reference + re-host the audio in the demo
  storage tier + flip `chime_status` to available (per user decision 2).
- **Pass-rate (#4-feature):** add a score-band to `SelectionSpec` (`AND s.score BETWEEN 70 AND 95`), flip the tiebreak to
  `score ASC` (prefer lower), 100% only as fallback; re-capture.
- **HARD internal gate (before any customer media lands in a demo):** fresh data-controller sign-off + a `safety.md`
  amendment covering raw audio + full documents + a voice/document anonymization decision (a voice cannot be
  token-scrubbed). Likely consumes **DEF-M10-01** (S3 read access).
**Out:** the language toggle (M241); the cockpit-UX regroup (M242).
**Depends on:** M237.  **Parallel with:** M238, M239, M243.  **Complexity:** large. **Note:** the re-capture needs prod
read (`~/.pgpass`).
**KB deps:** `corpus/ops/demo/session-clone-spec.md`, `corpus/ops/safety.md`, `corpus/architecture/ai_architecture.md`
(Chime/LiveKit), `corpus/ops/demo/content-stories-routes.md`.
**Delivers →** a **new media-substrate spec** under `corpus/ops/demo/` + a `corpus/ops/safety.md` §3.8 amendment (raw media).
**Open questions / HARD gate:** raw-media PII is a larger data-controller call than v2.5's scrubbed text; the internal
gate (sign-off + safety amendment + anonymization decision) **must clear before any customer audio lands in a demo**.

#### M241 — content-stories language  (`section`, opens with a pool-count go/no-go)
**Status:** `done` (closed-complete 2026-07-22) — the fourth post-barrier fix landed; a **clean complete
5-section close**. The **go/no-go pool query** (read-only `marco_read`, counts+labels only) returned **GO**: 11 of
12 requirement tuples carry both languages, and it **surfaced the core defect** — 11 of the 13 pinned sessions
were actually **Italian** yet the seeder hard-coded every clone to `english`. **§2 plumbing** adds `s.language`
to the `sourcing.go` SELECT + an optional language filter, a `Language` field on the fixture + the
`content_manifest.go` projection, and flips the write to **`cs.Language`** (empty→english belt-and-braces).
**§3** sourced **10 EN/IT counterparts** (fixture **13 → 23**, denominator **29 → 49**), 11/12 tuples now
bilingual, **INTERVIEW Italian-only** (EN interview passes all out-of-band, EN interview fails = 0 — release
risk **R2** realized, the EN-only-fallback-per-tuple user decision). **§4** the fail-closed **`ValidateLanguageConsistency`**
gate (wired into `--content-export`; a `lang_toggle` disagreeing with its own coverage — solo-marked-toggle /
bilingual-marked-solo / invalid-language — FAILS) + a TS mirror (`content-language.unit.spec.ts`), with teeth.
**§5** the cockpit **EN|IT segmented toggle** (`_LANG_JS` raw-string injection-free; per-row language pill;
byte-clean when the manifest has no language axis). **3 harden passes** mutation-verified every gate AND closed
the **core-bug write-side gap** the milestone lacked — no test asserted the seeded `sessions.language` column
carried `cs.Language`, so reverting the write to the hard-coded `english` passed EVERY Go suite;
`TestContentStorySeeder_WritesRealLanguage` closes it (the exact defect v2.6 "sound check" kills). The close
found 1 doc fix (a stale `29 → 49` denominator) + recorded 2 adversarial scenarios (both handled). rext
code-of-record tag `sound-check-m241-content-stories-language` @ `17beede` (created + pushed at close). Go 2005
test funcs (+6 M241); `test_cockpit.py` 142 (136 pass / 6 pre-existing academy+overlay, 0 new from M241, Fate-2
→ M244); TS 151 unit specs; flake gate 5/5 all stacks; deferral audit **YELLOW** (0 new deferrals); **0
platform-repo edits**; PII discipline held (structure/presence/label assertions, never a translated value).
**Delivered docs:** `content-stories-spec.md` (§2 language schema · §4 fail-closed both sides · §7.6 toggle) +
`session-clone-spec.md` §2.1 (pool query + `cs.Language` write).
**Goal:** each session is consumed in its intended language, with an EN/IT cockpit toggle.
**Shape:** `section` — opens with a read-only pool-count go/no-go query (R2).
**Scope — In:**
- **Read-only prod pool-count query FIRST** (IT sessions per requirement tuple — the interview-scarcity go/no-go). **(user decision 3)**
- add `s.language` to `sourcing.go` SELECT + optional filter; add a `language` field to the fixture +
  `content_manifest.go` projection (re-touch the `CanonicalFileMatchesProjection` honesty gate); use `cs.Language`
  instead of the hard-coded `sessLanguageEnglish`.
- source EN+IT pairs per tuple where IT exists; **EN-only fallback per tuple** where absent (toggle hidden/disabled
  there); cockpit toggle swaps the login-and-land target.
- Extend the **content-stories sweep** for language (assert structure/presence, **never** the translated value — P2
  forbids copy assertions).
**Out:** the row-layout/tab-selector cockpit-UX (M242).
**Depends on:** M240 (shares `stack-seeding` + the re-capture).  **Parallel:** none (serial after M240).  **Complexity:** medium.
**KB deps:** `corpus/ops/demo/session-clone-spec.md`, `corpus/ops/demo/content-stories-spec.md`, `corpus/ops/demo/coverage-protocol.md`.
**Delivers →** `corpus/ops/demo/content-stories-spec.md` (the language field + EN/IT toggle + the re-touched honesty gate).
**Open questions:** IT interview sessions may not exist (R2) — the pool query decides per-tuple coverage; EN-only fallback where absent.

#### M242 — cockpit UX  (`section`)
**Status:** `done` (closed-complete 2026-07-22) — the fifth post-barrier fix landed; a **clean complete
3-section close**, render/CSS only in the presenter cockpit (`cockpit.py`), **0 data/seed change, 0 manifest
schema change, 0 platform-repo edits**. **(1) row layout:** the Content-stories rows regroup by requirement
tuple `(sim_type, modality)` (non-sim products fall back to `label`) → one row per tuple —
`target label (+ modality pill) | passed login options | not-passed login options` side by side — with
symmetric **"No passing / No failing run"** empty markers (D3) and a presence-only (ai-labs) inline slot;
render-layer only, `_content_session_row` split into `_content_tuple_row` + `_content_login_cell` so the M241
EN/IT toggle's atomic `.session` cell was untouched. **(2) tab selector:** the `.tabs` bar moved out of
`main_body` into the white `<header>` (flex `.hwrap`, right, vertically centered); `_TAB_JS` is
placement-agnostic (no JS change); the **byte-identical-when-no-content-manifest invariant** is preserved via a
separate `else`-branch reproducing the verbatim pre-M242 header (mutation-STRONG). **(3) hero avatar by
user-type:** `_avatar_class` — manager = orange (`.b-manager`) / employee = indigo (default) / hiring candidate
= a net-new **teal** (`--cand #0f766e`/`--cand-soft #ccfbf1`, ~4.86:1 AA), **manager-wins** order (a hiring
recruiter reads as a manager). **2 harden passes** mutation-verified every render branch and found + fixed **2
toothless tests** (a wrong-column mutant + an unescaped-non-sim-title mutant) + added a WCAG AA-contrast pin. The
**close's adversarial pass** found + landed **1 latent D3-invariant gap** (D8): the verdict-column split crossed
the M241 per-cell language toggle so an **unbalanced bilingual tuple** (a verdict present in only one language)
could show a verdict header over a blank body — fixed client-side by `_LANG_JS.syncEmpty()` (re-derives the
per-column empty marker on toggle + on load, **0 server-markup change**), guarded mutation-verified; the
canonical seed is currently balanced (0 live occurrences), so it guards a latent regression. rext code-of-record
tag `sound-check-m242-cockpit-ux` @ `73d37d5` (re-pinned at close to the close-fix HEAD). `test_cockpit.py` 164
(158 pass / 6 pre-existing academy+overlay, 0 new from M242, Fate-2 → M244); full demo-stack suite **839 pass /
9 fail** (the standing set, 0 new); Go 2005 + TS 151 unchanged (Python-render-only); flake gate 5/5; deferral
audit **YELLOW** (0 new deferrals); **0 platform-repo edits**. **Delivered docs:** `cockpit-spec.md` (the v2.6
UX-pass section + the avatar-by-role glance note) + `content-stories-spec.md` (§7.2 tuple-regrouped row + the
render-helpers + empty-marker-under-toggle notes; §7.6 M242-delivered coexistence).
**Goal:** the Content-stories tab reads clearly and the heroes are legible by role.
**Shape:** `section`.
**Scope — In:**
- **(1) row layout** — regroup by requirement tuple `(sim_type, modality)` → `target | passed login options | not-passed
  login options` on one row (render-only; fields exist).
- **(2) tab selector** — move into the white header, right, vertically centered (restructure `cockpit.py` header to
  flex; **preserve the byte-identical-when-no-content-manifest invariant**).
- **(3) hero icon bg by user-type** (manager = orange / employee = indigo, reuse the badge palette; derive a candidate
  color = `is_hiring && vantage != manager`).
- Extend the **cockpit specs**.
**Out:** any data/seed change (M240/M241); platform edits.
**Depends on:** M240 + M241 (the row regroup wants the pass/fail variants + the language axis).  **Parallel:** none.  **Complexity:** medium.
**KB deps:** `corpus/ops/demo/cockpit-spec.md`, `corpus/ops/demo/content-stories-spec.md`.
**Delivers →** `corpus/ops/demo/cockpit-spec.md` + `corpus/ops/demo/content-stories-spec.md` (the row-regroup + header layout + role-color).
**Open questions:** none blocking.

#### M243 — assign-WRITE Playthrough  (`section`)  [realizes reserved M238]
**Status:** `done` (closed-complete 2026-07-22) — the sixth and last post-barrier fix landed the **one net-new
hero journey** and closed the **~10-routing `DEF-M235-03` / M204 assign-WRITE carry** that had ridden **5
releases** (fresh-dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20 with M243 as its "land-it-or-DROP" expiry — LANDED,
a deferral discharged as Fate-1). `pt-assignment-assign` is the **FIRST MUTATING Playthrough**: a manager
(Morgan / `pt-manager`) logs in → `/enterprise/assignments` → Skill Paths tab → opens the assign builder for a
member with **no** skill-path org-assignment → keyboard-picks a public-catalog skill path + accepts the pre-filled
deadline → **Assign** → and the write **LANDS** — a real `public.organization_assignments` row
(`app.createOrganizationAssignments`) **read back** through the real members surface as the assignable "Assign
Skill Path" affordance count dropping by **exactly one** (that member's cell flipped to the assigned title). This
is where the release's **anti-toothlessness thesis is sharpest**: the read-back FLIP is the proof, **not** a
closed modal — a silent write-failure leaves the count at `before` and the 20 s poll times out RED (fail-closed).
**No new seed CODE:** Org A (Meridian Labs, 40) pre-assigns skill paths to only a handful, so ~34 are
deterministic assign TARGETS; the precondition is DECLARED + enforced — UC1 names `seed.preconditions:
[public-catalog, org-unassigned-member]`, the latter added to `seed-worlds.yaml` in **lockstep** (a future
"assign-to-everyone" seed trips `ptvalidate`, not a mystery live failure). **antd-v6 lesson (D4):** the catalog
Select is an `rc-virtual-list` whose `role="option"` nodes carry the raw uuid + read non-visible, so the page
object commits the first real option by **keyboard** (`ArrowDown`+`Enter`). **Live GREEN at build** (demo-1, 7.9 s)
+ **DB-verified** (`organization_assignments` skill_path/active 6 → 7, the exact assign); the **cold re-drive on
`billion`** is **M244's** declared scope. **1 harden pass** closed the `isOnAssignments` orthogonal-dimension
parity gap (+4 unit tests) + empirically mutation-verified the read-back FLIP and the two Go honesty-teeth
(seed-rename reddens `ValidatesAgainstSeedWorlds`; UC1→TODO reddens `ManagerCoverageIsPresent`). The **close** was
**clean — 0 code fixes**; its adversarial pass verified the count-delta read-back **fail-closed** under
pagination/re-sort [S1], refetch-transient [S2], and antd mis-pick [S3] (S1 empirically clean on the real 40-member
roster). `ptvalidate` **7 products / 16 use cases / 16 live / 0 TODO** (was 15 live / 1 TODO — the last TODO
flipped); Go **2005** unchanged (manifest modified in place, +0 func); combined TS unit-spec run **73 → 77**
(url-shapes.unit 60 → 64, the `isOnAssignments` block); live `@pt:` specs **15 → 16**; flake **0** (77 5/5);
deferral audit **YELLOW** (0 new; the standing carry resolved); **0 platform-repo edits**. rext code-of-record tag
`sound-check-m243-assign-write-playthrough` @ `2ef5962` (unchanged by close — no code fix). **Delivered docs:**
`corpus/ops/demo/playthroughs.md` (count 15 → 16 / 0 TODO + the assign-WRITE subsection + the M243 page-object
bullet) + `README.md` + `CLAUDE.md` count updates.
**Goal:** the one net-new hero journey — a manager assigns content with a deadline and it lands.
**Shape:** `section`.
**Scope — In:**
- `playthroughs/manifest/assignment-monitoring.yaml` UC1 (`assign-and-track.UC1`, currently `TODO`).
- a new `/enterprise/assignments` page object.
- possibly a `pt-world` precondition (assignable content + target member) in lockstep with `seed-worlds.yaml`.
- the spec `e2e/tests/assignment-assign.spec.ts` tagged `@pt:...UC1`. Takes the corpus **15 → 16** live Playthroughs, 0 TODO.
**Out:** the re-prove-on-billion live drive (M244 executes it).
**Depends on:** M237 (fresh demo).  **Parallel with:** M238/M239/M240.  **Complexity:** medium. **Note:** needs a live
browser drive against a running demo.
**KB deps:** `corpus/ops/demo/playthroughs.md`.
**Delivers →** `corpus/ops/demo/playthroughs.md` (15 → 16 live Playthroughs; the assign-WRITE half of the M204 flow closes the ~10-routing DEF-M235-03/M204 carry).
**Open questions:** does the `assign` WRITE need a `pt-world` precondition co-authored with `seed-worlds.yaml`?

#### M244 — prove on billion  (`iterative`, the closer)  [realizes reserved M237]
**Status:** `done` — **gate MET 8/8 live on `billion` 2026-07-23** (closed-on-gate; the terminal milestone, merged
`--no-ff` into `release/02.60-sound-check` — the release→main merge + the `v2.6` tag are `/developer-kit:close-release`'s
job). Re-proved the whole feature cold reset-to-seed on the `billion` Tailscale VM, **0 platform edits**: **(a)**
ORG-CLEAN 0 surviving tokens · **(b)** content-stories **47/47** landed of the 49-pair denominator (2 Bunny-absent voice
**player** cells held presence-only) · **(c)** the **40 live-browser specs** green (24 stack-verify + **16/16 Playthroughs**,
96 cases in one clean full run) · **(d)** anon academy `/library`+`/free` twin renders real cards · **(e)** serve-reap
7→0 · **(f)** 3 v2.3 drift-carries incl. BURNIN-M221 (`/dev-up --public-host` graphql dev-2 live-cycled + tailnet-reachable)
· **(g)** interview plan-section alignment assertion (caught + fixed a real v1.3→v1.4 plan drift) · **(h)** all 6 v2.6 fixes
live + **p95 click→ACCESS 1.46 s / 1.31 s**. 27 iters (24 tiks / 3 toks — all HELD the strategy; the flat binary-per-gate
metric was diagnosed a coarse-metric artifact, not a stall). A real iter-25 finding fixed durably: the demo image must be
compiled from the **pinned** ref, not the highest fetched tag (the `launched_by` version-skew). Close: 2 doc
reconciliations, **0 code fixes**; deferral audit **YELLOW** (standing-8 demo-stack test debt + DEF-M239-01 re-fated
Fate-3 → close-release). rext consumption tag `sound-check-m244-content-sweep-robustness` @ `498b1a5`. **All 8 v2.6
milestones M237 → M244 CLOSED — the release is ready for `/developer-kit:close-release`.**
**Goal:** re-prove the whole feature — v2.5's headline AND every v2.6 fix — live on `billion`, cold reset-to-seed.
**Shape:** `iterative` — live-proof is measurement-driven (the M221/M236 lineage); iters until the gate.
**Exit gate:** on a cold reset-to-seed on `billion`: **(a)** `ORG-CLEAN` reports **0** surviving source-org tokens (or
each dispositioned) — **RUN FIRST**, read-only, before the bring-up; **(b)** content-stories `run-content-stories.sh`
green at the shipped harness with the CQ-1 grader fix + CQ-2 runner wiring + externally-sourced `EXPECTED_PAIRS`
(discharges CLOSE-D3); **(c)** the **40 live-browser specs** (24 stack-verify + 16 Playthroughs) execute green (T-3); **(d)** the anonymous academy
`/library`+`/free` twin renders real cards (S-1); **(e)** `DEF-M226-01` — the serve-reap self-resolution claim is
**actively tested or DROPPED**; **(f)** the 3 v2.3 drift-carries burned-in live (`BURNIN-M221` / `F-M220-4` /
`PROBE-M218-c3`); **(g)** the interview plan-section-id **alignment assertion** added + green (S-8/S-9); **(h)** every
v2.6 fix (academy course-start, talk-to-data live answer, library, content fidelity incl. media, language toggle,
cockpit UX) proven live; **p95 click→ACCESS < 5 s** hero vantages. **0 platform edits.**
**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` + `coverage-protocol.md` + `playthroughs.md`.
**Re-scope trigger:** 5 consecutive toks without a viable strategy → user-strategic-replan.
**Out:** new feature work (all built by M238–M243); content-seat latency (out of scope, per the v2.5 M236 precedent).
**Depends on:** M238, M239, M240, M241, M242, M243 (all fixes).  **Parallel:** none (terminal).  **Complexity:** medium (iterative).
**KB deps:** `corpus/ops/verification.md`, `corpus/ops/demo/tailscale-serve.md`, `corpus/ops/demo/coverage-protocol.md`, `corpus/ops/demo/playthroughs.md`.
**Delivers →** none (proof milestone; extends the coverage/playthrough manifests + burns in the carries).
**Open questions:** none blocking (the multi-part exit gate is the spec).

### Version plan (v2.6)

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v2.6** | **sound check** | Reliability / field-hardening — *make everything that's built actually get built + provisioned.* Barrier → parallel fixes → prove-on-billion; talk-to-data live via real AWS Bedrock creds; media ported behind a HARD PII gate; EN/IT language toggle; cockpit UX; the net-new assign-WRITE Playthrough | M237 (HARD go/no-go) → { M238 ∥ M239 ∥ M240 → M241 → M242 ∥ M243 } → M244 | 🟡 **IN DEVELOPMENT** (branch `release/02.60-sound-check`, designed 2026-07-20; **M237 + M238 + M239 CLOSED 2026-07-21, M240 + M241 + M242 + M243 CLOSED 2026-07-22, M244 prove-on-billion CLOSED 2026-07-23 (gate MET 8/8 live, closed-on-gate)**; all 8 M237→M244 closed, ready for `/developer-kit:close-release`; realizes reserved M237/M238; tooling + docs only, zero platform-repo edits; tag will be `v2.6`) |

### Execution graph
```
M237 clean stage (HARD go/no-go barrier — fix clone-freshness, re-triage on a fresh build)
  ├─▶ M238 academy reliability ──────┐
  ├─▶ M239 enterprise surfaces ──────┤
  ├─▶ M240 content-stories fidelity ─┐  │   (HARD media-safety gate)
  │      └─▶ M241 language ─────────┐ │  │
  │            └─▶ M242 cockpit-UX ─┤ │  │
  ├─▶ M243 assign-WRITE Playthrough ┼─┤
  └─────────────────────────────────▶ M244 prove-on-billion (iterative closer)
```
M238 ∥ M239 ∥ M240 ∥ M243 fan out off the M237 barrier; M241 is serial after M240 (shares `stack-seeding` + the
re-capture), M242 serial after M241 (wants the pass/fail variants + language axis). M244 is terminal — depends on all fixes.

### Risk map
- **R1 (blocks-quality) — raw-media PII.** Porting real customer voice + documents is a larger data-controller call than
  v2.5's scrubbed text; a voice cannot be token-scrubbed. **Mitigation:** M240's HARD internal gate (fresh sign-off + a
  `safety.md` raw-media amendment + a voice/document anonymization decision) **must clear before any customer audio lands
  in a demo**; the standing control is the VPN/tailnet scope (`safety.md` §3.8).
- **R2 (blocks-scope) — language scarcity.** IT interview sessions may not exist. **Mitigation:** M241 opens with a
  read-only pool-count query that decides per-tuple coverage; **EN-only fallback per tuple** where IT is absent (toggle
  hidden/disabled there). No blocking.
- **R3 (degrades-quality) — AWS Bedrock creds in the demo.** A new present-not-absent secret class for `app`.
  **Mitigation:** provision via `/stack-secrets` + extend the secret-coverage DNA; a secrets-posture note in `safety.md`
  (same class as the AI-provider keys, now present-not-absent for `app`).
- **R4 (dependency) — prod read + live billion.** M240/M241 re-capture + M244 re-prove need `~/.pgpass` prod read + a
  reachable `billion`. **Mitigation:** both confirmed available at the v2.5 close (`billion` up + reachable).
- **R5 (process) — v2.5 not pushed to origin.** The v2.6 branch was cut from **local** `main`; `main` + the `v2.5` tag
  are still local-only. **Mitigation:** flag to the user — do **not** auto-push; the user runs the v2.5 origin push +
  the v2.6 origin publish on their own cadence.

---

## Done — v2.5 "the playbill" (SHIPPED 2026-07-20, tag v2.5)

**Theme.** *Show the platform's content, played.* Two threads on the same mature machinery (the M35 Stories & Heroes
seeder fleet + the M43 cockpit + the M45 AI-fill engine + the M46 Directus serve-grants + the M202 Playthrough/coverage
proof harnesses): **A —** fill the empty ant-academy grid so it renders real content the way taxonomy/skill-path do;
**B —** a 2nd "Content stories" cockpit tab listing **played sessions** per content product, each with a login-and-land
**as-player** and **as-manager** action, cloned from **anonymized real production sessions**, source-pinned deterministically.

**User decisions baked in (2026-07-19):** one combined release · **real customer-session sourcing** accepted as the
data-controller's call, demos kept **VPN/tailnet-scoped** (not open-internet), release **amends `safety.md` Part 3** to
the honest posture (anonymized-real, VPN-bounded) · academy fill **production-faithful** (no "Draft" chip) · AI-labs +
the academy content-product section **scoped by the M231 spike**.

**Hard constraint (carried, unchanged):** **zero platform-repo edits** — a runtime-computed result page that won't
render from a seeded row routes to a sha-pinned `demopatch` or **escalates**; all stack-operating tooling lives in
`rosetta-extensions`.

**Shape — 8 milestones, spike-first, largely sequential:**
```
M229 ──→ M230 ──→ M231 (HARD go/no-go) ──→ M232 ──→ M233 ──→ M234 ──→ M235 ──→ M236
(A: academy)         (B: barrier)          (B: seeder→manifest→tab→prove→prove-live)
  M229 ∥ M231 research can overlap; M230 must land before M235's academy section
```

#### M229 — academy content-model re-ground  (`section`, small)
**Status:** `done` (closed-complete 2026-07-19) — corrected `ant-academy.md` (+ `frontend-tier.md`, `run_guide.md`,
`CLAUDE.md`) from the false "no backend / static JSON / only Clerk" model to the DB-authoritative catalog (grid →
academy subgraph over GraphQL → `emptyCatalogView()` on failure), added § The Content Model, and fixed the F4
mis-attribution. All code-verified. 4 docs, 0 platform edits, all Fate-1.
**Goal:** Correct the materially-stale, actively-misleading `ant-academy.md` — document the true DB-authoritative catalog
model + the demo empty-render root cause — BEFORE any fill code (the KB-fidelity prerequisite that mis-routed F4 for a
whole release when wrong).
**Scope — In:** rewrite `corpus/services/ant-academy.md` (remove *"Backend services: None / no GraphQL / static JSON"*;
document the v0.5.1 M7 DB-authoritative path `page.jsx → resolveCatalogView → getBackendCatalogView → academy subgraph`);
document WHY a demo grid renders 0 cards (unset `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` + empty app academy tables →
`emptyCatalogView()`); disambiguate the two "catalog" paths (grid READS app `internal/academy` via GraphQL; `build-catalog.mjs`
WRITES the unrelated `public/catalog.json` FS index); note the `ACADEMY_SHOW_DRAFTS`/`NODE_ENV=development → mergeDrafts()`
draft layer; correct the F4 mis-attribution in `frontend-tier.md`.
**Out:** any code/env change (M230); the Content-stories tab (Thread B).
**Depends on:** none.  **Parallel with:** M231.  **Complexity:** small.
**Delivers → `corpus/services/ant-academy.md`** (corrected, DB-authoritative).
**Open questions:** should `ant-academy.sh` wire `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` to the demo offset router regardless
of fill strategy? Is the academy subgraph even composed into the demo's offset Cosmo router?

#### M230 — academy demo-fill (production-faithful)  (`iterative`, medium)
**Status:** `done` (closed-incomplete/pragmatic 2026-07-19) — the Option C `academy-fs-published-fallback` demo-patch
(rext tag `playbill-m230-academy-fs-published`) is BUILT + runtime-proven (59 real cards, 0 Draft chips, exact
DB-authoritative code path, byte-clean revert; 14 unit tests, flake 3/3). Gate MET-BY-PROXY; the FORMAL cold-`/demo-up`
card-count sweep + the local next-web re-anchor + the `getPublicCatalogView` anon-routes follow-on are carried to
**M235/M236** (Fate-3, homed; see `m230-academy-demo-fill/carry-forward.md`). 0 platform edits.
**Goal:** Make the demo (and dev) ant-academy home grid render REAL academy content the way taxonomy/skill-path do —
**production-faithful, no "Draft" chip** (user decision) — closing the year-old F4 carry inside the zero-platform-edit wall.
**Exit gate:** on a cold `/demo-up`, the academy home grid renders real cards (count ≥ floor) for the employee vantage,
**no "Draft" chip**, via the real DB-authoritative GraphQL path (or a faithful equivalent), 0 prod-ejects, verified by the
coverage sweep on a **rendered-card count** (not the M53 port-serves + SSR-title check that let F4 slip).
**Iteration protocol:** `corpus/ops/verification.md` + `coverage-protocol.md`. The first tik decides the faithful path:
**Option C** (sha-pinned rext demo-patch restoring the M7 FS-as-published fallback on the ephemeral clone — `demopatch-spec.md`)
vs **Option B** (a net-new firewalled academy-content snapshot surface, capture→replay the public academy catalog rows into
the demo app DB + wire the endpoint + compose the subgraph). Draft-layer Option A is **rejected** (visible chip).
**Out:** any ant-academy platform-repo edit (routes to a demo-patch or escalates); an academy SESSION/progress model
(does not exist server-side — a Thread-B concern).
**Depends on:** M229.  **Complexity:** medium (large if Option B forced).
**Delivers → `corpus/ops/demo/frontend-tier.md`** (corrected F4 attribution + the shipped academy-fill mechanism); +
(conditional, Option B only) a new academy-content surface in `corpus/ops/snapshot-spec.md`.
**Open questions:** does prod academy content live in app `internal/academy` as firewallable public rows (needed for Option
B), and what is its public predicate? Is a demo-patch (Option C) sufficient + revert-clean?

#### M231 — content-stories feasibility spike + result-route map  (`section`, medium — HARD go/no-go)
**Status:** `done` (closed-complete 2026-07-19) — **Thread B is a GO.** Delivered `content-stories-routes.md` (result-route
map + prove-by-render classification + sourcing/anonymization contract + modality catalog + verdicts). Central risk
resolved: the sim result page reads a **persisted DB row** (plain SELECTs, no live recompute) → a cloned session
renders. Sim + Skill-path GO; **Interview GO behind a PostHog-flag demo-patch (D3→M232)**; **AI-labs OUT** (presence-only,
D4→M234); **Academy IN** (seedable chapter progress, D5→M234). Fixed 3 stale corpus claims inline (incl. the M228
intercepting-route misdiagnosis). 0 platform edits.
**Goal:** HARD go/no-go barrier (mirrors v2.4's M222): before building anything, discover the exact per-product player+manager
result routes, PROVE-BY-RENDER which land from seedable rows vs are runtime-computed-blank, confirm the **prod-session
sourcing mechanism** (read → pick interesting → pin by prod session-id), catalog public sims by modality, and rule
AI-labs + the academy section in/out.
**Scope — In:** enumerate per (product × {player, manager}) the exact result route (sim player `/sim/<slug>/result/<sessionId>`;
manager `/enterprise/activity-dashboard/<kind>/<id>/<userId>`; hiring `apps/hiring` scoreboard; interview `user_report`/`manager_report`;
skillpath legacy) + classify each by probe render (renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface);
confirm the **db-access read path** can select interesting real prod sessions per type + the **anonymization surface**
(which fields scrub cleanly, which free-text needs handling) + how to **pin by prod session-id**; catalog captured public
sims by task modality (confirm ≥2 voice + 1 code + 1 document assessment SOURCES exist); assess AI-labs feasibility (labs-api
nil) + the ant-academy "session" question.
**Out:** building the seeder/manifest/tab (M232–M234); any platform edit to make a runtime page render (routes to a demo-patch
or escalates — decided here).
**Depends on:** none (parallel to M229/M230).  **Complexity:** medium.
**Delivers → `corpus/ops/demo/content-stories-routes.md`** (manager-view eligibility matrix + per-product result-route map +
public-sim-by-modality catalog + AI-labs feasibility verdict + the prod-session sourcing+anonymization contract).
**Open questions:** does `/sim/.../result/<sessionId>` recompute live (unseedable) or read a persisted row a clone could seed?
which products actually HAVE a manager result route? is invoking the platform's own session-clone subcommand in-stack acceptable?

#### M232 — session-clone / sourcing seeder  (`section`, large)
**Status:** `done` (closed-complete 2026-07-19) — the ContentStorySeeder **COPIES real prod sessions** (feedback/
transcript/submission/interview report/skill node-ids) with a **best-effort PII scrub** (names/org→placeholders,
emails/phones/urls redacted), re-tenanted, non-manager-played, **source-pinned by prod session-id** (rext tag
`playbill-m232-sections-copyreal`); interview render flags via 2 sha-pinned demopatches; `safety.md` §3.8 amended to
the honest copy+scrub / **data-controller-accepted residual-risk** / VPN-scoped posture; deliverable `session-clone-spec.md`.
**A synthesize-first build was reworked to copy-real per the user's explicit decision** (weekly-limit interruption
recovered cleanly, no work lost). Guardrails flake 5/5, full suite GREEN. 0 platform edits.
**Goal:** Build the seeder that **copies real production sessions, anonymized where possible, re-tenanted into a manifest
org, non-manager-played, and source-pinned by prod session-id** — the deterministic realization of "clone real sessions."
**Scope — In:** read the selected real prod sessions (via the `db-access` read path, at authoring time) and reconstruct the
full seedable result substrate per session in the target org (`jobsimulation.sessions` + `public.local_jobsimulation_sessions`
mirror + `validation_attempt_results`/`_skill_results`/`_criterion_results` + `actors`/`interactions` transcript +
`interview_extraction_results.user_report`/`manager_report`), **passed + not-passed** via completion/score bands, all G14-valid
enums; **anonymize where possible** (structured fields scrubbed; free-text handled per M231's contract); net-new **code**
(roadrunner) + **document** (upload/PDF) assessment modalities; enforce **owner-is-player-vantage, never a manager seat**;
**pin the prod source session-id + the anonymization transform** in `seed-generation-manifest.yaml` (deterministic reseed);
**amend `corpus/ops/safety.md` Part 3** to the honest posture (content-story demos carry anonymized-real session data,
**VPN/tailnet-scoped**, source-pinned — the "nothing behind the door" guarantee gains a documented, bounded exception).
**Out:** the manifest projection + cockpit tab (M233/M234); playable voice/Chime/LiveKit recording artifacts (transcript-only,
assert at boundary); AI-labs sessions unless M231 ruled them feasible; opening demos wider than VPN/tailnet.
**Depends on:** M231.  **Complexity:** large.
**Delivers → `corpus/ops/demo/session-clone-spec.md`** (the copy+anonymize sourcing pattern, the VPN-scoped safety argument,
the source-pin contract, the no-manager-played rule, the type × modality × passed/not-passed matrix) + the `safety.md` Part 3 amendment.
**Open questions:** reuse existing hero seats as players or mint per-session anonymized player seats (brief leans mint; each
must map to a real seeded `public.users` row)? are a synthesized/scrubbed transcript + code-submission + document sufficient,
or must a real recording be playable (blob-mirroring is deferred, DEF-M10-01)?

#### M233 — content-stories manifest + honesty gate  (`section`, medium)
**Status:** `done` (closed-complete 2026-07-19) — merged `--no-ff` into `release/02.50-the-playbill`. Delivered the
**`content_products[]` projection** (`BuildContentProducts`, rext `stack-seeding/seeders/content_manifest.go`): per
content product, the played sessions each with **player+manager seat keys + result paths + `has_manager_view` +
per-product app_base + per-type icon** — SINGLE-SOURCED from the SAME content-session fixture the M232 seeder seeds
from (the player seat OWNS the seeded session by construction, D9). **Honesty-gated** (`CanonicalFileMatchesProjection`
+ a `HasTeeth` meta-test) so the checked-in `content-manifest.json` can't drift; **fail-closed** (`ValidateContentManifest`
drops-with-reason + fails loud — no fabricated links); emitted by **`stackseed --content-export`**. **Open question
RESOLVED (`#D-M233-1`):** a SEPARATE `content-manifest.json` (the stdlib-Python cockpit reads JSON, not YAML); the M232
source-pins stay folded in `seed-generation-manifest.yaml`'s `content_sessions` block. The load-bearing
flat-index-survives-drops seat single-source invariant is verified at both ends + pinned by
`TestContentProducts_FlatIndexSurvivesDrops`. Deliverable `content-stories-spec.md` (the manifest-schema half). rext tags
`playbill-m233-content-manifest` @ 9f0ab1c (build) + `playbill-m233-content-manifest-hardened` @ c30fee3 (harden — 100%
function coverage on the projector). Close review near-clean (1 fix: the `#D-M233-3` back-ref tag); deferral audit YELLOW/
0 blockers; flake 5/5; **0 platform edits.** The bring-up export wiring + cockpit render + `content-player-<idx>` seat
registration + non-simulation player-path builders are M234 (Fate-2, confirmed in M234's `In:` list).
**Goal:** Project a second, auditable `content_products[]` manifest block (peer to `population.orgs[]`) that pins each
session's prod source deterministically and is honesty-gated so it cannot drift from the seeded sessions.
**Scope — In:** a `content_products[]` projection (Simulation / Skill-path legacy / Skill-path new / AI-labs) each listing
pinned sessions with player+manager seat keys, result paths, `has_manager_view`, per-product app-base, and a per-type icon
key; project it from a preset via `stackseed --content-export` (or a 2nd block in `BuildCockpitManifest`), guarded by a
`CanonicalFileMatchesProjection`-style test (the D9 single-source discipline); fail-closed resolver when a pinned prod-source
id doesn't resolve in the replay (no-fabrication); fold pinned sources into the downloadable `seed-generation-manifest.yaml`.
**Out:** the cockpit render (M234); the seeder (M232).
**Depends on:** M232.  **Complexity:** medium.
**Delivers → `corpus/ops/demo/content-stories-spec.md`** (the manifest-schema half).
**Open questions:** one manifest with a 2nd block + client tab, or a separate `content-manifest.json` + endpoint (better
preserves D9 + the non-fatal bring-up)?

#### M234 — content-stories cockpit tab  (`section`, medium)
**Status:** `done` (closed-complete 2026-07-19) — merged `--no-ff` into `release/02.50-the-playbill`. The **render
half** of Content stories: `cockpit.py` gains the 2nd **"Content stories"** tab (client-side toggle, stdlib-only, no
manifest data in JS) reading the M233 `content-manifest.json` — per-product sections, per-session rows with
per-`sim_type` FontAwesome icons + up to two login-and-land CTAs (as-player fake-FAPI handshake iff
`player_result_path`; as-manager omitted where `has_manager_view=false`), per-product app-base routing
(web:3000/hiring:3001/academy:3077 + app-base fallback), **AI-labs presence-only** (no CTA), **academy direct-origin
link** (M53 cookie seam, not FAPI); absent/empty manifest ⇒ byte-identical pre-M234 page. `roster.go` appends
`content-player-<idx>` seats (one per distinct owner slot the projection references, after all heroes) single-sourced
with the UsersSeeder via the new `storyPopulationNames`; the `--roster-export` warning re-keyed on `RosterHeroCount`.
`up-injected.sh` wires `--content-export` + `--content-manifest` (non-fatal). **Renderer handles ALL dispositions —
unit-proven, not browser-proven** (Python 249 pass / 6 pre-existing fail / **0 new**; Go +8 test funcs 1931→1939 via
`git grep '^func Test'`; flake **5/5** both stacks). Close near-clean (2 record fixes: an `Adversarial review`
decisions subsection + 5 `(#M234-DK)` back-ref tags in §7); deferral audit **YELLOW / 0 blockers** (the 14-fail
demo-stack chronic homed at release-close, not re-woken); **0 platform edits**. rext tags `playbill-m234-content-tab`
@ 7f55eb4 (build) + `playbill-m234-content-tab-hardened` @ fd457bf (harden). Non-sim fixtures (ai-labs/academy/
skill-path) + prove-every-CTA-lands live = **M235** (Fate-2, verified homed in M235's `In:` + exit_gate).
**Goal:** Add the 2nd "Content stories" tab to `cockpit.py` beside "Org stories" — sections-per-content-product, a list of
played sessions each with per-type icons and TWO login-and-land CTAs (as-player / as-manager, manager omitted where
`has_manager_view=false`).
**Scope — In:** a client-side tab toggle in `render_page()` (reuse the shipped `_OVERLAY_JS` pattern; stdlib-only,
standalone-served); per-product sections rendering the M233 manifest; per-session rows with per-type FontAwesome icons; two
fake-FAPI deep-link CTAs per session (`?__clerk_identity=<seat>&redirect_url=<base><result-path>`), the `.actions` two-button
layout + `has_manager_view` omitempty; per-product app-base routing generalizing the `is_hiring`/`hiring_base` switch
(next-web :3000, apps/hiring :3001, academy :3077); mint/resolve per-session player seats via `roster.go` + Clerkenstein.
**Out:** any platform/next-web edit; making a runtime-computed result page render (M231's demo-patch/escalation decision).
**Depends on:** M233.  **Complexity:** medium.
**Delivers → `corpus/ops/demo/content-stories-spec.md`** (the cockpit-UX half: tabbed model, two-action contract, icon map, base routing).
**Open questions:** does the academy section deep-link to a content page (post-M230), map onto a skillpath session, or render
presence-only? confirm the per-(simId,userId) manager drill-down deep-link (M224 deferred it as "optional polish").

#### M235 — prove-it-lands: interesting sessions, player + manager  (`iterative`, large)
**Status:** `done` (closed-incomplete 2026-07-20 — pragmatic-close mandate; the LIVE `(session × action)`-lands gate routes to M236 by design)
**Goal:** Populate the tab with INTERESTING (not boring) real-shaped sessions per the brief and prove every cockpit action
lands on a non-empty, believable result page.

**Closure narrative (2026-07-20).** Iterative, **closed-incomplete** under the user's pragmatic-close mandate ("build
non-sim seeders, then close"). Everything the live proof *depends on* is built + unit-proven, **0 platform-repo edits**,
all in rext `stack-seeding` + the rosetta corpus. **Two user-blockers surfaced + resolved:** (1) **M235-01** — the
anonymization scrub removed **zero** names (8/9 fixtures leaked a real first name) because the capture sourced only the
empty `jobsimulation.actors` names, not the session owner's `public.users` identity where the candidate's first name
actually lives → hardened (owner-identity sourcing → `<<ACTOR_0>>`, token-split, word-boundary, a capture-time
`SurvivingToken` fail-closed post-condition) + re-captured 9 fixtures **provably clean** (0 leaked names, 545
placeholders) + a standing CI cleanliness tripwire (`#M235-B1`); (2) **M235-02** — the planned "coverage descriptor"
mechanism doesn't exist (dynamic-URL, cockpit-seat-reached result pages need NEW seat-login sweep plumbing authored +
calibrated against a live render) → user ruled "build non-sim seeders, then close" (`#M235-B2`). **Delivered:** the full
13-session simulation matrix (assessment PASSED = 2 voice / 1 code / 1 document; every type passed AND not-passed) + all
3 non-simulation sections (skill-path-legacy real progress + `local_skill_path_sessions` mirror; ai-labs presence-only;
academy `/library/<slug>` CTA) via a separate code-owned registry (`seeders/content_nonsim.go`) → the manifest projects
all **4 products / 18 sessions**, both honesty gates GREEN. rext code-of-record `playbill-m235-hardened @ 60eff14` (build
tags `-scrub-fix` / `-fixture-matrix` / `-nonsim-{skillpath,ailabs,academy}`). Whole-rext Go test funcs **1939 → 1974**
(+35, `git grep '^func Test'`); touched-pkg suites + honesty gates GREEN, `go vet` clean, **flake 5/5**; harden Pass 1+2
`--final` stabilized (0 bugs). Close near-clean (adversarial subsection + 2 back-ref tags, no rext code change), deferral
audit **YELLOW / 0 blockers** (chronic 14-fail demo-stack carry — M235's slice 6 `test_cockpit.py` — user-dispositioned →
v2.5 release-close; not re-woken). **carry-forward.md:** 3 clusters (LIVE proof + new seat-login plumbing · per-section
live-calibration checklists · M230 carry-forward live items), **all Fate-3 → M236, already applied** to M236's `In:`
(iter-08, commit `54eaefe`, user-authorized). No live proof faked; no platform edit.
**Exit gate:** on a cold reset-to-seed, every in-scope (session × action) logs in on the correct org and lands on a NON-EMPTY
result page for BOTH player and manager vantages, 0 ejects, with the assessment **2-voice / 1-code / 1-document PASSED** set
present and each type present in **passed AND not-passed** states; each product either passes or is declared with a documented
fate (AI-labs feasibility answered explicitly).
**Iteration protocol:** `corpus/ops/demo/playthroughs.md` + `coverage-protocol.md` — a Playthrough per (session × action) +
a coverage descriptor asserting non-zero rendered values (turns a blank clone RED — the M219/M222 mirror-table-vs-base-session
trap). Triage each blank landing to its true read-model; fix in seeder/manifest or route to a demo-patch / escalate.
**Out:** live-on-billion proof (M236); products M231 ruled out.
**Depends on:** M234 (+ M230 for the academy section).  **Complexity:** large.
**Delivers →** none (proof milestone; extends the coverage/playthrough manifests).
**Open questions:** if `/sim/.../result/<sessionId>` is runtime-blank, is landing "as player" on the seedable
`/profile/activities`|`/profile/skills` composed outcome acceptable, or is a demo-patch authorized? does not-passed render a
meaningful result page or blocked/empty?

#### M236 — prove-on-billion  (`iterative`, medium)
**Status:** `done` — **gate MET + MERGED 2026-07-20** (closed-on-gate; the merge block cleared by the close
continuation — user fate on the standing test-debt carry = **RE-BASELINE now, decide at release close**. The
carried "14, 6 of them pin drift" reproduces at 14 but falls to **8** on a clean stable-`main` clone set, with
**0** real defects and **0** pin drift — see `m236-prove-on-billion/rebaseline-standing-failures.md`)
**Goal:** Re-prove the whole feature live on the `billion` Tailscale VM (the house pattern that closed M215/M221/M226/M228) —
both cockpit tabs usable end-to-end from a 2nd machine on a cold reset-to-seed.
**Exit gate (as re-scoped by USER-BLOCKER-M236-01):** both tabs work live on `billion` — all **landable** (session × action)
pairs render real content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a
cold reset-to-seed, p95 click→ACCESS < 5 s **for the HERO vantages only**, 0 platform edits.
**Iteration protocol:** `corpus/ops/demo/coverage-protocol.md` + `playthroughs.md` (repointed from the hollow `verification.md`
ref — B5). Same billion-safety rules (one driver, no detached on-host scripts, assert from a tailnet peer, never kill a mid-build).
**Out:** new feature work (built by M235); content-seat latency (B2, out of scope for v2.5).
**Depends on:** M235.  **Complexity:** medium.
**Delivers →** none.

**Closure narrative (2026-07-20).** Iterative, **closed-on-gate** in **10 iters** (1 bootstrap tok + 9 tiks, single day).
**Gate MET cold on `billion`:** **29/29** landable (session × action) pairs render real content both vantages · **65** academy
course cards / 483 chapter links / **0** Draft chips · hero p95 **1.22 s** employee / **1.51 s** manager vs a 5 s budget, 5/5
ACCESS · reproduced on a cold reset-to-seed **with no intervention** · **0 platform-repo edits**, verified per-clone.

**The denominator was CORRECTED 31 → 29 mid-milestone (iter-07) — the target SHRANK, and this is not 31/31.** The 2
skill-path **manager** pairs point at a surface next-web **has not built** (`InsightsBySkillPathStudentSimulationsContainer`
renders the literal string "Coming soon", results table commented out, `userData` hardcoded null — no query touches the
seeded session), so under M233's fail-closed rule they are **not landable**, on the same ground that excludes AI-labs. *31 was
never a count of provable pairs.* The correction is argued inline in `overview.md` with the 31 struck through, not rewritten.
It also exposed a **false PASS**: the lighter of the two had been scoring green off a definition-only "Results for" header
(chrome served by a different query than the one that failed) — so the pre-correction reading was wrong in **both** directions
at once. Chain: `18 sessions + 15 manager views = 33 raw` → −2 skill-path → `31` → −2 ai-labs → **29 landable**.

**USER-BLOCKER-M236-01 (5 sub-findings, all user-resolved 2026-07-20):** the Phase-0b KB-fidelity audit returned **RED** on
spec grounds — the declared gate contained an unprovable clause ("tailnet-only" is false by construction: every demo publishes
on `0.0.0.0`), half the gate was unmeasurable (the content CTA emits no `data-login-as`, which *is* the ACCESS predicate), the
cited page-object did not exist (it is a next-web `.tsx` component, so the harness had to be **authored from scratch**), the
milestone had to consciously **reverse** a documented `skipPaths` rule, and its declared `iteration_protocol_ref` was hollow.
Resolved B1 drop-the-clause / B2 hero-only p95 / B3 accept the enlarged cluster / B4 amend the protocol / B5 repoint the refs.

**The milestone's most transferable finding — five wrong test assertions per one real product bug.** Of the defects that cost
iters, the majority were *the test being wrong*, not the product: a manager test that asserted the defective contract (which is
why it shipped), an interview manager view graded as a false FAIL against the wrong shape, a skill-path page graded as a scored
sim, an academy CTA whose unit test *required* a route that does not exist and so defended it. The final harden then found
**three more checks passing against a broken subject**: an aggregator reporting success on an empty run (0/0 is also
arithmetically 100%), the whole e2e suite passing by **collecting 0 tests** after a module-scope throw — silently taking **61
tests offline for 8 iters** — and a grader with **no negative tests at all**. Backfilled into `coverage-protocol.md`,
`latency-budget.md`, `content-stories-spec.md`. The rule: *ask of every layer that reports a number what it prints when
nothing happened.*

**Close (2026-07-20).** Review found 17 doc + 16 code + 6 test-coverage findings; **all fixed**. Notably: three docs still
asserted the skill-path manager surface exists (the claim that produced the 31) and two still routed the academy through
`app/cmd/academy-seed`, which iter-08 proved **moot on a demo** (no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` ⇒ the seeded rows have
no reader); `CLAUDE.md` asserted the manager route takes a `<userId>` when it takes a **membership** id — the exact defect
iter-05 spent an iter on. A full-suite sweep found the standing carry measured **19, not the briefed 14**: five *unnamed* stack-core failures
were the two cross-repo **doc-truth guards** (org-count: the preset has shipped **4** orgs since v2.4 M223 while docs, source,
and the guard's own test still said 3; and `DEMO_NO_ACADEMY_FILL` — the knob that **gates Thread A** — undocumented). Both
guards were red and **correct**, and had been read as noise for three milestones. Fixed → stack-core 5 → 0. Also centralized
the membership-key derivation (a bare literal at **9** sites, one of which writes the row and eight of which merely hope to
match it) after finding M236's own regression test for that defect was a **self-consistent tautology** — it derived the
expected value from the expression under test, so it could not fail. Both new pins mutation-verified. rext code-of-record
**`playbill-m236-close-fixes`**, pushed. Go **1974 → 1976**; stack-verify python **132 → 141**; harness specs **64 → 66**.

**The close blocker that held the merge — DISCHARGED 2026-07-20.** `/developer-kit:audit-deferrals` returned **RED**. The
standing pre-existing demo-stack failures were a genuine repeat-deferral across **10 milestones and 2 releases**, and their
declared destination — *the v2.4 release close* — **had already fired once without landing them** (v2.4 shipped them as a known
issue and re-anchored on v2.5), an **AGED_OUT** trigger no audit had recorded. M236 being the FINAL v2.5 milestone, there was
nowhere left to defer to. The set had also drifted under a fixed label (8 → 14) with the stated *class* changing from
stale-tests to `pre_sha256` pin drift, so the label was wrong in both directions. **User fate (2026-07-20): RE-BASELINE now,
decide at the release close** — executed, and the merge released.

**The authoritative count is 8 on macOS · 7 expected on Linux** — the clean-clone reading, re-measured at the v2.5 release
close, **0 real defects, 0 pin drift** (that diagnosis is **REFUTED**; its implied remedy — re-anchoring the "drifted" pins —
would have re-pinned five manifests to *patched* content and permanently disarmed the drift detector). **The count is
host-dependent: always state the host OS or it drifts again for exactly the reason it drifted the first time.** The other two
figures are audit trail, not measurements: **14** is the DIRTY-clone reading (6 of them were leftover applied demo patches
reporting themselves as test failures — they did not reproduce at the close, which independently confirms the `stack-demo`
clone set is pristine), and **19** folded in the 5 stack-core doc-truth-guard failures that were FIXED at the M236 close.
Sources: `m236-prove-on-billion/rebaseline-standing-failures.md` · `decisions.md` **CLOSE-D2** · the release `metrics.json`
→ `standing_failures`.

**Open questions:** none blocking.

### On the reserved M205 (updated 2026-07-19)
v2.4 discharged the recruiter/seeder half of vision **M205 "Hiring + tier gates."** v2.5 does NOT touch M205's residual
half (Stripe tier gates + ATS candidate-pipeline) — those stay a vision reservation. v2.5's Content-stories is a **NET-NEW**
content-vantage pillar, not part of M205.

## Done — shipped releases (v1.10b, v2.0 → v2.4) → **moved**

The full `## Done` sections for **v2.4 "casting call"** · **v2.3 "cue to cue"** · **v2.2 "panorama"** ·
**v2.1 "quick change"** · **v2.0 "opening night"** — and the interposed **v1.10b "fit-up"** backfill — now
live in **[`roadmap-archive-v2.0-v2.4.md`](roadmap-archive-v2.0-v2.4.md)**, split out at the v2.5 close
(2026-07-20, finding `KB-C`) under the `roadmap-legacy.md` precedent. This file had reached 2,079 lines /
203 KB, ~77% of it shipped history.

The one-paragraph-per-release summaries stay below in **§ Shipped releases** — that section is the index,
the archive is the detail. The retired v1.x major (M0 … M46) is in
[`roadmap-legacy.md`](roadmap-legacy.md).

---

## Shipped releases

- **v2.3 "cue to cue"** — **2026-07-15**, tag `v2.3`, **5 milestones (M217 → { M218 ∥ M219 ∥ M220 } → M221)**. The
  **presenter-speed** / field-hardening release: a presenter swaps heroes in **under 5 s** on a demo that comes up
  green, fully-loaded, and remotely reachable by default. Headline **click→ACCESS < 5 s** gate set at M218 and
  **re-proven live 8/8 on `billion`** over the tailnet, no flags (2.11 s / 1.31 s vs a ~39/38 s baseline, ~18×);
  remote default-on for demo; `safety.md` Part 3 (exposure axis); the ~24-instance **D17** status-artifact thread
  told honestly. **4 non-gate tail carries → v2.4.** Tooling + docs only, zero platform edits, 0 net-new direct
  deps (one indirect `x/crypto` patch). rext code-of-record `cue-to-cue-m221-final`; the `billion` demo LEFT LIVE.
  Records archived under [`releases/archive/02.30-cue-to-cue/`](releases/archive/02.30-cue-to-cue/).
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

_Last updated: 2026-07-20 (**v2.6 "sound check" DESIGNED + PROMOTED to active development** via
`/developer-kit:design-roadmap` — the **reliability / field-hardening release** [the v1.3b / v1.10b / v2.1 / v2.3
lineage]: *make everything that's built actually get built + provisioned.* **8 milestones M237 → M244**, barrier →
parallel fixes → prove-on-billion; branch `release/02.60-sound-check` cut from **local** `main`; tag `v2.6`; realizes
reserved M237→M244 + M238→M243; tooling + docs only, zero platform-repo edits. See § Active — v2.6. Prior:
2026-06-29 (**v1.10b "fit-up" DESIGNED + PROMOTED** via `/developer-kit:design-roadmap` — an
interposed **field-hardening backfill** [the v1.3b "dress rehearsal" lineage]; **7 milestones M47 → { M48 ∥ M49 } →
M50 → M51 → M52 → M53** re-opening the v1.x flat counter; branch `release/01.10b-fit-up` cut from `main`; tag
`v1.10.1`. Designed from the field review `.agentspace/annotation.md` + the M201 stale-clone finding [3 research
agents]. Re-grounds demo + corpus to current prod, fixes the from-scratch `/demo-up` issues + the v1.10 content
gaps, adds the AI-readiness showcase org, consolidates one auditable seed+gen manifest. **v2.0 "opening night"
PAUSED** until it ships. Tooling + docs only — zero platform-repo edits. Prior: 2026-06-28 **v2.0 "opening night"
DESIGNED + PROMOTED** — a NEW MAJOR opening the **Playthroughs** pillar; 4 milestones M201 ∥ M202 → { M203 ∥ M204 };
branch `release/02.00-opening-night`; from `spec-drafts/playthroughs/spec.md` v0.3.)_
