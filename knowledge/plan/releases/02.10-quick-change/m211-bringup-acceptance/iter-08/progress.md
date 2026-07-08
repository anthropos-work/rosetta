**Type:** tik (bring-up: cold /demo-up proof) — under TOK-01, move (4) "prove COLD". 0 code/corpus/rext
changes; the deliverable is the COLD proof + its recorded evidence. Protocol: re-pin consumption → tear down
warm dev → cold bring-up (reap-safe) → auto-verify + explicit sub-condition measurement.

# iter-08 — tik progress

## Execution log
1. **Pre-flight (Phase 0 survey).** stack-demo/platform + stack-demo/app confirmed at MERGED code (git log:
   "remove skiller service from docker-compose and repos.yml"; app "retire the app→skiller RPC path #989");
   no live demo containers; secrets dir present; Docker RAM ~9.7 GB (below the 12 GB UI rec — non-fatal).
   The rext consumption clone `stack-demo/rosetta-extensions` was STALE (v1.10.1) and the pin
   `quick-change-m209` is local-only (never pushed).
2. **Re-pin consumption clone.** Local-fetched the `quick-change-m209` tag from the authoring copy
   (`.agentspace/rosetta-extensions`) into `stack-demo/rosetta-extensions` (local never-pushed-tag op —
   allowed) + `git checkout quick-change-m209` (HEAD 2f06e78). ensure-clones later confirmed at bring-up:
   "rext pin: consuming rosetta-extensions @ quick-change-m209 (matches .agentspace/rext.tag)".
3. **Tore down the warm dev stack** (ONE stack at a time) — `make down` + a full-profile
   `docker compose -p anthropos down --remove-orphans` (11→0 containers, network removed) → freed Docker RAM.
4. **Cold `/demo-up` (reap-safe).** Launched `up-injected.sh 1` as a DETACHED nohup process
   (`.agentspace/scratch/work-m211/run-demo-up.sh` → `demo-up-1.log`), polled in-turn (heartbeats). Cold
   build completed **DONE-rc=0 in ~6.5 min** (16:46:09→16:52:35 — images were cached from a prior demo, so
   no 20-45-min rebuild).
5. **Sub-condition measurement (COLD):**
   - **(a) 4-subgraph / no skiller container:** 16 `demo-1-*` containers up, **no skiller container**
     (backend/cms/graphql/jobsimulation/skillpath/sentinel/storage/roadrunner/gotenberg/postgresql/redis +
     demo-specific directus/fake-bapi/fake-fapi/next-web-app/studio-desk). ✓
   - **(b) replay loads public.*:** "replayed taxonomy into demo-1: 10 tables, 330261 rows, reindexed
     [public.skill_embeddings … public.job_role_embeddings]"; "taxonomy rows=42790 ok". The migrated
     `public.*`-keyed cache HIT on a cold replay automatically. ✓
   - **(c) seed closure green:** stories seed (2 orgs × 3-hero trio) — "Audit: 49 writes, 67069 rows,
     prod=false, isolation: clean". Explicit `datadna measure-closure --stack demo-1` →
     **[PASS] seed-verified-skill-closure** (every seeded verified-skill node-id resolves in the replayed
     taxonomy). ✓
   - **(d) verify merged-assertion:** bring-up-tail autoverify — ✓ backend /api/health 200 on :18082, ✓
     sentinel.casbin_rules = 1304, ✓ directus.directus_collections = 21 (per-stack-local, not prod), ✓
     verify live: all liveness + readiness passed → "▶ autoverify demo-1: OK — verified-working." (no skiller
     schema/subgraph/container expected or found). ✓
   - **(f) 0 residual skiller:** public schema throughout; no skiller container, no skiller schema in
     replay/verify. ✓
6. **Observations (non-gate, non-fatal):** sim-embeddings replay skipped (rc=5, cache miss — no cached
   snapshot for this schema digest); this is the simulation-embeddings surface (vector search over sims),
   NOT the taxonomy and NOT a gate sub-condition — the demo is fully functional without it. Route forward as
   an optional cache-fill.

## Re-measurement (gate sub-conditions — COLD, demo)
| Sub-condition | Pre-iter (warm-only) | Post-iter (demo, COLD) |
|---|---|---|
| (a) compose / no-skiller | MET (warm) | **MET COLD** (demo-1: 16 containers, no skiller) |
| (b) replay loads public.* | MET (warm) | **MET COLD** (42,790 public skills, migrated cache HIT) |
| (c) seed closure green | MET (warm) | **MET COLD** (datadna measure-closure PASS on demo-1) |
| (d) verify merged-assertion | MET (warm) | **MET COLD** (autoverify demo-1 OK, no skiller) |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET (next: tik-09/10 against live demo-1) |
| (f) 0 residual skiller-schema refs | MET | **MET COLD** (public schema throughout demo-1) |
**Metric:** the DEMO half of the "both stacks GREEN cold" headline is **DONE** — a-d,f now COLD-PROVEN on
demo-1 (was warm-only). Remaining: (e) coverage+Playthroughs (against this live demo) + the cold `/dev-up`
half.

## Close — 2026-07-08

**Outcome:** Cold `/demo-up` GREEN on the merged platform (DONE-rc=0, ~6.5 min) — demo-1 up with 4-subgraph/
no-skiller compose, migrated `public.*` cache HIT on cold replay (42,790 skills), stories seed closure PASS
(datadna), autoverify merged-assertion OK, 0 skiller residue. The demo half of the headline cold proof is
complete + it's the live UI-tier substrate for (e).
**Type:** tik (bring-up: cold /demo-up proof)
**Status:** closed-fixed
**Gate:** NOT MET (demo half of the cold headline proven; (e) coverage/Playthroughs + the cold `/dev-up`
half remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (cold build came up green; RAM was borderline but sufficient after teardown) — (5) cap-reached: n (2nd tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-08 D1 (cold /demo-up GREEN cold on the merged platform — a-d,f COLD-proven), D2 (re-pin
consumption clone via local never-pushed tag fetch — the gate-critical prerequisite; the tag isn't on GitHub
and won't be until close-release)
**Side-deliverables:** none (no code/doc changes — a bring-up proof).
**Routes carried forward:** tik-09 → M42 coverage sweep (coverage-protocol.md presence gate) against live
demo-1 (offset 10000: next-web :13000, cockpit :17700). tik-10 → v2.0 Playthroughs against demo-1. OPTIONAL:
sim-embeddings cache-fill (non-gate; separate operator-confirmed capture). Keep demo-1 UP for tik-09/10.
**Lessons:** The migrated `public.*`-keyed taxonomy cache HITs automatically on a COLD replay (no manual
re-key needed at bring-up — iter-02's migration made the cold path self-serve). The demo cold path is fully
automated end-to-end on the merged platform (migrate-demo.sh's extensions+casbin+PG-wait already handle the
M25-D9 class) — a cold /demo-up needs only the consumption-clone re-pin as a manual prerequisite, which
close-release's real push will make automatic.
