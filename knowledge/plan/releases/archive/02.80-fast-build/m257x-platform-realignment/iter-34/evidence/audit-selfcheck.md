# iter-34 adversarial self-check of the iter-33/34 repair pass

Audited: the 11 claims written into 8 corpus files by the pending repair
(`git diff -- corpus/`), against platform `stack-demo/` (`platform` @ `2adcf71`, `app` @ `5ba17044`,
`next-web-app` @ `bb3313bc0`, `ant-academy` @ `9c3843cd`) and the rext authoring clone
`.agentspace/rosetta-extensions` @ `b2b46cb`.

**Verdict: 2 BLOCKERS · 8 minor.**

Every `file:line` anchor the repair wrote verified exact (again). Both BLOCKERS are in the *prose
around* an anchor — one is a parenthetical the repair newly authored, one is a sibling paragraph the
repair's own correction made false and did not sweep.

---

## BLOCKER 1 — `corpus/architecture/architecture_overview.md:246-248`

> "**Mistral is not in `app`'s routing at all** (its only platform reference is `terraform/ssm.tf:291`,
> for the cms studio markdown manager)"

The lead clause is true (Mistral is in neither `AIVendor` enum nor any `getClient`). **The parenthetical
is false**, and it is the load-bearing half: it tells a reader Mistral is a dangling SSM parameter with
no consumer. Mistral is **live Go code inside `app`**:

- `app/internal/cms/studio/markdownManager.go:11` — `import "github.com/anthropos-work/ai/mistral"`
- `app/internal/cms/studio/markdownManager.go:19` — `mistral.NewMistral(nil, os.Getenv("MISTRAL_API_KEY"))`
- `app/internal/cms/studio/studioManager.go:583` — `NewMarkdownManager(os.Getenv("MISTRAL_API_KEY"))`,
  the OCR path for uploaded Studio documents (`OCRProcess`, `Tokenize`, `CleanMarkdown`)
- also `app/internal/cms/studio/xlsx.go:13`, `app/go.mod:127`, `app/terraform/variables.tf:574`,
  `app/terraform/main.tf:535` (task-def secret), `platform/.env_example:38-39` (`MISTRAL_API_KEY`)

Two further problems in the same edit:

1. **Repo attribution.** `platform/` has no `terraform/` directory at all — `terraform/ssm.tf` is in
   `app`. The path is unqualified in a bullet whose previous citation was `app/internal/...`, so it
   reads as `app/terraform/ssm.tf`, which is right by accident.
2. **Self-contradiction inside the same file.** The repair removed Mistral from this bullet's provider
   list while `architecture_overview.md:36` still reads *"AI Providers: OpenAI, Anthropic, **Mistral**
   (EU-first routing)"* and `:200` still reads *"`ai` … (OpenAI, Azure, Anthropic, Bedrock, **Mistral**)"*.
   Those two lines are the accurate ones; the new bullet is the wrong one.

**Why it misdirects real work:** this bullet is the corpus's top-level AI-provider inventory, and it is
what a secrets/compliance/dependency sweep reads. Acting on it, an engineer would drop `MISTRAL_API_KEY`
from a stack `.env` (breaking Studio document OCR inside `backend`) or omit Mistral from an EU
sub-processor inventory — while Mistral is in fact processing *uploaded customer documents*.

**Fix:** replace the parenthetical with — *Mistral is absent from `app`'s AI **routing**, but is a live
dependency of the cms Studio OCR/markdown path (`app/internal/cms/studio/markdownManager.go:11,19`;
`studioManager.go:583`), provisioned via `app/terraform` SSM/task-def and `platform/.env_example:38-39`.*

---

## BLOCKER 2 — `corpus/architecture/security_compliance.md:154-155`

> "- AI providers are routed through EU endpoints first (Azure OpenAI EU, AWS Bedrock EU, Mistral EU)
> - US providers (OpenAI Direct, Anthropic Direct) used only as fallback
> - No customer data stored in US by default"

False at HEAD, and **left standing by a repair pass that edited this very file** while adding, one file
over, the warning *"⚠️ 'EU-first' is not 'EU-only' — one feature flag routes traffic to the US."*

