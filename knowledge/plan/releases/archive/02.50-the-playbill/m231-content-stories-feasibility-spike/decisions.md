# M231 — Decisions

## D1 — GO/NO-GO: Thread B is a GO (conditional per product)
The spike's central premise HOLDS: the simulation result page reads a **persisted DB row**, not a live recompute
(`jobsimulation/internal/graph/queries.resolvers.go:70` → plain Ent SELECTs of `validation_attempt_results` +
skill/criterion/check results; grading is an async submit-time worker the read path never invokes). A clone that
INSERTs the persisted fan-out renders a full result. **Verdict: GO** for the Simulation product (training/assessment/hiring)
and Skill-path; **conditional** for Interview (flag-gate); **AI-labs OUT**; **Academy IN**.

## D2 — Per-product render classification (the spike deliverable's spine)
| Product | Player | Manager | Verdict |
|---|---|---|---|
| Sim TRAINING/ASSESSMENT | renders-from-seed | renders-from-seed (`local_jobsimulation_sessions` mirror) | GO |
| Sim HIRING | renders-from-seed (`HiringResult` gate) | renders-from-seed (mirror; `apps/hiring` 2nd container, M224 eject) | GO |
| Sim INTERVIEW | renders-from-seed **+ PostHog-flag-gated** | renders-from-seed **+ flag-gated + admin-gated** | GO w/ demo-patch (D3) |
| Skill-path legacy | **runtime-computed-blank** (seedable via skillpath runtime rows) | renders-from-seed (`local_skill_path_session` mirror); `apps/hiring` = no-surface | GO |
| AI-labs | no result-render surface | no result-render surface | **OUT** (D4) |
| Academy | renders-from-seed (`academy_chapter_progress`) | (no manager result route; workforce academy insights TBD) | IN (D5) |

## D3 — INTERVIEW render requires the PostHog flags ON → demo-patch/flag-enablement  [Fate-3 → M232]
Both interview surfaces gate on `isExtractionEnabled = posthog.isFeatureEnabled('flag_interview_{player,manager}_report')`
(`AISimulationResultContainer.tsx:499-506`). A seeded `interview_extraction_results` row is **necessary but not
sufficient** — the flag must be ON in the demo, else the report hides. **Routing:** Fate-3 → **M232** (which builds the
interview substrate) must ensure the demo enables these PostHog flags (demo PostHog bootstrap, or a `demopatch` forcing
`isFeatureEnabled` true for the two interview flags). M235 (prove-it-lands) verifies the render. Recorded in M232's
overview under In. This is a normal spike output (a runtime-blank surface needing a demo-patch decision), SEVERITY warning.

## D4 — AI-labs ruled OUT (no seedable result-render surface)  [Fate-2 → roadmap already conditional]
`LabsAPIClient` is nil without `LABS_API_URL`; Create persists a `lab_sessions` row (status `booting`, no VM,
no `ide_url`/`preview_url`). `grade_result` is written only by an event FROM a live labs-api and is **not exposed by
GraphQL** at all. `/labs/[id]` reads LIVE from labs-api and throws without `LABS_API_URL`. → a seeded lab session cannot
render a "result". Wiring a live labs-api worker or adding a `gradeResult` GraphQL field are both out of the zero-edit
envelope (the latter is a platform edit → would ESCALATE). **Routing:** Fate-2 — the roadmap's M232 already excludes
AI-labs sessions "unless M231 ruled them feasible"; the spike rules them OUT, so M232's exclusion stands. **Scope
adjustment for M234:** the content-stories tab's planned "AI-labs" section should be **presence-only** (list the seeded
`lab_sessions` row as an activity/spend line in `/labs` + `/enterprise/labs`), NOT a played-result list. → Fate-3 → **M234**.

## D5 — Academy content-story is IN (backend-authoritative, seedable)  [depends on M230, Fate-2]
The "Clerk-only, no backend" premise is STALE: since ant-academy v0.5 M2 the academy WRITES progress to the platform
`app` academy backend over GraphQL. The "session" = per-user `academy_chapter_progress` + `academy_last_activity` rows,
seedable via `app/cmd/academy-seed` (fixtures starter/in-progress/completed). So the academy section renders REAL played
progress, not presence-only. **Dependency:** the chapters progress points at need CATALOG rows to render → depends on
**M230** academy demo-fill (Fate-2, already in the release). corpus fixed: `ant-academy.md` "no backend writes" claim
corrected inline (KB-7).

