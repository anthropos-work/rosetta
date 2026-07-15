# M219 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

---

## Pre-flight audits — S1 (the surface split)

**Phase 0b — KB-fidelity: YELLOW.** Satisfied inline by the M219 four-way census (below) rather than by a
separate `/developer-kit:audit-kb-fidelity` invocation: the census checked **every load-bearing behavioral claim**
in `corpus/services/ai-readiness.md` against the actual `app` + `next-web-app` source AND against a live
authenticated probe of the running `demo-1` on `billion` — a strictly stronger evidence base than a doc-vs-code
read. Recorded as **D-M219-1**. Findings are tracked as `KB-1`..`KB-3` below and are fixed in S1's doc pass.

| id | Claim | Verdict |
|----|-------|---------|
| KB-1 | `ai-readiness.md:162-176` — "the demo FE fires the data GET **WITHOUT** `?cycle=` … and never fires the `/cycles` list that supplies `latestClosedCycle.id`" | **STALE / MISATTRIBUTED.** True of the **legacy** page, false of the **current** one. See F-1/F-2. |
| KB-2 | `ai-readiness.md:159-160` — the live-recompute "**never completes** in the coverage harness budget" | **FALSIFIED on the current stack.** Measured **2.09 s**. See F-3. |
| KB-3 | `ai-readiness.md:99-106` §Surfaces — lists the manager dashboard + member onboarding, but does **not** name their routes and does not record that a **second, legacy** manager surface exists | **INCOMPLETE** — the omission is exactly what let every demo pointer land on the legacy page. See F-1. |

No **RED** (no load-bearing claim was both stale *and* unfixable before coding); the seeder contract at
`:120-192` re-verified ALIGNED, as the overview promised.

---

## The census (4 parallel read-only sweeps, 2026-07-14)

### F-1 — **The current/legacy split is REAL, and every demo pointer targets the LEGACY surface.**

The user's kickoff report ("not the old legacy ones") is **confirmed in code**.

| Vantage | Surface | Route | Status | Evidence |
|---------|---------|-------|--------|----------|
| Manager | `AIReadinessClient` (v3.0: HeroCard + Snapshot / How-we-measure / What-to-do-next tabs, **cycle-aware**) | **`/ai-readiness`** | **CURRENT** | the only route constant (`packages/core-js/src/constants/urls.ts:50` `AI_READINESS_URL`), and the only one the navbar links (`packages/ui/src/NavBar/useNavbarSections.tsx:253-260`); the only one e2e covers (`e2e/specs/web.ai-readiness.spec.ts`) |
| Manager | `AIReadinessContainer` → `AIReadinessView` (pre-v3.0 org-summary + team table, **no cycles, no archetypes, no people**) | `/enterprise/workforce/ai-readiness` | **LEGACY — an unlinked orphan** | **no nav entry, no tab, no redirect, no link anywhere**; `WorkforceNewClient.tsx:125-151` omits it from the workforce tab list; reads the cycle-less `?tag=` endpoint (`hooks/useWorkforceAIReadiness.ts:23-27` — *no `cycle` param exists in the hook*); no feature work since the v3.0 rewrite (`6479ac4c4`, 2026-06-19) |
| Employee | `AIReadinessHero` + `AIReadinessRailCard` | **`/home`** (embedded — **no route of its own**) | **CURRENT** | `HomeLeftContent.tsx:132`, `HomeRightContent.tsx:68` |

**The three demo pointers, all aimed at the legacy page:**

| Pointer | Value today | file:line |
|---------|-------------|-----------|
| Dana's cockpit `jump_to` | `/enterprise/workforce/ai-readiness` | `stack-seeding/presets/stories.seed.yaml:153` |
| The cockpit deep-link catalog's **only** readiness entry | `/enterprise/workforce/ai-readiness` | `stack-seeding/seeders/cockpit.go:87` |
| The coverage sweep's `AI_READINESS_PAGE` | `/enterprise/workforce/ai-readiness` | `stack-verify/e2e/lib/coverage-manifest.ts:468` |

There is **no** `end-user`-vantage readiness entry in the deep-link catalog at all, and **no** employee readiness
section in the coverage manifest. Aria's and Ben's `jump_to` are both the generic `/profile`
(`stories.seed.yaml:130,142`).