- `app/internal/jobsimulation/ai/ai.go:263-277` — `getClient` returns **`azureClientUs`** whenever the
  PostHog flag `flag_use_azure_us` evaluates true. That is a **flag switch, not a fallback**: no EU
  failure is required, and Azure US does not appear anywhere in this compliance list.
- The same flag is re-implemented at `ai.go:341-352` (`AudioTranscriptions`) and
  `app/internal/skillerai/ai.go:347`.
- "Anthropic Direct" is never used: the Anthropic client is **AWS Bedrock pinned `eu-west-1`**
  (`ai.go:85-95`); the only 429 override target is `Openai` (`ai.go:152-155`, `:299-301`).

**Direction: fails toward reassurance**, on a GDPR data-residency claim, sitting directly above
"No customer data stored in US by default". This is precisely the failure mode the repair's own new
text at `security_compliance.md:82-84` names ("wrong three times, each time in the same direction").

**Fix:** state the Azure-EU→Azure-US flag path explicitly in §EU Data Residency, drop "Anthropic Direct",
and cross-link `architecture_overview.md:242-248`.

---

## Minor findings

**m1 — `architecture_overview.md:245`, wrong anchor.** "direct OpenAI is the fallback on HTTP 429
(`:127-137`)". `ai.go:127-141` is `isThrottlingError`, the 429 **detector**. The OpenAI override is
`ai.go:152-155` (`if throttled { vendor = Openai }`), armed at `:166-168`; the Response path is
`:299-301` / `:325-327`. Claim true, citation points at the wrong half of the mechanism.

**m2 — `architecture_overview.md:284-285` vs `security_compliance.md:93`, cross-file contradiction the
repair created.** Both files were edited in this pass. security_compliance now says *"Ent privacy
policies auto-filter by organization on **31** schemas"*; architecture_overview — in the bullet that
links straight to that section — still says *"auto-filter **only the 30** schemas using
`OrganizationMixin{}`"*. The word "only" is now false.

**m3 — `security_compliance.md:72-78`, the 17-name list is off by one (toward alarm).**
I regenerated the list independently and confirm the arithmetic: 139 `.go` files / 135 declaring
`ent.Schema` (the 4 non-schema files are `database_types.go`, `mixin.go`, `skillpath_mixins.go`,
`skiller_mixins.go`); 30 `OrganizationMixin{}`; 7 `OrganizationIDMixin{}`; 18 with a plain
`organization_id` and neither mixin; and `org_membership.go` is indeed the only one of the 18 declaring
its own `Policy()` (`:172-188`, ending `privacy.AlwaysDenyRule()` on both Mutation and Query) —
all 17 listed filenames match mine exactly.

But **`academy_feedback.go` is policed**: it carries `UserMixin{}`, which declares a row-level
`Policy()` at `mixin.go:99` (`DenyIfNoUserInContext` → `FilterOwnerRule` → `AlwaysAllow`). It is
owner-scoped, not org-scoped, so it is not *un*policed. Truly-unpoliced is **16**, plus one
user-policed. Errs toward alarm, not reassurance — hence minor.

The **re-derivation recipe at `:86-87` reproduces the same defect** and adds one of its own:
`grep -L 'Organization\(ID\)\?Mixin{}' schema/*.go | xargs grep -l '"organization_id"'` returns **19**
files, not 18 — it also matches `skiller_mixins.go`, which is the *definition* of `OrganizationIDMixin`,
not a schema. And "then subtract any that declare their own `Policy()`" misses mixin-carried policies,
which is exactly how `academy_feedback.go` got onto the list. Recommend: `… | xargs grep -l 'ent.Schema'`
first, then subtract both own-`Policy()` and `UserMixin{}`/`OrganizationMixin{}` carriers.

**m4 — `security_compliance.md:79-80`.** "The remainder (the taxonomy, and other global reference data)
carry no org column by design." `skill_path_session.go` carries a plain nullable **`tenant_id`** with no
org mixin (`skillpath_mixins.go:24-26` documents the choice); it is policed by `UserMixin`, not by org.
The sentence's "no org column" is imprecise for that table.

