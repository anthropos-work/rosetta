# AUDITOR B — 7 files / 1559 lines

**Positive control:** all 7 read to final line; counts match `wc -l`
(ai-readiness 623 · ai_architecture 280 · security_compliance 259 · ai-labs 156 · clerk-integration 128 ·
customerio-sync 75 · architecture/README 38).

## BLOCKERS — 2 (both in text iter-46 wrote)

| # | site | the false claim | what is true |
|---|---|---|---|
| **B1** | `ai_architecture.md:51-54` | *"`ai_vendor` **unset or unrecognised** falls through the `default:` arm at `:113-115`. `AIVendor` is a **nullable** pointer on the sequence, so *unset* — not merely *mistyped* — is the ordinary way to reach it… An earlier revision named only the 'unrecognised string' case, which reads as a misconfiguration when it is in fact the default."* | **Both conjuncts fail.** (a) `simulation.Sequence.AIVendor` is a **value**, not a pointer — `grep -rn "simulation.Sequence{" --include="*.go" internal/` returns exactly one construction, `jobsimulation.go:1307`, setting `AIVendor: aiVendor` where `aiVendor := simulation.Openai` (`:1302`). The nullable pointer is on a **different** struct, the Directus DTO `collections.Sequence` (`:905 AIVendor *AIVendor`). (b) Therefore **unset never reaches the `default:` arm** — `:1302-1304` normalizes nil → `simulation.Openai`, which hits `case simulation.Openai:` at `simulator/ai/ai.go:58-59`. The `default:` arm is reached **only** by an unrecognised value — e.g. `azureglobal`, a real 5th enum member (`jobsimulation.go:971`) with no case — i.e. **exactly the "misconfiguration" reading this paragraph was written to retract.** Self-contradicts `:223` of the same file, which correctly states the unset content-side default is `openai` via `:1302`. The *outcome* claim (direct US OpenAI, first attempt, no error condition) survives; the mechanism and the "nullable pointer" premise do not |
| **B2** | `security_compliance.md:197-198` | *"`external_services.md:489` carries the provider row"* | `sed -n '489p'` → `// Types in app/__generated__/` — a Studio-Desk GraphQL-codegen comment inside a TypeScript fence. The **Anthropic Direct (first-party API)** row is at `external_services.md:533`. (`coursebuilder.md:48` *does* resolve and does say "the shipped path" ✓; the substantive compliance claim is independently correct — `coursebuilder/bedrock.go:109-113` verified) |

## MINORS — 11

| # | site | what is off |
|---|---|---|
| 1 | ai_architecture.md:23 | `getClient` "(`:259-289`)" — the function spans `259-288` |
| 2 | ai_architecture.md:210 | the `AIModel` enum block runs `977-991`; the cited `983-990` starts mid-enum and omits `anthropic-35/37/4-sonnet-aws` |
| 3 | ai_architecture.md:245 · security_compliance.md:216 | "the hardcoded switch at `criterion.go:127`" — the switch is `:125`, the LLM case `:126`; `:127` is `var p check.ParamsLLM`. The `validateLLM` call site is `:175`. (`:428`, `:168`, `:450-475`, the tmpl, temperature 0.0 and the `{check_id,feedback,success}` shape all verified exactly) |
| 4 | ai_architecture.md:84 | `:267,344` are both the bare literal `"flag_use_azure_us",`; the selection switch is `:259`, the 429 override `:154`/`:300` |
| 5 | security_compliance.md:73 | `UserMixin{}`'s `Policy()` is `mixin.go:99`; `:98` is its doc comment |
| 6 | ai-readiness.md:316 | `SHOW_SECONDARY_TABS` is at `:78` (typed `: boolean`), not `:69` |
| 7 | ai-readiness.md:438-440 | two anchors shifted, one naming the other's construct. **Both behavioural claims verified true** |
| 8 | ai-readiness.md:555 | `computeCycleTotals` is at `:260`. The query anchor `:285-287` is exact |
| 9 | ai-readiness.md:134/:481 vs :567 | bare basename `useAIReadiness.ts` denotes **two different files** (`components/ai-readiness/` vs `hooks/`). Each resolves, but only in one of the two |
| 10 | customerio-sync.md:19-33 | compose snippet omits `env_file:` and `networks:` |
| 11 | architecture/README.md:19 | bills `external_services.md` as covering "the WunderGraph Cosmo GraphQL gateway"; at `2adcf71` compose no longer defines a router service. Not false (the target doc does cover it) but the last un-fenced pointer to a removed local service in this set |

## Files read clean

- **`ai-labs.md` — 0 findings.** Every constant, route, webhook event, tier, interval, port and migration
  filename verified. Notable: the corpus is **right** where the platform's own header comment at
  `credits/cost.go:29` is **wrong** ("refine → 5"; actual `refine:1`).
- **`clerk-integration.md` — 0 findings.** Every SDK pin verified across app + four frontends + two mobile
  apps; `authn` absent from every checked-out `go.mod`/`go.sum` ✓.
- `customerio-sync.md` · `architecture/README.md` — clean apart from minors 10/11.
- **`security_compliance.md`** — the whole Layer-1 fence **re-measures exactly**: 139 `.go` / 135
  `ent.Schema` / 30 `OrganizationMixin{}` / 7 `OrganizationIDMixin{}` / **only 4 files declare any
  `Policy()`**; the doc's own `comm`-based derivation reproduces **18**, and the named 16 are precisely
  those 18 minus `org_membership.go` and `academy_feedback.go`. **iter-46's re-count is confirmed
  independently.**
- **`ai_architecture.md`** — outside B1 and minors 1-4, everything verified, including the studio slot
  table (`config_template.ini:39-40` = the two `gpt-4o` slots — **iter-46's #1 repair confirmed**).
- **`ai-readiness.md`** — outside minors 6-9, every one of ~40 platform and next-web anchors verified,
  plus the rext-side M254 re-anchor manifest — **iter-46's #14 repair confirmed**.

## Context volunteered for the other seats

`docker-compose.yml` @ `2adcf71` still defines `jobsimulation` (`:83`), `cms` (`:144`) and `roadrunner`
(`:281`), each carrying the default `graphql` profile, while the `graphql`/Cosmo-router service is **gone**.
