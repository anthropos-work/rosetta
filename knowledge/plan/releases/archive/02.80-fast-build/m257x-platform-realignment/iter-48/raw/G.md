# Seat G (diff) — iter-48

Subject: `cabc3b1` *"fix(M257x/48): the seven repaired by CLAIM, with the leak fence as the commit post-condition"*.
Ground truth read directly: `stack-demo/app` @ `5ba17044`, `stack-demo/platform` @ `2adcf71`, and — a source
this seat found and no prior pass appears to have used — **`stack-demo/app/studio/` is itself a full,
non-shallow git clone of `anthropos-studio-room`** (677 commits, first commit 2024-09-18), so Studio-Room
history *is* checkable locally.

## Coverage (hunks reviewed, files touched)

All 7 corpus/CLAUDE files in the commit, every hunk, every anchor:

| File | Hunks | What I re-derived from source |
|:--|:--|:--|
| `CLAUDE.md` | 2 (`:232`, `:246`) | anchor re-point `537 → 545` only |
| `corpus/architecture/ai_architecture.md` | 3 (`:42-68`, `:83-84`, `:96`) | the whole vendor-switch mechanism, arm by arm |
| `corpus/architecture/architecture_overview.md` | 1 (`:244-256`) | the four-exit scope claim + `getClient`/`isThrottlingError` anchors |
| `corpus/architecture/external_services.md` | 2 (`:139-146`, `:570-576`) | the Directus-compose history + the Studio-Room provider row |
| `corpus/architecture/security_compliance.md` | 3 (`:7`, `:183-188`, `:197-203`) | `coursebuilder/bedrock.go`, the `:489 → :541` re-derivation |
| `corpus/architecture/shared_libraries.md` | 1 (`:125`) | anchor re-point only |
| `corpus/ops/demo/coverage-protocol.md` | 1 (`:614-620`) | `aireadiness/readiness.go` + `live_snapshots.go` |

**Source files opened line-by-line** (not quoted from the commit message):
`internal/jobsimulation/simulator/ai/ai.go` (:20-26, :56-129) ·
`internal/cms/directus/collections/jobsimulation.go` (:834-915, :967-991, :1295-1345) ·
`internal/jobsimulation/ai/ai.go` (:27-33, :78-95, :125-170, :259-330) ·
`internal/skillerai/ai.go` (:191-206, :344-350) ·
`internal/coursebuilder/bedrock.go` (:95-120) · `main.go` (:754-764) ·
`internal/aireadiness/readiness.go` (:280-320, :761-800) · `internal/aireadiness/live_snapshots.go` (:50-130) ·
`internal/askengine/registry.go:526`, `internal/askengine/rules.md:746-773` ·
`studio/services/ai.py` (:1-2, :334-390, :490-540, :627-664, :704-716) ·
`studio/configs/{development,production,config_template}_config.ini` ·
`platform` compose at `a2a3ee6` and `a2a3ee6^`.

**Guards I ran myself** (not taken from the commit message): `anchor_construct_guard` → OK, 107 resolved
(130 unresolvable, incl. the `internal/…`-prefixed heads — those I checked by hand). `repair_leak_guard
--range HEAD~1..HEAD` → **GREEN**, 635 candidate shingles. `derived_value_guard` → OK.
`claim_twin_guard` → RED on 2 sites, both adjudicated by **seat A's** iter-48 report (not mine; not
double-booked here).

### Verified TRUE — do not re-spend budget on these next pass

Every mechanism the commit asserts for blockers #1, #2, #4, #5 checks out against source:

* `AIVendor *AIVendor` is on the **Directus DTO** at `jobsimulation.go:905`; the domain
  `simulation.Sequence` takes it **by value** (`:1343 AIVendor: aiVendor`, and `simulator/ai/ai.go:56`
  `switch sequence.AIVendor` with bare-value `case` arms would not compile against a pointer).
* nil→`simulation.Openai` at `:1302-1305`, **before** the one and only `simulation.Sequence{` literal in
  the entire repo at `:1307` (`grep -rn "simulation.Sequence{"` → exactly 1 hit). So *unset* reaches
  `case simulation.Openai:` and never the `default:` arm. Confirmed.
* Directus `AIVendor` enum = **five** members, `:970-974`; `Azureglobal` at **:971**. Switch = **four**
  cases at **:58 / :69 / :86 / :102**; `default:` at **:114-115** (the old `:113-115` included the
  `// OpenAI` comment line). `azureglobal` is the odd one out. Confirmed arm by arm.
  *(Caveat for the record: the proto `simulation` package is not vendored locally, so "the string values
  match" is inference — but `:1304`'s bare string conversion plus a working platform makes it safe, and
  nothing contradicts it.)*
