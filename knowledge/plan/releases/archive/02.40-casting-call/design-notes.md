# v2.4 "casting call" — roadmap proposal: a HIRING demo org on the presenter cockpit

**Designed 2026-07-15** via `/developer-kit:design-roadmap` (proposal stage — Phase 6 presentation; not yet
promoted to `roadmap.md`). Built on a 3-report research workflow (platform hiring model · rext seeder ·
cockpit/enterprise integration). This is a **NET-NEW** release: nothing hiring-specific is seeded today.

> **This release formally RE-OPENS v2.3's `D-DESIGN-4`** — *"the three story orgs are the three that already
> exist; there is no hiring org and none will be built."* The user's 2026-07-15 ask reverses it. D-DESIGN-4 was
> **right that a hiring org is net-new work**, but its stated blocker is **partially REFUTED by the research**: it
> claimed a hiring story "would need unmapped domain tables **+ the `hiring-app` frontend, which is not in the demo
> UI tier**." In fact the candidate-comparison surface **ships inside the dockerized `apps/web` (Workforce) app the
> demo already builds** — the Vercel-only `apps/hiring` product is **not** required — and the domain primitives
> (`organizations.is_hiring`, the `candidate` membership role, `jobsimulation.sessions` typed `SIMULATION_TYPE_HIRING`)
> **already exist in the platform**. The reversal is therefore tractable **within the zero-platform-edit wall**, but
> it is **not a clean section release**: two blind areas (the hiring read-model, and proof-by-rendering that the
> comparison surface is demo-servable) gate it, so the release **opens with an investigation-first spike**.

---

## Phase 0 preflight verdicts (recorded per the skill)

- **0a Deferral audit — YELLOW (proceed).** v2.3 closed with 4 non-gate tail carries signed off → v2.4
  (`F4` academy grid · `BURNIN-M221-dev-public-host` · `F-M220-4` · `PROBE-M218-c3-rerun`). **None of the four
  overlaps this hiring theme** — all are academy/live-infra follow-ups. They can ride v2.4 as a separate track or
  stay parked; this proposal scopes **only** the hiring theme and does not fold them in. No repeat-deferral
  pattern touches hiring.
- **0b KB blind-area check — RED for hiring (expected; drives M222).** Grep of `corpus/**` for a hiring model
  returns **no doc anchor**. Hiring exists in the corpus only as (i) a "distinct-frontend" line in
  `next-web-app.md` and (ii) the business KB (`kb-ant-business`). **The hiring model is a blind area** → the
  release must author it (`M222 Delivers → corpus/services/hiring.md`) **before** the seeder codes against the
  contract. This is exactly the failure mode the check exists to catch.

---

## 1. Gap analysis

### 1a. What EXISTS and is reusable (zero-platform-edit)

The strongest research finding is that **most of the hiring machine already exists** — in the platform *and* in
the rext seeder fleet — it is simply **never wired into a hiring org**:

