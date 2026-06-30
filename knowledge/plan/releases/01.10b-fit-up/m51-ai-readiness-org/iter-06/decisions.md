# iter-06 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1: triage — the iter-06 re-sweep residual is PERF-not-content (data confirmed in demo-1 DB) — 2026-06-30

**Context:** the iter-06 GATED re-sweep (with the committed iter-05 harness in effect, report generatedAt
21:29Z) read `failingSections=5, escapes=0, persona=0, reachable=66, frontier EXHAUSTED`. The 5:
`verification-funnel` + `talent-languages` (/enterprise/workforce), `ai-readiness-org-score` +
`ai-readiness-funnel` (/enterprise/workforce/ai-readiness), `activity-table` (/enterprise/activity-dashboard).

**The iter-05 org-agnostic manifest correction WORKED:** the 2 prior `cervato-systems.com` false-fails
(members-roster, assign-roster) are GONE from the residual — the Northwind roster now passes org-agnostically.
The 2 NEW failures are the iter-05 AI-readiness manifest descriptor I added (the milestone's headline
assertion) — they're a real, asserted residual now, not a regression.

**Diagnosis (per the protocol's DOM+network+log probe):**
- `/enterprise/workforce/ai-readiness` screenshot: the page renders the "AI Readiness" title + "Northwind
  Aviation" org + Dana (chrome + auth + org all correct) but the body is a skeleton/spinner — captured
  mid-load. The page is NOT in the `warmHeavyGrids` set (the warm primes /enterprise/members, activity-dashboard,
  workforce, assignments — but NOT the ai-readiness route), so it pays the FULL cold backing-query cost at the
  authoritative visit.
- `/enterprise/workforce` (Workforce Intelligence) screenshot: renders ALL the structure (title, the Live badge,
  every tab — Growth/Skills&Verification/Talent Pool/Assignments/Activity Log — and every card header: Skills
  growth over time, Top verified/mapped skills, Roles building expertise, …) but every DATA region is a
  skeleton. It IS in the warm set, but it has MULTIPLE independent per-card aggregate queries; the warm's
  `hasRealRows` returns true as soon as ANY real row paints (the tab nav / a card chrome) and stops early, so the
  per-card data queries are still cold at the authoritative visit. verification-funnel + talent-languages live on
  SUB-TABS (Skills&Verification / Talent Pool) — not even rendered in the default Growth view.
- DB confirms the data: demo-1 `public.ai_readiness_*` = 1 active cycle, 8 skills, 2 sims, 3 steps, **532
  completed user_step_progresses** (skill_mapping 199 → simulation 177 → interview 156 — a realistic ~78% funnel),
  live_snapshots=0 (correct — active cycle recomputes from signals). The graphql request log shows a heavy
  aggregate at **15.83s** cold (most queries sub-200ms; the org-scale aggregates are the slow tail).

**Choice:** classify all 5 as the org-scale cold-grid PERF-WALL (slow-not-empty, data-in-DB), addressable by
the harness WARM (cache-priming) — NOT a seed gap, NOT a content gap.

**Why:** the screenshots show full structure + correct chrome/org/auth + the DB has the data; the only missing
piece is the data REGIONS, which are skeleton because the cold aggregate query hasn't resolved at capture. This
is exactly the M46 "org-scale grid perf wall (slow GraphQL, not empty)" class — the documented fix is
cache-priming during warm so the authoritative visit reads hydrated.

## D2: the fix — extend warmHeavyGrids (add the ai-readiness route + wait for skeletons-to-CLEAR, not first-real-row) — 2026-06-30

**Context:** D1 shows two warm gaps: (a) the ai-readiness route isn't warmed at all; (b) the existing warm
early-stops on the first real row, so a multi-card page's per-card cold queries are still loading at the
authoritative visit.

**Options:** (a) widen the per-section authoritative re-assert poll to brute-force the slow grids — REJECTED:
the protocol explicitly bans "widen the poll to mask a slow query"; that masks slowness at assert-time rather
than priming the cache. (b) shrink the org below 200 — REJECTED: the protocol bans "shrink the org below the
org-scale premise just to pass" (and 200 IS the gate's premise). (c) extend the WARM (cache-priming): add
`/enterprise/workforce/ai-readiness` to the warm set + make `warmHeavyGrids` wait for the page to be
SUBSTANTIALLY hydrated (skeletons cleared / a stable real-row count), not just the first real row.

**Choice:** (c). Add the ai-readiness route to `warmHeavyGrids`; change `hasRealRows`→a "substantially
hydrated" check that requires the page's skeletons to have CLEARED (real rows present AND skeleton rows
absent/stable) so the warm doesn't bail while per-card aggregates are still cold; keep the generous ceiling +
best-effort/bounded property (a grid that never hydrates within the ceiling is left as-is → the authoritative
assert still FAILs honestly — the warm never masks a genuinely-empty grid).

