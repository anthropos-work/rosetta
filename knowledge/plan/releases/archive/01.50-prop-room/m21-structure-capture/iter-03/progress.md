**Type:** tik (second tik under TOK-01).

# M21 iter-03 — progress

## Work done
1. **Self-contained sources investigated + exhausted** — `stack-dev/cms/internal/directus/collections/*.go` is a
   lossy read-side app-view (field names + Go types + relational aliases, not Postgres types/registry). -> M21-D6.
2. **Operator sanctioned a prod structural read** (their words: the prod-DB skill should read the directus schema too;
   directus is in the same prod Postgres). Blocker RESOLVED; saved as a standing policy memory.
3. **Captured the real faithful structure for all 9 collections** via bounded read-only `pg_catalog` reads over the
   wired `postgres` MCP:
   - exact DDL per table (simulations 37 cols / skill_paths 28 / sequences 34 / resource 19 / roles 17 / sim_tasks 11
     / task_checks 8 / task_sub_checks 10 / sequences_roles 10) -- uuid / json / text / varchar(N) / timestamptz,
     defaults + the `sequences_roles_id_seq` serial. Column sets match the manifest exactly.
   - registry inventory: **9** `directus_collections` / **217** `directus_fields` / **43** `directus_relations`
     (20 of which point OUTSIDE the 9 -> the M23 referential-closure surface).
4. **Decisive digest finding (M21-D5 -> option B):** prod digest `6cd35278…` = the FULL **53-table** directus schema
   (27 `directus_*` system + **26** user collections). Our surface captures **9** of 26. So a bootstrapped stack +
   9-collection structure can never converge the whole-schema digest -> the cache MUST be **re-keyed per-surface**
   over only the captured content tables (option B) for the bootstrap+structure model to ever cache-hit. (Option A --
   capture all 26 collections + pin the Directus version -- is the alternative.) Architectural fork, touches the
   shared `pg.SchemaVersionSQL` keying (taxonomy too) -> surfaced to the operator.

## Close -- 2026-06-11

**Outcome:** structure-source blocker RESOLVED (operator-sanctioned prod structural read); the **real faithful
structure** for all 9 collections captured (DDL + registry inventory); the **digest-keying fork resolved in principle
toward option B** via the decisive 26-vs-9 / full-schema-digest evidence. furthest-passing-stage stays **2** -- the
artifact is not yet APPLIED to a live harness (that + the per-surface re-key is iter-04).
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n (tik) -- (3) re-scope: n -- (4) user-blocker: **y** (the
digest-keying architectural fork -- option A vs B -- touches the shared `pg.SchemaVersionSQL` keying used by taxonomy
+ has prod-safety/maintainability implications; the operator should steer it before iter-04 commits code) -- (5)
cap-reached: n (2 tiks) -- (6) protocol-stop: n -- Outcome: exit-4 (user-blocker)
**Decisions:** M21-D6 (source operator-gated, RESOLVED), M21-D5 refined (-> option B, the 26-vs-9 evidence); iter-local
in `decisions.md`.
**Side-deliverables:** standing policy memory (operator sanction of the prod structural read for M21).
**Routes carried forward:** `STRUCT-M21-iter04-apply` (apply the real structure -> stage 3), `STRUCT-M21-digest-keying`
(implement option B), `directus_files` wiring, M23 referential closure of the 20 dangling relations.
**Lessons:** the cache being keyed by the WHOLE-schema digest (not the captured-surface's tables) is invisible until
you read the source schema and count tables (26 collections, not 9). A staged-pipeline tik that "only reads, doesn't
apply" still produces a load-bearing architectural finding -- the source read WAS the deliverable here.
