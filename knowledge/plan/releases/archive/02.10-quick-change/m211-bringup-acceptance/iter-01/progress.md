**Type:** tok (bootstrap) — per build-mstone-iters Phase 0 rule 1 (iter-01 of M211).

# iter-01 — Bootstrap tok progress

## Phase 0b — Pre-flight KB-fidelity gate (milestone-once)
Ran a proportionate confirmation rather than a full re-audit: **M210 (closed 2026-07-08, same day)
was itself the corpus + skills re-ground and closed with a GREEN `/developer-kit:audit-kb-fidelity`
verdict (KB-1/2/3 resolved).** The orchestrator context explicitly bans re-deriving M208/M209/M210
verified facts. Spot-confirmed the standing verdict still holds at iter-01: grep of the tooling-queried
corpus paths (`corpus/ops`, `corpus/architecture`, `corpus/services`) for residual `skiller.<table>`
refs → **0 hits** (M210's deliverable intact). Verdict: **GREEN (inherited from M210, same-day,
spot-confirmed).** Recorded in `spec-notes.md § Pre-flight audits — iter-01`.

## Recon (grounds the strategy)
- Warm merged `stack-dev` is UP: 11 `anthropos-*` containers (backend, cms, gotenberg, graphql,
  jobsimulation, postgresql[healthy], redis[healthy], roadrunner, sentinel, skillpath, storage),
  4h uptime. **No skiller container** → sub-condition (a) MET; the merge composes 4 subgraphs.
- Taxonomy cache `.agentspace/snapshots/taxonomy/c75ce94d6a8021cad2915ddb4fb3dd4d/`: manifest.json +
  10 `skiller.*.copy` payloads (skills 42,790 rows / job_roles 22,470 / specializations 1,447 /
  categories 23 / skill_embeddings 42,790 / job_role_embeddings 18,919 / skill_translations 85,545 /
  job_role_translations 43,550 / job_role_skills 72,705 / job_role_categories 22). Manifest keys every
  table `"schema":"skiller"`, payload `skiller.X.copy`, plus internal `filter` SQL (`"skiller"."skills"`,
  `"skiller"."job_roles"`) + `public_via` refs → all need re-keying to `public` for the migration.
- rext authoring copy HEAD `2f06e78` = tag `quick-change-m209`; production code has **0** residual
  `skiller.<table>` refs (M209 delivered). Consumption pin `.agentspace/rext.tag` = `v1.10.1` (STALE).
- `demo-stack/migrate-demo.sh` ALREADY has the extensions-schema bootstrap (`CREATE SCHEMA extensions;
  CREATE EXTENSION vector/pgcrypto/pg_trgm SCHEMA extensions`), the bounded PG-readiness `wait_pg`, the
  sentinel-race defense, AND is M209-re-grounded (`for pair in app:public cms:cms
  jobsimulation:jobsimulation skillpath:skillpath`). → The M25-D9 class is already solved on the DEMO
  migrate path; the pre-surfaced gap is the DEV clean-bring-up path (platform `make reset-db`/`make
  migrate`, un-editable) which needs the same as a rext pre-migrate hook.

See `overview.md` for the strategy + baseline table; `decisions.md` for D1–D6.

## Close — 2026-07-08

**Outcome:** Authored TOK-01 "Warm-first cache-migrate, then cold-prove both stacks"; baseline 1/6 gate
sub-conditions met (compose), warm-only; next-tik direction set (target sub-condition (b): taxonomy
replay loads public.* ~42,790).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks never emit tok-fired) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (bootstrap tok continues into iter-02 tik within the same call)
**Decisions:** D1 (cache-migration mechanics), D2 (Phase 0b inherited-GREEN), D3 (demo path already M25-D9-solved; dev path is the gap), D4 (consumption re-pin plan), D5 (warm-first sequencing), D6 (recapture-prerequisite override by user cache-migration decision)
**Side-deliverables:** none (strategy authoring only)
**Routes carried forward:** none new (the whole gate is the forward queue; next-tik direction in overview.md)
**Lessons:** The demo migrate path is the working reference for the dev-side M25-D9 hook — mirror it, don't re-invent. The cache-migration is a mechanical re-key (manifest schema/payload/filter/public_via + 10 file renames + new digest key), gated on an empirical column-match, never fabrication.