### F-2 — **The `CycleID == nil` blocker DISSOLVES. No new demo-patch is needed.** *(overview item 1, 2nd blocker — REFUTED)*

The overview's premise was that the default GET bypasses the frozen seed "unless the frontend passes `?cycle=`".
**The current frontend DOES pass it** — `AIReadinessClient.tsx:137-138`:

```
const effectiveCycleId = selectedCycle ?? activeCycle?.id ?? latestClosedCycle?.id;
```

…gated on `cyclesQ.isFetched` (`:150-154`), i.e. it **waits for `/cycles`** before firing the data GET. Proven live
against `demo-1` on `billion` (authenticated as `dana-manager`): `GET /api/workforce/ai-readiness/cycles` →
`{"cycles":[{"id":"95d9fc3d…","status":"closed",…}]}`. So the current page resolves `latestClosedCycle.id` and takes
the **frozen** path.

The M51 iter-07 observation the corpus records (KB-1) — "the demo FE firing the data GET WITHOUT `?cycle=` … and
never firing the `/cycles` list" — is **exactly the behavior of the LEGACY page**, whose hook has no `cycle` param
and never calls `/cycles`. The demo was pointed at the legacy page; the probe watched the legacy page; the finding
was attributed to the platform. **It is a pointer bug, not a platform gap.**

⇒ **The sanctioned-hatch escalation is withdrawn.** Per `demopatch-spec.md §1`, a demo-patch is the *last* resort;
here the non-patch fix (point the demo at the current surface) is available and correct.

### F-3 — The live recompute is **2.09 s**, not "never completes". *(KB-2 — FALSIFIED)*

Authenticated timing probe, `demo-1` on `billion`, org Northwind (199 members), both `app` demo-patches applied
(the 24 ms frozen read is itself the proof that `app-aireadiness-snapshot-loadmembers` is live — unpatched it was
180 s):

| Path | Request | Result |
|------|---------|--------|
| **LIVE** (`buildLiveResponse`) | `GET /api/workforce/ai-readiness?include_people=true` | **HTTP 200 · 2.09 s · 304 KB** |
| **FROZEN** (`buildResponseFromSnapshots`) | `…?cycle=95d9fc3d…&include_people=true` | **HTTP 200 · 0.024 s · 95 KB** |

### F-4 — **The FROZEN path leaves 6+ sub-sections EMPTY. The LIVE path fills them.** *(overview item 8 — the core finding)*

Filled-ness of the two payloads, measured field-by-field against the section manifest (S1):

| API field | Sections it feeds | LIVE | FROZEN (what the demo serves today) |
|-----------|-------------------|------|--------------------------------------|
| `howWeMeasure.interview` | How-we-measure **B3** (the whole Step-3 block), **B3a** "How they use AI" (4 tiles), **B3b** "What holds them back", **B3c** "Strengths", **B3d** "Unexpected angles" | ✅ present | ❌ **NULL → all 5 sections ABSENT** |
| `people[].diagnosis` | Diagnose drawer: "Blocking analysis", **"Recommended actions"** | ✅ present | ❌ **missing → Recommended actions ABSENT** |
| `people[].sources` | Diagnose drawer: "Assessment sources" (2 cards) | ✅ present | ❌ **missing → grey "not started" cards** |
| `cycle` | HeroCard meta row (cycle dates) | ❌ NULL | ✅ present |
| `org.*`, `byTeam`, `people`, `howWeMeasure.{steps,skillInsights,simulations,cycleTotals}` | everything else | ✅ | ✅ |

Both paths fill: `org.score` (50 / 49), `org.members` 199, `archetypeCounts` (4), `capability/usageDistribution`
(4 each), `byTeam` (13 teams), `people` (199), `howWeMeasure.steps` (3), `.skillInsights` (2), `.simulations` (1).

⚠️ `howWeMeasure.cycleTotals.interviewQuestions = 0` on **both** paths — a genuine zero cell. → S2.

### F-5 — **Ben renders NOTHING. Aria renders only the compact archived card.** *(overview items 2+3 — the employee half of item 8)*

The employee surface is gated on an **active** cycle. Northwind's only cycle is **`closed`**, so:

