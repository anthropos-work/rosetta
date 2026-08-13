# iter-38 — the clause-5 fourth pass: 17 blockers, enumerated

Six auditors, 40 files, every in-scope file read top-to-bottom. Re-partitioned per §5 rule 18(b) so no
auditor inherited iters 33/34's boundaries. Combined verification volume reported by the auditors:
**~510 exact citations checked against source**, the live `demo-1` Postgres, or `docker-compose.yml`/
`repos.yml` at platform origin `2adcf71`.

| # | file:line | the false claim | ground truth |
|---|---|---|---|
| 1 | `security_compliance.md:7` + `:183` | *"scoring is deterministic (rubric-based), NOT AI-scored"* — the stated basis of the **EU AI Act Limited-Risk** classification | The only registered check engine is the model: `check.EngineLlm: NewLLMBulkChecker` (`v3/validator/validator.go:60-61`); `calculateSkillScore` (`skills.go:53-64`) counts LLM booleans. **A conjunction whose BOTH halves fail** |
| 2 | `security_compliance.md:119` | *"Organization switching requires re-authentication"* | Client-side `clerk.setActive({organization})` — a token re-mint. 8 call sites, no `signOut` |
| 3 | `architecture_overview.md:261` | prod flow *"Cosmo Router (port 5050)"* | `5050` was the deleted LOCAL compose host mapping; prod is `8080` (`graphql-wundergraph/terraform/locals.tf:8`) |
| 4 | `ai_architecture.md:7` + `:151` | *"simulation scoring is NOT done by AI"* | Same as #1 — **found independently by a second auditor in a different file** |
| 5 | `ai_architecture.md:159` | *"Thresholds: Level 1 ≥ 60, Level 2 ≥ 65, Level 3 ≥ 75, Level 4 ≥ 85, Level 5 ≥ 95"* | **No such ladder exists** in any repo. Real: `calculateCompetencyLevelScore` (`skills.go:40-51`) = `max(0, score*2-100)`, with a `// TODO fix this formula` |
| 6 | `service_taxonomy.md:207` | *"PWA via Serwist 9 (offline chapters)"* | Removed at ant-academy v0.5 M1; **the repo regression-fences the removal** (`next-scaffold.test.js:106,111`) |
| 7 | `backend.md:49` + `:136` | `internal/copilot` listed as a live package | Deleted at app `889ae776` (2026-05-19), 490 lines |
| 8 | `backend.md:50` | *"The labs-api client is currently wired as nil"* | Wired whenever `LABS_API_URL` is set (`main.go:735-738`). The same doc says so correctly 130 lines later |
| 9 | `backend.md:175` | *"9 `ai_readiness_*` ent tables"* | **13** (live DB + 13 schema files) |
| 10 | `skillpath.md:30-32` | *"the `SkillPathSessionService` surface … is served by `app`"* | **0** occurrences in Go source; the platform's own M506 record says *"No in-app RPC handler is served"* |
| 11 | `skillpath.md:62-64` | *"calls jobsimulation over Connect-RPC (`GetSessions`)"* | In-process `jobsimReader` returning `*ent.JobSimulationSession`; the RPC client was deleted at jobsim-in-app M709 |
| 12 | `coursebuilder.md:48` + `:87-89` | *"LLM usage — AWS Bedrock (eu-west-1)"*, *"routes stay unmounted if Bedrock creds are absent"* | `ANTHROPIC_API_KEY` ⇒ first-party API (`coursebuilder/bedrock.go:105-114`), and terraform makes it **required in production** |
| 13 | `coursebuilder.md:93` | env var `COURSEBUILDER_OPENAI_IMAGE_KEY` | Dead — deleted at app `68c24512`; the generator reads `OPENAI_KEY` (`main.go:816-819`) |
| 14 | `academy-backend.md:57-64` (+ `ant-academy.md:63,66`) | seven singular academy table names | Ent pluralizes: `academy_chapter_progresses`, `academy_certificates`, … `SELECT … FROM academy_certificate` errors |
| 15 | `hiring.md:52-53` | *"insights path … requires `is_hiring` true"* | `InsightsJobSimulationByMemberships` (`:1034-1080`) never reads it — gates on `OrgFeatureInsights` + membership status. **The doc contradicted itself at `:197-198`** |
| 16 | `hiring.md:78` | *"Clerk-only → the insights read-path won't treat the cohort as hiring"* | Same mechanism as #15 |
| 17 | `hiring.md:158` | `completion_status` values include `SIMULATION…` | Closed 5-value enum `pending/passed/failed/discarded/timedout`. The column has **no CHECK**, so a bad value INSERTs cleanly and the row vanishes at Ent scan |

