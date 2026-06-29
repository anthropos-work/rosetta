# M48 — progress

## Section checklist

### S1 — Drift survey (corpus/architecture + corpus/services vs current clones) ✅
- [x] surveyed `corpus/architecture/*` + `corpus/services/*` vs the M47-current clones (drift-survey agent) — gap list in `spec-notes.md`
- [x] flagged **member-AI-readiness** as the headline undocumented addition (→ S2)
- [x] recorded the gap list in `spec-notes.md`

### S2 — Member-AI-readiness flow doc (LOAD-BEARING — the M51 seeder contract) ✅
- [x] mapped the data model: 9 `ai_readiness_*` tables + the `organization_settings` gate + the 3-step scoring (30/40/30) + stage/score states (backend agent)
- [x] mapped the surfaces: manager dashboard (4 tabs + funnel) + member onboarding hero (frontend agent)
- [x] authored **NEW `corpus/services/ai-readiness.md`** — role, enablement, 3-step+scoring, data-model code-map, GraphQL+REST interface, surfaces, narratives, **+ the M51 seeding contract**
- [x] cross-linked from CLAUDE.md + backend.md + architecture_overview.md + next-web-app.md + service_taxonomy.md

### S3 — Reconcile the surveyed material drift ✅
- [x] backend.md — AI-readiness added to Recent Feature Additions (+ link)
- [x] architecture_overview.md — Backend/App responsibilities mention AI-readiness (+ link)
- [x] next-web-app.md — Workforce product row mentions the AI-readiness UI (+ link)
- [x] service_taxonomy.md — Backend/App row mentions the AI-readiness subsystem (+ link)
- [x] no deep per-service rewrite needed (docs matched code on spot-checks — D4, material-lag-first)

### S4 — Retire stale claims + cross-link ✅
- [x] ant-academy "in repos.yml / cloned by make init" claim corrected — CLAUDE.md + 3 spots in `services/ant-academy.md` (it's NOT in repos.yml today; **M49 #5 owns the repos.yml fix** — D3)
- [x] rext-tag staleness — NOT fixed here (M49 #1 owns the `.agentspace/rext.tag` source-of-truth); noted, no M48 edit
- [x] all new/edited cross-references verified resolving (fences balanced)

## Notes
- Docs-vs-code only — never touched the live demo → M48 ∥ M49 (clean).
- Zero escape-hatch deferrals. D2 (seed strategy) + D3 (repos.yml fix) are Fate-2 (M51 / M49 already own them).
