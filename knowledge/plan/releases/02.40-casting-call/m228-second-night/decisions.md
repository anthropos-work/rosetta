# M228 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions). The milestone-root strategy chain
lives as `## TOK-NN:` entries below; intra-iter decisions live in `iter-NN/decisions.md`._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## TOK-01: reprove-corrected-hiring-on-billion — 2026-07-17

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Re-prove the retuned 7-condition hiring gate + the 4 M227-correction render checks on billion by
**cutting billion's rext consumption clone from the M226 code-of-record (`casting-call-m226-c2-race-fix`) over to the
M227 corrections (`casting-call-m227-sections`), then running a default cold `/demo-up` and measuring from this Mac.**
The billion INFRA is fully proven at M226 (the serve-hiring front, the 4 hiring demo-patches, the recruiter vantage,
the count 5+45, the C2 harness race-fix — all folded into the M226 code-of-record, which `m227-sections` is built on
top of). **Only the DATA changed (M227, pure seed/content tooling — no new demo-patch, no platform image change).** The
first batch of tiks executes:

1. **Cold teardown of the M226 demo.** billion runs the M226 `demo-1` (17 containers Up 5h at `casting-call-m226-c2-race-fix`,
   devops-operated, 7 tailscale-serve fronts). Tear it down cleanly via the consumed tooling's teardown
   (`rosetta-demo down 1 --purge`) — incl. the serve reset + the ant-academy respawner reap (M221 F5/F5b/F12).
   Verify FROM THIS MAC that base+offset ports refuse + no survivor process/serve-front remains (M217 orphan lesson).
2. **Substrate cutover (rext-only).** Update billion's `/home/devops/panorama/stack-demo/rosetta-extensions` from
   `casting-call-m226-c2-race-fix` → **`casting-call-m227-sections`** (fetch + checkout + `.agentspace/rext.tag` SoT).
   No platform-clone changes needed (M227 touched no platform image / demo-patch). The 20 demo-1-* images stay cached.
3. **Default cold bring-up.** Run a **default `up-injected.sh 1` (NO FLAGS)** as `devops` on billion — remote reach is
   default-on (M220 D-DESIGN-3), auto-discovers `billion.taildc510.ts.net`. Run **synchronously** in a blocking ssh and
   WAIT (NEVER detach — a leaked build strands the demo-patches, M221/M217). **NEVER `docker builder prune`** (the M227
   LOCAL disaster: it evicted the go-build cache → 35-min cold compiles). Heartbeat by file-write/image activity.
4. **Measure the retuned 7-condition gate + 4 M227-correction render checks FROM THIS MAC** (the tailnet peer) — never
   on-host (SSL-artifact false-REDs). Conditions: (1) hiring org + is_hiring + **exactly 5 mgr/45 cand**; (2) recruiter
   comparison **≥ retuned floor 6 per each of 5 positions, each candidate on exactly 1 sim (~8/position)** (via
   `run-hiring-render.sh` COVERAGE_RENDER_GATE=1, `RENDER_GATE_FLOOR=6` default); (3) 2 candidate profiles usable +
   **external emails + gender-matched avatars visible**; (4) reads as hiring + **hiring-only content** (only the 5
   HIRING sims in the recruiter list, no training/assessment leakage); (5) **recruiter p95 click→ACCESS < 5 s** (via
   `run-latency.sh 1 recruiter`, gated on fresh-green autoverify.json); (6) coexists with 3 workforce orgs; (7) 0
   platform-repo edits. **The 4 M227-correction render checks (1-sim/candidate ~8/position, external emails, matched
   avatars, hiring-only) are the NEW live observations — never live-rendered before (M227's local re-prove was
   env-blocked).**
5. **Attribute before fixing.** For any failing condition, name the surface (a re-surfaced render gate R1, R4
   hydration, a Clerkenstein/seed-wiring gap, an OOM stall) using the protocol's diagnostic discipline (arithmetic
   signature per `latency-budget.md`; state the environment with every number). Fix via **tooling / a sha-pinned
   demo-patch re-proven live at final code**, tagged `casting-call-m228-*`. **A platform-repo edit is NEVER in bounds**
   — an un-patchable surface ESCALATES.
6. **Reproducibility.** Prove the gate on **2 clean default cold reset-to-seed cycles** (the M221/M226 "one cycle is
   provisional; a 2nd confirms reproducibly" precedent) before declaring gate-met.

**Rationale:** This is the M226 prove-on-billion recipe RE-RUN with the M227 corrections. The infra is proven; the only
new information is whether the corrected DATA renders right on the live cross-machine cold run — and specifically the 4
believability render checks that were only ever proven DETERMINISTICALLY (Go unit tests), never live-rendered (M227's
local re-prove was env-blocked → routed here). The strategy front-loads the rext-only cutover (cheap — M227 changed no
image/patch) so the first tik goes straight to a default cold bring-up + the first measurement. Orders measurement
before fixing so gaps are attributed to their surface, not guessed. Sharpest risks: R1 (a render check invisible
deterministically breaks live) + the billion 7.3 GiB memory floor under the 2-app demo (M226 held).

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** Gate = the retuned 7-condition live billion proof + the 4 M227-correction render checks,
reproducible on a cold reset-to-seed. **Starting value: 0/7 proven on billion at `casting-call-m227-sections`** — the
demo currently up is the M226 PRE-correction substrate (`casting-call-m226-c2-race-fix`: each candidate on all 5
positions ~43/position, org-domain emails, gender-blind avatars, non-hiring sim leakage). M228 proves the corrected
data (~8/position, external emails, matched avatars, hiring-only) live.

**Next-tik direction:** iter-02 (tik) executes steps 1–4: teardown the M226 demo (`rosetta-demo down 1 --purge` as
devops) + verify clean from this Mac, cut billion's rext → `casting-call-m227-sections`, run a default cold
`up-injected.sh 1` synchronously, then take the FIRST retuned-7-condition + 4-M227-render-check measurement from this
Mac. Attribute every failing condition to its surface before any fix. Handler: `PROVE-M228-iter02-first-corrected-cold-bringup`.

## PRAGMATIC-CLOSE-MANDATE — 2026-07-18 — close on iter-03's single cold cycle

**Decision (user, explicit via AskUserQuestion):** Close M228 on iter-03's cold-cycle 7/7 proof. Do NOT require
a 2nd independent cold reset-to-seed (iter-04) for the gate's "reproducibly" clause.

**Rationale.** All 7 gate conditions are MET on the cold-seed billion instance (tag
`casting-call-m228-hiring-scope-fix`), and all four M227 believability corrections are confirmed rendering live
(hiring-only 5 sims · external emails · 1-sim/candidate 8,8,9,9,8 · gender-matched avatars). The tik/tok loop
ALREADY exercised reproducibility in the strongest sense: iter-02's cold bring-up BROKE (F1 FK-crash + F2/F3
leak), iter-03 FIXED it and the fix was re-verified on a fresh cold bring-up + the full 7/7 UI re-measure. A 2nd
identical cold cycle (iter-04) would need a bring-up ON the billion VM, which the driver cannot trigger (no ssh
to the host; `rosetta-demo` reset-to-seed is host-locked there) — and the fixes are conclusively proven.

**Gate status:** MET (7/7 on 1 cold cycle; reproducibility accepted on tik/tok evidence). No `carry-forward.md`
required — this is a shipped-on-gate close, not a closed-incomplete close. iter-04 is NOT scheduled.

**Handler:** `PROVE-M228-close-on-iter03-cold-cycle`.
