# AI Architecture

This document describes the AI model inventory, provider routing, voice engine, recording architecture, and cost tracking across the Anthropos platform.

## High-Level Summary (For PMs & Non-Engineers)

Anthropos uses AI models from **multiple providers** to power its workplace simulations. AI actors in simulations can hold voice conversations, chat, analyze documents, and evaluate code — all powered by large language models. The platform routes AI requests through **EU providers first** for data residency compliance, falling back to US providers only when needed. Importantly, **simulation scoring is NOT done by AI** — it uses deterministic rubrics for EU AI Act compliance.

---

## AI Providers & Model Inventory

### Provider Routing Strategy

AI requests follow an **EU-first** routing policy:

1. **Azure OpenAI (EU)** — Primary for OpenAI models
2. **AWS Bedrock (EU)** — Primary for Anthropic models (e.g., `eu.anthropic.claude-*`)
3. **Mistral (EU)** — OCR and specialized tasks
4. **OpenAI Direct (US)** — Fallback
5. **Anthropic Direct (US)** — Fallback

EU data residency is non-negotiable. Customer data never leaves EU unless all EU providers are unavailable.

### Model Families

| Provider | Models | Routing |
|:---------|:-------|:--------|
| **OpenAI (Azure EU + Direct US)** | GPT-5.2, GPT-5.1, GPT-5, GPT-5-mini, GPT-5-nano, GPT-4.1, GPT-4.1-mini, O3, O4-mini | Azure EU primary |
| **Anthropic (Bedrock EU + Direct US)** | Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet | Bedrock EU primary |
| **Mistral (EU)** | Mistral OCR | EU only |
| **Speech** | GPT-4o Mini TTS, TTS v2 HD, TTS v2 | OpenAI |
| **Transcription** | GPT-4o Transcribe | OpenAI |
| **Embeddings** | Text Embedding 3 Small | OpenAI |

### Unified AI Library

All Go services access AI through the shared `ai` library, which provides:
- A single `ai.AI` interface across providers (OpenAI, Azure, Anthropic, Bedrock, Mistral)
- Per-provider client constructors that return provider token counts (`MetaData.Usage`)

> **EU-first routing/fallback and cost tracking are NOT in the `ai` library** — they live in the consuming services: routing/fallback in each service's `internal/ai/ai.go` (EU Azure default → US Azure via the PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429; Anthropic is always Bedrock `eu-west-1`), and cost tracking in `app/internal/aiusage/ai_usage.go` (fed by `Event_AiUsage` over Redis Streams). See [Shared Libraries → ai](shared_libraries.md#ai).

---

## AI Usage by Service

| Service | AI Use Case |
|:--------|:------------|
| **Jobsimulation** | Simulation conversations (chat + voice), document analysis, code evaluation |
| **Skiller** | Job role matching (embeddings + RAG), skill embeddings from 60K taxonomy (see [Vector storage](#vector-storage-in-skiller)) |
| **CMS** | Content generation, similarity matching, AI video (HeyGen), **and runs the full simulation generation pipeline** (Python studio-room embedded — see below) |
| **Studio-Desk** | Copilot AI assistant for content authoring (multi-provider chain: Azure OpenAI / OpenAI / Anthropic via `AI_PROVIDER_CHAIN`) |
| **Studio-Room** (Python) | Full simulation generation pipeline. **Runs as a subprocess inside the CMS container** (lives at `cms/studio/`, cloned from `anthropos-studio-room` via `cd cms && make init-studio`). |

### Studio-Room Generation Slots

The Python generation pipeline uses configurable model slots:

| Slot | Production | Experimental |
|:-----|:-----------|:-------------|
| FAST | gpt-4.1-mini | gpt-5-mini |
| STRICT | gpt-4.1 | gpt-5-mini |
| EXECUTION | gpt-4o | gpt-5.2 |
| CREATIVE | gpt-4o | gpt-5.2 |
| REASONING | — | gpt-5.2 |

### Embeddings & RAG (Skiller)

- **Model**: Text Embedding 3 Small (OpenAI), 1536-dim
- **Data**: Vectors for 60K skills and 18K job roles
- **Process**: RAG matches user input to taxonomy using OpenAI (Azure EU) or Anthropic (Bedrock EU)
- **Caching**: Redis for frequent matches

#### Vector storage in skiller

As of 2026-Q2 (migrations `20260417103036` and `20260417120309`), embeddings are stored in **dedicated tables**, not as columns on the entity tables:

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
2. **Route**: Selected model from CMS field (e.g., `gpt-5`, `gpt-4.1`, `anthropic-45-sonnet-aws`; default: `gpt-5` via Azure)
3. **Generate**: Per task type (voice/chat/code/document), AI generates responses or analysis
4. **Record**: LiveKit captures voice; AWS Chime captures video
5. **Score**: **Deterministic rubric scoring** (0-100 scale, NOT AI-scored) for EU AI Act compliance
6. **Insights**: AI generates post-session insights and feedback

### Evaluation System

Scoring is deliberately kept deterministic:
- Each skill has multiple criteria with binary checks (pass/fail)
- Rubric scores (0-100) map to competency levels (0-5)
- Thresholds: Level 1 ≥ 60, Level 2 ≥ 65, Level 3 ≥ 75, Level 4 ≥ 85, Level 5 ≥ 95
- This ensures the platform classifies as **Limited Risk** under the EU AI Act

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
- [Skiller Service](../services/skiller.md)
