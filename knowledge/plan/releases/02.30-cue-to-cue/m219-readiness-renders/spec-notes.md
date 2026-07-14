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
