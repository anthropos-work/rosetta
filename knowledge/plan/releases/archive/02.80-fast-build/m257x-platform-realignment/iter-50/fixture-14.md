# Fixture — the answer key of reading #9 (iter-49's 14), captured BEFORE any repair

**Why this file exists.** §5 rule 21: *the fixture is perishable* — a corpus carrying a known, anchored,
independently re-verified answer key exists exactly once, and repairing it destroys the only thing that
can falsify an instrument built to protect it. §5 rule 22's corollary extends it: *capture the answer key
of a reading that contradicts an earlier reading, even when — especially when — nothing will be repaired
from it.*

**Tree state this key is valid for:** rosetta `57dfbfd` (corpus byte-identical to `47c9b7d`).
Ground truth: `app 5ba17044` · `app/studio aeec036a` · `platform 2adcf714` · `next-web-app bb3313bc` ·
`sentinel 88bc5592` · `storage 4ce8ece5` · `messenger fa47850d` · `cms ca50c817` ·
`graphql-wundergraph 60c229f3` · `roadrunner 87d8d443` · `jobsimulation 462343b0` ·
`studio-desk 14a5442a` · `ant-academy 9c3843cd`.

**Full adjudication:** [`../iter-49/blocker-ledger.md`](../iter-49/blocker-ledger.md). This table is the
matchable index used by iter-50's paired-overlap computation — one row per finding, keyed by the
**anchor** and the **claim**, so a reading #10 finding can be scored `same` / `different` without a
narrative comparison.

| # | seat | anchor | claim in one line | class |
|---|---|---|---|---|
| 1 | A | `external_services.md:668` | *only the US LiveKit agent is suffixed* — misses `anthropos-agent-chain` | THIS (overshoot) |
| 2 | B | `ai-readiness.md:406-419` | active-cycle filter named as `keepStartedMembers`; it is `keepInCycleStep1` | THIS (wrong mechanism) |
| 3 | C | `security_compliance.md:76`,`:120` + `architecture_overview.md:299-300` | **31** org-filtered schemas; it is **32** (`Organization`'s own `Policy()`) | PRE |
| 4 | C | `shared_libraries.md:145` | authn *"imported by … app"* only; the cms + jobsimulation husks import `colony/authn` directly | PRE |
| 5 | D | `roadrunner.md:87` | `SubmissionPackage` submits a *batch*; it is one multi-**file** submission, one token | PRE |
| 6 | D | `roadrunner.md:49` | Asynq pool for *batch* submissions; one poll task per single submission (contradicts `:97`) | PRE |
| 7 | E | `hiring.md:66-68` | `CreateOrganizationSimInvitationLink` is *the very call the seeder uses*; the seeder never calls it | PRE |
| 8 | E | `hiring.md:113-115` | with Clerk-only wiring the seeder *cannot write the 5 positions*; it still writes all 5 | PRE |
| 9 | E | `hiring.md:293-296` (recurs `:33-34`, `:157-158`) | *`jobsimulation.sessions` still exists* stated unqualified in a demo/dev section; no local stack creates that schema | PRE |
| 10 | G | `external_services.md:672` | LiveKit + OpenAI Realtime *powers new sessions* — the claim retracted 4 lines from an edited line | THIS (paraphrase leak) |
| 11 | G | `jobsimulation.md:123`,`:126` | same refuted `flag_use_realtime_openai` claim — repaired in 1 of 3 places | THIS (paraphrase leak) |
| 12 | G | `jobsimulation.md:102` | mirror migration *back-fills then DROPs*; `SET "score"` = 0 hits | THIS (paraphrase leak) |
| 13 | G | `hiring.md:210` | `token` is the *only* required-and-undefaulted column; **four** are | THIS (overshoot) |
| 14 | G | `external_services.md` + `architecture_overview.md` + `security_compliance.md` | a **5th** EU-egress path via Studio-Room `openai`; selected by no shipped config | THIS (overshoot) |

**Split:** THIS (induced by iter-49's repair) = 7 → #1, 2, 10, 11, 12, 13, 14.
PRE (present before it) = 7 → #3, 4, 5, 6, 7, 8, 9.

**Per-seat yield at #9:** A 1 · B 1 · C 2 · D 2 · E 3 · **F 0** · G 5.
