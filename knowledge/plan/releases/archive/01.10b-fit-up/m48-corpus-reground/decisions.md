# M48 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1 — AI-readiness gets its own doc (`corpus/services/ai-readiness.md`), not a backend.md section
**Context:** the feature was undocumented; it needed a home. **Options:** (a) NEW standalone `services/ai-readiness.md`;
(b) a big section inside `backend.md`. **Choice:** (a) standalone, **+** a callout-with-link from `backend.md`'s
Recent-Feature-Additions. **Why:** it carries architectural weight (9 tables, a GraphQL schema, ~10 REST handlers, a
worker task, a member flow + a manager dashboard) and is the **M51 seeder contract** — it earns discoverability as a
top-level service doc (like jobsimulation/skillpath). backend.md keeps the one-paragraph "this subsystem lives here"
pointer.

## D2 — the seed STRATEGY (signals-true vs snapshot-direct) is M51's call (Fate 2)
**Context:** the doc's seeding contract notes two ways to seed: (a) write the underlying signals
(`user_skill_evidences` + jobsim sessions) and let the engine compute; (b) write `ai_readiness_user_step_progress`
+ upsert `ai_readiness_live_snapshots` directly. **Choice:** document BOTH in the contract; **defer the pick to
M51** (the AI-readiness showcase-org milestone — it owns the seeder). **Why:** it's a seeder-design decision, not a
doc decision; M51's `overview.md` already owns "build the AI-readiness seeder" → Fate 2, no edit needed.

## D3 — M48 fixes the ant-academy DOC truth; M49 #5 fixes `repos.yml` (Fate 2 split)
**Context:** CLAUDE.md + `services/ant-academy.md` (×3 spots) claimed ant-academy is "in `repos.yml` / cloned by
`make init`" — **false** (`stack-demo/platform/repos.yml` has 13 repos, no ant-academy). **Choice:** M48 corrects
the DOC to current reality (NOT in repos.yml; `make init` doesn't clone it) + flags **M49 #5** as the owner of the
actual `repos.yml` fix. **Why:** doc-truth is M48's job; the code fix (adding the repos.yml entry) is already in
M49's `In:` list → Fate 2. After M49 lands, M49 flips these doc claims to "Yes". (Accepted one-step doc churn —
each milestone's doc reflects the state at that milestone.)

## Adversarial review (Phase 2c, close)
- **Scenario — does the doc's seed contract actually make the manager dashboard render?** A seeder that writes
  `ai_readiness_live_snapshots` directly (the "snapshot-direct" path the draft doc offered) could leave the dashboard
  **empty**. **Verified against code:** the **active-cycle** dashboard `GetAIReadinessWithOptions` →
  `buildLiveResponse` → `computeOrgBreakdowns` (`ai_readiness.go:283-343`) **recomputes from signals**
  (`user_skill_evidences` + readiness jobsim sessions + the skills/sims config), and `keepStartedMembers` excludes
  members with no step-1 signal; `live_snapshots` is a materialized cache (`RefreshLiveSnapshots`) for Talk-to-Data
  SQL, **not** the dashboard source. Only the **closed-cycle** path (`buildResponseFromSnapshots`) reads frozen
  `ai_readiness_snapshots` directly. **Fixed the doc** (Fate 1): the seed strategy is now cycle-state-dependent —
  active ⇒ signals-true; closed ⇒ snapshot-direct. This sharpens the M51 contract + prevents a broken seeder.
- Corroboration: the platform repo's own KB (`app/knowledge/ai-readiness/overview.md`) confirms the model — added
  as the authoritative cross-reference in `ai-readiness.md`.

## Phase 1b — deferral re-audit verdict: GREEN (inline)
Two routings, both valid **Fate 2** (confirmed against the target overviews): **D2** (seed strategy → M51 — its
`overview.md` owns "the 3-step onboarding/evaluation seeder" + the enablement) + **D3** (ant-academy `repos.yml`
fix → M49 #5 — its `overview.md` lists the `repos.yml` entry). No repeat-deferrals; no escape-hatch. Ran inline
(docs-only milestone, trivial surface) rather than spawning `/developer-kit:audit-deferrals`.

## D4 — material-lag-first scope (the broad sweep was bounded)
The drift survey found the corpus broadly accurate EXCEPT the AI-readiness blind spot (critical) + the ant-academy
stale claim (material) + a few recent-feature omissions (minor: backend.md / next-web-app.md / architecture_overview
/ service_taxonomy — all fixed). No deep re-write of the per-service docs was warranted (they match current code on
the spot-checks). Scope held to material lag per the overview; no escape-hatch deferrals.