- `app/internal/workforce/readiness_steps.go:291-313` `queryActiveCycleEndDate` → `StatusEQ(StatusActive)` →
  `IsNotFound` → **`deadline = nil`**.
- `useAIReadiness.ts:48-62` `deriveMode`: `deadlinePassed = deadline ? … : true` ⇒ with a null deadline a
  **fully-done** member is `archived`; a **partially-done** member is `progress`.
- **Aria** (all 3 steps) → `archived` → `AIReadinessHero.tsx:83` returns null (hero hidden); `AIReadinessRailCard.tsx:56`
  renders → she gets **only the compact right-rail mini-card**, not the full done hero (ScoreGauge + 3 RecapRows).
- **Ben** (step 1 only) → `progress` → `AIReadinessHero.tsx:88` **`if (!air.deadline) return null;`** → hero hidden;
  RailCard requires `archived` → hidden. ⇒ **Ben has ZERO AI-readiness surface.** The "STARTED" hero — the entire
  point of the persona — is invisible.

**This is a SEEDER finding, not a render bug.** The platform is right: a member cannot be *mid-funnel* in a *closed*
cycle. The M51 design chose `closed` for the manager's sake and Ben was collateral.

### F-6 — The resolution: **seed an ACTIVE cycle alongside the closed one.**

Both vantages want opposite cycle states, and F-3 is what makes the conflict resolvable:

| | closed-only (today) | **+ ACTIVE cycle (S2)** |
|---|---|---|
| Dana (`/ai-readiness`) | frozen, 24 ms, **6 sections empty** (F-4) | live, **2.09 s**, **all sections filled** + HeroCard cycle dates (`resp.Cycle` is set on the active branch, `ai_readiness.go:298`) |
| Aria | compact archived rail-card | **full `done` hero** (ScoreGauge + 3 RecapRows + View details) |
| Ben | **nothing** | **full `progress` hero** (DueDate + 3 FunnelSteps: 1 done / 2 active / 3 locked) |
| The closed cycle | the only cycle | **retained** — a real cycle *history* in the CyclePill; the 199 frozen rows keep earning their keep |

Cost: the manager dashboard's data-load goes 24 ms → 2.09 s. Per **D-DESIGN-1** that load is **reported, not gated**
(M218 owns login speed, not grid speed), and 2.09 s for a 200-member org analytics recompute is honest. **Reported,
not hidden.** `ai_readiness_cycles` has a partial unique index of **one active cycle per org**, so 1 active + 1
closed is legal; `queryActiveCycleEndDate`'s `Only(ctx)` stays satisfied.

### F-7 — `app-aireadiness-snapshot-loadmembers` is **NOT dead**. *(overview item 1, 1st blocker — REFUTED)*

Measured, not assumed:

```
stack-demo/app @ v1.334.1  sha256 = b32169682a28…0ac4d29f  == manifest pre_sha256  → PRISTINE_AT_PIN, rc=0
stack-dev/app  @ v1.335.0  sha256 = dc9e167eda1a…eb7a333b  != manifest pre_sha256  → DRIFTED_ANCHOR_OK (self-heals), rc=0
anchor occurrences: exactly 1 in both clones
```

M217's close made the gate **self-healing** (anchor is the contract; the whole-file sha is only a baseline), so the
patch applies on both boxes. **No re-pin is required.** Its manifest *header comment* is stale (still claims the pin
is v1.315.0) → a comment fix in S2. Corroborated live: the frozen read answers in **24 ms** (unpatched: 180 s).

### F-8 — The `stories.seed.yaml` ACTIVE-vs-CLOSED comment is stale, as the overview said.
`stack-seeding/presets/stories.seed.yaml:112,117` still asserts "the cycle is ACTIVE … which is why the cycle is
active not closed"; the code writes `status='closed'` (`ai_readiness_config.go:98,143`). S2 makes the comment true
again — by seeding the active cycle the comment always described.

---

## Measured numbers (billion, demo-1, offset 10000)