## Where they were — the prediction, and its refutation

|  | files | blockers | per file |
|---|---|---|---|
| repaired by iter-34 (in clause-5 scope) | 8 | **11** | 1.375 |
| never opened by iter-34 | 32 | **6** | 0.19 |

The ~7.3× density ratio **reproduces** §5 rule 18's ~9×. But the pre-registered prediction — *"2-5 blockers
total; 0-1 across the 32 untouched"* — is **REFUTED on both count (17) and location (6)**, and the refutation
is the finding: rule 18 licenses **weighting**, not **narrowing**. A pass scoped to the 9 changed files, as
routed, would have found 11 of 17 and reported a clean sweep of everything else.

Two of the six untouched-file blockers (#4, #5) sit in `ai_architecture.md`, which two prior full-read
passes had already read and passed. **#4 is the same false claim as #1, in a different file, found
independently by a different auditor** — which is what a re-partition buys: correlated blind spots are a
property of how the corpus was divided, not only of who read it.

---

# The adversarial pass over THIS sweep — 6 self-inflicted + 7 collateral

§5 rule 18(a) is not ceremony: this is the third consecutive pass whose own repair introduced defects.

| # | what the sweep broke | class |
|---|---|---|
| S1 | `coursebuilder.md:66` still said the group unmounts "when the Bedrock-backed `Service` is nil (missing AWS creds)" — **18 lines below its own retraction** | the retracted claim restated in different words nearby |
| S2 | `hiring.md:57` over-corrected to *"`is_hiring` drives the client-side re-skin … not the read path"* — **both conjuncts false**: the re-skin is Clerk-derived (the very line cited as support says so), and the CONTENT-LIBRARY read path *does* branch on it (`resolver_cms_queries.go:95,210,258,295`) | over-correction past the truth |
| S3 | `hiring.md:50-55, 82` — a **half-applied edit**: a doubled *"The The"* and an orphaned predicate (*"…What actually gates that scoreboard / won't treat the cohort as hiring"*) that re-asserted the withdrawn claim and promised an answer never given | mechanical damage |
| S4 | `service_taxonomy.md:207` — *"no `public/manifest.json`"*: the manifest exists as `public/academy-manifest.json`, wired at `layout.jsx:132`; the app is still installable, just online-only | negative evidence tested the wrong filename |
| S5 | `security_compliance.md` + `ai_architecture.md` — cited `checkerEngines` as the mechanism. **It is stored and never read**; dispatch is the hardcoded switch at `criterion.go:127 → :428`. And `EngineTextDiff` **does** run (`:168`, `:450-475`, pure string compare), so *"the verdicts come from an LLM"* is false as a universal | a load-bearing citation pointing at a dead field — **on a compliance page** |
| S6 | `backend.md:173` — "the four a '9' omits" enumerated **five** tables | count/gloss mismatch |

Collateral found in files the sweep did not touch (same claims, elsewhere): `ant-academy.md:31,34,237`
(still advertises the Serwist PWA), `ai-readiness.md:253-254` (singular table names 90 lines below the rows
the sweep pluralized, **in a file the sweep had just edited**), `frontend_architecture.md:39,44` +
`external_services.md:372` (`:5050` as the PROD router port), and four `corpus/ops/**` docs
(out of clause-5 scope — routed, see the iter close).

**All six self-inflicted and all in-scope collateral were fixed (15 further edits, 10 files).** The
auditor verified ~50 of the sweep's new citations against source: **every `file:line` anchor resolved
correctly**; all six defects were in surrounding prose, over-correction, or a half-applied edit — the
distribution §5 rule 18 predicts, now observed a third time.