**m5 — `backend.md:35`, "appears **0 times** in the repo" is literally false.**
`SkillPathSessionService` appears **twice** in `app` — `app/CLAUDE.md:72` and
`app/knowledge/architecture.md:28`, both stale platform-side docs that still list it on the mux. The
*operative* claim is correct and verified: `main.go:1178-1218` registers exactly the six named handlers,
no skillpath handler, and no `skillpath…v1connect` package is imported anywhere in `app` (only
`proto/go/skillpath/v1` and `proto/go/domain/skillpath/v1/skillpathsession` type packages).
Worth rewording to "0 code references — the only two hits are `app`'s own stale CLAUDE.md/knowledge docs."

**m6 — `backend.md:35`, "the mux registers **exactly SIX** handlers."** `CMSService` is registered
**conditionally** (`main.go:1203`: `if cmsRPCServer != nil { mux.Handle(...) }` — the cms managers are
only built when the Directus edge is configured). A stack without `DIRECTUS_BASE_ADDR`-driven cms wiring
registers five. "Six" should be "up to six (CMSService is conditional)".

**m7 — `hiring.md:125`, splicing damage in the read-path table.** The repair rewrote row 7 to say
*"**Not a mirror**: `local_jobsimulation_session.go` no longer exists"* and to cite `intelligence.go:1820`
/ `:1846`. Row 6 immediately above it still reads
`| 6 | intelligence.go:1801 | Score ← ls.Score (**the mirror's score column**) |`.
Both halves are wrong at HEAD: `:1801` is `organizationAssignmentSessionMap := make(...)`; the score is
computed at `:1820` (`score := RoundFloat(float64(ls.Score), 0)`) and assigned at `:1846`; and there is
no mirror. One row of a 7-row table was corrected and its neighbour restates the retracted claim.

**m8 — `ai-readiness.md:400`, stale filename directly above the new insert.** The bullet cites
`computeOrgBreakdowns` at **`ai_readiness.go:283-343`**. The file is
`app/internal/aireadiness/readiness.go` (the insert two lines below uses the correct path), and
`computeOrgBreakdowns` is declared at `readiness.go:330`. `ai_readiness.go` does not exist at HEAD.

---

## Cleared — verified exact, no finding

