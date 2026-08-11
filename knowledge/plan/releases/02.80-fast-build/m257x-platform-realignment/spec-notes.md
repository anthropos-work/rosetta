---
milestone: M257x
status: archived
last_updated: 2026-08-11
---

# M257x — spec notes

## The recurring class this milestone must end

Three occurrences, one shape: **the platform consolidates a service into `app`, and rext keeps writing to the
schema that service owned.**

| release | service | how it surfaced |
|---|---|---|
| v2.1 | skiller → app | seeder broke |
| v2.7 | skillpath → app | seeder broke again; corpus asserted skillpath live Tier-1 in ~30 files |
| **now** | jobsimulation (+ cms, roadrunner "own no local schema") | **latent** — only because our clones are stale |

Each time the fix was re-derived from scratch. The deliverable that breaks the cycle is not the re-point — it
is the **fence** (clause 4) plus the **written procedure** (`corpus/ops/platform-alignment.md`).

## Measurement discipline inherited from this release

- **State the environment with every number** — the baseline mirror fence is now parameterised by host and
  FAILS a baseline-shaped claim that names no host (D120).
- **Prove a check can go RED before trusting it.** M256 found 43 checks that reported success without checking;
  M257 found the gate's own health check reading a dropped table behind a swallowed error.
- **A cold cycle is the only honest test.** B1 and B2 were both invisible to warm cycles for four days.

## Platform ground truth (verified 2026-07-31, Phase 0b KB-fidelity audit)

Read from **origin via the GitHub API** — no local clone was touched, and none of these came from the
stale `stack-dev/` clones. `platform` @ origin HEAD = **`1e8e75400c66dbd96abf4a1aca7e7a7cecaea497`**
(2026-07-30T08:26:40Z). Full report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md).

| fact | citation |
|---|---|
| `repos.yml` has **10** entries; `app` is the **only** `migrations: true`; `schema:` keys deleted from cms + jobsimulation; **no `ant-academy`** | `platform/repos.yml` @ `236771f103` |
| The fold commit is **`docs:`-only** — 3 files (`CLAUDE.md`, `README.md`, `repos.yml`); `docker-compose.yml` **untouched** | `platform` `236771f103` (2026-07-29T14:06:49Z) |
| cms / jobsimulation / roadrunner are **still in the default `graphql` profile** | `platform/docker-compose.yml:169/212`, `:108/165`, `:306/334` |
| `graphql` still `depends_on` jobsimulation + cms; `backend` still exports both RPC addrs | `platform/docker-compose.yml:20-25`, `:69`, `:75` |
| Only `/migrations: true/` repos migrate | `platform/Makefile:14` |
| Postgres image creates **no** schemas (pgvector only) | `platform/postgresql/Dockerfile` |
| jobsim data model re-created in `public` — **23 tables** | `app/terraform/migrations/20260722081626_jobsim_data_model.sql` |
| cms data model re-created in `public` — 6 tables | `app/terraform/migrations/20260724132049_cms_data_model.sql` |
| **`local_*` mirrors DROPPED** | `app/terraform/migrations/20260729133514.sql:62-63` |
| Canonical jobsim session table is **`job_simulation_sessions`**, not `sessions` | `app/terraform/migrations/20260722104506.sql:2`; `app/internal/data/ent/schema/job_simulation_session.go:7-9` |
| `app` embeds the studio pipeline | `app/Dockerfile.dev:24-26,38-41`; `app/.gitignore:78-79` |
| `cms` **still** embeds it too (addition, not move) | `cms/Makefile:11-17` |
| Supergraph is **ONE** subgraph (`backend`) | `graphql-wundergraph/supergraph-config-compose.yaml`, `-prod.yaml`, `schemas/`, `Dockerfile.dev:18-23`, `subgraphs.conf` |
| jobsimulation + cms subgraph wiring removed together | `graphql-wundergraph` `915da06c58` (2026-07-29T09:24:38Z) |

## Topic → doc → code triples (for fast re-audit at the exit gate)

| topic | corpus doc(s) | platform code | state |
|---|---|---|---|
| jobsimulation service | `services/jobsimulation.md`, `architecture/architecture_overview.md:143,151`, `architecture/service_taxonomy.md:58,112` | `repos.yml`, `docker-compose.yml:108-167`, `app/terraform/migrations/20260722081626_*.sql` | STALE |
| cms service / schema | `services/cms.md:26,113` | `repos.yml`, `docker-compose.yml:169-212`, `app/terraform/migrations/20260724132049_*.sql` | STALE |
| service→schema map | `architecture/architecture_overview.md:270-276`, `ops/platform_repo.md:92`, `ops/setup_guide.md:291-292,487,684` | `repos.yml`, `Makefile:14`, `postgresql/Dockerfile` | STALE |
| subgraph count (3 → **1**) | `architecture/external_services.md:307,312,332,379-384`, `architecture/service_taxonomy.md:288`, `services/graphql-wundergraph.md:9,16,94,103-105`, `services/backend.md:49,67-68`, `services/cms.md:104`, `services/README.md:30`, `architecture/architecture_overview.md:105,202,220,226`, `ops/update_guide.md:19` | `graphql-wundergraph/supergraph-config-*.yaml`, `schemas/`, `Dockerfile.dev:18-23` | STALE (16 places) |
| studio-room placement | `services/studio-room.md:15,25,27`, `services/cms.md:9,11`, `architecture/architecture_overview.md:13,23,54`, `architecture/ai_architecture.md:52,54`, `architecture/service_taxonomy.md:57,149-179`, `architecture/dependency_map.md:12,23,85` | `cms/Makefile:11-17` **+ `app/Dockerfile.dev:24-41`** | INCOMPLETE (app half absent) |
| `local_*` mirrors | `services/hiring.md:105,132-143,182`, `ops/seeding-spec.md:386-392,416,528-536` | `app/terraform/migrations/20260729133514.sql:62-63` | STALE (KB-1) |
| canonical session table | *(none — blind)* | `app/…/ent/schema/job_simulation_session.go:7-9` | CODE-ONLY (KB-2) |
| `ant-academy` in `repos.yml` | `ops/platform_repo.md:93`, `ops/setup_guide.md:297` (vs correct `CLAUDE.md:196`, `services/roadrunner.md:13`) | `repos.yml` (absent) | STALE |
| cold-init procedure | `ops/setup_guide.md:684-686` | only `app` migrates | STALE |

## Search-set integrity (do this before any absence claim)

The corpus is **clean** and the NUL-byte false-negative class does **not** apply: `find corpus -name '*.md'`
= `rg --files corpus/` = **88**; **0** NUL bytes; `file(1)` reports all 88 as text. A Phase 0b sub-agent
nonetheless claimed `services/next-web-app.md` was binary-skipped — **refuted** by 4 methods (0 NUL bytes,
`file(1)` UTF-8, Python decodes clean, `rg`/`grep` counts identical with and without `-a`). macOS `iconv`
is the lone dissenter and is not authoritative. **Always pass `-a` anyway, and prove the search set.**