| Metric | Value | How |
|--------|-------|-----|
| Northwind org id | `d7bb7482-28d5-50c5-b9f8-c1544975995b` | demo DB |
| Cycle | 1 × `closed`, `final_score=62`, `end_date=2026-07-12` | `ai_readiness_cycles` |
| Frozen snapshots / live snapshots | **199 / 200** | `ai_readiness_snapshots` / `_live_snapshots` |
| Skills / sims / steps / step-progress | **8 / 2 / 3 / 532** | the `ai_readiness_*` tables |
| Narratives / recommendations / translations | **0 / 0 / 0** | ← `ai_readiness_recommendations` is **absent from the corpus doc's data model** (S1 doc fix) |
| Dana's JWT `org_eid` | `d7bb7482-…` — **the real Northwind UUID** | FAPI handshake → the **FAPI is correct**; only the **BAPI** fabricates (F-11) |

---

## LIVE VERIFICATION of the ACTIVE-cycle fix (billion, demo-1, 2026-07-14)

The active cycle was applied to the running demo exactly as the M219 seeder now writes it
(`participants_filter` left to its column default — confirmed `{"all": true}`, so every member is in audience),
then both vantages were re-probed authenticated. **Measured, not assumed:**

| Probe | Before (closed-only) | **After (+ active cycle)** |
|-------|----------------------|---------------------------|
| **Ben** `aiReadinessUserPlanProgress` | `deadline: null` ⇒ `AIReadinessHero.tsx:88` returns null ⇒ **NOTHING RENDERS** | `deadline: 2026-08-13`, `completedSteps: 1`, `currentStep: 2`, `done: false` ⇒ mode **`progress`** ⇒ **the funnel renders** |
| **Aria** `aiReadinessUserPlanProgress` | `deadline: null` ⇒ mode `archived` ⇒ only the compact rail-card | `deadline: 2026-08-13`, `completedSteps: 3`, `done: true` ⇒ mode **`done`** ⇒ **the full hero** (ScoreGauge + 3 recap rows) |
| **Dana** `/cycles` | 1 (closed) | **2** — `active` "…— Q3" + `closed` ⇒ the FE's `activeCycle?.id` resolves ⇒ `?cycle=<active>` |
| **Dana** dashboard payload | `interview` NULL · `diagnosis` MISSING · `sources` MISSING | **`interview` PRESENT · `diagnosis` PRESENT · `sources` PRESENT** · `cycle` = the active one (so the HeroCard meta row fills too) |

**Timing, honestly.** The first (cold) live read was **2.09 s**; the post-warm `?cycle=<active>` read was **0.034 s**
— the live path warms a cache. Report the **cold** number (~2.1 s) as the cost, not the warm one.

**Funnel coherence checked** (item 8, persona self-consistency): `useAIReadinessProgram.ts:141`
`scores[n] = isDone ? step.score : null` — the FE nulls any step whose *status* is not `completed`, so Ben's
`activeIdx = scores.indexOf(null) = 1` ⇒ **step 1 done / step 2 active / step 3 locked**. The backend's stray
`score: 21` on his `not_started` interview step (it computes step scores from signals regardless of the progress
row) is correctly **ignored** by the member funnel. Not a render gap.

### Residual findings (reported, not silently absorbed)

| id | Finding | Fate |
|----|---------|------|
| **R-1** | `howWeMeasure.cycleTotals.interviewQuestions = 0` on **both** read paths — a genuine zero cell in the "Handled for you this cycle" tile. | open — see progress.md S2 |
| **R-2** | Step-1 (`skill_mapping`) scores are low for both heroes (Aria **5**/30, Ben **2**/30) — few of their `user_skill_evidences` match the org's 8 `ai_readiness_skills`. Aria totals 64/100. Believable but weak for a "Champion" showcase. **A believability weakness, not an empty section.** | open |
| **R-3** | Ben carries interview/sim **signals** (score 21) while his `ai_readiness_user_step_progresses` row says `not_started` — the session seeder and the funnel seeder disagree. Invisible on the member funnel (R-2 note above) but it does feed his manager-side tier3 score. | open |
| **R-4** | The active cycle was applied to the **running** demo by hand for this probe. The **seeder** path is unit-tested (13/13 Go pkgs green, incl. a future-`end_date` assert) but **has not yet been proven on a cold reset-to-seed**. | **PENDING — must be run before close** |

---

## THE THREE LIVE PROOFS (round 3, 2026-07-14) — executed, not asserted

