# iter-34 confirming pass — audit group A

Auditor: group A. Platform clone read-only at `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform`,
verified at origin HEAD `2adcf714bd877a205e8948f59a23db49b884c054` ("Merge pull request #23 … chore/drop-wundergraph",
2026-07-31). Peer clones (`app`, `cms`, `jobsimulation`, `messenger`, `sentinel`, `storage`, `studio-desk`,
`next-web-app`, `graphql-wundergraph`, `ant-academy`) read from the same `stack-demo/` set. The
`rosetta-extensions` authoring copy was read at `b2b46cb` (2026-08-01).

## Positive control — every assigned file read top to bottom, in full

| file | `wc -l` | last line actually read | status |
|---|---:|---:|---|
| `corpus/architecture/external_services.md` | 704 | 704 (`- [Architecture Overview](./architecture_overview.md) - System architecture`) | READ IN FULL |
| `corpus/services/studio-desk.md` | 419 | 419 (`- [demopatch-spec §8](../ops/demo/demopatch-spec.md) - the studio-desk source patches (additive-UI injection)`) | READ IN FULL |
| `corpus/services/ant-academy.md` | 397 | 397 (`- [External Services](../architecture/external_services.md) — Clerk integration details`) | READ IN FULL |
| `corpus/architecture/dependency_map.md` | 103 | 103 (`* \`GOTENBERG_URL=http://gotenberg:3200\` is injected via the backend's compose env.`) | READ IN FULL |
| `corpus/architecture/README.md` | 38 | 38 (`*   Document frontend changes in **frontend_architecture.md** as the monorepo evolves.`) | READ IN FULL |
| `corpus/services/skiller.md` | 55 | 55 (`* [Dependency Map](../architecture/dependency_map.md)`) | READ IN FULL |

No file UNREAD. Total 1,716 lines.

---

## Findings

### F-A1 — BLOCKER — `corpus/architecture/external_services.md:199-208`

**Verbatim (the load-bearing parts):**

> **⚠️ The `--local-content` re-point targets `cms`, NOT `backend` — measured, not inferred.** … and
> `rosetta-extensions/stack-injection/gen_injected_override.py:579-580` re-points **only the services in
> `DIRECTUS_DATA_CONSUMERS`** — which is `cms` — with `test_only_cms_is_repointed_not_other_services` asserting
> that `backend` must **not** carry the re-point. … Tracked as
> `FIX-M257x-iter23-backend-directus-not-repointed`

**What is actually true at HEAD.** That fix **already shipped**. This is the "work described as routed forward
that already shipped" class the addendum names, and it is wrong in three independent ways:

1. `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` — `.agentspace/rosetta-extensions/stack-injection/gen_injected_override.py:53`.
   `backend` **is** in the list. The doc's "which is `cms`" is false.
2. The re-point is emitted for every member of that list at
   `gen_injected_override.py:598-599` (`if with_directus and name in DIRECTUS_DATA_CONSUMERS:` →
   `env.append(f"DIRECTUS_BASE_ADDR={DIRECTUS_INNETWORK_ADDR}")`) — so `backend` **does** carry the re-point.
   The cited anchor `:579-580` is also wrong: line 579 is the M214 CORS emission (`fe_ports = (3000, 3001, 9000)`).
3. `test_only_cms_is_repointed_not_other_services` **no longer exists**. It was deleted and replaced by
   `test_backend_the_actual_reader_is_repointed` (`stack-injection/tests/test_injection.py:1005-1017`), whose own
   comment reads: *"This test replaces `test_only_cms_is_repointed_not_other_services`, which asserted the
   OPPOSITE … `backend` is now the direct, in-process Directus reader (`app/cms_reader_switch.go`), so it is the
   one service that MUST be re-pointed."* The negative test `test_non_readers_are_not_repointed`
   (`:1027-1035`) now explicitly excludes `backend` and `cms` from its non-reader list.

**Proof:** `rosetta-extensions` commit `f9ac72f` — *"fix(M257x/24): re-point the per-stack Directus at the
service that actually reads it"* (`git log -S'DIRECTUS_DATA_CONSUMERS = ("cms", "backend")'` returns exactly
that one commit).

**Why it misdirects real work.** The paragraph is written in the present tense as a *measured standing defect*
with an open FIX id. An engineer picking it up would re-implement a shipped fix, or — worse — would debug a demo
on the premise that `backend` reads **prod** Directus anonymously (the old 403-on-`directus_versions` symptom)
when at HEAD it reads the in-network per-stack Directus. The dated live-`demo-1` observation at `:204-206` is a
snapshot from *before* the fix and is presented as current state.

