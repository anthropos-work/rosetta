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

**Canonical reclassification applied** (each mention re-read in context, not blind sed):
- Supergraph = **3 subgraphs**: `backend` (app), `jobsimulation`, `cms`. skillpath removed from every subgraph
  list; "4 subgraphs" → "3". jobsim-in-app-coming note added where relevant (jobsim still standalone today).
- skillpath = **not a live service** — engine merged into `app` (M502→M507). Removed from live-service tables,
  port lists, compose profiles, depends_on lists, repo-clone lists, shared-lib importer lists; added to
  archived/merged tables. The skill-path engine's Redis-stream + CMS-RPC deps are now `app`-internal.

**Files edited (grep-verified 0 residual "4 subgraphs" / live-skillpath claims):**
- Architecture: `architecture_overview.md` (PM list, tier count 8→7, mermaid, service table→archived, content-vs-runtime
  callout, 4→3 ×4, schema-separation), `external_services.md` (gateway purpose + federation prose + mermaid 4→3 +
  depends_on + Dockerfile COPY + routing table + `docker compose ps`), `service_taxonomy.md` (service table→archived
  ×2, callout, subgraphs, aggregates, profile table), `dependency_map.md` (matrix rows, callout, shared-lib table,
  Redis-streams table, Flow 5), `shared_libraries.md` (colony/proto/authn/taxonomy importer lists + proto contract
  list), `architecture/README.md` (callout).
- Services: `README.md` (skillpath row Tier-1→archived done in §1; graphql one-liner 4→3), `graphql-wundergraph.md`
  (primary goal, key functions, schemas line, routing table, downstream/depends_on/CI lists), `backend.md`
  (federation 4→3 + live-de-risk note, RPC/consumer lists, downstream deps, Redis streams, Related), `cms.md`
  (callout, RPC/federation, upstream consumers, Related), `jobsimulation.md` (role, RPC consumers, dead-section),
  `sentinel.md` (caller list ×2), `messenger.md` (depends_on, `SKILLPATH_RPC_ADDR` gone-from-compose), `clerk-integration.md` (×2).
- Root: `CLAUDE.md` (app bullet gains skill-path engine + new domains, CMS NB, remove skillpath bullet, add to
  archived/merged + jobsim-next note, 4→3 ×2, safety disarm-list), `README.md` (core-backend row).
- General ops: `run_guide.md`, `platform_repo.md` (12→11 services, profiles, migrations, port table), `setup_guide.md`
  (repo table, starts list, migrations, migrate-dev schema list), `staging-sync.md` (15→14 repos, skip-worktree,
  Dockerfile.dev, schema map, migrate loop), `staging_from_dump.md` (clone list, consuming-service list),
  `staging-bringup.md` (repo dir tree), `safety.md` (demo authz-disarm list).

**Deliberately LEFT (not stale live-service claims):**
- `studio-desk.md` `/api/skillpath` route (studio-desk's own skill-path *builder* route → Directus; carved out for
  M249/M252/M253; not the skillpath service).
- `staging-bringup.md:241` + `staging_from_dump.md:113` role-does-not-exist dump quirk list (historical DB role
  names alongside already-dead skiller/chronos/skillsgateway — a dump artifact, not a live-service claim).
- `seeding-spec.md` "shared skillpaths myth" + "skillpath completed share" (skill-path SESSION *data* seeding —
  seeder-owned; M246 already re-pointed the seeder to `public.skill_path_sessions`, M250 owns further deltas).
- `snapshot-spec.md` cms `skillpath.go` + `publicSkillPaths` (skill-path *content*/CMS, not the service).
- `backend.md:125` `skillpaths/ Backend's view of skillpath data` (accurate app-internal code-map entry).
- `update_guide.md` M246 consolidation note (already correct, references M247 as the reconciliation owner).

**Out of scope (rext, doc-only milestone — D0):** `gen_injected_override.py` dormant skillpath key,
`test_injection.py`/`exposure_claim_guard.py` fixtures, `up-injected.sh` audit prose, and the dev-stack
`migrate-dev.sh`'s skillpath schema-creation — all rext files → M251/rext-hygiene (recorded in decisions.md D0).

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