Round 1 and round 2 both ended with the same three items outstanding, and both recorded them honestly as
PENDING. They are the reason the milestone was not closeable: **an unexecuted assert is not evidence**
(M218 D17). All three are now executed. Two of them found real defects.

### PROOF 1 — the coverage sweep, EXECUTED (it had never been run against a live demo)

**Why it had never run: the harness could not be pointed at the demo.** `run-coverage.sh` hardcoded its
app/fapi bases to `localhost`, and the demo under test lives on a remote tailnet VM. That is not a footnote
— it is a large part of *why* the readiness asserts sat unrun for four releases.

**F-14 — running the sweep ON the demo host is the WRONG VANTAGE, and it produces a systemic false-RED.**
The first execution (from `billion` itself) reported `failingSections=21, personaFailures=3, reachable=8/150`.
Every one of the 21 was `region selector "main" matched 0 elements`, on every page; the screenshot was a bare
loading spinner. The network capture gave the cause:

```
POST https://billion.taildc510.ts.net:15050/graphql  ::  net::ERR_SSL_PROTOCOL_ERROR   (every call)
```

A `--public-host` demo **bakes the MagicDNS origin into the frontend build**, so the app's own GraphQL client
calls `https://<magicdns>:15050/graphql`. But `docker-proxy` binds `0.0.0.0`, so a connection **from the demo
host** to its own `100.x` tailscale IP hits the kernel socket and **bypasses `tailscale serve`** — which is the
thing that terminates TLS. Plain HTTP then answers a TLS handshake → *"wrong version number"*. Measured: from
the host, https on `:13000`, `:15050` **and** `:18082` all fail TLS; **from a tailnet peer all three answer
307/200/200.** The demo was healthy the whole time.

⚠️ **Recorded because a 21-section RED that is really an SSL artifact is the kind of false-RED that gets
"fixed" by weakening asserts.** The asserts were right; the vantage was wrong. Same class as D17, inverted:
absence of *rendering* read as absence of *content*.

**The sweep from the correct vantage (a tailnet peer):**

| Vantage / seat | reachable | failingSections | persona | escapes | notReached | crossPort |
|---|---|---|---|---|---|---|
| manager `dana-manager` | 69 | **0** | **0** | 0 | 0 | 1 |
| employee `aria-completed` | 60 | **0** | 1 → **0** (R-5) | 0 | 0 | 1 |
| employee `ben-started` | 60 | **0** | 1 → **0** (R-5) | 0 | 0 | 1 |

**All 8 manager readiness sections PASS** (`hero`, `cycle-picker`, `matrix`, `by-tag`, `how-we-measure`,
`interview-findings`, `handled-for-you`, `what-to-do-next`), and **both employee readiness sections PASS**
(`ai-readiness-member-done` for Aria, `ai-readiness-member-progress` for Ben — the surface that previously
rendered *literally nothing*). The zero-cell fence is a genuine **value** assert, not a label assert:
`textMatch` `[1-9][\d,]*\s+(AI skills mapped|minutes saved)` with `minMatches: 3` — it passes, so all three
counters are non-zero.

`crossPort=1` on all three runs = ant-academy `:13077` → **HTTP 502**. That is **F-13**, caught independently
by the sweep (the running stack predated the F-13 fix).

### PROOF 3 (≡ R-4) — the ACTIVE-CYCLE SEEDER, from a COLD reset-to-seed

Round 2's live proof **inserted the active-cycle row by hand**; a fix proven by hand-inserting a row is not a
fix proven by the seeder. It is now proven from zero: the host's rext clone was at the M219 seeder, and cycle 2
was a genuine cold `down --purge` (data dir gone, `PG_VERSION` mtime new). Everything below therefore came *out
of the seeder*:

| Measured on the cold-seeded DB | Value |
|---|---|
| `ai_readiness_cycles` | **2 — one `active` (end 2026-08-13) + one `closed` (score 62)** ← R-4 |
| snapshots / step-progresses / skills / sims | 199 / 532 / 8 (weight 6.5) / 2 |
| `jobsimulation.interactions` / `actors` | **1346 / 312** ← R-1 (was a hard 0) |
| Aria: readiness skills held | **8/8**, weight 6.5/6.5 → `round(6.5/6.5*30)` = **30/30** = `frozen_step1`. Total **89, "champion"** ← R-2 (was 5/30) |
| Ben: readiness skills held | **3/8**, weight 3.0/6.5 → `round(3.0/6.5*30)` = **14/30** = `frozen_step1`. Stage 1, "standby" ← R-2 (was 2/30) |
| Ben sessions on a readiness sim | **0** · stage-1 members with readiness sessions: **0** ← R-3 (the fence is structural, not lucky) |
| sessions per readiness sim | interview **156** (== stage-3 count) · simulation **177** (== stage-2 21 + stage-3 156) — exact funnel math |

