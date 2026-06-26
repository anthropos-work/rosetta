# M45 — Generation engine · Retro

## Summary
The cheap-LLM batch profile generator — a YAML batch descriptor → realistic per-member profiles via
gpt-4o-mini, with the **CODE-owns-structure / AI-owns-content** boundary enforced by the existing taxonomy
closure gate. Built inside-out, fixtures-first (TOK-01) over **7 iters** (1 bootstrap tok + 6 tiks; the
re-scope trigger never fired). Closed **on-gate (5/5)** with **3 findings / 0 blocking**. The gate was MET on
the **first real LLM run** — an Azure gpt-4o-mini batch (EU-first Sweden; the direct `OPENAI_KEY` was
billing-dead/429), N=20 + a fresh demo-3 proof. The THIRD milestone of the M43→M46 extension, and the engine
M46 (org-scale fill) builds on. **The deliberate, user-acknowledged supply-chain inflection** — the FIRST new
third-party dep since v1.8 (`anthropos-work/ai@v1.40.1`), the thing this milestone is about. Zero canonical
platform-repo edits.

## Incidents This Cycle
- **No P0/P1 incidents, no flakes, no post-close regressions.** `go test -race ./...` clean; the 3×/5× flake
  gate (incl the timing-sensitive `--call-timeout` test) clean.
- **One real edge bug surfaced + fixed inline during harden Pass 4 (P2, NOT a post-close regression):** the
  `genEmail` degenerate-address edge — a separator-only local part (an empty first+last name yields `"."`)
  produced `.@domain`. `sanitizeEmailLocal` now collapses a separator-only result to `""` so `genEmail`
  falls through to the name, then to `member@…` — a valid address is ALWAYS produced (the
  CODE-owns-identity guarantee). Fix + regression test in the same commit (rext `d00c606`). Defensive reach
  (the upstream `env.Name == ""` drop already prevents the empty-name path in production), <10 lines, single
  subsystem — Fate 1.
- **Three issues the first REAL run surfaced (all fixed in iter-07, the gate-proving iter):**
  1. **No per-call timeout** (the stall class) — the run loop used `context.Background()` with no deadline;
     a long silent Azure batch could hang. FIX: a `--call-timeout` flag (default 60s) wraps each call so a
     hung endpoint fails FAST. (Plus a connectivity smoke-test before the batch.)
  2. **Intra-batch name duplication** — gpt-4o-mini is strongly name-sticky per mother-prompt (~9 distinct
     names / 20 on the raw run). FIX: the system prompt demands a varied/multicultural name + names the
     over-used Anglo defaults to avoid, and a re-roll carries an "avoid these already-used names" hint built
     from the batch's used-name set + the reserved heroes → **20/20 distinct names**.
  3. **`user_skills` CHECK violation (23514)** — generated claimed-skill rows left every provenance edge
     NULL, violating `user_skills_check_foreign_keys`. FIX: the seeder now writes ONE company + ONE
     current-role `user_experiences` row per generated member and ties each claimed skill to it (stays
     shallow — heroes keep the deep timelines).
- **An Azure env-name mismatch the real run exposed (not a bug, a config-naming gap):** the platform's
  services name Azure config differently (`AZURE_API_KEY`+`AZURE_ENDPOINT` = the prod Sweden/EU deployment;
  `AZURE_OPENAI_KEY`+`AZURE_OPENAI_ENDPOINT_URL` = eastus2; direct = `OPENAI_KEY` not `OPENAI_API_KEY`).
  FIX: `NewFromEnv` now resolves the Azure pair from a prioritized **EU-first** list + reads the direct key
  from both names. Values-blind preserved.

## What Went Well
- **Inside-out, fixtures-first (TOK-01) isolated the one empirical unknown.** Every layer except the model's
  English content is deterministic and was unit-proven against a fixture `ai.AI` (no key, no token) BEFORE
  the gate-proving run — so the only thing the real run had to discover was the cheap model's raw
  valid-JSON / resolution rate. The gate hit 5/5 on the first real batch.
- **The CODE-owns-structure / AI-owns-content boundary held under adversarial fuzz.** On a hallucination-heavy
  envelope every written role/skill node-id is a REAL one from the conn's pools, never a fabricated `J-…`/
  `K-…`; on-disk corruption drops cleanly (no panic). `datadna measure-closure` = `[PASS]` on the real run.
- **The cost discipline worked first try** — $0.0059 of the $0.10 ceiling (5.9%), and the $0 byte-identical
  re-seed (run 2 = 0 calls / 20 cache hits) proves the deterministic-expansion + prompt-hash-cache design.
- **The supply-chain inflection was clean.** Exactly 1 dep added (`ai v1.40.1`, all-permissive tree);
  the 5-pass harden added 0 further deps and `go mod tidy` is a no-op.

## What Didn't
- **The raw cheap model is name-sticky** — a bare seed change re-picks the same name; without the prompt
  + re-roll-hint fix the run gave ~9 distinct names / 20. A general lesson for any per-item LLM-generation
  engine (recorded in `ai-generation-spec.md`): re-rolls must carry an explicit avoid-set, not just a seed.
- **"Resolution" is measured against the `node_id` column, not the uuid `id`.** An early verification query
  joined on `id` and falsely read "all fabricated"; the resolvers SELECT `node_id` (the `K-`/`J-` NodeID).
  Noted so future closure checks don't repeat the false negative.
- **The cache-spec placement was deferred to the bootstrap tok** (overview said `corpus/ops/cache-spec.md`
  OR a section in `ai-generation-spec.md`); it landed at `corpus/ops/demo/cache-spec.md` per the
  demo-family index convention. The close caught two index links that still pointed at the old
  `corpus/ops/` path (via the Phase-8 cross-reference check) and fixed both.

## Carried Forward
- **None as a `closed-incomplete` carry-forward** — the gate is MET. The `Out:` items are owned by design
  (Fate-2, already in their `In:` lists), not punts:
  - **org-scale auto-fill to full org size → M46** (the ONLY remaining v1.10 milestone; its `In:` owns
    "count: auto-fill to org size … expanding to fill the remaining N members"). M45 proved the engine on a
    **bounded** N=20 batch; population-scale behaviour is M46's iterative gate.
  - **production-seeding key story → out of v1.10 release scope by design** (a documented Open question;
    `.env.local OPENAI_API_KEY` env var for now; orthogonal to v1.10's tooling-+-docs / zero-platform-edit
    boundary).

## Metrics Delta (from metrics.json)
- **Go test funcs (stack-seeding):** 567 → **677** (+110): the 5 new pkgs (services/ai 32 + blueprint 45 +
  batchcache 21 + cmd/gen-batch 23) + the GeneratedBatchSeeder/drop-not-fabricate/boundary-fuzz tests in
  seeders/, across iter-02..07 build + the 5-pass final harden. Module total 1406 → **1516**. clerkenstein/
  snapshot/alignment/secrets UNCHANGED (M45 touched stack-seeding only — verified empty clerkenstein diff).
- **Coverage (new pkgs):** services/ai 97.8% · blueprint 98.9% · batchcache 88.5% · cmd/gen-batch 93.0% ·
  seeders 97.3% (5-pass final harden; 0% delta at Pass 5, dimension scan exhausted).
- **Supply-chain:** +1 deliberate dep (`ai v1.40.1`). **Flakes:** 0. **Alignment:** 100%/100% (N/A change).
- **Cost (real gate run):** $0.0059 / $0.10 ceiling; $0 byte-identical re-seed.
