# Playthroughs — the readable map (by product × stream × proof depth)

**Purpose.** The tok-1 seed for the playthrough-sharpening milestone: what the 18 live Playthroughs
actually cover, by product and by vantage stream, and — the part the pass/fail count hides — **what
they prove and what they only look at.**

**Sources.** `rosetta-extensions/playthroughs/manifest/*.yaml` (8 files, 18 use cases, 0 TODO),
`playthroughs/e2e/tests/*.spec.ts` (18 browser specs + 5 unit specs), the M201 user-curated corpus
`knowledge/plan/releases/archive/02.00-opening-night/m201-manifest-corpus/manifest-draft.yaml`
(9 products / 28 use cases), M254 iter-10 live run on `billion` (18/18, 3.8 min).

**Status:** 18/18 GREEN live on billion (M254 iter-10, cold reset-to-seed, 2026-07-25). Everything
below is about what GREEN does and does not mean.

---

## 1. The headline, in three numbers

| | |
|---|---|
| **18 / 18 pass** | live on billion, 0 flake, 3.8 min suite (was 13 min before M254's `networkidle` fix) |
| **1 of 18 proves a WRITE** | `pt-assignment-assign` is the only Playthrough that mutates state and reads it back. The other **17 are render-presence proofs** wearing a journey's clothes. |
| **12 of the 28 curated use cases have no milestone home** | 16 of 28 are uncovered; M206 reserves 3, M207 reserves 1; **12 are un-homed** — nothing plans to build them. |

**One sentence:** the suite is green because it asks *"did the page render populated?"* almost
everywhere, and the platform can fail that question's *inverse* — *"does the thing actually work?"* —
without a single test going red.

---

## 2. By PRODUCT — the 18 live Playthroughs

| # | Product | Playthrough | Story / UC | Surface | Proof depth |
|---|---|---|---|---|---|
| 1 | **Profile & Skills** | `pt-profile-identity` | `profile.foundation.UC1` | apps/web `/profile` | render — hero's name visible |
| 2 | Profile & Skills | `pt-profile-verified` | `profile.verified.UC1` | apps/web `/profile/skills` | render — Spotlight + claimed-vs-verified non-empty |
| 3 | Profile & Skills | `pt-profile-growth` | `profile.growth.UC1` | apps/web `/profile/skills` | render — trajectory chart + gap stat |
| 4 | Profile & Skills | `pt-profile-timeline` | `profile.timeline.UC1` | apps/web `/profile` | render — Work + Education dated entries |
| 5 | **Skill Paths** | `pt-skillpath-legacy` | `skill-paths.legacy.UC1` | apps/web | **partial journey** — browse → open → start → *player opens*; asserts the step-completion control is **present**, not that completing it persists |
| 6 | **AI Simulations** | `pt-aisim-chat-launch` | `ai-simulations.chat.UC1` | apps/web | **launch boundary** — reaches `/sim/<slug>/start`; no turn, no completion, no result |
| 7 | **AI Readiness** | `pt-aireadiness-member-done` | `member-funnel.UC1` | apps/web `/home` | render — completed result + 3 steps |
| 8 | AI Readiness | `pt-aireadiness-member-progress` | `member-funnel.UC2` | apps/web `/home` | render — in-progress funnel + due date |
| 9 | AI Readiness | `pt-aireadiness-manager-dashboard` | `manager-dashboard.UC1` | apps/web `/ai-readiness` | render + **route assertion** (refuses the legacy orphan) |
| 10 | AI Readiness | `pt-aireadiness-manager-howwemeasure` | `manager-dashboard.UC2` | apps/web `/ai-readiness` | render — 3-step method + live interview findings |
| 11 | **Workforce Intelligence** | `pt-workforce-funnel` | `skills-funnel.UC1` | apps/web | render — mapped→verified funnel + chart |
| 12 | Workforce Intelligence | `pt-workforce-roster` | `roster.UC1` | apps/web `/enterprise/members` | render — populated rows |
| 13 | Workforce Intelligence | `pt-workforce-succession` | `talent-pool.UC1` | apps/web | render — succession / at-risk / mobility |
| 14 | **Assignment & Monitoring** | `pt-assignment-assign` | `assign-and-track.UC1` | apps/web | ✅ **WRITE + read-back** — assignable count drops by exactly one |
| 15 | Assignment & Monitoring | `pt-activity-drilldown` | `assign-and-track.UC2` | apps/web | render — aggregates + per-member drill-down |
| 16 | **Hiring** | `pt-hiring-recruiter-compare` | `recruiter-comparison.UC1` | **apps/hiring** | render + **app-identity assertion** (not ejected to workforce) |
| 17 | **Studio** | `pt-studio-advanced-generate` | `simulation-builders.advanced` | **studio-desk** | **generation boundary** — AI draft renders in the designer; generated content never asserted (P6) |
| 18 | Studio | `pt-studio-guided-generate` | `simulation-builders.guided` | **studio-desk** | render — Part-1 question + live-preview crest; the 5-part interview + final Generate are P6-out |

---

## 3. By STREAM (vantage) — who is proven, and how deep

| Stream | Hero seats | PTs | Apps touched | Deepest thing proven |
|---|---|---|---|---|
| **Employee / learner** | `pt-employee`, `pt-ai-completed`, `pt-ai-started` | **8** | apps/web | reaching a sim's launch route; opening a skill-path chapter player |
| **Manager** | `pt-manager`, `pt-ai-manager` | **7** | apps/web | **assigning a skill path (the one real write)** |
| **Recruiter** | `pt-recruiter` | **1** | apps/hiring | the Results scoreboard renders a comparable cohort |
| **Admin / content-creator** | `pt-manager` (reused) | **2** | studio-desk | the AI scenario draft renders in the advanced designer |
| **Anonymous / free / expired tiers** | — | **0** | — | nothing. Every Playthrough is `entitlement: enterprise`. |

### By app surface

| App | PTs | Note |
|---|---|---|
| `apps/web` (next-web) | **15** | |
| `apps/hiring` | **1** | a whole app on one test |
| `studio-desk` | **2** | first coverage ever, M252 |
| **`ant-academy`** | **0** | **zero Playthroughs.** Presence-only via the content-stories sweep. |

### Outcome shapes

| Shape | Count | Note |
|---|---|---|
| `success` | **18** | |
| `blocked` (a correct refusal — a gate/deny lands) | **0** | the manifest schema supports it; nothing uses it |
| `error` (a correct validation failure) | **0** | ditto |

**No Playthrough proves the platform correctly says *no*.** Every entitlement gate, permission deny,
and validation refusal is untested — and those are exactly the paths a demo audience trips.

---

## 4. The coverage gap vs the curated corpus (M201: 9 products / 28 UCs)

| Curated product | UCs | Covered | Missing |
|---|---|---|---|
| **Onboarding** | 5 | **0** | individual · ent-workforce-standard ×2 · ent-workforce-ai-readiness · ent-hiring |
| Skill Paths | 2 | 1 | `academy.UC1` → *M207-reserved* |
| AI Simulations | 3 | 1 | `code.UC1` · `interview.UC1` → *M206-reserved* |
| Profile & Skills Verification | 4 | 2 | `import.UC1` · `self-evaluation.UC1` (→ *M206*) |
| Workforce Intelligence | 5 | 4¹ | `organization-feedback.UC1` |
| **Org Admin & Settings** | 4 | **0** | roles · members · tags · feature-config |
| Assignment & Monitoring | 2 | **2** ✅ | — |
| Studio | 2 | **2** ✅ | — |
| **Talk to Data** | 1 | **0** | `query.UC1` |
| **Total** | **28** | **12** | **16** |

¹ `ai-readiness-monitoring.UC1` is discharged by the separate `ai-readiness` product's manager dashboard.

**Beyond the 18:** 6 live Playthroughs are net-new relative to the curated corpus (ai-readiness ×3
member/how-we-measure, hiring recruiter, profile identity + timeline).

### The un-homed 12

Of the 16 uncovered, **M206** reserves 3 (`ai-sim.code`, `ai-sim.interview`, `profile.self-evaluation`)
and **M207** reserves 1 (academy). **These 12 have no milestone home anywhere in roadmap or vision:**

1–5. **Onboarding** — all five journeys. *The first thing every real user does is untested end-to-end.*
6–9. **Org Admin & Settings** — roles, members, tags, feature-config. *Every one is a WRITE surface.*
10. `workforce-intelligence.organization-feedback.UC1`
11. `profile-skills.import.UC1`
12. `talk-to-data.query.UC1`

Plus, not in the curated corpus at all: **course-builder**, **AI Labs + credits** (M231: presence-only,
nil client), and the **server-owned academy** domain — three `app`-owned domains with zero journey coverage.

---

## 5. Why green can coexist with "things not working"

Five structural reasons, each independently sufficient:

1. **Render ≠ function (17/18).** A populated page proves the *seed* wrote rows and the *read path*
   resolves. It proves nothing about the write path, the engine, or the outcome. The demo's data is
   seeded — so the read path is the half that was never in doubt.
2. **No negative controls.** No Playthrough is proven to go RED when its outcome is absent. M219's
   whole lesson — *"a surface that renders is not the same as the RIGHT surface"* — is currently
   enforced in exactly one place (`LEGACY_AI_READINESS_URL`). Nowhere else is a false-green fenced.
3. **Journeys stop at the boundary.** `pt-aisim-chat-launch` stops at `/start`. `pt-skillpath-legacy`
   stops at "the completion control exists". `pt-studio-*` stops at "the draft rendered". The
   *interesting half of every journey is out of scope by design* (P6) — which is defensible, but it
   means the suite cannot see a broken engine.
4. **One entitlement, one org shape.** Everything is `enterprise` on `pt-world`. No free tier, no
   expired tier, no anon, no deny path, no second org shape.
5. **Whole surfaces at zero.** ant-academy 0, onboarding 0, org-admin 0, talk-to-data 0. Nothing red
   because nothing looks.

---

## 6. Speed anatomy (the "individually faster" half)

| | |
|---|---|
| Suite wall-clock | **3.8 min** on billion (M254 iter-10; was 13 min before the `networkidle`→`domcontentloaded` fix) |
| Browser Playthroughs | 18 |
| Unit specs in the same project | ~99 (`url-shapes` 671 LoC, `studio-builder-locators` 345, `ai-readiness-locators` 207, `stack-env` 176, `page-object` 103) — ~1 s total |
| Per-Playthrough average | **~12.6 s** — misleading; see the outlier row below |
| Concurrency | **`workers: 1`, `fullyParallel: false`** — hard-pinned in `playwright.config.ts` |
| Retries | 0 (deliberate — a flake is a defect) |
| Per-test timeout | 120 s default — **but three specs override it**: `studio-builder.spec.ts:45` = **300 s**, `:91` = **180 s**, `assignment-assign.spec.ts:43` = **240 s**. `expect` 15 s |
| **The outlier** | `pt-studio-advanced-generate` is a real **~2-3 min live-LLM round-trip**. 228 s is consistent with **studio-advanced ~120 s + 17 tests at ~4.5 s each** — the suite is dominated by ONE test, the ~12.6 s average is an artifact of it, and any flat suite-wall-clock target must budget the LLM lane separately |

> **⚠️ CORRECTED 2026-07-27** (adversarial plan review). The first version of this section claimed
> *"17 of 18 mutate nothing, so a read-only lane at `workers: N` is safe under the existing rationale."*
> **Both halves are false** — see below. The *"1 of 18 proves a WRITE"* headline is **accurate** (it
> means: mutates state **and reads it back**, the `playthroughs.md:169-172` sense); it is the derived
> *"17 mutate nothing"* that was wrong.

**One lever survives, one was false:**

- **PASS — per-seat `storageState` reuse.** All 18 log in from scratch across **6 distinct seats**.
  Reuse pays the handshake ~6x instead of ~18x — order **~30 s** — with `pt-profile-identity`
  retained as the one test that proves the handshake itself. Real, but small.
- **FAIL — "a read-only lane at `workers: N` is safe."** Wrong twice.
  **(a) The count.** `tests/skillpath-legacy.spec.ts:21-23` **self-declares** *"MUTATION: starting a
  path creates progress state"* (via `getOrCreateSkillPathSession`, a server-side create-on-read),
  echoed at `lib/skill-path-page.ts:47-48`. Three more are unclassified (`studio-builder.spec.ts` x2
  fire a real LLM generation; `aisim-chat-launch.spec.ts:61` clicks Start Simulation). Only **10 of
  18** carry an explicit "(no mutation)" note. Accurate statement: **17 of 18 are UNCLASSIFIED for
  mutation; at least 1 demonstrably mutates.**
  **(b) Postgres is not the binding surface.** Clerkenstein's fake FAPI holds **one global active
  seat / `signedIn` / `sessID` per stack** (`clerk-frontend/registry.go:67,75`; `server.go:100-105`).
  Every login hits `POST /v1/demo/select` -> `handleSelectIdentity` (`server.go:573-586`), which
  re-points the seat **and** clears `signedIn`/`sessID` **globally**. Under `workers: N`, worker 2's
  login signs worker 1 out mid-journey — and `server.go:454-466` reads `activeUserLocked()` with **no
  cookie input**, so **`storageState` reuse does not isolate it either**. Already recorded in-repo at
  `stack-verify/e2e/tests/m224-candidate-heroes.spec.ts:10-15` and `content-stories.spec.ts:128-130`.
  **(c)** The existing rationale sanctions only **stack-per-worker** or **per-worker seed partitions**
  (`spec.md:447-450`; `playthroughs.md:441-443`). Parallelism is an **enabler to be built** (a
  cookie-scoped Clerkenstein registry, or one fake-FAPI per worker), not a free lever.

---

## 7. What this map suggests for the milestone

Three axes, matching the ask:

- **Faster** — read-only parallel lane + per-seat `storageState` reuse; the 120 s timeout and 15 s
  expect ceiling are the flake-hiding budget, not the target.
- **More effective** — negative controls per Playthrough (prove it goes RED); convert boundary-stops
  into outcome assertions where the platform permits; add the first `blocked`/`error` outcomes.
- **Better coverage** — the un-homed 12, ranked: **onboarding** (5, every user's first journey) and
  **org-admin** (4, all WRITE surfaces) are the two clusters that most plausibly explain
  *"things not working but built up"*.

_Compiled 2026-07-27. Live-state anchor: M254 iter-10, billion, 18/18 GREEN._
