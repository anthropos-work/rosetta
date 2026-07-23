# M247 — Spec notes

Working notes per scope lane — the concrete file lists, redirect wording, fact-sheet outlines, and grep-verify
results captured as the re-ground lands.

## skillpath → merged-into-app redirect

Ground truth (current `stack-demo/app` @ origin/main + app CHANGELOG):
- **"skillpath-in-app" program, platform M502 → M507** (mirrors "skiller-in-app" v2.1). CHANGELOG anchors:
  M502 ent port → public; M503/M503b manager port + shadow-parity; M504 subscriber merge; M505 fold GraphQL
  surface into app subgraph (dormant); M506 caller cutover (in-process reads, drop skillpath RPC) + atomic
  router owner-swap; M507 decommission standalone service (commit `01fd340`).
- **Engine code** → `app/internal/skillpath/` (session.go, session_domain.go, repository/) + `app/internal/skillpaths/`.
- **Data** → `public.skill_path_sessions` (Ent `app/internal/data/ent/schema/skill_path_session.go` + `skillpath_mixins.go`);
  legacy `skillpath` schema = empty husk (0 rows; 561 rows in `public` at M246). askengine + all readers re-pointed.
- **GraphQL** → folded into `app`'s **backend** subgraph
  (`app/internal/web/backend/graphql/graph/schemas/skillpath_sessions.graphqls`, header: "M505 … folded into
  app's backend subgraph … router owner-swap at M506"). Subgraph count 4 → **3** (backend, jobsimulation, cms) —
  verified against `graphql-wundergraph/supergraph-config-{compose,prod}.yaml`.
- **Repos/compose** → skillpath ABSENT (repos.yml 10 repos; no compose service). Residual `SKILLPATH_STREAM=skillpath`.
- **Manager mirror** (durable, referenced by content-stories) → `public.local_skill_path_session`; per-user
  drill-down UNIMPLEMENTED ("Coming soon") — preserved in the redirect's Still-true section.

Delivered: `corpus/services/skillpath.md` → redirect (mirrors `skiller.md`); README moved skillpath from the
Tier-1 table to the archived/merged table. Inbound links (7 files, all plain no-anchor) still resolve.

## 3-subgraph reclassification sweep (~30 echo files)
_(TBD during build.)_

## Fact sheet — coursebuilder.md
_(TBD during build.)_

## Fact sheet — ai-labs.md (AI Labs + credits / v6.0 shared purse)
_(TBD during build.)_

## Fact sheet — askengine.md (Talk-to-Data)
_(TBD during build.)_

## Fact sheet — academy-backend.md (app-owned academy domain)
_(TBD during build.)_

## ai-readiness.md refresh (aireadiness-package refactor)
_(TBD during build.)_

## roadrunner ORPHANED→ARCHIVED resolution
_(TBD during build.)_