## D6 — Sourcing mechanism CONFIRMED viable; anonymization surface authored
- **Read path:** the wired `postgres` MCP / db-access read path selects interesting real prod sessions per type
  (confirmed live: ASSESSMENT 5,172 completed / TRAINING 1,799 / HIRING 1,679 / INTERVIEW 488 — all customer-scoped,
  all with persisted score + full child fan-out). **Pin by `jobsimulation.sessions.id` (uuid)** — the deterministic
  source-pin for the M232 manifest.
- **Public-anchoring rule:** the M232 sourcing query MUST inner-join `directus.simulations` on the public predicate so
  a cloned session's `sim_id` resolves in the demo (already snapshot-replayed). Ample public-anchored real sessions per
  product (ASSESSMENT 2,427 / TRAINING 549 / HIRING 395 / INTERVIEW 41-across-the-1-public-interview-sim).
- **Anonymization surface (classified without reading values):** structured IDs re-key/re-tenant; enums/numerics/
  timestamps kept; **free-text needs handling** = `actors.username/alias` (names), the LLM feedback fields on
  `validation_*`, `validation_criterion_results.input_data` (candidate submission), `interactions.action_payload`
  (transcript), `interview_extraction_results.{user_report,manager_report,summary}`. Full list in `content-stories-routes.md`.
- The actual copy + anonymize transform is M232 (Fate-2, already scoped). The release amends `safety.md` Part 3
  (anonymized-real, VPN-scoped) at M232 — confirmed here.

## D7 — Open question resolved: the platform's own `clone-session` subcommand
`jobsimulation cmd/clone_session.go` exposes `clone-session --session-id --user-id` → `CloneSession(ctx, sessionId, userId)`
— a platform-native session cloner that re-players a session to a NEW userId. **Invoking the built binary in-stack is
within the zero-platform-edit wall** (running the tool, not editing the repo). BUT it only re-players to a new userId —
it does NOT anonymize free-text or fully re-tenant (`organization_id`) — so M232 still needs a rext layer for
re-tenant + anonymization + the mirror co-write. Recorded for M232 (Fate-2 informs its build).

## D8 — The manager-view mirror trap generalizes beyond hiring
Both `local_jobsimulation_sessions` (sim/hiring manager) and `local_skill_path_session` (skill-path manager) are
`app`-side event-populated MIRRORS the manager scoreboards read — seeding only the runtime rows renders an empty
manager view. The content-stories seeder (M232) MUST co-write the mirror row for every manager-visible result. Documented
in `content-stories-routes.md`; pointer added to `skillpath.md` (KB-4).

## KB-fidelity findings (Phase 0b audit — full report in `kb-fidelity-audit.md`)
- **KB-1** `jobsimulation.md` silent on result read-model → consolidation home is `content-stories-routes.md` + pointer. Track.
- **KB-2** `jobsimulation.md` ports 8400/8401 vs actual :8080/:8081 (confirmed in repo CLAUDE.md). STALE, tangential. Track.
- **KB-3** `hiring.md` M228 "intercepting-route" → plain Ant `<Drawer>` on the leaf route. **FIXED INLINE.**
- **KB-4** `skillpath.md` missing the manager-side mirror. **FIXED INLINE** (pointer).
- **KB-5** `ai_architecture.md` code="via Roadrunner" → now in-process Judge0 (`jobsimulation/internal/runner/`). STALE. Track (deliverable carries correct fact).
- **KB-6** `roadrunner.md` describes a live jobsim consumer → orphaned. STALE. Track (architecture-doc pass).
- **KB-7** `ant-academy.md` "no backend writes at runtime" → backend-authoritative read/WRITE since v0.5 M2. **FIXED INLINE.**
- **KB-8** `backend.md` labs path `internal/labsession` → actual `internal/labs/session`; `grade_result` not GraphQL-exposed. STALE. Track.
