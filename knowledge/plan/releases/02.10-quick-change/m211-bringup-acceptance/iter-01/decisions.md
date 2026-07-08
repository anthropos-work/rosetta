# iter-01 — Decisions (bootstrap tok)

## D1 — Cache-migration mechanics (the load-bearing novel piece)
The user chose to MIGRATE the existing real taxonomy cache (`c75ce94…`, 42,790 real skills captured
under the old `skiller.*` schema) → re-key to `public.*`, NOT do a prod COPY-byte recapture (no
`marco_read` DSN locally; postgres MCP is query-only). The skiller→app merge was a pure schema-prefix
move (same data, same table names). Migration = a mechanical re-key of the manifest + payloads:
1. Each `tables[].schema`: `"skiller"` → `"public"`.
2. Each `tables[].payload`: `"skiller.X.copy"` → `"public.X.copy"` + rename the 10 payload files on disk.
3. Internal `tables[].filter` SQL refs: `"skiller"."skills"`/`"skiller"."job_roles"` → `"public"."…"`.
4. `tables[].public_via`: `"skiller.skills"`→`"public.skills"`, `"skiller.job_roles+skiller.skills"`→`"public.…"`.
5. The cache-dir key `schema_version` digest: M209 narrowed the digest, so the re-grounded replay computes
   a DIFFERENT key → the migrated cache must land under the NEW digest dir (resolve empirically in iter-02).
**GATE:** verify EMPIRICALLY the cached `skills` column-set (15 cols: id,name,description,aliases,is_soft,
parent,node_id,keywords,created_at,parent_node_id,updated_at,created_by_user,organization_id,deleted,
long_description) == merged `public.skills` columns BEFORE trusting the replay (`\d public.skills`). If a
genuine column drift makes it unreconcilable → escalate as `user-blocker`. NEVER fabricate rows.

## D2 — Phase 0b: inherited-GREEN, not a full re-audit
M210 (closed same day, 2026-07-08) WAS the corpus/skills re-ground and closed with a GREEN
audit-kb-fidelity (KB-1/2/3 resolved). Orchestrator bans re-deriving M210 facts. Ran a proportionate
spot-confirmation (0 residual `skiller.<table>` in tooling-queried corpus paths) → standing GREEN holds.
This is the skill's "already-run / inherit standing verdict" path for the milestone-once gate, applied
because the immediately-prior milestone delivered the audit today against this exact corpus.

## D3 — The M25-D9 class is already solved on the DEMO migrate path; the DEV path is the gap
`demo-stack/migrate-demo.sh` already bootstraps the `extensions` schema + vector/pg_trgm/pgcrypto
extensions, waits for PG readiness (`wait_pg`), defends the sentinel casbin race, and is M209-re-grounded
(app:public + 4 services). So the pre-surfaced M25-D9 fix does NOT need re-authoring for demo. The gap is
the DEV cold bring-up (main dev N=0 drives the platform Makefile's `make reset-db`/`make migrate`, which
M211 cannot edit) — the extensions-bootstrap + PG-wait must land as a rext DEV pre-migrate hook mirroring
migrate-demo.sh. iter-02 confirms whether the dev path already has such a hook or needs one.

## D4 — Consumption re-pin to the re-grounded tooling
`.agentspace/rext.tag` = `v1.10.1` is STALE (pre-merge, skiller-keyed). The bring-up must use the
re-grounded tooling. Plan: re-pin `.agentspace/rext.tag` → `quick-change-m209` (the tag at authoring HEAD
`2f06e78`), OR drive the bring-up directly from the authoring copy `.agentspace/rosetta-extensions/`.
Prefer re-pin (reproducible, matches the consumption contract). The final `v2.1` roll is close-release's
job; M211 needs the re-grounded tooling IN the loop now.

## D5 — Warm-first sequencing (speed), cold-prove for the gate
The merged warm stack-dev is UP (M208 de-risk). Iterate the fast inner loop (reset-db →
extensions-bootstrap → migrate → cache-migrated replay → seed → verify) on the warm stack to shorten the
fix→re-measure cycle and dodge the docker-reap hazard on long cold `make up`. The gate itself demands a
full COLD `/dev-up` + `/demo-up`, so a cold acceptance run is required — but only after the inner loop is
green warm. Long docker runs execute SYNCHRONOUSLY within a tik (or detached + polled same-turn), never
backgrounded-then-yielded (M208 reap lesson).

## D6 — Recapture-prerequisite (overview §Pre-surfaced) OVERRIDDEN by the user's cache-migration decision
The overview's "recapture via `stacksnap capture --dsn <restored prod pg_dump>`" path assumed a prod
source would be provisioned. The user confirmed proceeding WITHOUT prod access via the D1 cache-migration
instead. This is faithful to what the merge did (schema-prefix move on identical rows), not fabrication.
The MinRows floor (~42,790) still auto-catches an under-capture on replay.
