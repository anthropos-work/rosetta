# iter-50 — clause-5 reading #10: the blocker ledger

Seven auditors, the instrument frozen at iter-41's on every knob — same seat count, same briefing, same
6-way partition of the 40 files (**identical hand**, because the corpus did not change), all read
top-to-bottom under a per-file `wc -l` positive control, plus the diff seat. Ground-truth clones
**byte-identical to iter-49's**, all thirteen re-read at open.

**This reading repaired nothing and was taken on the tree reading #9 had just read** — rosetta `57dfbfd`,
whose 40 audited files are byte-identical to `47c9b7d`. It is the second half of the paired experiment
§5 rule 22 prescribes. The overlap analysis is in [`variance.md`](variance.md); this file adjudicates the
reading on its own terms.

**Blind.** No seat was told a prior reading existed, and every seat was barred from reading anything under
`knowledge/plan/**`. A seat that could see the answer key would measure agreement, not detection.

---

## The count

**7 raw → 7 unique → 7 held.** Per seat: **A 1 · B 0 · C 0 · D 1 · E 0 · F 0 · G 5.**

Reading #9 returned 14 on this same tree (`A 1 · B 1 · C 2 · D 2 · E 3 · F 0 · G 5`).

| pass | iter | auditors | blockers |
|---|---|---|---|
| 1–8 | 21/33 · 34 · 38 · 39 · 41 · 47 · 48 | grep-scoped → 7 | 25 · 13 · 11→17 · 37 · 18 · 7 · 12 |
| 9 | 49 | 7 — frozen | **14** |
| **10** | **50** | **7 — frozen, SAME TREE, no repair between** | **7** |

## The findings

**Class key:** **PRE** = present before iter-49's repair · **THIS** = manufactured by it.
**#9 column:** whether reading #9 named this same claim.

### Seat A — 1

| # | site | the false claim | what is true | class | #9 |
|---|---|---|---|---|---|
| 1 | `external_services.md:604-618` | the `platform/.env` AI block names `OPENAI_API_KEY`, `OPENAI_ORG_ID`, `AZURE_OPENAI_ENDPOINT`, `AZURE_OPENAI_DEPLOYMENT` | **Measured in `stack-demo/app`, `--include='*.go'`:** `OPENAI_KEY` **33** hits, `AZURE_OPENAI_ENDPOINT_URL` **24**, `OPENAI_API_KEY` **3** (all the studio-subprocess remap at `internal/cms/studio/studioManager.go:1067`), `AZURE_OPENAI_ENDPOINT` **0**, `AZURE_OPENAI_DEPLOYMENT` **0**, `OPENAI_ORG_ID` **0**. The one variable the app actually reads — `OPENAI_KEY` — is **absent from the block**. The corpus's own `CLAUDE.md` and `coursebuilder.md:98` already say `OPENAI_KEY` | **PRE** | **MISSED** |

A reader provisioning `.env` from this block gets a **keyless OpenAI client and an unread Azure endpoint**.
Independently re-measured at adjudication; the numbers above are this iteration's, not the seat's.

### Seat D — 1

| # | site | the false claim | what is true | class | #9 |
|---|---|---|---|---|---|
| 2 | `studio-room.md:388` | *"studio-room makes no GraphQL or Directus calls; **its only outbound API call is to the skills taxonomy service**"* | The pipeline's primary egress is to the AI providers — `app/studio/services/ai.py` instantiates OpenAI / AzureOpenAI / Anthropic clients (`:1-2, 383, 530, 664`). **Contradicted by this same file at `:36` and `:261-266`.** An egress allowlist built from this sentence blocks generation entirely | **PRE** | **MISSED** |

Held rather than downgraded: the sentence is unqualified, sits in the section a reader consults for exactly
this question, and the harm is concrete. The seat itself noted it reads as a scoping slip — that is an
argument about the *repair*, not about whether a reader would act on it.

### Seat G (diff) — 5 · the highest-yield seat, again

