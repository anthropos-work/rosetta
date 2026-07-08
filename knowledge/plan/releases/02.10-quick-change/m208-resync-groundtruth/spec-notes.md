# M208 — Spec notes

_Technical details accumulate here as the milestone is built._

## Pre-flight audits — Re-sync & merged-schema ground-truth
- Phase 0b KB-fidelity: **YELLOW** — report `kb-fidelity-audit.md` (sha at audit `e319d2f`). No blind
  area; pre-merge corpus staleness tracked as KB-1/2/3 (all Fate-2 → M210). Proceed.
- Topic → doc → code triples:
  - Merged shape → `corpus/services/{backend,skiller}.md` → `stack-dev/app/internal/data/ent/schema/{skill,jobrole,category,specialization,*_embeddings,skiller_mixins}.go` + `internal/rpc/skillerrpc/` (merge commit `1fc00c78`)
  - RPC re-point → `backend.md` + `dependency_map.md` → `platform@origin/main:docker-compose.yml` (`SKILLER_RPC_ADDR=http://backend:8083`)
  - 4-subgraph federation → `graphql-wundergraph.md` + `CLAUDE.md` → `platform@origin/main` router config
  - Re-sync ops → `update_guide.md`/`setup_guide.md`/`run_guide.md` → `stack-dev/platform/Makefile`

## Re-sync (make pull / refs)

Both stacks pulled to `origin/main` on 2026-07-08. `make pull` iterates the SIBLING repos from
`repos.yml` and does NOT touch `platform` itself, so platform was pulled directly first (which also
swaps in the merged `repos.yml`/compose before the sibling sweep runs).

**Before → After (short refs):**

| repo | stack-dev before → after | stack-demo before → after |
|---|---|---|
| **platform** | `5e1ae6b` → **`0808b92`** (2 ahead: rm skiller from compose+repos.yml+Make) | `5e1ae6b` → **`0808b92`** |
| **app** | `a848cccb` (v1.318) → **`c3c45e01`** (v1.334.1) — **86 commits** | `158a8394` → **`c3c45e01`** (v1.334.1) |
| cms | `57297a6` → `770ec3a` | `57297a6` → `770ec3a` |
| jobsimulation | `9f40604a` → `5d3003f9` | `9f40604a` → `5d3003f9` |
| graphql-wundergraph | `38f5d0a` → `c284453` (**`schemas/skiller.graphqls` deleted**; backend.graphqls +259) | `7ffe4f8` → `c284453` |
| next-web-app | `d689ecdea` → `23bdbb5db` | `928cc8e32` → `23bdbb5db` |
| studio-desk | `7a9ad78` → `f6320f8` | `7a9ad78` → `f6320f8` |
| sentinel / storage / messenger / roadrunner / skillpath | already current (unchanged) | already current (unchanged) |
| skiller (**vestigial**) | `b7a8950` (not in repos.yml — removed in §2) | `b7a8950` (removed in §2) |

Post-pull `platform` (both stacks): `repos.yml` has **0** skiller entries, `docker-compose.yml` has
**0** skiller services, and all four `SKILLER_RPC_ADDR` values are `http://backend:8083`.

Out of scope / untouched: `stack-dev/ant-academy` is **13 behind** but **not in `repos.yml`** (a
Clerk-free UI-tier native app, not part of the skiller merge) so `make pull` skips it — a UI-tier
concern for M211, not this merged-platform de-risk. `rosetta-extensions` (stack-demo) is a pinned-tag
clone, also not in `repos.yml` — M209's concern.

## Vestigial clone removal

`stack-dev/skiller` + `stack-demo/skiller` (both `b7a8950`, 8 MB each) — verified clean (no
uncommitted work) and referenced by **0** entries in either `repos.yml` — `rm -rf`'d. Post-removal
scan of `docker-compose.yml` confirms **no residual skiller container wiring** (no `context: ../skiller`,
no `http://skiller:8086`; all consumer RPC addrs point at `backend:8083`). A lingering clone was a
false signal only; nothing built or wired against it.

## Re-migrate against public

Two runs on stack-dev (override parked for both, restored at close). The **clean-slate reset-db run
(orchestrator-driven) is the authoritative de-risk**; the existing-volume run is complementary.