| Piece | Where it lives today | How we reuse it |
|-------|----------------------|-----------------|
| **Org-type discriminator** | `organizations.is_hiring` bool (`app/…/schema/organization.go:39`), wired end-to-end in the platform | Flip it `true` for the one hiring org (today hardcoded `false` at `stack-seeding/seeders/org.go:65`) |
| **Candidate identity** | `MembershipRole` enum has `candidate`; the platform auto-assigns `RoleCandidate` when a user joins an `is_hiring` org (`organization/manager.go:373`) | The 45 candidates are `role=candidate` memberships — **not** a net-new entity |
| **Manager / enterprise access** | `role=admin` + a Casbin `g2(org,user,admin)` grant + `g3` feature grant + Sentinel `m3` | The 5 managers reuse the **verbatim** workforce-manager seat (`roleForHero: manager→admin`) — Dan/Leah/Dana's mechanism |
| **The assessment record** | `jobsimulation.sessions` typed `SIMULATION_TYPE_HIRING`, `score` 0–100; comparability = many candidates sharing one `sim_id` | `JobsimSessionsSeeder` already writes scored HIRING sessions (the G14 valid-enum fix); change only `sim_id` selection to the 5 shared sims |
| **★ The candidate-COMPARISON surface** | **Ships in the dockerized `apps/web`**: `/enterprise/activity-dashboard → AI-Simulations tab → [simId]` (`InsightsByMembersContainer` + `useGetInsightsByMembers` → `GET_INSIGHTS_BY_JOB_SIMULATIONS_BY_MEMBERS`) = a sortable per-sim candidate leaderboard w/ `CandidateSearchInput`; plus `[simId]/[userId]` detail + Interview Insights | **No new frontend build.** This is the manager view that compares 45 candidates on the same 5 sims |
| **The narrative-gated-org pattern** | AI-readiness M51: `narrative:` discriminator on a story → a dedicated seeder set (`OrgSettings/AIReadinessConfig/AIReadinessFunnel`) | `narrative: hiring` gates a `HiringSeeder` set with **zero** changes to the seeder fleet's control flow |
| **★ The "N members, same shared sims, comparably scored" template** | `AIReadinessFunnelSeeder` resolves an org's shared sim refs once, writes a per-member scored session against them, rolls each into a per-member snapshot | 45 candidates × 5 shared sims is a **direct generalization** (2 shared → 5) of this exact seeder |
| **Cockpit + seat-switch** | `BuildCockpitManifest` / `WriteRoster` are pure blueprint projections of `stories.seed.yaml` | A 4th story's heroes **auto-appear** on the cockpit + Clerkenstein — no cockpit code change |
| **Reserved-tail sim refs** | `contentref.go reservedSimRefs/reservedAt` (the AI-readiness R-3 fix) | Reserve the 5 hiring sims disjoint so generic activity sessions can't collide/mis-score a candidate |
| **Closure + isolation gates** | `datadna measure-closure`, `main.go resetTables`, `isolation.Guard` | New hiring rows join `resetTables` + stay closure-measured across all orgs |
| **Real hiring CONTENT** | `directus.job_position` = 443 rows, **all public**; `directus.simulations` carries `type` + `job_position` FK — **715 HIRING sims (127 public, 306 with a job_position)** | The 5 "job positions" can be **REAL replayed content**, not placeholders |

### 1b. What is genuinely NET-NEW (all tooling / docs — NO platform edits)

1. A **blueprint field** `is_hiring: true` (and/or `narrative: hiring`) on the story/org spec — the `OrgSpec` has
   **no** hiring notion today (just `{Name, Slug, Industry, Narrative}`).
2. **`OrgSeeder` one-value change** — read `is_hiring` from the spec instead of the hardcoded `false`. Small, but
   **it is the load-bearing gate**.
3. A **`HiringSeeder`** (M51-analog): 5 managers (`admin`) + 45 candidates (`candidate`), and a
   **candidate-assessment funnel** giving all 45 scored `SIMULATION_TYPE_HIRING` sessions on the **SAME 5** shared
   `sim_id`s — today sessions are hash-distributed across a ~50-id pool, **never concentrated**.
4. **Reserve 5 shared Directus sim ids** as the "job positions" (disjoint reserved tail, the `ai_readiness_sims`
   pattern).
5. A **type-aware hiring-sim content reader** (`type='SIMULATION_TYPE_HIRING' AND job_position IS NOT NULL`) — the
   current `contentref` pool is **type-blind** (`SELECT id ORDER BY id LIMIT 50`), so there is no way today to pick
   the 5 hiring sims tied to job positions (the `readAIReadinessSkillPool` precedent for a dedicated pattern query).
