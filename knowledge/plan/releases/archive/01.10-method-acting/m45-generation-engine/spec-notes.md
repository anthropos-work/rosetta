# M45 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md)
(the generation-engine strand). M45 is **entirely in `rosetta-extensions` — zero platform-repo edits**. It
breaks the seeding module's 0-new-deps streak (the FIRST new third-party dep — a deliberate in-release
decision). The iteration protocol is the NEW `corpus/ops/demo/ai-generation-spec.md` (this milestone
authors it).

## `services/ai/` wrapper (the first new third-party dep)
TODO: a local wrapper around the shared `ai` library — EU-first routing → 429 fallback → usage tracking.
Do not reimplement provider routing (the shared `ai` lib + each consumer's wrapper own it); this is the
seeding-module's thin consumer.

## `blueprint.Batch` + `batch[]` + `EffectiveBatches()`
TODO: a `Batch` type + a `batch[]` field on `Story` / `StackSeed`; `EffectiveBatches()` does Go-template
per-member-prompt expansion at PARSE time — NO LLM at parse time (the LLM runs only in `cmd/gen-batch`).

## `cmd/gen-batch` CLI
TODO: gpt-4o-mini default; `--max-concurrent` semaphore (default 5); atomic `.tmp`→rename cache writes; a
`.lock` fence; a MANDATORY `--max-cost` ceiling (abort before exceeding). Emits `Event_AiUsage` + a
per-batch cost report.

## Prompt-hash cache (→ `cache-spec.md` or an ai-generation-spec.md Caching section)
TODO: `.agentspace/.batchcache/batch-${hash}/member-${i}.json`, keyed by the MOTHER prompt; cache key
EXTENDED with the taxonomy capture version (invalidate on re-replay). Atomic writes; the `.lock` fence.
Reproducibility target: an unchanged descriptor re-seeds byte-identical from cache at $0.

## `GeneratedBatchSeeder` (surface `'generated-batch'`)
TODO: `DependsOn` users+taxonomy, `PerStackIsolated`. Reads cached JSON; drives users/persona/profile rows
DETERMINISTICALLY via the existing resolvers (`resolveJobRoleRefs` / `resolveNamedSkillRefs`). The
CODE-owns-structure / AI-owns-content boundary: a non-resolving role/skill name DROPS (worst case a
SHALLOWER profile, never invalid); the closure gene (`datadna measure-closure`) stays green.

## Hero-name collision avoidance
TODO: reserved-hero-names in the prompt (the curated hero roster) + a post-gen collision re-roll fallback.
Exit gate requires ZERO generated name colliding with a hand-curated hero.

## Exit-gate measurement (per iter)
TODO: per-batch metrics — valid-JSON rate (≥95% + re-roll), taxonomy-resolution rate (every name resolves
or drops, closure green), collision count (=0), cost vs `--max-cost`. Plus the byte-identical-from-cache
re-seed at $0.

## Secrets (deferred for production)
`OPENAI_API_KEY` via `.env.local` env var (git-ignored) for now. NO platform-repo secrets store;
production-seeding key story deferred (an Open question).

## Pre-flight audits — iter-01
**`/developer-kit:audit-kb-fidelity --milestone=M45` → GREEN** (2026-06-26; report
[`kb-fidelity-audit.md`](kb-fidelity-audit.md)). All 5 KB-dependency topics PAIRED + ALIGNED; the 2 NEW
docs (`ai-generation-spec.md`, `cache-spec.md`) are intentional DOC-ONLY Delivers targets, not blind
areas. Load-bearing contracts confirmed against code/docs:
- the `ai` library is `github.com/anthropos-work/ai` (pinned `v1.40.1` across consumers, Go 1.25.0);
  interface `ai.AI` (`ChatCompletion`, …); constructors `openai.New`/`NewOpenAI` (direct),
  `openai.NewAzure` (Azure EU). `MetaData.Usage` carries **token counts only** — dollar cost is the
  CONSUMER's job (`app/internal/aiusage`, model→price switch, `Event_AiUsage` over Redis Streams).
- **EU-first routing + cost tracking are NOT in the lib** — they live in each consumer's
  `internal/ai/ai.go`. The M45 `services/ai/` wrapper is exactly such a thin consumer (do not
  reimplement provider internals).
- the resolver seam is real + drop-safe: `jobroleref.go forName()` → zero `jobRole{}` on no-match;
  `skillref_named.go` → empty pools on missing/unmatched taxonomy. A hallucinated name DROPS.
- the Seeder/Registry/DAG contract (`seeder/seeder.go` + `cmd/stackseed/main.go buildRegistry()`) takes
  a new `'generated-batch'` surface cleanly (DependsOn users+taxonomy, PerStackIsolated).
- supply-chain: `grep -rl anthropos-work/ai */go.mod` → empty (0-new-deps streak intact; M45 breaks it
  deliberately).

**Placement decision pending bootstrap tok:** overview's Delivers says `corpus/ops/cache-spec.md`; the
demo-family index convention puts demo docs under `corpus/ops/demo/`. Bootstrap tok resolves
placement (TOK-01).

## Gate-proving — iter-07 (the FIRST real LLM run)
**The env-name bug the real run surfaced.** `stack-demo/platform/.env`'s direct `OPENAI_KEY` is billing-dead
(HTTP 429 `insufficient_quota`). But the .env carries a LIVE Azure-OpenAI deployment the wrapper couldn't
reach: the platform's services name Azure config differently from the wrapper's original read
(`AZURE_OPENAI_ENDPOINT`, which is ABSENT). Authoritative platform names (verified in
`stack-demo/skiller/cmd/root.go` + `cms/studio`):
- `AZURE_OPENAI_KEY` + `AZURE_OPENAI_ENDPOINT_URL` → the **eastus2** deployment (skiller's read; the full URL
  is passed straight to `openai.NewAzure(key, url, nil)`).
- `AZURE_API_KEY` + `AZURE_ENDPOINT` → the **production Sweden (EU)** deployment (cms/studio's read).
- direct key is `OPENAI_KEY` (not `OPENAI_API_KEY`).

**The fix (iter-07):** `services/ai/ai.go` `NewFromEnv` now resolves the Azure pair from a prioritized,
**EU-first** list (`azureEnvPairs`: Sweden/EU `AZURE_API_KEY+AZURE_ENDPOINT` → eastus2
`AZURE_OPENAI_KEY+AZURE_OPENAI_ENDPOINT_URL` → legacy `AZURE_OPENAI_KEY+AZURE_OPENAI_ENDPOINT`), and reads
the direct key from `OPENAI_API_KEY` then `OPENAI_KEY`. Values-blind preserved. 4 new routing tests.

**Connectivity smoke-test FIRST (the stall-proofing).** Before the real batch, ONE completion fired via the
EU-first Azure routing with a SHORT timeout (gpt-4o-mini, 2.2s, valid JSON) — confirms a live deployment +
fails fast on a hung endpoint instead of stalling. The model/deployment proven + used for the whole gate:
**gpt-4o-mini via the EU-first Azure (Sweden) pair `AZURE_API_KEY`+`AZURE_ENDPOINT`.** (The watchdog stall in
the prior run-2 was a long silent Azure batch with NO per-call deadline — fixed below.)

**Three issues the real run surfaced (all fixed this iter):**
1. **No per-call timeout** — `cmd/gen-batch`'s run loop used `context.Background()` with no deadline (the
   stall class). FIX: a `--call-timeout` flag (default 60s) wraps each `CompleteJSON`; a hung endpoint fails
   FAST.
2. **Intra-batch name duplication** — gpt-4o-mini is strongly name-sticky per mother-prompt (a bare seed
   change re-picks the same name; the raw run gave ~9 distinct names / 20). FIX: (a) the system prompt now
   demands a varied/multicultural name + names the over-used Anglo defaults to avoid; (b) on a re-roll the
   user prompt carries an "avoid these already-used names" hint built from the batch's used-name set + the
   reserved heroes, so re-rolls actually DIVERGE. Result: **20/20 distinct names**. New unit test.
3. **`user_skills` CHECK violation (23514)** — the generated claimed-skill rows left every provenance edge
   NULL, violating `user_skills_check_foreign_keys` (≥1 of experience/education/… non-NULL). FIX: the
   `GeneratedBatchSeeder` now seeds **ONE company + ONE current-role `user_experiences` row per generated
   member** (reusing the ProfileSeeder helpers, FK-ordered COPY) and ties each claimed skill to that
   experience via `user_skill_experience`. Stays SHALLOW (1 exp = the current job; heroes keep the deep
   timelines) and makes the generated profile believably show a current role. A reproducible
   `stackseed --cache-root` flag points the seeder at the captured cache.

**The real gate measurement — ALL 5 DIMENSIONS PASS (EU-first Azure / Sweden, gpt-4o-mini, N=20):**
- **valid-JSON 100.0%** over 33 calls (pre-re-roll; the extra calls are the name-dedup re-rolls) — gate ≥95% **PASS**.
- **taxonomy-resolution + closure:** on the demo-3 re-seed (taxonomy replayed: 22,459 job_roles / 42,790
  skills / 72,685 role-skills), **47/47 claimed skills + 20/20 roles resolve to REAL `skiller.*.node_id`** (the
  resolvers SELECT the `node_id` column, e.g. `K-PYTHON-8B21` = "Python"); `datadna measure-closure --stack
  demo-3` = **`[PASS]`** (every seeded skill node-id resolves; **0 fabrication, closure GREEN**) — **PASS**.
- **hero-collisions 0** ('Maya Chen' never generated) — **PASS**.
- **cost $0.0059** of the $0.10 ceiling (5.9%) — **PASS**.
- **$0 byte-identical re-seed:** run 2 → **0 calls, $0.0000, 20 cache hits**, cache bytes unchanged — **PASS**.

**Believability proof (demo-3):** 20 generated members, **20/20 distinct multicultural names** (Aisling
O'Reilly, Amina Kone, Dmitri Petrov, Khalid Al-Farsi, Leandro Carvalho, Tariq Al-Mansouri, …), role-coherent
resolving skills (Backend Developer → Python/SQL/Node.js/Java; Frontend → JS/HTML/CSS), 20 current-role
experiences, **20/20 avatars**, **isolation CLEAN** (no shared/external writes — the firewall held). The
gate-met state is rext-tagged `method-acting-m45-iter07-gate` + bumped into the demo-3 consumption clone.
