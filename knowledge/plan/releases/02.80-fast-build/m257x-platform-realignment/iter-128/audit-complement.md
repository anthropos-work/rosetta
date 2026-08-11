# iter-128 — the complement triage's own accuracy, hand-measured

iter-124 audited its rules over **C1**. This run triages the **complement** — the 820 the previous run
named *"untriaged, not implied clean"* — so the rules must be re-audited **on this population** rather
than importing C1's error rate. Importing it would be the same substitution of inference for measurement
the milestone exists to catch.

**Invocation** (seed committed, not chosen after the fact):

```
/usr/bin/python3 iter-128/triage-complement.py /tmp/tier2.json --audit 30 --seed 128 --r4-only
```

**Denominator: 30 of the 329 R4 members of the complement.** The sample is drawn **R4-only** and
deliberately so: iter-124 established that the error is not spread but **concentrated in R4** (R3 audited
21/21 = 100 %, R4 6/9 = 66.7 %). A proportional sample of the whole complement would have drawn ~18 R3
and ~12 R4 and spent most of its power re-confirming the rule that was already clean. **This is a
different sampling frame from iter-124's, and the two accuracies are therefore not directly comparable
as like-for-like — stated, not smoothed over.**

## The 30, hand-classified

| # | site | rule says | hand says | agree |
|---|---|---|---|---|
| 1 | `sentinel.md:16` | cite | cite — compose has no `profiles:` on sentinel | ✅ |
| 2 | `architecture_overview.md:67` | cite | cite — Directus self-hosted, in compose + `external_services.md` | ✅ |
| 3 | `external_services.md:543` | cite | **drop** — *"When backend services add new GraphQL types or operations:"* is a colon lead-in to a command block; it asserts nothing | ❌ |
| 4 | `alignment_testing.md:42` | cite | cite — the gene = (capability × variant) encoding is in the rext DNA format | ✅ |
| 5 | `chronos.md:214` | cite | **drop** — *"**Scheduling a simulation timeout** (from Jobsimulation service):"* is a heading-shaped lead-in to a snippet | ❌ |
| 6 | `external_services.md:287` | cite | cite — no local Directus in the default posture; compose | ✅ |
| 7 | `ai-readiness.md:239` | cite | cite — the no-chart-fallback path is in next-web | ✅ |
| 8 | `cms.md:227` | cite | cite — Redis cache + Watermill streams | ✅ |
| 9 | `ai_architecture.md:28` | cite | cite — a negative claim, but decidable in `app` | ✅ |
| 10 | `shared_libraries.md:326` | cite | cite — the taxonomy lib really is `NodeID` + validators | ✅ |
| 11 | `external_services.md:800` | cite | **drop** — *"Configure S3 for file storage"* is an imperative checklist step, not an assertion | ❌ |
| 12 | `sentinel.md:91` | cite | cite — decidable in `sentinel` | ✅ |
| 13 | `service_taxonomy.md:495` | cite | cite — the three-service floor is in compose | ✅ |
| 14 | `security_compliance.md:264` | cite | cite — PostHog/product-analytics; borderline table label but decidable | ✅ |
| 15 | `frontend_architecture.md:93` | cite | cite — codegen from subgraph schemas | ✅ |
| 16 | `studio-room.md:143` | cite | cite — names the `GenMode` enum (escaped `ARTIFACT` only because it is CamelCase) | ✅ |
| 17 | `ant-academy.md:119` | cite | cite — no schema, no Atlas migrations | ✅ |
| 18 | `chronos.md:42` | cite | **hedge** — Chronos was removed from orchestration (`045857c`) and **is in no clone set**; not decidable from here | ❌ |
| 19 | `ai-readiness.md:256` | cite | cite — the seeder owns `session_count` | ✅ |
| 20 | `studio-desk.md:415` | cite | cite — the two demopatches are in rext | ✅ |
| 21 | `alignment_testing.md:60` | cite | cite — one subtest per gene, in rext | ✅ |
| 22 | `chronos.md:78` | cite | **hedge** — same as 18: `ChronosService` is in no clone set | ❌ |
| 23 | `dependency_map.md:79` | cite | cite — LiveKit voice / Chime recording, in `app` | ✅ |
| 24 | `db-backup.md:55` | cite | **hedge** — **`db-backup` is in NO clone set.** iter-123 measured exactly this class: its citations resolve nowhere reachable | ❌ |
| 25 | `architecture_overview.md:36` | cite | cite — the two Studio services | ✅ |
| 26 | `external_services.md:563` | cite | cite — `flag_use_azure_us` is in `app` | ✅ |
| 27 | `alignment_testing.md:12` | cite | cite — the differential-test definition matches the framework | ✅ |
| 28 | `external_services.md:230` | cite | cite — `app/internal/cms/` adds logic + caching | ✅ |
| 29 | `ai_architecture.md:73` | cite | cite — corpus-internal, but a decidable statement about a surviving claim | ✅ |
| 30 | `external_services.md:350` | cite | **drop** — *"Invalidate CMS service cache when content updates"* is an imperative checklist step | ❌ |