- **Item 1, `external_services.md:199-213`.** `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` at
  `gen_injected_override.py:53` ✓; emission `if with_directus and name in DIRECTUS_DATA_CONSUMERS:` /
  `env.append(...)` at `:598-599` ✓; `test_only_cms_is_repointed_not_other_services` has **0** live
  occurrences (only a back-reference in the replacement's comment) and
  `test_backend_the_actual_reader_is_repointed` is at `tests/test_injection.py:1005` ✓. Consistent with
  the `cms_reader_switch` paragraph at `:191-197`. No orphan found in the Directus section
  (all 60 `directus` mentions in the file reviewed).
  *Note, not a doc defect:* the rext **code comment** at `gen_injected_override.py:590-597` still reads
  *"re-point the data-plane consumer (cms) … cms is the only direct Directus consumer"* — stale against
  the tuple two lines below it. Worth fixing in rext, not in the corpus.
- **Item 2a, `ai-readiness.md:414-422`.** `GetAIReadinessWithOptions` at `readiness.go:289` ✓;
  cycle-scoped closed-cycle branch `:291-297` (doc says `:290-297`, i.e. includes the comment line) ✓;
  the default branch `if m.activeCycle(...) == nil { if closed := m.latestClosedCycle(...) ... }` at
  `:309-312` (doc says `:307-312`, incl. its 2 comment lines) ✓; `buildLiveResponse` fall-through at
  `:314` ✓. The correction is right and the retracted claim was genuinely false.
- **Item 2b, `ai-readiness.md:67-77`.** `AIReadinessClient.tsx` — grep for `posthog|featureFlag|
  AI_READINESS_FLAG` returns **zero** hits ✓; `const { orgEnabled } = useAiReadinessEnabled(true)` at
  `:133`, `const featureOn = orgEnabled === true` at `:134` ✓; `useAiReadinessEnabled`
  (`packages/graphql/src/hooks/aiReadiness/useAiReadinessEnabled.tsx`) does **not** consult PostHog
  itself ✓; the only `useFeatureFlagEnabled(AI_READINESS_FLAG)` call site repo-wide is
  `useAiReadinessActive.ts:22`, constant at `aiReadiness.constants.ts:26` ✓; that hook's consumers are
  `/onboarding` and `/reimport-profile` — member surfaces ✓. No contradiction with `:110-115`.
- **Item 4, `hiring.md:144-147, 160-163`.** `job_simulation_session.go` declares **no** anticheat field ✓;
  `anticheat_summary` exists in exactly one place in the entire migration set —
  `20250416091037.sql:5` on `local_jobsimulation_sessions` ✓ — and that table is dropped at
  `20260729133514.sql:62` ✓; `AnticheatResult.summary` is `field.Enum("summary")` at
  `ent/schema/anticheat_result.go:24` ✓; the FK re-point is `20260722104506.sql:53` ✓; the read path is
  a separate `anticheatSummariesBySession` query at `intelligence.go:1796` ✓. "Cosmo" correctly removed
  from `:265-268`. (Except m7 above.)
- **Item 6, `skillpath.md:68-86`.** `InsightsSkillPathByMemberships` at `intelligence.go:1144` ✓;
  `m.ent.SkillPathSession.Query()` at `:1159` with `SkillPathID` + `StatusIn(Active, Completed)` +
  `HasUserWith` + the `TenantIDIsNil() OR TenantID(org)` predicate through `:1169` ✓;
  `DROP TABLE "local_jobsimulation_sessions"` `:62` and `DROP TABLE "local_skill_path_sessions"` `:63` ✓;
  `20260729133514.sql` **is** the newest migration in `terraform/migrations/` ✓; no `local_*` Ent schema
  exists ✓. Related-Documentation link text was updated in step — no orphan.
- **Item 8, `clerk-integration.md:103`.** All four `@clerk/nextjs` pins are `^6.39.2`
  (`apps/web:10`, `apps/hiring:10`, `apps/integration:9`, `ant-academy/code/package.json:52`) ✓;
  `next-web-app/apps/mobile/package.json:6` = `~2.6.18`, `ant-academy/mobile/package.json:18` =
  `~2.19.36` ✓; "thirteen minor versions" (19−6) ✓; `app/go.mod:31` = `clerk-sdk-go/v2 v2.7.0` @
  `5ba17044` ✓. The table at `:96-101` is unaffected. *(Non-finding, FYI: `@clerk/types` is misaligned
  the same way — `^4.101.20` in web/hiring/integration vs `^4.60.0` in `apps/mobile:50`. The doc's
  narrower claim is not wrong, just not exhaustive.)*

---

## Positive control — coverage per file

| File | `wc -l` | Read |
|---|---:|---|
| `corpus/services/skillpath.md` | 103 | **full**, 1→103 |
| `corpus/services/backend.md` | 257 | **full**, 1→257 |
| `corpus/architecture/security_compliance.md` | 186 | **full**, 1→186 |
| `corpus/services/hiring.md` | 327 | **full**, 1→327 |
| `corpus/services/clerk-integration.md` | 128 | **full**, 1→127 |
| `corpus/architecture/architecture_overview.md` | 332 | **partial** — 25→69 + 200→331, plus a whole-file grep for `EU-first / EU-only / mistral / bedrock / azure / Mixin`. 1→24 and 70→199 not read line-by-line |
| `corpus/architecture/external_services.md` | 709 | **partial** — 126→153, 155→214, 230→259, 578→607, plus every one of the 60 `directus`-matching lines. Remainder not read line-by-line |
| `corpus/services/ai-readiness.md` | 613 | **partial** — 40→169, 385→474, plus a whole-file grep for `hardcoded / buildLiveResponse / CycleID`. Remainder not read line-by-line |

The three partials were scoped to the changed regions + a term sweep for every topic the corrections
touch. A full top-to-bottom read of those three (1,654 lines) is the residual gap in this audit.
