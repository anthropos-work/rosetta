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
5. **Mistral is nowhere in this path.** *Every* use of it in `app` is **OCR**, never generation: Go-side
   in the cms domain (`internal/cms/studio/markdownManager.go:30` — the constructor body
   `return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil` inside `func NewMarkdownManager` at `:29`,
   **@ `app` `ad9f3c49`**; this row cited an earlier line in that file until M257x iter-115 — a **doc-comment**
   line — *"It used to take aiKey and then IGNORE it"* — not code, exactly as
   the *AI Providers* section of [`external_services.md`](external_services.md) already said while this line went on asserting it —
   and `studioManager.go:583`), and Python-side
   in the in-image studio tree at `studio/tools/pdf2md.py:24` (`mistral-ocr-latest`) — a standalone CLI on
   neither the AI manager's path nor the generation pipeline's
   (`git -C app/studio grep -i mistral aeec036a`, 22 hits / 3 files).

**EU data residency still holds — but by a different mechanism than the ladder implied.** The posture
rests on the *default clients* being EU-resident (Azure EU, Bedrock `eu-west-1`), not on a policy that
walks EU options before US ones. **Two** levers inside the manager can send a request outside the EU,
**neither a region-health failover**: the `flag_use_azure_us` flag, and the 429 retry (which goes straight
to direct OpenAI **without** trying another EU provider first).

**`ANTHROPIC_API_KEY` is a third exit but is NOT within the manager** — an either/or backend switch for
**Course Builder** (`app/internal/coursebuilder/bedrock.go:106-113`), which never touches `AIManager`. An
earlier revision counted it as one of *"exactly three things … within the manager"*, a category error
rather than a miscount. *(It does **not** additionally "flip Studio-Room off Bedrock": Studio-Room was
never on Bedrock — `grep -rin 'bedrock\|boto3' app/studio/` returns **0** hits, and there the key is a
credential while the selector is the ini's `TARGET SERVICE`; see the provider row at
[`external_services.md:571`](external_services.md) (the Mistral row). Corrected M257x iter-48.)*

One layer *above* the manager, a sequence whose **`ai_vendor` is unset** reaches **direct US OpenAI on the
first attempt, with no error condition** — the residency-relevant route, and the one nothing in
`app/internal/jobsimulation/ai/ai.go` can show you. The nullable field is on the **Directus DTO**
(`app/internal/cms/directus/collections/jobsimulation.go:905`, `AIVendor *AIVendor`), not on the domain
`simulation.Sequence`; nil is replaced by `simulation.Openai` at `:1302-1305` **before** the sequence is
built at `:1307`, so the value reaching the vendor switch takes **`case simulation.Openai:`**
(`simulator/ai/ai.go:58-59`) — the same arm an explicit caller takes. Full per-line derivation:
[`external_services.md:619-629`](external_services.md) (item 4 of the four live EU-exit routes).

The switch's `default:` arm (`:114-115`) is a **separate** door: an *unrecognised* vendor string. The
Directus enum has **five** members and the switch **four** cases (`Openai` `:58`, `Azure` `:69`,
`AnthropicAws` `:86`, `Anthropic` `:102`), so `azureglobal` (`jobsimulation.go:971`) is the one with no
case. It resolves to `internalAi.Openai` as well, which is why the *outcome* coincides with the unset
route and the two are easy to conflate.

> An earlier revision of this passage called `AIVendor` *"a **nullable** pointer on the sequence"* and
> routed *unset* through the `default:` arm. Both halves are false, and both contradicted the per-line
> derivation linked above that the same passage cites as authoritative. The **outcome** claim — direct US
> OpenAI, first attempt, no error condition — survives unchanged. Corrected M257x iter-48.