**Build/compose (`make up`, cold, profile `graphql`) — rc=0 (~15 min).** The 86-commit merged `app`
monolith rebuilt clean (private-module pull OK via `GH_PAT`). Containers up: `graphql` (Cosmo/WunderGraph
router 0.275.0, host :5050), `jobsimulation`, `cms`, `skillpath`, `sentinel`, `storage`, `roadrunner`,
`gotenberg`, `postgresql`+`redis` (healthy) — plus `backend` (see the secrets finding). **No `skiller`
container** (running or stopped). `SKILLER_RPC_ADDR=http://backend:8083` on the merged compose (RPC
re-pointed to backend); backend `DB_CONNECTION` uses the default `public` search_path (no
`search_path=skiller`); the merged app subscribes to the `skiller` Redis stream **in-process**.

**Clean-slate schema proof (authoritative — `make reset-db` + `make migrate`, 149 app migrations from an
empty DB):** the merged migrations create the **full public taxonomy from scratch** — `public.skills`
(with an `organization_id` column — empirically confirms the M209 firewall `organization_id IS NULL`
predicate survives the merge), `public.job_roles`, `public.job_role_skills`, `public.skill_embeddings`,
`public.categories`, `public.specializations` — and **no `skiller` schema exists on a clean DB**
(`skiller_legacy=false`). `public.skills` = 0 rows after migrate (EXPECTED — taxonomy DATA comes from
snapshot replay, M209/M211, not migrations). This is the load-bearing merged-schema proof.

**Existing-volume run (complementary):** on the pre-existing DB volume (which already had `extensions` +
the taxonomy tables), `make migrate` applied clean — 1 pending app migration (`20260702141235`,
ai_readiness `track`) ok; cms/jobsim/skillpath none pending. With a dev `INVITATION_HMAC_SECRET` present,
`backend` came Up ("Web server started at :8082") and a **runtime smoke through the router confirmed the
4-subgraph federation SERVES end-to-end**: `{__typename}`→`Query`, `__type(name:"Skill")`→OBJECT, and
query fields `jobRoleMatch`/`similarJobRoles`/`jobRoleCount`/`mostPopularSkills`/`mostPopularCategories`/
`topJobRoleSkills` are served by the **backend** subgraph (router "ready to serve", **0 skiller mentions**).
NB: this run's "clean migrate" masked the `extensions` prerequisite below — the clean-slate run is why the
prerequisite is now known.

### Finding 1 — clean-bring-up `extensions` bootstrap + PG-readiness (M25-D9 class → M211 Fate-3; M209 Risk-2 cross-ref)
A clean `make reset-db` does **not** bootstrap the `extensions` schema (pgvector + `pg_trgm`) before
`make migrate`. The vector-column + trigram-index migrations then fail on an empty DB:
- app `20260518125439` — `CREATE TABLE ask_query_examples (… "embedding" extensions.vector(1536) …)` →
  `pq: schema "extensions" does not exist` (migrate2.log:2107).
- cms `20250116133510` — `CREATE TABLE similarities (… "small_embedding3" extensions.vector(1536) …)` →
  `pq: schema "extensions" does not exist` (migrate2.log:2121).
- app `20260623090000` — `skilltranslation_name_gin … USING GIN (name extensions.gin_trgm_ops)` needs the
  `extensions.gin_trgm_ops` opclass.
Manually `CREATE SCHEMA extensions; CREATE EXTENSION vector SCHEMA extensions; CREATE EXTENSION pg_trgm
SCHEMA extensions;` before migrate lets the taxonomy tables create (the `gin_trgm_ops` opclass resolution
still needs exact search_path handling — a bring-up-tooling detail). Secondary: a **PG-readiness race** on
reset-db (migrate runs before Postgres stabilizes → `connection reset by peer`; resetdb.log:14/17/20 — a
re-run succeeds). **This is the M25-D9 opportunistic item the overview flagged — it did NOT trivially fall
out as Fate-1; routed Fate-3 to M211** (the iterative bring-up-acceptance milestone owns extensions-bootstrap
+ PG-readiness on cold bring-up). The `extensions.`-qualified type detail **also** belongs in M209's Risk-2
note (the capture column list uses `extensions.vector` + `extensions.gin_trgm_ops`).

