# Skillpath Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of the **"skillpath-in-app"** program (platform milestones **M502 → M507**), the standalone `skillpath` Go
> microservice has been **merged into the `app` monolith** (the service the platform calls "backend") and then
> **decommissioned**. Skillpath no longer runs as a separate service — not in the local compose, not in the
> supergraph, not in production. This is the same pattern as the earlier [skiller-in-app merge](./skiller.md);
> skillpath was the next runtime engine consolidated into `app`.
>
> **Skillpath was always a runtime/session engine, never a content store** — it tracks per-user progression
> *state* (`SkillPathSession → ChapterSession → StepSession`, progress %, completion). The skill-path **content**
> it tracks against (title, cover, curators, chapters → steps, skills-to-verify, versioning) **remains owned by
> [CMS](./cms.md)** and is fetched by ID over Connect-RPC (`CMS_RPC_ADDR`). The consolidation moved the *engine*
> into `app`; it did not touch the content-vs-runtime split.
>
> Where everything went:
>
> * **Domain / engine** — the session manager + repository (`SessionManager`, the get-or-create + version-upgrade
>   logic, the jobsimulation-event subscriber) now live inside `app`: `app/internal/skillpath/`
>   (`session.go`, `session_domain.go`, `repository/`) and `app/internal/skillpaths/`. Ported at **M502/M503**
>   (manager port, dormant) and the subscriber merged at **M504**.
> * **Data** — runtime session state now lives in the **`public` schema** of the shared PostgreSQL database:
>   `public.skill_path_sessions` (Ent schema `app/internal/data/ent/schema/skill_path_session.go` +
>   `skillpath_mixins.go`). The old `skillpath` DB schema is **legacy — a decommissioned empty husk** (the table
>   was kept but holds 0 rows; runtime state is authoritative in `public`). `askengine` and every other reader was
>   re-pointed `skillpath.skill_path_sessions → public.skill_path_sessions`.
> * **RPC** — the `SkillPathSessionService` surface (`GetSkillPathSession`) is served by `app`; callers were cut
>   over to read sessions **in-process** and the `SKILLPATH_RPC_ADDR` was dropped from terraform (**M506** caller
>   cutover).
> * **GraphQL** — the skillpath subgraph was **removed** from the WunderGraph/Cosmo federation → the supergraph is
>   now **3 subgraphs** (backend/app, jobsimulation, cms). The skill-path session types/queries/mutations
>   (`getOrCreateSkillPathSession`, `skillPathActiveSessions`, `skillPathCompletedSessions`,
>   `completeSkillPathStep`, `uncompleteSkillPathStep`, `upgradeSkillPathSessionToLatest`,
>   `upgradeAllSkillPathSessionsToLatest`, and the deprecated `createSkillPathSession`) are **folded into `app`'s
>   `backend` subgraph** (`app/internal/web/backend/graphql/graph/schemas/skillpath_sessions.graphqls`). The fold
>   landed dormant at **M505**; the router owner-swap that routes `SkillPathSession` to `app` was the atomic
>   **M506** cutover.
> * **Infrastructure** — skillpath was removed from `repos.yml` (now **10 repos**, 0 skillpath) and from
>   docker-compose (no skillpath service); the standalone service + its terraform module were decommissioned at
>   **M507**. Only residual env plumbing remains (e.g. the `SKILLPATH_STREAM=skillpath` Redis-stream name).
> * **Repo** — the `skillpath` git repo still exists but is **legacy/decommissioned**, no longer deployed or
>   cloned by `make init`.
>
> For current documentation of this domain, see [Backend (`app`)](./backend.md).

## Still-true domain knowledge

* **Content-vs-runtime split (unchanged).** "Skillpath" the engine ≠ skill-path *content*. The content it runs
  against — chapters → steps, curators, the job-simulation steps, skills-to-verify, versioning — is owned by
  **[CMS](./cms.md)** (the `skill_paths` Directus collection) and fetched by ID over Connect-RPC (`CMS_RPC_ADDR`).
  This is the content-vs-runtime split documented in the [Service Taxonomy](../architecture/service_taxonomy.md).

* **Session model.** The engine owns a hierarchical session: `SkillPathSession → ChapterSession → StepSession`,
  each carrying `progress` (0–100), `status` (`pending|active|completed|archived`), `duration`, `version`, and
  `*_at` timestamps (shared via `SkillPathMixin`). The player skill-path page reads this runtime via
  `getOrCreateSkillPathSession` — a **get-OR-create** that auto-materializes a blank `pending` session on first
  read, so an unseeded skill path renders empty, not 404.

* **Event-driven step completion (unchanged).** The engine subscribes to the **jobsimulation Redis stream**
  (start/end events) to update `StepSession` status when a simulation completes, and additionally calls
  jobsimulation over Connect-RPC (`GetSessions`) on session create/upgrade to reconcile already-completed
  simulations. It publishes `EventSkillPathSessionUpdated` + `EventChapterStepSessionCompleted` to the
  `skillpath` Redis stream (consumed by `app`). All of this now runs in-process inside `app`.

* **The manager view reads an `app`-side MIRROR, not this runtime.** The **manager insights** surface
  (`insightsSkillPathByMemberships`, the `/enterprise/activity-dashboard/@tabs/skill-paths/[skillPathId]`
  scoreboard in `apps/web`) does **not** read the runtime session — it reads the `app`-side mirror table
  **`public.local_skill_path_session`** (`app/internal/organization/intelligence.go`; Ent schema
  `app/internal/data/ent/schema/local_skill_path_session.go` — `progress` 0-100, `status`, no `score`), the
  analog of hiring's `local_jobsimulation_sessions` mirror. **Seeding only the runtime session rows renders an
  empty manager scoreboard** — the mirror row must be co-written. `apps/hiring` has no skill-paths tab
  (no-surface). Full treatment: [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md).

* **The per-user drill-down one level deeper is UNIMPLEMENTED** (verified against `next-web-app` `origin/main`).
  The mirror above powers the *cohort* scoreboard at `…/skill-paths/[skillPathId]`
  (`InsightsBySkillPathStudentsContainer`) — a real table that genuinely requires the mirror row. The
  **per-member** route `…/skill-paths/[skillPathId]/[userId]`
  (`InsightsBySkillPathStudentSimulationsContainer`) is **not built**: `userData` is hardcoded `null`, its
  results table and totals block are commented out, and the body renders the literal string **"Coming soon"**.
  No query touches the seeded session there, so the page is byte-identical whether or not you seed. This is why
  the v2.5 content-stories gate is **player-link-only** for skill-path (denominator corrected 31 → 29).

## Related Documentation

* [Backend (`app`)](./backend.md) — where the skillpath engine now lives
* [CMS](./cms.md) — the content side of the content-vs-runtime split (owns the skill-path definitions)
* [Jobsimulation](./jobsimulation.md) — the peer runtime engine whose completion events drive step completion
* [`../ops/demo/content-stories-routes.md`](../ops/demo/content-stories-routes.md) — the manager-mirror + player-link-only treatment
* [Service Taxonomy](../architecture/service_taxonomy.md) · [Dependency Map](../architecture/dependency_map.md)
