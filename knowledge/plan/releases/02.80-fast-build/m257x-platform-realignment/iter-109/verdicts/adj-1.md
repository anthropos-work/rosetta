# Adjudicator 1 — seat A, readings #27 + #28

**Trees I read.** All 14 platform clones + both `rosetta-extensions` trees were re-verified at this
adjudication's open and every sha matched the brief's ground-truth table exactly
(`platform 0c91421d` · `app ad9f3c49` · `app/studio aeec036a` · `cms/studio aeec036a` ·
`next-web-app 8297c684` · `sentinel f2c46190` · `studio-desk 41ee3575` · `ant-academy 22df69dd` ·
`cms ca50c817` · `jobsimulation 462343b0` · `messenger fa47850d` · `storage 4ce8ece5` ·
`roadrunner 87d8d443` · `graphql-wundergraph 60c229f3` · rext **consumption** `09d06070` ·
rext **authoring** `680e8529`). No fetch, no pull, no checkout, no edit outside this file.
**No booking in this seat-group turned on a `rosetta-extensions` claim**, so the two-tree rule was
not load-bearing here; where I touched rext at all it was as a *negative*-control sweep and I used the
**pinned consumption clone `09d06070`**, saying so at the point of use.

**Instrument discipline.** Every absence claim below was taken with all three mechanisms the brief
names: `git grep` at each tree's own ref across all 15 git checkouts found by
`find stack-demo -maxdepth 3 -name .git` (so the two nested `studio` repos are read at `aeec036a`,
not through the host ref), **plus** a `.gitignore`-blind filesystem `grep -rI`, **plus** a live
positive control in the same pass. One absence I could not close from the clone set alone
(`TTS v2`) I closed against the **local Go module cache**, read-only — see B4.

---

## Verdicts

### r27-A B1 · r28-A B1 — same anchor, same predicate

```
r27-A B1 | corpus/architecture/ai_architecture.md:34 | UPHELD | IN-SCOPE | PREDICATE: The Mistral OCR client in `app` is constructed at markdownManager.go:19.
```
```
r28-A B1 | corpus/architecture/ai_architecture.md:34 | UPHELD | IN-SCOPE | PREDICATE: The Mistral OCR client in `app` is constructed at markdownManager.go:19.
```

evidence: I opened `stack-demo/app` at `ad9f3c49` (the block names no `app` ref — its only ref,
`aeec036a`, scopes the *studio-tree* grep at the end of the same bullet, so it grades at the
checkout). `internal/cms/studio/markdownManager.go:17-19` is the middle of the doc comment on
`NewMarkdownManager` — `:19` reads `// It used to take aiKey and then IGNORE it, reading
os.Getenv("MISTRAL_API_KEY")`. The construction is at **`:30`**
(`return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil`), inside `func NewMarkdownManager` at
`:29`; the field `ocr *mistralocr.Client` is at `:14` and the import at `:10`. The *other* anchor in
the same citation, `studioManager.go:583`, **does** resolve
(`markdownManager, err := NewMarkdownManager(os.Getenv("MISTRAL_API_KEY"))`), so exactly one of the
two is wrong.

This is also a **live corpus self-contradiction and not a retraction pair** (brief rule 5):
`corpus/architecture/external_services.md:565` states of this very citation *"the `:19` this row used
to cite is a **doc-comment** line, not code"* and cites `markdownManager.go:30`, and
`external_services.md:598` cites `:30` again — while `ai_architecture.md:34` goes on asserting `:19`
as live. One twin repaired, the other left standing.

Correction to r27's derivation (does not change the verdict): r27 states the file is
"byte-identical at `b948604`". It is not — `git diff b948604f ad9f3c49 -- internal/cms/studio/markdownManager.go`
is 25 insertions / 15 deletions, and at `b948604` `:19` was
`ai, err := mistral.NewMistral(nil, os.Getenv("MISTRAL_API_KEY"))`, i.e. correct. r28's account is
the accurate one. The verdict is unaffected because the sentence carries no `app` pin.