See [The three simulation model defaults](#the-three-simulation-model-defaults).

For the full per-line derivation see
[External Services → Routing: what is actually implemented](external_services.md#routing-what-is-actually-implemented)
— it is the canonical statement; do not restate it here.

### Model Families

The **Default client** column is the client the vendor const resolves to — *not* a position in a
fallback order (there is none; see above).

| Provider | Models | Default client |
|:---------|:-------|:---------------|
| **OpenAI (Azure EU + Direct US)** | GPT-5.4, GPT-5.4-mini, GPT-5.2, GPT-5.1, GPT-5, GPT-5-mini, GPT-5-nano, GPT-4.1, GPT-4.1-mini, O3, O4-mini | Azure **EU** (US via `flag_use_azure_us`). **Direct US OpenAI is NOT only a 429 retry target**: a simulation sequence with `ai_vendor` unset or set to `Openai` reaches it on the first attempt with no error condition, via `case simulation.Openai:` at `simulator/ai/ai.go:58-59` (the `default:` arm at `:114-115` is the separate *unrecognised-vendor* door, not the unset one). Corrected M257x iter-46, mechanism corrected iter-48 |
| **Anthropic (Bedrock EU + Direct US)** | **Claude 4.6 Sonnet**, Claude 4.5 Sonnet, Claude 4 Sonnet, Claude 3.7 Sonnet, Claude 3.5 Sonnet — **five families over six consts** (3.5 has both a Bedrock-EU and a direct-API const), enumerated from `app/internal/ai/anthropic/completion.go:20-30` **@ `app` `ad9f3c49`**; plus **Claude Opus 4.8** (`eu.anthropic.claude-opus-4-8`), which is not in that block — it is Course Builder's `DefaultAuthorModelID` at `internal/coursebuilder/bedrock.go:23`. **This row listed four families until M257x iter-115**, omitting 4.6 — and 4.6 (`eu.anthropic.claude-sonnet-4-6`, `completion.go:29`) is not dormant: it is the current production pin at `internal/askengine/bedrock.go:25`, `internal/jobsimulation/agent/report_agent.go:31` and `internal/coursebuilder/bedrock.go:29`. The row's construction rule is *enumerate the constants* (its sibling — the **OpenAI (Azure EU + Direct US)** row directly above in this same table — is an exact enumeration of `internal/ai/openai/config.go:8-26`), so an omission here is a defect and not an editorial cut | Bedrock `eu-west-1` — both `anthropic-aws` and `anthropic` map here. Direct US is reachable only *outside* this manager, by setting `ANTHROPIC_API_KEY` — for **Course Builder** it is the selector (key set → first-party API); for **Studio-Room**, which was never on Bedrock, it is only the credential the `anthropic` `TARGET SERVICE` needs ([`external_services.md:571`](external_services.md) (the Mistral row)) |
| **Mistral (EU)** | Mistral OCR (`mistral-ocr-latest`) | **OCR only** — the cms domain's Go client plus the in-image `studio/tools/pdf2md.py` CLI; reachable from neither the AI manager nor the generation pipeline |
| **Speech** | GPT-4o Mini TTS (`gpt-4o-mini-tts`) — **the whole set**, and it is also `DefaultModel` | Azure voice client (`CreateSpeech` is Azure-only). Enumerated from `app/internal/ai/speech.go:9-12` **@ `app` `ad9f3c49`**, which is the entire `SpeechModel` const block. **This row also listed "TTS v2 HD, TTS v2" until M257x iter-115.** Those two consts are real, but they belong to the **standalone** `github.com/anthropos-work/ai` module (`speech.go:12-13` @ `v1.40.1`, `TTSV2 = "tts-2"` / `TTSV2HD = "tts-2-hd"`) — a module **no repo a stack builds requires** since the fold at `1e457fa70` (see the **Unified AI Library** section below, which carries that fold); they were dropped in the fold and no caller in the clone set ever referenced them. Both live call sites use `ai.GPT4oMiniTTSS` |
| **Transcription** | GPT-4o Transcribe | Azure EU (US via `flag_use_azure_us`) |
| **Embeddings** | Text Embedding 3 Small | OpenAI |

### Unified AI Library

Go services access AI through one `ai.AI` interface — but **it is no longer a shared private module for any
service a stack builds.** ⚠️ **This paragraph was the repaired half of a pair for four readings.** The same
proposition stood unrepaired in the opening sentence of *AI Providers (External Intelligence)* in
[`external_services.md`](external_services.md) — *"All Go services access
AI through the shared `ai` module"* — until M257x iter-115 closed it; a reader who met that sentence first
got the pre-fold answer with no indication a correction existed. **When you repair this predicate, repair
both sites or neither** (`TOK-07` rule 3). `app` **folded the library into its own tree** at `1e457fa70` (2026-08-04,
*"refactor(ai): fold the ai library into app as internal/ai"*): at `app` `ad9f3c49` neither `app/go.mod` nor
`sentinel/go.mod` requires `github.com/anthropos-work/ai`, and the library lives at `app/internal/ai/`.
The only repos that still *require* the module are the frozen `cms` and `jobsimulation` husks, which nothing
clones or builds. ⚠️ **`sentinel` is measured here as a repo a stack clones, and since platform `766df6c`
it is not one** (v11.0, 2026-08-11 — the 8th merge; folded into `app` as `app/internal/sentinel/`, compose
service and `repos.yml` entry deleted together). The `sentinel/go.mod` reading is unchanged and still true —
it never required `ai` — but it now describes a **frozen** repo alongside `cms` and `jobsimulation`, not a
live one, so *"Go services"* in this section means **`app`, and only `app`**. What the interface provides is unchanged:
- A single `ai.AI` interface across providers (OpenAI, Azure, Anthropic, Bedrock, Mistral)
- Per-provider client constructors that return provider token counts (`MetaData.Usage`)

> **Vendor selection/fallback and cost tracking are NOT in the `ai` library** — they live in the consuming services: selection/fallback in each consumer's own wrapper — in `app` @ `5ba17044` that is **`app/internal/jobsimulation/ai/ai.go:267,344`** and **`app/internal/skillerai/ai.go:347`**, *not* a bare `app/internal/ai/ai.go`. **NB that file DOES exist** at `app` `ad9f3c49` — it is the folded `ai.AI` interface (`1e457fa70`), and this parenthetical asserted *"no such file"* until M257x iter-108; what remains true is only that **vendor selection does not live there**, which is the point the sentence is making. The Azure client defaults to EU and swaps to US on the PostHog flag `flag_use_azure_us`; direct OpenAI is the retry target on HTTP 429; Anthropic is always Bedrock `eu-west-1`. **These are three independent mechanisms, not rungs of an ordered ladder** — the ⚠️ under *Provider Routing Strategy* at the head of this file (`:15-17`) retracts that ladder, and this line went on publishing it 68 lines below, in `→` form, until M257x iter-48. Cost tracking is in `app/internal/aiusage/ai_usage.go` (fed by `Event_AiUsage` over Redis Streams). See [Shared Libraries → ai](shared_libraries.md#ai).

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
> **shipping** studio config — it *is* in two slots of the non-shipping scaffold this blockquote opens by
> telling you not to read: `configs/config_template.ini:39` (`EXECUTION_AI_STABLE_MODEL = azure, gpt-4o,
> none`) and `:40` (`CREATIVE_AI_STABLE_MODEL`). The unqualified form contradicted this same blockquote two
> lines earlier, which already says the template *"still carries"* `gpt-4o`; corrected M257x iter-46.
> Agrees with [`studio-room.md`](../services/studio-room.md#ai-service-configuration).

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
Player → LiveKit Room → GPT Realtime Agent (anthropos-agent [EU] / anthropos-agent-us)
```

> **The agents have repos, and they are now documented (v2.8 M257x iter-123).** The org holds **five**
> LiveKit agent repositories — `livekit-agent`, `livekit-agent-chain`, `livekit-agent-azure-us`,
> `livekit-agent-azure-eu`, `livekit-agent-azure-eu-fr` — none of which appear in `repos.yml`. Everything
> here documents the LiveKit *engine* and the *platform side* of the call; **the agent side now has a
> home**, [`org-repos.md` § 8](./org-repos.md), with the agent-name → repo map, the `livekit.toml` ids,
> and the mechanism.
>
> **This paragraph used to end *"none of which appear in `repos.yml`, in any corpus document, or in the
> deployment picture"* — a sentence that named all five while denying any corpus document names them, and
> then cited the table that enumerates them two lines later. Only the `repos.yml` clause was ever
> measured-true; corrected iter-123.** A universal quantifier falsified by its own sentence is the defect,
> not the staleness.
>
> **The finding the enumeration was hiding:** `azure-eu` and `azure-eu-fr` dispatch nothing not because
> they were abandoned but because they were **ABSORBED** — `livekit-agent` is now one agent serving four
> endpoints selected by dispatch metadata (`app`'s `metadata["endpoint"]`,
> `internal/jobsimulation/calls/livekit.go:148` ↔ `livekit-agent/src/agent.py:268`), and the three
> azure-* repos never read that key at all. The absorption is **incomplete**: US is still dual-routed, so
> `app` sets both `agentName="anthropos-agent-us"` and `metadata["endpoint"]="azure-us"`
> (`livekit.go:126-127`) and the receiving agent ignores the second.

- **Audio**: Real-time voice conversation, recorded as MP3
- **Transcript**: Generated from LiveKit conversation events
- **Configuration**: Voice engine is selectable per simulation in CMS (`livekitgptrealtime`)

### Legacy / Transitioning Engines

| Engine | Status | Description |
|:-------|:-------|:------------|
| `elevenlabs` | Active (legacy default) | ElevenLabs conversational agents; still used by the call/reply pipeline (`getJobSimulationCallSignedUrl` / `getJobSimulationCallConversationToken`) and transcript improvement |
| `gptrealtime` | Deprecated | Direct OpenAI Realtime without LiveKit |

**Engine choice is per SEQUENCE, from the CMS `voice_engine` field** — a 4-member enum on the authored
simulation (`app/internal/cms/directus/collections/jobsimulation.go:1079-1085`); when it is nil the content
layer supplies `gptrealtime` — **`app/internal/cms/directus/collections/jobsimulation.go:1594-1597` @ `app`
`ad9f3c49`**, the whole of `func voiceEngineFromDirectus` down to the closing brace of its nil branch
(`:1595-1596` is `if directusVoiceEngine == nil { return simulation.SimulationVoiceEngineGptrealtime }`).
**Pinned in full, and to `:1597` rather than `:1600`, deliberately (M257x iter-115):** this is one half of a
same-fact-different-pin pair with the **ElevenLabs** bullet under *External Dependencies* in
[`jobsimulation.md`](../services/jobsimulation.md), whose half spelled
the path `cms/directus/collections/…` — a path that exists in **no** clone. Fixing one side and leaving the
other is how the corpus acquires a self-contradiction, so both halves now carry the same path and the same
range. **ElevenLabs remains the active default** for the call/reply
pipeline and transcript improvement, so it is not yet fully replaced.

⚠️ **`flag_use_realtime_openai` does NOT select LiveKit over ElevenLabs, and gates no "new sessions"** —
this paragraph said so from 2026-06-01 until M257x iter-49. The flag is read **inside** `CreateAgentDispatch`,
i.e. on the LiveKit path the request has *already* entered (`calls/livekit.go:131-135` read, `:140-144`
effect). ⚠️ **The effect block does TWO things, and this sentence claimed one until M257x iter-115.** At
`app` `ad9f3c49` the block is four lines: `:142` `agentName = "anthropos-agent"` **and** `:143`
`agentEndpoint = "openai-hosted"`. The endpoint swap is the one the paragraph was written for; the
**agent-name reset is a real behavioural change on a US-located session** — `:126` has just set
`anthropos-agent-us`, and flipping the flag silently re-dispatches to the bare `anthropos-agent`. *"All it
does is swap the endpoint"* is an explicit completeness assertion about a block the sentence cites **by
line**, which is what made it checkable and what made it false. **This is
residency-relevant:** flipping it moves a live voice session off an Azure-EU endpoint — a path the routing
section above does not otherwise cover.

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
2. **Route**: Selected model from the CMS `ai_model` / `ai_vendor` fields per sequence (e.g. `gpt-5`, `gpt-4.1`, `anthropic-45-sonnet-aws` — the enum is `app/internal/cms/directus/collections/jobsimulation.go:983-990` @ `2035f9a` — re-pinned M257x iter-126). **There is no single "default model", and `gpt-5` is not a default anywhere.** Three distinct defaults apply at three different points — see below
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
  their own: `anthropic-aws` falls back to Claude 3.7 Sonnet on Bedrock and `anthropic` to
  Claude 3.5 Sonnet. **Both anchors named the file nowhere in this bullet until M257x iter-115** — they
  read as bare `:98-99` / `:110-111`, which is §5 rule 22's failure mode (an unowned line number rots
  silently). The file is `app/internal/jobsimulation/simulator/ai/ai.go` **@ `app` `ad9f3c49`**, where
  `:98-99` is `default: aiModel = anthropic.Anthropic37SonnetAWS20250219` and `:110-111` is
  `default: aiModel = anthropic.Anthropic35Sonnet20241022` — both re-derived and both true at that ref.
  **NB this bullet is not a claim about the platform's Anthropic model SET** (that is the **Anthropic
  (Bedrock EU + Direct US)** row of the *Model Families* table, which is where the 4.6/Opus-4.8 omission
  lived); these two are *fallback* arms and nothing more

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
