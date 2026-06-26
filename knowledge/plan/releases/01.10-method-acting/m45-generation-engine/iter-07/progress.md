# iter-07 — progress

**Type:** tik (under TOK-01, inside-out fixtures-first build — step 7: the real gate-proving run)

The FIRST real LLM run, moving the gate from **0/5 (fixture-only)** to a real, reproducible measurement.
The engine was CODE-COMPLETE (iter-02..06); this tik fired the real Azure gpt-4o-mini N=20 batch + the
demo-3 proof. (Run-2 of this tik stalled on a long silent Azure batch with no per-call deadline — the
watchdog killed it; this RETRY re-ran with frequent journal heartbeats + a per-call timeout, the fix.)

## What ran (per ai-generation-spec.md §4b)

1. **Connectivity smoke-test FIRST** (cheap, short timeout) — ONE completion via the EU-first Azure routing:
   gpt-4o-mini, 2.2s, valid JSON. Confirms a LIVE deployment (NOT a user-blocker) + fails fast on a hung
   endpoint. Model/deployment used for the whole gate: **gpt-4o-mini via EU-first Azure (Sweden)
   `AZURE_API_KEY`+`AZURE_ENDPOINT`** (the direct `OPENAI_KEY` is billing-dead — HTTP 429 insufficient_quota).
2. **Real N=20 batch** (`--max-cost 0.10 --max-concurrent 5 --call-timeout 60s`), heartbeated per ~60-90s.
3. **Hardened** three issues the run surfaced (per-call timeout, intra-batch name dedup, the `user_skills`
   CHECK), each with a fix in rext + a re-run.
4. **demo-3 re-seed** via the tagged `stackseed --cache-root` mechanism → measured resolution + closure.
5. **$0 byte-identical re-seed** (run 2 → 0 calls / $0 / byte-identical cache).

## The three fixes (all in rext, tagged `method-acting-m45-iter07-gate`)

- **Per-call timeout** (`cmd/gen-batch --call-timeout`, default 60s) — the stall-class fix.
- **Intra-batch name de-dup** — gpt-4o-mini is name-sticky per mother-prompt (raw run: ~9 distinct / 20).
  System-prompt diversity demand + a re-roll "avoid these names" hint → **20/20 distinct multicultural names**
  at $0.0059. New unit test `TestRun_IntraBatchNameDedupReroll`.
- **`user_skills` CHECK (23514)** — the claimed rows left every provenance edge NULL. FIX: seed ONE company +
  ONE current-role `user_experiences` per generated member; tie claimed skills to it via
  `user_skill_experience`. Plus a reproducible `stackseed --cache-root` flag.

## Close — 2026-06-26

**Outcome:** the gen-acceptance gate moved from **0/5 (fixture-only)** to **5/5 PASS on a real Azure
gpt-4o-mini N=20 batch + the demo-3 proof.** valid-JSON 100% (33/33) · resolution 47/47 skills + 20/20 roles
→ real `skiller.node_id`, `datadna measure-closure --stack demo-3` = `[PASS]` (closure GREEN, 0 fabrication) ·
0 hero-collisions · cost $0.0059 ≤ $0.10 · $0 byte-identical re-seed. Believability: 20/20 distinct
multicultural names, role-coherent resolving skills, 20 current-role experiences, 20/20 avatars, isolation
CLEAN.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (per-call timeout), D2 (intra-batch name dedup via prompt+hint), D3 (current-role experience for the FK CHECK), D4 (reproducible stackseed --cache-root) — see iter-07/decisions.md
**Side-deliverables:** none (all changes were in the iter's planned gate-proving scope).
**Routes carried forward:**
- M46 (org-scale fill): the per-story batch routing + scaling the batch beyond N=20 (M45 proves the engine on
  a bounded batch within the existing org — explicitly OUT of M45 per §5).
- Note (M46 quality polish, not a gate item): a prompt change is invisible to the cache key (key = mother
  prompt + capture-version, not the system prompt), so a prompt-hardening re-run needs a fresh cache root.
  Acceptable for dev-throwaway caches; flag if prompt evolution becomes frequent.
**Lessons:**
- A real LLM gate-proving run MUST heartbeat the journal during the batch + bound each call with a timeout —
  a long silent batch is the watchdog-stall class (the run-2 failure). This is now codified in the engine
  (`--call-timeout`) and the work-discipline.
- Cheap models are strongly name-sticky per-prompt; effective re-rolls need the model TOLD which names to
  avoid, not just a seed bump. (Generalizes to any per-item LLM-generation engine — recorded in the protocol.)
- "Resolution" is measured against the taxonomy's `node_id` column (the `K-`/`J-` NodeID), NOT the uuid `id`
  — a verification query joining on `id` falsely reads "all fabricated."
