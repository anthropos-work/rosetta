# Clause-5 ground truth — derived fresh 2026-08-01 against platform origin `2adcf71`

You are auditing the Rosetta **documentation corpus** against the **actual platform**. The platform clone
is at `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform` (checked out at origin HEAD
`2adcf71`). You may read it freely. **Read-only: do not edit any file anywhere.**

## The grading rule (this is the whole job)

> **BLOCKER = the claim is FALSE at platform origin HEAD *and* acting on it would misdirect real work.**

- A claim that is **true at HEAD** is NOT a blocker, even if it *feels* stale or old.
- **Explicitly-fenced historical or production-only content is NOT a blocker.** Docs legitimately describe
  prod (where e.g. `roadrunner` is still deployed) and legitimately narrate history. Only unfenced
  present-tense claims about how things are *now* can be blockers.
- A doc for a merged service that opens with a standing ⚠ banner saying so is **correct**, not a blocker.
- Grade each finding: **BLOCKER** / **minor** (true-but-confusing, wrong line number, dead link) /
  **not-a-finding**.

## THE AUTHORITATIVE SOURCE for who is merged / live / gone

`corpus/architecture/platform-migration-status.md` — its services table is machine-fenced against the
platform's own `repos.yml` by `stack-core/platform_alignment_guard.py`. **Treat that table as ground
truth.** If a doc you are reading disagrees with it about a service's state, the doc is wrong.

## Key facts, derived this session (verify anything you rely on)

1. **`app` is the backend monolith.** It serves cms, jobsimulation, skiller, skillpath **in-process**
   (`app/internal/{cms,jobsimulation,skiller,skillpath}/`, wired at `app/main.go:604`). Its compose
   service is named **`backend`** (`docker-compose.yml:28`). It is the **only** repo with migrations
   (`repos.yml:10-13`, `migrations: true`, `schema: public`).
2. **Every application table lives in the `public` schema.** `sentinel` is the one exception — it keeps
   its own `sentinel` schema via `search_path=sentinel` (`docker-compose.yml:18`) *despite*
   `migrations: false`.
3. **`app/internal/roadrunner/` DOES NOT EXIST.** The Judge0 runner was absorbed as
   `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:118`.
4. **⚠ THE BIG ONE — the GraphQL/Cosmo Router was DROPPED FROM LOCAL DEV at `2adcf71` (2026-07-31),
   mid-milestone.** `graphql-wundergraph` was deleted from **both** `repos.yml` and `docker-compose.yml`
   (by `b56d731` + `360efd4`). **There is no `graphql` service in a local stack any more.** Local dev now
   points **directly at `backend`** — `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=…:8082/graphql/query`
   (`docker-compose.yml:352`). The repo is **ARCHIVED on GitHub** (2026-07-30), though a router is still
   declared in *prod* terraform. Any doc that describes a local stack routing through Cosmo/the router, or
   lists `graphql` as a local container, or calls `graphql` the default profile's gateway, is describing a
   world that no longer exists locally. **This is the newest drift in the corpus and the least likely to
   have been swept — check for it everywhere.**
5. **The supergraph is ONE subgraph: `backend`.** `supergraph-config-prod.yaml` lists `backend` alone,
   `schemas/` holds `backend.graphqls` alone, `subgraphs.conf` = `BACKEND=v1.360.0`. Any claim of 2+
   subgraphs, or of a `cms`/`jobsimulation`/`skiller`/`skillpath` subgraph, is false.
6. **Compose services actually defined in `docker-compose.yml`:** `sentinel` (:5), `backend` (:28),
   `jobsimulation` (:83), `cms` (:144), `storage` (:189), `customerio-sync` (:220, own profile),
   `messenger` (:240, own profile), `roadrunner` (:281), `studio-desk` (:311, own profile),
   `next-web-app` (:344, `frontend` profile), `gotenberg` (:371). **`postgresql` and `redis` are NOT in
   `docker-compose.yml`** — they live in the *included* `common.yml` (`docker-compose.yml:1-2`,
   `include: - common.yml`). `directus` was removed from compose at `a2a3ee6` (2026-02-27).
7. **The husk containers still START, and that is deliberate.** `cms`, `jobsimulation` and `roadrunner`
   are merged in production but their compose services still exist and still run locally
   (`running_but_unfederated`), and consumers still carry `CMS_RPC_ADDR=http://cms:8091`,
   `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`, `ROADRUNNER_RPC_ADDR=http://roadrunner:10401`.
   **"Merged in production" is NOT "removed from compose"** — a doc saying the container still starts is
   correct. Teardown is platform **M810**.
8. **`SKILLER_RPC_ADDR=http://backend:8083`** everywhere — the env var survives, re-pointed at `backend`.
   That is a deliberate husk re-point, not a bug.
9. **Directus:** external, at `content.anthropos.work`. `DIRECTUS_BASE_ADDR` is set on the `cms` service
   in base compose (`:164-165`); on a rext demo stack **`backend` is the Directus reader** (the
   `cms_reader_switch` — `app` swaps its content reader to the in-process cms server), which is why
   `backend` needs `DIRECTUS_BASE_ADDR` in the demo override. A local stack only gets its own Directus via
   rext's `--local-content` cutover, never from the platform repo.
10. **Shared libraries** (private Go modules, never cloned by `make init`): `colony` (framework; **also
    contains `authn` as `colony/authn`** — the standalone `authn` repo is legacy), `proto`, `ai`,
    `taxonomy` (**the `NodeID` type ONLY — NOT the 60K-skill dataset**, which lives in `app`'s `public`
    schema).
11. **Archived on GitHub:** `jobsimulation` (2026-07-31), `skillpath` (2026-07-31),
    `graphql-wundergraph` (2026-07-30), `skiller` (2026-07-01), `intelligence` (2026-04-02).
    **NOT archived:** `cms`, `roadrunner`, **`chronos`** (a doc claiming chronos is archived is wrong).
12. **`storage` and `messenger` are named as the NEXT fold** — `app` PR #1103 (v9.0 "support-in-app").
    They are live today; a doc saying they are already merged would be wrong.
13. The Python generation pipeline's repo is **`anthropos-studio-room`**, not `studio-room`. It is pulled
    into the `app` image by CI and spawned as a subprocess from `app/internal/cms/studio/`. **Not a
    service, not a container, not in `repos.yml`.**
14. `ant-academy` is deliberately absent from `repos.yml` — run natively, never containerised.

## What has already been swept (do not re-report these as new)

Iters 21–23 swept the corpus for merged-service status claims (93 enumerated edits). Harden pass 6 then
added `ServiceDocStatusFence`, which mechanically holds every per-service doc to the migration map's
merged-states. So **"doc X says service Y is live when the map says merged"** is already fenced — if you
find one, it is a real finding (the fence would be broken), but it is not the expected shape.

**The expected shape of what remains is what a grep-based sweep structurally cannot find:** claims that
are wrong without using any of the words a sweep would grep for. The studio-room doc is the archetype —
it read as a live pipeline for five paragraphs and said "merged" only in paragraph six, and three
consecutive sweeps missed it because it never used the grepped vocabulary.

## Method — non-negotiable

**READ EACH ASSIGNED FILE IN FULL, TOP TO BOTTOM.** Do not grep, do not skim, do not sample. The whole
reason this pass exists is that three term-scoped sweeps produced a convergent-looking 11 → 5 → 2 curve
that was actually just exhausting its own vocabulary — a full read then found 53.

**Positive control, per file:** run `wc -l <file>` and state in your report, for each file, the line count
and that you read to that line. A file you did not read to the end must be reported as unread, not
guessed.
