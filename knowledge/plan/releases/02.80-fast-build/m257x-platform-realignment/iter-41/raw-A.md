# Auditor A — 7 files / 1716 lines (read-confirmed) — 4 blockers, 11 minors

## B-A1 `ai_architecture.md:104-105` [EDITED by iter-39]
"`gpt-4o` appears in NO `*_MODEL` slot of any studio config" — FALSE. It is in two:
`app/studio/configs/config_template.ini:39` EXECUTION_AI_STABLE_MODEL and `:40` CREATIVE_AI_STABLE_MODEL.
**Self-contradicting inside the SAME blockquote**: :102-103 says the template "still carries ... gpt-4o".
Intended true claim = "no SHIPPING config". As written the universal is false. CLASS: universal-quantifier.

## B-A2 `graphql-wundergraph.md:79` [EDITED] — ***THE :5050 CLAIM, INSIDE CLAUSE-5 SCOPE***
"**Ports**: host **5050** -> container 8080" — un-fenced, present tense, in Architecture & Code Map.
No 5050 at HEAD (`grep 5050 docker-compose.yml` -> exit 1). Prod is 8080->8080 (terraform locals.tf:8).
The SAME doc at :174-176 says localhost:5050 refuses connection; `external_services.md:348` retracts it.
>>> SELF-INDICTMENT OF iter-40: my claim-scoped sweep fixed :5050 at 8 sites in ops/.claude, and my
>>> in-scope uniformity check covered only 5 of the 8 claims — :5050 was NEVER grepped inside
>>> corpus/services|architecture. iter-40's headline "the 40 in-scope files are uniform on all of them"
>>> was therefore OVERSTATED: it was verified for 5 claims, asserted for 8.

## B-A3 `ai_architecture.md:43-46` [EDITED] — CROSS-FILE DRIFT OF iter-39's OWN C2 FIX
Says the "fourth exit" is an **unrecognised** `ai_vendor` string. The real fourth exit is an **UNSET**
vendor: nullable `AIVendor *AIVendor` (cms/directus/collections/jobsimulation.go:905) defaulted to
`openai` at :1302 -> direct US OpenAI on the FIRST attempt, no typo needed. `external_services.md:559-584`
has this right; ai_architecture.md restates the narrower claim. Also internally contradicted by its own :207.
CLASS: the exact cross-file drift class rule 19 was authored for, one pass later.

## B-A4 `external_services.md:788` [EDITED]
"Consistent with **:447** above" -> :447 is a table HEADER row. The correction is at :512. 65 lines off.

## MINORS: 11
## UNVERIFIABLE: gh/archive; org repo inventory (5 LiveKit agent repos); prod terraform outside cloned
## repos; private Go modules; prod-side infra (S3/CDN, prod Directus, Chime/LiveKit buckets).
## CONFIRMED CLEAN: whole supergraph ladder 5/4/3/1 + "3->1 not 2->1"; bba862f unmerged (rc=1);
## no `type Subscription` in backend.graphqls; taxonomy 42790/22470 live + manifest 18919; all 11
## academy_* names plural-verified live; TEMPLATE.md pure scaffold 0/0.
