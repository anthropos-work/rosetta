# Content Stories — per-product result-route map + prod-session sourcing contract

**The M231 feasibility-spike deliverable (v2.5 "the playbill", HARD go/no-go).** Before ANY Thread-B build
(the 2nd "Content stories" cockpit tab of played sessions per content product), this doc DISCOVERS + PROVES,
by code-cite and prod-DB evidence, the one question that gates the whole chain: **for each content product's
result page, does it render from a PERSISTED row a clone could seed, or recompute LIVE (unseedable → blank)?**

> **Verdict headline — Thread B is a GO** (conditional per product). The simulation result page reads a
> **persisted DB row**, not a live recompute — so a clone that INSERTs the persisted result fan-out renders a
> full result. Simulation (training/assessment/hiring) + Skill-path are GO; **Interview** is GO behind a
> PostHog-flag demo-patch; **AI-labs is OUT** (no seedable result-render surface); **Academy is IN** (backend-
> authoritative progress, seedable). All discovery is code-cited against the local platform checkout + prod-DB
> reads (read-only, public-vs-customer boundary honored — [`db-access.md`](../db-access.md)).

## For PMs — one paragraph

A "content story" is a *played session* a presenter can log into and see the result of — as the player who took
it, and as the manager who reviews it. The spike asked whether we can build those by copying real (anonymized)
production sessions into a demo. The answer is **yes for most products**, because the platform **stores** each
result in the database (the score, the skill breakdown, the feedback, the interview reports) and the result page
just **reads** those stored rows — it does not re-run the AI grader when you open the page. So a seeder that writes
the same rows produces the same result screen. The exceptions: **AI-labs** has no result screen to render at all
today (its grade isn't even exposed to the front-end), and **Interview** needs two feature-flags turned on in the
demo. **Academy** turned out to be seedable too — contrary to older docs, the academy now saves your progress to
the platform backend, so a "played" academy state is just a set of rows.

---

## 1. Per-product result-route map (player + manager) — prove-by-render classification

**Classification legend** (per the milestone's four-way probe):

| class | meaning | seedable? |
|---|---|---|
| **renders-from-seed** | the page reads a PERSISTED row; INSERT the row → it renders | ✅ yes, directly |
| **runtime-computed-blank** | the read path auto-materializes an empty state (or recomputes); an unseeded page renders blank, not 404 | ⚠️ seedable via the runtime's own persisted rows |
| **needs-demo-patch** | a persisted row is necessary but a config/flag gate hides it in a demo → a sha-pinned `demopatch` or flag-enablement is required | ⚠️ seed + patch |
| **no-surface** | there is no result route for this vantage | ❌ |

### The matrix

| Product | Player route | Player class | Manager route | Manager class |
|---|---|---|---|---|
| **Sim TRAINING** | `apps/web/…/sim/[slug]/result/[sessionId]/page.tsx` | renders-from-seed | `apps/web/…/enterprise/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx` (+`[userId]`) | renders-from-seed |
| **Sim ASSESSMENT** | same as TRAINING | renders-from-seed | same as TRAINING | renders-from-seed |
| **Sim HIRING** | `apps/hiring/…/sim/[slug]/result/[sessionId]/page.tsx` (`isHiring`, `HiringResult` gate) | renders-from-seed | `apps/hiring/…/@tabs/ai-simulations/[simId]/page.tsx` (+`[userId]`) | renders-from-seed |
| **Sim INTERVIEW** | `apps/{web,hiring}/…/sim/[slug]/result/[sessionId]` → `interviewExtractionUserReport` | **needs-demo-patch** (flag) | `…/@tabs/interviews/[simId]/[userId]/page.tsx` → `interviewExtractionManagerReport` (admin-gated) | **needs-demo-patch** (flag) |
| **Skill-path legacy** | `apps/web/…/skill-path/[skillPathId]/page.tsx` → `getOrCreateSkillPathSession` | **runtime-computed-blank** | `apps/web/…/@tabs/skill-paths/[skillPathId]/page.tsx` (+`[userId]`) | renders-from-seed |
| **Skill-path new (academy)** | `aiacademy.anthropos.work` chapter page (progress-driven) — see §6 | renders-from-seed (progress rows) | — (no academy manager result route today) | no-surface |
| **AI-labs** | `apps/web/…/labs/[id]/page.tsx` (reads LIVE from labs-api) — see §5 | **no-surface** (for a seeded result) | `apps/web/…/enterprise/labs` (activity/spend listing only) | no-surface (for a result) |

All simulation surfaces render through one shared component: `next-web-app/packages/ui/src/AISimulation/AISimulationResultContainer.tsx`.

### The central unknown — RESOLVED: simulation result = a PERSISTED READ, not a live recompute

- The render query is `GET_SIMULATION_RESULT` / `jobSimulationResult(sessionId)`, resolved in the **jobsimulation**
  subgraph at `jobsimulation/internal/graph/queries.resolvers.go:70`.
- The resolver does **plain Ent SELECTs**: it reads `validation_attempt_results.evaluation_status` straight off the
  persisted row, plus `validation_attempt_skill_results`, `validation_criterion_results`
  (`WithValidationChecksResults`), and `anticheat_results`. **No engine call, no LLM grading pass, no session
  replay on render.** Grading is an async Asynq worker that runs at *submit* time; the read path never invokes it.
- Front-end behavior: while `evaluationStatus === Pending` the page **polls every 2 s** (a spinner — this is what a
  *unseeded* session shows, forever); any terminal status stops the poll and renders. `passed` is derived
  client-side (`evaluationStatus === Passed`); `score` comes from the persisted `session.score`.
  `useRecalculateEvaluationResult` is wired ONLY to the retry button (a user-initiated mutation), never to render.
- Independent corroboration: the repo-local `seed-verified-skill` skill already lights up this exact UI by direct
  INSERT of `jobsimulation.sessions` + validation rows; and prod DB shows every completed session carries a
  persisted `score` + `result_status='completed'` + the full child fan-out (below).

**Result substrate to reconstruct per session** (all `jobsimulation` schema unless noted; verified per-type in prod):

```
sessions (score, status, completion_status, result_status, ended_at, validation_version[2|3])
  └─ validation_attempt_results (evaluation_status ← THE gate, score, success_threshold, *_summary)
       ├─ validation_attempt_skill_results (skill NodeID, score, competency_level_score, *_feedback)
       └─ validation_criterion_results (type=evaluation, title, skills, score, input_data, *_feedback)
            └─ validation_check_results (check_id, success, feedback, essential)
  ├─ actors (2–3 per session: the candidate + stakeholders)
  ├─ interactions (the transcript — action_type + action_payload)               [voice/chat/etc.]
  ├─ interview_extraction_results (user_report + manager_report JSON)           [INTERVIEW only]
  └─ public.local_jobsimulation_sessions  ← the MIRROR (in app, NOT jobsimulation) — see below
```

Prod-verified per-type (completed, org-scoped, sampled by pinned id): ASSESSMENT = 1 var + 3 actors + mirror
(score-0 no-shows have 0 var); HIRING = 1 var + 3 skill + 6 criterion + 2 actors + mirror; INTERVIEW = 1 var +
1 skill + 4–5 criterion + 2 actors + **1 interview_extraction_results** + mirror; TRAINING = 1 var + 1–3 skill +
4–5 criterion + 2–3 actors + mirror.

### The manager-view MIRROR trap — generalized (M219/M222, now beyond hiring)

The manager scoreboards do **NOT** read the runtime tables — they read an **`app`-side, event-populated MIRROR**:

- **Sim manager** (`insightsJobSimulationByMemberships`, `app/internal/web/backend/graphql/graph/resolver_queries.go:1088`,
  Casbin-gated `OrgFeatureInsights`) reads `app.public.local_jobsimulation_sessions.score`
  (`app/internal/organization/intelligence.go:1692/1801`; Ent `local_jobsimulation_session.go:52`), populated by
  the `app` Redis-Stream consumer (`JobSimulationSubscriber` → `updateOrCreateLocalSession`), NOT by writing
  `jobsimulation.sessions`.
- **Skill-path manager** (`insightsSkillPathByMemberships`, `resolver_queries.go:977`, same gate) reads
  `app.public.local_skill_path_session.progress` (`intelligence.go:997/1142`), the exact analog.

**⇒ seeding only the runtime rows renders an EMPTY manager scoreboard.** Every manager-visible result MUST co-write
the mirror row. In prod the mirror is ~1:1 with the source (`local_jobsimulation_sessions` 19,870 ≈
`jobsimulation.sessions` 19,873). This is the single sharpest seeding landmine for the manager vantage. (The
mirror trap is a **different surface** from the player result page in §1's proof — the player page reads
jobsimulation's own tables directly and needs no mirror row.)

### INTERVIEW is flag-gated → needs-demo-patch

Both interview surfaces gate on
`isExtractionEnabled = isInterview && (posthog.isFeatureEnabled('flag_interview_manager_report') || posthog.isFeatureEnabled('flag_interview_player_report'))`
(`AISimulationResultContainer.tsx:499-506`). A seeded `jobsimulation.interview_extraction_results` row (one row,
`user_report` + `manager_report` co-stored as JSON, resolvers `queries.resolvers.go:536`/`:563`) is **necessary
but not sufficient** — if the flags aren't ON in the demo's PostHog, the report hides. **Decision (M231 D3, routed
Fate-3 → M232):** the demo must enable the two `flag_interview_{player,manager}_report` flags (demo PostHog
bootstrap, or a sha-pinned `demopatch` forcing `isFeatureEnabled` true). This does NOT escalate — no platform edit.

### Skill-path player is get-or-create (runtime-computed-blank)

`skill-path/[skillPathId]/page.tsx` → `getOrCreateSkillPathSession` (`packages/graphql/src/query/skill-path.ts:404`)
is a **get-OR-create** that federates to the skillpath runtime; on read with no session it **auto-materializes a
fresh `pending` session at progress 0**. So an unseeded skill path renders *blank*, not 404. To show a meaningful
player result, seed the persisted skillpath runtime rows (`skillpath.skill_path_session` / `chapter_session` /
`step_session`) — progression state IS persisted; the read path just fabricates a blank one if absent.

### Structural correction — there are NO Next.js intercepting routes

Verified: **zero** `(.)`/`(..)`-prefixed dirs exist anywhere in `apps/`. The recruiter comparison "drawer" is a
plain Ant `<Drawer>` (`InsightsByMembersContainer.tsx:359`) rendered on the ordinary `[simId]/page.tsx` leaf route
— not an intercepting route. (The stale `hiring.md` M228 "intercepting-route-aware" claim was corrected in this
milestone.)

## 2. Manager-view eligibility matrix — which products HAVE a manager result route

| Product | Manager result route? | Read-model | Notes |
|---|---|---|---|
| Sim TRAINING / ASSESSMENT | ✅ yes | `local_jobsimulation_sessions` mirror | `apps/web` activity-dashboard → ai-simulations tab |
| Sim HIRING | ✅ yes | same mirror | `apps/hiring` only (a genuine hiring org **ejects** from `apps/web` → `apps/hiring`, `UserStatusContext.tsx:168-169`, M224) |
| Sim INTERVIEW | ✅ yes (flag+admin-gated) | `interview_extraction_results.manager_report` | admin gate `OrgActionAssignmentsWrite`; PostHog `flag_interview_manager_report` |
| Skill-path legacy | ✅ yes (`apps/web` only) | `local_skill_path_session` mirror | `apps/hiring` = **no-surface** (no skill-paths tab) |
| Skill-path new (academy) | ❌ no manager result route today | — | academy has no manager review surface (workforce academy insights TBD) |
| AI-labs | ❌ no (activity/spend listing only) | — | `grade_result` not GraphQL-exposed (§5) |

→ For the cockpit's per-session `as-manager` CTA, honor `has_manager_view`: TRUE for the four sim manager routes +
skill-path-legacy-in-`apps/web`; FALSE (omit the CTA) for academy + AI-labs.