### PROOF 4 — S5's Playthroughs actually RUN

`ptvalidate` only *statically* validated the manifest. Executed (same remote-vantage split — the Playthrough
runner had the identical `localhost` hardcoding, unfixed):

```
82 passed / 0 failed (4.8m)
ptreport: 14/15 passing (93.3%) — failing=0, unimplemented=1 (the declared TODO), unimplementable=0
  ai-readiness.manager-dashboard.UC1  PASS      ai-readiness.member-funnel.UC1  PASS
  ai-readiness.manager-dashboard.UC2  PASS      ai-readiness.member-funnel.UC2  PASS
```

The corpus's "14 live Playthroughs, 1 TODO" claim is now **executed**, not merely validated.

---

## R-5 (NEW — found by proof 1) — the readiness heroes' skills were the taxonomy's alphabetical head

The first employee sweep ever run on a Northwind seat failed persona `role-skills-coherence` on **both**
heroes. Two defects, one symptom.

**(a) Neither readiness role classified to a curated skill family.** M42e built the curated tier precisely to
keep the flat pool's `ORDER BY node_id` head (`15Five`, `3dcart`) out of a hero's profile — but it curated
exactly **two** families, `software` and `sales`, because those were the only orgs that existed. M51 then added
Northwind Aviation with **"Data Analyst"** and **"Operations Specialist"**, which match *neither*, so both
heroes fell through tier (2) into the flat pool. **Aria — a Data Analyst — claimed "24-hour dietary recall",
"2D Animation Software" and "3D Bioprinting in Dentistry".** The classifier's own comment blessed the
fall-through as *"no regression for an unclassified role"*; it is a silent regression for every role family
nobody thought to curate.

**(b) "Operations Specialist" IS NOT A PUBLIC JOB_ROLE.** The preset comment claimed *"resolving public role
(O6)"* — it does not exist in the taxonomy (which has Operations *Analyst* / *Manager* / *Engineer*, never
*Specialist*). So Ben had **no role at all**: no title on `/profile/skills`, no role-core skills, and therefore
**even his VERIFIED skills came off the flat head — he was "verified" in `15Five` and `17Track`.**

