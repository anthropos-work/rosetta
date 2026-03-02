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
- Unified interface across providers (OpenAI, Anthropic, Mistral, Azure)
- Automatic EU-first routing
- Centralized token tracking (input + output tokens, latency, cost per model)
- Provider fallback chain

---

## AI Usage by Service

| Service | AI Use Case |
|:--------|:------------|
| **Jobsimulation** | Simulation conversations (chat + voice), document analysis, code evaluation |
| **Skiller** | Job role matching (embeddings + RAG), skill embeddings from 60K taxonomy |
| **CMS** | Content generation, similarity matching, AI video (HeyGen) |
| **Studio-Desk** | Copilot AI assistant for content authoring |
| **Studio-Room** (Python) | Full simulation generation pipeline |

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

- **Model**: Text Embedding 3 Small (OpenAI)
- **Data**: Vectors for 60K skills and 18K job roles
- **Process**: RAG matches user input to taxonomy using OpenAI (Azure EU) or Anthropic (Bedrock EU)
- **Caching**: Redis for frequent matches

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

### Deprecated Engines

| Engine | Status | Description |
|:-------|:-------|:------------|
| `elevenlabs` | Deprecated | Standalone ElevenLabs; template duplicated per session |
| `gptrealtime` | Deprecated | Direct OpenAI Realtime without LiveKit |

LiveKit replaced ElevenLabs as the primary voice engine for better reliability and cost efficiency.

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

The shared `ai` Go library provides centralized cost tracking across all microservices:

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