**Suggested correction:** rewrite the callout in the past tense as *resolved at M257x iter-24*: state that
`DIRECTUS_DATA_CONSUMERS` is now `("cms", "backend")` (`gen_injected_override.py:53`), that the emission is at
`:598-599`, that the contract test is `test_backend_the_actual_reader_is_repointed`, and drop/retire the
`FIX-M257x-iter23-backend-directus-not-repointed` tracker.

---

### F-A2 — minor — `corpus/architecture/external_services.md:401-402`

**Verbatim:**
```
COPY cms/internal/graph/schemas/ /tmp/schemas/cms/
COPY jobsimulation/internal/graph/schemas/ /tmp/schemas/jobsimulation/
```
presented as a quote of *"The gateway's `Dockerfile.dev`"*, under a HISTORICAL fence (`:380-382`) that justifies
itself with *"Kept because the archived repo still contains these configs and a reader will meet them there."*

**True at HEAD:** the archived repo's `graphql-wundergraph/Dockerfile.dev` contains **only**
`COPY app/internal/web/backend/graphql/graph/schemas/ /tmp/schemas/backend/`, followed by the comment *"cms +
skillpath folded into the backend subgraph (cms-in-app / skillpath-in-app) — the backend SDL now owns the cms
content types + SkillPathSession, so there are no standalone subgraph SDLs."* The two quoted COPY lines are not
in the file the fence points a reader at.

**Grade:** minor — the section is fenced as historical, but the fence's own premise is false for this snippet.
**Correction:** replace the snippet with the file's actual contents, or mark it "as of the pre-fold revision".

---

### F-A3 — minor — `corpus/architecture/external_services.md:421-429` (`### Configuration`)

**Verbatim:** `ENVIRONMENT=compose  # or production` … *"**Build Context**: the platform monorepo (`context: ..`)
— not the upstream repo. This was changed from the old "git+url" build because the composition needs sibling
repos."*

**True at HEAD:** there is no `graphql` service in `docker-compose.yml`, so there is no local build context of
any kind (fact 4; `repos.yml` and `docker-compose.yml` both scrubbed by `b56d731`+`360efd4`). The HISTORICAL
fence at `:380-382` explicitly ends at *"the end of *Subgraph routing URLs*"* — i.e. this section is deliberately
**outside** it and reads as current local config.

**Grade:** minor (not blocker) — the file's opening banner, the Overview row at `:345` ("prod-only"), and the
fence three sections above all warn the reader first, so real misdirection risk is low.
**Correction:** extend the HISTORICAL fence to cover `### Configuration`, or re-scope it to "in production".

---

### F-A4 — minor — `corpus/architecture/external_services.md:456-457`

**Verbatim:** `// Queries in app/graphql/*.graphql` / `// Types in app/__generated__/`

**True at HEAD:** `studio-desk/app/graphql/` and `studio-desk/app/__generated__/` do not exist. The real paths
are `app/services/graphql/` and `app/services/__generated__/` (verified `ls studio-desk/app` and
`ls studio-desk/app/services`). `corpus/services/studio-desk.md:139` states the correct paths, so the corpus
contradicts itself.
**Correction:** `app/services/graphql/*` and `app/services/__generated__/`.

---

### F-A5 — minor — `corpus/architecture/external_services.md:694`

**Verbatim:** `> Consistent with :447 above, where the same correction is already recorded.`

**True at HEAD:** line 447 of this file is `const user = await client.query({` — a TypeScript sample, not a
correction. The correction it means (*"there is no `graphql` service locally any more; restart `backend`"*) is at
**`:481`**. Stale self-anchor.
**Correction:** `:481`.

---

### F-A6 — minor — `corpus/architecture/external_services.md:137`

**Verbatim:** line 137 begins with a space then `A freshly-` — no `>` prefix, mid-blockquote.

**Effect:** the sentence *"A freshly-built local stack reads its public content live from prod"* — a load-bearing
statement of the default posture — falls out of the ⚠️ callout when rendered, splitting the block in two.
**Correction:** prefix with `>`.

---

### F-A7 — minor — `corpus/architecture/external_services.md:163-165`

**Verbatim:** *"it starts **nine** containers — `postgresql` + `redis` (from the included `common.yml`,
profile-less so they always start) and seven application services, `sentinel` · `backend` · `jobsimulation` ·
`cms` · `storage` · `roadrunner` · `gotenberg`."*