### Finding 2 — `INVITATION_HMAC_SECRET` per-stack `.env` gap (→ M211 / `/stack-secrets`, NOT provisioned here)
The cold containerized `backend` `Exited(0)` at startup on a missing `INVITATION_HMAC_SECRET` (fatal:
`app/main.go:266` → `internal/invitations/token.go:23`; POSTHOG key also empty). This is a **per-stack
`.env` completeness gap, not merge-caused** — the key IS documented (`corpus/ops/secrets-spec.md:111` "the
app exits early when it is unset"; one of the platform `.env` 29 keys, line 98; in
`secretdna.DemoGeneratedKeys` which demo stacks auto-generate; app `docs/invitation-reminders.md` prescribes
`openssl rand -hex 32`). stack-dev's hand-assembled `.env` lacked it (its containerized `backend` is
normally not run — native-worktree dev). During the existing-volume run I did a Fate-1 dev-value add to
confirm the federation serves end-to-end (above), **then reverted it** per the orchestrator directive
("do NOT provision secrets here"); the `.env` is left in its original gap state. **Routed to M211 /
a `/stack-secrets` follow-up** — the sanctioned provisioner, not an ad-hoc edit. Because of this gap, the
clean-slate authoritative run could not verify the backend-serving subgraph LIVE (backend down); the
existing-volume run did.

## Merge fact-sheet

**Authoritative merged shape (verified 2026-07-08):**

- **Moved tables — now in `public`, names unchanged** (`skiller.X → public.X`): `skills`, `job_roles`,
  `categories`, `specializations`, `skill_embeddings`, `job_role_embeddings`, `skill_translations`,
  `job_role_translations`, `job_role_skills`, `job_role_categories` (+ the user-side `user_skills`,
  `user_skill_evidences`, `membership_skills`, …). Confirmed by the app ent schema
  (`internal/data/ent/schema/{skill,jobrole,category,specialization,*_embeddings,skiller_mixins}.go`),
  the port migrations (`terraform/migrations/…20260623090000_skiller_taxonomy_name_trgm_indexes.sql`,
  merge commit `1fc00c78 Deprecate skiller schema`), and `information_schema` on prod (both `public.*`
  and a **legacy** `skiller.*` mirror exist; `public.*` is authoritative).

- **Public predicate `organization_id IS NULL`** — the public taxonomy. Measured on prod (read-only
  postgres MCP, 2026-07-08):

  | metric (prod `public` schema) | count |
  |---|---|
  | `skills WHERE organization_id IS NULL` (**the public-skill count**) | **42,790** |
  | `skills` total (incl. 794 org-private) | 43,584 |
  | `job_roles WHERE organization_id IS NULL` | 22,490 |
  | `categories` (org NULL) | 23 |
  | `specializations` (org NULL) | 1,447 |
  | `skill_embeddings` | 43,584 |
  | legacy `skiller.skills` (deprecated mirror) | 43,584 |

  → the roadmap's "~42,763 public skills" assertion resolves to a **measured 42,790** (the live figure;
  taxonomy grows over time). The predicate + the merged `public.skills` shape are confirmed on prod.

- **RPC re-point**: `SKILLER_RPC_ADDR=http://backend:8083` (local, all 4 occurrences in the merged
  `docker-compose.yml`); `http://backend:8081` in prod terraform. The `SkillerService` surface is served
  by app (`internal/rpc/skillerrpc/`).

- **4 subgraphs** (skiller subgraph removed; `schemas/skiller.graphqls` deleted at `graphql-wundergraph@c284453`):
  **backend** (v1.332.0), **jobsimulation** (v0.252.1), **cms** (v0.254.2), **skillpath** (v0.32.3).

- **Local (cold) note**: a clean `make reset-db`+`make migrate` creates the `public` taxonomy tables
  **empty** (0 rows) — the 42,790 figure is the prod public taxonomy (populated per-stack by snapshot
  set-dress, M209/M211 scope). M208 confirms the local merged **shape** (tables created, `organization_id`
  column present, no `skiller` schema on a clean DB), not local population. **Prerequisite:** the clean
  bring-up needs the `extensions` schema (pgvector + `pg_trgm`) bootstrapped before migrate — see Finding 1.

## M25-D9 (opportunistic — surfaced on clean-slate → Fate-3 to M211)

**It surfaced** — but on the **clean-slate** `make reset-db` path, not the existing-volume migrate (which
masked it). It is the `extensions`-bootstrap + PG-readiness bring-up-ordering issue documented as **Finding
1** above: on a truly empty DB the merged vector/trigram migrations fail (`schema "extensions" does not
exist`; `extensions.gin_trgm_ops` opclass). It did **not** trivially fall out as a Fate-1 fix here (it's a
bring-up-tooling requirement, not a one-line migrate tweak), so per the charter's "resolve only if it falls
out naturally" it is **routed Fate-3 to M211** (bring-up acceptance owns extensions-bootstrap + PG-readiness)
with an M209 Risk-2 cross-ref for the `extensions.`-qualified capture column list. See Finding 1.