## Result — the error is concentrated the same way, and it runs the same direction

**23 of 30 agree = 76.7 %.** R4's measured accuracy on the complement is **76.7 %**, against **66.7 %**
on C1 — but see the sampling-frame caveat above before reading that as an improvement.

**All 7 disagreements run one way**: the rule assigns `cite`, the hand assigns `drop` (4) or `hedge` (3).
**Not one runs the other way, and none touches `fix`.** That is the direction R4 was *declared generous
in before either sample was drawn*, so both audits confirm the declared bias rather than discovering it.

**The two `drop` sub-classes are worth naming, because they are mechanical and a future rule could catch
them:** colon-terminated lead-ins to a code block (#3, #5) and imperative checklist steps (#11, #30).
Neither is an assertion; `R1` catches template placeholders and doc-purpose sentences but not these.

**The `hedge` sub-class is one thing, not three:** #18, #22 and #24 are all *"the subject is a repo no
clone set contains"* — `chronos` and `db-backup`. `R2`'s `UNCHECKABLE` list is keyed on vendor-internal
phrases and never learned about un-cloned first-party repos.

## The corrected split, with the correction shown rather than folded in

Applying R4's measured 23.3 % mis-assignment to its **329** complement members, split in the sampled
4-drop : 3-hedge ratio:

| fate | as printed | corrected estimate | note |
|---|---|---|---|
| **cite** | 815 (99.4 %) | **≈ 738 (90.0 %)** | R3's 486 members are **un-corrected — and un-audited here** |
| **drop** | 5 (0.6 %) | **≈ 49 (6.0 %)** | lead-ins + checklist steps |
| **hedge** | 0 (0.0 %) | **≈ 33 (4.0 %)** | almost entirely `chronos` + `db-backup` |
| **fix** | 0 | **0 is NOT a measurement — it is a FLOOR of unknown height** | see below |

## Three limitations, stated rather than hidden

1. **`fix = 0` over the complement means "nothing was read", not "nothing is false."** The triage cannot
   decide falsity — `fix` is a hand-adjudicated input. This run's reading was aimed at **C1**, by the
   consequence-ordering rule. **The complement has not been read for falsity at all, and its false-claim
   count is UNMEASURED.** Quoting `fix 0.0 %` without that sentence would be precisely the error
   iter-124's own `fix`-is-a-floor paragraph warns against.
2. **R3 (486 members, 59 % of the complement) was not audited here.** iter-124 measured it at 100 % on
   C1 and I carried that forward without re-testing it on this population. That is an assumption, and it
   is the largest un-tested block in the corrected number above.
3. **The sampling frame differs from iter-124's** (R4-only vs whole-class), so 76.7 % and 66.7 % are not
   a before/after pair and must not be quoted as a trend.