**True at HEAD:** the count **nine is correct** (verified: `backend` `:81`, `jobsimulation` `:140`, `cms` `:187`,
`storage` `:218`, `roadrunner` `:309`, `gotenberg` `:384` all carry `profiles: [graphql, …]`; `postgresql` +
`redis` are profile-less in `common.yml`). But `sentinel` (`docker-compose.yml:5-27`) has **no `profiles:` key at
all** — it is profile-less exactly like postgres/redis, not a `graphql`-profile member. The sentence attributes
it to the profile.
**Correction:** move `sentinel` into the profile-less clause: "`postgresql` + `redis` + `sentinel` (profile-less,
always started) and six `graphql`-profile services".

---

### F-A8 — minor — `corpus/services/studio-desk.md:64`

**Verbatim:** `│   ├── designer-sim/   # Simulation designer interface`

**True at HEAD:** there is no `studio-desk/app/designer-sim/`. The designer surfaces are
`app/simulation-builder/`, `app/sim-advanced-builder/` and `app/sim-guided-builder/` — the same three the doc's
own line 34 names correctly. Verified `ls studio-desk/app`. Every other entry in the tree (`core/`,
`builder-skill-path/`, `generation/`, `listing/`, `academy/`, `home/`, `skills/`, `shared/`, `services/`,
`assets/`) exists.
**Correction:** replace with the three builder dirs.

---

### F-A9 — minor — `corpus/services/studio-desk.md:38`

**Verbatim:** `   - GraphQL integration with CMS service`

**True at HEAD:** there is no standalone CMS service on the path — studio-desk talks to `backend:8082/graphql/query`
(`docker-compose.yml:318`/`:334`), and the Directus reader is the cms **domain** inside `backend`
(`app/cms_reader_switch.go`; `app/main.go:971-973` hard-fails without `DIRECTUS_BASE_ADDR`). The mermaid eleven
lines later (`:49`) states this correctly, so the bullet is a stale restatement rather than the doc's position.
**Correction:** "GraphQL integration with the cms domain inside `backend`".

---

### F-A10 — minor — `corpus/services/ant-academy.md:285-289` (and the `code/.env` references at `:301`, `:306`)

**Verbatim:** *"The **app's** env file is `code/.env`, not the repo root:"* followed by `cp .env.example .env`.

**True at HEAD:** the clone ships `code/.env.example` and `code/.env.local` — there is no `code/.env`
(`ls -a ant-academy/code`). Every other authority writes `.env.local`: the corpus `CLAUDE.md` ("the React app
reads only from `code/.env.local`"), `demo-stack/ant-academy.sh`, and `stacksecrets provision` (this doc's own
`:190` and `:196-199` say so). Because Next.js loads `.env.local` at **higher** precedence than `.env`, a reader
who follows this section on any stack that already has `.env.local` writes a file that is silently overridden —
which is precisely the M245 "dewire" failure mode described at `:196-199`.
**Grade:** minor (both files are read; only precedence differs).
**Correction:** `cp .env.example .env.local`, and change "the app's env file is `code/.env`" to `code/.env.local`.

---

### F-A11 — minor — `corpus/services/ant-academy.md:83-87`

**Verbatim:**
```
    subgraph Core["Core Backend (Tier 1, Docker)"]
        App[app]
        CMS[cms]
        Jobsim[jobsimulation]
    end
```
**True at HEAD:** `cms` and `jobsimulation` are `merged-into-app` / `running_but_unfederated` husks
(`platform-migration-status.md` rows `cms`, `jobsimulation`) — they own no schema, serve no subgraph, and nothing
in this diagram's flows reaches them. Drawn as unlabelled peers of `app` they read as live core services.
**Grade:** minor — the doc's prose gets the merge status right everywhere else and `Desk --> Core` is the only
edge into the box.
**Correction:** label them "(husk, merged into app)" or drop them from the diagram.

---

### F-A12 — minor — `corpus/architecture/README.md:19`

**Verbatim:** *"**[external_services.md]**: The third-party integrations — Clerk (auth), Directus (…), the
WunderGraph Cosmo GraphQL gateway, the AI providers, LiveKit (voice), and AWS Chime (recording) — how each is
configured and consumed."*

