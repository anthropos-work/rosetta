# AI Generation Engine — gen-acceptance protocol + the engine (v1.10 "method acting" M45)

> **The scalable-generation extension of the demo family.** Where the `PersonaSeeder` /
> `ProfileSeeder` / `Certificates`/`ProjectsSeeder` hand-author a small roster of curated heroes
> (`stories-spec.md`, `profile-completeness-spec.md`), the **generation engine** turns a high-level YAML
> **batch descriptor** into realistic **per-member profiles** at scale via a **cheap LLM**
> (`gpt-4o-mini`), behind a **prompt-hash cache** so reruns cost **$0** and reseed byte-identical. It is
> the **FIRST new third-party dependency** in the seeding module (the shared `ai` library) — a deliberate,
> in-release supply-chain decision. **Entirely in `rosetta-extensions/stack-seeding/` — zero platform-repo
> edits.**
>
> This doc is BOTH the **engine reference** AND the milestone's **iteration protocol** (the
> measure→fix→accept loop on generation quality). The cache mechanics are split into a companion
> [`cache-spec.md`](cache-spec.md).

---

## 1. The CODE-owns-structure / AI-owns-content boundary (the design's spine)

The whole engine rests on one line: **CODE owns structure, identity, and the closure gate; the AI owns
ONLY a per-member English content envelope.** The LLM **never** writes a DB row, **never** picks a
node-id, **never** decides identity.

| Concern | Owner | Why |
|---|---|---|
| The 7-table fan-out (users, personas, profile, work-history, education, …) | **CODE** | deterministic; the existing seeder machinery, unchanged |
| Avatars, emails, deterministic user IDs | **CODE** | identity is derived, not invented |
| Role / skill **node-id** resolution | **CODE** | the resolvers (`resolveJobRoleRefs` / `resolveNamedSkillRefs`) map a NAME → a **real public node-id** |
| Caching, cost/throttle, the `--max-cost` ceiling | **CODE** | reproducibility + the budget guard |
| The per-member JSON envelope (name, bio, education, roles, skill **names**, self-eval bias, location) | **AI** | plausible English content only |

**The seam (why a cheap model is safe).** The AI emits skill/role **names** as plain English strings.
Those names feed the **SAME resolvers the hand-curated heroes use**. A name that resolves → a real
node-id. A name that **doesn't** resolve (a hallucination, a typo, a too-niche skill) → it **DROPS**. The
worst case of a bad generation is therefore a **SHALLOWER profile** (fewer verified skills), **never an
invalid one**. The taxonomy **closure gene** (`datadna measure-closure`) stays **green** because every
node-id written was resolved from the real replayed public taxonomy — nothing is fabricated.

> This is the load-bearing invariant. Every tik measures it: **0 fabricated node-ids, closure green.**

---

## 2. Engine components (all in `stack-seeding/`)

### 2a. `services/ai/` — the thin wrapper (the new dep)

A local wrapper around the shared **`ai` library** (`github.com/anthropos-work/ai`, pinned `v1.40.1` — the
same pin the platform consumers use). The wrapper is a **thin consumer**, mirroring the platform's own
`internal/ai/ai.go` pattern (`shared_libraries.md` §ai + `ai_architecture.md`): it owns
**EU-first routing → 429 fallback → usage tracking**; it does **NOT** reimplement provider internals (the
`ai` lib owns transport).

- **Routing:** Azure-OpenAI EU first (`openai.NewAzure(apiKey, baseURL, apiVersion)`) when
  `AZURE_OPENAI_*` env is present; else direct OpenAI (`openai.New(apiKey)`) on `OPENAI_API_KEY`. On
  Azure HTTP 429 → fall back to direct OpenAI (the documented EU→direct fallback). gpt-4o-mini is the
  default model (`WithModel`), JSON mode forced (`WithResponseFormat(JSONObject)`), low temperature +
  fixed `WithSeed` for reproducibility.
- **Usage tracking:** every call returns `*MetaData` whose `Usage` is a `*openai.CompletionUsage`
  (`PromptTokens` / `CompletionTokens` / `TotalTokens`, all `int64`). The wrapper folds those into a
  running per-batch cost via a **model→price table** (mirroring `app/internal/aiusage`'s model→price
  switch — the lib itself does NOT compute dollars). It emits an `Event_AiUsage`-shaped cost record per
  call + a per-batch cost report.
- **Values-blind:** the key is read from the environment (`OPENAI_API_KEY` / `AZURE_OPENAI_KEY`); it is
  never logged, echoed, or written to a cache file.

### 2b. `blueprint.Batch` + `batch[]` + `EffectiveBatches()`

