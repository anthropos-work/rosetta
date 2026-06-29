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

## D4 — material-lag-first scope (the broad sweep was bounded)
The drift survey found the corpus broadly accurate EXCEPT the AI-readiness blind spot (critical) + the ant-academy
stale claim (material) + a few recent-feature omissions (minor: backend.md / next-web-app.md / architecture_overview
/ service_taxonomy — all fixed). No deep re-write of the per-service docs was warranted (they match current code on
the spot-checks). Scope held to material lag per the overview; no escape-hatch deferrals.