---

### r27-A B2

```
r27-A B2 | corpus/architecture/external_services.md:554 | UPHELD | IN-SCOPE | PREDICATE: All Go services access AI through the shared `ai` module `github.com/anthropos-work/ai`.
```

evidence: I read the `go.mod` of every Go repo in the clone set myself.
`stack-demo/app` @ `ad9f3c49` `go.mod:14-18` requires `analytics-go`, `colony v0.35.2`,
`proto v1.210.0`, `storage v0.15.2`, `taxonomy v1.2.0` — **no `github.com/anthropos-work/ai`**.
`stack-demo/sentinel` @ `f2c46190` `go.mod:8,9,21` requires colony + proto + taxonomy only, and uses
no AI at all. `git ls-tree ad9f3c49 internal/ai/` returns **17** entries (`ai.go`, `anthropic/`,
`openai/`, `speech.go`, …) and **67** `.go` files import
`github.com/anthropos-work/app/internal/ai`. The module survives as a `require` only in the frozen
`cms` (`go.mod:9`) and `jobsimulation` (`go.mod:11`) husks — neither has a compose service, and
`platform/repos.yml` @ `0c91421d` lists exactly **two** Go repos, `app` and `sentinel` (I read the
file; its own header comment says the merged repos "own no local schema, no compose service and no
clone entry here").

Not a pointer (brief rule 6): the `→ shared_libraries.md#ai` link in that paragraph is attached to
the *selection/cost-tracking* sentence, not to the "shared library" assertion, which stands alone.
Not a retraction pair (rule 5): I grepped `external_services.md` for `1e457fa|internal/ai|folded` —
**the fold is nowhere acknowledged in that file**, so `:554` is asserted live. Three sibling
documents were already repaired on exactly this proposition and now contradict it:
`ai_architecture.md:95-99` (*"it is no longer a shared private module for any service a stack
builds"*), `shared_libraries.md:3-10` (*"`ai` was folded into `app` as `app/internal/ai` at
`1e457fa70` … which dropped its module requirement"*), and `architecture/README.md:21` (*"only
**three** are imported as private modules … **`ai` is no longer one of them**"*).

---

### r27-A B3 · r28-A B2 — same anchor, same predicate

```
r27-A B3 | corpus/architecture/ai_architecture.md:87 | UPHELD | IN-SCOPE | PREDICATE: The platform's Anthropic models are Claude 4.5 / 4 / 3.7 / 3.5 Sonnet.
```
```
r28-A B2 | corpus/architecture/ai_architecture.md:87 | UPHELD | IN-SCOPE | PREDICATE: The platform's Anthropic models are Claude 4.5 / 4 / 3.7 / 3.5 Sonnet.
```

evidence: I enumerated the SET before comparing cardinalities (brief rule 4). `stack-demo/app` @
`ad9f3c49`, `internal/ai/anthropic/completion.go:20-30` is a **six**-const block over **five**
families:
`:21 Anthropic35SonnetAWS20241022` · `:22 Anthropic35Sonnet20241022` ·
`:23 Anthropic37SonnetAWS20250219` · `:24 Anthropic4SonnetAWS20250514` ·
`:25 Anthropic45SonnetAWS20251126` · **`:29 Anthropic46SonnetAWS20251126 = "eu.anthropic.claude-sonnet-4-6"`**.
The row lists four families; **Sonnet 4.6 is absent.** The sibling OpenAI row at `:86` is an *exact*
enumeration of `internal/ai/openai/config.go:8-26` (eleven consts, eleven listed), so the table's own
construction rule is "enumerate the constants" — the omission is not a stated editorial cut.

The omitted member is the platform's current production pin, not a dormant const:
`internal/askengine/bedrock.go:25 DefaultModelID`,
`internal/jobsimulation/agent/report_agent.go:31 defaultAgentModel` and
`internal/coursebuilder/bedrock.go:29 DefaultGraderModelID` all equal
`"eu.anthropic.claude-sonnet-4-6"`. `internal/coursebuilder/bedrock.go:23 DefaultAuthorModelID =
"eu.anthropic.claude-opus-4-8"` is missing from the row as well — and the row's own right-hand cell
names **Course Builder** as the reason the Direct-US door exists. Intra-corpus inconsistency too:
`external_services.md:564` gives both `eu.anthropic.claude-sonnet-4-6` and
`eu.anthropic.claude-opus-4-8` as the live Bedrock-EU models, in a document `ai_architecture.md:3`
calls the AI **model inventory**. The row carries no pin and no date, so it grades at the checkout.

Both seats book the same false proposition at the same anchor; r28 additionally names Opus 4.8. Two
omissions from one list are **one** predicate, not two.

---

### r27-A B4

```
r27-A B4 | corpus/architecture/ai_architecture.md:89 | UPHELD | IN-SCOPE | PREDICATE: The platform's speech-model set includes TTS v2 and TTS v2 HD.
```

evidence: `stack-demo/app` @ `ad9f3c49`, `internal/ai/speech.go:9-12` is the **whole** `SpeechModel`
const block: `GPT4oMiniTTSS SpeechModel = "gpt-4o-mini-tts"` and `DefaultModel = GPT4oMiniTTSS`.
Nothing else. Three-instrument sweep for `tts v2|tts-v2|ttsv2|tts_v2|tts-1|tts-hd`: `git grep` at
each of the **15** trees' own refs (positive control `tts` non-zero in 11 of them — `app` 164 lines,
`ant-academy` 1012, `next-web-app` 32, `studio-desk` 24, `jobsimulation` 23, both nested `studio`
checkouts 1 each at `aeec036a`) plus a `.gitignore`-blind `grep -rI` over the whole `stack-demo`
filesystem (control: `gpt-4o-mini-tts` matches `app/internal/ai/speech.go`,
`app/internal/jobsimwiring/wiring.go`, `app/terraform/{variables,ssm}.tf`, `platform/.env_example`).
**Every "TTS v2" hit in the entire clone set is a copy of this corpus line** — two
`rosetta-extensions` (pinned clone `09d06070`) test *fixtures* under
`stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/architecture/ai_architecture.md`. The
ant-academy hits are course prose about third-party `tts-1`/`tts-1-hd`, not platform models.

I closed the seat's stated residual, and it closes against the claim. The **standalone**
`github.com/anthropos-work/ai` module is in the local Go module cache
(`~/go/pkg/mod/github.com/anthropos-work/ai@v1.40.1`), and `speech.go:12-13` there really does declare
`TTSV2 = "tts-2"` and `TTSV2HD = "tts-2-hd"` — so that is where the corpus row came from. But those
consts were **dropped in the fold** into `app/internal/ai`, the module is required by no `go.mod` a
stack builds, and **no caller anywhere in the clone set ever referenced them**: `git grep
'TTSV2|TTSModel1|tts-2|tts-1'` across `app`, `jobsimulation`, `cms`, `messenger`, `storage` returns
only `app/internal/ai/speech.go:10-11`, and both live call sites use `ai.GPT4oMiniTTSS`
(`app/internal/jobsimulation/simulator/outbound/prompts/replyAudio.go:72`, and its husk twin). The
row carries no pin and no date; the same document announces the fold at `:95-99`, so it grades at the
checkout. Not ref-discipline.

---

### r27-A B5 · r28-A B5 — same anchor, same predicate

```
r27-A B5 | corpus/services/coursebuilder.md:48 | UPHELD | IN-SCOPE | PREDICATE: variables.tf:635-638 declares anthropic_api_key and main.tf:555 injects ANTHROPIC_API_KEY.
```
```
r28-A B5 | corpus/services/coursebuilder.md:48 | UPHELD | IN-SCOPE | PREDICATE: variables.tf:635-638 declares anthropic_api_key and main.tf:555 injects ANTHROPIC_API_KEY.
```

evidence: I opened all three anchors at **both** candidate refs in `stack-demo/app`.

| anchor | `b948604f` | `ad9f3c49` (checkout) |
|---|---|---|
| `terraform/variables.tf:635-638` | `variable "anthropic_api_key" {` at `:635`, `sensitive = true` at `:638` ✔ | a **cms-in-app comment block** (`:631-645`): *"empty → dormant, the standalone cms module stays the live path until the cutover (rollback path until M810)…"*. The variable is at **`:759-763`** ✘ |
| `terraform/ssm.tf:328-334` | `resource "aws_ssm_parameter" "anthropic_api_key"` at `:328` ✔ | identical at `:328-333` ✔ |
| `terraform/main.tf:555` | `"name": "ANTHROPIC_API_KEY"` ✔ | `"name": "DIRECTUS_BASE_ADDR"` (value `${var.directus_base_addr}` at `:556`). The `ANTHROPIC_API_KEY` injection is at **`:757-758`** ✘ |

**Why this is not ref-discipline.** The brief's rule 1 fixes a pin's scope at "the claim's own block —
a table cell, a wrapped sentence." The only pin in this bullet is *nested inside a parenthetical
attached to a different file*: `…reporting ModelBackendName() == "anthropic-api" (`:98-104`, logged at
`main.go:770` @ `app` `b948604` v1.366.0);`. I verified that pinned claim independently — `main.go:770`
@ `b948604` is exactly `logger.Info("coursebuilder model backend", "backend",
coursebuilder.ModelBackendName())`, so the pin is doing real work where it sits. It cannot reach a
later, separate sentence about terraform, which therefore names no ref and grades at the checkout.
There the sentence asserts two propositions about two file:line pairs and both are false.

The anchors do not merely drift by a few lines — they land on **different subjects** (a cms-in-app
secrets comment; a Directus base address), which is the maximally misleading failure mode: a reader
opening `main.tf:555` at HEAD sees a Directus variable and reads the whole production-key claim as
wrong. The substantive claim ("in production the key is required, so the shipped path is the direct
API") is in fact **true** at `ad9f3c49` — `variables.tf:759-763` is `sensitive = true` with no
default, `ssm.tf:328-333` creates the SecureString, `main.tf:757-758` injects it from the SSM ARN — so
what is booked is precisely the citation pair, and nothing more.

I record the honest counterweight both seats raised: if a markdown *bullet* is the pin's block rather
than a sentence, all five anchors verify at `b948604` and this is ref-discipline. I resolve it against
that reading because rule 1 names a sentence, and because this pin is syntactically narrower than a
sentence.

---

### r28-A B3

```
r28-A B3 | corpus/architecture/external_services.md:723 | UPHELD | IN-SCOPE | PREDICATE: The LiveKit EU/US split is on the endpoint only, not on the agent name.
```

evidence: `stack-demo/app` @ `ad9f3c49`, `internal/jobsimulation/calls/livekit.go` (the passage names
no `app` ref, so it grades at the checkout):

```
118  if location != nil && location.IsValid() {
119      if *location == LocationEu {
120          agentName = "anthropos-agent"
122          randIdx := rand.Intn(len(euAgentEndpoints))
123          agentEndpoint = euAgentEndpoints[randIdx]
124      } else {
125          // we use anthropos-agent-us and azure-us for the us location
126          agentName = fmt.Sprintf("anthropos-agent-%s", *location)
127          agentEndpoint = fmt.Sprintf("azure-%s", *location)
```

`agentName` is assigned **differently on the two location branches** — the split is on the agent name
as well as the endpoint. The same sentence says so eleven words earlier (*"the US one is suffixed
`anthropos-agent-us` (`:126`)"*), so this is a one-sentence self-contradiction with both halves
asserted live (brief rule 5), not a retraction.

The parenthetical is separately wrong about the endpoint set: EU does not resolve to `azure-eu`, it
resolves to a **random** member of `euAgentEndpoints = {"azure-eu", "azure-eu-fr"}` (`:101-104`,
`:122-123`); `azure-eu` is only the pre-branch default at `:111`.

The *other* half of the sentence is true and I checked it rather than assuming: `anthropos-agent-eu`
is **0** hits across all 15 trees at their own refs and **0** in a `.gitignore`-blind filesystem grep
of platform code — the only matches anywhere are corpus test fixtures inside the pinned rext clone
`09d06070` (`stack-core/tests/fixtures/**`), i.e. copies of this corpus. Positive control:
`anthropos-agent-us` matches `app/internal/jobsimulation/calls/livekit.go` and its husk twin.

---

### r28-A B4

```
r28-A B4 | corpus/architecture/ai_architecture.md:218 | UPHELD | IN-SCOPE | PREDICATE: The flag_use_realtime_openai arm's only effect is swapping the endpoint to openai-hosted.
```

evidence: the sentence cites its own effect block by line, and I opened it.
`stack-demo/app` @ `ad9f3c49`, `internal/jobsimulation/calls/livekit.go:140-144`:

```
140  if isRealtimeOpenaiFlagEnabled {
141      // if the feature flag is enabled, use the livekit with openai key and not azure key
142      agentName = "anthropos-agent"
143      agentEndpoint = "openai-hosted"
144  }
```

The block does **two** things. `:142` resets the agent name, which is a real behavioural change on a
US-located session — `:126` had just set `anthropos-agent-us`, and flipping the flag silently
re-dispatches to the bare `anthropos-agent`. "All it does is swap the … endpoint" is an explicit
completeness assertion about a block the sentence cites by line, and it is false. (The read anchor
`:131-135` and the effect anchor `:140-144` both resolve exactly, and the residency conclusion the
paragraph draws is unaffected — what is booked is the completeness claim alone.)

Distinct from B3: B3 denies that the *location* branch splits the agent name; B4 denies that the
*flag* branch touches it. Two propositions about two different code paths — not collapsed.

---

## PREDICATE ROLL-UP

```
P1 | The Mistral OCR client in `app` is constructed at markdownManager.go:19. | anchors: r27-A B1 @ corpus/architecture/ai_architecture.md:34, r28-A B1 @ corpus/architecture/ai_architecture.md:34
P2 | All Go services access AI through the shared `ai` module github.com/anthropos-work/ai. | anchors: r27-A B2 @ corpus/architecture/external_services.md:554
P3 | The platform's Anthropic models are Claude 4.5 / 4 / 3.7 / 3.5 Sonnet. | anchors: r27-A B3 @ corpus/architecture/ai_architecture.md:87, r28-A B2 @ corpus/architecture/ai_architecture.md:87
P4 | The platform's speech-model set includes TTS v2 and TTS v2 HD. | anchors: r27-A B4 @ corpus/architecture/ai_architecture.md:89
P5 | variables.tf:635-638 declares anthropic_api_key and main.tf:555 injects ANTHROPIC_API_KEY. | anchors: r27-A B5 @ corpus/services/coursebuilder.md:48, r28-A B5 @ corpus/services/coursebuilder.md:48
P6 | The LiveKit EU/US split is on the endpoint only, not on the agent name. | anchors: r28-A B3 @ corpus/architecture/external_services.md:723
P7 | The flag_use_realtime_openai arm's only effect is swapping the endpoint to openai-hosted. | anchors: r28-A B4 @ corpus/architecture/ai_architecture.md:218
```

Cross-reading collapses: **3 of the 7 predicates were booked by both seats** (P1, P3, P5), always at
the *same* anchor line — this seat-group's two readings converged rather than spreading. 10 bookings
→ 7 distinct predicates → 7 distinct anchor lines.

Note on P7 vs the other seat's grading: r27-A recorded the same fact as a **Minor**, not a blocker
(*"the residency point the sentence makes is correct"*). I adjudicate only what was booked as a
blocker, so P7 counts once, from r28-A. Had r27-A booked it, it would have collapsed onto P7 rather
than added an anchor.

---

BOOKED=10 UPHELD=10 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=10 DISTINCT-IN-SCOPE-PREDICATES=7