* `bedrock.go`: `func ModelBackendName()` **:98**, `return "anthropic-api"` **:100** ✅,
  `func newUnderlyingClient` **:109**, `NewAnthropicClientWithModel` **:111** ✅ — so
  `security_compliance.md:197-198`'s `:109-112` / `:100` are both **correct** (I initially miscounted and
  re-checked with `grep -n`; they are right). `main.go:762` = the `logger.Info("coursebuilder model
  backend"…)` line, so `:756-762` ✅.
* The `:489 → :541` re-derivation is **correct and honest**: pre-commit `external_services.md:489` is
  literally `// Types in app/__generated__/`, a TypeScript codegen comment, exactly as the commit says.
* Directus compose history: `platform a2a3ee6^:docker-compose.yml` has 7 `directus` hits — `:384
  image: directus/directus:10.10.1`, `:386 8055:8055`, `:409 ADMIN_PASSWORD=password`; `a2a3ee6` has **0**.
  `git log -S "admin@example.com" --all` → 0 commits. Every figure in the correction-of-the-correction ✅.
* All ten re-pointed anchors resolve to the right construct: `:545` = *"There is **no ordered EU-first
  fallback chain.**"*; `:541` = the **Anthropic Direct** provider row; `:569` = *"**Four** things can send a
  request outside the EU"*; `:577-587` = the per-line unset-`ai_vendor` derivation. Zero stale
  `external_services.md:537` / `:489` anchors survive anywhere in scope.
* `readiness.go:308-312` does span the nil-`CycleID` → `buildResponseFromSnapshots` branch (`:309` the
  `activeCycle == nil` guard, `:311` the snapshot return) ✅.
* Pre-existing anchors in the rewritten lines also hold: `ai.go:267` / `:344`, `skillerai/ai.go:347`,
  `isThrottlingError` `:129` applied at `:166` and `:325`, `WithRegion("eu-west-1")` `:87`,
  `openai.NewOpenAI(openaiKey)` `:80`, `getClient` `:259-289`, `ai.py:627-664` = `AnthropicProvider`.

## Blockers

| # | Site | The false claim (verbatim) | What is TRUE | Citation I personally opened |
|:--|:--|:--|:--|:--|
| **G-1** | `corpus/ops/demo/coverage-protocol.md:616` — a line **this commit wrote** | "the `ai_readiness_live_snapshots` read was gated behind a *closed* `CycleID`" | The read gated behind a closed `CycleID` is on **`ai_readiness_snapshots`**, not `ai_readiness_live_snapshots`. `buildResponseFromSnapshots` calls `ListAIReadinessSnapshots(ctx, orgID, cycle.ID)` and reads `Frozen*` columns. `ai_readiness_live_snapshots` is a **write-side materialized mirror for askengine ("Talk to Data")**, upserted by `RefreshLiveSnapshots` over the **active** population, and is never read by the dashboard path on any cycle gate. The file's own twin 12 lines below (`:628`) names the right table: *"a frozen `ai_readiness_snapshots` row per member"*. In a paragraph whose operative advice is *"seeding the snapshot table"*, naming the wrong table sends a seeder at the wrong table. | `app/internal/aireadiness/readiness.go:771-772` (`buildResponseFromSnapshots` → `ListAIReadinessSnapshots`); `app/internal/aireadiness/live_snapshots.go:54-57` (*"the queryable mirror the askengine ('Talk to Data') reads … the same way `ai_readiness_snapshots` backs 'in the last closed cycle' questions"*) and `:116` (`upsertLiveSnapshots`); `app/internal/askengine/registry.go:526`; `app/internal/askengine/rules.md:746,755`. Tree-wide `grep` confirms the only non-askengine consumers are the ent schema and the worker task. |
| **G-2** | `corpus/architecture/external_services.md:569`, **propagated as canonical by this commit** to `corpus/architecture/architecture_overview.md:248-249` and `corpus/architecture/security_compliance.md:186` | "**Four** things can send a request outside the EU, none of them a region-health failover" / "enumerates **four** ways a request leaves the EU" | The set is **not exhaustive — there is at least a fifth**, and this commit is the one that rewrote the item that admits Studio-Room into the enumeration. Studio-Room's selector is the ini `TARGET SERVICE`, whose legal values are `openai, azure, anthropic`. The corpus enumerates the `anthropic` value (as exit #3) but **not** the `openai` value, which builds a first-party `OpenAI(api_key=…)` client with no base-url override, against `OPENAI_ENDPOINT = https://api.openai.com/v1/…` — a US endpoint, selected by config, with no flag and no error condition. Structurally identical to the exit the list *does* carry. This is the same undercount defect as the commit's own blocker #3 (*"the two US paths" undercounts*), one scope level up. | `app/studio/services/ai.py:704-708` (`providers = {'openai': OpenAIProvider, 'azure': AzureProvider, 'anthropic': AnthropicProvider}`), `:334` + `:383` (`return OpenAI(api_key=self.api_key)` — no `base_url`), vs `:530-533` (`AzureOpenAI(azure_endpoint=self.endpoint…)`); `app/studio/configs/config_template.ini:28` (Azure endpoint = `…swede…cognitiveservices.azure.com`, i.e. EU) vs `:30-31` (`OPENAI_ENDPOINT = https://api.openai.com/v1/chat/completions`); `development_config.ini:25` and `production_config.ini:25` (`# TARGET SERVICE: openai, azure, anthropic`). |
| **G-3** | `corpus/ops/demo/stories-spec.md:598-599` — **pre-existing, not introduced here**; surfaced by this commit's own leak mandate | "M51's iters 03→06 built and then **falsified** the active-signals path — the live-recompute never completes in the coverage harness's budget" | FALSE against three corpus twins. `seeding-spec.md:496-498` quotes *that exact sentence* as the refuted form: *"**M219 FALSIFIED M51's headline strategy claim.** … **The live recompute completes in 2.09 s.** M219 measured it."* `services/ai-readiness.md:371` (*"M219 measured the live recompute at 2.09 s and refuted that"*) and `:449-450` (*"LIVE `GET /api/workforce/ai-readiness` → HTTP 200 · **2.09 s** · 304 KB. The M51-era 'translation-N+1 that never completes in-budget' is **not reproducible**…"*) say the same, and `CLAUDE.md:324` records it. `stories-spec.md:599` is the **last unfenced survivor**, and it is the paraphrase-sibling of the very paragraph blocker #7 repaired — same M51 AI-readiness narrative, different file. The `repair_leak_guard` is verbatim-shingle-based and reports GREEN on this commit (I ran it), which is precisely the paraphrase blind spot the brief warns about. | `grep -rn "never completes\|2.09 s" corpus .claude CLAUDE.md README.md` → `seeding-spec.md:498`, `stories-spec.md:599`, `ai-readiness.md:371,449,450,508`, `CLAUDE.md:324`. `grep -n "M219\|never completes" corpus/ops/demo/stories-spec.md` → the file mentions M219 eleven times and fences nine other things, but **not this**. |