| # | site | the false claim | what is true | class | #9 |
|---|---|---|---|---|---|
| 3 | `external_services.md:672` + `jobsimulation.md:123`,`:126` | *"LiveKit + OpenAI Realtime powers new sessions (gated by `flag_use_realtime_openai`)"* | The claim iter-49 **retracted** at `ai_architecture.md:207`, still standing at three sites — one of them **four lines below a hunk the same commit edited, in the same file**. Refuted by the flag's single read, `calls/livekit.go:131-135` | **THIS** (paraphrase leak) | **#10 + #11** |
| 4 | `hiring.md:296` (+ `:31`, `:158`) | *"`jobsimulation.sessions` still exists, frozen and unwritten, until M710"*, unqualified in the **Local development** section | Measured live: `demo-1-postgresql-1` has **no `jobsimulation` schema**; `to_regclass('jobsimulation.sessions')` → NULL. `repos.yml:17-19` `migrations: false`; app's only `CREATE SCHEMA` is `auth`. The cited `askengine/registry.go:192` is an **LLM name-alias map entry**, not a physical schema | **PRE** | **#9** |
| 5 | `hiring.md:241-242` | *"**Neither table exists** and there is no second subgraph, so **both halves collapsed**"* | Verbatim the claim iter-49 retracted **83 lines above** in the same file (`:157-159`: *"`jobsimulation.sessions` itself was NOT dropped … survives frozen until M710"*). **2 of 3 sites repaired.** The two sites are a mutual contradiction: one says it exists, one says it does not | **THIS** (leak) | **MISSED** |
| 6 | `dependency_map.md:19` | Storage: *"**No Postgres, no Redis:** storage owns no database (no `DB_CONNECTION`/`REDIS_ADDR` in compose…)"*, with `-` in the **Depends On** column. Marked *"Corrected M257x iter-49"* | `docker-compose.yml:213-217` declares `storage: depends_on: redis {service_healthy}`, `postgresql {service_healthy}` — verified at adjudication. The table's own header sources that column from those declarations. Contradicts the **unedited twin** `service_taxonomy.md:418`. The *runtime* half of the correction is right; the *dependency* half is an overshoot | **THIS** (overshoot) | **MISSED — in a hunk seat G reviewed and passed at #9** |
| 7 | `hiring.md:209-210` | `token` is *"the **only** required-and-undefaulted column in the table"* | **Four** are: `owner_id` `:6`, `sim_id` `:7`, `sim_type` `:10`, `token` `:13`, in the very DDL the sentence cites | **THIS** (overshoot) | **#13** |

---

## The split

| class | n | findings |
|---|---|---|
| **THIS** — manufactured by iter-49's repair | **4** | 3, 5, 6, 7 |
| **PRE** — present before it | **3** | 1, 2, 4 |

Reading #9, on the identical tree, split the same population **7 / 7**.

## What this reading positively CLEARED that reading #9 booked as a blocker

This is the half a normal ledger never records, and here it is the point.

| #9 finding | what reading #10 did |
|---|---|
| **#3** — *31 vs 32 org-filtered Ent schemas* (`security_compliance.md:76,:120` + `architecture_overview.md:299`) | **Three seats independently cleared it as an audited zero.** B re-ran the doc's own `comm`/`xargs` derivation; C re-derived `31 of 135` and `16+7=23`; G "re-derived independently — every figure matches" |
| **#2** — the active-cycle predicate is `keepInCycleStep1`, not `keepStartedMembers` | Seat B recorded the passage as an audited zero, having verified that the quoted SQL at `steps.go:915-938` is **verbatim** |
| **#5, #6** — two `roadrunner.md` batch-vs-single claims | Seat D read the file and recorded *"roadrunner.md's heavily-repaired banner survives intact"* |
| **#1** — the third LiveKit agent name | Seat A read `external_services.md` in full and did not raise it |
| **#14** — the 5th EU-egress path via Studio-Room `openai` | Seat A booked it a **MINOR**; seat G booked it **unverified** — *"latent, and the corpus says 'can'"*. A grading disagreement, not a detection failure |
| **#4, #7, #8, #12** | not raised |

**#9's #3 is CORRECT, and this reading's three clearances of it are wrong.** Measured at adjudication:
`schema/organization.go:56` declares `func (Organization) Policy()`, whose query rule is
`rule.FilterSameOrganizations()` (`:96`), and it uses neither mixin. The mixin count is **30** and exactly
**4** schemas declare their own `Policy()`. So the org-filtering set is **32**, not 31.

> **All three seats reproduced the doc's own arithmetic instead of testing the doc's own predicate.** The
> derivation shown in the text (*"the 30 mixin users plus `Membership`"*) is internally consistent and
> re-computes perfectly; it is the **set** that is incomplete. This is §5 rule 17 — *a count can be exactly
> right while the claim it supports is false; verify the PREDICATE, not just the count* — written by this
> milestone, and violated by three independent auditors in one pass, each of whom recorded the result as a
> **positively audited zero**. An audited zero that is wrong is worse than a silence: it is evidence
> pointing the wrong way.

## Clause 5

This reading did not return zero. Clause 5 is **NOT MET**.