6. A **snapshot extension to REPLAY `directus.job_position` rows** (all 443 public) + guarantee the 5 chosen HIRING
   sims are captured — today `job_position` is **structure + a synthesized public-read grant only** ("no replayed
   rows the expansion is simply NULL"). Without this, real hiring↔position content is not in a stack.
7. **Extend Clerkenstein** fake FAPI/BAPI to emit org `public_metadata.isHiring=true` — **tooling-owned, NOT a
   platform patch** — *only if* the candidate-facing hiring-flavored UX is in scope (see Q1). Today the roster
   emits org `public_metadata.eid` only.
8. **Cockpit deep-link catalog entries** for the recruiter surfaces (`/enterprise/activity-dashboard`, per-`[simId]`
   comparison) — **none exist today** — plus the hiring hero(es).
9. **A NET-NEW corpus doc** for the hiring model (`corpus/services/hiring.md`) + registrations in
   `seeding-spec.md` / `stories-spec.md` / `README.md` for the 4th story + the `is_hiring` gate.
10. **Coverage manifest hiring vantage** + a **hiring playthrough** (the reserved `playthroughs/manifest/hiring.yaml`
    slot) + the hiring org into the decoupled `pt-world` seed.

### 1c. Blind areas (no doc/code anchor — must be authored or proven-by-rendering first)

These are the reason the release is **not** a clean section stack. Each needs a `Delivers →` doc line or a
render-proof before code commits against it:

- **BA-1 — the hiring READ-model.** No corpus doc. The exact read-path that fills the activity-dashboard per-sim
  candidate list with a **comparable score** was *not* traced end-to-end. **Unproven whether the score/best-session
  comes purely from `jobsimulation.sessions.score` (which the seeder writes) or ALSO needs a
  `validation_attempt_results` / evaluation row per session** the current seeder may not populate — a potential
  **extra seeding surface (data-only, not a platform edit)**. → **`M222 Delivers → corpus/services/hiring.md`.**
- **BA-2 — does the comparison surface render for an `is_hiring` org?** Many `apps/web` pages **hardcode
  `isHiringOrg={false}`**; the exact set of surfaces that change vs stay identical under `isHiring` was **not
  enumerated**. The recruiter comparison view *appears* `isHiring`-independent (admin/enterprise-gated) but this was
  **inferred, not proven by rendering**. Also unresolved: the `isEnterprise=Boolean(organization)` (nav) vs
  `isEnterprise=!isHiringOrg && organizationId` (billing) **divergence** — the two definitions differ and the blast
  radius is unverified. → **M222 must prove by rendering.**
- **BA-3 — is the demo-servable surface actually in `apps/web`, or only `apps/hiring`?** Strong evidence it is
  `apps/web` (route files present; `Dockerfile.dev` builds `@anthropos/web-app`), but a **full route-by-route diff
  `apps/hiring` vs `apps/web` was not done**. If it turns out `apps/hiring`-only, showing it = containerizing a
  Vercel-only app = **large net-new + likely a platform edit → escalation.** This is the **make-or-break go/no-go**.
- **BA-4 — Job Fit Analysis / anti-cheating report / 0–5 competency breakdown** (named in the product KB) were
  **not located** as concrete tables/resolvers — unknown whether they need dedicated seeded data or are LLM-derived
  at read time.
- **BA-5 — the candidate end-user footprint.** Whether a seeded candidate needs a *rendering* `/profile` (the
  7-table verified-skill chain) or only needs to appear as a scored row in the recruiter's per-sim list — the 45's
  minimum viable data footprint is unresolved (couples to Q1).
- **BA-6 — does the cold-primed public snapshot contain ≥5 usable HIRING sim blueprints?** Probe confirms 715
  HIRING sims / 127 public *in prod*, but no anchor confirms the **captured** snapshot pool (private=false, ~304
  rows) carries ≥5 HIRING-typed sims tied to job positions vs being TRAINING-dominated.

### 1d. Open questions the USER must decide

Consolidated across the three reports; the 3 highest-leverage are surfaced in §5. The full set: scope of the
hiring "flavor" (recruiter-only vs candidate-facing) · which of the 5 managers + 45 candidates become cockpit
heroes · real-replayed vs synthesized 5 positions · full 5-sim grid (225 sessions) vs realistic funnel · whether to
seed `organization_sim_invitation_links` (faithful) or skip them (the comparison reads sessions, not links) ·
confirm this is its own release (the prior gap-analysis judged it "a separate release, not a milestone").

---

## 2. Milestone design

**Shape summary.** The **seeder half is section-shaped** (a direct M51 analog the research explicitly calls out).
But the release is **gated on two uncertain-path items** — reverse-engineering the undocumented hiring read-model
(BA-1) and proving-by-render the comparison surface is demo-servable (BA-2/BA-3). So it **opens with an
investigation-first spike (M222)**, keeps the enumerable seeder/content work as **section** milestones, and makes
the **render-proof (M224)** and the **live proof (M226)** **iterative** — because "seed → render → discover the
gate bypassed the seed → fix → re-render" is exactly the M219 measurement-driven loop, not a fixed checklist.

### M222 — Read the room *(the hiring-model spike + doc + the `is_hiring` gate)*
- **Shape: `section`** (small→medium). *Justify:* the deliverables ARE enumerable — a corpus doc, a set of
  render-probes, a go/no-go decision, and a one-value gate thread. What is uncertain is the *result* (which then
  re-scopes M224), not the task list. This is the "clean stage" analog: make the foundation real + prove the path
  before committing the risky work.
- **Goal:** author the missing hiring model doc, and **prove by rendering** — on a throwaway hand-seed against a
  live dockerized `apps/web` — that the recruiter comparison surface (a) exists in the demo-servable app and (b)
  renders a comparable score from seedable data. Land the `is_hiring` gate + `narrative: hiring` discriminator.
- **Scope In:** author `corpus/services/hiring.md` (the model: `is_hiring`, `candidate` role, hiring sims ↔
  `job_position`, `jobsimulation.sessions` HIRING, the read-path for the comparison surface, the `isHiringOrg`
  Clerk-publicMetadata derivation, the `isEnterprise` divergence blast radius) · add the blueprint `Org.IsHiring`
  field + `narrative: hiring` · thread `is_hiring=true` into `org.go` (the load-bearing one-value change) ·
  reserve the hiring-org deterministic `OrgID` · **render-probe answers to BA-1/BA-2/BA-3/BA-6** and record a
  **go/no-go + the exact seeder-output contract** M223/M224 build against (does the score render from
  `sessions.score` alone, or does it need a `validation_*`/eval row per session?).
- **Scope Out:** the full 50-person seed · the assessment funnel · cockpit heroes · any latency work.
- **Depends on:** nothing (release entry point). A green `/demo-up` on `billion` (v2.3 left one **live**) is a free
  render substrate.
- **KB deps:** `kb-ant-business` hiring.md / ai-interview.md · `corpus/ops/demo/stories-spec.md` (the 7-table chain)
  · `corpus/ops/snapshot-spec.md` · v2.3's `demopatch-spec.md`.
- **Delivers →** `corpus/services/hiring.md` **(BLIND AREA — BA-1/BA-2)**; registrations pending in
  `seeding-spec.md`/`stories-spec.md`/`README.md`.
- **Demo-patch?** **No patch built here** — but M222 *decides whether one is needed downstream*. If the render-probe
  shows the surface hard-gates on `isHiringOrg` or the score demands an un-seedable resolver path, that becomes an
  **M224 demo-patch** (the M219 precedent). If BA-3 comes back "`apps/hiring`-only," **escalate — do not proceed.**

### M223 — Casting the ensemble *(the `HiringSeeder`: 5 managers + 45 candidates + 5 shared positions + content)*
- **Shape: `section`** (medium→large). *Justify:* a direct M51 analog — 1 preset org + narrative-gated seeders +
  reset/closure wiring + a content reader/replay — all enumerable up front against the M222 contract.
- **Goal:** the hiring org exists in the seed: ≤50 people (**exactly 5 `admin` + 45 `candidate`, no `member`**),
  distinct from the 3 workforce orgs, with the **5 shared job positions** resolved to **real replayed content**.
- **Scope In:** a 4th `stories[]` entry (`narrative: hiring`, `RoleMix` ≈ 0.1 admin / 0.9 candidate) with a hero
  trio placeholder · a `HiringConfigSeeder` (AI-readiness-config analog) defining the org's **5 shared HIRING sims**
  · the **type-aware hiring-sim reader** (`type=HIRING AND job_position NOT NULL`) · the **snapshot extension to
  replay `directus.job_position` rows** + pin the 5 chosen HIRING sims · wire new hiring rows into `resetTables` +
  the closure gate + the isolation audit · the `candidate`-assessment **funnel seeder** (AI-readiness-funnel analog)
  writing each of the 45 candidates a scored `SIMULATION_TYPE_HIRING` session against **all 5** shared sims (or a
  funnel — see Q3), each rolled into a **comparable-but-DIFFERENTIATED** per-candidate score (a realistic spread,
  not a flat arc) · optionally seed `organization_sim_invitation_links` (see Q).
- **Scope Out:** the render proof (M224) · cockpit heroes (M224) · coverage/playthrough (M225).
- **Depends on:** M222 (the seeder-output contract + real-vs-synth decision + is_hiring gate).
- **KB deps:** `stories-spec.md` (7-table fan-out) · `seeding-spec.md` · `snapshot-spec.md` · `ai-readiness.md`
  (the funnel contract) · `corpus/services/hiring.md` (M222).
- **Delivers →** updates to `snapshot-spec.md` (the `job_position` replay surface) + `seeding-spec.md`/`stories-spec.md`
  (the 4th story + `is_hiring` gate).
- **Demo-patch?** **Pure seeding** (+ a tooling-owned snapshot extension). No platform-render gap here — this is
  data. *If* M222 found the read-path needs a per-session `validation_*`/eval row, that extra row is **added to the
  seeder (still data-only)**, not a patch.

### M224 — The callback *(cockpit heroes + enterprise access + Clerkenstein wiring + the comparison RENDERS)*
- **Shape: `iterative`** (large). *Justify:* the render-risk heart. Whether ~45 comparable, non-junk candidate rows
  actually paint on the manager's comparison surface is **measurement-driven** — the exact read-path is a blind area
  (BA-1) and the M219 lesson is that *a render gate silently bypasses the seed* (`CycleID==nil→buildLiveResponse`;
  the PostHog `flag-gate`; junk skills from a dry pool). You **seed → render → attribute the gap → fix (data-only
  seeding OR a sha-pinned demo-patch OR Clerkenstein wiring) → re-render.** A fixed `In:` list would be speculative.
- **Goal:** click **[Log in as]** the recruiter hero → land on `/enterprise/activity-dashboard` and see, for **each**
  of the 5 positions, **~45 rankable, comparable, non-junk** candidate rows; and (if in scope, Q1) the org reads as
  hiring, and a candidate hero renders a usable assessed `/profile`.
- **Exit gate:** on a **cold reset-to-seed**, the manager hero's comparison surface renders **≥40 comparable
  candidate rows per each of the 5 sims** with a **realistic non-degenerate score distribution**, **0 junk
  skills/roles/names** (closure green), **0 prod-eject escapes**, over **≥3 consecutive cold runs**. (Latency is
  *reported, not gated* here — gated at M226, per v2.3's D-DESIGN-1.)
- **Iteration protocol:** `corpus/ops/demo/coverage-protocol.md` (the live-browser measure→attribute→fix→re-measure
  loop) + `verification.md`.
- **Scope In:** the hero trio (**1 recruiter/manager** → `jump_to` the comparison surface; **2 candidate exemplars**
  — strong + weak on the same position — → `/profile`|`/profile/skills`) · **DeepLinkCatalog** entries for the
  comparison surface (a proper `[Log in as]` label; a `NeedsID` entry if the jump targets a specific `[simId]`) ·
  extend Clerkenstein FAPI/BAPI to emit org `public_metadata.isHiring` **(iff Q1 = candidate-flavored)** · drive the
  render loop to green.
- **Scope Out:** the coverage-sweep gate + playthrough (M225) · the live billion proof (M226).
- **Depends on:** M222, M223.
- **KB deps:** `coverage-protocol.md` · `cockpit-spec.md` · `clerkenstein.md` · `stories-spec.md` · `hiring.md`
  (M222) · `demopatch-spec.md`.
- **Delivers →** the render-path + hiring-vantage section into `corpus/services/hiring.md` + `cockpit-spec.md`.
- **Demo-patch? — LIKELY, one of two shapes (this is the release's D-DESIGN-2 order-of-preference call):**
  1. **`next-web-hiring-flag-gate`** — if a comparison surface hard-gates on `isHiringOrg` the way M219's member
     surface gated on an `undefined` PostHog flag. **Prefer the Clerkenstein `isHiring` wiring first** (tooling, no
     patch); only if a surface *still* won't mount does it route to a sha-pinned demo-patch (the M219 precedent).
  2. **A perf demo-patch** for the whole-org-hydration of the 45×5 compare table (the AI-readiness `loadMembers`
     180 s→19 ms precedent) — but **latency is not this milestone's gate**, so this may defer to M226.
  **A platform-repo edit is never in bounds; an un-patchable surface escalates.**

### M225 — Dress the set *(demo integration: set-dress replay, coverage sweep hiring vantage, hiring playthrough)*
- **Shape: `section`** (medium). *Justify:* enumerable — extend the auto-set-dress bring-up to replay `job_position`,
  author a hiring coverage manifest + a hiring playthrough, wire the org into `pt-world`. Reuses the M42 coverage +
  M202 playthrough machinery (never forked).
- **Goal:** the hiring org comes up **auto-set-dressed** on a default `/demo-up`, passes a **hiring-vantage coverage
  gate**, and has **one GREEN playthrough** proving the recruiter journey end-to-end.
- **Scope In:** fold the `job_position` replay + 5-sim capture into the auto-set-dress pass · a **hiring coverage
  manifest** wired into `manifestFor(vantage, expectedOrg, identityKey)` (org-conditional dispatch, the AI-readiness
  precedent) — candidate persona self-consistency (role↔skills↔score) + the compare-surface sections + 0 prod-eject
  · a **`playthroughs/manifest/hiring.yaml`** use case (recruiter compares candidates on a shared sim; optionally a
  candidate completes a hiring assessment) + the hiring org into the decoupled `pt-world` seed.
- **Scope Out:** the live cross-machine proof (M226).
- **Depends on:** M223 (frozen seed shape) + M224 (a rendering surface to sweep). **Note:** the manifest *authoring*
  can begin once M223 freezes the seed shape — a partial overlap with M224's render loop — but the coverage/playthrough
  **gate** cannot pass until M224 is green.
- **KB deps:** `coverage-protocol.md` · `playthroughs.md` · `frontend-tier.md` · `snapshot-spec.md`.
- **Delivers →** the hiring sections of `coverage-protocol.md` + `playthroughs.md`.
- **Demo-patch?** Pure tooling (manifests + seed + set-dress). No platform-render gap.

### M226 — Opening night *(prove it on billion)*
- **Shape: `iterative`** (large). *Justify:* the direct analogue is **M215 "prove-on-odyssey"** and **M221
  "prove-on-billion"** — the last breakages only surface on a live cross-machine run over the tailnet; the path is
  discovered, not enumerated.
- **Goal:** every requirement of this release verified **on the remote VM, over the tailnet, default `/demo-up`,
  no flags.**
- **Exit gate:** on `billion.taildc510.ts.net`, a default `/demo-up N` yields, **reproducibly on a cold
  reset-to-seed**: **(1)** the hiring org present, `is_hiring=true`, **exactly 5 managers + 45 candidates**;
  **(2)** the recruiter hero lands on the comparison surface and sees **≥40 comparable non-junk rows per each of
  the 5 positions**; **(3)** the 2 candidate heroes render usable assessed profiles (iff Q1); **(4)** the org reads
  as hiring (iff Q1); **(5)** **p95 click→ACCESS < 5 s** for the recruiter vantage (v2.3's gate extended to a 3rd
  measured path); **(6)** coexists with the 3 workforce orgs on the cockpit; **(7)** **0 platform-repo edits.**
- **Iteration protocol:** `corpus/ops/verification.md` + `coverage-protocol.md` + `latency-budget.md` (v2.3).
- **Depends on:** M222, M223, M224, M225.
- **KB deps:** `tailscale-serve.md` · `latency-budget.md` · `verification.md` · `safety.md`.
- **Delivers →** the live-proof record; any latency finding folds into `latency-budget.md` (a hiring 3rd vantage).
- **Demo-patch?** Whatever M224 pinned, **re-proven live at final code** (the M221 `REPROVE-…-at-final-code`
  discipline); a live-only perf gap may pin a perf demo-patch here.

### On the reserved M205

The vision reserves **M205 "Hiring + tier gates"** (the recruiter vantage + free→paid Stripe entitlement gates +
an ATS-style post-a-role→advance pipeline). This release **DISCHARGES the recruiter/seeder half of M205** but
**explicitly LEAVES OUT** (a) the **Stripe tier gates** (no mirror engine exists; spec §5.8) and (b) the
**ATS candidate-pipeline** (post-a-role / applications / advance-a-candidate) — which the platform **deliberately
does not model** ("Applicant Tracking: NOT job posting management… We integrate with ATS"). So M205's tier-gate half
stays a vision reservation (or is renamed); v2.4's hiring org is a **comparison-of-assessed-candidates** demo, **not**
a hiring-*product* build.

---

## 3. Version proposal

### v2.4 — "casting call"

- **Codename:** **`casting call`** — two-word, evocative, lowercase, in the release lineage's **stagecraft**
  family (body double → set dressing → dress rehearsal → quick change → cue to cue). A *casting call* is the theatre
  event where **many candidates audition for the same roles and the director compares them side by side** — a
  near-exact metaphor for the ask (45 candidates, the same 5 positions, a manager comparing them). Alternates if it
  doesn't sit right: **`open call`**, **`final callback`**, **`the shortlist`**.
- **One-line theme:** *The recruiter's vantage — a 4th, HIRING org where 45 candidates audition on the same 5
  positions and a manager compares them side by side, distinct from the three workforce orgs on the cockpit.*
- **Tag:** **`v2.4`** · branch `release/02.40-casting-call`.
- **Milestone numbers:** **M222 → M226** (the next free `Mxyy` band after v2.3's M221; assigned at design time per
  the `Mxyy` rule — the low-numbered reserved M205 is *consumed in intent*, see above, not re-used as a literal
  number).
- **Execution order — largely SEQUENTIAL** (each milestone consumes the prior's output; the render can't be proven
  before the seed exists, the coverage can't sweep before the render is green, the live proof needs everything):

```
v2.4 "casting call"

  M222 ─────→ M223 ─────→ M224 ─────→ M225 ─────→ M226
  read the    casting     the         dress       opening
  room        ensemble    callback    the set     night
  (spike +    (seeder +   (render,    (coverage   (prove on
   doc +       content +   iterative)  + playthru) billion,
   is_hiring   5 shared               ╎            iterative)
   gate)       positions)             ╎
   section     section    iterative   section──────╯ authoring
                                        can begin once M223 freezes
                                        the seed shape (partial overlap)

  M222 is a HARD go/no-go barrier: if its render-probe shows the comparison
  surface is apps/hiring-only (BA-3), the release ESCALATES rather than proceeding.
```

**Parallelism matrix:** minimal by design. The only real overlap is **M225 manifest *authoring* ∥ M224 render
loop** (yes-with-caveats: authoring can start on M223's frozen seed shape, but M225's *gate* waits on M224 green;
merge risk low — different files). Everything else is a hard data/render dependency chain. Iterative milestones
(M224, M226) do **not** parallelize with siblings — they surface cross-cutting issues.

---

## 4. Risks

| # | Risk | Severity | Mitigation |
|---|------|----------|------------|
| **R1** | **A platform HIRING render path bypasses the seed** — the M219 precedent, *twice over*: `CycleID==nil→buildLiveResponse` and the PostHog `flag-gate` both let a surface render **empty over correctly-seeded data**. If the comparison surface hard-gates on `isHiringOrg`, or the score demands a `validation_*`/eval resolver path the seed can't feed, "45 seeded candidates" shows as **0 rows**. | **blocks-milestone (M224)** | **M222 proves the read-path by RENDERING before M224 commits.** Prefer Clerkenstein `isHiring` wiring (tooling); a genuine platform-source render gate routes to a **sha-pinned demo-patch** (D-DESIGN-2 order); an un-patchable surface **escalates** (the 4-state `unimplementable-without-platform-edit`). |
| **R2** | **"Compare candidates on the same simulation" needs a platform build.** If the comparison view is **`apps/hiring`-only** (Vercel, no Dockerfile, not in `/demo-up`), showing it = containerizing a Vercel app = **large net-new + a platform edit** = out of the zero-edit wall. | **blocks-release** | **Strong research evidence it's already in the dockerized `apps/web`** (`InsightsByMembersContainer` + `useGetInsightsByMembers`, route files present, `Dockerfile.dev` builds `@anthropos/web-app`). But **inferred, not render-proven** (BA-2/BA-3) → **M222's go/no-go make-or-break.** If `apps/hiring`-only: **escalate, do not proceed.** |
| **R3** | **Believability — 45 junk-or-identical assessments.** The M219 close found junk skills ("24-hour dietary recall", `15Five`) when a claimed pool ran dry and topped up from a flat pool's alphabetical head; a flat growth-arc would make 45 candidates **identically scored** (un-rankable). | **degrades-quality (the demo's whole point)** | Reuse the M219 fix: skill ladder **role→curated→general→STOP, never pad**; every skill/role ref through the real resolvers (**closure green, never fabricated**); a **realistic non-degenerate score DISTRIBUTION** across the 45 (the funnel writes differentiated per-candidate scores, not a flat arc); a **hiring believability manifest** in M225's coverage sweep (persona self-consistency: candidate role↔skills↔score). |
| **R4** | **Whole-org-hydration latency** — the compare surface hydrates 45×5 candidate-sim results; the exact class of the AI-readiness `loadMembers` p95 wall (180 s→19 ms) that needed a perf demo-patch. Adds a **3rd measured vantage** to v2.3's <5 s access gate. | **degrades-quality → blocks-milestone (M226)** | Measure before blaming code (v2.3 method). If real, a **sha-pinned perf demo-patch** (the sanctioned hatch), pinned at M224, **re-proven live at final code** (M221 discipline) at M226. Latency is *reported* at M224, *gated* at M226. |
| **R5** | **`isHiring` blast radius** — `apps/web` has **two** `isEnterprise` definitions (`Boolean(organization)` nav vs `!isHiringOrg && organizationId` billing); flipping `isHiring` may hide/relabel a surface the compare demo depends on. Blast radius **unverified**. | **degrades-quality** | M222 **enumerates** the surfaces that change vs stay identical under `isHiring`; couples to Q1 (if recruiter-only, we may **not** flip `isHiring` at all and dodge the blast radius entirely). |
| **R6** | **Scope creep into the Hiring *product* / ATS.** The user may expect job ladders / candidate funnels / post-a-role — that's `apps/hiring` (Vercel-only) + ATS tables the platform **deliberately doesn't build**. | **scope** | **Hard line, stated up front:** this demo shows the **Workforce app's `/enterprise` comparison surface** for an `is_hiring` org, **NOT** the Hiring product. M205's tier-gate/pipeline half stays out. Confirm with the user (Q1 framing). |
| **R7** | **Snapshot content starvation** (BA-6) — the cold-primed public snapshot may be TRAINING-dominated and lack ≥5 HIRING sims tied to positions. | **degrades-quality** | M222 probes the captured pool; fallback is a dedicated hiring-sim pattern query (`readAIReadinessSkillPool` precedent) or, worst case, synthesized sims that **still resolve real skill/role refs** (closure preserved) — the Q3 real-vs-synth decision. |

---

## 5. The 3 open questions most worth asking the USER before building

**Q1 — Recruiter-comparison-only, or ALSO the candidate-facing hiring UX?** *(the biggest scope fork.)*
The manager comparison surface is **admin/enterprise-gated** and *appears* `isHiring`-independent — so the
"compare 45 candidates on 5 positions" demo may need **no** `isHiring` wiring at all (dodging R5 entirely). But if
you also want the **org and the candidates to read as *hiring*** (hiring onboarding/home, the "hiring" assignments
variant, a candidate hero logging into a hiring-flavored profile), that requires extending Clerkenstein to emit
`publicMetadata.isHiring` **and** may surface an `isHiringOrg`-gate demo-patch (R1). *Recommend:* start
**recruiter-comparison-only** (lowest risk, fully in the dockerized app); add the candidate flavor as a fast-follow
if wanted. **This decision sizes M224 and decides whether Clerkenstein work is in scope.**

**Q2 — Which of the 5 managers + 45 candidates become cockpit heroes?** Only **heroes** get a loggable
Clerkenstein seat; the other 47 are population rows. *Recommend a trio mirroring the existing stories:*
**1 recruiter/manager hero** (`vantage: manager` → jumps onto the comparison surface) **+ 2 candidate exemplars**
(a **strong** and a **weak** candidate on the *same* position → `/profile`, to headline the side-by-side compare).
Confirm the count and whether any candidate should be loggable at all.

**Q3 — Real replayed content vs synthesized, AND full grid vs funnel?** *(two coupled content decisions.)*
(a) The 5 "job positions" can be **REAL replayed public content** (extend the snapshot to replay `directus.job_position`
rows + pin 5 real HIRING sims — 715 exist, 127 public) **or SYNTHESIZED**. *Recommend real* (renders real content,
preserves closure) — pending the M222 BA-6 confirm that the cold snapshot carries ≥5 usable HIRING sims. (b) Does
each of the 45 candidates take **all 5** sims (**225 scored sessions**, full comparability on every position) or a
**realistic funnel** (fewer candidates deep into later positions)? The user said "assessed around the SAME 5" —
*recommend the full grid* for the strongest compare demo, unless a funnel reads more believably.

---

## Appendix — traceability

- Research: 3-report workflow (platform hiring model · rext Stories&Heroes seeder · cockpit/enterprise integration),
  2026-07-15. Prior gap-analysis that first flagged this as HIGH-severity net-new mapped to reserved M205:
  `.agentspace/scratch/roadmap-research-2026-07-13.md` (D1/BD-1).
- Reverses v2.3 `D-DESIGN-4`; consumes the recruiter/seeder half of the reserved vision `M205`.
- Hard constraints carried from the v1.x/v2.x lineage, **unchanged**: **zero platform-repo edits** (a platform-source
  wall routes to a sha-pinned `demopatch` per `corpus/ops/demo/demopatch-spec.md`, or escalates — never a repo edit);
  all stack-operating tooling lives in `rosetta-extensions` (authored in `.agentspace/rosetta-extensions/`, tagged,
  consumed per-stack at a pinned tag).
- **Next step:** this is a Phase-6 proposal. It is **not** promoted to `roadmap.md` / `state.md` and **no branch is
  cut** until the user resolves §5 and approves (design-roadmap Phase 6 → 7).