**True at HEAD:** still true *for production* (`graphql-wundergraph/terraform/main.tf:20` `= 1`), so this is not
false — but this index line is the only mention of the router in the file and presents it undifferentiated
alongside Clerk/Directus/AI as a current integration, with no hint that it was deleted from local dev at
`2adcf71` and that the repo is archived. Every other entry in this index carries a distinguishing qualifier.
**Grade:** minor — a reader is not told anything false, only nothing about the newest drift in the corpus.
**Correction:** append "(**production-only** since platform `2adcf71`; archived on GitHub — local dev calls
`backend` directly)".

---

## Files with NO findings

**`corpus/architecture/dependency_map.md` — clean.** Every citation verified exact against the clone:
`docker-compose.yml:70-80` (backend `depends_on` = redis/postgresql/sentinel/cms/storage), `:144` (cms), `:83`
(jobsimulation), `:281` (roadrunner), `:337-341` (studio-desk `depends_on` backend+cms), messenger `:255`
(`BACKEND_USERS_RPC_ADDR=http://backend:8083`), `:256` (`CMS_RPC_ADDR=http://cms:8091`), `:258`
(`JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`), `:265` (`SKILLER_RPC_ADDR=http://backend:8083`);
`repos.yml:14-19`; `AI_USAGE_STREAM=AI`, `CMS_STREAM`, `SKILLER_STREAM`, `SKILLPATH_STREAM`,
`JOBSIMULATION_STREAM`, `GOTENBERG_URL=http://gotenberg:3200` all present in backend's env block; the M809
re-point comment at `app/main.go:1196-1202`; `DIRECTUS_WEBHOOK_SECRET` + `POST /api/webhook/directus`
(`app/main.go:1079-1080`, `app/internal/web/backend/backend.go:324`); `app/internal/converter/gotenberg.go`.
The shared-library table at `:44-50` is exactly right, including the direct/indirect split — `go.mod` confirms
`taxonomy v1.2.0` **direct** in `app` and `messenger`, `// indirect` in `sentinel` and `storage`, and `ai`
**only** in `app`.

**`corpus/services/skiller.md` — clean.** The ⚠ banner is the correct shape per the grading rule, and every
derived fact underneath verified: `app/internal/rpc/skillerrpc/`; the Ent schemas `skill.go`, `jobrole.go`,
`skill_embeddings.go`, `job_role_embeddings.go`, `category.go`, `specialization.go`;
`app/internal/web/backend/graphql/graph/schemas/skiller_taxonomy.graphqls`, whose line 7 states verbatim that
`categoryTree` / `fullCategoryTree` "stay unported — no consumers"; `skill_embeddings` / `job_role_embeddings`
with `extensions.vector(1536)` and `USING IVFFLAT` indexes
(`terraform/migrations/20260615130000_skiller_taxonomy.sql:61,241,251`); `skill_translations` /
`job_role_translations`; and the **8** `ContentLanguage`s (`app/internal/content/language.go:12-19` — english,
italian, spanish, french, german, dutch, japanese, portuguese) — an exact count, correctly stated.

**Anchors and links checked across all six files** — all 14 relative links resolve, and all six deep anchors
(`#dependent-repos--how-they-integrate`, `#configuration-keys`, `#ai`, `#taxonomy`,
`#skiller-in-app-merge--fact-sheet-v21-quick-change`, `#the-content-model--db-authoritative-catalog-v051-m7`)
match real headings. `corpus/architecture/README.md` indexes all 10 sibling docs with no gaps.

---

## Counts

**1 BLOCKER, 11 minors.**

**How the group read:** **mixed, and the split is exactly the one the addendum predicted.** The single blocker is
in the *newest* text — the swept `external_services.md` — and is self-inflicted repair debris: a defect written
up as an open, present-tense, "measured not inferred" FIX that had already been fixed in the tooling
(`f9ac72f`, M257x iter-24) before this pass ran. Four more minors cluster in the same swept file, all in the
*surrounding prose* of correct anchors (a stale self-reference to `:447`, a fabricated Dockerfile quote, a
section left outside its own historical fence, a broken blockquote prefix). The never-edited files read
genuinely clean at the fact layer — `dependency_map.md` and `skiller.md` produced **zero** findings across ~40
verified citations and three exact counts, which reads as real cleanliness rather than under-detection; the four
minors in `studio-desk.md` / `ant-academy.md` / `architecture/README.md` are all stale-detail or presentation
issues (a directory that no longer exists, an env filename with the wrong precedence, husks drawn as live in a
diagram), none of them status rot.
