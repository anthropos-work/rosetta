# iter-53 — the UNION of readings #11 and #12

**This file is the perishable answer key** (§5 rule 21) for the tree at `0e35b1a`, together with the 14 raw
seat reports in `iter-53/raw/`. **Do not repair from it until iter-54 opens**, and do not consume the raw
reports as a fixture for anything else.

`n₁ = 32` (reading #11) · `n₂ = 26` (reading #12) · **matched `m = 12`** · **union = 46**

**Adjudication rule, stated before the matching was done and applied uniformly:** two findings MATCH iff they
assert **the same defect about the same passage** — same anchor *and* same predicate. Two judgment calls fell
out of it and are recorded so they can be re-adjudicated rather than trusted:
- `external_services.md:616` (#11 "OPENAI_API_KEY deleted from the block") and `:618` (#12
  "AZURE_OPENAI_KEY_EU/_US omitted") → **MATCHED.** Same block, same predicate (*the repaired `.env` AI block
  omits a key the source requires*), one repair closes both.
- `external_services.md:634` (#11 "the bare-name list is 3 of 6") and `:634` (#12 "the bare names are framed as
  `platform/.env` config; they are stripped") → **NOT matched.** Identical anchor, different predicates, two
  different repairs.

## Matched by BOTH readings (12) — highest confidence

| # | anchor | claim | induced by iter-52? |
|---|---|---|---|
| U01 | `services/backend.md:253` | a top-level `app/migrations/` dir "holding atlas.sum" does not exist | no |
| U02 | `services/ai-readiness.md:316` | `SHOW_SECONDARY_TABS` anchored at `AIReadinessClient.tsx:69` (an import); actual `:78` | no |
| U03 | `services/ai-readiness.md:459-460` | M219 blockquote's `:137-138` and `:150-154` both name the wrong construct | no |
| U04 | `services/graphql-wundergraph.md:128` | self-anchor `:84` points at the package.json-stub blockquote, not the build-context text | no |
| U05 | `services/cms.md:110` | `python-docx` listed as a studio requirement; absent from `requirements.txt` | no |
| U06 | `architecture/alignment_testing.md:193` | `gate.sh:61` is a comment; the coverage call is at `:69` | no |
| U07 | `services/README.md:20` | "three of the four" names roadrunner, which the same blockquote calls the fifth | no |
| U08 | `services/hiring.md:33` | "as the twins already said" attributes to `service_taxonomy.md:52` + `dependency_map.md:78` a claim neither makes | **YES** |
| U09 | `services/ant-academy.md:196` | "the four bring-up patches" — `ant-academy.sh` applies five | no |
| U10 | `architecture/security_compliance.md:68` | the refuted "30 use `OrganizationMixin{}`" left standing **above its own retraction** | **YES** |
| U11 | `architecture/external_services.md:616-618` | the repaired `.env` AI block omits keys the source requires (`OPENAI_API_KEY`; `AZURE_OPENAI_KEY_EU/_US`) | **YES** |
| U12 | `services/hiring.md:252` | self-anchor `:157-159` does not carry the fact attributed to it (it is at `:162-166`) | **YES** |

## Found by reading #11 only (20)

| # | anchor | claim | induced? |
|---|---|---|---|
| U13 | `architecture/external_services.md:409` | pre-drop router `depends_on` given as 2 services; compose had 4 | no |
| U14 | `architecture/ai_architecture.md:225` | "both recordings stored in S3" contradicts `media-substrate-spec`'s "never in prod S3" | no |
| U15 | `services/graphql-wundergraph.md:81` | self-cite `:174-176` lands on a blank line + heading; the text is at `:178` | no |
| U16 | `services/ai-readiness.md:47` | self-cite `:458` is a bare blockquote separator; the intended text is at `:484` | no |
| U17 | `services/messenger.md:110` | `SKILLPATH_STREAM` in the messenger env table; it is a backend var, absent from the messenger repo | no |
| U18 | `architecture/alignment_testing.md:460` | snapshot-fidelity operators are six not five; content carries five not four | no |
| U19 | `services/cms.md:181` | internal anchor `:37` points at the GraphQL bullet, not the Studio bullet | no |
| U20 | `services/gotenberg.md:7` | the Gotenberg PDF is a text-extraction intermediate, not "display and storage" | no |
| U21 | `services/next-web-app.md:98` | `GRAPHQL_SCHEMA_FOR_GEN` is read by nothing; `codegen.ts` hardcodes the URL | no |
| U22 | `services/next-web-app.md:14` | "reaches backend only via GraphQL" is false — 30 direct REST call sites | no |
| U23 | `services/clerkenstein.md:18` | rext sections given as 6; actual 11; `CLAUDE.md` says 9 | no |
| U24 | `architecture/shared_libraries.md:77` | "12 Connect-RPC services" omits the live, served `StorageService` | no |
| U25 | `architecture/dependency_map.md:58` | `backend` (and `cms`, `:61`) stream-consumer lists omit Messenger | no |
| U26 | `architecture/service_taxonomy.md:290` | "the only service the compose gives `DIRECTUS_BASE_ADDR`" — backend requires it too | no |
| U27 | `architecture/dependency_map.md:50` | taxonomy/ai/authn consumer rows omit cms + jobsimulation direct requires | no |
| U28 | `architecture/shared_libraries.md:126` | the `internal/ai/ai.go` wrapper path does not exist in `app` | no |
| U29 | `services/jobsimulation.md:95` | the jobsim RPC consumer list names Backend and omits cms; contradicts `:34-35` | no |
| U30 | `architecture/external_services.md:634` | `gen.py:45-47` "separate bare names" lists 3 of 6; omits one **on the cited line** | **YES** |
| U31 | `services/studio-room.md:389` | the outbound-call set omits Mistral; contradicts `dependency_map.md:25` | **YES** |
| U32 | `ops/platform-alignment.md:698` | the rule-24 ledger convicts under rule 17 a claim its own repair proved true | **YES** |

## Found by reading #12 only (14)

| # | anchor | claim | induced? |
|---|---|---|---|
| U33 | `services/skiller.md:19` | prod `SKILLER_RPC_ADDR` given two incompatible hostnames (vs `backend.md:195`) | no |
| U34 | `architecture/security_compliance.md:22` | Cosmo Router placed in public subnets; terraform says private | no |
| U35 | `architecture/ai_architecture.md:240` | "three simulation model defaults" — the source has six default-applying getters | no |
| U36 | `services/cms.md:61` | Mistral named as the Studio-Room Python pipeline provider; zero Python references | no |
| U37 | `services/roadrunner.md:21` | "no other platform repo references roadrunner at all" — false, and self-contradicted | no |
| U38 | `services/roadrunner.md:33` | `architecture_overview.md:188` names the Skiller row, not Jobsimulation (`:189`) | no |
| U39 | `services/hiring.md:147` | `jobsimulation_id` is a dropped-mirror column; the cohort key is `sim_id` | no |
| U40 | `architecture/service_taxonomy.md:52` | says cms/jobsimulation/skillpath schemas exist locally; `hiring.md:311` says none do | no |
| U41 | `services/jobsimulation.md:126` | `cms/directus/collections` path missing `internal/`; 5 sibling sites use `app/internal/cms/` | **YES** |
| U42 | `services/ant-academy.md:137` | the `emptyCatalogView()` literal is wrong — omits `bundles: PUBLIC_BUNDLES` + `catalogVersion` | no |
| U43 | `services/ant-academy.md:324` | `code/.env` vs `code/.env.local` — the file contradicts itself and `CLAUDE.md:255` | no |
| U44 | `architecture/external_services.md:577` | `ANTHROPIC_API_KEY` "supplies Studio-Room the credential" — the env strip forbids it | no |
| U45 | `architecture/external_services.md:634` | the bare `gen.py` env names are framed as `platform/.env` config; they are stripped | no |
| U46 | `services/jobsimulation.md:129` | the new flag description contradicts untouched `ai_architecture.md:210` ("all it does") | **YES** |

## Induced classification — method, so it can be re-derived

A finding is **INDUCED** iff its anchor falls inside an added-line range of the repair diff
`1255998..0e35b1a`, **or** the defect is a contradiction/omission the repair itself created (a deletion
leaves no added line, and a repair that adds a retraction while leaving the refuted claim standing creates a
contradiction that did not exist before the commit). The added-line ranges were computed mechanically:

```
git diff -U0 1255998..0e35b1a -- corpus/ CLAUDE.md .claude/ | awk '/^\+\+\+ b\//{f=substr($2,3)} /^@@/{...}'
```

**Induced = 9 of 46** — U08, U10, U11, U12, U30, U31, U32, U41, U46.
Two are judgment calls recorded as such: **U10** and **U11** anchor on lines the diff did not add, but the
*defect* is the repair's (a retraction placed under a surviving refutation; keys deleted from a block).
**U31** (`studio-room.md:389`) and **U32** (`platform-alignment.md:698`) sit one line outside an added hunk
in the passage the repair rewrote.
