**Type:** tik (fourth tik under TOK-01) — stages 5–6 (boot + serve anonymously, the gate).

# M21 iter-05 — progress

## Work done (live, against the iter-04 stage-4 harness + sanctioned prod MCP read)
1. **Read prod's public-access model:** Directus's hardcoded public-policy UUID `abf8a154-…` (created by bootstrap +
   its `(role=NULL,user=NULL)` access link); 6 public `read` perms (simulations/skill_paths filtered to
   `status=published`; roles/sequences/sequences_roles/directus_files unfiltered).
2. **Registered the content collections + replicated the public read permissions** on `abf8a154` (`iter-05/serve.sql`).
   Booted Directus → healthy (stage 5 PASSES), but anonymous GET → **403** (even ADMIN → 403, while `/collections`
   worked).
3. **Root-caused via the Directus logs:** `Collection "X" doesn't have a primary key column and will be ignored` — for
   every content table. The `iter-04/structure.sql` created tables from a column-only `pg_catalog` DDL **without
   PRIMARY KEY constraints**. The digest still converged (it's over column *types*, not constraints) and COPY worked,
   but **Directus refuses to serve a PK-less collection** → 403.
4. **Fixed:** captured + applied the real PKs (`iter-05/pks.sql`; `id` for all, `code` for `languages`). Digest stayed
   `6cd35278…` (PKs don't change column data_type). 0 PK warnings on reboot.
5. **THE GATE: anonymous `GET /items/simulations?limit=1` → HTTP 200 with a real published row** (id
   `008b139f-…`, title "Team Conflict Management: Speed vs. Documentation", status published). **Stage 6 PASSES** —
   all 6 stages demonstrated end-to-end.

## Re-measure
- Pre-iter furthest-passing-stage: **4**. Post-iter: **6** (boot + anonymous serve demonstrated). Delta **+2**.
- **Honest caveat (gate not yet met-by-tooling):** the structure was applied BY HAND (`structure.sql` + `pks.sql` +
  `serve.sql`), not by `stacksnap`. The milestone exit_gate explicitly says "**stacksnap** applies the captured
  structure" — so the 6-stage PATH is proven end-to-end, but the gate's automation clause needs the code-ification
  (`STRUCT-M21-codeify`). → M21-D9.
- **directus_fields NOT required for the gate:** Directus introspects DB columns once a collection is registered
  (`directus_collections`) AND has a PK. The serve recipe = structure(+PK) → register → public read permission.

## Close — 2026-06-11

**Outcome:** Cracked the milestone's flagged live-only risk: a booted per-stack Directus serves a captured public
simulation anonymously over HTTP (200 + real published row). The PK constraint is the load-bearing structure piece
Directus needs (a column-only DDL is insufficient). furthest-passing-stage **4 → 6 (demonstrated end-to-end)**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the exit_gate's "stacksnap applies the captured structure" clause needs the code-ification; the
6-stage path is demonstrated with a hand-applied structure)
**Phase 5 grading:** (1) gate-met: n (path demonstrated, but stacksnap doesn't yet apply the structure — code-ification pending) — (2) triggered-tok: n (tik) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** M21-D9 (the serve recipe: PK requirement + register + public permission; directus_fields not needed; gate demonstrated-not-automated).
**Side-deliverables:** `iter-05/pks.sql` (the 26 PK constraints) + `iter-05/serve.sql` (register + public-permission recipe) — the proven serve artifacts feeding the code-ification.
**Routes carried forward (→ iter-06 under TOK-01):**
  - `STRUCT-M21-codeify` — **now the critical path to the gate:** make `stacksnap` capture the structure (26-collection
    DDL **including PKs** + the directus_collections registration rows + the public read permissions) over `--dsn`
    into the snapshot, and APPLY it before the row replay in provision. This is what flips the gate from
    demonstrated → met.
  - Carried: `directus_files` ref capture; M23 referential closure (the 20 dangling relations — not needed for the
    basic gate, confirmed).
**Lessons:** a schema-digest match (column types) is necessary but NOT sufficient for serving — Directus needs the
PRIMARY KEY constraint, which the digest doesn't see. Any structure-capture that targets a *served* engine must
capture constraints (at least PKs), not just columns. (Generalizes: "digest converged" ≠ "engine will serve it".)