A new `Batch` type + a `batch[]` field on `Story` / `StackSeed`, normalized through **`EffectiveBatches()`**
— the exact mirror of the existing `EffectiveStories()` normalization layer. A batch descriptor is
high-level (count, a role/seniority mix, an industry/narrative frame, a self-eval-bias distribution); its
**`EffectiveBatches()` expansion is pure Go-template** per-member-prompt rendering — **NO LLM at parse
time.** The LLM runs **only** in `cmd/gen-batch`. Parsing a blueprint with a `batch[]` stays offline,
deterministic, and free.

The expansion produces, per member, a **MOTHER prompt** (the deterministic per-member prompt string) — the
cache key (§ [`cache-spec.md`](cache-spec.md)).

### 2c. `cmd/gen-batch` — the generation CLI

- **Default model** `gpt-4o-mini` (cheap); overridable.
- **`--max-concurrent`** semaphore (default 5) — bounded parallel generation.
- **`--max-cost`** — a **MANDATORY** dollar ceiling (the user's hard requirement). The run estimates +
  tracks running cost and **aborts before** a call would breach the ceiling. No batch ever runs uncapped.
- **Re-roll on malformed:** a call whose response is not valid JSON (or fails envelope-schema validation)
  is re-rolled up to a bounded retry count; persistent malformed → that member is dropped from the cache
  (never written half-formed).
- **Atomic cache writes** (`.tmp`→rename) + a **`.lock` fence** (§ `cache-spec.md`).
- Emits the **per-batch cost report** (calls, cache hits, tokens, dollars vs ceiling) + the per-member
  `Event_AiUsage` cost records.

### 2d. `GeneratedBatchSeeder` (surface `'generated-batch'`)

A new Seeder: `Surface() == "generated-batch"`, `DependsOn()` = users + taxonomy, `Isolation()` =
`PerStackIsolated`. It reads the **cached** per-member JSON and drives the users / persona / profile rows
**DETERMINISTICALLY** through the existing resolvers + seeder helpers — the same 7-table fan-out the
curated heroes use, just fed from cached AI content instead of hand-authored YAML. A non-resolving name
DROPS (§1). It writes **no** new content; it is a deterministic transform of `cache → DB rows`.

### 2e. Hero-name collision avoidance

The exit gate requires **ZERO** generated name colliding with a hand-curated hero. Two guards:
1. **Prompt-side:** the per-member prompt carries the **reserved-hero-names** roster (the curated heroes
   of the active stories) as an explicit "do not use these names" instruction.
2. **Code-side fallback:** a **post-gen collision check** re-rolls any generated name that still collides
   (case-insensitive, full-name). A name that can't be de-collided after the bounded re-roll budget drops
   the member (never writes a colliding identity).

---

## 3. The exit gate (verbatim from the milestone overview)

On a **real batch of N generated members**, the engine seeds end-to-end:
- the cheap model emits **valid JSON ≥95%** of calls (**re-roll** on malformed);
- **every** role/skill name resolves to a **REAL public-taxonomy node-id** via the existing resolvers
  (non-resolving names **drop**; the closure gene stays **green**);
- **ZERO** generated name collides with a hand-curated hero;
- total cost lands **within the declared `--max-cost` ceiling**.

**Reproducible:** an unchanged batch descriptor **re-seeds byte-identical from cache at $0.**

---

## 4. The iteration protocol (the measure → fix → accept loop)

M45 is **iterative** because the empirical heart — *"does `gpt-4o-mini` reliably emit valid JSON whose
skill names resolve?"* — is answerable **only by real runs**. Each tik measures, then hardens the prompt
or the code, then re-runs.

### 4a. The primary metric

The **gen-quality vector**, measured on a real batch of N:

| Sub-metric | Gate threshold | Source |
|---|---|---|
| **valid-JSON rate** | **≥ 95%** (re-roll on malformed; the rate is pre-re-roll, the honest model-quality number) | `cmd/gen-batch` run report |
| **taxonomy-resolution rate** | every name resolves-or-drops; **closure green**; **0 fabricated node-ids** | `GeneratedBatchSeeder` + `datadna measure-closure` |
| **hero-collision count** | **= 0** | post-gen collision check |
| **cost vs `--max-cost`** | **within** ceiling | per-batch cost report |
| **$0 re-seed reproducibility** | unchanged descriptor → **byte-identical** from cache, **$0** | second run diff |

The **gate is met** when ALL five hold on a real batch, reproducibly.

### 4b. Per-iter shape

Each tik:
1. **Probe / measure.** Run the engine on a fixed batch — first with **FIXTURE LLM responses**
   (deterministic, no key, no cost) for the unit-level sub-metrics, then (when the iter's hypothesis
   needs it) a **REAL** small batch (e.g. N=20, `--max-cost` ~$0.10) for the empirical valid-JSON /
   resolution rate. Record all five sub-metrics pre-iter.
2. **Diagnose.** Which sub-metric is short? Classify the failure (malformed-JSON shape, a skill-name
   family the resolver keeps dropping, a collision pattern, a cost overrun, a non-determinism in the
   cache).
3. **Harden.** Fix the prompt (envelope schema clarity, name constraints, the reserved-hero list) OR the
   code (re-roll logic, the resolver call path, the cache key, the cost guard).
4. **Re-measure.** Re-run; record the post-iter five-metric vector + the delta.
5. **Close** the iter (`closed-fixed` / `closed-fixed-partial` / `closed-no-lift`).

### 4c. Fixtures-first discipline

The engine must be **fully unit-testable without a key**: a fixture `ai.AI` (a fake implementing
`ChatCompletion`) returns canned envelopes (valid, malformed, colliding, non-resolving-skill) so every
code path — re-roll, drop, collision re-roll, cache write/read, cost accounting — is deterministic in CI.
The **real LLM batch** is run only to measure the empirical valid-JSON/resolution **rate** the gate
asserts (the one thing fixtures can't tell you). Build + unit-test against fixtures **first**; prove the
gate with a real capped batch **last**.

### 4d. Measurement conventions

- **valid-JSON rate is pre-re-roll** — it's the model-quality number, not the after-retry number (which
  is ~100% by construction). The gate's 95% is about the cheap model's raw reliability.
- **Resolution is drop-safe by design** — a low resolution rate is NOT a gate failure on its own (it
  yields shallower profiles); the gate failure is a **fabricated** node-id or a **broken closure**. Track
  the resolution rate as a quality signal, but the hard assert is *0 fabrications + closure green*.
- **Cost is always bounded** — even a passing run reports cost; the gate is "within ceiling", and the
  ceiling is mandatory on every invocation.
- **Reproducibility is a second run** — the $0 re-seed is proven by running the same descriptor twice and
  diffing the cache + the resulting rows byte-for-byte.

### 4e. Re-scope trigger

If the cheap model **cannot reach** the valid-JSON / resolution threshold after **~5 tiks** of prompt +
code hardening → escalate to a **user-strategic-replan** (model upgrade to a stronger/pricier model vs
scope reduction). The closure gate means the worst case is always a shallower-but-valid profile, so this
is a quality/cost trade-off decision for the user, never a correctness emergency.

### 4f. The key + cost (gate-proving)

`OPENAI_KEY` lives in `stack-demo/platform/.env` (read-only; **values-blind** — resolved into
`OPENAI_API_KEY` for `cmd/gen-batch`, never echoed). `gpt-4o-mini` makes a real N=20 batch cost
~**$0.005** — negligible, but **always** `--max-cost`-capped. The gate is proven on a real capped batch,
then re-seeded on a fresh demo via the tagged consumption clone.

---

## 4g. Org-scale lessons (v1.10 M46 — surfaced by the real ~600-member gate-proving)

Two failure modes are **structurally invisible** below ~2 batches / ~hundreds of members, so they only
surface when the engine is proven at org scale (M45's bounded N=20 cannot reach them):

- **The cache index is GLOBAL, not batch-local.** `cmd/gen-batch` selects/reads the cache by the member's
  position in the WHOLE `EffectiveBatches()` slice (`cache.Has(i)`/`Get(i)`), so the write MUST use the same
  global index — NOT the batch-local `Batch`-relative index (0..Count-1). With a single batch the two
  coincide; with **multiple story batches** a later batch's local index collides with an earlier batch's,
  so a local-index write OVERWRITES an earlier member's cached file and leaves the later slot empty — losing
  an entire later story's generated population. A **multi-batch** regression test is mandatory (a single-batch
  test can't catch it).
- **A cheap name-attractor model needs a deterministic disambiguator, not just LLM re-rolls.** gpt-4o-mini
  re-picks a small set of "distinctive" names hundreds of times across a 600-member org (the avoid-names
  re-roll hint can't scale to hundreds of taken names). The engine's last-attempt path therefore
  **deterministically disambiguates** a still-duplicate name (keep the generated first name + swap in a
  distinct surname keyed on the global index — cost-free, no extra LLM call, reproducible) so **name
  distinctness is guaranteed at any scale**, plus a larger avoid-hint (120) for more LLM-native divergence
  first. Accepting a duplicate on the last attempt (the pre-M46 behavior) destroys org-scale believability.
- **Distinctness MUST be enforced at SEED time too — the gen-time disambiguator only fires on a cache MISS
  (M46 iter-07).** The gen-time disambiguator above runs inside `genOne`, i.e. only when a member is being
  generated (a cache miss). A **`$0` cache-hit reseed** reads the cached envelope's `name`/`email_local`
  **verbatim** — so an EXISTING cache that pre-dates the disambiguator (or simply carries the model's raw
  attractor duplicates) would seed **duplicate identities**. The `GeneratedBatchSeeder` therefore applies the
  **same deterministic disambiguator at seed time** as a hard backstop, over **two distinctness axes**:
  - **name** — a cached name already taken (by a curated hero or an earlier generated member) is rewritten by
    `seeders.DisambiguateGeneratedName` (the **single source of truth** the gen-time path now also calls — one
    surname pool, one algorithm, so gen-time and seed-time AGREE on the same surname for the same global index);
  - **email** — name-distinctness alone does NOT imply email-distinctness: the cache carries duplicate
    `email_local`s AND two *distinct* names can derive the *same* local part ("Jinwoo Park" / "Jin-woo Park" →
    `jinwoo.park`). `public.users` enforces `UNIQUE(email)` (`user_basic_info_email_key`), so a colliding email
    aborts the **entire** generated-batch users COPY. The seeder indexes a colliding local part
    (`local+globalIdx@domain`, per-domain so two orgs never cross-collide).
  Both axes are deterministic + `$0` (no LLM) + reproducible (same cache → same distinct identities on every
  reseed → a FRESH `/demo-up` reproduces the same org). This is what makes the **`$0` cache-hit reseed**
  genuinely believable rather than a wall of duplicate names.

**Empirically proven on the real ~600-member batch (M46):** 0 hero-collisions at scale, 100% valid-JSON,
$0 cache-hit reseed (a complete 614-member cache reseeds **614/614 distinct names + 614/614 distinct emails**
at $0 via the seed-time backstop), and the **mandatory `--max-cost` guard correctly aborting at its ceiling**
(the re-roll/dedup overhead at scale is real, so a large org-fill is run with a generous cap or finished
across capped runs — the already-cached members reseed at $0, so finishing is cheap). **Population-math note
(M46 iter-07):** because the curated `UsersSeeder` ALSO seeds a full `size` synthetic body, an org with a
`fill: true` batch lands at ~2×`size` (heroes + a full curated body + a full generated fill), so the
gate-proving descriptor keeps `size` at 250 (Cervato) / 120 (Solvantis) to land a **believable ~500-member
headline org** (Cervato ≈ 498: ~250 curated + ~247 generated; Solvantis ≈ 237) rather than the ~1k a naive
`size: 500` would produce. The seeded population is believable (real distinct names/photos, role-coherent
skills, closure GREEN, 0 hero-collisions) and passes the employee-vantage M42 sweep + the manager persona /
cross-port checks.

> **Known platform limit at org scale (M46 iter-07 — re-scope-trigger).** The manager M42 sweep on this
> ~500-member org still fails 3 sections — `/enterprise/members`, `/enterprise/activity-dashboard`,
> `/enterprise/settings` — because the **federated GraphQL queries backing those enterprise grids never resolve
> in the harness window** (the Cosmo router logged 10–84 s latencies; the `organizationMembers` /
> activity-aggregation per-row resolver fan-out — `jobRole`/`targetRole`/`tags`/`lastActivityDate`/
> `organizationFeatures` × Sentinel authz — is an N+1 across subgraphs). This is **invariant to org size**
> (10.88 s @ 998 ≈ 10.5 s @ 500), so neither the resize nor a bounded warm-grid poll closes it; the manager
> gate last PASSED at ~221 members. The fix is a platform resolver change (forbidden by the zero-canonical-edit
> line) — **NOT** a seeding or harness change — so the M46 gate's "M42 sweep PASSES on a ~500 org" criterion is
> re-scoped: treat the enterprise members/activity grids as a documented org-scale platform-perf exception.

## 5. What's OUT (M45 scope boundary)

- **Org-scale auto-fill** to reach full org size → **M46** (M45 proves the engine + cache on a **bounded**
  batch).
- **Deep per-generated-member work-history/education timelines** → kept shallow (name + skills + bio +
  role); the curated heroes keep the deep timelines.
- **A platform-repo secrets store** → key via env var for now; production-seeding secrets deferred.
- **Any platform-repo edit** → the engine is entirely in `rosetta-extensions`.

---

## See also
- [`cache-spec.md`](cache-spec.md) — the prompt-hash cache (key, capture-version invalidation, atomic
  writes, the `.lock` fence).
- [`stories-spec.md`](stories-spec.md) — the curated verified-skill chain + the resolvers the generated
  names feed through.
- [`profile-completeness-spec.md`](profile-completeness-spec.md) — the M44 roster-completeness surfaces
  the generated members extend.
- [`seeding-spec.md`](../seeding-spec.md) — the seeding blueprint + the production-isolation boundary the
  `GeneratedBatchSeeder` lives within.
- [`shared_libraries.md`](../../architecture/shared_libraries.md) §ai +
  [`ai_architecture.md`](../../architecture/ai_architecture.md) — the shared `ai` library the `services/ai/`
  wrapper layers EU-first routing + usage tracking on.
