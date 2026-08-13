# iter-47 — adjudication

**Rule:** verify every reported blocker before accepting it. §5 — verify a claim before escalating it,
*including a claim made by an audit*: two of iter-22's handed corrections were themselves false, and
iter-01 refuted five inherited claims by measurement.

| id | auditor | re-verification performed by this iteration | verdict |
|---|---|---|---|
| **C1** | C | Read `architecture_overview.md:244-252` verbatim: it states *"the two US paths are a **feature flag** and a **429 retry target**"* and *"one feature flag routes traffic to the US."* Read the source it cites, `external_services.md:532`, verbatim: the **OpenAI Direct (US)** row states *"**two ways in**: (a) `vendor = Openai` from the caller — **including the case where the caller never chose, since a simulation sequence with `ai_vendor` unset defaults to `openai`** … (b) automatic on HTTP 429"*, and closes *"The 429 retry is the only automatic fallback — but it is **not** the only route to US OpenAI. **Path (a) gets there on the first attempt.**"* The two documents contradict each other, the platform source (`jobsimulation.go:1302` → `ai.go:80` `openai.NewOpenAI(openaiKey)`, no region override) supports `external_services.md`, and the claim is residency-relevant. | **HELD** |
| **B1** | B | Re-derived independently. `grep -rn "simulation.Sequence{" --include="*.go" internal/` → **exactly one** construction site, `internal/cms/directus/collections/jobsimulation.go:1307`. `:1302-1304` reads `aiVendor := simulation.Openai; if seq.AIVendor != nil { aiVendor = … }` — so nil is normalized to `Openai` **before** the struct is built, and `simulation.Sequence.AIVendor` is a value. The nullable pointer is on the DTO at `:905`. `AIVendor` has **five** members (`:967-973`) including `Azureglobal`, which `GetAIVendorAndModel` has no case for — so the `default:` arm is reachable only by an **unrecognised** value. iter-46's text asserts the opposite of all three facts. | **HELD** |
| **B2** | B | `sed -n '489p' corpus/architecture/external_services.md` → `// Types in app/__generated__/`. The **Anthropic Direct (first-party API)** row is at `:533`. The anchor was transcribed from iter-41's ledger without re-derivation — the precise failure `D-M257x-46-1` claimed to have eliminated. | **HELD** |
| **G1** | G | **Duplicate of B1** — same site, same refutation, reached independently by two auditors from different seats (B by full-read of the file, G by diff-read). Counted **once**. | **HELD (dup of B1)** |
| **G2** | G | Re-derived: `grep -rin 'bedrock\|boto3' stack-demo/app/studio/` → **0 hits**; `grep -n "class.*Provider" app/studio/services/ai.py` → `:231 AIProvider(ABC)`, `:334 OpenAIProvider`, `:490 AzureProvider`, `:627 AnthropicProvider`. Studio-Room has no Bedrock path to be flipped *off*. iter-46 rewrote this sentence and preserved the false conjunct. | **HELD** |
| **G3** | G | `sed -n '137,141p' corpus/architecture/external_services.md` still reads *"**all of that is false**; that service has never existed in the platform compose"* — verbatim the over-correction iter-46 repaired at `service_taxonomy.md`. Same git evidence (`a2a3ee6^` `:383`/`:384`/`:386`/`:409`) refutes it. | **HELD** |
| **G4** | G | `sed -n '84p' corpus/architecture/ai_architecture.md` still contains *"EU Azure default → US Azure via the PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429"*. `sed -n '15,17p'` of the **same file** reads *"**There is no ordered EU-first fallback chain** … **no such ladder exists in the code**"*. A file contradicting itself 68 lines apart, in the exact class this milestone measured as 13 of 18. | **HELD** |
| **G5** | G | `sed -n '613,617p' corpus/ops/demo/coverage-protocol.md` still reads *"the default AI-readiness dashboard GET **always** takes the live-recompute branch … gated behind a *closed* `CycleID`"*, in the present tense, 13 lines **above** the fence iter-46 added. Refuted by `readiness.go:309-312` (already re-derived at iter-46). | **HELD** |

## Result

**7 reported blockers across 3 seats (B×2, C×1, G×5); G1 ≡ B1 → 7 unique. All 7 re-verified by this
iteration against platform source. 7 of 7 HELD — none refuted on re-derivation.**

**Every one of the 7 sits in text M257x iter-46 wrote or rewrote.** Six full-read auditors covering all
40 files / 9,243 lines found **zero** blockers in text iter-46 did not touch.
