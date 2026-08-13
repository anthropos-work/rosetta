# iter-46 — decisions

## `D-M257x-46-1` — every "what is true" re-derived from platform source, never from the ledger's summary

TOK-02's diagnosis is that the repair method induces ~50% of what it fixes, and iter-41 measured the
mechanism. So no claim in this pass was rewritten from the blocker-ledger's prose. Each was re-measured
against `stack-demo/` at `app` @ `5ba17044`:

| blocker | re-derived from |
|---|---|
| #1 | `app/studio/configs/config_template.ini:39-40` — `gpt-4o` IS in two `*_MODEL` slots |
| #2, #3 | `app/internal/jobsimulation/simulator/ai/ai.go:56-59` (`case simulation.Openai:`) and `:113-115` (the `default:` arm) |
| #4 | `app/internal/coursebuilder/bedrock.go:100`, `:109-112` — `ModelBackendName() == "anthropic-api"` |
| #5 | the whole `internal/data/ent/schema/` directory, counted (below) |
| #7 | `git show a2a3ee6^:docker-compose.yml` — `:384`, `:386`, `:409` |
| #8, #9 | `studio-desk/package.json` (0 react), `platform/docker-compose.yml:311`, `:342` |
| #10, #11 | `sentinel/go.mod:3`, `sentinel/terraform/locals.tf:4-5` |
| #12 | `grep -c 5050 platform/docker-compose.yml` → **0** |
| #14 | `demo-stack/patches/app-aireadiness-snapshot-loadmembers.yaml:42`, `:33` |
| #15 | `platform/repos.yml:17`, `platform/docker-compose.yml:83` |
| #16 | `messenger/internal/flow/assignments.go:827-829` (`getSkillPath`) |
| #17 | `app/main.go:573` / `:604` / `:634` / `:1034` — four call sites, not one |

Two turned out to be *stronger* than the ledger recorded, and both changed the wording that shipped:
`#2`'s fourth exit is the **`default:` arm reached by a nullable-and-unset `AIVendor`** — the ordinary
path, not a misconfiguration — and there is a **fifth** route nobody counted, a caller explicitly
selecting `Openai`. `#4`'s `ANTHROPIC_API_KEY` path is selected by an **env var, not a flag**, so it is
not covered by the `flag_use_azure_us` caveat that sits four lines below it.

## `D-M257x-46-2` — `#5`'s base count settled by measurement, and the disagreement resolved rather than split

iter-41 left this **deliberately unsettled**: auditor C read 17+7=24, auditor E read 16+7=23. Re-measured
at `app` @ `5ba17044`:

```
139 .go files in internal/data/ent/schema/
 30 OrganizationMixin{}          → auto-filter by organization
  7 OrganizationIDMixin{}        → org column, no Policy()
 18 plain organization_id, neither mixin   (a 19th hit, skiller_mixins.go, is a mixin definition)
  4 files declare any Policy() at all: organization.go, mixin.go, user.go, org_membership.go
```

**E is right: 16.** Of the 18 plain-column schemas, `org_membership.go` polices itself and
`academy_feedback.go` is owner-filtered by `UserMixin`, leaving 16 — and the total unpoliced
org-carrying count is **16 + 7 = 23**, which is what now ships.

The defect was never the 16. It was the closing sentence *"the remainder … carry no org column by
design"*, which excluded the 7 **three lines after naming them as unpoliced** — a contradiction inside a
single blockquote, erring toward *"isolation is handled"*. That is the **fifth** consecutive failure of
this fence, and the error direction has now failed both ways, so the paragraph says so.

## `D-M257x-46-3` — the fence found three sites the repair left standing, which is the whole point

After the first repair pass all 18 anchored sites were fixed and `claim_twin_guard` still reported
**three**:

| site | what survived |
|---|---|
| `security_compliance.md:76` | the correction was **appended** while the original undercount sentence stayed |
| `security_compliance.md:183` | the EU Data Residency section restated the retracted claim; only `:7` had been fixed |
| `platform-migration-status.md:60` | the four-domain sentence still carried the shared-anchor phrasing |

Every one is *"repaired at one site, left standing at another"* — **the class iter-41 measured as 8 of
the 9 repair-induced blockers**. Under the previous method all three would have been committed and
counted as induced defects by the next full read. Here they were named before the commit and closed in
the same pass. **This is the first pass in the series where that class was caught by a machine rather
than by the next audit.**

## `D-M257x-46-4` — two live fence findings outside the 18 were repaired, not baselined

`markdown_structure_guard` flagged a stray unterminated ``` at the end of
`.claude/skills/stack-update/reference.md`, and `anchor_construct_guard` flagged `cms.md:171` citing
`:64`, a blank line (the `public`-schema statement it means is at `:27`). Neither is one of iter-41's 18.
Both were **fixed**, because the ratchet permits carrying them and a fence whose baseline accumulates the
findings nobody wanted to fix stops being a fence. The baseline now records **0 sites across all four**.

## `D-M257x-46-5` — GREEN was verified to mean "the corpus is clean", not "the fences broke"

§5 rule 8's exact failure mode, and the one this iteration is most exposed to: four fences went from 25
sites to 0 in a single pass. Three independent checks:

- **Reach did not shrink.** `anchor_construct_guard` resolves **101** anchors across 112 files (up, since
  repaired anchors now resolve); `derived_value_guard` measures **5** service docs; `markdown_structure_guard`
  scans **112**.
- **Both perishable fixtures still go RED.** `tests/fixtures/claim_twin/` (18 sites) and
  `tests/fixtures/mechanical/` (5 sites) — 53 tests, all passing, with the green twins still silent.
- **`stack-core` 491 tests / 14 failures**, exactly the pre-existing baseline.

Had the fences been green because they had stopped reaching the corpus, the fixtures would have gone
green with them. That is why iter-43 and iter-45 spent an iteration each capturing them.

## `D-M257x-46-6` — `#17` repaired by hand, as `D-M257x-45-3` said it would be

`anchor_construct_guard` cannot reach `#17`: `app/main.go:604` **resolves** and **names a construct** —
it is simply the wrong one for three of the four domains the sentence attaches it to. Catching that
requires deciding what a sentence claims. Repaired by hand, with all four call sites cited individually,
exactly as iter-45 routed it rather than tuning a fence until the answer key fit.

## `D-M257x-46-7` — clause 5 is NOT graded here

Four fences GREEN is **not** a clause-5 measurement, and reading it as one would repeat the mistake
iter-38 and iter-21 both paid for: a cheaper instrument returns an uncomparable number. Only TOK-02
step 5 — one full 7-auditor read at iter-41's frozen instrument — grades clause 5. The gate stays
**4 of 5** until that reading returns.

TOK-02's pre-registered prediction stands unmodified and testable: **the step-5 reading returns fewer
than 9.**

## `D-M257x-46-8` — the rext pin is still NOT moved

Unchanged from `D-M257x-45-9`: this iteration edits corpus prose and lowers a baseline JSON. No runtime
source changed, so the pin stays at `fast-build-m257x-iter-37` and clauses 1/2/4 are undisturbed.
