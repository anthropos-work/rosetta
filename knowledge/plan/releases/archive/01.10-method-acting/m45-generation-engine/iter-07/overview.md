---
iteration_type: tik
status: closed-fixed
---

# iter-07 — the FIRST real gpt-4o-mini gate-proving run

**Active strategy reference:** TOK-01 (inside-out fixtures-first build) — step 7: "Gate-proving: a real
N=20 capped batch (`OPENAI_KEY` from `stack-demo/platform/.env`, values-blind) + the $0 byte-identical
re-seed." Engine is CODE-COMPLETE (iter-02..06, fixture-proven). This tik fires the FIRST real LLM run.

**Re-survey (Step 0):** the gate is **0/5 EMPIRICALLY** — no real LLM call has happened. iter-06 closed
the last code component (GeneratedBatchSeeder) at the 5-tik cap; run-1 carried forward "NEXT CALL: the
EMPIRICAL gate-proving." Target is still untouched + still the right next thing. No substitution.

**Cluster / target identified:** the 5-vector gate must move from 0/5 (fixture-only) to a real
measurement. The one empirical unknown (the cheap model's raw valid-JSON / skill-name-resolution rate) is
answerable only by a real run. This tik runs it.

**Hypothesis:** the engine as built (DefaultBatchPromptTemplate forcing JSON mode + strict envelope, the
re-roll loop, the cost ceiling, the prompt-hash cache) will, on a real N=20 gpt-4o-mini batch:
- emit valid JSON ≥95% pre-re-roll (JSON-mode + the strict envelope instruction);
- produce skill/role names that resolve through the existing resolvers (the prompt says "use real,
  commonly-recognized names");
- collide with ZERO reserved heroes (the reserved-name prompt instruction + post-gen re-roll);
- land cost ≪ the $0.10 ceiling (~$0.005 expected);
- re-seed byte-identical at $0 on a second run (the cache).

**Expected lift:** gate 0/5 → measure all 5 on a real batch. Success = a clean measurement of each
dimension (whether or not all 5 PASS on the first run — a short dimension routes to a hardening tik).

**Phase plan (per ai-generation-spec.md §4b):**
1. Probe/measure — author a `gen-batch-20.seed.yaml` (N=20, real roles/seniority/bias), build the CLI,
   fire the real run (`OPENAI_API_KEY` exported from `OPENAI_KEY`, values-blind), capture the cost report
   + valid-JSON rate.
2. Diagnose — read the 20 cached envelopes: valid-JSON rate, collisions, cost vs ceiling.
3. (resolution + closure are measured against a live taxonomy-replayed stack in a later tik / Phase 4
   demo-proof; this tik measures the model-quality dimensions the cache alone proves).
4. Re-measure — the $0 byte-identical re-seed (run 2 → 0 calls, identical bytes).
5. Close.

**Escalation conditions:** valid-JSON < 95% after the built-in re-roll budget → harden the prompt next tik
(not a re-scope yet; that's after ~5 hardening tiks). A real correctness break (fabricated node-id /
broken closure on the demo-proof) → user-blocker.

**Acceptable close-no-lift outcomes:** if the real run reveals the model needs prompt hardening (valid-JSON
< 95%), the iter closes-fixed on the MEASUREMENT (0/5 → measured) and routes the hardening forward — the
measurement is the deliverable, the gate-pass may take another tik.
