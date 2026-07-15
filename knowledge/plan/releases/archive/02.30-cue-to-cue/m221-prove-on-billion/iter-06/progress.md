# iter-06 (tik) — the FINAL live demo cycle on `billion`, LEFT RUNNING

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). The confirming r4 cycle iter-05 said
remained — run ONCE, grade thoroughly (incl. the BROWSER grade of Dana), and **leave the stack UP** as the
final live demo (the user's explicit ask: ONE cycle, not a battery; a working live demo left on the box).

## The cycle — a DEFAULT cold `up-injected.sh 1`, NO FLAGS, at r4

- **rext pin:** `cue-to-cue-m221-r4` verified on `billion` (clone checkout **+** `.agentspace/rext.tag`, both) —
  graded == shipped.
- **Cold proof:** T0 `2026-07-15T06:55:51Z` < PG_VERSION mtime **from inside the container**
  `2026-07-15T07:00:20Z` ⇒ initdb re-ran ⇒ genuinely cold reset-to-seed.
- **Green proof:** `autoverify.json` `green:true, warnings:0, ts=2026-07-15T07:02:39Z` (THIS run). UP_INJECTED
  **EXIT=0**. Seed audit: 49 writes / 71,783 rows / prod=false / **isolation CLEAN**. 16 containers up.
- **Driven + graded from the tailnet peer** (this Mac), never on-host. A fast ~7-min cycle (warm Docker layers
  + warm snapshot cache; cold DATA).

## Close — 2026-07-15

**Outcome:** A DEFAULT cold no-flag cycle at r4 came up GREEN and **8/8 gate conditions MET**, browser-graded
from the peer (incl. Dana's `/ai-readiness` 900-char floor the prior cycle only DB-confirmed). **The F1 cascade
is fixed live** (taxonomy 0→42,790; personas/membership-skills/target-roles/ai-readiness-funnel all populate;
all 3 catalog surfaces replayed — directus self-unblocked via M21 auto-provision exactly as D-05e predicted).
The stack is **LEFT LIVE** as the final demo.

### The 8-condition gate — FINAL verdict

| # | Condition | Verdict | Evidence |
|---|-----------|---------|----------|
| 1 | p95 click→ACCESS < 5 s, BOTH heroes, over the tailnet | **MET** | maya-thriving (employee→/profile) **p95 2.11 s** (p50 0.82); dan-manager (manager→/enterprise/workforce) **p95 1.31 s** (p50 0.82); **ACCESS 5/5 each**; tailnet HTTPS origin. (Reused the M218 `stack-verify/e2e` latency spec — see F-M221-06b for the http-scheme wrapper caveat.) |
| 2 | Full catalog — taxonomy + directus + sim-embeddings, no skipped surface | **MET** | setdress "taxonomy=replayed directus=replayed sim-embeddings=replayed" (none skipped); `public.skills=42,790`, `job_roles=22,470`, `directus_collections=21`, `skill_embeddings=42,790`, `job_role_embeddings=18,919`. |
| 3 | 3 orgs incl. AI-readiness | **MET** | Cervato Systems / Solvantis / Northwind Aviation; `ai_readiness_cycles=2` (Northwind). |
| 4 | Dana's FILLED AI-readiness page — **BROWSER-graded** | **MET** | `/ai-readiness` from the peer: **YOUR AI READINESS 66/100**, 4-archetype matrix (Champion/Hidden-Talent/Standby/Explorer), 13-team breakdown, Steps Completion 156/200; **all 3 tabs filled** (Snapshot 1,745 ch · How-we-measure 1,629 ch · What-to-do-next 2,305 ch — each **>> the 900 floor**); interview dimension surfaced (Step 3 AI Interview 78% · 78 interview-hrs); `interview_aggregated_reports=1`, `ai_readiness_snapshots=199`, `user_step_progresses=532` (DB). |
| 5 | Ben STARTED visible on his dashboard (`/home`) | **MET** | browser: "Discover your AI Readiness — Due Aug 14", **Skill Mapping DONE / AI Simulation START / AI Interview LOCKED** — the STARTED workflow; DB: only `skill_mapping` completed. |
| 6 | Aria COMPLETED renders | **MET** | browser `/home`: "Your AI Readiness: **AI Champion · Completed Jul 9 · 89/100**", all 3 steps scored (30/30 · 33/40 · 26/30); DB: all 3 steps completed. |
| 7 | Remote-by-default — no flag, tailnet host auto-discovered | **MET** | log "public-host **AUTO-DISCOVERED** billion.taildc510.ts.net (all 6 tailscale rungs · cert MINTED)", **no flag passed**; `tailscale serve` fronting **6 HTTPS surfaces** (13000 next-web / 13077 academy / 15050 / 17700 cockpit / 18082 backend / 19000 studio); reachable from the peer. |
| 8 | Zero platform-repo edits | **MET** (live-nuance disclosed) | 12/13 clones CLEAN incl. **next-web-app CLEAN** (all 6 demopatches reverted after the image baked); `cms ?? studio/` = the disclosed pre-existing embedded repo (D-05h). **ant-academy carries 3 tracked files** (`next.config.js` dev-origins + predev-regen `catalog.json`/`content/index.md`) — the LIVE **native** academy's ephemeral run-state, reverted-clean on `--stop` (F5b, field-proven iter-05). No durable platform SOURCE edit; the canonical `anthropos-work` repos are untouched. Because the demo is **LEFT LIVE** per the user's ask, the native academy necessarily runs patched to accept the tailnet origin. |

**Scoreboard: 8/8 MET.**

### M219 readiness fold-in — all MET

- **No junk skills:** 132 distinct claimed skill names org-wide, **0 orphans, 0 junk** (all real taxonomy names).
- **Hero role titles resolve:** every hero's `job_role_name` resolves to a real taxonomy role with a valid `J-*`
  node-id (maya=DevOps Engineer, dan/dana=Engineering Manager, aria=Data Analyst, ben=Business Operations
  Analyst, sara/nick=Account Executive, leah=Sales Manager, tom=Backend Developer) — browser-confirmed rendering.
  (`job_role_title` denorm cache is null 540/541 — a disclosed data-detail the UI doesn't depend on, not a
  resolution failure.)
- **`ai_readiness_cycles == 2`:** MET (both Northwind).
- **`interview_aggregated_reports` non-empty:** MET (=1).
- **Readiness sections filled (900-char floor):** MET (Dana browser, all tabs >> 900).
- **Frozen and live cycles agree:** MET — frozen (closed) score **62** vs live (active Q3) **66**, both mid-60s,
  AI-Champion-dominant (the single-round score fix in r4 holds).

### F10 — fully field-exercised (closes the iter-04 residual)

- **`assert_ports_free`:** occupied `:13000` → correctly "ALREADY IN USE" + rc=1 (the tri-state MF-2 path — a
  root-owned docker-proxy listener is NOT false-reported free); free ports → rc=0.
- **demopatch freshness gate:** live clone `status=pristine` (reads real state); an isolated `/tmp` sandbox with a
  deliberately broken anchor → `status=drifted` → `apply` fired **"G2 REFUSE: the manifest anchor was NOT FOUND
  … a REAL semantic break"** — the abort path iter-04 never forced. (apply+self-heal was already field-proven by
  this cycle's bring-up: "demo-patches: all applied (none refused, none skipped)".)

### Known-gap + finding (routed forward — Fate-3)

- **F4 (academy empty grid):** CONFIRMED unchanged — `#courseGrid` renders **0 course cards** + "No adventures"
  empty-state (bodyLen 348) despite `catalog.json` serving **2,705** entries (HTTP 200). A client-side render
  defect in the **ant-academy** repo; **not a gate condition**, not demopatch-shaped without render-path
  investigation. Recorded as a known cosmetic gap; did not block leaving the demo live.
- **F-M221-06b (harness):** `run-latency.sh` hardcodes `http://${HOST}:${COCKPIT_PORT}` for the cockpit URL,
  which BREAKS against the M220 cockpit-HTTPS-fronting flip on a remote demo (`http://…:17700` → 400 "HTTP
  request to an HTTPS server"). Worked around by driving the SAME spec (`tests/latency.spec.ts` + `lib/latency.ts`)
  directly with `https://…:17700` (green gate verified independently) — a fix, not a fork. **Route a rext
  scheme-fix forward** (make the wrapper use `https` when the demo is public-host fronted).

**Type:** tik
**Status:** closed-fixed
**Gate:** **MET** — 8/8 conditions hold on this cold r4 cycle, browser-graded from the tailnet peer (incl. the
Dana 900-char floor). Together with iter-05's prior r4 cycle this satisfies the gate's "reproducibly on a cold
reset-to-seed" pragmatically (the user's ONE-cycle mandate for this final live demo).
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (tik #5 of 5) — (6) protocol-stop: n — **Outcome: GATE MET → the milestone's exit gate is
reached; the stack is LEFT LIVE as the final demo.**
**Decisions:** D-M221-06a..06e (iter-06/decisions.md).
**Left LIVE (not torn down):** 16 containers + cockpit(:17700) + academy(:13077) natives + `tailscale serve`
(6 HTTPS surfaces); registry `demo-1=up`; hostlock **released** (the tooling's normal post-up state). Peer-reach
confirmed at close: cockpit HTTP 200 (all hero CTAs), app 307 (unauth→handshake, correct), backend `/api/health`
200; the maya/dan/dana/ben/aria hero logins all drove end-to-end. **No background tasks armed** (the ONLY thing
running is the demo stack).
  - **Live URLs:** presenter cockpit **https://billion.taildc510.ts.net:17700** (pick a hero → [Log in as]) ·
    app **https://billion.taildc510.ts.net:13000** · academy https://billion.taildc510.ts.net:13077.

## Ledger
- iter-06 (tik): FINAL cold no-flag r4 cycle on billion — **8/8 gate MET** (browser-graded incl. Dana 900-char),
  M219 readiness all MET, F10 fully field-exercised, F4 confirmed known-gap + F-M221-06b harness finding routed
  forward. **Stack LEFT LIVE** (cockpit :17700 / app :13000, peer-reachable). **Gate: MET.**