## Minors

1. **`ai_architecture.md:45-48`, `:84`, `external_services.md:575` — "Studio-Room was **never** on Bedrock"
   is TRUE, but the evidence the corpus cites cannot establish it.** The cited proof is *"`grep -rin
   'bedrock\|boto3' app/studio/` returns **0** hits"* — a HEAD-only check. This commit writes the rule
   against exactly that inference 430 lines earlier in the same file
   (`external_services.md:145-146`: *"a check against `docker-compose.yml` at HEAD can establish 'does not
   exist now'; it cannot establish 'never'"*). I ran the check that **does** establish it:
   `stack-demo/app/studio/` is a full non-shallow clone of `anthropos-studio-room` (677 commits, no
   `.git/shallow`, first commit 2024-09-18 `cf0ca3d`), and
   `git log --all -S bedrock` / `-S boto3` / `-S Bedrock` each return **0 commits**. Fix is one clause,
   not a retraction: cite the history search instead of the HEAD grep. **Not booked as a blocker only
   because the strong claim happens to be true** — the epistemic self-contradiction inside one commit is
   real and is the exact over-correction class the brief names.
2. **`ai_architecture.md:96` — "this line went on publishing it 68 lines below" no longer computes.**
   68 was right at the pre-commit numbering (the line was `:84`, the ⚠️ at `:15-17`). This commit added 12
   lines above it, moving it to `:96`; a reader counting today gets 80.
3. **`coverage-protocol.md:619` — "thirteen lines above its own retraction" reproduces under no reading.**
   Pre-commit `:614` → the ⚠️ at `:628` is 14 (13 only if you count to `:627`); post-commit `:614` → `:631`
   is 17. A newly-written derived count that is stale at the moment of writing.
4. **`coverage-protocol.md:600-609` — the M51 case study around the fenced sentence was left unqualified.**
   *"logged **ZERO completions** across the entire backend log"* and *"**0 hits = the server can't produce
   the response in-budget**"* still read as standing instruction directly above the one sentence this
   commit put into the past tense, with no pointer to `services/ai-readiness.md:449-450`'s *"not
   reproducible on the app tag the demo builds today · 2.09 s"*. Same root as G-3; fencing one sentence
   of a ten-line paragraph leaves the paragraph telling the old story.
5. **`dependency_map.md:25` and `:91`, `services/cms.md:61` — Studio-Room's provider set is listed as
   "OpenAI, Anthropic, **Mistral**".** At `app@5ba17044` `studio/services/ai.py:706-708` registers exactly
   three providers — `openai`, `azure`, `anthropic` — and `mistral` appears **only** as an unused line in
   `studio/requirements.txt:8` (0 code references; `grep -rin mistral studio/` → 1 hit, that line).
   Pre-existing and outside the diff, but it is the provider set this commit re-derived, and it also
   omits `azure` — the one Studio-Room actually uses in both dev and prod config.

## Leak check

Scope as instructed: `corpus/**`, `.claude/skills/**`, `CLAUDE.md`, `README.md`; `knowledge/**` excluded.
Mechanical fence run independently: `repair_leak_guard --range HEAD~1..HEAD` → **GREEN** (635 shingles,
112 files, 317 296 words). Every finding below G-3 came from paraphrase hunting, which that fence cannot do.

| # | Claim this commit changed | grep run | What survived |
|:--|:--|:--|:--|
| 1 | anchor `external_services.md:537` → `:545`, and `:489` → `:541` | `grep -rn "external_services\.md:[0-9]*"` over scope | **Clean.** 10 citations, all `:545` / `:569` / `:541` / `:577-587`; zero `:537`, zero `:489`. Each target line opened and confirmed to carry the named construct. |
| 2 | *"flips Course Builder / **Studio-Room** off Bedrock"* | `grep -rni "studio.room\|studio_room" … \| grep -i "bedrock\|anthropic"`; plus `grep -rn "Bedrock"` over scope (24 hits, all read) | **Clean of the conjunct** — the three repaired sites are the only Studio-Room×Bedrock sentences and all three now say *"never on Bedrock"*. But see Minor 1 (evidence too weak for "never") and Minor 5 (`dependency_map.md:25` still carries a wrong Studio-Room provider set — a *different* claim, so not a leak, but adjacent). |
| 3 | the retracted `→` ladder, *"EU Azure default → US Azure → direct-OpenAI"* | `grep -rn "Azure.*→.*OpenAI\|EU.first routing\|EU-first fallback\|fallback ladder\|fallback chain"` | **Clean in the platform-describing docs.** `shared_libraries.md:127`'s *"Azure→direct-OpenAI fallback on HTTP 429"* is the 429 leg only and is fenced by `:124` — correct, not the ladder. `ops/demo/ai-generation-spec.md:50,55,331` and `demo/README.md:273` / `CLAUDE.md:343` still say *"EU-first routing"*, but of **rext's own `services/ai/` wrapper**, a different codebase — out of the retraction's subject. Flagging for the next pass only if that wrapper is in scope. |
| 4 | *"unset `ai_vendor` falls through the `default:` arm at `:113-115`"* | `grep -rn "113-115\|:113"`; `grep -rn "ai_vendor"` (10 hits, all read) | **Clean.** No `:113-115` anywhere. `ai_architecture.md:241` independently says *"`simulator/ai/ai.go:114-115` sets `aiVendor = internalAi.Openai`"* for the **unmatched-vendor** case — consistent with the repair, not a leak. `:235`'s content-side default (`:1297` model / `:1302` vendor) agrees with the new `:50-57` text. |
| 5 | the *"has never existed"* Directus over-correction | `grep -rn "never existed\|has never been\|never was"` | **Clean.** `service_taxonomy.md:303` and `external_services.md:143` both now quote the phrase *as the refuted form*; `ops/verification.md:599` and `platform-alignment.md:538` are unrelated/ledger prose. The two correction blocks are mutually consistent (same commit `a2a3ee6`, same three line cites, same email caveat). |
| 6 | *"the two US paths"* / *"exactly three things"* undercount | `grep -rn "two US path\|two ways out\|three things\|leaves the EU\|outside the EU\|send a request"` | **The old form is gone**, and the surviving *"two US paths"* at `architecture_overview.md:247` and `security_compliance.md:183` are now correctly scoped *"inside the AI manager"*. **But the replacement count is itself short — see blocker G-2.** |
| 7 | the aireadiness *"always takes the live-recompute branch"* present-tense claim | `grep -rn "buildLiveResponse\|live-recompute\|buildResponseFromSnapshots\|ai_readiness_live_snapshots\|ai_readiness_snapshots"` (21 hits, all read) | **The exact sentence is gone.** Its **paraphrase-sibling survives unfenced at `stories-spec.md:599`** → **blocker G-3**, and the case study around the repaired sentence was left unqualified → Minor 4. |

**Over-correction sweep.** Four candidate over-corrections examined; one confirmed, none fatal:
`"never on Bedrock"` (Minor 1 — over-strong for its cited evidence, but the claim is true);
`"there is no Directus service in the platform compose"` (correctly narrowed to *at `2adcf71`* — good);
`"Both halves are false"` at `ai_architecture.md:66` (both halves genuinely are false — verified);
`"the `default:` arm … is a **separate** door"` (correct, and the passage carefully preserves the surviving
outcome claim rather than retracting it wholesale — this is the best-executed correction in the commit).
