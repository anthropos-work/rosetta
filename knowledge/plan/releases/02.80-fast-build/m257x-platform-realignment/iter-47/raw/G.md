# AUDITOR G — adversarial diff read of iter-46's own repair

**Scope:** 30 hunks / 17 files — `git diff 29eb414..301d61a -- corpus/ CLAUDE.md .claude/`.
All 17 files' claims re-derived independently. Platform pin confirmed `2adcf71`, `app` `5ba17044`.

## BLOCKERS — 5

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| **G1** *(≡ B1)* | ai_architecture.md:51-54 | *"`ai_vendor` unset or unrecognised falls through the `default:` arm… `AIVendor` is a **nullable** pointer on the sequence"* | Both halves wrong. `simulation.Sequence.AIVendor` is a **value**; the nullable pointer is on the Directus DTO (`jobsimulation.go:905`), one layer up. `:1302-1305` normalizes nil → `simulation.Openai` **before** the sole construction site `:1307`, so unset lands on `case simulation.Openai:` — the doc's own item 2. Contradicts `external_services.md:532`, **the very site this paragraph calls "the full per-line derivation"** | new text over-corrects; anchor names wrong construct |
| **G2** | ai_architecture.md:42-43 | *"it flips Course Builder and **Studio-Room** off Bedrock onto Anthropic's first-party API"* | **Studio-Room was never on Bedrock.** `grep -rin 'bedrock\|boto3' app/studio/` → **0 hits**. Three providers only: `ai.py:334 OpenAIProvider`, `:490 AzureProvider`, `:627 AnthropicProvider`. And the key is a **credential, not the selector** — the selector is the ini's `TARGET SERVICE`, as `external_services.md:533` states correctly | new text over-corrects (rewrote the line, preserved a false conjunct) |
| **G3** | external_services.md:139 | *"**all of that is false**; that service **has never existed** in the platform compose"* | **The exact over-correction iter-46 repaired at `service_taxonomy.md:296-303` — left standing at the twin.** `git show a2a3ee6^:docker-compose.yml` → `:383 directus:`, `:384 image: directus/directus:10.10.1`, `:386 - 8055:8055`, `:409 ADMIN_PASSWORD=password`. Also: verifying against HEAD cannot establish "never existed" | repaired at one site, standing at another |
| **G4** | ai_architecture.md:84 | *"EU Azure default → US Azure via the PostHog flag `flag_use_azure_us` → direct-OpenAI on HTTP 429"* | The ordered-arrow chain iter-46 rewrote at `CLAUDE.md:246`, `architecture_overview.md:244-247` and `security_compliance.md:181-185` — surviving **68 lines below this same file's own fence** at `:15-17` (*"no such ladder exists in the code"*). The mechanics are independent: `ai.go:263-277` (flag) vs `:129/:166/:325` (429) | repaired at one site, standing at another (**same file**) + internal contradiction |
| **G5** | ops/demo/coverage-protocol.md:614-616 | *"the default AI-readiness dashboard GET **always** takes the live-recompute branch … so seeding the snapshot table would NOT short-circuit the default call"* | Refuted by `app/internal/aireadiness/readiness.go:309-312`: on the **nil-CycleID** path, no active cycle + a latest closed one → `buildResponseFromSnapshots`. Exactly what iter-46 wrote **13 lines below** at `:627-632` — but its fence (*"read the paragraph **below** as the contemporaneous iter-07 finding"*) does not cover this bullet, which stands in present tense and carries a routing consequence | repaired at one site, standing at another (**same file**) |

## SURVIVING OLD FORMS — the leak check

| claim iter-46 repaired | still publishing the old form |
|---|---|
| the EU-first ordered ladder | **`ai_architecture.md:84`** (G4). The `ai-generation-spec.md` / `demo/README.md` / `CLAUDE.md:343` hits are about **rext's own** wrapper — verified genuine at `stack-seeding/services/ai/ai.go:120-152`. **Not a leak** |
| "Anthropic Direct is not used at all" | clean |
| the Directus "has never existed" over-correction | **`external_services.md:139`** (G3) |
| Studio-Room conflated with Bedrock | **`ai_architecture.md:100`** — same conflation, unchanged context row |
| aireadiness live-path "hardcoded" | **`coverage-protocol.md:614-616`** (G5) |
| Studio-Desk "React" | clean |
| Tier-2 "not in main docker-compose" | clean — remaining hits are Ant Academy, which genuinely is not in compose |
| sentinel Go 1.25 → 1.26 | clean, and **not** a partial repair: `messenger`/`storage`/`roadrunner` really are `go 1.25.0`; only `app`/`cms`/`jobsimulation`/`sentinel` are 1.26 |
| `app/main.go:604` attached to all four domains | clean |
| jobsimulation listed as removed | clean |
| D-07 re-anchor "outstanding" | clean |
| host-`5050` router mapping | clean — `grep -c 5050 docker-compose.yml` = 0 |
| `assignments.go:815` → `:828` | clean |

## MINORS — 5

1. **`service_taxonomy.md:150-153`** — the new Technology cell spans **four physical lines** inside a GFM
   table (`:151`/`:152` contain no `|`). **The row will not render as one row.**
2. `graphql-wundergraph.md:81` — cites `:174-176`; the sentence is at `:178`.
3. `security_compliance.md:266` — *"**Both** bullets above"* now governs **three** bullets after the
   consequence bullet was moved up.
4. `ai-readiness.md:44` — cites `:458`; the past-tense re-point is at `:455`.
5. `readiness.go:308-312` vs `:307-312`/`:314` — 1-line drift across two docs.

## Verdict on iter-46's repair

**Net positive, but the induction pattern held: 5 of 5 blockers trace to the repair itself — 2 newly
written over-corrections, 3 half-applied edits.**

What was verified **sound** is the larger half: **every re-derived number is exact.** All four `app/main.go`
wiring anchors; the whole Ent-schema count chain (139 / 30 / 7 / 18 / only-4-`Policy()`, and both 18−2=16
and 16+7=23 check, with the two subtractions correctly justified); sentinel Go **1.26** and **256/128**; the
Directus un-over-correction exact to the line, with its one hedge right (`git log --all -S"admin@example.com"`
returns nothing); `gpt-4o` in a `*_MODEL` slot **only** at `config_template.ini:39,40`; studio-desk 0
react/0 tsx; and the `8080 → 8080` port change corroborated independently.

**The failure mode is narrow and mechanical: iter-46 re-derived NUMBERS rigorously and reasoned about
MECHANISMS loosely.** G1 and G2 are both in the *same newly-written prose block*, where the author extended
a correction past what the cited derivation site supported — and in both cases `external_services.md`, the
source it names, already had it right. G3–G5 are pure leaks, findable by grepping the old string.

**Recommended disposition:** fix G3/G4/G5 mechanically (one line each); **rewrite `ai_architecture.md:42-56`
against `external_services.md:532-533` rather than de novo.** That block is the third consecutive iteration
in which the paragraph explaining a correction became the next iteration's finding.