**Why:** this is cache-priming, the protocol-sanctioned lever for the org-scale perf wall (it makes the cold
query RESOLVE + CACHE during warm so the authoritative visit reads a hydrated grid) — fundamentally different
from widening the assert-time poll (which the protocol bans because it masks slowness without proving the data
rendered in budget). It STRENGTHENS the gate: a page that's genuinely empty/slow-beyond-ceiling still fails. It
stays honest because the authoritative assert is unchanged — only the warm (the cache-primer) is deepened.

## D3: the deepened warm did NOT lift — root cause is a PLATFORM-side AI-readiness response-build perf wall (RE-SCOPE / user-blocker) — 2026-06-30

**Context:** the iter-06 re-sweep WITH the deepened skeletons-cleared warm + the ai-readiness route added to
the warm set held at `failingSections=5` UNCHANGED — the same 5 (verification-funnel + talent-languages on
/enterprise/workforce; ai-readiness-org-score + ai-readiness-funnel; activity-table). The warm did not clear
the AI-readiness page.

**Deep diagnosis (the decisive evidence):**
- `GET /api/workforce/ai-readiness` **NEVER completes** — across the ENTIRE demo-1 backend log there is NOT a
  single completed GET (only OPTIONS preflights). The React Query fetch is aborted (signal) every time the
  page's cold query outruns the budget.
- The `ai_readiness_refresh` BACKGROUND worker (which calls the same `computeOrgBreakdowns` + snapshot upsert)
  fails repeatedly with `context deadline exceeded` (4× in the log). So the slowness is server-side, not a
  client/harness artifact — and it's why `ai_readiness_live_snapshots` is empty (0 rows): the materializer
  itself times out.
- ALL the individual AI-readiness SQL queries are ms-fast: `queryReadinessSimScores` EXPLAIN ANALYZE = **1.4ms**
  (jobsimulation.sessions is 1579 rows, fully indexed incl. a `sim_id,organization_id,status` composite);
  `queryUserAISkills` is a simple indexed `ANY()` scan; `loadMembers`/`hydrateMembers` are batch (no per-member
  RPC). So this is NOT the M46 index-bound members-grid wall — post-seed FK indexes won't help.
- The handler path: with NO CycleID (the default dashboard call) `GetAIReadinessWithOptions` ALWAYS takes
  `buildLiveResponse` → `computeOrgBreakdowns` (the live recompute) — it does NOT read the materialized
  `ai_readiness_live_snapshots` on the default call. So even populating the snapshot mirror from the seed would
  NOT short-circuit the default dashboard GET (the snapshot path is gated behind a CLOSED `CycleID`).
- skiller logs during the slow window show a STORM of `_entities resolving Entity "Skill": get skill
  translation <uuid>/english ... context canceled` — the response-build resolves skill-name translations
  through the federated GraphQL `_entities` path (the `withSkillerLang` translation resolution over the
  aggregate's skill set), a per-skill round-trip fan-out that (combined with the live recompute) blows the
  request deadline. Same N+1 CLASS as the M46 per-object Sentinel RPC, but in the AI-readiness translation path.

**Choice:** classify the residual as a **PLATFORM-side AI-readiness response-build perf wall** (the live
recompute + the per-skill federated translation fan-out exceed the request deadline at 200 members), and
ESCALATE per the milestone's `Re-scope trigger` — NOT a seed/harness/index fix.

**Why this is the milestone's declared escalation:** the M51 overview `Re-scope trigger` reads *"If enabling
the dashboard or seeding the funnel cannot be done without a platform edit → escalate
(`unimplementable-without-platform-edit`), never edit the platform."* The only fixes that would make the GET
complete in-budget are platform-read-path edits: (a) make the default dashboard call read the materialized
`ai_readiness_live_snapshots` (a snapshot short-circuit on the active-cycle default), OR (b) batch the per-skill
translation resolution (a DataLoader, the M46 platform finding's twin), OR (c) a new demo-patch to the
build-scratch `app` clone relaxing the translation path (the `app-targetrole-authz-skip` precedent — but a
NEW, substantial read-path demo-patch, a tooling investment that is a USER decision, not an inline iter fix).
The invariant is explicit: zero platform-repo edits. So iter-06 escalates the decision rather than editing the
platform or building a major new demo-patch unilaterally.

**Note on what iter-05+iter-06 DID achieve (kept, not reverted):** the org-agnostic email correction cleared
2 real false-fails; the AI-readiness manifest descriptor now ASSERTS the milestone's headline deliverable (so
the gate PROVES it when it renders); the cockpit jump_to lands the manager hero on the dashboard; the
content-presence/skeletons-cleared `warmHeavyGrids` is a genuine harness strengthening (it WILL help any
heavy grid whose wall is cold-cache rather than a hard server deadline). The 5 residual sections are the
platform perf wall, surfaced + root-caused, escalated for the user's platform-vs-demo-patch decision.
