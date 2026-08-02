# AI Architecture

This document describes the AI model inventory, provider routing, voice engine, recording architecture, and cost tracking across the Anthropos platform.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos uses AI models from **multiple providers** to power its workplace simulations. AI actors in simulations can hold voice conversations, chat, analyze documents, and evaluate code — all powered by large language models. Data stays in the EU because the **default provider clients are EU-resident** (Azure EU, AWS Bedrock `eu-west-1`) — **not** because requests walk an EU-first fallback ladder; there is no such ladder, and only a short, enumerable set of narrow levers can send a request to a US endpoint — none of them a region-health failover. See [Provider Routing Strategy](#provider-routing-strategy). **⚠️ The often-repeated claim that "simulation scoring is NOT done by AI" is false at platform HEAD** — the rubric *arithmetic* is deterministic, but **most** of the per-check pass/fail verdicts it aggregates come from an LLM. *Most*, not all — deterministic `EngineTextDiff` checks are the exception, and "all verdicts are AI" is the opposite error. See [Evaluation System](#evaluation-system).

---

## AI Providers & Model Inventory

### Provider Routing Strategy

> **⚠️ There is no ordered EU-first fallback chain.** The corpus published one for several releases
> ("Azure → Bedrock → Mistral → OpenAI → Anthropic"); **no such ladder exists in the code**, and
> **Mistral is not in the AI manager at all**.

What is actually implemented, in `app/internal/jobsimulation/ai/ai.go` (and mirrored in
`app/internal/skillerai/ai.go`):

1. **The caller picks the vendor.** `ChatCompletion` / `Response` take a `vendor AIVendor` argument and
   hand it to `getClient` (`:259-289`) — a plain `switch` over the four vendor consts at `:30-33`
   (`azure`, `openai`, `anthropic-aws`, `anthropic`). No ordering, no probing, no chain.
2. **`azure` resolves to the EU deployment by default**, swapped to the US one only when the PostHog
   flag `flag_use_azure_us` evaluates true; a *failed* flag lookup logs and **keeps the EU client**.
   That is a deliberate flag flip, not a health-based failover.
3. **`anthropic-aws` and `anthropic` both return the same Bedrock client**, constructed with
   `config.WithRegion("eu-west-1")`. There is **no** US-direct Anthropic branch in this manager.
4. **The one automatic fallback is 429-only.** `isThrottlingError` sets `vendor = Openai` for the next
   retry attempt — i.e. **direct US OpenAI**. Nothing else (timeout, 5xx, region outage) moves a
   request off its vendor.
5. **Mistral is nowhere in this path.** Its only use in `app` is OCR inside the cms domain
   (`internal/cms/studio/markdownManager.go:19`, `studioManager.go:583`).

**EU data residency still holds — but by a different mechanism than the ladder implied.** The posture
rests on the *default clients* being EU-resident (Azure EU, Bedrock `eu-west-1`), not on a policy that
walks EU options before US ones. Within the manager exactly three things can send a request outside the
EU, **none of them a region-health failover**: the `flag_use_azure_us` flag, the 429 retry (which goes
straight to direct OpenAI **without** trying another EU provider first), and setting `ANTHROPIC_API_KEY`,
which flips Course Builder / Studio-Room off Bedrock onto Anthropic's first-party API.

One layer *above* the manager there is a fourth exit: an unrecognised `ai_vendor` string on a sequence
falls through to `internalAi.Openai` — **direct US OpenAI** — in
`app/internal/jobsimulation/simulator/ai/ai.go:114-115`. See
[The three simulation model defaults](#the-three-simulation-model-defaults).

For the full per-line derivation see
[External Services → Routing: what is actually implemented](external_services.md#routing-what-is-actually-implemented)
— it is the canonical statement; do not restate it here.

### Model Families

The **Default client** column is the client the vendor const resolves to — *not* a position in a
fallback order (there is none; see above).

| Provider | Models | Default client |
|:---------|:-------|:---------------|
| **OpenAI (Azure EU + Direct US)** | GPT-5.4, GPT-5.4-mini, GPT-5.2, GPT-5.1, GPT-5, GPT-5-mini, GPT-5-nano, GPT-4.1, GPT-4.1-mini, O3, O4-mini | Azure **EU** (US only via `flag_use_azure_us`; direct US OpenAI only on a 429 retry) |
| **Anthropic (Bedrock EU + Direct US)** | Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` (Course Builder / Studio-Room) |
| **Mistral (EU)** | Mistral OCR | cms-domain OCR **only** — not reachable from the AI manager |
| **Speech** | GPT-4o Mini TTS, TTS v2 HD, TTS v2 | Azure voice client (`CreateSpeech` is Azure-only) |
| **Transcription** | GPT-4o Transcribe | Azure EU (US via `flag_use_azure_us`) |
| **Embeddings** | Text Embedding 3 Small | OpenAI |

### Unified AI Library

All Go services access AI through the shared `ai` library, which provides:
- A single `ai.AI` interface across providers (OpenAI, Azure, Anthropic, Bedrock, Mistral)
- Per-provider client constructors that return provider token counts (`MetaData.Usage`)

> **Vendor selection/fallback and cost tracking are NOT in the `ai` library** — they live in the consuming services: selection/fallback in each consumer's own wrapper — in `app` @ `5ba17044` that is **`app/internal/jobsimulation/ai/ai.go:267,344`** and **`app/internal/skillerai/ai.go:347`**, *not* a bare `app/internal/ai/ai.go` (**no such file** — that path now resolves only inside the frozen `jobsimulation` husk repo). EU Azure default → US Azure via the PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429; Anthropic is always Bedrock `eu-west-1`, and cost tracking in `app/internal/aiusage/ai_usage.go` (fed by `Event_AiUsage` over Redis Streams). See [Shared Libraries → ai](shared_libraries.md#ai).

---

## AI Usage by Service

| Service | AI Use Case |
|:--------|:------------|
| **jobsimulation domain** *(in `app`)* | Simulation conversations (chat + voice), document analysis, code evaluation |
| **Backend (`app`)** | Job role matching (embeddings + RAG), skill embeddings over the taxonomy (**≥42,790 skills** — the measured *public* subset; see [the "60K / 18K" figures](shared_libraries.md#taxonomy-figures)) — the merged skiller domain, July 2026 (see [Vector storage](#vector-storage-merged-skiller-domain)); plus Talk to Data (Bedrock) |
| **cms domain** *(in `app`)* | Content generation, similarity matching, AI video (HeyGen), **and runs the full simulation generation pipeline** (Python studio-room embedded — see below) |
| **Studio-Desk** | Copilot AI assistant for content authoring (multi-provider chain: Azure OpenAI / OpenAI / Anthropic via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** (Python) | Full simulation generation pipeline. **Runs as a subprocess inside the `app` (backend) container** since cms-in-app — the `anthropos-studio-room` project is pulled into the image by CI. |

### Studio-Room Generation Slots

The Python generation pipeline uses configurable model slots, one per `{MODE}_AI_{BRANCH}_MODEL` key of
`configs/{env}_config.ini` (`service, model, thinking`). Measured at `app` HEAD against the shipping
`studio/configs/production_config.ini:26-36` — **the `stable` and `experimental` branches are currently
identical**, and `development_config.ini:26-36` is identical to both:

| Slot | Stable (`production_config.ini`) | Experimental | Thinking |
|:-----|:---------------------------------|:-------------|:---------|
| FAST | azure / gpt-5-mini | azure / gpt-5-mini | none |
| STRICT | azure / gpt-5-mini | azure / gpt-5-mini | none |
| EXECUTION | azure / gpt-5.4 | azure / gpt-5.4 | none |
| CREATIVE | azure / gpt-5.4 | azure / gpt-5.4 | low |
| REASONING | azure / gpt-5.4 | azure / gpt-5.4 | medium |

> **Do not read the slot table off `configs/config_template.ini`.** That file is a non-shipping scaffold
> and still carries the older `gpt-4.1-mini` / `gpt-4.1` / `gpt-4o` / `o3` stable column; the corpus
> asserted a hybrid of it for several releases. **`gpt-5.2` appears in no studio config at all** (only as a
> pricing entry in `studio/services/ai.py:356,508`), and `gpt-4o` appears in no `*_MODEL` slot of any
> studio config. Agrees with [`studio-room.md`](../services/studio-room.md#ai-service-configuration).

### Embeddings & RAG (Backend `app` — merged skiller domain)

- **Model**: Text Embedding 3 Small (OpenAI), 1536-dim
- **Data**: `public.skill_embeddings` = **42,790** rows and `public.job_role_embeddings` = **18,919** rows over the *public* taxonomy (`organization_id IS NULL`, measured 2026-06-29)
- **Process**: RAG matches user input to taxonomy using OpenAI (Azure EU) or Anthropic (Bedrock EU)
- **Caching**: Redis for frequent matches

> **⚠️ Do not write "60K skills and 18K job roles" here.** They are not measurements and they fail in
> two *different* ways: **"18K roles" is REFUTED** — the public subset alone is **22,470** job roles and
> public ⊆ total, so prod holds **≥22,470**; 18,919 is the *job-role-embedding* row count, transcribed
> onto the role count. **"60K skills" is UNVERIFIED, not refuted** — a public-only capture cannot see
> org-private skills, so **42,790 is a floor, never "the total"**. Canonical statement:
> [Shared Libraries → the "60K / 18K" figures](shared_libraries.md#taxonomy-figures).

#### Vector storage (merged skiller domain)

As of 2026-Q2 (skiller migrations `20260417103036` and `20260417120309`), embeddings are stored in **dedicated tables**, not as columns on the entity tables. Since the skiller→app merge (July 2026) these tables live in the **`public` schema** owned by `app` (ported from the legacy `skiller` schema):

```
job_role_embeddings(
  id BIGSERIAL PK,
  job_role_id UUID FK → job_roles.id,
  small_embedding3 extensions.vector(1536),
  -- IVFFLAT index on small_embedding3
)

skill_embeddings(
  id BIGSERIAL PK,
  skill_id UUID FK → skills.id,
  small_embedding3 extensions.vector(1536),
  -- IVFFLAT index on small_embedding3
)
```

The previous denormalized `small_embedding3` columns on `job_roles` and `skills` were dropped in the same migration. The `extensions` schema (which houses the `pgvector` extension) must exist before applying these migrations — this is handled in `corpus/ops/setup_guide.md`.

---

## Voice Architecture

### Active Engine: LiveKit + GPT Realtime

The primary voice engine uses **LiveKit rooms** with **OpenAI GPT Realtime** agents:

```
Player → LiveKit Room → GPT Realtime Agent (anthropos-agent-eu / anthropos-agent-us)
```

> **The agents have repos, and this corpus has never named one (v2.8 M257x).** The org holds **five**
> LiveKit agent repositories — `livekit-agent`, `livekit-agent-chain`, `livekit-agent-azure-us`,
> `livekit-agent-azure-eu`, `livekit-agent-azure-eu-fr` — none of which appear in `repos.yml`, in any
> corpus document, or in the deployment picture below. Everything here documents the LiveKit *engine* and
> the *platform side* of the call; the agent process itself is undocumented. `azure-eu` and `azure-eu-fr`
> were measured at M257x iter-01 as dispatching nothing. Enumerated in
> [`platform-migration-status.md` §3](./platform-migration-status.md).

- **Audio**: Real-time voice conversation, recorded as MP3
- **Transcript**: Generated from LiveKit conversation events
- **Configuration**: Voice engine is selectable per simulation in CMS (`livekitgptrealtime`)

### Legacy / Transitioning Engines

| Engine | Status | Description |
|:-------|:-------|:------------|
| `elevenlabs` | Active (legacy default) | ElevenLabs conversational agents; still used by the call/reply pipeline (`getJobSimulationCallSignedUrl` / `getJobSimulationCallConversationToken`) and transcript improvement |
| `gptrealtime` | Deprecated | Direct OpenAI Realtime without LiveKit |

LiveKit + OpenAI Realtime is the engine for **new** sessions (gated by the `flag_use_realtime_openai` PostHog flag); **ElevenLabs remains the active default** for the call/reply pipeline and transcript improvement, so it is not yet fully replaced.

---

## Recording Architecture

Two parallel recording systems capture simulation sessions:

| System | Captures | Format | Purpose |
|:-------|:---------|:-------|:--------|
| **LiveKit** | Voice only | MP3 | Audio transcript and voice recording |
| **AWS Chime SDK** | Camera + screensharing + mic | Composited MP4 (grid view) | Full video record of simulation |

Both recordings are stored in S3 and linked to the simulation session.

---

## Simulation AI Flow

1. **Load**: Simulation definition fetched from CMS (actors, tasks, rubrics, AI model config per sequence)
2. **Route**: Selected model from the CMS `ai_model` / `ai_vendor` fields per sequence (e.g. `gpt-5`, `gpt-4.1`, `anthropic-45-sonnet-aws` — the enum is `app/internal/cms/directus/collections/jobsimulation.go:983-990`). **There is no single "default model", and `gpt-5` is not a default anywhere.** Three distinct defaults apply at three different points — see below
3. **Generate**: Per task type (voice/chat/code/document), AI generates responses or analysis
4. **Record**: LiveKit captures voice; AWS Chime captures video
5. **Score**: deterministic rubric *arithmetic* (0-100 scale) over **per-check verdicts that are mostly LLM-produced** (the `EngineTextDiff` minority is deterministic) — see [Evaluation System](#evaluation-system); this is NOT an AI-free scoring path
6. **Insights**: AI generates post-session insights and feedback

### The three simulation model defaults

The corpus asserted a single *"default: `gpt-5` via Azure"* for several releases. **`gpt-5` is not a default
in any of the three places a default is applied** — measured at `app` HEAD:

| Where the default fires | Vendor | Model | Source |
|:------------------------|:-------|:------|:-------|
| **Content side** — the CMS `ai_model` / `ai_vendor` field is left unset on a sequence | `openai` | **`gpt-5.1`** | `app/internal/cms/directus/collections/jobsimulation.go:1297` (model), `:1302` (vendor) |
| **Runtime routing** — `GetAIVendorAndModel` gets a model string it does not recognise | as selected | **`gpt-4.1`** | `app/internal/jobsimulation/simulator/ai/ai.go:65-66` (OpenAI arm), `:82-83` (Azure arm), `:126-127` (unmatched-vendor arm) |
| **Scoring / validation** — not configurable at all, hardcoded | **Azure** | **`gpt-4.1`** (summarize: `gpt-4.1-mini`) | `simulator/ai/ai.go:20-26`, `GetValidationAIVendorAndModel` |

Two further details the old one-liner erased:

- **The unmatched-*vendor* fallback is direct OpenAI, not Azure** — `simulator/ai/ai.go:114-115` sets
  `aiVendor = internalAi.Openai`, and `Openai` and `Azure` are distinct clients. An unrecognised vendor
  string therefore leaves the EU-resident default (Azure EU) for a US endpoint, so it is a
  data-residency-relevant fallback, not a cosmetic one
- **The `gpt-4.1` model default only holds inside the OpenAI/Azure/unmatched arms.** The Anthropic arms have
  their own: `anthropic-aws` falls back to Claude 3.7 Sonnet on Bedrock (`:98-99`) and `anthropic` to
  Claude 3.5 Sonnet (`:110-111`)

### Evaluation System

**The ARITHMETIC is deterministic. The inputs to it are not** — measured at `app` HEAD, M257x iter-38:

- Each skill has multiple criteria with binary checks (pass/fail), and **most of those verdicts are judged
  by an LLM**. The dispatch is a hardcoded switch, not the `checkerEngines` map — that map is stored and
  **never read** (`internal/jobsimulation/simulator/validation/v3/validator/validator.go:43,60-61,595`),
  so do not cite it as the mechanism. `basevalidator/criterion.go:127` routes LLM checks to `validateLLM`
  → `NewLLMBulkChecker(c.logger)` (`:428`), which sends `basevalidator/templates/checkValidationBulk.tmpl`
  at temperature 0.0 and reads back `{"check_id", "feedback", "success"}`
- **The exception, stated because "all verdicts are AI" would be the opposite error:** `EngineTextDiff`
  checks are evaluated deterministically — `criterion.go:168` runs `validateCodeDiff` concurrently with the
  LLM pass and `:450-475` sets `success` from a pure string comparison. A code simulation therefore mixes
  deterministic and LLM-judged checks in one score
- `calculateSkillScore` (`v3/validator/skills.go:53-64`) then counts those booleans and `:75` computes
  `passed / total * 100` — deterministic, over predominantly AI-produced atoms
- **There is no 60/65/75/85/95 threshold ladder.** The corpus asserted one for several releases; it does
  not exist in `app`, `cms`, `jobsimulation` or `next-web-app`. The real conversion is
  `calculateCompetencyLevelScore` (`v3/validator/skills.go:40-51`): `20` when `score < 60 && isPassed`,
  `100` at `>= 100`, else `max(0, score*2-100)` — and it carries a `// TODO fix this formula` comment.
  The 0-100 ↔ N-level mapping is a plain division (`app/internal/skill/skill.go:617-623`
  `convertLevelTo100`; frontend `packages/ui/src/Competency/CompetencyReadLevel.tsx:18`)
- **⚠️ Therefore this section does NOT support a Limited-Risk classification** under the EU AI Act. See
  [Security & Compliance → EU AI Act](./security_compliance.md#eu-ai-act).

---

## Cost Tracking

Cost is tracked centrally in the backend `app` service (`internal/aiusage/ai_usage.go`), fed by `Event_AiUsage` messages that the AI-consuming services publish over Redis Streams (the shared `ai` library itself only returns provider token counts):

- **Tokens**: Input and output token counts per request
- **Latency**: Request duration per model
- **Cost**: Estimated cost per model per request
- **Aggregation**: Available per service, per model, and per time period

---

## Related Documentation
- [Architecture Overview](./architecture_overview.md)
- [Security & Compliance](./security_compliance.md)
- [External Services](./external_services.md)
- [Jobsimulation Service](../services/jobsimulation.md)
- [Backend Service](../services/backend.md) — owns the embeddings/matching domain (former [skiller](../services/skiller.md), merged July 2026)
