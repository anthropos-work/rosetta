**Type:** tik — under TOK-01. Protocol: overview §Iteration protocol (run bring-up phase → triage → route
fix → re-measure) + verification.md.

# iter-02 — tik progress

## Execution log
1. **Empirical column-match (D1 gate):** merged `public.skills` = 16 cols; cached capture = 15. The delta
   is exactly one added column `ts_search` (`GENERATED ALWAYS AS (setweight(to_tsvector(...name/aliases/
   keywords...)) STORED`, tsvector). Across all 10 tables the ONLY merged-but-not-cached columns are the
   `ts_search` generateds on skills/job_roles/skill_translations/job_role_translations. All NOT-NULL-no-
   default columns are present in the capture. **Drift is fully reconcilable** — a COPY must omit generated
   columns, and the cached manifest column lists (pre-`ts_search`) already omit it. → D1 gate PASSED; no
   `user-blocker`.
2. **Build:** `go build ./cmd/stacksnap` from the re-grounded authoring copy (HEAD 2f06e78 =
   quick-change-m209) → RC 0.
3. **Digest probe:** `stacksnap capture --surface taxonomy --dsn localhost:5432 --dry-run` → merged-public
   schema digest = **`5afc0bccf1df7ef538b643321fc6362f`** (was `c75ce94…` under skiller). Filter SQL
   confirms `"public"."skills"` etc. This is the key a cold `/dev-up` will probe → the migrated cache must
   live here.
4. **Cache-migration:** new dir `.agentspace/snapshots/taxonomy/5afc0bccf1df7ef538b643321fc6362f/`;
   hardlinked the 10 payloads under `public.X.copy` names (link count 2 → zero extra disk, old `c75ce94…`
   cache preserved); transformed manifest (schema `skiller`→`public` on all 10 tables; payload/filter/
   public_via re-keyed; `schema_version`=new digest; migration note appended). **Payload BYTES unchanged →
   the manifest sha256 checksums stay valid** (replay verifies them before writing — passed).
5. **Replay:** `stacksnap replay --surface taxonomy --stack "" --dsn localhost:5432` → rc 0, "10 table(s)
   cleared, 10 table(s), **330,261 row(s) loaded**, reindexed [public.skill_embeddings.small_embedding3
   public.job_role_embeddings.small_embedding3]", 108s.
6. **Verify:** `public.skills` = 42,790 (all `organization_id IS NULL`); `ts_search` filled on all 42,790;
   job_roles 22,470 / skill_embeddings 42,790 / job_role_skills 72,705 / specializations 1,447 /
   categories 23; sample skills = real names + valid node_ids (K-INSDES-33ED, K-IOTSEC-80F2, K-ORANET-78FF).
7. **Consumption re-pin (D4):** `.agentspace/rext.tag` v1.10.1 → `quick-change-m209` (gitignored config;
   `git status` clean → no rosetta tree change). The bring-up loop now consumes the re-grounded tooling.

## Re-measurement (gate sub-conditions, warm)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) 4-subgraph / no-skiller compose | MET | MET |
| (b) replay loads public.* (~42,790) | NOT MET | **MET** (330,261 rows; skills 42,790; ts_search filled) |
| (c) seed closure green | NOT MET | NOT MET |
| (d) verify merged-assertion | NOT MET | NOT MET |
| (e) M42 coverage + v2.0 Playthroughs | NOT MET | NOT MET |
| (f) 0 residual skiller refs in queried paths | code/corpus clean | +migrated cache/tooling now public.* |
**Metric:** gate sub-conditions met (warm): 1/6 → **2/6** (delta +1, met expected lift).

## Close — 2026-07-08

**Outcome:** Cache-migration + replay → `public.*` taxonomy loaded into the warm merged stack (42,790
public skills, 330,261 total rows, ts_search auto-generated); consumption re-pinned to the re-grounded
tooling. Sub-condition (b) GREEN. Metric 1/6 → 2/6.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (2/6 sub-conditions; composite gate needs all-6 on BOTH stacks cold)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-02 D1 (drift reconcilable — empirical), D2 (hardlink zero-copy migration), D3 (key under real probed digest for cold-hit), D4 (re-pin now)
**Side-deliverables:** none (re-pin was planned prep under TOK-01, not a side-discovery)
**Routes carried forward:** none new — next-tik direction (iter-03): sub-condition (c) seed closure green (run the seeder against the now-taxonomy-loaded warm stack + measure datadna closure).
**Lessons:** The user's cache-migration is proven mechanically faithful: a schema-prefix re-key + hardlink (bytes unchanged → checksums hold) + key under the live-probed merged digest → replay hits + loads the real 42,790-row public taxonomy at $0, no prod access, no fabrication. The added `ts_search` is the ONLY drift and is self-healing (GENERATED). This pattern is the template for any future skiller→public re-key.