**Fix (rext `7ed67d3`):** new `data` + `operations` curated categories, allow-lists hand-picked and verified to
resolve against the live public taxonomy (the ops family also contains real junk for this persona — "Lean NOx
Traps", "Scheduling irrigations" — so hand-picked, never pattern-matched). Ben → **Operations Analyst**, which
resolves; it carries no role-skills of its own, so the new curated ops pool supplies them. Fixed in **both** the
demo preset and `pt-world` (which shipped the same two roles).

**The fence — the actual point.** Not "these two roles now classify" (patching two strings leaves the *next*
org to rediscover this on a presenter's screen), but `TestShippedPresets_EveryHeroRoleClassifies`: **every role
that any SHIPPED preset actually seeds must classify to a real curated category**, read from the real presets,
never a fixture. **Proven RED** against the pre-fix classifier — it names all four heroes across both presets —
and GREEN after. Ordering is load-bearing and pinned (`Operations Analyst` → ops, not data; `Sales Operations
Analyst` → sales; `Data Scientist`/`Engineering Manager` → software, unmoved).

**Why it survived four releases:** nothing ever asserted it. The employee sweep had **never been run on a
Northwind seat** — the readiness surface was a declared **BLIND AREA**, which is exactly what M219 existed to
close.

## F-13 — "started" meant "a pid existed one second ago"

Found on the live host, and independently confirmed by the sweep's `crossPort` 502. Two absence-read-as-success
defects on the bring-up's **own reporting path**:

1. The node check tested **existence**, not the engines requirement. `command -v node` passed with **Node 18**;
   ant-academy declares `engines: node >=22`, ships Next 16, and its toolchain imports `node:util`'s
   `styleText` (Node ≥ 20) — so `next dev` died at import on **every** bring-up. *A version requirement that is
   only checked for existence is not checked at all.* (The box had v22.22.1 under nvm, unused — a
   session-detached daemon never sources `nvm.sh`.)
2. The liveness probe polled `kill -0 $pid` for three seconds, then printed *"started (pid N)"*. The pid was
   alive at t=1s and gone by t=5s — **a probe that cannot outlive the thing it is probing is not a probe.**

The stack reported the academy as started, graded GREEN, and served a bare **502** on the AI Academy link for
the entire life of the stack. Fixed (rext `f803999`): resolve a satisfying node under `~/.nvm` or **fail loud**;
poll the **port**, not the pid.

---

## R-5b — the curated allow-list was a NO-OP for the very role it was written for

**The cold reseed is what caught it.** R-5 shipped, cycle 1 came up green, and the cold-seeded DB said:

| Hero | Role | `job_role_skills` | Result on the cold reseed |
|---|---|---|---|
| Aria | Data Analyst | **10** | **CLEAN** — Business Intelligence, Data Analysis, Looker, PowerBI, Agentic Coding, ML… |
| Ben | Operations Analyst | **0** | **STILL** "verified" in `15Five`/`17Track`; still claiming `24-hour dietary recall`, `2Checkout`, `2D Animation` |

…with the `operations` curated pool **fully populated and unused**.

**Root cause.** `skillsForRole`'s fallback ladder was **role → FLAT**. A public `job_role` can *exist* and
carry **zero** `job_role_skills` — "Operations Analyst" is exactly that — so `byRole` is empty and the
function handed back `r.flat`: the `ORDER BY node_id` head of the whole taxonomy, i.e. the junk pool the
curated tier exists to keep out of a profile.

And it defeated the curated tier **silently**. `combinedNamedPool` layers role → curated → flat, **but it
draws its ROLE tier from `skillsForRole`** — so for such a role **tier 1 was already the junk**, and it filled
the entire quota before the curated tier was ever consulted. Adding a curated allow-list for the role's family
therefore changed *nothing at all*. Both twins carried it (`namedSkillRefs` and `taxonomyRefs`), which is why
even Ben's **verified** chain certified him in `15Five`.

**Fix (rext `e284338`):** the ladder is **role → CURATED → flat**. Flat remains the last resort (never
fabricate, never fail to fill) but is now genuinely last. `TestSkillsForRole_ZeroRoleSkillsFallsToCuratedNotFlat`
fences the real shape — a role that *exists with zero role-skills* — **proven RED** against the old ladder
(*"fell through to the FLAT pool: got 15Five"*).

> ### ⚠️ The lesson — the D17 shape, a fourth time, and the sharpest instance of it
>
> **The classifier test and the pool-resolution test were BOTH GREEN while the seeder still wrote junk.**
> They proved the curated pool *resolved*. Neither proved anything ever *read* it.
>
> **A test that proves a thing exists is not a test that proves it is used.** The only thing that caught this
> was seeding a real database from zero and looking at what came out. That is the whole argument for the cold
> reset-to-seed battery, and it is why "the unit tests are green" was never going to be enough for this
> milestone.

### The D17 tally for this milestone (it bit SEVEN times before M219, and FOUR more inside it)

| # | The artifact that outlived the thing it described |
|---|---|
| 1 | A fence that asserted the **LEGACY** strings — and passed (round 2, `01f2644`). |
| 2 | A start check that read *"a pid exists"* as *"the service is up"* — the academy 502'd for the life of every stack (**F-13**). |
| 3 | My own R-2 verification query: `IN (SELECT skill_id FROM ai_readiness_skills)` silently bound `skill_id` to the **outer** scope (the table's column is `node_id`), so the `IN` was always true. It returned a plausible number and would have **passed** — only the impossible `36 > 8` exposed it. |
| 4 | My first render probe declared **"RENDERED at 4s"** on `main>0 && bodyLen>200` while the body still read *"Loading AI Readiness…"*. A success condition that does not exclude the loading state is not a success condition. |
| 5 | **R-5b**: two green tests proving a pool *resolved*, while nothing *read* it. |
