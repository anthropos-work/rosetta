# Auditor E — 6 files / 1443 lines — 3 blockers, 15 minors (all 3 in security_compliance.md)

## B-E1 `security_compliance.md:175-176` [EDITED] — a RESIDENCY claim, false at HEAD
"'Anthropic Direct' is **not used at all** — Anthropic is reached exclusively through AWS Bedrock eu-west-1".
FALSE: `app/internal/coursebuilder/bedrock.go:108-112` routes EVERY coursebuilder model call to
`askengine.NewAnthropicClientWithModel` (api.anthropic.com, first-party) whenever `ANTHROPIC_API_KEY` is set;
`bedrock.go:43-45` says so in words; `ModelBackendName()` returns "anthropic-api"; second key at
`cms/studio/studioManager.go:1059`. The cited `:85-95` anchor resolves but covers only jobsimulation/ai.
CLASS: universal quantifier ("exclusively"/"not at all") generalised from one package to the platform.

## B-E2 `security_compliance.md:76, :83-84` [EDITED] — ***DOUBLE-FIND with auditor C*** (arch_overview:288-291)
"**16** carry an organization_id with **no policy of any kind**" + "The remainder ... carry no org column by design".
RE-MEASURED BY THE ITERATION ITSELF:
  - 7 schemas use `OrganizationIDMixin{}`: category, jobrole, similarity, skill, specialization,
    studio_document, studio_task.
  - `OrganizationIDMixin` declares **0** `Policy()`; each of the 7 declares **0** own `Policy()` and does
    NOT carry `OrganizationMixin{}`. So all 7 carry organization_id with NO policy of any kind.
  - Only FOUR files in the whole schema dir declare any Policy(): mixin.go, organization.go,
    org_membership.go, user.go.
>>> THE DOC CONTRADICTS ITSELF SEVEN LINES APART: `:69` already names OrganizationIDMixin as
>>> "a plain nullable organization_id column with **no policy**" — and `:76` then excludes those very 7
>>> from the count of unpoliced schemas. Errs toward "isolation is handled" — the DANGEROUS direction,
>>> and the FIFTH consecutive failure of this fence.
  - Base count unsettled between auditors: C measured 17+7=24, E measured 16+7=23. NOT settled here.
  - E RESOLVED the denominator ambiguity that D-M257x-39-3 rested on: 135 is correct; the rival 112 is a
    GREP ARTIFACT (`grep '^\tent.Schema$'` misses 23 gofmt one-liners `struct{ ent.Schema }`). So half of
    D-M257x-39-3's stated reason for refusing the edit is now refuted.
  - Also independently corroborates C: `organization.go` declares its own org-filtering Policy(), so "31"
    is 32.

## B-E3 `security_compliance.md:205` [EDITED] — an orphaned bullet the retraction forbids
The EU-AI-Act section is a BULLET LIST with the retraction blockquote spliced into its middle. The list
then RESUMES at :205: "This classification means transparency obligations only, not the strict
requirements of High Risk systems" — stating the operative legal consequence as settled fact, immediately
after :202 says "**Do not cite this section as evidence of a Limited-Risk classification**".
ADJUDICATED PRECISELY (E over-read :7, which DOES defer properly to counsel): the defect is not that the
corpus asserts a classification wholesale — it is the single trailing bullet, orphaned by the splice,
drawing the legal consequence FROM the classification the same section just forbade relying on.
>>> This CORRECTS `D-M257x-39-4`, which recorded that the one-way-door check PASSED and that "neither file
>>> now asserts a legal conclusion". The deferral is present; it is not exclusive.
CLASS: mechanical (blockquote spliced into a list, trailing member orphaned).

## E's notable CLEAN result: `hiring.md` — repaired TWICE and defective after both — is now CLEAN.
~40 exact anchors resolve; 0 CHECK constraints on completion_status; 87 published HIRING sims; etc.
